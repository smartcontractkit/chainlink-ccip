package commitrmnocb

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func (p *Plugin) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// Across all chains we require at least 2F+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (p *Plugin) Observation(
	ctx context.Context, outCtx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)

	observation := Observation{}
	switch nextState {
	case SelectingRangesForReport:
		offRampNextSeqNums := p.ObserveOffRampNextSeqNums(ctx)
		observation = Observation{
			OnRampMaxSeqNums:   offRampNextSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			FChain:             p.ObserveFChain(),
		}

	case BuildingReport:
		observation = Observation{
			MerkleRoots: p.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			GasPrices:   []cciptypes.GasPriceChain{},
			TokenPrices: []cciptypes.TokenPrice{},
			FChain:      p.ObserveFChain(),
		}

	case WaitingForReportTransmission:
		observation = Observation{
			OffRampNextSeqNums: p.ObserveOffRampNextSeqNums(ctx),
			FChain:             p.ObserveFChain(),
		}

	default:
		p.lggr.Warnw("Unexpected state", "state", nextState)
		return observation.Encode()
	}

	p.lggr.Infow("Observation", "observation", observation)
	return observation.Encode()
}

// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp
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

// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
func (p *Plugin) ObserveMerkleRoots(ctx context.Context, ranges []ChainRange) []cciptypes.MerkleRootChain {
	var roots []cciptypes.MerkleRootChain
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
				root, err := p.computeMerkleRoot(ctx, msgs)
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

// computeMerkleRoot computes the merkle root of a list of messages
func (p *Plugin) computeMerkleRoot(ctx context.Context, msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
	var hashes [][32]byte
	sort.Slice(msgs, func(i, j int) bool { return msgs[i].Header.SequenceNumber < msgs[j].Header.SequenceNumber })

	for i, msg := range msgs {
		// Assert there are no sequence number gaps in msgs
		if i > 0 {
			if msg.Header.SequenceNumber != msgs[i-1].Header.SequenceNumber+1 {
				return [32]byte{}, fmt.Errorf("found non-consecutive sequence numbers when computing merkle root, "+
					"gap between sequence nums %d and %d, messages: %v", msgs[i-1].Header.SequenceNumber,
					msg.Header.SequenceNumber, msgs)
			}
		}

		msgHash, err := p.msgHasher.Hash(ctx, msg)
		if err != nil {
			msgID := hex.EncodeToString(msg.Header.MessageID[:])
			p.lggr.Warnw("failed to hash message", "msg", msg, "msg_id", msgID, "err", err)
			return cciptypes.Bytes32{}, fmt.Errorf("failed to hash message with id %s: %w", msgID, err)
		}

		hashes = append(hashes, msgHash)
	}

	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), hashes)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to construct merkle tree from %d leaves: %w", len(hashes), err)
	}

	root := tree.Root()
	p.lggr.Infow("computeMerkleRoot: Computed merkle root", "root", hex.EncodeToString(root[:]))

	return root, nil
}

func (p *Plugin) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
