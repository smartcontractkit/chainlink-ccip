package sequences

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/token_pool"
	tpseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

func (a *EVMAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:configure-token-for-transfers",
		tpops.Version,
		"Configure a token for cross-chain transfers for an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			if !common.IsHexAddress(input.TokenPoolAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address %q is not a valid hex address", input.TokenPoolAddress)
			}
			if !common.IsHexAddress(input.RegistryAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("registry address %q is not a valid hex address", input.RegistryAddress)
			}

			tpAddr := common.HexToAddress(input.TokenPoolAddress)
			if tpAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("token pool address is zero address")
			}

			trAddr := common.HexToAddress(input.RegistryAddress)
			if trAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("token admin registry address is zero address")
			}

			externalAdmin := common.Address{}
			if input.ExternalAdmin != "" {
				if !common.IsHexAddress(input.ExternalAdmin) {
					return sequences.OnChainOutput{}, fmt.Errorf("external admin address %q is not a valid hex address", input.ExternalAdmin)
				}

				externalAdmin = common.HexToAddress(input.ExternalAdmin)
			}

			token, err := cldf_ops.ExecuteOperation(b,
				tpops.GetToken,
				chain,
				evm_contract.FunctionInput[struct{}]{
					ChainSelector: input.ChainSelector,
					Address:       tpAddr,
					Args:          struct{}{},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address via GetToken operation: %w", err)
			}

			configureReport, err := cldf_ops.ExecuteSequence(b,
				tpseq.ConfigureTokenPoolForRemoteChains,
				chain,
				tpseq.ConfigureTokenPoolForRemoteChainsInput{
					TokenPoolAddress: tpAddr,
					RemoteChains:     input.RemoteChains,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
			}

			registerReport, err := cldf_ops.ExecuteSequence(b,
				tarseq.RegisterToken,
				chain,
				tarseq.RegisterTokenInput{
					ChainSelector:             input.ChainSelector,
					TokenAdminRegistryAddress: trAddr,
					TokenPoolAddress:          tpAddr,
					ExternalAdmin:             externalAdmin,
					TokenAddress:              token.Output,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token on chain %d: %w", input.ChainSelector, err)
			}

			var result sequences.OnChainOutput
			result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
			result.Addresses = append(result.Addresses, registerReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)
			result.BatchOps = append(result.BatchOps, registerReport.Output.BatchOps...)

			return result, nil
		})
}

func (a *EVMAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("address %q is not a valid hex address", ref.Address)
	}

	return common.HexToAddress(ref.Address).Bytes(), nil
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

	token, err := cldf_ops.ExecuteOperation(e.OperationsBundle,
		tpops.GetToken,
		chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       tpAddr,
			Args:          struct{}{},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get token address via GetToken operation: %w", err)
	}

	return token.Output.Bytes(), nil
}

func (a *EVMAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) (uint8, error) {
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

	tp, err := token_pool.NewTokenPool(tpAddr, chain.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to instantiate token pool contract: %w", err)
	}
	return tp.GetTokenDecimals(&bind.CallOpts{Context: e.GetContext()})
}

func (a *EVMAdapter) DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
	// For EVM chains, the token pool address is not derived from the token address, so we can return the token pool address as is.
	return tokenPool, nil
}

