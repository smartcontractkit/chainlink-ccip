package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcmsevmsdk "github.com/smartcontractkit/mcms/sdk/evm"
	mcms_types "github.com/smartcontractkit/mcms/types"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

type MCMSBuilder interface {
	Input() *common_utils.MCMSInput
	DeriveTimelocks(e cldf_deployment.Environment) (map[mcms_types.ChainSelector]string, error)
	DeriveChainMetaData(e cldf_deployment.Environment) (map[mcms_types.ChainSelector]mcms_types.ChainMetadata, error)
}

func ResolveMCMS(e cldf_deployment.Environment, args MCMSBuilder) (MCMSBuildParams, error) {
	in := args.Input()
	if in == nil {
		return MCMSBuildParams{}, nil
	}
	if err := in.Validate(); err != nil {
		return MCMSBuildParams{}, fmt.Errorf("invalid MCMS input: %w", err)
	}

	tl, err := args.DeriveTimelocks(e)
	if err != nil {
		return MCMSBuildParams{}, fmt.Errorf("derive timelocks: %w", err)
	}
	meta, err := args.DeriveChainMetaData(e)
	if err != nil {
		return MCMSBuildParams{}, fmt.Errorf("derive chain metadata: %w", err)
	}

	return MCMSBuildParams{
		MCMSInput:         *in, // value-copy to discourage later mutation
		TimelockAddresses: tl,
		ChainMetadata:     meta,
	}, nil
}

type EVMMCMBuilder struct {
	in *common_utils.MCMSInput
}

func (b EVMMCMBuilder) Input() *common_utils.MCMSInput {
	return b.in
}

func (b EVMMCMBuilder) DeriveTimelocks(e cldf_deployment.Environment) (map[mcms_types.ChainSelector]string, error) {
	in := b.Input()
	if in == nil {
		return nil, nil
	}
	evm := e.BlockChains.EVMChains()
	if len(evm) == 0 {
		return nil, fmt.Errorf("no EVM chains found in environment")
	}
	out := make(map[mcms_types.ChainSelector]string, len(evm))
	for sel := range evm {
		addrs, err := datastore_utils.FindAndFormatEachRef(
			e.DataStore,
			[]datastore.AddressRef{{
				ChainSelector: sel,
				Type:          in.TimelockAddressRef.Type,
				Version:       in.TimelockAddressRef.Version,
				Qualifier:     in.TimelockAddressRef.Qualifier,
				Labels:        in.TimelockAddressRef.Labels,
			}},
			datastore_utils.ToEVMAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("chain %d: %w", sel, err)
		}
		if len(addrs) != 1 {
			return nil, fmt.Errorf("chain %d: expected 1 address, got %d", sel, len(addrs))
		}
		out[mcms_types.ChainSelector(sel)] = addrs[0].String()
	}
	return out, nil
}

func (b EVMMCMBuilder) DeriveChainMetaData(e cldf_deployment.Environment) (map[mcms_types.ChainSelector]mcms_types.ChainMetadata, error) {
	in := b.Input()
	if in == nil {
		return nil, nil
	}
	evm := e.BlockChains.EVMChains()
	if len(evm) == 0 {
		return nil, fmt.Errorf("no EVM chains found in environment")
	}
	out := make(map[mcms_types.ChainSelector]mcms_types.ChainMetadata, len(evm))
	for sel, chain := range evm {
		inspector := mcmsevmsdk.NewInspector(chain.Client)
		addrs, err := datastore_utils.FindAndFormatEachRef(
			e.DataStore,
			[]datastore.AddressRef{{
				ChainSelector: sel,
				Type:          in.MCMSAddressRef.Type,
				Version:       in.MCMSAddressRef.Version,
				Qualifier:     in.MCMSAddressRef.Qualifier,
				Labels:        in.MCMSAddressRef.Labels,
			}},
			datastore_utils.ToEVMAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("chain %d: %w", sel, err)
		}
		if len(addrs) != 1 {
			return nil, fmt.Errorf("chain %d: expected 1 MCMS address, got %d", sel, len(addrs))
		}
		addr := addrs[0].String()
		opCount, err := inspector.GetOpCount(e.GetContext(), addr)
		if err != nil {
			return nil, fmt.Errorf("chain %d: get op count for %s: %w", sel, addr, err)
		}
		out[mcms_types.ChainSelector(sel)] = mcms_types.ChainMetadata{
			StartingOpCount: opCount,
			MCMAddress:      addr,
		}
	}
	return out, nil
}

func NewEVMMCMBuilder(in *common_utils.MCMSInput) EVMMCMBuilder {
	return EVMMCMBuilder{in: in}
}
