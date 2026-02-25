package ccv

import (
	"encoding/json"
	"fmt"
	"maps"

	jsonpatch "github.com/evanphx/json-patch/v5"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// OffchainConfigs contains generated configurations for offchain services.
// Uses the types from the aggregator and indexer packages directly.
type OffchainConfigs struct {
	// Aggregators maps service identifier (e.g., "default-aggregator") to generated committee config.
	Aggregators map[string]*offchain.Committee `json:"aggregators,omitempty"`
	// Indexers maps service identifier (e.g., "indexer") to generated verifier config.
	Indexers map[string]*offchain.IndexerGeneratedConfig `json:"indexers,omitempty"`
	// TokenVerifiers maps service identifier (e.g., "default-token-verifier") to generated token verifier config.
	TokenVerifiers map[string]*offchain.TokenVerifierConfig `json:"tokenVerifiers,omitempty"`
	// NOPJobs maps NOP alias to a map of job ID to job info.
	// This tracks the full lifecycle of jobs including proposal history and JD IDs.
	NOPJobs shared.NOPJobs `json:"nopJobs,omitempty"`
}

// CCVEnvMetadata represents the expected structure of env_metadata.json for CCV.
// OffchainConfigs stores generated configs after scanning on-chain state.
type CCVEnvMetadata struct {
	OffchainConfigs *OffchainConfigs `json:"offchainConfigs,omitempty"`
}

func loadCCVEnvMetadata(ds datastore.DataStore) (*CCVEnvMetadata, error) {
	envMeta, err := ds.EnvMetadata().Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get env metadata: %w", err)
	}
	return parseCCVEnvMetadata(envMeta.Metadata)
}

func parseCCVEnvMetadata(metadata any) (*CCVEnvMetadata, error) {
	if metadata == nil {
		return &CCVEnvMetadata{}, nil
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal env metadata: %w", err)
	}

	var ccvMeta CCVEnvMetadata
	if err := json.Unmarshal(data, &ccvMeta); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CCV env metadata: %w", err)
	}

	return &ccvMeta, nil
}

// SaveAggregatorConfig saves an aggregator committee config to the datastore under the given service identifier.
func SaveAggregatorConfig(ds datastore.MutableDataStore, serviceIdentifier string, cfg *offchain.Committee) error {
	ccvMeta, err := loadOrCreateCCVEnvMetadata(ds)
	if err != nil {
		return err
	}

	if ccvMeta.OffchainConfigs == nil {
		ccvMeta.OffchainConfigs = &OffchainConfigs{}
	}
	if ccvMeta.OffchainConfigs.Aggregators == nil {
		ccvMeta.OffchainConfigs.Aggregators = make(map[string]*offchain.Committee)
	}

	ccvMeta.OffchainConfigs.Aggregators[serviceIdentifier] = cfg

	return saveCCVEnvMetadata(ds, ccvMeta)
}

// SaveIndexerConfig saves an indexer generated config to the datastore under the given service identifier.
func SaveIndexerConfig(ds datastore.MutableDataStore, serviceIdentifier string, cfg *offchain.IndexerGeneratedConfig) error {
	ccvMeta, err := loadOrCreateCCVEnvMetadata(ds)
	if err != nil {
		return err
	}

	if ccvMeta.OffchainConfigs == nil {
		ccvMeta.OffchainConfigs = &OffchainConfigs{}
	}
	if ccvMeta.OffchainConfigs.Indexers == nil {
		ccvMeta.OffchainConfigs.Indexers = make(map[string]*offchain.IndexerGeneratedConfig)
	}

	ccvMeta.OffchainConfigs.Indexers[serviceIdentifier] = cfg

	return saveCCVEnvMetadata(ds, ccvMeta)
}

// SaveTokenVerifierConfig saves a token verifier config to the datastore under the given service identifier.
func SaveTokenVerifierConfig(ds datastore.MutableDataStore, serviceIdentifier string, cfg *offchain.TokenVerifierConfig) error {
	ccvMeta, err := loadOrCreateCCVEnvMetadata(ds)
	if err != nil {
		return err
	}

	if ccvMeta.OffchainConfigs == nil {
		ccvMeta.OffchainConfigs = &OffchainConfigs{}
	}
	if ccvMeta.OffchainConfigs.TokenVerifiers == nil {
		ccvMeta.OffchainConfigs.TokenVerifiers = make(map[string]*offchain.TokenVerifierConfig)
	}

	ccvMeta.OffchainConfigs.TokenVerifiers[serviceIdentifier] = cfg

	return saveCCVEnvMetadata(ds, ccvMeta)
}

