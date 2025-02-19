package ocrtypecodec

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
)

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
		t.Errorf("Decoded JSON and Protobuf objects differ %#v != %#v", jsonDec, protoDec)
	}
	return result
}

// Helper functions for pretty-printing results

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

// genQuery generates a Protobuf Query object with the specified number of signatures and lane updates.
func genQuery(numSigs int, numLaneUpdates int) *ocrtypecodecpb.CommitQuery {
	// Generate ECDSA Signatures
	signatures := make([]*ocrtypecodecpb.SignatureEcdsa, numSigs)
	for i := 0; i < numSigs; i++ {
		signatures[i] = &ocrtypecodecpb.SignatureEcdsa{
			R: randomBytes(32), // Simulating 32-byte ECDSA R value
			S: randomBytes(32), // Simulating 32-byte ECDSA S value
		}
	}

	// Generate Lane Updates
	laneUpdates := make([]*ocrtypecodecpb.DestChainUpdate, numLaneUpdates)
	for i := 0; i < numLaneUpdates; i++ {
		laneUpdates[i] = &ocrtypecodecpb.DestChainUpdate{
			LaneSource: &ocrtypecodecpb.SourceChainMeta{
				SourceChainSelector: uint64(rand.Intn(1000)), // Random chain selector
				OnrampAddress:       randomBytes(20),         // Simulated 20-byte address
			},
			SeqNumRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(rand.Intn(10000)),
				MaxMsgNr: uint64(rand.Intn(10000) + 100), // Ensure max > min
			},
			Root: randomBytes(32), // Simulated 32-byte Merkle root
		}
	}

	// Construct the Protobuf Query object
	query := &ocrtypecodecpb.CommitQuery{
		MerkleRootQuery: &ocrtypecodecpb.MerkleRootQuery{
			RetryRmnSignatures: true,
			RmnSignatures: &ocrtypecodecpb.ReportSignatures{
				Signatures:  signatures,
				LaneUpdates: laneUpdates,
			},
		},
	}

	return query
}

// genObservation generates a randomized Observation struct.
func genObservation(
	numMerkleRoots, numSeqNumChains, numSigners int,
	numTokenPrices, numFeeQuoterUpdates int, numFeeComponents int, numContractNames int,
) *ocrtypecodecpb.CommitObservation {
	return &ocrtypecodecpb.CommitObservation{
		MerkleRootObs: genMerkleRootObservation(numMerkleRoots, numSeqNumChains, numSigners),
		TokenPriceObs: genTokenPriceObservation(numTokenPrices, numFeeQuoterUpdates),
		ChainFeeObs:   genChainFeeObservation(numFeeComponents),
		DiscoveryObs:  genDiscoveryObservation(numContractNames),
		FChain:        genFChain(6),
	}
}

// genOutcome generates a randomized Outcome struct.
func genOutcome(
	numRanges, numRoots, numSigners, numSeqNumChains, numGasPrices, numTokenPrices int) *ocrtypecodecpb.CommitOutcome {
	return &ocrtypecodecpb.CommitOutcome{
		MerkleRootOutcome: genMerkleRootOutcome(numRanges, numRoots, numSigners, numSeqNumChains),
		TokenPriceOutcome: genTokenPriceOutcome(numTokenPrices),
		ChainFeeOutcome:   genChainFeeOutcome(numGasPrices),
		MainOutcome:       genMainOutcome(),
	}
}

// genMerkleRootOutcome generates a MerkleRootOutcome.
func genMerkleRootOutcome(numRanges, numRoots, numSigners, numSeqNumChains int) *ocrtypecodecpb.MerkleRootOutcome {
	return &ocrtypecodecpb.MerkleRootOutcome{
		OutcomeType:                     int32(rand.Intn(10)),
		RangesSelectedForReport:         genChainRanges(numRanges),
		RootsToReport:                   genMerkleRootChains(numRoots),
		RmnEnabledChains:                genBoolMap(5),
		OffRampNextSeqNums:              genSeqNumChains(numSeqNumChains),
		ReportTransmissionCheckAttempts: uint32(rand.Intn(10)),
		RmnReportSignatures:             genEcdsaSignatures(3),
		RmnRemoteCfg:                    genRemoteConfig(numSigners),
	}
}

