package v1

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var dataGenerators = []dataGenerator{
	smallGen,
	// medGen,
	// largeGen,
	// xLargeGen,
}

func TestCommitQuery(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.commitQuery(),
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeQuery(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeQuery(i.(committypes.Query)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeQuery(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeQuery(i.(committypes.Query)) },
		)
		results = append(results, result)
	}
	fmt.Println(resultDataArray(results))
}

func TestCommitObservation(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.commitObservation(),
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeObservation(i.(committypes.Observation)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeObservation(i.(committypes.Observation)) },
		)
		results = append(results, result)
	}
	fmt.Println(resultDataArray(results))
}

func TestCommitOutcome(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.commitOutcome(),
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeOutcome(i.(committypes.Outcome)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeOutcome(i.(committypes.Outcome)) },
		)
		results = append(results, result)
	}
	fmt.Println(resultDataArray(results))
}

func TestExecObservation(t *testing.T) {
	jsonCodec := NewExecCodecJSON()
	protoCodec := NewExecCodecProto()

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.execObservation(),
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeObservation(i.(exectypes.Observation)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeObservation(i.(exectypes.Observation)) },
		)
		results = append(results, result)
	}
	fmt.Println(resultDataArray(results))
}

func TestExecOutcome(t *testing.T) {
	jsonCodec := NewExecCodecJSON()
	protoCodec := NewExecCodecProto()

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.execOutcome(),
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeOutcome(i.(exectypes.Outcome)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeOutcome(i.(exectypes.Outcome)) },
		)
		results = append(results, result)
	}
	fmt.Println(resultDataArray(results))
}

// ------------------------------------------------

func runBenchmark(
	t *testing.T,
	name string,
	obj interface{},
	decodeJSONFunc func([]byte) (interface{}, error),
	encodeJSONFunc func(interface{}) ([]byte, error),
	decodeProtoFunc func([]byte) (interface{}, error),
	encodeProtoFunc func(interface{}) ([]byte, error),
) resultData {
	result := resultData{name: name}

	tStart := time.Now()
	jsonEnc, err := encodeJSONFunc(obj)
	require.NoError(t, err)
	result.jsonEncodingTime = time.Since(tStart)
	tStart = time.Now()
	jsonDec, err := decodeJSONFunc(jsonEnc)
	result.jsonDecodingTime = time.Since(tStart)
	require.NoError(t, err)
	result.jsonEncodingDataLength = len(jsonEnc)

	tStart = time.Now()
	protoEnc, err := encodeProtoFunc(obj)
	require.NoError(t, err)
	result.protoEncodingTime = time.Since(tStart)
	tStart = time.Now()
	protoDec, err := decodeProtoFunc(protoEnc)
	result.protoDecodingTime = time.Since(tStart)
	require.NoError(t, err)
	result.protoEncodingDataLength = len(protoEnc)

	// sanity check
	if !reflect.DeepEqual(jsonDec, protoDec) {
		jsonDecStr, _ := json.Marshal(jsonDec)
		protoDecStr, _ := json.Marshal(protoDec)
		require.JSONEq(t, string(jsonDecStr), string(protoDecStr))
	}

	return result
}

type resultDataArray []resultData

func (r resultDataArray) String() string {
	if len(r) == 0 {
		return "No results available"
	}

	// Table header
	header := []string{"Name", "JSON Enc", "Proto Enc", "JSON Dec", "Proto Dec", "JSON Size", "Proto Size"}
	columnWidths := []int{0, 20, 20, 20, 20, 12, 12}

	for _, entry := range r {
		if columnWidths[0] < len(entry.name) {
			columnWidths[0] = len(entry.name) + 1
		}
	}

	// Table separator
	separator := strings.Repeat("-", sum(columnWidths)+len(columnWidths)*3)

	// Format header row
	var sb strings.Builder
	sb.WriteString(separator + "\n")
	sb.WriteString(formatRow(header, columnWidths) + "\n")
	sb.WriteString(separator + "\n")

	// Format data rows
	for _, data := range r {
		row := []string{
			data.name,
			data.jsonEncodingTime.String(),
			data.protoEncodingTime.String(),
			data.jsonDecodingTime.String(),
			data.protoDecodingTime.String(),
			fmt.Sprintf("%d", data.jsonEncodingDataLength),
			fmt.Sprintf("%d", data.protoEncodingDataLength),
		}
		sb.WriteString(formatRow(row, columnWidths) + "\n")
	}

	sb.WriteString(separator)
	return sb.String()
}

