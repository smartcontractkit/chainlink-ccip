package ccip

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

var leafDomainSeparator = [32]byte{}

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

var reportSequence uint64 = 1

func CreateReportContext(sequence uint64) [3][32]byte {
	return [3][32]byte{
		config.ConfigDigest,
		[32]byte(binary.BigEndian.AppendUint64(config.Empty24Byte[:], sequence)),
		common.MakeRandom32ByteArray(),
	}
}

func ParseSequenceNumber(ctx [3][32]byte) uint64 {
	return binary.BigEndian.Uint64(ctx[1][24:])
}

func ReportSequence() uint64 {
	return reportSequence
}

func NextCommitReportContext() [3][32]byte {
	reportSequence++
	return CreateReportContext(reportSequence)
}

func CreateNextMessage(ctx context.Context, solanaGoClient *rpc.Client) (ccip_router.Any2SolanaRampMessage, [32]byte, error) {
	nextSeq, err := NextSequenceNumber(ctx, solanaGoClient, config.EvmSourceChainStatePDA)
	if err != nil {
		return ccip_router.Any2SolanaRampMessage{}, [32]byte{}, err
	}
	msg := CreateDefaultMessageWith(config.EvmChainSelector, nextSeq)

	hash, err := HashEvmToSolanaMessage(msg, config.OnRampAddress)
	return msg, [32]byte(hash), err
}

func NextSequenceNumber(ctx context.Context, solanaGoClient *rpc.Client, sourceChainStatePDA solana.PublicKey) (uint64, error) {
	var chainStateAccount ccip_router.SourceChain
	err := common.GetAccountDataBorshInto(ctx, solanaGoClient, sourceChainStatePDA, config.DefaultCommitment, &chainStateAccount)
	return chainStateAccount.State.MinSeqNr, err
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
		Sender:        []byte{1, 2, 3},
		Data:          []byte{4, 5, 6},
		LogicReceiver: config.CcipLogicReceiver,
		ExtraArgs: ccip_router.SolanaExtraArgs{
			ComputeUnits:     1000,
			IsWritableBitmap: 3, // [true, true, false]
			Accounts: []solana.PublicKey{
				config.ReceiverExternalExecutionConfigPDA,
				config.ReceiverTargetAccountPDA,
				solana.SystemProgramID,
			},
		},
		OnRampAddress: config.OnRampAddress,
	}
	return message
}

func MakeEvmToSolanaMessage(tokenReceiver solana.PublicKey, logicReceiver solana.PublicKey, evmChainSelector uint64, solanaChainSelector uint64, data []byte) (ccip_router.Any2SolanaRampMessage, [32]byte, error) {
	msg := CreateDefaultMessageWith(evmChainSelector, 1)
	msg.Header.DestChainSelector = solanaChainSelector
	msg.TokenReceiver = tokenReceiver
	msg.LogicReceiver = logicReceiver
	msg.Data = data

	hash, err := HashEvmToSolanaMessage(msg, config.OnRampAddress)
	msg.Header.MessageId = [32]byte(hash)
	return msg, msg.Header.MessageId, err
}

func HashEvmToSolanaMessage(msg ccip_router.Any2SolanaRampMessage, onRampAddress []byte) ([]byte, error) {
	hash := sha256.New()

	hash.Write(leafDomainSeparator[:])
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
	if _, err := hash.Write(msg.TokenReceiver[:]); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.LogicReceiver[:]); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.SequenceNumber); err != nil {
		return nil, err
	}
	extraArgsBytes, borshErr := bin.MarshalBorsh(msg.ExtraArgs)
	if borshErr != nil {
		return nil, borshErr
	}
	if _, err := hash.Write(extraArgsBytes); err != nil {
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
	tokenAmountsBytes, err := bin.MarshalBorsh(msg.TokenAmounts)
	if err != nil {
		return nil, err
	}
	if _, err := hash.Write(tokenAmountsBytes); err != nil {
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

func HashSolanaToAnyMessage(msg ccip_router.Solana2AnyRampMessage) ([]byte, error) {
	hash := sha256.New()

	hash.Write(leafDomainSeparator[:])
	hash.Write([]byte("Solana2AnyMessageHashV1"))

	if err := binary.Write(hash, binary.BigEndian, msg.Header.SourceChainSelector); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.DestChainSelector); err != nil {
		return nil, err
	}
	if _, err := hash.Write(config.CcipRouterProgram[:]); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.Sender[:]); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.SequenceNumber); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.Nonce); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.FeeToken[:]); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.FeeTokenAmount); err != nil {
		return nil, err
	}
	if _, err := hash.Write([]byte{uint8(len(msg.Receiver))}); err != nil { //nolint:gosec
		return nil, err
	}
	if _, err := hash.Write(msg.Receiver); err != nil {
		return nil, err
	}
	dataLen := uint16(len(msg.Data)) //nolint:gosec // max U16 larger than solana transaction size
	if err := binary.Write(hash, binary.BigEndian, dataLen); err != nil {
		return nil, err
	}
	if _, err := hash.Write(msg.Data); err != nil {
		return nil, err
	}
	tokenAmountsBytes, borshErr := bin.MarshalBorsh(msg.TokenAmounts)
	if borshErr != nil {
		return nil, borshErr
	}
	if _, err := hash.Write(tokenAmountsBytes); err != nil {
		return nil, err
	}
	extraArgsBytes, err := bin.MarshalBorsh(msg.ExtraArgs)
	if err != nil {
		return nil, err
	}
	if _, err := hash.Write(extraArgsBytes); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
