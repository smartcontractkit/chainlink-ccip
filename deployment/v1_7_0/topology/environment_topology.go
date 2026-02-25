package topology

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

// EnvironmentTopology holds all environment-specific configuration that cannot be inferred
// from the datastore. This serves as the single source of truth for the desired state of both off-chain
// (job specs) and on-chain (committee contracts) configuration.
type EnvironmentTopology struct {
	IndexerAddress []string                      `toml:"indexer_address"`
	PyroscopeURL   string                        `toml:"pyroscope_url"`
	Monitoring     MonitoringConfig              `toml:"monitoring"`
	NOPTopology    *NOPTopology                  `toml:"nop_topology"`
	ExecutorPools  map[string]ExecutorPoolConfig `toml:"executor_pools"`
}

// NOPTopology defines the node operator structure and committee membership.
// This is the single source of truth for both off-chain (job specs) and
// on-chain (committee contracts) configuration.
// NOPs are stored as a slice to preserve declaration order, which determines
// the node index for CL node assignment.
type NOPTopology struct {
	NOPs       []NOPConfig                `toml:"nops"`
	Committees map[string]CommitteeConfig `toml:"committees"`

	nopIndex map[string]int // internal index lookups, built on first access
}

// GetNOP returns the NOPConfig for the given alias.
func (t *NOPTopology) GetNOP(alias string) (NOPConfig, bool) {
	t.ensureIndex()
	idx, ok := t.nopIndex[alias]
	if !ok {
		return NOPConfig{}, false
	}
	return t.NOPs[idx], true
}

// GetNOPIndex returns the index of the NOP with the given alias.
// The index corresponds to the declaration order in the topology config,
// which maps to CL node indices.
func (t *NOPTopology) GetNOPIndex(alias string) (int, bool) {
	t.ensureIndex()
	idx, ok := t.nopIndex[alias]
	return idx, ok
}

// SetNOPSignerAddress sets the signer address for the NOP with the given alias.
// Returns true if the NOP was found and updated, false otherwise.
func (t *NOPTopology) SetNOPSignerAddress(alias, family, addr string) bool {
	t.ensureIndex()
	idx, ok := t.nopIndex[alias]
	if !ok {
		return false
	}

	if t.NOPs[idx].SignerAddressByFamily == nil {
		t.NOPs[idx].SignerAddressByFamily = make(map[string]string)
	}

	t.NOPs[idx].SignerAddressByFamily[family] = addr
	return true
}

// HasNOP returns true if a NOP with the given alias exists.
func (t *NOPTopology) HasNOP(alias string) bool {
	t.ensureIndex()
	_, ok := t.nopIndex[alias]
	return ok
}

// ensureIndex builds the internal alias-to-index map if not already built.
func (t *NOPTopology) ensureIndex() {
	if t.nopIndex != nil {
		return
	}
	t.nopIndex = make(map[string]int, len(t.NOPs))
	for i, nop := range t.NOPs {
		t.nopIndex[nop.Alias] = i
	}
}

// NOPConfig defines a Node Operator.
// Each NOP runs exactly one node. The NOP alias serves as the node/executor ID.
// For production: SignerAddress is resolved from e.Nodes at deployment time.
// For devenv: SignerAddress can be set directly in the config.
type NOPConfig struct {
	Alias                 string            `toml:"alias"`
	Name                  string            `toml:"name"`
	SignerAddressByFamily map[string]string `toml:"signer_address_by_family,omitempty"`
	Mode                  shared.NOPMode    `toml:"mode,omitempty"`
}

func (n *NOPConfig) GetMode() shared.NOPMode {
	if n.Mode == "" {
		return shared.NOPModeCL
	}
	return n.Mode
}

// CommitteeConfig defines a committee and its per-chain membership.
// It contains off-chain (aggregators, chain configs) and committee-wide
// on-chain (storage_locations) configuration.
type CommitteeConfig struct {
	Qualifier        string                          `toml:"qualifier"`
	VerifierVersion  *semver.Version                 `toml:"verifier_version"`
	StorageLocations []string                        `toml:"storage_locations,omitempty"`
	ChainConfigs     map[string]ChainCommitteeConfig `toml:"chain_configs"`
	Aggregators      []AggregatorConfig              `toml:"aggregators"`
}

// ChainCommitteeConfig defines committee membership and on-chain parameters
// for a specific chain. FeeAggregator and AllowlistAdmin are per-chain because
// address formats differ across chain families (EVM, Canton, etc.).
type ChainCommitteeConfig struct {
	NOPAliases     []string `toml:"nop_aliases"`
	Threshold      uint8    `toml:"threshold"`
	FeeAggregator  string   `toml:"fee_aggregator,omitempty"`
	AllowlistAdmin string   `toml:"allowlist_admin,omitempty"`
}

