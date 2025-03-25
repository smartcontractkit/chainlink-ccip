use anchor_lang::prelude::*;
use program::PingPongDemo;

use ccip_router::cpi;
use example_ccip_receiver::Any2SVMMessage;

declare_id!("PPbZmYFf5SPAM9Jhm9mNmYoCwT7icPYVKAfJoMCQovU");

const ANCHOR_DISCRIMINATOR: usize = 8; // nr of bytes for the anchor discriminator.

use crate::context::*;

#[program]
pub mod ping_pong_demo {
    use anchor_spl::token_interface;

    use crate::helpers::BigEndianU256;

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

        let seeds = &[seeds::CCIP_SEND_SIGNER, &[ctx.bumps.ccip_send_signer]];
        let signer_seeds = &[&seeds[..]];
        let approve_ctx = CpiContext::new_with_signer(
            ctx.accounts.fee_token_program.to_account_info(),
            token_interface::Approve {
                to: ctx.accounts.fee_token_ata.to_account_info(), // ATA from which transfers are approved
                delegate: ctx.accounts.router_fee_billing_signer.to_account_info(), // delegate who can transfer
                authority: ctx.accounts.ccip_send_signer.to_account_info(), // owner of the ATA
            },
            signer_seeds,
        );
        token_interface::approve(approve_ctx, u64::MAX)
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

    pub fn start_ping_pong<'info>(ctx: Context<Start>) -> Result<()> {
        ctx.accounts.config.is_paused = false;

        let accounts = cpi::accounts::CcipSend {
            config: ctx.accounts.ccip_router_config.to_account_info(),
            dest_chain_state: ctx.accounts.ccip_router_dest_chain_state.to_account_info(),
            nonce: ctx.accounts.ccip_router_nonce.to_account_info(),
            authority: ctx.accounts.ccip_send_signer.to_account_info(),
            system_program: ctx.accounts.system_program.to_account_info(),
            fee_token_mint: ctx.accounts.fee_token_mint.to_account_info(),
            fee_token_user_associated_account: ctx.accounts.fee_token_ata.to_account_info(),
            fee_token_receiver: ctx.accounts.ccip_router_fee_receiver.to_account_info(),
            fee_billing_signer: ctx
                .accounts
                .ccip_router_fee_billing_signer
                .to_account_info(),
            fee_quoter: ctx.accounts.fee_quoter.to_account_info(),
            fee_quoter_config: ctx.accounts.fee_quoter_config.to_account_info(),
            fee_quoter_dest_chain: ctx.accounts.fee_quoter_dest_chain.to_account_info(),
            fee_quoter_billing_token_config: ctx
                .accounts
                .fee_quoter_billing_token_config
                .to_account_info(),
            fee_quoter_link_token_config: ctx
                .accounts
                .fee_quoter_link_token_config
                .to_account_info(),
            rmn_remote: ctx.accounts.rmn_remote.to_account_info(),
            rmn_remote_curses: ctx.accounts.rmn_remote_curses.to_account_info(),
            rmn_remote_config: ctx.accounts.rmn_remote_config.to_account_info(),
            token_pools_signer: ctx.accounts.token_pools_signer.to_account_info(),
            fee_token_program: ctx.accounts.fee_token_program.to_account_info(),
        };

        helpers::respond(
            *ctx.accounts.config,
            ctx.accounts.ccip_router_program.to_account_info(),
            accounts,
            BigEndianU256::from(1), // start ping-ponging from the beginning
            ctx.bumps.ccip_send_signer,
        )
    }

    ///////////////////////////
    // Only callable by CCIP //
    ///////////////////////////
    pub fn ccip_receive(ctx: Context<PingPong>, message: Any2SVMMessage) -> Result<()> {
        if ctx.accounts.config.is_paused {
            return Ok(()); // stop ping-ponging as it is paused, but do not fail
        }

        let next_count = BigEndianU256::try_from(message.data)?.incr();

        let accounts = cpi::accounts::CcipSend {
            config: ctx.accounts.ccip_router_config.to_account_info(),
            dest_chain_state: ctx.accounts.ccip_router_dest_chain_state.to_account_info(),
            nonce: ctx.accounts.ccip_router_nonce.to_account_info(),
            authority: ctx.accounts.ccip_send_signer.to_account_info(),
            system_program: ctx.accounts.system_program.to_account_info(),
            fee_token_mint: ctx.accounts.fee_token_mint.to_account_info(),
            fee_token_user_associated_account: ctx.accounts.fee_token_ata.to_account_info(),
            fee_token_receiver: ctx.accounts.ccip_router_fee_receiver.to_account_info(),
            fee_billing_signer: ctx
                .accounts
                .ccip_router_fee_billing_signer
                .to_account_info(),
            fee_quoter: ctx.accounts.fee_quoter.to_account_info(),
            fee_quoter_config: ctx.accounts.fee_quoter_config.to_account_info(),
            fee_quoter_dest_chain: ctx.accounts.fee_quoter_dest_chain.to_account_info(),
            fee_quoter_billing_token_config: ctx
                .accounts
                .fee_quoter_billing_token_config
                .to_account_info(),
            fee_quoter_link_token_config: ctx
                .accounts
                .fee_quoter_link_token_config
                .to_account_info(),
            rmn_remote: ctx.accounts.rmn_remote.to_account_info(),
            rmn_remote_curses: ctx.accounts.rmn_remote_curses.to_account_info(),
            rmn_remote_config: ctx.accounts.rmn_remote_config.to_account_info(),
            token_pools_signer: ctx.accounts.token_pools_signer.to_account_info(),
            fee_token_program: ctx.accounts.fee_token_program.to_account_info(),
        };

        helpers::respond(
            *ctx.accounts.config,
            ctx.accounts.ccip_router_program.to_account_info(),
            accounts,
            next_count,
            ctx.bumps.ccip_send_signer,
        )
    }
}

