package adapters

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	evm_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"

	tpBindingsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"

	bnmOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmERC677Ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
)

var (
	_ tokens.TokenFeeAdapter = &TokenAdapter{}
	_ tokens.TokenAdapter    = &TokenAdapter{}
)

// TokenAdapter handles EVM token pools at version 2.0.0.
// It embeds EVMPoolAdapter for shared methods (DeriveTokenAddress,
// ManualRegistration) and overrides the methods that have genuinely
// different v2.0.0 logic (DeriveTokenDecimals with ERC20 fallback,
// SetTokenPoolRateLimits with batched default + fast-finality TPRL buckets,
// DeployTokenPoolForToken with its own deploy sequences).
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a v2.0.0 TokenAdapter with pre-wired PoolOps.
// DeployTokenPoolSeq is nil because DeployTokenPoolForToken is fully overridden.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops: &poolOpsV200{},
		},
	}
}

func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.ConfigureTokenForTransfers
}

func (t *TokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-2.0-adapter:deploy-token-pool-for-token",
		cciputils.Version_2_0_0,
		"Deploy a 2.0.0 token pool for a token on an EVM chain",
		func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			if input.TokenPoolVersion == nil {
				return sequences.OnChainOutput{}, errors.New("TokenPoolVersion is required")
			}

			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}

			threshold := big.NewInt(0)
			thresholdProvided := input.ThresholdAmountForAdditionalCCVs != ""
			if thresholdProvided {
				var ok bool
				threshold, ok = new(big.Int).SetString(input.ThresholdAmountForAdditionalCCVs, 10)
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid ThresholdAmountForAdditionalCCVs %q: must be a decimal integer string", input.ThresholdAmountForAdditionalCCVs)
				}
			}

			var rateLimitAdmin common.Address
			if input.RateLimitAdmin != "" {
				if !common.IsHexAddress(input.RateLimitAdmin) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid RateLimitAdmin address %q", input.RateLimitAdmin)
				}
				rateLimitAdmin = common.HexToAddress(input.RateLimitAdmin)
			}

			var feeAggregator common.Address
			if input.FeeAggregator != "" {
				if !common.IsHexAddress(input.FeeAggregator) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeAggregator address %q", input.FeeAggregator)
				}
				feeAggregator = common.HexToAddress(input.FeeAggregator)
			}

			var tokenAddr string
			if input.TokenRef != nil && input.TokenRef.Address != "" {
				tokenAddr = input.TokenRef.Address
			}
			if input.TokenRef != nil && input.TokenRef.Qualifier != "" {
				storedRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, *input.TokenRef, input.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("token with ref %+v not found in datastore: %w", *input.TokenRef, err)
				}
				if tokenAddr != "" && storedRef.Address != tokenAddr {
					return sequences.OnChainOutput{}, fmt.Errorf("provided token address %q does not match datastore address %q", tokenAddr, storedRef.Address)
				}
				if tokenAddr == "" {
					tokenAddr = storedRef.Address
				}
			}
			if tokenAddr == "" {
				return sequences.OnChainOutput{}, errors.New("token address must be provided either directly or via a datastore reference")
			}

			qualifier := input.TokenPoolQualifier
			if qualifier == "" {
				qualifier = tokenAddr
			}
			poolType := deployment.ContractType(input.PoolType)

			grantMintBurnRoles := func(poolRef datastore.AddressRef) (*mcms_types.BatchOperation, error) {
				if !t.EVMTokenBase.IsBurnMintPoolType(poolType.String()) {
					return nil, nil
				}

				tokenRef, lookupErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Address:       tokenAddr,
				}, input.ChainSelector, datastore_utils.FullRef)
				if lookupErr != nil || !t.EVMTokenBase.IsBurnMintTokenType(tokenRef.Type.String()) {
					return nil, nil
				}

				poolAddr := common.HexToAddress(poolRef.Address)
				if poolAddr == (common.Address{}) {
					return nil, errors.New("token pool address is zero")
				}

				grantInput := contract.FunctionInput[common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(tokenAddr),
					Args:          poolAddr,
				}
				var writes []contract.WriteOutput
				if t.EVMTokenBase.IsBurnMintERC677TokenType(tokenRef.Type.String()) {
					var grantErr error
					writes, grantErr = bnmERC677Ops.PrepareGrantMintAndBurnRoles(
						b,
						evmChain,
						grantInput,
						common.HexToAddress(input.TimelockAddress),
					)
					if grantErr != nil {
						return nil, fmt.Errorf("failed to grant mint/burn roles to pool %s for token %s: %w", poolAddr, tokenAddr, grantErr)
					}
				} else {
					grantReport, grantErr := cldf_ops.ExecuteOperation(b,
						bnmOps.GrantMintAndBurnRoles, evmChain, grantInput)
					if grantErr != nil {
						return nil, fmt.Errorf("failed to grant mint/burn roles to pool %s for token %s: %w", poolAddr, tokenAddr, grantErr)
					}
					writes = append(writes, grantReport.Output)
				}

				batchOp, bErr := contract.NewBatchOperationFromWrites(writes)
				if bErr != nil {
					return nil, fmt.Errorf("failed to create batch operation for role grants: %w", bErr)
				}
				return &batchOp, nil
			}

			matches := input.ExistingDataStore.Addresses().Filter(
				datastore.AddressRefByType(datastore.ContractType(input.PoolType)),
				datastore.AddressRefByChainSelector(input.ChainSelector),
				datastore.AddressRefByQualifier(qualifier),
				datastore.AddressRefByVersion(input.TokenPoolVersion),
			)
			if len(matches) > 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("multiple token pools found in datastore with type %q, version %q, qualifier %q on chain %d",
					input.PoolType, input.TokenPoolVersion, qualifier, input.ChainSelector)
			}
			if len(matches) == 1 {
				b.Logger.Info("Token pool already deployed at address:", matches[0].Address)
				// A previous partial run can leave the pool in datastore before
				// the token grants it burn/mint rights. Keep DeployTokenPoolForToken
				// declarative: after it runs, the token/pool authority relationship
				// should be correct whether the pool was just deployed or reused.
				var result sequences.OnChainOutput
				batchOp, err := grantMintBurnRoles(matches[0])
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				if batchOp != nil {
					result.BatchOps = append(result.BatchOps, *batchOp)
				}

				// Reconcile any dynamic-config fields the caller explicitly supplied
				// (router, rate-limit admin, fee aggregator, additional-CCVs
				// threshold). ConfigureTokenPool reads current values and only
				// emits a write when they differ, so re-runs with the same inputs
				// are no-ops. Fields the caller leaves unset (zero/empty) retain
				// their current on-chain values.
				if input.RouterRef != nil || rateLimitAdmin != (common.Address{}) || feeAggregator != (common.Address{}) || thresholdProvided {
					poolAddr := common.HexToAddress(matches[0].Address)
					configureInput := evm_tokens.ConfigureTokenPoolInput{
						ChainSelector:    input.ChainSelector,
						TokenPoolAddress: poolAddr,
						RateLimitAdmin:   rateLimitAdmin,
						FeeAggregator:    feeAggregator,
					}
					if input.RouterRef != nil {
						resolved, err := t.EVMTokenBase.ResolveRouterAddress(input.ExistingDataStore, input.ChainSelector, input.RouterRef)
						if err != nil {
							return sequences.OnChainOutput{}, err
						}
						configureInput.RouterAddress = resolved
					}
					if thresholdProvided {
						hooksReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAdvancedPoolHooks, evmChain, contract.FunctionInput[struct{}]{
							ChainSelector: input.ChainSelector,
							Address:       poolAddr,
						})
						if err != nil {
							return sequences.OnChainOutput{}, fmt.Errorf("failed to read advanced pool hooks address from existing token pool %s on chain %d: %w", poolAddr, input.ChainSelector, err)
						}
						configureInput.AdvancedPoolHooks = hooksReport.Output
						configureInput.ThresholdAmountForAdditionalCCVs = threshold
					}
					configureReport, err := cldf_ops.ExecuteSequence(b, evm_tokens.ConfigureTokenPool, evmChain, configureInput)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to reconcile dynamic config for existing token pool %s on chain %d: %w", poolAddr, input.ChainSelector, err)
					}
					result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)
				}

				return result, nil
			}

			tokenContract, err := erc20.NewERC20(common.HexToAddress(tokenAddr), evmChain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to bind ERC20 at %s: %w", tokenAddr, err)
			}
			tokenDecimals, err := tokenContract.Decimals(&bind.CallOpts{Context: b.GetContext()})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get decimals for token at %s: %w", tokenAddr, err)
			}

			resolvedRouter, err := t.EVMTokenBase.ResolveRouterAddress(input.ExistingDataStore, input.ChainSelector, input.RouterRef)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			rmnProxyRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(rmnproxyops.ContractType),
			}, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find RMN proxy in datastore for chain %d: %w", input.ChainSelector, err)
			}

			var allowlist []common.Address
			for _, addr := range input.Allowlist {
				if !common.IsHexAddress(addr) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid allowlist address: %s", addr)
				}
				allowlist = append(allowlist, common.HexToAddress(addr))
			}

			internalInput := evm_tokens.DeployTokenPoolInput{
				ChainSel:                         input.ChainSelector,
				TokenPoolType:                    datastore.ContractType(input.PoolType),
				TokenPoolVersion:                 input.TokenPoolVersion,
				TokenSymbol:                      qualifier,
				RateLimitAdmin:                   rateLimitAdmin,
				FeeAggregator:                    feeAggregator,
				ThresholdAmountForAdditionalCCVs: threshold,
				ConstructorArgs: evm_tokens.ConstructorArgs{
					Token:    common.HexToAddress(tokenAddr),
					Decimals: tokenDecimals,
					RMNProxy: common.HexToAddress(rmnProxyRef.Address),
					Router:   resolvedRouter,
				},
				AdvancedPoolHooksConfig: evm_tokens.AdvancedPoolHooksConfig{
					Allowlist: allowlist,
				},
			}

			var deployOutput sequences.OnChainOutput

			switch {
			case t.EVMTokenBase.IsBurnMintPoolType(poolType.String()):
				report, execErr := cldf_ops.ExecuteSequence(b, evm_tokens.DeployBurnMintTokenPool, evmChain, internalInput)
				if execErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token pool on chain %d: %w", input.ChainSelector, execErr)
				}
				deployOutput = report.Output
			case t.EVMTokenBase.IsLockReleasePoolType(poolType.String()):
				report, execErr := cldf_ops.ExecuteSequence(b, evm_tokens.DeployLockReleaseTokenPool, evmChain, internalInput)
				if execErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy lock release token pool on chain %d: %w", input.ChainSelector, execErr)
				}
				deployOutput = report.Output
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported 2.0.0 token pool type: %s", input.PoolType)
			}

			var result sequences.OnChainOutput
			result.Addresses = append(result.Addresses, deployOutput.Addresses...)
			result.BatchOps = append(result.BatchOps, deployOutput.BatchOps...)

			if t.EVMTokenBase.IsBurnMintPoolType(poolType.String()) && len(deployOutput.Addresses) >= 1 {
				batchOp, err := grantMintBurnRoles(deployOutput.Addresses[0])
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				if batchOp != nil {
					result.BatchOps = append(result.BatchOps, *batchOp)
				}
			}

			return result, nil
		},
	)
}

