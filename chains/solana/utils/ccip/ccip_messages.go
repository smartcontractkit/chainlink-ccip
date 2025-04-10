package ccip

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/crypto/sha3"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

const GenericExtraArgsV2Tag = "181dcf10"
const SVMExtraArgsV1Tag = "1f3b3aba"

var leafDomainSeparator = [32]byte{}

func HashCommitReport(ctx [2][32]byte, report ccip_offramp.CommitInput) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
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
	return hash.Sum(nil), nil
}

var reportSequence uint64 = 1

func CreateReportContext(sequence uint64) [2][32]byte {
	return [2][32]byte{
		config.ConfigDigest,
		[32]byte(binary.BigEndian.AppendUint64(config.Empty24Byte[:], sequence)),
	}
}

func ParseSequenceNumber(ctx [2][32]byte) uint64 {
	return binary.BigEndian.Uint64(ctx[1][24:])
}

func ReportSequence() uint64 {
	return reportSequence
}

func NextCommitReportContext() [2][32]byte {
	reportSequence++
	return CreateReportContext(reportSequence)
}

func CreateNextMessage(ctx context.Context, solanaGoClient *rpc.Client, remainingAccounts []solana.PublicKey) (ccip_offramp.Any2SVMRampMessage, [32]byte, error) {
	nextSeq, err := NextSequenceNumber(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA)
	if err != nil {
		return ccip_offramp.Any2SVMRampMessage{}, [32]byte{}, err
	}
	msg := CreateDefaultMessageWith(config.EvmChainSelector, nextSeq)

	hash, err := HashAnyToSVMMessage(msg, config.OnRampAddress, remainingAccounts)
	return msg, [32]byte(hash), err
}

func NextSequenceNumber(ctx context.Context, solanaGoClient *rpc.Client, sourceChainStatePDA solana.PublicKey) (uint64, error) {
	var chainStateAccount ccip_offramp.SourceChain
	err := common.GetAccountDataBorshInto(ctx, solanaGoClient, sourceChainStatePDA, config.DefaultCommitment, &chainStateAccount)
	return chainStateAccount.State.MinSeqNr, err
}

func CreateDefaultMessageWith(sourceChainSelector uint64, sequenceNumber uint64) ccip_offramp.Any2SVMRampMessage {
	sourceHash, _ := hex.DecodeString("4571dc5d4711693551f54a96307bf71121e2a1abd21d8ae04b8e05f447821064")
	var messageID [32]byte
	copy(messageID[:], sourceHash)

	message := ccip_offramp.Any2SVMRampMessage{
		Header: ccip_offramp.RampMessageHeader{
			MessageId:           messageID,
			SourceChainSelector: sourceChainSelector,
			DestChainSelector:   config.SvmChainSelector,
			SequenceNumber:      sequenceNumber,
			Nonce:               0,
		},
		Sender: []byte{1, 2, 3},
		Data:   []byte{4, 5, 6},
		ExtraArgs: ccip_offramp.Any2SVMRampExtraArgs{
			ComputeUnits:     1000,
			IsWritableBitmap: GenerateBitMapForIndexes([]int{0, 1}),
		},
	}
	return message
}

// Remaining accounts is passed separately as they're conceptually part of the message so they must be hashed alongside it,
// but they are not embedded in the message itself, as it would be redundant with `remaining_accounts`.
func MakeAnyToSVMMessage(tokenReceiver solana.PublicKey, chainSelector uint64, solanaChainSelector uint64, data []byte, msgAccounts []solana.PublicKey) (ccip_offramp.Any2SVMRampMessage, [32]byte, error) {
	msg := CreateDefaultMessageWith(chainSelector, 1)
	msg.Header.DestChainSelector = solanaChainSelector
	msg.TokenReceiver = tokenReceiver
	msg.Data = data

	hash, err := HashAnyToSVMMessage(msg, config.OnRampAddress, msgAccounts)
	msg.Header.MessageId = [32]byte(hash)
	return msg, msg.Header.MessageId, err
}

func HashAnyToSVMMessage(msg ccip_offramp.Any2SVMRampMessage, onRampAddress []byte, msgAccounts []solana.PublicKey) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()

	hash.Write(leafDomainSeparator[:])
	hash.Write([]byte("Any2SVMMessageHashV1"))

	if err := binary.Write(hash, binary.BigEndian, msg.Header.SourceChainSelector); err != nil {
		return nil, err
	}
	if err := binary.Write(hash, binary.BigEndian, msg.Header.DestChainSelector); err != nil {
		return nil, err
	}
	// Push OnRamp Size to ensure that the hash is unique
	if err := binary.Write(hash, binary.BigEndian, uint16(len(onRampAddress))); err != nil { //nolint:gosec
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
	if err := binary.Write(hash, binary.BigEndian, uint16(len(msg.Sender))); err != nil { //nolint:gosec
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

	for _, acc := range msgAccounts {
		if _, err := hash.Write(acc[:]); err != nil {
			return nil, err
		}
	}

	return hash.Sum(nil), nil
}

// merkleFrom computes the Merkle root from a slice of byte slices.
func MerkleFrom(leaves [][32]byte) ([32]byte, error) {
	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), leaves)
	if err != nil {
		return [32]byte{}, err
	}
	return tree.Root(), nil
}

func HashSVMToAnyMessage(msg ccip_router.SVM2AnyRampMessage) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()

	hash.Write(leafDomainSeparator[:])
	hash.Write([]byte("SVM2AnyMessageHashV1"))

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
	if err := binary.Write(hash, binary.BigEndian, msg.FeeValueJuels); err != nil {
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

// GenerateBitMapForIndexes generates a bitmap for the given indexes.
func GenerateBitMapForIndexes(indexes []int) uint64 {
	var bitmap uint64

	for _, index := range indexes {
		bitmap |= 1 << index
	}

	return bitmap
}

func SerializeExtraArgs(data interface{}, tag string) ([]byte, error) {
	tagBytes, err := hex.DecodeString(tag)
	if err != nil {
		return nil, err
	}
	v, err := bin.MarshalBorsh(data)
	return append(tagBytes, v...), err
}

func DeserializeExtraArgs(obj interface{}, data []byte, tag string) error {
	tagBytes, err := hex.DecodeString(tag)
	if err != nil {
		return err
	}

	if !bytes.Equal(data[:4], tagBytes) {
		return fmt.Errorf("Mismatched tag: %s != %s", hex.EncodeToString(data[:4]), tag)
	}

	err = bin.UnmarshalBorsh(obj, data[4:])
	return err
}
