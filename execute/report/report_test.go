package report

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// mustMakeBytes parses a given string into a byte array, any error causes a panic. Pass in an empty string for a
// random byte array.
// nolint:unparam // surly this will be useful at some point...
func mustMakeBytes(byteStr string) cciptypes.Bytes32 {
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

func TestMustMakeBytes(t *testing.T) {
	type args struct {
		byteStr string
	}
	tests := []struct {
		name      string
		args      args
		want      cciptypes.Bytes32
		willPanic bool
	}{
		{
			name: "empty - deterministic random",
			args: args{byteStr: ""},
			want: cciptypes.Bytes32{
				0x01, 0x94, 0xfd, 0xc2, 0xfa, 0x2f, 0xfc, 0xc0, 0x41, 0xd3, 0xff, 0x12, 0x4, 0x5b, 0x73, 0xc8,
				0x6e, 0x4f, 0xf9, 0x5f, 0xf6, 0x62, 0xa5, 0xee, 0xe8, 0x2a, 0xbd, 0xf4, 0x4a, 0x2d, 0xb, 0x75,
			},
		},
		{
			name: "constant",
			args: args{byteStr: "0x0000000000000000000000000000000000000000000000000000000000000000"},
			want: cciptypes.Bytes32{},
		},
		{
			name: "constant2",
			args: args{byteStr: "0x0102030405060708090102030405060708090102030405060708090102030405"},
			want: cciptypes.Bytes32{
				0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
				0x8, 0x9, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x1, 0x2, 0x3, 0x4, 0x5,
			},
		},
		{
			name:      "panic - wrong size",
			args:      args{byteStr: "0x00000"},
			willPanic: true,
		},
		{
			name:      "panic - wrong format",
			args:      args{byteStr: "lorem ipsum"},
			willPanic: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.willPanic {
				assert.Panics(t, func() {
					mustMakeBytes(tt.args.byteStr)
				})
				return
			}

			got := mustMakeBytes(tt.args.byteStr)
			assert.Equal(t, tt.want, got, "mustMakeBytes() = %v, want %v", got, tt.want)
			fmt.Println(got)
		})
	}
}

// assertMerkleRoot computes the source messages merkle root, then computes a verification with the proof, then compares
// the roots.
func assertMerkleRoot(
	t *testing.T,
	hasher cciptypes.MessageHasher,
	execReport cciptypes.ExecutePluginReportSingleChain,
	commitReport plugintypes.ExecutePluginCommitData,
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
		},
	}
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
) plugintypes.ExecutePluginCommitData {
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

	commitReport := plugintypes.ExecutePluginCommitData{
		//MerkleRoot:          root,
		SourceChain:         cciptypes.ChainSelector(srcChain),
		SequenceNumberRange: sequenceNumberRange,
		Timestamp:           time.UnixMilli(timestamp),
		BlockNum:            uint64(block),
		Messages:            messages,
		ExecutedMessages:    executed,
	}

	// calculate merkle root
	root := rootOverride
	if root.IsEmpty() {
		tree, err := ConstructMerkleTree(context.Background(), hasher, commitReport, logger.Nop())
		if err != nil {
			panic(fmt.Sprintf("unable to construct merkle tree: %s", err))
		}
		commitReport.MerkleRoot = tree.Root()
	}

	return commitReport
}

// breakCommitReport by adding an extra message. This causes the report to have an unexpected number of messages.
func breakCommitReport(
	commitReport plugintypes.ExecutePluginCommitData,
) plugintypes.ExecutePluginCommitData {
	commitReport.Messages = append(commitReport.Messages, cciptypes.Message{})
	return commitReport
}

// setMessageData at the given index to the given size. This function will panic if the index is out of range.
func setMessageData(
	idx int, size uint64, commitReport plugintypes.ExecutePluginCommitData,
) plugintypes.ExecutePluginCommitData {
	if len(commitReport.Messages) < idx {
		panic("message index out of range")
	}
	commitReport.Messages[idx].Data = make([]byte, size)
	return commitReport
}

type badHasher struct{}

func (bh badHasher) Hash(context.Context, cciptypes.Message) (cciptypes.Bytes32, error) {
	return cciptypes.Bytes32{}, fmt.Errorf("bad hasher")
}

