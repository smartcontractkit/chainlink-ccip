package chainaccessor

import (
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
