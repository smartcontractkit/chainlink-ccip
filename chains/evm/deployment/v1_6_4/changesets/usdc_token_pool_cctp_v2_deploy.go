package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type USDCTokenPoolCCTPV2DeployInputPerChain struct {
	ChainSelector  uint64
	TokenMessenger common.Address
	Token          common.Address
	Allowlist      []common.Address
	RMNProxy       common.Address
	Router         common.Address
}

type USDCTokenPoolCCTPV2DeployInput struct {
	ChainInputs []USDCTokenPoolCCTPV2DeployInputPerChain
	MCMS        mcms.Input
}

func USDCTokenPoolCCTPV2DeployChangeset() deployment.ChangeSetV2[USDCTokenPoolCCTPV2DeployInput] {
	return cldf.CreateChangeSet(usdcTokenPoolCCTPV2DeployApply(), usdcTokenPoolCCTPV2DeployVerify())
}

func usdcTokenPoolCCTPV2DeployApply() func(cldf.Environment, USDCTokenPoolCCTPV2DeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input USDCTokenPoolCCTPV2DeployInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		for _, perChainInput := range input.ChainInputs {
			sequenceInput := sequences.USDCTokenPoolCCTPV2DeploySequenceInput{
				ChainSelector:  perChainInput.ChainSelector,
				TokenMessenger: perChainInput.TokenMessenger,
				Token:          perChainInput.Token,
				Allowlist:      perChainInput.Allowlist,
				RMNProxy:       perChainInput.RMNProxy,
				Router:         perChainInput.Router,
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolCCTPV2DeploySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolCCTPV2 on chain %d: %w", perChainInput.ChainSelector, err)
			}
			for _, r := range report.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
			reports = append(reports, report.ExecutionReports...)
		}
		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(input.MCMS)
	}
}

func usdcTokenPoolCCTPV2DeployVerify() func(cldf.Environment, USDCTokenPoolCCTPV2DeployInput) error {
	return func(e cldf.Environment, input USDCTokenPoolCCTPV2DeployInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
