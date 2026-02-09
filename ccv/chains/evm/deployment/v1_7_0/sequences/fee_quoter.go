package sequences

import (
	"encoding/json"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	seq1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	fq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
)

const (
	LinkFeeMultiplierPercent uint8 = 90
)

type FeeQuoterUpdate struct {
	ChainSelector                 uint64
	ExistingAddresses             []datastore.AddressRef
	ConstructorArgs               fqops.ConstructorArgs
	PriceUpdates                  fqops.PriceUpdates
	DestChainConfigs              []fqops.DestChainConfigArgs
	TokenTransferFeeConfigUpdates fqops.ApplyTokenTransferFeeConfigUpdatesArgs
	AuthorizedCallerUpdates       fqops.AuthorizedCallerArgs
}

func (fqu FeeQuoterUpdate) IsEmpty() (bool, error) {
	empty := FeeQuoterUpdate{}
	// marshal into json
	emptyBytes, err := json.Marshal(empty)
	if err != nil {
		return false, fmt.Errorf("failed to marshal empty FeeQuoterUpdate: %w", err)
	}
	inputBytes, err := json.Marshal(fqu)
	if err != nil {
		return false, fmt.Errorf("failed to marshal FeeQuoterUpdate: %w", err)
	}
	return string(emptyBytes) == string(inputBytes), nil
}

