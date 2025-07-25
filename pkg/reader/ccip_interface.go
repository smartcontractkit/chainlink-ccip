package reader

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/chainaccessor"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

var (
	ErrChainAccessorNotFound  = errors.New("chain accessor not found")
	ErrContractReaderNotFound = errors.New("contract reader not found")
	ErrContractWriterNotFound = errors.New("contract writer not found")
)

// ContractAddresses is a map of contract names across all chain selectors and their address.
// Currently only one contract per chain per name is supported.
type ContractAddresses map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress

// ChainConfigSnapshot represents the complete configuration state of the chain
type ChainConfigSnapshot struct {
	Offramp   OfframpConfig
	RMNProxy  RMNProxyConfig
	RMNRemote RMNRemoteConfig
	FeeQuoter FeeQuoterConfig
	OnRamp    OnRampConfig
	Router    RouterConfig
	CurseInfo CurseInfo
}

type OnRampConfig struct {
	DynamicConfig   getOnRampDynamicConfigResponse
	DestChainConfig onRampDestChainConfig
}

type FeeQuoterConfig struct {
	StaticConfig feeQuoterStaticConfig
}

type RMNRemoteConfig struct {
	DigestHeader    rmnDigestHeader
	VersionedConfig versionedConfig
}

type OfframpConfig struct {
	CommitLatestOCRConfig OCRConfigResponse
	ExecLatestOCRConfig   OCRConfigResponse
	StaticConfig          offRampStaticChainConfig
	DynamicConfig         offRampDynamicChainConfig
}

type RMNProxyConfig struct {
	RemoteAddress []byte
}

type RouterConfig struct {
	WrappedNativeAddress cciptypes.Bytes
}

func (ca ContractAddresses) Append(contract string, chain cciptypes.ChainSelector, address []byte) ContractAddresses {
	resp := ca
	if resp == nil {
		resp = make(ContractAddresses)
	}
	if resp[contract] == nil {
		resp[contract] = make(map[cciptypes.ChainSelector]cciptypes.UnknownAddress)
	}
	resp[contract][chain] = address
	return resp
}

// StaticSourceChainConfig stores the static parts of SourceChainConfig
// that don't change frequently and are safe to cache.
type StaticSourceChainConfig struct {
	Router                    []byte
	IsEnabled                 bool
	IsRMNVerificationDisabled bool
	OnRamp                    cciptypes.UnknownAddress
}

// ToSourceChainConfig converts a CachedSourceChainConfig to a full SourceChainConfig
// by adding the provided sequence number.
func (s StaticSourceChainConfig) ToSourceChainConfig(minSeqNr uint64) cciptypes.SourceChainConfig {
	return cciptypes.SourceChainConfig{
		Router:                    s.Router,
		IsEnabled:                 s.IsEnabled,
		IsRMNVerificationDisabled: s.IsRMNVerificationDisabled,
		OnRamp:                    s.OnRamp,
		MinSeqNr:                  minSeqNr,
	}
}

func (s StaticSourceChainConfig) check() (bool /* enabled */, error) {
	// The chain may be set in CCIPHome's ChainConfig map but not hooked up yet in the offramp.
	if !s.IsEnabled {
		return false, nil
	}
	// This may happen due to some sort of regression in the codec that unmarshals
	// chain data -> go struct.
	if len(s.OnRamp) == 0 {
		return false, fmt.Errorf(
			"onRamp misconfigured/didn't unmarshal: %x",
			s.OnRamp,
		)
	}

	if len(s.Router) == 0 {
		return false, fmt.Errorf("router is empty: %v", s.Router)
	}

	return s.IsEnabled, nil
}

func NewCCIPChainReader(
	ctx context.Context,
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
	addrCodec cciptypes.AddressCodec,
) (CCIPReader, error) {
	reader, err := newCCIPChainReaderInternal(
		ctx,
		lggr,
		contractReaders,
		contractWriters,
		destChain,
		offrampAddress,
		addrCodec,
	)
	if err != nil {
		return nil, err
	}
	return NewObservedCCIPReader(
		reader,
		lggr,
		destChain,
	), nil
}

