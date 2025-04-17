// Fixed seeds - different contexts must use different PDA seeds
pub const CONFIG: &[u8] = b"config";
pub const ALLOWED_PRICE_UPDATER: &[u8] = b"allowed_price_updater";
pub const DEST_CHAIN: &[u8] = b"dest_chain";
pub const FEE_BILLING_TOKEN_CONFIG: &[u8] = b"fee_billing_token_config";
pub const PER_CHAIN_PER_TOKEN_CONFIG: &[u8] = b"per_chain_per_token_config";
pub const SOURCE_CHAIN: &[u8] = b"source_chain_state";
pub const COMMIT_REPORT: &[u8] = b"commit_report";
pub const REFERENCE_ADDRESSES: &[u8] = b"reference_addresses";
pub const STATE: &[u8] = b"state";
pub const CURSES: &[u8] = b"curses";
pub const DEST_CHAIN_STATE: &[u8] = b"dest_chain_state";
pub const NONCE: &[u8] = b"nonce";
pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";

// arbitrary messaging signer
pub const EXTERNAL_EXECUTION_CONFIG: &[u8] = b"external_execution_config";
// token pool interaction signer
pub const EXTERNAL_TOKEN_POOLS_SIGNER: &[u8] = b"external_token_pools_signer";
// signer for billing fee token transfer
pub const FEE_BILLING_SIGNER: &[u8] = b"fee_billing_signer";

// token specific
pub const TOKEN_ADMIN_REGISTRY: &[u8] = b"token_admin_registry";
pub const CCIP_TOKENPOOL_CONFIG: &[u8] = b"ccip_tokenpool_config";
pub const CCIP_TOKENPOOL_SIGNER: &[u8] = b"ccip_tokenpool_signer";
pub const TOKEN_POOL_CONFIG: &[u8] = b"ccip_tokenpool_chainconfig";