func Test_buildSingleChainReport_Errors(t *testing.T) {
	lggr := logger.Test(t)

	type args struct {
		report plugintypes.ExecutePluginCommitData
		hasher cciptypes.MessageHasher
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name:    "token data mismatch",
			wantErr: "token data length mismatch: got 2, expected 0",
			args: args{
				report: plugintypes.ExecutePluginCommitData{
					TokenData: make([][][]byte, 2),
				},
			},
		},
		{
			name:    "wrong number of messages",
			wantErr: "unexpected number of messages: expected 1, got 2",
			args: args{
				report: plugintypes.ExecutePluginCommitData{
					TokenData:           make([][][]byte, 2),
					SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
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
				report: plugintypes.ExecutePluginCommitData{
					TokenData:           make([][][]byte, 2),
					SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(101)),
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
				report: plugintypes.ExecutePluginCommitData{
					TokenData:           make([][][]byte, 1),
					SourceChain:         1111,
					SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					Messages: []cciptypes.Message{
						{
							Header: cciptypes.RampMessageHeader{
								SourceChainSelector: 2222,
								SequenceNumber:      cciptypes.SeqNum(100),
							},
						},
					},
				},
				hasher: badHasher{},
			},
		},
		{
			name:    "bad hasher",
			wantErr: "unable to hash message (1234567, 100): bad hasher",
			args: args{
				report: plugintypes.ExecutePluginCommitData{
					TokenData:           make([][][]byte, 1),
					SourceChain:         1234567,
					SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					Messages: []cciptypes.Message{
						{
							Header: cciptypes.RampMessageHeader{
								SourceChainSelector: 1234567,
								SequenceNumber:      cciptypes.SeqNum(100),
							},
						},
					},
				},
				hasher: badHasher{},
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

			ctx := context.Background()
			msgs := make(map[int]struct{})
			for i := 0; i < len(tt.args.report.Messages); i++ {
				msgs[i] = struct{}{}
			}
			_, err := buildSingleChainReportHelper(ctx, lggr, resolvedHasher, tt.args.report, msgs)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func Test_Builder_Build(t *testing.T) {
	hasher := mocks.NewMessageHasher()
	codec := mocks.NewExecutePluginJSONReportCodec()
	lggr := logger.Test(t)
	tokenDataReader := tdr{mode: good}

	type args struct {
		reports       []plugintypes.ExecutePluginCommitData
		maxReportSize uint64
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
				reports: []plugintypes.ExecutePluginCommitData{},
			},
			expectedExecReports:   0,
			expectedCommitReports: 0,
		},
		{
			name: "half report",
			args: args{
				maxReportSize: 2300,
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
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
				reports: []plugintypes.ExecutePluginCommitData{
					breakCommitReport(makeTestCommitReport(hasher, 10, 1, 101, 1000, 10101010102,
						cciptypes.Bytes32{}, // generate a correct root.
						nil)),
				},
			},
			wantErr: "unable to add a single chain report",
		},
		{
			name: "invalid merkle root",
			args: args{
				reports: []plugintypes.ExecutePluginCommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						mustMakeBytes(""), // random root
						nil),
				},
			},
			wantErr: "merkle root mismatch: expected 0x00000000000000000",
		},
		{
			name: "skip over one large messages",
			args: args{
				maxReportSize: 10000,
				reports: []plugintypes.ExecutePluginCommitData{
					setMessageData(5, 20000,
						makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
							cciptypes.Bytes32{}, // generate a correct root.
							nil)),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{9},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 106, 107, 108, 109},
		},
		{
			name: "skip over two large messages",
			args: args{
				maxReportSize: 10000,
				reports: []plugintypes.ExecutePluginCommitData{
					setMessageData(8, 20000,
						setMessageData(5, 20000,
							makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
								cciptypes.Bytes32{}, // generate a correct root.
								nil))),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{8},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 106, 107, 109},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			// look for error in Add or Build
			foundError := false

			builder := NewBuilder(ctx, lggr, hasher, tokenDataReader, codec, tt.args.maxReportSize, 0)
			var updatedMessages []plugintypes.ExecutePluginCommitData
			for _, report := range tt.args.reports {
				updatedMessage, err := builder.Add(report)
				if err != nil && tt.wantErr != "" {
					if strings.Contains(err.Error(), tt.wantErr) {
						foundError = true
						break
					}
				}
				require.NoError(t, err)
				updatedMessages = append(updatedMessages, updatedMessage)
			}
			if foundError {
				return
			}
			execReports, err := builder.Build()
			if tt.wantErr != "" {
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Len(t, execReports, tt.expectedExecReports)
			//require.Len(t, commitReports, tt.expectedCommitReports)
			for i, execReport := range execReports {
				require.Lenf(t, execReport.Messages, tt.expectedExecThings[i],
					"Unexpected number of messages, iter %d", i)
				require.Lenf(t, execReport.OffchainTokenData, tt.expectedExecThings[i],
					"Unexpected number of token data, iter %d", i)
				require.NotEmptyf(t, execReport.Proofs, "Proof should not be empty.")
				assertMerkleRoot(t, hasher, execReport, tt.args.reports[i])
			}
			// If the last report is partially executed, the executed messages can be checked.
			if len(updatedMessages) > 0 && len(tt.lastReportExecuted) > 0 {
				lastReport := updatedMessages[len(updatedMessages)-1]
				require.ElementsMatch(t, tt.lastReportExecuted, lastReport.ExecutedMessages)
			}
		})
	}
}