// formatRow formats a row with padding for each column
func formatRow(fields []string, widths []int) string {
	var parts []string
	for i, field := range fields {
		parts = append(parts, fmt.Sprintf("%-*s", widths[i], field))
	}
	return "| " + strings.Join(parts, " | ") + " |"
}

// sum calculates the total width of all columns
func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

type resultData struct {
	name                    string
	jsonEncodingTime        time.Duration
	protoEncodingTime       time.Duration
	jsonDecodingTime        time.Duration
	protoDecodingTime       time.Duration
	jsonEncodingDataLength  int
	protoEncodingDataLength int
}

type dataGenerator struct {
	name                 string
	numRmnNodes          int
	numSourceChains      int
	numPricedTokens      int
	numContractsPerChain int
	numMessagesPerChain  int
	numTokensPerMsg      int
}

var (
	smallGen = dataGenerator{
		name:                 "small",
		numRmnNodes:          4,
		numSourceChains:      8,
		numPricedTokens:      8,
		numContractsPerChain: 12,
		numMessagesPerChain:  4,
		numTokensPerMsg:      1,
	}

	// The following data generators are disabled to prevent long-running tests,
	// but they can be re-enabled to benchmark any future changes.

	//medGen = dataGenerator{
	//	name:                 "medium",
	//	numRmnNodes:          16,
	//	numSourceChains:      64,
	//	numPricedTokens:      28,
	//	numContractsPerChain: 18,
	//	numMessagesPerChain:  16,
	//	numTokensPerMsg:      4,
	//}
	//
	//largeGen = dataGenerator{
	//	name:                 "large",
	//	numRmnNodes:          32,
	//	numSourceChains:      256,
	//	numPricedTokens:      64,
	//	numContractsPerChain: 18,
	//	numMessagesPerChain:  64,
	//	numTokensPerMsg:      8,
	//}
	//
	//xLargeGen = dataGenerator{
	//	name:                 "xlarge",
	//	numRmnNodes:          64,
	//	numSourceChains:      512,
	//	numPricedTokens:      128,
	//	numContractsPerChain: 32,
	//	numMessagesPerChain:  256,
	//	numTokensPerMsg:      64,
	//}
)

func (d *dataGenerator) commitQuery() committypes.Query {
	sigs := make([]*serialization.EcdsaSignature, d.numRmnNodes)
	for i := 0; i < d.numRmnNodes; i++ {
		sigs[i] = &serialization.EcdsaSignature{
			R: randomBytes(32),
			S: randomBytes(32),
		}
	}

	laneUpdates := make([]*serialization.FixedDestLaneUpdate, d.numSourceChains)
	for i := 0; i < d.numSourceChains; i++ {
		laneUpdates[i] = &serialization.FixedDestLaneUpdate{
			LaneSource: &serialization.LaneSource{
				SourceChainSelector: rand.Uint64(),
				OnrampAddress:       randomBytes(40),
			},
			ClosedInterval: &serialization.ClosedInterval{
				MinMsgNr: rand.Uint64(),
				MaxMsgNr: rand.Uint64(),
			},
			Root: randomBytes(32),
		}
	}

	return committypes.Query{
		MerkleRootQuery: merkleroot.Query{
			RetryRMNSignatures: rand.Uint32()%2 == 0,
			RMNSignatures: &rmn.ReportSignatures{
				Signatures:  sigs,
				LaneUpdates: laneUpdates,
			},
		},
		TokenPriceQuery: tokenprice.Query{},
		ChainFeeQuery:   chainfee.Query{},
	}
}

