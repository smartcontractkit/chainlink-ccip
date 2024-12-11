package ccipocr3

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ExecutePluginReport struct {
	ChainReports []ExecutePluginReportSingleChain `json:"chainReports"`
}

type ExecutePluginReportSingleChain struct {
	SourceChainSelector ChainSelector `json:"sourceChainSelector"`
	Messages            []Message     `json:"messages"`
	OffchainTokenData   [][][]byte    `json:"offchainTokenData"`
	Proofs              []Bytes32     `json:"proofs"`
	ProofFlagBits       BigInt        `json:"proofFlagBits"`
}

type CommitReportWithMessages struct {
	MerkleRootChain
	MessagesToExecute []Message
}

// ExecuteReportInfo contains metadata needed by transmitter and contract
// writer.
type ExecuteReportInfo struct {
	// TODO: add real info, this is a placeholder.
	Roots            []Bytes32
	CommitReportData []MerkleRootChain
}

// Encode v1 execute report info. Very basic versioning in the first byte to
// allow for future encoding optimizations.
func (eri ExecuteReportInfo) Encode() ([]byte, error) {
	data, err := json.Marshal(eri)
	data = append([]byte{1}, data...)
	return data, err
}

// DecodeExecuteReportInfo is a version aware decode function for the execute
// report info bytes.
func DecodeExecuteReportInfo(data []byte) (ExecuteReportInfo, error) {
	switch data[0] {
	case 1:
		var result ExecuteReportInfo
		dec := json.NewDecoder(bytes.NewReader(data[1:]))
		dec.DisallowUnknownFields()
		err := dec.Decode(&result)
		return result, err
	default:
		return ExecuteReportInfo{}, fmt.Errorf("unknown execute report info version (%d)", data[0])
	}
}
