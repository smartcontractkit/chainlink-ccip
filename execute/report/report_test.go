package report

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	testhelpersrand "github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	gasmock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func randomAddress() string {
	b := make([]byte, 20)
	_, _ = crand.Read(b) // Assignment for errcheck. Only used in tests so we can ignore.
	return cciptypes.Bytes(b).String()
}

// mustMakeBytes parses a given string into a byte array, any error causes a panic. Pass in an empty string for a
// random byte array.
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

// assertMerkleRoot computes the source messages merkle root,
// then computes a verification with the proof, then compares
// the roots.
func assertMerkleRoot(
	t *testing.T,
	hasher cciptypes.MessageHasher,
	execReport cciptypes.ExecutePluginReportSingleChain,
	commitReport exectypes.CommitData,
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
			Nonce:               nonce,
		},
	}
}

// makeMessageWithSender is the same as makeMessage but overrides the sender to the provided input.
//
//nolint:unparam
func makeMessageWithSender(
	src cciptypes.ChainSelector,
	seqNum cciptypes.SeqNum,
	nonce uint64,
	sender cciptypes.UnknownAddress,
) cciptypes.Message {
	base := makeMessage(src, seqNum, nonce)
	base.Sender = sender
	return base
}

// makeTestCommitReportWithSenders is the same as makeTestCommitReport but overrides the senders.
func makeTestCommitReportWithSenders(
	hasher cciptypes.MessageHasher,
	numMessages,
	srcChain,
	firstSeqNum,
	block int,
	timestamp int64,
	senders []cciptypes.UnknownAddress,
	rootOverride cciptypes.Bytes32,
	executed []cciptypes.SeqNum,
) exectypes.CommitData {
	if len(senders) == 0 || len(senders) != numMessages {
		panic("wrong number of senders provided")
	}

	data := makeTestCommitReport(hasher, numMessages, srcChain, firstSeqNum,
		block, timestamp, senders[0], rootOverride, executed, false)

	for i := range data.Messages {
		data.Messages[i].Sender = senders[i]
	}

	return data
}

// makeTestCommitReport creates a basic commit report with messages given different parameters. This function
// will panic if the input parameters are inconsistent.
// if zeroNonces is set to true, nonces will be zero for all messages.
// if its false, they will increment starting from 1.
func makeTestCommitReport(
	hasher cciptypes.MessageHasher,
	numMessages,
	srcChain,
	firstSeqNum,
	block int,
	timestamp int64,
	sender cciptypes.UnknownAddress,
	rootOverride cciptypes.Bytes32,
	executed []cciptypes.SeqNum,
	zeroNonces bool,
) exectypes.CommitData {
	sequenceNumberRange :=
		cciptypes.NewSeqNumRange(cciptypes.SeqNum(firstSeqNum), cciptypes.SeqNum(firstSeqNum+numMessages-1))

	for _, e := range executed {
		if !sequenceNumberRange.Contains(e) {
			panic("executed message out of range")
		}
	}
	var messages []cciptypes.Message
	for i := 0; i < numMessages; i++ {
		var msg cciptypes.Message
		if zeroNonces {
			msg = makeMessage(
				cciptypes.ChainSelector(srcChain),
				cciptypes.SeqNum(i+firstSeqNum),
				0,
			)
		} else {

			msg = makeMessage(
				cciptypes.ChainSelector(srcChain),
				cciptypes.SeqNum(i+firstSeqNum),
				uint64(i)+1,
			)
		}

		msg.Sender = sender
		messages = append(messages, msg)
	}

	var hashes []cciptypes.Bytes32
	for i := 0; i < len(messages); i++ {
		hash, err := hasher.Hash(context.Background(), messages[i])
		if err != nil {
			panic(fmt.Sprintf("unable to hash message: %s", err))
		}
		hashes = append(hashes, hash)
	}

	messageTokenData := make([]exectypes.MessageTokenData, numMessages)
	for i := 0; i < len(messages); i++ {
		messageTokenData[i] = exectypes.MessageTokenData{
			TokenData: []exectypes.TokenData{
				{
					Ready: true,
					Data:  []byte(fmt.Sprintf("data %d", i)),
				},
			},
		}
	}

	commitReport := exectypes.CommitData{
		//MerkleRoot:          root,
		SourceChain:         cciptypes.ChainSelector(srcChain),
		SequenceNumberRange: sequenceNumberRange,
		Timestamp:           time.UnixMilli(timestamp),
		BlockNum:            uint64(block),
		Messages:            messages,
		Hashes:              hashes,
		ExecutedMessages:    executed,
		MessageTokenData:    messageTokenData,
	}

	// calculate merkle root
	root := rootOverride
	if root.IsEmpty() {
		tree, err := ConstructMerkleTree(commitReport, logger.Nop())
		if err != nil {
			panic(fmt.Sprintf("unable to construct merkle tree: %s", err))
		}
		commitReport.MerkleRoot = tree.Root()
	}

	return commitReport
}