// genChainRanges generates a slice of ChainRange.
func genChainRanges(n int) []*ocrtypecodecpb.ChainRange {
	ranges := make([]*ocrtypecodecpb.ChainRange, n)
	for i := 0; i < n; i++ {
		ranges[i] = &ocrtypecodecpb.ChainRange{
			ChainSel: rand.Uint64(),
			SeqNumRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(rand.Intn(1000)),
				MaxMsgNr: uint64(rand.Intn(1000) + 100), // Ensure max > min
			},
		}
	}
	return ranges
}

// genEcdsaSignatures generates a slice of ECDSA signatures.
func genEcdsaSignatures(n int) []*ocrtypecodecpb.SignatureEcdsa {
	sigs := make([]*ocrtypecodecpb.SignatureEcdsa, n)
	for i := 0; i < n; i++ {
		sigs[i] = &ocrtypecodecpb.SignatureEcdsa{
			R: randomBytes(32),
			S: randomBytes(32),
		}
	}
	return sigs
}

// genTokenPriceOutcome generates a TokenPriceOutcome.
func genTokenPriceOutcome(numTokenPrices int) *ocrtypecodecpb.TokenPriceOutcome {
	return &ocrtypecodecpb.TokenPriceOutcome{
		TokenPrices: genTokenPriceMap2(numTokenPrices),
	}
}

// genTokenPriceMap generates a map of token prices.
func genTokenPriceMap2(n int) map[string][]byte {
	m := make(map[string][]byte)
	for i := 0; i < n; i++ {
		m[genRandomString(5)] = randomBytes(32)
	}
	return m
}

// genChainFeeOutcome generates a ChainFeeOutcome.
func genChainFeeOutcome(numGasPrices int) *ocrtypecodecpb.ChainFeeOutcome {
	return &ocrtypecodecpb.ChainFeeOutcome{
		GasPrices: genGasPrices(numGasPrices),
	}
}

// genGasPrices generates a slice of GasPriceChain.
func genGasPrices(n int) []*ocrtypecodecpb.GasPriceChain {
	prices := make([]*ocrtypecodecpb.GasPriceChain, n)
	for i := 0; i < n; i++ {
		prices[i] = &ocrtypecodecpb.GasPriceChain{
			ChainSel: rand.Uint64(),
			GasPrice: randomBytes(32),
		}
	}
	return prices
}

// genMainOutcome generates a MainOutcome.
func genMainOutcome() *ocrtypecodecpb.MainOutcome {
	return &ocrtypecodecpb.MainOutcome{
		InflightPriceOcrSequenceNumber: rand.Uint64(),
		RemainingPriceChecks:           int32(rand.Intn(5)),
	}
}

// genMerkleRootObservation generates a MerkleRootObservation.
func genMerkleRootObservation(numMerkleRoots, numSeqNumChains, numSigners int) *ocrtypecodecpb.MerkleRootObservation {
	return &ocrtypecodecpb.MerkleRootObservation{
		MerkleRoots:        genMerkleRootChains(numMerkleRoots),
		RmnEnabledChains:   genBoolMap(5),
		OnRampMaxSeqNums:   genSeqNumChains(numSeqNumChains),
		OffRampNextSeqNums: genSeqNumChains(numSeqNumChains),
		RmnRemoteConfig:    genRemoteConfig(numSigners),
		FChain:             genFChain(5),
	}
}

// genMerkleRootChains generates a slice of MerkleRootChain.
func genMerkleRootChains(n int) []*ocrtypecodecpb.MerkleRootChain {
	chains := make([]*ocrtypecodecpb.MerkleRootChain, n)
	for i := 0; i < n; i++ {
		chains[i] = &ocrtypecodecpb.MerkleRootChain{
			ChainSel:      rand.Uint64(),
			OnRampAddress: randomBytes(20),
			SeqNumsRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: rand.Uint64(),
				MaxMsgNr: rand.Uint64(),
			},
			MerkleRoot: randomBytes(32),
		}
	}
	return chains
}