// AggregatorConfig defines an aggregator instance for HA setups.
type AggregatorConfig struct {
	Name                         string `toml:"name"`
	Address                      string `toml:"address"`
	InsecureAggregatorConnection bool   `toml:"insecure_connection"`
}

// ExecutorPoolConfig defines executor pool membership and configuration.
type ExecutorPoolConfig struct {
	NOPAliases        []string      `toml:"nop_aliases"`
	ExecutionInterval time.Duration `toml:"execution_interval"`
	IndexerQueryLimit uint64        `toml:"indexer_query_limit"`
	BackoffDuration   time.Duration `toml:"backoff_duration"`
	LookbackWindow    time.Duration `toml:"lookback_window"`
	ReaderCacheExpiry time.Duration `toml:"reader_cache_expiry"`
	MaxRetryDuration  time.Duration `toml:"max_retry_duration"`
	WorkerCount       int           `toml:"worker_count"`
	NtpServer         string        `toml:"ntp_server"`
}

// MonitoringConfig provides monitoring configuration shared across services.
type MonitoringConfig struct {
	Enabled  bool           `toml:"Enabled"`
	Type     string         `toml:"Type"`
	Beholder BeholderConfig `toml:"Beholder"`
}

// BeholderConfig wraps OpenTelemetry configuration for the beholder client.
type BeholderConfig struct {
	InsecureConnection       bool    `toml:"InsecureConnection"`
	CACertFile               string  `toml:"CACertFile"`
	OtelExporterGRPCEndpoint string  `toml:"OtelExporterGRPCEndpoint"`
	OtelExporterHTTPEndpoint string  `toml:"OtelExporterHTTPEndpoint"`
	LogStreamingEnabled      bool    `toml:"LogStreamingEnabled"`
	MetricReaderInterval     int64   `toml:"MetricReaderInterval"`
	TraceSampleRatio         float64 `toml:"TraceSampleRatio"`
	TraceBatchTimeout        int64   `toml:"TraceBatchTimeout"`
}

// LoadEnvironmentTopology loads an EnvironmentTopology from a TOML file.
func LoadEnvironmentTopology(path string) (*EnvironmentTopology, error) {
	data, err := os.ReadFile(path) //nolint:gosec // G304: path is provided by trusted caller
	if err != nil {
		return nil, fmt.Errorf("failed to read environment topology file: %w", err)
	}

	var cfg EnvironmentTopology
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment topology TOML: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("environment topology validation failed: %w", err)
	}

	return &cfg, nil
}

// WriteEnvironmentTopology writes an EnvironmentTopology to a TOML file.
func WriteEnvironmentTopology(path string, cfg EnvironmentTopology) error {
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("environment topology validation failed: %w", err)
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal environment topology to TOML: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("failed to write environment topology file: %w", err)
	}

	return nil
}

// Validate validates the EnvironmentTopology.
func (c *EnvironmentTopology) Validate() error {
	if len(c.IndexerAddress) == 0 {
		return fmt.Errorf("indexer_address is required")
	}
	// Ensure indexer addresses are unique
	addressSet := make(map[string]struct{}, len(c.IndexerAddress))
	for _, addr := range c.IndexerAddress {
		if addr == "" {
			return fmt.Errorf("indexer_address is required")
		}
		if _, exists := addressSet[addr]; exists {
			return fmt.Errorf("duplicate indexer_address found: %q", addr)
		}
		addressSet[addr] = struct{}{}
	}

	if err := c.NOPTopology.Validate(); err != nil {
		return fmt.Errorf("nop_topology validation failed: %w", err)
	}

	for poolName, pool := range c.ExecutorPools {
		if err := pool.Validate(poolName, c.NOPTopology); err != nil {
			return fmt.Errorf("executor_pool %q validation failed: %w", poolName, err)
		}
	}

	return nil
}

// Validate validates the NOPTopology.
func (t *NOPTopology) Validate() error {
	seen := make(map[string]struct{}, len(t.NOPs))
	for _, nop := range t.NOPs {
		if nop.Alias == "" {
			return fmt.Errorf("NOP alias is required")
		}
		if _, exists := seen[nop.Alias]; exists {
			return fmt.Errorf("duplicate NOP alias %q", nop.Alias)
		}
		seen[nop.Alias] = struct{}{}
		if nop.Name == "" {
			return fmt.Errorf("NOP %q name is required", nop.Alias)
		}
	}

	for qualifier, committee := range t.Committees {
		if committee.Qualifier != qualifier {
			return fmt.Errorf("committee qualifier mismatch: key %q != qualifier %q", qualifier, committee.Qualifier)
		}
		if err := committee.Validate(t); err != nil {
			return fmt.Errorf("committee %q validation failed: %w", qualifier, err)
		}
	}

	return nil
}

