package verifier

import (
	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// sourceState represents the state of a single source reader
type sourceState struct {
	chainSelector       cciptypes.ChainSelector
	reader              SourceReader
	verificationTaskCh  <-chan common.VerificationTask
	verificationErrorCh chan VerificationError
}

func newSourceState(chainSelector cciptypes.ChainSelector, reader SourceReader) *sourceState {
	return &sourceState{
		chainSelector:       chainSelector,
		reader:              reader,
		verificationTaskCh:  reader.VerificationTaskChannel(),
		verificationErrorCh: make(chan VerificationError, 100), // TODO: Make configurable
	}
}
