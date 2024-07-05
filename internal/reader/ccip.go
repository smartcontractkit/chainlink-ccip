package reader

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"golang.org/x/sync/errgroup"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/pkg/crconsts"

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
	) ([]cciptypes.CCIPMsg, error)

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
	contractReaders map[cciptypes.ChainSelector]types.ContractReader
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter
	destChain       cciptypes.ChainSelector
}

func NewCCIPChainReader(
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
) *CCIPChainReader {
	return &CCIPChainReader{
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
	panic("implement me")
}

func (r *CCIPChainReader) ExecutedMessageRanges(
	ctx context.Context, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNumRange, error) {
	if err := r.validateReaderExistence(source, dest); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (r *CCIPChainReader) MsgsBetweenSeqNums(
	ctx context.Context, chain cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.CCIPMsg, error) {
	if err := r.validateReaderExistence(chain); err != nil {
		return nil, err
	}

	seq, err := r.contractReaders[chain].QueryKey(
		ctx,
		crconsts.ContractNameOnRamp,
		query.KeyFilter{
			Key: crconsts.EventNameCCIPSendRequested,
			Expressions: []query.Expression{
				{
					Primitive: &primitives.Comparator{
						Name: crconsts.EventAttributeSequenceNumber,
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
					BoolExpression: query.BoolExpression{},
				},
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
		&cciptypes.CCIPMsg{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query onRamp: %w", err)
	}

	msgs := make([]cciptypes.CCIPMsg, 0)
	for _, item := range seq {
		msg, ok := item.Data.(cciptypes.CCIPMsg)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to CCIPMsg", item.Data)
		}
		msgs = append(msgs, msg)
	}

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

	for chain, cfg := range sourceConfigs {
		if cfg.OnRamp == "" {
			return false, fmt.Errorf("onRamp address not found for chain %d", chain)
		}

		// Bind the onRamp contract address to the reader.
		// If the same address exists -> no-op
		// If the address is changed -> updates the address, overwrites the existing one
		// If the contract not binded -> binds to the new address
		if err := r.contractReaders[chain].Bind(ctx, []types.BoundContract{
			{
				Address: cfg.OnRamp,
				Name:    crconsts.ContractNameOnRamp,
				Pending: false,
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

	eg := new(errgroup.Group)
	for chainSel := range r.contractReaders {
		eg.Go(func() error {
			resp := sourceChainConfig{}
			err := r.contractReaders[r.destChain].GetLatestValue(
				ctx,
				crconsts.ContractNameOffRamp,
				crconsts.FunctionNameGetSourceChainConfig,
				map[string]any{
					"sourceChainSelector": chainSel,
				},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get source chain config: %w", err)
			}
			res[chainSel] = resp
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

type sourceChainConfig struct {
	OnRamp   string `json:"onRamp"`
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
