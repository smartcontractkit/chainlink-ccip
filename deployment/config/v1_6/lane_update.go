package v1_6

import "github.com/smartcontractkit/chainlink/deployment/common/proposalutils"

type UpdateLanesInput struct {
	UpdateFeeQuoterDestsConfig map[uint64]FeeQuoterDestChainConfig
	UpdateOnRampDestsConfig    map[uint64]UpdateOnRampDestsInput
	UpdateOffRampSourcesConfig map[uint64]UpdateOffRampSourcesInput
	UpdateRouterRampsConfig    map[uint64]UpdateRouterDestInput
	ExtraConfigs               ExtraConfigs
	MCMS                       *proposalutils.TimelockConfig
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
