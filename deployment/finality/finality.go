package finality

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Bit layout mirrors FinalityCodec.sol:
//
//	bits 31..17  reserved flags (accepted on-chain but unused today)
//	bit  16      WAIT_FOR_SAFE_FLAG
//	bits 15..0   block depth  (0 = wait-for-finality when no flags set)
const (
	blockDepthBits = 16
	safeFlagBit    = 1 << blockDepthBits // 0x0001_0000
)

// RawWaitForFinality is the on-chain [4]byte encoding of WAIT_FOR_FINALITY_FLAG (0x00000000).
// Use this for comparisons against raw on-chain values instead of `[4]byte{}`.
var RawWaitForFinality = [4]byte{}

// Config is the human-readable, serializable representation of the on-chain
// bytes4 finality configuration stored in verifiers, token pools, and executors.
//
// Exactly one "mode" should be active for *requested* finality, but *allowed*
// finality (what this config represents) may combine modes so that the
// contract accepts multiple request types.
type Config struct {
	WaitForFinality bool   `json:"waitForFinality" yaml:"waitForFinality"`
	WaitForSafe     bool   `json:"waitForSafe"     yaml:"waitForSafe"`
	BlockDepth      uint16 `json:"blockDepth"      yaml:"blockDepth"`
}

// Raw encodes the config into the on-chain [4]byte representation
// matching FinalityCodec.sol.
//
// WaitForFinality is the zero-value sentinel (0x00000000) and is therefore
// implicit: it is "allowed" whenever the raw value is 0x00000000 OR when
// the caller explicitly sets the flag. Because 0x00000000 already means
// "wait for finality" on-chain, the flag is not encoded into the wire
// format -- it only exists for readability in the Go/YAML layer.
func (c Config) Raw() [4]byte {
	var v uint32
	if c.WaitForSafe {
		v |= safeFlagBit
	}
	v |= uint32(c.BlockDepth)
	var out [4]byte
	binary.BigEndian.PutUint32(out[:], v)
	return out
}

// FromRaw decodes an on-chain [4]byte value into a Config.
//
// WaitForFinality is set to true only when raw == 0x00000000.
func FromRaw(raw [4]byte) Config {
	v := binary.BigEndian.Uint32(raw[:])
	return Config{
		WaitForFinality: v == 0,
		WaitForSafe:     v&safeFlagBit != 0,
		BlockDepth:      uint16(v & 0xFFFF),
	}
}

// IsZero returns true when all fields are at their zero value (unconfigured).
func (c Config) IsZero() bool {
	return !c.WaitForFinality && !c.WaitForSafe && c.BlockDepth == 0
}

// Validate rejects configurations that are structurally invalid.
func (c Config) Validate() error {
	if c.IsZero() {
		return errors.New("finality config is empty; set at least one mode")
	}
	if c.WaitForSafe && c.BlockDepth > 0 {
		return nil // allowed finality can combine safe + block depth
	}
	if c.WaitForFinality && (c.WaitForSafe || c.BlockDepth > 0) {
		return fmt.Errorf(
			"WaitForFinality is the zero-value sentinel and cannot be combined with other modes "+
				"(WaitForSafe=%v, BlockDepth=%d)", c.WaitForSafe, c.BlockDepth)
	}
	return nil
}
