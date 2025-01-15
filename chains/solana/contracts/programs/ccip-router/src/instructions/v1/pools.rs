use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token_2022::spl_token_2022::{self, instruction::transfer_checked, state::Mint};
use anchor_spl::token_interface::TokenAccount;
use solana_program::{
    address_lookup_table::state::AddressLookupTable, instruction::Instruction,
    program::invoke_signed,
};
use solana_program::{program::get_return_data, program_pack::Pack};

use crate::{
    CcipRouterError, ExternalExecutionConfig, TokenAdminRegistry, CCIP_TOKENPOOL_CONFIG,
    CCIP_TOKENPOOL_SIGNER, FEE_BILLING_TOKEN_CONFIG, TOKEN_ADMIN_REGISTRY_SEED,
    TOKEN_POOL_BILLING_SEED, TOKEN_POOL_CONFIG_SEED,
};

pub const CCIP_POOL_V1_RET_BYTES: usize = 8;
pub const CCIP_LOCK_OR_BURN_V1_RET_BYTES: u32 = 32;
const MIN_TOKEN_POOL_ACCOUNTS: usize = 12; // see TokenAccounts struct for all required accounts

pub fn calculate_token_pool_account_indices(
    i: usize,
    start_indices: &[u8],
    remaining_accounts_count: usize,
) -> Result<(usize, usize)> {
    // account set = [start...end)
    let start: usize = start_indices[i] as usize;
    let end: usize = if i == start_indices.len() - 1 {
        remaining_accounts_count
    } else {
        (start_indices[i + 1]) as usize
    };

    // validate indexes and account lengths
    // start < end: prevent overflow
    // end <= MAX, index not exceeded
    // end - start >= MIN_TOKEN_POOL_ACCOUNTS, ensure there are enough accounts
    require!(
        start < end && end <= remaining_accounts_count && end - start >= MIN_TOKEN_POOL_ACCOUNTS,
        CcipRouterError::InvalidInputsTokenIndices
    );

    Ok((start, end))
}

pub(super) struct TokenAccounts<'a> {
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

pub(super) fn validate_and_parse_token_accounts<'info>(
    user: Pubkey,
    chain_selector: u64,
    router: Pubkey,
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
            &[TOKEN_ADMIN_REGISTRY_SEED, mint.key().as_ref()],
            &router,
        );
        require!(
            token_admin_registry.key() == expected_token_admin_registry,
            CcipRouterError::InvalidInputsTokenAdminRegistryAccounts
        );

        // check pool program + pool config + pool signer
        let (expected_pool_config, _) = Pubkey::find_program_address(
            &[CCIP_TOKENPOOL_CONFIG, mint.key().as_ref()],
            &pool_program.key(),
        );
        let (expected_pool_signer, _) = Pubkey::find_program_address(
            &[CCIP_TOKENPOOL_SIGNER, mint.key().as_ref()],
            &pool_program.key(),
        );
        require!(
            *pool_config.owner == pool_program.key()
                && pool_config.key() == expected_pool_config
                && pool_signer.key() == expected_pool_signer,
            CcipRouterError::InvalidInputsPoolAccounts
        );

        let (expected_fee_token_config, _) =
            Pubkey::find_program_address(&[FEE_BILLING_TOKEN_CONFIG, mint.key.as_ref()], &router);
        require!(
            fee_token_config.key() == expected_fee_token_config,
            CcipRouterError::InvalidInputsConfigAccounts
        );

        // check token accounts
        require!(
            *mint.owner == token_program.key(),
            CcipRouterError::InvalidInputsTokenAccounts
        );
        require!(
            user_token_account.key()
                == get_associated_token_address_with_program_id(
                    &user,
                    &mint.key(),
                    &token_program.key()
                )
                && pool_token_account.key()
                    == get_associated_token_address_with_program_id(
                        &pool_signer.key(),
                        &mint.key(),
                        &token_program.key()
                    ),
            CcipRouterError::InvalidInputsTokenAccounts
        );

        // check per token per chain configs
        // billing: configured via CCIP router/fee quoter
        // chain config: configured via pool
        let (expected_billing_config, _) = Pubkey::find_program_address(
            &[
                TOKEN_POOL_BILLING_SEED,
                chain_selector.to_le_bytes().as_ref(),
                mint.key().as_ref(),
            ],
            &router,
        );
        let (expected_pool_chain_config, _) = Pubkey::find_program_address(
            &[
                TOKEN_POOL_CONFIG_SEED,
                chain_selector.to_le_bytes().as_ref(),
                mint.key().as_ref(),
            ],
            &pool_program.key(),
        );
        require!(
            token_billing_config.key() == expected_billing_config, // TODO: determine if this can be zero key for optional billing config?
            CcipRouterError::InvalidInputsConfigAccounts
        );
        require!(
            pool_chain_config.key() == expected_pool_chain_config,
            CcipRouterError::InvalidInputsConfigAccounts
        );

        // Check Lookup Table Address configured in TokenAdminRegistry
        let token_admin_registry_account: Account<TokenAdminRegistry> =
            Account::try_from(token_admin_registry)?;
        require!(
            token_admin_registry_account.lookup_table == lookup_table.key(),
            CcipRouterError::InvalidInputsLookupTableAccounts
        );

        // Check Lookup Table Entries
        let lookup_table_data = &mut &lookup_table.data.borrow()[..];
        let lookup_table_account: AddressLookupTable =
            AddressLookupTable::deserialize(lookup_table_data)
                .map_err(|_| CcipRouterError::InvalidInputsLookupTableAccounts)?;

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
                CcipRouterError::InvalidInputsLookupTableAccounts
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
                    CcipRouterError::InvalidInputsLookupTableAccountWritable
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

