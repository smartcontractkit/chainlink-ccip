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

    /// Initialize the global config account.
    /// Call this just once.
    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        ctx.accounts.global_config.set_inner(state::GlobalConfig {
            router,
            owner: ctx.accounts.authority.key(),
        });
        Ok(())
    }

    ////////////////////////////
    // Only callable by owner //
    ////////////////////////////

    /// Initialize the chain's config account.
    /// Call this once for each chain you want to ping-pong with.
    pub fn initialize_chain(
        ctx: Context<InitializeChain>,
        counterpart_chain_selector: u64,
        counterpart_address: Vec<u8>,
        is_paused: bool,
        extra_args: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts.config.set_inner(state::Config {
            counterpart_chain_selector,
            counterpart_address: counterpart_address.try_into()?,
            is_paused,
            extra_args,
            fee_token_mint: ctx.accounts.fee_token_mint.key(),
        });
        Ok(())
    }

    /// Initializes the ATA for the fee token and approve the Router for transferring from it.
    /// Call this once for each token you want to pay CCIP fees with.
    pub fn initialize_fee_token(
        ctx: Context<InitializeFeeToken>,
        _counterpart_chain_selector: u64,
    ) -> Result<()> {
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

    pub fn set_counterpart(
        ctx: Context<UpdateConfig>,
        counterpart_chain_selector: u64,
        counterpart_address: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts.config.counterpart_chain_selector = counterpart_chain_selector;
        ctx.accounts.config.counterpart_address = counterpart_address.try_into()?;
        Ok(())
    }

    pub fn set_paused(
        ctx: Context<UpdateConfig>,
        _counterpart_chain_selector: u64,
        pause: bool,
    ) -> Result<()> {
        ctx.accounts.config.is_paused = pause;
        Ok(())
    }

    pub fn set_extra_args(
        ctx: Context<UpdateReallocConfig>,
        _counterpart_chain_selector: u64,
        extra_args: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts.config.extra_args = extra_args;
        Ok(())
    }

    pub fn start_ping_pong(ctx: Context<Start>, counterpart_chain_selector: u64) -> Result<()> {
        msg!("Starting PingPong on chain {}", counterpart_chain_selector);

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
            ctx.accounts.global_config.clone().into_inner(),
            ctx.accounts.config.clone().into_inner(),
            accounts,
            BigEndianU256::from(1), // start ping-ponging from the beginning
            ctx.bumps.ccip_send_signer,
        )
    }

    ///////////////////////////
    // Only callable by CCIP //
    ///////////////////////////
    pub fn ccip_receive(ctx: Context<PingPong>, message: Any2SVMMessage) -> Result<()> {
        let received_counter = BigEndianU256::try_from(message.data)?;

        msg!(
            "Received a ping on chain {} with counter {}",
            message.source_chain_selector,
            received_counter
        );

        if ctx.accounts.config.is_paused {
            return Ok(()); // stop ping-ponging as it is paused, but do not fail
        }

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
            ctx.accounts.global_config.clone().into_inner(),
            ctx.accounts.config.clone().into_inner(),
            accounts,
            received_counter.incr(),
            ctx.bumps.ccip_send_signer,
        )
    }
}

mod helpers {
    use solana_program::{instruction::Instruction, program::invoke_signed};

    use super::*;

    pub const CCIP_SEND_DISCRIMINATOR: [u8; 8] = [108, 216, 134, 191, 249, 234, 33, 84];

    pub fn respond(
        global_config: state::GlobalConfig,
        config: state::Config,
        accounts: cpi::accounts::CcipSend,
        ping_pong_count: BigEndianU256,
        signer_bump: u8,
    ) -> Result<()> {
        let data_bytes: [u8; 32] = ping_pong_count.into();
        let message = ccip_router::messages::SVM2AnyMessage {
            receiver: config.counterpart_address.bytes().to_vec(),
            data: data_bytes.to_vec(),
            token_amounts: vec![], // no token transfer
            fee_token: config.fee_token_mint,
            extra_args: config.extra_args,
        };
        let token_indices: &[u8] = &[]; // empty token indexes vec

        let seeds = &[seeds::CCIP_SEND_SIGNER, &[signer_bump]];
        let signer_seeds = &[&seeds[..]];

        let acc_infos = accounts.to_account_infos().to_vec();

        let acc_metas: Vec<AccountMeta> = acc_infos
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == accounts.authority.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let mut data = CCIP_SEND_DISCRIMINATOR.to_vec();
        data.extend_from_slice(config.counterpart_chain_selector.to_le_bytes().as_ref());
        data.extend_from_slice(&message.try_to_vec().unwrap());
        data.extend_from_slice(&token_indices.try_to_vec().unwrap());

        let instruction = Instruction {
            program_id: global_config.router,
            accounts: acc_metas,
            data,
        };

        invoke_signed(&instruction, &acc_infos, signer_seeds)?;

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

    impl std::fmt::Display for BigEndianU256 {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            if self.high == 0 {
                write!(f, "{}", self.low)
            } else {
                write!(f, "BigEndianU256(high: {}, low: {})", self.high, self.low)
            }
        }
    }
}

