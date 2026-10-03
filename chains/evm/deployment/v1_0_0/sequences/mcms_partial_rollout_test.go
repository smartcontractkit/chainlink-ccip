package sequences

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

func TestMergePartialRolloutConfigAddsNextSignerBatch(t *testing.T) {
	t.Parallel()

	desired := types.Config{
		Quorum:  2,
		Signers: syntheticSignerAddresses(25),
	}

	current := types.Config{
		Quorum:  2,
		Signers: syntheticSignerAddresses(5),
	}

	remaining := 20

	got, err := mergePartialRolloutConfig(current, desired, &remaining)
	require.NoError(t, err)
	require.Len(t, got.Signers, 25)
	require.Zero(t, remaining)
}

func TestMergePartialRolloutConfigIsIdempotent(t *testing.T) {
	t.Parallel()

	desired := types.Config{
		Quorum:  2,
		Signers: syntheticSignerAddresses(10),
	}

	remaining := 20

	got, err := mergePartialRolloutConfig(desired, desired, &remaining)
	require.NoError(t, err)
	require.Len(t, got.Signers, len(desired.Signers))
	require.Equal(t, 20, remaining)
}

func TestMergePartialRolloutConfigRejectsUnexpectedSigner(t *testing.T) {
	t.Parallel()

	desired := types.Config{Signers: syntheticSignerAddresses(1)}
	current := types.Config{Signers: []common.Address{common.HexToAddress("0x00000000000000000000000000000000000000ff")}}
	remaining := 20

	_, err := mergePartialRolloutConfig(current, desired, &remaining)
	require.Error(t, err)
}

func TestMergePartialRolloutConfigPreservesOrderAndAvoidsDuplicates(t *testing.T) {
	t.Parallel()

	firstSigner := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	secondSigner := common.HexToAddress("0x00000000000000000000000000000000000000bb")
	thirdSigner := common.HexToAddress("0x00000000000000000000000000000000000000cc")

	desired := types.Config{
		Quorum:  2,
		Signers: []common.Address{firstSigner, secondSigner, thirdSigner},
	}

	current := types.Config{
		Quorum:  2,
		Signers: []common.Address{secondSigner, firstSigner},
	}
	remaining := 20

	got, err := mergePartialRolloutConfig(current, desired, &remaining)
	require.NoError(t, err)
	require.Equal(t, []common.Address{secondSigner, firstSigner, thirdSigner}, got.Signers)
}

func TestMergePartialRolloutConfigRejectsIncompatibleStructure(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		current types.Config
		desired types.Config
	}{
		"quorum mismatch": {
			current: types.Config{Quorum: 1},
			desired: types.Config{Quorum: 2},
		},
		"group count mismatch": {
			current: types.Config{Quorum: 2},
			desired: types.Config{Quorum: 2, GroupSigners: []types.Config{{Quorum: 1}}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			remaining := 20
			_, err := mergePartialRolloutConfig(test.current, test.desired, &remaining)
			require.Error(t, err)
		})
	}
}

func TestMergePartialRolloutConfigConvergesForProductionShapedConfig(t *testing.T) {
	t.Parallel()

	desired := types.Config{
		Quorum: 2,
		GroupSigners: []types.Config{
			{
				Quorum:       2,
				Signers:      syntheticSignerAddressesFrom(1, 24),
				GroupSigners: []types.Config{},
			},
			{
				Quorum:       2,
				Signers:      syntheticSignerAddressesFrom(25, 24),
				GroupSigners: []types.Config{},
			},
			{
				Quorum:       2,
				Signers:      syntheticSignerAddressesFrom(49, 24),
				GroupSigners: []types.Config{},
			},
		},
	}

	current := types.Config{
		Quorum: 2,
		GroupSigners: []types.Config{
			{Quorum: 2, GroupSigners: []types.Config{}},
			{Quorum: 2, GroupSigners: []types.Config{}},
			{Quorum: 2, GroupSigners: []types.Config{}},
		},
	}

	for proposal := 0; proposal < 4; proposal++ {
		remaining := 20
		updated, err := mergePartialRolloutConfig(current, desired, &remaining)
		require.NoError(t, err)
		current = updated
	}

	require.Equal(t, desired, current)
}

func syntheticSignerAddresses(count int) []common.Address {
	return syntheticSignerAddressesFrom(1, count)
}

func syntheticSignerAddressesFrom(start, count int) []common.Address {
	signers := make([]common.Address, count)
	for i := range signers {
		signers[i] = common.BigToAddress(big.NewInt(int64(start + i)))
	}

	return signers
}
