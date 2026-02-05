package sequences

import (
	"encoding/hex"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	seq1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"

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

var (
	EVMFamilySelector, _ = hex.DecodeString("2812d52c")
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

	// CreateFeeQuoterUpdateInputFromV160 creates FeeQuoterUpdate input by importing configuration from FeeQuoter v1.6.0
	CreateFeeQuoterUpdateInputFromV160 = cldf_ops.NewSequence(
		"fetches-feequoter-config-values-from-v1.6.0",
		semver.MustParse("1.7.0"),
		"Creates FeeQuoterUpdate input by importing configuration from FeeQuoter v1.6.0",
		func(b cldf_ops.Bundle, chain evm.Chain, input deploy.FeeQuoterUpdateInput) (output FeeQuoterUpdate, err error) {
			// get feeQuoter 1.6 address
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
			fqOutput := metadataForFq16[0].Metadata.(seq1_6.FeeQuoterImportConfigSequenceOutput)
			// is feeQuoter going to be deployed or fetched from existing addresses?
			feeQuoterRef := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fqops.ContractType,
				fqops.Version,
				"",
			)
			isNewDeployment := datastore_utils.IsAddressRefEmpty(feeQuoterRef)

			if isNewDeployment {
				output.ConstructorArgs = fqops.ConstructorArgs{
					StaticConfig: fqops.StaticConfig{
						LinkToken:         fqOutput.StaticCfg.LinkToken,
						MaxFeeJuelsPerMsg: fqOutput.StaticCfg.MaxFeeJuelsPerMsg,
					},
					PriceUpdaters: fqOutput.PriceUpdaters,
				}
			} else {
				output.AuthorizedCallerUpdates = fqops.AuthorizedCallerArgs{
					AddedCallers: fqOutput.PriceUpdaters,
				}
			}
			for remoteChain, cfg := range fqOutput.RemoteChainCfgs {
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
						LinkFeeMultiplierPercent:    90,
					},
				}
				tokenTransferFeeCfgs := make([]fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs, 0)
				for token, transferCfg := range cfg.TokenTransferFeeCfgs {
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
				if isNewDeployment {
					output.ConstructorArgs.DestChainConfigArgs = append(output.ConstructorArgs.DestChainConfigArgs, outDestchainCfg)
					output.ConstructorArgs.TokenTransferFeeConfigArgs = append(output.ConstructorArgs.TokenTransferFeeConfigArgs,
						fqops.TokenTransferFeeConfigArgs{
							DestChainSelector:       remoteChain,
							TokenTransferFeeConfigs: tokenTransferFeeCfgs,
						})
				} else {
					output.DestChainConfigs = append(output.DestChainConfigs, outDestchainCfg)
					output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs = append(output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs,
						fqops.TokenTransferFeeConfigArgs{
							DestChainSelector:       remoteChain,
							TokenTransferFeeConfigs: tokenTransferFeeCfgs,
						})
				}
			}
			return output, nil
		})

	// CreateFeeQuoterUpdateInputFromV150 creates FeeQuoterUpdate input by importing configuration from PriceRegistry v1.5.0 and EVM2EVMOnRamp v1.5.0
	CreateFeeQuoterUpdateInputFromV150 = cldf_ops.NewSequence(
		"fetches-feequoter-config-values-from-v1.5.0",
		semver.MustParse("1.7.0"),
		"Creates FeeQuoterUpdate input by importing configuration from PriceRegistry v1.5.0 and EVM2EVMOnRamp v1.5.0",
		func(b cldf_ops.Bundle, chain evm.Chain, input deploy.FeeQuoterUpdateInput) (output FeeQuoterUpdate, err error) {
			// get ro
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

			// is feeQuoter going to be deployed or fetched from existing addresses?
			feeQuoterRef := datastore_utils.GetAddressRef(
				input.ExistingAddresses,
				input.ChainSelector,
				fqops.ContractType,
				fqops.Version,
				"",
			)
			isNewDeployment := datastore_utils.IsAddressRefEmpty(feeQuoterRef)
			var staticCfg fqops.StaticConfig
			var destChainCfgs []fqops.DestChainConfigArgs
			var tokenTransferFeeConfigArgs []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs
			var tokenTransferFeeConfigArgsForAll []fqops.TokenTransferFeeConfigArgs
			for _, meta := range onRampMetadata {
				onRampCfg := meta.Metadata.(seq1_5.OnRampImportConfigSequenceOutput)
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
				destChainCfgs = append(destChainCfgs, fqops.DestChainConfigArgs{
					DestChainSelector: onRampCfg.RemoteChainSelector,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true, // if the chain is supported on OnRamp, we should enable it on FeeQuoter
						MaxDataBytes:                onRampCfg.DynamicConfig.MaxDataBytes,
						MaxPerMsgGasLimit:           onRampCfg.DynamicConfig.MaxPerMsgGasLimit,
						DestGasOverhead:             onRampCfg.DynamicConfig.DestGasOverhead,
						DestGasPerPayloadByteBase:   uint8(onRampCfg.DynamicConfig.DestGasPerPayloadByte),
						ChainFamilySelector:         [4]byte(EVMFamilySelector), // as this is evm, safe to assume chain family selector is same as chain selector
						DefaultTokenFeeUSDCents:     onRampCfg.DynamicConfig.DefaultTokenFeeUSDCents,
						DefaultTokenDestGasOverhead: onRampCfg.DynamicConfig.DefaultTokenDestGasOverhead,
						DefaultTxGasLimit:           uint32(onRampCfg.StaticConfig.DefaultTxGasLimit),
						NetworkFeeUSDCents:          networkFeeUSDCents,
						LinkFeeMultiplierPercent:    90,
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
			if isNewDeployment {
				output.ConstructorArgs = fqops.ConstructorArgs{
					StaticConfig: fqops.StaticConfig{
						LinkToken:         staticCfg.LinkToken,
						MaxFeeJuelsPerMsg: staticCfg.MaxFeeJuelsPerMsg,
					},
					DestChainConfigArgs:        destChainCfgs,
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgsForAll,
					// TODO: what to do with price updaters for 1.5 if there is no 1.6 lanes here
					PriceUpdaters: []common.Address{},
				}
			} else {
				output.DestChainConfigs = destChainCfgs
				output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs = tokenTransferFeeConfigArgsForAll
			}
			return output, nil
		})
)
