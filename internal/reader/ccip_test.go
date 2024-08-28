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

	chainreadermocks "github.com/smartcontractkit/chainlink-ccip/mocks/cl-common/chainreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

func TestCCIPChainReader_getSourceChainsConfig(t *testing.T) {
	sourceCRs := make(map[cciptypes.ChainSelector]*chainreadermocks.MockChainReader)
	for _, chain := range []cciptypes.ChainSelector{chainA, chainB} {
		sourceCRs[chain] = chainreadermocks.NewMockChainReader(t)
	}

	destCR := chainreadermocks.NewMockChainReader(t)

	destCR.On(
		"GetLatestValue",
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(func(args mock.Arguments) {
		sourceChain := args.Get(4).(map[string]any)["sourceChainSelector"].(cciptypes.ChainSelector)
		v := args.Get(5).(*sourceChainConfig)
		v.OnRamp = []byte(fmt.Sprintf("onramp-%d", sourceChain))
	}).Return(nil)

	ccipReader := NewCCIPChainReader(
		logger.Test(t),
		map[cciptypes.ChainSelector]types.ContractReader{
			chainA: sourceCRs[chainA],
			chainB: sourceCRs[chainB],
			chainC: destCR,
		}, nil, chainC,
	)

	ctx := context.Background()
	cfgs, err := ccipReader.getSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB})
	assert.NoError(t, err)
	assert.Len(t, cfgs, 2)
	assert.Equal(t, []byte("onramp-1"), cfgs[chainA].OnRamp)
	assert.Equal(t, []byte("onramp-2"), cfgs[chainB].OnRamp)
}
