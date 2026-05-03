package sequences

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/tokenimpl"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DeployToken = cldf_ops.NewSequence(
	"deploy-token",
	common_utils.Version_1_0_0,
	"Deploy given type of token contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found among provided chains", input.ChainSelector)
		}

		preMint := big.NewInt(0)
		if input.PreMint != nil {
			preMint = tokenapi.ScaleTokenAmount(new(big.Int).SetUint64(*input.PreMint), input.Decimals)
		}

		externalAdmin := common.Address{}
		if input.ExternalAdmin != "" && !common.IsHexAddress(input.ExternalAdmin) {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid external admin address: %s", input.ExternalAdmin)
		} else {
			externalAdmin = common.HexToAddress(input.ExternalAdmin)
		}

		ccipAdmin := common.Address{}
		if input.CCIPAdmin != "" && !common.IsHexAddress(input.CCIPAdmin) {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid CCIP admin address: %s", input.CCIPAdmin)
		} else {
			ccipAdmin = common.HexToAddress(input.CCIPAdmin)
		}

		tokenImpl, ok := tokenimpl.Get(input.Type)
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type: %s", input.Type)
		}
		tokenRefr, deployWrites, err := tokenImpl.Deploy(b, chain, input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		tokenAddr, err := datastore_utils_evm.ToEVMAddress(tokenRefr)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid token address reference: %w", err)
		}

		caps := tokenImpl.Capabilities()
		recv := common.Address{}
		if len(input.Senders) >= 1 && preMint.Cmp(big.NewInt(0)) > 0 && caps.SupportsPreMint {
			address := input.Senders[0]
			if !common.IsHexAddress(address) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid pre-mint recipient address: %s", address)
			}
			recv = common.HexToAddress(address)
			if recv == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("pre-mint recipient address cannot be the zero address")
			}
			if len(input.Senders) != 1 {
				b.Logger.Warnf("Multiple sender addresses provided, but adapter only supports one. Only the first address will receive the tokens: %s", address)
			}
		}

		writes := append([]contract.WriteOutput{}, deployWrites...)
		if recv != (common.Address{}) && caps.SupportsPreMint {
			transferWrites, err := tokenImpl.Transfer(b, chain, tokenAddr, recv, preMint)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to mint pre-mint tokens: %w", err)
			}
			writes = append(writes, transferWrites...)
		}
		if input.ExternalAdmin != "" && caps.SupportsAdminRole {
			grantWrites, err := tokenImpl.GrantAdminRole(b, chain, tokenAddr, externalAdmin)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to grant admin role to %s: %w", input.ExternalAdmin, err)
			}
			writes = append(writes, grantWrites...)
		}
		if input.CCIPAdmin != "" && caps.SupportsCCIPAdmin {
			adminWrites, err := tokenImpl.SetCCIPAdmin(b, chain, tokenAddr, ccipAdmin)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set CCIP admin: %w", err)
			}
			writes = append(writes, adminWrites...)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{tokenRefr},
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
