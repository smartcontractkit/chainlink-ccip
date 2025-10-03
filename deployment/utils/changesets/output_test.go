package changesets_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
)

func TestWithDatastore(t *testing.T) {
	b := changesets.NewOutputBuilder(deployment.Environment{})
	out, err := b.WithDataStore(datastore.NewMemoryDataStore()).Build(mcms.Input{})
	require.NoError(t, err, "Build should not error")
	require.NotNil(t, out.DataStore, "DataStore should be set in ChangesetOutput")
}

func TestWithReports(t *testing.T) {
	b := changesets.NewOutputBuilder(deployment.Environment{})
	reports := []operations.Report[any, any]{
		{},
	}
	out, err := b.WithReports(reports).Build(mcms.Input{})
	require.NoError(t, err, "Build should not error")
	require.Len(t, out.Reports, 1, "Reports should be set in ChangesetOutput")
}

func TestWithWriteOutputs(t *testing.T) {
	tests := []struct {
		desc     string
		execInfo *contract.ExecInfo
	}{
		{
			execInfo: nil,
			desc:     "Tx not executed",
		},
		{
			execInfo: &contract.ExecInfo{
				Hash: common.HexToHash("0x02").Hex(),
			},
			desc: "Tx executed",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ds := datastore.NewMemoryDataStore()
			err := ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: 5009297550715157269,
				Type:          "Timelock",
				Version:       semver.MustParse("1.0.0"),
				Address:       common.HexToAddress("0x01").Hex(),
			})
			require.NoError(t, err)
			err = ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: 5009297550715157269,
				Type:          "MCM",
				Version:       semver.MustParse("1.0.0"),
				Address:       common.HexToAddress("0x02").Hex(),
			})
			require.NoError(t, err)

			b := changesets.NewOutputBuilder(deployment.Environment{
				DataStore: ds.Seal(),
			})
			out, err := b.WithWriteOutputs([]contract.WriteOutput{
				{
					ChainSelector: 5009297550715157269,
					ExecInfo:      test.execInfo,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: json.RawMessage{},
					},
				},
			}).Build(mcms.Input{
				OverridePreviousRoot: false,
				ValidUntil:           2756219818,
				TimelockDelay:        mcms_types.NewDuration(3 * time.Hour),
				TimelockAction:       mcms_types.TimelockActionSchedule,
				Description:          "Proposal",
				MCMSAddressRef: datastore.AddressRef{
					Type:    "MCM",
					Version: semver.MustParse("1.0.0"),
				},
				TimelockAddressRef: datastore.AddressRef{
					Type:    "Timelock",
					Version: semver.MustParse("1.0.0"),
				},
			})
			require.NoError(t, err, "Build should not error")
			if test.execInfo == nil {
				require.Len(t, out.MCMSTimelockProposals, 1, "Proposal should exist")
				require.Equal(t, uint64(OP_COUNT), out.MCMSTimelockProposals[0].ChainMetadata[5009297550715157269].StartingOpCount)
			} else {
				require.Len(t, out.MCMSTimelockProposals, 0, "Proposal should not exist")
			}
		})
	}
}
