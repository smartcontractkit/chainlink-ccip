package blockchainutils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/mcms"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/testhelpers"
)

// ProcessMCMSProposalsWithTimelockForAnvil processes MCMS timelock proposals by sending transactions from the specified timelock addresses.
// This is particularly useful for anvil chains with auto-impersonation enabled so that transactions can be sent without needing private keys.
func ProcessMCMSProposalsWithTimelockForAnvil(ctx context.Context, bcsInput []*blockchain.Input, props []mcms.TimelockProposal) error {
	bcs, err := getAnvilBlockchainsMapBySelector(bcsInput)
	if err != nil {
		return fmt.Errorf("could not get anvil blockchains map: %w", err)
	}
	for _, prop := range props {
		for _, op := range prop.Operations {
			bc, exists := bcs[uint64(op.ChainSelector)]
			if !exists {
				return fmt.Errorf("blockchain client for chain selector %d not found", op.ChainSelector)
			}
			tlAddr := prop.TimelockAddresses[op.ChainSelector]
			chainId, success := big.NewInt(0).SetString(bc.ChainID, 10)
			if !success {
				return fmt.Errorf("invalid chain ID: %s", bc.ChainID)
			}
			// Use WS URL if available, otherwise fallback to HTTP URL for HTTP-only mode
			rpcURL := bc.Out.Nodes[0].ExternalHTTPUrl
			if rpcURL == "" {
				return fmt.Errorf("no http RPC URL found for chain with selector %d", op.ChainSelector)
			}
			ec, err := ethclient.Dial(rpcURL)
			if err != nil {
				return fmt.Errorf("dial ethclient: %w", err)
			}
			defer ec.Close()
			for _, tx := range op.Transactions {
				err = testhelpers.SendImpersonatedTx(ctx, ec, rpcURL, tlAddr, tx.To, tx.Data)
				if err != nil {
					return fmt.Errorf("could not send impersonated tx to %s on chain with id %d from %s: %w", tx.To, chainId.Uint64(), tlAddr, err)
				}
			}
		}
	}
	return nil
}

func getAnvilBlockchainsMapBySelector(bcs []*blockchain.Input) (map[uint64]*blockchain.Input, error) {
	result := make(map[uint64]*blockchain.Input)
	for _, bc := range bcs {
		if bc.Type == "anvil" {
			networkInfo, err := chainsel.GetChainDetailsByChainIDAndFamily(bc.ChainID, chainsel.FamilyEVM)
			if err != nil {
				return nil, err
			}
			result[networkInfo.ChainSelector] = bc
		}
	}
	return result, nil
}
