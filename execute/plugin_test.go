package execute

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/internal/reader/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// makeMessage creates a message deterministically derived from the given inputs.
func makeMessage(src cciptypes.ChainSelector, num cciptypes.SeqNum, nonce uint64) cciptypes.Message {
	var placeholderID cciptypes.Bytes32
	n, err := rand.New(rand.NewSource(int64(src) * int64(num) * int64(nonce))).Read(placeholderID[:])
	if n != 32 {
		panic(fmt.Sprintf("Unexpected number of bytes read for placeholder id: want 32, got %d", n))
	}
	if err != nil {
		panic(fmt.Sprintf("Error reading random bytes: %v", err))
	}

	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:           placeholderID,
			SourceChainSelector: src,
			SequenceNumber:      num,
			MsgHash:             cciptypes.Bytes32{},
			Nonce:               nonce,
		},
	}
}

// mustParseByteStr parses a given string into a byte array, any error causes a panic. Pass in an empty string for a
// random byte array.
// nolint:unparam // surly this will be useful at some point...
func mustParseByteStr(byteStr string) cciptypes.Bytes32 {
	if byteStr == "" {
		var randomBytes cciptypes.Bytes32
		n, err := rand.New(rand.NewSource(0)).Read(randomBytes[:])
		if n != 32 {
			panic(fmt.Sprintf("Unexpected number of bytes read for placeholder id: want 32, got %d", n))
		}
		if err != nil {
			panic(fmt.Sprintf("Error reading random bytes: %v", err))
		}
		return randomBytes
	}
	b, err := cciptypes.NewBytes32FromString(byteStr)
	if err != nil {
		panic(err)
	}
	return b
}

func Test_getPendingExecutedReports(t *testing.T) {
	tests := []struct {
		name    string
		reports []plugintypes.CommitPluginReportWithMeta
		ranges  map[cciptypes.ChainSelector][]cciptypes.SeqNumRange
		want    plugintypes.ExecutePluginCommitObservations
		want1   time.Time
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			reports: nil,
			ranges:  nil,
			want:    plugintypes.ExecutePluginCommitObservations{},
			want1:   time.Time{},
			wantErr: assert.NoError,
		},
		{
			name: "single non-executed report",
			reports: []plugintypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				1: nil,
			},
			want: plugintypes.ExecutePluginCommitObservations{
				1: []plugintypes.ExecutePluginCommitDataWithMessages{
					{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						ExecutedMessages:    nil,
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
					}},
				},
			},
			want1:   time.UnixMilli(10101010101),
			wantErr: assert.NoError,
		},
		{
			name: "single half-executed report",
			reports: []plugintypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				1: {
					cciptypes.NewSeqNumRange(1, 3),
					cciptypes.NewSeqNumRange(7, 8),
				},
			},
			want: plugintypes.ExecutePluginCommitObservations{
				1: []plugintypes.ExecutePluginCommitDataWithMessages{
					{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 7, 8},
					}},
				},
			},
			want1:   time.UnixMilli(10101010101),
			wantErr: assert.NoError,
		},
		{
			name: "last timestamp",
			reports: []plugintypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report:    cciptypes.CommitPluginReport{},
				},
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(9999999999999999),
					Report:    cciptypes.CommitPluginReport{},
				},
			},
			ranges:  map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{},
			want:    plugintypes.ExecutePluginCommitObservations{},
			want1:   time.UnixMilli(9999999999999999),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReader := mocks.NewCCIPReader()
			mockReader.On(
				"CommitReportsGTETimestamp", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			).Return(tt.reports, nil)
			for k, v := range tt.ranges {
				mockReader.On("ExecutedMessageRanges", mock.Anything, k, mock.Anything, mock.Anything).Return(v, nil)
			}

			// CCIP Reader mocks:
			// once:
			//      CommitReportsGTETimestamp(ctx, dest, ts, 1000) -> ([]cciptypes.CommitPluginReportWithMeta, error)
			// for each chain selector:
			//      ExecutedMessageRanges(ctx, selector, dest, seqRange) -> ([]cciptypes.SeqNumRange, error)

			got, got1, err := getPendingExecutedReports(context.Background(), mockReader, 123, time.Now())
			if !tt.wantErr(t, err, "getPendingExecutedReports(...)") {
				return
			}
			assert.Equalf(t, tt.want, got, "getPendingExecutedReports(...)")
			assert.Equalf(t, tt.want1, got1, "getPendingExecutedReports(...)")
		})
	}
}

