package sequences

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpOpsV1_5_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/token_pool"
	tpSeqV1_5_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences/token_pool"
	tpOpsV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	tpSeqV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tpV1_5_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tpV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *EVMAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:configure-token-for-transfers",
		cciputils.Version_1_6_0,
		"Configure a token for cross-chain transfers for an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			if !common.IsHexAddress(input.TokenPoolAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address %q is not a valid hex address", input.TokenPoolAddress)
			}

			tpAddr := common.HexToAddress(input.TokenPoolAddress)
			if tpAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("token pool address is zero address")
			}

			externalAdmin := common.Address{}
			if input.ExternalAdmin != "" {
				if !common.IsHexAddress(input.ExternalAdmin) {
					return sequences.OnChainOutput{}, fmt.Errorf("external admin address %q is not a valid hex address", input.ExternalAdmin)
				}

				externalAdmin = common.HexToAddress(input.ExternalAdmin)
			}

			tarAddress, err := a.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			filters := datastore.AddressRef{
				ChainSelector: input.ChainSelector,
				Address:       tpAddr.String(),
			}
			fullTpRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, filters, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool in datastore using ref (%+v): %w", filters, err)
			}

			tokenAddress, err := a.GetTokenAddressFromFullTokenPoolRef(b, chain, fullTpRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool ref (%+v): %w", fullTpRef, err)
			}

			switch fullTpRef.Version.String() {
			case tpOpsV1_5_1.Version.String():
				if configureReport, err := cldf_ops.ExecuteSequence(b,
					tpSeqV1_5_1.ConfigureTokenPoolForRemoteChains, chain,
					tpSeqV1_5_1.ConfigureTokenPoolForRemoteChainsInput{
						TokenPoolAddress: tpAddr,
						TokenPoolVersion: fullTpRef.Version,
						RemoteChains:     input.RemoteChains,
					},
				); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
				} else {
					result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
					result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)
				}
			case tpOpsV1_6_1.Version.String():
				if configureReport, err := cldf_ops.ExecuteSequence(b,
					tpSeqV1_6_1.ConfigureTokenPoolForRemoteChains, chain,
					tpSeqV1_6_1.ConfigureTokenPoolForRemoteChainsInput{
						TokenPoolAddress: tpAddr,
						TokenPoolVersion: fullTpRef.Version,
						RemoteChains:     input.RemoteChains,
					},
				); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
				} else {
					result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
					result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)
				}
			default:
				return sequences.OnChainOutput{}, fmt.Errorf(
					"unsupported token pool version %s for token pool at address %q on chain selector %d",
					fullTpRef.Version.String(), tpAddr.String(), input.ChainSelector,
				)
			}

			registerReport, err := cldf_ops.ExecuteSequence(b,
				tarseq.RegisterToken,
				chain,
				tarseq.RegisterTokenInput{
					ChainSelector:             input.ChainSelector,
					TokenAdminRegistryAddress: tarAddress,
					TokenPoolAddress:          tpAddr,
					ExternalAdmin:             externalAdmin,
					TokenAddress:              tokenAddress,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token on chain %d: %w", input.ChainSelector, err)
			}
			result.Addresses = append(result.Addresses, registerReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, registerReport.Output.BatchOps...)

			return result, nil
		})
}

func (a *EVMAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	addrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, poolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find token pool in datastore using ref (%+v): %w", poolRef, err)
	}

	addrRaw, err := a.AddressRefToBytes(addrRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	tpAddr := common.BytesToAddress(addrRaw)
	if tpAddr == (common.Address{}) {
		return nil, errors.New("token pool address is zero address")
	}

	tokenAddr, err := a.GetTokenAddressFromFullTokenPoolRef(e.OperationsBundle, chain, addrRef)
	if err != nil {
		return nil, fmt.Errorf("failed to get token address from token pool ref (%+v): %w", addrRef, err)
	}

	return tokenAddr.Bytes(), nil
}

