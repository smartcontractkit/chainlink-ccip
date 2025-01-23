package ccipocr3

// FeeQuoterDestChainConfig represents the configuration of a destination chain in the FeeQuoter contract
type FeeQuoterDestChainConfig struct {
	IsEnabled                         bool    // Whether this destination chain is enabled
	MaxNumberOfTokensPerMsg           uint16  // Maximum number of distinct ERC20 token transferred per message
	MaxDataBytes                      uint32  // Maximum payload data size in bytes
	MaxPerMsgGasLimit                 uint32  // Maximum gas limit for messages targeting EVMs
	DestGasOverhead                   uint32  // Gas charged on top of the gasLimit to cover destination chain costs
	DestGasPerPayloadByteBase         uint32  // Destination gas charged per byte of payload to receiver by default
	DestGasPerPayloadByteHigh         uint32  // Destination gas charged per byte of payload over the threshold
	DestGasPerPayloadByteThreshold    uint32  // Threshold of payload byte size over which the high rate applies
	DestDataAvailabilityOverheadGas   uint32  // Extra data availability gas charged, e.g., for OCR
	DestGasPerDataAvailabilityByte    uint16  // Gas charged per byte of message data needing availability
	DestDataAvailabilityMultiplierBps uint16  // Multiplier for data availability gas in bps
	DefaultTokenFeeUSDCents           uint16  // Default token fee charged per token transfer
	DefaultTokenDestGasOverhead       uint32  // Default gas charged to execute token transfer on destination
	DefaultTxGasLimit                 uint32  // Default gas limit for a transaction
	GasMultiplierWeiPerEth            uint64  // Multiplier for gas costs, 1e18 based (11e17 = 10% extra cost)
	NetworkFeeUSDCents                uint32  // Flat network fee for messages, in multiples of 0.01 USD
	GasPriceStalenessThreshold        uint32  // Maximum time for gas price to be valid (0 means disabled)
	EnforceOutOfOrder                 bool    // Enforce the allowOutOfOrderExecution extraArg to be true
	ChainFamilySelector               [4]byte // Selector identifying the destination chain's family
}

// HasNonEmptyDAGasParams returns true if the destination chain has non-empty data availability gas parameters
func (c FeeQuoterDestChainConfig) HasNonEmptyDAGasParams() bool {
	return c.DestDataAvailabilityOverheadGas != 0 && c.DestGasPerDataAvailabilityByte != 0 &&
		c.DestDataAvailabilityMultiplierBps != 0
}

type DataAvailabilityGasConfig struct {
	DestDataAvailabilityOverheadGas   uint32 // Extra data availability gas charged, e.g., for OCR
	DestGasPerDataAvailabilityByte    uint16 // Gas charged per byte of message data needing availability
	DestDataAvailabilityMultiplierBps uint16 // Multiplier for data availability gas in bps
}
