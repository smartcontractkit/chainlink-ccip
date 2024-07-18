package reader

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

type CCIP interface {
	// CommitReportsGTETimestamp reads the requested chain starting at a given timestamp
	// and finds all ReportAccepted up to the provided limit.
	CommitReportsGTETimestamp(
		ctx context.Context,
		dest cciptypes.ChainSelector,
		ts time.Time,
		limit int,
	) ([]plugintypes.CommitPluginReportWithMeta, error)

	// ExecutedMessageRanges reads the destination chain and finds which messages are executed.
	// A slice of sequence number ranges is returned to express which messages are executed.
	ExecutedMessageRanges(
		ctx context.Context,
		source, dest cciptypes.ChainSelector,
		seqNumRange cciptypes.SeqNumRange,
	) ([]cciptypes.SeqNumRange, error)

	// MsgsBetweenSeqNums reads the provided chains.
	// Finds and returns ccip messages submitted between the provided sequence numbers.
	// Messages are sorted ascending based on their timestamp and limited up to the provided limit.
	MsgsBetweenSeqNums(
		ctx context.Context,
		chain cciptypes.ChainSelector,
		seqNumRange cciptypes.SeqNumRange,
	) ([]cciptypes.Message, error)

	// NextSeqNum reads the destination chain.
	// Returns the next expected sequence number for each one of the provided chains.
	// TODO: if destination was a parameter, this could be a capability reused across plugin instances.
	NextSeqNum(ctx context.Context, chains []cciptypes.ChainSelector) (seqNum []cciptypes.SeqNum, err error)

	// GasPrices reads the provided chains gas prices.
	GasPrices(ctx context.Context, chains []cciptypes.ChainSelector) ([]cciptypes.BigInt, error)

	// Sync can be used to perform frequent syncing operations inside the reader implementation.
	// Returns a bool indicating whether something was updated.
	Sync(ctx context.Context) (bool, error)

	// Close closes any open resources.
	Close(ctx context.Context) error
}

var (
	ErrContractReaderNotFound = errors.New("contract reader not found")
	ErrContractWriterNotFound = errors.New("contract writer not found")
)

// TODO: unit test the implementation when the actual contract reader and writer interfaces are finalized and mocks
// can be generated.
type CCIPChainReader struct {
	lggr            logger.Logger
	contractReaders map[cciptypes.ChainSelector]types.ContractReader
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter
	destChain       cciptypes.ChainSelector
}

func NewCCIPChainReader(
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
) *CCIPChainReader {
	return &CCIPChainReader{
		lggr:            lggr,
		contractReaders: contractReaders,
		contractWriters: contractWriters,
		destChain:       destChain,
	}
}

func (r *CCIPChainReader) CommitReportsGTETimestamp(
	ctx context.Context, dest cciptypes.ChainSelector, ts time.Time, limit int,
) ([]plugintypes.CommitPluginReportWithMeta, error) {
	if err := r.validateReaderExistence(dest); err != nil {
		return nil, err
	}

	dataTyp := cciptypes.CommitPluginReport{}
	iter, err := r.contractReaders[dest].QueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameCommitReportAccepted,
			Expressions: []query.Expression{
				{
					Primitive: &primitives.Timestamp{
						Timestamp: uint64(ts.Unix()),
						Operator:  primitives.Gte,
					},
				},
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortByTimestamp(query.Asc)},
			Limit:  query.Limit{Count: uint64(limit)},
		},
		&dataTyp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}

	reports := make([]plugintypes.CommitPluginReportWithMeta, 0)
	for _, item := range iter {
		report, is := (item.Data).(cciptypes.CommitPluginReport)
		if !is {
			return nil, fmt.Errorf("unexpected type %T while expecting a commit report", item)
		}

		blockNum, err := strconv.ParseUint(item.Head.Identifier, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse block number %s: %w", item.Head.Identifier, err)
		}

		reports = append(reports, plugintypes.CommitPluginReportWithMeta{
			Report:    report,
			Timestamp: time.Unix(int64(item.Timestamp), 0),
			BlockNum:  blockNum,
		})
	}

	return reports, nil
}

