package adapters

import (
	"math/big"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

const (
	DefaultQualifier = "default"
)

func defaultDeployContractParams() ccvadapters.DeployContractParams {
	usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
	if !ok {
		panic("invalid usdPerLink constant")
	}
	usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
	if !ok {
		panic("invalid usdPerWeth constant")
	}
	return ccvadapters.DeployContractParams{
		RMNRemote: ccvadapters.RMNRemoteDeployParams{
			Version: rmn_remote.Version,
		},
		OffRamp: ccvadapters.OffRampDeployParams{
			Version:                   offramp.Version,
			GasForCallExactCheck:      5_000,
			MaxGasBufferToUpdateState: 12_000,
		},
		OnRamp: ccvadapters.OnRampDeployParams{
			Version:               onramp.Version,
			MaxUSDCentsPerMessage: 100_00, // 100.00 USD
		},
		FeeQuoter: ccvadapters.FeeQuoterDeployParams{
			Version:                        fee_quoter.Version,
			MaxFeeJuelsPerMsg:              new(big.Int).Mul(big.NewInt(2e2), big.NewInt(1e18)),
			LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
			WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
			USDPerLINK:                     usdPerLink,
			USDPerWETH:                     usdPerWeth,
		},
		MockReceivers: defaultMockReceiverParams(),
	}
}

func defaultFinalityConfig() finality.Config {
	return finality.Config{BlockDepth: 1}
}

func defaultExecutorParams(feeAggr string) []ccvadapters.ExecutorDeployParams {
	return []ccvadapters.ExecutorDeployParams{
		{Version: executor.Version,
			MaxCCVsPerMsg: 10,
			DynamicConfig: ccvadapters.ExecutorDynamicDeployConfig{
				FeeAggregator:         feeAggr,
				AllowedFinalityConfig: defaultFinalityConfig(),
				CcvAllowlistEnabled:   false,
			},
			Qualifier: DefaultQualifier,
		},
	}
}

func defaultMockReceiverParams() []ccvadapters.MockReceiverDeployParams {
	return []ccvadapters.MockReceiverDeployParams{
		{
			Version: mock_receiver.Version,
			RequiredVerifiers: []datastore.AddressRef{
				{
					Qualifier: DefaultQualifier,
					Type:      datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
					Version:   versioned_verifier_resolver.Version,
				},
			},
			AllowedFinalityConfig: defaultFinalityConfig(),
			Qualifier:             DefaultQualifier,
		},
	}
}
