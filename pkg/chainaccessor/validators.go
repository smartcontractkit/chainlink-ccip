package chainaccessor

import (
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func ValidateSendRequestedEvent(
	ev *SendRequestedEvent, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) error {
	if ev == nil {
		return fmt.Errorf("send requested event is nil")
	}

	if ev.Message.Header.DestChainSelector != dest {
		return fmt.Errorf("msg dest chain is not the expected queried one")
	}
	if ev.DestChainSelector != dest {
		return fmt.Errorf("dest chain is not the expected queried one")
	}

	if ev.Message.Header.SourceChainSelector != source {
		return fmt.Errorf("source chain is not the expected queried one")
	}

	if ev.SequenceNumber != ev.Message.Header.SequenceNumber {
		return fmt.Errorf("event sequence number does not match the message sequence number %d != %d",
			ev.SequenceNumber, ev.Message.Header.SequenceNumber)
	}

	if ev.SequenceNumber < seqNumRange.Start() || ev.SequenceNumber > seqNumRange.End() {
		return fmt.Errorf("send requested event sequence number is not in the expected range")
	}

	if ev.Message.Header.MessageID.IsEmpty() {
		return fmt.Errorf("message ID is zero")
	}

	if len(ev.Message.Receiver) == 0 {
		return fmt.Errorf("empty receiver address: %s", ev.Message.Receiver.String())
	}

	if ev.Message.Sender.IsZeroOrEmpty() {
		return fmt.Errorf("invalid sender address: %s", ev.Message.Sender.String())
	}

	if ev.Message.FeeTokenAmount.IsEmpty() {
		return fmt.Errorf("fee token amount is zero")
	}

	if ev.Message.FeeToken.IsZeroOrEmpty() {
		return fmt.Errorf("invalid fee token: %s", ev.Message.FeeToken.String())
	}

	return nil
}