func (d *dataGenerator) commitObservation() committypes.Observation {
	fChain := make(map[cciptypes.ChainSelector]int, d.numSourceChains)
	merkleRoots := genMerkleRootChain(d.numSourceChains)
	rmnEnabledChains := genRmnEnabledChains(d.numSourceChains)
	onRampMaxSeqNums := genSeqNumChain(d.numSourceChains)
	offRampNextSeqNums := genSeqNumChain(d.numSourceChains)
	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, d.numSourceChains)
	nativeTokenPrices := make(map[cciptypes.ChainSelector]cciptypes.BigInt, d.numSourceChains)
	chainFeeUpdates := make(map[cciptypes.ChainSelector]chainfee.Update, d.numSourceChains)

	for i := 0; i < d.numSourceChains; i++ {
		fChain[cciptypes.ChainSelector(rand.Uint64())] = rand.Intn(256)

		feeComponents[cciptypes.ChainSelector(rand.Uint64())] = types.ChainFeeComponents{
			ExecutionFee:        randBigInt().Int,
			DataAvailabilityFee: randBigInt().Int,
		}

		nativeTokenPrices[cciptypes.ChainSelector(rand.Uint64())] = randBigInt()

		chainFeeUpdates[cciptypes.ChainSelector(rand.Uint64())] = chainfee.Update{
			ChainFee: chainfee.ComponentsUSDPrices{
				ExecutionFeePriceUSD: randBigInt().Int,
				DataAvFeePriceUSD:    randBigInt().Int,
			},
			Timestamp: time.Now().UTC(),
		}
	}

	feedTokenPrices := make(map[cciptypes.UnknownEncodedAddress]cciptypes.BigInt, d.numPricedTokens)
	feeQuoterTokenUpdates := make(map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig, d.numPricedTokens)

	for i := 0; i < d.numPricedTokens; i++ {
		feedTokenPrices[cciptypes.UnknownEncodedAddress(genRandomString(40))] = randBigInt()

		feeQuoterTokenUpdates[cciptypes.UnknownEncodedAddress(genRandomString(40))] = cciptypes.TimestampedBig{
			Timestamp: time.Now().UTC(),
			Value:     randBigInt(),
		}
	}

	return committypes.Observation{
		MerkleRootObs: merkleroot.Observation{
			MerkleRoots:        merkleRoots,
			RMNEnabledChains:   rmnEnabledChains,
			OnRampMaxSeqNums:   onRampMaxSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			RMNRemoteConfig:    genRmnRemoteConfig(d.numRmnNodes),
			FChain:             fChain,
		},
		TokenPriceObs: tokenprice.Observation{
			FeedTokenPrices:       feedTokenPrices,
			FeeQuoterTokenUpdates: feeQuoterTokenUpdates,
			FChain:                fChain,
			Timestamp:             time.Now().UTC(),
		},
		ChainFeeObs: chainfee.Observation{
			FeeComponents:     feeComponents,
			NativeTokenPrices: nativeTokenPrices,
			ChainFeeUpdates:   chainFeeUpdates,
			FChain:            fChain,
			TimestampNow:      time.Now().UTC(),
		},
		DiscoveryObs: genDiscoveryObservation(d.numSourceChains, d.numContractsPerChain),
		FChain:       fChain,
	}
}

