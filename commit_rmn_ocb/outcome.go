package commitrmnocb

import (
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Outcome TODO: doc
func (p *Plugin) Outcome(
	outCtx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	_, nextState := p.decodeOutcome(outCtx.PreviousOutcome)
	commitQuery, err := DecodeCommitPluginQuery(query)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	consensusObservation, err := p.getConsensusObservation(aos)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	switch nextState {
	case SelectingRangesForReport:
		return p.ReportRangesOutcome(commitQuery, consensusObservation)

	case BuildingReport:
		return nil, nil

	case WaitingForReportTransmission:
		return nil, nil

	default:
		p.log.Warnw("Unexpected state", "state", nextState)
		return nil, nil
	}
}

func (p *Plugin) ReportRangesOutcome(query CommitQuery, consensusObservation ConsensusObservation) (ocr3types.Outcome, error) {
	return nil, nil
}

//// ReportRangesOutcome TODO: doc
//func (p *Plugin) ReportRangesOutcome(query CommitQuery, consensusObservation ConsensusObservation) (ocr3types.Outcome, error) {
//	nextRanges := getRangesForNextReport(
//		p.log,
//		query.RmnOnRampMaxSeqNums,
//		aggregatedObservation.OnRampMaxSeqNums,
//		aggregatedObservation.OffRampMaxSeqNums,
//	)
//
//	outcome := CommitPluginOutcome{
//		OutcomeType:             ReportIntervalsSelected,
//		RangesSelectedForReport: nextRanges,
//	}
//
//	return outcome.Encode()
//}
//
//// TODO: doc
//func getRangesForNextReport(
//	log logger.Logger,
//	rmnOnRampMaxSeqNums,
//	observedOnRampMaxSeqNums,
//	observedOffRampMaxSeqNums []plugintypes.SeqNumChain,
//) []ChainRange {
//	rangesToReport := make([]ChainRange, 0, len(observedOffRampMaxSeqNums))
//
//	rmnOnRampMaxSeqNumsMap := seqNumChainArrayToMap(rmnOnRampMaxSeqNums)
//	observedOnRampMaxSeqNumsMap := seqNumChainArrayToMap(observedOnRampMaxSeqNums)
//	observedOffRampMaxSeqNumsMap := seqNumChainArrayToMap(observedOffRampMaxSeqNums)
//
//	for chainSel, offRampMaxSeqNum := range observedOffRampMaxSeqNumsMap {
//		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
//		if !exists {
//			continue
//		}
//
//		if rmnOnRampMaxSeqNum, exists := rmnOnRampMaxSeqNumsMap[chainSel]; exists {
//			onRampMaxSeqNum = min(onRampMaxSeqNum, rmnOnRampMaxSeqNum)
//		}
//
//		if offRampMaxSeqNum > onRampMaxSeqNum {
//			log.Warnw("Found an offRampMaxSeqNum greater than an onRampMaxSeqNum",
//				"offRampMaxSeqNum", offRampMaxSeqNum,
//				"onRampMaxSeqNum", onRampMaxSeqNum,
//				"chainSelector", chainSel)
//			continue
//		} else if offRampMaxSeqNum == onRampMaxSeqNum {
//			continue
//		} else {
//			chainRange := ChainRange{
//				ChainSel:    chainSel,
//				SeqNumRange: [2]cciptypes.SeqNum{offRampMaxSeqNum, onRampMaxSeqNum},
//			}
//			rangesToReport = append(rangesToReport, chainRange)
//		}
//	}
//
//	return rangesToReport
//}
//
//func seqNumChainArrayToMap(seqNumChains []plugintypes.SeqNumChain) map[cciptypes.ChainSelector]cciptypes.SeqNum {
//	chainToSeqNum := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
//	for _, seqNumChain := range seqNumChains {
//		chainToSeqNum[seqNumChain.ChainSel] = seqNumChain.SeqNum
//	}
//
//	return chainToSeqNum
//}

// getConsensusObservation TODO: doc
func (p *Plugin) getConsensusObservation(aos []types.AttributedObservation) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)
	fChains := consensus(p.log, aggObs.FChain, p.fChainConsensus)
	// TODO: doc, fix
	var fTokenChain int
	var fDestChain int

	merkleRootConsensusFn := func(chain cciptypes.ChainSelector, roots []cciptypes.Bytes32) (cciptypes.Bytes32, error) {
		return p.merkleRootConsensus(chain, roots, fChains)
	}

	gasPricesConsensusFn := func(chain cciptypes.ChainSelector, prices []cciptypes.BigInt) (cciptypes.BigInt, error) {
		return p.gasPriceConsensus(chain, prices, fChains)
	}

	tokenPricesConsensusFn := func(tokenID types.Account, prices []cciptypes.BigInt) (cciptypes.BigInt, error) {
		return p.tokenPriceConsensus(tokenID, prices, fTokenChain)
	}

	onRampMaxSeqNumsConsensusFn := func(chain cciptypes.ChainSelector, seqNums []cciptypes.SeqNum) (cciptypes.SeqNum, error) {
		return p.onRampMaxSeqNumsConsensus(chain, seqNums, fChains)
	}

	offRampMaxSeqNumsConsensusFn := func(chain cciptypes.ChainSelector, seqNums []cciptypes.SeqNum) (cciptypes.SeqNum, error) {
		return p.offRampMaxSeqNumsConsensus(chain, seqNums, fDestChain)
	}

	consensusObs := ConsensusObservation{
		MerkleRoots:       consensus(p.log, aggObs.MerkleRoots, merkleRootConsensusFn),
		GasPrices:         consensus(p.log, aggObs.GasPrices, gasPricesConsensusFn),
		TokenPrices:       consensus(p.log, aggObs.TokenPrices, tokenPricesConsensusFn),
		OnRampMaxSeqNums:  consensus(p.log, aggObs.OnRampMaxSeqNums, onRampMaxSeqNumsConsensusFn),
		OffRampMaxSeqNums: consensus(p.log, aggObs.OffRampMaxSeqNums, offRampMaxSeqNumsConsensusFn),
		FChain:            fChains,
	}

	return consensusObs, nil
}

