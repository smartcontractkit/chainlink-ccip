package adapters

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
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

	// connectedChainsCache memoizes the result of ConnectedChains per chain selector
	// to avoid duplicate (potentially expensive) RPC work when the method is called
	// multiple times for the same chain within the same adapter instance.
	connectedChainsCache map[uint64][]uint64
	connectedChainsMu    sync.Mutex
}

func (ci *ConfigImportAdapter) InitializeAdapter(e cldf.Environment, chainSelector uint64) error {
	fqRef, err := seq1_6.GetFeeQuoterAddress(
		e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSelector),
			datastore.AddressRefByType(datastore.ContractType(fqops.ContractType)),
		),
		chainSelector, semver.MustParse("2.0.0"))
	if err != nil {
		return fmt.Errorf("failed to find fee quoter contract ref for chain %d: %w", chainSelector, err)
	}
	ci.FeeQuoter, err = evm_datastore_utils.ToEVMAddress(fqRef)
	if err != nil {
		return fmt.Errorf("failed to convert fee quoter address to EVM address for chain %d: %w", chainSelector, err)
	}
	routerRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:          datastore.ContractType(routerops.ContractType),
		Version:       routerops.Version,
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
		Version: onrampops.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find onramp contract ref for chain %d: %w", chainSelector, err)
	}
	ci.OnRamp = onRampRef
	offRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(offrampops.ContractType),
		Version: offrampops.Version,
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
	remoteChains, err := ci.ConnectedChains(e, chainsel)
	if err != nil {
		return nil, fmt.Errorf("failed to get connected chains for chain %d: %w", chainsel, err)
	}
	// get all supported tokens from token admin registry
	return adapters1_5.GetSupportedTokensPerRemoteChain(e.GetContext(), e.Logger, ci.TokenAdminReg, chain, remoteChains)
}

func (ci *ConfigImportAdapter) ConnectedChains(e cldf.Environment, chainsel uint64) ([]uint64, error) {
	// Fast path: return cached result if available to avoid duplicate RPC work.
	ci.connectedChainsMu.Lock()
	if ci.connectedChainsCache == nil {
		ci.connectedChainsCache = make(map[uint64][]uint64)
	}
	if cached, ok := ci.connectedChainsCache[chainsel]; ok {
		// Return a copy to prevent callers from mutating the cached slice.
		result := make([]uint64, len(cached))
		copy(result, cached)
		ci.connectedChainsMu.Unlock()
		return result, nil
	}
	ci.connectedChainsMu.Unlock()

	var connected []uint64
	laneResolver := adapters1_2.LaneVersionResolver{}
	remoteChainToVersionMap, _, err := laneResolver.DeriveLaneVersionsForChain(e, chainsel)
	if err != nil {
		return nil, fmt.Errorf("failed to derive lane versions for chain %d: %w", chainsel, err)
	}
	for destSel, version := range remoteChainToVersionMap {
		if version.Equal(semver.MustParse("1.6.0")) {
			connected = append(connected, destSel)
		}
	}

	// Cache the computed result for subsequent calls.
	ci.connectedChainsMu.Lock()
	if ci.connectedChainsCache == nil {
		ci.connectedChainsCache = make(map[uint64][]uint64)
	}
	cached := make([]uint64, len(connected))
	copy(cached, connected)
	ci.connectedChainsCache[chainsel] = cached
	ci.connectedChainsMu.Unlock()

	return connected, nil
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
