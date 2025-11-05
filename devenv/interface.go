package ccip

import (
	"context"
	"math/big"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"

	nodeset "github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"
)

/*
This package contains interfaces for devenv to load chain-specific product implementations for CCIP 1.6
*/

// CCIP16ProductConfiguration includes all the interfaces that if implemented allows us to run a standard test suite for 2+ chains
// it deploys network-specific infrastructure, configures both CL nodes and contracts and returns
// operations for testing and SLA/Metrics assertions.
type CCIP16ProductConfiguration interface {
	Chains
	Observable
	OnChainConfigurable
	OffChainConfigurable
}

// Observable pushes Loki streams and exposes Prometheus metrics and returns queries to assert SLAs.
type Observable interface {
	// ExposeMetrics exposes Prometheus metrics for the given source and destination chain IDs.
	ExposeMetrics(ctx context.Context, source, dest uint64, chainIDs, wsURLs []string) ([]string, *prometheus.Registry, error)
}

// Chains provides methods to interact with a set of chains that have CCIP deployed.
type Chains interface {
	// SetCLDF sets CLDF environment
	SetCLDF(e *deployment.Environment)
	// GetEOAReceiverAddress gets an EOA receiver address for the provided chain selector.
	GetEOAReceiverAddress(ctx context.Context, chainSelector uint64) ([]byte, error)
	// SendMessage sends a CCIP message from src to dest with the specified message options.
	SendMessage(ctx context.Context, src, dest uint64, fields any, opts any) error
	// GetExpectedNextSequenceNumber gets an expected sequence number for message with "from" and "to" selectors
	GetExpectedNextSequenceNumber(ctx context.Context, from, to uint64) (uint64, error)
	// WaitOneSentEventBySeqNo waits until exactly one event for CCIP message sent is emitted on-chain
	WaitOneSentEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error)
	// WaitOneExecEventBySeqNo waits until exactly one event for CCIP execution state change is emitted on-chain
	WaitOneExecEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error)
	// GetTokenBalance gets the balance of an account for a token on a chain
	GetTokenBalance(ctx context.Context, chainSelector uint64, address, tokenAddress []byte) (*big.Int, error)
}

// OnChainConfigurable defines methods that allows devenv to
// deploy, configure Chainlink product and connect on-chain part with other chains.
type OnChainConfigurable interface {
	// DeployContractsForSelector configures contracts for chain X
	// returns all the contract addresses and metadata as datastore.DataStore
	DeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*nodeset.Input, selector uint64) (datastore.DataStore, error)
	// ConnectContractsWithSelectors connects this chain onRamp to one or multiple offRamps for remote selectors (other chains)
	ConnectContractsWithSelectors(ctx context.Context, e *deployment.Environment, selector uint64, remoteSelectors []uint64) error
}

// OffChainConfigurable defines methods that allows to
// deploy a local blockchain network for tests and configure CL nodes for Chainlink product.
type OffChainConfigurable interface {
	// DeployLocalNetwork deploy local node of network X
	DeployLocalNetwork(ctx context.Context, bcs *blockchain.Input) (*blockchain.Output, error)
	// ConfigureNodes configure CL nodes from blockchain data
	// returns a piece of TOML config as a string that the framework inject into final configuration
	ConfigureNodes(ctx context.Context, blockchain *blockchain.Input) (string, error)
	// FundNodes Fund Chainlink nodes for some amount of native/LINK currency
	// using chain-specific clients or CLDF
	FundNodes(ctx context.Context, cls []*nodeset.Input, bc *blockchain.Input, linkAmount, nativeAmount *big.Int) error
}