// Validate validates the CommitteeConfig.
func (c *CommitteeConfig) Validate(topology *NOPTopology) error {
	if len(c.Aggregators) == 0 {
		return fmt.Errorf("at least one aggregator is required")
	}

	for _, agg := range c.Aggregators {
		if agg.Name == "" {
			return fmt.Errorf("aggregator name is required")
		}
		if agg.Address == "" {
			return fmt.Errorf("aggregator %q address is required", agg.Name)
		}
	}

	for chainSelector, chainCfg := range c.ChainConfigs {
		if len(chainCfg.NOPAliases) == 0 {
			return fmt.Errorf("chain %q requires at least one NOP", chainSelector)
		}
		for _, alias := range chainCfg.NOPAliases {
			if !topology.HasNOP(alias) {
				return fmt.Errorf("chain %q references unknown NOP alias %q", chainSelector, alias)
			}
		}
		if chainCfg.Threshold == 0 {
			return fmt.Errorf("chain %q threshold must be greater than 0", chainSelector)
		}
		if int(chainCfg.Threshold) > len(chainCfg.NOPAliases) {
			return fmt.Errorf("chain %q threshold %d exceeds NOP count %d", chainSelector, chainCfg.Threshold, len(chainCfg.NOPAliases))
		}
	}

	return nil
}

// Validate validates the ExecutorPoolConfig.
func (p *ExecutorPoolConfig) Validate(poolName string, topology *NOPTopology) error {
	if len(p.NOPAliases) == 0 {
		return fmt.Errorf("executor pool requires at least one NOP")
	}

	for _, alias := range p.NOPAliases {
		if !topology.HasNOP(alias) {
			return fmt.Errorf("executor pool references unknown NOP alias %q", alias)
		}
	}

	return nil
}

// GetNOPsForPool returns the NOP aliases for a given executor pool.
func (c *EnvironmentTopology) GetNOPsForPool(poolName string) ([]string, error) {
	pool, ok := c.ExecutorPools[poolName]
	if !ok {
		return nil, fmt.Errorf("executor pool %q not found", poolName)
	}
	return pool.NOPAliases, nil
}

// GetNOPsForCommittee returns the NOP aliases that are members of a committee on any chain.
func (c *EnvironmentTopology) GetNOPsForCommittee(committeeQualifier string) ([]string, error) {
	committee, ok := c.NOPTopology.Committees[committeeQualifier]
	if !ok {
		return nil, fmt.Errorf("committee %q not found", committeeQualifier)
	}

	nopSet := make(map[string]struct{})
	for _, chainCfg := range committee.ChainConfigs {
		for _, alias := range chainCfg.NOPAliases {
			nopSet[alias] = struct{}{}
		}
	}

	nops := make([]string, 0, len(nopSet))
	for alias := range nopSet {
		nops = append(nops, alias)
	}

	return nops, nil
}

// GetCommitteesForNOP returns the committee qualifiers that include a NOP on any chain.
func (c *EnvironmentTopology) GetCommitteesForNOP(nopAlias string) []string {
	var committees []string
	for qualifier, committee := range c.NOPTopology.Committees {
		for _, chainCfg := range committee.ChainConfigs {
			if slices.Contains(chainCfg.NOPAliases, nopAlias) {
				committees = append(committees, qualifier)
				break
			}
		}
	}
	return committees
}

// GetPoolsForNOP returns the executor pool names that include a NOP.
func (c *EnvironmentTopology) GetPoolsForNOP(nopAlias string) []string {
	var pools []string
	for poolName, pool := range c.ExecutorPools {
		if slices.Contains(pool.NOPAliases, nopAlias) {
			pools = append(pools, poolName)
		}
	}
	return pools
}

func (c *EnvironmentTopology) GetAggregatorNamesForCommittee(name string) ([]string, error) {
	for _, committee := range c.NOPTopology.Committees {
		if strings.EqualFold(committee.Qualifier, name) {
			names := make([]string, len(committee.Aggregators))
			for i, agg := range committee.Aggregators {
				names[i] = agg.Name
			}
			return names, nil
		}
	}
	return nil, fmt.Errorf("no aggregators found for committee: %s", name)
}
