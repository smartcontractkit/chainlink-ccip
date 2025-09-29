package v1_6

import (
	changeset_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

type ConnectChainsConfig struct {
	Lanes []LaneConfig
	MCMS  *changeset_utils.MCMSParams
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
	UpdateFeeQuoterDestsConfig map[uint64]FeeQuoterDestChainConfig
	UpdateOnRampDestsConfig    map[uint64]UpdateOnRampDestsInput
	UpdateOffRampSourcesConfig map[uint64]UpdateOffRampSourcesInput
	UpdateRouterRampsConfig    map[uint64]UpdateRouterDestInput
	ExtraConfigs               ExtraConfigs
	MCMS                       *changeset_utils.MCMSParams
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
