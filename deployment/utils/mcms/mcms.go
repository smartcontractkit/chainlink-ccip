package mcms

import (
	"errors"
	"fmt"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// Input is the configuration for the MCMS proposal.
type Input struct {
	// OverridePreviousRoot indicates whether to override the root of the MCMS contract.
	OverridePreviousRoot bool
	// ValidUntil is a unix timestamp indicating when the proposal expires.
	// Root can't be set or executed after this time.
	ValidUntil uint32
	// TimelockDelay is the amount of time each operation in the proposal must wait before it can be executed.
	TimelockDelay mcms_types.Duration
	// TimelockAction is the action to perform on the timelock contract (schedule, bypass, or cancel).
	TimelockAction mcms_types.TimelockAction
	// MCMSAddressRef is a reference to the MCMS contract address in the datastore.
	MCMSAddressRef datastore.AddressRef
	// TimelockAddressRef is a reference to the timelock contract address in the datastore.
	TimelockAddressRef datastore.AddressRef
	// Description is a human-readable description of the proposal.
	Description string
}

// Validate validates the MCMS input.
func (c *Input) Validate() error {
	if c.TimelockAction != mcms_types.TimelockActionSchedule &&
		c.TimelockAction != mcms_types.TimelockActionBypass &&
		c.TimelockAction != mcms_types.TimelockActionCancel {
		return fmt.Errorf("invalid timelock action: %s", c.TimelockAction)
	}

	if datastore_utils.IsAddressRefEmpty(c.MCMSAddressRef) {
		return errors.New("mcms address ref is empty")
	}

	if datastore_utils.IsAddressRefEmpty(c.TimelockAddressRef) {
		return errors.New("timelock address ref is empty")
	}

	if c.ValidUntil == 0 {
		return errors.New("valid until timestamp must be set")
	}

	return nil
}
