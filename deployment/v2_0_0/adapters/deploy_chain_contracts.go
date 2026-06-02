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
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
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
	CcvAllowlistEnabled   bool
	AllowedFinalityConfig finality.Config `json:"allowedFinalityConfig" yaml:"allowedFinalityConfig"`
}

type ExecutorDeployParams struct {
	Version       *semver.Version
	MaxCCVsPerMsg uint8
	DynamicConfig ExecutorDynamicDeployConfig
	Qualifier     string
}

type MockReceiverDeployParams struct {
	Version               *semver.Version
	RequiredVerifiers     []datastore.AddressRef
	OptionalVerifiers     []datastore.AddressRef
	OptionalThreshold     uint8
	AllowedFinalityConfig finality.Config `json:"allowedFinalityConfig" yaml:"allowedFinalityConfig"`
	Qualifier             string
}

// RMNRemoteDeployParamsOverrides holds optional RMN remote deploy overrides.
// Unset pointer fields use adapter defaults at apply time.
type RMNRemoteDeployParamsOverrides struct {
	Version   *semver.Version `json:"version,omitempty" yaml:"version,omitempty"`
	LegacyRMN *string         `json:"legacyRMN,omitempty" yaml:"legacyRMN,omitempty"`
}

// OffRampDeployParamsOverrides holds optional off-ramp deploy overrides.
type OffRampDeployParamsOverrides struct {
	Version                   *semver.Version `json:"version,omitempty" yaml:"version,omitempty"`
	GasForCallExactCheck      *uint16         `json:"gasForCallExactCheck,omitempty" yaml:"gasForCallExactCheck,omitempty"`
	MaxGasBufferToUpdateState *uint32         `json:"maxGasBufferToUpdateState,omitempty" yaml:"maxGasBufferToUpdateState,omitempty"`
}

// OnRampDeployParamsOverrides holds optional on-ramp deploy overrides.
type OnRampDeployParamsOverrides struct {
	Version               *semver.Version `json:"version,omitempty" yaml:"version,omitempty"`
	FeeAggregator         *string         `json:"feeAggregator,omitempty" yaml:"feeAggregator,omitempty"`
	MaxUSDCentsPerMessage *uint32         `json:"maxUSDCentsPerMessage,omitempty" yaml:"maxUSDCentsPerMessage,omitempty"`
}

// FeeQuoterDeployParamsOverrides holds optional fee quoter deploy overrides.
type FeeQuoterDeployParamsOverrides struct {
	Version                        *semver.Version `json:"version,omitempty" yaml:"version,omitempty"`
	MaxFeeJuelsPerMsg              *big.Int        `json:"maxFeeJuelsPerMsg,omitempty" yaml:"maxFeeJuelsPerMsg,omitempty"`
	LINKPremiumMultiplierWeiPerEth *uint64         `json:"linkPremiumMultiplierWeiPerEth,omitempty" yaml:"linkPremiumMultiplierWeiPerEth,omitempty"`
	WETHPremiumMultiplierWeiPerEth *uint64         `json:"wethPremiumMultiplierWeiPerEth,omitempty" yaml:"wethPremiumMultiplierWeiPerEth,omitempty"`
	USDPerLINK                     *big.Int        `json:"usdPerLINK,omitempty" yaml:"usdPerLINK,omitempty"`
	USDPerWETH                     *big.Int        `json:"usdPerWETH,omitempty" yaml:"usdPerWETH,omitempty"`
}

