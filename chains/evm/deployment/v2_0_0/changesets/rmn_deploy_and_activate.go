package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// ActivateRMNCfg configures the ActivateRMN changeset.
type ActivateRMNCfg struct {
	ChainSels []uint64
	// CurseAdmins are optional additional authorized callers (cursers) added at RMN deploy
	// time, keyed by chain selector. The Ultra Fast Curse RBACTimelock is always included.
	CurseAdmins map[uint64][]common.Address
}

var ActivateRMN = func(mcmsRegistry *changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[ActivateRMNCfg]] {
	return cldf_deployment.CreateChangeSet(
		func(e cldf_deployment.Environment, input changesets.WithMCMS[ActivateRMNCfg]) (cldf_deployment.ChangesetOutput, error) {
			return applyDeployAndActivateRMN(e, mcmsRegistry, input)
		},
		validateActivateRMN,
	)
}

func validateActivateRMN(e cldf_deployment.Environment, input changesets.WithMCMS[ActivateRMNCfg]) error {
	if len(input.Cfg.ChainSels) == 0 {
		return fmt.Errorf("at least one chain selector is required")
	}
	if err := validateActivateRMNCurseAdmins(input.Cfg.CurseAdmins, input.Cfg.ChainSels); err != nil {
		return err
	}
	evmChains := e.BlockChains.EVMChains()
	for _, sel := range input.Cfg.ChainSels {
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))
		if err := validateActivateRMNAddresses(addresses, sel); err != nil {
			return err
		}
	}
	return nil
}

func validateActivateRMNCurseAdmins(curseAdmins map[uint64][]common.Address, chainSels []uint64) error {
	if len(curseAdmins) == 0 {
		return nil
	}
	chainSet := make(map[uint64]struct{}, len(chainSels))
	for _, sel := range chainSels {
		chainSet[sel] = struct{}{}
	}
	for sel, addrs := range curseAdmins {
		if _, ok := chainSet[sel]; !ok {
			return fmt.Errorf("curse admins configured for chain %d which is not in ChainSels", sel)
		}
		seen := make(map[common.Address]struct{}, len(addrs))
		for _, addr := range addrs {
			if addr == (common.Address{}) {
				return fmt.Errorf("curse admin address cannot be zero for chain %d", sel)
			}
			if _, dup := seen[addr]; dup {
				return fmt.Errorf("duplicate curse admin %s for chain %d", addr.Hex(), sel)
			}
			seen[addr] = struct{}{}
		}
	}
	return nil
}

func applyDeployAndActivateRMN(
	e cldf_deployment.Environment,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	input changesets.WithMCMS[ActivateRMNCfg],
) (cldf_deployment.ChangesetOutput, error) {
	evmChains := e.BlockChains.EVMChains()
	outputDS := datastore.NewMemoryDataStore()
	allBatchOps := make([]mcms_types.BatchOperation, 0, len(input.Cfg.ChainSels))

	for _, sel := range input.Cfg.ChainSels {
		chain := evmChains[sel]
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployAndActivateRMN, chain, sequences.ActivateRMNInput{
			ChainSelector:     sel,
			ExistingAddresses: addresses,
			CurseAdmins:       input.Cfg.CurseAdmins[sel],
		})
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to activate RMN on chain %d: %w", sel, err)
		}

		for _, ref := range report.Output.Addresses {
			if addErr := outputDS.Addresses().Add(ref); addErr != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add address to datastore: %w", addErr)
			}
		}
		allBatchOps = append(allBatchOps, report.Output.BatchOps...)
	}

	return changesets.NewOutputBuilder(e, mcmsRegistry).
		WithDataStore(outputDS).
		WithBatchOps(allBatchOps).
		Build(input.MCMS)
}

func validateActivateRMNAddresses(addresses []datastore.AddressRef, chainSelector uint64) error {
	if ref := datastore_utils.GetAddressRef(
		addresses,
		chainSelector,
		common_utils.RBACTimelock,
		mcms_ops.MCMSVersion,
		common_utils.RMNTimelockQualifier,
	); ref.Address == "" {
		return fmt.Errorf(
			"ownership transfer requires RMNMCMS RBACTimelock (qualifier %q) in datastore for chain %d",
			common_utils.RMNTimelockQualifier, chainSelector,
		)
	}
	if ref := datastore_utils.GetAddressRef(
		addresses,
		chainSelector,
		common_utils.RBACTimelock,
		mcms_ops.MCMSVersion,
		common_utils.UltraFastCurseMCMSQualifier,
	); ref.Address == "" {
		return fmt.Errorf(
			"Ultra Fast Curse RBACTimelock (qualifier %q) not found in datastore for chain %d",
			common_utils.UltraFastCurseMCMSQualifier, chainSelector,
		)
	}
	return nil
}
