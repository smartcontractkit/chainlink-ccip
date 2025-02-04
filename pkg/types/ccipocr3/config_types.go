package ccipocr3

import (
	"fmt"
)

type OCRConfigResponse struct {
	OCRConfig OCRConfig
}

type OCRConfig struct {
	ConfigInfo   ConfigInfo
	Signers      [][]byte
	Transmitters [][]byte
}

type ConfigInfo struct {
	ConfigDigest                   [32]byte
	F                              uint8
	N                              uint8
	IsSignatureVerificationEnabled bool
}

type RMNDigestHeader struct {
	DigestHeader Bytes32
}

// FeeQuoterStaticConfig is used to parse the response from the feeQuoter contract's getStaticConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/FeeQuoter.sol#L946
//
//nolint:lll // It's a URL.
type FeeQuoterStaticConfig struct {
	MaxFeeJuelsPerMsg  BigInt `json:"maxFeeJuelsPerMsg"`
	LinkToken          []byte `json:"linkToken"`
	StalenessThreshold uint32 `json:"stalenessThreshold"`
}

// selectorsAndConfigs wraps the return values from getAllsourceChainConfigs.
type SelectorsAndConfigs struct {
	Selectors          []uint64            `mapstructure:"F0"`
	SourceChainConfigs []SourceChainConfig `mapstructure:"F1"`
}

// sourceChainConfig is used to parse the response from the offRamp contract's getSourceChainConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L94
//
//nolint:lll // It's a URL.
type SourceChainConfig struct {
	Router    []byte
	IsEnabled bool
	MinSeqNr  uint64
	OnRamp    UnknownAddress
}

func (scc SourceChainConfig) Check() (bool /* enabled */, error) {
	// The chain may be set in CCIPHome's ChainConfig map but not hooked up yet in the offramp.
	if !scc.IsEnabled {
		return false, nil
	}
	// This may happen due to some sort of regression in the codec that unmarshals
	// chain data -> go struct.
	if len(scc.OnRamp) == 0 {
		return false, fmt.Errorf(
			"onRamp misconfigured/didn't unmarshal: %x",
			scc.OnRamp,
		)
	}
	return scc.IsEnabled, nil
}

// OffRampStaticChainConfig is used to parse the response from the offRamp contract's getStaticConfig method.
// See: <chainlink repo>/contracts/src/v0.8/ccip/offRamp/OffRamp.sol:StaticConfig
type OffRampStaticChainConfig struct {
	ChainSelector        ChainSelector `json:"chainSelector"`
	GasForCallExactCheck uint16        `json:"gasForCallExactCheck"`
	RmnRemote            []byte        `json:"rmnRemote"`
	TokenAdminRegistry   []byte        `json:"tokenAdminRegistry"`
	NonceManager         []byte        `json:"nonceManager"`
}

// OffRampDynamicChainConfig maps to DynamicConfig in OffRamp.sol
type OffRampDynamicChainConfig struct {
	FeeQuoter                               []byte `json:"feeQuoter"`
	PermissionLessExecutionThresholdSeconds uint32 `json:"permissionLessExecutionThresholdSeconds"`
	IsRMNVerificationDisabled               bool   `json:"isRMNVerificationDisabled"`
	MessageInterceptor                      []byte `json:"messageInterceptor"`
}

// We're wrapping the onRampDynamicConfig this way to map to on-chain return type which is a named struct
// https://github.com/smartcontractkit/chainlink/blob/12af1de88238e0e918177d6b5622070417f48adf/contracts/src/v0.8/ccip/onRamp/OnRamp.sol#L328
//
//nolint:lll
type GetOnRampDynamicConfigResponse struct {
	DynamicConfig OnRampDynamicConfig `json:"dynamicConfig"`
}

// See DynamicChainConfig in OnRamp.sol
type OnRampDynamicConfig struct {
	FeeQuoter              []byte `json:"feeQuoter"`
	ReentrancyGuardEntered bool   `json:"reentrancyGuardEntered"`
	MessageInterceptor     []byte `json:"messageInterceptor"`
	FeeAggregator          []byte `json:"feeAggregator"`
	AllowListAdmin         []byte `json:"allowListAdmin"`
}

// VersionedConfigRemote is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L167-L169
type VersionedConfigRemote struct {
	Version uint32 `json:"version"`
	Config  Config `json:"config"`
}

// Config is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L49-L53
type Config struct {
	RMNHomeContractConfigDigest Bytes32  `json:"rmnHomeContractConfigDigest"`
	Signers                     []Signer `json:"signers"`
	F                           uint64   `json:"f"` // previously: MinSigners
}

// signer is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L42-L45
type Signer struct {
	OnchainPublicKey []byte `json:"onchainPublicKey"`
	NodeIndex        uint64 `json:"nodeIndex"`
}
