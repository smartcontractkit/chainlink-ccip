// Package strategy provides per-token-contract-type behaviors for EVM
// token deployment and pool wiring. A strategy encapsulates everything
// specific to one token contract type (e.g. BurnMintERC20, TIP-20):
// how to deploy the token, how to grant pool roles after a pool is
// deployed, how to grant an external admin role, and which capabilities
// the token contract supports.
//
// Strategies are looked up by ContractType from the singleton Registry
// and are independent of pool version, so adding a new token type makes
// it available to every pool-version adapter that consults the registry.
package strategy

import (
	"github.com/ethereum/go-ethereum/common"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// Capabilities reports the optional flow steps a token contract type
// participates in. The orchestrating sequence reads these flags to
// decide whether to invoke the corresponding step.
//
// ParticipatesInPoolRoleGrant is declared explicitly rather than inferred
// from a no-op GrantPoolRoles return so that intentional non-participation
// (e.g. plain ERC20) is distinguishable from a strategy bug that returned
// no writes by accident.
type Capabilities struct {
	SupportsAdminRole           bool
	SupportsCCIPAdmin           bool
	SupportsPreMint             bool
	ParticipatesInPoolRoleGrant bool
}

// EVMTokenStrategy encapsulates everything specific to one EVM token
// contract type. Implementations are registered with the singleton
// Registry keyed by ContractType.
type EVMTokenStrategy interface {
	// ContractType returns the deployment.ContractType used as the
	// registry key.
	ContractType() deployment.ContractType

	// Capabilities returns the static feature flags for this token type.
	Capabilities() Capabilities

	// Deploy performs the token contract deployment, returning the
	// resulting datastore reference and any token-side write outputs
	// produced during deployment. Implementations wrap either an
	// Operation (via contract.MaybeDeployContract) or a Sequence
	// (via cldf_ops.ExecuteSequence) as appropriate for the token type.
	Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (
		datastore.AddressRef, []evm_contract.WriteOutput, error)

	// GrantPoolRoles emits the writes that authorize a freshly-deployed
	// pool to mint/burn (or its TIP-20 issuer-role equivalent) against
	// this token. Returns (nil, nil) for token types that do not
	// participate in pool role granting; ParticipatesInPoolRoleGrant
	// is the authoritative flag, callers should consult it first.
	GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain,
		token, pool common.Address, chainSelector uint64) (
		[]evm_contract.WriteOutput, error)

	// GrantExternalAdmin grants the default-admin or contract-specific
	// admin role to externalAdmin. Implementations return (nil, nil)
	// for token types whose Capabilities.SupportsAdminRole is false;
	// callers should consult that flag first.
	GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain,
		token, externalAdmin common.Address, chainSelector uint64) (
		[]evm_contract.WriteOutput, error)
}
