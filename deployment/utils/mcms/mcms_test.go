package mcms_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

func TestInput_Validate(t *testing.T) {
	tests := []struct {
		desc        string
		input       mcms.Input
		expectedErr string
	}{
		{
			desc: "happy path - schedule",
			input: mcms.Input{
				TimelockAction: types.TimelockActionSchedule,
			},
		},
		{
			desc: "happy path - bypass",
			input: mcms.Input{
				TimelockAction: types.TimelockActionBypass,
			},
		},
		{
			desc: "happy path - cancel",
			input: mcms.Input{
				TimelockAction: types.TimelockActionCancel,
			},
		},
		{
			desc: "invalid action",
			input: mcms.Input{
				TimelockAction: "InvalidAction",
			},
			expectedErr: "invalid timelock action: InvalidAction",
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
