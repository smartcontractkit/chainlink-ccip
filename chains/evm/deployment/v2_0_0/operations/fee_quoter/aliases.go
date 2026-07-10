package fee_quoter

import (
	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
)

// Backward-compatible aliases for downstream consumers (e.g. chainlink deployment)
// that imported the pre-ops-v2 fee_quoter package surface.

var FeeQuoterABI = gobindings.FeeQuoterABI
var FeeQuoterBin = gobindings.FeeQuoterBin

type (
	AuthorizedCallerArgs                    = gobindings.AuthorizedCallersAuthorizedCallerArgs
	DestChainConfig                         = gobindings.FeeQuoterDestChainConfig
	DestChainConfigArgs                     = gobindings.FeeQuoterDestChainConfigArgs
	GasPriceUpdate                          = gobindings.InternalGasPriceUpdate
	PriceUpdates                            = gobindings.InternalPriceUpdates
	StaticConfig                            = gobindings.FeeQuoterStaticConfig
	TimestampedPackedUint224                = gobindings.InternalTimestampedPackedUint224
	TokenPriceUpdate                        = gobindings.InternalTokenPriceUpdate
	TokenTransferFeeConfig                  = gobindings.FeeQuoterTokenTransferFeeConfig
	TokenTransferFeeConfigArgs              = gobindings.FeeQuoterTokenTransferFeeConfigArgs
	TokenTransferFeeConfigRemoveArgs        = gobindings.FeeQuoterTokenTransferFeeConfigRemoveArgs
	TokenTransferFeeConfigSingleTokenArgs   = gobindings.FeeQuoterTokenTransferFeeConfigSingleTokenArgs
)
