package multinode

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/config"
)

func NewMultiNode(ctx context.Context, lggr logger.Logger, chainConfig *config.ChainScoped) (client.Client, error) {
	multiNode, err := client.NewEvmClient(chainConfig.EVM().NodePool(), chainConfig.EVM(), chainConfig.EVM().NodePool().Errors(), lggr, chainConfig.EVM().ChainID(), chainConfig.Nodes(), chainConfig.EVM().ChainType())
	if err != nil {
		lggr.Errorw("Error creating MultiNode", "err", err)
		return nil, err
	}

	err = multiNode.Dial(ctx)
	if err != nil {
		lggr.Errorw("Error dialing MultiNode", "err", err)
		return nil, err
	}

	return multiNode, nil
}
