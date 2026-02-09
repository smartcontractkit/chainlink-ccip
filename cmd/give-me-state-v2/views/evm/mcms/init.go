package mcms

import (
	"give-me-state-v2/views"
)

func init() {
	// Register MCMS views
	// These contracts are ManyChainMultiSig instances with different roles
	// They all share the same interface: owner() and getConfig()

	// Register with version 1.0.0 (from mainnet_address_refs.json)
	views.Register("evm", "BypasserManyChainMultiSig", "1.0.0", ViewMCMS)
	views.Register("evm", "CancellerManyChainMultiSig", "1.0.0", ViewMCMS)
	views.Register("evm", "ProposerManyChainMultiSig", "1.0.0", ViewMCMS)
	views.Register("evm", "ManyChainMultiSig", "1.0.0", ViewMCMS)

	// Also register without version as fallback
	views.Register("evm", "BypasserManyChainMultiSig", "", ViewMCMS)
	views.Register("evm", "CancellerManyChainMultiSig", "", ViewMCMS)
	views.Register("evm", "ProposerManyChainMultiSig", "", ViewMCMS)
	views.Register("evm", "ManyChainMultiSig", "", ViewMCMS)

	// Register RBACTimelock view
	views.Register("evm", "RBACTimelock", "1.0.0", ViewRBACTimelock)
	views.Register("evm", "RBACTimelock", "", ViewRBACTimelock)

	// Register CallProxy view (minimal - no public read functions)
	views.Register("evm", "CallProxy", "1.0.0", ViewCallProxy)
	views.Register("evm", "CallProxy", "", ViewCallProxy)
}