func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokens.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.MigrateLockReleasePoolLiquidity
}

func (t *TokenAdapter) SetAllowedFinalityConfig(e *deployment.Environment) *cldf_ops.Sequence[tokens.SetAllowedFinalityConfigSequenceInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.SetAllowedFinalityConfigForTokenPools
}

func (t *TokenAdapter) SetTokenTransferFee(e *deployment.Environment) *cldf_ops.Sequence[tokens.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.SetTokenTransferFeeConfigForTokenPools
}

func (t *TokenAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) tokens.TokenTransferFeeConfig {
	return tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(src, dst)
}

func (t *TokenAdapter) GetOnchainTokenTransferFeeConfig(e deployment.Environment, poolAddress string, src uint64, dst uint64) (tokens.TokenTransferFeeConfig, error) {
	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return tokens.TokenTransferFeeConfig{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	args := token_pool.GetTokenTransferFeeConfigArgs{DestChainSelector: dst}
	if !common.IsHexAddress(poolAddress) {
		return tokens.TokenTransferFeeConfig{}, fmt.Errorf("invalid pool address: %s", poolAddress)
	}

	addr := common.HexToAddress(poolAddress)
	if addr == (common.Address{}) {
		return tokens.TokenTransferFeeConfig{}, errors.New("pool address cannot be the zero address")
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		token_pool.GetTokenTransferFeeConfig,
		chain,
		contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
			ChainSelector: src,
			Address:       addr,
			Args:          args,
		},
	)
	if err != nil {
		return tokens.TokenTransferFeeConfig{}, fmt.Errorf("failed to get on-chain token transfer fee config for pool %s on chain selector %d for dest chain selector %d: %w", poolAddress, src, dst, err)
	}

	return tokens.TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: report.Output.FinalityTransferFeeBps,
		CustomFinalityTransferFeeBps:  report.Output.FastFinalityTransferFeeBps,
		DefaultFinalityFeeUSDCents:    report.Output.FinalityFeeUSDCents,
		CustomFinalityFeeUSDCents:     report.Output.FastFinalityFeeUSDCents,
		DestBytesOverhead:             report.Output.DestBytesOverhead,
		DestGasOverhead:               report.Output.DestGasOverhead,
		IsEnabled:                     report.Output.IsEnabled,
	}, nil
}