// GetAggregatorConfig retrieves an aggregator committee config from the datastore by service identifier.
func GetAggregatorConfig(ds datastore.DataStore, serviceIdentifier string) (*offchain.Committee, error) {
	ccvMeta, err := loadCCVEnvMetadata(ds)
	if err != nil {
		return nil, err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.Aggregators == nil {
		return nil, fmt.Errorf("no aggregator configs found")
	}

	cfg, ok := ccvMeta.OffchainConfigs.Aggregators[serviceIdentifier]
	if !ok {
		return nil, fmt.Errorf("aggregator config %q not found", serviceIdentifier)
	}

	return cfg, nil
}

// GetIndexerConfig retrieves an indexer generated config from the datastore by service identifier.
func GetIndexerConfig(ds datastore.DataStore, serviceIdentifier string) (*offchain.IndexerGeneratedConfig, error) {
	ccvMeta, err := loadCCVEnvMetadata(ds)
	if err != nil {
		return nil, err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.Indexers == nil {
		return nil, fmt.Errorf("no indexer configs found")
	}

	cfg, ok := ccvMeta.OffchainConfigs.Indexers[serviceIdentifier]
	if !ok {
		return nil, fmt.Errorf("indexer config %q not found", serviceIdentifier)
	}

	return cfg, nil
}

// GetTokenVerifierConfig retrieves a token verifier config from the datastore by service identifier.
func GetTokenVerifierConfig(ds datastore.DataStore, serviceIdentifier string) (*offchain.TokenVerifierConfig, error) {
	ccvMeta, err := loadCCVEnvMetadata(ds)
	if err != nil {
		return nil, err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.TokenVerifiers == nil {
		return nil, fmt.Errorf("no token verifier configs found")
	}

	cfg, ok := ccvMeta.OffchainConfigs.TokenVerifiers[serviceIdentifier]
	if !ok {
		return nil, fmt.Errorf("token verifier config %q not found", serviceIdentifier)
	}

	return cfg, nil
}

// SaveJob saves a single job to the datastore.
func SaveJob(ds datastore.MutableDataStore, job shared.JobInfo) error {
	ccvMeta, err := loadOrCreateCCVEnvMetadata(ds)
	if err != nil {
		return err
	}

	if ccvMeta.OffchainConfigs == nil {
		ccvMeta.OffchainConfigs = &OffchainConfigs{}
	}
	if ccvMeta.OffchainConfigs.NOPJobs == nil {
		ccvMeta.OffchainConfigs.NOPJobs = make(shared.NOPJobs)
	}
	if ccvMeta.OffchainConfigs.NOPJobs[job.NOPAlias] == nil {
		ccvMeta.OffchainConfigs.NOPJobs[job.NOPAlias] = make(map[shared.JobID]shared.JobInfo)
	}

	ccvMeta.OffchainConfigs.NOPJobs[job.NOPAlias][job.JobID] = job

	return saveCCVEnvMetadata(ds, ccvMeta)
}

// SaveJobs saves multiple jobs to the datastore in a single operation.
func SaveJobs(ds datastore.MutableDataStore, jobs []shared.JobInfo) error {
	if len(jobs) == 0 {
		return nil
	}

	ccvMeta, err := loadOrCreateCCVEnvMetadata(ds)
	if err != nil {
		return err
	}

	if ccvMeta.OffchainConfigs == nil {
		ccvMeta.OffchainConfigs = &OffchainConfigs{}
	}
	if ccvMeta.OffchainConfigs.NOPJobs == nil {
		ccvMeta.OffchainConfigs.NOPJobs = make(shared.NOPJobs)
	}

	for _, job := range jobs {
		if ccvMeta.OffchainConfigs.NOPJobs[job.NOPAlias] == nil {
			ccvMeta.OffchainConfigs.NOPJobs[job.NOPAlias] = make(map[shared.JobID]shared.JobInfo)
		}
		ccvMeta.OffchainConfigs.NOPJobs[job.NOPAlias][job.JobID] = job
	}

	return saveCCVEnvMetadata(ds, ccvMeta)
}

// GetJob retrieves a specific job from the datastore.
func GetJob(ds datastore.DataStore, nopAlias shared.NOPAlias, jobID shared.JobID) (*shared.JobInfo, error) {
	ccvMeta, err := loadCCVEnvMetadata(ds)
	if err != nil {
		return nil, err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.NOPJobs == nil {
		return nil, fmt.Errorf("no jobs found")
	}

	nopJobs, ok := ccvMeta.OffchainConfigs.NOPJobs[nopAlias]
	if !ok {
		return nil, fmt.Errorf("no jobs found for NOP %q", nopAlias)
	}

	job, ok := nopJobs[jobID]
	if !ok {
		return nil, fmt.Errorf("job %q not found for NOP %q", jobID, nopAlias)
	}

	return &job, nil
}

// GetJobsByNOP retrieves all jobs for a given NOP alias.
func GetJobsByNOP(ds datastore.DataStore, nopAlias shared.NOPAlias) (map[shared.JobID]shared.JobInfo, error) {
	ccvMeta, err := loadCCVEnvMetadata(ds)
	if err != nil {
		return nil, err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.NOPJobs == nil {
		return nil, fmt.Errorf("no jobs found")
	}

	nopJobs, ok := ccvMeta.OffchainConfigs.NOPJobs[nopAlias]
	if !ok {
		return nil, fmt.Errorf("no jobs found for NOP %q", nopAlias)
	}

	return nopJobs, nil
}

// GetAllJobs retrieves all jobs from the datastore.
func GetAllJobs(ds datastore.DataStore) (shared.NOPJobs, error) {
	ccvMeta, err := loadCCVEnvMetadata(ds)
	if err != nil {
		return nil, err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.NOPJobs == nil {
		return make(shared.NOPJobs), nil
	}

	return ccvMeta.OffchainConfigs.NOPJobs, nil
}

// GetJobByExternalID finds a job by its external job ID.
func GetJobByExternalID(ds datastore.DataStore, externalJobID string) (*shared.JobInfo, error) {
	allJobs, err := GetAllJobs(ds)
	if err != nil {
		return nil, err
	}

	for _, nopJobs := range allJobs {
		for _, job := range nopJobs {
			if job.ExternalJobID == externalJobID {
				return &job, nil
			}
		}
	}

	return nil, fmt.Errorf("job with external ID %q not found", externalJobID)
}

// GetJobByJDJobID finds a job by JD's job ID.
func GetJobByJDJobID(ds datastore.DataStore, jdJobID string) (*shared.JobInfo, error) {
	allJobs, err := GetAllJobs(ds)
	if err != nil {
		return nil, err
	}

	for _, nopJobs := range allJobs {
		for _, job := range nopJobs {
			if job.JDJobID == jdJobID {
				return &job, nil
			}
		}
	}

	return nil, fmt.Errorf("job with JD job ID %q not found", jdJobID)
}

// DeleteJob removes a job from the datastore.
func DeleteJob(ds datastore.MutableDataStore, nopAlias shared.NOPAlias, jobID shared.JobID) error {
	ccvMeta, err := loadOrCreateCCVEnvMetadata(ds)
	if err != nil {
		return err
	}

	if ccvMeta.OffchainConfigs == nil || ccvMeta.OffchainConfigs.NOPJobs == nil {
		return nil
	}

	nopJobs, ok := ccvMeta.OffchainConfigs.NOPJobs[nopAlias]
	if !ok {
		return nil
	}

	if _, ok := nopJobs[jobID]; !ok {
		return nil
	}

	delete(ccvMeta.OffchainConfigs.NOPJobs[nopAlias], jobID)

	if len(ccvMeta.OffchainConfigs.NOPJobs[nopAlias]) == 0 {
		delete(ccvMeta.OffchainConfigs.NOPJobs, nopAlias)
	}

	return replaceCCVEnvMetadata(ds, ccvMeta)
}

// CollectOrphanedJobs finds jobs that should be revoked based on the current topology.
// This returns jobs that exist in the metadata but are no longer expected (e.g., NOP removed from committee).
func CollectOrphanedJobs(
	ds datastore.DataStore,
	scope shared.JobScope,
	expectedJobsByNOP map[shared.NOPAlias]map[shared.JobID]bool,
	scopedNOPs map[shared.NOPAlias]bool,
	environmentNOPs map[shared.NOPAlias]bool,
) ([]shared.JobInfo, error) {
	allJobs, err := GetAllJobs(ds)
	if err != nil {
		return nil, fmt.Errorf("failed to get all jobs for cleanup: %w", err)
	}

	orphaned := make([]shared.JobInfo, 0)
	for nopAlias, nopJobs := range allJobs {
		if scopedNOPs != nil && !scopedNOPs[nopAlias] {
			continue
		}

		for jobID, job := range nopJobs {
			if !scope.IsJobInScope(jobID) {
				continue
			}

			nopExpectedJobs := expectedJobsByNOP[nopAlias]
			shouldRevoke := nopExpectedJobs == nil || !nopExpectedJobs[jobID]
			if environmentNOPs != nil && !environmentNOPs[nopAlias] {
				shouldRevoke = true
			}

			if shouldRevoke && job.LatestStatus() != shared.JobProposalStatusRevoked {
				orphaned = append(orphaned, job)
			}
		}
	}

	return orphaned, nil
}

// CleanupOrphanedJobs removes jobs from the metadata after they have been revoked.
func CleanupOrphanedJobs(
	ds datastore.MutableDataStore,
	jobs []shared.JobInfo,
) error {
	for _, job := range jobs {
		if err := DeleteJob(ds, job.NOPAlias, job.JobID); err != nil {
			return fmt.Errorf("failed to delete job %q for NOP %q: %w", job.JobID, job.NOPAlias, err)
		}
	}
	return nil
}

// replaceCCVEnvMetadata replaces the CCV metadata completely (not merge).
// This is needed for delete operations since JSON Merge Patch doesn't remove missing keys.
func replaceCCVEnvMetadata(ds datastore.MutableDataStore, ccvMeta *CCVEnvMetadata) error {
	// Get existing metadata to preserve non-CCV fields
	var existingMeta map[string]any
	if envMeta, err := ds.EnvMetadata().Get(); err == nil && envMeta.Metadata != nil {
		data, err := json.Marshal(envMeta.Metadata)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(data, &existingMeta); err != nil {
			return err
		}
	} else {
		existingMeta = make(map[string]any)
	}

	// Marshal the CCV metadata
	ccvData, err := json.Marshal(ccvMeta)
	if err != nil {
		return err
	}

	var ccvMap map[string]any
	if err := json.Unmarshal(ccvData, &ccvMap); err != nil {
		return err
	}

	// Replace CCV-specific fields (offchainConfigs) completely
	maps.Copy(existingMeta, ccvMap)

	return ds.EnvMetadata().Set(datastore.EnvMetadata{Metadata: existingMeta})
}

func loadOrCreateCCVEnvMetadata(ds datastore.MutableDataStore) (*CCVEnvMetadata, error) {
	envMeta, err := ds.EnvMetadata().Get()
	if err != nil {
		return &CCVEnvMetadata{}, nil
	}
	return parseCCVEnvMetadata(envMeta.Metadata)
}

func saveCCVEnvMetadata(ds datastore.MutableDataStore, ccvMeta *CCVEnvMetadata) error {
	var base json.RawMessage = []byte(`{}`)

	if envMeta, err := ds.EnvMetadata().Get(); err == nil && envMeta.Metadata != nil {
		b, err := json.Marshal(envMeta.Metadata)
		if err != nil {
			return err
		}
		base = b
	}

	patch, err := json.Marshal(ccvMeta)
	if err != nil {
		return err
	}

	merged, err := jsonpatch.MergePatch(base, patch)
	if err != nil {
		return err
	}

	var result map[string]any
	if err := json.Unmarshal(merged, &result); err != nil {
		return err
	}

	return ds.EnvMetadata().Set(datastore.EnvMetadata{Metadata: result})
}
