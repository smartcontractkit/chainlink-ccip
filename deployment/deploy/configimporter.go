package deploy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
)

type ImportConfigInput struct {
	PerChainInputs []ImportConfigPerChainInput
	AdapterVersion *semver.Version
}

type ImportConfigPerChainInput struct {
	ChainSelector uint64
	RemoteChains  []uint64
	Tokens        []common.Address
}
