package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/tokenimpl"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var RevokeTokenAdminRole = cldf_ops.NewSequence(
	"evm-revoke-token-admin-role",
	cciputils.Version_1_0_0,
	"Revoke an admin role from an EVM token",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.RevokeTokenAdminRoleSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found among provided EVM chains", input.ChainSelector)
		}
		if input.AdminAddress == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("admin address is required")
		}
		if !common.IsHexAddress(input.AdminAddress) {
			return sequences.OnChainOutput{}, fmt.Errorf("admin address %q is not a valid hex address", input.AdminAddress)
		}
		adminAddress := common.HexToAddress(input.AdminAddress)
		if adminAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("admin address cannot be the zero address")
		}

		tokenAddress, err := datastore_utils_evm.ToEVMAddress(input.TokenRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token ref to EVM address: %w", err)
		}
		if tokenAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("token address cannot be the zero address")
		}

		tokenImpl, ok := tokenimpl.Get(cldf_deployment.ContractType(input.TokenRef.Type))
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type %q for token ref %s", input.TokenRef.Type, datastore_utils.SprintRef(input.TokenRef))
		}
		adminRoleToken, ok := tokenImpl.(tokenimpl.AdminRoleToken)
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("token type %q does not support admin role checks", input.TokenRef.Type)
		}

		ctx := b.GetContext()
		hasAdminRole, err := adminRoleToken.HasAdminRole(ctx, chain, tokenAddress, adminAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for %s on token %s: %w", adminAddress.Hex(), tokenAddress.Hex(), err)
		}
		if !hasAdminRole {
			b.Logger.Infof("admin %s does not have an admin role on token %s on chain %d; skipping revoke", adminAddress.Hex(), tokenAddress.Hex(), input.ChainSelector)
			return sequences.OnChainOutput{}, nil
		}

		candidates := []common.Address{chain.DeployerKey.From}
		if common.IsHexAddress(input.TimelockAddress) {
			candidates = append(candidates, common.HexToAddress(input.TimelockAddress))
		}
		for _, user := range chain.Users {
			if user != nil {
				candidates = append(candidates, user.From)
			}
		}

		hasRemainingAdmin := false
		for _, candidate := range candidates {
			if candidate == (common.Address{}) || candidate == adminAddress {
				continue
			}
			hasAdminRole, err := adminRoleToken.HasAdminRole(ctx, chain, tokenAddress, candidate)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for candidate %s on token %s: %w", candidate.Hex(), tokenAddress.Hex(), err)
			}
			if hasAdminRole {
				hasRemainingAdmin = true
				break
			}
		}
		if !hasRemainingAdmin {
			return sequences.OnChainOutput{}, fmt.Errorf("refusing to revoke admin role from %s on token %s because no remaining admin could be confirmed", adminAddress.Hex(), tokenAddress.Hex())
		}

		revokeWrites, err := tokenImpl.RevokeAdminRole(b, chain, tokenAddress, adminAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke admin role from %s on token %s: %w", adminAddress.Hex(), tokenAddress.Hex(), err)
		}
		batchOp, err := evm_contract.NewBatchOperationFromWrites(revokeWrites)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for token admin role revocation: %w", err)
		}
		if len(batchOp.Transactions) == 0 {
			return sequences.OnChainOutput{}, nil
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)