// TODO: Use this to test the verifyReport function.
type badCodec struct{}

func (bc badCodec) Encode(ctx context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	return nil, fmt.Errorf("bad codec")
}

func (bc badCodec) Decode(ctx context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	return cciptypes.ExecutePluginReport{}, fmt.Errorf("bad codec")
}

func Test_execReportBuilder_verifyReport(t *testing.T) {
	type fields struct {
		encoder            cciptypes.ExecutePluginCodec
		maxReportSizeBytes uint64
		accumulated        validationMetadata
	}
	type args struct {
		execReport cciptypes.ExecutePluginReportSingleChain
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedLog      string
		expectedIsValid  bool
		expectedMetadata validationMetadata
		expectedError    string
	}{
		{
			name: "empty report",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{},
			},
			fields: fields{
				maxReportSizeBytes: 1000,
			},
			expectedIsValid: true,
			expectedMetadata: validationMetadata{
				encodedSizeBytes: 120,
			},
		},
		{
			name: "good report",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
						makeMessage(1, 101, 0),
						makeMessage(1, 102, 0),
						makeMessage(1, 103, 0),
					},
				},
			},
			fields: fields{
				maxReportSizeBytes: 10000,
			},
			expectedIsValid: true,
			expectedMetadata: validationMetadata{
				encodedSizeBytes: 1633,
			},
		},
		{
			name: "oversized report",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
						makeMessage(1, 101, 0),
						makeMessage(1, 102, 0),
						makeMessage(1, 103, 0),
					},
				},
			},
			fields: fields{
				maxReportSizeBytes: 1000,
			},
			expectedLog: "invalid report, report size exceeds limit",
		},
		{
			name: "oversized report - accumulated size",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
						makeMessage(1, 101, 0),
						makeMessage(1, 102, 0),
						makeMessage(1, 103, 0),
					},
				},
			},
			fields: fields{
				accumulated: validationMetadata{
					encodedSizeBytes: 1000,
				},
				maxReportSizeBytes: 2000,
			},
			expectedLog: "invalid report, report size exceeds limit",
		},
		{
			name: "bad token data reader",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
						makeMessage(1, 101, 0),
					},
				},
			},
			fields: fields{
				encoder: badCodec{},
			},
			expectedError: "unable to encode report",
			expectedLog:   "unable to encode report",
		},
	}
	for _, tt := range tests {
		tt := tt
		lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Select token data reader mock.
			var resolvedEncoder cciptypes.ExecutePluginCodec
			if tt.fields.encoder != nil {
				resolvedEncoder = tt.fields.encoder
			} else {
				resolvedEncoder = mocks.NewExecutePluginJSONReportCodec()
			}

			b := &execReportBuilder{
				ctx:                context.Background(),
				lggr:               lggr,
				encoder:            resolvedEncoder,
				maxReportSizeBytes: tt.fields.maxReportSizeBytes,
				accumulated:        tt.fields.accumulated,
			}
			isValid, metadata, err := b.verifyReport(context.Background(), tt.args.execReport)
			if tt.expectedError != "" {
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}
			if tt.expectedLog != "" {
				found := false
				for _, log := range logs.All() {
					fmt.Println(log.Message)
					found = found || strings.Contains(log.Message, tt.expectedLog)
				}
				assert.True(t, found, "expected log not found")
			}
			assert.Equalf(t, tt.expectedIsValid, isValid, "verifyReport(...)")
			assert.Equalf(t, tt.expectedMetadata, metadata, "verifyReport(...)")
		})
	}
}

