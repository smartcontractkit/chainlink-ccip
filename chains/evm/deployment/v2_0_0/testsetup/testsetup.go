package testsetup

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
)

// evmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector EVM")) = 0x2812d52c.
// Duplicated here to avoid an import cycle (testsetup ← adapters ← cctp ← testsetup).
var evmFamilySelector = [4]byte{0x28, 0x12, 0xd5, 0x2c}

// CreateBasicFeeQuoterDestChainConfig creates a basic fee quoter dest chain config with reasonable defaults for testing
func CreateBasicFeeQuoterDestChainConfig() lanes.FeeQuoterDestChainConfig {
	return lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DefaultTokenFeeUSDCents:     25,
		DestGasPerPayloadByteBase:   16,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		ChainFamilySelector:         binary.BigEndian.Uint32(evmFamilySelector[:]),
		V2Params: &lanes.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: 90,
			USDPerUnitGas:            big.NewInt(1e6),
		},
	}
}

// CreateBasicFeeQuoterDestChainConfigOverrides creates adapter-level FeeQuoter dest chain overrides for v2.0.0 lane tests.
func CreateBasicFeeQuoterDestChainConfigOverrides() changesetadapters.FeeQuoterDestChainConfigOverrides {
	return changesetadapters.FeeQuoterDestChainConfigOverrides{
		IsEnabled:                   new(true),
		MaxDataBytes:                new(uint32(30_000)),
		MaxPerMsgGasLimit:           new(uint32(3_000_000)),
		DestGasOverhead:             new(uint32(300_000)),
		DefaultTokenFeeUSDCents:     new(uint16(25)),
		DestGasPerPayloadByteBase:   new(uint8(16)),
		DefaultTokenDestGasOverhead: new(uint32(90_000)),
		DefaultTxGasLimit:           new(uint32(200_000)),
		NetworkFeeUSDCents:          new(uint16(10)),
		ChainFamilySelector:         evmFamilySelector,
		LinkFeeMultiplierPercent:    new(uint8(90)),
	}
}

// CreateBasicExecutorDestChainConfig creates a basic executor dest chain config with reasonable defaults for testing
func CreateBasicExecutorDestChainConfig() lanes.ExecutorDestChainConfig {
	return lanes.ExecutorDestChainConfig{
		USDCentsFee: 50,
		Enabled:     true,
	}
}

// CreateBasicCommitteeVerifierRemoteChainConfig creates a basic committee verifier remote chain config with reasonable defaults for testing
func CreateBasicCommitteeVerifierRemoteChainConfig() lanes.CommitteeVerifierRemoteChainConfig {
	return lanes.CommitteeVerifierRemoteChainConfig{
		AllowlistEnabled:   false,
		FeeUSDCents:        50,
		GasForVerification: 50_000,
		PayloadSizeBytes:   6*64 + 2*32,
		SignatureConfig: lanes.CommitteeVerifierSignatureQuorumConfig{
			Signers:   []string{common.HexToAddress("0x01").String()},
			Threshold: 1,
		},
	}
}

func AdapterCommitteeVerifierRemoteChainConfig() changesetadapters.CommitteeVerifierRemoteChainConfig {
	return changesetadapters.CommitteeVerifierRemoteChainConfig{
		AllowlistEnabled:   false,
		FeeUSDCents:        50,
		GasForVerification: 50_000,
		PayloadSizeBytes:   6*64 + 2*32,
		SignatureConfig: changesetadapters.CommitteeVerifierSignatureQuorumConfig{
			Signers:   []string{common.HexToAddress("0x01").String()},
			Threshold: 1,
		},
	}
}

// CreateBasicTokenTransferFeeConfig creates a basic token transfer fee config with reasonable defaults for testing
func CreateBasicTokenTransferFeeConfig() tokens.TokenTransferFeeConfig {
	return tokens.TokenTransferFeeConfig{
		IsEnabled:                     true,
		DestGasOverhead:               200_000,
		DestBytesOverhead:             32,
		DefaultFinalityFeeUSDCents:    100, // $1.00
		CustomFinalityFeeUSDCents:     200, // $2.00
		DefaultFinalityTransferFeeBps: 100, // 2%
		CustomFinalityTransferFeeBps:  100, // 1%
	}
}

// CreateRateLimiterConfig creates a rate limiter config
func CreateRateLimiterConfig(rate int64, capacity int64) tokens.RateLimiterConfig {
	if rate == 0 && capacity == 0 {
		return tokens.RateLimiterConfig{
			IsEnabled: false,
			Rate:      big.NewInt(0),
			Capacity:  big.NewInt(0),
		}
	}
	return tokens.RateLimiterConfig{
		IsEnabled: true,
		Rate:      big.NewInt(rate),
		Capacity:  big.NewInt(capacity),
	}
}

func CreateRateLimiterConfigFloatInput(rate float64, capacity float64) *tokens.RateLimiterConfigFloatInput {
	if rate == 0 && capacity == 0 {
		return &tokens.RateLimiterConfigFloatInput{
			IsEnabled: false,
			Rate:      0,
			Capacity:  0,
		}
	}
	return &tokens.RateLimiterConfigFloatInput{
		IsEnabled: true,
		Rate:      rate,
		Capacity:  capacity,
	}
}

