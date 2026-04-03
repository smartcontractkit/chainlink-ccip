package adapters

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/token_pool"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences/token_pool"
	v1_6_0_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
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

var _ tokensapi.TokenAdapter = &TokenAdapter{}

// TokenAdapter handles EVM token pools at version 1.5.1.
// It embeds EVMTokenBase for shared methods and defines all pool-specific
// methods using only v1.5.1 bindings -- no version switches.
type TokenAdapter struct {
	evm1_0_0.EVMTokenBase
}

func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.1-adapter:configure-token-for-transfers",
		tpOps.Version,
		"Configure a v1.5.1 token pool for cross-chain transfers on an EVM chain",
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

			tarAddress, err := getTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector, &t.EVMTokenBase)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokenAddress, err := getTokenFromPool(b, chain, tpAddr)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from pool at %s: %w", tpAddr, err)
			}

			configureReport, err := cldf_ops.ExecuteSequence(b,
				tpSeq.ConfigureTokenPoolForRemoteChains, chain,
				tpSeq.ConfigureTokenPoolForRemoteChainsInput{
					TokenPoolAddress: tpAddr,
					TokenPoolVersion: tpOps.Version,
					RemoteChains:     input.RemoteChains,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
			}
			result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)

			registerReport, err := cldf_ops.ExecuteSequence(b,
				tarseq.RegisterToken, chain,
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

func (t *TokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	addrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, poolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find token pool in datastore using ref (%+v): %w", poolRef, err)
	}

	addrRaw, err := t.AddressRefToBytes(addrRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	tpAddr := common.BytesToAddress(addrRaw)
	if tpAddr == (common.Address{}) {
		return nil, errors.New("token pool address is zero address")
	}

	tokenAddr, err := getTokenFromPool(e.OperationsBundle, chain, tpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get token address from token pool ref (%+v): %w", addrRef, err)
	}

	return tokenAddr.Bytes(), nil
}

func (t *TokenAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, _ []byte) (uint8, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	addrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, poolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return 0, fmt.Errorf("failed to find token pool in datastore using ref (%+v): %w", poolRef, err)
	}

	addrRaw, err := t.AddressRefToBytes(addrRef)
	if err != nil {
		return 0, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	tpAddr := common.BytesToAddress(addrRaw)
	if tpAddr == (common.Address{}) {
		return 0, errors.New("token pool address is zero address")
	}

	pool, err := token_pool.NewTokenPool(tpAddr, chain.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
	}

	return pool.GetTokenDecimals(&bind.CallOpts{Context: e.GetContext()})
}

func (t *TokenAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokensapi.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.1-adapter:set-token-pool-rate-limits",
		tpOps.Version,
		"Set rate limits for a v1.5.1 token pool on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.TPRLRemotes) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tokenPoolAddrBytes, err := t.AddressRefToBytes(input.TokenPoolRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", err)
			}
			tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
			if tokenPoolAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref (%+v) is zero address", input.TokenPoolRef)
			}

			if input.SkipIfMissingPermissions {
				timelockFltr := datastore.AddressRef{Type: datastore.ContractType(cciputils.RBACTimelock), ChainSelector: chain.Selector, Qualifier: cciputils.CLLQualifier}
				timelockAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, timelockFltr, chain.Selector, datastore_utils_evm.ToEVMAddress)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to find timelock address for chain %d: %w", chain.Selector, err)
				}
				poolOwner, rlAdmin, err := getTokenPoolAdmins(b.GetContext(), &chain, tokenPoolAddr)
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

			report, err := cldf_ops.ExecuteOperation(b,
				tpOps.SetChainRateLimiterConfig, chain,
				evm_contract.FunctionInput[tpOps.SetChainRateLimiterConfigArgs]{
					ChainSelector: chain.Selector,
					Address:       tokenPoolAddr,
					Args: tpOps.SetChainRateLimiterConfigArgs{
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
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
			}

			batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			result.BatchOps = append(result.BatchOps, batchOp)
			return result, nil
		})
}

