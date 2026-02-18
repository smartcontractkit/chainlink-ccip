package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

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

	routerRef, err := findExactlyOneRef(e, sel, datastore.ContractType(router_ops.ContractType), router_ops.Version)
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

	onRampRef, err := findExactlyOneRef(e, sel, datastore.ContractType(onramp_ops.ContractType), onramp_ops.Version)
	if err != nil {
		return fmt.Errorf("cannot resolve EVM2EVMOnRamp to discover LINK: %w", err)
	}

	chain, ok := e.BlockChains.EVMChains()[sel]
	if !ok {
		return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
	}

	onRampAddr := common.HexToAddress(onRampRef.Address)
	onRampContract, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, chain.Client)
	if err != nil {
		return fmt.Errorf("failed to bind EVM2EVMOnRamp at %s: %w", onRampAddr.Hex(), err)
	}

	opts := &bind.CallOpts{Context: e.OperationsBundle.GetContext()}
	staticCfg, err := onRampContract.GetStaticConfig(opts)
	if err != nil {
		return fmt.Errorf("failed to call getStaticConfig on EVM2EVMOnRamp %s: %w", onRampAddr.Hex(), err)
	}

	if staticCfg.LinkToken == (common.Address{}) {
		return fmt.Errorf("EVM2EVMOnRamp %s returned zero address for LinkToken", onRampAddr.Hex())
	}

	e.Logger.Infof("Discovered LinkToken on chain %d: %s (via EVM2EVMOnRamp %s)", sel, staticCfg.LinkToken.Hex(), onRampAddr.Hex())

	return outputDs.Addresses().Add(datastore.AddressRef{
		Address:       staticCfg.LinkToken.Hex(),
		ChainSelector: sel,
		Type:          linkType,
		Version:       link_token.Version,
	})
}

func findExactlyOneRef(
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

	var valid []datastore.AddressRef
	for _, ref := range refs {
		if common.IsHexAddress(ref.Address) {
			valid = append(valid, ref)
		}
	}

	switch len(valid) {
	case 0:
		return datastore.AddressRef{}, fmt.Errorf("no %s v%s found on chain %d", contractType, version, sel)
	case 1:
		return valid[0], nil
	default:
		return datastore.AddressRef{}, fmt.Errorf("expected exactly one %s v%s on chain %d, found %d", contractType, version, sel, len(valid))
	}
}
