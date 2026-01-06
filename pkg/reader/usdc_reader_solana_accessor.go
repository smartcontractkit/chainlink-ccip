package reader

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

// solanaUSDCMessageReaderAccessor is a USDCReader for Solana that delegates calls to the Solana ChainAccessor
// to fetch onchain information.
type solanaUSDCMessageReaderAccessor struct {
	lggr          logger.Logger
	chainAccessor cciptypes.ChainAccessor
}

func NewSolanaUSDCReaderAccessor(
	ctx context.Context,
	lggr logger.Logger,
	chainAccessors map[cciptypes.ChainSelector]cciptypes.ChainAccessor,
	addrCodec cciptypes.AddressCodec,
	chainSelector cciptypes.ChainSelector,
	bytesAddress cciptypes.UnknownAddress,
) (USDCMessageReader, error) {
	accessor, err := getChainAccessor(chainAccessors, chainSelector)
	if err != nil {
		return nil, fmt.Errorf("get chain accessor for Solana USDCReader: %w", err)
	}

	addressStr, err := addrCodec.AddressBytesToString(bytesAddress, chainSelector)
	if err != nil {
		return nil, fmt.Errorf("unable to convert address bytes to string: %w, address: %v", err, bytesAddress)
	}

	lggr.Debugw("Syncing contract to accessor",
		"chainSelector", chainSelector,
		"contractName", consts.ContractNameUSDCTokenPool,
		"address", addressStr,
	)

	err = accessor.Sync(ctx, consts.ContractNameUSDCTokenPool, bytesAddress)
	if err != nil {
		return nil, fmt.Errorf("sync contract to Solana USDC accessor: %w", err)
	}

	return solanaUSDCMessageReaderAccessor{
		lggr:          lggr,
		chainAccessor: accessor,
	}, nil
}

// MessagesByTokenID retrieves the CCTP MessageSent events.
func (u solanaUSDCMessageReaderAccessor) MessagesByTokenID(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[cciptypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[cciptypes.MessageTokenID]cciptypes.Bytes, error) {
	if len(tokens) == 0 {
		return map[cciptypes.MessageTokenID]cciptypes.Bytes{}, nil
	}

	u.lggr.Debugw("Searching for Solana CCTP USDC logs", "numExpected", len(tokens))
	messagesByID, err := u.chainAccessor.MessagesByTokenID(ctx, source, dest, tokens)
	if err != nil {
		return nil, fmt.Errorf("error fetching messages by token ID from chain accessor for chain %d: %w", source, err)
	}

	return messagesByID, nil
}
