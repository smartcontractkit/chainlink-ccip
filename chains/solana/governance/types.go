// Package governance holds the Solana MCMS/governance helpers shared by the CCIP 1.6 and 2.0
// deployment code. It must not import chainlink-ccip's deployment module or any MCMS gobindings:
// the mcms library registers the decoders for those programs, and a second registration panics.
package governance

import (
	"github.com/gagliardetto/solana-go"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// The string values are written to datastores and must stay byte-identical.
const (
	TimelockProgramType         cldf_datastore.ContractType = "RBACTimelockProgram"
	McmProgramType              cldf_datastore.ContractType = "ManyChainMultiSigProgram"
	AccessControllerProgramType cldf_datastore.ContractType = "AccessControllerProgram"
	// special type for Solana that encodes PDA seed usage
	ProposerSeed     cldf_datastore.ContractType = "ProposerSeed"
	CancellerSeed    cldf_datastore.ContractType = "CancellerSeed"
	BypasserSeed     cldf_datastore.ContractType = "BypasserSeed"
	RBACTimelockSeed cldf_datastore.ContractType = "RBACTimelockSeed"
	// access controller accounts
	ProposerAccessControllerAccount  cldf_datastore.ContractType = "ProposerAccessControllerAccount"
	ExecutorAccessControllerAccount  cldf_datastore.ContractType = "ExecutorAccessControllerAccount"
	CancellerAccessControllerAccount cldf_datastore.ContractType = "CancellerAccessControllerAccount"
	BypasserAccessControllerAccount  cldf_datastore.ContractType = "BypasserAccessControllerAccount"
)

// Ref types of the MCMS instances.
const (
	RBACTimelock               cldf_datastore.ContractType = "RBACTimelock"
	ProposerManyChainMultisig  cldf_datastore.ContractType = "ProposerManyChainMultiSig"
	CancellerManyChainMultisig cldf_datastore.ContractType = "CancellerManyChainMultiSig"
	BypasserManyChainMultisig  cldf_datastore.ContractType = "BypasserManyChainMultiSig"
)

// TransferOwnershipParams contains common parameters for transferring ownership of a program.
type TransferOwnershipParams struct {
	Program      solana.PublicKey
	CurrentOwner solana.PublicKey
	NewOwner     solana.PublicKey
}
