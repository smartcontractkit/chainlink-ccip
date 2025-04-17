package ccip

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
)

// Events - temporary event struct to decode
// anchor-go does not support events
// https://github.com/fragmetric-labs/solana-anchor-go does but requires upgrade to anchor >= v0.30.0
type EventCCIPMessageSent struct {
	Discriminator            [8]byte
	DestinationChainSelector uint64
	SequenceNumber           uint64
	Message                  ccip_router.SVM2AnyRampMessage
}

type EventCommitReportAccepted struct {
	Discriminator [8]byte
	Report        *ccip_offramp.MerkleRoot `bin:"optional"`
	PriceUpdates  ccip_offramp.PriceUpdates
}

type EventCommitReportPDAClosed struct {
	Discriminator       [8]byte
	SourceChainSelector uint64
	MerkleRoot          [32]byte
}

type EventTransmitted struct {
	Discriminator  [8]byte
	OcrPluginType  uint8
	ConfigDigest   [32]byte
	SequenceNumber uint64
}

type EventExecutionStateChanged struct {
	Discriminator       [8]byte
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageID           [32]byte
	MessageHash         [32]byte
	State               ccip_offramp.MessageExecutionState
}

type EventSkippedAlreadyExecutedMessage struct {
	Discriminator       [8]byte
	SourceChainSelector uint64
	SequenceNumber      uint64
}

type EventConfigSet struct {
	Discriminator [8]byte
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       [][20]uint8
	Transmitters  []solana.PublicKey
	F             uint8
}

// FeeQuoter-specific event
type UsdPerTokenUpdated struct {
	Discriminator [8]byte
	Token         solana.PublicKey
	Value         [28]byte
	Timestamp     int64
}

// FeeQuoter-specific event
type TokenPriceUpdateIgnored struct {
	Discriminator [8]byte
	Token         solana.PublicKey
	Value         [28]byte
}

// FeeQuoter-specific event
type UsdPerUnitGasUpdated struct {
	Discriminator [8]byte
	DestChain     uint64
	Value         [28]byte
	Timestamp     int64
}

type OwnershipTransferRequested struct {
	Discriminator [8]byte
	From          solana.PublicKey
	To            solana.PublicKey
}

type OwnershipTransferred struct {
	Discriminator [8]byte
	From          solana.PublicKey
	To            solana.PublicKey
}

type PremiumMultiplierWeiPerEthUpdated struct {
	Discriminator              [8]byte
	Token                      solana.PublicKey
	PremiumMultiplierWeiPerEth uint64
}

type TokenTransferFeeConfigUpdated struct {
	Discriminator            [8]byte
	DestinationChainSelector uint64
	Token                    solana.PublicKey
	TokenTransferFeeConfig   fee_quoter.TokenTransferFeeConfig
}

type EventOfframpConfigSet struct {
	Discriminator              [8]byte
	SvmChainSelector           uint64
	EnableManualExecutionAfter int64
}

type EventOfframpReferenceAddressesSet struct {
	Discriminator      [8]byte
	Router             solana.PublicKey
	FeeQuoter          solana.PublicKey
	OfframpLookupTable solana.PublicKey
	RMNRemote          solana.PublicKey
}

type EventRouterConfigSet struct {
	Discriminator    [8]byte
	SvmChainSelector uint64
	FeeQuoter        solana.PublicKey
	RMNRemote        solana.PublicKey
	LinkTokenMint    solana.PublicKey
	FeeAggregator    solana.PublicKey
}

type EventFeeQuoterConfigSet struct {
	Discriminator      [8]byte
	MaxFeeJuelsPerMsg  bin.Uint128
	LinkTokenMint      solana.PublicKey
	LinkTokenDecimals  uint8
	Onramp             solana.PublicKey
	DefaultCodeVersion fee_quoter.CodeVersion
}

type EventReferenceAddressesSet struct {
	Discriminator      [8]byte
	Router             solana.PublicKey
	FeeQuoter          solana.PublicKey
	OfframpLookupTable solana.PublicKey
}
