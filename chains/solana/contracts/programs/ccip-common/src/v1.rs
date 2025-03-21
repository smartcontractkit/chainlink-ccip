use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use solana_program::address_lookup_table::state::AddressLookupTable;

use crate::{router_accounts::TokenAdminRegistry, seed, CommonCcipError};

pub struct TokenAccounts<'a> {
    pub user_token_account: &'a AccountInfo<'a>,
    pub token_billing_config: &'a AccountInfo<'a>,
    pub pool_chain_config: &'a AccountInfo<'a>,
    pub pool_program: &'a AccountInfo<'a>,
    pub pool_config: &'a AccountInfo<'a>,
    pub pool_token_account: &'a AccountInfo<'a>,
    pub pool_signer: &'a AccountInfo<'a>,
    pub token_program: &'a AccountInfo<'a>,
    pub mint: &'a AccountInfo<'a>,
    pub fee_token_config: &'a AccountInfo<'a>,
    pub remaining_accounts: &'a [AccountInfo<'a>],
}

pub fn validate_and_parse_token_accounts<'info>(
    token_receiver: Pubkey,
    chain_selector: u64,
    router: Pubkey,
    fee_quoter: Pubkey,
    accounts: &'info [AccountInfo<'info>],
) -> Result<TokenAccounts> {
    // accounts based on user or chain
    let (user_token_account, remaining_accounts) = accounts.split_first().unwrap();
    let (token_billing_config, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (pool_chain_config, remaining_accounts) = remaining_accounts.split_first().unwrap();

    // constant accounts for any pool interaction
    let (lookup_table, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (token_admin_registry, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (pool_program, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (pool_config, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (pool_token_account, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (pool_signer, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (token_program, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (mint, remaining_accounts) = remaining_accounts.split_first().unwrap();
    let (fee_token_config, remaining_accounts) = remaining_accounts.split_first().unwrap();

    // Account validations (using remaining_accounts does not facilitate built-in anchor checks)
    {
        // Check Token Admin Registry
        let (expected_token_admin_registry, _) = Pubkey::find_program_address(
            &[seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
            &router,
        );
        require_eq!(
            token_admin_registry.key(),
            expected_token_admin_registry,
            CommonCcipError::InvalidInputsTokenAdminRegistryAccounts
        );

        // check pool program + pool config + pool signer
        let (expected_pool_config, _) = Pubkey::find_program_address(
            &[seed::CCIP_TOKENPOOL_CONFIG, mint.key().as_ref()],
            &pool_program.key(),
        );
        let (expected_pool_signer, _) = Pubkey::find_program_address(
            &[seed::CCIP_TOKENPOOL_SIGNER, mint.key().as_ref()],
            &pool_program.key(),
        );
        require_eq!(
            *pool_config.owner,
            pool_program.key(),
            CommonCcipError::InvalidInputsPoolAccounts
        );
        require_eq!(
            pool_config.key(),
            expected_pool_config,
            CommonCcipError::InvalidInputsPoolAccounts
        );
        require_eq!(
            pool_signer.key(),
            expected_pool_signer,
            CommonCcipError::InvalidInputsPoolAccounts
        );

        let (expected_fee_token_config, _) = Pubkey::find_program_address(
            &[seed::FEE_BILLING_TOKEN_CONFIG, mint.key.as_ref()],
            &fee_quoter,
        );
        require_eq!(
            fee_token_config.key(),
            expected_fee_token_config,
            CommonCcipError::InvalidInputsConfigAccounts
        );

        // check token accounts
        require_eq!(
            *mint.owner,
            token_program.key(),
            CommonCcipError::InvalidInputsTokenAccounts
        );
        require_eq!(
            user_token_account.key(),
            get_associated_token_address_with_program_id(
                &token_receiver,
                &mint.key(),
                &token_program.key()
            ),
            CommonCcipError::InvalidInputsTokenAccounts
        );
        require_eq!(
            pool_token_account.key(),
            get_associated_token_address_with_program_id(
                &pool_signer.key(),
                &mint.key(),
                &token_program.key()
            ),
            CommonCcipError::InvalidInputsTokenAccounts
        );

        // check per token per chain configs
        // billing: configured via CCIP fee quoter
        // chain config: configured via pool
        let (expected_billing_config, _) = Pubkey::find_program_address(
            &[
                seed::PER_CHAIN_PER_TOKEN_CONFIG,
                chain_selector.to_le_bytes().as_ref(),
                mint.key().as_ref(),
            ],
            &fee_quoter,
        );
        let (expected_pool_chain_config, _) = Pubkey::find_program_address(
            &[
                seed::TOKEN_POOL_CONFIG,
                chain_selector.to_le_bytes().as_ref(),
                mint.key().as_ref(),
            ],
            &pool_program.key(),
        );
        require_eq!(
            token_billing_config.key(),
            expected_billing_config, // TODO: determine if this can be zero key for optional billing config?
            CommonCcipError::InvalidInputsConfigAccounts
        );
        require_eq!(
            pool_chain_config.key(),
            expected_pool_chain_config,
            CommonCcipError::InvalidInputsConfigAccounts
        );

        // Check Lookup Table Address configured in TokenAdminRegistry
        let token_admin_registry_account: Account<TokenAdminRegistry> =
            Account::try_from(token_admin_registry)?;
        require_eq!(
            token_admin_registry_account.lookup_table,
            lookup_table.key(),
            CommonCcipError::InvalidInputsLookupTableAccounts
        );

        // Check Lookup Table Entries
        let lookup_table_data = &mut &lookup_table.data.borrow()[..];
        let lookup_table_account: AddressLookupTable =
            AddressLookupTable::deserialize(lookup_table_data)
                .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;

        // reconstruct + validate expected values in token pool lookup table
        // base set of constant accounts (9)
        // + additional constant accounts (remaining_accounts) that are not required but may be used for additional token pool functionality (like CPI)
        let required_entries = [
            lookup_table,
            token_admin_registry,
            pool_program,
            pool_config,
            pool_token_account,
            pool_signer,
            token_program,
            mint,
            fee_token_config,
        ];
        {
            // validate pool addresses
            let mut expected_keys: Vec<Pubkey> = required_entries.iter().map(|x| x.key()).collect();
            let mut remaining_keys: Vec<Pubkey> =
                remaining_accounts.iter().map(|x| x.key()).collect();
            expected_keys.append(&mut remaining_keys);
            require!(
                lookup_table_account.addresses.as_ref() == expected_keys,
                CommonCcipError::InvalidInputsLookupTableAccounts
            );
        }
        {
            // validate pool address writable
            // token admin registry contains an array (binary) of indexes that are writable
            // check that the writability of the passed accounts match the writable configuration (using indexes)
            let mut expected_is_writable: Vec<bool> =
                required_entries.iter().map(|x| x.is_writable).collect();
            let mut remaining_is_writable: Vec<bool> =
                remaining_accounts.iter().map(|x| x.is_writable).collect();
            expected_is_writable.append(&mut remaining_is_writable);
            for (i, is_writable) in expected_is_writable.iter().enumerate() {
                require_eq!(
                    token_admin_registry_writable::is(&token_admin_registry_account, i as u8),
                    *is_writable,
                    CommonCcipError::InvalidInputsLookupTableAccountWritable
                );
            }
        }
    }

    Ok(TokenAccounts {
        user_token_account,
        token_billing_config,
        pool_chain_config,
        pool_program,
        pool_config,
        pool_token_account,
        pool_signer,
        token_program,
        mint,
        fee_token_config,
        remaining_accounts,
    })
}

pub mod token_admin_registry_writable {
    use crate::router_accounts::TokenAdminRegistry;

    // set writable inserts bits from left to right
    // index 0 is left-most bit
    pub fn set(tar: &mut TokenAdminRegistry, index: u8) {
        match index < 128 {
            true => {
                tar.writable_indexes[0] |= 1 << (127 - index);
            }
            false => {
                tar.writable_indexes[1] |= 1 << (255 - index);
            }
        }
    }

    pub fn is(tar: &TokenAdminRegistry, index: u8) -> bool {
        match index < 128 {
            true => tar.writable_indexes[0] & 1 << (127 - index) != 0,
            false => tar.writable_indexes[1] & 1 << (255 - index) != 0,
        }
    }

    pub fn reset(tar: &mut TokenAdminRegistry) {
        tar.writable_indexes = [0, 0];
    }

    #[cfg(test)]
    mod tests {
        use super::*;
        use solana_program::pubkey::Pubkey;

        #[test]
        fn set_writable() {
            let state = &mut TokenAdminRegistry {
                version: 0,
                administrator: Pubkey::default(),
                pending_administrator: Pubkey::default(),
                lookup_table: Pubkey::default(),
                writable_indexes: [0, 0],
                mint: Pubkey::default(),
            };

            set(state, 0);
            set(state, 128);
            assert_eq!(state.writable_indexes[0], 2u128.pow(127));
            assert_eq!(state.writable_indexes[1], 2u128.pow(127));

            reset(state);
            assert_eq!(state.writable_indexes[0], 0);
            assert_eq!(state.writable_indexes[1], 0);

            set(state, 0);
            set(state, 2);
            set(state, 127);
            set(state, 128);
            set(state, 2 + 128);
            set(state, 255);
            assert_eq!(
                state.writable_indexes[0],
                2u128.pow(127) + 2u128.pow(127 - 2) + 2u128.pow(0)
            );
            assert_eq!(
                state.writable_indexes[1],
                2u128.pow(127) + 2u128.pow(127 - 2) + 2u128.pow(0)
            );
        }

        #[test]
        fn check_writable() {
            let state = &TokenAdminRegistry {
                version: 0,
                administrator: Pubkey::default(),
                pending_administrator: Pubkey::default(),
                lookup_table: Pubkey::default(),
                writable_indexes: [
                    2u128.pow(127 - 7) + 2u128.pow(127 - 2) + 2u128.pow(127 - 4),
                    2u128.pow(127 - 8) + 2u128.pow(127 - 56) + 2u128.pow(127 - 100),
                ],
                mint: Pubkey::default(),
            };

            assert!(!is(state, 0));
            assert!(!is(state, 128));
            assert!(!is(state, 255));

            assert_eq!(state.writable_indexes[0].count_ones(), 3);
            assert_eq!(state.writable_indexes[1].count_ones(), 3);

            assert!(is(state, 7));
            assert!(is(state, 2));
            assert!(is(state, 4));
            assert!(is(state, 128 + 8));
            assert!(is(state, 128 + 56));
            assert!(is(state, 128 + 100));
        }
    }
}
