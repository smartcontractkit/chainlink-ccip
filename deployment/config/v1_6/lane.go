package v1_6

import "github.com/smartcontractkit/chainlink/deployment/common/proposalutils"

type AddLanesConfig struct {
	Lanes []LaneConfig
	MCMS  *proposalutils.TimelockConfig
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
