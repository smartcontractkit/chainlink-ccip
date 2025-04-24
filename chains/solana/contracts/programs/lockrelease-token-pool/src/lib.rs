use anchor_lang::prelude::*;
use anchor_lang::solana_program::program::invoke_signed;
use anchor_spl::token_2022::spl_token_2022::{self, instruction::transfer_checked};
use base_token_pool::{common::*, rate_limiter::*};

mod context;
use crate::context::*;

declare_id!("8eqh8wppT9c5rw4ERqNCffvU6cNFJWff9WmkcYtmGiqC");

#[program]
pub mod lockrelease_token_pool {

    use super::*;

    pub fn initialize(
        ctx: Context<InitializeTokenPool>,
        router: Pubkey,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        ctx.accounts.state.set_inner(State {
            version: 1,
            config: BaseConfig::init(
                &ctx.accounts.mint,
                ctx.program_id.key(),
                ctx.accounts.authority.key(),
                router,
                rmn_remote,
            ),
        });
        Ok(())
    }

    /// Returns the program type (name) and version.
    /// Used by offchain code to easily determine which program & version is being interacted with.
    ///
    /// # Arguments
    /// * `ctx` - The context
    pub fn type_version(_ctx: Context<Empty>) -> Result<String> {
        let response = env!("CCIP_BUILD_TYPE_VERSION").to_string();
        msg!("{}", response);
        Ok(response)
    }

    pub fn transfer_ownership(ctx: Context<SetConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts.state.config.transfer_ownership(proposed_owner)
    }

    // shared func signature with other programs
    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts.state.config.accept_ownership()
    }

    // set_router changes the expected signers for mint/release + burn/lock method calls
    // this is used to update the router address
    pub fn set_router(ctx: Context<SetConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .config
            .set_router(new_router, ctx.program_id)
    }

    // permissionless method to set a pool's `state.version` value, only when
    // it was not set during the original initialization of the pool
    pub fn initialize_state_version(
        ctx: Context<InitializeStateVersion>,
        _mint: Pubkey,
    ) -> Result<()> {
        ctx.accounts.state.version = 1;
        Ok(())
    }

    // initialize remote config (with no remote pools as it must be zero sized)
    pub fn init_chain_remote_config(
        ctx: Context<InitializeChainConfig>,
        remote_chain_selector: u64,
        mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        require!(
            cfg.pool_addresses.is_empty(),
            CcipTokenPoolError::NonemptyPoolAddressesInit
        );

        ctx.accounts
            .chain_config
            .base
            .set(remote_chain_selector, mint, cfg)
    }

    // edit remote config
    pub fn edit_chain_remote_config(
        ctx: Context<EditChainConfigDynamicSize>,
        remote_chain_selector: u64,
        mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .base
            .set(remote_chain_selector, mint, cfg)
    }

    // Add remote pool addresses
    pub fn append_remote_pool_addresses(
        ctx: Context<AppendRemotePoolAddresses>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        addresses: Vec<RemoteAddress>,
    ) -> Result<()> {
        ctx.accounts.chain_config.base.append_remote_pool_addresses(
            remote_chain_selector,
            _mint,
            addresses,
        )
    }

    // set rate limit
    pub fn set_chain_rate_limit(
        ctx: Context<SetChainRateLimit>,
        remote_chain_selector: u64,
        mint: Pubkey,
        inbound: RateLimitConfig,
        outbound: RateLimitConfig,
    ) -> Result<()> {
        ctx.accounts.chain_config.base.set_chain_rate_limit(
            remote_chain_selector,
            mint,
            inbound,
            outbound,
        )
    }

    // delete chain config
    pub fn delete_chain_config(
        _ctx: Context<DeleteChainConfig>,
        remote_chain_selector: u64,
        mint: Pubkey,
    ) -> Result<()> {
        emit!(RemoteChainRemoved {
            chain_selector: remote_chain_selector,
            mint
        });
        Ok(())
    }

    pub fn configure_allow_list(
        ctx: Context<AddToAllowList>,
        add: Vec<Pubkey>,
        enabled: bool,
    ) -> Result<()> {
        ctx.accounts.state.config.list_enabled = enabled;
        let list = &mut ctx.accounts.state.config.allow_list;
        for key in add {
            require!(
                !list.contains(&key),
                CcipTokenPoolError::AllowlistKeyAlreadyExisted
            );
            list.push(key);
        }

        Ok(())
    }

    pub fn remove_from_allow_list(
        ctx: Context<RemoveFromAllowlist>,
        remove: Vec<Pubkey>,
    ) -> Result<()> {
        let list = &mut ctx.accounts.state.config.allow_list;
        for key in remove {
            require!(
                list.contains(&key),
                CcipTokenPoolError::AllowlistKeyDidNotExist
            );
            list.retain(|k| k != &key);
        }

        Ok(())
    }

    pub fn release_or_mint_tokens(
        ctx: Context<TokenOfframp>,
        release_or_mint: ReleaseOrMintInV1,
    ) -> Result<ReleaseOrMintOutV1> {
        let parsed_amount = to_svm_token_amount(
            release_or_mint.amount,
            ctx.accounts.chain_config.base.remote.decimals,
            ctx.accounts.state.config.decimals,
        )?;

        let BaseChain {
            remote,
            inbound_rate_limit,
            ..
        } = &mut ctx.accounts.chain_config.base;

        validate_release_or_mint(
            &release_or_mint,
            parsed_amount,
            ctx.accounts.state.config.mint,
            &remote.pool_addresses,
            inbound_rate_limit,
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
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

    pub fn lock_or_burn_tokens(
        ctx: Context<TokenOnramp>,
        lock_or_burn: LockOrBurnInV1,
    ) -> Result<LockOrBurnOutV1> {
        validate_lock_or_burn(
            &lock_or_burn,
            ctx.accounts.state.config.mint,
            &mut ctx.accounts.chain_config.base.outbound_rate_limit,
            ctx.accounts.state.config.list_enabled,
            &ctx.accounts.state.config.allow_list,
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
        )?;

        lock_tokens(ctx.accounts.authority.key(), lock_or_burn)?;

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.base.remote.token_address.clone(),
            dest_pool_data: {
                let mut abi_encoded_decimals = vec![0u8; 32];
                abi_encoded_decimals[31] = ctx.accounts.state.config.decimals;
                abi_encoded_decimals
            },
        })
    }

    pub fn set_rebalancer(ctx: Context<SetConfig>, rebalancer: Pubkey) -> Result<()> {
        ctx.accounts.state.config.set_rebalancer(rebalancer)
    }

    pub fn set_can_accept_liquidity(ctx: Context<SetConfig>, allow: bool) -> Result<()> {
        ctx.accounts.state.config.set_can_accept_liquidity(allow)
    }

    pub fn provide_liquidity(ctx: Context<TokenTransfer>, amount: u64) -> Result<()> {
        require!(
            ctx.accounts.state.config.can_accept_liquidity,
            CcipTokenPoolError::LiquidityNotAccepted
        );
        transfer_tokens(
            ctx.accounts.token_program.key(),
            ctx.accounts.pool_token_account.to_account_info(), // to
            ctx.accounts.remote_token_account.to_account_info(), // from
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.bumps.pool_signer,
            amount,
            ctx.accounts.mint.decimals,
        )
    }

    // withdraw liquidity can be used to transfer liquidity from one pool to another by setting the `rebalancer` to the calling pool
    pub fn withdraw_liquidity(ctx: Context<TokenTransfer>, amount: u64) -> Result<()> {
        transfer_tokens(
            ctx.accounts.token_program.key(),
            ctx.accounts.remote_token_account.to_account_info(), // to
            ctx.accounts.pool_token_account.to_account_info(),   // from
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.bumps.pool_signer,
            amount,
            ctx.accounts.mint.decimals,
        )
    }
}

