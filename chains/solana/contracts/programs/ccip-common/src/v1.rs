use anchor_lang::prelude::*;
use ethnum::U256;
use solana_program::address_lookup_table::state::AddressLookupTable;

use crate::{
    context::TokenAccountsValidationContext, router_accounts::TokenAdminRegistry, seed,
    CommonCcipError,
};

pub const MIN_TOKEN_POOL_ACCOUNTS: usize = 13; // see TokenAccounts struct for all required accounts
const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);
const EVM_PRECOMPILE_SPACE: u32 = 1024;

pub struct TokenAccounts<'a> {
    pub user_token_account: &'a AccountInfo<'a>,
    pub token_billing_config: &'a AccountInfo<'a>,
    pub pool_chain_config: &'a AccountInfo<'a>,
    pub lookup_table: &'a AccountInfo<'a>,
    pub token_admin_registry: &'a AccountInfo<'a>,
    pub pool_program: &'a AccountInfo<'a>,
    pub pool_config: &'a AccountInfo<'a>,
    pub pool_token_account: &'a AccountInfo<'a>,
    pub pool_signer: &'a AccountInfo<'a>,
    pub token_program: &'a AccountInfo<'a>,
    pub mint: &'a AccountInfo<'a>,
    pub fee_token_config: &'a AccountInfo<'a>,
    pub ccip_router_pool_signer: &'a AccountInfo<'a>,
    // as this one is optional, it doesn't count for the MIN_TOKEN_POOL_ACCOUNTS
    pub ccip_offramp_pool_signer: Option<&'a AccountInfo<'a>>,

    pub ccip_router_pool_signer_bump: u8,
    pub ccip_offramp_pool_signer_bump: u8,

    pub remaining_accounts: &'a [AccountInfo<'a>],
}

pub fn validate_and_parse_token_accounts<'info>(
    token_receiver: Pubkey,
    chain_selector: u64,
    router: Pubkey,
    fee_quoter: Pubkey,
    offramp: Option<Pubkey>, // id of the offramp program that called this function, when the caller is not the router
    raw_acc_infos: &'info [AccountInfo<'info>],
) -> Result<TokenAccounts> {
    // The program_id here is provided solely to satisfy the interface of try_accounts.
    // Note: All program IDs for PDA derivation are explicitly defined in the account context
    // (TokenAccountsValidationContext) via seeds and program attributes.
    // Therefore, the value of program_id (set here to Pubkey::default()) is effectively unused.
    // Changes in environment-specific program addresses will not affect the PDA derivation.
    let program_id = Pubkey::default();

    let mut accounts = raw_acc_infos;

    let ccip_offramp_pool_signer = if offramp.is_some() {
        // When and only when the offramp program is calling this function, the first account is the offramp pool signer.
        // Then, the rest of the accounts are the same as when the router program is calling this function.
        accounts = &accounts[1..];
        Some(&raw_acc_infos[0])
    } else {
        None
    };
    let mut input_accounts = accounts;

    let bumps = &mut <TokenAccountsValidationContext as anchor_lang::Bumps>::Bumps::default();
    let reallocs = &mut std::collections::BTreeSet::new();

    // leveraging Anchor's account context validation
    // Instead of manually checking each account (ownership, PDA derivation, constraints),
    // we're using Anchor's `try_accounts` to perform these validations based on the
    // constraints defined in the `TokenAccountsValidationContext` account context struct
    TokenAccountsValidationContext::try_accounts(
        &program_id,
        &mut input_accounts,
        &[
            token_receiver.as_ref(),
            &chain_selector.to_le_bytes(),
            router.as_ref(),
            fee_quoter.as_ref(),
        ]
        .concat(),
        bumps,
        reallocs,
    )?;

    let mut accounts_iter = accounts.iter();

    // accounts based on user or chain
    let user_token_account = next_account_info(&mut accounts_iter)?;
    let token_billing_config = next_account_info(&mut accounts_iter)?;
    let pool_chain_config = next_account_info(&mut accounts_iter)?;

    // constant accounts for any pool interaction
    let lookup_table = next_account_info(&mut accounts_iter)?;
    let token_admin_registry = next_account_info(&mut accounts_iter)?;
    let pool_program = next_account_info(&mut accounts_iter)?;
    let pool_config = next_account_info(&mut accounts_iter)?;
    let pool_token_account = next_account_info(&mut accounts_iter)?;
    let pool_signer = next_account_info(&mut accounts_iter)?;
    let token_program = next_account_info(&mut accounts_iter)?;
    let mint = next_account_info(&mut accounts_iter)?;
    let fee_token_config = next_account_info(&mut accounts_iter)?;
    let ccip_router_pool_signer = next_account_info(&mut accounts_iter)?;

    // As the offramp signer is optional (only used when the offramp program is calling this function),
    // the context cannot be used to validate it, as it could try to apply those validations to the wrong account
    let ccip_offramp_pool_signer_bump = if let Some(offramp) = offramp {
        let (expected_offramp_pool_signer, bump) = Pubkey::find_program_address(
            &[
                seed::EXTERNAL_TOKEN_POOLS_SIGNER,
                pool_program.key().as_ref(),
            ],
            &offramp,
        );
        let acc_info = ccip_offramp_pool_signer.unwrap();
        require_keys_eq!(
            acc_info.key(),
            expected_offramp_pool_signer,
            CommonCcipError::InvalidInputsPoolAccounts
        );
        bump
    } else {
        0
    };

    // collect remaining accounts
    let remaining_accounts = accounts_iter.as_slice();

    // Additional validations that can't be expressed in the account context
    {
        // Check Lookup Table Address configured in TokenAdminRegistry
        let token_admin_registry_account: Account<TokenAdminRegistry> =
            Account::try_from(token_admin_registry)?;
        require_keys_eq!(
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
        // base set of constant accounts (10)
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
            ccip_router_pool_signer,
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
        lookup_table,
        token_admin_registry,
        pool_program,
        pool_config,
        pool_token_account,
        pool_signer,
        token_program,
        mint,
        fee_token_config,
        ccip_router_pool_signer,
        ccip_offramp_pool_signer,
        ccip_router_pool_signer_bump: bumps.ccip_router_pool_signer,
        ccip_offramp_pool_signer_bump,
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

// address validation helpers based on the chain family selector
pub fn validate_evm_address(address: &[u8]) -> Result<()> {
    require_eq!(address.len(), 32, CommonCcipError::InvalidEVMAddress);

    let address: U256 = U256::from_be_bytes(
        address
            .try_into()
            .map_err(|_| CommonCcipError::InvalidEncoding)?,
    );
    require!(address <= U160_MAX, CommonCcipError::InvalidEVMAddress);
    if let Ok(small_address) = TryInto::<u32>::try_into(address) {
        require_gte!(
            small_address,
            EVM_PRECOMPILE_SPACE,
            CommonCcipError::InvalidEVMAddress
        )
    };
    Ok(())
}

pub fn validate_svm_address(address: &[u8], address_must_be_nonzero: bool) -> Result<()> {
    require_eq!(address.len(), 32, CommonCcipError::InvalidSVMAddress);
    require!(
        !address_must_be_nonzero || address.iter().any(|b| *b != 0),
        CommonCcipError::InvalidSVMAddress
    );

    Pubkey::try_from_slice(address)
        .map(|_| ())
        .map_err(|_| CommonCcipError::InvalidSVMAddress.into())
}
