package report

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas"
)

var _ ExecReportBuilder = &execReportBuilder{}

type ExecReportBuilder interface {
	Add(report exectypes.CommitData) (exectypes.CommitData, error)
	Build() ([]cciptypes.ExecutePluginReportSingleChain, error)
}

func NewBuilder(
	ctx context.Context,
	logger logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader exectypes.TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	estimateProvider gas.EstimateProvider,
	nonces map[cciptypes.ChainSelector]map[string]uint64,
	destChainSelector cciptypes.ChainSelector,
	maxReportSizeBytes uint64,
	maxGas uint64,
) ExecReportBuilder {
	return &execReportBuilder{
		ctx:  ctx,
		lggr: logger,

		tokenDataReader:  tokenDataReader,
		encoder:          encoder,
		hasher:           hasher,
		estimateProvider: estimateProvider,
		sendersNonce:     nonces,
		expectedNonce:    make(map[cciptypes.ChainSelector]map[string]uint64),

		destChainSelector:  destChainSelector,
		maxReportSizeBytes: maxReportSizeBytes,
		maxGas:             maxGas,
	}
}

// validationMetadata contains all metadata needed to accumulate results across multiple reports and messages.
type validationMetadata struct {
	encodedSizeBytes uint64
	gas              uint64
}

func (vm validationMetadata) accumulate(other validationMetadata) validationMetadata {
	var result validationMetadata
	result.encodedSizeBytes = vm.encodedSizeBytes + other.encodedSizeBytes
	result.gas = vm.gas + other.gas
	return result
}

type execReportBuilder struct {
	ctx  context.Context // TODO: remove context from builder so that it can be pure?
	lggr logger.Logger

	// Providers
	tokenDataReader  exectypes.TokenDataReader
	encoder          cciptypes.ExecutePluginCodec
	hasher           cciptypes.MessageHasher
	estimateProvider gas.EstimateProvider
	sendersNonce     map[cciptypes.ChainSelector]map[string]uint64

	// Config
	destChainSelector  cciptypes.ChainSelector
	maxReportSizeBytes uint64
	maxGas             uint64

	// State
	accumulated validationMetadata
	// expectedNonce is used to track nonces for multiple messages from the same sender.
	expectedNonce map[cciptypes.ChainSelector]map[string]uint64

	// Result
	execReports []cciptypes.ExecutePluginReportSingleChain
}

func (b *execReportBuilder) Add(
	commitReport exectypes.CommitData,
) (exectypes.CommitData, error) {
	execReport, updatedReport, err := b.buildSingleChainReport(b.ctx, commitReport)

	// No messages fit into the report, move to next report
	if errors.Is(err, ErrEmptyReport) {
		return commitReport, nil
	}
	if err != nil {
		return commitReport, fmt.Errorf("unable to add a single chain report: %w", err)
	}

	b.execReports = append(b.execReports, execReport)

	return updatedReport, nil
}

func (b *execReportBuilder) Build() ([]cciptypes.ExecutePluginReportSingleChain, error) {
	b.lggr.Infow(
		"selected commit reports for execution report",
		"numReports", len(b.execReports),
		"sizeBytes", b.accumulated.encodedSizeBytes,
		"maxSize", b.maxReportSizeBytes)
	return b.execReports, nil
}
