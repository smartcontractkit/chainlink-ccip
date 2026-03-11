package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type EVMDeployChainContractsAdapter struct{}

var _ ccvadapters.DeployChainContractsAdapter = (*EVMDeployChainContractsAdapter)(nil)

var evmDeployChainContracts = cldf_ops.NewSequence(
	"evm-deploy-chain-contracts",
	semver.MustParse("2.0.0"),
	"Wraps the EVM DeployChainContracts sequence with chain-agnostic input conversion",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input ccvadapters.DeployChainContractsInput) (seq_core.OnChainOutput, error) {
		evmChains := chains.EVMChains()
		evmChain, ok := evmChains[input.ChainSelector]
		if !ok {
			return seq_core.OnChainOutput{}, fmt.Errorf("EVM chain not found for selector %d", input.ChainSelector)
		}

		evmInput, err := toEVMDeployInput(input)
		if err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to convert deploy input to EVM types: %w", err)
		}

		report, err := cldf_ops.ExecuteSequence(b, sequences.DeployChainContracts, evmChain, evmInput)
		if err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to execute EVM deploy chain contracts sequence: %w", err)
		}

		return report.Output, nil
	},
)

func (a *EVMDeployChainContractsAdapter) DeployChainContracts() *cldf_ops.Sequence[ccvadapters.DeployChainContractsInput, seq_core.OnChainOutput, chain.BlockChains] {
	return evmDeployChainContracts
}

func toEVMDeployInput(input ccvadapters.DeployChainContractsInput) (sequences.DeployChainContractsInput, error) {
	create2Factory, err := parseRequiredHexAddress(input.DeployerContract, "DeployerContract")
	if err != nil {
		return sequences.DeployChainContractsInput{}, err
	}

	committeeVerifiers, err := convertCommitteeVerifiers(input.ContractParams.CommitteeVerifiers)
	if err != nil {
		return sequences.DeployChainContractsInput{}, err
	}

	executors, err := convertExecutors(input.ContractParams.Executors)
	if err != nil {
		return sequences.DeployChainContractsInput{}, err
	}

	var legacyRMN common.Address
	if input.ContractParams.RMNRemote.LegacyRMN != "" {
		legacyRMN, err = parseHexAddress(input.ContractParams.RMNRemote.LegacyRMN, "RMNRemote.LegacyRMN")
		if err != nil {
			return sequences.DeployChainContractsInput{}, err
		}
	}

	var onRampFeeAgg common.Address
	if input.ContractParams.OnRamp.FeeAggregator != "" {
		onRampFeeAgg, err = parseHexAddress(input.ContractParams.OnRamp.FeeAggregator, "OnRamp.FeeAggregator")
		if err != nil {
			return sequences.DeployChainContractsInput{}, err
		}
	}

	mockReceivers := convertMockReceivers(input.ContractParams.MockReceivers)

	return sequences.DeployChainContractsInput{
		ChainSelector:     input.ChainSelector,
		CREATE2Factory:    create2Factory,
		ExistingAddresses: input.ExistingAddresses,
		DeployTestRouter:  input.DeployTestRouter,
		DeployerKeyOwned:  input.DeployerKeyOwned,
		ContractParams: sequences.ContractParams{
			RMNRemote: sequences.RMNRemoteParams{
				Version:   input.ContractParams.RMNRemote.Version,
				LegacyRMN: legacyRMN,
			},
			OffRamp: sequences.OffRampParams{
				Version:                   input.ContractParams.OffRamp.Version,
				GasForCallExactCheck:      input.ContractParams.OffRamp.GasForCallExactCheck,
				MaxGasBufferToUpdateState: input.ContractParams.OffRamp.MaxGasBufferToUpdateState,
			},
			CommitteeVerifiers: committeeVerifiers,
			OnRamp: sequences.OnRampParams{
				Version:               input.ContractParams.OnRamp.Version,
				FeeAggregator:         onRampFeeAgg,
				MaxUSDCentsPerMessage: input.ContractParams.OnRamp.MaxUSDCentsPerMessage,
			},
			FeeQuoter: sequences.FeeQuoterParams{
				Version:                        input.ContractParams.FeeQuoter.Version,
				MaxFeeJuelsPerMsg:              input.ContractParams.FeeQuoter.MaxFeeJuelsPerMsg,
				LINKPremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoter.LINKPremiumMultiplierWeiPerEth,
				WETHPremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoter.WETHPremiumMultiplierWeiPerEth,
				USDPerLINK:                     input.ContractParams.FeeQuoter.USDPerLINK,
				USDPerWETH:                     input.ContractParams.FeeQuoter.USDPerWETH,
			},
			Executors:     executors,
			MockReceivers: mockReceivers,
		},
	}, nil
}

