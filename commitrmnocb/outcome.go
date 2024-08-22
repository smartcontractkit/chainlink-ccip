package commitrmnocb

import (
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/sharedtypes"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// Outcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (p *Plugin) Outcome(
	outCtx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)
	commitQuery := Query{}

	consensusObservation, err := getConsensusObservation(
		p.lggr,
		p.reportingCfg.F,
		p.cfg.DestChain,
		p.cfg.OffchainConfig,
		previousOutcome.LastPricesUpdate,
		aos,
	)

	if err != nil {
		return ocr3types.Outcome{}, err
	}

	outcome := Outcome{}

	switch nextState {
	case SelectingRangesForReport:
		outcome = ReportRangesOutcome(commitQuery, consensusObservation)

	case BuildingReport:
		outcome = buildReport(commitQuery, consensusObservation)

	case WaitingForReportTransmission:
		outcome = checkForReportTransmission(
			p.lggr, p.cfg.MaxReportTransmissionCheckAttempts, previousOutcome, consensusObservation)

	default:
		p.lggr.Warnw("Unexpected state in Outcome", "state", nextState)
		return outcome.Encode()
	}

	p.lggr.Infow("Commit Plugin Outcome", "outcome", outcome, "oid", p.nodeID)
	return outcome.Encode()
}

// ReportRangesOutcome determines the sequence number ranges for each chain to build a report from in the next round
// TODO: ensure each range is below a limit
func ReportRangesOutcome(
	query Query,
	consensusObservation ConsensusObservation,
) Outcome {
	rangesToReport := make([]plugintypes.ChainRange, 0)

	rmnOnRampMaxSeqNumsMap := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, seqNumChain := range query.RmnOnRampMaxSeqNums {
		rmnOnRampMaxSeqNumsMap[seqNumChain.ChainSel] = seqNumChain.SeqNum
	}

	observedOnRampMaxSeqNumsMap := consensusObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consensusObservation.OffRampNextSeqNums

	for chainSel, offRampNextSeqNum := range observedOffRampNextSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if rmnOnRampMaxSeqNum, exists := rmnOnRampMaxSeqNumsMap[chainSel]; exists {
			onRampMaxSeqNum = min(onRampMaxSeqNum, rmnOnRampMaxSeqNum)
		}

		if offRampNextSeqNum <= onRampMaxSeqNum {
			chainRange := plugintypes.ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: [2]cciptypes.SeqNum{offRampNextSeqNum, onRampMaxSeqNum},
			}
			rangesToReport = append(rangesToReport, chainRange)
		}
	}

	outcome := Outcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
	}

	return outcome
}

// Given a set of observed merkle roots, gas prices and token prices, and roots from RMN, construct a report
// to transmit on-chain
func buildReport(
	_ Query,
	consensusObservation ConsensusObservation,
) Outcome {
	roots := maps.Values(consensusObservation.MerkleRoots)

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	outcome := Outcome{
		OutcomeType:   outcomeType,
		RootsToReport: roots,
		GasPrices:     consensusObservation.GasPricesArray(),
		TokenPrices:   consensusObservation.TokenPricesArray(),
	}

	return outcome
}

// checkForReportTransmission checks if the OffRamp has an updated set of max seq nums compared to the seq nums that
// were observed when the most recent report was generated. If an update to these max seq sums is detected, it means
// that the previous report has been transmitted, and we output ReportTransmitted to dictate that a new report
// generation phase should begin. If no update is detected, and we've exhausted our check attempts, output
// ReportTransmissionFailed to signify we stop checking for updates and start a new report generation phase. If no
// update is detected, and we haven't exhausted our check attempts, output ReportInFlight to signify that we should
// check again next round.
func checkForReportTransmission(
	lggr logger.Logger,
	maxReportTransmissionCheckAttempts uint,
	previousOutcome Outcome,
	consensusObservation ConsensusObservation,
) Outcome {

	offRampUpdated := false
	for _, previousSeqNumChain := range previousOutcome.OffRampNextSeqNums {
		if currentSeqNum, exists := consensusObservation.OffRampNextSeqNums[previousSeqNumChain.ChainSel]; exists {
			if previousSeqNumChain.SeqNum != currentSeqNum {
				offRampUpdated = true
				break
			}
		}
	}

	if offRampUpdated {
		return Outcome{
			OutcomeType: ReportTransmitted,
		}
	}

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= maxReportTransmissionCheckAttempts {
		lggr.Warnw("Failed to detect report transmission")
		return Outcome{
			OutcomeType: ReportTransmissionFailed,
		}
	}

	return Outcome{
		OutcomeType:                     ReportInFlight,
		OffRampNextSeqNums:              previousOutcome.OffRampNextSeqNums,
		ReportTransmissionCheckAttempts: previousOutcome.ReportTransmissionCheckAttempts + 1,
	}
}

