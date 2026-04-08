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

	evm_tokens "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"

	bnmOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmDripOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"

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
// SetTokenPoolRateLimits with batch/custom finality, DeployTokenPoolForToken
// with its own deploy sequences).
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

// ConfigureTokenForTransfersSequence returns the sequence for configuring an EVM token with a 2.0.0 token pool.
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
			if input.ThresholdAmountForAdditionalCCVs != "" {
				var ok bool
				threshold, ok = new(big.Int).SetString(input.ThresholdAmountForAdditionalCCVs, 10)
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid ThresholdAmountForAdditionalCCVs %q: must be a decimal integer string", input.ThresholdAmountForAdditionalCCVs)
				}
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
				return sequences.OnChainOutput{}, nil
			}

			tokenContract, err := erc20.NewERC20(common.HexToAddress(tokenAddr), evmChain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to bind ERC20 at %s: %w", tokenAddr, err)
			}
			tokenDecimals, err := tokenContract.Decimals(&bind.CallOpts{Context: b.GetContext()})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get decimals for token at %s: %w", tokenAddr, err)
			}

			routerRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(router.ContractType),
			}, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find router in datastore for chain %d: %w", input.ChainSelector, err)
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
				ThresholdAmountForAdditionalCCVs: threshold,
				ConstructorArgs: evm_tokens.ConstructorArgs{
					Token:    common.HexToAddress(tokenAddr),
					Decimals: tokenDecimals,
					RMNProxy: common.HexToAddress(rmnProxyRef.Address),
					Router:   common.HexToAddress(routerRef.Address),
				},
				AdvancedPoolHooksConfig: evm_tokens.AdvancedPoolHooksConfig{
					Allowlist: allowlist,
				},
			}

			poolType := deployment.ContractType(input.PoolType)
			var deployOutput sequences.OnChainOutput

			switch {
			case isBurnMintPoolType(poolType):
				report, execErr := cldf_ops.ExecuteSequence(b, evm_tokens.DeployBurnMintTokenPool, evmChain, internalInput)
				if execErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token pool on chain %d: %w", input.ChainSelector, execErr)
				}
				deployOutput = report.Output
			case isLockReleasePoolType(poolType):
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

			if isBurnMintPoolType(poolType) && len(deployOutput.Addresses) >= 1 {
				toknRef, lookupErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Address:       tokenAddr,
				}, input.ChainSelector, datastore_utils.FullRef)
				if lookupErr == nil && isBurnMintTokenType(toknRef.Type) {
					poolRef := deployOutput.Addresses[0]
					poolAddr := common.HexToAddress(poolRef.Address)
					if poolAddr == (common.Address{}) {
						return sequences.OnChainOutput{}, errors.New("deployed token pool address is zero")
					}

					grantReport, grantErr := cldf_ops.ExecuteOperation(b,
						bnmOps.GrantMintAndBurnRoles, evmChain,
						contract.FunctionInput[common.Address]{
							ChainSelector: input.ChainSelector,
							Address:       common.HexToAddress(tokenAddr),
							Args:          poolAddr,
						},
					)
					if grantErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint/burn roles to pool %s for token %s: %w", poolAddr, tokenAddr, grantErr)
					}

					batchOp, bErr := contract.NewBatchOperationFromWrites([]contract.WriteOutput{grantReport.Output})
					if bErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for role grants: %w", bErr)
					}
					result.BatchOps = append(result.BatchOps, batchOp)
				}
			}

			return result, nil
		},
	)
}

// DeriveTokenDecimals has v2.0.0-specific logic: it falls back to ERC20.Decimals()
// when the pool's GetTokenDecimals fails (e.g., proxy pools like USDCTokenPoolProxy).
func (t *TokenAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, tokenBytes []byte) (uint8, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	getTokenDecimalsReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetTokenDecimals, evmChain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(poolRef.Address),
	})
	if err == nil {
		return getTokenDecimalsReport.Output, nil
	}
	poolErr := err

	tokenAddr := common.BytesToAddress(tokenBytes)
	if tokenAddr.Cmp(common.Address{}) == 0 {
		getTokenReport, getTokErr := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, evmChain, contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       common.HexToAddress(poolRef.Address),
		})
		if getTokErr != nil {
			return 0, fmt.Errorf("failed to get token decimals from token pool with address %s on %s: %w", poolRef.Address, evmChain, poolErr)
		}
		tokenAddr = getTokenReport.Output
	}

	tokenContract, newErr := erc20.NewERC20(tokenAddr, evmChain.Client)
	if newErr != nil {
		return 0, fmt.Errorf("failed to get token decimals from token pool with address %s on %s: %w; failed to bind erc20 at token %s: %w", poolRef.Address, evmChain, poolErr, tokenAddr.Hex(), newErr)
	}
	decimals, erc20Err := tokenContract.Decimals(&bind.CallOpts{Context: e.GetContext()})
	if erc20Err != nil {
		return 0, fmt.Errorf("failed to get token decimals from token pool with address %s on %s: %w; erc20.decimals on token %s also failed: %w", poolRef.Address, evmChain, poolErr, tokenAddr.Hex(), erc20Err)
	}
	return decimals, nil
}

