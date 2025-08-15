package executor

import (
	"context"
	"fmt"
	"sync"
)

func (e *Executor) run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// 1. We read messages from all sources for any new incoming work.
	// Immediately we perform checks to see if it's our turn to execute the message, and skip additional processing if not.
	for _, reader := range e.MessageSources {
		go func(r MessageReader) {
			wg.Add(1)
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Println("Executor shutting down message reader...")
					return
				case msg := <-reader.SubscribeMessages(ctx):
					// Check if it's our turn to execute the message, if not send it to the timed message channel
					isLeader, delay := e.isMyTurn(msg)
					if !isLeader {
						e.timedMessageCh.SendMessage(msg, delay)
						continue
					}

					if e.Transmitters[msg.Header.DestChainSelector] == nil {
						fmt.Printf("No chain writer for destination chain selector %s, skipping message %s\n",
							msg.Header.DestChainSelector, msg.Header.MessageID)
						continue
					}
					e.messageCh <- msg
					fmt.Printf("Executor received message %s, it's our turn to execute it\n", msg.Header.MessageID)
				case msg := <-e.timedMessageCh.Messages():
					// We don't need to check if it's our turn to execute the message here, we just need to check if it has been executed
					// TODO: the above

					e.messageCh <- msg
				}
			}
		}(reader)
	}

	transmitterPayloadCh := make(chan TransmitterPayload)

	// 2. Now that we know we have messages that we are assigned to execute, try to get attestations for them.
	//
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Executor shutting down attestation poller...")
				return
			case msg := <-e.messageCh:

				// TODO: We need some abstraction here to loop and handle retries when attestations are not yet ready
				attestations, err := e.AttestationReader.GetAttestations(ctx, msg)
				if err != nil {
					// handle error, log it, etc.
					continue
				}

				transmitterPayloadCh <- TransmitterPayload{
					Message:      msg,
					Attestations: attestations,
				}
			}
		}
	}()

	// 3. Once we have a quorum of attestations for a specific message, write the payload to the destination chain.
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Executor shutting down contract transmitter...")
				return
			case payloads := <-transmitterPayloadCh:
				transmitter, exists := e.Transmitters[payloads.Message.Header.DestChainSelector]
				if !exists {
					// could not find a transmitter for dest chain, this shouldn't happen because the executor would not have accepted work for this message
					// if there's an rpc error, we would handle that down the line
					// handle error
					continue
				}

				encodedMessage, _ := e.messageCodec.Encode(ctx, payloads.Message)
				proofs := make([][]byte, len(payloads.Attestations))
				for _, attestation := range payloads.Attestations {
					proofs = append(proofs, attestation.Proof)
				}

				err := transmitter.Transmit(ctx, encodedMessage, proofs, nil)
				if err != nil {
					// handle error
					continue
				}
			}
		}
	}()

	// Wait for service to stop.
	done := false
	for !done {
		select {
		case <-e.stopCh:
			cancel()
			done = true
		case <-ctx.Done():
			done = true
		}
	}

	wg.Wait()
}