// getConsensusObservation Combine the list of observations into a single consensus observation
func getConsensusObservation(
	lggr logger.Logger,
	F int,
	destChain cciptypes.ChainSelector,
	offchainCfg pluginconfig.CommitOffchainConfig,
	lastPricesUpdate time.Time,
	aos []types.AttributedObservation,
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)
	fChains := fChainConsensus(lggr, F, aggObs.FChain)
	timestampConsensus := TimestampSortedMiddle(aggObs.Timestamps)

	fDestChain, exists := fChains[destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", destChain)
	}

	feedPrices := feedPricesConsensus(lggr, aggObs.FeedTokenPrices, fDestChain)
	registryPrices := registryPricesConsensus(lggr, aggObs.PriceRegistryTokenUpdates, fDestChain)
	tokePrices := selectTokens(
		lggr,
		feedPrices,
		registryPrices,
		timestampConsensus,
		lastPricesUpdate,
		offchainCfg.TokenPriceBatchWriteFrequency,
		offchainCfg.TokenInfo,
	)
	consensusObs := ConsensusObservation{
		MerkleRoots: merkleRootConsensus(lggr, aggObs.MerkleRoots, fChains),
		// TODO: use consensus of observed gas prices
		GasPrices:          make(map[cciptypes.ChainSelector]cciptypes.BigInt),
		TokenPrices:        tokePrices,
		OnRampMaxSeqNums:   onRampMaxSeqNumsConsensus(lggr, aggObs.OnRampMaxSeqNums, fChains),
		OffRampNextSeqNums: offRampMaxSeqNumsConsensus(lggr, aggObs.OffRampNextSeqNums, fDestChain),
		FChain:             fChains,
	}

	return consensusObs, nil
}

// Checks which tokens need to be updated based on the observed token prices and the price registry updates
// if time passed since the last update is greater than the stale threshold, update all tokens
// otherwise calculate deviation between the price registry and feed and include deviated token prices
func selectTokens(
	lggr logger.Logger,
	medianizedFeedTokenPrices map[types.Account]cciptypes.BigInt,
	medianizedRegistryPrices map[types.Account]cciptypes.BigInt,
	consensusTimestamp time.Time,
	lastPricesUpdate time.Time,
	updateFrequency commonconfig.Duration,
	tokenInfo map[types.Account]pluginconfig.TokenInfo,
) map[types.Account]cciptypes.BigInt {
	tokenPrices := make(map[types.Account]cciptypes.BigInt)
	// if the time since the last update is greater than the update frequency, update all tokens
	if consensusTimestamp.Sub(lastPricesUpdate) > updateFrequency.Duration() {
		return medianizedFeedTokenPrices
	}

	// otherwise, calculate the deviation between the feed and the price registry
	for token, feedPrice := range medianizedFeedTokenPrices {
		registryPrice, exists := medianizedRegistryPrices[token]
		if !exists {
			lggr.Warnf("could not find registry price for token %s", token)
			continue
		}

		ti, ok := tokenInfo[token]
		if !ok {
			lggr.Warnf("could not find token info for token %s", token)
			continue
		}

		deviation := ti.DeviationPPB.Int64()
		if Deviates(feedPrice.Int, registryPrice.Int, deviation) {
			tokenPrices[token] = feedPrice
		}
	}

	return make(map[types.Account]cciptypes.BigInt)
}

