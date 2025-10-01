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
	Source                     ChainDefinition
	Dest                       ChainDefinition
	IsDisabled                 bool
	ExtraConfigs               ExtraConfigs
	MCMS                       *utils.MCMSInput
}
