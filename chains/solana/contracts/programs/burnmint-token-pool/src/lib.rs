use anchor_lang::{prelude::*, solana_program::program::invoke_signed};
use anchor_spl::{
    token::spl_token::{self, instruction::MAX_SIGNERS},
    token_2022::spl_token_2022::{
        self,
        instruction::{burn, mint_to},
    },
};
use solana_program::program_pack::Pack;

use base_token_pool::{common::*, rate_limiter::*};
declare_id!("41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB");

pub mod context;
use crate::context::*;

#[program]
pub mod burnmint_token_pool {

    use super::*;

    pub fn init_global_config(
        ctx: Context<InitGlobalConfig>,
        router_address: Pubkey,
        rmn_address: Pubkey,
    ) -> Result<()> {
        ctx.accounts.config.set_inner(PoolConfig {
            self_served_allowed: false,
            version: 1,
            router: router_address,
            rmn_remote: rmn_address,
        });

        emit!(GlobalConfigUpdated {
            self_served_allowed: ctx.accounts.config.self_served_allowed,
            router: ctx.accounts.config.router,
            rmn_remote: ctx.accounts.config.rmn_remote,
        });

        Ok(())
    }

    pub fn update_self_served_allowed(
        ctx: Context<UpdateGlobalConfig>,
        self_served_allowed: bool,
    ) -> Result<()> {
        ctx.accounts.config.self_served_allowed = self_served_allowed;

        emit!(GlobalConfigUpdated {
            self_served_allowed: ctx.accounts.config.self_served_allowed,
            router: ctx.accounts.config.router,
            rmn_remote: ctx.accounts.config.rmn_remote,
        });

        Ok(())
    }

    pub fn update_default_router(
        ctx: Context<UpdateGlobalConfig>,
        router_address: Pubkey,
    ) -> Result<()> {
        ctx.accounts.config.router = router_address;

        emit!(GlobalConfigUpdated {
            self_served_allowed: ctx.accounts.config.self_served_allowed,
            router: ctx.accounts.config.router,
            rmn_remote: ctx.accounts.config.rmn_remote,
        });

        Ok(())
    }

    pub fn update_default_rmn(ctx: Context<UpdateGlobalConfig>, rmn_address: Pubkey) -> Result<()> {
        ctx.accounts.config.rmn_remote = rmn_address;

        emit!(GlobalConfigUpdated {
            self_served_allowed: ctx.accounts.config.self_served_allowed,
            router: ctx.accounts.config.router,
            rmn_remote: ctx.accounts.config.rmn_remote,
        });

        Ok(())
    }

    pub fn initialize(ctx: Context<InitializeTokenPool>) -> Result<()> {
        ctx.accounts.state.set_inner(State {
            version: 1,
            config: BaseConfig::init(
                &ctx.accounts.mint,
                ctx.program_id.key(),
                ctx.accounts.authority.key(),
                ctx.accounts.config.router,
                ctx.accounts.config.rmn_remote,
            ),
        });

        Ok(())
    }

