use anchor_lang::prelude::*;
use anchor_spl::token_2022::spl_token_2022::{self, instruction::transfer_checked, state::Mint};
use ccip_common::v1::MIN_TOKEN_POOL_ACCOUNTS;
use solana_program::{instruction::Instruction, program::invoke_signed};
use solana_program::{program::get_return_data, program_pack::Pack};

use crate::CcipRouterError;

pub const CCIP_LOCK_OR_BURN_V1_RET_BYTES: u32 = 32;

pub fn calculate_token_pool_account_indices(
    i: usize,
    start_indices: &[u8],
    remaining_accounts_count: usize,
) -> Result<(usize, usize)> {
    require!(
        i < start_indices.len(),
        CcipRouterError::InvalidInputsTokenIndices
    );

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

pub fn transfer_token<'info>(
    amount: u64,
    token_program: &AccountInfo,
    mint: &AccountInfo<'info>,
    from: &AccountInfo<'info>,
    to: &AccountInfo<'info>,
    signer: &AccountInfo<'info>,
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn token_pool_indices_underflow_overflow() {
        let empty_start_indices: [u8; 0] = [];
        let full_start_indices = [1u8, 2, 3];
        let arbitrary_account_count = 5;

        assert_eq!(
            calculate_token_pool_account_indices(0, &empty_start_indices, arbitrary_account_count),
            Err(CcipRouterError::InvalidInputsTokenIndices.into())
        );

        assert_eq!(
            calculate_token_pool_account_indices(
                full_start_indices.len(),
                &full_start_indices,
                arbitrary_account_count
            ),
            Err(CcipRouterError::InvalidInputsTokenIndices.into())
        );
    }
}
