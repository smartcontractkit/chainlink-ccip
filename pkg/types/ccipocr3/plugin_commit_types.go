package ccipocr3

import "bytes"

// CommitPluginReport contains the necessary information to commit CCIP
// messages from potentially many source chains, to a single destination chain.
//
// It must consist of either:
//
//  1. a non-empty MerkleRoots array, or
//  2. a non-empty PriceUpdates array
//
// If neither of the above is provided the report is considered empty and should
// not be transmitted on-chain.
//
// In the event the MerkleRoots array is non-empty, it may also contain
// RMNSignatures, if RMN is configured for some lanes involved in the commitment.
// A report with RMN signatures but without merkle roots is invalid.
type CommitPluginReport struct {
	MerkleRoots  []MerkleRootChain `json:"merkleRoots"`
	PriceUpdates PriceUpdates      `json:"priceUpdates"`

	// RMNSignatures are the ECDSA signatures from the RMN signing nodes on the RMNReport structure.
	// For more details see the contract here: https://github.com/smartcontractkit/chainlink/blob/7ba0f37134a618375542079ff1805fe2224d7916/contracts/src/v0.8/ccip/interfaces/IRMNV2.sol#L8-L12
	//nolint:lll // it's a url
	RMNSignatures []RMNECDSASignature `json:"rmnSignatures"`
}

// IsEmpty returns true if the CommitPluginReport is empty
func (r CommitPluginReport) IsEmpty() bool {
	return len(r.MerkleRoots) == 0 &&
		len(r.PriceUpdates.TokenPriceUpdates) == 0 &&
		len(r.PriceUpdates.GasPriceUpdates) == 0 &&
		len(r.RMNSignatures) == 0
}

// MerkleRootChain Mirroring https://github.com/smartcontractkit/chainlink/blob/cd5c78959575f593b27fd83d8766086d0c678487/contracts/src/v0.8/ccip/libraries/Internal.sol#L356-L362
//
//nolint:lll // it's a url
type MerkleRootChain struct {
	ChainSel      ChainSelector  `json:"chain"`
	OnRampAddress UnknownAddress `json:"onRampAddress"`
	SeqNumsRange  SeqNumRange    `json:"seqNumsRange"`
	MerkleRoot    Bytes32        `json:"merkleRoot"`
}

func (m MerkleRootChain) Equals(other MerkleRootChain) bool {
	return m.ChainSel == other.ChainSel &&
		bytes.Equal(m.OnRampAddress, other.OnRampAddress) &&
		m.SeqNumsRange == other.SeqNumsRange &&
		m.MerkleRoot == other.MerkleRoot
}

type PriceUpdates struct {
	TokenPriceUpdates []TokenPrice    `json:"tokenPriceUpdates"`
	GasPriceUpdates   []GasPriceChain `json:"gasPriceUpdates"`
}
