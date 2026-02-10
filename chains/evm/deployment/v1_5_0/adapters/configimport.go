package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	tokenadminops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	seq1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var GetTokensPaginationSize = uint64(20)

type ConfigImportAdapter struct {
	OnRamp        map[uint64]common.Address
	OffRamp       map[uint64]common.Address
	TokenAdminReg common.Address
	PriceRegistry common.Address
	Router        common.Address
}

func (ci *ConfigImportAdapter) InitializeAdapter(e cldf.Environment, sel uint64) error {
	ci.OnRamp = make(map[uint64]common.Address)
	onRampRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(onrampops.ContractType)),
		datastore.AddressRefByVersion(onrampops.Version),
		datastore.AddressRefByChainSelector(sel),
	)

	if len(onRampRefs) == 0 {
		return fmt.Errorf("failed to get onramp ref for chain %d", sel)
	}
	chain := e.BlockChains.EVMChains()[sel]
	for _, ref := range onRampRefs {
		onRampC, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(common.HexToAddress(ref.Address), chain.Client)
		if err != nil {
			return fmt.Errorf("failed to instantiate onramp contract for chain %d: %w", sel, err)
		}
		staticCfg, err := onRampC.GetStaticConfig(nil)
		if err != nil {
			return err
		}
		ci.OnRamp[staticCfg.DestChainSelector] = common.HexToAddress(ref.Address)
	}
	offRampRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(offrampops.ContractType)),
		datastore.AddressRefByVersion(offrampops.Version),
		datastore.AddressRefByChainSelector(sel),
	)
	if len(offRampRefs) == 0 {
		return fmt.Errorf("failed to get offramp ref for chain %d", sel)
	}
	for _, ref := range offRampRefs {
		offRampC, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(common.HexToAddress(ref.Address), chain.Client)
		if err != nil {
			return fmt.Errorf("failed to instantiate offramp contract for chain %d: %w", sel, err)
		}
		staticCfg, err := offRampC.GetStaticConfig(nil)
		if err != nil {
			return err
		}
		ci.OffRamp[staticCfg.SourceChainSelector] = common.HexToAddress(ref.Address)
	}
	tokenAdminRegRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(tokenadminops.ContractType),
		Version: tokenadminops.Version,
	}, sel, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find token admin registry contract ref for chain %d: %w", sel, err)
	}
	ci.TokenAdminReg = tokenAdminRegRef
	priceRegistryRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(priceregistryops.ContractType),
		Version: priceregistryops.Version,
	}, sel, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find price registry contract ref for chain %d: %w", sel, err)
	}
	ci.PriceRegistry = priceRegistryRef
	routerRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType("Router"),
		Version: semver.MustParse("1.2.0"),
	}, sel, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find router contract ref for chain %d: %w", sel, err)
	}
	ci.Router = routerRef
	return nil
}

func (ci *ConfigImportAdapter) ConnectedChains(e cldf.Environment, chainsel uint64) ([]uint64, error) {
	var connected []uint64
	// to ensure deduplication in case there are multiple onramps addresses in datastore for the same remote chain selector
	var mapConnectedChains = make(map[uint64]bool)
	chain, ok := e.BlockChains.EVMChains()[chainsel]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found in environment", chainsel)
	}
	routerC, err := router.NewRouter(ci.Router, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate router contract at %s on chain %d: %w", ci.Router.String(), chain.Selector, err)
	}
	for destSel, onrampForDest := range ci.OnRamp {
		onRamp, err := routerC.GetOnRamp(nil, destSel)
		if err != nil {
			return nil, fmt.Errorf("failed to get onramp for dest chain %d from router at %s on chain %d: %w", destSel, ci.Router.String(), chain.Selector, err)
		}
		// if the onramp address from the router doesn't match the onramp address we have, then this chain is not actually connected with 1.5
		if onRamp == onrampForDest && !mapConnectedChains[destSel] {
			connected = append(connected, destSel)
			mapConnectedChains[destSel] = true
		}
	}
	return connected, nil
}

func (ci *ConfigImportAdapter) SupportedTokensPerRemoteChain(e cldf.Environment, chainsel uint64) (map[uint64][]common.Address, error) {
	chain, ok := e.BlockChains.EVMChains()[chainsel]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found in environment", chainsel)
	}
	// get all supported tokens from token admin registry
	return GetSupportedTokensPerRemoteChain(ci.TokenAdminReg, chain)
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
			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, chains,
				seq1_5.OnRampImportConfigSequence,
				seq1_5.OnRampImportConfigSequenceInput{
					ChainSelector:           chainSelector,
					OnRampsPerRemoteChain:   ci.OnRamp,
					SupportedTokensPerChain: in.TokensPerRemoteChain,
					PriceRegistry:           ci.PriceRegistry,
				}, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to import onramp config for chain %d: %w", chainSelector, err)
			}
			result, err = sequences.RunAndMergeSequence(b, chains,
				seq1_5.OffRampImportConfigSequence,
				seq1_5.OffRampImportConfigSequenceInput{
					ChainSelector:          chainSelector,
					OffRampsPerRemoteChain: ci.OffRamp,
				}, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to import offramp config for chain %d: %w", chainSelector, err)
			}
			return result, nil
		})
}

func GetSupportedTokensPerRemoteChain(tokenAdminRegAddr common.Address, chain evm.Chain) (map[uint64][]common.Address, error) {
	// get all supported tokens from token admin registry
	tokenAdminRegC, err := token_admin_registry.NewTokenAdminRegistry(tokenAdminRegAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token admin registry contract at %s on chain %d: %w", tokenAdminRegAddr.String(), chain.Selector, err)
	}
	startIndex := uint64(0)
	allTokens := make([]common.Address, 0)
	for {
		fetchedTokens, err := tokenAdminRegC.GetAllConfiguredTokens(nil, startIndex, GetTokensPaginationSize)
		if err != nil {
			return nil, err
		}
		allTokens = append(allTokens, fetchedTokens...)
		startIndex += GetTokensPaginationSize
		if uint64(len(fetchedTokens)) < GetTokensPaginationSize {
			break
		}
	}
	pools, err := tokenAdminRegC.GetPools(nil, allTokens)
	if err != nil {
		return nil, fmt.Errorf("failed to get pools for tokens from token admin registry at %s on chain %d: %w", tokenAdminRegAddr.String(), chain.Selector, err)
	}
	tokensPerRemoteChain := make(map[uint64][]common.Address)
	for _, poolAddr := range pools {
		// there is no supported pool for this token
		if poolAddr == (common.Address{}) {
			continue
		}
		tokenPoolC, err := token_pool.NewTokenPool(poolAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to instantiate token pool contract at %s on chain %d: %w", poolAddr.String(), chain.Selector, err)
		}
		chains, err := tokenPoolC.GetSupportedChains(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get supported chains from token pool at %s on chain %d: %w", poolAddr.String(), chain.Selector, err)
		}
		tokenAddr, err := tokenPoolC.GetToken(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get token address from token pool at %s on chain %d: %w", poolAddr.String(), chain.Selector, err)
		}
		for _, remoteChain := range chains {
			tokensPerRemoteChain[remoteChain] = append(tokensPerRemoteChain[remoteChain], tokenAddr)
		}
	}
	return tokensPerRemoteChain, nil
}
