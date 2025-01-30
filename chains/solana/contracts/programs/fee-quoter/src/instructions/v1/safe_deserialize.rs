/// Methods in this module are used to deserialize AccountInfo into the state structs
use anchor_lang::prelude::*;

use crate::context::seed::{FEE_BILLING_TOKEN_CONFIG, PER_CHAIN_PER_TOKEN_CONFIG};
use crate::state::{BillingTokenConfig, BillingTokenConfigWrapper, PerChainPerTokenConfig};
use crate::FeeQuoterError;

pub fn per_chain_per_token_config<'info>(
    account: &'info AccountInfo<'info>,
    token: Pubkey,
    dest_chain_selector: u64,
) -> Result<PerChainPerTokenConfig> {
    let (expected, _) = Pubkey::find_program_address(
        &[
            PER_CHAIN_PER_TOKEN_CONFIG,
            dest_chain_selector.to_le_bytes().as_ref(),
            token.key().as_ref(),
        ],
        &crate::ID,
    );
    require_keys_eq!(account.key(), expected, FeeQuoterError::InvalidInputs);
    let account = Account::<PerChainPerTokenConfig>::try_from(account)?;
    require_eq!(
        account.version,
        1, // the v1 version of the onramp will always be tied to version 1 of the state
        FeeQuoterError::InvalidInputs
    );
    Ok(account.into_inner())
}

// Returns Ok(None) when parsing the ZERO address, which is a valid input from users
// specifying a token that has no Billing config.
pub fn billing_token_config<'info>(
    account: &'info AccountInfo<'info>,
    token: Pubkey,
) -> Result<Option<BillingTokenConfig>> {
    if account.key() == Pubkey::default() {
        return Ok(None);
    }

    let (expected, _) =
        Pubkey::find_program_address(&[FEE_BILLING_TOKEN_CONFIG, token.as_ref()], &crate::ID);
    require_keys_eq!(account.key(), expected, FeeQuoterError::InvalidInputs);
    let account = Account::<BillingTokenConfigWrapper>::try_from(account)?;
    require_eq!(
        account.version,
        1, // the v1 version of the onramp will always be tied to version 1 of the state
        FeeQuoterError::InvalidInputs
    );
    Ok(Some(account.into_inner().config))
}