func (d *dataGenerator) commitOutcome() committypes.Outcome {
	rmnReportSigs := make([]cciptypes.RMNECDSASignature, d.numRmnNodes)
	for i := 0; i < d.numRmnNodes; i++ {
		rmnReportSigs[i] = cciptypes.RMNECDSASignature{
			R: randomBytes32(),
			S: randomBytes32(),
		}
	}

	tokenPrices := make(cciptypes.TokenPriceMap)
	for i := 0; i < d.numPricedTokens; i++ {
		tokenPrices[cciptypes.UnknownEncodedAddress(genRandomString(40))] = randBigInt()
	}

	gasPrices := make([]cciptypes.GasPriceChain, d.numSourceChains)
	for i := 0; i < d.numSourceChains; i++ {
		gasPrices[i] = cciptypes.GasPriceChain{
			ChainSel: cciptypes.ChainSelector(rand.Uint64()),
			GasPrice: randBigInt(),
		}
	}

	return committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType:                     merkleroot.OutcomeType(rand.Int() % 128),
			RangesSelectedForReport:         genChainRanges(d.numSourceChains),
			RootsToReport:                   genMerkleRootChain(d.numSourceChains),
			RMNEnabledChains:                genRmnEnabledChains(d.numSourceChains),
			OffRampNextSeqNums:              genSeqNumChain(d.numSourceChains),
			ReportTransmissionCheckAttempts: uint(rand.Intn(128)),
			RMNReportSignatures:             rmnReportSigs,
			RMNRemoteCfg:                    genRmnRemoteConfig(d.numRmnNodes),
		},
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: tokenPrices,
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: gasPrices,
		},
		MainOutcome: committypes.MainOutcome{
			InflightPriceOcrSequenceNumber: cciptypes.SeqNum(rand.Uint64()),
			RemainingPriceChecks:           rand.Int() % 256,
		},
	}
}

func (d *dataGenerator) execObservation() exectypes.Observation {
	discoveryObs := genDiscoveryObservation(d.numSourceChains, d.numContractsPerChain)

	msgs := genMessages(d.numMessagesPerChain, d.numTokensPerMsg)

	tokenDataObservations := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData)
	msgHashObservations := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32)
	nonces := make(map[cciptypes.ChainSelector]map[string]uint64)
	msgObservations := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message)
	for i := 0; i < d.numSourceChains; i++ {
		chainSel := cciptypes.ChainSelector(rand.Uint64())
		nonces[chainSel] = map[string]uint64{
			genRandomString(5): rand.Uint64(),
		}

		tokenDataObservations[chainSel] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData, d.numMessagesPerChain)
		msgHashObservations[chainSel] = make(map[cciptypes.SeqNum]cciptypes.Bytes32, d.numMessagesPerChain)
		msgObservations[chainSel] = make(map[cciptypes.SeqNum]cciptypes.Message, d.numMessagesPerChain)

		for j := 0; j < d.numMessagesPerChain; j++ {
			tokenDataObservations[chainSel][cciptypes.SeqNum(rand.Uint64())] = exectypes.MessageTokenData{
				TokenData: []exectypes.TokenData{
					{
						Ready:     rand.Int()%2 == 0,
						Data:      randomBytes(128),
						Error:     nil,
						Supported: false,
					},
				},
			}

			msgObservations[chainSel][cciptypes.SeqNum(rand.Uint64())] = msgs[0]
			msgHashObservations[chainSel][cciptypes.SeqNum(rand.Uint64())] = randomBytes32()
		}

	}

	commitReports := make(map[cciptypes.ChainSelector][]exectypes.CommitData)
	for i := 0; i < d.numSourceChains; i++ {
		commitReports[cciptypes.ChainSelector(rand.Uint64())] = []exectypes.CommitData{
			{
				SourceChain:   cciptypes.ChainSelector(rand.Uint64()),
				OnRampAddress: randomBytes(40),
				Timestamp:     time.Now().UTC(),
				BlockNum:      rand.Uint64(),
				MerkleRoot:    randomBytes32(),
				SequenceNumberRange: cciptypes.NewSeqNumRange(
					cciptypes.SeqNum(rand.Uint64()),
					cciptypes.SeqNum(rand.Uint64()),
				),
				ExecutedMessages: genSeqNums(d.numMessagesPerChain / 2),
				Messages:         msgs,
				Hashes:           genBytes32Slice(d.numMessagesPerChain),
				MessageTokenData: []exectypes.MessageTokenData{
					{
						TokenData: []exectypes.TokenData{
							{
								Ready:     rand.Int()%2 == 0,
								Data:      randomBytes(128),
								Error:     nil,
								Supported: false,
							},
						},
					},
				},
			},
		}
	}

	return exectypes.Observation{
		CommitReports: commitReports,
		Messages:      msgObservations,
		Hashes:        msgHashObservations,
		TokenData:     tokenDataObservations,
		Nonces:        nonces,
		Contracts:     discoveryObs,
		FChain:        discoveryObs.FChain,
	}
}

