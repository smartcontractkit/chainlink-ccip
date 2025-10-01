package v1_6

import (
	utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

type ConnectChainsConfig struct {
	Lanes []LaneConfig
	MCMS  *utils.MCMSInput
}
type LaneConfig struct {
	Source       ChainDefinition
	Dest         ChainDefinition
	IsDisabled   bool
	TestRouter   bool
	ExtraConfigs ExtraConfigs
}

type ExtraConfigs struct {
	OnRampVersion []byte
}

type UpdateLanesInput struct {
	Selector                   uint64
	RemoteSelector             uint64
	IsDisabled                 bool
	TestRouter                 bool
	IsRMNVerificationDisabled  bool
	AllowListEnabled           bool
	SrcOnRamp                  []byte
	DestOffRamp                []byte
	UpdateFeeQuoterDestsConfig FeeQuoterDestChainConfig
	UpdateFeeQuoterPrices      FeeQuoterPriceUpdatePerSource
	ExtraConfigs               ExtraConfigs
	MCMS                       *utils.MCMSInput
}

