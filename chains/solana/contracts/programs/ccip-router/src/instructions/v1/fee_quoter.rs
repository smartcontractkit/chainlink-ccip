use anchor_lang::prelude::*;
use anchor_spl::token_interface;
use solana_program::{program::invoke_signed, system_instruction};

use crate::{
    seed::FEE_BILLING_SIGNER, BillingTokenConfig, CcipRouterError, DestChain,
    PerChainPerTokenConfig, SVM2AnyMessage, SVMTokenAmount,
};

pub fn fee_for_msg(
    message: &SVM2AnyMessage,
    dest_chain: &DestChain,
    fee_token_config: &BillingTokenConfig,
    additional_token_configs: &[Option<BillingTokenConfig>],
    additional_token_configs_for_dest_chain: &[PerChainPerTokenConfig],
) -> Result<SVMTokenAmount> {
    Ok(SVMTokenAmount {
        token: Pubkey::default(),
        amount: 1,
    })
}

pub fn wrap_native_sol<'info>(
    token_program: &AccountInfo<'info>,
    from: &mut Signer<'info>,
    to: &mut InterfaceAccount<'info, token_interface::TokenAccount>,
    amount: u64,
    signer_bump: u8,
) -> Result<()> {
    require!(
        // guarantee that if caller is a PDA with data it won't get garbage-collected
        from.data_is_empty() || from.get_lamports() > amount,
        CcipRouterError::InsufficientLamports
    );

    invoke_signed(
        &system_instruction::transfer(&from.key(), &to.key(), amount),
        &[from.to_account_info(), to.to_account_info()],
        &[&[FEE_BILLING_SIGNER, &[signer_bump]]],
    )?;

    let seeds = &[FEE_BILLING_SIGNER, &[signer_bump]];
    let signer_seeds = &[&seeds[..]];
    let account = to.to_account_info();
    let sync: anchor_spl::token_2022::SyncNative = anchor_spl::token_2022::SyncNative { account };
    let cpi_ctx = CpiContext::new_with_signer(token_program.to_account_info(), sync, signer_seeds);

    token_interface::sync_native(cpi_ctx)
}

pub fn transfer_fee<'info>(
    fee: &SVMTokenAmount,
    token_program: AccountInfo<'info>,
    transfer: token_interface::TransferChecked<'info>,
    decimals: u8,
    signer_bump: u8,
) -> Result<()> {
    require!(
        fee.token == transfer.mint.key(),
        CcipRouterError::InvalidInputs
    ); // TODO use more specific error

    do_billing_transfer(token_program, transfer, fee.amount, decimals, signer_bump)
}

pub fn do_billing_transfer<'info>(
    token_program: AccountInfo<'info>,
    transfer: token_interface::TransferChecked<'info>,
    amount: u64,
    decimals: u8,
    signer_bump: u8,
) -> Result<()> {
    let seeds = &[FEE_BILLING_SIGNER, &[signer_bump]];
    let signer_seeds = &[&seeds[..]];
    let cpi_ctx = CpiContext::new_with_signer(token_program, transfer, signer_seeds);
    token_interface::transfer_checked(cpi_ctx, amount, decimals)
}
