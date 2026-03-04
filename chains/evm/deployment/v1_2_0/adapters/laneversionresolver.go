package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type LaneVersionResolver struct{}

// DeriveLaneVersionsForChain derives the versions of the lanes for a given chain by looking at the router's offramps and onramps.
// It returns a map of remote chain selector to lane version, a list of unique lane versions, and an error if any.
func (r *LaneVersionResolver) DeriveLaneVersionsForChain(e cldf.Environment, chainSel uint64) (map[uint64]*semver.Version, []*semver.Version, error) {
	// get the router
	routerAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:          datastore.ContractType(routerops.ContractType),
		Version:       routerops.Version,
		ChainSelector: chainSel,
	}, chainSel, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return nil, nil, err
	}
	chain, ok := e.BlockChains.EVMChains()[chainSel]
	if !ok {
		return nil, nil, fmt.Errorf("EVM chain with selector %d not found in environment", chainSel)
	}
	// get the offRamps
	routerC, err := router.NewRouter(routerAddr, chain.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to bind router contract at address %s: %w", routerAddr.Hex(), err)
	}
	// all lanes are bi-directional, so we can just check the offramps to determine which chains this router can talk to
	offRamps, err := routerC.GetOffRamps(&bind.CallOpts{
		Context: e.GetContext(),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get offramps from router at address %s for chain %d: %w", routerAddr.Hex(), chainSel, err)
	}
	remoteChains := make(map[uint64]struct{})
	for _, offRamp := range offRamps {
		if offRamp.OffRamp == (common.Address{}) {
			continue
		}
		remoteChains[offRamp.SourceChainSelector] = struct{}{}
	}
	versions := make(map[string]*semver.Version)
	laneVersionForRemoteChain := make(map[uint64]*semver.Version)
	// for all remote chains, find the onRamp and check its version , if unique add it to the list of versions to import
	for remoteChain := range remoteChains {
		onRamp, err := routerC.GetOnRamp(&bind.CallOpts{
			Context: e.GetContext(),
		}, remoteChain)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get onramp for remote chain %d from router at address %s for chain %d: %w", remoteChain, routerAddr.Hex(), chainSel, err)
		}
		if onRamp == (common.Address{}) {
			continue
		}
		_, version, err := utils.TypeAndVersion(onRamp, chain.Client)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get version for onramp at address %s on chain %d: %w", onRamp.Hex(), chainSel, err)
		}
		if _, exists := versions[version.String()]; !exists {
			versions[version.String()] = version
		}
		laneVersionForRemoteChain[remoteChain] = version
	}
	if len(versions) == 0 {
		return nil, nil, fmt.Errorf("version not found for any onramps connected to router at address %s for chain %d", routerAddr.Hex(), chainSel)
	}
	versionList := make([]*semver.Version, 0, len(versions))
	for _, version := range versions {
		versionList = append(versionList, version)
	}
	return laneVersionForRemoteChain, versionList, nil
}
