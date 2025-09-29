package changesets

import (
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// TODO : move this to a common chain agnostic package
type MCMSOperatorConfig struct {
	// OverridePreviousRoot indicates whether to override the root of the MCMS contract.
	// if unset, default behavior is to not override the root.
	OverridePreviousRoot bool
	// TimelockDelay is the amount of time each operation in the proposal must wait before it can be executed.
	// if unset, default behavior is 3 hours.
	TimelockDelay mcms_types.Duration
	// TimelockAction is the action to perform on the timelock contract (schedule, bypass, or cancel).
	// if unset, default behavior is to schedule.
	TimelockAction mcms_types.TimelockAction
}