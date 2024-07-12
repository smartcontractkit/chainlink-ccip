package commitrmnocb

import (
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
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

	aggregatedObservation, err := p.aggregateObservations(aos)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	switch nextState {
	case SelectingRangesForReport:
		return p.ReportRangesOutcome(commitQuery, aggregatedObservation)

	case BuildingReport:
		return nil, nil

	case WaitingForReportTransmission:
		return nil, nil

	default:
		p.log.Warnw("Unexpected state", "state", nextState)
		return nil, nil
	}
}

// ReportRangesOutcome TODO: doc
func (p *Plugin) ReportRangesOutcome(query CommitQuery, aggregatedObservation CommitPluginObservation) (ocr3types.Outcome, error) {
	nextRanges := getRangesForNextReport(
		p.log,
		query.RmnOnRampMaxSeqNums,
		aggregatedObservation.OnRampMaxSeqNums,
		aggregatedObservation.OffRampMaxSeqNums,
	)

	outcome := CommitPluginOutcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: nextRanges,
	}

	return outcome.Encode()
}

// TODO: doc
func getRangesForNextReport(
	log logger.Logger,
	rmnOnRampMaxSeqNums,
	observedOnRampMaxSeqNums,
	observedOffRampMaxSeqNums []plugintypes.SeqNumChain,
) []ChainRange {
	rangesToReport := make([]ChainRange, 0, len(observedOffRampMaxSeqNums))

	rmnOnRampMaxSeqNumsMap := seqNumChainArrayToMap(rmnOnRampMaxSeqNums)
	observedOnRampMaxSeqNumsMap := seqNumChainArrayToMap(observedOnRampMaxSeqNums)
	observedOffRampMaxSeqNumsMap := seqNumChainArrayToMap(observedOffRampMaxSeqNums)

	for chainSel, offRampMaxSeqNum := range observedOffRampMaxSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if rmnOnRampMaxSeqNum, exists := rmnOnRampMaxSeqNumsMap[chainSel]; exists {
			onRampMaxSeqNum = min(onRampMaxSeqNum, rmnOnRampMaxSeqNum)
		}

		if offRampMaxSeqNum > onRampMaxSeqNum {
			log.Warnw("Found an offRampMaxSeqNum greater than an onRampMaxSeqNum",
				"offRampMaxSeqNum", offRampMaxSeqNum,
				"onRampMaxSeqNum", onRampMaxSeqNum,
				"chainSelector", chainSel)
			continue
		} else if offRampMaxSeqNum == onRampMaxSeqNum {
			continue
		} else {
			chainRange := ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: [2]cciptypes.SeqNum{offRampMaxSeqNum, onRampMaxSeqNum},
			}
			rangesToReport = append(rangesToReport, chainRange)
		}
	}

	return rangesToReport
}

func seqNumChainArrayToMap(seqNumChains []plugintypes.SeqNumChain) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	chainToSeqNum := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, seqNumChain := range seqNumChains {
		chainToSeqNum[seqNumChain.ChainSel] = seqNumChain.SeqNum
	}

	return chainToSeqNum
}

// TODO: doc
func (p *Plugin) getConsensusObservation2(aos []types.AttributedObservation) (ConsensusObservation, error) {
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

func majorityConsensus[V, K comparable](f func(K) (int, error)) func(K, []V) (V, error) {
	return func(key K, values []V) (V, error) {
		counts := counts(values)

		var candidate V
		candidateAssigned := false
		n, err := f(key)
		if err != nil {
			return candidate, err
		}
		for elem, count := range counts {
			if count > n {
				if candidateAssigned {
					return candidate, fmt.Errorf("found multiple elems with more than f occurrences")
				}

				candidate = elem
				candidateAssigned = true
			}
		}

		if candidateAssigned {
			return candidate, nil
		}

		return candidate, fmt.Errorf("did not find an elem with more than f occurrences")
	}
}

// TODO: doc
func (p *Plugin) getConsensusObservation(aos []types.AttributedObservation) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)
	fChains := p.getFChainConsensus(aggObs.FChain)
	// TODO: doc, fix
	var fTokenChain int
	var fDestChain int

	consensusObs := ConsensusObservation{
		MerkleRoots:       p.getMerkleRootConsensus(aggObs.MerkleRoots, fChains),
		GasPrices:         p.getGasPricesConsensus(aggObs.GasPrices, fChains),
		TokenPrices:       p.getTokenPricesConsensus(aggObs.TokenPrices, fTokenChain),
		OnRampMaxSeqNums:  p.getOnRampMaxSeqNumsConsensus(aggObs.OffRampMaxSeqNums, fChains),
		OffRampMaxSeqNums: p.getOffRampMaxSeqNumsConsensus(aggObs.OffRampMaxSeqNums, fDestChain),
		FChain:            fChains,
	}

	return consensusObs, nil
}

