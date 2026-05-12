package tokenimpl

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// CapabilitySet reports the optional flow steps a token contract type
// participates in. The orchestrating sequence reads these flags to
// decide whether to invoke the corresponding step.
type CapabilitySet struct {
	// ParticipatesInPoolRoleGrant is true when the token requires token-side
	// role grants for the pool to operate; GrantPoolRoles must emit those writes.
	ParticipatesInPoolRoleGrant bool

	// SupportsAdminRole is true when the token exposes a manageable admin or
	// default-admin role; GrantAdminRole and RevokeAdminRole must implement it.
	SupportsAdminRole bool

	// SupportsCCIPAdmin is true when the token has a token-level CCIP admin;
	// SetCCIPAdmin must emit the write that updates it.
	SupportsCCIPAdmin bool

	// SupportsPreMint is true when the token can mint during deployment and
	// transfer those tokens afterward to the configured recipient.
	SupportsPreMint bool
}

// Token encapsulates everything specific to one EVM token contract type.
type Token interface {
	// ContractType returns the deployment.ContractType used as the registry key.
	ContractType() deployment.ContractType

	// Capabilities returns the static feature flags for this token type.
	Capabilities() CapabilitySet

	// RevokeAdminRole revokes the default-admin or contract-specific admin
	// role from user. Callers should consult SupportsAdminRole first.
	RevokeAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error)

	// HasAdminRole returns whether user currently has the default-admin or
	// contract-specific admin role. Callers should consult SupportsAdminRole first.
	HasAdminRole(ctx context.Context, chain evm.Chain, token, user common.Address) (bool, error)

	// KnownAdminRoleHolders returns current admin role holders that can be
	// reconstructed from token-specific onchain state. It is best effort and is
	// used as an additional safety check before revoking an admin.
	KnownAdminRoleHolders(ctx context.Context, chain evm.Chain, token common.Address) ([]common.Address, error)

	// GrantAdminRole grants the default-admin or contract-specific
	// admin role to user. Returns an error for token types whose
	// Capabilities.SupportsAdminRole is false; callers should consult
	// that flag first.
	GrantAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error)

	// GrantPoolRoles emits the writes that authorize a freshly-deployed pool
	// to mint/burn (or its TIP-20 issuer-role equivalent) against this token.
	// proposalExecutor is the MCMS timelock (or zero when unused); BurnMintERC677
	// uses it for PrepareGrantMintAndBurnRoles. Other token types ignore it.
	// Returns an error for token types that don't participate in pool role
	// granting; ParticipatesInPoolRoleGrant is the authoritative flag, callers
	// should consult it first.
	GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool, proposalExecutor common.Address) ([]contract.WriteOutput, error)

	// SetCCIPAdmin sets the token-level CCIP admin where the token contract
	// supports one. Callers should consult SupportsCCIPAdmin first.
	SetCCIPAdmin(b cldf_ops.Bundle, chain evm.Chain, token, admin common.Address) ([]contract.WriteOutput, error)

	// Transfer emits the writes that transfer already-scaled token units from
	// the deployer to to, typically for post-deploy pre-mint distribution.
	Transfer(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error)

	// Deploy performs the token contract deployment, returning the
	// resulting datastore reference and any token-side write outputs
	// produced during deployment. Implementations may call lower-level
	// deployment operations or helpers, but batching is handled by the
	// outer token deployment sequence.
	Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error)
}
