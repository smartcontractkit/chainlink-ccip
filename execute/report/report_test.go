package report

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

	commitReport := plugintypes.ExecutePluginCommitDataWithMessages{
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
		tree, err := ConstructMerkleTree(context.Background(), hasher, commitReport)
		if err != nil {
			panic(fmt.Sprintf("unable to construct merkle tree: %s", err))
		}
		commitReport.MerkleRoot = tree.Root()
	}

	return commitReport
}

// breakCommitReport by adding an extra message. This causes the report to have an unexpected number of messages.
func breakCommitReport(
	commitReport plugintypes.ExecutePluginCommitDataWithMessages,
) plugintypes.ExecutePluginCommitDataWithMessages {
	commitReport.Messages = append(commitReport.Messages, cciptypes.Message{})
	return commitReport
}

// SetMessageData at the given index to the given size. This function will panic if the index is out of range.
func SetMessageData(
	idx int, size uint64, commitReport plugintypes.ExecutePluginCommitDataWithMessages,
) plugintypes.ExecutePluginCommitDataWithMessages {
	if len(commitReport.Messages) < idx {
		panic("message index out of range")
	}
	commitReport.Messages[idx].Data = make([]byte, size)
	return commitReport
}

// TODO: better than this
type tdr struct{}

func (t tdr) ReadTokenData(
	ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error,
) {
	return nil, nil
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

/*
// TODO: Use this to test the verifyReport function.
type badCodec struct{}

func (bc badCodec) Encode(ctx context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	return nil, fmt.Errorf("bad codec")
}

func (bc badCodec) Decode(ctx context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	return cciptypes.ExecutePluginReport{}, fmt.Errorf("bad codec")
}
*/

func Test_buildSingleChainReport_Errors(t *testing.T) {
	lggr := logger.Test(t)

	type args struct {
		report          plugintypes.ExecutePluginCommitDataWithMessages
		hasher          cciptypes.MessageHasher
		tokenDataReader types.TokenDataReader
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
			var resolvedTokenDataReader types.TokenDataReader
			if tt.args.tokenDataReader != nil {
				resolvedTokenDataReader = tt.args.tokenDataReader
			} else {
				resolvedTokenDataReader = tdr{}
			}

			ctx := context.Background()
			msgs := make(map[int]struct{})
			for i := 0; i < len(tt.args.report.Messages); i++ {
				msgs[i] = struct{}{}
			}
			_, err := buildSingleChainReport(
				ctx, lggr, resolvedHasher, resolvedTokenDataReader, tt.args.report, msgs)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func Test_Builder_Build(t *testing.T) {
	hasher := mocks.NewMessageHasher()
	codec := mocks.NewExecutePluginJSONReportCodec()
	lggr := logger.Test(t)
	var tokenDataReader tdr

	type args struct {
		reports       []plugintypes.ExecutePluginCommitDataWithMessages
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
			wantErr: "unable to add a single chain report",
		},
		{
			name: "invalid merkle root",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						mustMakeBytes(""), // random root
						nil),
				},
			},
			wantErr: "merkle root mismatch: expected 0x00000000000000000",
		},
		// TODO: A test that requires skipping over a large message because only a smaller message fits in the report.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			// look for error in Add or Build
			foundError := false

			builder := NewBuilder(ctx, lggr, hasher, tokenDataReader, codec, tt.args.maxReportSize, 0)
			var updatedMessages []plugintypes.ExecutePluginCommitDataWithMessages
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
