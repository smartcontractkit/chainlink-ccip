package v1_6

import "math/big"

type ChainDefinition struct {
	// ConnectionConfig holds configuration for connection.
	ConnectionConfig
	// Selector is the chain selector of this chain.
	Selector uint64
	// GasPrice defines the USD price (18 decimals) per unit gas for this chain as a destination.
	GasPrice *big.Int
	// TokenPrices define the USD price (18 decimals) per 1e18 of the smallest token denomination for various tokens on this chain.
	TokenPrices map[string]*big.Int
	// FeeQuoterDestChainConfig is the configuration to be applied on source chain when this chain is a destination.
	FeeQuoterDestChainConfig FeeQuoterDestChainConfig
}

type ConnectionConfig struct {
	// RMNVerificationDisabled is true if we do not want the RMN to bless messages FROM this chain.
	RMNVerificationDisabled bool
	// AllowListEnabled is true if we want an allowlist to dictate who can send messages TO this chain.
	AllowListEnabled bool
}

type FeeQuoterDestChainConfig struct {
	IsEnabled                         bool
	MaxNumberOfTokensPerMsg           uint16
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DestGasOverhead                   uint32
	DestGasPerPayloadByteBase         uint8
	DestGasPerPayloadByteHigh         uint8
	DestGasPerPayloadByteThreshold    uint16
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	ChainFamilySelector               uint32
	EnforceOutOfOrder                 bool
	DefaultTokenFeeUSDCents           uint16
	DefaultTokenDestGasOverhead       uint32
	DefaultTxGasLimit                 uint32
	GasMultiplierWeiPerEth            uint64
	GasPriceStalenessThreshold        uint32
	NetworkFeeUSDCents                uint32
}

type FeeQuoterPriceUpdatePerSource struct {
	TokenPrices map[string]*big.Int // token address -> price
	GasPrices   map[uint64]*big.Int // dest chain -> gas price
}