pub mod context {
    use anchor_spl::associated_token::AssociatedToken;
    use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

    use super::*;

    use crate::state::*;

    pub mod seeds {
        pub const GLOBAL_CONFIG: &[u8] = b"global_config";
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
    pub struct Initialize<'info> {
        #[account(
            init,
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
            space = ANCHOR_DISCRIMINATOR + GlobalConfig::INIT_SPACE,
            payer = authority,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(mut)]
        pub authority: Signer<'info>,

        pub system_program: Program<'info, System>,

        #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
        pub program: Program<'info, PingPongDemo>,

        // Initialization only allowed by program upgrade authority
        #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ PingPongDemoError::Unauthorized)]
        pub program_data: Account<'info, ProgramData>,
    }

    #[derive(Accounts)]
    #[instruction(
        counterpart_chain_selector: u64,
        _counterpart_address: Vec<u8>,
        _is_paused: bool,
        extra_args: Vec<u8>,
    )]
    pub struct InitializeChain<'info> {
        #[account(
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(
            init,
            seeds = [seeds::CONFIG, counterpart_chain_selector.to_le_bytes().as_ref()],
            bump,
            space = Config::get_space(&extra_args),
            payer = authority,
        )]
        pub config: Account<'info, Config>,

        pub fee_token_mint: InterfaceAccount<'info, Mint>,

        #[account(
            mut,
            constraint = global_config.owner == authority.key() @ PingPongDemoError::Unauthorized,
        )]
        pub authority: Signer<'info>,

        pub system_program: Program<'info, System>,
    }

    #[derive(Accounts)]
    #[instruction(counterpart_chain_selector: u64)]
    pub struct InitializeFeeToken<'info> {
        #[account(
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(
            seeds = [seeds::CONFIG, counterpart_chain_selector.to_le_bytes().as_ref()],
            bump,
        )]
        pub config: Account<'info, Config>,

        #[account(
            seeds = [ccip_seeds::FEE_BILLING_SIGNER],
            bump,
            seeds::program = global_config.router,
        )]
        /// CHECK
        pub router_fee_billing_signer: UncheckedAccount<'info>,

        pub fee_token_program: Interface<'info, TokenInterface>,

        #[account(
            owner = fee_token_program.key(),
            address = config.fee_token_mint,
        )]
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

        #[account(mut, address = global_config.owner @ PingPongDemoError::Unauthorized)]
        pub authority: Signer<'info>,

        pub associated_token_program: Program<'info, AssociatedToken>,

        pub system_program: Program<'info, System>,
    }

    #[derive(Accounts)]
    #[instruction(counterpart_chain_selector: u64)]
    pub struct UpdateConfig<'info> {
        #[account(
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(
            mut,
            seeds = [seeds::CONFIG, counterpart_chain_selector.to_le_bytes().as_ref()],
            bump,
        )]
        pub config: Account<'info, Config>,

        #[account(
            mut,
            address = global_config.owner @ PingPongDemoError::Unauthorized,
        )]
        pub authority: Signer<'info>,
    }

    #[derive(Accounts)]
    #[instruction(counterpart_chain_selector: u64, extra_args: Vec<u8>)]
    pub struct UpdateReallocConfig<'info> {
        #[account(
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(
            mut,
            seeds = [seeds::CONFIG, counterpart_chain_selector.to_le_bytes().as_ref()],
            bump,
            realloc = Config::get_space(&extra_args),
            realloc::payer = authority,
            realloc::zero = false,
        )]
        pub config: Account<'info, Config>,

        #[account(
            mut,
            address = global_config.owner @ PingPongDemoError::Unauthorized,
        )]
        pub authority: Signer<'info>,

        pub system_program: Program<'info, System>,
    }

    #[derive(Accounts)]
    #[instruction(counterpart_chain_selector: u64)]
    pub struct Start<'info> {
        #[account(
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(
            mut,
            seeds = [seeds::CONFIG, counterpart_chain_selector.to_le_bytes().as_ref()],
            bump,
        )]
        pub config: Account<'info, Config>,

        #[account(
            mut,
            address = global_config.owner @ PingPongDemoError::Unauthorized,
        )]
        pub authority: Signer<'info>,

        #[account(
            mut,
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
            address = global_config.router,
        )]
        /// CHECK
        pub ccip_router_program: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_config: UncheckedAccount<'info>,
        /// CHECK
        #[account(mut)]
        pub ccip_router_dest_chain_state: UncheckedAccount<'info>,
        /// CHECK
        #[account(mut)]
        pub ccip_router_nonce: UncheckedAccount<'info>,
        /// CHECK
        #[account(mut)]
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
        #[account(mut)]
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
        owner = global_config.router, // this guarantees that it was initialized
        seeds = [
            ccip_seeds::ALLOWED_OFFRAMP,
            message.source_chain_selector.to_le_bytes().as_ref(),
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = global_config.router,
    )]
        pub allowed_offramp: UncheckedAccount<'info>,

        #[account(
            seeds = [seeds::GLOBAL_CONFIG],
            bump,
        )]
        pub global_config: Account<'info, GlobalConfig>,

        #[account(
            seeds = [seeds::CONFIG, message.source_chain_selector.to_le_bytes().as_ref()],
            bump,
        )]
        pub config: Account<'info, Config>,

        #[account(
            mut,
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
            address = global_config.router,
        )]
        /// CHECK
        pub ccip_router_program: UncheckedAccount<'info>,
        /// CHECK
        pub ccip_router_config: UncheckedAccount<'info>,
        /// CHECK
        #[account(mut)]
        pub ccip_router_dest_chain_state: UncheckedAccount<'info>,
        /// CHECK
        #[account(mut)]
        pub ccip_router_nonce: UncheckedAccount<'info>,
        /// CHECK
        #[account(mut)]
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
        #[account(mut)]
        pub token_pools_signer: UncheckedAccount<'info>,

        pub system_program: Program<'info, System>,
    }
}