func (a *EVMAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	addrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, poolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return 0, fmt.Errorf("failed to find token pool in datastore using ref (%+v): %w", poolRef, err)
	}

	addrRaw, err := a.AddressRefToBytes(addrRef)
	if err != nil {
		return 0, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	tpAddr := common.BytesToAddress(addrRaw)
	if tpAddr == (common.Address{}) {
		return 0, errors.New("token pool address is zero address")
	}

	type TokenPoolDecimalReader interface {
		GetTokenDecimals(*bind.CallOpts) (uint8, error)
	}

	var pool TokenPoolDecimalReader
	switch addrRef.Version.String() {
	case tpOpsV1_5_1.Version.String():
		if tp, err := tpV1_5_1.NewTokenPool(tpAddr, chain.Client); err != nil {
			return 0, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
		} else {
			pool = tp
		}
	case tpOpsV1_6_1.Version.String():
		if tp, err := tpV1_6_1.NewTokenPool(tpAddr, chain.Client); err != nil {
			return 0, fmt.Errorf("failed to instantiate token pool v1.6.1 contract: %w", err)
		} else {
			pool = tp
		}
	default:
		return 0, fmt.Errorf("unsupported token pool version %s for token pool at address %q on chain selector %d", addrRef.Version.String(), tpAddr.Hex(), addrRef.ChainSelector)
	}

	return pool.GetTokenDecimals(&bind.CallOpts{Context: e.GetContext()})
}

