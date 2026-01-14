package blockchainutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/mcms"
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
				err = sendImpersonatedTx(ec, rpcURL, tlAddr, tx.To, tx.Data)
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

func sendImpersonatedTx(ec *ethclient.Client, rpcURL string, from, to string, data []byte) error {
	fromAddr := common.HexToAddress(from)
	toAddr := common.HexToAddress(to)
	msg := ethereum.CallMsg{
		From: fromAddr,
		To:   &toAddr,
		Data: data,
	}
	estGas, err := ec.EstimateGas(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("estimate gas: %w", err)
	}

	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("suggest gas: %w", err)
	}
	log.Info().
		Str("from", from).
		Str("to", to).
		Uint64("estGas", estGas).
		Str("gasPrice", gasPrice.String()).
		Msg("Impersonated tx details")

	// 🔹 2. Build JSON-RPC params for eth_sendTransaction
	txParams := map[string]interface{}{
		"from": from,
		"to":   to,
		"data": fmt.Sprintf("0x%x", data),
	}

	bal, err := ec.BalanceAt(context.Background(), fromAddr, nil)
	if err != nil {
		return fmt.Errorf("fetch balance: %w", err)
	}
	requiredBal := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(estGas))
	if bal.Cmp(requiredBal) < 0 {
		err = setImpersonatedBalance(rpcURL, from, new(big.Int).Mul(requiredBal, big.NewInt(2)))
		if err != nil {
			return fmt.Errorf("set impersonated balance: %w", err)
		}
	}

	rpcReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendTransaction",
		"params":  []interface{}{txParams},
		"id":      1,
	}

	reqBody, _ := json.Marshal(rpcReq)
	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("rpc post: %w", err)
	}
	defer resp.Body.Close()

	var rpcResp struct {
		Result string          `json:"result"`
		Error  json.RawMessage `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	if rpcResp.Error != nil {
		return fmt.Errorf("rpc error: %s", rpcResp.Error)
	}
	log.Info().
		Str("txHash", rpcResp.Result).
		Msg("Impersonated tx sent successfully")
	return nil
}

func setImpersonatedBalance(rpcURL, addr string, weiBalance *big.Int) error {
	body := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "anvil_setBalance",
		"params":  []interface{}{addr, fmt.Sprintf("0x%x", weiBalance)},
		"id":      1,
	}
	b, _ := json.Marshal(body)
	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("anvil_setBalance call failed: %w", err)
	}
	defer resp.Body.Close()
	return nil
}