// poolOpsV200 implements PoolOps using v2.0.0 token pool bindings.
// Only GetToken and Version are called at runtime (by the inherited
// DeriveTokenAddress and ManualRegistration); the other methods are
// stubs because the v2.0.0 adapter overrides the methods that use them.
type poolOpsV200 struct{}

func (p *poolOpsV200) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		token_pool.GetToken, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v2.0.0: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV200) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		token_pool.GetTokenDecimals, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v2.0.0: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV200) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (common.Address, common.Address, error) {
	pool, err := token_pool.NewTokenPoolContract(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v2.0.0 contract at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	owner, err := pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	cfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get dynamic config of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	return owner, cfg.RateLimitAdmin, nil
}

func (p *poolOpsV200) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokens.TPRLRemotes) ([]contract.WriteOutput, error) {
	var writes []contract.WriteOutput
	if !input.AllowedFinalityConfig.IsZero() {
		currentFinalityConfig, err := cldf_ops.ExecuteOperation(b, token_pool.GetAllowedFinalityConfig, chain, contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          struct{}{},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get allowed finality config for token pool at %s on chain %d: %w", poolAddr.Hex(), input.ChainSelector, err)
		}
		if input.AllowedFinalityConfig.Raw() != currentFinalityConfig.Output {
			setFinalityReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetAllowedFinalityConfig, chain, contract.FunctionInput[[4]byte]{
				ChainSelector: input.ChainSelector,
				Address:       poolAddr,
				Args:          input.AllowedFinalityConfig.Raw(),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to set allowed finality config on token pool at %s on chain %d: %w", poolAddr.Hex(), input.ChainSelector, err)
			}
			writes = append(writes, setFinalityReport.Output)
		}
	}

	args := make([]token_pool.RateLimitConfigArgs, 0, len(input.RateLimitBuckets))
	for _, bucket := range input.RateLimitBuckets {
		args = append(args, token_pool.RateLimitConfigArgs{
			RemoteChainSelector: input.RemoteChainSelector,
			FastFinality:        bucket.FastFinality,
			OutboundRateLimiterConfig: token_pool.Config{
				IsEnabled: bucket.OutboundRateLimiterConfig.IsEnabled,
				Capacity:  bucket.OutboundRateLimiterConfig.Capacity,
				Rate:      bucket.OutboundRateLimiterConfig.Rate,
			},
			InboundRateLimiterConfig: token_pool.Config{
				IsEnabled: bucket.InboundRateLimiterConfig.IsEnabled,
				Capacity:  bucket.InboundRateLimiterConfig.Capacity,
				Rate:      bucket.InboundRateLimiterConfig.Rate,
			},
		})
	}

	if len(args) > 0 {
		rateLimitsReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, chain, contract.FunctionInput[[]token_pool.RateLimitConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          args,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to set rate limit config on pool %s: %w", poolAddr, err)
		}
		writes = append(writes, rateLimitsReport.Output)
	}

	if len(writes) == 0 {
		return nil, nil
	}

	return writes, nil
}