// getFChainConsensus TODO: doc
func (p *Plugin) getFChainConsensus(aggFChains map[cciptypes.ChainSelector][]int) map[cciptypes.ChainSelector]int {
	fChainConsensus := make(map[cciptypes.ChainSelector]int)
	for chainSelector, fChainValues := range aggFChains {
		// TODO: is this the correct f to use here (p.reportingCfg.F)?
		consensusF, err := onlyValueWithMoreThanFOccurences(fChainValues, p.reportingCfg.F)
		if err != nil {
			continue
		}
		fChainConsensus[chainSelector] = consensusF
	}

	return fChainConsensus
}

// getMerkleRootConsensus TODO: doc
func (p *Plugin) getMerkleRootConsensus(
	merkleRoots map[cciptypes.ChainSelector][]cciptypes.Bytes32,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.Bytes32 {
	merkleRootConsensus := make(map[cciptypes.ChainSelector]cciptypes.Bytes32)

	for chainSelector, roots := range merkleRoots {
		if f, exists := fChains[chainSelector]; exists {
			consensusRoot, err := onlyValueWithMoreThanFOccurences(roots, f)
			if err != nil {
				// TODO: log error
				continue
			}
			merkleRootConsensus[chainSelector] = consensusRoot
		}
	}

	return merkleRootConsensus
}

// getGasPricesConsensus TODO: doc
// Explain why we use the f value we do
func (p *Plugin) getGasPricesConsensus(
	gasPrices map[cciptypes.ChainSelector][]cciptypes.BigInt,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	gasPricesConsensus := make(map[cciptypes.ChainSelector]cciptypes.BigInt)

	for chainSelector, prices := range gasPrices {
		if f, exists := fChains[chainSelector]; exists {
			if len(prices) < 2*f+1 {
				p.log.Warnw(
					"not enough gas price observations",
					"chainSelector", chainSelector,
					"prices", prices,
					"expected", 2*f+1,
					"got", len(prices),
				)
				continue
			}

			gasPricesConsensus[chainSelector] = slicelib.BigIntSortedMiddle(prices)
		} else {
			// TODO: log
		}
	}

	return gasPricesConsensus
}

// getTokenPricesConsensus TODO: doc
// Explain why we use the f value we do
func (p *Plugin) getTokenPricesConsensus(
	tokenPrices map[types.Account][]cciptypes.BigInt,
	fTokenChain int,
) map[types.Account]cciptypes.BigInt {
	tokenPricesConsensus := make(map[types.Account]cciptypes.BigInt)

	for tokenId, prices := range tokenPrices {
		if len(prices) < 2*fTokenChain+1 {
			p.log.Warnw(
				"not enough token price observations",
				"tokenId", tokenId,
				"prices", prices,
				"expected", 2*fTokenChain+1,
				"got", len(prices),
			)
			continue
		}

		tokenPricesConsensus[tokenId] = slicelib.BigIntSortedMiddle(prices)
	}

	return tokenPricesConsensus
}

// getOnRampMaxSeqNumsConsensus TODO: doc
// Explain why we use the f value we do
func (p *Plugin) getOnRampMaxSeqNumsConsensus(
	onRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	onRampMaxSeqNumsConsensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chainSelector, onRampMaxSeqNums := range onRampMaxSeqNumsByChain {
		if f, exists := fChains[chainSelector]; exists {
			if len(onRampMaxSeqNums) < 2*f+1 {
				p.log.Warnw(
					"not enough onRampMaxSeqNums observations",
					"chainSelector", chainSelector,
					"onRampMaxSeqNums", onRampMaxSeqNums,
					"expected", 2*f+1,
					"got", len(onRampMaxSeqNums),
				)
				continue
			}

			sort.Slice(onRampMaxSeqNums, func(i, j int) bool { return i > j })
			onRampMaxSeqNumsConsensus[chainSelector] = onRampMaxSeqNums[f]
		} else {
			// TODO: log
		}
	}

	return onRampMaxSeqNumsConsensus
}

// getOffRampMaxSeqNumsConsensus TODO: doc
// Explain why we use the f value we do
func (p *Plugin) getOffRampMaxSeqNumsConsensus(
	offRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fDestChain int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	offRampMaxSeqNumsConsensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chainSelector, offRampMaxSeqNums := range offRampMaxSeqNumsByChain {
		consensusSeqNum, err := onlyValueWithMoreThanFOccurences(offRampMaxSeqNums, fDestChain)
		if err != nil {
			// TODO: log
			continue
		}

		offRampMaxSeqNumsConsensus[chainSelector] = consensusSeqNum
	}

	return offRampMaxSeqNumsConsensus
}

func onlyValueWithMoreThanFOccurencesWithTag[T comparable](tag string, elems []T, f int) (T, error) {
	counts := counts(elems)
	var candidate T
	candidateAssigned := false
	for elem, count := range counts {
		if count > f {
			if candidateAssigned {
				return candidate, fmt.Errorf("%s: found multiple elems with more than f (%d) occurrences", tag, f)
			}

			candidate = elem
			candidateAssigned = true
		}
	}

	if candidateAssigned {
		return candidate, nil
	}

	return candidate, fmt.Errorf("%s: did not find an elem with more than f (%d) occurrences", tag, f)
}

// onlyValueWithMoreThanFOccurences TODO: doc
func onlyValueWithMoreThanFOccurences[T comparable](elems []T, f int) (T, error) {
	counts := counts(elems)
	var candidate T
	candidateAssigned := false
	for elem, count := range counts {
		if count > f {
			if candidateAssigned {
				return candidate, fmt.Errorf("found multiple elems with more than f occurrences")
			}

			candidate = elem
			candidateAssigned = true
		}
	}

	if candidateAssigned {
		return candidate, nil
	}

	return candidate, fmt.Errorf("did not find an elem with more than f occurrences")
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

// aggregateObservations TODO: doc
func (p *Plugin) aggregateObservations(aos []types.AttributedObservation) (CommitPluginObservation, error) {
	decodedObservations := make([]CommitPluginObservation, 0, len(aos))
	for _, ao := range aos {
		decodedObservation, err := DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			return CommitPluginObservation{}, err
		}
		decodedObservations = append(decodedObservations, decodedObservation)
	}

	fChains := fChainConsensus(decodedObservations)
	fChainDest, ok := fChains[p.cfg.DestChain]
	if !ok {
		return CommitPluginObservation{}, fmt.Errorf("missing destination chain %d in fChain config", p.cfg.DestChain)
	}

	gasPrices := gasPricesConsensus(p.log, decodedObservations, fChains)
	p.log.Debugw("gas prices consensus", "gasPrices", gasPrices)

	tokenPrices := tokenPricesConsensus(decodedObservations, fChainDest)
	p.log.Debugw("token prices consensus", "tokenPrices", tokenPrices)

	onRampMaxSeqNums := OnRampMaxSeqNumsConsensus(p.log, fChains, decodedObservations)
	offRampMaxSeqNums := OffRampMaxSeqNumsConsensus(p.log, fChainDest, decodedObservations)

	// TODO: merkle root consensus

	return CommitPluginObservation{
		MerkleRoots:       nil,
		GasPrices:         gasPrices,
		TokenPrices:       tokenPrices,
		OnRampMaxSeqNums:  onRampMaxSeqNums,
		OffRampMaxSeqNums: offRampMaxSeqNums,
		FChain:            fChains,
	}, nil
}

// fChainConsensus comes to consensus on the plugin config based on the observations.
// We cannot trust the state of a single follower, so we need to come to consensus on the config.
// TODO: (Ryan) not sure this is 100% secure
func fChainConsensus(
	observations []CommitPluginObservation, // observations from all followers
) map[cciptypes.ChainSelector]int {
	// Come to consensus on fChain.
	// Use the fChain observed by most followers for each chain.
	fChainCounts := make(map[cciptypes.ChainSelector]map[int]int) // {chain: {fChain: count}}
	for _, obs := range observations {
		for chain, fChain := range obs.FChain {
			if _, exists := fChainCounts[chain]; !exists {
				fChainCounts[chain] = make(map[int]int)
			}
			fChainCounts[chain][fChain]++
		}
	}
	consensusFChain := make(map[cciptypes.ChainSelector]int)
	for chain, counts := range fChainCounts {
		maxCount := 0
		for fChain, count := range counts {
			if count > maxCount {
				maxCount = count
				consensusFChain[chain] = fChain
			}
		}
	}

	return consensusFChain
}

// tokenPricesConsensus returns the median price for tokens that have at least 2f_chain+1 observations.
// TODO: (Ryan) why is fchain the fchain for dest? Are token prices only read by oracles that support the dest chain?
func tokenPricesConsensus(observations []CommitPluginObservation, fChainDest int) []cciptypes.TokenPrice {
	pricesPerToken := make(map[types.Account][]cciptypes.BigInt)
	for _, obs := range observations {
		for _, price := range obs.TokenPrices {
			if _, exists := pricesPerToken[price.TokenID]; !exists {
				pricesPerToken[price.TokenID] = make([]cciptypes.BigInt, 0)
			}
			pricesPerToken[price.TokenID] = append(pricesPerToken[price.TokenID], price.Price)
		}
	}

	// Keep the median
	consensusPrices := make([]cciptypes.TokenPrice, 0)
	for token, prices := range pricesPerToken {
		if len(prices) < 2*fChainDest+1 {
			continue
		}
		consensusPrices = append(consensusPrices, cciptypes.NewTokenPrice(token, slicelib.BigIntSortedMiddle(prices).Int))
	}

	sort.Slice(consensusPrices, func(i, j int) bool { return consensusPrices[i].TokenID < consensusPrices[j].TokenID })
	return consensusPrices
}

// gasPricesConsensus TODO: doc
func gasPricesConsensus(
	lggr logger.Logger, observations []CommitPluginObservation, fChains map[cciptypes.ChainSelector]int,
) []cciptypes.GasPriceChain {
	// Group the observed gas prices by chain.
	gasPricePerChain := make(map[cciptypes.ChainSelector][]cciptypes.BigInt)
	for _, obs := range observations {
		for _, gasPrice := range obs.GasPrices {
			if _, exists := gasPricePerChain[gasPrice.ChainSel]; !exists {
				gasPricePerChain[gasPrice.ChainSel] = make([]cciptypes.BigInt, 0)
			}
			gasPricePerChain[gasPrice.ChainSel] = append(gasPricePerChain[gasPrice.ChainSel], gasPrice.GasPrice)
		}
	}

	// Keep the median
	consensusGasPrices := make([]cciptypes.GasPriceChain, 0)
	for chain, gasPrices := range gasPricePerChain {
		fChain := fChains[chain]
		if len(gasPrices) < 2*fChain+1 {
			lggr.Warnw("not enough gas price observations", "chain", chain, "gasPrices", gasPrices)
			continue
		}

		consensusGasPrices = append(
			consensusGasPrices,
			cciptypes.NewGasPriceChain(slicelib.BigIntSortedMiddle(gasPrices).Int, chain),
		)
	}

	sort.Slice(
		consensusGasPrices,
		func(i, j int) bool { return consensusGasPrices[i].ChainSel < consensusGasPrices[j].ChainSel },
	)
	return consensusGasPrices
}

// OnRampMaxSeqNumsConsensus groups the observed max seq nums across all followers per chain.
// Orders the sequence numbers and selects the one at the index of destination chain fChain.
//
// For example:
//
//	seqNums: [1, 1, 1, 10, 10, 10, 10, 10, 10]
//	fChain: 4
//	result: 10
//
// Selecting seqNums[fChain] ensures:
//   - At least one honest node has seen this value, so adversary cannot bias the value lower which would cause reverts
//   - If an honest oracle reports sorted_min[f] which happens to be stale i.e. that oracle has a delayed view
//     of the chain, then the report will revert onchain but still succeed upon retry
//   - We minimize the risk of naturally hitting the error condition minSeqNum > maxSeqNum due to oracles
//     delayed views of the chain (would be an issue with taking sorted_mins[-f])
func OnRampMaxSeqNumsConsensus(
	lggr logger.Logger, fChains map[cciptypes.ChainSelector]int, observations []CommitPluginObservation,
) []plugintypes.SeqNumChain {
	observedSeqNumsPerChain := make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
	for _, obs := range observations {
		for _, maxSeqNum := range obs.OnRampMaxSeqNums {
			if _, exists := observedSeqNumsPerChain[maxSeqNum.ChainSel]; !exists {
				observedSeqNumsPerChain[maxSeqNum.ChainSel] = make([]cciptypes.SeqNum, 0)
			}
			observedSeqNumsPerChain[maxSeqNum.ChainSel] =
				append(observedSeqNumsPerChain[maxSeqNum.ChainSel], maxSeqNum.SeqNum)
		}
	}

	seqNums := make([]plugintypes.SeqNumChain, 0, len(observedSeqNumsPerChain))
	for ch, observedSeqNums := range observedSeqNumsPerChain {
		fChain := fChains[ch]
		if len(observedSeqNums) < 2*fChain+1 {
			lggr.Warnw("not enough observations for chain", "chain", ch, "observedSeqNums", observedSeqNums)
			continue
		}

		sort.Slice(observedSeqNums, func(i, j int) bool { return observedSeqNums[i] < observedSeqNums[j] })
		seqNums = append(seqNums, plugintypes.NewSeqNumChain(ch, observedSeqNums[fChain]))
	}

	sort.Slice(seqNums, func(i, j int) bool { return seqNums[i].ChainSel < seqNums[j].ChainSel })
	return seqNums
}

// OffRampMaxSeqNumsConsensus TODO: doc
func OffRampMaxSeqNumsConsensus(
	lggr logger.Logger, fChainDest int, observations []CommitPluginObservation,
) []plugintypes.SeqNumChain {
	observedSeqNumsPerChain := make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
	for _, obs := range observations {
		for _, maxSeqNum := range obs.OffRampMaxSeqNums {
			if _, exists := observedSeqNumsPerChain[maxSeqNum.ChainSel]; !exists {
				observedSeqNumsPerChain[maxSeqNum.ChainSel] = make([]cciptypes.SeqNum, 0)
			}
			observedSeqNumsPerChain[maxSeqNum.ChainSel] =
				append(observedSeqNumsPerChain[maxSeqNum.ChainSel], maxSeqNum.SeqNum)
		}
	}

	seqNums := make([]plugintypes.SeqNumChain, 0, len(observedSeqNumsPerChain))
	for ch, observedSeqNums := range observedSeqNumsPerChain {
		if len(observedSeqNums) < 2*fChainDest+1 {
			lggr.Warnw("not enough observations for chain", "chain", ch, "observedSeqNums", observedSeqNums)
			continue
		}

		sort.Slice(observedSeqNums, func(i, j int) bool { return observedSeqNums[i] < observedSeqNums[j] })
		// This isn't right, should be most occurrences, not median
		seqNums = append(seqNums, plugintypes.NewSeqNumChain(ch, observedSeqNums[fChainDest]))
	}

	sort.Slice(seqNums, func(i, j int) bool { return seqNums[i].ChainSel < seqNums[j].ChainSel })
	return seqNums
}
