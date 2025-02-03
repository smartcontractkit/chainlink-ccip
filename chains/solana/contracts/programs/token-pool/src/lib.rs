use std::ops::Deref;

use anchor_lang::prelude::*;
use anchor_lang::solana_program::{instruction::Instruction, program::invoke_signed};
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id,
    token_2022::spl_token_2022::{
        self,
        instruction::{burn, mint_to, transfer_checked},
        solana_zk_token_sdk::curve25519::scalar::Zeroable,
    },
};
use spl_math::uint::U256;

declare_id!("GRvFSLwR7szpjgNEZbGe4HtxfJYXqySXuuRUAJDpu4WH");

mod context;
use crate::context::*;

mod rate_limiter;
use crate::rate_limiter::*;

#[program]
pub mod token_pool {
    use super::*;

    pub fn initialize(
        ctx: Context<InitializeTokenPool>,
        pool_type: PoolType,
        // The signer for mint/release + burn/lock method calls.
        ramp_authority: Pubkey,
    ) -> Result<()> {
        let token_info = ctx.accounts.mint.to_account_info();

        let config = &mut ctx.accounts.config;
        config.pool_type = pool_type;
        config.token_program = *token_info.owner;
        config.mint = ctx.accounts.mint.key();
        config.decimals = ctx.accounts.mint.decimals;
        config.pool_signer = ctx.accounts.pool_signer.key();
        config.pool_token_account = get_associated_token_address_with_program_id(
            &config.pool_signer,
            &config.mint,
            &config.token_program,
        );
        config.owner = ctx.accounts.authority.key();
        config.ramp_authority = ramp_authority;

        Ok(())
    }

