use anchor_lang::prelude::*;
use base_token_pool::{common::*, rate_limiter::RateLimitConfig};

declare_id!("JuCcZ4smxAYv9QHJ36jshA7pA3FuQ3vQeWLUeAtZduJ");

mod context;
use crate::context::*;

#[program]
pub mod test_token_pool {
    use anchor_lang::solana_program::{instruction::Instruction, program::invoke_signed};
    use burnmint_token_pool::{burn_tokens, mint_tokens};
    use lockrelease_token_pool::{lock_tokens, release_tokens};

    use super::*;

    pub fn initialize(
        ctx: Context<InitializeTokenPool>,
        pool_type: PoolType,
        router: Pubkey,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        ctx.accounts.state.set_inner(State {
            pool_type,
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

    pub fn release_or_mint_tokens<'info>(
        ctx: Context<'_, '_, '_, 'info, TokenOfframp<'info>>,
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

        match ctx.accounts.state.pool_type {
            PoolType::LockAndRelease => release_tokens(
                ctx.accounts.token_program.key(),
                ctx.accounts.receiver_token_account.to_account_info(),
                ctx.accounts.pool_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.pool_signer.to_account_info(),
                ctx.bumps.pool_signer,
                release_or_mint,
                parsed_amount,
                ctx.accounts.mint.decimals,
            )?,
            PoolType::BurnAndMint => mint_tokens(
                ctx.accounts.token_program.key(),
                ctx.accounts.receiver_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.pool_signer.to_account_info(),
                ctx.bumps.pool_signer,
                release_or_mint,
                parsed_amount,
            )?,
            PoolType::Wrapped => {
                // The External Execution Config Account is used to sign the CPI instruction
                let signer = &ctx.accounts.pool_signer;
                // first remaining accounts is the wrapped program
                require!(
                    !ctx.remaining_accounts.is_empty(),
                    CcipTokenPoolError::InvalidInputs
                );
                let (wrapped_program, remaining_accounts) =
                    ctx.remaining_accounts.split_first().unwrap();

                // The accounts of the user that will be used in the CPI instruction, none of them are signers
                // They need to specify if mutable or not, but none of them is allowed to init, realloc or close
                // note: CPI signer is always first account
                let mut acc_infos = signer.to_account_infos();
                acc_infos.extend_from_slice(&[
                    ctx.accounts.pool_token_account.to_account_info(),
                    ctx.accounts.mint.to_account_info(),
                    ctx.accounts.token_program.to_account_info(),
                    ctx.accounts.receiver_token_account.to_account_info(),
                ]);
                acc_infos.extend_from_slice(remaining_accounts);

                let acc_metas: Vec<AccountMeta> = acc_infos
                    .to_vec()
                    .iter()
                    .flat_map(|acc_info| {
                        // Check signer from PDA External Execution config
                        let is_signer = acc_info.key() == signer.key();
                        acc_info.to_account_metas(Some(is_signer))
                    })
                    .collect();

                let mut data = RELEASE_MINT.to_vec();
                data.extend_from_slice(&release_or_mint.try_to_vec()?);
                let ix = Instruction {
                    program_id: wrapped_program.key(),
                    accounts: acc_metas,
                    data,
                };

                let seeds = &[
                    POOL_SIGNER_SEED,
                    &ctx.accounts.mint.key().to_bytes(),
                    &[ctx.bumps.pool_signer],
                ];
                invoke_signed(&ix, &acc_infos, &[&seeds[..]])?;
            }
        };

        Ok(ReleaseOrMintOutV1 {
            destination_amount: parsed_amount,
        })
    }

    pub fn lock_or_burn_tokens<'info>(
        ctx: Context<'_, '_, '_, 'info, TokenOnramp<'info>>,
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

        match ctx.accounts.state.pool_type {
            PoolType::LockAndRelease => lock_tokens(ctx.accounts.authority.key(), lock_or_burn)?,
            PoolType::BurnAndMint => burn_tokens(
                ctx.accounts.token_program.key(),
                ctx.accounts.pool_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.pool_signer.to_account_info(),
                ctx.bumps.pool_signer,
                ctx.accounts.authority.key(),
                lock_or_burn,
            )?,
            PoolType::Wrapped => {
                // The External Execution Config Account is used to sign the CPI instruction
                let signer = &ctx.accounts.pool_signer;
                // first remaining accounts is the wrapped program
                require!(
                    !ctx.remaining_accounts.is_empty(),
                    CcipTokenPoolError::InvalidInputs
                );
                let (wrapped_program, remaining_accounts) =
                    ctx.remaining_accounts.split_first().unwrap();

                // The accounts of the user that will be used in the CPI instruction, none of them are signers
                // They need to specify if mutable or not, but none of them is allowed to init, realloc or close
                // note: CPI signer is always first account
                let mut acc_infos = signer.to_account_infos();
                acc_infos.extend_from_slice(&[
                    ctx.accounts.pool_token_account.to_account_info(),
                    ctx.accounts.mint.to_account_info(),
                    ctx.accounts.token_program.to_account_info(),
                ]);
                acc_infos.extend_from_slice(remaining_accounts);

                let acc_metas: Vec<AccountMeta> = acc_infos
                    .to_vec()
                    .iter()
                    .flat_map(|acc_info| {
                        // Check signer from PDA External Execution config
                        let is_signer = acc_info.key() == signer.key();
                        acc_info.to_account_metas(Some(is_signer))
                    })
                    .collect();

                let mut data = LOCK_BURN.to_vec();
                data.extend_from_slice(&lock_or_burn.try_to_vec()?);
                let ix = Instruction {
                    program_id: wrapped_program.key(),
                    accounts: acc_metas,
                    data,
                };

                let seeds = &[
                    POOL_SIGNER_SEED,
                    &ctx.accounts.mint.key().to_bytes(),
                    &[ctx.bumps.pool_signer],
                ];
                invoke_signed(&ix, &acc_infos, &[&seeds[..]])?
            }
        };

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.base.remote.token_address.clone(),
            dest_pool_data: {
                let mut abi_encoded_decimals = vec![0u8; 32];
                abi_encoded_decimals[31] = ctx.accounts.state.config.decimals;
                abi_encoded_decimals
            },
        })
    }
}

// NOTE: accounts derivations must be native to program - will cause ownership check issues if imported
// wraps functionality from shared Config and Chain types
#[account]
#[derive(InitSpace)]
pub struct State {
    pub pool_type: PoolType,
    pub config: BaseConfig,
}

#[account]
#[derive(InitSpace)]
pub struct ChainConfig {
    pub base: BaseChain,
}

#[repr(u8)]
#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub enum PoolType {
    LockAndRelease,
    BurnAndMint,
    Wrapped, // wrap forwards the CPI call to an arbitrary program, assumes CPI call will handle programs pass to the pool
}
