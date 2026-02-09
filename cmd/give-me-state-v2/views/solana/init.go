package solana

import (
	"give-me-state-v2/views"
)

func init() {
	// FeeQuoter
	views.Register("svm", "FeeQuoter", "1.6.0", ViewFeeQuoter)
	views.Register("svm", "FeeQuoter", "", ViewFeeQuoter)

	// Router
	views.Register("svm", "Router", "1.6.0", ViewRouter)
	views.Register("svm", "Router", "", ViewRouter)

	// OffRamp
	views.Register("svm", "OffRamp", "1.6.0", ViewOffRamp)
	views.Register("svm", "OffRamp", "", ViewOffRamp)

	// RMNRemote
	views.Register("svm", "RMNRemote", "1.6.0", ViewRMNRemote)
	views.Register("svm", "RMNRemote", "", ViewRMNRemote)

	// TokenPools
	views.Register("svm", "BurnMintTokenPool", "1.6.0", ViewTokenPool)
	views.Register("svm", "BurnMintTokenPool", "", ViewTokenPool)
	views.Register("svm", "LockReleaseTokenPool", "1.6.0", ViewTokenPool)
	views.Register("svm", "LockReleaseTokenPool", "", ViewTokenPool)

	// SPL Tokens
	views.Register("svm", "SPLTokens", "1.6.0", ViewSPLToken)
	views.Register("svm", "SPLTokens", "1.0.0", ViewSPLToken)
	views.Register("svm", "SPLTokens", "", ViewSPLToken)
	views.Register("svm", "SPL2022Tokens", "1.6.0", ViewSPLToken)
	views.Register("svm", "SPL2022Tokens", "", ViewSPLToken)

	// LINK Token
	views.Register("svm", "LinkToken", "1.6.0", ViewLinkToken)
	views.Register("svm", "LinkToken", "", ViewLinkToken)

	// Receiver
	views.Register("svm", "Receiver", "1.6.0", ViewReceiver)
	views.Register("svm", "Receiver", "", ViewReceiver)

	// Remote Source/Dest (lane configs)
	views.Register("svm", "RemoteSource", "1.6.0", ViewRemoteSource)
	views.Register("svm", "RemoteSource", "", ViewRemoteSource)
	views.Register("svm", "RemoteDest", "1.6.0", ViewRemoteDest)
	views.Register("svm", "RemoteDest", "", ViewRemoteDest)

	// MCMS
	views.Register("svm", "MCM", "1.6.0", ViewMCMConfig)
	views.Register("svm", "MCM", "", ViewMCMConfig)
	views.Register("svm", "AccessController", "1.6.0", ViewAccessController)
	views.Register("svm", "AccessController", "", ViewAccessController)

	// Timelock
	views.Register("svm", "Timelock", "1.6.0", ViewTimelock)
	views.Register("svm", "Timelock", "", ViewTimelock)
	views.Register("svm", "RBACTimelock", "1.6.0", ViewTimelock)
	views.Register("svm", "RBACTimelock", "", ViewTimelock)

	// TokenPoolLookupTable (generic for now)
	views.Register("svm", "TokenPoolLookupTable", "1.6.0", ViewGenericAccount)
	views.Register("svm", "TokenPoolLookupTable", "", ViewGenericAccount)

	// SVMSignerRegistry (generic for now)
	views.Register("svm", "SVMSignerRegistry", "1.6.0", ViewGenericAccount)
	views.Register("svm", "SVMSignerRegistry", "", ViewGenericAccount)
}
