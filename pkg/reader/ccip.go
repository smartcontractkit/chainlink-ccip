package reader

import (
	readerinternal "github.com/smartcontractkit/chainlink-ccip/internal/reader"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type CCIPReader = readerinternal.CCIP

func NewCCIPReader(
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
) CCIPReader {
	return readerinternal.NewCCIPChainReader(lggr, contractReaders, contractWriters, destChain)
}