    pub fn transfer_ownership(ctx: Context<SetConfig>, proposed_owner: Pubkey) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(
            proposed_owner != config.owner && proposed_owner != Pubkey::zeroed(),
            CcipTokenPoolError::InvalidInputs
        );
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    /// Must be signed by the proposed owner
    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.owner = std::mem::take(&mut config.proposed_owner);
        config.proposed_owner = Pubkey::new_from_array([0; 32]);
        Ok(())
    }

    /// `set_ramp_authority` changes the expected signer for mint/release + burn/lock method calls
    /// this is used to update the router address
    pub fn set_ramp_authority(ctx: Context<SetConfig>, new_authority: Pubkey) -> Result<()> {
        require!(
            new_authority != Pubkey::zeroed(),
            CcipTokenPoolError::InvalidInputs
        );

        let old_authority = ctx.accounts.config.ramp_authority;
        ctx.accounts.config.ramp_authority = new_authority;
        emit!(RouterUpdated {
            old_authority,
            new_authority
        });
        Ok(())
    }

    /// Remote pools must be empty as it must be zero sized, but they can be immediately
    /// added with a call to the `append_remote_pool_addresses` instruction
    pub fn init_chain_remote_config(
        ctx: Context<InitializeChainConfig>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        ctx.accounts.chain_config.remote = cfg;

        emit!(RemoteChainConfigured {
            chain_selector: remote_chain_selector,
            token: ctx.accounts.chain_config.remote.token_address.clone(),
            pool_addresses: ctx.accounts.chain_config.remote.pool_addresses.clone(),
            previous_token: None,
            previous_pool_addresses: None,
        });

        Ok(())
    }

    /// Remote pools can be modified arbitrarily, the account space will be dynamically reallocated.
    pub fn edit_chain_remote_config(
        ctx: Context<EditChainConfigDynamicSize>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        let old_mint = ctx.accounts.chain_config.remote.token_address.clone();
        let old_pools = ctx.accounts.chain_config.remote.pool_addresses.clone();

        ctx.accounts.chain_config.remote = cfg;

        emit!(RemoteChainConfigured {
            chain_selector: remote_chain_selector,
            token: ctx.accounts.chain_config.remote.token_address.clone(),
            previous_token: Some(old_mint),
            pool_addresses: ctx.accounts.chain_config.remote.pool_addresses.clone(),
            previous_pool_addresses: Some(old_pools),
        });

        Ok(())
    }

    pub fn append_remote_pool_addresses(
        ctx: Context<AppendRemotePoolAddresses>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        addresses: Vec<RemoteAddress>,
    ) -> Result<()> {
        let old_pools = ctx.accounts.chain_config.remote.pool_addresses.clone();

        ctx.accounts
            .chain_config
            .remote
            .pool_addresses
            .append(&mut addresses.clone());

        emit!(RemotePoolsAppended {
            chain_selector: remote_chain_selector,
            pool_addresses: ctx.accounts.chain_config.remote.pool_addresses.clone(),
            previous_pool_addresses: old_pools,
        });

        Ok(())
    }

    pub fn set_chain_rate_limit(
        ctx: Context<EditChainConfigStaticSize>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        inbound: RateLimitConfig,
        outbound: RateLimitConfig,
    ) -> Result<()> {
        validate_token_bucket_config(&inbound)?;
        validate_token_bucket_config(&outbound)?;

        ctx.accounts.chain_config.inbound_rate_limit.cfg = inbound.clone();
        ctx.accounts.chain_config.outbound_rate_limit.cfg = outbound.clone();

        emit!(RateLimitConfigured {
            chain_selector: remote_chain_selector,
            outbound_rate_limit: outbound,
            inbound_rate_limit: inbound,
        });

        Ok(())
    }

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

    pub fn release_or_mint_tokens<'info>(
        ctx: Context<'_, '_, '_, 'info, TokenOfframp<'info>>,
        // All information necessary to perform a release or mint operation.
        release_or_mint: ReleaseOrMintInputV1,
    ) -> Result<ReleaseOrMintOutputV1> {
        let parsed_amount = to_svm_token_amount(
            release_or_mint.amount,
            ctx.accounts.chain_config.remote.decimals,
            ctx.accounts.config.decimals,
        )?;

        let ChainConfig {
            remote,
            inbound_rate_limit,
            ..
        } = &mut *ctx.accounts.chain_config;

        validate_release_or_mint(
            &release_or_mint,
            parsed_amount,
            ctx.accounts.config.mint,
            &remote.pool_addresses,
            inbound_rate_limit,
        )?;

        let receiver = release_or_mint.receiver;

        match ctx.accounts.config.pool_type {
            PoolType::LockAndRelease => release_tokens_to_receiver(ctx, parsed_amount, receiver)?,
            PoolType::BurnAndMint => mint_tokens_to_receiver(ctx, parsed_amount, receiver)?,
            PoolType::Wrapped => invoke_external_program_offramp(ctx, &release_or_mint)?,
        };

        Ok(ReleaseOrMintOutputV1 {
            destination_amount: parsed_amount,
        })
    }

    pub fn lock_or_burn_tokens<'info>(
        ctx: Context<'_, '_, '_, 'info, TokenOnramp<'info>>,
        // All information necessary to perform a lock or burn operation.
        lock_or_burn: LockOrBurnInV1,
    ) -> Result<LockOrBurnOutV1> {
        validate_lock_or_burn(
            &lock_or_burn,
            ctx.accounts.config.mint,
            &mut ctx.accounts.chain_config.outbound_rate_limit,
        )?;

        match ctx.accounts.config.pool_type {
            PoolType::LockAndRelease => {
                // Sender -> token pool (occurs outside pool). Tokens are held.
                emit!(Locked {
                    sender: ctx.accounts.authority.key(),
                    amount: lock_or_burn.amount,
                });
            }
            PoolType::BurnAndMint => burn_held_tokens(&ctx, lock_or_burn.amount)?,
            PoolType::Wrapped => invoke_external_program_onramp(&ctx, &lock_or_burn)?,
        };

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.remote.token_address.clone(),
            dest_pool_data: RemoteAddress::ZERO,
        })
    }
}

fn validate_lock_or_burn(
    lock_or_burn_in: &LockOrBurnInV1,
    config_mint: Pubkey,
    outbound_rate_limit: &mut RateLimitTokenBucket,
) -> Result<()> {
    require!(
        config_mint == lock_or_burn_in.local_token,
        CcipTokenPoolError::InvalidToken
    );

    // Ensures the rate limit is respected.
    outbound_rate_limit.consume(lock_or_burn_in.amount)
}

