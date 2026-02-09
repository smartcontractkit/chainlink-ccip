package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	adapters1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	tokenadminops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type ConfigImportAdapter struct {
	FeeQuoter     common.Address
	OnRamp        common.Address
	OffRamp       common.Address
	Router        common.Address
	TokenAdminReg common.Address
}

func (ci *ConfigImportAdapter) InitializeAdapter(e cldf.Environment, chainSelector uint64) error {
	fqRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:          datastore.ContractType(fqops.ContractType),
		Version:       fqops.Version,
		ChainSelector: chainSelector,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find fee quoter contract ref for chain %d: %w", chainSelector, err)
	}
	ci.FeeQuoter = fqRef
	routerRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:          datastore.ContractType("Router"),
		Version:       semver.MustParse("1.2.0"),
		ChainSelector: chainSelector,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find router contract ref for chain %d: %w", chainSelector, err)
	}
	ci.Router = routerRef
	tokenAdminRegRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:          datastore.ContractType(tokenadminops.ContractType),
		Version:       tokenadminops.Version,
		ChainSelector: chainSelector,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find token admin registry contract ref for chain %d: %w", chainSelector, err)
	}
	ci.TokenAdminReg = tokenAdminRegRef
	onRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(onrampops.ContractType),
		Version: semver.MustParse("1.6.0"),
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find onramp contract ref for chain %d: %w", chainSelector, err)
	}
	ci.OnRamp = onRampRef
	offRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(offrampops.ContractType),
		Version: semver.MustParse("1.6.0"),
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find offramp contract ref for chain %d: %w", chainSelector, err)
	}
	ci.OffRamp = offRampRef
	return nil
}

func (ci *ConfigImportAdapter) SupportedTokensPerRemoteChain(e cldf.Environment, chainsel uint64) (map[uint64][]common.Address, error) {
	chain, ok := e.BlockChains.EVMChains()[chainsel]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found in environment", chainsel)
	}
	// get all supported tokens from token admin registry
	return adapters1_5.GetSupportedTokensPerRemoteChain(ci.TokenAdminReg, chain)
}

func (ci *ConfigImportAdapter) ConnectedChains(e cldf.Environment, chainsel uint64) ([]uint64, error) {
	chain, ok := e.BlockChains.EVMChains()[chainsel]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found in environment", chainsel)
	}
	routerAddr := ci.Router
	if routerAddr == (common.Address{}) {
		return nil, fmt.Errorf("router address not initialized for chain %d", chainsel)
	}
	// get all offRamps from router to find connected chains
	routerC, err := router.NewRouter(ci.Router, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate router contract at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	offRamps, err := routerC.GetOffRamps(&bind.CallOpts{
		Context: e.GetContext(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramps from router at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	connectedChains := make([]uint64, 0)
	for _, offRamp := range offRamps {
		// if the offramp's address matches our offramp, then we are connected to the source chain via 1.6
		if offRamp.OffRamp == ci.OffRamp {
			// get the onRamp on router for the source chain and check if it matches our onRamp, if it does then we are connected to that chain
			// lanes are always bi-directional so source and destination chain selectors are interchangeable for the purpose of finding connected chains
			onRamp, err := routerC.GetOnRamp(&bind.CallOpts{
				Context: e.GetContext(),
			}, offRamp.SourceChainSelector)
			if err != nil {
				return nil, fmt.Errorf("failed to get on ramp for source chain selector %d from router at %s on chain %d: %w", offRamp.SourceChainSelector, routerAddr.String(), chain.Selector, err)
			}
			if onRamp != ci.OnRamp {
				continue
			}
			connectedChains = append(connectedChains, offRamp.SourceChainSelector)
		}
	}
	return connectedChains, nil
}

func (ci *ConfigImportAdapter) SequenceImportConfig() *cldf_ops.Sequence[api.ImportConfigPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"seq-config-import",
		semver.MustParse("1.0.0"),
		"Imports configuration for specified chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.ImportConfigPerChainInput) (output sequences.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			chainSelector := in.ChainSelector
			b.Logger.Infof("Importing configuration for chain %d (%s)", chainSelector, evmChain.Name())
			// read FQ config from onchain
			fqAddress := ci.FeeQuoter
			if fqAddress == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("fee quoter address not initialized for chain %d", chainSelector)
			}
			var result sequences.OnChainOutput
			// fetch fee quoter config
			result, err = sequences.RunAndMergeSequence(b, chains,
				seq1_6.FeeQuoterImportConfigSequence,
				seq1_6.FeeQuoterImportConfigSequenceInput{
					Address:              fqAddress,
					ChainSelector:        chainSelector,
					RemoteChains:         in.RemoteChains,
					TokensPerRemoteChain: in.TokensPerRemoteChain,
				}, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to import fee quoter config on chain %d: %w", chainSelector, err)
			}
			// fetch onramp config
			onRampAddress := ci.OnRamp
			if onRampAddress == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("onramp address not initialized for chain %d", chainSelector)
			}
			result, err = sequences.RunAndMergeSequence(b, chains,
				seq1_6.OnRampImportConfigSequence,
				seq1_6.OnRampImportConfigSequenceInput{
					Address:       onRampAddress,
					ChainSelector: chainSelector,
					RemoteChains:  in.RemoteChains,
				}, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to import onramp config on chain %d: %w", chainSelector, err)
			}
			// fetch offramp config
			offRampAddress := ci.OffRamp
			if offRampAddress == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("offramp address not initialized for chain %d", chainSelector)
			}
			result, err = sequences.RunAndMergeSequence(b, chains,
				seq1_6.OffRampImportConfigSequence,
				seq1_6.OffRampImportConfigSequenceInput{
					Address:       offRampAddress,
					ChainSelector: chainSelector,
					RemoteChains:  in.RemoteChains,
				}, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to import offramp config on chain %d: %w", chainSelector, err)
			}
			return result, nil
		},
	)
}
