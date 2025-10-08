package v1_0

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	MCMSVersion = semver.MustParse("1.0.0")
)

type Deployer interface {
	DeployMCMS() *cldf_ops.Sequence[MCMSDeploymentConfigPerChain, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type DeployerRegistry struct {
	mu        sync.Mutex
	deployers map[string]Deployer
}

func (r *DeployerRegistry) RegisterDeployer(chainFamily string, version *semver.Version, deployer Deployer) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	if _, exists := r.deployers[id]; exists {
		panic("Deployer already registered for " + id)
	}
	r.deployers[id] = deployer
}

func NewDeployerRegistry() *DeployerRegistry {
	return &DeployerRegistry{
		mu:        sync.Mutex{},
		deployers: make(map[string]Deployer),
	}
}

func (r *DeployerRegistry) GetDeployer(chainFamily string, version *semver.Version) (Deployer, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	deployer, ok := r.deployers[id]
	return deployer, ok
}
