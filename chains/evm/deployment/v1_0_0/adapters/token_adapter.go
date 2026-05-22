package adapters

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	bnmERC20Ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmDripOpsV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	v1_0_0_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	bnmDripOpsV1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	_ tokensapi.TokenRefResolver = &EVMTokenBase{}
	_ tokensapi.TokenAdapter     = &EVMTokenBase{}
)

// EVMTokenBase provides version-agnostic EVM token adapter methods that are
// shared across all pool versions (v1.5.1, v1.6.0, v1.6.1, v2.0.0).
// It is also registered at v1.0.0 so callers that only need token deployment
// can obtain a valid adapter without importing a pool-version package.
type EVMTokenBase struct{}

func (a *EVMTokenBase) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("address %q is not a valid hex address", ref.Address)
	}

	return common.HexToAddress(ref.Address).Bytes(), nil
}

func (a *EVMTokenBase) DeployToken() *cldf_ops.Sequence[tokensapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return v1_0_0_seq.DeployToken
}

func (a *EVMTokenBase) DeployTokenVerify(e deployment.Environment, input tokensapi.DeployTokenInput) error {
	tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore,
		datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(input.Type),
			Qualifier:     input.Symbol,
		},
		input.ChainSelector,
		datastore_utils.FullRef,
	)
	if err == nil {
		e.Logger.Info("Token already deployed at address:", tokenAddr.Address)
		return nil
	}

	if err := utils.ValidateEVMAddress(input.CCIPAdmin, "CCIPAdmin"); err != nil {
		return err
	}
	if err := utils.ValidateEVMAddress(input.ExternalAdmin, "ExternalAdmin"); err != nil {
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

// UpdateAuthorities transfers token pool ownership to the timelock via MCMS.
// It creates a self-contained EVMTransferOwnershipAdapter within the sequence
// closure so it works correctly regardless of how the embedding struct is initialized.
func (a *EVMTokenBase) UpdateAuthorities() *cldf_ops.Sequence[tokensapi.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	return cldf_ops.NewSequence(
		"evm-base:update-authorities",
		cciputils.Version_1_0_0,
		"Transfer token pool ownership to timelock on EVM chain",
		func(b cldf_ops.Bundle, e *deployment.Environment, input tokensapi.UpdateAuthoritiesInput) (sequences.OnChainOutput, error) {
			chain, ok := e.BlockChains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			adapter := &EVMTransferOwnershipAdapter{}
			if err := adapter.InitializeTimelockAddress(*e, mcms.Input{Qualifier: cciputils.CLLQualifier}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize timelock address for chain %d: %w", input.ChainSelector, err)
			}
			timelockAddr, err := a.GetTimelockAddressCLL(e.DataStore, chain.Selector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get timelock address for chain %d: %w", input.ChainSelector, err)
			}

			ownershipInput := deployops.TransferOwnershipPerChainInput{
				ChainSelector: chain.Selector,
				CurrentOwner:  chain.DeployerKey.From.Hex(),
				ProposedOwner: timelockAddr.Hex(),
				ContractRef:   []datastore.AddressRef{input.TokenPoolRef},
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, e.BlockChains, adapter.SequenceTransferOwnershipViaMCMS(), ownershipInput, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership on chain %d: %w", input.ChainSelector, err)
			}
			result, err = sequences.RunAndMergeSequence(b, e.BlockChains, adapter.SequenceAcceptOwnership(), ownershipInput, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership on chain %d: %w", input.ChainSelector, err)
			}

			return result, nil
		})
}

func (a *EVMTokenBase) ResolveTokenPoolRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, _ datastore.DataStore, chainSelector uint64, address string) (datastore.AddressRef, error) {
	var poolAddress common.Address
	if !common.IsHexAddress(address) {
		return datastore.AddressRef{}, fmt.Errorf("pool address %q is not a valid hex address", address)
	} else {
		poolAddress = common.HexToAddress(address)
	}

	chain, ok := chains.EVMChains()[chainSelector]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	tv, err := cldf_ops.ExecuteOperation(b,
		type_and_version.GetTypeAndVersion, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       poolAddress,
		},
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get typeAndVersion for pool %s: %w", address, err)
	}

	// NOTE: for EVM, the token pool qualifier is typically set to the token address
	// although this is not a hard requirement. We attempt to pull the token address
	// from the token pool on a best-effort basis, but if this fails for any reason,
	// then we will fallback to a placeholder qualifier. At the time of this writing
	// every token pool contract has a getToken( ) function with the same ABI across
	// all versions (v1.5.x, v1.6.x, v2.x.x) so we use the v1.5.x generated bindings
	// here for simplicity instead of overcomplicating the code with a switch on the
	// pool version.
	qualifier := fmt.Sprintf("%s-%s", poolAddress, tv.Output.Type)
	if token, err := cldf_ops.ExecuteOperation(b,
		token_pool.GetToken, chain,
		evm_contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       poolAddress,
		},
	); err != nil {
		b.Logger.Warnf("failed to get token address from pool at %s: %v; using fallback qualifier %s", poolAddress, err, qualifier)
	} else {
		qualifier = token.Output.Hex()
	}

	return datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(tv.Output.Type),
		Qualifier:     qualifier,
		Version:       tv.Output.Version,
		Address:       poolAddress.Hex(),
	}, nil
}

