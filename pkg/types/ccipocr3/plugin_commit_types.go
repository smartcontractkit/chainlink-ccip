package ccipocr3

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
	PriceUpdates         PriceUpdates      `json:"priceUpdates"`
	BlessedMerkleRoots   []MerkleRootChain `json:"blessedMerkleRoots"`
	UnblessedMerkleRoots []MerkleRootChain `json:"unblessedMerkleRoots"`

	// RMNSignatures are the ECDSA signatures from the RMN signing nodes on the RMNReport structure.
	// For more details see the contract here: https://github.com/smartcontractkit/chainlink/blob/7ba0f37134a618375542079ff1805fe2224d7916/contracts/src/v0.8/ccip/interfaces/IRMNV2.sol#L8-L12
	//nolint:lll // it's a url
	RMNSignatures []RMNECDSASignature `json:"rmnSignatures"`
}

// IsEmpty returns true if the CommitPluginReport is empty.
// NOTE: A report is considered empty when core fields are missing (MerkleRoots, TokenPrices, GasPriceUpdates).
func (r CommitPluginReport) IsEmpty() bool {
	return len(r.BlessedMerkleRoots) == 0 &&
		len(r.UnblessedMerkleRoots) == 0 &&
		len(r.PriceUpdates.TokenPriceUpdates) == 0 &&
		len(r.PriceUpdates.GasPriceUpdates) == 0
}

func (r CommitPluginReport) HasNoRoots() bool {
	return len(r.BlessedMerkleRoots) == 0 && len(r.UnblessedMerkleRoots) == 0
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

// String returns a string representation of the MerkleRootChain
func (m MerkleRootChain) String() string {
	return fmt.Sprintf("MerkleRoot(chain:%d, seqNumsRange:%s, merkleRoot:%s, onRamp:%s)",
		m.ChainSel, m.SeqNumsRange, m.MerkleRoot, m.OnRampAddress)
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

// CommitReportInfo is the info data that will be sent with the along with the report
// It will be used to determine if the report should be accepted or not
type CommitReportInfo struct {
	// RemoteF Max number of faulty RMN nodes; f+1 signers are required to verify a report.
	RemoteF     uint64            `json:"remoteF"`
	MerkleRoots []MerkleRootChain `json:"merkleRoots"`
	PriceUpdates
}

func (cri CommitReportInfo) Encode() ([]byte, error) {
	data, err := json.Marshal(cri)
	data = append([]byte{01}, data...)
	return data, err
}

// DecodeCommitReportInfo is a version aware decode function for the commit
// report info bytes.
func DecodeCommitReportInfo(data []byte) (CommitReportInfo, error) {
	if len(data) == 0 {
		return CommitReportInfo{}, nil
	}

	switch data[0] {
	case 1:
		var result CommitReportInfo
		dec := json.NewDecoder(bytes.NewReader(data[1:]))
		dec.DisallowUnknownFields()
		err := dec.Decode(&result)
		return result, err
	default:
		return CommitReportInfo{}, fmt.Errorf("unknown execute report info version (%d)", data[0])
	}
}
