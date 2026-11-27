package contractreader_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip/consts"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	chainreadermocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

func TestExtendedContractReader(t *testing.T) {
	const contractName = "testContract"
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	extCr := contractreader.NewExtendedContractReader(cr)

	bindings := extCr.GetBindings(contractName)
	assert.Len(t, bindings, 0)

	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x123"}}).Return(nil)
	cr.On("Unbind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x123"}}).Return(nil)
	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x124"}}).Return(nil)
	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x125"}}).Return(fmt.Errorf("some err"))

	err := extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x123"}})
	assert.NoError(t, err)

	// ignored since 0x123 already exists
	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x123"}})
	assert.NoError(t, err)

	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x124"}})
	assert.NoError(t, err)

	// Bind fails
	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x125"}})
	assert.Error(t, err)

	bindings = extCr.GetBindings(contractName)
	assert.Len(t, bindings, 1)
	assert.Equal(t, "0x124", bindings[0].Binding.Address)
}

func TestExtendedContractReader_AllowMultiBindingForAggregator(t *testing.T) {
	const contractName = consts.ContractNamePriceAggregator
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	extCr := contractreader.NewExtendedContractReader(cr)

	bindings := extCr.GetBindings(contractName)
	assert.Len(t, bindings, 0)

	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x123"}}).Return(nil)
	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x124"}}).Return(nil)
	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x125"}}).Return(fmt.Errorf("some err"))

	err := extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x123"}})
	assert.NoError(t, err)

	// ignored since 0x123 already exists
	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x123"}})
	assert.NoError(t, err)

	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x124"}})
	assert.NoError(t, err)

	// Bind fails
	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x125"}})
	assert.Error(t, err)

	bindings = extCr.GetBindings(contractName)
	assert.Len(t, bindings, 2)
	assert.Equal(t, "0x123", bindings[0].Binding.Address)
	assert.Equal(t, "0x124", bindings[1].Binding.Address)
}

func TestDoubleWrap(t *testing.T) {
	var cr contractreader.ContractReaderFacade

	wrapped := contractreader.NewExtendedContractReader(cr)
	require.NotEqual(t, &cr, &wrapped)

	doubleWrapped := contractreader.NewExtendedContractReader(cr)
	require.Equal(t, wrapped, doubleWrapped)
}