mod helpers {
    use super::*;

    pub fn respond<'info>(
        config: state::Config,
        router_program: AccountInfo<'info>,
        accounts: cpi::accounts::CcipSend<'info>,
        ping_pong_count: BigEndianU256,
        signer_bump: u8,
    ) -> Result<()> {
        let data_bytes: [u8; 32] = ping_pong_count.into();
        let message = ccip_router::messages::SVM2AnyMessage {
            receiver: config.counterpart_address.to_vec(), // TODO trim address here
            data: data_bytes.to_vec(),
            token_amounts: vec![], // no token transfer
            fee_token: config.fee_token_mint,
            extra_args: vec![], // no extra args
        };

        let seeds = &[seeds::CCIP_SEND_SIGNER, &[signer_bump]];
        let signer_seeds = &[&seeds[..]];

        let cpi_ctx = CpiContext::new_with_signer(router_program, accounts, signer_seeds);
        cpi::ccip_send(cpi_ctx, config.counterpart_chain_selector, message, vec![])?;

        Ok(())
    }

    // The ping-pong count value is transmitted as an abi-encoded u256.
    // This means that it's an unsigned integer, 32 bytes long, big-endian encoded.
    pub struct BigEndianU256 {
        high: u128,
        low: u128,
    }

    impl BigEndianU256 {
        pub fn incr(&self) -> Self {
            if self.low == u128::MAX {
                if self.high == u128::MAX {
                    // Overflow, just wrap around.
                    // u256 is so big though that this should never happen.
                    return BigEndianU256 { high: 0, low: 0 };
                }
                BigEndianU256 {
                    high: self.high + 1,
                    low: 0,
                }
            } else {
                BigEndianU256 {
                    high: self.high,
                    low: self.low + 1,
                }
            }
        }
    }

    impl From<u128> for BigEndianU256 {
        fn from(value: u128) -> Self {
            BigEndianU256 {
                high: 0,
                low: value,
            }
        }
    }

    impl TryFrom<Vec<u8>> for BigEndianU256 {
        type Error = anchor_lang::error::Error;

        fn try_from(value: Vec<u8>) -> std::result::Result<Self, Self::Error> {
            require_eq!(value.len(), 32, PingPongDemoError::InvalidMessageDataLength);
            let bytes: [u8; 32] = value.try_into().unwrap();
            Ok(BigEndianU256::from(bytes))
        }
    }

    impl From<[u8; 32]> for BigEndianU256 {
        fn from(value: [u8; 32]) -> Self {
            BigEndianU256 {
                high: u128::from_be_bytes(value[..16].try_into().unwrap()),
                low: u128::from_be_bytes(value[16..].try_into().unwrap()),
            }
        }
    }

    impl From<BigEndianU256> for [u8; 32] {
        fn from(value: BigEndianU256) -> Self {
            let mut bytes = [0u8; 32];
            bytes[..16].copy_from_slice(&value.high.to_be_bytes());
            bytes[16..].copy_from_slice(&value.low.to_be_bytes());
            bytes
        }
    }
}

pub mod context {
    use anchor_spl::associated_token::AssociatedToken;
    use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

    use super::*;

    use crate::state::*;

    pub mod seeds {
        pub const CONFIG: &[u8] = b"config";
        pub const NAME_VERSION: &[u8] = b"name_version";
        pub const CCIP_SEND_SIGNER: &[u8] = b"ccip_send_signer";
    }

