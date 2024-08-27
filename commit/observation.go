package commit

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
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
		offRampNextSeqNums := p.observer.ObserveOffRampNextSeqNums(ctx)
		observation = Observation{
			// TODO: observe OnRamp max seq nums. The use of offRampNextSeqNums here effectively disables batching,
			// e.g. the ranges selected for each chain will be [x, x] (e.g. [46, 46]), which means reports will only
			// contain one message per chain. Querying the OnRamp contract requires changes to reader.CCIP, which will
			// need to be done in a future change.
			OnRampMaxSeqNums:   offRampNextSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			FChain:             p.observer.ObserveFChain(),
		}

	case BuildingReport:
		observation = Observation{
			MerkleRoots: p.observer.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			GasPrices:   p.observer.ObserveGasPrices(ctx),
			TokenPrices: p.observer.ObserveTokenPrices(ctx),
			FChain:      p.observer.ObserveFChain(),
		}

	case WaitingForReportTransmission:
		observation = Observation{
			OffRampNextSeqNums: p.observer.ObserveOffRampNextSeqNums(ctx),
			FChain:             p.observer.ObserveFChain(),
		}

	default:
		p.lggr.Errorw("Unexpected state", "state", nextState)
		return observation.Encode()
	}

	p.lggr.Infow("Observation", "observation", observation)
	return observation.Encode()
}

type Observer interface {
	// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp
	ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain

	// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
	ObserveMerkleRoots(ctx context.Context, ranges []plugintypes.ChainRange) []cciptypes.MerkleRootChain

	ObserveTokenPrices(ctx context.Context) []cciptypes.TokenPrice

	ObserveGasPrices(ctx context.Context) []cciptypes.GasPriceChain

	ObserveFChain() map[cciptypes.ChainSelector]int
}

type ObserverImpl struct {
	lggr         logger.Logger
	homeChain    reader.HomeChain
	nodeID       commontypes.OracleID
	chainSupport ChainSupport
	ccipReader   reader.CCIP
	msgHasher    cciptypes.MessageHasher
}

// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp
func (o ObserverImpl) ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	supportsDestChain, err := o.chainSupport.SupportsDestChain(o.nodeID)
	if err != nil {
		o.lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return nil
	}

	if !supportsDestChain {
		return nil
	}

	sourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}
	offRampNextSeqNums, err := o.ccipReader.NextSeqNum(ctx, sourceChains)
	if err != nil {
		o.lggr.Warnw("call to NextSeqNum failed", "err", err)
		return nil
	}

	if len(offRampNextSeqNums) != len(sourceChains) {
		o.lggr.Errorf("call to NextSeqNum returned unexpected number of seq nums, got %d, expected %d",
			len(offRampNextSeqNums), len(sourceChains))
		return nil
	}

	result := make([]plugintypes.SeqNumChain, len(sourceChains))
	for i := range sourceChains {
		result[i] = plugintypes.SeqNumChain{ChainSel: sourceChains[i], SeqNum: offRampNextSeqNums[i]}
	}

	return result
}

// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
func (o ObserverImpl) ObserveMerkleRoots(
	ctx context.Context,
	ranges []plugintypes.ChainRange,
) []cciptypes.MerkleRootChain {
	var roots []cciptypes.MerkleRootChain
	supportedChains, err := o.chainSupport.SupportedChains(o.nodeID)
	if err != nil {
		o.lggr.Warnw("call to supportedChains failed", "err", err)
		return nil
	}

	for _, chainRange := range ranges {
		if supportedChains.Contains(chainRange.ChainSel) {
			msgs, err := o.ccipReader.MsgsBetweenSeqNums(ctx, chainRange.ChainSel, chainRange.SeqNumRange)
			if err != nil {
				o.lggr.Warnw("call to MsgsBetweenSeqNums failed", "err", err)
				continue
			}
			root, err := o.computeMerkleRoot(ctx, msgs)
			if err != nil {
				o.lggr.Warnw("call to computeMerkleRoot failed", "err", err)
				continue
			}
			merkleRoot := cciptypes.MerkleRootChain{
				ChainSel:     chainRange.ChainSel,
				SeqNumsRange: chainRange.SeqNumRange,
				MerkleRoot:   root,
			}
			roots = append(roots, merkleRoot)
		}
	}

	return roots
}

// computeMerkleRoot computes the merkle root of a list of messages
func (o ObserverImpl) computeMerkleRoot(ctx context.Context, msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
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

		msgHash, err := o.msgHasher.Hash(ctx, msg)
		if err != nil {
			msgID := hex.EncodeToString(msg.Header.MessageID[:])
			o.lggr.Warnw("failed to hash message", "msg", msg, "msg_id", msgID, "err", err)
			return cciptypes.Bytes32{}, fmt.Errorf("failed to hash message with id %s: %w", msgID, err)
		}

		hashes = append(hashes, msgHash)
	}

	// TODO: Do not hard code the hash function, it should be derived from the message hasher
	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), hashes)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to construct merkle tree from %d leaves: %w", len(hashes), err)
	}

	root := tree.Root()
	o.lggr.Infow("computeMerkleRoot: Computed merkle root", "root", hex.EncodeToString(root[:]))

	return root, nil
}

func (o ObserverImpl) ObserveTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
	return []cciptypes.TokenPrice{}
}

func (o ObserverImpl) ObserveGasPrices(ctx context.Context) []cciptypes.GasPriceChain {
	return []cciptypes.GasPriceChain{}
}

func (o ObserverImpl) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := o.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		o.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

// Interface compliance check
var _ Observer = (*ObserverImpl)(nil)