fn validate_release_or_mint(
    release_or_mint_in: &ReleaseOrMintInputV1,
    parsed_amount: u64,
    config_mint: Pubkey,
    configured_pool_addresses: &[RemoteAddress],
    inbound_rate_limit: &mut RateLimitTokenBucket,
) -> Result<()> {
    require_eq!(
        config_mint,
        release_or_mint_in.local_token,
        CcipTokenPoolError::InvalidToken
    );

    let pool_is_supported = !configured_pool_addresses.is_empty()
        && !release_or_mint_in.source_pool_address.is_empty()
        && configured_pool_addresses.contains(&release_or_mint_in.source_pool_address);

    require!(
        pool_is_supported,
        CcipTokenPoolError::InvalidSourcePoolAddress
    );

    // Ensures the rate limit is respected.
    inbound_rate_limit.consume(parsed_amount)
}

#[account]
#[derive(InitSpace)]
pub struct Config {
    pub version: u8,
    pub pool_type: PoolType,

    pub token_program: Pubkey,
    pub mint: Pubkey,
    pub decimals: u8,
    // PDA for token ownership
    pub pool_signer: Pubkey,
    // associated token account for pool_signer + token
    pub pool_token_account: Pubkey,

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,
    // signer for CCIP calls
    ramp_authority: Pubkey,
}

#[account]
#[derive(InitSpace)]
pub struct ExternalExecutionConfig {}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnInV1 {
    pub receiver: RemoteAddress,
    pub remote_chain_selector: u64,
    pub original_sender: Pubkey,
    // local solana amount to lock/burn, denominated in the source token's decimals,
    // using Solana's native type for token amounts (u64)
    pub amount: u64,
    pub local_token: Pubkey,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnOutV1 {
    pub dest_token_address: RemoteAddress,
    pub dest_pool_data: RemoteAddress,
}

/// All information necessary to perform a release or mint operation. As it's an input to
/// the offramp process, the sender is remote and the receiver local.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintInputV1 {
    pub original_sender: RemoteAddress,
    pub remote_chain_selector: u64,
    pub receiver: Pubkey,
    /// Encoded as u256, denominated in the source token's decimals.
    pub amount: [u8; 32],
    pub local_token: Pubkey,
    /// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
    /// expected pool address for the given remoteChainSelector.
    pub source_pool_address: RemoteAddress,
    /// Opaque data received from the source pool to process the release or mint.
    pub source_pool_data: RemoteAddress,
    /// @dev WARNING: offchainTokenData is untrusted data.
    ///  Opaque offchain data to process the release or mint
    pub offchain_token_data: RemoteAddress,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintOutputV1 {
    /// The number of tokens released or minted on the destination chain, denominated in the local token's decimals.
    /// This value is expected to be equal to the ReleaseOrMintInV1.amount in the case where the source and destination
    /// chain have the same number of decimals.
    pub destination_amount: u64,
}

#[repr(u8)]
#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub enum PoolType {
    LockAndRelease,
    BurnAndMint,
    /// Wrap forwards the CPI call to an arbitrary program, assumes CPI call will handle programs passed to the pool.
    Wrapped,
}

#[account]
#[derive(InitSpace)]
pub struct ChainConfig {
    pub remote: RemoteConfig,
    pub inbound_rate_limit: RateLimitTokenBucket,
    pub outbound_rate_limit: RateLimitTokenBucket,
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RemoteConfig {
    /// Allowed remote pool addresses. It's stored as an unbounded vector to ensure it's always possible
    /// to upgrade the config with new allowed addresses, without invalidating in-flight messages.
    #[max_len(0)]
    pub pool_addresses: Vec<RemoteAddress>,
    pub token_address: RemoteAddress,
    /// This value (remote decimals) must be tracked to convert amounts properly from remote to local types.
    pub decimals: u8,
}

impl RemoteConfig {
    pub fn space(&self) -> usize {
        Self::INIT_SPACE
            + self
                .pool_addresses
                .iter()
                .map(RemoteAddress::space)
                .sum::<usize>()
    }
}

/// Represents an address from a non-SVM remote chain. In the case of EVM chains, it's ABI encoded.
#[derive(Default, InitSpace, Clone, AnchorDeserialize, AnchorSerialize, PartialEq, Eq)]
pub struct RemoteAddress {
    #[max_len(64)]
    pub address: Vec<u8>,
}

impl RemoteAddress {
    const ZERO: Self = Self {
        address: Vec::new(),
    };

    pub fn space(&self) -> usize {
        // vector prefix
        4
        // data length in bytes
        + self.address.len()
    }
}

impl Deref for RemoteAddress {
    type Target = Vec<u8>;