// TODO: better than this
type tdr struct{}

func (t tdr) ReadTokenData(
	ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error,
) {
	return nil, nil
}

// breakCommitReport by adding an extra message. This causes the report to have an unexpected number of messages.
func breakCommitReport(
	commitReport plugintypes.ExecutePluginCommitDataWithMessages,
) plugintypes.ExecutePluginCommitDataWithMessages {
	commitReport.Messages = append(commitReport.Messages, cciptypes.Message{})
	return commitReport
}

// makeTestCommitReport creates a basic commit report with messages given different parameters. This function
// will panic if the input parameters are inconsistent.
func makeTestCommitReport(
	hasher cciptypes.MessageHasher,
	numMessages,
	srcChain,
	firstSeqNum,
	block int,
	timestamp int64,
	rootOverride cciptypes.Bytes32,
	executed []cciptypes.SeqNum,
) plugintypes.ExecutePluginCommitDataWithMessages {
	sequenceNumberRange :=
		cciptypes.NewSeqNumRange(cciptypes.SeqNum(firstSeqNum), cciptypes.SeqNum(firstSeqNum+numMessages-1))

	for _, e := range executed {
		if !sequenceNumberRange.Contains(e) {
			panic("executed message out of range")
		}
	}
	var messages []cciptypes.Message
	for i := 0; i < numMessages; i++ {
		messages = append(messages, makeMessage(
			cciptypes.ChainSelector(srcChain),
			cciptypes.SeqNum(i+firstSeqNum),
			uint64(i)))
	}

	report := plugintypes.ExecutePluginCommitDataWithMessages{
		ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
			//MerkleRoot:          root,
			SourceChain:         cciptypes.ChainSelector(srcChain),
			SequenceNumberRange: sequenceNumberRange,
			Timestamp:           time.UnixMilli(timestamp),
			BlockNum:            uint64(block),
			ExecutedMessages:    executed,
		},
		Messages: messages,
	}

	// calculate merkle root
	root := rootOverride
	if root.IsEmpty() {
		tree, err := constructMerkleTree(context.Background(), hasher, report)
		if err != nil {
			panic(fmt.Sprintf("unable to construct merkle tree: %s", err))
		}
		report.MerkleRoot = tree.Root()
	}

	return report
}

// assertMerkleRoot computes the source messages merkle root, then computes a verification with the proof, then compares
// the roots.
func assertMerkleRoot(
	t *testing.T,
	hasher cciptypes.MessageHasher,
	execReport cciptypes.ExecutePluginReportSingleChain,
	commitReport plugintypes.ExecutePluginCommitDataWithMessages,
) {
	keccak := hashutil.NewKeccak()
	// Generate merkle root from commit report messages
	var leafHashes [][32]byte
	for _, msg := range commitReport.Messages {
		hash, err := hasher.Hash(context.Background(), msg)
		require.NoError(t, err)
		leafHashes = append(leafHashes, hash)
	}
	tree, err := merklemulti.NewTree(keccak, leafHashes)
	require.NoError(t, err)
	merkleRoot := tree.Root()

	// Generate merkle root from exec report messages and proof
	ctx := context.Background()
	var leaves [][32]byte
	for _, msg := range execReport.Messages {
		hash, err := hasher.Hash(ctx, msg)
		require.NoError(t, err)
		leaves = append(leaves, hash)
	}
	proofCast := make([][32]byte, len(execReport.Proofs))
	for i, p := range execReport.Proofs {
		copy(proofCast[i][:], p[:32])
	}
	var proof merklemulti.Proof[[32]byte]
	proof.Hashes = proofCast
	proof.SourceFlags = slicelib.BitFlagsToBools(execReport.ProofFlagBits.Int, len(leaves)+len(proofCast)-1)
	recomputedMerkleRoot, err := merklemulti.VerifyComputeRoot(hashutil.NewKeccak(),
		leaves,
		proof)
	assert.NoError(t, err)
	assert.NotNil(t, recomputedMerkleRoot)

	// Compare them
	assert.Equal(t, merkleRoot, recomputedMerkleRoot)
}

