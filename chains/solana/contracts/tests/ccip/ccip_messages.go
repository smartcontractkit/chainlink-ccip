package contracts

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
)

func HashCommitReport(ctx [3][32]byte, report ccip_router.CommitInput) ([]byte, error) {
	hash := sha256.New()
	encodedReport, err := bin.MarshalBorsh(report)
	if err != nil {
		return nil, err
	}

	reportLen := uint16(len(encodedReport)) //nolint:gosec // max U16 larger than solana transaction size
	if err := binary.Write(hash, binary.LittleEndian, reportLen); err != nil {
		return nil, err
	}
	if _, err := hash.Write(encodedReport); err != nil {
		return nil, err
	}
	if _, err := hash.Write(ctx[0][:]); err != nil {
		return nil, err
	}
	if _, err := hash.Write(ctx[1][:]); err != nil {
		return nil, err
	}
	if _, err := hash.Write(ctx[2][:]); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func CreateNextMessage(ctx context.Context, solanaGoClient *rpc.Client, t *testing.T) (ccip_router.Any2SolanaRampMessage, [32]byte) {
	nextSeq := NextSequenceNumber(ctx, solanaGoClient, config.EvmChainStatePDA, t)
	msg := CreateDefaultMessageWith(config.EvmChainSelector, nextSeq)

	hash, err := HashEvmToSolanaMessage(msg, config.OnRampAddress)
	require.NoError(t, err)
	return msg, [32]byte(hash)
}

func NextSequenceNumber(ctx context.Context, solanaGoClient *rpc.Client, chainStatePDA solana.PublicKey, t *testing.T) uint64 {
	var chainStateAccount ccip_router.ChainState
	err := utils.GetAccountDataBorshInto(ctx, solanaGoClient, chainStatePDA, config.DefaultCommitment, &chainStateAccount)
	require.NoError(t, err)
	return chainStateAccount.SourceChain.State.MinSeqNr
}

func CreateDefaultMessageWith(sourceChainSelector uint64, sequenceNumber uint64) ccip_router.Any2SolanaRampMessage {
	sourceHash, _ := hex.DecodeString("4571dc5d4711693551f54a96307bf71121e2a1abd21d8ae04b8e05f447821064")
	var messageID [32]byte
	copy(messageID[:], sourceHash)

	message := ccip_router.Any2SolanaRampMessage{
		Header: ccip_router.RampMessageHeader{
			MessageId:           messageID,
			SourceChainSelector: sourceChainSelector,
			DestChainSelector:   config.SolanaChainSelector,
			SequenceNumber:      sequenceNumber,
			Nonce:               0,
		},
		Sender:   []byte{1, 2, 3},
		Data:     []byte{4, 5, 6},
		Receiver: config.ReceiverExternalExecutionConfigPDA,
		ExtraArgs: ccip_router.SolanaExtraArgs{
			ComputeUnits: 1000,
			Accounts: []ccip_router.SolanaAccountMeta{
				{Pubkey: config.CcipReceiverProgram},
				{Pubkey: config.ReceiverTargetAccountPDA, IsWritable: true},
				{Pubkey: solana.SystemProgramID, IsWritable: false},
			},
		},
	}
	return message
}

func MakeEvmToSolanaMessage(t *testing.T, ccipReceiver solana.PublicKey, evmChainSelector uint64, solanaChainSelector uint64, data []byte) (ccip_router.Any2SolanaRampMessage, [32]byte) {
	msg := CreateDefaultMessageWith(evmChainSelector, 1)
	msg.Header.DestChainSelector = solanaChainSelector
	msg.Receiver = ccipReceiver
	msg.Data = data

	hash, err := HashEvmToSolanaMessage(msg, config.OnRampAddress)
	require.NoError(t, err)
	msg.Header.MessageId = [32]byte(hash)
	return msg, msg.Header.MessageId
}

func HashEvmToSolanaMessage(msg ccip_router.Any2SolanaRampMessage, onRampAddress []byte) ([]byte, error) {
	hash := sha256.New()

	hash.Write([]byte("Any2SolanaMessageHashV1"))

	if err := binary.Write(hash, binary.BigEndian, msg.Header.SourceChainSelector); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.DestChainSelector); err != nil {
		return nil, err
	}
	// Push OnRamp Size to ensure that the hash is unique
	if _, err := hash.Write([]byte{uint8(len(onRampAddress))}); err != nil { //nolint:gosec
		return nil, err
	}
	if _, err := hash.Write(onRampAddress); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.Header.MessageId[:]); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.Receiver[:]); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.SequenceNumber); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.ExtraArgs.ComputeUnits); err != nil {
		return nil, err
	}
	// Push accounts size
	if _, err := hash.Write([]byte{uint8(len(msg.ExtraArgs.Accounts))}); err != nil { //nolint:gosec
		return nil, err
	}
	accountsBytes, err := bin.MarshalBorsh(msg.ExtraArgs.Accounts)
	if err != nil {
		return nil, err
	}
	if _, err := hash.Write(accountsBytes); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.Nonce); err != nil {
		return nil, err
	}
	// Push Sender Size to ensure that the hash is unique
	if _, err := hash.Write([]byte{uint8(len(msg.Sender))}); err != nil { //nolint:gosec
		return nil, err
	}
	if _, err := hash.Write(msg.Sender); err != nil {
		return nil, err
	}
	// Push Data Size to ensure that the hash is unique
	dataLen := uint16(len(msg.Data)) //nolint:gosec // max U16 larger than solana transaction size
	if err := binary.Write(hash, binary.BigEndian, dataLen); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.Data); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// hashPair hashes two byte slices and returns the result as a byte slice.
func hashPair(a, b []byte) []byte {
	h := sha256.New()
	if bytes.Compare(a, b) < 0 {
		h.Write(a)
		h.Write(b)
	} else {
		h.Write(b)
		h.Write(a)
	}
	return h.Sum(nil)
}

// merkleFrom computes the Merkle root from a slice of byte slices.
func MerkleFrom(data [][]byte) []byte {
	if len(data) == 1 {
		return data[0]
	}

	hash := hashPair(data[0], data[1])

	for i := 2; i < len(data); i++ {
		hash = hashPair(hash, data[i])
	}

	return hash
}
