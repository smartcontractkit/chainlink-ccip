package adapters

type DeployLombardInput[LocalContract any, RemoteContract any] struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// LombardVerifier is set of addresses comprising the LombardVerifier system.
	LombardVerifier []LocalContract
	// RMN is the address of the RMN contract.
	RMN LocalContract
	// Router is the address of the Router contract.
	Router LocalContract
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// StorageLocations is the set of storage locations for the LombardVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// Bridge is the address of the Bridge contract.
	Bridge string
}
