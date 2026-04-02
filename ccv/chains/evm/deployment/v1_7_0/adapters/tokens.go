package adapters

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_tokens "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/token_pool"
	evmutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnmOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmDripOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	evm16seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"

	ownershipAdapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
)

var _ tokens.TokenAdapter = &TokenAdapter{}

// TokenAdapter is the adapter for EVM tokens using 2.0.0 token pools.
type TokenAdapter struct{}

// ConfigureTokenForTransfersSequence returns the sequence for configuring an EVM token with a 2.0.0 token pool.
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.ConfigureTokenForTransfers
}

// AddressRefToBytes returns an EVM address reference as an EVM address.
func (t *TokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}

// DeriveTokenAddress derives the token address from a token pool reference, returning it as an EVM address.
func (t *TokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	getTokenReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, evmChain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(poolRef.Address),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", poolRef.Address, evmChain, err)
	}

	return t.AddressRefToBytes(datastore.AddressRef{
		Address: getTokenReport.Output.Hex(),
	})
}

func (t *TokenAdapter) DeployToken() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm16seq.DeployToken
}

func (t *TokenAdapter) DeployTokenVerify(e deployment.Environment, input tokens.DeployTokenInput) error {
	tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
		ChainSelector: input.ChainSelector,
		Type:          datastore.ContractType(input.Type),
		Qualifier:     input.Symbol,
	}, input.ChainSelector, datastore_utils.FullRef)
	if err == nil {
		e.OperationsBundle.Logger.Info("Token already deployed at address:", tokenAddr.Address)
		return nil
	}

	if err := evmutils.ValidateEVMAddress(input.CCIPAdmin, "CCIPAdmin"); err != nil {
		return err
	}
	if err := evmutils.ValidateEVMAddress(input.ExternalAdmin, "ExternalAdmin"); err != nil {
		return err
	}

	if input.Decimals > 18 {
		return fmt.Errorf("EVM tokens cannot have more than 18 decimals, got %d", input.Decimals)
	}

	if input.PreMint != nil && input.Supply != nil && *input.Supply != 0 && *input.PreMint > *input.Supply {
		return fmt.Errorf("pre-mint amount cannot be greater than max supply, got pre-mint %d and supply %d", *input.PreMint, *input.Supply)
	}

	return nil
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

			// Parse threshold for additional CCVs
			threshold := big.NewInt(0)
			if input.ThresholdForAdditionalCCVs != "" {
				var ok bool
				threshold, ok = new(big.Int).SetString(input.ThresholdForAdditionalCCVs, 10)
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid ThresholdForAdditionalCCVs %q: must be a decimal integer string", input.ThresholdForAdditionalCCVs)
				}
			}

			// Resolve token address
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

			// Skip if pool already exists in datastore
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

			// Get token decimals
			tokenContract, err := erc20.NewERC20(common.HexToAddress(tokenAddr), evmChain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to bind ERC20 at %s: %w", tokenAddr, err)
			}
			tokenDecimals, err := tokenContract.Decimals(&bind.CallOpts{Context: b.GetContext()})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get decimals for token at %s: %w", tokenAddr, err)
			}

			// Find router and RMN proxy from datastore
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

			// Convert string allowlist to common.Address
			var allowlist []common.Address
			for _, addr := range input.Allowlist {
				allowlist = append(allowlist, common.HexToAddress(addr))
			}

			// Build internal deploy input
			internalInput := evm_tokens.DeployTokenPoolInput{
				ChainSel:         input.ChainSelector,
				TokenPoolType:    datastore.ContractType(input.PoolType),
				TokenPoolVersion: input.TokenPoolVersion,
				TokenSymbol:      qualifier,
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

			// Dispatch based on pool type
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

			// Grant mint/burn roles if the token is a BurnMint type and the pool is BurnMint
			if isBurnMintPoolType(poolType) && len(deployOutput.Addresses) >= 1 {
				toknRef, lookupErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Address:       tokenAddr,
				}, input.ChainSelector, datastore_utils.FullRef)
				if lookupErr == nil && isBurnMintTokenType(toknRef.Type) {
					// The first address in pool deploy output is the token pool
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

func (t *TokenAdapter) SetPool() *cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
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

	// Fallback to fetch token's decimals directly, this is useful for proxy pools like USDCTokenPoolProxy which does not expose the GetTokenDecimals function.
	tokenAddr := common.BytesToAddress(token)
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

func (t *TokenAdapter) DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
	return tokenPool, nil
}

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

			// 2.0.0 pools support separate default/custom finality rate limits via a batch call
			args := []token_pool.RateLimitConfigArgs{
				{
					RemoteChainSelector:      input.RemoteChainSelector,
					CustomBlockConfirmations: false,
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
					CustomBlockConfirmations: true,
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

func (t *TokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokens.ManualRegistrationSequenceInput, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-2.0-adapter:manual-registration",
		cciputils.Version_2_0_0,
		"Manually register a token on the TokenAdminRegistry for EVM with 2.0.0 token pools",
		func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.ManualRegistrationSequenceInput) (sequences.OnChainOutput, error) {
			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}

			// Find token admin registry
			tarRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				Type:          datastore.ContractType(tarops.ContractType),
				ChainSelector: evmChain.Selector,
				Version:       tarops.Version,
			}, evmChain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find TAR on chain %d: %w", evmChain.Selector, err)
			}
			tarAddr := common.HexToAddress(tarRef.Address)

			// Resolve token address from TokenRef, with fallback to pool
			tokenRef := input.TokenRef
			if tokenRef.Address == "" {
				if found, findErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, tokenRef, evmChain.Selector, datastore_utils.FullRef); findErr == nil {
					tokenRef = found
				} else {
					b.Logger.Warnf("token address could not be resolved using TokenRef (%+v): %v; falling back to TokenPoolRef", tokenRef, findErr)
					poolRef, poolErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, input.TokenPoolRef, evmChain.Selector, datastore_utils.FullRef)
					if poolErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve token pool ref (%+v): %w", input.TokenPoolRef, poolErr)
					}
					tokenAddrReport, getErr := cldf_ops.ExecuteOperation(b, token_pool.GetToken, evmChain, contract.FunctionInput[struct{}]{
						ChainSelector: evmChain.Selector,
						Address:       common.HexToAddress(poolRef.Address),
					})
					if getErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to derive token address from pool %s: %w", poolRef.Address, getErr)
					}
					tokenRef = datastore.AddressRef{
						ChainSelector: evmChain.Selector,
						Address:       tokenAddrReport.Output.Hex(),
					}
				}
			}

			tokenAddr := common.HexToAddress(tokenRef.Address)
			if tokenAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("resolved token address is zero for ref %+v", tokenRef)
			}

			if !common.IsHexAddress(input.ProposedOwner) {
				return sequences.OnChainOutput{}, fmt.Errorf("proposed owner address %q is not a valid hex address", input.ProposedOwner)
			}
			proposedOwner := common.HexToAddress(input.ProposedOwner)
			if proposedOwner == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("proposed owner address cannot be the zero address")
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, chains,
				tarseq.ManualRegistrationSequence,
				tarseq.ManualRegistrationSequenceInput{
					AdminAddress:  proposedOwner,
					ChainSelector: evmChain.Selector,
					TokenAddress:  tokenAddr,
					Address:       tarAddr,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to manually register token on chain %d: %w", evmChain.Selector, err)
			}

			return result, nil
		})
}