// DeployContractParamsOverrides holds optional contract deploy overrides.
// Unset pointer fields use adapter defaults at apply time.
type DeployContractParamsOverrides struct {
	RMNRemote     *RMNRemoteDeployParamsOverrides `json:"rmnRemote,omitempty" yaml:"rmnRemote,omitempty"`
	OffRamp       *OffRampDeployParamsOverrides   `json:"offRamp,omitempty" yaml:"offRamp,omitempty"`
	OnRamp        *OnRampDeployParamsOverrides    `json:"onRamp,omitempty" yaml:"onRamp,omitempty"`
	FeeQuoter     *FeeQuoterDeployParamsOverrides `json:"feeQuoter,omitempty" yaml:"feeQuoter,omitempty"`
	Executors     *[]ExecutorDeployParams         `json:"executors,omitempty" yaml:"executors,omitempty"`
	MockReceivers *[]MockReceiverDeployParams     `json:"mockReceivers,omitempty" yaml:"mockReceivers,omitempty"`
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

type DeployChainContractsOutput struct {
	sequences.OnChainOutput
	RefsToTransferOwnership []datastore.AddressRef
}

// DeployChainResolvedAddresses holds addresses resolved from the datastore (or deployed
// during resolution) before the main DeployChainContracts sequence runs.
type DeployChainResolvedAddresses struct {
	DeployerContract string
	NewAddressRefs   []datastore.AddressRef
}

// BuildDeployContractParamsInput carries topology-derived data and optional user overrides.
type BuildDeployContractParamsInput struct {
	ChainSelector      uint64
	CommitteeVerifiers []CommitteeVerifierDeployParams
	Defaults           DeployContractParams
	Overrides          *DeployContractParamsOverrides
}

// ApplyDeployContractParamsOverrides merges optional user overrides onto adapter defaults.
func ApplyDeployContractParamsOverrides(params DeployContractParams, overrides *DeployContractParamsOverrides) DeployContractParams {
	if overrides == nil {
		return params
	}
	if overrides.RMNRemote != nil {
		o := overrides.RMNRemote
		params.RMNRemote.Version = utils.CoalescePtr(o.Version, params.RMNRemote.Version)
		params.RMNRemote.LegacyRMN = utils.Coalesce(o.LegacyRMN, params.RMNRemote.LegacyRMN)
	}
	if overrides.OffRamp != nil {
		o := overrides.OffRamp
		params.OffRamp.Version = utils.CoalescePtr(o.Version, params.OffRamp.Version)
		params.OffRamp.GasForCallExactCheck = utils.Coalesce(o.GasForCallExactCheck, params.OffRamp.GasForCallExactCheck)
		params.OffRamp.MaxGasBufferToUpdateState = utils.Coalesce(o.MaxGasBufferToUpdateState, params.OffRamp.MaxGasBufferToUpdateState)
	}
	if overrides.OnRamp != nil {
		o := overrides.OnRamp
		params.OnRamp.Version = utils.CoalescePtr(o.Version, params.OnRamp.Version)
		params.OnRamp.FeeAggregator = utils.Coalesce(o.FeeAggregator, params.OnRamp.FeeAggregator)
		params.OnRamp.MaxUSDCentsPerMessage = utils.Coalesce(o.MaxUSDCentsPerMessage, params.OnRamp.MaxUSDCentsPerMessage)
	}
	if overrides.FeeQuoter != nil {
		o := overrides.FeeQuoter
		params.FeeQuoter.Version = utils.CoalescePtr(o.Version, params.FeeQuoter.Version)
		params.FeeQuoter.MaxFeeJuelsPerMsg = utils.CoalescePtr(o.MaxFeeJuelsPerMsg, params.FeeQuoter.MaxFeeJuelsPerMsg)
		params.FeeQuoter.LINKPremiumMultiplierWeiPerEth = utils.Coalesce(o.LINKPremiumMultiplierWeiPerEth, params.FeeQuoter.LINKPremiumMultiplierWeiPerEth)
		params.FeeQuoter.WETHPremiumMultiplierWeiPerEth = utils.Coalesce(o.WETHPremiumMultiplierWeiPerEth, params.FeeQuoter.WETHPremiumMultiplierWeiPerEth)
		params.FeeQuoter.USDPerLINK = utils.CoalescePtr(o.USDPerLINK, params.FeeQuoter.USDPerLINK)
		params.FeeQuoter.USDPerWETH = utils.CoalescePtr(o.USDPerWETH, params.FeeQuoter.USDPerWETH)
	}
	if overrides.Executors != nil {
		params.Executors = *overrides.Executors
	}
	if overrides.MockReceivers != nil {
		params.MockReceivers = *overrides.MockReceivers
	}
	return params
}

type DeployChainContractsAdapter interface {
	GetDefaultDeployContractParams(chainSelector uint64) DeployContractParams
	ResolveDeployAddresses(e deployment.Environment, chainSelector uint64) (DeployChainResolvedAddresses, error)
	BuildDeployContractParams(input BuildDeployContractParamsInput) (DeployContractParams, error)
	DeployChainContracts() *cldf_ops.Sequence[DeployChainContractsInput, DeployChainContractsOutput, cldf_chain.BlockChains]
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