func (a *EVMAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokensapi.RateLimiterConfigInputs, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:set-token-pool-rate-limits",
		tpops.Version,
		"Set rate limits for a token pool across multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.RateLimiterConfigInputs) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			tpAddress, err := a.FindLatestTokenPoolAddress(input.ExistingDataStore, input.ChainSelector, input.TokenPoolQualifier, input.PoolType)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool with qualifier %q on chain %d: %w", input.TokenPoolQualifier, input.ChainSelector, err)
			}

			report, err := cldf_ops.ExecuteOperation(b, tpops.SetChainRateLimiterConfig, chain, evm_contract.FunctionInput[tpops.SetChainRateLimiterConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       tpAddress,
				Args: tpops.SetChainRateLimiterConfigArgs{
					OutboundRateLimitConfig: token_pool.RateLimiterConfig{
						IsEnabled: input.OutboundRateLimiterConfig.IsEnabled,
						Capacity:  input.OutboundRateLimiterConfig.Capacity,
						Rate:      input.OutboundRateLimiterConfig.Rate},
					InboundRateLimitConfig: token_pool.RateLimiterConfig{
						IsEnabled: input.InboundRateLimiterConfig.IsEnabled,
						Capacity:  input.InboundRateLimiterConfig.Capacity,
						Rate:      input.InboundRateLimiterConfig.Rate},
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

func (a *EVMAdapter) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:manual-registration",
		tarops.Version,
		"Manually register a token and token pool on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ManualRegistrationInput) (sequences.OnChainOutput, error) {
			store := datastore.NewMemoryDataStore()
			for _, addr := range input.ExistingAddresses {
				if err := store.AddressRefStore.Upsert(addr); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to upsert address %v: %w", addr, err)
				}
			}
			ds := store.Seal()

			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tokenAdminRegistryAddress, err := a.GetTokenAdminRegistryAddress(ds, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokenPoolAddress, err := a.FindLatestTokenPoolAddress(ds, chain.Selector, input.RegisterTokenConfigs.TokenPoolQualifier, input.RegisterTokenConfigs.PoolType)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool with qualifier %q on chain %d: %w", input.RegisterTokenConfigs.TokenPoolQualifier, input.ChainSelector, err)
			}

			token, err := cldf_ops.ExecuteOperation(b,
				tpops.GetToken,
				chain,
				evm_contract.FunctionInput[struct{}]{
					ChainSelector: input.ChainSelector,
					Address:       tokenPoolAddress,
					Args:          struct{}{},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address via GetToken operation: %w", err)
			}

			proposedOwner := input.RegisterTokenConfigs.ProposedOwner
			if !common.IsHexAddress(proposedOwner) {
				return sequences.OnChainOutput{}, fmt.Errorf("proposed owner address %q is not a valid hex address", proposedOwner)
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, chains,
				tarseq.ManualRegistrationSequence,
				tarseq.ManualRegistrationSequenceInput{
					AdminAddress:  common.HexToAddress(proposedOwner),
					ChainSelector: input.ChainSelector,
					TokenAddress:  token.Output,
					Address:       tokenAdminRegistryAddress,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to manually register token on chain %d: %w", input.ChainSelector, err)
			}

			return result, nil
		})
}

func (a *EVMAdapter) DeployToken() *cldf_ops.Sequence[tokensapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DeployToken
}

func (a *EVMAdapter) DeployTokenVerify(e deployment.Environment, in any) error {
	input := in.(tokensapi.DeployTokenInput)

	tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
		ChainSelector: input.ChainSelector,
		Type:          datastore.ContractType(input.Type),
		Qualifier:     input.Symbol,
	}, input.ChainSelector, datastore_utils.FullRef)
	if err == nil {
		e.OperationsBundle.Logger.Info("Token already deployed at address:", tokenAddr.Address)
		return nil
	}

	// Validate EVM addresses from chain-agnostic input
	if err := utils.ValidateEVMAddress(input.CCIPAdmin, "CCIPAdmin"); err != nil {
		return err
	}
	if err := utils.ValidateEVMAddress(input.ExternalAdmin, "ExternalAdmin"); err != nil {
		return err
	}
	// ensuring that decimals is not more than 18
	if input.Decimals > 18 {
		return fmt.Errorf("EVM tokens cannot have more than 18 decimals, got %d", input.Decimals)
	}
	// ensuring that supply and pre-mint are not negative
	if input.Supply != nil && input.Supply.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("token supply cannot be negative, got %v", *input.Supply)
	}
	if input.PreMint != nil && input.PreMint.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("token pre-mint cannot be negative, got %v", *input.PreMint)
	}

	return nil
}

