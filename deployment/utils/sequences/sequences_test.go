package sequences_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

func newBundle(t *testing.T) cldf_ops.Bundle {
	t.Helper()
	lggr := logger.Test(t)
	return cldf_ops.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		cldf_ops.NewMemoryReporter(),
	)
}

// outputSequence returns a sequence that emits a fixed output (or error).
func outputSequence(output sequences.OnChainOutput, err error) *cldf_ops.Sequence[struct{}, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"test-sequence",
		semver.MustParse("1.0.0"),
		"test sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ struct{}) (sequences.OnChainOutput, error) {
			return output, err
		},
	)
}

var emptyChains = cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{})

// ---- RunAndMergeSequence ----

func TestRunAndMergeSequence_MergesAddressesAndBatchOps(t *testing.T) {
	initial := sequences.OnChainOutput{
		Addresses: []datastore.AddressRef{{Address: "0xAA", Type: "TypeA"}},
		BatchOps:  []mcms_types.BatchOperation{{ChainSelector: 1}},
	}
	seq := outputSequence(sequences.OnChainOutput{
		Addresses: []datastore.AddressRef{{Address: "0xBB", Type: "TypeB"}},
		BatchOps:  []mcms_types.BatchOperation{{ChainSelector: 2}},
	}, nil)

	result, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, initial)

	require.NoError(t, err)
	assert.Len(t, result.Addresses, 2)
	assert.Len(t, result.BatchOps, 2)
}

func TestRunAndMergeSequence_MergesContractMetadata(t *testing.T) {
	initial := sequences.OnChainOutput{
		Metadata: sequences.Metadata{
			Contracts: []datastore.ContractMetadata{{Address: "0xAA", ChainSelector: 1}},
		},
	}
	seq := outputSequence(sequences.OnChainOutput{
		Metadata: sequences.Metadata{
			Contracts: []datastore.ContractMetadata{{Address: "0xBB", ChainSelector: 2}},
		},
	}, nil)

	result, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, initial)

	require.NoError(t, err)
	assert.Len(t, result.Metadata.Contracts, 2)
}

func TestRunAndMergeSequence_SetsChainMetadataWhenAggregateHasNone(t *testing.T) {
	chain := &datastore.ChainMetadata{ChainSelector: 99}
	seq := outputSequence(sequences.OnChainOutput{
		Metadata: sequences.Metadata{Chain: chain},
	}, nil)

	result, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, sequences.OnChainOutput{})

	require.NoError(t, err)
	require.NotNil(t, result.Metadata.Chain)
	assert.Equal(t, uint64(99), result.Metadata.Chain.ChainSelector)
}

func TestRunAndMergeSequence_ErrorsOnConflictingChainMetadata(t *testing.T) {
	initial := sequences.OnChainOutput{
		Metadata: sequences.Metadata{Chain: &datastore.ChainMetadata{ChainSelector: 1}},
	}
	seq := outputSequence(sequences.OnChainOutput{
		Metadata: sequences.Metadata{Chain: &datastore.ChainMetadata{ChainSelector: 2}},
	}, nil)

	_, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, initial)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "conflicting chain metadata")
}

func TestRunAndMergeSequence_SetsEnvMetadataWhenAggregateHasNone(t *testing.T) {
	seq := outputSequence(sequences.OnChainOutput{
		Metadata: sequences.Metadata{Env: &datastore.EnvMetadata{}},
	}, nil)

	result, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, sequences.OnChainOutput{})

	require.NoError(t, err)
	require.NotNil(t, result.Metadata.Env)
}

func TestRunAndMergeSequence_ErrorsOnConflictingEnvMetadata(t *testing.T) {
	initial := sequences.OnChainOutput{
		Metadata: sequences.Metadata{Env: &datastore.EnvMetadata{}},
	}
	seq := outputSequence(sequences.OnChainOutput{
		Metadata: sequences.Metadata{Env: &datastore.EnvMetadata{}},
	}, nil)

	_, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, initial)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "conflicting env metadata")
}

func TestRunAndMergeSequence_PropagatesSequenceError(t *testing.T) {
	seq := outputSequence(sequences.OnChainOutput{}, errors.New("inner failure"))

	_, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, sequences.OnChainOutput{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "inner failure")
}

func TestRunAndMergeSequence_EmptyAggregateAndEmptyOutput(t *testing.T) {
	seq := outputSequence(sequences.OnChainOutput{}, nil)

	result, err := sequences.RunAndMergeSequence(newBundle(t), emptyChains, seq, struct{}{}, sequences.OnChainOutput{})

	require.NoError(t, err)
	assert.Empty(t, result.Addresses)
	assert.Empty(t, result.BatchOps)
	assert.Empty(t, result.Metadata.Contracts)
}

// ---- WriteMetadataToDatastore ----

func TestWriteMetadataToDatastore_WritesAllFields(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	err := sequences.WriteMetadataToDatastore(ds, sequences.Metadata{
		Contracts: []datastore.ContractMetadata{
			{Address: "0xAA", ChainSelector: 42},
			{Address: "0xBB", ChainSelector: 42},
		},
		Chain: &datastore.ChainMetadata{ChainSelector: 42},
		Env:   &datastore.EnvMetadata{},
	})
	require.NoError(t, err)

	contracts, err := ds.ContractMetadata().Fetch()
	require.NoError(t, err)
	assert.Len(t, contracts, 2)

	_, err = ds.ChainMetadata().Get(datastore.NewChainMetadataKey(42))
	require.NoError(t, err)

	_, err = ds.EnvMetadata().Get()
	require.NoError(t, err)
}

func TestWriteMetadataToDatastore_NilChainAndEnvAreNoOps(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	// No Chain, no Env — should not error
	err := sequences.WriteMetadataToDatastore(ds, sequences.Metadata{})
	require.NoError(t, err)
}

func TestWriteMetadataToDatastore_UpsertOverwritesExistingContractMetadata(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	first := sequences.Metadata{
		Contracts: []datastore.ContractMetadata{{Address: "0xAA", ChainSelector: 1, Metadata: []byte(`{"v":1}`)}},
	}
	second := sequences.Metadata{
		Contracts: []datastore.ContractMetadata{{Address: "0xAA", ChainSelector: 1, Metadata: []byte(`{"v":2}`)}},
	}

	require.NoError(t, sequences.WriteMetadataToDatastore(ds, first))
	require.NoError(t, sequences.WriteMetadataToDatastore(ds, second))

	all, err := ds.ContractMetadata().Fetch()
	require.NoError(t, err)
	// Upsert should collapse to 1 record (same key).
	assert.Len(t, all, 1)
}
