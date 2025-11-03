package deploy

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	MCMSVersion              = semver.MustParse("1.0.0")
	singletonRegistry        *DeployerRegistry
	once                     sync.Once
	singletonAdapterRegistry *TransferOwnershipAdapterRegistry
	chainAdapterOnce         sync.Once
)

type Deployer interface {
	DeployChainContracts() *cldf_ops.Sequence[ContractDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains]
	DeployMCMS() *cldf_ops.Sequence[MCMSDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains]
	SetOCR3Config() *cldf_ops.Sequence[SetOCR3ConfigInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type DeployerRegistry struct {
	mu        sync.Mutex
	deployers map[string]Deployer
}

func (r *DeployerRegistry) RegisterDeployer(chainFamily string, version *semver.Version, deployer Deployer) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	if _, exists := r.deployers[id]; !exists {
		r.deployers[id] = deployer
	}
}

func newDeployerRegistry() *DeployerRegistry {
	return &DeployerRegistry{
		mu:        sync.Mutex{},
		deployers: make(map[string]Deployer),
	}
}

// GetRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetRegistry() *DeployerRegistry {
	once.Do(func() {
		singletonRegistry = newDeployerRegistry()
	})
	return singletonRegistry
}

func (r *DeployerRegistry) GetDeployer(chainFamily string, version *semver.Version) (Deployer, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	deployer, ok := r.deployers[id]
	return deployer, ok
}

func (r *DeployerRegistry) Blockchain(e cldf.Environment, chainSelector uint64) (cldf_chain.BlockChain, error) {
	allChains := e.BlockChains.All()
	for sel, chain := range allChains {
		if sel == chainSelector {
			return chain, nil
		}
	}
	return nil, fmt.Errorf("no blockchain found in environment for selector %d", chainSelector)
}

func (r *DeployerRegistry) ExistingAddressesForChain(e cldf.Environment, chainSelector uint64) []datastore.AddressRef {
	// filter addresses for the given chain selector
	return e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector))
}

type TransferOwnershipAdapter interface {
	InitializeTimelockAddress(e cldf.Environment, input mcms.Input) error
	SequenceTransferOwnershipViaMCMS() *cldf_ops.Sequence[TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	SequenceAcceptOwnership() *cldf_ops.Sequence[TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	ShouldAcceptOwnershipWithTransferOwnership(e cldf.Environment, in TransferOwnershipPerChainInput) (bool, error)
}