func (p *poolOpsV200) SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) (contract.WriteOutput, error) {
	pool, err := token_pool.NewTokenPoolContract(poolAddr, chain.Client)
	if err != nil {
		return contract.WriteOutput{}, fmt.Errorf("failed to instantiate token pool v2.0.0 contract at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	cfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: b.GetContext()})
	if err != nil {
		return contract.WriteOutput{}, fmt.Errorf("failed to get dynamic config of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	if newAdmin == cfg.RateLimitAdmin {
		b.Logger.Info("Rate limit admin is already set to the desired address; no update needed")
		return contract.WriteOutput{}, nil
	}

	res, err := cldf_ops.ExecuteOperation(b,
		token_pool.SetDynamicConfig, chain,
		contract.FunctionInput[token_pool.SetDynamicConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: token_pool.SetDynamicConfigArgs{
				RateLimitAdmin: newAdmin,
				FeeAdmin:       cfg.FeeAdmin,
				Router:         cfg.Router,
			},
		},
	)
	if err != nil {
		return contract.WriteOutput{}, fmt.Errorf("SetDynamicConfig v2.0.0: %w", err)
	}

	return res.Output, nil
}

func (p *poolOpsV200) GetCurrentInboundRateLimit(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, ff bool) (tokens.RateLimiterConfig, error) {
	// Call the contract binding directly rather than cldf_ops Read: the framework caches read
	// reports by input hash, and earlier sequences in the same Apply run may have read this
	// same lane while it was still uninitialized — caching that stale result.
	tp, err := tpBindingsV2_0_0.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return tokens.RateLimiterConfig{}, fmt.Errorf("failed to instantiate v2.0.0 token pool contract at %s: %w", poolAddr.Hex(), err)
	}
	state, err := tp.GetCurrentRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, remoteSelector, ff)
	if err != nil {
		return tokens.RateLimiterConfig{}, fmt.Errorf("failed to get inbound rate limiter state for remote chain %d (fastFinality=%v): %w", remoteSelector, ff, err)
	}
	return tokens.RateLimiterConfig{
		IsEnabled: state.InboundRateLimiterState.IsEnabled,
		Capacity:  state.InboundRateLimiterState.Capacity,
		Rate:      state.InboundRateLimiterState.Rate,
	}, nil
}

func (p *poolOpsV200) Version() *semver.Version {
	return cciputils.Version_2_0_0
}
