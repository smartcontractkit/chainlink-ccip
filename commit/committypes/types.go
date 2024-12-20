package committypes

import (
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Query struct {
	MerkleRootQuery merkleroot.Query `json:"merkleRootQuery"`
	TokenPriceQuery tokenprice.Query `json:"tokenPriceQuery"`
	ChainFeeQuery   chainfee.Query   `json:"chainFeeQuery"`
}

func (q Query) Encode() ([]byte, error) {
	return json.Marshal(q)
}

func DecodeCommitPluginQuery(encodedQuery []byte) (Query, error) {
	q := Query{}
	err := json.Unmarshal(encodedQuery, &q)
	return q, err
}

type Observation struct {
	MerkleRootObs merkleroot.Observation          `json:"merkleObs"`
	TokenPriceObs tokenprice.Observation          `json:"tokenObs"`
	ChainFeeObs   chainfee.Observation            `json:"chainFeeObs"`
	DiscoveryObs  dt.Observation                  `json:"discoveryObs"`
	FChain        map[cciptypes.ChainSelector]int `json:"fChain"`
}

func (obs Observation) Encode() ([]byte, error) {
	encodedObservation, err := json.Marshal(obs)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Observation: %w", err)
	}

	return encodedObservation, nil
}

func DecodeCommitPluginObservation(encodedObservation []byte) (Observation, error) {
	o := Observation{}
	err := json.Unmarshal(encodedObservation, &o)
	return o, err
}

type Outcome struct {
	MerkleRootOutcome merkleroot.Outcome `json:"merkleRootOutcome"`
	TokenPriceOutcome tokenprice.Outcome `json:"tokenPriceOutcome"`
	ChainFeeOutcome   chainfee.Outcome   `json:"chainFeeOutcome"`
}

// Encode encodes an Outcome deterministically
func (o Outcome) Encode() ([]byte, error) {
	// Sort all lists to ensure deterministic serialization
	o.MerkleRootOutcome.Sort()
	encodedOutcome, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Outcome: %w", err)
	}

	return encodedOutcome, nil
}

func DecodeOutcome(b []byte) (Outcome, error) {
	if len(b) == 0 {
		return Outcome{}, nil
	}

	o := Outcome{}
	if err := json.Unmarshal(b, &o); err != nil {
		return Outcome{}, fmt.Errorf("decode outcome: %w", err)
	}

	return o, nil
}
