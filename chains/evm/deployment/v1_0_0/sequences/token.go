package sequences

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

func tokenSupportsAdminRole(tokenType deployment.ContractType) bool {
	switch tokenType {
	case burn_mint_erc20.ContractType,
		burn_mint_erc20_with_drip.ContractType:
		return true
	default:
		return false
	}
}

func tokenSupportsCCIPAdmin(tokenType deployment.ContractType) bool {
	switch tokenType {
	case burn_mint_erc20.ContractType,
		burn_mint_erc20_with_drip.ContractType:
		return true
	default:
		return false
	}
}

func tokenSupportsPreMint(tokenType deployment.ContractType) bool {
	switch tokenType {
	case burn_mint_erc20.ContractType, burn_mint_erc20_with_drip.ContractType:
		return true
	default:
		return false
	}
}

var DeployToken = cldf_ops.NewSequence(
	"deploy-token",
	common_utils.Version_1_0_0,
	"Deploy given type of token contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenInput) (sequences.OnChainOutput, error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		chain := chains.EVMChains()[input.ChainSelector]
		var err error
		var tokenRef datastore.AddressRef
		qualifier := input.Symbol

		maxSupply := big.NewInt(0)
		if input.Supply != nil {
			maxSupply = tokenapi.ScaleTokenAmount(new(big.Int).SetUint64(*input.Supply), input.Decimals)
		}

		preMint := big.NewInt(0)
		if input.PreMint != nil {
			preMint = tokenapi.ScaleTokenAmount(new(big.Int).SetUint64(*input.PreMint), input.Decimals)
		}

		switch input.Type {
		case erc20.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, erc20.Deploy, chain, contract.DeployInput[erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *common_utils.Version_1_0_0),
				ChainSelector:  chain.Selector,
				Args: erc20.ConstructorArgs{
					Name:   input.Name,
					Symbol: input.Symbol,
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC20 token: %w", err)
			}

		case burn_mint_erc20.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, burn_mint_erc20.Deploy, chain, contract.DeployInput[burn_mint_erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20.ContractType, *common_utils.Version_1_0_0),
				ChainSelector:  chain.Selector,
				Args: burn_mint_erc20.ConstructorArgs{
					Name:      input.Name,
					Symbol:    input.Symbol,
					Decimals:  input.Decimals,
					MaxSupply: maxSupply,
					PreMint:   preMint,
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintERC20 token: %w", err)
			}

		case burn_mint_erc20_with_drip.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain, contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *common_utils.Version_1_0_0),
				ChainSelector:  chain.Selector,
				Args: burn_mint_erc20_with_drip.ConstructorArgs{
					Name:      input.Name,
					Symbol:    input.Symbol,
					Decimals:  input.Decimals,
					MaxSupply: maxSupply,
					PreMint:   preMint,
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintERC20WithDrip token: %w", err)
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type: %s", input.Type)
		}

		tokenAddr := common.HexToAddress(tokenRef.Address)
		addresses = append(addresses, tokenRef)

		if tokenSupportsPreMint(input.Type) && preMint.Cmp(big.NewInt(0)) > 0 && len(input.Senders) > 0 {
			firstSender := input.Senders[0]
			if !common.IsHexAddress(firstSender) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid sender address: %s", firstSender)
			}
			tokReceiver := common.HexToAddress(firstSender)
			if tokReceiver == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("refusing to transfer pre-minted tokens to the zero address")
			}
			if len(input.Senders) > 1 {
				b.Logger.Warnf("Multiple senders provided but only the first one (%s) will receive the pre-minted tokens", tokReceiver.Hex())
			}
			transferReport, err := cldf_ops.ExecuteOperation(b, erc20.Transfer, chain, contract.FunctionInput[erc20.TransferArgs]{
				ChainSelector: chain.Selector,
				Address:       tokenAddr,
				Args: erc20.TransferArgs{
					Receiver: tokReceiver,
					Amount:   preMint,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer pre-minted tokens to sender %s: %w", tokReceiver.Hex(), err)
			}
			writes = append(writes, transferReport.Output)
		}

		if input.CCIPAdmin != "" && tokenSupportsCCIPAdmin(input.Type) {
			setCCIPAdminReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.SetCCIPAdmin, chain, contract.FunctionInput[string]{
				ChainSelector: chain.Selector,
				Address:       tokenAddr,
				Args:          input.CCIPAdmin,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set CCIP admin: %w", err)
			}
			writes = append(writes, setCCIPAdminReport.Output)
		}

		if input.ExternalAdmin != "" && tokenSupportsAdminRole(input.Type) {
			token, err := bnm_erc20_bindings.NewBurnMintERC20(tokenAddr, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
			}
			role, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: b.GetContext()})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get default admin role constant: %w", err)
			}

			grantReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantAdminRole, chain, contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
				ChainSelector: chain.Selector,
				Address:       tokenAddr,
				Args: burn_mint_erc20.RoleAssignment{
					Role: role,
					To:   common.HexToAddress(input.ExternalAdmin),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to grant admin role to %s: %w", input.ExternalAdmin, err)
			}
			writes = append(writes, grantReport.Output)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
