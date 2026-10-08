package mcms_test

import (
	"testing"
	"time"

	"github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
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
				ValidUntil:     uint32(time.Now().UTC().Add(2 * time.Hour).Unix()), // some time in the future
			},
		},
		{
			desc: "happy path - bypass",
			input: mcms.Input{
				TimelockAction: types.TimelockActionBypass,
				ValidUntil:     uint32(time.Now().UTC().Add(2 * time.Hour).Unix()), // some time in the future
			},
		},
		{
			desc: "happy path - cancel",
			input: mcms.Input{
				TimelockAction: types.TimelockActionCancel,
				ValidUntil:     uint32(time.Now().UTC().Add(2 * time.Hour).Unix()), // some time in the future
			},
		},
		{
			desc: "invalid action",
			input: mcms.Input{
				TimelockAction: "InvalidAction",
				ValidUntil:     uint32(time.Now().UTC().Add(2 * time.Hour).Unix()), // some time in the future
			},
			expectedErr: "invalid timelock action: InvalidAction",
		},
		{
			desc: "validUntil in the past",
			input: mcms.Input{
				TimelockAction: types.TimelockActionSchedule,
				ValidUntil:     189999999, // some time in the future
			},
			expectedErr: "failed to validate MCMS input: ValidUntil must be in the future",
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

func TestInput_PopulateDefaults(t *testing.T) {
	t.Run("empty input defaults to schedule and DefaultValidDuration from now", func(t *testing.T) {
		input := mcms.Input{}
		before := time.Now().Add(mcms.DefaultValidDuration).Unix()
		require.NoError(t, input.PopulateDefaults())
		after := time.Now().Add(mcms.DefaultValidDuration).Unix()

		require.Equal(t, types.TimelockActionSchedule, input.TimelockAction)
		require.GreaterOrEqual(t, int64(input.ValidUntil), before)
		require.LessOrEqual(t, int64(input.ValidUntil), after)
		require.NoError(t, input.Validate())
	})

	t.Run("explicit values are kept", func(t *testing.T) {
		validUntil := uint32(time.Now().Add(365 * 24 * time.Hour).Unix())
		input := mcms.Input{
			TimelockAction: types.TimelockActionBypass,
			ValidUntil:     validUntil,
		}
		require.NoError(t, input.PopulateDefaults())

		require.Equal(t, types.TimelockActionBypass, input.TimelockAction)
		require.Equal(t, validUntil, input.ValidUntil)
	})
}
