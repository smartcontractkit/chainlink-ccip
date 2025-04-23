package ccipocr3

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

// AccessorMetadata contains metadata about the chain accessor.
type AccessorMetadata struct {
	ChainSelector ChainSelector
	Contracts     map[string]UnknownAddress
}

// ChainAccessor for all direct chain access. The accessor is responsible for
// in addition to direct access to the chain, this interface also translates
// onchain representations of data to the plugin representation.
type ChainAccessor interface {
	AllAccessors
	SourceAccessor
	DestinationAccessor
	RMNAccessor
}

// AllAccessors contains functionality that is available to all types of accessors.
type AllAccessors interface {
	// Metadata associated with the chain accessor. Replaces GetContractAddress.
	Metadata() AccessorMetadata

	// TODO: Discovery/Binding functions here?

	// GetContractAddress returns the contract address that is registered for the provided contract name and chain.
	// WARNING: This function will fail if the oracle does not support the requested chain.
	//
	// Deprecated: use Metadata() instead.
	GetContractAddress(contractName string) ([]byte, error)

	// GetAllConfig looks up all configurations available to the accessor. This
	// function replaces prepareBatchConfigRequests.
	//
	// This includes the following contracts:
	// - Router
	// - OffRamp
	// - OnRamp
	// - FeeQuoter
	// - RMNProxy
	// - RMNRemote
	// - CurseInfo
	//
	// Access Type: Method(many, see code)
	// Contract: Many
	// Confidence: Unconfirmed
	/*
		// TODO: ChainConfig object is huge with many private components. Need to move them all here
		// to avoid circular dependencies.
		GetAllConfig(
			ctx context.Context,
		) (ChainConfig, error)
	*/

	// GetChainFeeComponents Returns all fee components for given chains if corresponding
	// chain writer is available.
	//
	// Access Type: ChainWriter
	// Contract: N/A
	// Confidence: N/A
	GetChainFeeComponents(
		ctx context.Context,
	) map[ChainSelector]ChainFeeComponents

	// GetDestChainFeeComponents seems redundant. If the error is needed lets
	// add it to GetChainFeeComponents.
	//
	// Deprecated: use GetChainFeeComponents instead.
	GetDestChainFeeComponents(
		ctx context.Context,
	) (types.ChainFeeComponents, error)

	// Sync can be used to perform frequent syncing operations inside the reader implementation.
	// Returns a bool indicating whether something was updated.
	Sync(ctx context.Context, contracts ContractAddresses) error
}

// DestinationAccessor contains all functions typically associated by the destination chain.
type DestinationAccessor interface {

	// CommitReportsGTETimestamp reads CommitReportAccepted events starting from a given timestamp.
	// The number of results are limited according to limit.
	//
	// Access Type: Event(CommitReportAccepted)
	// Contract: OffRamp
	// Confidence: Unconfirmed
	CommitReportsGTETimestamp(
		ctx context.Context,
		ts time.Time,
		limit int,
	) ([]CommitPluginReportWithMeta, error)

	// ExecutedMessages looks for ExecutionStateChanged events for each sequence
	// in a given range. The presence of these events indicates that an attempt to
	// execute the message has been made, which the system considers "executed".
	// A slice of all executed sequence numbers is returned.
	//
	// Access Type: Event(ExecutionStateChanged)
	// Contract: OffRamp
	// Confidence: Unconfirmed, Finalized
	ExecutedMessages(
		ctx context.Context,
		ranges map[ChainSelector]SeqNumRange,
		confidence ConfidenceLevel,
	) (map[ChainSelector][]SeqNum, error)

	// NextSeqNum reads the source chain config to get the next expected
	// sequence number for the given source chains.
	//
	// Access Type: Method(GetSourceChainConfig)
	// Contract: OffRamp
	// Confidence: Unconfirmed
	//
	// TODO: Have a single "GetSourceChainConfigs" method for the log poller to
	//       use for cached values, and the plugin to use for non-cached values.
	//       Rename to "GetSourceChainConfig".
	//      See Design Doc (NextSeqNum) for notes.
	NextSeqNum(ctx context.Context, sources []ChainSelector) (
		seqNum map[ChainSelector]SeqNum, err error)

	// Nonces for all provided selector/address pairs. Addresses must be encoded
	// according to the source chain requirements by using the AddressCodec.
	//
	// Access Type: Method(GetInboundNonce)
	// Contract: NonceManager
	// Confidence: Unconfirmed
	Nonces(
		ctx context.Context,
		addresses map[ChainSelector][]UnknownEncodedAddress,
	) (map[ChainSelector]map[string]uint64, error)

	// GetChainFeePriceUpdate Gets latest chain fee price update for the provided chains.
	//
	// Access Type: Method(GetChainFeePriceUpdate)
	// Contract: FeeQuoter
	// Confidence: Unconfirmed
	GetChainFeePriceUpdate(
		ctx context.Context,
		selectors []ChainSelector,
	) map[ChainSelector]TimestampedBig

	// GetLatestPriceSeqNr returns the latest price sequence number for the destination chain.
	// Not to confuse with the sequence number of the messages. This is the OCR sequence number.
	//
	// Access Type: Method(GetLatestPriceSequenceNumber)
	// Contract: OffRamp
	// Confidence: Unconfirmed
	GetLatestPriceSeqNr(ctx context.Context) (uint64, error)

	// GetOffRampConfigDigest returns the offramp config digest for the provided plugin type.
	//
	// TODO: this should go away, it's a call to the ConfigPoller.
	GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error)

	// GetOffRampSourceChainsConfig returns the sourceChains config for all the provided source chains.
	// If a config was not found it will be missing from the returned map.
	//
	// TODO: this should go away, it's a call to the ConfigPoller.
	GetOffRampSourceChainsConfig(
		ctx context.Context,
		sourceChains []ChainSelector,
	) (map[ChainSelector]SourceChainConfig, error)
}