// NOTE: accounts derivations must be native to program - will cause ownership check issues if imported
// wraps functionality from shared Config and Chain types
#[account]
#[derive(InitSpace)]
pub struct State {
    pub version: u8,
    pub config: BaseConfig,
}

#[account]
#[derive(InitSpace)]
pub struct ChainConfig {
    pub base: BaseChain,
}

pub fn lock_tokens(sender: Pubkey, lock_or_burn: LockOrBurnInV1) -> Result<()> {
    // receiver -> token pool (occurs outside pool)
    // hold tokens
    emit!(Locked {
        sender,
        amount: lock_or_burn.amount,
        mint: lock_or_burn.local_token,
    });

    Ok(())
}

#[allow(clippy::too_many_arguments)]
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
    transfer_tokens(
        token_program,
        receiver_token_account,
        pool_token_account,
        mint.clone(),
        pool_signer.clone(),
        pool_signer_bump,
        parsed_amount,
        decimals,
    )?;

    emit!(Released {
        sender: pool_signer.key(),
        recipient: release_or_mint.receiver,
        amount: parsed_amount,
        mint: mint.key(),
    });

    Ok(())
}

#[allow(clippy::too_many_arguments)]
pub fn transfer_tokens<'a>(
    token_program: Pubkey,
    to: AccountInfo<'a>,
    from: AccountInfo<'a>,
    mint: AccountInfo<'a>,
    pool_signer: AccountInfo<'a>,
    pool_signer_bump: u8,
    parsed_amount: u64,
    decimals: u8,
) -> Result<()> {
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.transfer.html
    let mut ix = transfer_checked(
        &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
        &from.key(),
        &mint.key(),
        &to.key(),
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
    invoke_signed(&ix, &[from, mint, to, pool_signer.clone()], &[&seeds[..]]).map_err(|e| e.into())
}
