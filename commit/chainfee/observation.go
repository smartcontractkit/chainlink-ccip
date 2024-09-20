package chainfee

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func (p *processor) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

func (p *processor) ObserveFeeComponents(ctx context.Context) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	return p.ccipReader.GetAllChainsFeeComponents(ctx)
}
