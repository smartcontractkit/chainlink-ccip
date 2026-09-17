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

func (tokenBurnMintERC20Transparent) HasAdminRole(b operations.Bundle, chain evm.Chain, token, user common.Address) (bool, error) {
	return hasDefaultAdminRoleBurnMintERC20Transparent(b, chain, token, user)
}

func (tokenBurnMintERC20Transparent) GrantAdminRole(b operations.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return beginDefaultAdminTransferBurnMintERC20Transparent(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20Transparent) AcceptAdminRole(b operations.Bundle, chain evm.Chain, token common.Address) ([]contract.WriteOutput, error) {
	return acceptDefaultAdminTransferBurnMintERC20Transparent(b, chain, token)
}

func (tokenBurnMintERC20Transparent) PendingAdminRoleTarget(b operations.Bundle, chain evm.Chain, token common.Address) (common.Address, error) {
	return pendingDefaultAdminBurnMintERC20Transparent(b, chain, token)
}

func (tokenBurnMintERC20Transparent) GrantPoolRoles(b operations.Bundle, chain evm.Chain, token, pool, _ common.Address) ([]contract.WriteOutput, error) {
	return grantMintAndBurnRolesBurnMintERC20Transparent(b, chain, token, pool)
}

func (tokenBurnMintERC20Transparent) SetCCIPAdmin(b operations.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	return setCCIPAdminBurnMintERC20Transparent(b, chain, token, ccipAdmin)
}

func (tokenBurnMintERC20Transparent) Transfer(b operations.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	// NOTE: BurnMintERC20Transparent inherits from a standard ERC20Upgradeable implementation
	// with an ABI-compatible transfer(address,uint256), so the plain ERC20 transfer helper works
	// unchanged against the proxy address.
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

// Deploy performs the 3-step composite deploy described in CCIP-13516: deploy the
// BurnMintERC20Transparent implementation, deploy a TransparentUpgradeableProxy pointing at
// it, then call initialize(...) through the proxy. The returned AddressRef points at the
// proxy address - that's what token/pool refs resolve to, per the ticket's deploy shape.
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

	proxyRef, err := contract.MaybeDeployContract(b, burn_mint_erc20_transparent.DeployProxy, chain,
		contract.DeployInput[burn_mint_erc20_transparent.ProxyConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_transparent.ContractType, *utils.Version_1_0_0),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: burn_mint_erc20_transparent.ProxyConstructorArgs{
				Logic:        implAddr,
				InitialOwner: proxyAdmin,
				Data:         []byte{},
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy TransparentUpgradeableProxy: %w", err)
	}
	proxyAddr, err := datastore_utils_evm.ToEVMAddress(proxyRef)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("invalid proxy address reference: %w", err)
	}

	// defaultAdmin mirrors what BurnMintERC20's constructor implicitly does for msg.sender (the
	// deployer): it becomes the DEFAULT_ADMIN_ROLE holder, the default ccipAdmin, and the preMint
	// recipient. The generic DeployToken sequence then moves preMint/ccipAdmin to their intended
	// holders (Transfer/SetCCIPAdmin) exactly as it does for BurnMintERC20.
	initReport, err := operations.ExecuteOperation(b, burn_mint_erc20_transparent.Initialize, chain,
		contract.FunctionInput[burn_mint_erc20_transparent.InitializeArgs]{
			ChainSelector: chain.Selector,
			Address:       proxyAddr,
			Args: burn_mint_erc20_transparent.InitializeArgs{
				Name:         in.Name,
				Symbol:       in.Symbol,
				Decimals:     in.Decimals,
				MaxSupply:    maxSupply,
				PreMint:      preMint,
				DefaultAdmin: chain.DeployerKey.From,
			},
		},
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to initialize BurnMintERC20Transparent proxy: %w", err)
	}

	return proxyRef, []contract.WriteOutput{initReport.Output}, nil
}
