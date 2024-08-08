package commitrmnocb

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func (p *Plugin) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// Across all chains we require at least 2F+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (p *Plugin) Observation(
	ctx context.Context, outCtx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)

	observation := CommitPluginObservation{}
	switch nextState {
	case SelectingRangesForReport:
		offRampNextSeqNums := p.ObserveOffRampNextSeqNums(ctx)
		observation = CommitPluginObservation{
			OnRampMaxSeqNums:   offRampNextSeqNums, // TODO: change
			OffRampNextSeqNums: offRampNextSeqNums,
			FChain:             p.ObserveFChain(),
		}

	case BuildingReport:
		observation = CommitPluginObservation{
			MerkleRoots: p.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			GasPrices:   p.ObserveGasPrices(ctx),
			TokenPrices: p.ObserveTokenPrices(ctx),
			FChain:      p.ObserveFChain(),
		}

	case WaitingForReportTransmission:
		observation = CommitPluginObservation{
			OffRampNextSeqNums: p.ObserveOffRampNextSeqNums(ctx),
			FChain:             p.ObserveFChain(),
		}

	default:
		p.lggr.Warnw("Unexpected state", "state", nextState)
		return types.Observation{}, nil
	}

	p.lggr.Infow("Observation", "observation", observation)
	return observation.Encode()
}

// ObserveOnRampMaxSeqNums Simply add NewMsgScanBatchSize to the offRampNextSeqNums
// TODO: read from the source chain OnRamps to get their OnRampMaxSeqNums
func (p *Plugin) ObserveOnRampMaxSeqNums(offRampNextSeqNums []plugintypes.SeqNumChain) []plugintypes.SeqNumChain {
	onRampMaxSeqNums := make([]plugintypes.SeqNumChain, len(offRampNextSeqNums))
	copy(onRampMaxSeqNums, offRampNextSeqNums)

	for i := range onRampMaxSeqNums {
		onRampMaxSeqNums[i] = plugintypes.NewSeqNumChain(
			onRampMaxSeqNums[i].ChainSel,
			onRampMaxSeqNums[i].SeqNum+cciptypes.SeqNum(p.cfg.NewMsgScanBatchSize),
		)
	}

	return onRampMaxSeqNums
}

// ObserveOffRampNextSeqNums TODO: impl
func (p *Plugin) ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	supportsDestChain, err := p.supportsDestChain(p.nodeID)
	if err != nil {
		p.lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return nil
	}

	if supportsDestChain {
		sourceChains := p.knownSourceChainsSlice()
		offRampNextSeqNums, err := p.ccipReader.NextSeqNum(ctx, sourceChains)
		if err != nil {
			p.lggr.Warnw("call to NextSeqNum failed", "err", err)
			return nil
		}

		if len(offRampNextSeqNums) != len(sourceChains) {
			p.lggr.Warnf("call to NextSeqNum returned unexpected number of seq nums, got %d, expected %d",
				len(offRampNextSeqNums), len(sourceChains))
			return nil
		}

		result := make([]plugintypes.SeqNumChain, len(sourceChains))
		for i := range sourceChains {
			result[i] = plugintypes.SeqNumChain{ChainSel: sourceChains[i], SeqNum: offRampNextSeqNums[i]}
		}

		return result
	}

	return nil
}

// ObserveMerkleRoots TODO: doc
func (p *Plugin) ObserveMerkleRoots(ctx context.Context, ranges []ChainRange) []cciptypes.MerkleRootChain {
	roots := make([]cciptypes.MerkleRootChain, 0)
	supportedChains, err := p.supportedChains(p.nodeID)
	if err != nil {
		p.lggr.Warnw("call to supportedChains failed", "err", err)
		return nil
	}

	for _, chainRange := range ranges {
		if supportedChains.Contains(chainRange.ChainSel) {
			msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, chainRange.ChainSel, chainRange.SeqNumRange)
			if err != nil {
				p.lggr.Warnw("call to MsgsBetweenSeqNums failed", "err", err)
			} else {
				root, err := computeMerkleRoot(msgs)
				if err != nil {
					p.lggr.Warnw("call to computeMerkleRoot failed", "err", err)
				} else {
					merkleRoot := cciptypes.MerkleRootChain{
						ChainSel:     chainRange.ChainSel,
						SeqNumsRange: chainRange.SeqNumRange,
						MerkleRoot:   root,
					}
					roots = append(roots, merkleRoot)
				}
			}
		}
	}

	return roots
}