    // This method transfers the mint authority of the mint, so it does a CPI to the Token Program.
    // It is only defined in the burn and mint program as the mint authority is only used for minting tokens.
    pub fn transfer_mint_authority_to_multisig<'info>(
        ctx: Context<'_, '_, 'info, 'info, TransferMintAuthority<'info>>,
    ) -> Result<()> {
        let mint = &ctx.accounts.mint;

        let old_mint_authority = mint
            .mint_authority
            .ok_or(CcipBnMTokenPoolError::FixedMintToken)?;

        let new_mint_authority = ctx.accounts.new_multisig_mint_authority.key();

        require_keys_neq!(
            old_mint_authority,
            new_mint_authority,
            CcipBnMTokenPoolError::MintAuthorityAlreadySet
        );

        // The new Mint Authority must be a multisig account that contains the pool signer as one of its signers.
        let token_program_id = &ctx.accounts.state.config.token_program.key();

        let multisig_data = &mut &ctx.accounts.new_multisig_mint_authority.data.borrow()[..];
        // then check that the multisig account is a valid multisig account
        let multisig_account: MultisigAccount = if token_program_id == &spl_token_2022::ID {
            spl_token_2022::state::Multisig::unpack_from_slice(multisig_data)
                .map_err(|_| CcipBnMTokenPoolError::InvalidToken2022Multisig)?
                .into()
        } else if token_program_id == &spl_token::ID {
            spl_token::state::Multisig::unpack_from_slice(multisig_data)
                .map_err(|_| CcipBnMTokenPoolError::InvalidSPLTokenMultisig)?
                .into()
        } else {
            // Reject any unsupported token programs
            return Err(CcipBnMTokenPoolError::UnsupportedTokenProgram.into());
        };

        let n = multisig_account.n as usize;
        let m = multisig_account.m as usize;

        let signer_count = multisig_account
            .signers
            .iter()
            .filter(|s| *s == &ctx.accounts.pool_signer.key())
            .count();

        validate_multisig_config(n, m, signer_count)?;

        // Transfer the mint authority to the new mint authority using the corresponding Token Program. It can be token 22 or token SPL
        let ix = spl_token_2022::instruction::set_authority(
            token_program_id,
            &ctx.accounts.mint.key(),
            Some(&new_mint_authority),
            spl_token_2022::instruction::AuthorityType::MintTokens,
            &old_mint_authority,
            &[ctx.accounts.pool_signer.key],
        )?;

        let seeds = &[
            POOL_SIGNER_SEED,
            &ctx.accounts.mint.key().to_bytes(),
            &[ctx.bumps.pool_signer],
        ];

        let account_infos: Vec<AccountInfo> =
            if old_mint_authority != ctx.accounts.pool_signer.key() {
                // If the old mint authority is not the pool signer, we need to include the multisig account in the instruction
                let multisig_signer_account = ctx
                    .remaining_accounts
                    .iter()
                    .find(|acc| acc.key() == old_mint_authority)
                    .ok_or(CcipBnMTokenPoolError::InvalidMultisig)?;
                vec![
                    ctx.accounts.mint.to_account_info(),
                    multisig_signer_account.to_account_info(),
                    ctx.accounts.pool_signer.to_account_info(),
                ]
            } else {
                vec![
                    ctx.accounts.mint.to_account_info(),
                    ctx.accounts.pool_signer.to_account_info(),
                ]
            };

        invoke_signed(&ix, &account_infos, &[&seeds[..]])?;

        emit!(MintAuthorityTransferred {
            mint: ctx.accounts.mint.key(),
            old_mint_authority,
            new_mint_authority,
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
    pub fn set_router(ctx: Context<AdminUpdateTokenPool>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .config
            .set_router(new_router, ctx.program_id)
    }

    pub fn set_rmn(ctx: Context<AdminUpdateTokenPool>, rmn_address: Pubkey) -> Result<()> {
        ctx.accounts.state.config.set_rmn(rmn_address)
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

    // set rate limit admin
    pub fn set_rate_limit_admin(
        ctx: Context<SetRateLimitAdmin>,
        _mint: Pubkey,
        new_rate_limit_admin: Pubkey,
    ) -> Result<()> {
        ctx.accounts
            .state
            .config
            .set_rate_limit_admin(new_rate_limit_admin)
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
        // Cache initial length
        let initial_list_len = list.len();
        // Collect all keys to remove into a HashSet for O(1) lookups
        let keys_to_remove: std::collections::HashSet<Pubkey> = remove.into_iter().collect();
        // Perform a single pass through the list
        list.retain(|k| !keys_to_remove.contains(k));

        // We don't store repeated keys, so the keys_to_remove should match the removed keys
        require_eq!(
            initial_list_len,
            list.len().checked_add(keys_to_remove.len()).unwrap(),
            CcipTokenPoolError::AllowlistKeyDidNotExist
        );

        Ok(())
    }

    pub fn release_or_mint_tokens<'info>(
        ctx: Context<'_, '_, 'info, 'info, TokenOfframp<'info>>,
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

        // For a burnmint pool, the mint authority should never be empty (that would mean a fixed supply and prevent minting)
        // If not set, then the mint operation will fail
        require!(
            ctx.accounts.mint.mint_authority.is_some(),
            CcipBnMTokenPoolError::InvalidMultisig
        );

        let mint_authority = ctx.accounts.mint.mint_authority.unwrap();

        let multisig = if mint_authority != ctx.accounts.pool_signer.key() {
            let multisig_account = ctx
                .remaining_accounts
                .iter()
                .find(|acc| acc.key() == mint_authority)
                .ok_or(CcipBnMTokenPoolError::InvalidMultisig)?;

            Some(multisig_account)
        } else {
            None
        };

        mint_tokens(
            ctx.accounts.token_program.key(),
            release_or_mint,
            parsed_amount,
            MintTokenAccountsInfo {
                receiver_token_account: &ctx.accounts.receiver_token_account.to_account_info(),
                mint: &ctx.accounts.mint.to_account_info(),
                pool_signer: &ctx.accounts.pool_signer.to_account_info(),
                pool_signer_bump: ctx.bumps.pool_signer,
                multisig,
            },
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

        burn_tokens(
            ctx.accounts.token_program.key(),
            ctx.accounts.pool_token_account.to_account_info(),
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.bumps.pool_signer,
            ctx.accounts.authority.key(),
            lock_or_burn,
        )?;

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

// Same interface for Multisig for both the Token SPL and Token 2022 programs.
pub struct MultisigAccount {
    /// Number of signers required
    pub m: u8,
    /// Number of valid signers
    pub n: u8,
    /// Is `true` if this structure has been initialized
    pub is_initialized: bool,
    /// Signer public keys
    pub signers: [Pubkey; MAX_SIGNERS],
}

impl From<spl_token_2022::state::Multisig> for MultisigAccount {
    fn from(multisig: spl_token_2022::state::Multisig) -> Self {
        Self {
            m: multisig.m,
            n: multisig.n,
            is_initialized: multisig.is_initialized,
            signers: multisig.signers,
        }
    }
}

impl From<spl_token::state::Multisig> for MultisigAccount {
    fn from(multisig: spl_token::state::Multisig) -> Self {
        Self {
            m: multisig.m,
            n: multisig.n,
            is_initialized: multisig.is_initialized,
            signers: multisig.signers,
        }
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

pub fn burn_tokens<'a>(
    token_program: Pubkey,
    pool_token_account: AccountInfo<'a>,
    mint: AccountInfo<'a>,
    pool_signer: AccountInfo<'a>,
    pool_signer_bump: u8,
    sender: Pubkey,
    lock_or_burn: LockOrBurnInV1,
) -> Result<()> {
    // receiver -> token pool (occurs outside pool)
    // burn tokens held in pool
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.burn.html
    let mut ix = burn(
        &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
        &pool_token_account.key(),
        &mint.key(),
        &pool_signer.key(),
        &[],
        lock_or_burn.amount,
    )?;
    ix.program_id = token_program; // set to user specified program

    let seeds = &[
        POOL_SIGNER_SEED,
        &mint.key().to_bytes(),
        &[pool_signer_bump],
    ];
    invoke_signed(
        &ix,
        &[pool_token_account, mint.clone(), pool_signer],
        &[&seeds[..]],
    )?;

    emit!(Burned {
        sender,
        amount: lock_or_burn.amount,
        mint: mint.key(),
    });

    Ok(())
}

pub struct MintTokenAccountsInfo<'a, 'b> {
    pub receiver_token_account: &'a AccountInfo<'b>,
    pub mint: &'a AccountInfo<'b>,
    pub pool_signer: &'a AccountInfo<'b>,
    pub pool_signer_bump: u8,
    pub multisig: Option<&'a AccountInfo<'b>>,
}

pub fn mint_tokens(
    token_program: Pubkey,
    release_or_mint: ReleaseOrMintInV1,
    parsed_amount: u64,
    accounts: MintTokenAccountsInfo,
) -> Result<()> {
    // mint to receiver
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.mint_to.html
    let mut ix = mint_to(
        &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
        &accounts.mint.key(),
        &accounts.receiver_token_account.key(),
        &accounts.multisig.unwrap_or(accounts.pool_signer).key(),
        &[accounts.pool_signer.key],
        parsed_amount,
    )?;
    ix.program_id = token_program.key(); // set to user specified program

    let seeds = &[
        POOL_SIGNER_SEED,
        &accounts.mint.key().to_bytes(),
        &[accounts.pool_signer_bump],
    ];

    let account_infos: Vec<AccountInfo> = if accounts.multisig.is_some() {
        vec![
            accounts.receiver_token_account.clone(),
            accounts.mint.clone(),
            accounts.multisig.unwrap().clone(),
            accounts.pool_signer.clone(),
        ]
    } else {
        vec![
            accounts.receiver_token_account.clone(),
            accounts.mint.clone(),
            accounts.pool_signer.clone(),
        ]
    };

    invoke_signed(&ix, &account_infos, &[&seeds[..]])?;

    emit!(Minted {
        sender: accounts.pool_signer.key(),
        recipient: release_or_mint.receiver,
        amount: parsed_amount,
        mint: accounts.mint.key(),
    });

    Ok(())
}

pub fn validate_multisig_config(n: usize, m: usize, signer_count: usize) -> Result<()> {
    // Multisig must have at least two signers
    require_gte!(
        n,
        2,
        CcipBnMTokenPoolError::MultisigMustHaveAtLeastTwoSigners
    );

    // Threshold must be at least one
    require_gte!(
        m,
        1,
        CcipBnMTokenPoolError::MultisigMustHaveMoreThanOneSigner
    );

    // Threshold must not exceed the total number of signers
    require_gte!(n, m, CcipBnMTokenPoolError::InvalidMultisigThreshold);

    // Pool signer must appear at least m times to satisfy the threshold on its own
    require_gte!(
        signer_count,
        m,
        CcipBnMTokenPoolError::PoolSignerNotInMultisig
    );

    // Pool signer must not appear more than (n - m) times
    require_gte!(
        n - m,
        signer_count,
        CcipBnMTokenPoolError::InvalidMultisigThresholdTooHigh
    );

    Ok(())
}
