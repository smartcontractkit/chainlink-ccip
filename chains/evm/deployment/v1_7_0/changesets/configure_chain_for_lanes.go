package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type RemoteChainConfig struct {
	AllowTrafficFrom            bool
	CCIPMessageSource           changesets.TypeAndVersion
	DefaultCCVOffRamps          []changesets.TypeAndVersion
	LaneMandatedCCVOffRamps     []changesets.TypeAndVersion
	DefaultCCVOnRamps           []changesets.TypeAndVersion
	LaneMandatedCCVOnRamps      []changesets.TypeAndVersion
	DefaultExecutor             changesets.TypeAndVersion
	CommitOnRampDestChainConfig sequences.CommitOnRampDestChainConfig
	FeeQuoterDestChainConfig    fee_quoter_v2.DestChainConfig
}

type ConfigureChainForLanesCfg struct {
	ChainSel     uint64
	RemoteChains map[uint64]RemoteChainConfig
}

func (c ConfigureChainForLanesCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var ConfigureChainForLanes = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.ConfigureChainForLanesInput,
	evm.Chain,
	ConfigureChainForLanesCfg,
]{
	Sequence: sequences.ConfigureChainForLanes,
	Describe: func(in sequences.ConfigureChainForLanesInput, dep evm.Chain) string {
		remoteChains := make([]uint64, 0, len(in.RemoteChains))
		for chainSel := range in.RemoteChains {
			remoteChains = append(remoteChains, chainSel)
		}
		return fmt.Sprintf("Configure remote chain connnections on %s: %v", dep, remoteChains)
	},
	ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureChainForLanesCfg) (sequences.ConfigureChainForLanesInput, error) {
		routerAddr, err := getSingleAddress(e, cfg.ChainSel, datastore.ContractType(router.ContractType), semver.MustParse("1.2.0"))
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, err
		}
		ccvProxyAddr, err := getSingleAddress(e, cfg.ChainSel, datastore.ContractType(ccv_proxy.ContractType), semver.MustParse("1.7.0"))
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, err
		}
		commitOnRampAddr, err := getSingleAddress(e, cfg.ChainSel, datastore.ContractType(commit_onramp.ContractType), semver.MustParse("1.7.0"))
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, err
		}
		feeQuoterAddr, err := getSingleAddress(e, cfg.ChainSel, datastore.ContractType(fee_quoter_v2.ContractType), semver.MustParse("1.7.0"))
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, err
		}
		ccvAggregatorAddr, err := getSingleAddress(e, cfg.ChainSel, datastore.ContractType(ccv_aggregator.ContractType), semver.MustParse("1.7.0"))
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, err
		}

		remoteChains := make(map[uint64]sequences.RemoteChainConfig, len(cfg.RemoteChains))
		for remoteChainSel, remoteConfig := range cfg.RemoteChains {
			ccipMessageSourceAddr, err := getSingleAddress(e, remoteChainSel, remoteConfig.CCIPMessageSource.Type, remoteConfig.CCIPMessageSource.Version)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, err
			}
			defaultCCVOffRamps, err := resolveRefs(e, remoteChainSel, remoteConfig.DefaultCCVOffRamps)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, err
			}
			laneMandatedCCVOffRamps, err := resolveRefs(e, remoteChainSel, remoteConfig.LaneMandatedCCVOffRamps)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, err
			}
			defaultCCVOnRamps, err := resolveRefs(e, remoteChainSel, remoteConfig.DefaultCCVOnRamps)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, err
			}
			laneMandatedCCVOnRamps, err := resolveRefs(e, remoteChainSel, remoteConfig.LaneMandatedCCVOnRamps)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, err
			}
			defaultExecutorAddr, err := getSingleAddress(e, remoteChainSel, remoteConfig.DefaultExecutor.Type, remoteConfig.DefaultExecutor.Version)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, err
			}

			remoteChains[remoteChainSel] = sequences.RemoteChainConfig{
				AllowTrafficFrom: remoteConfig.AllowTrafficFrom,
				// TODO: Current CCIPMessageSource handling is a hack that we need to resolve ASAP.
				// We need the cldf.BlockChain interface to support a method that converts an address from datastore to bytes
				// e.g.
				// for chainSel, chain := range e.BlockChains.All() {
				//   fmt.Println(chain.AddressToBytes(addresses[chainSel]))
				// }
				// Right now, we just assume that the remote chain is EVM, which is not correct.
				CCIPMessageSource:           common.LeftPadBytes(common.HexToAddress(ccipMessageSourceAddr).Bytes(), 32),
				DefaultCCVOffRamps:          defaultCCVOffRamps,
				LaneMandatedCCVOffRamps:     laneMandatedCCVOffRamps,
				DefaultCCVOnRamps:           defaultCCVOnRamps,
				LaneMandatedCCVOnRamps:      laneMandatedCCVOnRamps,
				DefaultExecutor:             common.HexToAddress(defaultExecutorAddr),
				CommitOnRampDestChainConfig: remoteConfig.CommitOnRampDestChainConfig,
				FeeQuoterDestChainConfig:    remoteConfig.FeeQuoterDestChainConfig,
			}
		}

		return sequences.ConfigureChainForLanesInput{
			ChainSelector: cfg.ChainSel,
			Router:        common.HexToAddress(routerAddr),
			CCVProxy:      common.HexToAddress(ccvProxyAddr),
			CommitOnRamp:  common.HexToAddress(commitOnRampAddr),
			FeeQuoter:     common.HexToAddress(feeQuoterAddr),
			CCVAggregator: common.HexToAddress(ccvAggregatorAddr),
			RemoteChains:  remoteChains,
		}, nil
	},
	ResolveDep: changesets.ResolveEVMChainDep[ConfigureChainForLanesCfg],
})

// Helper to fetch a single address for a given selector, type, and version
func getSingleAddress(e cldf_deployment.Environment, selector uint64, typ datastore.ContractType, version *semver.Version) (string, error) {
	addrs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(typ),
		datastore.AddressRefByVersion(version),
	)
	if len(addrs) != 1 {
		return "", fmt.Errorf("expected to find exactly one %s %s on chain with selector %d, found %d", typ, version, selector, len(addrs))
	}
	return addrs[0].Address, nil
}

// Helper to map []TypeAndVersion to []common.Address
func resolveRefs(e cldf_deployment.Environment, selector uint64, refs []changesets.TypeAndVersion) ([]common.Address, error) {
	addrs := make([]common.Address, 0, len(refs))
	for _, ref := range refs {
		addr, err := getSingleAddress(e, selector, ref.Type, ref.Version)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, common.HexToAddress(addr))
	}
	return addrs, nil
}
