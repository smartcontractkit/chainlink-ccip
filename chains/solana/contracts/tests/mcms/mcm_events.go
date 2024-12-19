package contracts

import (
	"github.com/gagliardetto/solana-go"
)

// Events - temporary event struct to decode
// anchor-go does not support events
// https://github.com/fragmetric-labs/solana-anchor-go does but requires upgrade to anchor >= v0.30.0

// NewRoot represents an event emitted when a new root is set
type NewRoot struct {
	Root       [32]byte // root
	ValidUntil uint32   // valid_until

	// Metadata fields
	MetadataChainID              uint64           // metadata_chain_id
	MetadataMultisig             solana.PublicKey // metadata_multisig
	MetadataPreOpCount           uint64           // metadata_pre_op_count
	MetadataPostOpCount          uint64           // metadata_post_op_count
	MetadataOverridePreviousRoot bool             // metadata_override_previous_root
}

const numGroups = 32

// ConfigSet represents an event emitted when a new config is set
type ConfigSet struct {
	// Note: Rust comment indicates signers are omitted due to memory overflow
	GroupParents  [numGroups]byte // group_parents
	GroupQuorums  [numGroups]byte // group_quorums
	IsRootCleared bool            // is_root_cleared
}

// OpExecuted represents an event emitted when an op is successfully executed
type OpExecuted struct {
	Nonce uint64           // nonce
	To    solana.PublicKey // to
	Data  []byte           // data: Vec<u8>
}