pub mod state {
    use super::*;

    #[account]
    #[derive(InitSpace)]
    pub struct GlobalConfig {
        pub owner: Pubkey,  // The owner of the contract.
        pub router: Pubkey, // The address of the CCIP router contract.
    }

    #[account]
    #[derive(InitSpace)]
    pub struct Config {
        pub counterpart_chain_selector: u64, // The chain ID of the counterpart ping pong contract.
        pub counterpart_address: CounterpartAddress, // The contract address of the counterpart ping pong contract.
        pub is_paused: bool,                         // Pause ping-ponging.
        pub fee_token_mint: Pubkey, // The fee token used to pay for CCIP transactions.

        // Extra args on ccip_send
        #[max_len(0)]
        pub extra_args: Vec<u8>,
    }

    impl Config {
        pub fn get_space(extra_args: &[u8]) -> usize {
            ANCHOR_DISCRIMINATOR + Config::INIT_SPACE + extra_args.len()
        }
    }

    #[derive(AnchorDeserialize, AnchorSerialize, Clone, Copy, InitSpace)]
    pub struct CounterpartAddress {
        pub bytes: [u8; 64],
        pub len: u8,
    }

    impl CounterpartAddress {
        pub fn bytes(&self) -> &[u8] {
            &self.bytes[..self.len as usize]
        }
    }

    impl TryFrom<Vec<u8>> for CounterpartAddress {
        type Error = PingPongDemoError;

        fn try_from(value: Vec<u8>) -> std::result::Result<Self, Self::Error> {
            let mut address = CounterpartAddress {
                bytes: [0u8; 64],
                len: 0,
            };
            if value.is_empty() || value.len() > address.bytes.len() {
                return Err(PingPongDemoError::InvalidCounterpartAddress);
            }

            address.len = value.len() as u8;
            address.bytes[0..address.len as usize].copy_from_slice(&value);

            Ok(address)
        }
    }
}

#[error_code]
pub enum PingPongDemoError {
    #[msg("Unauthorized")]
    Unauthorized,
    #[msg("Invalid message data length")]
    InvalidMessageDataLength,
    #[msg("Invalid counterpart address")]
    InvalidCounterpartAddress,
}
