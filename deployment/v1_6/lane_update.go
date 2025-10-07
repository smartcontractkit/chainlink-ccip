package v1_6

import "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

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
	Source       ChainDefinition
	Dest         ChainDefinition
	IsDisabled   bool
	TestRouter   bool
	ExtraConfigs ExtraConfigs
	MCMS         *mcms.Input
}