var (
	// SequenceFeeQuoterUpdate is a sequence that deploys or fetches existing FeeQuoter contract
	// and does the following if the corresponding input is provided -
	// 1. applies destination chain config updates
	// 2. price updates
	// 3. token transfer fee config updates
	// 4. authorized caller updates
	SequenceFeeQuoterUpdate = cldf_ops.NewSequence(
		"fee-quoter-v1.7.0:update-sequence",
		semver.MustParse("1.7.0"),
		"Deploys or fetches existing FeeQuoter contract and applies destination chain config updates and price updates",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input FeeQuoterUpdate) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
			}

			// deploy fee quoter or fetch existing fee quoter address
			feeQuoterRef, err := contract.MaybeDeployContract(
				b, fqops.Deploy, chain, contract.DeployInput[fqops.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(fqops.ContractType, *fqops.Version),
					ChainSelector:  chain.Selector,
					Args:           input.ConstructorArgs,
				}, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			if feeQuoterRef.Address == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy or "+
					"fetch FeeQuoter on chain %s", chain.String())
			}
			writes := make([]contract.WriteOutput, 0)
			output.Addresses = append(output.Addresses, feeQuoterRef)
			fqAddr := common.HexToAddress(feeQuoterRef.Address)
			// ApplyDestChainConfigUpdates on FeeQuoter
			if len(input.DestChainConfigs) > 0 {
				feeQuoterReport, err := cldf_ops.ExecuteOperation(
					b, fqops.ApplyDestChainConfigUpdates, chain,
					contract.FunctionInput[[]fqops.DestChainConfigArgs]{
						ChainSelector: chain.Selector,
						Address:       fqAddr,
						Args:          input.DestChainConfigs,
					})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain "+
						"config updates to FeeQuoter(%s) on chain %s: %w", fqAddr.Hex(), chain, err)
				}
				writes = append(writes, feeQuoterReport.Output)
			}
			// update price
			if len(input.PriceUpdates.GasPriceUpdates) > 0 || len(input.PriceUpdates.TokenPriceUpdates) > 0 {
				feeQuoterUpdatePricesReport, err := cldf_ops.ExecuteOperation(
					b, fqops.UpdatePrices, chain, contract.FunctionInput[fqops.PriceUpdates]{
						ChainSelector: chain.Selector,
						Address:       fqAddr,
						Args: fqops.PriceUpdates{
							GasPriceUpdates:   input.PriceUpdates.GasPriceUpdates,
							TokenPriceUpdates: input.PriceUpdates.TokenPriceUpdates,
						},
					})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to update gas prices on "+
						"FeeQuoter(%s) on chain %s: %w", fqAddr.Hex(), chain, err)
				}
				writes = append(writes, feeQuoterUpdatePricesReport.Output)
			}
			// TokenTransferFeeConfigUpdates on FeeQuoter
			if len(input.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs) > 0 ||
				len(input.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs) > 0 {
				feeQuoterTokenTransferFeeConfigReport, err := cldf_ops.ExecuteOperation(
					b, fqops.ApplyTokenTransferFeeConfigUpdates, chain,
					contract.FunctionInput[fqops.ApplyTokenTransferFeeConfigUpdatesArgs]{
						ChainSelector: chain.Selector,
						Address:       fqAddr,
						Args: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
							TokenTransferFeeConfigArgs:   input.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs,
							TokensToUseDefaultFeeConfigs: input.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs,
						},
					})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee "+
						"config updates to FeeQuoter(%s) on chain %s: %w", fqAddr.Hex(), chain, err)
				}
				writes = append(writes, feeQuoterTokenTransferFeeConfigReport.Output)
			}
			// ApplyAuthorizedCallerUpdates on FeeQuoter
			if len(input.AuthorizedCallerUpdates.AddedCallers) > 0 ||
				len(input.AuthorizedCallerUpdates.RemovedCallers) > 0 {
				feeQuoterAuthorizedCallerReport, err := cldf_ops.ExecuteOperation(
					b, fqops.ApplyAuthorizedCallerUpdates, chain,
					contract.FunctionInput[fqops.AuthorizedCallerArgs]{
						ChainSelector: chain.Selector,
						Address:       fqAddr,
						Args:          input.AuthorizedCallerUpdates,
					})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller "+
						"updates to FeeQuoter(%s) on chain %s: %w", fqAddr.Hex(), chain, err)
				}
				writes = append(writes, feeQuoterAuthorizedCallerReport.Output)
			}
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			output.BatchOps = []mcms_types.BatchOperation{batch}
			return output, nil
		},
	)

	// CreateFeeQuoterUpdateInputFromV163 creates FeeQuoterUpdate input by importing configuration from FeeQuoter v1.6.0
	CreateFeeQuoterUpdateInputFromV163 = cldf_ops.NewSequence(
		"fetches-feequoter-config-values-from-v1.6.3",
		semver.MustParse("1.7.0"),
		"Creates FeeQuoterUpdate input by importing configuration from FeeQuoter v1.6.3",
		func(b cldf_ops.Bundle, chain evm.Chain, input deploy.FeeQuoterUpdateInput) (output FeeQuoterUpdate, err error) {
			// check if FeeQuoter v1.6.3 is present in existing addresses, if not, we return empty output
			// it means there is no existing fee quoter deployed from v1.6.3 deployment, and we can skip the config import from v1.6.3
			fq16Ref := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fq1_6.ContractType,
				fq1_6.Version,
				"",
			)
			if datastore_utils.IsAddressRefEmpty(fq16Ref) {
				return FeeQuoterUpdate{}, nil
			}
			output.ChainSelector = input.ChainSelector
			output.ExistingAddresses = input.ExistingAddresses
			// get feeQuoter 1.6 address meta
			metadataForFq16, err := datastore_utils.FilterContractMetaByContractTypeAndVersion(
				input.ExistingAddresses,
				input.ContractMeta,
				fq1_6.ContractType,
				fq1_6.Version,
				"",
				input.ChainSelector,
			)
			if err != nil {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to get FeeQuoter 1.6.3 address: %w", err)
			}
			if len(metadataForFq16) == 0 {
				return FeeQuoterUpdate{}, fmt.Errorf("no metadata found for FeeQuoter v1.6.3 on chain selector %d", input.ChainSelector)
			}
			if len(metadataForFq16) > 1 {
				return FeeQuoterUpdate{}, fmt.Errorf("multiple metadata entries found for FeeQuoter v1.6.3 on chain selector %d", input.ChainSelector)
			}
			// Convert metadata to typed struct if needed
			fqOutput, err := datastore_utils.ConvertMetadataToType[seq1_6.FeeQuoterImportConfigSequenceOutput](metadataForFq16[0].Metadata)
			if err != nil {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to convert metadata to "+
					"FeeQuoterImportConfigSequenceOutput for chain selector %d: %w", input.ChainSelector, err)
			}
			// is feeQuoter going to be deployed or fetched from existing addresses?
			feeQuoterRef := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fqops.ContractType,
				fqops.Version,
				"",
			)
			isNewFQ17Deployment := datastore_utils.IsAddressRefEmpty(feeQuoterRef)
			tokenTransferFeeConfigArgs := make([]fee_quoter.FeeQuoterTokenTransferFeeConfigArgs, 0)
			allDestChainConfigs := make([]fqops.DestChainConfigArgs, 0)
			for remoteChain, cfg := range fqOutput.RemoteChainCfgs {
				if !cfg.DestChainCfg.IsEnabled {
					continue
				}
				destChainConfig := cfg.DestChainCfg
				outDestchainCfg := fqops.DestChainConfigArgs{
					DestChainSelector: remoteChain,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   destChainConfig.IsEnabled,
						MaxDataBytes:                destChainConfig.MaxDataBytes,
						MaxPerMsgGasLimit:           destChainConfig.MaxPerMsgGasLimit,
						DestGasOverhead:             destChainConfig.DestGasOverhead,
						DestGasPerPayloadByteBase:   destChainConfig.DestGasPerPayloadByteBase,
						ChainFamilySelector:         destChainConfig.ChainFamilySelector,
						DefaultTokenFeeUSDCents:     destChainConfig.DefaultTokenFeeUSDCents,
						DefaultTokenDestGasOverhead: destChainConfig.DefaultTokenDestGasOverhead,
						DefaultTxGasLimit:           destChainConfig.DefaultTxGasLimit,
						NetworkFeeUSDCents:          uint16(destChainConfig.NetworkFeeUSDCents),
						LinkFeeMultiplierPercent:    LinkFeeMultiplierPercent,
					},
				}
				tokenTransferFeeCfgs := make([]fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs, 0)
				for token, transferCfg := range cfg.TokenTransferFeeCfgs {
					if !transferCfg.IsEnabled {
						continue
					}
					tokenTransferFeeCfgs = append(tokenTransferFeeCfgs, fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						Token: token,
						TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
							FeeUSDCents:       transferCfg.MinFeeUSDCents,
							DestGasOverhead:   transferCfg.DestGasOverhead,
							DestBytesOverhead: transferCfg.DestBytesOverhead,
							IsEnabled:         transferCfg.IsEnabled,
						},
					})
				}
				tokenTransferFeeConfigArgs = append(tokenTransferFeeConfigArgs, fee_quoter.FeeQuoterTokenTransferFeeConfigArgs{
					DestChainSelector:       remoteChain,
					TokenTransferFeeConfigs: tokenTransferFeeCfgs,
				})
				allDestChainConfigs = append(allDestChainConfigs, outDestchainCfg)
			}
			if isNewFQ17Deployment {
				output.ConstructorArgs = fqops.ConstructorArgs{
					StaticConfig: fqops.StaticConfig{
						LinkToken:         fqOutput.StaticCfg.LinkToken,
						MaxFeeJuelsPerMsg: fqOutput.StaticCfg.MaxFeeJuelsPerMsg,
					},
					PriceUpdaters:              fqOutput.PriceUpdaters,
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgs,
					DestChainConfigArgs:        allDestChainConfigs,
				}
			} else {
				output.AuthorizedCallerUpdates = fqops.AuthorizedCallerArgs{
					AddedCallers: fqOutput.PriceUpdaters,
				}
				output.TokenTransferFeeConfigUpdates = fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgs,
				}
				output.DestChainConfigs = allDestChainConfigs
			}
			return output, nil
		})

	// CreateFeeQuoterUpdateInputFromV150 creates FeeQuoterUpdate input by importing configuration from PriceRegistry v1.5.0 and EVM2EVMOnRamp v1.5.0
	CreateFeeQuoterUpdateInputFromV150 = cldf_ops.NewSequence(
		"fetches-feequoter-config-values-from-v1.5.0",
		semver.MustParse("1.7.0"),
		"Creates FeeQuoterUpdate input by importing configuration from PriceRegistry v1.5.0 and EVM2EVMOnRamp v1.5.0",
		func(b cldf_ops.Bundle, chain evm.Chain, input deploy.FeeQuoterUpdateInput) (output FeeQuoterUpdate, err error) {
			// get addressref for onramp 1.5.0
			onRampRef := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				onrampops.ContractType,
				onrampops.Version,
				"",
			)
			// if there is no address ref for onRamp 1.5.0, it means onRamp 1.5.0 is not deployed and we can skip the config import from onRamp 1.5.0
			if datastore_utils.IsAddressRefEmpty(onRampRef) {
				return FeeQuoterUpdate{}, nil
			}
			// get address meta for onRamp 1.5.0 to read the config values from onRamp 1.5.0
			onRampMetadata, err := datastore_utils.FilterContractMetaByContractTypeAndVersion(
				input.ExistingAddresses,
				input.ContractMeta,
				onrampops.ContractType,
				onrampops.Version,
				"",
				input.ChainSelector,
			)
			if err != nil {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to get EVM2EVMOnRamp v1.5.0 address: %w", err)
			}
			if len(onRampMetadata) == 0 {
				return FeeQuoterUpdate{}, fmt.Errorf("no metadata found for EVM2EVMOnRamp v1.5.0 on chain selector %d", input.ChainSelector)
			}
			// get the commit stores and that will act like price updaters for fee quoter
			var commitStoreRefs []datastore.AddressRef
			for _, addressRef := range input.ExistingAddresses {
				if addressRef.Type == "CommitStore" &&
					addressRef.Version == semver.MustParse("1.5.0") &&
					addressRef.ChainSelector == input.ChainSelector {
					commitStoreRefs = append(commitStoreRefs, addressRef)
				}
			}

			if len(commitStoreRefs) == 0 {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to get commit store ref for chain %d", input.ChainSelector)
			}
			var priceUpdaters []common.Address
			for _, ref := range commitStoreRefs {
				priceUpdaters = append(priceUpdaters, common.HexToAddress(ref.Address))
			}
			output.ChainSelector = input.ChainSelector
			output.ExistingAddresses = input.ExistingAddresses
			// is feeQuoter going to be deployed or fetched from existing addresses?
			feeQuoter17Ref := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fqops.ContractType,
				fqops.Version,
				"",
			)
			isNewFQ17Deployment := datastore_utils.IsAddressRefEmpty(feeQuoter17Ref)
			var staticCfg fqops.StaticConfig
			var destChainCfgs []fqops.DestChainConfigArgs
			var tokenTransferFeeConfigArgs []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs
			var tokenTransferFeeConfigArgsForAll []fqops.TokenTransferFeeConfigArgs
			for _, meta := range onRampMetadata {
				// Convert metadata to typed struct if needed
				onRampCfg, err := datastore_utils.ConvertMetadataToType[seq1_5.OnRampImportConfigSequenceOutput](meta.Metadata)
				if err != nil {
					return FeeQuoterUpdate{}, fmt.Errorf("failed to convert metadata to "+
						"OnRampImportConfigSequenceOutput for chain selector %d: %w", input.ChainSelector, err)
				}
				if staticCfg.LinkToken == (common.Address{}) {
					staticCfg = fqops.StaticConfig{
						LinkToken:         onRampCfg.StaticConfig.LinkToken,
						MaxFeeJuelsPerMsg: onRampCfg.StaticConfig.MaxNopFeesJuels,
					}
				}
				var networkFeeUSDCents uint16
				// NetworkFeeUSDCents is same across all feetokens in the same chain, so we can just take it from the first onRamp config
				for _, feeTokenCfg := range onRampCfg.FeeTokenConfig {
					if feeTokenCfg.NetworkFeeUSDCents != 0 {
						networkFeeUSDCents = uint16(feeTokenCfg.NetworkFeeUSDCents)
						break
					}
				}
				chainFamilySelectorBytes := utils.GetSelectorHex(onRampCfg.RemoteChainSelector)
				// Safely convert ChainFamilySelector from []byte to [4]byte
				var chainFamilySelector [4]byte
				if len(chainFamilySelectorBytes) < 4 {
					return FeeQuoterUpdate{}, fmt.Errorf("ChainFamilySelector has invalid length %d (expected 4) for remote chain selector %d", len(chainFamilySelectorBytes), onRampCfg.RemoteChainSelector)
				}
				copy(chainFamilySelector[:], chainFamilySelectorBytes[:4])
				destChainCfgs = append(destChainCfgs, fqops.DestChainConfigArgs{
					DestChainSelector: onRampCfg.RemoteChainSelector,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true, // if the chain is supported on OnRamp, we should enable it on FeeQuoter
						MaxDataBytes:                onRampCfg.DynamicConfig.MaxDataBytes,
						MaxPerMsgGasLimit:           onRampCfg.DynamicConfig.MaxPerMsgGasLimit,
						DestGasOverhead:             onRampCfg.DynamicConfig.DestGasOverhead,
						DestGasPerPayloadByteBase:   uint8(onRampCfg.DynamicConfig.DestGasPerPayloadByte),
						ChainFamilySelector:         chainFamilySelector,
						DefaultTokenFeeUSDCents:     onRampCfg.DynamicConfig.DefaultTokenFeeUSDCents,
						DefaultTokenDestGasOverhead: onRampCfg.DynamicConfig.DefaultTokenDestGasOverhead,
						DefaultTxGasLimit:           uint32(onRampCfg.StaticConfig.DefaultTxGasLimit),
						NetworkFeeUSDCents:          networkFeeUSDCents,
						LinkFeeMultiplierPercent:    LinkFeeMultiplierPercent,
					},
				})
				for token, tokenCfg := range onRampCfg.TokenTransferFeeConfig {
					tokenTransferFeeConfigArgs = append(tokenTransferFeeConfigArgs, fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						Token: token,
						TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
							FeeUSDCents:       tokenCfg.MinFeeUSDCents,
							DestGasOverhead:   tokenCfg.DestGasOverhead,
							DestBytesOverhead: tokenCfg.DestBytesOverhead,
							IsEnabled:         tokenCfg.IsEnabled,
						},
					})
				}
				tokenTransferFeeConfigArgsForAll = append(tokenTransferFeeConfigArgsForAll, fqops.TokenTransferFeeConfigArgs{
					DestChainSelector:       onRampCfg.RemoteChainSelector,
					TokenTransferFeeConfigs: tokenTransferFeeConfigArgs,
				})
			}
			if isNewFQ17Deployment {
				output.ConstructorArgs = fqops.ConstructorArgs{
					StaticConfig: fqops.StaticConfig{
						LinkToken:         staticCfg.LinkToken,
						MaxFeeJuelsPerMsg: staticCfg.MaxFeeJuelsPerMsg,
					},
					DestChainConfigArgs:        destChainCfgs,
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgsForAll,
					PriceUpdaters:              priceUpdaters,
				}
			} else {
				output.DestChainConfigs = destChainCfgs
				output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs = tokenTransferFeeConfigArgsForAll
				output.AuthorizedCallerUpdates = fqops.AuthorizedCallerArgs{
					AddedCallers: priceUpdaters,
				}
			}
			return output, nil
		})
)

