use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token_interface::TokenAccount;
use solana_program::program::get_return_data;
use solana_program::{
    address_lookup_table::state::AddressLookupTable, instruction::Instruction,
    program::invoke_signed,
};

use super::messages::router_state;
use super::messages::ReleaseOrMintInV1;

use crate::{seed, CcipOfframpError};

pub const CCIP_POOL_V1_RET_BYTES: usize = 8;
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
        CcipOfframpError::InvalidInputsTokenIndices
    );

    Ok((start, end))
}

pub(super) struct TokenAccounts<'a> {
    // The fields prefixed with `_` are not used by the offramp directly,
    // but are present in the token lookup table anyway for other pool-related functionality.
    // So, the code here still validates that they are correct, even if those values are not
    // then read by this program.
    pub user_token_account: &'a AccountInfo<'a>,
    pub _token_billing_config: &'a AccountInfo<'a>,
    pub pool_chain_config: &'a AccountInfo<'a>,
    pub pool_program: &'a AccountInfo<'a>,
    pub pool_config: &'a AccountInfo<'a>,
    pub pool_token_account: &'a AccountInfo<'a>,
    pub pool_signer: &'a AccountInfo<'a>,
    pub token_program: &'a AccountInfo<'a>,
    pub mint: &'a AccountInfo<'a>,
    pub _fee_token_config: &'a AccountInfo<'a>,
    pub remaining_accounts: &'a [AccountInfo<'a>],
}

#[derive(Accounts)]
#[instruction(token_receiver: Pubkey, chain_selector: u64, router: Pubkey, fee_quoter: Pubkey)]
pub struct TokenPoolAccounts<'info> {
    /**
     * accounts based on user or chain
     */
    /// CHECK: User Token Account
    #[account(
        constraint = user_token_account.key() == get_associated_token_address_with_program_id(
            &token_receiver.key(),
            &mint.key(),
            &token_program.key()
        ) @ CcipOfframpError::InvalidInputsTokenAccounts
    )]
    pub user_token_account: AccountInfo<'info>,

    // TODO: determine if this can be zero key for optional billing config?
    /// CHECK: Per chain token billing config
    // billing: configured via CCIP fee quoter
    // chain config: configured via pool
    #[account(
        seeds = [
            fee_quoter::context::seed::PER_CHAIN_PER_TOKEN_CONFIG,
            chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref(),
        ],
        seeds::program = fee_quoter.key(),
        bump
    )]
    pub token_billing_config: AccountInfo<'info>,

    /// CHECK: Pool chain config
    #[account(
        seeds = [
            seed::TOKEN_POOL_CONFIG,
            chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref(),
        ],
        seeds::program = pool_program.key(),
        bump
    )]
    pub pool_chain_config: AccountInfo<'info>,

    /**
     * constant accounts for any pool interaction
     */
    /// CHECK: Lookup table
    pub lookup_table: AccountInfo<'info>,

    /// CHECK: Token admin registry
    #[account(
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        seeds::program = router.key(),
        bump,
        constraint = *token_admin_registry.to_account_info().owner == router.key() @ CcipOfframpError::InvalidInputsTokenAdminRegistryAccounts,
    )]
    pub token_admin_registry: AccountInfo<'info>,

    /// CHECK: Pool program
    pub pool_program: AccountInfo<'info>,

    // todo: PDA constraint violation will emit AccountConstraintViolation error instead of InvalidInputsPoolAccounts
    /// CHECK: Pool config
    #[account(
        seeds = [seed::CCIP_TOKENPOOL_CONFIG, mint.key().as_ref()],
        seeds::program = pool_program.key(),
        bump,
        owner = pool_program.key() @ CcipOfframpError::InvalidInputsPoolAccounts
    )]
    pub pool_config: AccountInfo<'info>,

    /// CHECK: Pool token account
    #[account(
        address = get_associated_token_address_with_program_id(
            &pool_signer.key(),
            &mint.key(),
            &token_program.key()
        ) @ CcipOfframpError::InvalidInputsTokenAccounts
    )]
    pub pool_token_account: AccountInfo<'info>,

    // todo: PDA constraint violation will emit AccountConstraintViolation error instead of InvalidInputsPoolAccounts
    /// CHECK: Pool signer
    #[account(
        seeds = [seed::CCIP_TOKENPOOL_SIGNER, mint.key().as_ref()],
        seeds::program = pool_program.key(),
        bump
    )]
    pub pool_signer: AccountInfo<'info>,

    /// CHECK: Token program
    pub token_program: AccountInfo<'info>,

    /// CHECK: Mint
    #[account(owner = token_program.key() @ CcipOfframpError::InvalidInputsTokenAccounts)]
    pub mint: AccountInfo<'info>,

    // todo: PDA constraint violation will emit AccountConstraintViolation error instead of InvalidInputsConfigAccounts
    /// CHECK: Fee token config
    #[account(
        seeds = [
            fee_quoter::context::seed::FEE_BILLING_TOKEN_CONFIG,
            mint.key().as_ref()
        ],
        seeds::program = fee_quoter.key(),
        bump
    )]
    pub fee_token_config: AccountInfo<'info>,
}

