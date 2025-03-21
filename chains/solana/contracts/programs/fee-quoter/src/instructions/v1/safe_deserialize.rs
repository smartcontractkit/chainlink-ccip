/// Methods in this module are used to deserialize AccountInfo into the state structs
use anchor_lang::prelude::*;
use ccip_common::seed;

use crate::state::{BillingTokenConfig, BillingTokenConfigWrapper, PerChainPerTokenConfig};
use crate::FeeQuoterError;

/// Returns Ok(None) when there's no chain-token configuration pair. The user must still
/// specify the correct PDA, to ensure the configuration isn't ignored in the case it exists.
pub fn per_chain_per_token_config<'info>(
    account: &'info AccountInfo<'info>,
    token: Pubkey,
    dest_chain_selector: u64,
) -> Result<Option<PerChainPerTokenConfig>> {
    let (expected, _) = Pubkey::find_program_address(
        &[
            seed::PER_CHAIN_PER_TOKEN_CONFIG,
            dest_chain_selector.to_le_bytes().as_ref(),
            token.key().as_ref(),
        ],
        &crate::ID,
    );
    require_keys_eq!(
        account.key(),
        expected,
        FeeQuoterError::InvalidInputsPerChainPerTokenConfig
    );
    let account = match Account::<PerChainPerTokenConfig>::try_from(account) {
        Ok(account) => {
            require_eq!(
                account.version,
                1, // the v1 version of the onramp will always be tied to version 1 of the state
                FeeQuoterError::InvalidVersion
            );
            Some(account.into_inner())
        }
        Err(e) if e == ErrorCode::AccountNotInitialized.into() => None,
        Err(e) => return Err(e),
    };

    Ok(account)
}

/// Returns Ok(None) when there's no token specific billing config. The user must still
/// specify the correct PDA, to ensure the configuration isn't ignored in the case it exists.
pub fn billing_token_config<'info>(
    account: &'info AccountInfo<'info>,
    token: Pubkey,
) -> Result<Option<BillingTokenConfig>> {
    let (expected, _) = Pubkey::find_program_address(
        &[seed::FEE_BILLING_TOKEN_CONFIG, token.as_ref()],
        &crate::ID,
    );
    require_keys_eq!(
        account.key(),
        expected,
        FeeQuoterError::InvalidInputsBillingTokenConfig
    );

    let config = match Account::<BillingTokenConfigWrapper>::try_from(account) {
        Ok(account) => {
            require_eq!(
                account.version,
                1, // the v1 version of the onramp will always be tied to version 1 of the state
                FeeQuoterError::InvalidVersion
            );
            Some(account.into_inner().config)
        }
        Err(e) if e == ErrorCode::AccountNotInitialized.into() => None,
        Err(e) => return Err(e),
    };

    Ok(config)
}
