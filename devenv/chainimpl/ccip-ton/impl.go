package ccip_ton

import (
	"context"

	ccipTon "github.com/smartcontractkit/chainlink-ton/devenv-impl"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// CCIP16TON wraps the external chainlink-ton devenv-impl to add missing interface methods
type CCIP16TON struct {
	*ccipTon.CCIP16TON
}

// NewEmptyCCIP16TON creates a new wrapped CCIP16TON instance
func NewEmptyCCIP16TON() *CCIP16TON {
	return &CCIP16TON{
		CCIP16TON: ccipTon.NewEmptyCCIP16TON(),
	}
}

// LinkPingPongContracts is a no-op for TON - PingPong not yet implemented
func (m *CCIP16TON) LinkPingPongContracts(ctx context.Context, e *deployment.Environment, selector uint64, remoteSelectors []uint64) error {
	// PingPong contracts not yet implemented for TON
	return nil
}