func (a *EVMAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokensapi.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:set-token-pool-rate-limits",
		cciputils.Version_1_6_0,
		"Set rate limits for a token pool across multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.TPRLRemotes) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			// NOTE: the top level changeset will fully populate `input.TokenPoolRef` BEFORE calling this sequence,
			// so we can safely assume all its fields will be accounted for and avoid re-querying the datastore. We
			// use `AddressRefToBytes(*)` as a shortcut to avoid writing the same `common.IsHexAddress` followed by
			// `common.HexToAddress` boilerplate.
			tokenPoolAddrBytes, err := a.AddressRefToBytes(input.TokenPoolRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", err)
			}
			tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
			if tokenPoolAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref (%+v) is zero address", input.TokenPoolRef)
			}

			// NOTE: there are certain situations where we don't have access to change the rate limits (e.g. the customer is the
			// pool owner and rate limit admin). In these cases, we should NOT attempt to set the rate limit or generate an MCMS
			// batch since we know for certain that the transaction will fail. Instead, we log a warning and skip the rate limit
			// configuration for that chain (which effectively turns this into a no-op).
			if input.SkipIfMissingPermissions {
				timelockFltr := datastore.AddressRef{Type: datastore.ContractType(cciputils.RBACTimelock), ChainSelector: chain.Selector, Qualifier: cciputils.CLLQualifier}
				timelockAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, timelockFltr, chain.Selector, datastore_utils_evm.ToEVMAddress)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to find timelock address for chain %d: %w", chain.Selector, err)
				}
				poolOwner, rlAdmin, err := a.GetTokenPoolAdmins(b.GetContext(), &chain, input.TokenPoolRef)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool admins for token pool ref (%+v) on chain %d: %w", input.TokenPoolRef, chain.Selector, err)
				}
				isRateLimitAdmin := rlAdmin == timelockAddr || rlAdmin == chain.DeployerKey.From
				isPoolOwner := poolOwner == timelockAddr || poolOwner == chain.DeployerKey.From
				if !isRateLimitAdmin && !isPoolOwner {
					b.Logger.Warnf(
						"Timelock address %q and deployer address %q are not the owner or rate limit admin for token pool at address %q on chain selector %d. Skipping rate limiter config for this chain.",
						timelockAddr.Hex(), chain.DeployerKey.From.Hex(), tokenPoolAddr.Hex(), chain.Selector,
					)
					return sequences.OnChainOutput{}, nil
				}
			}

			var output evm_contract.WriteOutput
			switch input.TokenPoolRef.Version.String() {
			case tpOpsV1_5_1.Version.String():
				if report, err := cldf_ops.ExecuteOperation(b,
					tpOpsV1_5_1.SetChainRateLimiterConfig, chain,
					evm_contract.FunctionInput[tpOpsV1_5_1.SetChainRateLimiterConfigArgs]{
						ChainSelector: chain.Selector,
						Address:       tokenPoolAddr,
						Args: tpOpsV1_5_1.SetChainRateLimiterConfigArgs{
							OutboundRateLimitConfig: token_pool.RateLimiterConfig{
								IsEnabled: input.DefaultFinalityOutboundRateLimiterConfig.IsEnabled,
								Capacity:  input.DefaultFinalityOutboundRateLimiterConfig.Capacity,
								Rate:      input.DefaultFinalityOutboundRateLimiterConfig.Rate,
							},
							InboundRateLimitConfig: token_pool.RateLimiterConfig{
								IsEnabled: input.DefaultFinalityInboundRateLimiterConfig.IsEnabled,
								Capacity:  input.DefaultFinalityInboundRateLimiterConfig.Capacity,
								Rate:      input.DefaultFinalityInboundRateLimiterConfig.Rate,
							},
							RemoteChainSelector: input.RemoteChainSelector,
						},
					}); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
				} else {
					output = report.Output
				}
			case tpOpsV1_6_1.Version.String():
				if report, err := cldf_ops.ExecuteOperation(b,
					tpOpsV1_6_1.SetChainRateLimiterConfig, chain,
					evm_contract.FunctionInput[tpOpsV1_6_1.SetChainRateLimiterConfigArgs]{
						ChainSelector: chain.Selector,
						Address:       tokenPoolAddr,
						Args: tpOpsV1_6_1.SetChainRateLimiterConfigArgs{
							OutboundConfig: tpOpsV1_6_1.Config{
								IsEnabled: input.DefaultFinalityOutboundRateLimiterConfig.IsEnabled,
								Capacity:  input.DefaultFinalityOutboundRateLimiterConfig.Capacity,
								Rate:      input.DefaultFinalityOutboundRateLimiterConfig.Rate,
							},
							InboundConfig: tpOpsV1_6_1.Config{
								IsEnabled: input.DefaultFinalityInboundRateLimiterConfig.IsEnabled,
								Capacity:  input.DefaultFinalityInboundRateLimiterConfig.Capacity,
								Rate:      input.DefaultFinalityInboundRateLimiterConfig.Rate,
							},
							RemoteChainSelector: input.RemoteChainSelector,
						},
					}); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
				} else {
					output = report.Output
				}
			default:
				return sequences.OnChainOutput{}, fmt.Errorf(
					"unsupported token pool version %s for token pool with ref (%+v) on chain selector %d",
					input.TokenPoolRef.Version.String(), input.TokenPoolRef, input.ChainSelector,
				)
			}

			batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			result.BatchOps = append(result.BatchOps, batchOp)
			return result, nil
		})
}

