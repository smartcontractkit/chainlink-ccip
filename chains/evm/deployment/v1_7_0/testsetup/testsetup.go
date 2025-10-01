package testsetup

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// CreateBasicFeeQuoterDestChainConfig creates a basic fee quoter dest chain config with reasonable defaults for testing
func CreateBasicFeeQuoterDestChainConfig() fee_quoter_v2.DestChainConfig {
	return fee_quoter_v2.DestChainConfig{
		IsEnabled:                         true,
		MaxNumberOfTokensPerMsg:           10,
		MaxDataBytes:                      30_000,
		MaxPerMsgGasLimit:                 3_000_000,
		DestGasOverhead:                   300_000,
		DefaultTokenFeeUSDCents:           25,
		DestGasPerPayloadByteBase:         16,
		DestGasPerPayloadByteHigh:         40,
		DestGasPerPayloadByteThreshold:    3000,
		DestDataAvailabilityOverheadGas:   100,
		DestGasPerDataAvailabilityByte:    16,
		DestDataAvailabilityMultiplierBps: 1,
		DefaultTokenDestGasOverhead:       90_000,
		DefaultTxGasLimit:                 200_000,
		GasMultiplierWeiPerEth:            11e17, // Gas multiplier in wei per eth is scaled by 1e18, so 11e17 is 1.1 = 110%
		NetworkFeeUSDCents:                10,
		ChainFamilySelector:               [4]byte{0x28, 0x12, 0xd5, 0x2c}, // EVM
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
		CCVAggregator: sequences.CCVAggregatorParams{
			Version: semver.MustParse("1.7.0"),
		},
		CommitteeVerifier: sequences.CommitteeVerifierParams{
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
		},
		CCVProxy: sequences.CCVProxyParams{
			Version:       semver.MustParse("1.7.0"),
			FeeAggregator: common.HexToAddress("0x01"),
		},
		ExecutorOnRamp: sequences.ExecutorOnRampParams{
			Version:       semver.MustParse("1.7.0"),
			MaxCCVsPerMsg: 10,
		},
		FeeQuoter: sequences.FeeQuoterParams{
			Version:                        semver.MustParse("1.7.0"),
			MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
			TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
			LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
			WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
			USDPerLINK:                     usdPerLink,
			USDPerWETH:                     usdPerWeth,
		},
	}
}

// CreateEnvironment creates a test deployment environment with the given EVM chains
func CreateEnvironment(t *testing.T, evmChains map[uint64]provider.SimChainProviderConfig) (deployment.Environment, error) {
	lggr, err := logger.New()
	if err != nil {
		return deployment.Environment{}, fmt.Errorf("failed to create logger: %w", err)
	}

	bundle := operations.NewBundle(
		func() context.Context { return t.Context() },
		lggr,
		operations.NewMemoryReporter(),
	)

	blockChains := make([]chain.BlockChain, 0, len(evmChains))
	for chainSel, cfg := range evmChains {
		blockChain, err := provider.NewSimChainProvider(t, chainSel, cfg).Initialize(t.Context())
		if err != nil {
			return deployment.Environment{}, fmt.Errorf("failed to create chain provider: %w", err)
		}
		blockChains = append(blockChains, blockChain)
	}

	return deployment.Environment{
		GetContext:       func() context.Context { return t.Context() },
		Logger:           lggr,
		OperationsBundle: bundle,
		BlockChains:      chain.NewBlockChainsFromSlice(blockChains),
		DataStore:        datastore.NewMemoryDataStore().Seal(),
	}, nil
}