// genSeqNumChains generates a slice of SeqNumChain.
func genSeqNumChains(n int) []*ocrtypecodecpb.SeqNumChain {
	chains := make([]*ocrtypecodecpb.SeqNumChain, n)
	for i := 0; i < n; i++ {
		chains[i] = &ocrtypecodecpb.SeqNumChain{
			ChainSel: rand.Uint64(),
			SeqNum:   rand.Uint64(),
		}
	}
	return chains
}

// genRemoteConfig generates an RmnRemoteConfig.
func genRemoteConfig(numSigners int) *ocrtypecodecpb.RmnRemoteConfig {
	return &ocrtypecodecpb.RmnRemoteConfig{
		ContractAddress:  randomBytes(20),
		ConfigDigest:     randomBytes(32),
		Signers:          genRemoteSigners(numSigners),
		FSign:            rand.Uint64(),
		ConfigVersion:    uint32(rand.Intn(100)),
		RmnReportVersion: randomBytes(32),
	}
}

// genRemoteSigners generates a slice of RemoteSignerInfo.
func genRemoteSigners(n int) []*ocrtypecodecpb.RemoteSignerInfo {
	signers := make([]*ocrtypecodecpb.RemoteSignerInfo, n)
	for i := 0; i < n; i++ {
		signers[i] = &ocrtypecodecpb.RemoteSignerInfo{
			OnchainPublicKey: randomBytes(20),
			NodeIndex:        rand.Uint64(),
		}
	}
	return signers
}

// genTokenPriceObservation generates a TokenPriceObservation.
func genTokenPriceObservation(numTokenPrices, numFeeQuoterUpdates int) *ocrtypecodecpb.TokenPriceObservation {
	return &ocrtypecodecpb.TokenPriceObservation{
		FeedTokenPrices:       genFeedPriceMap(numTokenPrices),
		FeeQuoterTokenUpdates: genFeeQuoterUpdates(numFeeQuoterUpdates),
		FChain:                genFChain(5),
		Timestamp:             timestamppb.Now(),
	}
}

// genTokenPriceMap generates a map of token prices.
func genTokenPriceMap(n int) map[uint64][]byte {
	m := make(map[uint64][]byte)
	for i := 0; i < n; i++ {
		m[rand.Uint64()] = randomBytes(32)
	}
	return m
}

// genFeedPriceMap generates a map of token prices.
func genFeedPriceMap(n int) map[string][]byte {
	m := make(map[string][]byte)
	for i := 0; i < n; i++ {
		m[genRandomString(16)] = randomBytes(32)
	}
	return m
}

// genFeeQuoterUpdates generates a map of fee quoter updates.
func genFeeQuoterUpdates(n int) map[string]*ocrtypecodecpb.TimestampedBig {
	m := make(map[string]*ocrtypecodecpb.TimestampedBig)
	for i := 0; i < n; i++ {
		m[genRandomString(5)] = &ocrtypecodecpb.TimestampedBig{
			Timestamp: timestamppb.Now(),
			Value:     randomBytes(32),
		}
	}
	return m
}

// genChainFeeObservation generates a ChainFeeObservation.
func genChainFeeObservation(numFeeComponents int) *ocrtypecodecpb.ChainFeeObservation {
	return &ocrtypecodecpb.ChainFeeObservation{
		FeeComponents:     genFeeComponents(numFeeComponents),
		NativeTokenPrices: genTokenPriceMap(5),
		ChainFeeUpdates:   genChainFeeUpdates(5),
		FChain:            genFChain(5),
		TimestampNow:      timestamppb.Now(),
	}
}

