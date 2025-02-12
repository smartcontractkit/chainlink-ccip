package ocrtypecodec

import (
	"math/rand"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
)

// genQuery generates a Protobuf Query object with the specified number of signatures and lane updates.
func genQuery(numSigs int, numLaneUpdates int) *ocrtypecodecpb.Query {
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
	query := &ocrtypecodecpb.Query{
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
func genObservation(numMerkleRoots, numSeqNumChains, numSigners int, numTokenPrices, numFeeQuoterUpdates int, numFeeComponents int, numContractNames int) *ocrtypecodecpb.CommitObservation {
	return &ocrtypecodecpb.CommitObservation{
		MerkleRootObs: genMerkleRootObservation(numMerkleRoots, numSeqNumChains, numSigners),
		TokenPriceObs: genTokenPriceObservation(numTokenPrices, numFeeQuoterUpdates),
		ChainFeeObs:   genChainFeeObservation(numFeeComponents),
		DiscoveryObs:  genDiscoveryObservation(numContractNames),
		FChain:        genFChain(5),
	}
}

// genOutcome generates a randomized Outcome struct.
func genOutcome(numRanges, numRoots, numSigners, numSeqNumChains, numGasPrices, numTokenPrices int) *ocrtypecodecpb.CommitOutcome {
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
			ChainSel: uint64(rand.Uint64()),
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
			ChainSel: uint64(rand.Uint64()),
			GasPrice: randomBytes(32),
		}
	}
	return prices
}

// genMainOutcome generates a MainOutcome.
func genMainOutcome() *ocrtypecodecpb.MainOutcome {
	return &ocrtypecodecpb.MainOutcome{
		InflightPriceOcrSequenceNumber: uint64(rand.Uint64()),
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
	rand.Read(b)
	return b
}
