package ccipocr3

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

// UnknownAddress represents a raw address with an unknown encoding.
type UnknownAddress []byte

// NewUnknownAddressFromHex creates a new UnknownAddress from a hex string.
func NewUnknownAddressFromHex(s string) (UnknownAddress, error) {
	b, err := NewBytesFromString(s)
	if err != nil {
		return nil, err
	}
	return UnknownAddress(b), nil
}

// String returns the hex representation of the unknown address.
func (a UnknownAddress) String() string {
	return Bytes(a).String()
}

func (a UnknownAddress) MarshalJSON() ([]byte, error) {
	return Bytes(a).MarshalJSON()
}

func (a *UnknownAddress) UnmarshalJSON(data []byte) error {
	return (*Bytes)(a).UnmarshalJSON(data)
}

// IsZeroOrEmpty returns true if the address contains 0 bytes or if all the bytes are 0.
func (a UnknownAddress) IsZeroOrEmpty() bool {
	if len(a) == 0 {
		return true // empty
	}

	for _, b := range a {
		if b != 0 {
			return false // zero
		}
	}
	return true
}

// UnknownEncodedAddress represents an encoded address with an unknown encoding.
type UnknownEncodedAddress string

type Bytes []byte

func NewBytesFromString(s string) (Bytes, error) {
	if len(s) < 2 {
		return nil, fmt.Errorf("Bytes must be of at least length 2 (i.e, '0x' prefix): %s", s)
	}

	if !strings.HasPrefix(s, "0x") {
		return nil, fmt.Errorf("Bytes must start with '0x' prefix: %s", s)
	}

	b, err := hex.DecodeString(s[2:])
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return Bytes(b), nil
}

func (b Bytes) String() string {
	return "0x" + hex.EncodeToString(b)
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.String())), nil
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	v := string(data)
	if len(v) < 4 {
		return fmt.Errorf("bytes must be of at least length 4 (i.e, '\"0x\"'): %s", v)
	}

	// trim the start and end double quotes
	v = v[1 : len(v)-1]

	if !strings.HasPrefix(v, "0x") {
		return fmt.Errorf("bytes must start with '0x' prefix: %s", v)
	}

	// Decode everything after the '0x' prefix.
	bs, err := hex.DecodeString(v[2:])
	if err != nil {
		return fmt.Errorf("failed to decode hex: %w", err)
	}

	*b = bs
	return nil
}

type Bytes32 [32]byte

func NewBytes32FromString(s string) (Bytes32, error) {
	if len(s) > 66 { // "0x" + 64 hex chars
		return Bytes32{}, fmt.Errorf("Bytes32 must be at most 32 bytes (64 hex chars) long: %s", s)
	}

	if !strings.HasPrefix(s, "0x") {
		return Bytes32{}, fmt.Errorf("Bytes32 must start with '0x' prefix: %s", s)
	}

	b, err := hex.DecodeString(s[2:])
	if err != nil {
		return Bytes32{}, fmt.Errorf("failed to decode hex: %w", err)
	}

	var res Bytes32
	copy(res[:], b)
	return res, nil
}

func (b Bytes32) String() string {
	return "0x" + hex.EncodeToString(b[:])
}

func (b Bytes32) IsEmpty() bool {
	return b == Bytes32{}
}

func (b Bytes32) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.String())), nil
}

func (b *Bytes32) UnmarshalJSON(data []byte) error {
	v := string(data)
	if len(v) < 4 {
		return fmt.Errorf("invalid Bytes32: %s", v)
	}
	v = v[1 : len(v)-1] // trim quotes

	if !strings.HasPrefix(v, "0x") {
		return fmt.Errorf("bytes must start with '0x' prefix: %s", v)
	}
	v = v[2:] // trim 0x prefix

	bCp, err := hex.DecodeString(v)
	if err != nil {
		return err
	}

	copy(b[:], bCp)
	return nil
}

type BigInt struct {
	*big.Int
}

func NewBigInt(i *big.Int) BigInt {
	return BigInt{Int: i}
}

func NewBigIntFromInt64(i int64) BigInt {
	return BigInt{Int: big.NewInt(i)}
}

func (b BigInt) Bytes() []byte {
	if b.Int == nil {
		return []byte{}
	}
	return b.Int.Bytes()
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	if b.Int == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, b.String())), nil
}

func (b *BigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}

	if len(p) < 2 {
		return fmt.Errorf("invalid BigInt: %s", p)
	}
	p = p[1 : len(p)-1]

	z := big.NewInt(0)
	_, ok := z.SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}

func (b BigInt) IsEmpty() bool {
	return b.Int == nil
}

func (b BigInt) IsPositive() bool {
	return b.Int != nil && b.Int.Cmp(big.NewInt(0)) > 0
}
