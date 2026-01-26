package adapters

import "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"

type DeployLombardInput[LocalContract any, RemoteContract any] struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// LombardVerifier is set of addresses comprising the LombardVerifier system.
	LombardVerifier []LocalContract
	Token           string
	TokenPool       LocalContract
	// TokenAdminRegistry is the address of the TokenAdminRegistry contract.
	TokenAdminRegistry LocalContract
	// RMN is the address of the RMN contract.
	RMN LocalContract
	// Router is the address of the Router contract.
	Router LocalContract
	// RemoteChains is the set of remote chains to configure on the CCTPVerifier contract.
	RemoteChains map[uint64]RemoteLombardChainConfig[LocalContract, RemoteContract]
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// StorageLocations is the set of storage locations for the LombardVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// Bridge is the address of the Bridge contract.
	Bridge string
	// RateLimitAdmin is the address allowed to update token pool rate limits.
	RateLimitAdmin string
}

type RemoteLombardChainConfig[LocalContract any, RemoteContract any] struct {
	// TokenPoolConfig configures the token pool for the remote chain.
	TokenPoolConfig tokens.RemoteChainConfig[RemoteContract, LocalContract]
	RemoteDomain    LombardRemoteDomain[RemoteContract]
}

// LombardRemoteDomain identifies Lombard-specific parameters for a remote chain.
type LombardRemoteDomain[RemoteContract any] struct {
	AllowedCaller RemoteContract
	LChainId      uint32
}
