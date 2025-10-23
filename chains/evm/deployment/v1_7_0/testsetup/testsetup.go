package testsetup

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// CreateBasicFeeQuoterDestChainConfig creates a basic fee quoter dest chain config with reasonable defaults for testing
func CreateBasicFeeQuoterDestChainConfig() fee_quoter.DestChainConfig {
	return fee_quoter.DestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DefaultTokenFeeUSDCents:     25,
		DestGasPerPayloadByteBase:   16,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		ChainFamilySelector:         [4]byte{0x28, 0x12, 0xd5, 0x2c}, // EVM
	}
}

func CreateBasicExecutorDestChainConfig() executor.RemoteChainConfig {
	return executor.RemoteChainConfig{
		UsdCentsFee:            50,
		BaseExecGas:            100_000,
		DestAddressLengthBytes: 20,
		Enabled:                true,
	}
}

// CreateBasicContractParams creates a basic set of contract deployment params with reasonable defaults for testing
func CreateBasicContractParams() sequences.ContractParams {
	usdPerLink, _ := new(big.Int).SetString("15000000000000000000", 10)   // $15
	usdPerWeth, _ := new(big.Int).SetString("2000000000000000000000", 10) // $2000

	return sequences.ContractParams{
		RMNRemote: sequences.RMNRemoteParams{
			Version: semver.MustParse("1.6.0"),
		},
		OffRamp: sequences.OffRampParams{
			Version: semver.MustParse("1.7.0"),
		},
		CommitteeVerifier: []sequences.CommitteeVerifierParams{
			{
				Version:       semver.MustParse("1.7.0"),
				FeeAggregator: common.HexToAddress("0x01"),
				SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
					Threshold: 1,
					Signers: []common.Address{
						common.HexToAddress("0x02"),
						common.HexToAddress("0x03"),
						common.HexToAddress("0x04"),
						common.HexToAddress("0x05"),
					},
				},
				StorageLocation: "https://test.chain.link.fake",
				Qualifier:       "alpha",
			},
		},
		OnRamp: sequences.OnRampParams{
			Version:       semver.MustParse("1.7.0"),
			FeeAggregator: common.HexToAddress("0x01"),
		},
		Executor: sequences.ExecutorParams{
			Version:               semver.MustParse("1.7.0"),
			MaxCCVsPerMsg:         10,
			MinBlockConfirmations: 5,
		},
		FeeQuoter: sequences.FeeQuoterParams{
			Version:                        semver.MustParse("1.7.0"),
			MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
			LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
			WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
			USDPerLINK:                     usdPerLink,
			USDPerWETH:                     usdPerWeth,
		},
		MockReceivers: []sequences.MockReceiverParams{
			{
				Version: semver.MustParse("1.7.0"),
				RequiredVerifiers: []datastore.AddressRef{
					{
						// ChainSelector we don't know here but should still work.
						Type:      datastore.ContractType(committee_verifier.ContractType),
						Version:   semver.MustParse("1.7.0"),
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
