// Package registrations registers all known EVM token contract strategies
// with the singleton strategy.Registry at init time. Adapters that need
// per-token-type behavior pull in this package via a blank import; all
// known token types become available in one line.
//
// Adding a new EVM token contract type means adding one strategy struct
// and one Register call in init below. No edits to pool adapters or
// deploy sequences are required.
package registrations

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	drip_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

func init() {
	r := strategy.GetRegistry()
	r.RegisterEVM(erc20Strategy{})
	r.RegisterEVM(burnMintERC20Strategy{})
	r.RegisterEVM(burnMintERC20WithDripStrategy{})
	r.RegisterEVM(burnMintERC20WithDripV150Strategy{})
	r.RegisterEVM(tip20Strategy{})
}

// ---------- ERC20 (plain, not CCIP-aware) ----------

type erc20Strategy struct{}

func (erc20Strategy) ContractType() deployment.ContractType { return erc20.ContractType }

func (erc20Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{}
}

func (erc20Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	ref, err := contract.MaybeDeployContract(b, erc20.Deploy, chain, contract.DeployInput[erc20.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *common_utils.Version_1_0_0),
		ChainSelector:  chain.Selector,
		Args: erc20.ConstructorArgs{
			Name:   in.Name,
			Symbol: in.Symbol,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy ERC20 token: %w", err)
	}
	return ref, nil, nil
}

func (erc20Strategy) GrantPoolRoles(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address, _ uint64) ([]contract.WriteOutput, error) {
	return nil, nil
}

func (erc20Strategy) GrantExternalAdmin(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address, _ uint64) ([]contract.WriteOutput, error) {
	return nil, nil
}

// ---------- BurnMintERC20 (v1.0.0) ----------

type burnMintERC20Strategy struct{}

func (burnMintERC20Strategy) ContractType() deployment.ContractType {
	return burn_mint_erc20.ContractType
}

func (burnMintERC20Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             true,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (burnMintERC20Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	maxSupply, preMint := scaledSupplyAndPreMint(in)
	ref, err := contract.MaybeDeployContract(b, burn_mint_erc20.Deploy, chain, contract.DeployInput[burn_mint_erc20.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20.ContractType, *common_utils.Version_1_0_0),
		ChainSelector:  chain.Selector,
		Args: burn_mint_erc20.ConstructorArgs{
			Name:      in.Name,
			Symbol:    in.Symbol,
			Decimals:  in.Decimals,
			MaxSupply: maxSupply,
			PreMint:   preMint,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20 token: %w", err)
	}
	return ref, nil, nil
}

func (burnMintERC20Strategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMMintAndBurnRoles(b, chain, token, pool, chainSelector)
}

func (burnMintERC20Strategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMDefaultAdminRole(b, chain, token, externalAdmin, chainSelector)
}

// ---------- BurnMintERC20WithDrip (v1.0.0) ----------

type burnMintERC20WithDripStrategy struct{}

func (burnMintERC20WithDripStrategy) ContractType() deployment.ContractType {
	return burn_mint_erc20_with_drip.ContractType
}

func (burnMintERC20WithDripStrategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             true,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (burnMintERC20WithDripStrategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	maxSupply, preMint := scaledSupplyAndPreMint(in)
	ref, err := contract.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain, contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *common_utils.Version_1_0_0),
		ChainSelector:  chain.Selector,
		Args: burn_mint_erc20_with_drip.ConstructorArgs{
			Name:      in.Name,
			Symbol:    in.Symbol,
			Decimals:  in.Decimals,
			MaxSupply: maxSupply,
			PreMint:   preMint,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20WithDrip token: %w", err)
	}
	return ref, nil, nil
}

func (burnMintERC20WithDripStrategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMMintAndBurnRoles(b, chain, token, pool, chainSelector)
}

func (burnMintERC20WithDripStrategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMDefaultAdminRole(b, chain, token, externalAdmin, chainSelector)
}

// ---------- BurnMintERC20WithDrip (v1.5.0) ----------
//
// Pre-mint is unsupported because the v1.5.0 constructor takes neither
// supply nor decimals; matches the historical tokenSupportsPreMint table.

type burnMintERC20WithDripV150Strategy struct{}

func (burnMintERC20WithDripV150Strategy) ContractType() deployment.ContractType {
	return drip_v150.ContractType
}

func (burnMintERC20WithDripV150Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             false,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (burnMintERC20WithDripV150Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	ref, err := contract.MaybeDeployContract(b, drip_v150.Deploy, chain, contract.DeployInput[drip_v150.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(drip_v150.ContractType, *drip_v150.Version),
		ChainSelector:  chain.Selector,
		Args: drip_v150.ConstructorArgs{
			Name:   in.Name,
			Symbol: in.Symbol,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20WithDrip (v1.5.0) token: %w", err)
	}
	return ref, nil, nil
}

func (burnMintERC20WithDripV150Strategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMMintAndBurnRoles(b, chain, token, pool, chainSelector)
}

func (burnMintERC20WithDripV150Strategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMDefaultAdminRole(b, chain, token, externalAdmin, chainSelector)
}

// ---------- TIP-20 ----------
//
// Tempo-only token; deployed via a factory sequence rather than
// MaybeDeployContract. CCIPAdmin and pre-mint do not apply.

type tip20Strategy struct{}

func (tip20Strategy) ContractType() deployment.ContractType { return tip20.ContractType }

func (tip20Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           false,
		SupportsPreMint:             false,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (tip20Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	// Initial admin must be the deployer so subsequent ops (e.g. GrantIssuerRole)
	// pass IsAllowedCaller; ExternalAdmin receives DEFAULT_ADMIN_ROLE in a
	// follow-up grant performed by the orchestrating sequence.
	report, err := cldf_ops.ExecuteSequence(b, tip20.Deploy, chain, tip20.FactoryDeployArgs{
		QuoteToken: common.Address{},
		Currency:   "",
		Salt:       [32]byte{},
		Symbol:     in.Symbol,
		Admin:      chain.DeployerKey.From,
		Name:       in.Name,
	})
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy TIP20 token via factory: %w", err)
	}
	if len(report.Output.Addresses) == 0 {
		return datastore.AddressRef{}, nil, errors.New("no address returned from TIP20 factory deployment")
	}
	return report.Output.Addresses[0], nil, nil
}

func (tip20Strategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, tip20.GrantIssuerRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       token,
		Args:          pool,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant TIP-20 issuer role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (tip20Strategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, tip20.GrantAdminRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       token,
		Args:          externalAdmin,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant TIP-20 admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

// ---------- shared helpers ----------

func scaledSupplyAndPreMint(in tokensapi.DeployTokenInput) (*big.Int, *big.Int) {
	maxSupply := big.NewInt(0)
	if in.Supply != nil {
		maxSupply = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.Supply), in.Decimals)
	}
	preMint := big.NewInt(0)
	if in.PreMint != nil {
		preMint = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.PreMint), in.Decimals)
	}
	return maxSupply, preMint
}

// grantBnMMintAndBurnRoles is shared by all BnM-family strategies.
// Historically the BnM, BnM+Drip (v1.0.0), and BnM+Drip (v1.5.0) types
// all dispatch to burn_mint_erc20.GrantMintAndBurnRoles (the v1.0.0 op);
// preserved here verbatim.
func grantBnMMintAndBurnRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantMintAndBurnRoles, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       token,
		Args:          pool,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant mint and burn roles: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func grantBnMDefaultAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(token, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
	}
	role, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: b.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role constant: %w", err)
	}
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantAdminRole, chain, contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
		ChainSelector: chainSelector,
		Address:       token,
		Args: burn_mint_erc20.RoleAssignment{
			Role: role,
			To:   externalAdmin,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant default admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}
