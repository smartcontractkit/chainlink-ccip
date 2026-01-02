package mcms_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

func TestInput_Validate(t *testing.T) {
	mcmsAddressRef := datastore.AddressRef{
		Type:    "MCM",
		Version: semver.MustParse("1.0.0"),
	}
	timelockAddressRef := datastore.AddressRef{
		Type:    "Timelock",
		Version: semver.MustParse("1.0.0"),
	}
	validUntil := uint32(3759765795)

	tests := []struct {
		desc        string
		input       mcms.Input
		expectedErr string
	}{
		{
			desc: "happy path - schedule",
			input: mcms.Input{
				TimelockAction:     types.TimelockActionSchedule,
				MCMSAddressRef:     mcmsAddressRef,
				TimelockAddressRef: timelockAddressRef,
				ValidUntil:         validUntil,
			},
		},
		{
			desc: "happy path - bypass",
			input: mcms.Input{
				TimelockAction:     types.TimelockActionBypass,
				MCMSAddressRef:     mcmsAddressRef,
				TimelockAddressRef: timelockAddressRef,
				ValidUntil:         validUntil,
			},
		},
		{
			desc: "happy path - cancel",
			input: mcms.Input{
				TimelockAction:     types.TimelockActionCancel,
				MCMSAddressRef:     mcmsAddressRef,
				TimelockAddressRef: timelockAddressRef,
				ValidUntil:         validUntil,
			},
		},
		{
			desc: "invalid action",
			input: mcms.Input{
				TimelockAction: "InvalidAction",
			},
			expectedErr: "invalid timelock action: InvalidAction",
		},
		{
			desc: "invalid mcms address ref",
			input: mcms.Input{
				TimelockAction: types.TimelockActionSchedule,
				MCMSAddressRef: datastore.AddressRef{},
			},
			expectedErr: "mcms address ref is empty",
		},
		{
			desc: "invalid timelock address ref",
			input: mcms.Input{
				TimelockAction:     types.TimelockActionSchedule,
				MCMSAddressRef:     mcmsAddressRef,
				TimelockAddressRef: datastore.AddressRef{},
			},
			expectedErr: "timelock address ref is empty",
		},
		{
			desc: "invalid valid until timestamp",
			input: mcms.Input{
				TimelockAction:     types.TimelockActionSchedule,
				MCMSAddressRef:     mcmsAddressRef,
				TimelockAddressRef: timelockAddressRef,
			},
			expectedErr: "valid until timestamp must be set",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.input.Validate()
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}
