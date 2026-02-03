package v1_6

import (
	"call-orchestrator-demo/views"
)

func init() {
	// Register v1.6 views
	// Note: All views use manual decoding via common utilities
	views.Register("evm", "FeeQuoter", "1.6.0", ViewFeeQuoter)
	views.Register("evm", "FeeQuoter", "1.6.3", ViewFeeQuoter) // Same ABI as 1.6.0
	views.Register("evm", "OnRamp", "1.6.0", ViewOnRamp)
	views.Register("evm", "OffRamp", "1.6.0", ViewOffRamp)
	views.Register("evm", "RegistryModuleOwnerCustom", "1.6.0", ViewRegistryModuleOwnerCustom)
	views.Register("evm", "NonceManager", "1.6.0", ViewNonceManager)
	views.Register("evm", "RMNRemote", "1.6.0", ViewRMNRemote)
	views.Register("evm", "BurnMintTokenPool", "1.6.0", ViewBurnMintTokenPool)
	views.Register("evm", "LockReleaseTokenPool", "1.6.0", ViewLockReleaseTokenPool)

	// Home contracts (deployed only on Ethereum mainnet)
	views.Register("evm", "CCIPHome", "1.6.0", ViewCCIPHome)
	views.Register("evm", "RMNHome", "1.6.0", ViewRMNHome)
}
