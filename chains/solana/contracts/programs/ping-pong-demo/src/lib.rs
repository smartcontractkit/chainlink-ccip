use anchor_lang::prelude::*;
use program::PingPongDemo;

declare_id!("PPbZmYFf5SPAM9Jhm9mNmYoCwT7icPYVKAfJoMCQovU");

const ANCHOR_DISCRIMINATOR: usize = 8; // nr of bytes for the anchor discriminator.

use crate::context::*;

#[program]
pub mod ping_pong_demo {
    use anchor_spl::{
        token,
        token_2022::{spl_token_2022::instruction::approve_checked, Approve},
        token_interface,
    };
    use solana_program::program::invoke_signed;

    use super::*;

    pub fn initialize<'info>(
        ctx: Context<Initialize>,
        router: Pubkey,
        counterpart_chain_selector: u64,
        counterpart_address: [u8; 64],
        is_paused: bool,
        default_gas_limit: u64,
        out_of_order_execution: Pubkey,
    ) -> Result<()> {
        ctx.accounts.config.set_inner(state::Config {
            router,
            counterpart_chain_selector,
            counterpart_address,
            is_paused,
            default_gas_limit,
            out_of_order_execution,

            fee_token_mint: ctx.accounts.fee_token_mint.key(),

            // Set owner to the upgrade authority. The context constraints enforce
            // that the signer of the transaction (i.e. the authority) is the upgrade authority
            owner: ctx.accounts.authority.key(),
        });

        ctx.accounts.name_version.set_inner(state::NameVersion {
            name: "ping-pong-demo".to_string(),
            version: "1.5.0".to_string(), // TODO confirm if we keep EVM version here
        });

        // TODO approve the router to transfer max of the fee token
        {
            let approve_ctx: CpiContext<'_, '_, '_, 'info, Approve<'info>> = CpiContext::new(
                ctx.accounts.fee_token_program.to_account_info(),
                token_interface::Approve {
                    to: todo!(),
                    delegate: todo!(),
                    authority: todo!(),
                },
            );
            token_interface::approve(approve_ctx, 1000)?;
            // let mut ix = approve_checked(
            //     &ctx.accounts.fee_token_program.key(),
            //     &ctx.accounts.fee_token_ata.key(),
            //     &ctx.accounts.fee_token_mint.key(),
            //     &ctx.accounts.router_fee_billing_signer.key(),
            //     &ctx.accounts.ccip_send_signer.key(),
            //     &[&ctx.accounts.ccip_send_signer.key()],
            //     u64::MAX,
            //     ctx.accounts.fee_token_mint.decimals,
            // )?;
            // invoke_signed(
            //     &ix,
            //     &[
            //         ctx.accounts.fee_token_ata.to_account_info(),
            //         // self_ata.clone(),
            //         ctx.accounts.fee_token_mint.to_account_info(),
            //         // token_mint.clone(),
            //         ctx.accounts.,
            //         to_signer.clone(),
            //         self_signer.clone(),
            //     ],
            //     &[seeds],
            // )?;
        }

        helpers::approve_router_debits()
    }

    ////////////////////////////
    // Only callable by owner //
    ////////////////////////////

    pub fn set_counterpart(
        ctx: Context<UpdateConfig>,
        counterpart_chain_selector: u64,
        counterpart_address: [u8; 64],
    ) -> Result<()> {
        ctx.accounts.config.counterpart_chain_selector = counterpart_chain_selector;
        ctx.accounts.config.counterpart_address = counterpart_address;
        Ok(())
    }

    pub fn set_paused(ctx: Context<UpdateConfig>, pause: bool) -> Result<()> {
        ctx.accounts.config.is_paused = pause;
        Ok(())
    }

    pub fn set_out_of_order_execution(
        ctx: Context<UpdateConfig>,
        out_of_order_execution: Pubkey,
    ) -> Result<()> {
        ctx.accounts.config.out_of_order_execution = out_of_order_execution;
        Ok(())
    }

    pub fn start_ping_pong(ctx: Context<UpdateConfig>) -> Result<()> {
        ctx.accounts.config.is_paused = false;
        helpers::respond(1)
    }

    ///////////////////////////
    // Only callable by CCIP //
    ///////////////////////////
}

mod helpers {
    use super::*;