func TestFinalityViolation(t *testing.T) {
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	cr.EXPECT().HealthReport().Return(map[string]error{"farmerwolfcabbagegoat": types.ErrFinalityViolated}).Times(3)
	cr.EXPECT().QueryKey(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	cr.EXPECT().GetLatestValue(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	cr.EXPECT().BatchGetLatestValues(mock.Anything, mock.Anything).Return(nil, nil)

	wrapped := contractreader.NewExtendedContractReader(cr)

	_, err := wrapped.QueryKey(
		t.Context(),
		types.BoundContract{},
		query.KeyFilter{},
		query.LimitAndSort{},
		nil)
	require.ErrorIs(t, err, contractreader.ErrFinalityViolated)

	err = wrapped.GetLatestValue(
		t.Context(),
		"",
		primitives.Finalized,
		nil,
		nil)
	require.ErrorIs(t, err, contractreader.ErrFinalityViolated)

	_, err = wrapped.BatchGetLatestValues(
		t.Context(),
		types.BatchGetLatestValuesRequest{})
	require.ErrorIs(t, err, contractreader.ErrFinalityViolated)
}

func TestExtendedContractReader_PreventRebinding(t *testing.T) {
	const contractName = consts.ContractNameOnRamp
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	extCr := contractreader.NewExtendedContractReader(cr)

	// Only one bind call.
	cr.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil).
		Times(1)

	// 2 calls to bind, the second should be a no-op.
	contracts := []types.BoundContract{{Name: contractName, Address: "0x123"}}
	err := extCr.Bind(context.Background(), contracts)
	assert.NoError(t, err)
	err = extCr.Bind(context.Background(), contracts)
	assert.NoError(t, err)

	bindings := extCr.GetBindings(contractName)
	assert.Len(t, bindings, 1)
	assert.Equal(t, "0x123", bindings[0].Binding.Address)

	// Assert that `Unbind` is not called.
	cr.AssertExpectations(t)
}

// TestExtendedContractReader_Rebinding tests that Unbind is called when rebinding.
func TestExtendedContractReader_Rebinding(t *testing.T) {
	const contractName = consts.ContractNameOnRamp
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	extCr := contractreader.NewExtendedContractReader(cr)

	c1 := []types.BoundContract{{Name: contractName, Address: "0x123"}}
	c2 := []types.BoundContract{{Name: contractName, Address: "0x456"}}

	// 2 bind calls, the second rebinds to a new address.
	cr.EXPECT().Bind(mock.Anything, c1).Return(nil).Times(1)
	cr.EXPECT().Bind(mock.Anything, c2).Return(nil).Times(1)
	// 1 unbind call
	cr.EXPECT().Unbind(mock.Anything, c1).Return(nil).Times(1)

	// 2 calls to bind, the second should trigger an unbind.
	err := extCr.Bind(context.Background(), c1)
	assert.NoError(t, err)
	err = extCr.Bind(context.Background(), c2)
	assert.NoError(t, err)

	bindings := extCr.GetBindings(contractName)
	assert.Len(t, bindings, 1)
	assert.Equal(t, "0x456", bindings[0].Binding.Address)

	// Assert that all expectations were met.
	cr.AssertExpectations(t)
}

func TestExtendedContractReader_QueryKeyDedup(t *testing.T) {
	type event struct {
		SeqNum uint64
		Value  *big.Int
	}

	row := func(block byte, logIndex int, tx byte, seqNum uint64) types.Sequence {
		return types.Sequence{
			Cursor: fmt.Sprintf("%d-%d-0x%02x", block, logIndex, tx),
			TxHash: []byte{tx},
			Head:   types.Head{Height: fmt.Sprint(block), Hash: []byte{block}},
			Data:   &event{SeqNum: seqNum, Value: big.NewInt(int64(seqNum))},
		}
	}
	seqNums := func(seqs []types.Sequence) []uint64 {
		out := make([]uint64, 0, len(seqs))
		for _, s := range seqs {
			out = append(out, s.Data.(*event).SeqNum)
		}
		return out
	}
	// twice stores each message twice under different log indexes.
	twice := func(from, to uint64) []types.Sequence {
		var rows []types.Sequence
		for s := from; s <= to; s++ {
			rows = append(rows, row(1, int(s)*10, byte(s), s), row(1, int(s)*10+4, byte(s), s))
		}
		return rows
	}

	tests := []struct {
		name       string
		rows       []types.Sequence
		limit      uint64
		expSeqNums []uint64
		expLimits  []uint64
		expWarn    bool
		expErr     error
	}{
		{
			name:       "no duplicates is a single untouched query",
			rows:       []types.Sequence{row(1, 0, 1, 697), row(1, 1, 2, 698), row(2, 0, 3, 699)},
			limit:      2,
			expSeqNums: []uint64{697, 698},
			expLimits:  []uint64{2},
		},
		{
			name:       "duplicates filling the page trigger a re-query",
			rows:       twice(697, 698),
			limit:      2,
			expSeqNums: []uint64{697, 698},
			expLimits:  []uint64{2, 4},
			expWarn:    true,
		},
		{
			name:       "result is cut to the requested limit",
			rows:       twice(697, 700),
			limit:      3,
			expSeqNums: []uint64{697, 698, 699},
			expLimits:  []uint64{3, 6},
			expWarn:    true,
		},
		{
			name:       "duplicates below the limit need no re-query",
			rows:       twice(697, 698),
			limit:      10,
			expSeqNums: []uint64{697, 698},
			expLimits:  []uint64{10},
			expWarn:    true,
		},
		{
			name:       "no limit",
			rows:       twice(697, 698),
			limit:      0,
			expSeqNums: []uint64{697, 698},
			expLimits:  []uint64{0},
			expWarn:    true,
		},
		{
			name:       "same tx and block with different data are kept",
			rows:       []types.Sequence{row(1, 0, 1, 697), row(1, 1, 1, 698)},
			limit:      2,
			expSeqNums: []uint64{697, 698},
			expLimits:  []uint64{2},
		},
		{
			name:       "same data in another block or tx is kept",
			rows:       []types.Sequence{row(1, 0, 1, 697), row(2, 0, 1, 697), row(2, 1, 2, 697)},
			limit:      3,
			expSeqNums: []uint64{697, 697, 697},
			expLimits:  []uint64{3},
		},
		{
			name: "sequences without tx hash are kept",
			rows: []types.Sequence{
				{Cursor: "a", Data: &event{SeqNum: 697}},
				{Cursor: "b", Data: &event{SeqNum: 697}},
			},
			limit:      2,
			expSeqNums: []uint64{697, 697},
			expLimits:  []uint64{2},
		},
		{
			name: "error when duplicates still fill the page after the last re-query",
			rows: func() []types.Sequence {
				rows := make([]types.Sequence, 0, 32)
				for i := range 32 {
					rows = append(rows, row(1, i, 1, 697))
				}
				return append(rows, row(1, 32, 2, 698))
			}(),
			limit:     2,
			expLimits: []uint64{2, 4, 8, 16},
			expErr:    contractreader.ErrDuplicatesExceedLimit,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var limits []uint64
			cr := chainreadermocks.NewMockContractReaderFacade(t)
			cr.EXPECT().HealthReport().Return(nil)
			cr.EXPECT().QueryKey(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				RunAndReturn(func(
					_ context.Context, _ types.BoundContract, _ query.KeyFilter, ls query.LimitAndSort, _ any,
				) ([]types.Sequence, error) {
					limits = append(limits, ls.Limit.Count)
					if ls.Limit.Count == 0 || ls.Limit.Count > uint64(len(tc.rows)) {
						return tc.rows, nil
					}
					return tc.rows[:ls.Limit.Count], nil
				})

			lggr, logs := logger.TestObserved(t, zapcore.WarnLevel)
			extCr := contractreader.NewExtendedContractReader(cr, contractreader.WithLogger(lggr))

			seqs, err := extCr.QueryKey(
				t.Context(),
				types.BoundContract{Name: consts.ContractNameOnRamp, Address: "0x123"},
				query.KeyFilter{Key: consts.EventNameCCIPMessageSent},
				query.NewLimitAndSort(query.CountLimit(tc.limit), query.NewSortBySequence(query.Asc)),
				&event{},
			)

			assert.Equal(t, tc.expLimits, limits)
			if tc.expErr != nil {
				require.ErrorIs(t, err, tc.expErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expSeqNums, seqNums(seqs))
			assert.Equal(t, tc.rows[0].Cursor, seqs[0].Cursor, "the first stored row of a log is the one kept")
			assert.Equal(t, tc.expWarn, logs.Len() > 0)
		})
	}
}

func TestExtendedContractReader_QueryKeyDedupPropagatesError(t *testing.T) {
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	cr.EXPECT().HealthReport().Return(nil)
	cr.EXPECT().QueryKey(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, assert.AnError).Once()

	_, err := contractreader.NewExtendedContractReader(cr).QueryKey(
		t.Context(), types.BoundContract{}, query.KeyFilter{}, query.LimitAndSort{}, nil)
	require.ErrorIs(t, err, assert.AnError)
}