// TODO: better than this
type tdr struct {
	mode tokenDataMode
}

type tokenDataMode int

const (
	noop tokenDataMode = iota + 1
	good
	bad
	notReady
)

func (t tdr) ReadTokenData(
	ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error,
) {
	switch t.mode {
	case noop:
		return nil, nil
	case good:
		return [][]byte{{0x01, 0x02, 0x03}}, nil
	case bad:
		return nil, fmt.Errorf("bad token data reader")
	case notReady:
		return nil, ErrNotReady
	default:
		panic("mode should be 1, 2 or 3")
	}
}

func Test_execReportBuilder_checkMessage(t *testing.T) {
	type fields struct {
		tokenDataReader types.TokenDataReader
	}
	type args struct {
		idx        int
		execReport plugintypes.ExecutePluginCommitData
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedData     plugintypes.ExecutePluginCommitData
		expectedStatus   messageStatus
		expectedMetadata validationMetadata
		expectedError    string
		expectedLog      string
	}{
		{
			name:          "empty",
			expectedError: "message index out of range",
			expectedLog:   "message index out of range",
		},
		{
			name: "already executed",
			args: args{
				idx: 0,
				execReport: plugintypes.ExecutePluginCommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
					ExecutedMessages: []cciptypes.SeqNum{100},
				},
			},
			expectedStatus: AlreadyExecuted,
			expectedLog:    "message already executed",
		},
		{
			name: "bad token data",
			args: args{
				idx: 0,
				execReport: plugintypes.ExecutePluginCommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
				},
			},
			fields: fields{
				tokenDataReader: tdr{mode: bad},
			},
			expectedStatus: TokenDataFetchError,
			expectedLog:    "unable to read token data - unknown error",
		},
		{
			name: "token data not ready",
			args: args{
				idx: 0,
				execReport: plugintypes.ExecutePluginCommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
				},
			},
			fields: fields{
				tokenDataReader: tdr{mode: notReady},
			},
			expectedStatus: TokenDataNotReady,
			expectedLog:    "unable to read token data - token data not ready",
		},
		{
			name: "good token data is cached",
			args: args{
				idx: 0,
				execReport: plugintypes.ExecutePluginCommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
				},
			},
			fields: fields{
				tokenDataReader: tdr{mode: good},
			},
			expectedStatus: ReadyToExecute,
			expectedData: plugintypes.ExecutePluginCommitData{
				Messages: []cciptypes.Message{
					makeMessage(1, 100, 0),
				},
				TokenData: [][][]byte{{{0x01, 0x02, 0x03}}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)

			// Select token data reader mock.
			var resolvedTokenDataReader types.TokenDataReader
			if tt.fields.tokenDataReader != nil {
				resolvedTokenDataReader = tt.fields.tokenDataReader
			} else {
				resolvedTokenDataReader = tdr{mode: good}
			}

			b := &execReportBuilder{
				lggr:            lggr,
				tokenDataReader: resolvedTokenDataReader,
			}
			data, status, err := b.checkMessage(context.Background(), tt.args.idx, tt.args.execReport)
			if tt.expectedError != "" {
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}
			if tt.expectedLog != "" {
				found := false
				for _, log := range logs.All() {
					fmt.Println(log.Message)
					found = found || strings.Contains(log.Message, tt.expectedLog)
				}
				assert.True(t, found, "expected log not found")
			}
			assert.Equalf(t, tt.expectedStatus, status, "checkMessage(...)")
			// If expected data not provided, we expect the result to be the same as the input.
			if reflect.DeepEqual(tt.expectedData, plugintypes.ExecutePluginCommitData{}) {
				assert.Equalf(t, tt.args.execReport, data, "checkMessage(...)")
			} else {
				assert.Equalf(t, tt.expectedData, data, "checkMessage(...)")
			}
		})
	}
}
