package reader

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

func TestCCIPChainReader_getSourceChainsConfig(t *testing.T) {
	sourceCRs := make(map[cciptypes.ChainSelector]*mocks.ContractReaderMock)
	for _, chain := range []cciptypes.ChainSelector{chainA, chainB} {
		sourceCRs[chain] = mocks.NewContractReaderMock()
	}

	destCR := mocks.NewContractReaderMock()

	destCR.On(
		"GetLatestValue",
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		mock.Anything,
		mock.Anything,
	).Run(func(args mock.Arguments) {
		sourceChain := args.Get(3).(map[string]any)["sourceChainSelector"].(cciptypes.ChainSelector)
		v := args.Get(4).(*sourceChainConfig)
		v.OnRamp = []byte(fmt.Sprintf("onramp-%d", sourceChain))
	}).Return(nil)

	ccipReader := NewCCIPChainReader(
		logger.Test(t),
		map[cciptypes.ChainSelector]types.ContractReader{
			chainC: destCR,
		}, nil, chainC,
	)

	ctx := context.Background()
	cfgs, err := ccipReader.getSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB})
	assert.NoError(t, err)
	assert.Len(t, cfgs, 2)
	assert.Equal(t, "onramp-1", string(cfgs[chainA].OnRamp))
	assert.Equal(t, "onramp-2", string(cfgs[chainB].OnRamp))
}
