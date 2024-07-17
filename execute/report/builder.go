package report

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/execute/temp"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

var _ ExecReportBuilder = &execReportBuilder{}

type ExecReportBuilder interface {
	Add(report plugintypes.ExecutePluginCommitDataWithMessages) (plugintypes.ExecutePluginCommitDataWithMessages, error)
	Build() ([]cciptypes.ExecutePluginReportSingleChain, error)
}

func NewBuilder(
	ctx context.Context,
	logger logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader temp.TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	maxReportSizeBytes uint64,
	maxGas uint64,
) ExecReportBuilder {
	return &execReportBuilder{
		ctx:  ctx,
		lggr: logger,

		tokenDataReader: tokenDataReader,
		encoder:         encoder,
		hasher:          hasher,

		maxReportSizeBytes: maxReportSizeBytes,
		maxGas:             maxGas,
	}
}

type execReportBuilder struct {
	ctx  context.Context // TODO: remove context from builder so that it can be pure?
	lggr logger.Logger

	// Providers
	tokenDataReader temp.TokenDataReader
	encoder         cciptypes.ExecutePluginCodec
	hasher          cciptypes.MessageHasher

	// Config
	maxReportSizeBytes uint64
	maxGas             uint64

	// State
	reportSizeBytes uint64
	gas             uint64

	// Result
	execReports []cciptypes.ExecutePluginReportSingleChain
}

func (b *execReportBuilder) Add(commitReport plugintypes.ExecutePluginCommitDataWithMessages) (plugintypes.ExecutePluginCommitDataWithMessages, error) {

	// TODO: buildSingleChainReportMaxSize needs to be part of the builder in order to access state.
	execReport, encodedSize, updatedReport, err :=
		buildSingleChainReportMaxSize(b.ctx, b.lggr, b.hasher, b.tokenDataReader, b.encoder,
			commitReport,
			int(b.maxReportSizeBytes-b.reportSizeBytes))

	// No messages fit into the report, move to next report
	if errors.Is(err, ErrEmptyReport) {
		return commitReport, nil
	}
	if err != nil {
		return commitReport, fmt.Errorf("unable to build single chain report: %w", err)
	}

	b.reportSizeBytes += uint64(encodedSize)
	b.execReports = append(b.execReports, execReport)

	/*
		// partially executed report detected, stop adding more reports.
		// TODO: do not break if messages were intentionally skipped.
		if len(updatedReport.Messages) != len(updatedReport.ExecutedMessages) {
			break
		}
	*/
	return updatedReport, nil
}

func (b *execReportBuilder) Build() ([]cciptypes.ExecutePluginReportSingleChain, error) {
	return b.execReports, nil
}