// NewCCIPReaderWithExtendedContractReaders can be used when you want to directly provide contractreader.Extended
func NewCCIPReaderWithExtendedContractReaders(
	ctx context.Context,
	lggr logger.Logger,
	extendedContractReaders map[cciptypes.ChainSelector]contractreader.Extended,
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
	addrCodec cciptypes.AddressCodec,
) CCIPReader {
	cr, err := newCCIPChainReaderInternal(ctx, lggr, nil, contractWriters, destChain, offrampAddress, addrCodec)
	if err != nil {
		// Panic here since right now this is only called from tests in core
		panic(fmt.Errorf("failed to create CCIP reader: %w", err))
	}
	var cas = make(map[cciptypes.ChainSelector]cciptypes.ChainAccessor)
	for ch, extendedCr := range extendedContractReaders {
		cr.WithExtendedContractReader(ch, extendedCr)
		accessor, err := chainaccessor.NewDefaultAccessor(
			lggr,
			ch,
			extendedCr,
			contractWriters[ch],
			addrCodec,
		)
		if err != nil {
			// Panic here since this is only called from tests in core
			panic(fmt.Errorf("failed to create chain accessor for %s: %w", ch, err))
		}
		cas[ch] = accessor
	}

	cr.accessors = cas
	return cr
}

