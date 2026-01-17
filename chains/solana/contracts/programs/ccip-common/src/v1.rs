use anchor_lang::prelude::*;
use anchor_lang::Discriminator;
use ethnum::U256;
use solana_program::address_lookup_table::state::AddressLookupTable;

use crate::router_accounts::TokenAdminRegistryV1;
use crate::{
    context::TokenAccountsValidationContext, router_accounts::TokenAdminRegistry, seed,
    CommonCcipError,
};

pub const MIN_TOKEN_POOL_ACCOUNTS: usize = 13; // see TokenAccounts struct for all required accounts
const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);
const EVM_PRECOMPILE_SPACE: u32 = 1024;
pub const V1_TOKEN_ADMIN_REGISTRY_SIZE: usize = 169; // for migration v1->v2 of the TokenAdminRegistry, which adds the `supports_auto_derivation` field.

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
) -> Result<TokenAccounts<'info>> {
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
        let token_admin_registry_account = load_token_admin_registry_checked(token_admin_registry)?;
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
            if token_admin_registry_account.supports_auto_derivation {
                // When auto-derivation is supported, there may be more accounts than what the lookup table contains.
                // On those extra accounts, given that the token pool supports derivation, we also trust that the pool
                // performs the required checks.
                let n = lookup_table_account.addresses.len();
                require!(
                    lookup_table_account.addresses.as_ref() == expected_keys[..n].to_vec(),
                    CommonCcipError::InvalidInputsLookupTableAccounts
                );
            } else {
                // When the pool does not support auto-derivation, the lookup table must contain all the expected accounts,
                // and CCIP is expected to validate them.
                require!(
                    lookup_table_account.addresses.as_ref() == expected_keys,
                    CommonCcipError::InvalidInputsLookupTableAccounts
                );
            }
        }
        {
            // validate pool address writable
            // token admin registry contains an array (binary) of indexes that are writable
            // check that the writability of the passed accounts match the writable configuration (using indexes)
            let mut expected_is_writable: Vec<bool> =
                required_entries.iter().map(|x| x.is_writable).collect();

            let mut remaining_is_writable: Vec<bool> =
                if token_admin_registry_account.supports_auto_derivation {
                    // when auto-derivation is supported, we only validate the accounts present in the LUT
                    // as the pool is expected to validate the rest of the accounts. The LUT must contain
                    // the required entries, plus it may contain some additional ones that are static,
                    // and finally the auto-derivation may add more accounts at the end (static or message-dependant).
                    let lut_len = lookup_table_account.addresses.len();
                    let end = lut_len.checked_sub(required_entries.len()).unwrap();

                    remaining_accounts[..end]
                        .iter()
                        .map(|x| x.is_writable)
                        .collect()
                } else {
                    // when auto-derivation is not supported, we validate all remaining accounts
                    // as they are all expected to be in the LUT
                    remaining_accounts.iter().map(|x| x.is_writable).collect()
                };

            expected_is_writable.append(&mut remaining_is_writable);
            for (i, is_writable) in expected_is_writable.iter().enumerate() {
                let expected =
                    token_admin_registry_writable::is(&token_admin_registry_account, i as u8);
                let actual = *is_writable;
                if expected != actual {
                    msg!(
                        "Expected account at index {} to be writable {}, but it is {}",
                        i,
                        expected,
                        actual
                    );
                    return Err(CommonCcipError::InvalidInputsLookupTableAccountWritable.into());
                }
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

pub fn load_token_admin_registry_checked<'info>(
    token_admin_registry: &'info AccountInfo<'info>,
) -> Result<TokenAdminRegistry> {
    Ok(
        if token_admin_registry.data_len() == V1_TOKEN_ADMIN_REGISTRY_SIZE {
            load_v1_token_admin_registry(token_admin_registry)?
        } else {
            // this uses Anchor's built-in deserialization that already has ownership and discriminator checks
            Account::<TokenAdminRegistry>::try_from(token_admin_registry)?.into_inner()
        },
    )
}

fn load_v1_token_admin_registry(
    token_admin_registry: &AccountInfo<'_>,
) -> Result<TokenAdminRegistry> {
    const ANCHOR_DISCRIMINATOR_SIZE: usize = 8;

    let borrowed_data = token_admin_registry.try_borrow_data()?;
    let (discriminator, data) = borrowed_data.split_at(ANCHOR_DISCRIMINATOR_SIZE);

    require_keys_eq!(
        token_admin_registry.owner.key(),
        crate::ID, // ccip-common crate must have the same ID as ccip-router, which is the actual program that owns it
        CommonCcipError::InvalidInputsTokenAdminRegistryAccounts
    );

    require!(
        TokenAdminRegistry::DISCRIMINATOR == discriminator,
        CommonCcipError::InvalidInputsTokenAdminRegistryAccounts
    );

    let token_admin_registry_v1 = TokenAdminRegistryV1::deserialize(&mut &data[..])
        .map_err(|_| CommonCcipError::InvalidInputsTokenAdminRegistryAccounts)?;

    TokenAdminRegistry::try_from(token_admin_registry_v1)
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
                supports_auto_derivation: false,
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
                supports_auto_derivation: false,
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

#[cfg(test)]
mod tests {
    use super::*;

    fn get_expected() -> TokenAdminRegistry {
        TokenAdminRegistry {
            version: 1,
            administrator: Pubkey::new_unique(),
            pending_administrator: Pubkey::new_unique(),
            lookup_table: Pubkey::new_unique(),
            writable_indexes: [1, 2],
            mint: Pubkey::new_unique(),
            supports_auto_derivation: false,
        }
    }

    fn to_v1_bytes(t: &TokenAdminRegistry) -> Vec<u8> {
        let mut bytes = t.try_to_vec().unwrap();
        bytes.pop(); // remove the `supports_auto_derivation` field, as this is not part of the v1 data
        bytes
    }

    fn to_acc_info<'a>(
        key: &'a Pubkey,
        lamports: &'a mut u64,
        data: &'a mut [u8],
    ) -> AccountInfo<'a> {
        AccountInfo::new(key, false, true, lamports, data, &crate::ID, false, 0)
    }

    #[test]
    fn test_valid_load_v1_token_admin_registry() {
        let expected = get_expected();

        let bytes = to_v1_bytes(&expected);

        let mut data = vec![0u8; V1_TOKEN_ADMIN_REGISTRY_SIZE];
        data[0..8].copy_from_slice(&TokenAdminRegistry::DISCRIMINATOR);
        data[8..].copy_from_slice(bytes.as_slice());

        let (mut lamports, key) = (0u64, Pubkey::new_unique());
        let account_info = to_acc_info(&key, &mut lamports, &mut data);
        let token_admin_registry = load_v1_token_admin_registry(&account_info).unwrap();

        assert_eq!(token_admin_registry, expected);
    }

    #[test]
    fn test_invalid_discriminator() {
        let source = get_expected();
        let bytes = to_v1_bytes(&source);

        let mut data = vec![0u8; V1_TOKEN_ADMIN_REGISTRY_SIZE];
        data[0..8].copy_from_slice(&TokenAdminRegistry::DISCRIMINATOR);
        data[2] = 2; // change the discriminator to an invalid one
        data[8..].copy_from_slice(bytes.as_slice());

        let (mut lamports, key) = (0u64, Pubkey::new_unique());
        let account_info = to_acc_info(&key, &mut lamports, &mut data);

        let result = load_v1_token_admin_registry(&account_info);
        assert!(result.is_err());
    }

    #[test]
    fn test_invalid_version() {
        for v in [0, 2, 255] {
            let mut source = get_expected();
            source.version = v; // change the version to an invalid one for this function
            let bytes = to_v1_bytes(&source);

            let mut data = vec![0u8; V1_TOKEN_ADMIN_REGISTRY_SIZE];
            data[0..8].copy_from_slice(&TokenAdminRegistry::DISCRIMINATOR);
            data[8..].copy_from_slice(bytes.as_slice());

            let (mut lamports, key) = (0u64, Pubkey::new_unique());
            let account_info = to_acc_info(&key, &mut lamports, &mut data);

            let result = load_v1_token_admin_registry(&account_info);
            assert!(result.is_err(), "Version {} should not be valid", v);
        }
    }

    #[test]
    fn test_invalid_owner() {
        let source = get_expected();
        let bytes = to_v1_bytes(&source);

        let mut data = vec![0u8; V1_TOKEN_ADMIN_REGISTRY_SIZE];
        data[0..8].copy_from_slice(&TokenAdminRegistry::DISCRIMINATOR);
        data[8..].copy_from_slice(bytes.as_slice());

        let (mut lamports, key) = (0u64, Pubkey::new_unique());
        let mut account_info = to_acc_info(&key, &mut lamports, &mut data);
        let invalid_owner = Pubkey::new_unique();
        account_info.owner = &invalid_owner; // change the owner to an invalid one

        let result = load_v1_token_admin_registry(&account_info);
        assert!(result.is_err());
    }
}
