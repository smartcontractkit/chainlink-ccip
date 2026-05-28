package tokens

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	adaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var DeployTokenPool = cldf_ops.NewSequence(
	"deploy-token-pool",
	utils.Version_2_0_0,
	"Deploy a 2.0.0 token pool for a token on an EVM chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
		}

		// Validate required deployment inputs
		poolutil := adaptersV1_0_0.EVMTokenBase{}
		if input.TokenPoolVersion == nil {
			return sequences.OnChainOutput{}, errors.New("TokenPoolVersion is required")
		}
		if input.TokenRef == nil {
			return sequences.OnChainOutput{}, errors.New("TokenRef is required")
		}

		// Parse the token ref as an EVM address
		tokenAddress, err := poolutil.ParseNonZeroAddressRef(input.ExistingDataStore, input.TokenRef.Clone(), chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve token address from ref: %w", err)
		}

		// If no pool qualifier is provided, then fall back to using the token address
		poolQualifier := input.TokenPoolQualifier
		if poolQualifier == "" {
			poolQualifier = tokenAddress.Hex()
		}

		// NOTE: the datastore uses the type, selector, qualifier, and version of an address
		// ref to uniquely identify records, so the query below should only match one record
		// at most. If multiple records are returned, then this would indicate an issue with
		// the datastore's data integrity. If no matches are returned, then the ref does not
		// exist and we proceed with the deployment.
		var tokenPoolAddress common.Address
		matches := input.ExistingDataStore.Addresses().Filter(
			datastore.AddressRefByType(datastore.ContractType(input.PoolType)),
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByQualifier(poolQualifier),
			datastore.AddressRefByVersion(input.TokenPoolVersion),
		)
		if len(matches) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"multiple token pools found in datastore with type '%s', version '%s', qualifier '%s' on chain with selector %d",
				input.PoolType, input.TokenPoolVersion.String(), poolQualifier, chain.Selector,
			)
		}
		if len(matches) == 1 {
			tokenPoolAddress = common.HexToAddress(matches[0].Address)
		}

		// If the token pool is already deployed, then apply any dynamic configuration updates the
		// caller gave (e.g. router, rate-limit admin, fee aggregator, additional-CCVs threshold).
		// This allows the seq to be re-run idempotently with an updated config without needing to
		// tear down and re-deploy the pool.
		if tokenPoolAddress != (common.Address{}) {
			b.Logger.Infof("Token pool already deployed on chain %d at address %q - updating dynamic pool config if needed", chain.Selector, tokenPoolAddress.Hex())
			configureInput := ConfigureTokenPoolInput{}

			// Populate configureInput with any dynamic config fields that the caller entered.
			// Empty/zero values are ignored and result in no change to those fields on-chain.
			if input.ThresholdAmountForAdditionalCCVs != "" {
				threshold, ok := new(big.Int).SetString(input.ThresholdAmountForAdditionalCCVs, 10)
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid ThresholdAmountForAdditionalCCVs '%s': must be a decimal integer string", input.ThresholdAmountForAdditionalCCVs)
				}
				report, err := cldf_ops.ExecuteOperation(b,
					token_pool.GetAdvancedPoolHooks, chain,
					contract.FunctionInput[struct{}]{
						ChainSelector: chain.Selector,
						Address:       tokenPoolAddress,
					},
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to read advanced pool hooks address from existing token pool %s on chain %d: %w", tokenPoolAddress, chain.Selector, err)
				}
				configureInput.ThresholdAmountForAdditionalCCVs = threshold
				configureInput.AdvancedPoolHooks = report.Output
			}
			if input.RateLimitAdmin != "" {
				if !common.IsHexAddress(input.RateLimitAdmin) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid RateLimitAdmin address '%s'", input.RateLimitAdmin)
				} else {
					configureInput.RateLimitAdmin = common.HexToAddress(input.RateLimitAdmin)
				}
			}
			if input.FeeAggregator != "" {
				if !common.IsHexAddress(input.FeeAggregator) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeAggregator address '%s'", input.FeeAggregator)
				} else {
					configureInput.FeeAggregator = common.HexToAddress(input.FeeAggregator)
				}
			}
			if input.RouterRef != nil {
				if routerAddr, err := poolutil.ResolveRouterAddress(input.ExistingDataStore, chain.Selector, input.RouterRef); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve router address for ref (%s): %w", datastore_utils.SprintRef(*input.RouterRef), err)
				} else {
					configureInput.RouterAddress = routerAddr
				}
			}

			// If the caller did not provide any dynamic config fields to update, then
			// skip the configure step and return early.
			if configureInput == (ConfigureTokenPoolInput{}) {
				return sequences.OnChainOutput{Addresses: matches}, nil
			} else {
				configureInput.TokenPoolAddress = tokenPoolAddress
				configureInput.ChainSelector = chain.Selector
			}

			// ConfigureTokenPool reads current values and only emits a write when they
			// differ so reruns with the same inputs are no-ops. Fields that the caller
			// leaves unset (zero/empty) retain their current on-chain values.
			output := sequences.OnChainOutput{Addresses: matches}
			if report, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, configureInput); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to reconcile dynamic config for existing token pool %s on chain %d: %w", tokenPoolAddress, chain.Selector, err)
			} else {
				output.Addresses = append(output.Addresses, report.Output.Addresses...)
				output.BatchOps = report.Output.BatchOps
			}

			// Set the allowed finality config (if applicable) - this does not produce a
			// batch if the current on-chain config already matches the requested config
			if !input.AllowedFinalityConfig.IsZero() {
				if report, err := cldf_ops.ExecuteSequence(b, SetAllowedFinalityConfigForTokenPools, chains, tokens.SetAllowedFinalityConfigSequenceInput{
					Settings: map[string]finality.Config{tokenPoolAddress.Hex(): input.AllowedFinalityConfig},
					Selector: chain.Selector,
				}); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set allowed finality config for existing token pool %s on chain %d: %w", tokenPoolAddress, chain.Selector, err)
				} else {
					output.Addresses = append(output.Addresses, report.Output.Addresses...)
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				}
			}

			return output, nil
		}

		// Infer pool deployment inputs
		tokenDecimals, err := poolutil.ERC20Decimals(b, input.ExistingDataStore, chain, tokenAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals for token at address '%s': %w", tokenAddress, err)
		}
		rmnProxyAddr, err := poolutil.GetRMNProxyAddress(input.ExistingDataStore, chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve rmn proxy address for chain selector %d: %w", chain.Selector, err)
		}
		routerAddr, err := poolutil.ResolveRouterAddress(input.ExistingDataStore, chain.Selector, input.RouterRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve router address for chain selector %d: %w", chain.Selector, err)
		}
		allowlist, err := poolutil.ParseAddressStrings(input.Allowlist)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to parse allowlist: %w", err)
		}

		// Resolve pool configuration inputs
		var rateLimitAdmin common.Address
		if input.RateLimitAdmin != "" {
			if !common.IsHexAddress(input.RateLimitAdmin) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid RateLimitAdmin address '%s'", input.RateLimitAdmin)
			} else {
				rateLimitAdmin = common.HexToAddress(input.RateLimitAdmin)
			}
		}
		var feeAggregator common.Address
		if input.FeeAggregator != "" {
			if !common.IsHexAddress(input.FeeAggregator) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeAggregator address '%s'", input.FeeAggregator)
			} else {
				feeAggregator = common.HexToAddress(input.FeeAggregator)
			}
		}
		thresholdCCV := big.NewInt(0)
		if input.ThresholdAmountForAdditionalCCVs != "" {
			if threshold, ok := new(big.Int).SetString(input.ThresholdAmountForAdditionalCCVs, 10); !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid ThresholdAmountForAdditionalCCVs '%s': must be a decimal integer string", input.ThresholdAmountForAdditionalCCVs)
			} else {
				thresholdCCV = threshold
			}
		}

		// Build the pool deployment input
		tokenPoolType := datastore.ContractType(input.PoolType)
		internalInput := DeployTokenPoolInput{
			TokenPoolVersion:                 input.TokenPoolVersion,
			TokenPoolType:                    tokenPoolType,
			ChainSel:                         chain.Selector,
			TokenSymbol:                      poolQualifier,
			RateLimitAdmin:                   rateLimitAdmin,
			FeeAggregator:                    feeAggregator,
			ThresholdAmountForAdditionalCCVs: thresholdCCV,
			ConstructorArgs: ConstructorArgs{
				Token:    tokenAddress,
				Decimals: tokenDecimals,
				RMNProxy: rmnProxyAddr,
				Router:   routerAddr,
			},
			AdvancedPoolHooksConfig: AdvancedPoolHooksConfig{
				Allowlist: allowlist,
			},
		}

		// Deploy the desired pool contract
		output := sequences.OnChainOutput{Addresses: matches}
		switch {
		case poolutil.IsLockReleasePoolType(tokenPoolType.String()):
			if report, err := cldf_ops.ExecuteSequence(b, DeployLockReleaseTokenPool, chain, internalInput); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy lock release token pool on chain %d: %w", chain.Selector, err)
			} else {
				output.Addresses = append(output.Addresses, report.Output.Addresses...)
				output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
			}
		case poolutil.IsBurnMintPoolType(tokenPoolType.String()):
			if report, err := cldf_ops.ExecuteSequence(b, DeployBurnMintTokenPool, chain, internalInput); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token pool on chain %d: %w", chain.Selector, err)
			} else {
				output.Addresses = append(output.Addresses, report.Output.Addresses...)
				output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
			}
		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for chain with selector %d", input.PoolType, chain.Selector)
		}

		// Configure the finality config (if applicable)
		if !input.AllowedFinalityConfig.IsZero() && len(output.Addresses) > 0 {
			poolRef := output.Addresses[0]
			if report, err := cldf_ops.ExecuteSequence(b, SetAllowedFinalityConfigForTokenPools, chains, tokens.SetAllowedFinalityConfigSequenceInput{
				Settings: map[string]finality.Config{poolRef.Address: input.AllowedFinalityConfig},
				Selector: chain.Selector,
			}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set allowed finality config for token pool (%s) on chain %d: %w", datastore_utils.SprintRef(poolRef), chain.Selector, err)
			} else {
				output.Addresses = append(output.Addresses, report.Output.Addresses...)
				output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
			}
		}

		// Return the deployment output
		return output, nil
	},
)
