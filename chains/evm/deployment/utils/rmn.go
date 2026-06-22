package utils

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

const (
	rmnContractType       = "RMN"
	rmnRemoteContractType = "RMNRemote"
)

// ActiveRMNVersion returns the version of the active RMN on the given chain.
// It reads the address from RMNProxy when available, otherwise falls back to RMN/RMNRemote in the datastore.
func ActiveRMNVersion(e cldf.Environment, selector uint64) (*semver.Version, error) {
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return nil, fmt.Errorf("no EVM chain found for selector %d", selector)
	}

	if version, err := activeRMNVersionFromProxy(e, selector, chain); err == nil {
		return version, nil
	}

	return activeRMNVersionFromDatastore(e, selector, chain)
}

func activeRMNVersionFromProxy(e cldf.Environment, selector uint64, chain cldf_evm.Chain) (*semver.Version, error) {
	rmnProxyRef := datastore.AddressRef{
		Type:    datastore.ContractType(rmnproxyops.ContractType),
		Version: semver.MustParse("1.0.0"),
	}
	rmnProxyAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, rmnProxyRef, selector, evmds.ToEVMAddress)
	if err != nil {
		return nil, err
	}
	rmnProxyC, err := rmn_proxy_contract.NewRMNProxy(rmnProxyAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate RMNProxy contract at %s on chain %d: %w", rmnProxyAddr, selector, err)
	}
	rmnAddr, err := rmnProxyC.GetARM(&bind.CallOpts{Context: e.GetContext()})
	if err != nil {
		return nil, err
	}
	if rmnAddr == (common.Address{}) {
		return nil, fmt.Errorf("RMNProxy on chain %d has no active RMN set", selector)
	}
	_, version, err := TypeAndVersion(rmnAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to get type and version from RMN at %s on chain %d: %w", rmnAddr, selector, err)
	}
	return version, nil
}

func activeRMNVersionFromDatastore(e cldf.Environment, selector uint64, chain cldf_evm.Chain) (*semver.Version, error) {
	for _, contractType := range []string{rmnRemoteContractType, rmnContractType} {
		refs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(selector),
			datastore.AddressRefByType(datastore.ContractType(contractType)),
		)
		if len(refs) == 0 {
			continue
		}
		addr := common.HexToAddress(refs[0].Address)
		_, version, err := TypeAndVersion(addr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to get type and version from %s at %s on chain %d: %w", contractType, addr, selector, err)
		}
		return version, nil
	}
	return nil, fmt.Errorf("no active RMN found on chain %d", selector)
}
