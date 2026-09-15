// zkccv is the local verifier and executor for the Succinct ZK CCV proof of concept.
//
// For one send transaction it waits until SP1Helios anchors a block at or above the message block, builds the
// witness, and executes the message on the destination OffRamp.
package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const versionTagPreimage = "SuccinctZKVerifier 2.0.0"

func versionTag() [4]byte {
	var tag [4]byte
	copy(tag[:], crypto.Keccak256([]byte(versionTagPreimage))[:4])
	return tag
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: zkccv <fixture|run> [flags]")
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "fixture":
		err = fixtureCmd(os.Args[2:])
	case "run":
		err = runCmd(os.Args[2:])
	default:
		err = fmt.Errorf("unknown command %q", os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Fixture is the JSON consumed by the Foundry test.
type Fixture struct {
	AnchorBlockNumber  uint64 `json:"anchorBlockNumber"`
	AnchorBlockHash    string `json:"anchorBlockHash"`
	MessageBlockNumber uint64 `json:"messageBlockNumber"`
	ReceiptsRoot       string `json:"receiptsRoot"`
	MessageID          string `json:"messageId"`
	OnRamp             string `json:"onRamp"`
	EncodedMessage     string `json:"encodedMessage"`
	Envelope           string `json:"envelope"`
}

// fixtureCmd builds a witness for a past send transaction against a chosen anchor block and writes it as JSON.
func fixtureCmd(args []string) error {
	fs := flag.NewFlagSet("fixture", flag.ExitOnError)
	srcRPC := fs.String("src-rpc", "", "source chain RPC URL")
	txHash := fs.String("tx", "", "send transaction hash on the source chain")
	anchorOffset := fs.Uint64("anchor-offset", 3, "anchor block = message block + offset")
	out := fs.String("out", "fixture.json", "output path")
	if err := fs.Parse(args); err != nil {
		return err
	}
	ctx := context.Background()
	src, err := ethclient.DialContext(ctx, *srcRPC)
	if err != nil {
		return err
	}
	info, err := loadMessage(ctx, src, common.HexToHash(*txHash))
	if err != nil {
		return err
	}
	anchor := info.BlockNumber + *anchorOffset
	anchorHeader, err := src.HeaderByNumber(ctx, new(big.Int).SetUint64(anchor))
	if err != nil {
		return err
	}
	envelope, receiptsRoot, err := buildEnvelope(ctx, src, info, anchor, anchorHeader.Hash())
	if err != nil {
		return err
	}
	fixture := Fixture{
		AnchorBlockNumber:  anchor,
		AnchorBlockHash:    anchorHeader.Hash().Hex(),
		MessageBlockNumber: info.BlockNumber,
		ReceiptsRoot:       receiptsRoot.Hex(),
		MessageID:          info.MessageID.Hex(),
		OnRamp:             info.OnRamp.Hex(),
		EncodedMessage:     "0x" + hex.EncodeToString(info.EncodedMessage),
		Envelope:           "0x" + hex.EncodeToString(envelope),
	}
	data, err := json.MarshalIndent(fixture, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(*out, data, 0o644); err != nil {
		return err
	}
	fmt.Printf("fixture written to %s: message block %d, anchor %d, envelope %d bytes\n", *out, info.BlockNumber, anchor, len(envelope))
	return nil
}

// runCmd verifies and executes one message end to end.
func runCmd(args []string) error {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	srcRPC := fs.String("src-rpc", "", "source chain RPC URL")
	dstRPC := fs.String("dst-rpc", "", "destination chain RPC URL")
	txHash := fs.String("tx", "", "send transaction hash on the source chain")
	helios := fs.String("helios", "", "SP1Helios address on the destination chain")
	offRamp := fs.String("offramp", "", "OffRamp address on the destination chain")
	ccv := fs.String("ccv", "", "ZK verifier resolver address on the destination chain")
	poll := fs.Duration("poll", 20*time.Second, "Helios polling interval")
	if err := fs.Parse(args); err != nil {
		return err
	}
	key := os.Getenv("POC_PK")
	if key == "" {
		return fmt.Errorf("POC_PK is not set")
	}
	ctx := context.Background()
	src, err := ethclient.DialContext(ctx, *srcRPC)
	if err != nil {
		return err
	}
	dst, err := ethclient.DialContext(ctx, *dstRPC)
	if err != nil {
		return err
	}

	info, err := loadMessage(ctx, src, common.HexToHash(*txHash))
	if err != nil {
		return err
	}
	fmt.Printf("message %s in source block %d, tx index %d, log index %d\n", info.MessageID, info.BlockNumber, info.TxIndex, info.LogIndex)

	anchor, anchorHash, err := waitForAnchor(ctx, newHeliosReader(dst, common.HexToAddress(*helios)), info.BlockNumber, *poll)
	if err != nil {
		return err
	}
	fmt.Printf("anchor block %d hash %s\n", anchor, anchorHash)

	envelope, receiptsRoot, err := buildEnvelope(ctx, src, info, anchor, anchorHash)
	if err != nil {
		return err
	}
	fmt.Printf("witness built: receipts root %s, envelope %d bytes\n", receiptsRoot, len(envelope))

	state, execTx, err := executeOnOffRamp(ctx, dst, key, common.HexToAddress(*offRamp), common.HexToAddress(*ccv), info.EncodedMessage, envelope)
	if err != nil {
		return err
	}
	fmt.Printf("executed in %s with state %d (2 = SUCCESS)\n", execTx, state)
	return nil
}

// buildEnvelope assembles the header chain and receipt proof and encodes them.
func buildEnvelope(ctx context.Context, src *ethclient.Client, info *MessageInfo, anchor uint64, anchorHash common.Hash) ([]byte, common.Hash, error) {
	headers, err := buildHeaderChain(ctx, src, anchor, info.BlockNumber, anchorHash)
	if err != nil {
		return nil, common.Hash{}, err
	}
	nodes, receiptsRoot, err := buildReceiptProof(ctx, src, info.BlockNumber, info.TxIndex)
	if err != nil {
		return nil, common.Hash{}, err
	}
	messageHeader, err := src.HeaderByNumber(ctx, new(big.Int).SetUint64(info.BlockNumber))
	if err != nil {
		return nil, common.Hash{}, err
	}
	if receiptsRoot != messageHeader.ReceiptHash {
		return nil, common.Hash{}, fmt.Errorf("rebuilt receipts root %s differs from header root %s", receiptsRoot, messageHeader.ReceiptHash)
	}
	if info.TxIndex > 0xffff {
		return nil, common.Hash{}, fmt.Errorf("tx index %d does not fit in two bytes", info.TxIndex)
	}
	witness := &Witness{
		AnchorBlockNumber: anchor,
		Headers:           headers,
		TxIndex:           uint16(info.TxIndex),
		LogIndex:          info.LogIndex,
		ProofNodes:        nodes,
	}
	envelope, err := witness.Encode(versionTag())
	if err != nil {
		return nil, common.Hash{}, err
	}
	return envelope, receiptsRoot, nil
}
