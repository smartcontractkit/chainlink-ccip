package verifier

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// run is the verifier main loop.
//
// Implementation of a pipeline:
//   - Read events from blockchain.
//   - Transform events to the verifier payload.
//   - * Write raw message to storage.
//   - Dispatch payloads to verifiers.
//   - Collect results from verifiers as they are available.
//   - Write verifier attestations to storage.
func (s *Verifier) run() {
	// Start reader
	s.started = true
	if s.poller == nil {
		log.Println("No poller configured, verifier will not process any work.")
		// TODO: surface an error.
		return
	}
	if s.transformer == nil {
		log.Println("No transformer configured, verifier will not process any work.")
		// TODO: surface an error.
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	workCh := make(chan Work)
	messageCh := make(chan StoredMessage)
	handlerPayloadCh := make(chan HandlerPayload)
	resultCh := make(chan Result)

	// produce data
	go func() {
		wg.Add(1)
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				log.Println("Verifier shutting down poller...")
				return
			case data := <-s.poller.Watch(ctx):
				fmt.Println(data)
				workCh <- data
			}
		}
	}()

	// transform data
	go func() {
		wg.Add(1)
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				log.Println("Verifier shutting down transformer...")
				return
			case work := <-workCh:
				payload := s.transformer.Transform(work)
				fmt.Println("Payload transformed:", payload)
				handlerPayloadCh <- payload
				// TODO:: send StoredMessage to 'messageCh', it can be written ahead of verification.
			}
		}
	}()

	// run handlers
	go func() {
		wg.Add(1)
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				log.Println("Verifier shutting down handlers...")
				return
			case payload := <-handlerPayloadCh:
				log.Printf("Running handlers for payload: %v\n", payload)
				// Dispatch payload to each handler.
				for _, handler := range s.handlers {
					go handler(ctx, payload, resultCh)
				}
			}
		}
	}()

	// write results
}