func Test_selectReport(t *testing.T) {
	hasher := mocks.NewMessageHasher()
	codec := mocks.NewExecutePluginJSONReportCodec()
	lggr := logger.Test(t)
	var tokenDataReader tdr

	type args struct {
		reports       []plugintypes.ExecutePluginCommitDataWithMessages
		maxReportSize int
	}
	tests := []struct {
		name                  string
		args                  args
		expectedExecReports   int
		expectedCommitReports int
		expectedExecThings    []int
		lastReportExecuted    []cciptypes.SeqNum
		wantErr               string
	}{
		{
			name: "empty report",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{},
			},
			expectedExecReports:   0,
			expectedCommitReports: 0,
		},
		{
			name: "half report",
			args: args{
				maxReportSize: 2300,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{5},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104},
		},
		{
			name: "full report",
			args: args{
				maxReportSize: 10000,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{10},
		},
		{
			name: "two reports",
			args: args{
				maxReportSize: 15000,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
				},
			},
			expectedExecReports:   2,
			expectedCommitReports: 0,
			expectedExecThings:    []int{10, 20},
		},
		{
			name: "one and half reports",
			args: args{
				maxReportSize: 8500,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
				},
			},
			expectedExecReports:   2,
			expectedCommitReports: 1,
			expectedExecThings:    []int{10, 10},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		},
		{
			name: "exactly one report",
			args: args{
				maxReportSize: 4200,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						nil),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{10},
			lastReportExecuted:    []cciptypes.SeqNum{},
		},
		{
			name: "execute remainder of partially executed report",
			args: args{
				maxReportSize: 2500,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						[]cciptypes.SeqNum{100, 101, 102, 103, 104}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{5},
		},
		{
			name: "partially execute remainder of partially executed report",
			args: args{
				maxReportSize: 2050,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						[]cciptypes.SeqNum{100, 101, 102, 103, 104}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{4},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 105, 106, 107, 108},
		},
		{
			name: "execute remainder of sparsely executed report",
			args: args{
				maxReportSize: 3500,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						[]cciptypes.SeqNum{100, 102, 104, 106, 108}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{5},
		},
		{
			name: "partially execute remainder of partially executed sparse report",
			args: args{
				maxReportSize: 2050,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						cciptypes.Bytes32{}, // generate a correct root.
						[]cciptypes.SeqNum{100, 102, 104, 106, 108}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{4},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 105, 106, 107, 108},
		},
		{
			name: "broken report",
			args: args{
				maxReportSize: 10000,
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					breakCommitReport(makeTestCommitReport(hasher, 10, 1, 101, 1000, 10101010102,
						cciptypes.Bytes32{}, // generate a correct root.
						nil)),
				},
			},
			wantErr: "unable to build a single chain report",
		},
		{
			name: "invalid merkle root",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						mustParseByteStr(""), // random root
						nil),
				},
			},
			wantErr: "merkle root mismatch: expected 0x00000000000000000",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			execReports, commitReports, err :=
				selectReport(ctx, lggr, hasher, codec, tokenDataReader, tt.args.reports, tt.args.maxReportSize)
			if tt.wantErr != "" {
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Len(t, execReports, tt.expectedExecReports)
			require.Len(t, commitReports, tt.expectedCommitReports)
			for i, execReport := range execReports {
				require.Lenf(t, execReport.Messages, tt.expectedExecThings[i],
					"Unexpected number of messages, iter %d", i)
				require.Lenf(t, execReport.OffchainTokenData, tt.expectedExecThings[i],
					"Unexpected number of token data, iter %d", i)
				require.NotEmptyf(t, execReport.Proofs, "Proof should not be empty.")
				assertMerkleRoot(t, hasher, execReport, tt.args.reports[i])
			}
			// If the last report is partially executed, the executed messages can be checked.
			if len(execReports) > 0 && len(tt.lastReportExecuted) > 0 {
				lastReport := commitReports[len(commitReports)-1]
				require.ElementsMatch(t, tt.lastReportExecuted, lastReport.ExecutedMessages)
			}
		})
	}
}

type badHasher struct{}

func (bh badHasher) Hash(context.Context, cciptypes.Message) (cciptypes.Bytes32, error) {
	return cciptypes.Bytes32{}, fmt.Errorf("bad hasher")
}

type badTokenDataReader struct{}

func (btdr badTokenDataReader) ReadTokenData(
	_ context.Context, _ cciptypes.ChainSelector, _ cciptypes.SeqNum,
) ([][]byte, error) {
	return nil, fmt.Errorf("bad token data reader")
}

type badCodec struct{}

func (bc badCodec) Encode(ctx context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	return nil, fmt.Errorf("bad codec")
}

