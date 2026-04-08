package adapters

import (
	"fmt"
	"math/big"
	"sync"

	"dario.cat/mergo"
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type CommitteeVerifierDeployParams struct {
	Version          *semver.Version
	FeeAggregator    string
	AllowlistAdmin   string
	StorageLocations []string
	Qualifier        string
}

type RMNRemoteDeployParams struct {
	Version   *semver.Version
	LegacyRMN string
}

type OffRampDeployParams struct {
	Version                   *semver.Version
	GasForCallExactCheck      uint16
	MaxGasBufferToUpdateState uint32
}

type OnRampDeployParams struct {
	Version               *semver.Version
	FeeAggregator         string
	MaxUSDCentsPerMessage uint32
}

type FeeQuoterDeployParams struct {
	Version                        *semver.Version
	MaxFeeJuelsPerMsg              *big.Int
	LINKPremiumMultiplierWeiPerEth uint64
	WETHPremiumMultiplierWeiPerEth uint64
	USDPerLINK                     *big.Int
	USDPerWETH                     *big.Int
}

type ExecutorDynamicDeployConfig struct {
	FeeAggregator         string
	MinBlockConfirmations uint16
	CcvAllowlistEnabled   bool
}

type ExecutorDeployParams struct {
	Version       *semver.Version
	MaxCCVsPerMsg uint8
	DynamicConfig ExecutorDynamicDeployConfig
	Qualifier     string
}

type MockReceiverDeployParams struct {
	Version                   *semver.Version
	RequiredVerifiers         []datastore.AddressRef
	OptionalVerifiers         []datastore.AddressRef
	OptionalThreshold         uint8
	MinimumBlockConfirmations uint16
	Qualifier                 string
}

type DeployContractParams struct {
	RMNRemote          RMNRemoteDeployParams
	OffRamp            OffRampDeployParams
	CommitteeVerifiers []CommitteeVerifierDeployParams
	OnRamp             OnRampDeployParams
	FeeQuoter          FeeQuoterDeployParams
	Executors          []ExecutorDeployParams
	MockReceivers      []MockReceiverDeployParams
}

// MergeWithOverrideIfNotEmpty merges source into a copy of d. Only non-empty source fields overwrite
func (d DeployContractParams) MergeWithOverrideIfNotEmpty(source DeployContractParams) (DeployContractParams, error) {
	result := d
	if err := mergo.Merge(&result, &source, mergo.WithOverride); err != nil {
		return DeployContractParams{}, fmt.Errorf("failed to merge DeployContractParams: %w", err)
	}
	return result, nil
}

type DeployChainContractsInput struct {
	ChainSelector     uint64
	DeployerContract  string
	DeployTestRouter  bool
	ExistingAddresses []datastore.AddressRef
	ContractParams    DeployContractParams
	// DeployerKeyOwned, when true, skips the transfer-ownership step so that
	// contracts remain owned by the deployer key. By default (false) the
	// sequence looks up the existing CLLCCIP RBACTimelock in ExistingAddresses
	// and transfers ownership of product contracts to it, failing fast if the
	// required MCMS instances are not found.
	DeployerKeyOwned bool
}

type DeployChainConfigCreatorInput struct {
	ChainSelector      uint64
	ExistingAddresses  []datastore.AddressRef
	ContractMeta       []datastore.ContractMetadata
	UserProvidedConfig DeployContractParams
}

type DeployChainContractsAdapter interface {
	// SetContractParamsFromImportedConfig is used when ImportConfig is true in DeployChainContractsInput.
	// It should read the necessary contract parameters from the datastore contract metadata based on the chain selector and
	// return them in the same format as DeployContractParams for use in the deployment sequence.
	SetContractParamsFromImportedConfig() *cldf_ops.Sequence[DeployChainConfigCreatorInput, DeployContractParams, cldf_chain.BlockChains]
	DeployChainContracts() *cldf_ops.Sequence[DeployChainContractsInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type DeployChainContractsRegistry struct {
	mu                  sync.Mutex
	adapters            map[string]DeployChainContractsAdapter
	configImporters     map[string]deploy.ConfigImporter
	laneVersionResolver map[string]deploy.LaneVersionResolver
}

var (
	singletonDeployChainContractsRegistry *DeployChainContractsRegistry
	deployChainContractsRegistryOnce      sync.Once
)

func NewDeployChainContractsRegistry() *DeployChainContractsRegistry {
	return &DeployChainContractsRegistry{
		adapters:            make(map[string]DeployChainContractsAdapter),
		configImporters:     make(map[string]deploy.ConfigImporter),
		laneVersionResolver: make(map[string]deploy.LaneVersionResolver),
	}
}

func GetDeployChainContractsRegistry() *DeployChainContractsRegistry {
	deployChainContractsRegistryOnce.Do(func() {
		singletonDeployChainContractsRegistry = NewDeployChainContractsRegistry()
	})
	return singletonDeployChainContractsRegistry
}

func (r *DeployChainContractsRegistry) Register(family string, a DeployChainContractsAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *DeployChainContractsRegistry) RegisterConfigImporter(family string, version *semver.Version, importer deploy.ConfigImporter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	if _, exists := r.configImporters[id]; !exists {
		r.configImporters[id] = importer
	}
}

func (r *DeployChainContractsRegistry) RegisterLaneVersionResolver(family string, resolver deploy.LaneVersionResolver) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.laneVersionResolver[family]; !exists {
		r.laneVersionResolver[family] = resolver
	}
}

func (r *DeployChainContractsRegistry) GetLaneVersionResolver(sel uint64) (deploy.LaneVersionResolver, bool) {
	family, err := chainsel.GetSelectorFamily(sel)
	if err != nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	resolver, ok := r.laneVersionResolver[family]
	return resolver, ok
}

func (r *DeployChainContractsRegistry) Get(family string) (DeployChainContractsAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
}

func (r *DeployChainContractsRegistry) GetConfigImporter(chainSel uint64, version *semver.Version) (deploy.ConfigImporter, bool) {
	family, err := chainsel.GetSelectorFamily(chainSel)
	if err != nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	importer, ok := r.configImporters[id]
	return importer, ok
}

func (r *DeployChainContractsRegistry) GetByChain(chainSelector uint64) (DeployChainContractsAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no deploy chain contracts adapter registered for chain family %q", family)
	}
	return adapter, nil
}