    fn deref(&self) -> &Self::Target {
        &self.address
    }
}

pub fn to_svm_token_amount(
    // LE encoded u256
    incoming_amount_bytes: [u8; 32],
    incoming_decimal: u8,
    local_decimal: u8,
) -> Result<u64> {
    let mut incoming_amount = U256::from_little_endian(&incoming_amount_bytes);

    // handle differences in decimals by multipling/dividing by 10^N
    match incoming_decimal.cmp(&local_decimal) {
        std::cmp::Ordering::Less => {
            incoming_amount = incoming_amount
                .checked_mul(U256::exp10((local_decimal - incoming_decimal) as usize))
                .ok_or(CcipTokenPoolError::InvalidTokenAmountConversion)?;
        }
        std::cmp::Ordering::Equal => {
            // nothing needed if decimals are equal
        }
        std::cmp::Ordering::Greater => {
            incoming_amount = incoming_amount
                .checked_div(U256::exp10((incoming_decimal - local_decimal) as usize))
                .ok_or(CcipTokenPoolError::InvalidTokenAmountConversion)?;
        }
    }

    // check to prevent panic
    require!(
        incoming_amount <= U256::from(u64::MAX),
        CcipTokenPoolError::InvalidTokenAmountConversion
    );
    Ok(incoming_amount.as_u64())
}

fn release_tokens_to_receiver<'info>(
    ctx: Context<'_, '_, '_, 'info, TokenOfframp<'info>>,
    token_amount: u64,
    receiver: Pubkey,
) -> Result<()> {
    // Build the instruction to transfer from pool -> receiver
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.transfer.html
    let mut ix = transfer_checked(
        // use spl-token-2022 to build the instruction - change program later (`transfer_checked`
        // doesn't actually transfer anything, it just constructs the instruction.)
        &spl_token_2022::ID,
        &ctx.accounts.pool_token_account.key(),
        &ctx.accounts.mint.key(),
        &ctx.accounts.receiver_token_account.key(),
        &ctx.accounts.pool_signer.key(),
        &[],
        token_amount,
        ctx.accounts.mint.decimals,
    )?;
    // Replace with user specified program
    ix.program_id = ctx.accounts.token_program.key();

    let seeds = &[
        seed::CCIP_TOKENPOOL_SIGNER,
        &ctx.accounts.mint.key().to_bytes(),
        &[ctx.bumps.pool_signer],
    ];
    invoke_signed(
        &ix,
        &[
            ctx.accounts.pool_token_account.to_account_info(),
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.receiver_token_account.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
        ],
        &[&seeds[..]],
    )?;

    emit!(Released {
        sender: ctx.accounts.pool_signer.key(),
        recipient: receiver,
        amount: token_amount,
    });

    Ok(())
}

fn mint_tokens_to_receiver<'info>(
    ctx: Context<'_, '_, '_, 'info, TokenOfframp<'info>>,
    token_amount: u64,
    receiver: Pubkey,
) -> Result<()> {
    // Instruction to mint to receiver
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.mint_to.html
    let mut ix = mint_to(
        // use spl-token-2022 to build the instruction - change program later (`mint_to`
        // doesn't actually transfer anything, it just constructs the instruction.)
        &spl_token_2022::ID,
        &ctx.accounts.mint.key(),
        &ctx.accounts.receiver_token_account.key(),
        &ctx.accounts.pool_signer.key(),
        &[],
        token_amount,
    )?;
    // Replace with user specified program
    ix.program_id = ctx.accounts.token_program.key();

    let seeds = &[
        seed::CCIP_TOKENPOOL_SIGNER,
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
        recipient: receiver,
        amount: token_amount,
    });

    Ok(())
}

pub fn burn_held_tokens<'info>(
    ctx: &Context<'_, '_, '_, 'info, TokenOnramp<'info>>,
    token_amount: u64,
) -> Result<()> {
    // Receiver -> token pool (occurs outside pool).
    // https://docs.rs/spl-token-2022/latest/spl_token_2022/instruction/fn.burn.html
    // Builds instruction to burn tokens held in the pool.
    let mut ix = burn(
        // use spl-token-2022 to build the instruction - change program later (`mint_to`
        // doesn't actually transfer anything, it just constructs the instruction.)
        &spl_token_2022::ID,
        &ctx.accounts.pool_token_account.key(),
        &ctx.accounts.mint.key(),
        &ctx.accounts.pool_signer.key(),
        &[],
        token_amount,
    )?;
    // Replace with user specified program
    ix.program_id = ctx.accounts.token_program.key();

    let seeds = &[
        seed::CCIP_TOKENPOOL_SIGNER,
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
        amount: token_amount,
    });
    Ok(())
}

