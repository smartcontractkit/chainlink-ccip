use anchor_lang::prelude::*;
use anchor_lang::solana_program::program::invoke_signed;
use anchor_spl::token_2022::spl_token_2022::{self, instruction::transfer_checked};
use example_burnmint_token_pool::{common::*, rate_limiter::*};

mod context;
use crate::context::*;

declare_id!("TokenPooL11111111111111111111111LockReLease");

#[program]
pub mod example_lockrelease_token_pool {

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

        release_tokens(
            ctx.accounts.token_program.key(),
            ctx.accounts.receiver_token_account.to_account_info(),
            ctx.accounts.pool_token_account.to_account_info(),
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.bumps.pool_signer,
            release_or_mint,
            parsed_amount,
            ctx.accounts.mint.decimals,
        )?;

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

        lock_tokens(ctx.accounts.authority.key(), lock_or_burn)?;

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.remote.token_address.clone(),
            dest_pool_data: RemoteAddress::ZERO,
        })
    }
}

pub fn lock_tokens<'a>(sender: Pubkey, lock_or_burn: LockOrBurnInV1) -> Result<()> {
    // receiver -> token pool (occurs outside pool)
    // hold tokens
    emit!(Locked {
        sender,
        amount: lock_or_burn.amount,
    });

    Ok(())
}

pub fn release_tokens<'a>(
    token_program: Pubkey,
    receiver_token_account: AccountInfo<'a>,
    pool_token_account: AccountInfo<'a>,

    mint: AccountInfo<'a>,
    pool_signer: AccountInfo<'a>,
    pool_signer_bump: u8,
    release_or_mint: ReleaseOrMintInV1,
    parsed_amount: u64,
    decimals: u8,
) -> Result<()> {
    // transfer from pool -> receiver
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.transfer.html
    let mut ix = transfer_checked(
        &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
        &pool_token_account.key(),
        &mint.key(),
        &receiver_token_account.key(),
        &pool_signer.key(),
        &[],
        parsed_amount,
        decimals,
    )?;
    ix.program_id = token_program; // set to user specified program

    let seeds = &[
        POOL_SIGNER_SEED,
        &mint.key().to_bytes(),
        &[pool_signer_bump],
    ];
    invoke_signed(
        &ix,
        &[
            pool_token_account,
            mint,
            receiver_token_account,
            pool_signer.clone(),
        ],
        &[&seeds[..]],
    )?;

    emit!(Released {
        sender: pool_signer.key(),
        recipient: release_or_mint.receiver,
        amount: parsed_amount,
    });

    Ok(())
}
