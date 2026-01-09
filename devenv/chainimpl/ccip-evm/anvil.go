package ccip_evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/mcms"
)

// ProcessMCMSProposalsWithTimelockForAnvil processes MCMS timelock proposals by sending transactions from the specified timelock addresses.
// This is particularly useful for anvil chains with auto-impersonation enabled so that transactions can be sent without needing private keys.
func ProcessMCMSProposalsWithTimelockForAnvil(ctx context.Context, bcsInput []*blockchain.Input, props []mcms.TimelockProposal) error {
	l := zerolog.Ctx(ctx)
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
			rpcURL := bc.Out.Nodes[0].ExternalWSUrl
			if rpcURL == "" {
				rpcURL = bc.Out.Nodes[0].ExternalHTTPUrl
				l.Info().Str("URL", rpcURL).Msg("Using HTTP URL for ETH client (HTTP-only mode)")
			}
			c, _, _, err := ETHClient(ctx, rpcURL, &GasSettings{
				FeeCapMultiplier: 2,
				TipCapMultiplier: 2,
			})
			if err != nil {
				return fmt.Errorf("could not create basic eth client: %w", err)
			}
			err = checkFundingAndFundIfNeeded(ctx, c, tlAddr, 0.2)
			if err != nil {
				return fmt.Errorf("could not fund timelock address %s: %w", tlAddr, err)
			}
			fromAddress := common.HexToAddress(tlAddr)
			for _, tx := range op.Transactions {
				nonce, err := c.PendingNonceAt(ctx, fromAddress)
				if err != nil {
					return err
				}
				feeCap, err := c.SuggestGasPrice(context.Background())
				if err != nil {
					return err
				}
				tipCap, err := c.SuggestGasTipCap(context.Background())
				if err != nil {
					return err
				}
				toAddress := common.HexToAddress(tx.To)
				ethTx := types.NewTx(&types.DynamicFeeTx{
					ChainID:   chainId,
					Nonce:     nonce,
					To:        &toAddress,
					Gas:       DefaultNativeTransferGasLimit,
					GasFeeCap: feeCap,
					GasTipCap: tipCap,
					Data:      tx.Data,
				})
				l.Info().Uint64("chainID", chainId.Uint64()).
					Uint64("nonce", nonce).
					Str("to", tx.To).
					Msg("Sending transaction from timelock")
				err = c.SendTransaction(ctx, ethTx)
				if err != nil {
					return fmt.Errorf("could not send transaction: %w", err)
				}
				if _, err := bind.WaitMined(ctx, c, ethTx); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func checkFundingAndFundIfNeeded(ctx context.Context, c *ethclient.Client, recipientAddress string, amountInEth float64) error {
	l := zerolog.Ctx(ctx)
	recipient := common.HexToAddress(recipientAddress)
	amount := new(big.Float).Mul(big.NewFloat(amountInEth), big.NewFloat(1e18))
	amountWei, _ := amount.Int(nil)
	balance, err := c.BalanceAt(ctx, recipient, nil)
	if err != nil {
		return fmt.Errorf("could not get balance for address %s: %w", recipientAddress, err)
	}
	if balance.Cmp(amountWei) >= 0 {
		l.Info().
			Str("Addr", recipientAddress).
			Str("Balance", balance.String()).
			Msg("No funding needed for address")
		return nil
	}
	l.Info().
		Str("Addr", recipientAddress).
		Str("Balance", balance.String()).
		Str("Needed", amountWei.String()).
		Msg("Funding address")
	// Use default Anvil key for local chain 1337, otherwise use PRIVATE_KEY env var
	privateKey := getNetworkPrivateKey()
	if err := FundNodeEIP1559(ctx, c, privateKey, recipientAddress, amountInEth); err != nil {
		return fmt.Errorf("failed to fund CL nodes on src chain: %w", err)
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
