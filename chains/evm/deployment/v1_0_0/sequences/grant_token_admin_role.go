package sequences

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
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
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var GrantTokenAdminRole = cldf_ops.NewSequence(
	"evm-grant-token-admin-role",
	cciputils.Version_1_0_0,
	"Grant an admin role on an EVM token",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.GrantTokenAdminRoleSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found among provided EVM chains", input.ChainSelector)
		}
		if !datastore_utils.IsAddressRefFullyPopulated(input.TokenRef) {
			return sequences.OnChainOutput{}, fmt.Errorf("token ref is incomplete: %v", input.TokenRef)
		}

		tokenAddress, err := datastore_utils_evm.ToEVMAddress(input.TokenRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token ref to EVM address: %w", err)
		}
		if tokenAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("token address cannot be the zero address")
		}

		if !common.IsHexAddress(input.AdminAddress) {
			return sequences.OnChainOutput{}, fmt.Errorf("admin address %q is not a valid hex address", input.AdminAddress)
		}
		grantAddress := common.HexToAddress(input.AdminAddress)
		if grantAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("admin address cannot be the zero address")
		}

		timelockAddress := common.Address{}
		if input.TimelockAddress != "" {
			if !common.IsHexAddress(input.TimelockAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("timelock address %q is not a valid hex address", input.TimelockAddress)
			}
			timelockAddress = common.HexToAddress(input.TimelockAddress)
		}

		tokenImpl, ok := tokenimpl.Get(cldf_deployment.ContractType(input.TokenRef.Type))
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type %q for token address %q on chain %d", input.TokenRef.Type, input.TokenRef.Address, input.TokenRef.ChainSelector)
		}
		if !tokenImpl.Capabilities().SupportsAdminRole {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s on chain %d with type %s does not support admin role management", tokenAddress.Hex(), input.ChainSelector, input.TokenRef.Type)
		}

		timelockHasAdminRole := false
		if timelockAddress != (common.Address{}) {
			if hasRole, err := tokenImpl.HasAdminRole(b, chain, tokenAddress, timelockAddress); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for timelock %s on token %s: %w", timelockAddress.Hex(), tokenAddress.Hex(), err)
			} else {
				timelockHasAdminRole = hasRole
			}
		}
		deployerHasAdminRole := false
		if chain.DeployerKey.From != (common.Address{}) {
			if hasRole, err := tokenImpl.HasAdminRole(b, chain, tokenAddress, chain.DeployerKey.From); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for deployer %s on token %s: %w", chain.DeployerKey.From.Hex(), tokenAddress.Hex(), err)
			} else {
				deployerHasAdminRole = hasRole
			}
		}
		if !timelockHasAdminRole && !deployerHasAdminRole {
			b.Logger.Warnf("neither timelock %s nor deployer %s has an admin role on token %s on chain %d; skipping grant since there is no account with sufficient permissions to perform the operation", timelockAddress.Hex(), chain.DeployerKey.From.Hex(), tokenAddress.Hex(), input.ChainSelector)
			return sequences.OnChainOutput{}, nil
		}

		targetHasAdminRole, err := tokenImpl.HasAdminRole(b, chain, tokenAddress, grantAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check admin role for target %s on token %s: %w", grantAddress.Hex(), tokenAddress.Hex(), err)
		}
		if targetHasAdminRole {
			b.Logger.Infof("target %s already has an admin role on token %s on chain %d; skipping grant", grantAddress.Hex(), tokenAddress.Hex(), input.ChainSelector)
			return sequences.OnChainOutput{}, nil
		}

		writes, err := tokenImpl.GrantAdminRole(b, chain, tokenAddress, grantAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant admin role to %s on token %s: %w", grantAddress.Hex(), tokenAddress.Hex(), err)
		}

		if len(writes) == 0 {
			b.Logger.Infof("no writes generated for granting admin role to %s on token %s on chain %d; skipping operation", grantAddress.Hex(), tokenAddress.Hex(), input.ChainSelector)
			return sequences.OnChainOutput{}, nil
		}

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for token admin role grant: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)
