package adapters_test

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	offrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// importConfigAddressRefs are address refs for SetContractParamsFromImportedConfig.
// Chain 5009297550715157269 has OnRamp 1.6.0 + OffRamp 1.6.0 (and optionally RMN 1.5.0).
// Chain 4949039107694359620 has only RMN 1.5.0 (v1.6 path returns empty; v1.5 path fills LegacyRMN and GasForCallExactCheck).
var (
	legacyRMNAddress1       = common.HexToAddress("0x8888888888888888888888888888888888888888")
	legacyRMNAddress2       = common.HexToAddress("0x9999999999999999999999999999999999999999")
	feeAggregatorAddress    = common.HexToAddress("0xFEE0000000000000000000000000000000000001")
	importConfigAddressRefs = []datastore.AddressRef{
		{Address: "0x6666666666666666666666666666666666666666", ChainSelector: 5009297550715157269, Type: datastore.ContractType(onrampops_v160.ContractType), Version: onrampops_v160.Version},
		{Address: "0x7777777777777777777777777777777777777777", ChainSelector: 5009297550715157269, Type: datastore.ContractType(offrampops_v160.ContractType), Version: offrampops_v160.Version},
		{Address: legacyRMNAddress1.String(), ChainSelector: 5009297550715157269, Type: datastore.ContractType(rmnops1_5.ContractType), Version: semver.MustParse("1.5.0")},
		{Address: legacyRMNAddress2.String(), ChainSelector: 4949039107694359620, Type: datastore.ContractType(rmnops1_5.ContractType), Version: semver.MustParse("1.5.0")},
	}

	// importConfigContractMetadata holds OnRamp 1.6.0 and OffRamp 1.6.0 metadata for the chain that has both.
	importConfigContractMetadata = []datastore.ContractMetadata{
		{
			Address:       "0x6666666666666666666666666666666666666666",
			ChainSelector: 5009297550715157269,
			Metadata: seq1_6.OnRampImportConfigSequenceOutput{
				DestChainCfgs: map[uint64]onrampops_v160.GetDestChainConfigResult{},
				StaticConfig: onrampops_v160.StaticConfig{
					ChainSelector:      5009297550715157269,
					RmnRemote:          common.HexToAddress("0x8888888888888888888888888888888888888888"),
					NonceManager:       common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
					TokenAdminRegistry: common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
				},
				DynamicConfig: onrampops_v160.DynamicConfig{
					FeeQuoter:          common.HexToAddress("0x1111111111111111111111111111111111111111"),
					FeeAggregator:      feeAggregatorAddress,
					AllowlistAdmin:     common.Address{},
					MessageInterceptor: common.Address{},
				},
			},
		},
		{
			Address:       "0x7777777777777777777777777777777777777777",
			ChainSelector: 5009297550715157269,
			Metadata: seq1_6.OffRampImportConfigSequenceOutput{
				SourceChainCfgs: map[uint64]offrampops_v160.SourceChainConfig{},
				StaticConfig: offrampops_v160.StaticConfig{
					ChainSelector:        5009297550715157269,
					GasForCallExactCheck: 6000,
					RmnRemote:            common.HexToAddress("0x8888888888888888888888888888888888888888"),
					TokenAdminRegistry:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
					NonceManager:         common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				},
				DynamicConfig: offrampops_v160.DynamicConfig{},
			},
		},
	}
)

// dummyDeployContractParams returns params with Executors, CommitteeVerifiers, OnRamp, and OffRamp
// populated with dummy values. OnRamp/OffRamp use zero/empty so that when merged as source into
// generated config they do not overwrite the imported OnRamp/OffRamp/RMN (per MergeWithOverrideIfNotEmpty behavior).
func dummyDeployContractParams() ccvadapters.DeployContractParams {
	v170 := semver.MustParse("2.0.0")
	return ccvadapters.DeployContractParams{
		OnRamp: ccvadapters.OnRampDeployParams{
			Version:       v170,
			FeeAggregator: "0xDummyOnRampFeeAgg",
		},
		OffRamp: ccvadapters.OffRampDeployParams{
			Version: v170,
			// GasForCallExactCheck etc. left zero so merge keeps generated (imported) values.
		},
		Executors: []ccvadapters.ExecutorDeployParams{
			{
				Version:       v170,
				MaxCCVsPerMsg: 3,
				DynamicConfig: ccvadapters.ExecutorDynamicDeployConfig{
					FeeAggregator:         "0xDummyExecutorFeeAgg",
					AllowedFinalityConfig: finality.Config{BlockDepth: 5},
					CcvAllowlistEnabled:   true,
				},
				Qualifier: "dummy-exec",
			},
		},
		CommitteeVerifiers: []ccvadapters.CommitteeVerifierDeployParams{
			{
				Version:          v170,
				FeeAggregator:    "0xDummyCVFeeAgg",
				StorageLocations: []string{"https://dummy.store"},
				Qualifier:        "dummy-cv",
			},
		},
	}
}