func (a *EVMAdapter) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:manual-registration",
		cciputils.Version_1_6_0,
		"Manually register a token and token pool on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ManualRegistrationSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tokenAdminRegistryAddress, err := a.GetTokenAdminRegistryAddress(input.ExistingDataStore, chain.Selector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", chain.Selector, err)
			}

			// NOTE: the token address is derived using the following steps:
			//   1. if it is already present in the TokenRef, then skip the datastore altogether and simply use the given address
			//   2. if the token address is not present, then use whatever fields are defined in TokenRef to lookup the address in the DS
			//   3. if step #2 produces multiple tokens or an error, then attempt to resolve the token address from the TokenPoolRef
			//   4. if we still can't derive it from the TokenPoolRef, then give up
			tokenRef := input.TokenRef
			if tokenRef.Address == "" {
				if tokRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, tokenRef, chain.Selector, datastore_utils.FullRef); err != nil {
					b.Logger.Warnf("token address could not be resolved using TokenRef (%+v): %v", tokenRef, err)
					b.Logger.Warnf("attempting to resolve token address using TokenPoolRef instead: (%+v)", input.TokenPoolRef)

					tokenPoolRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, input.TokenPoolRef, chain.Selector, datastore_utils.FullRef)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("token pool could not be resolved using TokenPoolRef (%+v): %w", input.TokenPoolRef, err)
					}
					tokenPoolAddrBytes, err := a.AddressRefToBytes(tokenPoolRef)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", err)
					}
					tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
					if tokenPoolAddr == (common.Address{}) {
						return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref (%+v) is zero address", tokenPoolRef)
					}
					tokenAddr, err := a.GetTokenAddressFromFullTokenPoolRef(b, chain, tokenPoolRef)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool ref (%+v): %w", tokenPoolRef, err)
					}

					tokenRef = datastore.AddressRef{
						ChainSelector: chain.Selector,
						Address:       tokenAddr.Hex(),
					}
				} else {
					tokenRef = tokRef
				}
			}

			tokenAddrBytes, err := a.AddressRefToBytes(tokenRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token address ref to bytes: %w", err)
			}
			tokenAddr := common.BytesToAddress(tokenAddrBytes)
			if tokenAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token address for ref (%+v) is zero address", tokenRef)
			}

			proposedOwnerAddrString := input.ProposedOwner
			if !common.IsHexAddress(proposedOwnerAddrString) {
				return sequences.OnChainOutput{}, fmt.Errorf("proposed owner address %q is not a valid hex address", proposedOwnerAddrString)
			}
			proposedOwnerAddr := common.HexToAddress(proposedOwnerAddrString)
			if proposedOwnerAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("proposed owner address cannot be the zero address")
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, chains,
				tarseq.ManualRegistrationSequence,
				tarseq.ManualRegistrationSequenceInput{
					AdminAddress:  proposedOwnerAddr,
					ChainSelector: chain.Selector,
					TokenAddress:  tokenAddr,
					Address:       tokenAdminRegistryAddress,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to manually register token on chain %d: %w", chain.Selector, err)
			}

			return result, nil
		})
}

func (a *EVMAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:deploy-token-pool-for-token",
		cciputils.Version_1_6_0,
		"Deploy a token pool for a token on an EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
			out, err := cldf_ops.ExecuteSequence(b, DeployTokenPool, chains, input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token pool on chain %d: %w", input.ChainSelector, err)
			}

			var result sequences.OnChainOutput
			result.Addresses = append(result.Addresses, out.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)

			toknFilterDS := datastore.AddressRef{ChainSelector: input.ChainSelector}
			if input.TokenRef.Address != "" {
				toknFilterDS.Address = input.TokenRef.Address
			}
			if input.TokenRef.Qualifier != "" {
				toknFilterDS.Qualifier = input.TokenRef.Qualifier
			}
			if input.TokenRef.Type != "" {
				toknFilterDS.Type = input.TokenRef.Type
			}
			toknRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, toknFilterDS, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token address for symbol %q on chain %d: %w", input.TokenRef.Qualifier, input.ChainSelector, err)
			}

			// For a BnM token + BnM token pool, we need to grant the pool mint and burn roles on the token
			isToknTypeBnM := toknRef.Type.String() == bnmERC20ops.ContractType.String()
			isPoolTypeBnM := input.PoolType == cciputils.BurnMintTokenPool.String()
			if isPoolTypeBnM && isToknTypeBnM && len(out.Output.Addresses) >= 1 {
				// NOTE: the pool ref isn't in the datastore yet so we need to fetch it from
				// the DeployTokenPool sequence output. It is assumed that the sequence will
				// return at least one AddressRef{} if the pool was deployed. The 1st ref is
				// assumed to be the token pool ref. If the token pool was already in the DS
				// (i.e. no addresses were returned from the seq) then we skip this step and
				// assume that permissions were already setup.
				poolRef := out.Output.Addresses[0]

				poolAddrBytes, err := a.AddressRefToBytes(poolRef)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to convert deployed token pool address ref to bytes: %w", err)
				}

				toknAddrBytes, err := a.AddressRefToBytes(toknRef)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token address ref to bytes: %w", err)
				}

				poolAddr := common.BytesToAddress(poolAddrBytes)
				if poolAddr == (common.Address{}) {
					return sequences.OnChainOutput{}, errors.New("deployed token pool address is zero address")
				}

				toknAddr := common.BytesToAddress(toknAddrBytes)
				if toknAddr == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("token address for symbol %q is zero address", input.TokenRef.Qualifier)
				}

				chain, ok := chains.EVMChains()[input.ChainSelector]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
				}

				report, err := cldf_ops.ExecuteOperation(b,
					bnmERC20ops.GrantMintAndBurnRoles,
					chain,
					evm_contract.FunctionInput[common.Address]{
						ChainSelector: input.ChainSelector,
						Address:       toknAddr,
						Args:          poolAddr,
					},
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenRef.Qualifier, input.ChainSelector, err)
				}

				batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{report.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for granting mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenRef.Qualifier, input.ChainSelector, err)
				}

				result.BatchOps = append(result.BatchOps, batchOp)
			}

			return result, nil
		},
	)
}

