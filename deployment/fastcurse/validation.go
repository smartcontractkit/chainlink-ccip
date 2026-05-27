package fastcurse

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
)

type laneKey struct {
	source uint64
	target uint64
}

type curseActionInfo struct {
	version         string
	chainSelector   uint64
	subjectSelector uint64
}

// validateBidirectionalV16Cursing validates that v1.6 lanes are only cursed/uncursed bidirectionally.
//
// Business rule:
// - v1.6 lane actions (non-global) must include the reverse direction (any version is acceptable)
// - global curses are excluded
func validateBidirectionalV16Cursing(cfg RMNCurseConfig) error {
	allLaneActions := make(map[laneKey]curseActionInfo)
	v16LaneActions := make(map[laneKey]curseActionInfo)

	for _, action := range cfg.CurseActions {
		if action.IsGlobalCurse {
			continue
		}
		if action.Version == nil {
			continue
		}
		if action.ChainSelector == action.SubjectChainSelector {
			continue
		}

		key := laneKey{source: action.ChainSelector, target: action.SubjectChainSelector}
		actionInfo := curseActionInfo{
			version:         action.Version.String(),
			chainSelector:   action.ChainSelector,
			subjectSelector: action.SubjectChainSelector,
		}
		allLaneActions[key] = actionInfo

		if isV16(action.Version) {
			v16LaneActions[key] = actionInfo
		}
	}

	var unidirectional []curseActionInfo
	for key, actionInfo := range v16LaneActions {
		reverseKey := laneKey{source: key.target, target: key.source}
		if _, ok := allLaneActions[reverseKey]; ok {
			continue
		}
		unidirectional = append(unidirectional, actionInfo)
	}

	if len(unidirectional) > 0 {
		return formatUnidirectionalV16Error(unidirectional)
	}
	return nil
}

func isV16(v *semver.Version) bool {
	return v != nil && v.Major() == 1 && v.Minor() == 6
}

func formatUnidirectionalV16Error(unidirectional []curseActionInfo) error {
	if len(unidirectional) == 1 {
		lane := unidirectional[0]
		return fmt.Errorf(
			"unidirectional v1.6 lane cursing is not allowed: chain %d -> %d (version %s). v1.6 lanes must be cursed bidirectionally to prevent requests from getting stuck indefinitely",
			lane.chainSelector,
			lane.subjectSelector,
			lane.version,
		)
	}

	var b strings.Builder
	fmt.Fprintf(
		&b,
		"unidirectional v1.6 lane cursing is not allowed for %d lanes. v1.6 lanes must be cursed bidirectionally to prevent requests from getting stuck indefinitely:\n",
		len(unidirectional),
	)
	for i, lane := range unidirectional {
		fmt.Fprintf(&b, "  %d. Chain %d -> %d (version %s)\n", i+1, lane.chainSelector, lane.subjectSelector, lane.version)
	}
	return fmt.Errorf("%s", b.String())
}
