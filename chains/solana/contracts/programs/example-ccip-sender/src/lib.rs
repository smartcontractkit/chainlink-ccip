#![allow(unexpected_cfgs)]

use anchor_lang::prelude::*;
use anchor_spl::token_2022::spl_token_2022::{self, instruction::transfer_checked};
use solana_program::{
    instruction::Instruction,
    program::{get_return_data, invoke, invoke_signed},
};

use ccip_router::messages::{GetFeeResult, SVM2AnyMessage, SVMTokenAmount};

declare_id!("4LfBQWYaU6zQZbDyYjX8pbY4qjzrhoumUFYZEZEqMNhJ");

#[cfg(target_os = "solana")]
#[global_allocator]
static ALLOC: smalloc::Smalloc<
    { solana_program::entrypoint::HEAP_START_ADDRESS as usize },
    { solana_program::entrypoint::HEAP_LENGTH as usize },
    16,
    1024,
> = smalloc::Smalloc::new();

pub mod context;
use context::*;

pub mod state;
use state::*;

pub mod tokens;
use tokens::*;

pub const CHAIN_CONFIG_SEED: &[u8] = b"remote_chain_config";
pub const CCIP_SENDER: &[u8] = b"ccip_sender";

pub const CCIP_SEND_DISCRIMINATOR: [u8; 8] = [108, 216, 134, 191, 249, 234, 33, 84]; // ccip_send
pub const CCIP_GET_FEE_DISCRIMINATOR: [u8; 8] = [115, 195, 235, 161, 25, 219, 60, 29]; // get_fee

/// This program an example of a CCIP Sender Program.
/// Used to test CCIP Router ccip_send.
#[program]
pub mod example_ccip_sender {
    use std::collections::BTreeSet;

