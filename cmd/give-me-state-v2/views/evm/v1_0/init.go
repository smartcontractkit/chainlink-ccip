package v1_0

import (
	"give-me-state-v2/views"
)

func init() {
	// Register v1.0 views
	// Note: ARMProxy is the same contract as RMNProxy
	views.Register("evm", "ARMProxy", "1.0.0", ViewRMNProxy)

	// LinkToken (ERC20 with minters/burners)
	views.Register("evm", "LinkToken", "1.0.0", ViewLinkToken)

	// StaticLinkToken (simpler wrapped/static LINK token)
	views.Register("evm", "StaticLinkToken", "1.0.0", ViewStaticLinkToken)

	// CapabilitiesRegistry v1.0
	views.Register("evm", "CapabilitiesRegistry", "1.0.0", ViewCapabilitiesRegistry)
}
