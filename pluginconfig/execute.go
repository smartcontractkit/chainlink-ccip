package pluginconfig

import (
	"time"

	"github.com/smartcontractkit/libocr/commontypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type ExecutePluginConfig struct {
	// DestChain is the ccip destination chain configured for the execute DON.
	DestChain cciptypes.ChainSelector `json:"destChain"`

	// ObserverInfo is a map of oracle IDs to ObserverInfo.
	ObserverInfo map[commontypes.OracleID]ObserverInfo `json:"observerInfo"`

	// MessageVisibilityInterval is the time interval for which the messages are visible by the plugin.
	MessageVisibilityInterval time.Duration `json:"messageVisibilityInterval"`
}

type ObserverInfo struct {
	// Writer indicates that the node can contribute by sending reports to the destination chain.
	// Being a Writer guarantees that the node can also read from the destination chain.
	Writer bool `json:"writer"`

	// Reads define the chains that the current node can read from.
	Reads []cciptypes.ChainSelector `json:"reads"`
}