type CCIPReader interface {
	// CommitReportsGTETimestamp reads the destination chain starting at a given timestamp
	// and finds all ReportAccepted up to the provided limit.
	CommitReportsGTETimestamp(ctx context.Context,
		ts time.Time,
		confidence primitives.ConfidenceLevel,
		limit int,
	) ([]cciptypes.CommitPluginReportWithMeta, error)

	// ExecutedMessages finds executed messages for all source chains/ranges provided on a single destination chain.
	// A map of source chain to slice of sequence numbers is returned to express which seqnrs have executed.
	ExecutedMessages(
		ctx context.Context,
		rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
		confidence primitives.ConfidenceLevel,
	) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error)

	// MsgsBetweenSeqNums reads the provided chains, finds and returns ccip messages
	// submitted between the provided sequence numbers. Messages are sorted ascending based on
	// their timestamp and limited up to the provided limit.
	MsgsBetweenSeqNums(
		ctx context.Context,
		chain cciptypes.ChainSelector,
		seqNumRange cciptypes.SeqNumRange,
	) ([]cciptypes.Message, error)

	// LatestMsgSeqNum reads the source chain and returns the latest finalized message sequence number.
	LatestMsgSeqNum(ctx context.Context, chain cciptypes.ChainSelector) (cciptypes.SeqNum, error)

	// GetExpectedNextSequenceNumber reads the source chain and returns the expected next onRamp sequence number.
	// DEPRECATED: Not used and should be removed.
	GetExpectedNextSequenceNumber(
		ctx context.Context,
		sourceChainSelector cciptypes.ChainSelector,
	) (cciptypes.SeqNum, error)

	// NextSeqNum reads the destination chain.
	// Returns the next expected sequence number for each one of the provided chains.
	NextSeqNum(ctx context.Context, chains []cciptypes.ChainSelector) (
		seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error)

	// GetContractAddress returns the contract address that is registered for the provided contract name and chain.
	GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error)

	// Nonces fetches all nonces for the provided selector/address pairs. Addresses are a string encoded raw address,
	// it must be encoding according to the source chain requirements with typeconv.AddressBytesToString.
	Nonces(
		ctx context.Context,
		addressesByChain map[cciptypes.ChainSelector][]string,
	) (map[cciptypes.ChainSelector]map[string]uint64, error)

	// GetChainsFeeComponents Returns all fee components for given chains if corresponding
	// chain writer is available. It will log errors and return a subset of the requested chains.
	GetChainsFeeComponents(
		ctx context.Context,
		chains []cciptypes.ChainSelector,
	) map[cciptypes.ChainSelector]types.ChainFeeComponents

	// GetDestChainFeeComponents Reads the fee components for the destination chain.
	// DEPRECATED: Not used and should be removed.
	GetDestChainFeeComponents(ctx context.Context) (types.ChainFeeComponents, error)

	// GetWrappedNativeTokenPriceUSD Gets the wrapped native token price in USD for the provided chains.
	// It reads the prices from the FeeQuoter contract on each chain.
	// If it encounters an issue (e.g. chain not supported) it logs the error and skips the chain.
	GetWrappedNativeTokenPriceUSD(
		ctx context.Context,
		selectors []cciptypes.ChainSelector,
	) map[cciptypes.ChainSelector]cciptypes.BigInt

	// GetChainFeePriceUpdate Gets latest chain fee price update for the provided chains from the destination.
	// If it encounters an issue (e.g. dest chain not supported) it logs the error and returns an empty result.
	GetChainFeePriceUpdate(
		ctx context.Context,
		selectors []cciptypes.ChainSelector,
	) map[cciptypes.ChainSelector]cciptypes.TimestampedBig

	// GetRMNRemoteConfig reads the RMN remote contract on the destination chain and returns the remote config.
	GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error)

	// GetRmnCurseInfo returns rmn curse/pausing information about the provided chains
	// from the destination chain RMN remote contract. Caller should be able to access destination.
	GetRmnCurseInfo(ctx context.Context) (CurseInfo, error)

	// DiscoverContracts will discover as many addresses as possible based on which addresses are already known.
	// Initially only the offramp address is known so in the first round the oracles that support the destination chain
	// are reading the offramp to discover onRamps, dest routers, rmn remote, nonce manager and fee quoter addresses.
	// On the next round the oracles that support the source chains should be able to access the onRamps since their
	// addresses are known. That's how they eventually discover source fee quoters and source routers.
	// They are returned per contract and source chain selector.
	// allChains is needed because there is no way to enumerate all chain selectors on Solana. We'll attempt to
	// fetch the source config from the offramp for each of them.
	DiscoverContracts(ctx context.Context,
		supportedChains, allChains []cciptypes.ChainSelector) (ContractAddresses, error)

	// LinkPriceUSD gets the LINK price in 1e-18 USDs from the FeeQuoter contract on the destination chain.
	// For example, if the price is 1 LINK = 10 USD, this function will return 10e18 (10 * 1e18). You can think of this
	// function returning the price of LINK not in USD, but in a small denomination of USD, similar to returning
	// the price of ETH not in ETH but in wei (1e-18 ETH).
	// DEPRECATED: Not used and should be removed.
	LinkPriceUSD(ctx context.Context) (cciptypes.BigInt, error)

	// Sync can be used to perform frequent syncing operations inside the reader implementation.
	// NOTE: this method may make network calls.
	//
	// NOTE: You should ensure that Sync is called deterministically for every oracle in the DON (e.g. in the Outcome).
	Sync(ctx context.Context, contracts ContractAddresses) error

	// GetLatestPriceSeqNr returns the latest price sequence number for the destination chain.
	// Not to confuse with the sequence number of the messages. This is the OCR sequence number.
	GetLatestPriceSeqNr(ctx context.Context) (uint64, error)

	// GetOffRampConfigDigest returns the offramp config digest for the provided plugin type.
	GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error)

	// GetOffRampSourceChainsConfig returns the source chain static configs for all the provided source chains.
	// This method returns StaticSourceChainConfig objects which deliberately exclude MinSeqNr.
	// If a config was not found it will be missing from the returned map.
	GetOffRampSourceChainsConfig(ctx context.Context, sourceChains []cciptypes.ChainSelector,
	) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error)

	Close() error
}
