package chainconfig

import (
	"encoding/json"
	"errors"
	"math/big"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ChainConfig holds configuration that is stored in the onchain CCIPConfig.sol
// configuration contract, specifically the `bytes config` field of the ChainConfig
// solidity struct.
type ChainConfig struct {
	// GasPriceDeviationPPB is the maximum deviation in parts per billion that the
	// gas price of this chain is allowed to deviate from the last written gas price
	// on-chain before we write a new gas price.
	GasPriceDeviationPPB cciptypes.BigInt `json:"gasPriceDeviationPPB"`

	// DAGasPriceDeviationPPB is the maximum deviation in parts per billion that the
	// data-availability gas price of this chain is allowed to deviate from the last
	// written data-availability gas price on-chain before we write a new data-availability
	// gas price.
	// This is only applicable for some chains, such as L2's.
	DAGasPriceDeviationPPB cciptypes.BigInt `json:"daGasPriceDeviationPPB"`

	// OptimisticConfirmations is the number of confirmations of a chain event before
	// it is considered optimistically confirmed (i.e not necessarily finalized).
	OptimisticConfirmations uint32 `json:"optimisticConfirmations"`

	// ChainFeeDeviationDisabled is a flag to disable deviation-based reporting. If true, we will only report
	// prices based on the heartbeat.
	ChainFeeDeviationDisabled bool `json:"chainFeeDeviationDisabled"`
}

func (cc ChainConfig) Validate() error {
	if cc.GasPriceDeviationPPB.Int.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("GasPriceDeviationPPB not set or negative")
	}

	// No validation for DAGasPriceDeviationPPB as it is optional
	// and only applicable to L2's.

	if cc.OptimisticConfirmations == 0 {
		return errors.New("OptimisticConfirmations not set")
	}

	return nil
}

// EncodeChainConfig encodes a ChainConfig into bytes using JSON.
func EncodeChainConfig(cc ChainConfig) ([]byte, error) {
	return json.Marshal(cc)
}

// DecodeChainConfig JSON decodes a ChainConfig from bytes.
func DecodeChainConfig(encodedChainConfig []byte) (ChainConfig, error) {
	var cc ChainConfig
	if err := json.Unmarshal(encodedChainConfig, &cc); err != nil {
		return cc, err
	}
	return cc, nil
}