// CreateBasicContractParams creates a basic set of contract deployment params with reasonable defaults for testing
func CreateBasicContractParams() sequences.ContractParams {
	usdPerLink, _ := new(big.Int).SetString("15000000000000000000", 10)   // $15
	usdPerWeth, _ := new(big.Int).SetString("2000000000000000000000", 10) // $2000

	return sequences.ContractParams{
		OffRamp: sequences.OffRampParams{
			Version:                   offramp.Version,
			GasForCallExactCheck:      5_000,
			MaxGasBufferToUpdateState: 12_000,
		},
		CommitteeVerifiers: []sequences.CommitteeVerifierParams{
			{
				Version:          committee_verifier.Version,
				FeeAggregator:    common.HexToAddress("0x01"),
				StorageLocations: []string{"https://test.chain.link.fake"},
				Qualifier:        "alpha",
			},
		},
		OnRamp: sequences.OnRampParams{
			Version:               onramp.Version,
			FeeAggregator:         common.HexToAddress("0x01"),
			MaxUSDCentsPerMessage: 100_00, // 100.00 USD
		},
		Executors: []sequences.ExecutorParams{
			{
				Version:       executor.Version,
				MaxCCVsPerMsg: 10,
				DynamicConfig: executor.DynamicConfig{
					FeeAggregator:         common.HexToAddress("0x01"),
					AllowedFinalityConfig: finality.Config{BlockDepth: 1}.Raw(),
					CcvAllowlistEnabled:   false,
				},
				Qualifier: "default",
			},
			{
				Version:       executor.Version,
				MaxCCVsPerMsg: 10,
				DynamicConfig: executor.DynamicConfig{
					FeeAggregator:         common.HexToAddress("0x01"),
					AllowedFinalityConfig: finality.Config{BlockDepth: 1}.Raw(),
					CcvAllowlistEnabled:   false,
				},
				Qualifier: "custom",
			},
		},
		FeeQuoter: sequences.FeeQuoterParams{
			Version:                        fee_quoter.Version,
			MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
			LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
			WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
			USDPerLINK:                     usdPerLink,
			USDPerWETH:                     usdPerWeth,
		},
		MockReceivers: []sequences.MockReceiverParams{
			{
				Version: mock_receiver.Version,
				RequiredVerifiers: []datastore.AddressRef{
					{
						// ChainSelector we don't know here but should still work.
						Type:      datastore.ContractType(committee_verifier.ContractType),
						Version:   committee_verifier.Version,
						Qualifier: "alpha",
					},
				},
			},
		},
	}
}

// BundleWithFreshReporter returns a new bundle with a fresh reporter.
// It takes the context function and logger from the inputted bundle.
// You may want to use this if performing state checks using operations,
// as you may inadvertently pull a report when you really want to re-check on-chain state.
func BundleWithFreshReporter(bundle operations.Bundle) operations.Bundle {
	return operations.NewBundle(bundle.GetContext, bundle.Logger, operations.NewMemoryReporter())
}

// UltraFastCurseMCMSRefs returns datastore refs seeding an Ultra Fast Curse RBACTimelock for the
// given chain.
//
// DeployChainContracts resolves this timelock and passes it as the RMN's curse admin, so every
// deployment needs one present. It is only ever used as a constructor argument — the sequence never
// calls the timelock — so tests can stand in a plain address ref instead of deploying a real MCMS
// instance. The address just has to be non-zero, since AuthorizedCallers rejects the zero address.
func UltraFastCurseMCMSRefs(chainSelector uint64) []datastore.AddressRef {
	return []datastore.AddressRef{
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(common_utils.RBACTimelock),
			Version:       mcms_ops.MCMSVersion,
			Qualifier:     common_utils.UltraFastCurseMCMSQualifier,
			Address:       common.HexToAddress("0x00000000000000000000000000000000000c0e5e").Hex(),
		},
	}
}

// SeedUltraFastCurseMCMS adds the Ultra Fast Curse RBACTimelock ref from UltraFastCurseMCMSRefs to
// an in-memory datastore, for tests that drive DeployChainContracts through a changeset or adapter
// (those read ExistingAddresses from the environment datastore rather than taking them directly).
func SeedUltraFastCurseMCMS(ds *datastore.MemoryDataStore, chainSelector uint64) error {
	for _, ref := range UltraFastCurseMCMSRefs(chainSelector) {
		if err := ds.Addresses().Add(ref); err != nil {
			return err
		}
	}

	return nil
}

// WithUltraFastCurseMCMS returns a sealed datastore containing everything in base plus the Ultra
// Fast Curse RBACTimelock ref for each given chain. Use it for tests that drive
// DeployChainContracts through a changeset, which reads ExistingAddresses off the environment
// datastore.
func WithUltraFastCurseMCMS(base datastore.DataStore, chainSelectors ...uint64) (datastore.DataStore, error) {
	ds := datastore.NewMemoryDataStore()
	if base != nil {
		if err := ds.Merge(base); err != nil {
			return nil, err
		}
	}
	for _, sel := range chainSelectors {
		if err := SeedUltraFastCurseMCMS(ds, sel); err != nil {
			return nil, err
		}
	}

	return ds.Seal(), nil
}