pub fn transfer_token<'info>(
    amount: u64,
    token_program: &AccountInfo,
    mint: &AccountInfo<'info>,
    from: &AccountInfo<'info>,
    to: &AccountInfo<'info>,
    signer: &Account<'info, ExternalExecutionConfig>,
    seeds: &[&[u8]],
) -> std::result::Result<(), ProgramError> {
    let mint_data = Mint::unpack(*mint.try_borrow_data()?)?;
    let mut transfer_ix = transfer_checked(
        &spl_token_2022::ID, // SDK requires spl-token or spl-token-2022 (cannot handle arbitrary token program)
        &from.key(),
        &mint.key(),
        &to.key(),
        &signer.key(),
        &[],
        amount,
        mint_data.decimals, // parse decimals from token account
    )?;
    transfer_ix.program_id = token_program.key(); // set token program in case custom
    invoke_signed(
        &transfer_ix,
        &[
            from.to_account_info(),
            mint.to_account_info(),
            to.to_account_info(),
            signer.to_account_info(),
        ],
        &[seeds],
    )
}

// ToTxData implements an interface that returns instruction data
pub trait ToTxData {
    fn to_tx_data(&self) -> Vec<u8>;
}

pub fn interact_with_pool<D: ToTxData>(
    pool_program: Pubkey,
    signer: Pubkey,
    acc_infos: Vec<AccountInfo>,
    data: D,
    seeds: &[&[u8]],
) -> std::result::Result<Vec<u8>, ProgramError> {
    let acc_metas: Vec<AccountMeta> = acc_infos
        .to_vec()
        .iter()
        .flat_map(|acc_info| {
            // Check signer from PDA External Execution config
            let is_signer = acc_info.key() == signer;
            acc_info.to_account_metas(Some(is_signer))
        })
        .collect();

    let ix = Instruction {
        program_id: pool_program,
        accounts: acc_metas,
        data: data.to_tx_data(),
    };

    // CPI call to pool is expected to return data using set_return_data
    // anchor does this automatically but is limited to max 256 bytes
    // https://github.com/coral-xyz/anchor/blob/0109f4a3cf4117570f627e3ae465b6247d504f69/lang/syn/src/codegen/program/handlers.rs#L113
    invoke_signed(&ix, &acc_infos, &[seeds])?;

    // parse return data
    // https://github.com/coral-xyz/anchor/blob/0109f4a3cf4117570f627e3ae465b6247d504f69/lang/syn/src/codegen/program/cpi.rs#L83
    let (_, data) = get_return_data().unwrap();
    Ok(data)
}

pub fn get_balance<'a>(token_account: &'a AccountInfo<'a>) -> Result<u64> {
    let mut acc: InterfaceAccount<TokenAccount> = InterfaceAccount::try_from(token_account)?;
    acc.reload()?; // reload state to ensure latest balance
    Ok(acc.amount)
}

// pack u64 into LE u256 for cross-chain amount
pub fn u64_to_le_u256(v: u64) -> [u8; 32] {
    let mut out: [u8; 32] = [0; 32];
    out[..8].copy_from_slice(v.to_le_bytes().as_slice());
    out
}

pub mod token_admin_registry_writable {
    use crate::TokenAdminRegistry;

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
            };

            assert_eq!(is(state, 0), false);
            assert_eq!(is(state, 128), false);
            assert_eq!(is(state, 255), false);

            assert_eq!(state.writable_indexes[0].count_ones(), 3);
            assert_eq!(state.writable_indexes[1].count_ones(), 3);

            assert_eq!(is(state, 7), true);
            assert_eq!(is(state, 2), true);
            assert_eq!(is(state, 4), true);
            assert_eq!(is(state, 128 + 8), true);
            assert_eq!(is(state, 128 + 56), true);
            assert_eq!(is(state, 128 + 100), true);
        }
    }
}
