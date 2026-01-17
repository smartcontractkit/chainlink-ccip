package chainaccessor

import (
	"encoding/binary"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

func processConfigResults(
	lggr logger.Logger,
	resultsChainSelector cciptypes.ChainSelector,
	destChainSelector cciptypes.ChainSelector,
	batchResult types.BatchGetLatestValuesResult,
) (cciptypes.ChainConfigSnapshot, error) {
	config := cciptypes.ChainConfigSnapshot{}

	for contract, results := range batchResult {
		var err error
		switch contract.Name {
		case consts.ContractNameOffRamp:
			config.Offramp, err = processOfframpResults(results)
		case consts.ContractNameRMNProxy:
			config.RMNProxy, err = processRMNProxyResults(results)
		case consts.ContractNameRMNRemote:
			config.RMNRemote, config.CurseInfo, err = processRMNRemoteResults(results, destChainSelector)
		case consts.ContractNameFeeQuoter:
			config.FeeQuoter, err = processFeeQuoterResults(results)
		case consts.ContractNameOnRamp:
			// Only process OnRamp results for source chains
			if resultsChainSelector != destChainSelector {
				config.OnRamp, err = processOnRampResults(results)
			}
		case consts.ContractNameRouter:
			// Only process Router results for source chains
			if resultsChainSelector != destChainSelector {
				config.Router, err = processRouterResults(results)
			}
		default:
			lggr.Warnw("Unhandled contract in batch results", "chain", resultsChainSelector, "contract", contract.Name)
		}
		if err != nil {
			return cciptypes.ChainConfigSnapshot{}, fmt.Errorf("process %s results: %w", contract.Name, err)
		}
	}

	return config, nil
}

// processSourceChainConfigResults extracts and processes source chain config results from the batch
func processSourceChainConfigResults(
	lggr logger.Logger,
	batchResult types.BatchGetLatestValuesResult,
	standardOffRampRequestCount int,
	filteredSourceChains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.SourceChainConfig {
	sourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	// Find the OffRamp results
	for contract, results := range batchResult {
		if contract.Name == consts.ContractNameOffRamp && len(results) > standardOffRampRequestCount {
			// Extract just the source chain results (everything after standard results)
			sourceChainResults := results[standardOffRampRequestCount:]

			if len(sourceChainResults) != len(filteredSourceChains) {
				lggr.Warnw("Source chain result count mismatch",
					"expected", len(filteredSourceChains),
					"got", len(sourceChainResults))
			} else {
				// Process each source chain result
				for i, chain := range filteredSourceChains {
					if i >= len(sourceChainResults) {
						continue
					}

					v, err := sourceChainResults[i].GetResult()
					if err != nil {
						lggr.Errorw("Failed to get source chain config from result",
							"chain", chain,
							"error", err)
						continue
					}

					cfg, ok := v.(*cciptypes.SourceChainConfig)
					if !ok {
						lggr.Errorw("Invalid result type from GetSourceChainConfig",
							"chain", chain,
							"type", fmt.Sprintf("%T", v))
						continue
					}

					// Store the config
					sourceConfigs[chain] = *cfg
				}
			}

			break // Found and processed the OffRamp results
		}
	}

	return sourceConfigs
}

func processOfframpResults(
	results []types.BatchReadResult) (cciptypes.OfframpConfig, error) {

	if len(results) != 4 {
		return cciptypes.OfframpConfig{}, fmt.Errorf("expected 4 offramp results, got %d", len(results))
	}

	config := cciptypes.OfframpConfig{}

	// Define processors for each expected result
	processors := []resultProcessor{
		// CommitLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for CommitLatestOCRConfig: %T", val)
			}
			config.CommitLatestOCRConfig = *typed
			return nil
		},
		// ExecLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for ExecLatestOCRConfig: %T", val)
			}
			config.ExecLatestOCRConfig = *typed
			return nil
		},
		// StaticConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OffRampStaticChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for StaticConfig: %T", val)
			}
			config.StaticConfig = *typed
			return nil
		},
		// DynamicConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OffRampDynamicChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for DynamicConfig: %T", val)
			}
			config.DynamicConfig = *typed
			return nil
		},
	}

	// Process each result with its corresponding processor
	for i, result := range results {
		val, err := result.GetResult()
		if err != nil {
			return cciptypes.OfframpConfig{}, fmt.Errorf("get offramp result %d: %w", i, err)
		}

		if err := processors[i](val); err != nil {
			return cciptypes.OfframpConfig{}, fmt.Errorf("process result %d: %w", i, err)
		}
	}

	return config, nil
}

func processOnRampResults(results []types.BatchReadResult) (cciptypes.OnRampConfig, error) {
	if len(results) != 2 {
		return cciptypes.OnRampConfig{}, fmt.Errorf("expected 2 OnRamp results, got %d", len(results))
	}

	var config cciptypes.OnRampConfig

	// Process DynamicConfig
	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.OnRampConfig{}, fmt.Errorf("get OnRamp dynamic config result: %w", err)
	}

	dynamicConfig, ok := val.(*cciptypes.GetOnRampDynamicConfigResponse)
	if !ok {
		return cciptypes.OnRampConfig{}, fmt.Errorf("invalid type for OnRamp dynamic config: %T", val)
	}
	config.DynamicConfig = *dynamicConfig

	// Process DestChainConfig
	val, err = results[1].GetResult()
	if err != nil {
		return cciptypes.OnRampConfig{}, fmt.Errorf("get OnRamp dest chain config result: %w", err)
	}

	destConfig, ok := val.(*cciptypes.OnRampDestChainConfig)
	if !ok {
		return cciptypes.OnRampConfig{}, fmt.Errorf("invalid type for OnRamp dest chain config: %T", val)
	}
	config.DestChainConfig = *destConfig

	return config, nil
}