func (a *EVMAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:deploy-token-pool-for-token",
		tpops.Version,
		"Deploy a token pool for a token on an EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
			out, err := cldf_ops.ExecuteSequence(b, DeployTokenPool, chains, input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token pool on chain %d: %w", input.ChainSelector, err)
			}

			var result sequences.OnChainOutput
			result.Addresses = append(result.Addresses, out.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)

			toknFilterDS := datastore.AddressRef{ChainSelector: input.ChainSelector, Qualifier: input.TokenSymbol}
			toknRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, toknFilterDS, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token address for symbol %q on chain %d: %w", input.TokenSymbol, input.ChainSelector, err)
			}

			// For a BnM token + BnM token pool, we need to grant the pool mint and burn roles on the token
			isToknTypeBnM := toknRef.Type.String() == bnmERC20ops.ContractType.String()
			isPoolTypeBnM := input.PoolType == cciputils.BurnMintTokenPool.String()
			if isPoolTypeBnM && isToknTypeBnM {
				// NOTE: the pool ref isn't in the datastore yet so
				// we locate it from the sequence output addresses.
				poolRef, foundIt := datastore.AddressRef{}, false
				for _, addrRef := range out.Output.Addresses {
					isPoolRef := addrRef.ChainSelector == input.ChainSelector &&
						addrRef.Qualifier == input.TokenPoolQualifier &&
						addrRef.Type.String() == input.PoolType &&
						addrRef.Address != ""

					if isPoolRef {
						poolRef = addrRef
						foundIt = true
						break
					}
				}

				if !foundIt {
					return sequences.OnChainOutput{}, fmt.Errorf("deployed token pool address for qualifier %q on chain %d not found in output addresses", input.TokenPoolQualifier, input.ChainSelector)
				}

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
					return sequences.OnChainOutput{}, fmt.Errorf("token address for symbol %q is zero address", input.TokenSymbol)
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
					return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenSymbol, input.ChainSelector, err)
				}

				batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{report.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for granting mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenSymbol, input.ChainSelector, err)
				}

				result.BatchOps = append(result.BatchOps, batchOp)
			}

			return result, nil
		},
	)
}

func (a *EVMAdapter) RegisterToken() *cldf_ops.Sequence[tokensapi.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:register-token",
		tarops.Version,
		"Register a token and its pool on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.RegisterTokenInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tarAddress, err := a.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokAddress, err := a.FindOneTokenAddress(input.ExistingDataStore, input.ChainSelector, input.TokenSymbol)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address for symbol %q on chain %d: %w", input.TokenSymbol, input.ChainSelector, err)
			}

			extnAdmin := common.Address{}
			if input.TokenAdmin != "" {
				if !common.IsHexAddress(input.TokenAdmin) {
					return sequences.OnChainOutput{}, fmt.Errorf("token admin address %q is not a valid hex address", input.TokenAdmin)
				}

				extnAdmin = common.HexToAddress(input.TokenAdmin)
			}

			// INFO: by default, operations are cached if they are called with the same ID, version,
			// description, and input params. For this sequence (and token expansion in general), we
			// need to be extra cautious with this behavior. Many token API sequence implementations
			// call the ExecuteOperation function assuming that they are getting the latest on-chain
			// data back. However, if many of these seqeuences are combined to create one changeset,
			// (e.g. TokenExpansion) then downstream sequences will receive stale data from upstream
			// ones due to caching causing issues. Here is an example scenario:
			// --
			//   1. sequence A calls GET token config with some payload (config is empty)
			//   2. sequence B calls SET token config (config is no longer empty now)
			//   3. sequence C calls GET token config with the same exact payload (the empty cached config is returned, which is stale)
			// --
			// Unfortunately, this sequence is subject to the issue described above, and there is no
			// clean way to disable the default caching mechanism at present. To circumvent this for
			// now, we intentionally avoid the use of ExecuteOperation in favor of fetching directly
			// from the contract. This type of workaround only needs to be added if there's a chance
			// that there's a changeset that does read -> write -> read, so a function like GetToken
			// is not affected since the value is immutable after deployment. For completeness, here
			// is an example of what should *NOT* be done:
			// --
			//   ```go
			//     cached, err := operations.ExecuteOperation(b, tarops.GetTokenConfig, chain,
			//     	contract.FunctionInput[common.Address]{
			//     		ChainSelector: input.ChainSelector,
			//     		Address:       tarAddress,
			//     		Args:          tokAddress,
			//     	})
			//   ```
			// --
			// Reference: https://docs.cld.cldev.sh/guides/changesets/operations-api
			// --
			tar, err := token_admin_registry.NewTokenAdminRegistry(tarAddress, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to bind to token admin registry at address %q on chain %d: %w", tarAddress, input.ChainSelector, err)
			}
			cfg, err := tar.GetTokenConfig(&bind.CallOpts{Context: b.GetContext()}, tokAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config for token %q from token admin registry at address %q: %w", tokAddress, tarAddress, err)
			}

			// INFO: there are two cases to consider here:
			// --
			//   1. The token pool is ALREADY set on-chain. In this case, we simply reuse the
			//      current on-chain value. This makes the RegisterToken sequence idempotent.
			// --
			//   2. The token pool is NOT set on-chain. In this case, the token pool address
			//      will be set to the zero address below. This causes the RegisterToken seq
			//      to skip calling `SetPool`, which is OK since the token pool expansion cs
			//      will call `SetPool` again later on.
			tpAddress := cfg.TokenPool

			report, err := cldf_ops.ExecuteSequence(b,
				tarseq.RegisterToken,
				chain,
				tarseq.RegisterTokenInput{
					ChainSelector:             input.ChainSelector,
					TokenAdminRegistryAddress: tarAddress,
					TokenPoolAddress:          tpAddress,
					ExternalAdmin:             extnAdmin,
					TokenAddress:              tokAddress,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token on chain %d: %w", input.ChainSelector, err)
			}

			result.Addresses = append(result.Addresses, report.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, report.Output.BatchOps...)
			return result, nil
		})
}

