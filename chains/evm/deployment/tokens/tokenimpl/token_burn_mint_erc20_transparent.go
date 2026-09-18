package tokenimpl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_transparent"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenBurnMintERC20Transparent struct{}

func (tokenBurnMintERC20Transparent) ContractType() deployment.ContractType {
	return burn_mint_erc20_transparent.ContractType
}

func (tokenBurnMintERC20Transparent) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		// BurnMintERC20Transparent uses OZ's AccessControlDefaultAdminRulesUpgradeable, where
		// grantRole/revokeRole revert unconditionally for DEFAULT_ADMIN_ROLE. The role can only
		// move via beginDefaultAdminTransfer (GrantAdminRole, deployer-signed) followed by
		// acceptDefaultAdminTransfer (AcceptAdminRole, signed by the new admin themselves) - see
		// UsesAsyncRoleManagement. RevokeAdminRole has no faithful equivalent here and stays
		// unsupported: revokeRole/renounceRole for DEFAULT_ADMIN_ROLE only succeed when
		// abandoning adminship entirely (pending transfer target == address(0)), not when
		// removing one holder while another is mid-transfer.
		SupportsAdminRole:       true,
		UsesAsyncRoleManagement: true,
		SupportsCCIPAdmin:       true,
		SupportsPreMint:         true,
	}
}

func (tokenBurnMintERC20Transparent) RevokeAdminRole(b operations.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("RevokeAdminRole is not supported for BurnMintERC20TransparentToken: DEFAULT_ADMIN_ROLE transfer requires the 2-step beginDefaultAdminTransfer/acceptDefaultAdminTransfer flow (see Capabilities.UsesAsyncRoleManagement)")
}

// HasAdminRole checks DEFAULT_ADMIN_ROLE (the zero bytes32 constant for every OZ AccessControl
// contract, so no read is needed to look it up).
func (tokenBurnMintERC20Transparent) HasAdminRole(b operations.Bundle, chain evm.Chain, token, user common.Address) (bool, error) {
	report, err := operations.ExecuteOperation(
		b, burn_mint_erc20_transparent.HasRole, chain,
		contract.FunctionInput[burn_mint_erc20_transparent.RoleAssignment]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args: burn_mint_erc20_transparent.RoleAssignment{
				Role: [32]byte{},
				To:   user,
			},
		},
		operations.WithRetryConfig(getRetryConfig[burn_mint_erc20_transparent.RoleAssignment](b, chain, token.Hex())),
	)
	if err != nil {
		return false, fmt.Errorf("failed to check default admin role for %s: %w", user.Hex(), err)
	}

	return report.Output, nil
}

