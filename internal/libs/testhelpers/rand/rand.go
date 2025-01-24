package rand

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
	n, err := rand.Int(rand.Reader, big.NewInt(256))
	if err != nil {
		panic(err)
	}
	return uint32(n.Int64())
}

func RandomUint64() uint64 {
	n, err := rand.Int(rand.Reader, new(big.Int).SetUint64(^uint64(0)))
	if err != nil {
		panic(err)
	}
	return n.Uint64()
}

func RandomInt64() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(1e18))
	if err != nil {
		panic(err)
	}
	return n.Int64()
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
