package token_pools

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	token_deployments "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

// ensureSetRouterPersists checks, before anything is sent or proposed, that the token pool
// program at programID actually persists set_router. Programs before solana-v1.6.2 declared
// AdminUpdateTokenPool.state without mut, so they accept set_router and silently discard the
// write. readOnlyStateProbe must be a set_router instruction whose state account meta is
// read-only (what the v0.1.1 bindings generate): simulating it against a program >= 1.6.2 fails
// with Anchor's ConstraintMut (error 2000), while a legacy program executes it successfully.
// Signature verification is disabled for the simulation so the same probe works when the
// authority is an MCMS signer rather than the deployer key.
func ensureSetRouterPersists(ctx context.Context, chain cldf_solana.Chain, programID solana.PublicKey, readOnlyStateProbe solana.Instruction, poolTypeName string) error {
	tx, err := solana.NewTransaction([]solana.Instruction{readOnlyStateProbe}, solana.Hash{}, solana.TransactionPayer(chain.DeployerKey.PublicKey()))
	if err != nil {
		return fmt.Errorf("failed to build set router probe transaction for %s token pool: %w", poolTypeName, err)
	}
	// The validator sanitizes the wire format before simulating, so the transaction must carry
	// one signature slot per required signer even though none of them is verified.
	tx.Signatures = make([]solana.Signature, tx.Message.Header.NumRequiredSignatures)
	res, err := chain.Client.SimulateTransactionWithOpts(ctx, tx, &rpc.SimulateTransactionOpts{
		SigVerify:              false,
		ReplaceRecentBlockhash: true,
		Commitment:             rpc.CommitmentConfirmed,
	})
	if err != nil {
		return fmt.Errorf("failed to simulate set router probe for %s token pool program %s: %w", poolTypeName, programID, err)
	}
	if res == nil || res.Value == nil {
		return fmt.Errorf("set router probe for %s token pool program %s returned no simulation result", poolTypeName, programID)
	}
	if res.Value.Err == nil {
		return fmt.Errorf(
			"%s token pool program %s does not persist set_router: token pool programs before solana-v1.6.2 accept the instruction without writing the state; upgrade the program",
			poolTypeName, programID,
		)
	}
	for _, log := range res.Value.Logs {
		if strings.Contains(log, "ConstraintMut") {
			return nil
		}
	}
	return fmt.Errorf(
		"set router probe for %s token pool program %s failed for an unexpected reason: %v (logs: %s)",
		poolTypeName, programID, res.Value.Err, strings.Join(res.Value.Logs, " | "),
	)
}

// get diff of pool addresses
func poolDiff(existingPoolAddresses []base_token_pool.RemoteAddress, newPoolAddresses []base_token_pool.RemoteAddress) []base_token_pool.RemoteAddress {
	var result []base_token_pool.RemoteAddress
	// for every new address, check if it exists in the existing pool addresses
	for _, newAddr := range newPoolAddresses {
		exists := false
		for _, existingAddr := range existingPoolAddresses {
			if bytes.Equal(existingAddr.Address, newAddr.Address) {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, newAddr)
		}
	}
	return result
}

type TokenPoolTransferOwnershipInput struct {
	Program   solana.PublicKey
	NewOwner  solana.PublicKey
	TokenMint solana.PublicKey
}

type SetPoolRouterInput struct {
	Program   solana.PublicKey
	TokenMint solana.PublicKey
	NewRouter solana.PublicKey
}

type Params struct {
	TokenPool solana.PublicKey
	TokenMint solana.PublicKey
	// SPLToken or SPLToken2022
	TokenProgramID solana.PublicKey
	// Only used for certain ops
	RMNRemote        solana.PublicKey
	Router           solana.PublicKey
	NewMintAuthority solana.PublicKey
	OldMintAuthority solana.PublicKey
}

type RemoteChainConfig struct {
	TokenPool solana.PublicKey
	TokenMint solana.PublicKey
	// SPLToken or SPLToken2022
	TokenProgramID            solana.PublicKey
	RemoteSelector            uint64
	RemoteTokenAddress        []byte
	RemotePoolAddress         []byte
	RemoteDecimals            uint8
	ForceOverrideRemoteConfig bool
	InboundRateLimiterConfig  token_deployments.RateLimiterConfig
	OutboundRateLimiterConfig token_deployments.RateLimiterConfig
}

type initGlobalCfgParams struct {
	PoolTypeLabel string // e.g. common_utils.BurnMintTokenPool.String()
	LogName       string // e.g. "BurnMintTokenPool"
	SetProgramID  func(solana.PublicKey)
	BuildInitIx   func(configPDA solana.PublicKey, upgradeAuthority solana.PublicKey, programData solana.PublicKey) (solana.Instruction, error)
}

// initGlobalConfigTokenPool initializes the token pool global config if not initialized.
// If upgradeAuthority != deployer => produces MCMS batch op, otherwise sends tx directly.
func initGlobalConfigTokenPool(
	b operations.Bundle,
	chain cldf_solana.Chain,
	input Params,
	p initGlobalCfgParams,
) (sequences.OnChainOutput, error) {
	p.SetProgramID(input.TokenPool)

	programData, err := utils.GetSolProgramData(chain.Client, input.TokenPool)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	upgradeAuthority, err := utils.GetUpgradeAuthority(chain.Client, input.TokenPool)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	configPDA, _, _ := state.FindConfigPDA(input.TokenPool)

	// Check if already initialized.
	var chainConfig base_token_pool.BaseConfig
	_ = chain.GetAccountDataBorshInto(b.GetContext(), configPDA, &chainConfig)
	if !chainConfig.TokenProgram.IsZero() {
		b.Logger.Info(p.LogName+" global config already initialized for token pool:", input.TokenPool.String())
		return sequences.OnChainOutput{}, nil
	}

	ixn, err := p.BuildInitIx(configPDA, upgradeAuthority, programData.Address)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	// If deployer isn't upgrade authority, create MCMS batch op.
	if upgradeAuthority != chain.DeployerKey.PublicKey() {
		batch, err := utils.BuildMCMSBatchOperation(
			chain.Selector,
			[]solana.Instruction{ixn},
			input.TokenPool.String(),
			p.PoolTypeLabel,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batch}}, nil
	}

	// Otherwise execute directly.
	if err := chain.Confirm([]solana.Instruction{ixn}); err != nil {
		return sequences.OnChainOutput{}, err
	}

	return sequences.OnChainOutput{}, nil
}

type PoolInitializeOut struct {
	sequences.OnChainOutput
	Initializer solana.PublicKey
}