func (t *TokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.1-adapter:manual-registration",
		tpOps.Version,
		"Manually register a token and token pool on EVM chains using v1.5.1 pools",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ManualRegistrationSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tarAddress, err := getTokenAdminRegistryAddress(input.ExistingDataStore, chain.Selector, &t.EVMTokenBase)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", chain.Selector, err)
			}

			tokenRef := input.TokenRef
			if tokenRef.Address == "" {
				if tokRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, tokenRef, chain.Selector, datastore_utils.FullRef); err != nil {
					b.Logger.Warnf("token address could not be resolved using TokenRef (%+v): %v", tokenRef, err)
					b.Logger.Warnf("attempting to resolve token address using TokenPoolRef instead: (%+v)", input.TokenPoolRef)

					tokenPoolRef, poolErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, input.TokenPoolRef, chain.Selector, datastore_utils.FullRef)
					if poolErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("token pool could not be resolved using TokenPoolRef (%+v): %w", input.TokenPoolRef, poolErr)
					}
					tokenPoolAddrBytes, addrErr := t.AddressRefToBytes(tokenPoolRef)
					if addrErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", addrErr)
					}
					tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
					if tokenPoolAddr == (common.Address{}) {
						return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref (%+v) is zero address", tokenPoolRef)
					}
					tokenAddr, getErr := getTokenFromPool(b, chain, tokenPoolAddr)
					if getErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool ref (%+v): %w", tokenPoolRef, getErr)
					}

					tokenRef = datastore.AddressRef{
						ChainSelector: chain.Selector,
						Address:       tokenAddr.Hex(),
					}
				} else {
					tokenRef = tokRef
				}
			}

			tokenAddrBytes, err := t.AddressRefToBytes(tokenRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token address ref to bytes: %w", err)
			}
			tokenAddr := common.BytesToAddress(tokenAddrBytes)
			if tokenAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token address for ref (%+v) is zero address", tokenRef)
			}

			if !common.IsHexAddress(input.ProposedOwner) {
				return sequences.OnChainOutput{}, fmt.Errorf("proposed owner address %q is not a valid hex address", input.ProposedOwner)
			}
			proposedOwnerAddr := common.HexToAddress(input.ProposedOwner)
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
					Address:       tarAddress,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to manually register token on chain %d: %w", chain.Selector, err)
			}

			return result, nil
		})
}

func (t *TokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.1-adapter:deploy-token-pool-for-token",
		tpOps.Version,
		"Deploy a v1.5.1 token pool for a token on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			out, err := cldf_ops.ExecuteSequence(b, v1_6_0_seq.DeployTokenPool, chains, input)
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

			isToknTypeBnM := toknRef.Type.String() == bnmERC20ops.ContractType.String()
			isPoolTypeBnM := input.PoolType == cciputils.BurnMintTokenPool.String()
			if isPoolTypeBnM && isToknTypeBnM && len(out.Output.Addresses) >= 1 {
				poolRef := out.Output.Addresses[0]

				poolAddrBytes, addrErr := t.AddressRefToBytes(poolRef)
				if addrErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to convert deployed token pool address ref to bytes: %w", addrErr)
				}
				toknAddrBytes, addrErr := t.AddressRefToBytes(toknRef)
				if addrErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token address ref to bytes: %w", addrErr)
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

				report, execErr := cldf_ops.ExecuteOperation(b,
					bnmERC20ops.GrantMintAndBurnRoles, chain,
					evm_contract.FunctionInput[common.Address]{
						ChainSelector: input.ChainSelector,
						Address:       toknAddr,
						Args:          poolAddr,
					},
				)
				if execErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenRef.Qualifier, input.ChainSelector, execErr)
				}

				batchOp, bErr := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{report.Output})
				if bErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for granting mint and burn roles: %w", bErr)
				}
				result.BatchOps = append(result.BatchOps, batchOp)
			}

			return result, nil
		},
	)
}

// getTokenFromPool calls GetToken on a v1.5.1 token pool to derive the token address.
func getTokenFromPool(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		tpOps.GetToken, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get token address via GetToken v1.5.1 operation: %w", err)
	}
	return res.Output, nil
}

// getTokenPoolAdmins returns the owner and rate limit admin of a v1.5.1 token pool.
func getTokenPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (poolOwner common.Address, rlAdmin common.Address, err error) {
	pool, err := token_pool.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
	}

	poolOwner, err = pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	rlAdmin, err = pool.GetRateLimitAdmin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get rate limit admin of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}

	return poolOwner, rlAdmin, nil
}

// getTokenAdminRegistryAddress looks up the TAR (v1.5.0) address from the datastore.
func getTokenAdminRegistryAddress(ds datastore.DataStore, selector uint64, base *evm1_0_0.EVMTokenBase) (common.Address, error) {
	filters := datastore.AddressRef{
		Type:          datastore.ContractType(tarops.ContractType),
		ChainSelector: selector,
		Version:       tarops.Version,
	}
	ref, err := datastore_utils.FindAndFormatRef(ds, filters, selector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token admin registry address on chain %d: %w", selector, err)
	}
	addr, err := base.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}
	return common.BytesToAddress(addr), nil
}