func processRMNProxyResults(results []types.BatchReadResult) (cciptypes.RMNProxyConfig, error) {
	if len(results) != 1 {
		return cciptypes.RMNProxyConfig{}, fmt.Errorf("expected 1 RMN proxy result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.RMNProxyConfig{}, fmt.Errorf("get RMN proxy result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return cciptypes.RMNProxyConfig{
			RemoteAddress: *bytes,
		}, nil
	}

	return cciptypes.RMNProxyConfig{}, fmt.Errorf("invalid type for RMN proxy remote address: %T", val)
}

func processRMNRemoteResults(
	results []types.BatchReadResult,
	destChainSelector cciptypes.ChainSelector,
) (
	cciptypes.RMNRemoteConfig,
	cciptypes.CurseInfo,
	error,
) {
	config := cciptypes.RMNRemoteConfig{}

	if len(results) != 3 {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("expected 3 RMN remote results, got %d", len(results))
	}

	// Process DigestHeader
	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("get RMN remote digest header result: %w", err)
	}

	typed, ok := val.(*cciptypes.RMNDigestHeader)
	if !ok {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("invalid type for RMN remote digest header: %T", val)
	}
	config.DigestHeader = *typed

	// Process VersionedConfig
	val, err = results[1].GetResult()
	if err != nil {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("get RMN remote versioned config result: %w", err)
	}

	vconf, ok := val.(*cciptypes.VersionedConfig)
	if !ok {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("invalid type for RMN remote versioned config: %T", val)
	}
	config.VersionedConfig = *vconf

	// Process CursedSubjects
	val, err = results[2].GetResult()
	if err != nil {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("get RMN remote cursed subjects result: %w", err)
	}

	c, ok := val.(*cciptypes.RMNCurseResponse)
	if !ok {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("invalid type for RMN remote cursed subjects: %T", val)
	}
	curseInfo := *getCurseInfoFromCursedSubjects(
		mapset.NewSet(c.CursedSubjects...),
		destChainSelector,
	)

	return config, curseInfo, nil
}

func processFeeQuoterResults(results []types.BatchReadResult) (cciptypes.FeeQuoterConfig, error) {
	if len(results) != 1 {
		return cciptypes.FeeQuoterConfig{}, fmt.Errorf("expected 1 fee quoter result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.FeeQuoterConfig{}, fmt.Errorf("get fee quoter result: %w", err)
	}

	if typed, ok := val.(*cciptypes.FeeQuoterStaticConfig); ok {
		return cciptypes.FeeQuoterConfig{
			StaticConfig: *typed,
		}, nil
	}

	return cciptypes.FeeQuoterConfig{}, fmt.Errorf("invalid type for fee quoter static config: %T", val)
}

func processRouterResults(results []types.BatchReadResult) (cciptypes.RouterConfig, error) {
	if len(results) != 1 {
		return cciptypes.RouterConfig{}, fmt.Errorf("expected 1 router result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.RouterConfig{}, fmt.Errorf("get router wrapped native result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return cciptypes.RouterConfig{
			WrappedNativeAddress: cciptypes.Bytes(*bytes),
		}, nil
	}

	return cciptypes.RouterConfig{}, fmt.Errorf("invalid type for router wrapped native address: %T", val)
}

func getCurseInfoFromCursedSubjects(
	cursedSubjectsSet mapset.Set[[16]byte],
	destChainSelector cciptypes.ChainSelector,
) *cciptypes.CurseInfo {
	curseInfo := &cciptypes.CurseInfo{
		CursedSourceChains: make(map[cciptypes.ChainSelector]bool, cursedSubjectsSet.Cardinality()),
		CursedDestination: cursedSubjectsSet.Contains(cciptypes.GlobalCurseSubject) ||
			cursedSubjectsSet.Contains(chainSelectorToBytes16(destChainSelector)),
		GlobalCurse: cursedSubjectsSet.Contains(cciptypes.GlobalCurseSubject),
	}

	for _, cursedSubject := range cursedSubjectsSet.ToSlice() {
		if cursedSubject == cciptypes.GlobalCurseSubject {
			continue
		}

		chainSelector := cciptypes.ChainSelector(binary.BigEndian.Uint64(cursedSubject[8:]))
		if chainSelector == destChainSelector {
			continue
		}

		curseInfo.CursedSourceChains[chainSelector] = true
	}
	return curseInfo
}

func chainSelectorToBytes16(chainSel cciptypes.ChainSelector) [16]byte {
	var result [16]byte
	// Convert the uint64 to bytes and place it in the last 8 bytes of the array
	binary.BigEndian.PutUint64(result[8:], uint64(chainSel))
	return result
}

// resultProcessor defines a function type for processing individual results
type resultProcessor func(interface{}) error
