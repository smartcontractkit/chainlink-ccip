package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/governance"
)

const (
	TimelockProgramType         cldf_deployment.ContractType = cldf_deployment.ContractType(governance.TimelockProgramType)
	McmProgramType              cldf_deployment.ContractType = cldf_deployment.ContractType(governance.McmProgramType)
	AccessControllerProgramType cldf_deployment.ContractType = cldf_deployment.ContractType(governance.AccessControllerProgramType)
	// special type for Solana that encodes PDA seed usage
	ProposerSeed     cldf_deployment.ContractType = cldf_deployment.ContractType(governance.ProposerSeed)
	CancellerSeed    cldf_deployment.ContractType = cldf_deployment.ContractType(governance.CancellerSeed)
	BypasserSeed     cldf_deployment.ContractType = cldf_deployment.ContractType(governance.BypasserSeed)
	RBACTimelockSeed cldf_deployment.ContractType = cldf_deployment.ContractType(governance.RBACTimelockSeed)
	// access controller accounts
	ProposerAccessControllerAccount  cldf_deployment.ContractType = cldf_deployment.ContractType(governance.ProposerAccessControllerAccount)
	ExecutorAccessControllerAccount  cldf_deployment.ContractType = cldf_deployment.ContractType(governance.ExecutorAccessControllerAccount)
	CancellerAccessControllerAccount cldf_deployment.ContractType = cldf_deployment.ContractType(governance.CancellerAccessControllerAccount)
	BypasserAccessControllerAccount  cldf_deployment.ContractType = cldf_deployment.ContractType(governance.BypasserAccessControllerAccount)
	// tokens
	SPLTokens     cldf_deployment.ContractType = "SPLTokens"
	SPL2022Tokens cldf_deployment.ContractType = "SPL2022Tokens"
)

// Common parameters for transferring ownership of a program
type TransferOwnershipParams = governance.TransferOwnershipParams

func BuildMCMSBatchOperation(
	chainSelector uint64,
	ixns []solana.Instruction,
	programID string,
	contractType string) (types.BatchOperation, error) {
	return governance.BuildMCMSBatchOperation(chainSelector, ixns, programID, contractType)
}

// ErrMCMSInstanceNotFound is returned when the datastore has no ref for a Solana timelock or MCM instance.
var ErrMCMSInstanceNotFound = governance.ErrMCMSInstanceNotFound

// MCMSInstance is a Solana timelock or MCM instance resolved from the datastore.
type MCMSInstance = governance.MCMSInstance

// ResolveMCMSInstance resolves a Solana timelock or MCM instance at 1.6.0. instanceType is the
// "<program>.<seed>" ref type (RBACTimelock or one of the *ManyChainMultiSig types). When that ref
// is missing, the instance is rebuilt from the program ref and the seed ref. It returns
// ErrMCMSInstanceNotFound when neither the instance ref nor the seed ref exists.
func ResolveMCMSInstance(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	instanceType cldf_deployment.ContractType,
	qualifier string) (MCMSInstance, error) {
	return governance.ResolveMCMSInstance(
		existingAddresses, chainSelector, cldf_datastore.ContractType(instanceType), common_utils.Version_1_6_0, qualifier)
}

func GetTimelockSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	qualifier string) (solana.PublicKey, error) {
	return governance.GetTimelockSignerPDA(existingAddresses, chainSelector, common_utils.Version_1_6_0, qualifier)
}

func GetMCMSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	signerType cldf_deployment.ContractType,
	qualifier string) (solana.PublicKey, error) {
	return governance.GetMCMSignerPDA(
		existingAddresses, chainSelector, cldf_datastore.ContractType(signerType), common_utils.Version_1_6_0, qualifier)
}

func FundSolanaAccounts(
	ctx context.Context,
	accounts []solana.PublicKey,
	solAmount uint64,
	solanaGoClient *rpc.Client,
) error {
	var sigs = make([]solana.Signature, 0, len(accounts))
	for _, account := range accounts {
		sig, err := solanaGoClient.RequestAirdrop(
			ctx,
			account,
			solAmount*solana.LAMPORTS_PER_SOL,
			rpc.CommitmentFinalized)
		if err != nil {
			return err
		}
		sigs = append(sigs, sig)
	}

	const timeout = 100 * time.Second
	const pollInterval = 500 * time.Millisecond

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	remaining := len(sigs)
	for remaining > 0 {
		select {
		case <-timeoutCtx.Done():
			return errors.New("unable to find transaction within timeout")
		case <-ticker.C:
			statusRes, sigErr := solanaGoClient.GetSignatureStatuses(ctx, true, sigs...)
			if sigErr != nil {
				return sigErr
			}
			if statusRes == nil {
				return errors.New("Status response is nil")
			}
			if statusRes.Value == nil {
				return errors.New("Status response value is nil")
			}

			unfinalizedCount := 0
			for _, res := range statusRes.Value {
				if res == nil || res.ConfirmationStatus == rpc.ConfirmationStatusFinalized {
					unfinalizedCount++
				}
			}
			remaining = unfinalizedCount
		}
	}
	return nil
}

// FundFromAddressIxs transfers SOL from the given address to each provided account and waits for confirmations.
func FundFromAddressIxs(solChain cldf_solana.Chain, from solana.PublicKey, accounts []solana.PublicKey, amount uint64) ([]solana.Instruction, error) {
	var ixs []solana.Instruction
	for _, account := range accounts {
		// Create a transfer instruction using the provided builder.
		ix, err := system.NewTransferInstruction(
			amount,
			from,    // funding account (sender)
			account, // recipient account
		).ValidateAndBuild()
		if err != nil {
			return nil, fmt.Errorf("failed to create transfer instruction: %w", err)
		}
		ixs = append(ixs, ix)
	}

	return ixs, nil
}

// FundFromDeployerKey transfers SOL from the deployer to each provided account and waits for confirmations.
func FundFromDeployerKey(solChain cldf_solana.Chain, accounts []solana.PublicKey, amount uint64) error {
	ixs, err := FundFromAddressIxs(solChain, solChain.DeployerKey.PublicKey(), accounts, amount*solana.LAMPORTS_PER_SOL)
	if err != nil {
		return fmt.Errorf("failed to create transfer instructions: %w", err)
	}
	err = solChain.Confirm(ixs)
	if err != nil {
		return fmt.Errorf("failed to confirm transaction: %w", err)
	}
	return nil
}
