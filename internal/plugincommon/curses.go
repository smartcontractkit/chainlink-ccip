package plugincommon

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// IsReportCursed helps make sure logs and errors are consistent in commit and execute plugins.
func IsReportCursed(
	ctx context.Context,
	lggr logger.Logger,
	ccipReader reader.CCIPReader,
	destChain ccipocr3.ChainSelector,
	touchedChains []ccipocr3.ChainSelector,
) (bool, error) {
	// No error is returned in case there are other things for the report.
	if len(touchedChains) == 0 {
		return false, nil
	}

	curseInfo, err := ccipReader.GetRmnCurseInfo(ctx, destChain, touchedChains)
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
	filtered := curseInfo.NonCursedSourceChains(touchedChains)
	if len(filtered) != len(touchedChains) {
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
