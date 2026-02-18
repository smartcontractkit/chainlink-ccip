package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	fee_quoter_ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	router_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	onramp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DiscoverTokensCfg struct {
	ChainSelectors []uint64
}

var DiscoverTokens = cldf_deployment.CreateChangeSet(applyDiscoverTokens, validateDiscoverTokens)

func validateDiscoverTokens(e cldf_deployment.Environment, cfg DiscoverTokensCfg) error {
	if len(cfg.ChainSelectors) == 0 {
		return fmt.Errorf("at least one chain selector is required")
	}
	evmChains := e.BlockChains.EVMChains()
	seen := make(map[uint64]bool, len(cfg.ChainSelectors))
	for _, sel := range cfg.ChainSelectors {
		if seen[sel] {
			return fmt.Errorf("duplicate chain selector %d", sel)
		}
		seen[sel] = true
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
	}
	return nil
}

func applyDiscoverTokens(e cldf_deployment.Environment, cfg DiscoverTokensCfg) (cldf_deployment.ChangesetOutput, error) {
	outputDs := datastore.NewMemoryDataStore()

	for _, sel := range cfg.ChainSelectors {
		if err := maybeDiscoverWETH(e, sel, outputDs); err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to discover WETH on chain %d: %w", sel, err)
		}

		if err := maybeDiscoverLINK(e, sel, outputDs); err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to discover LINK on chain %d: %w", sel, err)
		}
	}

	return cldf_deployment.ChangesetOutput{DataStore: outputDs}, nil
}

func maybeDiscoverWETH(
	e cldf_deployment.Environment,
	sel uint64,
	outputDs datastore.MutableDataStore,
) error {
	wethType := datastore.ContractType(weth.ContractType)

	existing := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(sel),
		datastore.AddressRefByType(wethType),
		datastore.AddressRefByVersion(weth.Version),
	)
	if len(existing) > 0 {
		e.Logger.Infof("WETH9 already exists on chain %d at %s, skipping", sel, existing[0].Address)
		return nil
	}

	routerRef, err := findFirstRef(e, sel, datastore.ContractType(router_ops.ContractType), router_ops.Version)
	if err != nil {
		return fmt.Errorf("cannot resolve Router to discover WETH: %w", err)
	}

	chain, ok := e.BlockChains.EVMChains()[sel]
	if !ok {
		return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
	}

	routerAddr := common.HexToAddress(routerRef.Address)
	routerContract, err := router.NewRouter(routerAddr, chain.Client)
	if err != nil {
		return fmt.Errorf("failed to bind Router at %s: %w", routerAddr.Hex(), err)
	}

	opts := &bind.CallOpts{Context: e.OperationsBundle.GetContext()}
	wethAddr, err := routerContract.GetWrappedNative(opts)
	if err != nil {
		return fmt.Errorf("failed to call getWrappedNative on Router %s: %w", routerAddr.Hex(), err)
	}

	if wethAddr == (common.Address{}) {
		return fmt.Errorf("Router %s returned zero address for wrappedNative", routerAddr.Hex())
	}

	e.Logger.Infof("Discovered WETH9 on chain %d: %s (via Router %s)", sel, wethAddr.Hex(), routerAddr.Hex())

	return outputDs.Addresses().Add(datastore.AddressRef{
		Address:       wethAddr.Hex(),
		ChainSelector: sel,
		Type:          wethType,
		Version:       weth.Version,
	})
}

func maybeDiscoverLINK(
	e cldf_deployment.Environment,
	sel uint64,
	outputDs datastore.MutableDataStore,
) error {
	linkType := datastore.ContractType(link_token.ContractType)

	existing := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(sel),
		datastore.AddressRefByType(linkType),
		datastore.AddressRefByVersion(link_token.Version),
	)
	if len(existing) > 0 {
		e.Logger.Infof("LinkToken already exists on chain %d at %s, skipping", sel, existing[0].Address)
		return nil
	}

	chain, ok := e.BlockChains.EVMChains()[sel]
	if !ok {
		return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
	}

	linkAddr, onRampErr := discoverLINKFromOnRamp(e, sel, chain.Client)
	if onRampErr != nil {
		e.Logger.Infof("EVM2EVMOnRamp not available on chain %d, falling back to FeeQuoter: %v", sel, onRampErr)

		var fqErr error
		linkAddr, fqErr = discoverLINKFromFeeQuoter(e, sel, chain.Client)
		if fqErr != nil {
			return fmt.Errorf("cannot discover LINK from either EVM2EVMOnRamp (%v) or FeeQuoter (%v)", onRampErr, fqErr)
		}
	}

	return outputDs.Addresses().Add(datastore.AddressRef{
		Address:       linkAddr.Hex(),
		ChainSelector: sel,
		Type:          linkType,
		Version:       link_token.Version,
	})
}