// MergeFeeQuoterUpdateOutputs merges FeeQuoterUpdate outputs from the v1.6.3 and v1.5.0 import
// sequences into a single update. output16 is the base; output15 supplements it. Where both
// provide values (e.g. ConstructorArgs, dest chain configs, token transfer fee configs),
// output16 takes precedence and output15 fills in only missing entries.
func MergeFeeQuoterUpdateOutputs(output16, output15 FeeQuoterUpdate) (FeeQuoterUpdate, error) {
	result := output16

	// ConstructorArgs: use output15 if output16 is empty
	if result.ConstructorArgs.IsEmpty() {
		result.ConstructorArgs = output15.ConstructorArgs
	} else {
		// merge the dest chainConfig args
		result.ConstructorArgs.DestChainConfigArgs = mergeDestChainConfigs(
			result.ConstructorArgs.DestChainConfigArgs,
			output15.ConstructorArgs.DestChainConfigArgs)
		resultPriceUpdatersMap := make(map[common.Address]bool)
		for _, updater := range result.ConstructorArgs.PriceUpdaters {
			resultPriceUpdatersMap[updater] = true
		}
		for _, updater := range output15.ConstructorArgs.PriceUpdaters {
			if !resultPriceUpdatersMap[updater] {
				result.ConstructorArgs.PriceUpdaters = append(result.ConstructorArgs.PriceUpdaters, updater)
				resultPriceUpdatersMap[updater] = true
			}
		}
		result.ConstructorArgs.TokenTransferFeeConfigArgs = mergeTokenTransferFeeConfigArgs(
			result.ConstructorArgs.TokenTransferFeeConfigArgs,
			output15.ConstructorArgs.TokenTransferFeeConfigArgs)
		result.ConstructorArgs.PriceUpdaters = maps.Keys(resultPriceUpdatersMap)
	}

	result.DestChainConfigs = mergeDestChainConfigs(result.DestChainConfigs, output15.DestChainConfigs)

	// TokenTransferFeeConfigUpdates: merge by DestChainSelector, output16 takes precedence for duplicates
	result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs = mergeTokenTransferFeeConfigArgs(
		result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs,
		output15.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs)

	// TokensToUseDefaultFeeConfigs: merge by DestChainSelector and Token
	if len(result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs) == 0 {
		result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs = output15.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs
	} else {
		// Create a map of (DestChainSelector, Token) pairs from output16
		tokenRemoveMap := make(map[string]bool)
		for _, cfg := range result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs {
			key := fmt.Sprintf("%d:%s", cfg.DestChainSelector, cfg.Token.Hex())
			tokenRemoveMap[key] = true
		}
		// Add configs from output15 that don't exist in output16
		for _, cfg := range output15.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs {
			key := fmt.Sprintf("%d:%s", cfg.DestChainSelector, cfg.Token.Hex())
			if !tokenRemoveMap[key] {
				result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs = append(result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs, cfg)
			}
			// If it exists in both, output16's value is already used (takes precedence)
		}
	}

	// AuthorizedCallerUpdates: merge unique entries from both outputs
	result.AuthorizedCallerUpdates = mergePriceUpdaters(result.AuthorizedCallerUpdates, output15.AuthorizedCallerUpdates)

	return result, nil
}

