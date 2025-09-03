package verifier

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common"
)

// VerificationError represents an error that occurred during message verification
type VerificationError struct {
	Task      common.VerificationTask
	Error     error
	Timestamp time.Time
}

// Verifier defines the interface for message verification logic
type Verifier interface {
	// VerifyMessage performs the actual verification of a message asynchronously
	// The verifier should handle the task in its own goroutine and send results to ccvDataCh
	// Any verification errors should be sent to verificationErrorCh
	// This enables different verification flows based on finality, priority, etc.
	VerifyMessage(ctx context.Context, task common.VerificationTask, ccvDataCh chan<- common.CCVData, verificationErrorCh chan<- VerificationError)
}

// VerificationCoordinator orchestrates the verification workflow
// Reads messages from multiple SourceReaders and processes them using a Verifier
type VerificationCoordinator struct {
	// Basic operation channels
	ccvDataCh chan common.CCVData
	stopCh    chan struct{}
	doneCh    chan struct{}

	// Core components
	verifier Verifier
	// N Channels producing Any2AnyVerifierMessage
	sourceStates map[cciptypes.ChainSelector]*sourceState
	storage      common.OffchainStorageWriter

	// Configuration
	config CoordinatorConfig
	lggr   logger.Logger

	// State management
	mu      sync.RWMutex
	started bool
	stopped bool
}

// Option is the functional option type for VerificationCoordinator
type Option func(*VerificationCoordinator)

// WithVerifier sets the verifier implementation
func WithVerifier(verifier Verifier) Option {
	return func(vc *VerificationCoordinator) {
		vc.verifier = verifier
	}
}

// WithSourceReaders sets multiple source readers
func WithSourceReaders(sourceReaders map[cciptypes.ChainSelector]SourceReader) Option {
	return func(vc *VerificationCoordinator) {
		if vc.sourceStates == nil {
			vc.sourceStates = make(map[cciptypes.ChainSelector]*sourceState)
		}
		for chainSelector, reader := range sourceReaders {
			vc.sourceStates[chainSelector] = newSourceState(chainSelector, reader)
		}
	}
}

// AddSourceReader adds a single source reader to the existing map
func AddSourceReader(chainSelector cciptypes.ChainSelector, reader SourceReader) Option {
	return func(vc *VerificationCoordinator) {
		if vc.sourceStates == nil {
			vc.sourceStates = make(map[cciptypes.ChainSelector]*sourceState)
		}
		vc.sourceStates[chainSelector] = newSourceState(chainSelector, reader)
	}
}

// WithStorage sets the storage writer
func WithStorage(storage common.OffchainStorageWriter) Option {
	return func(vc *VerificationCoordinator) {
		vc.storage = storage
	}
}

// WithConfig sets the coordinator configuration
func WithConfig(config CoordinatorConfig) Option {
	return func(vc *VerificationCoordinator) {
		vc.config = config
	}
}

// WithLogger sets the logger
func WithLogger(lggr logger.Logger) Option {
	return func(vc *VerificationCoordinator) {
		vc.lggr = lggr
	}
}

// NewVerificationCoordinator creates a new verification coordinator
func NewVerificationCoordinator(opts ...Option) (*VerificationCoordinator, error) {
	//TODO: Make channel size configurable
	vc := &VerificationCoordinator{
		ccvDataCh:    make(chan common.CCVData, 1000), // Channel for CCVData results
		stopCh:       make(chan struct{}),
		doneCh:       make(chan struct{}),
		sourceStates: make(map[cciptypes.ChainSelector]*sourceState),
	}

	// Apply all options
	for _, opt := range opts {
		opt(vc)
	}

	// Validate required components
	if err := vc.validate(); err != nil {
		return nil, fmt.Errorf("invalid coordinator configuration: %w", err)
	}

	return vc, nil
}

// Start begins the verification coordinator processing
func (vc *VerificationCoordinator) Start(ctx context.Context) error {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	if vc.started {
		return fmt.Errorf("coordinator already started")
	}

	if vc.stopped {
		return errors.New("coordinator stopped")
	}

	// Start all source readers
	for chainSelector, state := range vc.sourceStates {
		if err := state.reader.Start(ctx); err != nil {
			return fmt.Errorf("failed to start source reader for chain %d: %w", chainSelector, err)
		}
	}

	vc.started = true

	// Start processing loop
	go vc.run(ctx)

	vc.lggr.Infow("VerificationCoordinator started",
		"coordinatorID", vc.config.VerifierID,
	)

	return nil
}

// Stop stops the verification coordinator processing
func (vc *VerificationCoordinator) Stop() error {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	if vc.stopped {
		return nil
	}

	vc.stopped = true
	vc.started = false
	close(vc.stopCh)

	// Stop all source readers and close error channels
	for chainSelector, state := range vc.sourceStates {
		if err := state.reader.Stop(); err != nil {
			vc.lggr.Errorw("Error stopping source reader", "error", err, "chainSelector", chainSelector)
		}
		// Close the per-source error channel
		close(state.verificationErrorCh)
	}

	// Wait for processing to finish
	<-vc.doneCh

	vc.lggr.Infow("VerificationCoordinator stopped")

	return nil
}