// feedPricesConsensus returns the median of the feed token prices for each token given all observed prices
func feedPricesConsensus(
	lggr logger.Logger,
	feedTokenPrices map[types.Account][]cciptypes.BigInt,
	fDestChain int,
) map[types.Account]cciptypes.BigInt {
	tokenPrices := make(map[types.Account]cciptypes.BigInt)
	for token, prices := range feedTokenPrices {
		if len(prices) < 2*fDestChain+1 {
			lggr.Warnf("could not reach consensus on feed token prices for token %s ", token)
			continue
		}
		tokenPrices[token] = BigIntSortedMiddle(prices)
	}
	return tokenPrices
}

// registryPricesConsensus returns the median of the price registry token prices for each token given all observed updates
func registryPricesConsensus(
	lggr logger.Logger,
	priceRegistryTokenUpdates map[types.Account][]sharedtypes.NumericalUpdate,
	fDestChain int,
) map[types.Account]cciptypes.BigInt {
	tokenPrices := make(map[types.Account]cciptypes.BigInt)
	for token, updates := range priceRegistryTokenUpdates {
		if len(updates) < 2*fDestChain+1 {
			lggr.Warnf("could not reach consensus on price registry token updates for token %s ", token)
			continue
		}
		// for each update get the median for the token price
		var prices []cciptypes.BigInt
		//var timestamps []time.Time
		for _, update := range updates {
			prices = append(prices, update.Value)
			//timestamps = append(timestamps, update.Timestamp)
		}
		medianPrice := BigIntSortedMiddle(prices)
		//medianTimestamp := TimestampSortedMiddle(timestamps)
		tokenPrices[token] = cciptypes.NewBigInt(medianPrice.Int)
	}

	return tokenPrices
}

// Given a mapping from chains to a list of merkle roots, return a mapping from chains to a single consensus merkle
// root. The consensus merkle root for a given chain is the merkle root with the most observations that was observed at
// least fChain times.
func merkleRootConsensus(
	lggr logger.Logger,
	rootsByChain map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.MerkleRootChain {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.MerkleRootChain)

	for chain, roots := range rootsByChain {
		if fChain, exists := fChains[chain]; exists {
			root, count := mostFrequentElem(roots)

			if count <= fChain {
				// TODO: metrics
				lggr.Warnf("failed to reach consensus on a merkle root for chain %d "+
					"because no single merkle root was observed more than the expected fChain (%d) times, found "+
					"merkle root %d observed by only %d oracles, all observed merkle roots: %v",
					chain, fChain, root, count, roots)
			} else {
				consensus[chain] = root
			}
		} else {
			// TODO: metrics
			lggr.Warnf("merkleRootConsensus: fChain not found for chain %d", chain)
		}
	}

	return consensus
}

// Given a mapping from chains to a list of max seq nums on their corresponding OnRamp, return a mapping from chains
// to a single max seq num. The consensus max seq num for a given chain is the f'th lowest max seq num if the number
// of max seq num observations is greater or equal than 2f+1, where f is the FChain of the corresponding source chain.
func onRampMaxSeqNumsConsensus(
	lggr logger.Logger,
	onRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, onRampMaxSeqNums := range onRampMaxSeqNumsByChain {
		if fChain, exists := fChains[chain]; exists {
			if len(onRampMaxSeqNums) < 2*fChain+1 {
				// TODO: metrics
				lggr.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
					"because we did not receive more than 2fChain+1 observed sequence numbers, 2fChain+1: %d, "+
					"len(onRampMaxSeqNums): %d, onRampMaxSeqNums: %v",
					chain, 2*fChain+1, len(onRampMaxSeqNums), onRampMaxSeqNums)
			} else {
				sort.Slice(onRampMaxSeqNums, func(i, j int) bool { return onRampMaxSeqNums[i] < onRampMaxSeqNums[j] })
				consensus[chain] = onRampMaxSeqNums[fChain]
			}
		} else {
			// TODO: metrics
			lggr.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
				"because there was no consensus fChain value for this chain", chain)
		}
	}

	return consensus
}

