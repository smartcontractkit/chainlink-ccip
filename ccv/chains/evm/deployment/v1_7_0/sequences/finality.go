package sequences

// WaitForFinalityConfig is the finality configuration that instructs a pool or verifier to wait for
// full (hard) finality before processing a message. Corresponds to WAIT_FOR_FINALITY_FLAG in Solidity.
var WaitForFinalityConfig = [4]byte{}

// BlockDepthFinalityConfig encodes a block depth as a bytes4 finality configuration.
// The encoding mirrors FinalityCodec._encodeBlockDepth in Solidity: the depth is stored as a
// big-endian uint32 value in the lower 16 bits, leaving the upper 16 bits (flags) as zero.
func BlockDepthFinalityConfig(blockDepth uint16) [4]byte {
	return [4]byte{0, 0, byte(blockDepth >> 8), byte(blockDepth)}
}