func (a *EVMAdapter) SetPool() *cldf_ops.Sequence[tokensapi.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-adapter:set-pool",
		tarops.Version,
		"Set the pool for a token across multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.SetPoolInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tokenAdminRegistryAddress, err := a.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokenPoolAddress, err := a.FindLatestTokenPoolAddress(input.ExistingDataStore, chain.Selector, input.TokenPoolQualifier, input.PoolType)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool with qualifier %q on chain %d: %w", input.TokenPoolQualifier, input.ChainSelector, err)
			}

			token, err := cldf_ops.ExecuteOperation(b,
				tpops.GetToken,
				chain,
				evm_contract.FunctionInput[struct{}]{
					ChainSelector: input.ChainSelector,
					Address:       tokenPoolAddress,
					Args:          struct{}{},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address via GetToken operation: %w", err)
			}

			report, err := cldf_ops.ExecuteOperation(b,
				tarops.SetPool,
				chain,
				evm_contract.FunctionInput[tarops.SetPoolArgs]{
					Address:       tokenAdminRegistryAddress,
					ChainSelector: input.ChainSelector,
					Args: tarops.SetPoolArgs{
						TokenPoolAddress: tokenPoolAddress,
						TokenAddress:     token.Output,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set pool for token %q on chain selector %d: %w", input.TokenSymbol, input.ChainSelector, err)
			}

			batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}

			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batchOp},
			}, nil
		},
	)
}

func (a *EVMAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokensapi.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO: implement me
	return nil
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

func (a *EVMAdapter) FindOneTokenAddress(ds datastore.DataStore, chainSelector uint64, tokenSymbol string) (common.Address, error) {
	filters := datastore.AddressRef{
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	}

	ref, err := datastore_utils.FindAndFormatRef(ds, filters, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token address for symbol %q on chain %d: %w", tokenSymbol, chainSelector, err)
	}

	addr, err := a.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	return common.BytesToAddress(addr), nil
}

func (a *EVMAdapter) FindLatestTokenPoolAddress(ds datastore.DataStore, chainSelector uint64, qualifier string, poolType string) (common.Address, error) {
	// Define the version range
	minVersion := semver.MustParse("1.5.0") // inclusive
	maxVersion := semver.MustParse("1.7.0") // exclusive

	// Get all matching token pool addresses
	refs := ds.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(poolType)),
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByQualifier(qualifier),
	)

	// Use the latest version found within the specified range
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
		return common.Address{}, fmt.Errorf("no token pool found for type %q with qualifier %q on chain %d", poolType, qualifier, chainSelector)
	}

	// Convert the address reference to bytes
	addrBytes, err := a.AddressRefToBytes(latestRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	// Construct the token pool instance
	return common.BytesToAddress(addrBytes), nil
}
