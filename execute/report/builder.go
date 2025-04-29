package report

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

var _ ExecReportBuilder = &execReportBuilder{}

type ExecReportBuilder interface {
	Add(ctx context.Context, report exectypes.CommitData) (exectypes.CommitData, error)
	Build() ([]cciptypes.ExecutePluginReport, [][]exectypes.CommitData, error)
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

// WithMultipleReports configures whether the builder should generate more than one report.
// WARNING: this feature only supports out of order messages.
// TODO: Move Nonce management to the Build function to support Nonce consistency across multiple reports.
func WithMultipleReports(enabled bool) Option {
	return func(erb *execReportBuilder) {
		erb.multipleReportsEnabled = enabled
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
	checks                 []Check
	destChainSelector      cciptypes.ChainSelector
	maxReportSizeBytes     uint64
	maxGas                 uint64
	maxMessages            uint64
	maxSingleChainReports  uint64
	multipleReportsEnabled bool

	// State
	// accumulated keeps track of per exec report metadata as commit reports are added
	// to the builder. It is used to limit what gets included in a report based on the
	// builder configuration. Metadata is added to the slice as multiple reports are
	// generated, with the last value in the slice being the currently building report.
	accumulated []validationMetadata

	// Result
	// execReports is the final output of the builder.
	execReports []cciptypes.ExecutePluginReport
	// commitReports is the updated CommitData based on the reports being built. If
	// multiple reports are being built, it's possible for the same CommitData object
	// to be included multiple times.
	commitReports [][]exectypes.CommitData
}

// checkInitialize initializes builder state if needed.
func (b *execReportBuilder) checkInitialize() {
	if b.accumulated == nil {
		b.accumulated = append(b.accumulated, validationMetadata{})
	}
	if b.execReports == nil {
		b.execReports = append(b.execReports, cciptypes.ExecutePluginReport{})
	}
	if b.commitReports == nil {
		b.commitReports = append(b.commitReports, []exectypes.CommitData{})
	}
}

// Add an exec report for as many messages as possible in the given commit report.
// The commit report with updated metadata is returned. It reflects which messages
// were selected for the exec report.
//
// TODO: Finish this function to generate multiple reports..
func (b *execReportBuilder) Add(
	ctx context.Context,
	commitReport exectypes.CommitData,
) (exectypes.CommitData, error) {
	b.checkInitialize()

	// Check which messages are ready to execute and update the report with additional metadata needed for execution.
	readyMessages, err := b.checkMessages(ctx, commitReport)
	if err != nil {
		return commitReport,
			fmt.Errorf("unable to check message: %w", err)
	}
	if len(readyMessages) == 0 {
		return commitReport, ErrEmptyReport
	}

	execReport, updatedReport, err := b.buildSingleChainReport(ctx, commitReport, readyMessages)

	// No messages fit into the exec report
	if errors.Is(err, ErrEmptyReport) {
		if !b.multipleReportsEnabled {
			return commitReport, nil
		}
		// start a new exec report and try again.
		// TODO: start a new exec report and try again.
		panic("multiple reports enabled but not implemented")
	}
	if err != nil {
		return commitReport, fmt.Errorf("unable to add a single chain report: %w", err)
	}

	index := len(b.execReports) - 1
	b.execReports[index].ChainReports = append(b.execReports[index].ChainReports, execReport)
	b.commitReports[index] = append(b.commitReports[index], updatedReport)

	// TODO: Check if there are any remaining readyMessages and try to build the next report.

	return updatedReport, nil
}

func (b *execReportBuilder) Build() (
	[]cciptypes.ExecutePluginReport, [][]exectypes.CommitData, error,
) {
	var results []cciptypes.ExecutePluginReport
	if len(b.execReports) != len(b.commitReports) {
		return nil, nil, fmt.Errorf(
			"expected the same number of exec and commit reports, got %d and %d",
			len(b.execReports),
			len(b.commitReports),
		)
	}

	numSingleChainReports := len(b.execReports)

	for _, report := range b.execReports {
		results = append(results, report)
		if !b.multipleReportsEnabled {
			// skip remaining reports if they aren't enabled.
			break
		}
	}

	// this now happens when adding the reports, so we don't need to do it here
	/*
		// Check if limiting is required.
		if b.maxSingleChainReports != 0 && uint64(len(b.execReports)) > b.maxSingleChainReports {
			b.lggr.Infof(
				"limiting number of reports to maxReports from %d to %d",
				len(b.execReports),
				b.maxSingleChainReports,
			)

			if b.multipleReportsEnabled {
				for len(b.execReports) > int(b.maxSingleChainReports) {
					results = append(results, b.execReports[:b.maxSingleChainReports])
					b.execReports = b.execReports[b.maxSingleChainReports:]
				}
				if len(b.execReports) > 0 {
					results = append(results, b.execReports)
				}
			} else {
				// Otherwise, truncate
				b.execReports = b.execReports[:b.maxSingleChainReports]
				b.commitReports = b.commitReports[:b.maxSingleChainReports]
				results = append(results, b.execReports)
				numSingleChainReports = len(b.execReports)
			}
		}
	*/

	// TODO: Sort reports into a canonical form prior to returning.

	b.lggr.Infow(
		"selected commit reports for execution report",
		"numReports", len(b.execReports),
		"numSingleChainReports", numSingleChainReports,
		// this is before any truncation
		// TODO: info for all of the reports?
		//"sizeBytes", b.accumulated.encodedSizeBytes,
		"maxSize", b.maxReportSizeBytes)
	return results, b.commitReports, nil
}
