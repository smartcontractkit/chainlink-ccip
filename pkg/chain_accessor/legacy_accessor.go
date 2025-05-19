package chain_accessor

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type LegacyAccessor struct {
	chainSelector  cciptypes.ChainSelector
	contractReader contractreader.Extended
	contractWriter types.ContractWriter
	addrCodec      cciptypes.AddressCodec
}

func NewLegacyAccessor(chainSelector cciptypes.ChainSelector, contractReader contractreader.Extended, contractWriter types.ContractWriter, addrCodec cciptypes.AddressCodec) *LegacyAccessor {
	// for now, all chains use the same cr/cw based legacy accessor
	return &LegacyAccessor{
		chainSelector:  chainSelector,
		contractReader: contractReader,
		contractWriter: contractWriter,
		addrCodec:      addrCodec,
	}
}

func (l LegacyAccessor) Metadata() cciptypes.AccessorMetadata {
	return cciptypes.AccessorMetadata{
		ChainSelector: l.chainSelector,
	}
}

func (l LegacyAccessor) GetContractAddress(contractName string) ([]byte, error) {
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

func (l LegacyAccessor) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetDestChainFeeComponents(ctx context.Context) (types.ChainFeeComponents, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) Sync(ctx context.Context, contracts cciptypes.ContractAddresses) error {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) MsgsBetweenSeqNums(ctx context.Context, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange) ([]cciptypes.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) LatestMsgSeqNum(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetExpectedNextSequenceNumber(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetTokenPriceUSD(ctx context.Context, address cciptypes.UnknownAddress) (cciptypes.TimestampedUnixBig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetFeeQuoterDestChainConfig(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.FeeQuoterDestChainConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) CommitReportsGTETimestamp(ctx context.Context, ts time.Time, limit int) ([]cciptypes.CommitPluginReportWithMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) ExecutedMessages(ctx context.Context, ranges map[cciptypes.ChainSelector][]cciptypes.SeqNumRange, confidence cciptypes.ConfidenceLevel) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) NextSeqNum(ctx context.Context, sources []cciptypes.ChainSelector) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) Nonces(ctx context.Context, addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.TimestampedBig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetOffRampSourceChainsConfig(ctx context.Context, sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	//TODO implement me
	panic("implement me")
}