func (p *Plugin) offRampMaxSeqNumsConsensus(
	chain cciptypes.ChainSelector,
	offRampMaxSeqNums []cciptypes.SeqNum,
	fDestChain int,
) (cciptypes.SeqNum, error) {
	seqNum, count := mostFrequentElem(offRampMaxSeqNums)
	if count <= fDestChain {
		return 0, fmt.Errorf("could not reach consensus on offRampMaxSeqNums for chain %d "+
			"because we did not receive a sequence number that was observed by at least f (%d) oracles, "+
			"offRampMaxSeqNums: %v", chain, fDestChain, offRampMaxSeqNums)
	}

	return seqNum, nil
}

func (p *Plugin) onRampMaxSeqNumsConsensus(
	chain cciptypes.ChainSelector,
	onRampMaxSeqNums []cciptypes.SeqNum,
	fChains map[cciptypes.ChainSelector]int,
) (cciptypes.SeqNum, error) {
	if f, exists := fChains[chain]; exists {
		if len(onRampMaxSeqNums) < 2*f+1 {
			return 0, fmt.Errorf("could not reach consensus on onRampMaxSeqNums for chain %d "+
				"because we did not receive more than 2f+1 observed sequence numbers, 2f+1: %d, "+
				"len(onRampMaxSeqNums): %d, onRampMaxSeqNums: %v",
				chain, 2*f+1, len(onRampMaxSeqNums), onRampMaxSeqNums)
		}

		sort.Slice(onRampMaxSeqNums, func(i, j int) bool { return i > j })
		return onRampMaxSeqNums[f], nil
	} else {
		return 0, fmt.Errorf("could not reach consensus on onRampMaxSeqNums for chain %d "+
			"because there was no consensus f value for this chain", chain)
	}
}

