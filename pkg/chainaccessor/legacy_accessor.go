package chainaccessor

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// LegacyAccessor is an implementation of cciptypes.ChainAccessor that allows the CCIPReader
// to cutover and migrate away from depending directly on contract reader and contract writer.
type LegacyAccessor struct {
	lggr           logger.Logger
	chainSelector  cciptypes.ChainSelector
	contractReader contractreader.Extended
	contractWriter types.ContractWriter
	addrCodec      cciptypes.AddressCodec
}

var _ cciptypes.ChainAccessor = (*LegacyAccessor)(nil)

func NewLegacyAccessor(
	lggr logger.Logger,
	chainSelector cciptypes.ChainSelector,
	contractReader contractreader.Extended,
	contractWriter types.ContractWriter,
	addrCodec cciptypes.AddressCodec,
) cciptypes.ChainAccessor {
	return &LegacyAccessor{
		lggr:           lggr,
		chainSelector:  chainSelector,
		contractReader: contractReader,
		contractWriter: contractWriter,
		addrCodec:      addrCodec,
	}
}

func (l *LegacyAccessor) Metadata() cciptypes.AccessorMetadata {
	allBindings := l.contractReader.GetAllBindings()
	contracts := make(map[string]cciptypes.UnknownAddress, len(allBindings))
	for contractName, binding := range allBindings {
		addressBytes, err := l.addrCodec.AddressStringToBytes(binding[0].Binding.Address, l.chainSelector)
		if err != nil {
			l.lggr.Errorf("failed to convert address to bytes : %v", err)
			continue
		}
		contracts[contractName] = addressBytes
	}

	return cciptypes.AccessorMetadata{
		ChainSelector: l.chainSelector,
		Contracts:     contracts,
	}
}

func (l *LegacyAccessor) GetContractAddress(contractName string) ([]byte, error) {
	bindings := l.contractReader.GetBindings(contractName)
	if len(bindings) != 1 {
		return nil, fmt.Errorf("expected one binding for the %s contract, got %d", contractName, len(bindings))
	}

	addressBytes, err := l.addrCodec.AddressStringToBytes(bindings[0].Binding.Address, l.chainSelector)
	if err != nil {
		return nil, fmt.Errorf("convert address %s to bytes: %w", bindings[0].Binding.Address, err)
	}

	return addressBytes, nil
}

func (l *LegacyAccessor) GetChainFeeComponents(
	ctx context.Context,
) map[cciptypes.ChainSelector]cciptypes.ChainFeeComponents {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetDestChainFeeComponents(ctx context.Context) (types.ChainFeeComponents, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) Sync(ctx context.Context, contracts cciptypes.ContractAddresses) error {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) MsgsBetweenSeqNums(
	ctx context.Context,
	dest cciptypes.ChainSelector,
	seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) LatestMsgSeqNum(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetExpectedNextSequenceNumber(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetTokenPriceUSD(
	ctx context.Context,
	address cciptypes.UnknownAddress,
) (cciptypes.BigInt, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetFeeQuoterDestChainConfig(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.FeeQuoterDestChainConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) ExecutedMessages(
	ctx context.Context,
	ranges map[cciptypes.ChainSelector]cciptypes.SeqNumRange,
	confidence cciptypes.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) NextSeqNum(
	ctx context.Context,
	sources []cciptypes.ChainSelector,
) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) Nonces(
	ctx context.Context,
	addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetOffRampSourceChainsConfig(
	ctx context.Context,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}
