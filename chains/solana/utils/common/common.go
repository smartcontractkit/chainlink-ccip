package common

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var ZeroAddress = [32]byte{}

func MakeRandom32ByteArray() [32]byte {
	a := make([]byte, 32)
	if _, err := rand.Read(a); err != nil {
		panic(err) // should never panic but check in case
	}
	return [32]byte(a)
}

func Uint64ToLE(chain uint64) []byte {
	chainLE := make([]byte, 8)
	binary.LittleEndian.PutUint64(chainLE, chain)
	return chainLE
}

func ToPadded64Bytes(input []byte) (result [64]byte) {
	if len(input) > 64 {
		panic("input is too long")
	}
	copy(result[:], input[:])
	return result
}

func To28BytesLE(value uint64) [28]byte {
	le := make([]byte, 28)
	binary.LittleEndian.PutUint64(le, value)
	return [28]byte(le)
}

func To28BytesBE(value uint64) [28]byte {
	be := make([]byte, 28)
	binary.BigEndian.PutUint64(be[20:], value)
	return [28]byte(be)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Discriminator(namespace, name string) []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s:%s", namespace, name)))
	return h.Sum(nil)[:8]
}

func GetBlockTime(ctx context.Context, client *rpc.Client, commitment rpc.CommitmentType) (*solana.UnixTimeSeconds, error) {
	block, err := client.GetBlockHeight(ctx, commitment)
	if err != nil {
		return nil, fmt.Errorf("failed to get block height: %w", err)
	}

	blockTime, err := client.GetBlockTime(ctx, block)
	if err != nil {
		return nil, fmt.Errorf("failed to get block time: %w", err)
	}

	return blockTime, nil
}