// breakCommitReport by adding an extra message. This causes the report to have an unexpected number of messages.
func breakCommitReport(
	commitReport exectypes.CommitData,
) exectypes.CommitData {
	commitReport.Messages = append(commitReport.Messages, cciptypes.Message{})
	return commitReport
}

// setMessageData at the given index to the given size. This function will panic if the index is out of range.
func setMessageData(
	idx int, size uint64, commitReport exectypes.CommitData,
) exectypes.CommitData {
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
		report exectypes.CommitData
		hasher cciptypes.MessageHasher
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name:    "token data mismatch",
			wantErr: "token data length mismatch: got 2, expected 1",
			args: args{
				report: exectypes.CommitData{
					MessageTokenData:    make([]exectypes.MessageTokenData, 2),
					SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{}},
					},
				},
			},
		},
		{
			name:    "wrong number of messages",
			wantErr: "unexpected number of messages: expected 1, got 2",
			args: args{
				report: exectypes.CommitData{
					MessageTokenData:    make([]exectypes.MessageTokenData, 2),
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
				report: exectypes.CommitData{
					MessageTokenData:    make([]exectypes.MessageTokenData, 2),
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
					Hashes: []cciptypes.Bytes32{
						{0x1},
						{0x2},
					},
				},
			},
		},
		{
			name:    "source mismatch",
			wantErr: "unexpected source chain: expected 1111, got 2222",
			args: args{
				report: exectypes.CommitData{
					MessageTokenData:    make([]exectypes.MessageTokenData, 1),
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
					Hashes: []cciptypes.Bytes32{
						{0x1},
					},
				},
				hasher: badHasher{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			msgs := make(map[int]struct{})
			for i := 0; i < len(tt.args.report.Messages); i++ {
				msgs[i] = struct{}{}
			}

			test := func(readyMessages map[int]struct{}) {
				_, err := buildSingleChainReportHelper(lggr, tt.args.report, readyMessages)
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			}

			// Test with pre-built and all messages.
			test(msgs)
			test(nil)
		})
	}
}

