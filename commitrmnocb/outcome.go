package commitrmnocb

import (
	"fmt"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/utils"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
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
		outcome = ReportRangesOutcome(commitQuery, consensusObservation, previousOutcome)

	case BuildingReport:
		outcome = buildReport(commitQuery, consensusObservation, p.cfg.OffchainConfig, previousOutcome.LastPricesUpdate)

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
	previousOutcome Outcome,
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
		LastPricesUpdate:        previousOutcome.LastPricesUpdate,
	}

	return outcome
}

// Given a set of observed merkle roots, gas prices and token prices, and roots from RMN, construct a report
// to transmit on-chain
func buildReport(
	_ Query,
	consensusObservation ConsensusObservation,
	offchainCfg pluginconfig.CommitOffchainConfig,
	lastUpdate time.Time,
) Outcome {
	roots := maps.Values(consensusObservation.MerkleRoots)

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	var lastPricesUpdate time.Time
	tokenPrices := consensusObservation.TokenPricesArray()

	// If we're updating all token prices, set the last prices update to the observed timestamp,
	// otherwise keep it as the last update time.
	if len(offchainCfg.TokenInfo) == len(tokenPrices) {
		lastPricesUpdate = consensusObservation.Timestamp
	} else {
		lastPricesUpdate = lastUpdate
	}

	outcome := Outcome{
		OutcomeType:      outcomeType,
		RootsToReport:    roots,
		GasPrices:        consensusObservation.GasPricesArray(),
		TokenPrices:      consensusObservation.TokenPricesArray(),
		LastPricesUpdate: lastPricesUpdate,
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
			OutcomeType:      ReportTransmitted,
			LastPricesUpdate: previousOutcome.LastPricesUpdate,
		}
	}

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= maxReportTransmissionCheckAttempts {
		lggr.Warnw("Failed to detect report transmission")
		return Outcome{
			OutcomeType:      ReportTransmissionFailed,
			LastPricesUpdate: previousOutcome.LastPricesUpdate,
		}
	}

	return Outcome{
		OutcomeType:                     ReportInFlight,
		OffRampNextSeqNums:              previousOutcome.OffRampNextSeqNums,
		ReportTransmissionCheckAttempts: previousOutcome.ReportTransmissionCheckAttempts + 1,
		LastPricesUpdate:                previousOutcome.LastPricesUpdate,
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
	timestampConsensus := utils.MedianTimestamp(aggObs.Timestamps)

	fDestChain, exists := fChains[destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", destChain)
	}

	feedPrices := feedPricesConsensus(lggr, aggObs.FeedTokenPrices, fDestChain)
	registryPrices := registryPricesConsensus(lggr, aggObs.PriceRegistryTokenPrices, fDestChain)
	tokenPrices := selectTokens(
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
		TokenPrices:        tokenPrices,
		OnRampMaxSeqNums:   onRampMaxSeqNumsConsensus(lggr, aggObs.OnRampMaxSeqNums, fChains),
		OffRampNextSeqNums: offRampMaxSeqNumsConsensus(lggr, aggObs.OffRampNextSeqNums, fDestChain),
		FChain:             fChains,
		Timestamp:          timestampConsensus,
	}

	return consensusObs, nil
}

// selectTokens checks which tokens need to be updated based on the observed token prices and the price registry updates
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
	nextPricesUpdate := lastPricesUpdate.Add(updateFrequency.Duration())
	if consensusTimestamp.After(nextPricesUpdate) {
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
		if utils.Deviates(feedPrice.Int, registryPrice.Int, deviation) {
			tokenPrices[token] = feedPrice
		}
	}

	return tokenPrices
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
		tokenPrices[token] = utils.MedianBigInt(prices)
	}
	return tokenPrices
}

// registryPricesConsensus returns the median of the price registry token prices for each
// token given all observed updates
func registryPricesConsensus(
	lggr logger.Logger,
	priceRegistryPrices map[types.Account][]cciptypes.BigInt,
	fDestChain int,
) map[types.Account]cciptypes.BigInt {
	tokenPrices := make(map[types.Account]cciptypes.BigInt)
	for token, prices := range priceRegistryPrices {
		if len(prices) < 2*fDestChain+1 {
			lggr.Warnf("could not reach consensus on price registry token prices for token %s ", token)
			continue
		}
		medianPrice := utils.MedianBigInt(prices)
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
			root, count := utils.MostFrequentElem(roots)

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
		seqNum, count := utils.MostFrequentElem(offRampMaxSeqNums)
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
		fChain, count := utils.MostFrequentElem(fValues)
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