type SourceAccessor interface {
	// MsgsBetweenSeqNums returns all messages being sent to the provided dest
	// chain, between the provided sequence numbers. Messages are sorted ascending
	// based on their timestamp.
	//
	// Access Type: Event(CCIPMessageSent)
	// Contract: OnRamp
	// Confidence: Finalized
	MsgsBetweenSeqNums(
		ctx context.Context,
		dest ChainSelector,
		seqNumRange SeqNumRange,
	) ([]Message, error)

	// LatestMsgSeqNum returns the sequence number associated with the most
	// recent message being sent to a given destination.
	//
	// Access Type: Event(CCIPMessageSent)
	// Contract: OnRamp
	// Confidence: Finalized
	//
	// TODO: Rename function in CAL to LatestMessageTo(dest) and return the entire message.
	//       See Design Doc (LatestMsgSeqNum) for notes.
	LatestMsgSeqNum(ctx context.Context, dest ChainSelector) (SeqNum, error)

	// GetExpectedNextSequenceNumber returns the expected next sequence number
	// messages being sent to the provided destination.
	//
	// Access Type: Method(GetExpectedNextSequenceNumber)
	// Contract: OnRamp
	// Confidence: Unconfirmed
	GetExpectedNextSequenceNumber(
		ctx context.Context,
		dest ChainSelector,
	) (SeqNum, error)

	// GetTokenPriceUSD looks up a token price in USD. The address value should
	// be retrieved from a configuration cache (i.e. ConfigPoller).
	//
	// Access Type: Method(GetTokenPrice)
	// Contract: FeeQuoter
	// Confidence: Unconfirmed
	//
	// Notes: This function is new and serves as a general price interface for
	//        both LinkPriceUSD and GetWrappedNativeTokenPriceUSD.
	//        See Design Doc (Combined Token Price Helper) for notes.
	GetTokenPriceUSD(
		ctx context.Context,
		address UnknownAddress,
	) (BigInt, error)

	// GetFeeQuoterDestChainConfig returns the fee quoter destination chain config.
	//
	// Access Type: Method(GetDestChainConfig)
	// Contract: FeeQuoter
	// Confidence: Unconfirmed
	//
	// Notes: This is a new general purpose function needed to implement
	//        GetMedianDataAvailabilityGasConfig.
	GetFeeQuoterDestChainConfig(
		ctx context.Context,
		dest ChainSelector,
	) (FeeQuoterDestChainConfig, error)
}

type RMNAccessor interface {
	// GetRMNRemoteConfig
	//
	// TODO: This function should be removed and replaced with direct access to
	//       the ConfigPoller. It also manages chain reader contract binding
	//       directly instead of using DiscoverContracts.
	GetRMNRemoteConfig(ctx context.Context) (RemoteConfig, error)

	// GetRmnCurseInfo returns rmn curse/pausing information about the provided chains
	// from the destination chain RMN remote contract. Caller should be able to access destination.
	GetRmnCurseInfo(ctx context.Context) (CurseInfo, error)
}

////////////////////////////////////////////////////////////////
// TODO: Find a better location for the types below this line //
//       For the purpose of designing these interfaces, the   //
//       location is not critical.                            //
////////////////////////////////////////////////////////////////

// Random types. These are defined here mainly to bring focus to types which should
// probably be removed or replaced.

// ConfidenceLevel represents how likely it is that the data could be impacted by a reorg.
type ConfidenceLevel = primitives.ConfidenceLevel

const (
	// Finalized data is not expected to change, even if there is a reorg.
	Finalized ConfidenceLevel = "finalized"

	// Unconfirmed data may be modified by a reorged.
	Unconfirmed ConfidenceLevel = "unconfirmed"
)

