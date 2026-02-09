package v1_6_1

import (
	"give-me-state-v2/views"
)

func init() {
	// Register v1.6.1 views
	// These use manual decoding via common utilities
	views.Register("evm", "BurnMintTokenPool", "1.6.1", ViewBurnMintTokenPool)
	views.Register("evm", "LockReleaseTokenPool", "1.6.1", ViewLockReleaseTokenPool)
}
