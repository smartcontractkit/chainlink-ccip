package chainconfig

import (
	"encoding/json"
	"errors"
	"math/big"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ChainConfig holds configuration that is stored in the onchain CCIPConfig.sol
// configuration contract, specifically the `bytes config` field of the ChainConfig
// solidity struct.
type ChainConfig struct {
	// GasPriceDeviationPPB is the maximum deviation in parts per billion that the
	// gas price being reported to this chain is allowed to deviate from the last
	// written gas price before we write a new gas price.
	//
	// Note that this applies to ALL source chains that can send messages to this chain.
	// For example, if the ChainConfig for Ethereum has a GasPriceDeviationPPB of 1e9 (100%),
	// then the gas price for ALL source chains that send messages to Ethereum
	// will be allowed to deviate by up to 100% from the last written gas price on-chain
	// before we write a new gas price.
	GasPriceDeviationPPB cciptypes.BigInt `json:"gasPriceDeviationPPB"`

	// DAGasPriceDeviationPPB is the maximum deviation in parts per billion that the
	// data-availability gas prices being reported to this chain are allowed to deviate
	// from the last written data-availability gas price on-chain before we write a new
	// data-availability gas price.
	//
	// Note that this applies to ALL source chains that can send messages to this chain.
	// For example, if the ChainConfig for Ethereum has a DAGasPriceDeviationPPB of 5e8 (50%),
	// then the data-availability gas price for ALL source chains that send messages to Ethereum
	// will be allowed to deviate by up to 50% from the last written data-availability gas price on-chain
	// before we write a new data-availability gas price.
	//
	// This is only applicable for some chains, such as L2's.
	DAGasPriceDeviationPPB cciptypes.BigInt `json:"daGasPriceDeviationPPB"`

	// OptimisticConfirmations is the number of confirmations of a chain event before
	// it is considered optimistically confirmed (i.e not necessarily finalized).
	// Deprecated: this is no longer used, setting it will have no effect.
	OptimisticConfirmations uint32 `json:"optimisticConfirmations"`

	// ChainFeeDeviationDisabled is a flag to disable deviation-based reporting of gas prices
	// on this chain. If true, we will only report prices based on the heartbeat (configured
	// in the commit plugin offchain config).
	ChainFeeDeviationDisabled bool `json:"chainFeeDeviationDisabled"`
}

func (cc ChainConfig) Validate() error {
	if cc.GasPriceDeviationPPB.Int == nil {
		return errors.New("GasPriceDeviationPPB not set")
	}

	if cc.GasPriceDeviationPPB.Int.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("GasPriceDeviationPPB not set or negative")
	}

	if cc.DAGasPriceDeviationPPB.Int == nil {
		return errors.New("DAGasPriceDeviationPPB not set")
	}

	if cc.DAGasPriceDeviationPPB.Int.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("DAGasPriceDeviationPPB not set or negative")
	}

	// No validation for OptimisticConfirmations as it is deprecated
	// and no longer used.

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