// GrantAdminRole starts the 2-step DEFAULT_ADMIN_ROLE transfer. It always executes synchronously
// (deployer-signed): it's onlyRole(DEFAULT_ADMIN_ROLE) and the deployer holds that role right
// after Initialize.
func (tokenBurnMintERC20Transparent) GrantAdminRole(b operations.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(
		b, burn_mint_erc20_transparent.BeginDefaultAdminTransfer, chain,
		contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          externalAdmin,
		},
		operations.WithRetryConfig(getRetryConfig[common.Address](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to begin default admin transfer: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

// AcceptAdminRole prepares (but, when called by the deployer, never directly executes) the
// acceptDefaultAdminTransfer() call. Callers must only invoke this when the pending transfer
// target is the chain's timelock - see Capabilities.UsesAsyncRoleManagement callers. The
// resulting WriteOutput is left unexecuted (folded into the caller's MCMS batch/proposal) so the
// timelock can execute it itself, since only it can satisfy the pendingDefaultAdmin check.
// NOTE: intentionally NOT retried - retrying an unexecuted, proposal-only write is meaningless
// (there's nothing to retry: the deployer was never going to sign it in the first place).
func (tokenBurnMintERC20Transparent) AcceptAdminRole(b operations.Bundle, chain evm.Chain, token common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(
		b, burn_mint_erc20_transparent.AcceptDefaultAdminTransfer, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          struct{}{},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare accept default admin transfer: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func (tokenBurnMintERC20Transparent) PendingAdminRoleTarget(b operations.Bundle, chain evm.Chain, token common.Address) (common.Address, error) {
	report, err := operations.ExecuteOperation(
		b, burn_mint_erc20_transparent.PendingDefaultAdmin, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          struct{}{},
		},
		operations.WithRetryConfig(getRetryConfig[struct{}](b, chain, token.Hex())),
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get pending default admin: %w", err)
	}

	return report.Output, nil
}

func (tokenBurnMintERC20Transparent) GrantPoolRoles(b operations.Bundle, chain evm.Chain, token, pool, _ common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(
		b, burn_mint_erc20_transparent.GrantMintAndBurnRoles, chain,
		contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          pool,
		},
		operations.WithRetryConfig(getRetryConfig[common.Address](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to grant mint and burn roles: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func (tokenBurnMintERC20Transparent) SetCCIPAdmin(b operations.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(
		b, burn_mint_erc20_transparent.SetCCIPAdmin, chain,
		contract.FunctionInput[string]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          ccipAdmin.Hex(),
		},
		operations.WithRetryConfig(getRetryConfig[string](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set CCIP admin: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func (tokenBurnMintERC20Transparent) Transfer(b operations.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	// NOTE: BurnMintERC20Transparent inherits from a standard ERC20Upgradeable implementation
	// with an ABI-compatible transfer(address,uint256), so the plain ERC20 transfer helper works
	// unchanged against the proxy address.
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

// Deploy performs the composite deploy described in CCIP-13516: deploy the
// BurnMintERC20Transparent implementation, then deploy a TransparentUpgradeableProxy pointing at
// it with ABI-encoded initialize(...) call data as the proxy constructor's `_data` argument - OZ's
// ERC1967Proxy delegatecalls that to the implementation immediately during construction, so the
// proxy is deployed and initialized atomically in one transaction. The returned AddressRef points
// at the proxy address - that's what token/pool refs resolve to, per the ticket's deploy shape.
func (tokenBurnMintERC20Transparent) Deploy(b operations.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	maxSupply := big.NewInt(0)
	if in.Supply != nil {
		maxSupply = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.Supply), in.Decimals)
	}

	preMint := big.NewInt(0)
	if in.PreMint != nil {
		preMint = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.PreMint), in.Decimals)
	}

	// Proxy admin / upgrade authority. By the time DeployTokenInput reaches this adapter, the
	// TokenExpansion changeset has already resolved ExternalAdmin to the chain timelock when it
	// was left unset (see deployment/tokens/token_expansion.go). We still refuse to deploy with a
	// zero-address proxy admin here (e.g. direct/unit-test callers that bypass that changeset),
	// since TransparentUpgradeableProxy's ProxyAdmin uses OZ Ownable, which cannot have its owner
	// set to the zero address, and doing so would permanently strand the proxy without an upgrader.
	if !common.IsHexAddress(in.ExternalAdmin) {
		return datastore.AddressRef{}, nil, fmt.Errorf("invalid or empty external admin address %q: BurnMintERC20TransparentToken requires a resolved proxy admin (ExternalAdmin-else-timelock)", in.ExternalAdmin)
	}
	proxyAdmin := common.HexToAddress(in.ExternalAdmin)

	implRef, err := contract.MaybeDeployContract(b, burn_mint_erc20_transparent.DeployImplementation, chain,
		contract.DeployInput[burn_mint_erc20_transparent.ImplConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_transparent.ImplementationContractType, *utils.Version_1_0_0),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args:           burn_mint_erc20_transparent.ImplConstructorArgs{},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20Transparent implementation: %w", err)
	}
	implAddr, err := datastore_utils_evm.ToEVMAddress(implRef)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("invalid implementation address reference: %w", err)
	}

	// defaultAdmin mirrors what BurnMintERC20's constructor implicitly does for msg.sender (the
	// deployer): it becomes the DEFAULT_ADMIN_ROLE holder, the default ccipAdmin, and the preMint
	// recipient. The generic DeployToken sequence then moves preMint/ccipAdmin to their intended
	// holders (Transfer/SetCCIPAdmin) exactly as it does for BurnMintERC20.
	initData, err := burn_mint_erc20_transparent.EncodeInitializeCallData(burn_mint_erc20_transparent.InitializeArgs{
		Name:         in.Name,
		Symbol:       in.Symbol,
		Decimals:     in.Decimals,
		MaxSupply:    maxSupply,
		PreMint:      preMint,
		DefaultAdmin: chain.DeployerKey.From,
	})
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to encode initialize call data: %w", err)
	}

	proxyRef, err := contract.MaybeDeployContract(b, burn_mint_erc20_transparent.DeployProxy, chain,
		contract.DeployInput[burn_mint_erc20_transparent.ProxyConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_transparent.ContractType, *utils.Version_1_0_0),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: burn_mint_erc20_transparent.ProxyConstructorArgs{
				Logic:        implAddr,
				InitialOwner: proxyAdmin,
				Data:         initData,
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy TransparentUpgradeableProxy: %w", err)
	}

	return proxyRef, nil, nil
}
