package adapters

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	onramp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"golang.org/x/sync/errgroup"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct {
	// A map from src chain selector -> dst chain selector -> onRamp address. The address
	// will be nil if we have not fetched it from the chain yet. It'll be zero if we HAVE
	// looked it up but no connection exists between src <> dst. It'll be nonzero if some
	// connection between src <> dst exists.
	onRampCache map[uint64]map[uint64]*common.Address

	// A mutex to protect access to the onRampCache
	onRampMutex *sync.RWMutex

	// The max number of concurrent goroutines to use when building the cache. By default
	// the cache is constructed serially.
	concurrency int
}

func NewFeesAdapter(concurrency *int) *FeesAdapter {
	cnc := 1 // by default, construct the cache serially with a single goroutine
	if concurrency != nil {
		cnc = *concurrency
	}

	return &FeesAdapter{
		onRampCache: map[uint64]map[uint64]*common.Address{},
		onRampMutex: &sync.RWMutex{},
		concurrency: cnc,
	}
}

func (a *FeesAdapter) setOnRampAddressInCache(src uint64, dst uint64, addr *common.Address) {
	a.onRampMutex.Lock()
	defer a.onRampMutex.Unlock()

	if _, ok := a.onRampCache[src]; !ok {
		a.onRampCache[src] = map[uint64]*common.Address{dst: addr}
	} else {
		a.onRampCache[src][dst] = addr
	}
}

func (a *FeesAdapter) getOnRampAddressInCache(src uint64, dst uint64) *common.Address {
	a.onRampMutex.RLock()
	defer a.onRampMutex.RUnlock()

	entry, ok := a.onRampCache[src]
	if !ok {
		return nil
	}

	addr, ok := entry[dst]
	if !ok {
		return nil
	}

	return addr
}

// cacheOnRampAddress populates the onRampCache for the given src <> dst chain selectors
//
// NOTE: on v1.5 a single chain can have many EVM2EVMOnRamp contracts deployed, each one for a different destination
// chain. As a result, we'll need need to iterate over all the EVM2EVMOnRamp contracts deployed on the source chain,
// query the static config of each one, and compare the `DestChainSelector` field in the result with the one that we
// are provided to this function. This process can be quite expensive if there are many OnRamp contracts deployed on
// a single chain, which is the case for ETH testnet (249 on ramps at the time of this writing), and ETH mainnet (62
// onramps at the time of this writing). To resolve this we use caching + goroutines to speed up the process. It may
// also be worth investigating the possibility of using muticall3 in the future to batch these calls together.
func (a *FeesAdapter) cacheOnRampAddress(e cldf.Environment, src uint64, dst uint64) error {
	// If we already have the src <> dst onramp address cached, return early
	if a.getOnRampAddressInCache(src, dst) != nil {
		return nil
	}

	// Fetch all v1.5 OnRamp addresses for the source chain from the datastore
	filters := []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]{
		datastore.AddressRefByType(datastore.ContractType(onramp.ContractType)),
		datastore.AddressRefByVersion(onramp.Version),
		datastore.AddressRefByChainSelector(src),
	}

	// This will be used to short-circuit the goroutines when we find the onramp
	errShortCircuit := errors.New("onramp found")

	// Construct the onramp address cache concurrently - we stop as soon as we find the address
	grp, grpCtx := errgroup.WithContext(e.GetContext())
	grp.SetLimit(a.concurrency)
	for _, addrRef := range e.DataStore.Addresses().Filter(filters...) {
		addr, sel := addrRef.Address, addrRef.ChainSelector
		if !common.IsHexAddress(addr) {
			return fmt.Errorf("invalid OnRamp address %s for chain selector %d", addr, sel)
		}

		chain, ok := e.BlockChains.EVMChains()[sel]
		if !ok {
			return fmt.Errorf("chain with selector %d not defined", sel)
		}

		address := common.HexToAddress(addr)
		grp.Go(func() error {
			onramp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(address, chain.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate OnRamp contract at address %s on chain selector %d: %w", address, sel, err)
			}

			config, err := onramp.GetStaticConfig(&bind.CallOpts{Context: grpCtx})
			if err != nil {
				return fmt.Errorf("failed to get OnRamp static config for address %s on chain selector %d: %w", address, sel, err)
			}

			a.setOnRampAddressInCache(sel, config.DestChainSelector, &address)
			if config.DestChainSelector == dst {
				return errShortCircuit
			}

			return nil
		})
	}

	// Wait for all goroutines to finish
	if err := grp.Wait(); err != nil && !errors.Is(err, errShortCircuit) {
		return fmt.Errorf("failed to cache OnRamp addresses for src %d and dst %d: %w", src, dst, err)
	}

	// Cache a zero address if we did not find any onramp for the given src <> dst
	if a.getOnRampAddressInCache(src, dst) == nil {
		a.setOnRampAddressInCache(src, dst, &common.Address{})
	}

	// At this point, the cache will have an entry for src <> dst
	return nil
}