func (d *dataGenerator) execOutcome() exectypes.Outcome {
	commitReports := make([]exectypes.CommitData, d.numSourceChains)
	msgs := genMessages(d.numMessagesPerChain, d.numTokensPerMsg)

	for i := 0; i < d.numSourceChains; i++ {
		commitReports[i] = exectypes.CommitData{
			SourceChain:   cciptypes.ChainSelector(rand.Uint64()),
			OnRampAddress: randomBytes(40),
			Timestamp:     time.Now().UTC(),
			BlockNum:      rand.Uint64(),
			MerkleRoot:    randomBytes32(),
			SequenceNumberRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(rand.Uint64()),
				cciptypes.SeqNum(rand.Uint64()),
			),
			ExecutedMessages: genSeqNums(d.numMessagesPerChain / 2),
			Messages:         msgs,
			Hashes:           genBytes32Slice(d.numMessagesPerChain),
			MessageTokenData: []exectypes.MessageTokenData{
				{
					TokenData: []exectypes.TokenData{
						{
							Ready:     rand.Int()%2 == 0,
							Data:      randomBytes(128),
							Error:     nil,
							Supported: false,
						},
					},
				},
			},
		}
	}

	chainReports := make([]cciptypes.ExecutePluginReportSingleChain, d.numSourceChains)
	for i := 0; i < d.numSourceChains; i++ {
		tokenData := make([][][]byte, 0)
		for j := 0; j < d.numTokensPerMsg/4; j++ {
			tokenData = append(tokenData,
				[][]byte{randomBytes(32), randomBytes(32), randomBytes(32), randomBytes(32)})
		}

		chainReports[i] = cciptypes.ExecutePluginReportSingleChain{
			SourceChainSelector: cciptypes.ChainSelector(rand.Uint64()),
			Messages:            msgs,
			OffchainTokenData:   tokenData,
			Proofs:              genBytes32Slice(d.numMessagesPerChain),
			ProofFlagBits:       randBigInt(),
		}
	}

	return exectypes.Outcome{
		State:         exectypes.PluginState(genRandomString(128)),
		CommitReports: commitReports,
		Report: cciptypes.ExecutePluginReport{
			ChainReports: chainReports,
		},
	}
}

// ------------------------------------------------

func genDiscoveryObservation(numSourceChains, numContractsPerChain int) dt.Observation {
	fChain := make(map[cciptypes.ChainSelector]int, numSourceChains)
	for i := 0; i < numSourceChains; i++ {
		fChain[cciptypes.ChainSelector(rand.Uint64())] = rand.Intn(256)
	}

	contractAddresses := make(reader.ContractAddresses, numContractsPerChain)
	for i := 0; i < numContractsPerChain; i++ {
		contractName := fmt.Sprintf("contract-%d", i)
		contractAddresses[contractName] = make(map[cciptypes.ChainSelector]cciptypes.UnknownAddress, numSourceChains)
		for j := 0; j < numSourceChains; j++ {
			contractAddresses[contractName][cciptypes.ChainSelector(rand.Uint64())] = randomBytes(40)
		}
	}

	return dt.Observation{
		FChain:    fChain,
		Addresses: contractAddresses,
	}
}

