package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type RemoteChainConfig struct {
	AllowTrafficFrom                 bool
	CCIPMessageSource                datastore.AddressRef
	CCIPMessageDest                  datastore.AddressRef
	DefaultInboundCCVs               []datastore.AddressRef
	LaneMandatedInboundCCVs          []datastore.AddressRef
	DefaultOutboundCCVs              []datastore.AddressRef
	LaneMandatedOutboundCCVs         []datastore.AddressRef
	DefaultExecutor                  datastore.AddressRef
	CommitteeVerifierDestChainConfig sequences.CommitteeVerifierDestChainConfig
	FeeQuoterDestChainConfig         fee_quoter.DestChainConfig
}

type ConfigureChainForLanesCfg struct {
	ChainSel     uint64
	RemoteChains map[uint64]RemoteChainConfig
	// CommitteeVerifiers are the committee verifiers that will be configured by the CS.
	CommitteeVerifiers []datastore.AddressRef
	MCMSArgs           *mcms.Input
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
			Type:          datastore.ContractType(onramp.ContractType),
			Version:       semver.MustParse("1.7.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve onRamp ref: %w", err)
		}

		var committeeVerifiers []common.Address
		for _, cvRef := range cfg.CommitteeVerifiers {
			committeeVerifierAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, cvRef, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve committee verifier ref: %w", err)
			}
			committeeVerifiers = append(committeeVerifiers, committeeVerifierAddr)
		}

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
			Type:          datastore.ContractType(offramp.ContractType),
			Version:       semver.MustParse("1.7.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve offramp ref: %w", err)
		}

		remoteChains := make(map[uint64]sequences.RemoteChainConfig, len(cfg.RemoteChains))
		for remoteChainSel, remoteConfig := range cfg.RemoteChains {
			ExecutorAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteConfig.DefaultExecutor, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve executor ref: %w", err)
			}
			defaultOutboundCCVs := make([]common.Address, len(remoteConfig.DefaultOutboundCCVs))
			for i, ref := range remoteConfig.DefaultOutboundCCVs {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve outbound ccv ref: %w", err)
				}
				defaultOutboundCCVs[i] = addr
			}
			defaultInboundCCVs := make([]common.Address, len(remoteConfig.DefaultInboundCCVs))
			for i, ref := range remoteConfig.DefaultInboundCCVs {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve inbound ccv ref: %w", err)
				}
				defaultInboundCCVs[i] = addr
			}
			laneMandatedOutboundCCVs := make([]common.Address, len(remoteConfig.LaneMandatedOutboundCCVs))
			for i, ref := range remoteConfig.LaneMandatedOutboundCCVs {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve lane mandated outbound ccv ref: %w", err)
				}
				laneMandatedOutboundCCVs[i] = addr
			}
			laneMandatedInboundCCVs := make([]common.Address, len(remoteConfig.LaneMandatedInboundCCVs))
			for i, ref := range remoteConfig.LaneMandatedInboundCCVs {
				addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
				if err != nil {
					return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve lane mandated inbound ccv ref: %w", err)
				}
				laneMandatedInboundCCVs[i] = addr
			}

			// TODO: CCIPMessageSource/Dest handling via ToEVMAddressBytes is a hack, assumes the remote chain is also EVM.
			// Usage of cross-family changesets will resolve this issue, and this changeset will eventually be deprecated.
			ccipMessageSourceAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteConfig.CCIPMessageSource, remoteChainSel, evm_datastore_utils.ToEVMAddressBytes)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve CCIPMessageSource ref: %w", err)
			}
			ccipMessageDestAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteConfig.CCIPMessageDest, remoteChainSel, evm_datastore_utils.ToEVMAddressBytes)
			if err != nil {
				return sequences.ConfigureChainForLanesInput{}, fmt.Errorf("failed to resolve CCIPMessageDest ref: %w", err)
			}

			remoteChains[remoteChainSel] = sequences.RemoteChainConfig{
				AllowTrafficFrom:                 remoteConfig.AllowTrafficFrom,
				DefaultExecutor:                  ExecutorAddr,
				DefaultInboundCCVs:               defaultInboundCCVs,
				LaneMandatedInboundCCVs:          laneMandatedInboundCCVs,
				DefaultOutboundCCVs:              defaultOutboundCCVs,
				LaneMandatedOutboundCCVs:         laneMandatedOutboundCCVs,
				CCIPMessageSource:                ccipMessageSourceAddr,
				CCIPMessageDest:                  ccipMessageDestAddr,
				CommitteeVerifierDestChainConfig: remoteConfig.CommitteeVerifierDestChainConfig,
				FeeQuoterDestChainConfig:         remoteConfig.FeeQuoterDestChainConfig,
			}
		}

		return sequences.ConfigureChainForLanesInput{
			ChainSelector:      cfg.ChainSel,
			Router:             routerAddr,
			OnRamp:             onRampAddr,
			CommitteeVerifiers: committeeVerifiers,
			FeeQuoter:          feeQuoterAddr,
			OffRamp:            offRampAddr,
			RemoteChains:       remoteChains,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[ConfigureChainForLanesCfg],
})
