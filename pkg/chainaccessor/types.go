package chainaccessor

import cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

// SendRequestedEvent represents the contents of the event emitted by the CCIP onramp when a message is sent.
type SendRequestedEvent struct {
	DestChainSelector cciptypes.ChainSelector
	SequenceNumber    cciptypes.SeqNum
	Message           cciptypes.Message
}