func (a *FeesAdapter) getOnRampAddress(e cldf.Environment, src uint64, dst uint64) (common.Address, error) {
	err := a.cacheOnRampAddress(e, src, dst)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to cache OnRamp address for src %d and dst %d: %w", src, dst, err)
	}

	maybeAddr := a.getOnRampAddressInCache(src, dst)
	if maybeAddr == nil {
		return common.Address{}, fmt.Errorf("impossible state reached - OnRamp address for src %d and dst %d not found in cache after caching", src, dst)
	}

	cacheAddr := *maybeAddr
	if cacheAddr == (common.Address{}) {
		return common.Address{}, fmt.Errorf("no OnRamp address found for src %d and dst %d", src, dst)
	}

	return cacheAddr, nil
}

func (a *FeesAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) fees.TokenTransferFeeArgs {
	minFeeUSDCents := uint32(25)

	// NOTE: we validate that src != dst so only one of these if statements will execute
	if src == chain_selectors.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 50
	}
	if dst == chain_selectors.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 150
	}

	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: 32,
		DestGasOverhead:   90_000,
		MinFeeUSDCents:    minFeeUSDCents,
		MaxFeeUSDCents:    math.MaxUint32,
		DeciBps:           0,
		IsEnabled:         true,
	}
}

func (a *FeesAdapter) GetOnchainTokenTransferFeeConfig(e cldf.Environment, src uint64, dst uint64, address string) (fees.TokenTransferFeeArgs, error) {
	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	onRampAddr, err := a.getOnRampAddress(e, src, dst)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get OnRamp address for src %d and dst %d: %w", src, dst, err)
	}

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, chain.Client)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to instantiate OnRamp contract at address %s on chain selector %d: %w", onRampAddr.Hex(), src, err)
	}

	if !common.IsHexAddress(address) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid contract address: %s", address)
	}

	// This gets the token transfer fee config for the given token from the EVM2EVMOnRamp contract
	// https://sepolia.etherscan.io/address/0xf9765c80F6448e6d4d02BeF4a6b4152131A2F513#code#F1#L719
	cfg, err := onRamp.GetTokenTransferFeeConfig(&bind.CallOpts{Context: e.GetContext()}, common.HexToAddress(address))
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get on-chain token transfer fee config for src %d, dst %d, and address %s: %w", src, dst, address, err)
	}

	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead,
		DestGasOverhead:   cfg.DestGasOverhead,
		MinFeeUSDCents:    cfg.MinFeeUSDCents,
		MaxFeeUSDCents:    cfg.MaxFeeUSDCents,
		IsEnabled:         cfg.IsEnabled,
		DeciBps:           cfg.DeciBps,
	}, nil
}

func (a *FeesAdapter) SetTokenTransferFee(e cldf.Environment) *operations.Sequence[fees.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		semver.MustParse("1.5.0"),
		"Set token transfer fee config on the OnRamp 1.5.0 contract across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			src := input.Selector

			for dst, dstCfg := range input.Settings {
				onRampAddr, err := a.getOnRampAddress(e, src, dst)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get OnRamp address for src %d and dst %d: %w", src, dst, err)
				}

				updatesByChain := onramp.SetTokenTransferFeeConfigInput{}
				for rawTokenAddress, feeCfg := range dstCfg {
					if !common.IsHexAddress(rawTokenAddress) {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					}

					token := common.HexToAddress(rawTokenAddress)
					if feeCfg == nil {
						updatesByChain.TokensToUseDefaultFeeConfigs = append(updatesByChain.TokensToUseDefaultFeeConfigs, token)
					} else {
						updatesByChain.TokenTransferFeeConfigArgs = append(
							updatesByChain.TokenTransferFeeConfigArgs,
							evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
								DestBytesOverhead: feeCfg.DestBytesOverhead,
								DestGasOverhead:   feeCfg.DestGasOverhead,
								MinFeeUSDCents:    feeCfg.MinFeeUSDCents,
								MaxFeeUSDCents:    feeCfg.MaxFeeUSDCents,
								DeciBps:           feeCfg.DeciBps,
								Token:             token,

								// NOTE: Aggregate rate limit should be false by default, we do
								// not do lane based rate limits anymore, we limit on pools now
								AggregateRateLimitEnabled: false,
							},
						)
					}
				}

				if len(updatesByChain.TokensToUseDefaultFeeConfigs) == 0 && len(updatesByChain.TokenTransferFeeConfigArgs) == 0 {
					continue
				}

				result, err = sequences.RunAndMergeSequence(b, chains,
					evmseq.OnRampSetTokenTransferFeeConfigSequence,
					evmseq.OnRampSetTokenTransferFeeConfigSequenceInput{
						UpdatesByChain: updatesByChain,
						ChainSelector:  input.Selector,
						Address:        onRampAddr,
					},
					result,
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set token transfer fee config on OnRamp contract at %s: %w", onRampAddr.Hex(), err)
				}
			}

			return result, nil
		},
	)
}
