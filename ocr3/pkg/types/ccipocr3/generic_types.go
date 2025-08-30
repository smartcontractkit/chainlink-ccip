package ccipocr3

import (
	"math/big"

	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Deprecated: Use ccipocr3common.TokenPrice instead.
type TokenPrice = ccipocr3common.TokenPrice

// Deprecated: Use ccipocr3common.TokenPriceMap instead.
type TokenPriceMap = ccipocr3common.TokenPriceMap

// Deprecated: Use ccipocr3common.NewTokenPrice instead.
func NewTokenPrice(tokenID UnknownEncodedAddress, price *big.Int) TokenPrice {
	return ccipocr3common.NewTokenPrice(tokenID, price)
}

// Deprecated: Use ccipocr3common.GasPriceChain instead.
type GasPriceChain = ccipocr3common.GasPriceChain

// Deprecated: Use ccipocr3common.NewGasPriceChain instead.
func NewGasPriceChain(gasPrice *big.Int, chainSel ChainSelector) GasPriceChain {
	return ccipocr3common.NewGasPriceChain(gasPrice, chainSel)
}

// Deprecated: Use ccipocr3common.SeqNum instead.
type SeqNum = ccipocr3common.SeqNum

// Deprecated: Use ccipocr3common.NewSeqNumRange instead.
func NewSeqNumRange(start, end SeqNum) SeqNumRange {
	return ccipocr3common.NewSeqNumRange(start, end)
}

// Deprecated: Use ccipocr3common.SeqNumRange instead.
type SeqNumRange = ccipocr3common.SeqNumRange

// Deprecated: Use ccipocr3common.ChainSelector instead.
type ChainSelector = ccipocr3common.ChainSelector

// Deprecated: Use ccipocr3common.Message instead.
type Message = ccipocr3common.Message

// Deprecated: Use ccipocr3common.RampMessageHeader instead.
type RampMessageHeader = ccipocr3common.RampMessageHeader

// Deprecated: Use ccipocr3common.RampTokenAmount instead.
type RampTokenAmount = ccipocr3common.RampTokenAmount