func discoverLINKFromOnRamp(e cldf_deployment.Environment, sel uint64, client bind.ContractBackend) (common.Address, error) {
	onRampRef, err := findFirstRef(e, sel, datastore.ContractType(onramp_ops.ContractType), onramp_ops.Version)
	if err != nil {
		return common.Address{}, err
	}

	onRampAddr := common.HexToAddress(onRampRef.Address)
	onRampContract, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to bind EVM2EVMOnRamp at %s: %w", onRampAddr.Hex(), err)
	}

	opts := &bind.CallOpts{Context: e.OperationsBundle.GetContext()}
	staticCfg, err := onRampContract.GetStaticConfig(opts)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call getStaticConfig on EVM2EVMOnRamp %s: %w", onRampAddr.Hex(), err)
	}

	if staticCfg.LinkToken == (common.Address{}) {
		return common.Address{}, fmt.Errorf("EVM2EVMOnRamp %s returned zero address for LinkToken", onRampAddr.Hex())
	}

	e.Logger.Infof("Discovered LinkToken on chain %d: %s (via EVM2EVMOnRamp %s)", sel, staticCfg.LinkToken.Hex(), onRampAddr.Hex())
	return staticCfg.LinkToken, nil
}

func discoverLINKFromFeeQuoter(e cldf_deployment.Environment, sel uint64, client bind.ContractBackend) (common.Address, error) {
	fqType := datastore.ContractType(fee_quoter_ops.ContractType)

	// Search any FeeQuoter on this chain regardless of version.
	// The v1.7.0 gobinding is ABI-compatible with v1.6.x for getStaticConfig().LinkToken.
	fqRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(sel),
		datastore.AddressRefByType(fqType),
	)

	for _, ref := range fqRefs {
		if !common.IsHexAddress(ref.Address) {
			continue
		}

		fqAddr := common.HexToAddress(ref.Address)
		fqContract, err := fee_quoter.NewFeeQuoter(fqAddr, client)
		if err != nil {
			e.Logger.Warnf("Failed to bind FeeQuoter at %s: %v, trying next", fqAddr.Hex(), err)
			continue
		}

		opts := &bind.CallOpts{Context: e.OperationsBundle.GetContext()}
		staticCfg, err := fqContract.GetStaticConfig(opts)
		if err != nil {
			e.Logger.Warnf("Failed to call getStaticConfig on FeeQuoter %s: %v, trying next", fqAddr.Hex(), err)
			continue
		}

		if staticCfg.LinkToken == (common.Address{}) {
			e.Logger.Debugf("FeeQuoter %s returned zero LinkToken, skipping", fqAddr.Hex())
			continue
		}

		e.Logger.Infof("Discovered LinkToken on chain %d: %s (via FeeQuoter %s v%s)", sel, staticCfg.LinkToken.Hex(), fqAddr.Hex(), ref.Version)
		return staticCfg.LinkToken, nil
	}

	return common.Address{}, fmt.Errorf("no FeeQuoter with valid LinkToken found on chain %d", sel)
}

func findFirstRef(
	e cldf_deployment.Environment,
	sel uint64,
	contractType datastore.ContractType,
	version *semver.Version,
) (datastore.AddressRef, error) {
	refs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(sel),
		datastore.AddressRefByType(contractType),
		datastore.AddressRefByVersion(version),
	)

	for _, ref := range refs {
		if common.IsHexAddress(ref.Address) && common.HexToAddress(ref.Address) != (common.Address{}) {
			return ref, nil
		}
	}

	return datastore.AddressRef{}, fmt.Errorf("no %s v%s found on chain %d", contractType, version, sel)
}