func (t *TokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	return cldf_ops.NewSequence(
		"evm-2.0-adapter:update-authorities",
		cciputils.Version_2_0_0,
		"Transfer token pool ownership to timelock on EVM chain",
		func(b cldf_ops.Bundle, e *deployment.Environment, input tokens.UpdateAuthoritiesInput) (sequences.OnChainOutput, error) {
			evmChain, ok := e.BlockChains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}

			timelockRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(cciputils.RBACTimelock),
				ChainSelector: evmChain.Selector,
				Qualifier:     cciputils.CLLQualifier,
			}, evmChain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find timelock for chain %d: %w", input.ChainSelector, err)
			}

			adapter := &ownershipAdapters.EVMTransferOwnershipAdapter{}
			if err := adapter.InitializeTimelockAddress(*e, mcms.Input{Qualifier: cciputils.CLLQualifier}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize timelock address: %w", err)
			}

			ownershipInput := deployapi.TransferOwnershipPerChainInput{
				ChainSelector: evmChain.Selector,
				CurrentOwner:  evmChain.DeployerKey.From.Hex(),
				ProposedOwner: timelockRef.Address,
				ContractRef:   []datastore.AddressRef{input.TokenPoolRef},
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, e.BlockChains,
				adapter.SequenceTransferOwnershipViaMCMS(), ownershipInput, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership on chain %d: %w", input.ChainSelector, err)
			}

			result, err = sequences.RunAndMergeSequence(b, e.BlockChains,
				adapter.SequenceAcceptOwnership(), ownershipInput, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership on chain %d: %w", input.ChainSelector, err)
			}

			return result, nil
		})
}

func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokens.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.MigrateLockReleasePoolLiquidity
}

// isBurnMintPoolType returns true for 2.0.0 burn/mint pool types.
func isBurnMintPoolType(poolType deployment.ContractType) bool {
	return poolType == cciputils.BurnMintTokenPool ||
		poolType == cciputils.BurnFromMintTokenPool ||
		poolType == cciputils.BurnWithFromMintTokenPool
}

// isLockReleasePoolType returns true for 2.0.0 lock/release pool types.
func isLockReleasePoolType(poolType deployment.ContractType) bool {
	return poolType == cciputils.LockReleaseTokenPool ||
		poolType == deployment.ContractType("SiloedLockReleaseTokenPool")
}

// isBurnMintTokenType returns true for EVM burn/mint token types that support role management.
func isBurnMintTokenType(typ datastore.ContractType) bool {
	return typ.String() == bnmOps.ContractType.String() || typ.String() == bnmDripOps.ContractType.String()
}