func mergeTokenTransferFeeConfigArgs(args1, args2 []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs) []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs {
	result := args1
	// TokenTransferFeeConfigArgs: merge by DestChainSelector
	if len(result) == 0 {
		result = args2
	} else {
		// Create a map of dest chain selectors from output16
		tokenConfigMap := make(map[uint64]int)
		for i, cfg := range result {
			tokenConfigMap[cfg.DestChainSelector] = i
		}
		// Add configs from output15 that don't exist in output16
		for _, cfg := range args2 {
			if _, exists := tokenConfigMap[cfg.DestChainSelector]; !exists && len(cfg.TokenTransferFeeConfigs) > 0 {
				result = append(result, cfg)
			}
			// If it exists in both, output16's value is already used (takes precedence)
		}
	}
	return result
}

func mergePriceUpdaters(updaters1, updaters2 fqops.AuthorizedCallerArgs) fqops.AuthorizedCallerArgs {
	result := updaters1
	// AddedCallers: merge unique addresses from both outputs
	addedCallersMap := make(map[common.Address]bool)
	for _, addr := range result.AddedCallers {
		addedCallersMap[addr] = true
	}
	for _, addr := range updaters2.AddedCallers {
		if !addedCallersMap[addr] {
			result.AddedCallers = append(result.AddedCallers, addr)
			addedCallersMap[addr] = true
		}
	}
	// RemovedCallers: merge unique addresses from both outputs
	removedCallersMap := make(map[common.Address]bool)
	for _, addr := range result.RemovedCallers {
		removedCallersMap[addr] = true
	}
	for _, addr := range updaters2.RemovedCallers {
		if !removedCallersMap[addr] {
			result.RemovedCallers = append(result.RemovedCallers, addr)
			removedCallersMap[addr] = true
		}
	}
	return result
}

func mergeDestChainConfigs(cfgs1, cfgs2 []fqops.DestChainConfigArgs) []fqops.DestChainConfigArgs {
	// Create a map of dest chain selectors from cfgs1
	destChainMap := make(map[uint64]fqops.DestChainConfigArgs)
	for _, cfg := range cfgs1 {
		destChainMap[cfg.DestChainSelector] = cfg
	}
	result := cfgs1
	// Add configs from cfgs2 that don't exist in cfgs1
	for _, cfg := range cfgs2 {
		if _, exists := destChainMap[cfg.DestChainSelector]; !exists {
			result = append(result, cfg)
		}
		// If it exists in both, cfgs1's value is already used (takes precedence)
	}
	if len(destChainMap) == 0 {
		return nil
	}
	return result
}
