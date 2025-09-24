package changesets_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	changeset "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

func TestWithDatastore(t *testing.T) {
	b := changeset.NewOutputBuilder()
	out, err := b.WithDataStore(datastore.NewMemoryDataStore()).Build(changeset.MCMSBuildParams{})
	require.NoError(t, err, "Build should not error")
	require.NotNil(t, out.DataStore, "DataStore should be set in ChangesetOutput")
}

func TestWithReports(t *testing.T) {
	b := changeset.NewOutputBuilder()
	reports := []operations.Report[any, any]{
		{},
	}
	out, err := b.WithReports(reports).Build(changeset.MCMSBuildParams{})
	require.NoError(t, err, "Build should not error")
	require.Len(t, out.Reports, 1, "Reports should be set in ChangesetOutput")
}

func TestWithWriteOutputs(t *testing.T) {
	tests := []struct {
		desc     string
		executed bool
	}{
		{
			executed: false,
			desc:     "Tx not executed",
		},
		{
			executed: true,
			desc:     "Tx executed",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			b := changeset.NewOutputBuilder()
			out, err := b.WithWriteOutputs([]contract.WriteOutput{
				{
					ChainSelector: 5009297550715157269,
					Executed:      test.executed,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: json.RawMessage{},
					},
				},
			}).Build(changeset.MCMSBuildParams{
				Description: "Proposal",
				MCMSInput: changeset.MCMSInput{
					OverridePreviousRoot: false,
					ValidUntil:           2756219818,
					TimelockDelay:        mcms_types.NewDuration(3 * time.Hour),
					TimelockAction:       mcms_types.TimelockActionSchedule,
				},

				TimelockAddresses: map[mcms_types.ChainSelector]string{
					5009297550715157269: common.HexToAddress("0x02").Hex(),
				},
				ChainMetadata: map[mcms_types.ChainSelector]mcms_types.ChainMetadata{
					5009297550715157269: {
						StartingOpCount: 5,
						MCMAddress:      common.HexToAddress("0x03").Hex(),
					},
				},
			})
			require.NoError(t, err, "Build should not error")
			if !test.executed {
				require.Len(t, out.MCMSTimelockProposals, 1, "Proposal should exist")
			} else {
				require.Len(t, out.MCMSTimelockProposals, 0, "Proposal should not exist")
			}
		})
	}
}
