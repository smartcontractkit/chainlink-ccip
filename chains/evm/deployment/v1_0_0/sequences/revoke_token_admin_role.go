package sequences

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/tokenimpl"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var RevokeTokenAdminRole = cldf_ops.NewSequence(
	"evm-revoke-token-admin-role",
	cciputils.Version_1_0_0,
	"Revoke an admin role from an EVM token",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.RevokeTokenAdminRoleSequenceInput) (sequences.OnChainOutput, error) {
		// Validate the chain
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found among provided EVM chains", input.ChainSelector)
		}

		// Validate the token address
		tokenAddress, err := datastore_utils_evm.ToEVMAddress(input.TokenRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token ref to EVM address: %w", err)
		}
		if tokenAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("token address cannot be the zero address")
		}

		// Validate timelock address (if it exists)
		timelockAddress := common.Address{}
		if input.TimelockAddress != "" {
			if !common.IsHexAddress(input.TimelockAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("timelock address %q is not a valid hex address", input.TimelockAddress)
			} else {
				timelockAddress = common.HexToAddress(input.TimelockAddress)
			}
		}

		// Validate the fallback address (if provided)
		fallbackAddress := common.Address{}
		if input.FallbackAddress != "" {
			if !common.IsHexAddress(input.FallbackAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("fallback address %q is not a valid hex address", input.FallbackAddress)
			} else {
				fallbackAddress = common.HexToAddress(input.FallbackAddress)
			}
		}

		// If the token does not support admin role management, then we skip the revocation and return
		// no operations since proceeding would ultimately result in a failure.
		tokenImpl, ok := tokenimpl.Get(cldf_deployment.ContractType(input.TokenRef.Type))
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type %q for token address %q on chain %d", input.TokenRef.Type, input.TokenRef.Address, input.TokenRef.ChainSelector)
		}
		if !tokenImpl.Capabilities().SupportsAdminRole {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s on chain %d with type %s does not support admin role management", tokenAddress.Hex(), input.ChainSelector, input.TokenRef.Type)
		}

		// This operation will be run by either timelock or the deployer key, so we need to ensure that
		// the account running the operation has sufficient access to perform the operation. If this is
		// not the case, then we return no operations and log a warning instead of returning an error.
		timelockHasAdminRole := false
		if timelockAddress != (common.Address{}) {
			if hasRole, err := tokenImpl.HasAdminRole(b.GetContext(), chain, tokenAddress, timelockAddress); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for timelock %s on token %s: %w", timelockAddress.Hex(), tokenAddress.Hex(), err)
			} else {
				timelockHasAdminRole = hasRole
			}
		}
		deployerHasAdminRole := false
		if chain.DeployerKey.From != (common.Address{}) {
			if hasRole, err := tokenImpl.HasAdminRole(b.GetContext(), chain, tokenAddress, chain.DeployerKey.From); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for deployer %s on token %s: %w", chain.DeployerKey.From.Hex(), tokenAddress.Hex(), err)
			} else {
				deployerHasAdminRole = hasRole
			}
		}
		if !timelockHasAdminRole && !deployerHasAdminRole {
			b.Logger.Warnf("neither timelock %s nor deployer %s has an admin role on token %s on chain %d; skipping revoke since there is no account with sufficient permissions to perform the operation", timelockAddress.Hex(), chain.DeployerKey.From.Hex(), tokenAddress.Hex(), input.ChainSelector)
			return sequences.OnChainOutput{}, nil
		}

		// If the user does not provide an AdminAddress, then the top-level changeset will attempt
		// to set it to timelock. If timelock isn't found in the datastore, then we fall back back
		// to the deployer key in this sequence.
		revokeAddress := chain.DeployerKey.From
		if input.AdminAddress != "" {
			if !common.IsHexAddress(input.AdminAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("admin address %q is not a valid hex address", input.AdminAddress)
			} else {
				revokeAddress = common.HexToAddress(input.AdminAddress)
			}
		}
		if revokeAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("admin address cannot be the zero address")
		}

		// If the fallback address is unspecified OR the fallback and revoke addresses are the
		// same, then we skip the grant operation and assume the user wants to bypass all role
		// protections. If the fallback and revoke addresses differ, then we the admin role is
		// granted to the fallback address (if needed), and we proceed with revoking the admin
		// role from the revoke address. The fallback address is a safety measure that ensures
		// there is at least one account with the admin role on the token contract.
		var writes []evm_contract.WriteOutput
		if fallbackAddress != (common.Address{}) && fallbackAddress != revokeAddress {
			fallbackAccountHasAdminRole, err := tokenImpl.HasAdminRole(b.GetContext(), chain, tokenAddress, fallbackAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for fallback address %s on token %s: %w", fallbackAddress.Hex(), tokenAddress.Hex(), err)
			}
			if !fallbackAccountHasAdminRole {
				if output, err := tokenImpl.GrantAdminRole(b, chain, tokenAddress, fallbackAddress); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to grant admin role to fallback address %s on token %s: %w", fallbackAddress.Hex(), tokenAddress.Hex(), err)
				} else {
					writes = append(writes, output...)
				}
			} else {
				b.Logger.Infof("fallback address %s already has an admin role on token %s on chain %d; skipping grant", fallbackAddress.Hex(), tokenAddress.Hex(), input.ChainSelector)
			}
		} else {
			b.Logger.Warnf("no fallback address provided or fallback address is the same as revoke address for token %s on chain %d; skipping grant operation and proceeding directly to revoke (this can be dangerous if the revoke address is the only account with admin role on the token)", tokenAddress.Hex(), input.ChainSelector)
		}

		// If the account that we want to revoke admin access from does NOT have the admin role
		// then we skip the revocation to save gas and avoid unnecessary transactions.
		revokeAccountHasAdminRole, err := tokenImpl.HasAdminRole(b.GetContext(), chain, tokenAddress, revokeAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for %s on token %s: %w", revokeAddress.Hex(), tokenAddress.Hex(), err)
		}
		if revokeAccountHasAdminRole {
			if output, err := tokenImpl.RevokeAdminRole(b, chain, tokenAddress, revokeAddress); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke admin role from %s on token %s: %w", revokeAddress.Hex(), tokenAddress.Hex(), err)
			} else {
				writes = append(writes, output...)
			}
		} else {
			b.Logger.Infof("admin %s does not have an admin role on token %s on chain %d; skipping revoke", revokeAddress.Hex(), tokenAddress.Hex(), input.ChainSelector)
		}

		if len(writes) == 0 {
			b.Logger.Infof("no writes generated for revoking admin role from %s on token %s on chain %d; skipping operation", revokeAddress.Hex(), tokenAddress.Hex(), input.ChainSelector)
			return sequences.OnChainOutput{}, nil
		}

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for token admin role revocation: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)
