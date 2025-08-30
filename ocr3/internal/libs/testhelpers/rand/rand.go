package rand

import (
	"crypto/rand"
	"encoding/hex"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func RandomBytes32() cciptypes.Bytes32 {
	var b cciptypes.Bytes32
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	return b
}

func RandomUint32() uint32 {
	// Generate random bytes directly to avoid big.Int truncation
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// Convert bytes to uint32, then mod by 256 to match original behavior
	var result uint32
	for i, byt := range b {
		result |= uint32(byt) << (8 * i)
	}
	return result % 256
}

func RandomUint64() uint64 {
	// Generate random bytes directly to avoid big.Int truncation
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// Convert bytes to uint64
	var result uint64
	for i, byt := range b {
		result |= uint64(byt) << (8 * i)
	}
	return result
}

func RandomInt64() int64 {
	// Generate random bytes and convert to int64, constrained to 1e18
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// Convert bytes to uint64 first, then to int64 with constraint
	var result uint64
	for i, byt := range b {
		result |= uint64(byt) << (8 * i)
	}
	// Ensure result is positive and within 1e18 range
	return int64(result % uint64(1e18))
}

func RandomPrefix() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func RandomReportVersion() cciptypes.Bytes32 {
	versions := []cciptypes.Bytes32{
		stringToBytes32("RMN_V1_6_ANY2EVM_REPORT"),
		stringToBytes32("RMN_V2_0_ANY2EVM_REPORT"),
		stringToBytes32("RMN_V1_6_EVM2EVM_REPORT"),
	}
	return versions[RandomUint32()%uint32(len(versions))]
}

func stringToBytes32(s string) cciptypes.Bytes32 {
	var result cciptypes.Bytes32
	copy(result[:], s)
	return result
}

func RandomAddress() cciptypes.UnknownEncodedAddress {
	b := RandomAddressBytes()
	return cciptypes.UnknownEncodedAddress(b.String())
}

func RandomAddressBytes() cciptypes.Bytes {
	return cciptypes.Bytes(RandomBytes(20))
}
