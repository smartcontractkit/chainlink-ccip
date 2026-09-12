package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const onRampEventABI = `[{"type":"event","name":"CCIPMessageSent","inputs":[
{"name":"destChainSelector","type":"uint64","indexed":true},
{"name":"sender","type":"address","indexed":true},
{"name":"messageId","type":"bytes32","indexed":true},
{"name":"feeToken","type":"address","indexed":false},
{"name":"tokenAmountBeforeTokenPoolFees","type":"uint256","indexed":false},
{"name":"encodedMessage","type":"bytes","indexed":false},
{"name":"receipts","type":"tuple[]","indexed":false,"components":[
  {"name":"issuer","type":"address"},{"name":"destGasLimit","type":"uint32"},{"name":"destBytesOverhead","type":"uint32"},
  {"name":"feeTokenAmount","type":"uint256"},{"name":"extraArgs","type":"bytes"}]},
{"name":"verifierBlobs","type":"bytes[]","indexed":false}]}]`

const offRampABI = `[{"type":"function","name":"execute","inputs":[
{"name":"encodedMessage","type":"bytes"},{"name":"ccvs","type":"address[]"},{"name":"verifierResults","type":"bytes[]"},
{"name":"gasLimitOverride","type":"uint32"}],"outputs":[]}]`

const heliosABI = `[
{"type":"function","name":"latestExecutionBlockNumber","inputs":[],"outputs":[{"type":"uint256"}],"stateMutability":"view"},
{"type":"function","name":"executionBlockHashes","inputs":[{"type":"uint256"}],"outputs":[{"type":"bytes32"}],"stateMutability":"view"},
{"type":"function","name":"executionReceiptsRoots","inputs":[{"type":"uint256"}],"outputs":[{"type":"bytes32"}],"stateMutability":"view"}]`

var (
	ccipMessageSentTopic       = crypto.Keccak256Hash([]byte("CCIPMessageSent(uint64,address,bytes32,address,uint256,bytes,(address,uint32,uint32,uint256,bytes)[],bytes[])"))
	executionStateChangedTopic = crypto.Keccak256Hash([]byte("ExecutionStateChanged(uint64,uint64,bytes32,uint8,bytes)"))
)

func mustABI(src string) abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(src))
	if err != nil {
		panic(err)
	}
	return parsed
}

// MessageInfo describes one CCIP message as sent on the source chain.
type MessageInfo struct {
	MessageID      common.Hash
	EncodedMessage []byte
	OnRamp         common.Address
	BlockNumber    uint64
	TxIndex        uint
	LogIndex       uint8
}

// loadMessage reads the CCIPMessageSent log out of the send transaction receipt.
func loadMessage(ctx context.Context, src *ethclient.Client, txHash common.Hash) (*MessageInfo, error) {
	receipt, err := src.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("receipt %s: %w", txHash, err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("send transaction %s reverted", txHash)
	}
	parsed := mustABI(onRampEventABI)
	for i, lg := range receipt.Logs {
		if len(lg.Topics) != 4 || lg.Topics[0] != ccipMessageSentTopic {
			continue
		}
		fields, err := parsed.Events["CCIPMessageSent"].Inputs.NonIndexed().Unpack(lg.Data)
		if err != nil {
			return nil, fmt.Errorf("unpack CCIPMessageSent: %w", err)
		}
		if i > 0xff {
			return nil, fmt.Errorf("log index %d does not fit in one byte", i)
		}
		return &MessageInfo{
			MessageID:      lg.Topics[3],
			EncodedMessage: fields[2].([]byte),
			OnRamp:         lg.Address,
			BlockNumber:    receipt.BlockNumber.Uint64(),
			TxIndex:        receipt.TransactionIndex,
			LogIndex:       uint8(i),
		}, nil
	}
	return nil, fmt.Errorf("no CCIPMessageSent log in transaction %s", txHash)
}

// HeliosReader reads anchored blocks from the SP1Helios contract on the destination chain.
type HeliosReader struct {
	client  *ethclient.Client
	address common.Address
	abi     abi.ABI
}

func newHeliosReader(client *ethclient.Client, address common.Address) *HeliosReader {
	return &HeliosReader{client: client, address: address, abi: mustABI(heliosABI)}
}

func (h *HeliosReader) call(ctx context.Context, method string, args ...interface{}) ([]interface{}, error) {
	data, err := h.abi.Pack(method, args...)
	if err != nil {
		return nil, err
	}
	out, err := h.client.CallContract(ctx, ethereum.CallMsg{To: &h.address, Data: data}, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}
	return h.abi.Unpack(method, out)
}