// run is the main processing loop
/*
src1 -> CCIPMessageSentEvent --processMessageEvent--->
src2 -> CCIPMessageSentEvent --processMessageEvent--->   ccvDatach --> Storage/Aggregator
src3 -> CCIPMessageSentEvent --processMessageEvent--->
*/
func (vc *VerificationCoordinator) run(ctx context.Context) {
	defer close(vc.doneCh)

	// Start goroutines for each source state
	var wg sync.WaitGroup
	for _, state := range vc.sourceStates {
		wg.Add(1)
		go vc.processSourceMessages(ctx, &wg, state)

		// Start error processing goroutine for each source
		wg.Add(1)
		go vc.processSourceErrors(ctx, &wg, state)
	}

	// Main loop - focus solely on ccvDataCh processing and storage
	for {
		select {
		case <-ctx.Done():
			vc.lggr.Infow("VerificationCoordinator processing stopped due to context cancellation")
			wg.Wait() // Wait for all source goroutines to finish
			return
		case <-vc.stopCh:
			vc.lggr.Infow("VerificationCoordinator processing stopped due to stop signal")
			wg.Wait() // Wait for all source goroutines to finish
			return
		case ccvData, ok := <-vc.ccvDataCh:
			if !ok {
				vc.lggr.Infow("CCVData channel closed, stopping processing")
				wg.Wait() // Wait for all source goroutines to finish
				return
			}

			// Store CCVData to offchain storage
			//TODO: handle errors/retries?
			// Do we want to store async as well?
			if err := vc.storage.StoreCCVData(ctx, []common.CCVData{ccvData}); err != nil {
				vc.lggr.Errorw("Error storing CCV data",
					"error", err,
					"messageID", ccvData.MessageID,
					"sequenceNumber", ccvData.SequenceNumber,
					"sourceChain", ccvData.SourceChainSelector,
				)
			} else {
				vc.lggr.Debugw("CCV data stored successfully",
					"messageID", ccvData.MessageID,
					"sequenceNumber", ccvData.SequenceNumber,
					"sourceChain", ccvData.SourceChainSelector,
				)
			}
		}
	}
}

// processSourceMessages handles message processing for a single source state in its own goroutine
func (vc *VerificationCoordinator) processSourceMessages(ctx context.Context, wg *sync.WaitGroup, state *sourceState) {
	defer wg.Done()
	chainSelector := state.chainSelector

	vc.lggr.Debugw("Starting source message processor", "chainSelector", chainSelector)
	defer vc.lggr.Debugw("Source message processor stopped", "chainSelector", chainSelector)

	for {
		select {
		case <-ctx.Done():
			vc.lggr.Debugw("Source message processor stopped due to context cancellation", "chainSelector", chainSelector)
			return
		case <-vc.stopCh:
			vc.lggr.Debugw("Source message processor stopped due to stop signal", "chainSelector", chainSelector)
			return
		case verificationTask, ok := <-state.verificationTaskCh:
			if !ok {
				vc.lggr.Errorw("Message channel closed for source", "chainSelector", chainSelector)
				return
			}
			// Process message event using the verifier asynchronously
			go vc.verifier.VerifyMessage(ctx, verificationTask, vc.ccvDataCh, state.verificationErrorCh)
		}
	}
}

// processSourceErrors handles error processing for a single source state in its own goroutine
func (vc *VerificationCoordinator) processSourceErrors(ctx context.Context, wg *sync.WaitGroup, state *sourceState) {
	defer wg.Done()
	chainSelector := state.chainSelector

	vc.lggr.Debugw("Starting source error processor", "chainSelector", chainSelector)
	defer vc.lggr.Debugw("Source error processor stopped", "chainSelector", chainSelector)

	for {
		select {
		case <-ctx.Done():
			vc.lggr.Debugw("Source error processor stopped due to context cancellation", "chainSelector", chainSelector)
			return
		case <-vc.stopCh:
			vc.lggr.Debugw("Source error processor stopped due to stop signal", "chainSelector", chainSelector)
			return
		case verificationError, ok := <-state.verificationErrorCh:
			if !ok {
				vc.lggr.Infow("Verification error channel closed for source", "chainSelector", chainSelector)
				return
			}

			// Handle verification errors for this specific source
			header := verificationError.Task.Message.Header
			vc.lggr.Errorw("Verification error received",
				"error", verificationError.Error,
				"messageID", header.MessageID,
				"sequenceNumber", header.SequenceNumber,
				"sourceChain", header.SourceChainSelector,
				"destChain", header.DestChainSelector,
				"timestamp", verificationError.Timestamp,
				"chainSelector", chainSelector,
			)

			//TODO: Add source-specific error handling strategies:
			// For Source-specific error handling we'll need some source-specific configuration
			// - Source-specific retry logic for transient errors
			// - Source-specific dead letter queues
			// - Source-specific metrics and alerting
		}
	}
}