// genFeeComponents generates a map of ChainFeeComponents.
func genFeeComponents(n int) map[uint64]*ocrtypecodecpb.ChainFeeComponents {
	m := make(map[uint64]*ocrtypecodecpb.ChainFeeComponents)
	for i := 0; i < n; i++ {
		m[rand.Uint64()] = &ocrtypecodecpb.ChainFeeComponents{
			ExecutionFee:        randomBytes(32),
			DataAvailabilityFee: randomBytes(32),
		}
	}
	return m
}

// genChainFeeUpdates generates a map of ChainFeeUpdates.
func genChainFeeUpdates(n int) map[uint64]*ocrtypecodecpb.ChainFeeUpdate {
	m := make(map[uint64]*ocrtypecodecpb.ChainFeeUpdate)
	for i := 0; i < n; i++ {
		m[rand.Uint64()] = &ocrtypecodecpb.ChainFeeUpdate{
			ChainFee: &ocrtypecodecpb.ComponentsUSDPrices{
				ExecutionFeePriceUsd: randomBytes(32),
				DataAvFeePriceUsd:    randomBytes(32),
			},
			Timestamp: timestamppb.Now(),
		}
	}
	return m
}

// genDiscoveryObservation generates a DiscoveryObservation.
func genDiscoveryObservation(numContractNames int) *ocrtypecodecpb.DiscoveryObservation {
	return &ocrtypecodecpb.DiscoveryObservation{
		FChain:        genFChain(5),
		ContractNames: genContractAddresses(numContractNames),
	}
}

// genContractAddresses generates a ContractNameChainAddresses.
func genContractAddresses(n int) *ocrtypecodecpb.ContractNameChainAddresses {
	addresses := make(map[string]*ocrtypecodecpb.ChainAddressMap)
	for i := 0; i < n; i++ {
		addresses[genRandomString(5)] = genChainAddressMap(5)
	}
	return &ocrtypecodecpb.ContractNameChainAddresses{Addresses: addresses}
}

// genChainAddressMap generates a ChainAddressMap.
func genChainAddressMap(n int) *ocrtypecodecpb.ChainAddressMap {
	m := make(map[uint64][]byte)
	for i := 0; i < n; i++ {
		m[rand.Uint64()] = randomBytes(20)
	}
	return &ocrtypecodecpb.ChainAddressMap{ChainAddresses: m}
}

// genFChain generates a map of uint64 to int32.
func genFChain(n int) map[uint64]int32 {
	m := make(map[uint64]int32)
	for i := 0; i < n; i++ {
		m[rand.Uint64()] = int32(rand.Intn(100))
	}
	return m
}

// genBoolMap generates a map of uint64 to bool.
func genBoolMap(n int) map[uint64]bool {
	m := make(map[uint64]bool)
	for i := 0; i < n; i++ {
		m[rand.Uint64()] = rand.Intn(2) == 1
	}
	return m
}

// genRandomString generates a random string of the given length.
func genRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// randomBytes generates a random byte slice of the given length.
func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// genExecObservation generates a randomized ExecObservation for benchmarking.
func genExecObservation(
	numCommitReports, numMessagesPerChain, numTokenDataPerChain, numNoncesPerChain, numCostlyMessages int,
) *ocrtypecodecpb.ExecObservation {
	return &ocrtypecodecpb.ExecObservation{
		CommitReports:         genCommitReports(numCommitReports),
		SeqNumsToMessages:     genMessages(numMessagesPerChain),
		MessageHashes:         genMessageHashes(numMessagesPerChain),
		TokenDataObservations: genTokenDataObservations(numTokenDataPerChain),
		CostlyMessages:        genBytes32Slice(numCostlyMessages),
		Nonces:                genNonces(numNoncesPerChain),
		Contracts:             genDiscoveryObservation(32),
		FChain:                genFChain(5),
	}
}