    // TODO EVM uses u256
    pub fn respond(ping_pong_count: u128) -> Result<()> {
        // TODO implement
        Ok(())
    }

    pub fn approve_router_debits() -> Result<()> {
        // TODO implement
        Ok(())
    }
}

pub mod context {
    use anchor_spl::associated_token::AssociatedToken;
    use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

    use super::*;

    use crate::state::*;

    mod seeds {
        pub const CONFIG: &[u8] = b"config";
        pub const NAME_VERSION: &[u8] = b"name_version";
        pub const CCIP_SEND_SIGNER: &[u8] = b"ccip_send_signer";
    }

    mod ccip_seeds {
        pub const FEE_BILLING_SIGNER: &[u8] = b"fee_billing_signer";
    }

    #[derive(Accounts)]
    #[instruction(router: Pubkey)]
    pub struct Initialize<'info> {
        #[account(
            init,
            seeds = [seeds::CONFIG],
            bump,
            space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
            payer = authority,
        )]
        pub config: Account<'info, Config>,

        #[account(
            init,
            seeds = [seeds::NAME_VERSION],
            bump,
            space = ANCHOR_DISCRIMINATOR + NameVersion::INIT_SPACE,
            payer = authority,
        )]
        pub name_version: Account<'info, NameVersion>, // TODO define if we keep this or remove it

        #[account(
            seeds = [ccip_seeds::FEE_BILLING_SIGNER],
            bump,
            seeds::program = router,
        )]
        pub router_fee_billing_signer: UncheckedAccount<'info>,

        // #[account(
        //     mut,
        //     associated_token::authority = router_fee_billing_signer.key(),
        //     associated_token::mint = fee_token_mint.key(),
        //     associated_token::token_program = fee_token_program.key(),
        // )]
        // pub ccip_fee_token_receiver: InterfaceAccount<'info, TokenAccount>,
        pub fee_token_program: Interface<'info, TokenInterface>,

        pub fee_token_mint: InterfaceAccount<'info, Mint>,

        #[account(
            init,
            payer = authority,
            associated_token::mint = fee_token_mint,
            associated_token::authority = ccip_send_signer,
            associated_token::token_program = fee_token_program,
        )]
        pub fee_token_ata: InterfaceAccount<'info, TokenAccount>,

        #[account(
            seeds = [seeds::CCIP_SEND_SIGNER],
            bump,
        )]
        pub ccip_send_signer: UncheckedAccount<'info>,

        #[account(mut)]
        pub authority: Signer<'info>,

        pub associated_token_program: Program<'info, AssociatedToken>,

        pub system_program: Program<'info, System>,

        #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
        pub program: Program<'info, PingPongDemo>,

        // Initialization only allowed by program upgrade authority
        #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ PingPongDemoError::Unauthorized)]
        pub program_data: Account<'info, ProgramData>,
    }

    #[derive(Accounts)]
    pub struct UpdateConfig<'info> {
        #[account(
            mut,
            seeds = [seeds::CONFIG],
            bump,
        )]
        pub config: Account<'info, Config>,

        #[account(
            mut,
            address = config.owner @ PingPongDemoError::Unauthorized,
        )]
        pub authority: Signer<'info>,
    }
}

pub mod state {
    use super::*;

    #[account]
    #[derive(InitSpace)]
    pub struct Config {
        pub owner: Pubkey, // The owner of the contract.

        pub router: Pubkey, // The address of the CCIP router contract.

        pub counterpart_chain_selector: u64, // The chain ID of the counterpart ping pong contract.
        pub counterpart_address: [u8; 64], // The contract address of the counterpart ping pong contract.
        pub is_paused: bool,               // Pause ping-ponging.
        pub fee_token_mint: Pubkey,        // The fee token used to pay for CCIP transactions.

        // Extra args on ccip_send
        pub default_gas_limit: u64, // Default gas limit used for EVMExtraArgsV2 construction.
        pub out_of_order_execution: Pubkey, // Allowing out of order execution.
    }

    #[account]
    #[derive(InitSpace)]
    pub struct NameVersion {
        #[max_len(32)]
        pub name: String,
        #[max_len(10)]
        pub version: String,
    }
}

#[error_code]
pub enum PingPongDemoError {
    #[msg("Unauthorized")]
    Unauthorized,
}