func (h *HeliosReader) LatestExecutionBlockNumber(ctx context.Context) (uint64, error) {
	out, err := h.call(ctx, "latestExecutionBlockNumber")
	if err != nil {
		return 0, err
	}
	return out[0].(*big.Int).Uint64(), nil
}

func (h *HeliosReader) ExecutionBlockHash(ctx context.Context, blockNumber uint64) (common.Hash, error) {
	out, err := h.call(ctx, "executionBlockHashes", new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return common.Hash{}, err
	}
	return out[0].([32]byte), nil
}

// waitForAnchor polls Helios until a block at or above messageBlock is anchored and returns the lowest one seen.
func waitForAnchor(ctx context.Context, helios *HeliosReader, messageBlock uint64, poll time.Duration) (uint64, common.Hash, error) {
	seen := map[uint64]bool{}
	for {
		latest, err := helios.LatestExecutionBlockNumber(ctx)
		if err != nil {
			return 0, common.Hash{}, err
		}
		seen[latest] = true
		best := uint64(0)
		for n := range seen {
			if n >= messageBlock && (best == 0 || n < best) {
				best = n
			}
		}
		if best != 0 {
			hash, err := helios.ExecutionBlockHash(ctx, best)
			if err != nil {
				return 0, common.Hash{}, err
			}
			if hash != (common.Hash{}) {
				return best, hash, nil
			}
		}
		fmt.Printf("helios latest anchored block %d, waiting for >= %d\n", latest, messageBlock)
		select {
		case <-ctx.Done():
			return 0, common.Hash{}, ctx.Err()
		case <-time.After(poll):
		}
	}
}

// executeOnOffRamp submits OffRamp.execute and returns the execution state from ExecutionStateChanged.
func executeOnOffRamp(ctx context.Context, dst *ethclient.Client, key string, offRamp, ccv common.Address, encodedMessage, verifierResults []byte) (uint8, common.Hash, error) {
	parsed := mustABI(offRampABI)
	data, err := parsed.Pack("execute", encodedMessage, []common.Address{ccv}, [][]byte{verifierResults}, uint32(0))
	if err != nil {
		return 0, common.Hash{}, err
	}
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(key, "0x"))
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("private key: %w", err)
	}
	from := crypto.PubkeyToAddress(privateKey.PublicKey)

	msg := ethereum.CallMsg{From: from, To: &offRamp, Data: data}
	if _, err := dst.CallContract(ctx, msg, nil); err != nil {
		return 0, common.Hash{}, fmt.Errorf("execute simulation reverted: %w", err)
	}
	gas, err := dst.EstimateGas(ctx, msg)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("estimate gas: %w", err)
	}
	chainID, err := dst.ChainID(ctx)
	if err != nil {
		return 0, common.Hash{}, err
	}
	nonce, err := dst.PendingNonceAt(ctx, from)
	if err != nil {
		return 0, common.Hash{}, err
	}
	gasPrice, err := dst.SuggestGasPrice(ctx)
	if err != nil {
		return 0, common.Hash{}, err
	}
	tx := types.NewTx(&types.LegacyTx{Nonce: nonce, To: &offRamp, Gas: gas * 13 / 10, GasPrice: new(big.Int).Mul(gasPrice, big.NewInt(2)), Data: data})
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return 0, common.Hash{}, err
	}
	if err := dst.SendTransaction(ctx, signed); err != nil {
		return 0, common.Hash{}, fmt.Errorf("send execute: %w", err)
	}
	fmt.Printf("execute sent: %s\n", signed.Hash())

	var receipt *types.Receipt
	for receipt == nil {
		time.Sleep(3 * time.Second)
		receipt, err = dst.TransactionReceipt(ctx, signed.Hash())
		if err != nil && err != ethereum.NotFound {
			return 0, signed.Hash(), err
		}
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return 0, signed.Hash(), fmt.Errorf("execute transaction reverted")
	}
	for _, lg := range receipt.Logs {
		if len(lg.Topics) == 4 && lg.Topics[0] == executionStateChangedTopic && len(lg.Data) >= 32 {
			return lg.Data[31], signed.Hash(), nil
		}
	}
	return 0, signed.Hash(), fmt.Errorf("no ExecutionStateChanged log in execute receipt")
}