// genCommitReports generates a map of CommitObservations.
func genCommitReports(n int) map[uint64]*ocrtypecodecpb.CommitObservations {
	commitReports := make(map[uint64]*ocrtypecodecpb.CommitObservations, n)
	for i := 0; i < n; i++ {
		chainSel := rand.Uint64()
		commitReports[chainSel] = &ocrtypecodecpb.CommitObservations{
			CommitData: genCommitData(rand.Intn(5) + 1), // 1 to 5 commit reports per chain
		}
	}
	return commitReports
}

func genTokenDataObservations(n int) *ocrtypecodecpb.TokenDataObservations {
	tokenData := make(map[uint64]*ocrtypecodecpb.SeqNumToTokenData, n)
	for i := 0; i < n; i++ {
		data := make(map[uint64]*ocrtypecodecpb.MessageTokenData, rand.Intn(5))
		for j := 0; j < rand.Intn(5); j++ {
			data[rand.Uint64()] = genMessageTokenDataEntry()
		}
		tokenData[rand.Uint64()] = &ocrtypecodecpb.SeqNumToTokenData{TokenData: data}
	}
	return &ocrtypecodecpb.TokenDataObservations{TokenData: tokenData}
}

// genCommitData generates a slice of CommitData.
func genCommitData(n int) []*ocrtypecodecpb.CommitData {
	commits := make([]*ocrtypecodecpb.CommitData, n)
	for i := 0; i < n; i++ {
		commits[i] = &ocrtypecodecpb.CommitData{
			SourceChain:         rand.Uint64(),
			OnRampAddress:       randomBytes(20),
			Timestamp:           uint64(time.Now().Unix()),
			BlockNum:            rand.Uint64(),
			MerkleRoot:          randomBytes(32),
			SequenceNumberRange: &ocrtypecodecpb.SeqNumRange{MinMsgNr: rand.Uint64(), MaxMsgNr: rand.Uint64() + 100},
			ExecutedMessages:    genSeqNums(rand.Intn(10)),
			Messages:            genMessageSlice(rand.Intn(10)),
			Hashes:              genBytes32Slice(rand.Intn(10)),
			CostlyMessages:      genBytes32Slice(rand.Intn(5)),
			MessageTokenData:    genMessageTokenData(rand.Intn(10)),
		}
	}
	return commits
}

// genSeqNums generates a slice of sequence numbers.
func genSeqNums(n int) []uint64 {
	seqNums := make([]uint64, n)
	for i := 0; i < n; i++ {
		seqNums[i] = rand.Uint64()
	}
	return seqNums
}

// genMessageSlice generates a slice of messages.
func genMessageSlice(n int) []*ocrtypecodecpb.Message {
	messages := make([]*ocrtypecodecpb.Message, n)
	for i := 0; i < n; i++ {
		messages[i] = genMessage()
	}
	return messages
}

// genMessageTokenData generates a slice of MessageTokenData.
func genMessageTokenData(n int) []*ocrtypecodecpb.MessageTokenData {
	tokenData := make([]*ocrtypecodecpb.MessageTokenData, n)
	for i := 0; i < n; i++ {
		tokenData[i] = genMessageTokenDataEntry()
	}
	return tokenData
}

// genMessages generates a map of chain selectors to message observations.
func genMessages(n int) map[uint64]*ocrtypecodecpb.SeqNumToMessage {
	messages := make(map[uint64]*ocrtypecodecpb.SeqNumToMessage, n)
	for i := 0; i < n; i++ {
		chainSel := rand.Uint64()
		messages[chainSel] = &ocrtypecodecpb.SeqNumToMessage{Messages: genMessageMap(rand.Intn(10) + 1)}
	}
	return messages
}

// genMessageMap generates a map of sequence numbers to messages.
func genMessageMap(n int) map[uint64]*ocrtypecodecpb.Message {
	msgs := make(map[uint64]*ocrtypecodecpb.Message, n)
	for i := 0; i < n; i++ {
		msgs[rand.Uint64()] = genMessage()
	}
	return msgs
}

