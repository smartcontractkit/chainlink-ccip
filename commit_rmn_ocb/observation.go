package commitrmnocb

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

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

	switch nextState {
	case SelectingRangesForReport:
		return CommitPluginObservation{
			OnRampMaxSeqNums:  p.ObserveOnRampMaxSeqNums(),
			OffRampMaxSeqNums: p.ObserveOffRampMaxSeqNums(),
			FChain:            p.ObserveFChain(),
		}.Encode()

	case BuildingReport:
		return CommitPluginObservation{
			MerkleRoots: p.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			GasPrices:   p.ObserveGasPrices(ctx),
			TokenPrices: p.ObserveTokenPrices(ctx),
			FChain:      p.ObserveFChain(),
		}.Encode()

	case WaitingForReportTransmission:
		return CommitPluginObservation{
			OffRampMaxSeqNums: p.ObserveOffRampMaxSeqNums(),
			FChain:            p.ObserveFChain(),
		}.Encode()

	default:
		p.log.Warnw("Unexpected state", "state", nextState)
		return types.Observation{}, nil
	}
}

// ObserveOnRampMaxSeqNums TODO: doc
func (p *Plugin) ObserveOnRampMaxSeqNums() []plugintypes.SeqNumChain {
	onRampMaxSeqNums, err := p.onChain.GetOnRampMaxSeqNums()
	if err != nil {
		p.log.Warnw("call to GetOnRampMaxSeqNums failed", "err", err)
	}

	return onRampMaxSeqNums
}

// ObserveOffRampMaxSeqNums TODO: doc
func (p *Plugin) ObserveOffRampMaxSeqNums() []plugintypes.SeqNumChain {
	offRampMaxSeqNums, err := p.onChain.GetOffRampMaxSeqNums()
	if err != nil {
		p.log.Warnw("call to GetOffRampMaxSeqNums failed", "err", err)
	}

	return offRampMaxSeqNums
}

// ObserveMerkleRoots TODO: doc
func (p *Plugin) ObserveMerkleRoots(ctx context.Context, ranges []ChainRange) []MerkleRoot {
	roots := make([]MerkleRoot, len(ranges))
	for _, chainRange := range ranges {
		msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, chainRange.ChainSel, chainRange.SeqNumRange)
		if err != nil {
			p.log.Warnw("call to MsgsBetweenSeqNums failed", "err", err)
		} else {
			root, err := computeMerkleRoot(msgs)
			if err != nil {
				p.log.Warnw("call to computeMerkleRoot failed", "err", err)
			} else {
				merkleRoot := MerkleRoot{
					ChainSel:    chainRange.ChainSel,
					SeqNumRange: chainRange.SeqNumRange,
					RootHash:    root,
				}
				roots = append(roots, merkleRoot)
			}
		}
	}

	return roots
}

// computeMerkleRoot Compute the merkle root of a list of messages
func computeMerkleRoot(msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
	msgSeqNumToHash := make(map[cciptypes.SeqNum]cciptypes.Bytes32)
	seqNums := make([]cciptypes.SeqNum, len(msgs))

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
	chains := p.sourceChains()
	if len(chains) == 0 {
		return []cciptypes.GasPriceChain{}
	}

	gasPrices, err := p.ccipReader.GasPrices(ctx, chains)
	if err != nil {
		p.log.Warnw("failed to get gas prices", "err", err)
		return []cciptypes.GasPriceChain{}
	}

	if len(gasPrices) != len(chains) {
		p.log.Warnw(
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
		p.log.Warnw("call to ObserveTokenPrices failed", "err", err)
	}
	return tokenPrices
}

// ObserveTokenPricesHelper TODO: doc
func (p *Plugin) observeTokenPricesHelper(ctx context.Context) ([]cciptypes.TokenPrice, error) {
	if p.cfg.TokenPricesObserver {
		tokenPrices, err := p.tokenPricesReader.GetTokenPricesUSD(ctx, p.cfg.PricedTokens)
		if err != nil {
			return nil, err
		}

		if len(tokenPrices) != len(p.cfg.PricedTokens) {
			return nil, fmt.Errorf("token prices length mismatch: got %d, expected %d",
				len(tokenPrices), len(p.cfg.PricedTokens))
		}

		tokenPricesUSD := make([]cciptypes.TokenPrice, 0, len(p.cfg.PricedTokens))
		for i, token := range p.cfg.PricedTokens {
			tokenPricesUSD = append(tokenPricesUSD, cciptypes.NewTokenPrice(token, tokenPrices[i]))
		}

		return tokenPricesUSD, nil
	}

	return nil, nil
}

func (p *Plugin) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.log.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
