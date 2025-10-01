package v1_6

import (
	changeset_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

type ConnectChainsConfig struct {
	Lanes []LaneConfig
	MCMS  *changeset_utils.MCMSInput
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
	UpdateRouterRampsConfig    UpdateRouterDestInput
	ExtraConfigs               ExtraConfigs
	MCMS                       *changeset_utils.MCMSInput
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

type UpdateRouterDestInput struct {
	OffRampUpdates map[uint64]bool
	OnRampUpdates  map[uint64]bool
}
