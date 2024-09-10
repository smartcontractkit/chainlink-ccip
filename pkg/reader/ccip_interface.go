package reader

import (
	"context"
	"errors"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

var (
	ErrContractReaderNotFound = errors.New("contract reader not found")
	ErrContractWriterNotFound = errors.New("contract writer not found")
)

func NewCCIPChainReader(
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
) CCIPReader {
	return newCCIPChainReaderInternal(
		lggr,
		contractReaders,
		contractWriters,
		destChain,
		offrampAddress,
	)
}

// NewCCIPReaderWithExtendedContractReaders can be used when you want to directly provide contractreader.Extended
func NewCCIPReaderWithExtendedContractReaders(
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
) CCIPReader {
	cr := newCCIPChainReaderInternal(lggr, nil, contractWriters, destChain, offrampAddress)
	for ch, extendedCr := range contractReaders {
		cr.WithExtendedContractReader(ch, extendedCr)
	}
	return cr
}

type CCIPReader interface {
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

	// GetExpectedNextSequenceNumber returns the next sequence number to be used
	// in the onramp.
	GetExpectedNextSequenceNumber(
		ctx context.Context,
		sourceChainSelector, destChainSelector cciptypes.ChainSelector,
	) (cciptypes.SeqNum, error)

	// NextSeqNum reads the destination chain.
	// Returns the next expected sequence number for each one of the provided chains.
	// TODO: if destination was a parameter, this could be a capability reused across plugin instances.
	NextSeqNum(ctx context.Context, chains []cciptypes.ChainSelector) (seqNum []cciptypes.SeqNum, err error)

	// GetContractAddress returns the contract address that is registered for the provided contract name and chain.
	GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error)

	// Nonces fetches all nonces for the provided selector/address pairs. Addresses are a string encoded raw address,
	// it must be encoding according to the destination chain requirements with typeconv.AddressBytesToString.
	Nonces(
		ctx context.Context,
		source, dest cciptypes.ChainSelector,
		addresses []string,
	) (map[string]uint64, error)

	// GasPrices reads the provided chains gas prices.
	GasPrices(ctx context.Context, chains []cciptypes.ChainSelector) ([]cciptypes.BigInt, error)

	// Sync can be used to perform frequent syncing operations inside the reader implementation.
	// Returns a bool indicating whether something was updated.
	Sync(ctx context.Context) (bool, error)

	// Close closes any open resources.
	Close(ctx context.Context) error
}
