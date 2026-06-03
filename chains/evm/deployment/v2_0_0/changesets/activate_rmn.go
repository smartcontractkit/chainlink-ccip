package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// ActivateRMNCfg configures the ActivateRMN changeset for one EVM chain.
type ActivateRMNCfg struct {
	ChainSel uint64
}

func (c ActivateRMNCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var ActivateRMN = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.ActivateRMNInput,
	evm.Chain,
	ActivateRMNCfg,
]{
	Sequence: sequences.ActivateRMN,
	ResolveInput: func(e cldf_deployment.Environment, cfg ActivateRMNCfg) (sequences.ActivateRMNInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		if err := validateActivateRMNAddresses(addresses, cfg.ChainSel); err != nil {
			return sequences.ActivateRMNInput{}, err
		}
		return sequences.ActivateRMNInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[ActivateRMNCfg],
})

// ValidateActivateRMNCfg checks durable-pipeline configuration for ActivateRMN.
func ValidateActivateRMNCfg(cfg ActivateRMNCfg) error {
	if cfg.ChainSel == 0 {
		return fmt.Errorf("chain selector is required")
	}
	return nil
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
			"ultra Fast Curse RBACTimelock (qualifier %q) not found in datastore for chain %d",
			common_utils.UltraFastCurseMCMSQualifier, chainSelector,
		)
	}
	return nil
}
