package chainaccessor

import (
	"context"
	"fmt"
	"math/big"
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

// Helper function to create a valid SendRequestedEvent for testing
func createValidSendRequestedEvent(seqNum cciptypes.SeqNum) *SendRequestedEvent {
	return &SendRequestedEvent{
		DestChainSelector: chainB,
		SequenceNumber:    seqNum,
		Message: cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				SourceChainSelector: chainA,
				DestChainSelector:   chainB,
				SequenceNumber:      seqNum,
				MessageID:           cciptypes.Bytes32{byte(seqNum)},
			},
			Sender:          cciptypes.UnknownAddress("sender"),
			Receiver:        cciptypes.UnknownAddress("receiver"),
			FeeToken:        cciptypes.UnknownAddress("feeToken"),
			FeeTokenAmount:  cciptypes.NewBigInt(big.NewInt(100)),
		},
	}
}

func TestMsgsBetweenSeqNums(t *testing.T) {
	tests := []struct {
		name                 string
		seqNumRange          cciptypes.SeqNumRange
		destChainSelector    cciptypes.ChainSelector
		sequences            []types.Sequence
		expectedError        bool
		expectedMsgCount     int
		validateTxHash       func(t *testing.T, msgs []cciptypes.Message)
	}{
		{
			name:              "TxHash populated from item.TxHash",
			seqNumRange:       cciptypes.NewSeqNumRange(1, 3),
			destChainSelector: chainB,
			sequences: []types.Sequence{
				{
					Cursor: "100-1-0xabc123",
					TxHash: []byte{0xab, 0xcd, 0xef, 0x12, 0x34, 0x56},
					Data:   createValidSendRequestedEvent(1),
				},
				{
					Cursor: "100-2-0xdef456",
					TxHash: []byte{0xde, 0xad, 0xbe, 0xef, 0x78, 0x90},
					Data:   createValidSendRequestedEvent(2),
				},
			},
			expectedError:    false,
			expectedMsgCount: 2,
			validateTxHash: func(t *testing.T, msgs []cciptypes.Message) {
				require.Len(t, msgs, 2)
				// TxHash should be populated from item.TxHash with 0x prefix
				assert.Equal(t, "0xabcdef123456", msgs[0].Header.TxHash)
				assert.Equal(t, "0xdeadbeef7890", msgs[1].Header.TxHash)
			},
		},
		{
			name:              "Mixed: some with item.TxHash, some without",
			seqNumRange:       cciptypes.NewSeqNumRange(1, 3),
			destChainSelector: chainB,
			sequences: []types.Sequence{
				{
					Cursor: "100-1-0xcursor111",
					TxHash: []byte{0x11, 0x22, 0x33}, // Has TxHash
					Data:   createValidSendRequestedEvent(1),
				},
				{
					Cursor: "100-2-0xcursor222",
					TxHash: nil, // No TxHash
					Data:   createValidSendRequestedEvent(2),
				},
				{
					Cursor: "100-3-0xcursor333",
					TxHash: []byte{0x44, 0x55, 0x66}, // Has TxHash
					Data:   createValidSendRequestedEvent(3),
				},
			},
			expectedError:    false,
			expectedMsgCount: 3,
			validateTxHash: func(t *testing.T, msgs []cciptypes.Message) {
				require.Len(t, msgs, 3)
				assert.Equal(t, "0x112233", msgs[0].Header.TxHash) // From item.TxHash
				assert.Equal(t, "", msgs[1].Header.TxHash)          // Empty - no TxHash provided
				assert.Equal(t, "0x445566", msgs[2].Header.TxHash) // From item.TxHash
			},
		},
		{
			name:              "Empty TxHash when item.TxHash is not provided",
			seqNumRange:       cciptypes.NewSeqNumRange(1, 1),
			destChainSelector: chainB,
			sequences: []types.Sequence{
				{
					Cursor: "100-1-0xabc123",
					TxHash: nil, // No TxHash
					Data:   createValidSendRequestedEvent(1),
				},
			},
			expectedError:    false,
			expectedMsgCount: 1,
			validateTxHash: func(t *testing.T, msgs []cciptypes.Message) {
				require.Len(t, msgs, 1)
				// TxHash should be empty when item.TxHash is not provided
				assert.Equal(t, "", msgs[0].Header.TxHash)
			},
		},
		{
			name:              "Filter out invalid sequence number",
			seqNumRange:       cciptypes.NewSeqNumRange(1, 2),
			destChainSelector: chainB,
			sequences: []types.Sequence{
				{
					Cursor: "100-1-0xabc123",
					TxHash: []byte{0xab, 0xcd},
					Data:   createValidSendRequestedEvent(1),
				},
				{
					Cursor: "100-2-0xdef456",
					TxHash: []byte{0xde, 0xef},
					Data:   createValidSendRequestedEvent(10), // Out of range
				},
			},
			expectedError:    false,
			expectedMsgCount: 1,
			validateTxHash: func(t *testing.T, msgs []cciptypes.Message) {
				require.Len(t, msgs, 1)
				assert.Equal(t, "0xabcd", msgs[0].Header.TxHash)
			},
		},
		{
			name:              "Wrong data type in sequence",
			seqNumRange:       cciptypes.NewSeqNumRange(1, 2),
			destChainSelector: chainB,
			sequences: []types.Sequence{
				{
					Cursor: "100-1-0xabc123",
					TxHash: []byte{0xab, 0xcd},
					Data:   "invalid data type", // Not SendRequestedEvent
				},
			},
			expectedError:    true,
			expectedMsgCount: 0,
			validateTxHash:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockReader := reader_mocks.NewMockExtended(t)
			mockWriter := writer_mocks.NewMockContractWriter(t)
			codec := internal.NewMockAddressCodecHex(t)

			accessor := &DefaultAccessor{
				lggr:           logger.Test(t),
				chainSelector:  chainA,
				contractReader: mockReader,
				contractWriter: mockWriter,
				addrCodec:      codec,
			}

			// Setup mock expectations
			mockReader.On("ExtendedQueryKey", mock.Anything, mock.Anything,
				mock.Anything, mock.Anything, mock.Anything).
				Return(tt.sequences, nil).Once()

			// Setup GetBindings mock for GetContractAddress call
			// The address string will be converted to bytes by the codec
			onRampAddressStr := "0x1234567890abcdef"
			onRampAddress := cciptypes.UnknownAddress{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}
			bindings := []contractreader.ExtendedBoundContract{
				{
					Binding: types.BoundContract{
						Name:    consts.ContractNameOnRamp,
						Address: onRampAddressStr,
					},
				},
			}
			mockReader.On("GetBindings", consts.ContractNameOnRamp).
				Return(bindings).Once()

			// Execute test
			msgs, err := accessor.MsgsBetweenSeqNums(context.Background(), tt.destChainSelector, tt.seqNumRange)

			// Verify results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, msgs, tt.expectedMsgCount)
				if tt.validateTxHash != nil {
					tt.validateTxHash(t, msgs)
				}
				// Verify OnRamp is always set
				for _, msg := range msgs {
					assert.Equal(t, onRampAddress, msg.Header.OnRamp)
				}
			}

			// Verify all mock expectations were met
			mockReader.AssertExpectations(t)
		})
	}
}
