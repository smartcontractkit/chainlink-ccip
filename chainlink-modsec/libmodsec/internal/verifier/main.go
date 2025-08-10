package verifier

import (
	"context"
	"fmt"
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
	if s.writer == nil {
		fmt.Println("No writer configured, verifier will not process any work.")
		// TODO: surface an error.
		return
	}
	if s.transformer == nil {
		fmt.Println("No transformer configured, verifier will not process any work.")
		// TODO: surface an error.
		return
	}
	if s.signer == nil {
		fmt.Println("No signer configured, verifier will not process any work.")
		// TODO: surface an error.
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	transformCh := make(chan Work)
	handlerPayloadCh := make(chan HandlerPayload)
	writePayloadCh := make(chan HandlerPayload)
	writeAttestationCh := make(chan Attestation)

	// Read blockchain data for new work.
	go func() {
		wg.Add(1)
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				fmt.Println("Verifier shutting down poller...")
				return
			case data := <-s.reader.Next(ctx):
				fmt.Println(data)
				transformCh <- data // to transformer
			}
		}
	}()

	// Transform work into the handler payload.
	go func() {
		wg.Add(1)
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				fmt.Println("Verifier shutting down transformer...")
				return
			case work := <-transformCh:
				payload := s.transformer.Transform(work)
				fmt.Println("Payload transformed:", payload)
				handlerPayloadCh <- payload // to handler dispatch.
				writePayloadCh <- payload   // to writer.
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
				fmt.Println("Verifier shutting down handlers...")
				return
			case payload := <-handlerPayloadCh:
				fmt.Printf("Running handlers for payload: %v\n", payload)
				// Dispatch payload to each handler.
				for _, handler := range s.handlers {
					// TODO: signing needs to fit in before writing.
					go handler(ctx, payload, writeAttestationCh)
				}
			}
		}
	}()

	// write results
	go func() {
		wg.Add(1)
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				fmt.Println("Verifier shutting down handlers...")
				return
			case payload := <-writePayloadCh:
				for _, w := range s.writer {
					if err := w.WriteMessage(ctx, payload); err != nil {
						// TODO: surface an error.
						fmt.Println("Error writing payload.")
					}
				}
			case attestation := <-writeAttestationCh:
				for _, w := range s.writer {
					if err := w.WriteAttestation(ctx, attestation); err != nil {
						// TODO: surface an error.
						fmt.Println("Error writing payload.")
					}
				}
			}
		}
	}()

	// Wait for service to stop.
	done := false
	for !done {
		select {
		case <-s.stopCh:
			cancel()
			done = true
		case <-ctx.Done():
			done = true
		}
	}

	wg.Wait()
}
