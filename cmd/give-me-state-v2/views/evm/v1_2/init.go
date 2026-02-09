package v1_2

import (
	"give-me-state-v2/views"
)

func init() {

	// Register v1.2 views
	views.Register("evm", "Router", "1.2.0", ViewRouter)
	views.Register("evm", "PriceRegistry", "1.2.0", ViewPriceRegistry)
}