// Given a mapping from chains to a list of max seq nums on the OffRamp, return a mapping from chains
// to a single max seq num. The consensus max seq num for a given chain is the max seq num with the most observations
// that was observed at least f times, where f is the FChain of the dest chain.
func offRampMaxSeqNumsConsensus(
	lggr logger.Logger,
	offRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fDestChain int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, offRampMaxSeqNums := range offRampMaxSeqNumsByChain {
		seqNum, count := mostFrequentElem(offRampMaxSeqNums)
		if count <= fDestChain {
			// TODO: metrics
			lggr.Warnf("could not reach consensus on offRampMaxSeqNums for chain %d "+
				"because we did not receive a sequence number that was observed by at least fChain (%d) oracles, "+
				"offRampMaxSeqNums: %v", chain, fDestChain, offRampMaxSeqNums)
		} else {
			consensus[chain] = seqNum
		}
	}

	return consensus
}

// Given a mapping from chains to a list of FChain values for each chain, return a mapping from chains
// to a single FChain. The consensus FChain for a given chain is the FChain with the most observations
// that was observed at least f times, where f is the F of the DON (p.reportingCfg.F).
func fChainConsensus(
	lggr logger.Logger,
	F int,
	fChainValues map[cciptypes.ChainSelector][]int,
) map[cciptypes.ChainSelector]int {
	consensus := make(map[cciptypes.ChainSelector]int)

	for chain, fValues := range fChainValues {
		fChain, count := mostFrequentElem(fValues)
		if count < F {
			// TODO: metrics
			lggr.Warnf("failed to reach consensus on fChain values for chain %d because no single fChain "+
				"value was observed more than the expected %d times, found fChain value %d observed by only %d oracles, "+
				"fChain values: %v",
				chain, F, fChain, count, fValues)
		}

		consensus[chain] = fChain
	}

	return consensus
}

// Given a list of elems, return the elem that occurs most frequently and how often it occurs
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

// Given a list of elems, return a map from elems to how often they occur in the given list
func counts[T comparable](elems []T) map[T]int {
	m := make(map[T]int)
	for _, elem := range elems {
		m[elem]++
	}

	return m
}

func TimestampSortedMiddle(timestamps []time.Time) time.Time {
	if len(timestamps) == 0 {
		return time.Time{}
	}
	valsCopy := make([]time.Time, len(timestamps))
	copy(valsCopy[:], timestamps[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Before(valsCopy[j])
	})

	return valsCopy[len(valsCopy)/2]
}

// Deviates checks if x1 and x2 deviates based on the provided ppb (parts per billion)
// ppb is calculated based on the smaller value of the two
// e.g, if x1 > x2, deviation_parts_per_billion = ((x1 - x2) / x2) * 1e9
func Deviates(x1, x2 *big.Int, ppb int64) bool {
	// if x1 == 0 or x2 == 0, deviates if x2 != x1, to avoid the relative division by 0 error
	if x1.BitLen() == 0 || x2.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	diff := big.NewInt(0).Sub(x1, x2) // diff = x1-x2
	diff.Mul(diff, big.NewInt(1e9))   // diff = diff * 1e9
	// dividing by the smaller value gives consistent ppb regardless of input order, and supports >100% deviation.
	if x1.Cmp(x2) > 0 {
		diff.Div(diff, x2)
	} else {
		diff.Div(diff, x1)
	}
	return diff.CmpAbs(big.NewInt(ppb)) > 0 // abs(diff) > ppb
}

// BigIntSortedMiddle returns the middle number after sorting the provided numbers. nil is returned if the provided slice is empty.
// If length of the provided slice is even, the right-hand-side value of the middle 2 numbers is returned.
// The objective of this function is to always pick within the range of values reported by honest nodes when we have 2f+1 values.
func BigIntSortedMiddle(vals []cciptypes.BigInt) cciptypes.BigInt {
	if len(vals) == 0 {
		return cciptypes.BigInt{}
	}

	valsCopy := make([]cciptypes.BigInt, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j].Int) == -1
	})
	return valsCopy[len(valsCopy)/2]
}
