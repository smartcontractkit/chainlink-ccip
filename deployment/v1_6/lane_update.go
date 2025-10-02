package v1_6

import (
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type ConnectChainsConfig struct {
	Lanes []LaneConfig
	MCMS  *mcms.Input
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
	UpdateFeeQuoterDestsConfig FeeQuoterDestChainConfig
	UpdateFeeQuoterPrices      FeeQuoterPriceUpdatePerSource
	UpdateOnRampDestsConfig    UpdateOnRampDestsInput
	UpdateOffRampSourcesConfig UpdateOffRampSourcesInput
	ExtraConfigs               ExtraConfigs
	MCMS                       *mcms.Input
}

type UpdateOnRampDestsInput struct {
	IsEnabled        bool
	TestRouter       bool
	AllowListEnabled bool
}

type UpdateOffRampSourcesInput struct {
	IsEnabled                 bool
	TestRouter                bool
	IsRMNVerificationDisabled bool
	OnRamp                    []byte
}
