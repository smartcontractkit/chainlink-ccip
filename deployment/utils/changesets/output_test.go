package changesets_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestWithDatastore(t *testing.T) {
	b := changesets.NewOutputBuilder(deployment.Environment{}, changesets.GetRegistry())
	out, err := b.WithDataStore(datastore.NewMemoryDataStore()).Build(mcms.Input{})
	require.NoError(t, err, "Build should not error")
	require.NotNil(t, out.DataStore, "DataStore should be set in ChangesetOutput")
}

func TestWithReports(t *testing.T) {
	b := changesets.NewOutputBuilder(deployment.Environment{}, changesets.GetRegistry())
	reports := []operations.Report[any, any]{
		{},
	}
	out, err := b.WithReports(reports).Build(mcms.Input{})
	require.NoError(t, err, "Build should not error")
	require.Len(t, out.Reports, 1, "Reports should be set in ChangesetOutput")
}

func TestWithBatchOps(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	err := ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 5009297550715157269,
		Type:          "Timelock",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x01",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 5009297550715157269,
		Type:          "MCM",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x02",
	})
	require.NoError(t, err)

	registry := changesets.GetRegistry()
	b := changesets.NewOutputBuilder(deployment.Environment{
		DataStore: ds.Seal(),
	}, registry)
	batchOps := []mcms_types.BatchOperation{
		{
			ChainSelector: 5009297550715157269,
			Transactions: []mcms_types.Transaction{
				{
					To:               "0x01",
					Data:             []byte("0xdeadbeef"),
					AdditionalFields: json.RawMessage{},
				},
			},
		},
	}
	out, err := b.WithBatchOps(batchOps).Build(mcms.Input{
		OverridePreviousRoot: false,
		ValidUntil:           2756219818,
		TimelockDelay:        mcms_types.NewDuration(3 * time.Hour),
		TimelockAction:       mcms_types.TimelockActionSchedule,
		Description:          "Proposal",
	})
	require.NoError(t, err, "Build should not error")
	require.Len(t, out.MCMSTimelockProposals, 1, "Proposal should exist")
	require.Equal(t, uint64(OP_COUNT), out.MCMSTimelockProposals[0].ChainMetadata[5009297550715157269].StartingOpCount)
}