func Test_Builder_Build(t *testing.T) {
	hasher := mocks.NewMessageHasher()
	codec := mocks.NewExecutePluginJSONReportCodec()
	lggr := logger.Test(t)
	sender, err := cciptypes.NewUnknownAddressFromHex(randomAddress())
	require.NoError(t, err)
	defaultNonces := map[cciptypes.ChainSelector]map[string]uint64{
		1: {
			sender.String(): 0,
		},
		2: {
			sender.String(): 0,
		},
	}
	tenSenders := make([]cciptypes.UnknownAddress, 10)
	for i := range tenSenders {
		tenSenders[i], err = cciptypes.NewUnknownAddressFromHex(randomAddress())
		require.NoError(t, err)
	}

	type args struct {
		reports       []exectypes.CommitData
		nonces        map[cciptypes.ChainSelector]map[string]uint64
		maxReportSize uint64
		maxGasLimit   uint64
		maxMessages   uint64
		maxReports    uint64
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
				reports: []exectypes.CommitData{},
			},
			expectedExecReports:   0,
			expectedCommitReports: 0,
		},
		{
			name: "half report",
			args: args{
				maxReportSize: 2854,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
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
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
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
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
				},
			},
			expectedExecReports:   2,
			expectedCommitReports: 0,
			expectedExecThings:    []int{10, 20},
		},
		{
			name: "one and half reports",
			args: args{
				maxReportSize: 9700,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
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
				maxReportSize: 4800,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
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
				maxReportSize: 2654,
				maxGasLimit:   10000000,
				nonces: map[cciptypes.ChainSelector]map[string]uint64{
					1: {
						sender.String(): 5,
					},
				},
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						[]cciptypes.SeqNum{100, 101, 102, 103, 104}, // executed
						false, // zeroNonces
					),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{5},
		},
		{
			name: "partially execute remainder of partially executed report",
			args: args{
				maxReportSize: 2270,
				maxGasLimit:   10000000,
				nonces: map[cciptypes.ChainSelector]map[string]uint64{
					1: {
						sender.String(): 5,
					},
				},
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						[]cciptypes.SeqNum{100, 101, 102, 103, 104}, // executed
						false, // zeroNonces
					),
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
				maxGasLimit:   10000000,
				nonces: map[cciptypes.ChainSelector]map[string]uint64{
					1: {
						tenSenders[1].String(): 1,
						tenSenders[3].String(): 3,
						tenSenders[5].String(): 5,
						tenSenders[7].String(): 7,
						tenSenders[9].String(): 9,
					},
				},
				reports: []exectypes.CommitData{
					makeTestCommitReportWithSenders(hasher, 10, 1, 100, 999, 10101010101,
						tenSenders,
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
				maxReportSize: 2408,
				maxGasLimit:   10000000,
				nonces: map[cciptypes.ChainSelector]map[string]uint64{
					1: {
						tenSenders[1].String(): 1,
						tenSenders[3].String(): 3,
						tenSenders[5].String(): 5,
						tenSenders[7].String(): 7,
						tenSenders[9].String(): 9,
					},
				},
				reports: []exectypes.CommitData{
					makeTestCommitReportWithSenders(hasher, 10, 1, 100, 999, 10101010101,
						tenSenders,
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
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					breakCommitReport(makeTestCommitReport(hasher, 10, 1, 101, 1000, 10101010102,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					)),
				},
			},
			wantErr: "unable to add a single chain report",
		},
		{
			name: "invalid merkle root",
			args: args{
				nonces: defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						mustMakeBytes(""), // random root
						nil,               // executed
						false,             // zeroNonces
					),
				},
			},
			wantErr: "merkle root mismatch: expected 0x00000000000000000",
		},
		{
			name: "skip over one large message in OOO",
			args: args{
				maxReportSize: 10000,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					setMessageData(5, 20000, // message at index 5 is too big.
						makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
							sender,
							cciptypes.Bytes32{}, // generate a correct root.
							nil,                 // executed
							true,                // zeroNonces must be set otherwise the skip logic will fail the nonce check.
						)),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{9},
			// seq num 105 is skipped due to size.
			// because of out of order execution, the rest can be safely included.
			lastReportExecuted: []cciptypes.SeqNum{100, 101, 102, 103, 104, 106, 107, 108, 109},
		},
		{
			name: "skip over one large messages in ordered execution",
			args: args{
				maxReportSize: 10000,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					setMessageData(5, 20000,
						makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
							sender,
							cciptypes.Bytes32{}, // generate a correct root.
							nil,                 // executed
							false,               // zeroNonces
						)),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			// message at index 5 is the large one, and because its skipped, so are the rest, due to nonce check.
			expectedExecThings: []int{5},
			// seq num 105 is initially skipped due to size.
			// the rest are skipped due to skipped nonce checks.
			lastReportExecuted: []cciptypes.SeqNum{100, 101, 102, 103, 104},
		},
		{
			name: "skip over two large messages in OOO",
			args: args{
				maxReportSize: 10000,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					setMessageData(8, 20000,
						setMessageData(5, 20000,
							makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
								sender,
								cciptypes.Bytes32{}, // generate a correct root.
								nil,                 // executed
								true,                // zeroNonces must be set otherwise the skip logic will fail the nonce check.
							))),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{8},
			// seq num 105 and 108 are skipped due to size.
			// because of out of order execution, the rest can be safely included.
			lastReportExecuted: []cciptypes.SeqNum{100, 101, 102, 103, 104, 106, 107, 109},
		},
		{
			name: "skip over two large messages in ordered execution",
			args: args{
				maxReportSize: 10000,
				maxGasLimit:   10000000,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					setMessageData(8, 20000,
						setMessageData(5, 20000,
							makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
								sender,
								cciptypes.Bytes32{}, // generate a correct root.
								nil,                 // executed
								false,               // zeroNonces
							))),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{5},
			// seq num 105 is skipped due to size.
			// the rest are skipped due to skipped nonce checks.
			lastReportExecuted: []cciptypes.SeqNum{100, 101, 102, 103, 104},
		},
		{
			name: "report num message limiting",
			args: args{
				maxReportSize: 10000000,
				maxGasLimit:   10000000,
				maxMessages:   3,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,                 // executed
						false,               // zeroNonces
					),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{3},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102},
		},
		{
			name: "report and size limiting (skip over two large messages)",
			args: args{
				maxReportSize: 10000,
				maxGasLimit:   10000000,
				maxMessages:   3,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					setMessageData(1, 90000,
						setMessageData(3, 90000,
							makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
								sender,
								cciptypes.Bytes32{},
								nil,  // executed
								true, // zeroNonces must be set otherwise the skip logic will fail the nonce check.
							))),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{3},
			lastReportExecuted:    []cciptypes.SeqNum{100, 102, 104},
		},
		{
			name: "two reports limited to one",
			args: args{
				maxReportSize: 15000,
				maxGasLimit:   10000000,
				maxReports:    1,
				nonces:        defaultNonces,
				reports: []exectypes.CommitData{
					makeTestCommitReport(hasher, 10, 1, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,
						false,
					),
					makeTestCommitReport(hasher, 20, 2, 100, 999, 10101010101,
						sender,
						cciptypes.Bytes32{}, // generate a correct root.
						nil,
						false,
					),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{10},
		},
	}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			// look for error in Add or Build
			foundError := false

			ep := gasmock.NewMockEstimateProvider(t)
			ep.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(uint64(0)).Maybe()
			ep.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(uint64(0)).Maybe()

			builder := NewBuilder(
				lggr,
				hasher,
				codec,
				ep,
				1, // destChainSelector
				mockAddrCodec,
				WithMaxReportSizeBytes(tt.args.maxReportSize),
				WithMaxGas(tt.args.maxGasLimit),
				WithMaxMessages(tt.args.maxMessages),
				WithMaxSingleChainReports(tt.args.maxReports),
				WithExtraMessageCheck(CheckNonces(tt.args.nonces, mockAddrCodec)),
			)

			var updatedMessages []exectypes.CommitData
			for _, report := range tt.args.reports {
				updatedMessage, err := builder.Add(ctx, report)
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
			execReports, commitReports, err := builder.Build()
			if tt.wantErr != "" {
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Len(t, execReports, tt.expectedExecReports)
			require.Equal(t, len(execReports), len(commitReports))
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
		estimateProvider   cciptypes.EstimateProvider
		maxReportSizeBytes uint64
		maxGas             uint64
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
				maxGas:             1000000,
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
				maxGas:             1000000,
				estimateProvider: func() cciptypes.EstimateProvider {
					// values taken from evm.EstimateProvider
					mockep := gasmock.NewMockEstimateProvider(t)
					mockep.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(uint64(119920)).Times(4)
					mockep.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(uint64(2560)).Times(1)
					return mockep
				}(),
			},
			expectedIsValid: true,
			expectedMetadata: validationMetadata{
				encodedSizeBytes: 1717,
				gas:              482_240,
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
				maxGas:             1000000,
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
				maxGas:             1000000,
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

			ep := tt.fields.estimateProvider
			if ep == nil {
				mockep := gasmock.NewMockEstimateProvider(t)
				mockep.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(uint64(0)).Maybe()
				mockep.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(uint64(0)).Maybe()
				ep = mockep
			}

			b := newBuilderInternal(
				lggr,
				nil,
				resolvedEncoder,
				ep,
				1,
				internal.NewMockAddressCodecHex(t),
				WithMaxReportSizeBytes(tt.fields.maxReportSizeBytes),
				WithMaxGas(tt.fields.maxGas),
			)
			b.accumulated = tt.fields.accumulated

			isValid, metadata, err := b.verifyReport(context.Background(), tt.args.execReport)
			if tt.expectedError != "" {
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}
			require.NoError(t, err)
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

func Test_execReportBuilder_checkMessage(t *testing.T) {
	type args struct {
		idx        int
		nonces     map[cciptypes.ChainSelector]map[string]uint64
		execReport exectypes.CommitData
	}
	tests := []struct {
		name             string
		args             args
		expectedData     exectypes.CommitData
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
				execReport: exectypes.CommitData{
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
			name: "token data not ready",
			args: args{
				idx: 0,
				execReport: exectypes.CommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: false, Data: nil}}},
					},
				},
			},
			expectedStatus: TokenDataNotReady,
			expectedLog:    "unable to read token data - token data not ready",
		},
		{
			name: "good token data is cached",
			args: args{
				idx: 0,
				execReport: exectypes.CommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{0x01, 0x02, 0x03}}}},
					},
				},
			},
			expectedStatus: ReadyToExecute,
			expectedData: exectypes.CommitData{
				Messages: []cciptypes.Message{
					makeMessage(1, 100, 0),
				},
				MessageTokenData: []exectypes.MessageTokenData{
					{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{0x01, 0x02, 0x03}}}},
				},
			},
		},
		{
			name: "good - no token data - 1 msg",
			args: args{
				idx: 0,
				execReport: exectypes.CommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					},
				},
			},
			expectedStatus: ReadyToExecute,
			expectedData: exectypes.CommitData{
				Messages: []cciptypes.Message{
					makeMessage(1, 100, 0),
				},
				MessageTokenData: []exectypes.MessageTokenData{
					{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
				},
			},
		},
		{
			name: "good - no token data - 2nd msg pads slice",
			args: args{
				idx: 1,
				execReport: exectypes.CommitData{
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 0),
						makeMessage(1, 101, 0),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					},
				},
			},
			expectedStatus: ReadyToExecute,
			expectedData: exectypes.CommitData{
				Messages: []cciptypes.Message{
					makeMessage(1, 100, 0),
					makeMessage(1, 101, 0),
				},
				MessageTokenData: []exectypes.MessageTokenData{
					{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
				},
			},
		},
		{
			name: "missing chain nonce",
			args: args{
				idx: 0,
				execReport: exectypes.CommitData{
					SourceChain: 1,
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 1),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					},
				},
			},
			expectedStatus: MissingNoncesForChain,
		},
		{
			name: "missing sender nonce",
			args: args{
				idx: 0,
				nonces: map[cciptypes.ChainSelector]map[string]uint64{
					1: {},
				},
				execReport: exectypes.CommitData{
					SourceChain: 1,
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 1),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					},
				},
			},
			expectedStatus: MissingNonce,
		},
		{
			name: "invalid sender nonce",
			args: args{
				idx: 0,
				nonces: map[cciptypes.ChainSelector]map[string]uint64{
					1: {
						"0x": 99,
					},
				},
				execReport: exectypes.CommitData{
					SourceChain: 1,
					Messages: []cciptypes.Message{
						makeMessage(1, 100, 1),
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					},
				},
			},
			expectedStatus: InvalidNonce,
		},
		{
			name: "pseudo deleted",
			args: args{
				idx: 0,
				execReport: exectypes.CommitData{
					SourceChain: 1,
					Messages: []cciptypes.Message{
						{}, // pseudodeleted
					},
					MessageTokenData: []exectypes.MessageTokenData{
						{TokenData: []exectypes.TokenData{{Ready: true, Data: []byte{}}}},
					},
				},
			},
			expectedStatus: PseudoDeleted,
		},
	}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)

			ep := gasmock.NewMockEstimateProvider(t)
			ep.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(uint64(0)).Maybe()
			ep.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(uint64(0)).Maybe()

			b := newBuilderInternal(
				lggr,
				nil,
				nil,
				nil,
				1,
				internal.NewMockAddressCodecHex(t),
				WithExtraMessageCheck(CheckNonces(tt.args.nonces, mockAddrCodec)),
			)
			data, status, err := b.checkMessage(context.Background(), tt.args.idx, tt.args.execReport)
			if tt.expectedError != "" {
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}
			require.NoError(t, err)
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
			if reflect.DeepEqual(tt.expectedData, exectypes.CommitData{}) {
				assert.Equalf(t, tt.args.execReport, data, "checkMessage(...)")
			} else {
				assert.Equalf(t, tt.expectedData, data, "checkMessage(...)")
			}
		})
	}
}

