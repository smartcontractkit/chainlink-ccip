package changesets

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcmssdksol "github.com/smartcontractkit/mcms/sdk/solana"
	"github.com/smartcontractkit/mcms/types"
)

type TimelockConfig struct {
	MinDelay                  time.Duration        `json:"minDelay"` // delay for timelock worker to execute the transfers.
	MCMSAction                types.TimelockAction `json:"mcmsAction"`
	OverrideRoot              bool                 `json:"overrideRoot"`                        // if true, override the previous root with the new one.
	TimelockQualifierPerChain map[uint64]string    `json:"timelockQualifierPerChain,omitempty"` // optional qualifier to fetch timelock address from datastore
	ValidDuration             time.Duration        `json:"validDuration" yaml:"validDuration"`
}

type RMNRemoteSetEventAuthoritiesChangesetInput struct {
	ChainSelector uint64 `json:"chainSelector"`
	UpgradeConfig struct {
		MCMS TimelockConfig `json:"mcms"`
	} `json:"upgradeConfig"`
}

type MCMSRole string

const (
	BypasserManyChainMultisig  cldf.ContractType = "BypasserManyChainMultiSig"
	CancellerManyChainMultisig cldf.ContractType = "CancellerManyChainMultiSig"
	ProposerManyChainMultisig  cldf.ContractType = "ProposerManyChainMultiSig"
	ManyChainMultisig          cldf.ContractType = "ManyChainMultiSig"
	RBACTimelock               cldf.ContractType = "RBACTimelock"
	CallProxy                  cldf.ContractType = "CallProxy"

	// roles
	ProposerRole  MCMSRole = "PROPOSER"
	BypasserRole  MCMSRole = "BYPASSER"
	CancellerRole MCMSRole = "CANCELLER"

	// LinkToken is the burn/mint link token. It should be used everywhere for
	// new deployments. Corresponds to
	// https://github.com/smartcontractkit/chainlink/blob/develop/core/gethwrappers/shared/generated/link_token/link_token.go#L34
	LinkToken cldf.ContractType = "LinkToken"
	// StaticLinkToken represents the (very old) non-burn/mint link token.
	// It is not used in new deployments, but still exists on some chains
	// and has a distinct ABI from the new LinkToken.
	// Corresponds to the ABI
	// https://github.com/smartcontractkit/chainlink/blob/develop/core/gethwrappers/generated/link_token_interface/link_token_interface.go#L34
	StaticLinkToken cldf.ContractType = "StaticLinkToken"
	// mcms Solana specific
	ManyChainMultisigProgram         cldf.ContractType = "ManyChainMultiSigProgram"
	RBACTimelockProgram              cldf.ContractType = "RBACTimelockProgram"
	AccessControllerProgram          cldf.ContractType = "AccessControllerProgram"
	ProposerAccessControllerAccount  cldf.ContractType = "ProposerAccessControllerAccount"
	ExecutorAccessControllerAccount  cldf.ContractType = "ExecutorAccessControllerAccount"
	CancellerAccessControllerAccount cldf.ContractType = "CancellerAccessControllerAccount"
	BypasserAccessControllerAccount  cldf.ContractType = "BypasserAccessControllerAccount"
)

func (role MCMSRole) String() string {
	return string(role)
}

func (tc *TimelockConfig) validateCommon() error {
	// if MCMSAction is not set, default to timelock.Schedule
	if tc.MCMSAction == "" {
		tc.MCMSAction = types.TimelockActionSchedule
	}
	if tc.MCMSAction != types.TimelockActionSchedule &&
		tc.MCMSAction != types.TimelockActionCancel &&
		tc.MCMSAction != types.TimelockActionBypass {
		return fmt.Errorf("invalid MCMS type %s", tc.MCMSAction)
	}
	return nil
}

func (tc *TimelockConfig) ValidateSolana(e cldf.Environment, chainSelector uint64) error {
	err := tc.validateCommon()
	if err != nil {
		return err
	}

	validateContract := func(contractType cldf.ContractType) error {
		timelockID, err := cldf.SearchAddressBook(e.ExistingAddresses, chainSelector, contractType)
		if err != nil {
			return fmt.Errorf("%s not present on the chain %w", contractType, err)
		}
		// Make sure addresses are correctly parsed. Format is: "programID.PDASeed"
		_, _, err = mcmssdksol.ParseContractAddress(timelockID)
		if err != nil {
			return fmt.Errorf("failed to parse timelock address: %w", err)
		}
		return nil
	}

	err = validateContract(RBACTimelock)
	if err != nil {
		return err
	}

	switch tc.MCMSAction {
	case types.TimelockActionSchedule:
		err = validateContract(ProposerManyChainMultisig)
		if err != nil {
			return err
		}
	case types.TimelockActionCancel:
		err = validateContract(CancellerManyChainMultisig)
		if err != nil {
			return err
		}
	case types.TimelockActionBypass:
		err = validateContract(BypasserManyChainMultisig)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid MCMS action %s", tc.MCMSAction)
	}

	return nil
}

// DeployTokenGovernor returns a changeset that deploys a TokenGovernor contract
func RMNRemoteSetEventAuthoritiesSequenceChangeset() cldf.ChangeSetV2[RMNRemoteSetEventAuthoritiesChangesetInput] {
	return cldf.CreateChangeSet(rmnSetEventAuthoritiesApply, rmnSetEventAuthoritiesVerify)
}

func rmnSetEventAuthoritiesVerify(env cldf.Environment, input RMNRemoteSetEventAuthoritiesChangesetInput) error {
	if err := deployment.IsValidChainSelector(input.ChainSelector); err != nil {
		return fmt.Errorf("invalid chain selector: %d - %w", input.ChainSelector, err)
	}
	if !env.BlockChains.Exists(input.ChainSelector) {
		return fmt.Errorf("chain with selector %d does not exist", input.ChainSelector)
	}

	// Validate UpgradeConfig
	if err := input.UpgradeConfig.MCMS.ValidateSolana(env, input.ChainSelector); err != nil {
		return fmt.Errorf("invalid MCMS configuration: %w", err)
	}

	return nil
}

func rmnSetEventAuthoritiesApply(e cldf.Environment, input RMNRemoteSetEventAuthoritiesChangesetInput) (cldf.ChangesetOutput, error) {
	reports := make([]cldf_ops.Report[any, any], 0)

	// Prepare sequence input
	seqInput := sequences.RMNRemoteSetEventAuthoritiesSequenceInput{
		DataStore: e.DataStore,
		Selector:  input.ChainSelector,
	}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.SetRMNRemoteEventAuthorities, e.BlockChains, seqInput)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to set rmn event authorities: %w", err)
	}

	reports = append(reports, report.ExecutionReports...)

	// Create the datastore with the addresses from the report
	ds := datastore.NewMemoryDataStore()
	for _, addr := range report.Output.Addresses {
		if err := ds.Addresses().Add(addr); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to add address to datastore: %w", err)
		}
	}

	return cldf.ChangesetOutput{
		DataStore: ds,
		Reports:   reports,
	}, nil
}