// computeMerkleRoot Compute the merkle root of a list of messages
func computeMerkleRoot(msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
	msgSeqNumToHash := make(map[cciptypes.SeqNum]cciptypes.Bytes32)
	seqNums := make([]cciptypes.SeqNum, 0)

	for _, msg := range msgs {
		seqNum := msg.Header.SequenceNumber
		seqNums = append(seqNums, seqNum)
		msgSeqNumToHash[seqNum] = msg.Header.MsgHash
	}

	sort.Slice(seqNums, func(i, j int) bool { return seqNums[i] < seqNums[j] })

	// Assert there are no gaps in the seq num range
	if len(seqNums) >= 2 {
		for i := 1; i < len(seqNums); i++ {
			if seqNums[i] != seqNums[i-1]+1 {
				return [32]byte{}, fmt.Errorf("found non-consecutive sequence numbers when computing merkle root, "+
					"gap between seq nums %d and %d, messages: %v", seqNums[i-1], seqNums[i], msgs)
			}
		}
	}

	treeLeaves := make([][32]byte, 0)
	for _, seqNum := range seqNums {
		msgHash, ok := msgSeqNumToHash[seqNum]
		if !ok {
			return [32]byte{}, fmt.Errorf("msg hash not found for seq num %d", seqNum)
		}
		treeLeaves = append(treeLeaves, msgHash)
	}

	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), treeLeaves)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to construct merkle tree from %d leaves: %w", len(treeLeaves), err)
	}

	return tree.Root(), nil
}

// ObserveGasPrices TODO: doc
func (p *Plugin) ObserveGasPrices(ctx context.Context) []cciptypes.GasPriceChain {
	// TODO: Should this be sourceChains or supportedChains?
	chains := p.knownSourceChainsSlice()
	if len(chains) == 0 {
		return []cciptypes.GasPriceChain{}
	}

	gasPrices, err := p.ccipReader.GasPrices(ctx, chains)
	if err != nil {
		p.lggr.Warnw("failed to get gas prices", "err", err)
		return []cciptypes.GasPriceChain{}
	}

	if len(gasPrices) != len(chains) {
		p.lggr.Warnw(
			"gas prices length mismatch",
			"len(gasPrices)", len(gasPrices),
			"len(chains)", len(chains),
		)
		return []cciptypes.GasPriceChain{}
	}

	gasPricesGwei := make([]cciptypes.GasPriceChain, 0, len(chains))
	for i, chain := range chains {
		gasPricesGwei = append(gasPricesGwei, cciptypes.NewGasPriceChain(gasPrices[i].Int, chain))
	}

	return gasPricesGwei
}

// ObserveTokenPrices TODO: doc
func (p *Plugin) ObserveTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
	tokenPrices, err := p.observeTokenPricesHelper(ctx)
	if err != nil {
		p.lggr.Warnw("call to ObserveTokenPrices failed", "err", err)
	}
	return tokenPrices
}

// ObserveTokenPricesHelper TODO: doc
func (p *Plugin) observeTokenPricesHelper(ctx context.Context) ([]cciptypes.TokenPrice, error) {
	if supportTPChain, err := p.supportsTokenPriceChain(); err == nil && supportTPChain {
		tokens := maps.Keys(p.cfg.OffchainConfig.PriceSources)

		tokenPrices, err := p.tokenPricesReader.GetTokenPricesUSD(ctx, tokens)
		if err != nil {
			return nil, fmt.Errorf("get token prices: %w", err)
		}

		if len(tokenPrices) != len(tokens) {
			return nil, fmt.Errorf("internal critical error token prices length mismatch: got %d, want %d",
				len(tokenPrices), len(tokens))
		}

		tokenPricesUSD := make([]cciptypes.TokenPrice, 0, len(tokens))
		for i, token := range tokens {
			tokenPricesUSD = append(tokenPricesUSD, cciptypes.NewTokenPrice(token, tokenPrices[i]))
		}

		return tokenPricesUSD, nil
	}

	return nil, nil
}

func (p *Plugin) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
