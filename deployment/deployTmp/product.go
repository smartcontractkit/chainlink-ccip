package deploytmp

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	MCMSVersion       = semver.MustParse("1.0.0")
	singletonRegistry *DeployerTmpRegistry
	once              sync.Once
	chainAdapterOnce  sync.Once
)

type DeployerTmp interface {
	UpdateMCMSConfig() *cldf_ops.Sequence[UpdateMCMSConfigInputPerChainWithSelector, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type DeployerTmpRegistry struct {
	mu        sync.Mutex
	deployers map[string]DeployerTmp
}

func (r *DeployerTmpRegistry) RegisterDeployer(chainFamily string, version *semver.Version, deployer DeployerTmp) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	if _, exists := r.deployers[id]; !exists {
		r.deployers[id] = deployer
	}
}

func newDeployerTmpRegistry() *DeployerTmpRegistry {
	return &DeployerTmpRegistry{
		mu:        sync.Mutex{},
		deployers: make(map[string]DeployerTmp),
	}
}

// GetRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetRegistry() *DeployerTmpRegistry {
	once.Do(func() {
		singletonRegistry = newDeployerTmpRegistry()
	})
	return singletonRegistry
}

func (r *DeployerTmpRegistry) GetDeployer(chainFamily string, version *semver.Version) (DeployerTmp, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	deployer, ok := r.deployers[id]
	return deployer, ok
}

func (r *DeployerTmpRegistry) Blockchain(e cldf.Environment, chainSelector uint64) (cldf_chain.BlockChain, error) {
	allChains := e.BlockChains.All()
	for sel, chain := range allChains {
		if sel == chainSelector {
			return chain, nil
		}
	}
	return nil, fmt.Errorf("no blockchain found in environment for selector %d", chainSelector)
}

func (r *DeployerTmpRegistry) ExistingAddressesForChain(e cldf.Environment, chainSelector uint64) []datastore.AddressRef {
	// filter addresses for the given chain selector
	return e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector))
}
