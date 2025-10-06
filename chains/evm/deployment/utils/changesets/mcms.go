package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcmsevmsdk "github.com/smartcontractkit/mcms/sdk/evm"
	mcms_types "github.com/smartcontractkit/mcms/types"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
)

// TODO : move this to a common chain agnostic package
type MCMSInput struct {
	// OverridePreviousRoot indicates whether to override the root of the MCMS contract.
	OverridePreviousRoot bool
	// ValidUntil is a unix timestamp indicating when the proposal expires.
	// Root can't be set or executed after this time.
	ValidUntil uint32
	// TimelockDelay is the amount of time each operation in the proposal must wait before it can be executed.
	TimelockDelay mcms_types.Duration
	// TimelockAction is the action to perform on the timelock contract (schedule, bypass, or cancel).
	TimelockAction mcms_types.TimelockAction
	// MCMSAddressRef is a reference to the MCMS contract address in the datastore.
	MCMSAddressRef *datastore.AddressRef
	// TimelockAddressRef is a reference to the timelock contract address in the datastore.
	TimelockAddressRef *datastore.AddressRef
}

// TODO : need to put more validation here
func (c *MCMSInput) Validate() error {
	if c.TimelockAction != mcms_types.TimelockActionSchedule &&
		c.TimelockAction != mcms_types.TimelockActionBypass &&
		c.TimelockAction != mcms_types.TimelockActionCancel {
		return fmt.Errorf("invalid timelock action: %s", c.TimelockAction)
	}
	return nil
}

type MCMSBuilder interface {
	Input() *MCMSInput
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
	in *MCMSInput
}

func (b EVMMCMBuilder) Input() *MCMSInput {
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

func NewEVMMCMBuilder(in *MCMSInput) EVMMCMBuilder {
	return EVMMCMBuilder{in: in}
}