func (p *Plugin) merkleRootConsensus(
	chain cciptypes.ChainSelector,
	roots []cciptypes.Bytes32,
	fChains map[cciptypes.ChainSelector]int,
) (cciptypes.Bytes32, error) {
	if f, exists := fChains[chain]; exists {
		root, count := mostFrequentElem(roots)

		if count <= f {
			return cciptypes.Bytes32{}, fmt.Errorf("failed to reach consensus on a merkle root for chain %d "+
				"because no single merkle root was observed more than the expected %d times, found merkle root %d "+
				"observed by only %d oracles, observed merkle roots: %v",
				chain, f, root, count, roots)
		}

		return root, nil
	} else {
		return cciptypes.Bytes32{}, fmt.Errorf("merkleRootConsensus: fChain not found for chain %d", chain)
	}
}

func (p *Plugin) tokenPriceConsensus(
	tokenID types.Account,
	prices []cciptypes.BigInt,
	fTokenChain int,
) (cciptypes.BigInt, error) {
	if len(prices) < 2*fTokenChain+1 {
		return cciptypes.BigInt{}, fmt.Errorf("could not reach consensus on token prices for token %s because "+
			"we did not receive more than 2f+1 observed prices, 2f+1: %d, len(prices): %d, prices: %v",
			tokenID, 2*fTokenChain+1, len(prices), prices)
	}

	return slicelib.BigIntSortedMiddle(prices), nil
}

func (p *Plugin) gasPriceConsensus(
	chain cciptypes.ChainSelector,
	prices []cciptypes.BigInt,
	fChains map[cciptypes.ChainSelector]int,
) (cciptypes.BigInt, error) {
	if f, exists := fChains[chain]; exists {
		if len(prices) < 2*f+1 {
			return cciptypes.BigInt{}, fmt.Errorf("could not reach consensus on gas prices for chain %d "+
				"because we did not receive more than 2f+1 observed prices, 2f+1: %d, len(prices): %d, prices: %v",
				chain, 2*f+1, len(prices), prices)
		}

		return slicelib.BigIntSortedMiddle(prices), nil
	} else {
		return cciptypes.BigInt{}, fmt.Errorf("could not reach consensus on gas prices for chain %d because "+
			"there was no consensus f value for this chain", chain)
	}
}

// fChainConsensus TODO: doc
func (p *Plugin) fChainConsensus(chain cciptypes.ChainSelector, fValues []int) (int, error) {
	f, count := mostFrequentElem(fValues)

	if count < p.reportingCfg.F {
		return 0, fmt.Errorf("failed to reach consensus on fChain values for chain %d because no single f "+
			"value was observed more than the expected %d times, found f value %d observed by only %d oracles, "+
			"f values: %v",
			chain, p.reportingCfg.F, f, count, fValues)
	}

	return f, nil
}

// mostFrequentElem TODO: doc
func mostFrequentElem[T comparable](elems []T) (T, int) {
	var mostFrequentElem T

	counts := counts(elems)
	maxCount := 0

	for elem, count := range counts {
		if count > maxCount {
			mostFrequentElem = elem
			maxCount = count
		}
	}

	return mostFrequentElem, maxCount
}

// consensus TODO: doc, rename
func consensus[V, K comparable](log logger.Logger, elems map[K][]V, consensus func(K, []V) (V, error)) map[K]V {
	consensusMap := make(map[K]V)
	for key, values := range elems {
		consensusValue, err := consensus(key, values)
		if err != nil {
			log.Warnw("consensus error", "err", err)
			continue
		}
		consensusMap[key] = consensusValue
	}

	return consensusMap
}

// counts TODO: doc
func counts[T comparable](elems []T) map[T]int {
	m := make(map[T]int)
	for _, elem := range elems {
		if _, exists := m[elem]; exists {
			m[elem]++
		} else {
			m[elem] = 1
		}
	}

	return m
}
