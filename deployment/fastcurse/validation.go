package fastcurse

import (
	"fmt"
	"strings"
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

// validateVersions validates that curse actions specify a version.
// A version is required to select the correct adapter implementation.
func validateVersions(cfg RMNCurseConfig) error {
	for _, action := range cfg.CurseActions {
		if action.Version == nil {
			if action.IsGlobalCurse {
				return fmt.Errorf("global curse action missing version: chain %d", action.ChainSelector)
			}
			return fmt.Errorf("lane curse action missing version: chain %d -> %d", action.ChainSelector, action.SubjectChainSelector)
		}
	}
	return nil
}

// validateBidirectionalLaneActions validates that lane curses/uncurses are only applied bidirectionally.
//
// Business rule:
// - non-global lane actions must include the reverse direction (any version is acceptable)
// - global curses are excluded
// - actions with an explicit Subject are excluded, since they aren't tied to a real
//   lane between two chain selectors
func validateBidirectionalLaneActions(cfg RMNCurseConfig) error {
	allLaneActions := make(map[laneKey]curseActionInfo)

	for _, action := range cfg.CurseActions {
		if action.IsGlobalCurse || action.Subject != nil {
			continue
		}
		if action.Version == nil {
			return fmt.Errorf(
				"lane curse action missing version: chain %d -> %d",
				action.ChainSelector,
				action.SubjectChainSelector,
			)
		}

		allLaneActions[laneKey{source: action.ChainSelector, target: action.SubjectChainSelector}] = curseActionInfo{
			version:         action.Version.String(),
			chainSelector:   action.ChainSelector,
			subjectSelector: action.SubjectChainSelector,
		}
	}

	var unidirectional []curseActionInfo
	for key, actionInfo := range allLaneActions {
		reverseKey := laneKey{source: key.target, target: key.source}
		if _, ok := allLaneActions[reverseKey]; ok {
			continue
		}
		unidirectional = append(unidirectional, actionInfo)
	}

	if len(unidirectional) > 0 {
		return formatUnidirectionalLaneError(unidirectional)
	}
	return nil
}

func formatUnidirectionalLaneError(unidirectional []curseActionInfo) error {
	if len(unidirectional) == 1 {
		lane := unidirectional[0]
		return fmt.Errorf(
			"unidirectional lane cursing is not allowed: chain %d -> %d (version %s). lanes must be cursed bidirectionally to prevent requests from getting stuck indefinitely",
			lane.chainSelector,
			lane.subjectSelector,
			lane.version,
		)
	}

	var b strings.Builder
	fmt.Fprintf(
		&b,
		"unidirectional lane cursing is not allowed for %d lanes. lanes must be cursed bidirectionally to prevent requests from getting stuck indefinitely:\n",
		len(unidirectional),
	)
	for i, lane := range unidirectional {
		fmt.Fprintf(&b, "  %d. Chain %d -> %d (version %s)\n", i+1, lane.chainSelector, lane.subjectSelector, lane.version)
	}
	return fmt.Errorf("%s", b.String())
}
