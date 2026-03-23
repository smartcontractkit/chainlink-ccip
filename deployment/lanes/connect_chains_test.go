package lanes

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateChainDefinition(t *testing.T) {
	t.Parallel()

	t.Run("valid input passes", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, validateChainDefinition(ChainDefinition{
			Selector: 1,
		}))
	})

	programmaticFields := []struct {
		name string
		def  ChainDefinition
	}{
		{"OnRamp", ChainDefinition{OnRamp: []byte{0x01}}},
		{"OffRamp", ChainDefinition{OffRamp: []byte{0x01}}},
		{"Router", ChainDefinition{Router: []byte{0x01}}},
		{"FeeQuoter", ChainDefinition{FeeQuoter: []byte{0x01}}},
		{"FeeQuoterDestChainConfig", ChainDefinition{FeeQuoterDestChainConfig: FeeQuoterDestChainConfig{IsEnabled: true}}},
		{"FeeQuoterVersion", ChainDefinition{FeeQuoterVersion: semver.MustParse("1.6.0")}},
	}

	for _, tc := range programmaticFields {
		t.Run(tc.name+" is rejected", func(t *testing.T) {
			t.Parallel()
			err := validateChainDefinition(tc.def)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.name)
			assert.Contains(t, err.Error(), "must not be set by the caller")
		})
	}
}
