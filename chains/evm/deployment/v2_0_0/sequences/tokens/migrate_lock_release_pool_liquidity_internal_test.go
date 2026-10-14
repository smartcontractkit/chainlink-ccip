package tokens

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// TestValidatePoolPair_SubtypeMatch pins the source/destination subtype agreement. The migration
// path is selected from the old pool's subtype, and the two v2.0.0 subtypes expose different
// lockbox ABIs, so a mismatched pair must be rejected before any read is attempted.
func TestValidatePoolPair_SubtypeMatch(t *testing.T) {
	oldPool := common.HexToAddress("0x0000000000000000000000000000000000000001")
	newPool := common.HexToAddress("0x0000000000000000000000000000000000000002")
	oldVersion := semver.MustParse("1.6.1")

	tests := []struct {
		name        string
		oldType     cldf.ContractType
		newType     cldf.ContractType
		expectedErr string
	}{
		{
			name:    "Success - plain source to plain destination",
			oldType: utils.LockReleaseTokenPool,
			newType: utils.LockReleaseTokenPool,
		},
		{
			name:    "Success - siloed source to siloed destination",
			oldType: utils.SiloedLockReleaseTokenPool,
			newType: utils.SiloedLockReleaseTokenPool,
		},
		{
			name:        "Failure - siloed source to plain destination",
			oldType:     utils.SiloedLockReleaseTokenPool,
			newType:     utils.LockReleaseTokenPool,
			expectedErr: "must both be siloed or both be unsiloed",
		},
		{
			name:        "Failure - plain source to siloed destination",
			oldType:     utils.LockReleaseTokenPool,
			newType:     utils.SiloedLockReleaseTokenPool,
			expectedErr: "must both be siloed or both be unsiloed",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validatePoolPair(oldPool, newPool,
				type_and_version.TypeAndVersion{Type: tc.oldType, Version: oldVersion},
				type_and_version.TypeAndVersion{Type: tc.newType, Version: utils.Version_2_0_0},
			)
			if tc.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.expectedErr)
		})
	}
}
