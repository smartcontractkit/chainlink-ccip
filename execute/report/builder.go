package report

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var _ ExecReportBuilder = &execReportBuilder{}

type ExecReportBuilder interface {
	Add(ctx context.Context, report exectypes.CommitData) (exectypes.CommitData, error)
	Build() ([]cciptypes.ExecutePluginReportSingleChain, []exectypes.CommitData, error)
}

// Option that can be passed to the builder.
type Option func(erb *execReportBuilder)

// WithMaxGas limits how much gas can be used during execution.
func WithMaxGas(maxGas uint64) func(*execReportBuilder) {
	return func(erb *execReportBuilder) {
		erb.maxGas = maxGas
	}
}

// WithMaxReportSizeBytes configures the maximum report size.
func WithMaxReportSizeBytes(maxReportSizeBytes uint64) Option {
	return func(erb *execReportBuilder) {
		erb.maxReportSizeBytes = maxReportSizeBytes
	}
}

// WithMaxMessages configures the number of messages allowed to be in a report.
func WithMaxMessages(maxMessages uint64) Option {
	return func(erb *execReportBuilder) {
		erb.maxMessages = maxMessages
	}
}

// WithMaxSingleChainReports configures the number of reports when building the final result.
func WithMaxSingleChainReports(maxSingleChainReports uint64) Option {
	return func(erb *execReportBuilder) {
		erb.maxSingleChainReports = maxSingleChainReports
	}
}

// WithExtraMessageCheck adds additional message checks to the default ones.
func WithExtraMessageCheck(check Check) Option {
	return func(erb *execReportBuilder) {
		erb.checks = append(erb.checks, check)
	}
}

func newBuilderInternal(
	logger logger.Logger,
	hasher cciptypes.MessageHasher,
	encoder cciptypes.ExecutePluginCodec,
	estimateProvider cciptypes.EstimateProvider,
	destChainSelector cciptypes.ChainSelector,
	addressCodec cciptypes.AddressCodec,
	options ...Option,
) *execReportBuilder {
	defaultChecks := []Check{
		CheckIfPseudoDeleted(),
		CheckAlreadyExecuted(),
		CheckTokenData(),
	}

	builder := &execReportBuilder{
		lggr:              logger,
		checks:            defaultChecks,
		encoder:           encoder,
		hasher:            hasher,
		estimateProvider:  estimateProvider,
		destChainSelector: destChainSelector,
		addressCodec:      addressCodec,
	}

	for _, option := range options {
		if option != nil {
			option(builder)
		}
	}

	return builder
}

// NewBuilder constructs the report builder.
func NewBuilder(
	logger logger.Logger,
	hasher cciptypes.MessageHasher,
	encoder cciptypes.ExecutePluginCodec,
	estimateProvider cciptypes.EstimateProvider,
	destChainSelector cciptypes.ChainSelector,
	addressCodec cciptypes.AddressCodec,
	options ...Option,
) ExecReportBuilder {
	return newBuilderInternal(logger, hasher, encoder, estimateProvider, destChainSelector, addressCodec, options...)
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
	lggr logger.Logger

	// Providers
	encoder          cciptypes.ExecutePluginCodec
	hasher           cciptypes.MessageHasher
	estimateProvider cciptypes.EstimateProvider
	addressCodec     cciptypes.AddressCodec

	// Config
	checks                []Check
	destChainSelector     cciptypes.ChainSelector
	maxReportSizeBytes    uint64
	maxGas                uint64
	maxMessages           uint64
	maxSingleChainReports uint64

	// State
	accumulated validationMetadata

	// Result
	execReports   []cciptypes.ExecutePluginReportSingleChain
	commitReports []exectypes.CommitData
}

// Add an exec report for as many messages as possible in the given commit report.
// The commit report with updated metadata is returned, it reflects which messages
// were selected for the exec report.
func (b *execReportBuilder) Add(
	ctx context.Context,
	commitReport exectypes.CommitData,
) (exectypes.CommitData, error) {
	execReport, updatedReport, err := b.buildSingleChainReport(ctx, commitReport)

	// No messages fit into the report, move to next report
	if errors.Is(err, ErrEmptyReport) {
		return commitReport, nil
	}
	if err != nil {
		return commitReport, fmt.Errorf("unable to add a single chain report: %w", err)
	}

	b.execReports = append(b.execReports, execReport)
	b.commitReports = append(b.commitReports, updatedReport)

	return updatedReport, nil
}

func (b *execReportBuilder) Build() (
	[]cciptypes.ExecutePluginReportSingleChain, []exectypes.CommitData, error,
) {
	if len(b.execReports) != len(b.commitReports) {
		return nil, nil, fmt.Errorf(
			"expected the same number of exec and commit reports, got %d and %d",
			len(b.execReports),
			len(b.commitReports),
		)
	}

	// Check if limiting is required.
	if b.maxSingleChainReports != 0 && uint64(len(b.execReports)) > b.maxSingleChainReports {
		b.lggr.Infof(
			"limiting number of reports to maxReports from %d to %d",
			len(b.execReports),
			b.maxSingleChainReports,
		)
		b.execReports = b.execReports[:b.maxSingleChainReports]
		b.commitReports = b.commitReports[:b.maxSingleChainReports]
	}

	b.lggr.Infow(
		"selected commit reports for execution report",
		"numReports", len(b.execReports),
		"sizeBytes", b.accumulated.encodedSizeBytes,
		"maxSize", b.maxReportSizeBytes)
	return b.execReports, b.commitReports, nil
}