pub(super) fn validate_and_parse_token_accounts<'info>(
    token_receiver: Pubkey,
    chain_selector: u64,
    router: Pubkey,
    fee_quoter: Pubkey,
    accounts: &'info [AccountInfo<'info>],
) -> Result<TokenAccounts<'info>> {
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

    // collect remaining accounts
    let remaining_accounts = accounts_iter.as_slice();

    let program_id = crate::id();

    let mut input_accounts = accounts;
    let mut bumps = <TokenPoolAccounts as anchor_lang::Bumps>::Bumps::default();
    let mut reallocs = std::collections::BTreeSet::new();

    TokenPoolAccounts::try_accounts(
        &program_id,
        &mut input_accounts,
        &[
            token_receiver.as_ref(),
            &chain_selector.to_le_bytes(),
            router.as_ref(),
            fee_quoter.as_ref(),
        ]
        .concat(),
        &mut bumps,
        &mut reallocs,
    )?;

    // Additional validations
    {
        // Check Lookup Table Address configured in TokenAdminRegistry
        // For that, deserialize the TokenAdminRegistry first. It has already been checked that it is(can't be deserialized in the account context)
        // the right PDA address and that its owner is the Router program, so it's safe to deserialize the data directly
        let token_admin_registry_data = &mut &token_admin_registry.data.borrow()[..];
        let token_admin_registry_account =
            router_state::TokenAdminRegistry::try_deserialize(token_admin_registry_data)
                .map_err(|_| CcipOfframpError::InvalidInputsTokenAdminRegistryAccounts)?;
        require_keys_eq!(
            token_admin_registry_account.lookup_table,
            lookup_table.key(),
            CcipOfframpError::InvalidInputsLookupTableAccounts
        );

        // Check Lookup Table Entries
        let lookup_table_data = &mut &lookup_table.data.borrow()[..];
        let lookup_table_account = AddressLookupTable::deserialize(lookup_table_data)
            .map_err(|_| CcipOfframpError::InvalidInputsLookupTableAccounts)?;

        // reconstruct + validate expected values in token pool lookup table
        // base set of constant accounts (9)
        // + additional constant accounts (remaining_accounts) that are not required but may be used for additional token pool functionality (like CPI)
        let required_entries = [
            lookup_table.clone(),
            token_admin_registry.clone(),
            pool_program.clone(),
            pool_config.clone(),
            pool_token_account.clone(),
            pool_signer.clone(),
            token_program.clone(),
            mint.clone(),
            fee_token_config.clone(),
        ];
        {
            // validate pool addresses
            let mut expected_keys: Vec<Pubkey> =
                required_entries.iter().map(|acc| acc.key()).collect();
            let mut remaining_keys: Vec<Pubkey> =
                remaining_accounts.iter().map(|acc| acc.key()).collect();
            expected_keys.append(&mut remaining_keys);
            require!(
                lookup_table_account.addresses.as_ref() == expected_keys,
                CcipOfframpError::InvalidInputsLookupTableAccounts
            );
        }
        {
            // validate pool address writable
            // token admin registry contains an array (binary) of indexes that are writable
            // check that the writability of the passed accounts match the writable configuration (using indexes)
            let mut expected_is_writable: Vec<bool> =
                required_entries.iter().map(|acc| acc.is_writable).collect();
            let mut remaining_is_writable: Vec<bool> = remaining_accounts
                .iter()
                .map(|acc| acc.is_writable)
                .collect();
            expected_is_writable.append(&mut remaining_is_writable);
            for (i, is_writable) in expected_is_writable.iter().enumerate() {
                require_eq!(
                    token_admin_registry_writable::is(&token_admin_registry_account, i as u8),
                    *is_writable,
                    CcipOfframpError::InvalidInputsLookupTableAccountWritable
                );
            }
        }
    }

    Ok(TokenAccounts {
        user_token_account,
        _token_billing_config: token_billing_config,
        pool_chain_config,
        pool_program,
        pool_config,
        pool_token_account,
        pool_signer,
        token_program,
        mint,
        _fee_token_config: fee_token_config,
        remaining_accounts,
    })
}

pub(super) fn interact_with_pool(
    pool_program: Pubkey,
    signer: Pubkey,
    acc_infos: Vec<AccountInfo>,
    data: ReleaseOrMintInV1,
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

mod token_admin_registry_writable {
    use super::super::pools::router_state::TokenAdminRegistry;

    pub(super) fn is(tar: &TokenAdminRegistry, index: u8) -> bool {
        match index < 128 {
            true => tar.writable_indexes[0] & 1 << (127 - index) != 0,
            false => tar.writable_indexes[1] & 1 << (255 - index) != 0,
        }
    }

    #[cfg(test)]
    mod tests {
        use super::*;
        use solana_program::pubkey::Pubkey;

        // set writable inserts bits from left to right
        // index 0 is left-most bit
        fn set(tar: &mut TokenAdminRegistry, index: u8) {
            match index < 128 {
                true => {
                    tar.writable_indexes[0] |= 1 << (127 - index);
                }
                false => {
                    tar.writable_indexes[1] |= 1 << (255 - index);
                }
            }
        }

        fn reset(tar: &mut TokenAdminRegistry) {
            tar.writable_indexes = [0, 0];
        }

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
