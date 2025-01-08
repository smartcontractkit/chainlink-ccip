package mcms

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"time"

	"golang.org/x/crypto/sha3"
)

func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
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
