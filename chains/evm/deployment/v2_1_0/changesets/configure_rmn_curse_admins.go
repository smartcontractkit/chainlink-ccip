package changesets

import (
	"fmt"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// ConfigureRMNCurseAdminsCfg is the configuration for the ConfigureRMNCurseAdmins changeset.
// It manages the authorized callers (curse admins) on an already-deployed RMN contract.
type ConfigureRMNCurseAdminsCfg struct {
	ChainSel uint64
	// RMN is an address ref used to look up the deployed RMN contract in the datastore.
	RMN  datastore.AddressRef
	Args rmnops.AuthorizedCallerArgs
}

func (c ConfigureRMNCurseAdminsCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var ConfigureRMNCurseAdmins = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.ConfigureRMNCurseAdminsInput,
	evm.Chain,
	ConfigureRMNCurseAdminsCfg,
]{
	Sequence: sequences.ConfigureRMNCurseAdmins,
	ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureRMNCurseAdminsCfg) (sequences.ConfigureRMNCurseAdminsInput, error) {
		rmnAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.RMN, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureRMNCurseAdminsInput{}, fmt.Errorf("failed to resolve RMN ref on chain with selector %d: %w", cfg.ChainSel, err)
		}
		return sequences.ConfigureRMNCurseAdminsInput{
			ChainSelector: cfg.ChainSel,
			RMNAddress:    rmnAddr,
			Args:          cfg.Args,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[ConfigureRMNCurseAdminsCfg],
})
