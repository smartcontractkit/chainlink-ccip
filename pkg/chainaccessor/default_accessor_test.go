package chainaccessor

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	writer_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/types"
	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

var (
	chainA = cciptypes.ChainSelector(1)
	chainB = cciptypes.ChainSelector(2)
)

func TestCCIPChainReader_CreateExecutedMessagesKeyFilter(t *testing.T) {
	var (
		range1 = cciptypes.NewSeqNumRange(1, 2)
		range2 = cciptypes.NewSeqNumRange(5, 7)
		range3 = cciptypes.NewSeqNumRange(10, 15)
	)
	testCases := []struct {
		name               string
		seqNrRangesByChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange
		confidence         primitives.ConfidenceLevel
		expectedCount      uint64
		expected           query.KeyFilter
	}{
		{
			name: "simple example",
			seqNrRangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				chainA: {range1},
			},
			confidence:    primitives.Finalized,
			expectedCount: 2,
			expected: query.KeyFilter{
				Key: consts.EventNameExecutionStateChanged,
				Expressions: []query.Expression{
					{
						BoolExpression: query.BoolExpression{
							BoolOperator: query.AND,
							Expressions: []query.Expression{
								{
									Primitive: &primitives.Comparator{
										Name: consts.EventAttributeSequenceNumber,
										ValueComparators: []primitives.ValueComparator{
											{Value: range1.Start(), Operator: primitives.Gte},
											{Value: range1.End(), Operator: primitives.Lte},
										},
									},
								},
								{
									Primitive: &primitives.Comparator{
										Name: consts.EventAttributeSourceChain,
										ValueComparators: []primitives.ValueComparator{
											{Value: chainA, Operator: primitives.Eq},
										},
									},
								},
							},
						},
					},
					{
						Primitive: &primitives.Comparator{
							Name:             consts.EventAttributeState,
							ValueComparators: []primitives.ValueComparator{{Value: 0, Operator: primitives.Gt}},
						},
					},
					{Primitive: &primitives.Confidence{ConfidenceLevel: primitives.Finalized}},
				},
			},
		},
		{
			name: "multiChain simple example",
			seqNrRangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				chainA: {range1},
				chainB: {range2},
			},
			confidence:    primitives.Finalized,
			expectedCount: 5,
			expected: query.KeyFilter{
				Key: consts.EventNameExecutionStateChanged,
				Expressions: []query.Expression{
					{
						BoolExpression: query.BoolExpression{
							BoolOperator: query.OR,
							Expressions: []query.Expression{
								{
									BoolExpression: query.BoolExpression{
										BoolOperator: query.AND,
										Expressions: []query.Expression{
											{
												Primitive: &primitives.Comparator{
													Name: consts.EventAttributeSequenceNumber,
													ValueComparators: []primitives.ValueComparator{
														{Value: range1.Start(), Operator: primitives.Gte},
														{Value: range1.End(), Operator: primitives.Lte},
													},
												},
											},
											{
												Primitive: &primitives.Comparator{
													Name: consts.EventAttributeSourceChain,
													ValueComparators: []primitives.ValueComparator{
														{Value: chainA, Operator: primitives.Eq},
													},
												},
											},
										},
									},
								},
								{
									BoolExpression: query.BoolExpression{
										BoolOperator: query.AND,
										Expressions: []query.Expression{
											{
												Primitive: &primitives.Comparator{
													Name: consts.EventAttributeSequenceNumber,
													ValueComparators: []primitives.ValueComparator{
														{Value: range2.Start(), Operator: primitives.Gte},
														{Value: range2.End(), Operator: primitives.Lte},
													},
												},
											},
											{
												Primitive: &primitives.Comparator{
													Name: consts.EventAttributeSourceChain,
													ValueComparators: []primitives.ValueComparator{
														{Value: chainB, Operator: primitives.Eq},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Primitive: &primitives.Comparator{
							Name:             consts.EventAttributeState,
							ValueComparators: []primitives.ValueComparator{{Value: 0, Operator: primitives.Gt}},
						},
					},
					{Primitive: &primitives.Confidence{ConfidenceLevel: primitives.Finalized}},
				},
			},
		},
		{
			name: "multichain multi range example",
			seqNrRangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				chainA: {range1, range2, range3},
				chainB: {range2, range3},
			},
			confidence:    primitives.Finalized,
			expectedCount: 20,
			expected: query.KeyFilter{
				Key: consts.EventNameExecutionStateChanged,
				Expressions: []query.Expression{
					{
						BoolExpression: query.BoolExpression{
							BoolOperator: query.OR,
							Expressions: []query.Expression{
								{
									BoolExpression: query.BoolExpression{
										BoolOperator: query.AND,
										Expressions: []query.Expression{
											{
												BoolExpression: query.BoolExpression{
													BoolOperator: query.OR,
													Expressions: []query.Expression{
														{
															Primitive: &primitives.Comparator{
																Name: consts.EventAttributeSequenceNumber,
																ValueComparators: []primitives.ValueComparator{
																	{Value: range1.Start(), Operator: primitives.Gte},
																	{Value: range1.End(), Operator: primitives.Lte},
																},
															},
														},
														{
															Primitive: &primitives.Comparator{
																Name: consts.EventAttributeSequenceNumber,
																ValueComparators: []primitives.ValueComparator{
																	{Value: range2.Start(), Operator: primitives.Gte},
																	{Value: range2.End(), Operator: primitives.Lte},
																},
															},
														},
														{
															Primitive: &primitives.Comparator{
																Name: consts.EventAttributeSequenceNumber,
																ValueComparators: []primitives.ValueComparator{
																	{Value: range3.Start(), Operator: primitives.Gte},
																	{Value: range3.End(), Operator: primitives.Lte},
																},
															},
														},
													},
												},
											},
											{
												Primitive: &primitives.Comparator{
													Name: consts.EventAttributeSourceChain,
													ValueComparators: []primitives.ValueComparator{
														{Value: chainA, Operator: primitives.Eq},
													},
												},
											},
										},
									},
								},
								{
									BoolExpression: query.BoolExpression{
										BoolOperator: query.AND,
										Expressions: []query.Expression{
											{
												BoolExpression: query.BoolExpression{
													BoolOperator: query.OR,
													Expressions: []query.Expression{
														{
															Primitive: &primitives.Comparator{
																Name: consts.EventAttributeSequenceNumber,
																ValueComparators: []primitives.ValueComparator{
																	{Value: range2.Start(), Operator: primitives.Gte},
																	{Value: range2.End(), Operator: primitives.Lte},
																},
															},
														},
														{
															Primitive: &primitives.Comparator{
																Name: consts.EventAttributeSequenceNumber,
																ValueComparators: []primitives.ValueComparator{
																	{Value: range3.Start(), Operator: primitives.Gte},
																	{Value: range3.End(), Operator: primitives.Lte},
																},
															},
														},
													},
												},
											},
											{
												Primitive: &primitives.Comparator{
													Name: consts.EventAttributeSourceChain,
													ValueComparators: []primitives.ValueComparator{
														{Value: chainB, Operator: primitives.Eq},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Primitive: &primitives.Comparator{
							Name:             consts.EventAttributeState,
							ValueComparators: []primitives.ValueComparator{{Value: 0, Operator: primitives.Gt}},
						},
					},
					{Primitive: &primitives.Confidence{ConfidenceLevel: primitives.Finalized}},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			output, count := createExecutedMessagesKeyFilter(tt.seqNrRangesByChain, tt.confidence)
			//assert.ElementsMatch(t, tt.expected, output, "unequal values")
			if !reflect.DeepEqual(tt.expected, output) {
				t.Errorf("createExecutedMessagesKeyFilter() got = %+v, want %+v", output, tt.expected)
			}
			assert.Equal(t, tt.expectedCount, count)
		})
	}
}

func TestDefaultAccessor_GetSourceChainsConfig(t *testing.T) {
	ctx := context.Background()
	destChain := cciptypes.ChainSelector(3)
	sourceChainA := cciptypes.ChainSelector(1)
	sourceChainB := cciptypes.ChainSelector(2)
	sourceChains := []cciptypes.ChainSelector{sourceChainA, sourceChainB}

	// Create mocked extended reader for destination chain
	mockExtendedReader := reader_mocks.NewMockExtended(t)

	// Mock ExtendedBatchGetLatestValues to return source chain configurations
	mockExtendedReader.EXPECT().ExtendedBatchGetLatestValues(
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).RunAndReturn(func(
		ctx context.Context,
		request contractreader.ExtendedBatchGetLatestValuesRequest,
		allowStale bool,
	) (types.BatchGetLatestValuesResult, []string, error) {
		results := make(types.BatchGetLatestValuesResult, 0)
		for contractName, batch := range request {
			var contractResults []types.BatchReadResult
			for _, readReq := range batch {
				res := types.BatchReadResult{
					ReadName: readReq.ReadName,
				}

				// Handle source chain config requests
				if readReq.ReadName == consts.MethodNameGetSourceChainConfig {
					params := readReq.Params.(map[string]any)
					sourceChain := params["sourceChainSelector"].(cciptypes.ChainSelector)
					v := readReq.ReturnVal.(*cciptypes.SourceChainConfig)

					fromString, err := cciptypes.NewBytesFromString(fmt.Sprintf(
						"0x%d000000000000000000000000000000000000000", sourceChain),
					)
					require.NoError(t, err)
					v.OnRamp = cciptypes.UnknownAddress(fromString)
					v.IsEnabled = true
					v.Router = fromString
					res.SetResult(v, nil)
				}
				contractResults = append(contractResults, res)
			}
			contractKey := types.BoundContract{Name: contractName, Address: "0x123"}
			results[contractKey] = contractResults
		}
		return results, nil, nil
	})

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	mockContractWriter := writer_mocks.NewMockContractWriter(t)
	accessor, err := NewDefaultAccessor(
		logger.Test(t),
		destChain,
		mockExtendedReader,
		mockContractWriter,
		mockAddrCodec,
	)
	require.NoError(t, err)

	// Get source chain configs
	_, sourceChainConfigs, err := accessor.GetAllConfigsLegacy(ctx, destChain, sourceChains)
	assert.NoError(t, err)
	assert.Len(t, sourceChainConfigs, 2)

	// Verify source chain A configuration
	assert.Contains(t, sourceChainConfigs, sourceChainA)
	cfgA := sourceChainConfigs[sourceChainA]
	assert.Equal(t, "0x1000000000000000000000000000000000000000", cfgA.OnRamp.String())
	assert.True(t, cfgA.IsEnabled)

	// Verify router address for chain A
	expectedRouterA, err := cciptypes.NewBytesFromString("0x1000000000000000000000000000000000000000")
	require.NoError(t, err)
	assert.Equal(t, []byte(expectedRouterA), cfgA.Router)

	// Verify source chain B configuration
	assert.Contains(t, sourceChainConfigs, sourceChainB)
	cfgB := sourceChainConfigs[sourceChainB]
	assert.Equal(t, "0x2000000000000000000000000000000000000000", cfgB.OnRamp.String())
	assert.True(t, cfgB.IsEnabled)

	// Verify router address for chain B (compare as bytes directly since it's raw bytes)
	expectedRouterB, err := cciptypes.NewBytesFromString("0x2000000000000000000000000000000000000000")
	require.NoError(t, err)
	assert.Equal(t, []byte(expectedRouterB), cfgB.Router)
}
