package verifier_tags

// TODO potentially update schema with 2-byte prefix + 2 byte version.
var (
	cctpVerifierV2      = [4]byte{0x91, 0xb3, 0x33, 0x8e} // bytes4(keccak256("CCTPVerifier 2.1.0"))
	committeeVerifierV2 = [4]byte{0xe9, 0xa0, 0x5a, 0x20} // bytes4(keccak256("CommitteeVerifier 2.0.0"))
	lombardVerifierV2   = [4]byte{0x5b, 0x92, 0x53, 0xce} // bytes4(keccak256("LombardVerifier 2.1.0"))
)

func CCTPVerifierV2() [4]byte {
	return cctpVerifierV2
}

func CommitteeVerifierV2() [4]byte {
	return committeeVerifierV2
}

func LombardVerifierV2() [4]byte {
	return lombardVerifierV2
}