    mod ccip_seeds {
        pub const FEE_BILLING_SIGNER: &[u8] = b"fee_billing_signer";
        pub const EXTERNAL_EXECUTION_CONFIG: &[u8] = b"external_execution_config";
        pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";
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
        /// CHECK
        pub router_fee_billing_signer: UncheckedAccount<'info>,

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
        /// CHECK
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

    #[derive(Accounts)]
    pub struct Start<'info> {
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

        #[account(
            seeds = [seeds::CCIP_SEND_SIGNER],
            bump,
        )]
        /// CHECK
        pub ccip_send_signer: UncheckedAccount<'info>,

        pub fee_token_program: Interface<'info, TokenInterface>,

        pub fee_token_mint: InterfaceAccount<'info, Mint>,

        #[account(
            mut,
            associated_token::mint = fee_token_mint,
            associated_token::authority = ccip_send_signer,
            associated_token::token_program = fee_token_program,
        )]
        pub fee_token_ata: InterfaceAccount<'info, TokenAccount>,

        #[account(
            address = config.router,
        )]
        /// CHECK
        pub ccip_router_program: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_config: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_dest_chain_state: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_nonce: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_fee_receiver: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_fee_billing_signer: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_config: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_dest_chain: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_billing_token_config: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_link_token_config: UncheckedAccount<'info>,
        /// CHECK
        pub rmn_remote: UncheckedAccount<'info>,
        /// CHECK
        pub rmn_remote_curses: UncheckedAccount<'info>,
        /// CHECK
        pub rmn_remote_config: UncheckedAccount<'info>,
        /// CHECK
        pub token_pools_signer: UncheckedAccount<'info>,

        pub system_program: Program<'info, System>,
    }

    #[derive(Accounts)]
    #[instruction(message: Any2SVMMessage)]
    pub struct PingPong<'info> {
        // Offramp CPI signer PDA must be first
        // It is not mutable, and thus cannot be used as payer of init/realloc of other PDAs.
        #[account(
            seeds = [ccip_seeds::EXTERNAL_EXECUTION_CONFIG],
            bump,
            seeds::program = offramp_program.key(),
        )]
        pub authority: Signer<'info>,

        /// CHECK offramp program: exists only to derive the allowed offramp PDA
        /// and the authority PDA. Must be second.
        pub offramp_program: UncheckedAccount<'info>,

        // PDA to verify that calling offramp is valid. Must be third. It is left up to the implementer to decide
        // how they want to persist the router address to verify that this is the correct account (e.g. in the top level of
        // a global config/state account for the receiver, which is what this example does, or hard-coded,
        // or stored in any other way in any other account).
        /// CHECK PDA of the router program verifying the signer is an allowed offramp.
        /// If PDA does not exist, the router doesn't allow this offramp
        #[account(
        owner = config.router, // this guarantees that it was initialized
        seeds = [
            ccip_seeds::ALLOWED_OFFRAMP,
            message.source_chain_selector.to_le_bytes().as_ref(),
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = config.router,
    )]
        pub allowed_offramp: UncheckedAccount<'info>,

        #[account(
            seeds = [seeds::CONFIG],
            bump,
        )]
        pub config: Account<'info, Config>,

        #[account(
            seeds = [seeds::CCIP_SEND_SIGNER],
            bump,
        )]
        /// CHECK
        pub ccip_send_signer: UncheckedAccount<'info>,

        pub fee_token_program: Interface<'info, TokenInterface>,

        pub fee_token_mint: InterfaceAccount<'info, Mint>,

        #[account(
            mut,
            associated_token::mint = fee_token_mint,
            associated_token::authority = ccip_send_signer,
            associated_token::token_program = fee_token_program,
        )]
        pub fee_token_ata: InterfaceAccount<'info, TokenAccount>,

        #[account(
            address = config.router,
        )]
        /// CHECK
        pub ccip_router_program: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_config: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_dest_chain_state: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_nonce: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_fee_receiver: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_fee_billing_signer: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_config: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_dest_chain: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_billing_token_config: UncheckedAccount<'info>,
        /// CHECK
        pub fee_quoter_link_token_config: UncheckedAccount<'info>,
        /// CHECK
        pub rmn_remote: UncheckedAccount<'info>,
        /// CHECK
        pub rmn_remote_curses: UncheckedAccount<'info>,
        /// CHECK
        pub rmn_remote_config: UncheckedAccount<'info>,
        /// CHECK
        pub token_pools_signer: UncheckedAccount<'info>,

        pub system_program: Program<'info, System>,
    }
}

pub mod state {
    use super::*;

    #[account]
    #[derive(InitSpace, Copy)]
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
    #[msg("Invalid message data length")]
    InvalidMessageDataLength,
}