// genMessage generates a single Message.
func genMessage() *ocrtypecodecpb.Message {
	return &ocrtypecodecpb.Message{
		Header:         genMessageHeader(),
		Sender:         randomBytes(20),
		Data:           randomBytes(50),
		Receiver:       randomBytes(20),
		ExtraArgs:      randomBytes(20),
		FeeToken:       randomBytes(20),
		FeeTokenAmount: randomBytes(32),
		FeeValueJuels:  randomBytes(32),
		TokenAmounts:   genRampTokenAmounts(rand.Intn(5)),
	}
}

// genMessageHeader generates a RampMessageHeader.
func genMessageHeader() *ocrtypecodecpb.RampMessageHeader {
	return &ocrtypecodecpb.RampMessageHeader{
		MessageId:           randomBytes(32),
		SourceChainSelector: rand.Uint64(),
		DestChainSelector:   rand.Uint64(),
		SequenceNumber:      rand.Uint64(),
		Nonce:               rand.Uint64(),
		MsgHash:             randomBytes(32),
		OnRamp:              randomBytes(20),
	}
}

func genRampTokenAmounts(n int) []*ocrtypecodecpb.RampTokenAmount {
	amounts := make([]*ocrtypecodecpb.RampTokenAmount, n)
	for i := 0; i < n; i++ {
		amounts[i] = &ocrtypecodecpb.RampTokenAmount{
			SourcePoolAddress: randomBytes(20),
			DestTokenAddress:  randomBytes(20),
			ExtraData:         randomBytes(32),
			Amount:            randomBytes(32),
			DestExecData:      randomBytes(32),
		}
	}
	return amounts
}

// genMessageHashes generates a map of message hashes.
func genMessageHashes(n int) map[uint64]*ocrtypecodecpb.SeqNumToBytes {
	hashes := make(map[uint64]*ocrtypecodecpb.SeqNumToBytes, n)
	for i := 0; i < n; i++ {
		chainSel := rand.Uint64()
		hashes[chainSel] = &ocrtypecodecpb.SeqNumToBytes{SeqNumToBytes: genSeqNumToBytes(rand.Intn(10))}
	}
	return hashes
}

// genSeqNumToBytes generates a map of sequence numbers to bytes32.
func genSeqNumToBytes(n int) map[uint64][]byte {
	result := make(map[uint64][]byte, n)
	for i := 0; i < n; i++ {
		result[rand.Uint64()] = randomBytes(32)
	}
	return result
}

// genMessageTokenDataEntry generates MessageTokenData.
func genMessageTokenDataEntry() *ocrtypecodecpb.MessageTokenData {
	tokenData := make([]*ocrtypecodecpb.TokenData, rand.Intn(5))
	for i := range tokenData {
		tokenData[i] = &ocrtypecodecpb.TokenData{
			Ready: rand.Intn(2) == 1,
			Data:  randomBytes(32),
		}
	}
	return &ocrtypecodecpb.MessageTokenData{TokenData: tokenData}
}

// genNonces generates nonce observations.
func genNonces(n int) map[uint64]*ocrtypecodecpb.StringAddrToNonce {
	nonces := make(map[uint64]*ocrtypecodecpb.StringAddrToNonce, n)
	for i := 0; i < n; i++ {
		chainSel := rand.Uint64()
		addrToNonce := make(map[string]uint64)
		for j := 0; j < rand.Intn(32); j++ {
			addrToNonce[genRandomString(5)] = rand.Uint64()
		}
		nonces[chainSel] = &ocrtypecodecpb.StringAddrToNonce{Nonces: addrToNonce}
	}
	return nonces
}

func genBytes32Slice(n int) [][]byte {
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = randomBytes(32)
	}
	return result
}