func (a *EVMTokenBase) ResolveTokenRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, _ datastore.DataStore, chainSelector uint64, address string) (datastore.AddressRef, error) {
	var tokenAddress common.Address
	if !common.IsHexAddress(address) {
		return datastore.AddressRef{}, fmt.Errorf("token address %q is not a valid hex address", address)
	} else {
		tokenAddress = common.HexToAddress(address)
	}

	chain, ok := chains.EVMChains()[chainSelector]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	symbolReport, err := cldf_ops.ExecuteOperation(b,
		erc20.GetSymbol, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       tokenAddress,
		},
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to read ERC20 symbol for token %s: %w", address, err)
	}

	// NOTE: at the moment, there's currently not an easy way to determine the exact
	// token type. For now, we simply return `ERC20Token` as the type for all tokens
	// but if needed we can add more specific logic here in the future.
	return datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(erc20.ContractType),
		Version:       cciputils.Version_1_0_0,
		Qualifier:     symbolReport.Output,
		Address:       tokenAddress.Hex(),
	}, nil
}

// Pool-specific stubs -- these are overridden by per-version adapters (v1.5.1, v1.6.1, v2.0.0).
// EVMTokenBase is registered at v1.0.0 so callers that only need token deployment (DeployToken,
// DeployTokenVerify) can obtain a valid adapter without importing a pool-version package.

func (a *EVMTokenBase) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) DeriveTokenPoolCounterpart(_ deployment.Environment, _ uint64, tokenPool []byte, _ []byte) ([]byte, error) {
	return tokenPool, nil
}

func (a *EVMTokenBase) DeriveTokenAddress(_ deployment.Environment, _ uint64, _ datastore.AddressRef) (string, error) {
	return "", fmt.Errorf("DeriveTokenAddress is not implemented on EVMTokenBase; use a pool-version adapter")
}

func (a *EVMTokenBase) DeriveTokenDecimals(_ deployment.Environment, _ uint64, _ datastore.AddressRef, _ []byte) (uint8, error) {
	return 0, fmt.Errorf("DeriveTokenDecimals is not implemented on EVMTokenBase; use a pool-version adapter")
}

func (a *EVMTokenBase) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokensapi.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

// ================================================================
// === Version-agnostic helpers for all EVM token/pool versions ===
// ================================================================

// IsBurnMintPoolType returns true if the pool type is one of the burn-mint variants (standard or with from-mint).
func (a *EVMTokenBase) IsBurnMintPoolType(poolType string) bool {
	return poolType == cciputils.BurnMintTokenPool.String() ||
		poolType == cciputils.BurnFromMintTokenPool.String() ||
		poolType == cciputils.BurnWithFromMintTokenPool.String()
}

// IsLockReleasePoolType returns true if the pool type is one of the lock-release variants (standard or siloed).
func (a *EVMTokenBase) IsLockReleasePoolType(poolType string) bool {
	return poolType == cciputils.LockReleaseTokenPool.String() ||
		poolType == cciputils.SiloedLockReleaseTokenPool.String()
}