fn invoke_external_program_offramp<'info>(
    ctx: Context<'_, '_, '_, 'info, TokenOfframp<'info>>,
    release_or_mint_in: &ReleaseOrMintInputV1,
) -> Result<()> {
    // The External Execution Config Account is used to sign the CPI instruction
    let signer = &ctx.accounts.pool_signer;
    require!(
        !ctx.remaining_accounts.is_empty(),
        CcipTokenPoolError::InvalidInputs
    );
    let (wrapped_program, remaining_accounts) = ctx.remaining_accounts.split_first().unwrap();

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
    data.extend_from_slice(&release_or_mint_in.try_to_vec()?);
    let ix = Instruction {
        program_id: wrapped_program.key(),
        accounts: acc_metas,
        data,
    };

    let seeds = &[
        seed::CCIP_TOKENPOOL_SIGNER,
        &ctx.accounts.mint.key().to_bytes(),
        &[ctx.bumps.pool_signer],
    ];
    invoke_signed(&ix, &acc_infos, &[&seeds[..]])?;

    Ok(())
}

fn invoke_external_program_onramp<'info>(
    ctx: &Context<'_, '_, '_, 'info, TokenOnramp<'info>>,
    lock_or_burn: &LockOrBurnInV1,
) -> Result<()> {
    // The External Execution Config Account is used to sign the CPI instruction
    let signer = &ctx.accounts.pool_signer;
    // first remaining accounts is the wrapped program
    require!(
        !ctx.remaining_accounts.is_empty(),
        CcipTokenPoolError::InvalidInputs
    );
    let (wrapped_program, remaining_accounts) = ctx.remaining_accounts.split_first().unwrap();

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
        seed::CCIP_TOKENPOOL_SIGNER,
        &ctx.accounts.mint.key().to_bytes(),
        &[ctx.bumps.pool_signer],
    ];
    invoke_signed(&ix, &acc_infos, &[&seeds[..]])?;

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;
    const BASE_VALUE: u64 = 1_000_000_000; // 1e9

    #[test]
    fn test_larger_incoming_decimal() {
        let mut u256_bytes = [0u8; 32];
        U256::from(BASE_VALUE * BASE_VALUE).to_little_endian(&mut u256_bytes);
        let local_val = to_svm_token_amount(u256_bytes, 18, 9).unwrap();
        assert!(local_val == BASE_VALUE);
    }

    #[test]
    fn test_smaller_incoming_decimal() {
        let mut u256_bytes = [0u8; 32];
        U256::from(BASE_VALUE).to_little_endian(&mut u256_bytes);
        let local_val = to_svm_token_amount(u256_bytes, 9, 18).unwrap();
        assert!(local_val == BASE_VALUE * BASE_VALUE);
    }

    #[test]
    fn test_equal_incoming_decimal() {
        let mut u256_bytes = [0u8; 32];
        U256::from(BASE_VALUE).to_little_endian(&mut u256_bytes);
        let local_val = to_svm_token_amount(u256_bytes, 9, 9).unwrap();
        assert!(local_val == BASE_VALUE);
    }

    #[test]
    fn test_u256_overflow() {
        let mut u256_bytes = [0u8; 32];
        U256::from(BASE_VALUE * BASE_VALUE).to_little_endian(&mut u256_bytes);
        let res = to_svm_token_amount(u256_bytes, 0, 18);
        assert!(res.is_err());
    }

    #[test]
    fn test_u256_divide_to_zero() {
        let mut u256_bytes = [0u8; 32];
        U256::from(BASE_VALUE).to_little_endian(&mut u256_bytes);
        let local_val = to_svm_token_amount(u256_bytes, 18, 0).unwrap();
        assert!(local_val == 0);
    }

    #[test]
    fn test_u64_overflow() {
        let mut u256_bytes = [0u8; 32];
        U256::from(u64::MAX)
            .checked_add(U256::from(1))
            .unwrap()
            .to_little_endian(&mut u256_bytes);
        let res = to_svm_token_amount(u256_bytes, 0, 0);
        assert!(res.is_err());
    }
}
