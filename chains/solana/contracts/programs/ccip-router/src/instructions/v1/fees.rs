use anchor_lang::prelude::*;
use anchor_spl::token_interface;
use solana_program::{program::invoke_signed, system_instruction};

use crate::context::{seed, CcipSend};
use crate::messages::SVM2AnyMessage;
use crate::seed::FEE_BILLING_SIGNER;
use crate::{CcipRouterError, SVMTokenAmount};

use super::pools::TokenAccounts;

/// Invokes Fee Quoter to:
/// - validate the message,
/// - calculate the fee in both the billing token and LINK, and
/// - assert that it doesn't exceed a maximum
///
/// # Returns
///  GetFeeResult with fee quoter's response
pub(super) fn get_fee_cpi<'info>(
    ctx: &Context<'_, '_, 'info, 'info, CcipSend<'info>>,
    dest_chain_selector: u64,
    message: &SVM2AnyMessage,
    accounts_per_token: &[TokenAccounts<'info>],
) -> Result<fee_quoter::messages::GetFeeResult> {
    let get_fee_seeds = &[seed::FEE_BILLING_SIGNER, &[ctx.bumps.fee_billing_signer]];

    let get_fee_signer = &[&get_fee_seeds[..]];

    let cpi_program = ctx.accounts.fee_quoter.to_account_info();

    let cpi_accounts = fee_quoter::cpi::accounts::GetFee {
        config: ctx.accounts.fee_quoter_config.to_account_info(),
        dest_chain: ctx.accounts.fee_quoter_dest_chain.to_account_info(),
        billing_token_config: ctx
            .accounts
            .fee_quoter_billing_token_config
            .to_account_info(),
        link_token_config: ctx.accounts.fee_quoter_link_token_config.to_account_info(),
    };

    let billing_token_config_accs: &mut Vec<AccountInfo<'info>> = &mut accounts_per_token
        .iter()
        .map(|a| a.fee_token_config.to_account_info())
        .collect();
    let per_chain_per_token_config_accs: &mut Vec<AccountInfo<'info>> = &mut accounts_per_token
        .iter()
        .map(|a| a.token_billing_config.to_account_info())
        .collect();
    let remaining_accounts = billing_token_config_accs;
    remaining_accounts.append(per_chain_per_token_config_accs);

    let cpi_ctx = CpiContext::new_with_signer(cpi_program, cpi_accounts, get_fee_signer)
        .with_remaining_accounts(remaining_accounts.to_vec());

    let cpi_message = fee_quoter::messages::SVM2AnyMessage {
        receiver: message.receiver.clone(),
        data: message.data.clone(),
        token_amounts: message
            .token_amounts
            .iter()
            .map(|x| fee_quoter::messages::SVMTokenAmount {
                token: x.token,
                amount: x.amount,
            })
            .collect(),
        fee_token: message.fee_token,
        extra_args: message.extra_args.clone(),
    };

    let fee = fee_quoter::cpi::get_fee(cpi_ctx, dest_chain_selector, cpi_message)?.get();

    Ok(fee)
}

pub fn transfer_and_wrap_native_sol<'info>(
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