func (r *CCIPChainReader) ExecutedMessageRanges(
	ctx context.Context, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNumRange, error) {
	if err := r.validateReaderExistence(dest); err != nil {
		return nil, err
	}

	type executionStateChangedEvent struct {
		sourceChainSelector cciptypes.ChainSelector
		sequenceNumber      cciptypes.SeqNum
		state               uint8
	}

	dataTyp := executionStateChangedEvent{}

	iter, err := r.contractReaders[dest].QueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameExecutionStateChanged,
			Expressions: []query.Expression{
				{
					// sequence numbers inside the range
					Primitive: &primitives.Comparator{
						Name: consts.EventAttributeSequenceNumber,
						ValueComparators: []primitives.ValueComparator{
							{
								Value:    seqNumRange.Start().String(),
								Operator: primitives.Gte,
							},
							{
								Value:    seqNumRange.End().String(),
								Operator: primitives.Lte,
							},
						},
					},
				},
				{
					// source chain
					Primitive: &primitives.Comparator{
						Name: consts.EventAttributeSourceChain,
						ValueComparators: []primitives.ValueComparator{
							{
								Value:    source.String(),
								Operator: primitives.Eq,
							},
						},
					},
				},
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortByTimestamp(query.Asc)},
		},
		&dataTyp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}

	executed := make([]cciptypes.SeqNumRange, 0)
	for _, item := range iter {
		stateChange, ok := item.Data.(*executionStateChangedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %T to executionStateChangedEvent", item.Data)
		}
		if stateChange.sourceChainSelector != source {
			return nil, fmt.Errorf("wrong cr query, unexpected source chain %d", stateChange.sourceChainSelector)
		}
		if stateChange.sequenceNumber < seqNumRange.Start() || stateChange.sequenceNumber > seqNumRange.End() {
			return nil, fmt.Errorf("wrong cr query, unexpected sequence number %d", stateChange.sequenceNumber)
		}
		if stateChange.state <= 1 {
			r.lggr.Debugw("execution state change status is %d, skipped",
				"seqNum", stateChange.sequenceNumber, "state", stateChange.state)
			continue
		}

		executed = append(executed, cciptypes.NewSeqNumRange(stateChange.sequenceNumber, stateChange.sequenceNumber))
	}

	return executed, nil
}

func (r *CCIPChainReader) MsgsBetweenSeqNums(
	ctx context.Context, sourceChainSelector cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	if err := r.validateReaderExistence(sourceChainSelector); err != nil {
		return nil, err
	}

	seq, err := r.contractReaders[sourceChainSelector].QueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPSendRequested,
			Expressions: []query.Expression{
				{
					Primitive: &primitives.Comparator{
						Name: consts.EventAttributeSequenceNumber,
						ValueComparators: []primitives.ValueComparator{
							{
								Value:    seqNumRange.Start().String(),
								Operator: primitives.Gte,
							},
							{
								Value:    seqNumRange.End().String(),
								Operator: primitives.Lte,
							},
						},
					},
				},
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{
				query.NewSortByTimestamp(query.Asc),
			},
			Limit: query.Limit{
				Count: uint64(seqNumRange.End() - seqNumRange.Start() + 1),
			},
		},
		&cciptypes.Message{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query onRamp: %w", err)
	}

	r.lggr.Infow("queried messages between sequence numbers",
		"numMsgs", len(seq),
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String(),
	)

	msgs := make([]cciptypes.Message, 0)
	for _, item := range seq {
		msg, ok := item.Data.(*cciptypes.Message)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}

		msgs = append(msgs, *msg)
	}

	r.lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

func (r *CCIPChainReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) ([]cciptypes.SeqNum, error) {
	cfgs, err := r.getSourceChainsConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("get source chains config: %w", err)
	}

	res := make([]cciptypes.SeqNum, 0, len(chains))
	for _, chain := range chains {
		cfg, exists := cfgs[chain]
		if !exists {
			return nil, fmt.Errorf("source chain config not found for chain %d", chain)
		}
		if cfg.MinSeqNr == 0 {
			return nil, fmt.Errorf("minSeqNr not found for chain %d", chain)
		}
		res = append(res, cciptypes.SeqNum(cfg.MinSeqNr))
	}

	return res, err
}

