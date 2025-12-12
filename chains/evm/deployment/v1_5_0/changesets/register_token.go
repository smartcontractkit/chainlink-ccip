package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type RegisterTokenCfg struct {
	ChainSel      uint64
	TokenPool     datastore.AddressRef
	Token         datastore.AddressRef
	ExternalAdmin common.Address
}

func (c RegisterTokenCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var RegisterToken = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.RegisterTokenInput,
	evm.Chain,
	RegisterTokenCfg,
]{
	Sequence: sequences.RegisterToken,
	ResolveInput: func(e cldf_deployment.Environment, cfg RegisterTokenCfg) (sequences.RegisterTokenInput, error) {
		tokenAdminRegistryRef := datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}
		tokenAdminRegistryAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, tokenAdminRegistryRef, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.RegisterTokenInput{}, fmt.Errorf("failed to find token admin registry address: %w", err)
		}
		tokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.Token, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.RegisterTokenInput{}, fmt.Errorf("failed to find token address: %w", err)
		}
		tokenPoolAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.TokenPool, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.RegisterTokenInput{}, fmt.Errorf("failed to find token pool address: %w", err)
		}
		return sequences.RegisterTokenInput{
			ChainSelector:             cfg.ChainSel,
			TokenPoolAddress:          tokenPoolAddress,
			TokenAdminRegistryAddress: tokenAdminRegistryAddress,
			TokenAddress:              tokenAddress,
			ExternalAdmin:             cfg.ExternalAdmin,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[RegisterTokenCfg],
})
