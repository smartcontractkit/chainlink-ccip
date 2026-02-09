package v1_5_1

import (
	"give-me-state-v2/views"
)

func init() {
	// Register v1.5.1 views
	views.Register("evm", "TokenPoolFactory", "1.5.1", ViewTokenPoolFactory)
	views.Register("evm", "BurnMintTokenPool", "1.5.1", ViewBurnMintTokenPool)
	views.Register("evm", "LockReleaseTokenPool", "1.5.1", ViewLockReleaseTokenPool)
	views.Register("evm", "BurnFromMintTokenPool", "1.5.1", ViewBurnFromMintTokenPool)
	views.Register("evm", "BurnWithFromMintTokenPool", "1.5.1", ViewBurnWithFromMintTokenPool)
}