func (r *CCIPChainReader) GasPrices(ctx context.Context, chains []cciptypes.ChainSelector) ([]cciptypes.BigInt, error) {
	if err := r.validateWriterExistence(chains...); err != nil {
		return nil, err
	}

	eg := new(errgroup.Group)
	gasPrices := make([]cciptypes.BigInt, len(chains))
	for i, chain := range chains {
		i, chain := i, chain
		eg.Go(func() error {
			gasPrice, err := r.contractWriters[chain].GetFeeComponents(ctx)
			if err != nil {
				return fmt.Errorf("failed to get gas price: %w", err)
			}
			gasPrices[i] = cciptypes.NewBigInt(gasPrice.ExecutionFee)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return gasPrices, nil
}

func (r *CCIPChainReader) Sync(ctx context.Context) (bool, error) {
	sourceConfigs, err := r.getSourceChainsConfig(ctx)
	if err != nil {
		return false, fmt.Errorf("get onramps: %w", err)
	}

	r.lggr.Infow("got source chain configs", "onramps", func() []string {
		var r []string
		for chainSelector, scc := range sourceConfigs {
			r = append(r, typeconv.AddressBytesToString(scc.OnRamp, uint64(chainSelector)))
		}
		return r
	}())

	for chain, cfg := range sourceConfigs {
		if len(cfg.OnRamp) == 0 {
			return false, fmt.Errorf("onRamp address not found for chain %d", chain)
		}

		// Bind the onRamp contract address to the reader.
		// If the same address exists -> no-op
		// If the address is changed -> updates the address, overwrites the existing one
		// If the contract not binded -> binds to the new address
		if err := r.contractReaders[chain].Bind(ctx, []types.BoundContract{
			{
				Address: typeconv.AddressBytesToString(cfg.OnRamp, uint64(chain)),
				Name:    consts.ContractNameOnRamp,
			},
		}); err != nil {
			return false, fmt.Errorf("bind onRamp: %w", err)
		}
	}

	return true, nil
}

func (r *CCIPChainReader) Close(ctx context.Context) error {
	return nil
}

// getSourceChainsConfig returns the offRamp contract's source chain configurations for each supported source chain.
func (r *CCIPChainReader) getSourceChainsConfig(
	ctx context.Context) (map[cciptypes.ChainSelector]sourceChainConfig, error) {
	if err := r.validateReaderExistence(r.destChain); err != nil {
		return nil, err
	}

	res := make(map[cciptypes.ChainSelector]sourceChainConfig)
	mu := new(sync.Mutex)

	eg := new(errgroup.Group)
	for chainSel := range r.contractReaders {
		if chainSel == r.destChain {
			continue
		}

		chainSel := chainSel
		eg.Go(func() error {
			resp := sourceChainConfig{}
			err := r.contractReaders[r.destChain].GetLatestValue(
				ctx,
				consts.ContractNameOffRamp,
				consts.MethodNameGetSourceChainConfig,
				primitives.Finalized,
				map[string]any{
					"sourceChainSelector": chainSel,
				},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get source chain config: %w", err)
			}
			mu.Lock()
			res[chainSel] = resp
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

type sourceChainConfig struct {
	OnRamp   []byte `json:"onRamp"`
	MinSeqNr uint64 `json:"minSeqNr"`
}

func (r *CCIPChainReader) validateReaderExistence(chains ...cciptypes.ChainSelector) error {
	for _, ch := range chains {
		_, exists := r.contractReaders[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrContractReaderNotFound)
		}
	}
	return nil
}

func (r *CCIPChainReader) validateWriterExistence(chains ...cciptypes.ChainSelector) error {
	for _, ch := range chains {
		_, exists := r.contractWriters[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrContractWriterNotFound)
		}
	}
	return nil
}

// Interface compliance check
var _ CCIP = (*CCIPChainReader)(nil)
