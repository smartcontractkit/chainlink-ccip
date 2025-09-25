package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
)

type RemoteChainConfig struct {
	AllowTrafficFrom                 bool
	CCIPMessageSource                datastore.AddressRef
	CCIPMessageDest                  datastore.AddressRef
	DefaultCCVOffRamps               []datastore.AddressRef
	LaneMandatedCCVOffRamps          []datastore.AddressRef
	DefaultCCVOnRamps                []datastore.AddressRef
	LaneMandatedCCVOnRamps           []datastore.AddressRef
	DefaultExecutor                  datastore.AddressRef
	CommitteeVerifierDestChainConfig sequences.CommitteeVerifierDestChainConfig
	FeeQuoterDestChainConfig         fee_quoter_v2.DestChainConfig
}

type ConfigureChainForLanesCfg struct {
	ChainSel     uint64
	RemoteChains map[uint64]RemoteChainConfig
	MCMSArgs     *changesets.MCMSInput
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
		staticAddrs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{
			{
				ChainSelector: cfg.ChainSel,
				Type:          datastore.ContractType(router.ContractType),
				Version:       semver.MustParse("1.2.0"),
			},
			{
				ChainSelector: cfg.ChainSel,
				Type:          datastore.ContractType(ccv_proxy.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
			{
				ChainSelector: cfg.ChainSel,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
			{
				ChainSelector: cfg.ChainSel,
				Type:          datastore.ContractType(fee_quoter_v2.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
			{
				ChainSelector: cfg.ChainSel,
				Type:          datastore.ContractType(ccv_aggregator.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
		}, datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve contract refs: %w", err)
		}
		routerAddr := staticAddrs[0]
		ccvProxyAddr := staticAddrs[1]
		committeeVerifierAddr := staticAddrs[2]
		feeQuoterAddr := staticAddrs[3]
		ccvAggregatorAddr := staticAddrs[4]

		remoteChains := make(map[uint64]sequences.RemoteChainConfig, len(cfg.RemoteChains))
		for remoteChainSel, remoteConfig := range cfg.RemoteChains {
			refs := []datastore.AddressRef{remoteConfig.DefaultExecutor}
			refs = append(refs, remoteConfig.DefaultCCVOffRamps...)
			refs = append(refs, remoteConfig.LaneMandatedCCVOffRamps...)
			refs = append(refs, remoteConfig.DefaultCCVOnRamps...)
			refs = append(refs, remoteConfig.LaneMandatedCCVOnRamps...)

			// Bookmarking for creating the final struct
			defaultCCVOffRampsEnd := 1 + len(remoteConfig.DefaultCCVOffRamps)
			laneMandatedCCVOffRampsEnd := defaultCCVOffRampsEnd + len(remoteConfig.LaneMandatedCCVOffRamps)
			defaultCCVOnRampsEnd := laneMandatedCCVOffRampsEnd + len(remoteConfig.DefaultCCVOnRamps)
			laneMandatedCCVOnRampsEnd := defaultCCVOnRampsEnd + len(remoteConfig.LaneMandatedCCVOnRamps)

			// Set the chain selector of every ref stored thus far to the chain that we are configuring
			for i := range refs {
				refs[i].ChainSelector = cfg.ChainSel
			}

			addrs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, refs, datastore_utils.ToEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve contract refs: %w", err)
			}

			// Resolve CCIPMessageSource separately since it is on the remote chain
			// TODO: Current CCIPMessageSource handling (via ToPaddedEVMAddress) is a hack that we need to resolve ASAP.
			// We need the cldf.BlockChain interface to support a method that converts an address from datastore to bytes
			// e.g.
			// for chainSel, chain := range e.BlockChains.All() {
			//   fmt.Println(chain.AddressToBytes(addresses[chainSel]))
			// }
			// Right now, we just assume that the remote chain is EVM, which is not correct.
			// Same goes for CCIPMessageDest.
			remoteConfig.CCIPMessageSource.ChainSelector = remoteChainSel
			remoteConfig.CCIPMessageDest.ChainSelector = remoteChainSel
			remoteAddrs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{
				remoteConfig.CCIPMessageSource,
				remoteConfig.CCIPMessageDest,
			}, datastore_utils.ToPaddedEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve CCIPMessageSource and CCIPMessageDest ref: %w", err)
			}

			remoteChains[remoteChainSel] = sequences.RemoteChainConfig{
				AllowTrafficFrom:                 remoteConfig.AllowTrafficFrom,
				DefaultExecutor:                  addrs[0],
				DefaultCCVOffRamps:               addrs[1:defaultCCVOffRampsEnd],
				LaneMandatedCCVOffRamps:          addrs[defaultCCVOffRampsEnd:laneMandatedCCVOffRampsEnd],
				DefaultCCVOnRamps:                addrs[laneMandatedCCVOffRampsEnd:defaultCCVOnRampsEnd],
				LaneMandatedCCVOnRamps:           addrs[defaultCCVOnRampsEnd:laneMandatedCCVOnRampsEnd],
				CCIPMessageSource:                remoteAddrs[0],
				CCIPMessageDest:                  remoteAddrs[1],
				CommitteeVerifierDestChainConfig: remoteConfig.CommitteeVerifierDestChainConfig,
				FeeQuoterDestChainConfig:         remoteConfig.FeeQuoterDestChainConfig,
			}
		}

		return sequences.ConfigureChainForLanesInput{
			ChainSelector:     cfg.ChainSel,
			Router:            routerAddr,
			CCVProxy:          ccvProxyAddr,
			CommitteeVerifier: committeeVerifierAddr,
			FeeQuoter:         feeQuoterAddr,
			CCVAggregator:     ccvAggregatorAddr,
			RemoteChains:      remoteChains,
		}, nil
	},
	ResolveDep: changesets.ResolveEVMChainDep[ConfigureChainForLanesCfg],
	ResolveMCMS: func(e cldf_deployment.Environment, cfg ConfigureChainForLanesCfg) (changesets.MCMSBuildParams, error) {
		return changesets.ResolveMCMS(e, changesets.NewEVMMCMBuilder(cfg.MCMSArgs))
	},
})
