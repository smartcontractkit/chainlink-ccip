package deploy

import (
	"github.com/Masterminds/semver/v3"
	ccipocr3 "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// OCR3 is specifically a 1.6.0 feature
var OCR3Version = *semver.MustParse("1.6.0")

// CapabilitiesRegistryVersion is specifically a 1.0.0 feature
var CapabilitiesRegistryVersion = *semver.MustParse("1.0.0")

type SetOCR3ConfigInput struct {
	ChainSelector uint64
	Datastore     datastore.DataStore
	Configs       map[ccipocr3.PluginType]OCR3ConfigArgs
}

type OCR3ConfigArgs struct {
	ConfigDigest                   [32]byte
	PluginType                     ccipocr3.PluginType
	F                              uint8
	IsSignatureVerificationEnabled bool
	Signers                        [][]byte
	Transmitters                   [][]byte
}

type SetOCR3ConfigArgs struct {
	HomeChainSel    uint64
	RemoteChainSels []uint64
	ConfigType      utils.ConfigType
	MCMS            mcms.Input
}
