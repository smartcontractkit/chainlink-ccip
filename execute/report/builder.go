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
// TODO: Add support for multiple reports with ordered messages.
func (b *execReportBuilder) Add(
	ctx context.Context,
	commitReport exectypes.CommitData,
) (exectypes.CommitData, error) {
	b.checkInitialize()

	if b.multipleReportsEnabled {
		// Check if all messages have zero nonces - if not, we should only process the first report
		for _, msg := range commitReport.Messages {
			if msg.Header.Nonce > 0 {
				b.lggr.Errorw("Found message with non-zero nonce when multiple reports are enabled",
					"sourceChain", msg.Header.SourceChainSelector,
					"messageID", msg.Header.MessageID,
					"nonce", msg.Header.Nonce)
				// If multiple reports are enabled, we can still process the report, but we need to limit it to a single report
				// If we didn't even build the first report yet, then break and process the first report only.
				if len(b.execReports) >= 1 && len(b.execReports[0].ChainReports) != 0 {
					b.execReports = []cciptypes.ExecutePluginReport{b.execReports[0]}
					b.commitReports = [][]exectypes.CommitData{b.commitReports[0]}
					b.accumulated = []validationMetadata{b.accumulated[0]}
					return commitReport, fmt.Errorf("messages with non-zero nonces detected, limiting to single report")
				}
			}
		}
	}

	currentCommitReport := commitReport
	// Main processing loop - continue until no more messages or reports can be created
	for {
		index := len(b.execReports) - 1

		// Check if we've reached max reports for the current exec report
		if b.maxSingleChainReports != 0 && len(b.execReports[index].ChainReports) >= int(b.maxSingleChainReports) {
			if !b.multipleReportsEnabled {
				return currentCommitReport, nil
			}

			// Start a new exec report
			b.execReports = append(b.execReports, cciptypes.ExecutePluginReport{})
			b.commitReports = append(b.commitReports, []exectypes.CommitData{})
			b.accumulated = append(b.accumulated, validationMetadata{})
			index = len(b.execReports) - 1
		}

		// Check which messages are ready to execute, excluding already processed ones
		readyMessages, err := b.checkMessages(ctx, currentCommitReport)
		if err != nil {
			return currentCommitReport, fmt.Errorf("unable to check messages: %w", err)
		}

		if len(readyMessages) == 0 {
			// No more messages to process
			return currentCommitReport, nil
		}

		execReport, updatedReport, err := b.buildSingleChainReport(ctx, currentCommitReport, readyMessages)

		if errors.Is(err, ErrEmptyReport) {
			// No messages fit into the report
			if !b.multipleReportsEnabled {
				return currentCommitReport, nil
			}

			// Try with a new exec report
			b.execReports = append(b.execReports, cciptypes.ExecutePluginReport{})
			b.commitReports = append(b.commitReports, []exectypes.CommitData{})
			b.accumulated = append(b.accumulated, validationMetadata{})
			index = len(b.execReports) - 1

			execReport, updatedReport, err = b.buildSingleChainReport(ctx, currentCommitReport, readyMessages)
			if errors.Is(err, ErrEmptyReport) {
				// Remove the empty report we just added
				b.execReports = b.execReports[:len(b.execReports)-1]
				b.commitReports = b.commitReports[:len(b.commitReports)-1]
				b.accumulated = b.accumulated[:len(b.accumulated)-1]
				return currentCommitReport, nil
			}
		}

		if err != nil {
			return currentCommitReport, fmt.Errorf("unable to add a single chain report: %w", err)
		}

		// Update our data structures
		b.execReports[index].ChainReports = append(b.execReports[index].ChainReports, execReport)
		b.commitReports[index] = append(b.commitReports[index], updatedReport)
		currentCommitReport = updatedReport

		// If multiple reports aren't enabled, we're done after the first report
		if !b.multipleReportsEnabled {
			break
		}

		// Check if we've processed all messages - if so, exit the loop
		readyMessagesLeft, _ := b.checkMessages(ctx, currentCommitReport)
		if len(readyMessagesLeft) == 0 {
			return currentCommitReport, nil
		}
	}

	return currentCommitReport, nil
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
