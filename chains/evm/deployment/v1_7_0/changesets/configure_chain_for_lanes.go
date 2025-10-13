package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/off_ramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/on_ramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
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
	FeeQuoterDestChainConfig         fee_quoter.DestChainConfig
}

type ConfigureChainForLanesCfg struct {
	ChainSel     uint64
	RemoteChains map[uint64]RemoteChainConfig
	MCMSArgs     *mcms.Input
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
	ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureChainForLanesCfg) (sequences.ConfigureChainForLanesInput, error) {
		routerAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: cfg.ChainSel,
			Type:          datastore.ContractType(router.ContractType),
			Version:       semver.MustParse("1.2.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve router ref: %w", err)
		}
		onRampAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: cfg.ChainSel,
			Type:          datastore.ContractType(on_ramp.ContractType),
			Version:       semver.MustParse("1.7.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve onRamp ref: %w", err)
		}
		committeeVerifierAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: cfg.ChainSel,
			Type:          datastore.ContractType(committee_verifier.ContractType),
			Version:       semver.MustParse("1.7.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)

		feeQuoterAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: cfg.ChainSel,
			Type:          datastore.ContractType(fee_quoter.ContractType),
			Version:       semver.MustParse("1.7.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve fee quoter ref: %w", err)
		}
		offRampAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: cfg.ChainSel,
			Type:          datastore.ContractType(off_ramp.ContractType),
			Version:       semver.MustParse("1.7.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve off ramp ref: %w", err)
		}

		remoteChains := make(map[uint64]sequences.RemoteChainConfig, len(cfg.RemoteChains))
		for remoteChainSel, remoteConfig := range cfg.RemoteChains {
			executorOnRampAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteConfig.DefaultExecutor, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve executor on ramp ref: %w", err)
			}
			defaultCCVOnRamps := make([]common.Address, len(remoteConfig.DefaultCCVOnRamps))
			for i, ref := range remoteConfig.DefaultCCVOnRamps {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve ccv on ramp ref: %w", err)
				}
				defaultCCVOnRamps[i] = addr
			}
			defaultCCVOffRamps := make([]common.Address, len(remoteConfig.DefaultCCVOffRamps))
			for i, ref := range remoteConfig.DefaultCCVOffRamps {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve ccv off ramp ref: %w", err)
				}
				defaultCCVOffRamps[i] = addr
			}
			laneMandatedCCVOnRamps := make([]common.Address, len(remoteConfig.LaneMandatedCCVOnRamps))
			for i, ref := range remoteConfig.LaneMandatedCCVOnRamps {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve lane mandated ccv on ramp ref: %w", err)
				}
				laneMandatedCCVOnRamps[i] = addr
			}
			laneMandatedCCVOffRamps := make([]common.Address, len(remoteConfig.LaneMandatedCCVOffRamps))
			for i, ref := range remoteConfig.LaneMandatedCCVOffRamps {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve lane mandated ccv off ramp ref: %w", err)
				}
				laneMandatedCCVOffRamps[i] = addr
			}

			// TODO: CCIPMessageSource/Dest handling via ToPaddedEVMAddress is a hack, assumes the remote chain is also EVM.
			// Usage of cross-family changesets will resolve this issue, and this changeset will eventually be deprecated.
			ccipMessageSourceAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteConfig.CCIPMessageSource, remoteChainSel, evm_datastore_utils.ToPaddedEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve CCIPMessageSource ref: %w", err)
			}
			ccipMessageDestAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteConfig.CCIPMessageDest, remoteChainSel, evm_datastore_utils.ToPaddedEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve CCIPMessageDest ref: %w", err)
			}

			remoteChains[remoteChainSel] = sequences.RemoteChainConfig{
				AllowTrafficFrom:                 remoteConfig.AllowTrafficFrom,
				DefaultExecutor:                  executorOnRampAddr,
				DefaultCCVOffRamps:               defaultCCVOffRamps,
				LaneMandatedCCVOffRamps:          laneMandatedCCVOffRamps,
				DefaultCCVOnRamps:                defaultCCVOnRamps,
				LaneMandatedCCVOnRamps:           laneMandatedCCVOnRamps,
				CCIPMessageSource:                ccipMessageSourceAddr,
				CCIPMessageDest:                  ccipMessageDestAddr,
				CommitteeVerifierDestChainConfig: remoteConfig.CommitteeVerifierDestChainConfig,
				FeeQuoterDestChainConfig:         remoteConfig.FeeQuoterDestChainConfig,
			}
		}

		return sequences.ConfigureChainForLanesInput{
			ChainSelector:     cfg.ChainSel,
			Router:            routerAddr,
			OnRamp:            onRampAddr,
			CommitteeVerifier: committeeVerifierAddr,
			FeeQuoter:         feeQuoterAddr,
			OffRamp:           offRampAddr,
			RemoteChains:      remoteChains,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[ConfigureChainForLanesCfg],
})
