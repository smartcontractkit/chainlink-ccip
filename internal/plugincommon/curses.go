package plugincommon

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// IsReportCursed fetches RMN CursesInfo from the destination and compares it
// to the reportSourceChains. If any source (or destination) are cursed, a nice
// warning is logged and true is returned.
//
// Used by commit and execute plugins to ensure consistent logs and errors.
func IsReportCursed(
	ctx context.Context,
	lggr logger.Logger,
	ccipReader reader.CCIPReader,
	reportSourceChains []cciptypes.ChainSelector,
) (bool, error) {
	// No error is returned in case there are other things for the report.
	if len(reportSourceChains) == 0 {
		return false, nil
	}

	curseInfo, err := ccipReader.GetRmnCurseInfo(ctx)
	if err != nil {
		return false, fmt.Errorf("error while fetching curse info: %w", err)
	}
	if curseInfo.GlobalCurse || curseInfo.CursedDestination {
		lggr.Warnw(
			"report not accepted due to RMN curse",
			"curseInfo", curseInfo,
		)
		return true, nil
	}
	filtered := curseInfo.NonCursedSourceChains(reportSourceChains)
	if len(filtered) != len(reportSourceChains) {
		var cursedSources []cciptypes.ChainSelector
		for cs, isCursed := range curseInfo.CursedSourceChains {
			if isCursed {
				cursedSources = append(cursedSources, cs)
			}
		}

		lggr.Warnw("report not accepted due to cursing, source chains were cursed during report generation",
			"cursedSourceChains", cursedSources,
		)
		return true, nil
	}
	return false, nil
}