// SetTokenPoolRateLimits has v2.0.0-specific logic: batch call with both
// default and custom finality rate limits.
func (t *TokenAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-2.0-adapter:set-token-pool-rate-limits",
		cciputils.Version_2_0_0,
		"Set rate limits for a 2.0.0 token pool on an EVM chain",
		func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.TPRLRemotes) (sequences.OnChainOutput, error) {
			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}

			tokenPoolAddrBytes, err := t.AddressRefToBytes(input.TokenPoolRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref: %w", err)
			}
			tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
			if tokenPoolAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref %+v is zero", input.TokenPoolRef)
			}

			args := []token_pool.RateLimitConfigArgs{
				{
					RemoteChainSelector:      input.RemoteChainSelector,
					FastFinality: false,
					OutboundRateLimiterConfig: token_pool.Config{
						IsEnabled: input.DefaultFinalityOutboundRateLimiterConfig.IsEnabled,
						Capacity:  input.DefaultFinalityOutboundRateLimiterConfig.Capacity,
						Rate:      input.DefaultFinalityOutboundRateLimiterConfig.Rate,
					},
					InboundRateLimiterConfig: token_pool.Config{
						IsEnabled: input.DefaultFinalityInboundRateLimiterConfig.IsEnabled,
						Capacity:  input.DefaultFinalityInboundRateLimiterConfig.Capacity,
						Rate:      input.DefaultFinalityInboundRateLimiterConfig.Rate,
					},
				},
				{
					RemoteChainSelector:      input.RemoteChainSelector,
					FastFinality: true,
					OutboundRateLimiterConfig: token_pool.Config{
						IsEnabled: input.CustomFinalityOutboundRateLimiterConfig.IsEnabled,
						Capacity:  input.CustomFinalityOutboundRateLimiterConfig.Capacity,
						Rate:      input.CustomFinalityOutboundRateLimiterConfig.Rate,
					},
					InboundRateLimiterConfig: token_pool.Config{
						IsEnabled: input.CustomFinalityInboundRateLimiterConfig.IsEnabled,
						Capacity:  input.CustomFinalityInboundRateLimiterConfig.Capacity,
						Rate:      input.CustomFinalityInboundRateLimiterConfig.Rate,
					},
				},
			}

			report, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, evmChain, contract.FunctionInput[[]token_pool.RateLimitConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddr,
				Args:          args,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limit config on pool %s: %w", tokenPoolAddr, err)
			}

			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
		})
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

func (p *poolOpsV200) GetToken(b cldf_ops.Bundle, ch evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		token_pool.GetToken, ch,
		contract.FunctionInput[struct{}]{
			ChainSelector: ch.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v2.0.0: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV200) GetTokenDecimals(_ context.Context, _ evm.Chain, _ common.Address) (uint8, error) {
	return 0, errors.New("poolOpsV200.GetTokenDecimals: not used; v2.0.0 adapter overrides DeriveTokenDecimals")
}

func (p *poolOpsV200) GetPoolAdmins(_ context.Context, _ *evm.Chain, _ common.Address) (common.Address, common.Address, error) {
	return common.Address{}, common.Address{}, errors.New("poolOpsV200.GetPoolAdmins: not used; v2.0.0 adapter overrides SetTokenPoolRateLimits")
}

func (p *poolOpsV200) SetRateLimiterConfig(_ cldf_ops.Bundle, _ evm.Chain, _ common.Address, _ uint64, _, _ tokens.RateLimiterConfig) (contract.WriteOutput, error) {
	return contract.WriteOutput{}, errors.New("poolOpsV200.SetRateLimiterConfig: not used; v2.0.0 adapter overrides SetTokenPoolRateLimits")
}

func (p *poolOpsV200) Version() *semver.Version {
	return cciputils.Version_2_0_0
}

func isBurnMintPoolType(poolType deployment.ContractType) bool {
	return poolType == cciputils.BurnMintTokenPool ||
		poolType == cciputils.BurnFromMintTokenPool ||
		poolType == cciputils.BurnWithFromMintTokenPool
}

func isLockReleasePoolType(poolType deployment.ContractType) bool {
	return poolType == cciputils.LockReleaseTokenPool ||
		poolType == siloed_lock_release_token_pool.ContractType
}

func isBurnMintTokenType(typ datastore.ContractType) bool {
	return typ.String() == bnmOps.ContractType.String() || typ.String() == bnmDripOps.ContractType.String()
}