func convertCommitteeVerifiers(params []ccvadapters.CommitteeVerifierDeployParams) ([]sequences.CommitteeVerifierParams, error) {
	result := make([]sequences.CommitteeVerifierParams, 0, len(params))
	for _, cv := range params {
		feeAgg, err := parseRequiredNonZeroHexAddress(cv.FeeAggregator, fmt.Sprintf("committee %q FeeAggregator", cv.Qualifier))
		if err != nil {
			return nil, err
		}

		var allowlistAdmin common.Address
		if cv.AllowlistAdmin != "" {
			allowlistAdmin, err = parseHexAddress(cv.AllowlistAdmin, fmt.Sprintf("committee %q AllowlistAdmin", cv.Qualifier))
			if err != nil {
				return nil, err
			}
		}

		result = append(result, sequences.CommitteeVerifierParams{
			Version:          cv.Version,
			FeeAggregator:    feeAgg,
			AllowlistAdmin:   allowlistAdmin,
			StorageLocations: cv.StorageLocations,
			Qualifier:        cv.Qualifier,
		})
	}
	return result, nil
}

func convertExecutors(params []ccvadapters.ExecutorDeployParams) ([]sequences.ExecutorParams, error) {
	result := make([]sequences.ExecutorParams, 0, len(params))
	for _, ep := range params {
		var feeAgg common.Address
		var err error
		if ep.DynamicConfig.FeeAggregator != "" {
			feeAgg, err = parseHexAddress(ep.DynamicConfig.FeeAggregator, fmt.Sprintf("executor %q DynamicConfig.FeeAggregator", ep.Qualifier))
			if err != nil {
				return nil, err
			}
		}

		result = append(result, sequences.ExecutorParams{
			Version:       ep.Version,
			MaxCCVsPerMsg: ep.MaxCCVsPerMsg,
			DynamicConfig: executor.SetDynamicConfigArgs{
				FeeAggregator:         feeAgg,
				MinBlockConfirmations: ep.DynamicConfig.MinBlockConfirmations,
				CcvAllowlistEnabled:   ep.DynamicConfig.CcvAllowlistEnabled,
			},
			Qualifier: ep.Qualifier,
		})
	}
	return result, nil
}

func convertMockReceivers(params []ccvadapters.MockReceiverDeployParams) []sequences.MockReceiverParams {
	result := make([]sequences.MockReceiverParams, 0, len(params))
	for _, mr := range params {
		result = append(result, sequences.MockReceiverParams{
			Version:           mr.Version,
			RequiredVerifiers: mr.RequiredVerifiers,
			OptionalVerifiers: mr.OptionalVerifiers,
			OptionalThreshold: mr.OptionalThreshold,
			Qualifier:         mr.Qualifier,
		})
	}
	return result
}

func parseHexAddress(hex, field string) (common.Address, error) {
	if !common.IsHexAddress(hex) {
		return common.Address{}, fmt.Errorf("%s: %q is not a valid hex address", field, hex)
	}
	return common.HexToAddress(hex), nil
}

func parseRequiredHexAddress(hex, field string) (common.Address, error) {
	if hex == "" {
		return common.Address{}, fmt.Errorf("%s is required", field)
	}
	return parseHexAddress(hex, field)
}

func parseRequiredNonZeroHexAddress(hex, field string) (common.Address, error) {
	addr, err := parseRequiredHexAddress(hex, field)
	if err != nil {
		return common.Address{}, err
	}
	if addr == (common.Address{}) {
		return common.Address{}, fmt.Errorf("%s cannot be zero address", field)
	}
	return addr, nil
}