////////////////////
// Helper methods //
////////////////////

func (a *EVMAdapter) GetTokenAdminRegistryAddress(ds datastore.DataStore, selector uint64) (common.Address, error) {
	filters := datastore.AddressRef{
		Type:          datastore.ContractType(tarops.ContractType),
		ChainSelector: selector,
		Version:       tarops.Version,
	}

	ref, err := datastore_utils.FindAndFormatRef(ds, filters, selector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token admin registry address on chain %d: %w", selector, err)
	}

	addr, err := a.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	return common.BytesToAddress(addr), nil
}

func (a *EVMAdapter) FindOneTokenAddress(ds datastore.DataStore, chainSelector uint64, partialRef *datastore.AddressRef) (common.Address, error) {
	filters := datastore.AddressRef{
		ChainSelector: chainSelector,
	}
	if partialRef != nil {
		filters.Address = partialRef.Address
		filters.Qualifier = partialRef.Qualifier
		filters.Type = partialRef.Type
	}

	ref, err := datastore_utils.FindAndFormatRef(ds, filters, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token address for ref %v on chain %d: %w", filters, chainSelector, err)
	}

	addr, err := a.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	return common.BytesToAddress(addr), nil
}

func (a *EVMAdapter) GetTokenAddressFromFullTokenPoolRef(b cldf_ops.Bundle, chain evm.Chain, populatedTokenPoolRef datastore.AddressRef) (common.Address, error) {
	tokenPoolAddressBytes, err := a.AddressRefToBytes(populatedTokenPoolRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", err)
	}

	tokenPoolAddress := common.BytesToAddress(tokenPoolAddressBytes)
	if tokenPoolAddress == (common.Address{}) {
		return common.Address{}, fmt.Errorf("token pool address for ref (%+v) is zero address", populatedTokenPoolRef)
	}

	switch populatedTokenPoolRef.Version.String() {
	case tpOpsV1_5_1.Version.String():
		if res, err := cldf_ops.ExecuteOperation(b,
			tpOpsV1_5_1.GetToken, chain,
			evm_contract.FunctionInput[struct{}]{
				ChainSelector: populatedTokenPoolRef.ChainSelector,
				Address:       tokenPoolAddress,
				Args:          struct{}{},
			},
		); err != nil {
			return common.Address{}, fmt.Errorf("failed to get token address via GetToken operation (version=%s): %w", tpOpsV1_5_1.Version.String(), err)
		} else {
			return res.Output, nil
		}
	case tpOpsV1_6_1.Version.String():
		if res, err := cldf_ops.ExecuteOperation(b,
			tpOpsV1_6_1.GetToken, chain,
			evm_contract.FunctionInput[struct{}]{
				ChainSelector: populatedTokenPoolRef.ChainSelector,
				Address:       tokenPoolAddress,
				Args:          struct{}{},
			}); err != nil {
			return common.Address{}, fmt.Errorf("failed to get token address via GetToken operation (version=%s): %w", tpOpsV1_6_1.Version.String(), err)
		} else {
			return res.Output, nil
		}
	default:
		return common.Address{}, fmt.Errorf(
			"unsupported token pool version %s for token pool at address %q on chain selector %d",
			populatedTokenPoolRef.Version.String(), tokenPoolAddress.Hex(), populatedTokenPoolRef.ChainSelector,
		)
	}
}

