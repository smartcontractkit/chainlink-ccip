package lanes

import (
	"encoding/binary"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type ChainDefinition struct {
	// Selector is the chain selector of this chain.
	// This is provided by the user
	Selector uint64
	// GasPrice defines the USD price (18 decimals) per unit gas for this chain as a destination.
	// This is provided by the user
	GasPrice *big.Int
	// TokenPrices define the USD price (18 decimals) per 1e18 of the smallest token denomination for various tokens on this chain.
	// This is provided by the user
	TokenPrices map[string]*big.Int
	// FeeQuoterDestChainConfig is the configuration to be applied on source chain when this chain is a destination.
	// This is provided by the user
	FeeQuoterDestChainConfig FeeQuoterDestChainConfig
	// RMNVerificationEnabled is true if we want the RMN to bless messages FROM this chain.
	// This is provided by the user
	RMNVerificationEnabled bool
	// AllowListEnabled is true if we want an allowlist to dictate who can send messages TO this chain.
	// This is provided by the user
	AllowListEnabled bool
	// AllowList is the list of addresses that are allowed to send messages TO this chain.
	// This is provided by the user
	AllowList []string
	// OnRamp is the address of the OnRamp contract on this chain.
	// This is populated programmatically
	OnRamp []byte
	// OffRamp is the address of the OffRamp contract on this chain.
	// This is populated programmatically
	OffRamp []byte
	// Router is the address of the Router contract on this chain.
	// This is populated programmatically
	Router []byte
	// FeeQuoter is the address of the FeeQuoter contract on this chain.
	// This is populated programmatically
	FeeQuoter []byte
}

type FeeQuoterDestChainConfig struct {
	IsEnabled                         bool
	MaxNumberOfTokensPerMsg           uint16
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DestGasOverhead                   uint32
	DestGasPerPayloadByteBase         uint32
	DestGasPerPayloadByteHigh         uint32
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
	NetworkFeeUSDCents                uint16
}

type ConnectChainsConfig struct {
	Lanes []LaneConfig
	MCMS  mcms.Input
}
type LaneConfig struct {
	ChainA       ChainDefinition
	ChainB       ChainDefinition
	Version      *semver.Version
	IsDisabled   bool
	TestRouter   bool
	ExtraConfigs ExtraConfigs
}

type ExtraConfigs struct {
	OnRampVersion []byte
}

type UpdateLanesInput struct {
	Source       *ChainDefinition
	Dest         *ChainDefinition
	IsDisabled   bool
	TestRouter   bool
	ExtraConfigs ExtraConfigs
}

func DefaultFeeQuoterDestChainConfig(configEnabled bool, selector uint64) FeeQuoterDestChainConfig {
	chainHex := utils.GetSelectorHex(selector)
	params := FeeQuoterDestChainConfig{
		IsEnabled:                         configEnabled,
		MaxNumberOfTokensPerMsg:           10,
		MaxDataBytes:                      30_000,
		MaxPerMsgGasLimit:                 3_000_000,
		DestGasOverhead:                   300_000,
		DefaultTokenFeeUSDCents:           25,
		DestGasPerPayloadByteBase:         16,
		DestGasPerPayloadByteHigh:         40,
		DestGasPerPayloadByteThreshold:    3000,
		DestDataAvailabilityOverheadGas:   100,
		DestGasPerDataAvailabilityByte:    16,
		DestDataAvailabilityMultiplierBps: 1,
		DefaultTokenDestGasOverhead:       90_000,
		DefaultTxGasLimit:                 200_000,
		GasMultiplierWeiPerEth:            11e17,
		NetworkFeeUSDCents:                10,
		ChainFamilySelector:               binary.BigEndian.Uint32(chainHex[:]),
	}
	family, _ := chain_selectors.GetSelectorFamily(selector)
	switch family {
	case chain_selectors.FamilyTon:
		params.MaxPerMsgGasLimit = 4_200_000_000 // 4_200_000_000 nano TON = 4.2 TON
	}
	return params
}

func DefaultGasPrice(selector uint64) *big.Int {
	family, _ := chain_selectors.GetSelectorFamily(selector)
	switch family {
	case chain_selectors.FamilyTon:
		return big.NewInt(2.12e9) // 1 TON ~2.13 USD -> 1 nanoTON = 2.13e−9 USD -> 1 nanoTON expressed in 1e18 (1 USD) = 2.13e9
	}
	// Gas price in USD (18 decimals) per unit of gas
	// 2e12 = $0.000002 per gas unit
	// With ~500,000 gas, this results in ~$1 USD fee per message
	return big.NewInt(2e12)
}
