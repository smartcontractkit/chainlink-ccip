package reader

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/crconsts"
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
		crconsts.ContractNameOffRamp,
		crconsts.FunctionNameGetSourceChainConfig,
		mock.Anything,
		mock.Anything,
	).Run(func(args mock.Arguments) {
		sourceChain := args.Get(3).(map[string]any)["sourceChainSelector"].(cciptypes.ChainSelector)
		v := args.Get(4).(*sourceChainConfig)
		v.OnRamp = fmt.Sprintf("onramp-%d", sourceChain)
	}).Return(nil)

	ccipReader := NewCCIPChainReader(
		map[cciptypes.ChainSelector]types.ContractReader{
			chainA: sourceCRs[chainA],
			chainB: sourceCRs[chainB],
			chainC: destCR,
		}, nil, chainC,
	)

	ctx := context.Background()
	cfgs, err := ccipReader.getSourceChainsConfig(ctx)
	assert.NoError(t, err)
	assert.Len(t, cfgs, 3)
	assert.Equal(t, "onramp-1", cfgs[chainA].OnRamp)
	assert.Equal(t, "onramp-2", cfgs[chainB].OnRamp)
	assert.Equal(t, "onramp-3", cfgs[chainC].OnRamp)
}
