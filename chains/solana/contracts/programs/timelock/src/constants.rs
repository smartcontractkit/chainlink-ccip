/// PDA seeds
pub const TIMELOCK_CONFIG_SEED: &[u8] = b"timelock_config";
pub const TIMELOCK_OPERATION_SEED: &[u8] = b"timelock_operation";
pub const TIMELOCK_BYPASSER_OPERATION_SEED: &[u8] = b"timelock_bypasser_operation";
pub const TIMELOCK_SIGNER_SEED: &[u8] = b"timelock_signer";
pub const TIMELOCK_BLOCKED_FUNCITON_SELECTOR_SEED: &[u8] = b"timelock_blocked_function_selector";

/// constants
pub const ANCHOR_DISCRIMINATOR: usize = 8;
pub const EMPTY_PREDECESSOR: [u8; 32] = [0; 32];
pub const TIMELOCK_ID_PADDED: usize = 32; // fixed size timelock id for distinguishing different timelock states
pub const MAX_SELECTORS: usize = 128; // max number of function selectors that can be blocked(arrayvec)
