package adapters

import (
	"context"
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"golang.org/x/sync/errgroup"

	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	tokenadminops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	seq1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	getTokensPaginationSize = uint64(20)
	// getSupportedTokensPoolConcurrency caps concurrent RPC calls when fetching supported tokens per pool.
	// Limits in-flight requests to avoid overwhelming the node/provider (rate limits, timeouts) and memory.
	getSupportedTokensPoolConcurrency = 10
)

type ConfigImportAdapter struct {
	OnRamp        map[uint64]common.Address
	OffRamp       map[uint64]common.Address
	TokenAdminReg common.Address
	PriceRegistry common.Address
	Router        common.Address

	// connectedChainsCache memoizes the result of ConnectedChains per chain selector
	// to avoid duplicate (potentially expensive) RPC work when the method is called
	// multiple times for the same chain within the same adapter instance.
	connectedChainsCache map[uint64][]uint64
	connectedChainsMu    sync.Mutex
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
	ci.OffRamp = make(map[uint64]common.Address)
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
		Type:    datastore.ContractType(routerops.ContractType),
		Version: routerops.Version,
	}, sel, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("failed to find router contract ref for chain %d: %w", sel, err)
	}
	ci.Router = routerRef
	return nil
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
		if version.Equal(semver.MustParse("1.5.0")) {
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
	return GetSupportedTokensPerRemoteChain(e.GetContext(), e.Logger, ci.TokenAdminReg, chain, remoteChains)
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

func GetSupportedTokensPerRemoteChain(ctx context.Context, l logger.Logger, tokenAdminRegAddr common.Address, chain evm.Chain, remoteChains []uint64) (map[uint64][]common.Address, error) {
	// get all supported tokens from token admin registry
	tokenAdminRegC, err := token_admin_registry.NewTokenAdminRegistry(tokenAdminRegAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token admin registry contract at %s on chain %d: %w", tokenAdminRegAddr.String(), chain.Selector, err)
	}
	startIndex := uint64(0)
	allTokens := make([]common.Address, 0)
	for {
		fetchedTokens, err := tokenAdminRegC.GetAllConfiguredTokens(nil, startIndex, getTokensPaginationSize)
		if err != nil {
			return nil, err
		}
		allTokens = append(allTokens, fetchedTokens...)
		startIndex += getTokensPaginationSize
		if uint64(len(fetchedTokens)) < getTokensPaginationSize {
			break
		}
	}

	tokensPerRemoteChain := make(map[uint64][]common.Address)
	var mu sync.Mutex
	grp, grpCtx := errgroup.WithContext(ctx)
	grp.SetLimit(getSupportedTokensPoolConcurrency)
	for _, tokenAddr := range allTokens {
		// there is no supported pool for this token
		if tokenAddr == (common.Address{}) {
			continue
		}
		tokenAddr := tokenAddr // capture loop variable
		grp.Go(func() error {
			poolAddr, err := tokenAdminRegC.GetPool(&bind.CallOpts{
				Context: grpCtx,
			}, tokenAddr)
			if err != nil {
				return fmt.Errorf("failed to get pool for token %s from token admin registry at %s on chain %d: %w",
					tokenAddr.String(), tokenAdminRegAddr.String(), chain.Selector, err)
			}
			if poolAddr == (common.Address{}) {
				// no pool configured for this token, skip
				return nil
			}
			tokenPoolC, err := token_pool.NewTokenPool(poolAddr, chain.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate token pool contract at %s on chain %d: %w", poolAddr.String(), chain.Selector, err)
			}

			// Cache the token address per pool so we only fetch it once, and
			// track when certain pool methods appear to be unsupported so we
			// can avoid repeated failed calls and warning spam.
			var (
				isSupportedChainUnsupported bool
			)

			for _, remoteChain := range remoteChains {
				// If we've already determined that IsSupportedChain
				// is unsupported for this pool, stop checking further chains.
				if isSupportedChainUnsupported {
					break
				}

				supported, err := tokenPoolC.IsSupportedChain(&bind.CallOpts{
					Context: grpCtx,
				}, remoteChain)
				if err != nil {
					// If we fail to check if the pool supports a remote chain,
					// assume this method isn't supported by this pool, log once,
					// and short-circuit to avoid failing the entire import and
					// spamming warnings for every remote chain.
					l.Warnf("failed to check if token pool at %s on chain %d supports remote chain %d: %v", poolAddr.String(), chain.Selector, remoteChain, err)
					isSupportedChainUnsupported = true
					break
				}
				if !supported {
					continue
				}

				mu.Lock()
				tokensPerRemoteChain[remoteChain] = append(tokensPerRemoteChain[remoteChain], tokenAddr)
				mu.Unlock()
			}
			return nil
		})
	}
	if err := grp.Wait(); err != nil {
		return nil, err
	}
	return tokensPerRemoteChain, nil
}
