use anchor_lang::prelude::*;
use anchor_spl::token::spl_token::native_mint;

use crate::messages::SVM2AnyMessage;
use crate::state::{BillingTokenConfigWrapper, Config, DestChain};
use crate::FeeQuoterError;

// Fixed seeds - different contexts must use different PDA seeds
pub const DEST_CHAIN_STATE_SEED: &[u8] = b"dest_chain_state";
pub const CONFIG_SEED: &[u8] = b"config";
pub const FEE_BILLING_TOKEN_CONFIG_SEED: &[u8] = b"fee_billing_token_config";
pub const TOKEN_POOL_BILLING_SEED: &[u8] = b"ccip_tokenpool_billing";

// valid_version validates that the passed in version is not 0 (uninitialized)
// and it is within the expected maximum supported version bounds
pub fn valid_version(v: u8, max_v: u8) -> bool {
    v != 0 && v <= max_v
}
pub fn uninitialized(v: u8) -> bool {
    v == 0
}

const MAX_CONFIG_V: u8 = 1;
const MAX_CHAINSTATE_V: u8 = 1;

#[derive(Accounts)]
#[instruction(destination_chain_selector: u64, message: SVM2AnyMessage)]
pub struct GetFee<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,
    #[account(
        seeds = [DEST_CHAIN_STATE_SEED, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub dest_chain: Account<'info, DestChain>,

    #[account(
        seeds = [FEE_BILLING_TOKEN_CONFIG_SEED,
            if message.fee_token == Pubkey::default() {
                native_mint::ID.as_ref() // pre-2022 WSOL
            } else {
                message.fee_token.as_ref()
            }
        ],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,
}