    use super::*;

    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .init(ctx.accounts.authority.key(), router)
    }

    pub fn ccip_send<'info>(
        ctx: Context<'_, '_, 'info, 'info, CcipSend<'info>>,
        dest_chain_selector: u64,
        token_amounts: Vec<SVMTokenAmount>,
        data: Vec<u8>,
        fee_token: Pubkey,
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        let message = SVM2AnyMessage {
            receiver: ctx.accounts.chain_config.recipient.clone(),
            data,
            token_amounts: token_amounts.clone(),
            fee_token,
            extra_args: ctx.accounts.chain_config.extra_args_bytes.clone(),
        };
        let seeds = &[CCIP_SENDER, &[ctx.bumps.ccip_sender]];

        // Check that no token is transferred more than once
        let mints: BTreeSet<Pubkey> = token_amounts.iter().map(|ta| ta.token).collect();
        require!(
            token_amounts.len() == mints.len(),
            SenderError::TransferTokenDuplicated
        );

        // process token accounts
        let (token_accounts, ccip_token_indexes) = parse_and_validate_token_pool_accounts(
            &token_amounts,
            &token_indexes,
            ctx.remaining_accounts,
        )?;
        for (i, acc) in token_accounts.iter().enumerate() {
            // transfer tokens to this contract and approve the router to take possession during message processing
            if acc.mint.key() == fee_token {
                continue; // skip this one, as it will be handled together with the fee later
            }
            transfer_to_self_and_approve(
                acc.program,
                acc.mint,
                acc.from_ata,
                acc.self_ata,
                &ctx.accounts.ccip_sender.to_account_info(),
                &ctx.accounts.ccip_fee_billing_signer.to_account_info(),
                seeds,
                token_amounts[i].amount,
                acc.decimals,
            )?;
        }

        // CPI: get fee from router
        let fee: GetFeeResult = {
            let mut acc_infos: Vec<AccountInfo> = [
                ctx.accounts.ccip_config.to_account_info(),
                ctx.accounts.ccip_dest_chain_state.to_account_info(),
                ctx.accounts.ccip_fee_quoter.to_account_info(),
                ctx.accounts.ccip_fee_quoter_config.to_account_info(),
                ctx.accounts.ccip_fee_quoter_dest_chain.to_account_info(),
                ctx.accounts
                    .ccip_fee_quoter_billing_token_config
                    .to_account_info(),
                ctx.accounts
                    .ccip_fee_quoter_link_token_config
                    .to_account_info(),
            ]
            .to_vec();

            // add tokens to fee quoter request
            let billing_token_config_accs: &mut Vec<AccountInfo<'info>> = &mut token_accounts
                .iter()
                .map(|a| a.fee_token_config.to_account_info())
                .collect();
            let per_chain_per_token_config_accs: &mut Vec<AccountInfo<'info>> = &mut token_accounts
                .iter()
                .map(|a| a.token_billing_config.to_account_info())
                .collect();
            acc_infos.append(billing_token_config_accs);
            acc_infos.append(per_chain_per_token_config_accs);

            let acc_metas: Vec<AccountMeta> = acc_infos
                .to_vec()
                .iter()
                .flat_map(|acc_info| acc_info.to_account_metas(None))
                .collect();

            let instruction = Instruction {
                program_id: ctx.accounts.ccip_router.key(),
                accounts: acc_metas,
                data: builder::instruction(
                    &message,
                    CCIP_GET_FEE_DISCRIMINATOR,
                    dest_chain_selector,
                ),
            };

            invoke(&instruction, &acc_infos)?;
            let (_, data) = get_return_data().unwrap();
            GetFeeResult::try_from_slice(&data)?
        };

        // if fee token is not native, transfer the fee amounts from sender to the program and approve the router
        if fee_token != Pubkey::default() {
            // if paying fees with a token that is also being transferred, consider the amount to for both operations
            let amount = token_amounts
                .iter()
                .find(|ta| ta.token == fee_token)
                .map_or(0, |ta| ta.amount)
                .checked_add(fee.amount)
                .expect("The fee + transfer amount of token {} is too large");

            transfer_to_self_and_approve(
                &ctx.accounts.ccip_fee_token_program.to_account_info(),
                &ctx.accounts.ccip_fee_token_mint.to_account_info(),
                &ctx.accounts.authority_fee_token_ata.to_account_info(),
                &ctx.accounts.ccip_fee_token_user_ata.to_account_info(),
                &ctx.accounts.ccip_sender.to_account_info(),
                &ctx.accounts.ccip_fee_billing_signer.to_account_info(),
                seeds,
                amount,
                ctx.accounts.ccip_fee_token_mint.decimals,
            )?;
        }

        // minimum set of ccipSend accounts
        let mut acc_infos: Vec<AccountInfo> = [
            ctx.accounts.ccip_config.to_account_info(),
            ctx.accounts.ccip_dest_chain_state.to_account_info(),
            ctx.accounts.ccip_sender_nonce.to_account_info(),
            ctx.accounts.ccip_sender.to_account_info(),
            ctx.accounts.system_program.to_account_info(),
            ctx.accounts.ccip_fee_token_program.to_account_info(),
            ctx.accounts.ccip_fee_token_mint.to_account_info(),
            ctx.accounts.ccip_fee_token_user_ata.to_account_info(),
            ctx.accounts.ccip_fee_token_receiver.to_account_info(),
            ctx.accounts.ccip_fee_billing_signer.to_account_info(),
            ctx.accounts.ccip_fee_quoter.to_account_info(),
            ctx.accounts.ccip_fee_quoter_config.to_account_info(),
            ctx.accounts.ccip_fee_quoter_dest_chain.to_account_info(),
            ctx.accounts
                .ccip_fee_quoter_billing_token_config
                .to_account_info(),
            ctx.accounts
                .ccip_fee_quoter_link_token_config
                .to_account_info(),
            ctx.accounts.ccip_rmn_remote.to_account_info(),
            ctx.accounts.ccip_rmn_remote_curses.to_account_info(),
            ctx.accounts.ccip_rmn_remote_config.to_account_info(),
        ]
        .to_vec();

        // pass token accounts along with ccipSend instruction
        for tokens in token_accounts {
            acc_infos.append(&mut tokens.pool_accounts.to_vec());
        }

        let acc_metas: Vec<AccountMeta> = acc_infos
            .to_vec()
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == ctx.accounts.ccip_sender.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let instruction = Instruction {
            program_id: ctx.accounts.ccip_router.key(),
            accounts: acc_metas,
            data: builder::instruction_with_token_indexes(
                &message,
                CCIP_SEND_DISCRIMINATOR,
                dest_chain_selector,
                &ccip_token_indexes,
            ),
        };

        let seeds = &[CCIP_SENDER, &[ctx.bumps.ccip_sender]];
        let signer = &[&seeds[..]];

        invoke_signed(&instruction, &acc_infos, signer)?;
        let (_, data) = get_return_data().unwrap();

        emit!(MessageSent {
            message_id: data.try_into().unwrap()
        });

        Ok(())
    }

    pub fn update_router(ctx: Context<UpdateConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .update_router(ctx.accounts.authority.key(), new_router)
    }

    pub fn transfer_ownership(ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .transfer_ownership(ctx.accounts.authority.key(), proposed_owner)
    }

    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts
            .state
            .accept_ownership(ctx.accounts.authority.key())
    }

    pub fn withdraw_tokens(ctx: Context<WithdrawTokens>, amount: u64, decimals: u8) -> Result<()> {
        let mut ix = transfer_checked(
            &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
            &ctx.accounts.program_token_account.key(),
            &ctx.accounts.mint.key(),
            &ctx.accounts.to_token_account.key(),
            &ctx.accounts.ccip_sender.key(),
            &[],
            amount,
            decimals,
        )?;
        ix.program_id = ctx.accounts.token_program.key(); // set to user specified program

        let seeds = &[CCIP_SENDER, &[ctx.bumps.ccip_sender]];
        invoke_signed(
            &ix,
            &[
                ctx.accounts.program_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.to_token_account.to_account_info(),
                ctx.accounts.ccip_sender.to_account_info(),
            ],
            &[&seeds[..]],
        )?;
        Ok(())
    }

    // creates initial chain config
    // only can be called for allowing a chain
    pub fn init_chain_config(
        ctx: Context<InitChainConfig>,
        _chain_selector: u64,
        recipient: Vec<u8>,
        extra_args_bytes: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .set_config(recipient, extra_args_bytes)
    }

    // updates the chain config parameters
    // configs cannot be changed through init_chain_config due to realloc needs
    pub fn update_chain_config(
        ctx: Context<UpdateChainConfig>,
        _chain_selector: u64,
        recipient: Vec<u8>,
        extra_args_bytes: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .set_config(recipient, extra_args_bytes)
    }

    // disables remote chain by closing chain config account
    // can be re-enabled with init_chain_config
    pub fn remove_chain_config(
        _ctx: Context<RemoveChainConfig>,
        _chain_selector: u64,
    ) -> Result<()> {
        Ok(())
    }
}

#[error_code]
pub enum SenderError {
    #[msg("The same token is being transferred more than once")]
    TransferTokenDuplicated,
}