func genMessages(numMsgs, numTokensPerMsg int) []cciptypes.Message {
	tokenAmounts := make([]cciptypes.RampTokenAmount, numTokensPerMsg)
	for i := 0; i < numTokensPerMsg; i++ {
		tokenAmounts[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: randomBytes(40),
			DestTokenAddress:  randomBytes(40),
			ExtraData:         randomBytes(40),
			Amount:            randBigInt(),
			DestExecData:      randomBytes(40),
		}
	}

	msgs := make([]cciptypes.Message, numMsgs)
	for i := 0; i < numMsgs; i++ {
		msgs[i] = cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				MessageID:           randomBytes32(),
				SourceChainSelector: cciptypes.ChainSelector(rand.Uint64()),
				DestChainSelector:   cciptypes.ChainSelector(rand.Uint64()),
				SequenceNumber:      cciptypes.SeqNum(rand.Uint64()),
				Nonce:               rand.Uint64(),
				MsgHash:             randomBytes32(),
				OnRamp:              randomBytes(40),
			},
			Sender:         randomBytes(40),
			Data:           randomBytes(256),
			Receiver:       randomBytes(40),
			ExtraArgs:      randomBytes(256),
			FeeToken:       randomBytes(40),
			FeeTokenAmount: randBigInt(),
			FeeValueJuels:  randBigInt(),
			TokenAmounts:   tokenAmounts,
		}
	}

	return msgs
}

func genMerkleRootChain(n int) []cciptypes.MerkleRootChain {
	mrcs := make([]cciptypes.MerkleRootChain, n)
	for i := 0; i < n; i++ {
		mrcs[i] = cciptypes.MerkleRootChain{
			ChainSel:      cciptypes.ChainSelector(rand.Uint64()),
			OnRampAddress: randomBytes(40),
			SeqNumsRange:  cciptypes.NewSeqNumRange(cciptypes.SeqNum(rand.Uint64()), cciptypes.SeqNum(rand.Uint64())),
			MerkleRoot:    randomBytes32(),
		}
	}
	return mrcs
}

func genRmnEnabledChains(n int) map[cciptypes.ChainSelector]bool {
	m := make(map[cciptypes.ChainSelector]bool)
	for i := 0; i < n; i++ {
		m[cciptypes.ChainSelector(rand.Uint64())] = rand.Int()%2 == 0
	}
	return m
}

func genSeqNumChain(n int) []plugintypes.SeqNumChain {
	chains := make([]plugintypes.SeqNumChain, n)
	for i := 0; i < n; i++ {
		chains[i] = plugintypes.SeqNumChain{
			ChainSel: cciptypes.ChainSelector(rand.Uint64()),
			SeqNum:   cciptypes.SeqNum(rand.Uint64()),
		}
	}
	return chains
}

func genRmnRemoteConfig(numSigners int) cciptypes.RemoteConfig {
	rmnSigners := make([]cciptypes.RemoteSignerInfo, numSigners)
	for i := 0; i < numSigners; i++ {
		rmnSigners[i] = cciptypes.RemoteSignerInfo{
			OnchainPublicKey: randomBytes(20),
			NodeIndex:        rand.Uint64(),
		}
	}

	return cciptypes.RemoteConfig{
		ContractAddress:  randomBytes(40),
		ConfigDigest:     randomBytes32(),
		Signers:          rmnSigners,
		FSign:            rand.Uint64(),
		ConfigVersion:    rand.Uint32(),
		RmnReportVersion: randomBytes32(),
	}
}

func randBigInt() cciptypes.BigInt {
	return cciptypes.NewBigInt(big.NewInt(rand.Int63()))
}

func genChainRanges(n int) []plugintypes.ChainRange {
	ranges := make([]plugintypes.ChainRange, n)
	for i := 0; i < n; i++ {
		ranges[i] = plugintypes.ChainRange{
			ChainSel: cciptypes.ChainSelector(rand.Uint64()),
			SeqNumRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(rand.Uint64()),
				cciptypes.SeqNum(rand.Uint64()),
			),
		}
	}
	return ranges
}

func genRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func randomBytes32() [32]byte {
	return [32]byte(randomBytes(32))
}

func genSeqNums(n int) []cciptypes.SeqNum {
	seqNums := make([]cciptypes.SeqNum, n)
	for i := 0; i < n; i++ {
		seqNums[i] = cciptypes.SeqNum(rand.Uint64())
	}
	return seqNums
}

func genBytes32Slice(n int) []cciptypes.Bytes32 {
	result := make([]cciptypes.Bytes32, n)
	for i := 0; i < n; i++ {
		result[i] = randomBytes32()
	}
	return result
}
