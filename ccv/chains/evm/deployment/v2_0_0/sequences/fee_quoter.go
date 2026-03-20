package sequences

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	seq1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	fq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
)

const (
	LinkFeeMultiplierPercent uint8  = 90
	NetworkFeeUSDCents       uint16 = 10
)

var (
	staticGasPriceByChainFamily = map[string]*big.Int{
		chain_selectors.FamilyAptos: big.NewInt(15e11),
		chain_selectors.FamilySui:   big.NewInt(15e11),
	}
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
	// marshal into JSON
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
		"fee-quoter-v2.0.0:update-sequence",
		semver.MustParse("2.0.0"),
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

	// CreateFeeQuoterUpdateInputFromV16x creates FeeQuoterUpdate input by importing configuration from FeeQuoter v1.6.x
	CreateFeeQuoterUpdateInputFromV16x = cldf_ops.NewSequence(
		"fetches-feequoter-config-values-from-v1.6.x",
		semver.MustParse("2.0.0"),
		"Creates FeeQuoterUpdate input by importing configuration from FeeQuoter v1.6.x",
		func(b cldf_ops.Bundle, chain evm.Chain, input deploy.FeeQuoterUpdateInput) (output FeeQuoterUpdate, err error) {
			// check if FeeQuoter v1.6.x is present in existing addresses, if not, we return empty output
			// it means there is no existing fee quoter deployed from v1.6.x deployment, and we can skip the config import from v1.6.x
			fq16AddressRef, err := seq1_6.GetFeeQuoterAddress(input.ExistingAddresses, input.ChainSelector, fqops.Version)
			if err != nil {
				if strings.Contains(err.Error(), "no fee quoter address found") {
					return FeeQuoterUpdate{}, nil
				}
				return FeeQuoterUpdate{}, fmt.Errorf("failed to get FeeQuoter 1.6.x address: %w", err)
			}
			output.ChainSelector = input.ChainSelector
			output.ExistingAddresses = input.ExistingAddresses

			// get feeQuoter 1.6 address meta
			metadataForFq16, err := datastore_utils.FilterContractMetaByContractTypeAndVersion(
				input.ExistingAddresses,
				input.ContractMeta,
				fq1_6.ContractType,
				fq16AddressRef.Version,
				"",
				input.ChainSelector,
			)
			if err != nil {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to get FeeQuoter 1.6.x address: %w", err)
			}
			if len(metadataForFq16) == 0 {
				return FeeQuoterUpdate{}, fmt.Errorf("no metadata found for FeeQuoter v1.6.x on chain selector %d", input.ChainSelector)
			}
			if len(metadataForFq16) > 1 {
				return FeeQuoterUpdate{}, fmt.Errorf("multiple metadata entries found for FeeQuoter v1.6.x on chain selector %d", input.ChainSelector)
			}
			// Convert metadata to typed struct if needed
			fqOutput, err := datastore_utils.ConvertMetadataToType[seq1_6.FeeQuoterImportConfigSequenceOutput](metadataForFq16[0].Metadata)
			if err != nil {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to convert metadata to "+
					"FeeQuoterImportConfigSequenceOutput for chain selector %d: %w", input.ChainSelector, err)
			}
			routerAddr := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				routerops.ContractType,
				routerops.Version,
				"",
			)
			if routerAddr.Address == "" {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to find router address ref for chain selector %d", input.ChainSelector)
			}

			// is feeQuoter going to be deployed or fetched from existing addresses?
			feeQuoterRef := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fqops.ContractType,
				fqops.Version,
				"",
			)
			isNewFQV2Deployment := datastore_utils.IsAddressRefEmpty(feeQuoterRef)
			tokenTransferFeeConfigArgs := make([]fqops.TokenTransferFeeConfigArgs, 0)
			allDestChainConfigs := make([]fqops.DestChainConfigArgs, 0)
			var providedRemoteChains map[uint64]struct{}
			if len(input.RemoteChainSelectors) > 0 {
				// initialize providedRemoteChains map if remote chains are provided in the input,
				// this means we only want to import config for those remote chains from 1.6
				providedRemoteChains = make(map[uint64]struct{})
				for _, remoteChain := range input.RemoteChainSelectors {
					providedRemoteChains[remoteChain] = struct{}{}
				}
			}
			for remoteChain, cfg := range fqOutput.RemoteChainCfgs {
				if !cfg.DestChainCfg.IsEnabled {
					continue
				}
				// check if the remote chain is connected with 1.6 deployment, if not, we skip importing config for that remote chain from FQ 1.6
				// this is to safeguard against having incorrect config import from 1.6
				version, err := adapters1_2.GetLaneVersionForRemoteChain(b.GetContext(), chain, remoteChain, common.HexToAddress(routerAddr.Address))
				if err != nil {
					return FeeQuoterUpdate{}, fmt.Errorf("failed to get lane version for remote chain %d: %w", remoteChain, err)
				}
				if version == nil || !version.Equal(semver.MustParse("1.6.0")) {
					continue
				}
				// if remote chains are provided in the input, we only import config for those remote chains,
				// otherwise we import config for all supported remote chains in 1.6
				if providedRemoteChains != nil {
					if _, exists := providedRemoteChains[remoteChain]; !exists {
						continue
					}
				}
				destChainConfig := cfg.DestChainCfg
				// check if gasprice stateness threashold is zero
				if destChainConfig.GasPriceStalenessThreshold == 0 {
					priceUpdates, err := HandleEmptyGasPriceStalenessThreshold(remoteChain, input)
					if err != nil {
						return FeeQuoterUpdate{}, fmt.Errorf("failed to handle empty gas price staleness threshold for remote chain %d: %w", remoteChain, err)
					}
					output.PriceUpdates.GasPriceUpdates = append(output.PriceUpdates.GasPriceUpdates, priceUpdates.GasPriceUpdates...)
					output.PriceUpdates.TokenPriceUpdates = append(output.PriceUpdates.TokenPriceUpdates, priceUpdates.TokenPriceUpdates...)
				}
				outDestchainCfg := fqops.DestChainConfigArgs{
					DestChainSelector: remoteChain,
					DestChainConfig: fqops.DestChainConfig{
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
				tokenTransferFeeCfgs := make([]fqops.TokenTransferFeeConfigSingleTokenArgs, 0)
				for token, transferCfg := range cfg.TokenTransferFeeCfgs {
					if !transferCfg.IsEnabled {
						continue
					}
					tokenTransferFeeCfgs = append(tokenTransferFeeCfgs, fqops.TokenTransferFeeConfigSingleTokenArgs{
						Token: token,
						TokenTransferFeeConfig: fqops.TokenTransferFeeConfig{
							FeeUSDCents:       transferCfg.MinFeeUSDCents,
							DestGasOverhead:   transferCfg.DestGasOverhead,
							DestBytesOverhead: transferCfg.DestBytesOverhead,
							IsEnabled:         transferCfg.IsEnabled,
						},
					})
				}
				tokenTransferFeeConfigArgs = append(tokenTransferFeeConfigArgs, fqops.TokenTransferFeeConfigArgs{
					DestChainSelector:       remoteChain,
					TokenTransferFeeConfigs: tokenTransferFeeCfgs,
				})
				allDestChainConfigs = append(allDestChainConfigs, outDestchainCfg)
			}
			if isNewFQV2Deployment {
				// if new deployment, adding deployer key as price updater so that
				// manual gas prices can be set right after deployment if needed
				output.ConstructorArgs = fqops.ConstructorArgs{
					StaticConfig: fqops.StaticConfig{
						LinkToken:         fqOutput.StaticCfg.LinkToken,
						MaxFeeJuelsPerMsg: fqOutput.StaticCfg.MaxFeeJuelsPerMsg,
					},
					PriceUpdaters:              append(fqOutput.PriceUpdaters, chain.DeployerKey.From),
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
		semver.MustParse("2.0.0"),
		"Creates FeeQuoterUpdate input by importing configuration from PriceRegistry v1.5.0 and EVM2EVMOnRamp v1.5.0",
		func(b cldf_ops.Bundle, chain evm.Chain, input deploy.FeeQuoterUpdateInput) (output FeeQuoterUpdate, err error) {
			// get address ref for onramp 1.5.0
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
			routerAddr := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				routerops.ContractType,
				routerops.Version,
				"",
			)
			if routerAddr.Address == "" {
				return FeeQuoterUpdate{}, fmt.Errorf("failed to find router address ref for chain selector %d", input.ChainSelector)
			}
			// get the commit stores and that will act like price updaters for fee quoter
			var commitStoreRefs []datastore.AddressRef
			for _, addressRef := range input.ExistingAddresses {
				if addressRef.Type == "CommitStore" &&
					addressRef.Version.Equal(semver.MustParse("1.5.0")) &&
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
			feeQuoterV2Ref := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fqops.ContractType,
				fqops.Version,
				"",
			)
			isNewFQv2Deployment := datastore_utils.IsAddressRefEmpty(feeQuoterV2Ref)

			var staticCfg fqops.StaticConfig
			var destChainCfgs []fqops.DestChainConfigArgs
			var tokenTransferFeeConfigArgsForAll []fqops.TokenTransferFeeConfigArgs
			var providedRemoteChains map[uint64]struct{}
			if len(input.RemoteChainSelectors) > 0 {
				// initialize providedRemoteChains map if remote chain selectors are provided in the input,
				// so that we can check against this map when importing config for each remote chain from onRamp 1.5.0
				providedRemoteChains = make(map[uint64]struct{})
				for _, remoteChain := range input.RemoteChainSelectors {
					providedRemoteChains[remoteChain] = struct{}{}
				}
			}
			for _, meta := range onRampMetadata {
				var tokenTransferFeeConfigArgs []fqops.TokenTransferFeeConfigSingleTokenArgs

				// Convert metadata to typed struct if needed
				onRampCfg, err := datastore_utils.ConvertMetadataToType[seq1_5.OnRampImportConfigSequenceOutput](meta.Metadata)
				if err != nil {
					return FeeQuoterUpdate{}, fmt.Errorf("failed to convert metadata to "+
						"OnRampImportConfigSequenceOutput for chain selector %d: %w", input.ChainSelector, err)
				}
				remoteChain := onRampCfg.RemoteChainSelector
				// check if the remote chain is connected with 1.5 deployment, if not, we skip importing config for that remote chain from OnRamp 1.5
				// this is to safeguard against having incorrect config import from 1.5
				version, err := adapters1_2.GetLaneVersionForRemoteChain(b.GetContext(), chain, remoteChain, common.HexToAddress(routerAddr.Address))
				if err != nil {
					return FeeQuoterUpdate{}, fmt.Errorf("failed to get lane version for remote chain %d: %w", remoteChain, err)
				}
				if version == nil || !version.Equal(semver.MustParse("1.5.0")) {
					continue
				}
				// if remote chains are provided in the input, we only import config for those remote chains,
				// otherwise we import config for all supported remote chains in the 1.5
				if providedRemoteChains != nil {
					if _, exists := providedRemoteChains[remoteChain]; !exists {
						continue
					}
				}
				if staticCfg.LinkToken == (common.Address{}) {
					staticCfg = fqops.StaticConfig{
						LinkToken:         onRampCfg.StaticConfig.LinkToken,
						MaxFeeJuelsPerMsg: onRampCfg.StaticConfig.MaxNopFeesJuels,
					}
				}
				chainFamilySelector := utils.GetSelectorHex(onRampCfg.RemoteChainSelector)

				destChainCfgs = append(destChainCfgs, fqops.DestChainConfigArgs{
					DestChainSelector: onRampCfg.RemoteChainSelector,
					DestChainConfig: fqops.DestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                onRampCfg.DynamicConfig.MaxDataBytes,
						MaxPerMsgGasLimit:           onRampCfg.DynamicConfig.MaxPerMsgGasLimit,
						DestGasOverhead:             onRampCfg.DynamicConfig.DestGasOverhead,
						DestGasPerPayloadByteBase:   uint8(onRampCfg.DynamicConfig.DestGasPerPayloadByte),
						ChainFamilySelector:         chainFamilySelector,
						DefaultTokenFeeUSDCents:     onRampCfg.DynamicConfig.DefaultTokenFeeUSDCents,
						DefaultTokenDestGasOverhead: onRampCfg.DynamicConfig.DefaultTokenDestGasOverhead,
						DefaultTxGasLimit:           uint32(onRampCfg.StaticConfig.DefaultTxGasLimit),
						NetworkFeeUSDCents:          NetworkFeeUSDCents,
						LinkFeeMultiplierPercent:    LinkFeeMultiplierPercent,
					},
				})
				for token, tokenCfg := range onRampCfg.TokenTransferFeeConfig {
					tokenTransferFeeConfigArgs = append(tokenTransferFeeConfigArgs, fqops.TokenTransferFeeConfigSingleTokenArgs{
						Token: token,
						TokenTransferFeeConfig: fqops.TokenTransferFeeConfig{
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
			if isNewFQv2Deployment {
				output.ConstructorArgs = fqops.ConstructorArgs{
					StaticConfig: fqops.StaticConfig{
						LinkToken:         staticCfg.LinkToken,
						MaxFeeJuelsPerMsg: staticCfg.MaxFeeJuelsPerMsg,
					},
					DestChainConfigArgs:        destChainCfgs,
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgsForAll,
					PriceUpdaters:              append(priceUpdaters, chain.DeployerKey.From),
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

// MergeFeeQuoterUpdateOutputs merges FeeQuoterUpdate outputs from the v1.6.x and v1.5.0 import
// sequences into a single update. output16 is the base; output15 supplements it. Where both
// provide values (e.g. ConstructorArgs, dest chain configs, token transfer fee configs),
// output16 takes precedence and output15 fills in only missing entries.
func MergeFeeQuoterUpdateOutputs(output16, output15 FeeQuoterUpdate) (FeeQuoterUpdate, error) {
	result := output16
	empty16, err := output16.IsEmpty()
	if err != nil {
		return FeeQuoterUpdate{}, fmt.Errorf("failed to check if output16 is empty: %w", err)
	}
	empty15, err := output15.IsEmpty()
	if err != nil {
		return FeeQuoterUpdate{}, fmt.Errorf("failed to check if output15 is empty: %w", err)
	}

	if empty16 && empty15 {
		return FeeQuoterUpdate{}, nil
	}

	// if output16 is empty, we can just return output15
	if empty16 {
		return output15, nil
	}
	// if output15 is empty, we can just return output16
	if empty15 {
		return output16, nil
	}
	// ConstructorArgs: use output15 if output16 is empty
	if IsConstructorArgsEmpty(result.ConstructorArgs) {
		result.ConstructorArgs = output15.ConstructorArgs
	} else {
		// merge the dest chainConfig args
		result.ConstructorArgs.DestChainConfigArgs = mergeDestChainConfigs(
			result.ConstructorArgs.DestChainConfigArgs,
			output15.ConstructorArgs.DestChainConfigArgs)

		resultPriceUpdatersMap := make(map[common.Address]bool)
		for _, updater := range output15.ConstructorArgs.PriceUpdaters {
			resultPriceUpdatersMap[updater] = true
		}
		for _, updater := range output16.ConstructorArgs.PriceUpdaters {
			resultPriceUpdatersMap[updater] = true
		}

		result.ConstructorArgs.PriceUpdaters = maps.Keys(resultPriceUpdatersMap)

		result.ConstructorArgs.TokenTransferFeeConfigArgs = mergeTokenTransferFeeConfigArgs(
			result.ConstructorArgs.TokenTransferFeeConfigArgs,
			output15.ConstructorArgs.TokenTransferFeeConfigArgs)
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

func mergeTokenTransferFeeConfigArgs(args1, args2 []fqops.TokenTransferFeeConfigArgs) []fqops.TokenTransferFeeConfigArgs {
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

// mergeDestChainConfigs merges two slices of DestChainConfigArgs, giving precedence to the first slice in case of duplicate DestChainSelectors.
func mergeDestChainConfigs(cfgs1, cfgs2 []fqops.DestChainConfigArgs) []fqops.DestChainConfigArgs {
	result := cfgs1

	// Create a map of dest chain selectors from cfgs1 which will be skipped when adding from cfgs2
	destChainMap := make(map[uint64]fqops.DestChainConfigArgs)
	for _, cfg := range cfgs1 {
		destChainMap[cfg.DestChainSelector] = cfg
	}

	// Add configs from cfgs2 that don't exist in cfgs1
	for _, cfg := range cfgs2 {
		if _, exists := destChainMap[cfg.DestChainSelector]; !exists {
			result = append(result, cfg)
		}
		// If it exists in both, cfgs1's value is already used (takes precedence)
	}

	return result
}

func IsConstructorArgsEmpty(a fqops.ConstructorArgs) bool {
	return (a.StaticConfig == fqops.StaticConfig{}) &&
		len(a.PriceUpdaters) == 0 &&
		len(a.TokenTransferFeeConfigArgs) == 0 &&
		len(a.DestChainConfigArgs) == 0
}

// HandleEmptyGasPriceStalenessThreshold handles the case when GasPriceStalenessThreshold is zero for a remote chain.
// It first looks for gas price for that remote chain in the input additional config (GasPricesPerRemoteChain).
// If found and valid, it adds that price to the output. If not found, it checks whether the chain family
// has a hardcoded static price (e.g. Aptos and Sui in staticGasPriceByChainFamily); if so, it uses that
// and adds it to the output. If the chain family has no hardcoded price, it returns empty price updates.
// Returns an error only for an invalid gas price string in config or failure to resolve the chain family.
// It is exported for testing.
func HandleEmptyGasPriceStalenessThreshold(remoteChain uint64, input deploy.FeeQuoterUpdateInput) (output fqops.PriceUpdates, err error) {
	var staticPrice *big.Int
	if input.AdditionalConfig != nil && input.AdditionalConfig.GasPricesPerRemoteChain != nil {
		gaspriceStr, ok := input.AdditionalConfig.GasPricesPerRemoteChain[remoteChain]
		if ok {
			var success bool
			staticPrice, success = new(big.Int).SetString(gaspriceStr, 10)
			if !success {
				return fqops.PriceUpdates{}, fmt.Errorf("invalid gas price %s for remote chain %d in input additional config", gaspriceStr, remoteChain)
			}
		}
	}
	if staticPrice == nil {
		// check if static gas price is already hard coded for the chain family
		chainFamily, err := chain_selectors.GetSelectorFamily(remoteChain)
		if err != nil {
			return fqops.PriceUpdates{}, fmt.Errorf("failed to get chain family for remote chain %d: %w", remoteChain, err)
		}
		var exists bool
		staticPrice, exists = staticGasPriceByChainFamily[chainFamily]
		if !exists || staticPrice == nil {
			return fqops.PriceUpdates{}, nil
		}
	}

	output.GasPriceUpdates = append(output.GasPriceUpdates, fqops.GasPriceUpdate{
		DestChainSelector: remoteChain,
		UsdPerUnitGas:     staticPrice,
	})
	return output, nil
}
