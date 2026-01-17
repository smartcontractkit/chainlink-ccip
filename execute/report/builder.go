package report

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
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

func WithMaxReportsCount(maxReportCount uint64) Option {
	return func(erb *execReportBuilder) {
		erb.maxReportCount = maxReportCount
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
	maxReportCount         uint64

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

	// Validate nonces for multiple reports mode
	if b.multipleReportsEnabled {
		if err := b.validateNoncesForMultipleReports(commitReport); err != nil {
			return commitReport, fmt.Errorf("multiple reports validate nonces: %w", err)
		}
	}

	currentCommitReport := commitReport
	// Main processing loop - continue until no more messages or reports can be created
	for {
		// Check which messages are ready to execute, excluding already processed ones
		readyMessages, err := b.extractReadyMessages(ctx, currentCommitReport)
		if err != nil {
			return currentCommitReport, fmt.Errorf("unable to extract ready messages: %w", err)
		}

		if len(readyMessages) == 0 {
			// No more messages to process
			return currentCommitReport, nil
		}

		// Check if we need to create a new exec report due to limits
		if b.shouldCreateNewExecReport() {
			if !b.multipleReportsEnabled {
				return currentCommitReport, nil
			}
			b.createNewExecReport()
		}

		if uint64(len(b.execReports)) > b.maxReportCount {
			b.lggr.Debugw("Reached maximum report count, stopping further processing",
				"maxReportCount", b.maxReportCount,
				"currentCount", len(b.execReports))
			return currentCommitReport, nil
		}

		// Try to build a report with current messages
		updatedReport, shouldContinue, err := b.tryBuildReport(ctx, currentCommitReport, readyMessages)
		if err != nil {
			return currentCommitReport, err
		}

		currentCommitReport = updatedReport

		if !shouldContinue {
			break
		}
	}

	return currentCommitReport, nil
}

// tryBuildReport attempts to build a single chain report and handle empty report cases
func (b *execReportBuilder) tryBuildReport(
	ctx context.Context,
	commitReport exectypes.CommitData,
	readyMessages map[int]struct{},
) (exectypes.CommitData, bool, error) {
	currentIndex := len(b.execReports) - 1

	execReport, updatedReport, err := b.buildSingleChainReport(ctx, commitReport, readyMessages)

	if err != nil {
		// If the report is empty, we handle it separately especially when multiple reports are enabled.
		if errors.Is(err, ErrEmptyReport) {
			return b.handleEmptyReport(ctx, commitReport, readyMessages)
		}
		return commitReport, false, fmt.Errorf("unable to add a single chain report: %w", err)
	}

	// Successfully built a report - update our data structures
	b.execReports[currentIndex].ChainReports = append(b.execReports[currentIndex].ChainReports, execReport)
	b.commitReports[currentIndex] = append(b.commitReports[currentIndex], updatedReport)

	// Determine if we should continue processing
	shouldContinue := b.multipleReportsEnabled
	return updatedReport, shouldContinue, nil
}

// handleEmptyReport deals with the case where no messages fit into the current report
func (b *execReportBuilder) handleEmptyReport(
	ctx context.Context,
	commitReport exectypes.CommitData,
	readyMessages map[int]struct{},
) (exectypes.CommitData, bool, error) {
	if !b.multipleReportsEnabled {
		return commitReport, false, nil
	}

	// Don't create new reports if we've reached the max count
	if uint64(len(b.execReports)) >= b.maxReportCount {
		return commitReport, false, nil
	}
	b.lggr.Debugw("Creating new exec report due to empty report")
	// Try with a new exec report
	b.createNewExecReport()
	currentIndex := len(b.execReports) - 1

	execReport, updatedReport, err := b.buildSingleChainReport(ctx, commitReport, readyMessages)
	if errors.Is(err, ErrEmptyReport) {
		// Remove the empty report we just added and stop processing
		b.removeLastExecReport()
		return commitReport, false, nil
	}

	if err != nil {
		return commitReport, false, fmt.Errorf("unable to add a single chain report: %w", err)
	}

	// Successfully built a report with the new exec report
	b.execReports[currentIndex].ChainReports = append(b.execReports[currentIndex].ChainReports, execReport)
	b.commitReports[currentIndex] = append(b.commitReports[currentIndex], updatedReport)

	return updatedReport, true, nil
}

// shouldCreateNewExecReport checks if we need to create a new exec report due to chain report limits
func (b *execReportBuilder) shouldCreateNewExecReport() bool {
	if b.maxSingleChainReports == 0 {
		return false
	}

	currentIndex := len(b.execReports) - 1
	return len(b.execReports[currentIndex].ChainReports) >= int(b.maxSingleChainReports)
}

func (b *execReportBuilder) validateNoncesForMultipleReports(commitReport exectypes.CommitData) error {
	for _, msg := range commitReport.Messages {
		if msg.Header.Nonce > 0 {
			b.lggr.Errorw("Found message with non-zero nonce when multiple reports are enabled",
				"sourceChain", msg.Header.SourceChainSelector,
				"messageID", msg.Header.MessageID,
				"nonce", msg.Header.Nonce)

			return fmt.Errorf("messages with non-zero nonces detected")
		}
	}
	return nil
}

func (b *execReportBuilder) createNewExecReport() {
	b.execReports = append(b.execReports, cciptypes.ExecutePluginReport{})
	b.commitReports = append(b.commitReports, []exectypes.CommitData{})
	b.accumulated = append(b.accumulated, validationMetadata{})
}

func (b *execReportBuilder) removeLastExecReport() {
	if len(b.execReports) == 0 {
		b.lggr.Error("Attempted to remove last exec report, but no reports exist")
		return
	}
	b.execReports = b.execReports[:len(b.execReports)-1]
	b.commitReports = b.commitReports[:len(b.commitReports)-1]
	b.accumulated = b.accumulated[:len(b.accumulated)-1]
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
		if !b.multipleReportsEnabled || uint64(len(results)) >= b.maxReportCount {
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
