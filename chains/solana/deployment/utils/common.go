package utils

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	"github.com/smartcontractkit/mcms/types"
)

const (
	TimelockProgramType cldf_deployment.ContractType = "RBACTimelockProgram"
	McmProgramType      cldf_deployment.ContractType = "ManyChainMultiSigProgram"
	// special type for Solana that encodes PDA seed usage
	ProposerSeed     cldf_deployment.ContractType = "ProposerSeed"
	CancellerSeed    cldf_deployment.ContractType = "CancellerSeed"
	BypasserSeed     cldf_deployment.ContractType = "BypasserSeed"
	RBACTimelockSeed cldf_deployment.ContractType = "RBACTimelockSeed"
)

// Common parameters for transferring ownership of a program
type TransferOwnershipParams struct {
	Program      solana.PublicKey
	CurrentOwner solana.PublicKey
	NewOwner     solana.PublicKey
}

func BuildMCMSBatchOperation(
	chainSelector uint64,
	ixns []solana.Instruction,
	programID string,
	contractType string) (types.BatchOperation, error) {
	txns := make([]types.Transaction, 0, len(ixns))
	for _, ixn := range ixns {
		data, err := ixn.Data()
		if err != nil {
			return types.BatchOperation{}, fmt.Errorf("failed to extract data: %w", err)
		}
		for _, account := range ixn.Accounts() {
			if account.IsSigner {
				account.IsSigner = false
			}
		}
		tx, err := mcms_solana.NewTransaction(
			programID,
			data,
			big.NewInt(0),  // e.g. value
			ixn.Accounts(), // pass along needed accounts
			contractType,   // some string identifying the target
			[]string{},     // any relevant metadata
		)
		if err != nil {
			return types.BatchOperation{}, fmt.Errorf("failed to create transaction: %w", err)
		}
		txns = append(txns, tx)
	}
	return types.BatchOperation{
		ChainSelector: types.ChainSelector(chainSelector),
		Transactions:  txns,
	}, nil
}

func GetTimelockSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	qualifier string) solana.PublicKey {
	// timelock seeds stored as a separate program type
	// qualifier will identify the correct timelock instance
	timelock := datastore.GetAddressRef(
		existingAddresses,
		chainSelector,
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0,
		qualifier,
	)
	id, seed, _ := mcms_solana.ParseContractAddress(timelock.Address)
	return state.GetTimelockSignerPDA(
		id,
		state.PDASeed([]byte(seed[:])),
	)
}

func GetMCMSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	signerType cldf_deployment.ContractType,
	qualifier string) solana.PublicKey {
	// mcm seeds stored as a separate program type
	// qualifier will identify the correct mcm instance
	mcm := datastore.GetAddressRef(
		existingAddresses,
		chainSelector,
		signerType,
		common_utils.Version_1_6_0,
		qualifier,
	)
	id, seed, _ := mcms_solana.ParseContractAddress(mcm.Address)
	return state.GetMCMSignerPDA(
		id,
		state.PDASeed([]byte(seed[:])),
	)
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