func (bc badCodec) Decode(ctx context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	return cciptypes.ExecutePluginReport{}, fmt.Errorf("bad codec")
}

func Test_buildSingleChainReport_Errors(t *testing.T) {
	lggr := logger.Test(t)

	type args struct {
		report          plugintypes.ExecutePluginCommitDataWithMessages
		maxMessages     int
		hasher          cciptypes.MessageHasher
		tokenDataReader TokenDataReader
		codec           cciptypes.ExecutePluginCodec
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name:    "wrong number of messages",
			wantErr: "unexpected number of messages: expected 1, got 2",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{}},
						{Header: cciptypes.RampMessageHeader{}},
					},
				},
			},
		},
		{
			name:    "wrong sequence numbers",
			wantErr: "sequence number 102 outside of report range [100 -> 101]",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(101)),
					},
					Messages: []cciptypes.Message{
						{
							Header: cciptypes.RampMessageHeader{
								SequenceNumber: cciptypes.SeqNum(100),
							},
						},
						{
							Header: cciptypes.RampMessageHeader{
								SequenceNumber: cciptypes.SeqNum(102),
							},
						},
					},
				},
			},
		},
		{
			name:    "source mismatch",
			wantErr: "unexpected source chain: expected 1111, got 2222",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1111,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 2222,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				hasher: badHasher{},
			},
		},
		{
			name:    "bad hasher",
			wantErr: "unable to hash message (1234567, 100): bad hasher",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1234567,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 1234567,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				hasher: badHasher{},
			},
		},
		{
			name:    "bad token data reader",
			wantErr: "unable to read token data for message 100: bad token data reader",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1234567,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 1234567,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				tokenDataReader: badTokenDataReader{},
			},
		},
		{
			name:    "bad codec",
			wantErr: "unable to encode report: bad codec",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1234567,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 1234567,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				codec: badCodec{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Select hasher mock.
			var resolvedHasher cciptypes.MessageHasher
			if tt.args.hasher != nil {
				resolvedHasher = tt.args.hasher
			} else {
				resolvedHasher = mocks.NewMessageHasher()
			}

			// Select token data reader mock.
			var resolvedTokenDataReader TokenDataReader
			if tt.args.tokenDataReader != nil {
				resolvedTokenDataReader = tt.args.tokenDataReader
			} else {
				resolvedTokenDataReader = tdr{}
			}

			// Select codec mock.
			var resolvedCodec cciptypes.ExecutePluginCodec
			if tt.args.codec != nil {
				resolvedCodec = tt.args.codec
			} else {
				resolvedCodec = mocks.NewExecutePluginJSONReportCodec()
			}

			ctx := context.Background()
			msgs := make(map[int]struct{})
			for i := 0; i < len(tt.args.report.Messages); i++ {
				msgs[i] = struct{}{}
			}
			execReport, size, _, err := buildSingleChainReport(
				ctx, lggr, resolvedHasher, resolvedTokenDataReader, resolvedCodec, tt.args.report, msgs)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			fmt.Println(execReport, size, err)
		})
	}
}

func TestPlugin_Close(t *testing.T) {
	p := &Plugin{}
	require.NoError(t, p.Close())
}

func TestPlugin_Query(t *testing.T) {
	p := &Plugin{}
	q, err := p.Query(context.Background(), ocr3types.OutcomeContext{})
	require.NoError(t, err)
	require.Equal(t, types.Query{}, q)
}

func TestPlugin_ObservationQuorum(t *testing.T) {
	p := &Plugin{}
	got, err := p.ObservationQuorum(ocr3types.OutcomeContext{}, nil)
	require.NoError(t, err)
	assert.Equal(t, ocr3types.QuorumFPlusOne, got)
}

func TestPlugin_ValidateObservation_NonDecodable(t *testing.T) {
	p := &Plugin{}
	err := p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte("not a valid observation"),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observation")
}