func (a *EVMAdapter) GetTokenPoolAdmins(ctx context.Context, chain *evm.Chain, ref datastore.AddressRef) (poolOwner common.Address, rlAdmin common.Address, err error) {
	type TokenPoolAdminReader interface {
		GetRateLimitAdmin(*bind.CallOpts) (common.Address, error)
		Owner(*bind.CallOpts) (common.Address, error)
	}

	addrBytes, err := a.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	addr := common.BytesToAddress(addrBytes)
	if addr == (common.Address{}) {
		return common.Address{}, common.Address{}, fmt.Errorf("token pool address for ref (%+v) is zero address", ref)
	}

	var pool TokenPoolAdminReader
	switch ref.Version.String() {
	case tpOpsV1_5_1.Version.String():
		if tp, err := tpV1_5_1.NewTokenPool(addr, chain.Client); err != nil {
			return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
		} else {
			pool = tp
		}
	case tpOpsV1_6_1.Version.String():
		if tp, err := tpV1_6_1.NewTokenPool(addr, chain.Client); err != nil {
			return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v1.6.1 contract: %w", err)
		} else {
			pool = tp
		}
	default:
		return common.Address{}, common.Address{}, fmt.Errorf("unsupported token pool version %s for token pool at address %q on chain selector %d", ref.Version.String(), addr.Hex(), ref.ChainSelector)
	}

	poolOwner, err = pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at address %q on chain %d: %w", addr.Hex(), chain.Selector, err)
	}
	rlAdmin, err = pool.GetRateLimitAdmin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get rate limit admin of token pool at address %q on chain %d: %w", addr.Hex(), chain.Selector, err)
	}

	return poolOwner, rlAdmin, nil
}

func (a *EVMAdapter) FindLatestAddressRef(ds datastore.DataStore, ref datastore.AddressRef) (common.Address, error) {
	// Define the version range
	minVersion := semver.MustParse("1.5.0") // inclusive
	maxVersion := semver.MustParse("2.0.0") // exclusive

	// Build the filter
	filter := []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]{}
	if ref.ChainSelector != 0 {
		filter = append(filter, datastore.AddressRefByChainSelector(ref.ChainSelector))
	}
	if ref.Qualifier != "" {
		filter = append(filter, datastore.AddressRefByQualifier(ref.Qualifier))
	}
	if ref.Version != nil {
		// NOTE: this shouldn't be set otherwise we won't be able to find the latest version within the specified range
		return common.Address{}, fmt.Errorf("ref version should not be set when finding the latest address ref, got version %s", ref.Version.String())
	}
	if ref.Address != "" {
		// NOTE: this shouldn't be set otherwise we'd always get zero or one result back, which defeats this function's purpose
		return common.Address{}, fmt.Errorf("ref address should not be set when finding the latest address ref, got address %q", ref.Address)
	}
	if ref.Type.String() != "" {
		filter = append(filter, datastore.AddressRefByType(ref.Type))
	}

	// Get all matching token pool addresses
	refs := ds.Addresses().Filter(filter...)

	// Use the latest version found within the specified range.
	var latestRef datastore.AddressRef
	latestVer := minVersion
	doesExist := false
	for _, ref := range refs {
		v := ref.Version
		if v == nil {
			continue
		}

		isInside := v.GreaterThanEqual(minVersion) && v.LessThan(maxVersion)
		isBetter := !doesExist || v.GreaterThan(latestVer)
		if isInside && isBetter {
			doesExist = true
			latestRef = ref
			latestVer = v
		}
	}

	// If no matching reference was found, then return an error
	if !doesExist {
		return common.Address{}, fmt.Errorf("no address found for ref (%+v) in version range [%s, %s)", ref, minVersion.String(), maxVersion.String())
	}

	// Convert the address reference to bytes
	addrBytes, err := a.AddressRefToBytes(latestRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	// Construct the token pool instance
	return common.BytesToAddress(addrBytes), nil
}

