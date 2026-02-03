package aptos

import (
	"call-orchestrator-demo/views"
)

func init() {
	// CCIP (main module - includes FeeQuoter, RMNRemote, TokenAdminRegistry, NonceManager, ReceiverRegistry)
	views.Register("aptos", "AptosCCIP", "1.6.0", ViewCCIP)
	views.Register("aptos", "AptosCCIP", "", ViewCCIP)

	// Router
	views.Register("aptos", "AptosRouter", "1.6.0", ViewRouter)
	views.Register("aptos", "AptosRouter", "", ViewRouter)

	// OnRamp
	views.Register("aptos", "AptosOnRamp", "1.6.0", ViewOnRamp)
	views.Register("aptos", "AptosOnRamp", "", ViewOnRamp)

	// OffRamp
	views.Register("aptos", "AptosOffRamp", "1.6.0", ViewOffRamp)
	views.Register("aptos", "AptosOffRamp", "", ViewOffRamp)

	// Token Pools
	views.Register("aptos", "AptosManagedTokenPool", "1.6.0", ViewTokenPool)
	views.Register("aptos", "AptosManagedTokenPool", "", ViewTokenPool)
	views.Register("aptos", "AptosRegulatedTokenPool", "1.6.0", ViewTokenPool)
	views.Register("aptos", "AptosRegulatedTokenPool", "", ViewTokenPool)
	views.Register("aptos", "AptosBurnMintTokenPool", "1.6.0", ViewTokenPool)
	views.Register("aptos", "AptosBurnMintTokenPool", "", ViewTokenPool)
	views.Register("aptos", "AptosLockReleaseTokenPool", "1.6.0", ViewTokenPool)
	views.Register("aptos", "AptosLockReleaseTokenPool", "", ViewTokenPool)

	// MCMS
	views.Register("aptos", "AptosManyChainMultisig", "1.6.0", ViewMCMS)
	views.Register("aptos", "AptosManyChainMultisig", "", ViewMCMS)

	// Tokens
	views.Register("aptos", "AptosManagedTokenType", "1.6.0", ViewToken)
	views.Register("aptos", "AptosManagedTokenType", "", ViewToken)
	views.Register("aptos", "LinkToken", "1.6.0", ViewToken)
	views.Register("aptos", "LinkToken", "", ViewToken)

	// Generic fallback for any other Aptos types
	views.Register("aptos", "AptosGeneric", "1.6.0", ViewGenericAccount)
	views.Register("aptos", "AptosGeneric", "", ViewGenericAccount)
}
