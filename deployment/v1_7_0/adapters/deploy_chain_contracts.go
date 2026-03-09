package adapters

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

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
	Version           *semver.Version
	RequiredVerifiers []datastore.AddressRef
	OptionalVerifiers []datastore.AddressRef
	OptionalThreshold uint8
	Qualifier         string
}

type MCMSInstanceDeployParams struct {
	Proposer         mcms_types.Config
	Bypasser         mcms_types.Config
	Canceller        mcms_types.Config
	TimelockMinDelay *big.Int
	TimelockAdmin    common.Address
	Label            *string
}

type MCMSDeployParams struct {
	CLLCCIP MCMSInstanceDeployParams
	RMNMCMS MCMSInstanceDeployParams
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

type DeployChainContractsInput struct {
	ChainSelector     uint64
	DeployerContract  string
	DeployTestRouter  bool
	ExistingAddresses []datastore.AddressRef
	ContractParams    DeployContractParams
	// MCMS configures deployment of CLLCCIP and RMNMCMS instances.
	// When non-nil, both instances are deployed and ownership of product contracts
	// is transferred to the CLLCCIP timelock.
	MCMS *MCMSDeployParams
}

type DeployChainContractsAdapter interface {
	DeployChainContracts() *cldf_ops.Sequence[DeployChainContractsInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type DeployChainContractsRegistry struct {
	mu       sync.Mutex
	adapters map[string]DeployChainContractsAdapter
}

var (
	singletonDeployChainContractsRegistry *DeployChainContractsRegistry
	deployChainContractsRegistryOnce      sync.Once
)

func NewDeployChainContractsRegistry() *DeployChainContractsRegistry {
	return &DeployChainContractsRegistry{
		adapters: make(map[string]DeployChainContractsAdapter),
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

func (r *DeployChainContractsRegistry) Get(family string) (DeployChainContractsAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
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