// genExecOutcome generates an ExecOutcome with configurable sizes for nested elements.
func genExecOutcome(
	numCommitReports int,
	numMessagesPerCommit int,
	numProofs int,
	numTokenDataEntries int,
) *ocrtypecodecpb.ExecOutcome {
	execOutcome := &ocrtypecodecpb.ExecOutcome{
		PluginState:   "ACTIVE", // Example plugin state
		CommitReports: make([]*ocrtypecodecpb.CommitData, numCommitReports),
		ExecutePluginReport: &ocrtypecodecpb.ExecutePluginReport{
			ChainReports: make([]*ocrtypecodecpb.ChainReport, numCommitReports),
		},
	}

	for i := 0; i < numCommitReports; i++ {
		commitData := &ocrtypecodecpb.CommitData{
			SourceChain:   rand.Uint64(),
			OnRampAddress: randomBytes(32),
			Timestamp:     uint64(time.Now().Unix()),
			BlockNum:      rand.Uint64(),
			MerkleRoot:    randomBytes(32),
			SequenceNumberRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: rand.Uint64(),
				MaxMsgNr: rand.Uint64(),
			},
			ExecutedMessages: make([]uint64, numMessagesPerCommit),
			Messages:         make([]*ocrtypecodecpb.Message, numMessagesPerCommit),
			Hashes:           make([][]byte, numMessagesPerCommit),
			CostlyMessages:   make([][]byte, numMessagesPerCommit),
			MessageTokenData: make([]*ocrtypecodecpb.MessageTokenData, numMessagesPerCommit),
		}

		for j := 0; j < numMessagesPerCommit; j++ {
			commitData.ExecutedMessages[j] = rand.Uint64()
			commitData.Messages[j] = &ocrtypecodecpb.Message{
				Header: &ocrtypecodecpb.RampMessageHeader{
					MessageId:           randomBytes(32),
					SourceChainSelector: rand.Uint64(),
					DestChainSelector:   rand.Uint64(),
					SequenceNumber:      rand.Uint64(),
					Nonce:               rand.Uint64(),
					MsgHash:             randomBytes(32),
					OnRamp:              randomBytes(32),
				},
				Sender:         randomBytes(20),
				Data:           randomBytes(64),
				Receiver:       randomBytes(20),
				ExtraArgs:      randomBytes(16),
				FeeToken:       randomBytes(20),
				FeeTokenAmount: randomBytes(8),
				FeeValueJuels:  randomBytes(8),
				TokenAmounts: []*ocrtypecodecpb.RampTokenAmount{
					{
						SourcePoolAddress: randomBytes(20),
						DestTokenAddress:  randomBytes(20),
						ExtraData:         randomBytes(10),
						Amount:            randomBytes(8),
						DestExecData:      randomBytes(10),
					},
				},
			}
			commitData.Hashes[j] = randomBytes(32)
			commitData.CostlyMessages[j] = randomBytes(32)
			commitData.MessageTokenData[j] = &ocrtypecodecpb.MessageTokenData{
				TokenData: []*ocrtypecodecpb.TokenData{
					{
						Ready: true,
						Data:  randomBytes(32),
					},
				},
			}
		}

		execOutcome.CommitReports[i] = commitData

		// Create corresponding ChainReport
		execOutcome.ExecutePluginReport.ChainReports[i] = &ocrtypecodecpb.ChainReport{
			SourceChainSelector: rand.Uint64(),
			Messages:            commitData.Messages,
			OffchainTokenData:   generateTokenData(numTokenDataEntries),
			Proofs:              generateProofs(numProofs),
			ProofFlagBits:       randomBytes(8),
		}
	}

	return execOutcome
}

// generateTokenData creates a list of token data entries.
func generateTokenData(numEntries int) []*ocrtypecodecpb.RepeatedBytes {
	tokenData := make([]*ocrtypecodecpb.RepeatedBytes, numEntries)
	for i := 0; i < numEntries; i++ {
		tokenData[i] = &ocrtypecodecpb.RepeatedBytes{
			Items: [][]byte{randomBytes(32)},
		}
	}
	return tokenData
}

// generateProofs creates a list of random proofs.
func generateProofs(numProofs int) [][]byte {
	proofs := make([][]byte, numProofs)
	for i := 0; i < numProofs; i++ {
		proofs[i] = randomBytes(32)
	}
	return proofs
}
