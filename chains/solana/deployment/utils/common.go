package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	"github.com/smartcontractkit/mcms/types"
)

const (
	TimelockProgramType         cldf_deployment.ContractType = "RBACTimelockProgram"
	McmProgramType              cldf_deployment.ContractType = "ManyChainMultiSigProgram"
	AccessControllerProgramType cldf_deployment.ContractType = "AccessControllerProgram"
	// special type for Solana that encodes PDA seed usage
	ProposerSeed     cldf_deployment.ContractType = "ProposerSeed"
	CancellerSeed    cldf_deployment.ContractType = "CancellerSeed"
	BypasserSeed     cldf_deployment.ContractType = "BypasserSeed"
	RBACTimelockSeed cldf_deployment.ContractType = "RBACTimelockSeed"
	// access controller accounts
	ProposerAccessControllerAccount  cldf_deployment.ContractType = "ProposerAccessControllerAccount"
	ExecutorAccessControllerAccount  cldf_deployment.ContractType = "ExecutorAccessControllerAccount"
	CancellerAccessControllerAccount cldf_deployment.ContractType = "CancellerAccessControllerAccount"
	BypasserAccessControllerAccount  cldf_deployment.ContractType = "BypasserAccessControllerAccount"
	// tokens
	SPLTokens     cldf_deployment.ContractType = "SPLTokens"
	SPL2022Tokens cldf_deployment.ContractType = "SPL2022Tokens"
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

// ErrMCMSInstanceNotFound is returned when the datastore has no ref for a Solana timelock or MCM instance.
var ErrMCMSInstanceNotFound = errors.New("solana mcms instance not found in datastore")

// mcmsInstanceRefTypes maps the "<program>.<seed>" ref type of a Solana MCMS instance to its
// seed-only ref type and its program ref type. Deploys write both instance refs, but some
// datastores only have the "<program>.<seed>" one.
var mcmsInstanceRefTypes = map[cldf_deployment.ContractType]struct {
	seed, program cldf_deployment.ContractType
}{
	common_utils.RBACTimelock:               {RBACTimelockSeed, TimelockProgramType},
	common_utils.ProposerManyChainMultisig:  {ProposerSeed, McmProgramType},
	common_utils.CancellerManyChainMultisig: {CancellerSeed, McmProgramType},
	common_utils.BypasserManyChainMultisig:  {BypasserSeed, McmProgramType},
}

// MCMSInstance is a Solana timelock or MCM instance resolved from the datastore.
type MCMSInstance struct {
	ProgramID solana.PublicKey
	Seed      state.PDASeed
	// MissingRefs are the refs for this instance the datastore lacks, derived from the ref it has.
	MissingRefs []cldf_datastore.AddressRef
}

// ResolveMCMSInstance resolves a Solana timelock or MCM instance. instanceType is the
// "<program>.<seed>" ref type (RBACTimelock or one of the *ManyChainMultiSig types). When that ref
// is missing, the instance is rebuilt from the program ref and the seed ref. It returns
// ErrMCMSInstanceNotFound when neither the instance ref nor the seed ref exists.
func ResolveMCMSInstance(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	instanceType cldf_deployment.ContractType,
	qualifier string) (MCMSInstance, error) {
	refTypes, ok := mcmsInstanceRefTypes[instanceType]
	if !ok {
		return MCMSInstance{}, fmt.Errorf("unsupported solana mcms instance type %q", instanceType)
	}
	instanceRef := datastore.GetAddressRef(existingAddresses, chainSelector, instanceType, common_utils.Version_1_6_0, qualifier)
	seedRef := datastore.GetAddressRef(existingAddresses, chainSelector, refTypes.seed, common_utils.Version_1_6_0, qualifier)

	if instanceRef.Address != "" {
		id, seed, err := mcms_solana.ParseContractAddress(instanceRef.Address)
		if err != nil {
			return MCMSInstance{}, fmt.Errorf("invalid %s ref %q on chain %d: %w", instanceType, instanceRef.Address, chainSelector, err)
		}
		out := MCMSInstance{ProgramID: id, Seed: state.PDASeed(seed)}
		seedStr := string(bytes.TrimRight(seed[:], "\x00"))
		switch seedRef.Address {
		case "":
			out.MissingRefs = append(out.MissingRefs, cldf_datastore.AddressRef{
				Address:       seedStr,
				ChainSelector: chainSelector,
				Type:          cldf_datastore.ContractType(refTypes.seed),
				Version:       common_utils.Version_1_6_0,
				Qualifier:     instanceRef.Qualifier,
			})
		case seedStr:
		default:
			return MCMSInstance{}, fmt.Errorf("%s ref %q and %s ref %q disagree on chain %d", instanceType, instanceRef.Address, refTypes.seed, seedRef.Address, chainSelector)
		}
		return out, nil
	}

	if seedRef.Address == "" {
		return MCMSInstance{}, fmt.Errorf("%w: no %s or %s ref on chain %d with qualifier %q", ErrMCMSInstanceNotFound, instanceType, refTypes.seed, chainSelector, qualifier)
	}
	programRef := datastore.GetAddressRef(existingAddresses, chainSelector, refTypes.program, common_utils.Version_1_6_0, "")
	if programRef.Address == "" {
		return MCMSInstance{}, fmt.Errorf("found %s ref but no %s ref on chain %d", refTypes.seed, refTypes.program, chainSelector)
	}
	id, err := solana.PublicKeyFromBase58(programRef.Address)
	if err != nil {
		return MCMSInstance{}, fmt.Errorf("invalid %s ref %q on chain %d: %w", refTypes.program, programRef.Address, chainSelector, err)
	}
	var seed state.PDASeed
	if len(seedRef.Address) > len(seed) {
		return MCMSInstance{}, fmt.Errorf("%s ref %q on chain %d is longer than %d bytes", refTypes.seed, seedRef.Address, chainSelector, len(seed))
	}
	copy(seed[:], seedRef.Address)
	return MCMSInstance{
		ProgramID: id,
		Seed:      seed,
		MissingRefs: []cldf_datastore.AddressRef{{
			Address:       mcms_solana.ContractAddress(id, mcms_solana.PDASeed(seed)),
			ChainSelector: chainSelector,
			Type:          cldf_datastore.ContractType(instanceType),
			Version:       common_utils.Version_1_6_0,
			Qualifier:     seedRef.Qualifier,
		}},
	}, nil
}

func GetTimelockSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	qualifier string) (solana.PublicKey, error) {
	timelock, err := ResolveMCMSInstance(existingAddresses, chainSelector, common_utils.RBACTimelock, qualifier)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return state.GetTimelockSignerPDA(timelock.ProgramID, timelock.Seed), nil
}

func GetMCMSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	signerType cldf_deployment.ContractType,
	qualifier string) (solana.PublicKey, error) {
	mcm, err := ResolveMCMSInstance(existingAddresses, chainSelector, signerType, qualifier)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return state.GetMCMSignerPDA(mcm.ProgramID, mcm.Seed), nil
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