// IsBurnMintTokenType returns true if the token type is one of the burn-mint variants (ERC20 or ERC677).
func (a *EVMTokenBase) IsBurnMintTokenType(tokenType string) bool {
	return tokenType == bnmERC20Ops.ContractType.String() ||
		tokenType == bnmDripOpsV1_0_0.ContractType.String() ||
		tokenType == bnmDripOpsV1_5_0.ContractType.String() ||
		a.IsBurnMintERC677TokenType(tokenType)
}

// IsBurnMintERC677TokenType returns true if the token type is one of the burn-mint ERC677 variants.
func (a *EVMTokenBase) IsBurnMintERC677TokenType(tokenType string) bool {
	return tokenType == cciputils.BurnMintToken.String() ||
		tokenType == cciputils.ERC677TokenHelper.String()
}

// resolveRouterAddress returns the router address to wire into the pool.
// If routerRef is nil, the chain's production Router is looked up in the datastore.
// If routerRef.Address is non-empty, it is used directly (no datastore lookup).
// Otherwise the ref is resolved against the datastore; ChainSelector is forced to
// the target chain and Type defaults to the production Router when unset, so callers
// targeting the TestRouter only need to set Type=router.TestRouterContractType.
func (a *EVMTokenBase) ResolveRouterAddress(ds datastore.DataStore, chainSelector uint64, routerRef *datastore.AddressRef) (common.Address, error) {
	ref := datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.ContractType),
	}
	if routerRef != nil {
		if routerRef.Address != "" {
			if !common.IsHexAddress(routerRef.Address) {
				return common.Address{}, fmt.Errorf("invalid RouterRef.Address %q: not a hex address", routerRef.Address)
			}
			addr := common.HexToAddress(routerRef.Address)
			if addr == (common.Address{}) {
				return common.Address{}, errors.New("RouterRef.Address resolves to the zero address")
			}
			return addr, nil
		}
		ref = *routerRef
		ref.ChainSelector = chainSelector
		if ref.Type == "" {
			ref.Type = datastore.ContractType(router.ContractType)
		}
	}
	resolved, err := datastore_utils.FindAndFormatRef(ds, ref, chainSelector, datastore_utils_evm.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find router (type=%q qualifier=%q) in datastore for chain %d: %w", ref.Type, ref.Qualifier, chainSelector, err)
	}
	return resolved, nil
}

// GetTokenAdminRegistryAddress looks up the TAR (v1.5.0) address from the datastore.
func (a *EVMTokenBase) GetTokenAdminRegistryAddress(ds datastore.DataStore, selector uint64) (common.Address, error) {
	filters := datastore.AddressRef{
		Type:          datastore.ContractType(tarops.ContractType),
		ChainSelector: selector,
		Version:       tarops.Version,
	}
	addr, err := datastore_utils.FindAndFormatRef(ds, filters, selector, datastore_utils_evm.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token admin registry address on chain %d: %w", selector, err)
	}
	return addr, nil
}

// GetTimelockAddressCLL looks up the timelock (RBACTimelock) address from the datastore using the CLL qualifier.
func (a *EVMTokenBase) GetTimelockAddressCLL(ds datastore.DataStore, selector uint64) (common.Address, error) {
	filters := datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(cciputils.RBACTimelock),
		Version:       cciputils.Version_1_0_0,
		Qualifier:     cciputils.CLLQualifier,
	}
	addr, err := datastore_utils.FindAndFormatRef(ds, filters, selector, datastore_utils_evm.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find timelock address on chain %d: %w", selector, err)
	}
	return addr, nil
}

// ParseAddressRef attempts to parse an address from the given ref. If ref.Address is non-empty, then the datastore
// lookup is skipped and the provided address is parsed as an EVM address and returned directly. Otherwise, the ref
// is resolved against the datastore and parsed as a hex address.
func (a *EVMTokenBase) ParseAddressRef(ds datastore.DataStore, ref datastore.AddressRef, sel uint64) (common.Address, error) {
	refAddr := ref.Address
	if refAddr != "" {
		if !common.IsHexAddress(refAddr) {
			return common.Address{}, fmt.Errorf("invalid address %q: not a hex address", refAddr)
		} else {
			return common.HexToAddress(refAddr), nil
		}
	}

	evmAddr, err := datastore_utils.FindAndFormatRef(ds, ref, sel, datastore_utils_evm.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to resolve address from datastore using ref filter (%s): %w", datastore_utils.SprintRef(ref), err)
	}

	return evmAddr, nil
}