func Test_checkSkippedNonces(t *testing.T) {
	sourceChain1 := cciptypes.ChainSelector(1)
	sender1 := cciptypes.UnknownAddress(testhelpersrand.RandomBytes(32))
	sender2 := cciptypes.UnknownAddress(testhelpersrand.RandomBytes(32))
	type args struct {
		execReport cciptypes.ExecutePluginReportSingleChain
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid nonces from same sender",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessageWithSender(sourceChain1, 100, 1, sender1),
						makeMessageWithSender(sourceChain1, 101, 2, sender1),
						makeMessageWithSender(sourceChain1, 102, 3, sender1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipped nonce from same sender",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessageWithSender(sourceChain1, 100, 1, sender1),
						// nonce 2 is skipped.
						makeMessageWithSender(sourceChain1, 101, 3, sender1),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "different nonce from different sender",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessageWithSender(sourceChain1, 100, 1, sender1),
						// nonce is not really skipped, different sender
						makeMessageWithSender(sourceChain1, 101, 3, sender2),
						makeMessageWithSender(sourceChain1, 102, 2, sender1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "out of order execution message from same sender",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessageWithSender(sourceChain1, 101, 1, sender1),
						// nonce == 0 => out of order execution
						makeMessageWithSender(sourceChain1, 102, 0, sender1),
						makeMessageWithSender(sourceChain1, 103, 2, sender1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple senders valid nonces",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessageWithSender(sourceChain1, 100, 1, sender1),
						makeMessageWithSender(sourceChain1, 101, 2, sender1),
						makeMessageWithSender(sourceChain1, 102, 1, sender2),
						makeMessageWithSender(sourceChain1, 103, 2, sender2),
						makeMessageWithSender(sourceChain1, 104, 3, sender1),
						makeMessageWithSender(sourceChain1, 105, 3, sender2),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple senders skipped nonce on one of them",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						// sender1 is fine
						makeMessageWithSender(sourceChain1, 100, 1, sender1),
						makeMessageWithSender(sourceChain1, 101, 2, sender1),
						// sender2 skips nonce 2
						makeMessageWithSender(sourceChain1, 102, 1, sender2),
						// note that seqNr _cannot_ be 103, because that would indicate
						// that the onchain message indeed had nonce 3, but it should have nonce 2.
						// so, essentially msg (seqNr=103, nonce=2) was skipped.
						makeMessageWithSender(sourceChain1, 104, 3, sender2),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "unordered nonces from same sender",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						// nonce 1 should be included in messages array before nonce 2.
						makeMessageWithSender(sourceChain1, 100, 2, sender1),
						makeMessageWithSender(sourceChain1, 101, 1, sender1),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "multiple senders one of them unordered nonces",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						// sender1 is fine
						makeMessageWithSender(sourceChain1, 100, 1, sender1),
						makeMessageWithSender(sourceChain1, 101, 2, sender1),
						// sender2 has unordered nonces
						makeMessageWithSender(sourceChain1, 102, 2, sender2),
						makeMessageWithSender(sourceChain1, 103, 1, sender2),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "all out of order execution",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						makeMessageWithSender(sourceChain1, 100, 0, sender1),
						makeMessageWithSender(sourceChain1, 101, 0, sender1),
						makeMessageWithSender(sourceChain1, 102, 0, sender1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple senders all out of order execution",
			args: args{
				execReport: cciptypes.ExecutePluginReportSingleChain{
					Messages: []cciptypes.Message{
						// sender1 is fine
						makeMessageWithSender(sourceChain1, 100, 0, sender1),
						makeMessageWithSender(sourceChain1, 101, 0, sender1),
						// sender2 is fine
						makeMessageWithSender(sourceChain1, 102, 0, sender2),
						makeMessageWithSender(sourceChain1, 103, 0, sender2),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAddrCodec := internal.NewMockAddressCodecHex(t)
			err := verifyReportNonceContinuity(mockAddrCodec, tt.args.execReport)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
