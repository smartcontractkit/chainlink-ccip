package chainconfig

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
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
	// Note that this is only applicable for some chains, such as L2's.
	DAGasPriceDeviationPPB cciptypes.BigInt `json:"daGasPriceDeviationPPB"`

	// TODO: do we want something like finality tag enabled / finality depth here?
	// FinalityDepth is the number of blocks on top of a block that need to be included
	// before a block is considered finalized.
	// If set to 0, finality tags will be used.
	FinalityDepth uint64 `json:"finalityDepth"`
}
