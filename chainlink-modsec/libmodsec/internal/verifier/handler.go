// Package verifier implements a service that monitors a channel for incoming
// messages and sends attestations to a results channel.
package verifier

import (
	"context"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

type Work struct {
	// Define your work payload here
	Data []byte // maybe nil if there were no messages for a particular block num

	// Block information. The same block data may be sent multiple times if there are multiple messages.
	SourceChain uint64
	BlockNum    uint64 // why do they use big int... at 100k  bps it would take 500k years to overflow uint64.
	BlockHash   []byte
}

type HandlerPayload struct {
	Work    Work
	Message modsectypes.Message
	//TODO encoded message as well.
}

type Result struct {
	// Define your result payload here
	Data []byte
}

// Handler is a function that processes incoming work from the work channel and sends results to the result channel.
// There may be no results for a particular piece of work, in which case the result channel should not be written to.
// For example, the verifier may have a rule to skip processing until a certain block depth is reached. In that case,
// the work may be received, but no result will be produced until later.
//
// The verifier should respect the context for cancellation. If the context is cancelled, the verifier should stop.
//
// TODO: should verifier be a full service with start/stop/cancel, a cache, maybe a db connection, etc?
type Handler func(ctx context.Context, payload HandlerPayload, result chan<- Result)
