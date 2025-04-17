use anchor_lang::prelude::*;
use anchor_spl::token_interface::TokenAccount;
use ccip_common::v1::MIN_TOKEN_POOL_ACCOUNTS;
use solana_program::program::get_return_data;
use solana_program::{instruction::Instruction, program::invoke_signed};

use super::messages::ReleaseOrMintInV1;

use crate::CcipOfframpError;

pub const CCIP_POOL_V1_RET_BYTES: usize = 8;

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
