package sequences

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy"
	// Defensive blank import: ensures all known EVM token strategies are
	// registered when this package is imported directly (e.g. by tests)
	// rather than transitively via an adapter package.
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy/registrations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
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
		chain := chains.EVMChains()[input.ChainSelector]

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

		strat, ok := strategy.GetRegistry().GetEVM(input.Type)
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type: %s", input.Type)
		}
		caps := strat.Capabilities()

		tokenRef, deployWrites, err := strat.Deploy(b, chain, input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		addresses := []datastore.AddressRef{tokenRef}
		writes := append([]contract.WriteOutput(nil), deployWrites...)
		tokenAddr := common.HexToAddress(tokenRef.Address)

		if caps.SupportsPreMint && preMint.Cmp(big.NewInt(0)) > 0 && len(input.Senders) > 0 {
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

		if input.CCIPAdmin != "" && caps.SupportsCCIPAdmin {
			setCCIPAdminReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.SetCCIPAdmin, chain, contract.FunctionInput[string]{
				ChainSelector: chain.Selector,
				Address:       tokenAddr,
				Args:          ccipAdmin.Hex(),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set CCIP admin: %w", err)
			}
			writes = append(writes, setCCIPAdminReport.Output)
		}

		if input.ExternalAdmin != "" && caps.SupportsAdminRole {
			adminWrites, err := strat.GrantExternalAdmin(b, chain, tokenAddr, externalAdmin, chain.Selector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to grant admin role to %s: %w", input.ExternalAdmin, err)
			}
			writes = append(writes, adminWrites...)
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