func TestPlugin_ValidateObservation_SupportedChainsError(t *testing.T) {
	p := &Plugin{}
	err := p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte(`{"oracleID": "0xdeadbeef"}`),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error finding supported chains by node: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_ValidateObservation_IneligibleObserver(t *testing.T) {
	lggr := logger.Test(t)

	p := &Plugin{
		homeChain: setupHomeChainPoller(lggr, []reader.ChainConfigInfo{
			{
				ChainSelector: 0,
				ChainConfig:   reader.HomeChainConfigMapper{},
			},
		}),
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
	}

	observation := plugintypes.NewExecutePluginObservation(nil, plugintypes.ExecutePluginMessageObservations{
		0: map[cciptypes.SeqNum]cciptypes.Message{
			1: {
				Header: cciptypes.RampMessageHeader{
					SourceChainSelector: 1,
				},
			},
		},
	})
	encoded, err := observation.Encode()
	require.NoError(t, err)
	err = p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observer reading eligibility: observer not allowed to read from chain 0")
}

func TestPlugin_ValidateObservation_ValidateObservedSeqNum_Error(t *testing.T) {
	lggr := logger.Test(t)

	p := &Plugin{
		homeChain: setupHomeChainPoller(lggr, []reader.ChainConfigInfo{
			{
				ChainSelector: 1,
				ChainConfig:   reader.HomeChainConfigMapper{},
			},
		}),
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
	}

	// Reports with duplicate roots.
	root := mustParseByteStr("")
	commitReports := map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
		1: {
			{
				ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
					MerkleRoot: root,
				},
			},
			{
				ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
					MerkleRoot: root,
				},
			},
		},
	}
	observation := plugintypes.NewExecutePluginObservation(commitReports, nil)
	encoded, err := observation.Encode()
	require.NoError(t, err)
	err = p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observed sequence numbers: duplicate merkle root")
}

