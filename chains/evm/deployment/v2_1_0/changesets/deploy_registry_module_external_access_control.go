package changesets

import (
	"fmt"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// DeployRegistryModuleExternalAccessControlCfg configures the DeployRegistryModuleExternalAccessControl changeset.
type DeployRegistryModuleExternalAccessControlCfg struct {
	// ChainSel is the selector of the EVM chain to deploy the module on.
	ChainSel uint64
}

// ChainSelector implements the sequences.WithChainSelector interface used to resolve the EVM chain dependency.
func (c DeployRegistryModuleExternalAccessControlCfg) ChainSelector() uint64 {
	return c.ChainSel
}

// DeployRegistryModuleExternalAccessControl deploys the RegistryModuleExternalAccessControl contract and registers it as a
// registry module on the chain's TokenAdminRegistry. The TokenAdminRegistry address is resolved from the datastore.
//
// Deployment is idempotent: an existing module recorded in the datastore is reused. Registration is idempotent too: the
// module is only added to the TokenAdminRegistry if it is not already registered. Because the TokenAdminRegistry is
// owned by the timelock on live chains, the registration is prepared as an MCMS transaction when the deployer key is not
// the owner.
var DeployRegistryModuleExternalAccessControl = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.DeployRegistryModuleExternalAccessControlInput,
	evm.Chain,
	DeployRegistryModuleExternalAccessControlCfg,
]{
	Sequence: sequences.DeployRegistryModuleExternalAccessControl,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployRegistryModuleExternalAccessControlCfg) (sequences.DeployRegistryModuleExternalAccessControlInput, error) {
		tokenAdminRegistryAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.DeployRegistryModuleExternalAccessControlInput{}, fmt.Errorf("failed to find token admin registry address: %w", err)
		}

		return sequences.DeployRegistryModuleExternalAccessControlInput{
			ChainSelector:      cfg.ChainSel,
			TokenAdminRegistry: tokenAdminRegistryAddress,
			ExistingAddresses:  e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel)),
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployRegistryModuleExternalAccessControlCfg],
})
