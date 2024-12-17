package eth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"golang.org/x/crypto/sha3"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
)

type Signer struct {
	PrivateKey []byte
	PubKey     []byte
	Address    [20]byte
}

func (s *Signer) Sign(msg []byte) (mcm.Signature, error) {
	if len(s.PrivateKey) != 32 {
		return mcm.Signature{}, errors.New("invalid private key length")
	}

	privateKey := secp256k1.PrivKeyFromBytes(s.PrivateKey)

	if len(msg) != 32 {
		return mcm.Signature{}, errors.New("message must be a 32-byte hash")
	}

	signature := ecdsa.SignCompact(privateKey, msg, false)

	vByte := signature[0] // V is already adjusted with +27 in SignCompact
	var rBytes, sBytes [32]byte
	copy(rBytes[:], signature[1:33])
	copy(sBytes[:], signature[33:65])

	sig := mcm.Signature{
		V: vByte,
		R: rBytes,
		S: sBytes,
	}
	return sig, nil
}

func (s Signer) String() string {
	return "0x" + hex.EncodeToString(s.Address[:])
}

func GetEvmSigners(privateKeys []string) ([]Signer, error) {
	signers := make([]Signer, 0, len(privateKeys))
	for _, key := range privateKeys {
		signer, err := GetSignerFromPk(key)
		if err != nil {
			return nil, fmt.Errorf("failed to get signer from private key: %v", err)
		}
		signers = append(signers, signer)
	}

	sort.Slice(signers, func(i, j int) bool {
		return string(signers[i].Address[:]) < string(signers[j].Address[:])
	})
	return signers, nil
}

func GetSignerFromPk(key string) (Signer, error) {
	pkBytes, err := hex.DecodeString(key)
	if err != nil {
		return Signer{}, fmt.Errorf("failed to decode private key: %v", err)
	}

	privateKey := secp256k1.PrivKeyFromBytes(pkBytes)
	publicKey := privateKey.PubKey()

	pubKeyBytes := publicKey.SerializeUncompressed()
	hash := Keccak256(pubKeyBytes[1:]) // skip the leading 0x04 byte
	var address [20]byte
	copy(address[:], hash[12:])

	return Signer{
		PrivateKey: pkBytes,
		PubKey:     pubKeyBytes,
		Address:    address,
	}, nil
}

// TODO: helper functions for n of signer tests on set_config - will be removed after refactor
// N is the order of the secp256k1 curve
var secp256k1N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)

// Returns a slice of private keys as raw hexadecimal strings without "0x" prefix
func GenerateEthPrivateKeys(n int) ([]string, error) {
	if n <= 0 {
		return nil, fmt.Errorf("number of keys must be positive")
	}

	privateKeys := make([]string, n)

	for i := 0; i < n; i++ {
		// Generate a valid private key
		pk, err := generateValidPrivateKey()
		if err != nil {
			return nil, fmt.Errorf("failed to generate private key: %v", err)
		}

		// Convert to hex without 0x prefix
		pkBytes := pk.Bytes()
		// Ensure the key is exactly 32 bytes by left-padding with zeros if necessary
		paddedPk := make([]byte, 32)
		copy(paddedPk[32-len(pkBytes):], pkBytes)
		privateKeys[i] = hex.EncodeToString(paddedPk)
	}

	return privateKeys, nil
}

func generateValidPrivateKey() (*big.Int, error) {
	zero := big.NewInt(0)

	for {
		// Generate 32 random bytes
		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			return nil, err
		}

		k := new(big.Int).SetBytes(b)

		// Check if private key is valid:
		// 1. k must be greater than 0
		// 2. k must be less than the curve order (secp256k1N)
		if k.Cmp(zero) > 0 && k.Cmp(secp256k1N) < 0 {
			return k, nil
		}
		// If invalid, continue loop to generate another key
	}
}

func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}
