package token_pools

import (
	"bytes"
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	token_deployments "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

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
	Program      solana.PublicKey
	NewOwner     solana.PublicKey
	TokenMint    solana.PublicKey
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
	_ = chain.GetAccountDataBorshInto(context.Background(), configPDA, &chainConfig)
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