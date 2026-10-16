package shared

import (
	"github.com/Masterminds/semver/v3"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// Version1_0_0 is reproduced from chainlink/deployment/version.go so the ported changesets keep
// recording exactly what they recorded before the move.
//
// It is known to be wrong: no Solana program has ever been at 1.0.0, and an audit of the live
// datastores found 44 Solana rows carrying it, 43 of them in the two staging environments. It is
// preserved here rather than corrected because correcting it re-keys production rows -- the
// datastore key is (chain, type, version, qualifier) -- which belongs with the qualifier migration
// and not inside a port whose whole value is being a provably mechanical copy.
var Version1_0_0 = *semver.MustParse("1.0.0")

// MCMS program types on Solana, reproduced from chainlink/deployment/common/types.
//
// These name the deployed programs, as distinct from the accounts and seeds that live under them
// (ProposerManyChainMultisig, RBACTimelock and friends), which chainlink-ccip/deployment/utils
// already declares and which the ported code takes from there.
const (
	ManyChainMultisigProgram cldf.ContractType = "ManyChainMultiSigProgram"
	RBACTimelockProgram      cldf.ContractType = "RBACTimelockProgram"
	AccessControllerProgram  cldf.ContractType = "AccessControllerProgram"
)

// Coalesce returns *p when p is non-nil and fallback otherwise. Reproduced from
// chainlink/deployment/ccip/internal/pointer, which is an internal package of the node module and
// so cannot be imported from here even if the dependency were acceptable.
func Coalesce[T any](p *T, fallback T) T {
	if p != nil {
		return *p
	}

	return fallback
}
