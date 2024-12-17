package mcms

import (
	"bytes"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/eth"
)

type McmConfigArgs struct {
	MultisigName    [32]uint8
	SignerAddresses [][20]uint8
	SignerGroups    []byte
	GroupQuorums    [32]uint8
	GroupParents    [32]uint8
	ClearRoot       bool
}

func NewValidMcmConfig(msigName [32]byte, signerPrivateKeys []string, signerGroups []byte, quorums []uint8, parents []uint8, clearRoot bool) (*McmConfigArgs, error) {
	if len(signerGroups) == 0 {
		return nil, fmt.Errorf("signerGroups cannot be empty")
	}

	signers, err := eth.GetEvmSigners(signerPrivateKeys)
	if err != nil {
		return nil, fmt.Errorf("failed to get test EVM signers: %w", err)
	}

	if len(signers) != len(signerGroups) {
		return nil, fmt.Errorf("number of signers (%d) does not match length of signerGroups (%d)", len(signers), len(signerGroups))
	}

	signerAddresses := make([][20]uint8, len(signers))
	for i, signer := range signers {
		signerAddresses[i] = signer.Address
	}

	var groupQuorums [32]uint8
	var groupParents [32]uint8

	copy(groupQuorums[:], quorums)
	copy(groupParents[:], parents)

	// Create new config vars to ensure atomic test configs
	newSignerAddresses := make([][20]uint8, len(signerAddresses))
	copy(newSignerAddresses, signerAddresses)

	newSignerGroups := make([]byte, len(signerGroups))
	copy(newSignerGroups, signerGroups)

	newGroupQuorums := groupQuorums
	newGroupParents := groupParents
	newClearRoot := clearRoot

	config := &McmConfigArgs{
		MultisigName: msigName,
	}
	config.SignerAddresses = newSignerAddresses
	config.SignerGroups = newSignerGroups
	config.GroupQuorums = newGroupQuorums
	config.GroupParents = newGroupParents
	config.ClearRoot = newClearRoot
	return config, nil
}

func FindInSortedList(list []solana.PublicKey, target solana.PublicKey) (int, bool) {
	return slices.BinarySearchFunc(list, target, func(a, b solana.PublicKey) int {
		return bytes.Compare(a.Bytes(), b.Bytes())
	})
}

func SafeToUint8(n int) (uint8, error) {
	if n < 0 || n > 255 {
		return 0, fmt.Errorf("value %d is outside uint8 range [0,255]", n)
	}
	return uint8(n), nil
}

func SafeToUint32(n int) (uint32, error) {
	if n < 0 || n > math.MaxUint32 {
		return 0, fmt.Errorf("value %d is outside uint32 range [0,%d]", n, math.MaxUint32)
	}
	return uint32(n), nil
}

func PadString32(input string) ([32]byte, error) {
	var result [32]byte
	inputBytes := []byte(input)
	inputLen := len(inputBytes)
	if inputLen > 32 {
		return result, errors.New("input string exceeds 32 bytes")
	}
	startPos := 32 - inputLen
	copy(result[startPos:], inputBytes)
	return result, nil
}

func UnpadString32(input [32]byte) string {
	startPos := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			startPos = i
			break
		}
	}
	return string(input[startPos:])
}

// simple salt generator that uses the current Unix timestamp(in mills)
func SimpleSalt() ([32]byte, error) {
	var salt [32]byte
	now := time.Now().UnixMilli()
	if now < 0 {
		return salt, fmt.Errorf("negative timestamp: %d", now)
	}
	// unix timestamp in millseconds
	binary.BigEndian.PutUint64(salt[:8], uint64(now))
	// Next 8 bytes: Crypto random
	randBytes := make([]byte, 8)
	if _, err := crypto_rand.Read(randBytes); err != nil {
		return salt, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	copy(salt[8:16], randBytes)
	return salt, nil
}
