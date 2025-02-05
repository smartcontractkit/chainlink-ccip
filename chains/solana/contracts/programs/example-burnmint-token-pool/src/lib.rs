use anchor_lang::prelude::*;

declare_id!("TokenPooL11111111111111111111111111BurnMint");

mod context;
use crate::context::*;

mod rate_limiter;
use crate::rate_limiter::*;

mod common;
use crate::common::*;

#[program]
pub mod example_burnmint_token_pool {
    use anchor_lang::solana_program::program::invoke_signed;
    use anchor_spl::token_2022::spl_token_2022::{
        self,
        instruction::{burn, mint_to},
    };

    use super::*;

    pub fn initialize<'info>(
        ctx: Context<InitializeTokenPool>,
        ramp_authority: Pubkey,
    ) -> Result<()> {
        ctx.accounts.state.config.init(
            &ctx.accounts.mint,
            ctx.program_id.key(),
            ctx.accounts.authority.key(),
            ramp_authority,
        )
    }

    pub fn transfer_ownership(ctx: Context<SetConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts.state.config.transfer_ownership(proposed_owner)
    }

    // shared func signature with other programs
    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts.state.config.accept_ownership()
    }

    // set_ramp_authority changes the expected signer for mint/release + burn/lock method calls
    // this is used to update the router address
    pub fn set_ramp_authority(ctx: Context<SetConfig>, new_authority: Pubkey) -> Result<()> {
        ctx.accounts.state.config.set_ramp_authority(new_authority)
    }

    // initialize remote config (with no remote pools as it must be zero sized)
    pub fn init_chain_remote_config(
        ctx: Context<InitializeChainConfig>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        ctx.accounts.chain_config.set(remote_chain_selector, cfg)
    }

    // edit remote config
    pub fn edit_chain_remote_config(
        ctx: Context<EditChainConfigDynamicSize>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        ctx.accounts.chain_config.set(remote_chain_selector, cfg)
    }

    // Add remote pool addresses
    pub fn append_remote_pool_addresses(
        ctx: Context<AppendRemotePoolAddresses>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        addresses: Vec<RemoteAddress>,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .append_remote_pool_addresses(remote_chain_selector, addresses)
    }

    // set rate limit
    pub fn set_chain_rate_limit(
        ctx: Context<SetChainRateLimit>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        inbound: RateLimitConfig,
        outbound: RateLimitConfig,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .set_chain_rate_limit(remote_chain_selector, inbound, outbound)
    }

    // delete chain config
    pub fn delete_chain_config(
        _ctx: Context<DeleteChainConfig>,
        remote_chain_selector: u64,
        _mint: Pubkey,
    ) -> Result<()> {
        emit!(RemoteChainRemoved {
            chain_selector: remote_chain_selector,
        });
        Ok(())
    }

    pub fn configure_allow_list(
        ctx: Context<AddToAllowList>,
        add: Vec<Pubkey>,
        enabled: bool,
    ) -> Result<()> {
        ctx.accounts
            .state
            .config
            .update_allow_list(Some(enabled), add, vec![])
    }

    pub fn remove_from_allow_list(ctx: Context<SetConfig>, remove: Vec<Pubkey>) -> Result<()> {
        ctx.accounts
            .state
            .config
            .update_allow_list(None, vec![], remove)
    }

    pub fn release_or_mint_tokens<'info>(
        ctx: Context<TokenOfframp>,
        release_or_mint: ReleaseOrMintInV1,
    ) -> Result<ReleaseOrMintOutV1> {
        let parsed_amount = to_svm_token_amount(
            release_or_mint.amount,
            ctx.accounts.chain_config.remote.decimals,
            ctx.accounts.state.config.decimals,
        )?;

        let ChainConfig {
            remote,
            inbound_rate_limit,
            ..
        } = &mut *ctx.accounts.chain_config;

        validate_release_or_mint(
            &release_or_mint,
            parsed_amount,
            ctx.accounts.state.config.mint,
            &remote.pool_addresses,
            inbound_rate_limit,
        )?;

        // mint to receiver
        // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.mint_to.html
        let mut ix = mint_to(
            &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
            &ctx.accounts.mint.key(),
            &ctx.accounts.receiver_token_account.key(),
            &ctx.accounts.pool_signer.key(),
            &[],
            parsed_amount,
        )?;
        ix.program_id = ctx.accounts.token_program.key(); // set to user specified program

        let seeds = &[
            POOL_SIGNER_SEED,
            &ctx.accounts.mint.key().to_bytes(),
            &[ctx.bumps.pool_signer],
        ];
        invoke_signed(
            &ix,
            &[
                ctx.accounts.receiver_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.pool_signer.to_account_info(),
            ],
            &[&seeds[..]],
        )?;

        emit!(Minted {
            sender: ctx.accounts.pool_signer.key(),
            recipient: release_or_mint.receiver,
            amount: parsed_amount,
        });

        Ok(ReleaseOrMintOutV1 {
            destination_amount: parsed_amount,
        })
    }

    pub fn lock_or_burn_tokens<'info>(
        ctx: Context<TokenOnramp>,
        lock_or_burn: LockOrBurnInV1,
    ) -> Result<LockOrBurnOutV1> {
        validate_lock_or_burn(
            &lock_or_burn,
            ctx.accounts.state.config.mint,
            &mut ctx.accounts.chain_config.outbound_rate_limit,
            &ctx.accounts.state.config.allow_list,
        )?;

        // receiver -> token pool (occurs outside pool)
        // burn tokens held in pool
        // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.burn.html
        let mut ix = burn(
            &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
            &ctx.accounts.pool_token_account.key(),
            &ctx.accounts.mint.key(),
            &ctx.accounts.pool_signer.key(),
            &[],
            lock_or_burn.amount,
        )?;
        ix.program_id = ctx.accounts.token_program.key(); // set to user specified program

        let seeds = &[
            POOL_SIGNER_SEED,
            &ctx.accounts.mint.key().to_bytes(),
            &[ctx.bumps.pool_signer],
        ];
        invoke_signed(
            &ix,
            &[
                ctx.accounts.pool_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.pool_signer.to_account_info(),
            ],
            &[&seeds[..]],
        )?;

        emit!(Burned {
            sender: ctx.accounts.authority.key(),
            amount: lock_or_burn.amount,
        });

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.remote.token_address.clone(),
            dest_pool_data: RemoteAddress::ZERO,
        })
    }
}

#[account]
#[derive(InitSpace)]
pub struct State {
    pub version: u8,
    pub config: Config,
}
