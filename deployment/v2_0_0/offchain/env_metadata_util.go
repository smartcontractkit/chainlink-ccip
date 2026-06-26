package offchain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/shared"
)

type OffchainConfigs struct {
	NOPJobs shared.NOPJobs `json:"nopJobs,omitempty"`
}

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

func loadOrCreateCCVEnvMetadata(ds datastore.MutableDataStore) (*CCVEnvMetadata, error) {
	envMeta, err := ds.EnvMetadata().Get()
	if err != nil {
		if errors.Is(err, datastore.ErrEnvMetadataNotSet) {
			return &CCVEnvMetadata{}, nil
		}
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

// persistCCVEnvMetadata persists CCV metadata using shallow merge at the
// offchainConfigs level. Known offchainConfigs keys (nopJobs) are replaced
// fully — removing stale nested entries when jobs are removed. Unknown sibling
// keys under offchainConfigs and non-CCV top-level keys are preserved.
func persistCCVEnvMetadata(ds datastore.MutableDataStore, ccvMeta *CCVEnvMetadata) error {
	var existingMeta map[string]any
	envMeta, err := ds.EnvMetadata().Get()
	if err != nil {
		if !errors.Is(err, datastore.ErrEnvMetadataNotSet) {
			return fmt.Errorf("failed to get env metadata: %w", err)
		}
		existingMeta = make(map[string]any)
	} else if envMeta.Metadata != nil {
		data, err := json.Marshal(envMeta.Metadata)
		if err != nil {
			return fmt.Errorf("failed to marshal existing env metadata: %w", err)
		}
		if err := json.Unmarshal(data, &existingMeta); err != nil {
			return fmt.Errorf("failed to unmarshal existing env metadata: %w", err)
		}
	} else {
		existingMeta = make(map[string]any)
	}

	if ccvMeta.OffchainConfigs != nil {
		existingOC, ok := existingMeta["offchainConfigs"].(map[string]any)
		if !ok {
			existingOC = make(map[string]any)
		}

		oc := ccvMeta.OffchainConfigs
		if oc.NOPJobs != nil {
			v, err := marshalToAny(oc.NOPJobs)
			if err != nil {
				return fmt.Errorf("failed to convert NOP jobs: %w", err)
			}
			existingOC["nopJobs"] = v
		}

		existingMeta["offchainConfigs"] = existingOC
	}

	return ds.EnvMetadata().Set(datastore.EnvMetadata{Metadata: existingMeta})
}

func marshalToAny(v any) (any, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal value: %w", err)
	}
	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal value: %w", err)
	}
	return result, nil
}

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

	return persistCCVEnvMetadata(ds, ccvMeta)
}

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

	return persistCCVEnvMetadata(ds, ccvMeta)
}

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

	return persistCCVEnvMetadata(ds, ccvMeta)
}

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

			if shouldRevoke {
				orphaned = append(orphaned, job)
			}
		}
	}

	return orphaned, nil
}

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
