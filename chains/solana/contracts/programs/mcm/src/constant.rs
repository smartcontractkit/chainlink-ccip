// Business-logic constants
pub const MAX_NUM_SIGNERS: usize = 180; // maximum number of signers supported
pub const NUM_GROUPS: usize = 32; // maximum number of signer groups supported

// fixed size msig id for distinguishing different multisig instances
pub const MULTISIG_ID_PADDED: usize = 32;

// PDA seeds
// Note: These seeds are not full seed, for unique seeds per instance, multisig_id should be appended
pub const SIGNER_SEED: &[u8] = b"multisig_signer"; // seed for dataless pda signing CPI
pub const CONFIG_SEED: &[u8] = b"multisig_config";
pub const CONFIG_SIGNERS_SEED: &[u8] = b"multisig_config_signers";
pub const ROOT_METADATA_SEED: &[u8] = b"root_metadata";
pub const ROOT_SIGNATURES_SEED: &[u8] = b"root_signatures";

pub const EXPIRING_ROOT_AND_OP_COUNT_SEED: &[u8] = b"expiring_root_and_op_count";
pub const SEEN_SIGNED_HASHES_SEED: &[u8] = b"seen_signed_hashes";

// Fixed sizes in bytes
pub const ANCHOR_DISCRIMINATOR: usize = 8;
