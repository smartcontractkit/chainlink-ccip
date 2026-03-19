package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	offrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

type EVMDeployChainContractsAdapter struct{}

var _ ccvadapters.DeployChainContractsAdapter = (*EVMDeployChainContractsAdapter)(nil)

var (
	evmDeployChainContracts = cldf_ops.NewSequence(
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
		})

	importConfigForDeployContracts = cldf_ops.NewSequence(
		"evm-import-config-for-deploy-chain-contracts",
		semver.MustParse("2.0.0"),
		"Reads contract parameters from the datastore based on the chain selector and returns them in the format needed for DeployChainContracts",
		func(b cldf_ops.Bundle, chains chain.BlockChains, input ccvadapters.DeployChainConfigCreatorInput) (output ccvadapters.DeployContractParams, err error) {
			evmChains := chains.EVMChains()
			evmChain, ok := evmChains[input.ChainSelector]
			if !ok {
				return ccvadapters.DeployContractParams{}, fmt.Errorf("EVM chain not found for selector %d", input.ChainSelector)
			}
			paramsFrom1_6_0, err := importConfigFromv1_6_0(b, evmChain, input)
			if err != nil {
				return output, fmt.Errorf("failed to import config from v1.6.0: %w", err)
			}
			paramsFrom1_5_0, err := importConfigFromv1_5_0(input)
			if err != nil {
				return output, fmt.Errorf("failed to import config from v1.5.0: %w", err)
			}
			output, err = paramsFrom1_5_0.MergeWithOverrideIfNotEmpty(paramsFrom1_6_0)
			return
		})
)

func (a *EVMDeployChainContractsAdapter) SetContractParamsFromImportedConfig() *cldf_ops.Sequence[ccvadapters.DeployChainConfigCreatorInput, ccvadapters.DeployContractParams, chain.BlockChains] {
	return importConfigForDeployContracts
}

func (a *EVMDeployChainContractsAdapter) DeployChainContracts() *cldf_ops.Sequence[ccvadapters.DeployChainContractsInput, seq_core.OnChainOutput, chain.BlockChains] {
	return evmDeployChainContracts
}

func importConfigFromv1_5_0(input ccvadapters.DeployChainConfigCreatorInput) (ccvadapters.DeployContractParams, error) {
	output := input.UserProvidedConfig
	// find legacy RMN address
	rmnAddr := datastore_utils.GetAddressRef(
		input.ExistingAddresses,
		input.ChainSelector,
		rmnops1_5.ContractType,
		semver.MustParse("1.5.0"),
		"",
	)
	if rmnAddr.Address != "" {
		output.RMNRemote.LegacyRMN = rmnAddr.Address
	}
	// set default value for gas for call exact check based on v1.5.0 deployments
	//https://github.com/smartcontractkit/ccip/blob/b5529a39311a2fd39cafceb62e4bb8f40eeb2e9e/contracts/src/v0.8/ccip/libraries/Internal.sol#L14C55-L14C60
	output.OffRamp.GasForCallExactCheck = 5000
	return output, nil
}

func importConfigFromv1_6_0(b cldf_ops.Bundle, chain evm.Chain, input ccvadapters.DeployChainConfigCreatorInput) (ccvadapters.DeployContractParams, error) {
	output := input.UserProvidedConfig
	// fetch onRamp 1.6.0 , if not found, it's possible onRamp wasn't deployed in v1.6.0 for this chain, so we just return
	onRampAddr := datastore_utils.GetAddressRef(
		input.ExistingAddresses,
		input.ChainSelector,
		onrampops_v160.ContractType,
		onrampops_v160.Version,
		"",
	)
	if onRampAddr.Address == "" {
		return output, nil
	}
	// if onRamp address is found, we assume the other onRamp metadata fields are also present
	metadataForonRamp16, err := datastore_utils.FilterContractMetaByContractTypeAndVersion(
		input.ExistingAddresses,
		input.ContractMeta,
		onrampops_v160.ContractType,
		onrampops_v160.Version,
		"",
		input.ChainSelector,
	)
	if err != nil {
		return output, fmt.Errorf("failed to get onRamp metadata for chain selector %d: %w", input.ChainSelector, err)
	}
	if len(metadataForonRamp16) == 0 {
		return output, fmt.Errorf("no metadata found for onRamp v1.6.0 on chain selector %d", input.ChainSelector)
	}
	if len(metadataForonRamp16) > 1 {
		return output, fmt.Errorf("multiple metadata entries found for onRamp v1.6.0 on chain selector %d", input.ChainSelector)
	}
	metadataForoffRamp16, err := datastore_utils.FilterContractMetaByContractTypeAndVersion(
		input.ExistingAddresses,
		input.ContractMeta,
		offrampops_v160.ContractType,
		offrampops_v160.Version,
		"",
		input.ChainSelector,
	)
	if err != nil {
		return output, fmt.Errorf("failed to get offRamp metadata for chain selector %d: %w", input.ChainSelector, err)
	}
	if len(metadataForoffRamp16) == 0 {
		return output, fmt.Errorf("no metadata found for offRamp v1.6.0 on chain selector %d", input.ChainSelector)
	}
	if len(metadataForoffRamp16) > 1 {
		return output, fmt.Errorf("multiple metadata entries found for offRamp v1.6.0 on chain selector %d", input.ChainSelector)
	}
	// Convert metadata to typed struct if needed
	onRampCfg16, err := datastore_utils.ConvertMetadataToType[seq1_6.OnRampImportConfigSequenceOutput](metadataForonRamp16[0].Metadata)
	if err != nil {
		return output, fmt.Errorf("failed to convert metadata to "+
			"OnRampImportConfigSequenceOutput for chain selector %d: %w", input.ChainSelector, err)
	}
	offRampCfg16, err := datastore_utils.ConvertMetadataToType[seq1_6.OffRampImportConfigSequenceOutput](metadataForoffRamp16[0].Metadata)
	if err != nil {
		return output, fmt.Errorf("failed to convert metadata to "+
			"OffRampImportConfigSequenceOutput for chain selector %d: %w", input.ChainSelector, err)
	}

	feeAggr := onRampCfg16.DynamicConfig.FeeAggregator.String()
	output.OnRamp.FeeAggregator = feeAggr

	for i := range output.Executors {
		output.Executors[i].DynamicConfig.FeeAggregator = feeAggr
	}
	for i := range output.CommitteeVerifiers {
		output.CommitteeVerifiers[i].FeeAggregator = feeAggr
	}
	output.OffRamp.GasForCallExactCheck = offRampCfg16.StaticConfig.GasForCallExactCheck
	return output, nil
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
			DynamicConfig: executor.DynamicConfig{
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
			Version:                   mr.Version,
			RequiredVerifiers:         mr.RequiredVerifiers,
			OptionalVerifiers:         mr.OptionalVerifiers,
			OptionalThreshold:         mr.OptionalThreshold,
			MinimumBlockConfirmations: mr.MinimumBlockConfirmations,
			Qualifier:                 mr.Qualifier,
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