func TestPlugin_Observation_BadPreviousOutcome(t *testing.T) {
	p := &Plugin{}
	_, err := p.Observation(context.Background(), ocr3types.OutcomeContext{
		PreviousOutcome: []byte("not a valid observation"),
	}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode previous outcome: invalid character")
}

func TestPlugin_Observation_EligibilityCheckFailure(t *testing.T) {
	lggr := logger.Test(t)
	p := &Plugin{
		homeChain:       setupHomeChainPoller(lggr, []reader.ChainConfigInfo{}),
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{},
	}

	_, err := p.Observation(context.Background(), ocr3types.OutcomeContext{}, nil)
	require.Error(t, err)
	// nolint:lll // error message
	assert.Contains(t, err.Error(), "unable to determine if the destination chain is supported: error getting supported chains: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_Outcome_BadObservationEncoding(t *testing.T) {
	p := &Plugin{}
	_, err := p.Outcome(ocr3types.OutcomeContext{}, nil,
		[]types.AttributedObservation{
			{
				Observation: []byte("not a valid observation"),
				Observer:    0,
			},
		})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observations: invalid character")
}

func TestPlugin_Outcome_BelowF(t *testing.T) {
	p := &Plugin{
		reportingCfg: ocr3types.ReportingPluginConfig{
			F: 1,
		},
	}
	_, err := p.Outcome(ocr3types.OutcomeContext{}, nil,
		[]types.AttributedObservation{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "below F threshold")
}

func TestPlugin_Outcome_HomeChainError(t *testing.T) {
	homeChain := reader_mock.NewHomeChain(t)
	homeChain.On("GetFChain", mock.Anything).Return(nil, fmt.Errorf("test error"))

	p := &Plugin{
		homeChain: homeChain,
	}
	_, err := p.Outcome(ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to get FChain: test error")
}

func TestPlugin_Outcome_CommitReportsMergeError(t *testing.T) {
	homeChain := reader_mock.NewHomeChain(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		10: 20,
	}
	homeChain.On("GetFChain", mock.Anything).Return(fChainMap, nil)

	p := &Plugin{
		homeChain: homeChain,
	}

	commitReports := map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
		1: {
			{
				ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
					MerkleRoot: mustParseByteStr(""),
				},
			},
			{
				ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
					MerkleRoot: mustParseByteStr(""),
				},
			},
		},
	}
	observation, err := plugintypes.NewExecutePluginObservation(commitReports, nil).Encode()
	require.NoError(t, err)
	_, err = p.Outcome(ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to merge commit report observations: no validator")
}

func TestPlugin_Outcome_MessagesMergeError(t *testing.T) {
	homeChain := reader_mock.NewHomeChain(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		10: 20,
	}
	homeChain.On("GetFChain", mock.Anything).Return(fChainMap, nil)

	p := &Plugin{
		homeChain: homeChain,
	}

	//map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message
	messages := map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
		1: {
			1: {
				Header: cciptypes.RampMessageHeader{
					SourceChainSelector: 1,
				},
			},
		},
	}
	observation, err := plugintypes.NewExecutePluginObservation(nil, messages).Encode()
	require.NoError(t, err)
	_, err = p.Outcome(ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to merge message observations: no validator")
}

func TestPlugin_Reports_UnableToParse(t *testing.T) {
	p := &Plugin{}
	_, err := p.Reports(0, ocr3types.Outcome("not a valid observation"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode outcome")
}

func TestPlugin_Reports_UnableToEncode(t *testing.T) {
	codec := mocks.NewExecutePluginCodec(t)
	codec.On("Encode", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("test error"))
	p := &Plugin{reportCodec: codec}
	report, err := plugintypes.NewExecutePluginOutcome(nil, cciptypes.ExecutePluginReport{}).Encode()
	require.NoError(t, err)

	_, err = p.Reports(0, report)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to encode report: test error")
}

func TestPlugin_ShouldAcceptAttestedReport_DoesNotDecode(t *testing.T) {
	codec := mocks.NewExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, fmt.Errorf("test error"))
	p := &Plugin{
		reportCodec: codec,
	}
	_, err := p.ShouldAcceptAttestedReport(context.Background(), 0, ocr3types.ReportWithInfo[[]byte]{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "decode commit plugin report: test error")
}

func TestPlugin_ShouldAcceptAttestedReport_NoReports(t *testing.T) {
	codec := mocks.NewExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, nil)
	p := &Plugin{
		lggr:        logger.Test(t),
		reportCodec: codec,
	}
	result, err := p.ShouldAcceptAttestedReport(context.Background(), 0, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.False(t, result)
}

func TestPlugin_ShouldAcceptAttestedReport_ShouldAccept(t *testing.T) {
	codec := mocks.NewExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{},
			},
		}, nil)
	p := &Plugin{
		lggr:        logger.Test(t),
		reportCodec: codec,
	}
	result, err := p.ShouldAcceptAttestedReport(context.Background(), 0, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.True(t, result)
}

func TestPlugin_ShouldTransmitAcceptReport_ElegibilityCheckFailure(t *testing.T) {
	lggr := logger.Test(t)
	p := &Plugin{
		homeChain:       setupHomeChainPoller(lggr, []reader.ChainConfigInfo{}),
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{},
	}

	_, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.Error(t, err)
	// nolint:lll // error message
	assert.Contains(t, err.Error(), "unable to determine if the destination chain is supported: error getting supported chains: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_ShouldTransmitAcceptReport_Ineligible(t *testing.T) {
	lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)
	p := &Plugin{
		lggr:         lggr,
		cfg:          pluginconfig.ExecutePluginConfig{DestChain: 1},
		reportingCfg: ocr3types.ReportingPluginConfig{OracleID: 2},
		homeChain:    setupHomeChainPoller(lggr, []reader.ChainConfigInfo{}),
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			2: {},
		},
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.False(t, shouldTransmit)

	messages := slicelib.Map(logs.All(), func(e observer.LoggedEntry) string {
		return e.Message
	})
	require.ElementsMatch(t, messages, []string{"not a destination writer, skipping report transmission"})
}

func TestPlugin_ShouldTransmitAcceptReport_DecodeFailure(t *testing.T) {
	homeChain := reader_mock.NewHomeChain(t)
	homeChain.On("GetSupportedChainsForPeer", mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)
	codec := mocks.NewExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, fmt.Errorf("test error"))

	p := &Plugin{
		lggr:         logger.Test(t),
		cfg:          pluginconfig.ExecutePluginConfig{DestChain: 1},
		reportingCfg: ocr3types.ReportingPluginConfig{OracleID: 2},
		reportCodec:  codec,
		homeChain:    homeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			2: {1},
		},
	}

	_, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.Error(t, err)
	require.ErrorContains(t, err, "decode commit plugin report: test error")
}

func TestPlugin_ShouldTransmitAcceptReport_Success(t *testing.T) {
	lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)
	homeChain := reader_mock.NewHomeChain(t)
	homeChain.On("GetSupportedChainsForPeer", mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)
	codec := mocks.NewExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, nil)

	p := &Plugin{
		lggr:         lggr,
		cfg:          pluginconfig.ExecutePluginConfig{DestChain: 1},
		reportingCfg: ocr3types.ReportingPluginConfig{OracleID: 2},
		reportCodec:  codec,
		homeChain:    homeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			2: {1},
		},
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.True(t, shouldTransmit)

	messages := slicelib.Map(logs.All(), func(e observer.LoggedEntry) string {
		return e.Message
	})
	require.ElementsMatch(t, messages, []string{"transmitting report"})
}
