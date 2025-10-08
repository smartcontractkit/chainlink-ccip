package changesets

import (
	changeset "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
)

// RMNCurseAction represent a curse action to be applied on a chain (ChainSelector) with a specific subject (SubjectToCurse)
// The curse action will by applied by calling the Curse method on the RMNRemote contract on the chain (ChainSelector)
type RMNCurseAction struct {
	ChainSelector  uint64
	SubjectToCurse rmnops.Subject
}

// CurseAction is a function that returns a list of RMNCurseAction to be applied on a chain
// CurseChain, CurseLane, CurseGloballyOnlyOnSource are examples of function implementing CurseAction
type CurseAction func() ([]RMNCurseAction, error)

type RMNCurseConfig struct {
	MCMS         *changeset.MCMSInput
	CurseActions []CurseAction
	// Use this if you need to include lanes that are not in sourcechain in the offramp. i.e. not yet migrated lane from 1.5
	IncludeNotConnectedLanes bool
	// Use this if you want to include curse subject even when they are already cursed (CurseChangeset) or already uncursed (UncurseChangeset)
	Force  bool
	Reason string
}
