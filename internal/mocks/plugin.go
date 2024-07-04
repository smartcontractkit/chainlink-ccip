package mocks

import (
	"context"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/mock"
)

type PluginMock struct {
	*mock.Mock
}

func NewPluginMock() *PluginMock {
	return &PluginMock{}
}

func (m *PluginMock) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	args := m.Called(ctx, outctx)
	return args.Get(0).(types.Query), args.Error(1)
}

func (m *PluginMock) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query) (types.Observation, error) {
	args := m.Called(ctx, outctx, query)
	return args.Get(0).(types.Observation), args.Error(1)
}

func (m *PluginMock) ValidateObservation(outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation) error {
	args := m.Called(outctx, query, ao)
	return args.Error(0)
}

func (m *PluginMock) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.OutcomeContext, error) {
	args := m.Called(outctx, query)
	return args.Get(0).(ocr3types.OutcomeContext), args.Error(1)
}

func (m *PluginMock) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	args := m.Called(outctx, query, aos)
	return args.Get(0).(ocr3types.Outcome), args.Error(1)
}

func (m *PluginMock) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	args := m.Called(seqNr, outcome)
	return args.Get(0).([]ocr3types.ReportWithInfo[[]byte]), args.Error(1)
}

func (m *PluginMock) ShouldAcceptAttestedReport(ctx context.Context, seqNr uint64, report ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	args := m.Called(ctx, seqNr, report)
	return args.Bool(0), args.Error(1)
}

func (m *PluginMock) ShouldTransmitAcceptedReport(ctx context.Context, seqNr uint64, report ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	args := m.Called(ctx, seqNr, report)
	return args.Bool(0), args.Error(1)
}

func (m *PluginMock) Close() error {
	args := m.Called()
	return args.Error(0)
}
