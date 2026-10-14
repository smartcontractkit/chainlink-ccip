package mcms

import (
	"fmt"
	"time"

	mcms_types "github.com/smartcontractkit/mcms/types"
)

// DefaultValidDuration is how long a proposal stays valid when the input does not set ValidUntil.
const DefaultValidDuration = 7 * 24 * time.Hour

type Input struct {
	// OverridePreviousRoot indicates whether to override the root of the MCMS contract.
	OverridePreviousRoot bool
	// ValidUntil is a unix timestamp indicating when the proposal expires.
	// Root can't be set or executed after this time.
	// Defaults to DefaultValidDuration from proposal generation when unset.
	ValidUntil uint32
	// NO LONGER USED. THIS VALUE AUTO RESOLVES.
	TimelockDelay mcms_types.Duration
	// TimelockAction is the action to perform on the timelock contract (schedule, bypass, or cancel).
	TimelockAction mcms_types.TimelockAction
	// Qualifier is a string used to qualify the MCMS + Timelock contract addresses.
	Qualifier string
	// Description is a human-readable description of the proposal.
	Description string
}

// TODO : need to put more validation here
func (c *Input) Validate() error {
	if c.TimelockAction != mcms_types.TimelockActionSchedule &&
		c.TimelockAction != mcms_types.TimelockActionBypass &&
		c.TimelockAction != mcms_types.TimelockActionCancel {
		return fmt.Errorf("invalid timelock action: %s", c.TimelockAction)
	}
	if c.ValidUntil <= 0 {
		return fmt.Errorf("failed to validate MCMS input: ValidUntil must be a positive unix timestamp")
	}
	// check if ValidUntil is in the past
	// current time in utc plus 10 minutes to account for any potential clock drift or delays in proposal creation
	// this is to prevent proposals from being created with a ValidUntil that is already expired
	if c.ValidUntil < uint32(time.Now().Add(10*time.Minute).UTC().Unix()) {
		return fmt.Errorf("failed to validate MCMS input: ValidUntil must be in the future")
	}
	return nil
}

func (c *Input) PopulateDefaults() error {
	if c.TimelockAction == "" {
		c.TimelockAction = mcms_types.TimelockActionSchedule
	}
	if c.ValidUntil == 0 {
		// The timelock salt is derived from ValidUntil, so a generation-time value also gives
		// regenerated proposals with identical payloads distinct operation IDs.
		c.ValidUntil = uint32(time.Now().Add(DefaultValidDuration).Unix()) //nolint:gosec // G115: Unix timestamp fits in uint32 until 2106
	}
	return nil
}