type ChainFeeComponents = types.ChainFeeComponents

//type ChainConfig = reader.ChainConfigSnapshot

type TimestampedBig struct {
	Timestamp time.Time `json:"timestamp"`
	Value     BigInt    `json:"value"`
}

// TimestampedUnixBig Maps to on-chain struct
// https://github.com/smartcontractkit/chainlink/blob/37f3132362ec90b0b1c12fb1b69b9c16c46b399d/contracts/src/v0.8/ccip/libraries/Internal.sol#L43-L47
//
//nolint:lll //url
type TimestampedUnixBig struct {
	Value     *big.Int `json:"value"`
	Timestamp uint32   `json:"timestamp"`
}

func NewTimestampedBig(value int64, timestamp time.Time) TimestampedBig {
	return TimestampedBig{
		Value:     BigInt{Int: big.NewInt(value)},
		Timestamp: timestamp,
	}
}

func TimeStampedBigFromUnix(input TimestampedUnixBig) TimestampedBig {
	return TimestampedBig{
		Value:     NewBigInt(input.Value),
		Timestamp: time.Unix(int64(input.Timestamp), 0),
	}
}

type CommitPluginReportWithMeta struct {
	Report    CommitPluginReport `json:"report"`
	Timestamp time.Time          `json:"timestamp"`
	BlockNum  uint64             `json:"blockNum"`
}

type CommitReportsByConfidenceLevel struct {
	Finalized   []CommitPluginReportWithMeta `json:"finalized"`
	Unfinalized []CommitPluginReportWithMeta `json:"unfinalized"`
}

// sourceChainConfig is used to parse the response from the offRamp contract's getSourceChainConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L94
//
//nolint:lll // It's a URL.
type SourceChainConfig struct {
	Router                    []byte // local router
	IsEnabled                 bool
	IsRMNVerificationDisabled bool
	MinSeqNr                  uint64
	OnRamp                    UnknownAddress
}

// ContractAddresses is a map of contract names across all chain selectors and their address.
// Currently only one contract per chain per name is supported.
type ContractAddresses map[string]map[ChainSelector]UnknownAddress

// CurseInfo contains cursing information that are fetched from the rmn remote contract.
type CurseInfo struct {
	// CursedSourceChains contains the cursed source chains.
	CursedSourceChains map[ChainSelector]bool
	// CursedDestination indicates that the destination chain is cursed.
	CursedDestination bool
	// GlobalCurse indicates that all chains are cursed.
	GlobalCurse bool
}

func (ci CurseInfo) NonCursedSourceChains(inputChains []ChainSelector) []ChainSelector {
	if ci.GlobalCurse {
		return nil
	}

	sourceChains := make([]ChainSelector, 0, len(inputChains))
	for _, ch := range inputChains {
		if !ci.CursedSourceChains[ch] {
			sourceChains = append(sourceChains, ch)
		}
	}
	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })

	return sourceChains
}

// GlobalCurseSubject Defined as a const in RMNRemote.sol
// Docs of RMNRemote:
// An active curse on this subject will cause isCursed() and isCursed(bytes16) to return true. Use this subject
// for issues affecting all of CCIP chains, or pertaining to the chain that this contract is deployed on, instead of
// using the local chain selector as a subject.
var GlobalCurseSubject = [16]byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
}

// RemoteConfig contains the configuration fetched from the RMNRemote contract.
type RemoteConfig struct {
	ContractAddress UnknownAddress     `json:"contractAddress"`
	ConfigDigest    Bytes32            `json:"configDigest"`
	Signers         []RemoteSignerInfo `json:"signers"`
	// F defines the max number of faulty RMN nodes; F+1 signers are required to verify a report.
	FSign            uint64  `json:"fSign"` // previously: MinSigners
	ConfigVersion    uint32  `json:"configVersion"`
	RmnReportVersion Bytes32 `json:"rmnReportVersion"` // e.g., keccak256("RMN_V1_6_ANY2EVM_REPORT")
}

func (r RemoteConfig) IsEmpty() bool {
	// NOTE: contract address will always be present, since the code auto populates it
	return r.ConfigDigest == (Bytes32{}) &&
		len(r.Signers) == 0 &&
		r.FSign == 0 &&
		r.ConfigVersion == 0 &&
		r.RmnReportVersion == (Bytes32{})
}

// RemoteSignerInfo contains information about a signer from the RMNRemote contract.
type RemoteSignerInfo struct {
	// The signer's onchain address, used to verify report signature
	OnchainPublicKey UnknownAddress `json:"onchainPublicKey"`
	// The index of the node in the RMN config
	NodeIndex uint64 `json:"nodeIndex"`
}