// CommitVerifier provides a basic verifier implementation
type CommitVerifier struct {
	config CoordinatorConfig
	signer MessageSigner
	lggr   logger.Logger
}

// NewCommitVerifier creates a new commit verifier
func NewCommitVerifier(config CoordinatorConfig, signer MessageSigner, lggr logger.Logger) *CommitVerifier {
	return &CommitVerifier{
		config: config,
		signer: signer,
		lggr:   lggr,
	}
}

func (cv *CommitVerifier) ValidateMessage(message common.Any2AnyVerifierMessage) error {
	//TODO: Implement message validation
	return nil
}

// sendVerificationError sends a verification error to the error channel with proper context handling
func (cv *CommitVerifier) sendVerificationError(ctx context.Context, verificationTask common.VerificationTask, err error, verificationErrorCh chan<- VerificationError) {
	verificationError := VerificationError{
		Task:      verificationTask,
		Error:     err,
		Timestamp: time.Now(),
	}

	select {
	case verificationErrorCh <- verificationError:
	case <-ctx.Done():
	}
}

// VerifyMessage implements the Verifier interface
func (cv *CommitVerifier) VerifyMessage(ctx context.Context, verificationTask common.VerificationTask, ccvDataCh chan<- common.CCVData, verificationErrorCh chan<- VerificationError) {
	message := verificationTask.Message
	header := message.Header

	// Validate that the message comes from a configured source chain
	sourceConfig, exists := cv.config.SourceConfigs[header.SourceChainSelector]
	if !exists {
		cv.sendVerificationError(ctx, verificationTask,
			fmt.Errorf("message source chain selector %d is not configured", header.SourceChainSelector),
			verificationErrorCh)
		return
	}

	cv.lggr.Debugw("Message event validation passed",
		"messageID", header.MessageID,
		"sequenceNumber", header.SequenceNumber,
		"sourceChain", header.SourceChainSelector,
		"destChain", header.DestChainSelector,
	)

	// This is where different flows based on finality can be implemented
	// For example:
	// - Immediate verification for finalized blocks
	// - Delayed verification for unfinalized blocks

	// Sign the message event
	signature, err := cv.signer.SignMessage(ctx, verificationTask)
	if err != nil {
		cv.sendVerificationError(ctx, verificationTask,
			fmt.Errorf("failed to sign message event: %w", err),
			verificationErrorCh)
		return
	}

	// Create CCV data
	ccvData := &common.CCVData{
		MessageID:             header.MessageID,
		SequenceNumber:        header.SequenceNumber,
		SourceChainSelector:   header.SourceChainSelector,
		DestChainSelector:     header.DestChainSelector,
		SourceVerifierAddress: sourceConfig.VerifierAddress,
		CCVData:               signature,
		BlobData:              []byte{},
		Timestamp:             time.Now().UnixMicro(),
		Message:               message, // Store the complete message
	}

	// Send CCVData to channel for storage
	select {
	case ccvDataCh <- *ccvData:
		cv.lggr.Debugw("CCV data sent to storage channel",
			"messageID", header.MessageID,
			"sequenceNumber", header.SequenceNumber,
			"sourceChain", header.SourceChainSelector,
			"destChain", header.DestChainSelector,
		)
	case <-ctx.Done():
		cv.lggr.Debugw("Context cancelled while sending CCV data",
			"messageID", header.MessageID,
			"sequenceNumber", header.SequenceNumber,
			"sourceChain", header.SourceChainSelector,
		)
	}
}

// validate checks that all required components are configured
func (vc *VerificationCoordinator) validate() error {
	if len(vc.sourceStates) == 0 {
		return fmt.Errorf("at least one source reader is required")
	}

	// Validate that all configured sources have corresponding readers
	for chainSelector := range vc.config.SourceConfigs {
		if _, exists := vc.sourceStates[chainSelector]; !exists {
			return fmt.Errorf("source reader not found for chain selector %d", chainSelector)
		}
	}

	if vc.verifier == nil {
		return fmt.Errorf("verifier is required")
	}

	if vc.storage == nil {
		return fmt.Errorf("storage writer is required")
	}

	if vc.lggr == nil {
		return fmt.Errorf("logger is required")
	}

	if vc.config.VerifierID == "" {
		return fmt.Errorf("coordinator ID cannot be empty")
	}

	// Note: verifier is now required and must be set via WithVerifier option

	return nil
}

// HealthCheck returns the current health status
func (vc *VerificationCoordinator) HealthCheck(ctx context.Context) error {
	vc.mu.RLock()
	defer vc.mu.RUnlock()

	if vc.stopped {
		return errors.New("coordinator stopped")
	}

	if !vc.started {
		return errors.New("coordinator not started")
	}

	// Check all source readers health
	for chainSelector, state := range vc.sourceStates {
		if err := state.reader.HealthCheck(ctx); err != nil {
			return fmt.Errorf("source reader unhealthy for chain %d: %w", chainSelector, err)
		}
	}

	return nil
}
