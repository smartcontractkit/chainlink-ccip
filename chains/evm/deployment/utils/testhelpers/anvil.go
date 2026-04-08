package testhelpers

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
)

func AnvilRPCCall(rpcURL string, method string, params []any) error {
	body := map[string]any{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal %s request: %w", method, err)
	}
	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("%s rpc post: %w", method, err)
	}
	defer resp.Body.Close()
	var rpcResp struct {
		Error json.RawMessage `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return fmt.Errorf("decode %s response: %w", method, err)
	}
	if rpcResp.Error != nil {
		return fmt.Errorf("%s rpc error: %s", method, rpcResp.Error)
	}
	return nil
}

// AnvilImpersonateAccount must be called before eth_sendTransaction with from=addr on Foundry Anvil;
// otherwise the node returns "No Signer available" (-32602).
func AnvilImpersonateAccount(rpcURL, addr string) error {
	return AnvilRPCCall(rpcURL, "anvil_impersonateAccount", []any{addr})
}

func SendImpersonatedTx(ec *ethclient.Client, rpcURL string, from, to string, data []byte) error {
	if err := AnvilImpersonateAccount(rpcURL, from); err != nil {
		return fmt.Errorf("anvil impersonate %s: %w", from, err)
	}
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
	txParams := map[string]any{
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
		err = SetImpersonatedBalance(rpcURL, from, new(big.Int).Mul(requiredBal, big.NewInt(2)))
		if err != nil {
			return fmt.Errorf("set impersonated balance: %w", err)
		}
	}

	rpcReq := map[string]any{
		"jsonrpc": "2.0",
		"method":  "eth_sendTransaction",
		"params":  []any{txParams},
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

func SetImpersonatedBalance(rpcURL, addr string, weiBalance *big.Int) error {
	body := map[string]any{
		"jsonrpc": "2.0",
		"method":  "anvil_setBalance",
		"params":  []any{addr, fmt.Sprintf("0x%x", weiBalance)},
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
