use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id, token_interface::Mint,
};
use rmn_remote::state::CurseSubject;
use spl_math::uint::U256;
use std::ops::Deref;

use crate::rate_limiter::{RateLimitConfig, RateLimitTokenBucket};

pub const ANCHOR_DISCRIMINATOR: usize = 8; // 8-byte anchor discriminator length
pub const POOL_CHAINCONFIG_SEED: &[u8] = b"ccip_tokenpool_chainconfig"; // seed used by CCIP to provide correct chain config to pool
pub const POOL_STATE_SEED: &[u8] = b"ccip_tokenpool_config";
pub const POOL_SIGNER_SEED: &[u8] = b"ccip_tokenpool_signer";

pub const EXTERNAL_TOKEN_POOLS_SIGNER: &[u8] = b"external_token_pools_signer";
pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";

pub const fn valid_version(v: u8, max_version: u8) -> bool {
    !uninitialized(v) && v <= max_version
}

pub const fn uninitialized(v: u8) -> bool {
    v == 0
}

// discriminators
#[allow(dead_code)]
pub const RELEASE_MINT: [u8; 8] = [0x14, 0x94, 0x71, 0xc6, 0xe5, 0xaa, 0x47, 0x30];
#[allow(dead_code)]
pub const LOCK_BURN: [u8; 8] = [0xc8, 0x0e, 0x32, 0x09, 0x2c, 0x5b, 0x79, 0x25];

#[derive(InitSpace, AnchorSerialize, AnchorDeserialize, Clone)]
pub struct BaseConfig {
    // token config
    pub token_program: Pubkey,
    pub mint: Pubkey,
    pub decimals: u8,
    pub pool_signer: Pubkey,        // PDA for token ownership
    pub pool_token_account: Pubkey, // associated token account for pool_signer + token

    // ownership
    pub owner: Pubkey,
    pub proposed_owner: Pubkey,
    pub rate_limit_admin: Pubkey,
    pub router_onramp_authority: Pubkey, // signer PDA for CCIP onramp calls
    pub router: Pubkey,                  // router address

    // rebalancing - only used for lock/release pools
    pub rebalancer: Pubkey,
    pub can_accept_liquidity: bool,

    // lockOrBurn allow list
    pub list_enabled: bool,
    #[max_len(0)]
    // max_len = 0 for InitSpace calculation, the actual length is calculated using the length of the allow_list. `realloc` should be used to handle the increase in account size accordingly
    pub allow_list: Vec<Pubkey>,

    // RMN Remote
    pub rmn_remote: Pubkey,
}

impl BaseConfig {
    pub fn init(
        mint: &InterfaceAccount<Mint>,
        pool_program: Pubkey,
        owner: Pubkey,
        router: Pubkey,
        rmn_remote: Pubkey,
    ) -> Self {
        let token_info = mint.to_account_info();
        let token_program = *token_info.owner;

        let (pool_signer, _) =
            Pubkey::find_program_address(&[POOL_SIGNER_SEED, mint.key().as_ref()], &pool_program);

        let (router_onramp_authority, _) = Pubkey::find_program_address(
            &[EXTERNAL_TOKEN_POOLS_SIGNER, pool_program.as_ref()],
            &router,
        );

        let pool_token_account =
            get_associated_token_address_with_program_id(&pool_signer, &mint.key(), &token_program);

        Self {
            token_program,
            mint: mint.key(),
            decimals: mint.decimals,
            pool_signer,
            pool_token_account,
            owner,
            proposed_owner: Pubkey::default(),
            rate_limit_admin: owner,
            router_onramp_authority,
            router,
            rebalancer: Pubkey::default(),
            can_accept_liquidity: false,
            list_enabled: false,
            allow_list: vec![],
            rmn_remote,
        }
    }

    pub fn transfer_ownership(&mut self, proposed_owner: Pubkey) -> Result<()> {
        require!(
            proposed_owner != self.owner && proposed_owner != Pubkey::default(),
            CcipTokenPoolError::InvalidInputs
        );
        self.proposed_owner = proposed_owner;

        emit!(OwnershipTransferRequested {
            mint: self.mint,
            from: self.owner,
            to: self.proposed_owner,
        });
        Ok(())
    }

    pub fn accept_ownership(&mut self) -> Result<()> {
        let old_owner = self.owner;
        // NOTE: take() resets proposed_owner to default
        self.owner = std::mem::take(&mut self.proposed_owner);

        emit!(OwnershipTransferred {
            mint: self.mint,
            from: old_owner,
            to: self.owner,
        });

        Ok(())
    }

    pub fn set_router(&mut self, new_router: Pubkey, pool_program: &Pubkey) -> Result<()> {
        require_keys_neq!(
            new_router,
            Pubkey::default(),
            CcipTokenPoolError::InvalidInputs
        );

        let old_router = self.router;
        self.router = new_router;
        (self.router_onramp_authority, _) = Pubkey::find_program_address(
            &[EXTERNAL_TOKEN_POOLS_SIGNER, pool_program.as_ref()],
            &new_router,
        );
        emit!(RouterUpdated {
            mint: self.mint,
            old_router,
            new_router,
        });
        Ok(())
    }

    pub fn set_rebalancer(&mut self, address: Pubkey) -> Result<()> {
        self.rebalancer = address;
        Ok(())
    }

    pub fn set_can_accept_liquidity(&mut self, allow: bool) -> Result<()> {
        self.can_accept_liquidity = allow;
        Ok(())
    }
}

#[derive(InitSpace, AnchorSerialize, AnchorDeserialize, Clone)]
pub struct BaseChain {
    pub remote: RemoteConfig,
    pub inbound_rate_limit: RateLimitTokenBucket,
    pub outbound_rate_limit: RateLimitTokenBucket,
}

impl BaseChain {
    pub fn set(
        &mut self,
        remote_chain_selector: u64,
        mint: Pubkey,
        new_cfg: RemoteConfig,
    ) -> Result<()> {
        let old_mint = self.remote.token_address.clone();
        let old_pools = self.remote.pool_addresses.clone();

        self.remote = new_cfg;

        emit!(RemoteChainConfigured {
            chain_selector: remote_chain_selector,
            mint,
            token: self.remote.token_address.clone(),
            previous_token: old_mint,
            pool_addresses: self.remote.pool_addresses.clone(),
            previous_pool_addresses: old_pools,
        });
        Ok(())
    }

    pub fn append_remote_pool_addresses(
        &mut self,
        remote_chain_selector: u64,
        mint: Pubkey,
        addresses: Vec<RemoteAddress>,
    ) -> Result<()> {
        let old_pools = self.remote.pool_addresses.clone();

        for address in addresses {
            if !self.remote.pool_addresses.contains(&address) {
                self.remote.pool_addresses.push(address);
            } else {
                return Err(CcipTokenPoolError::RemotePoolAddressAlreadyExisted.into());
            }
        }

        emit!(RemotePoolsAppended {
            chain_selector: remote_chain_selector,
            mint,
            pool_addresses: self.remote.pool_addresses.clone(),
            previous_pool_addresses: old_pools,
        });
        Ok(())
    }

    pub fn set_chain_rate_limit(
        &mut self,
        remote_chain_selector: u64,
        mint: Pubkey,
        inbound: RateLimitConfig,
        outbound: RateLimitConfig,
    ) -> Result<()> {
        self.inbound_rate_limit
            .set_token_bucket_config(inbound.clone())?;
        self.outbound_rate_limit
            .set_token_bucket_config(outbound.clone())?;

        emit!(RateLimitConfigured {
            chain_selector: remote_chain_selector,
            mint,
            outbound_rate_limit: outbound,
            inbound_rate_limit: inbound,
        });

        Ok(())
    }
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RemoteConfig {
    // Remote pool addresses are a vec to support multiple pool versions of the same token.
    // Although for a given remote chain, there should be one active pool address at a time,
    // tokens could have been sent from earlier versions of that pool.
    // On SVM, pools are expected to upgrade in place, but on EVM, earlier version have different addresses.
    // Tokens could be stuck in flight for a long time, we need to support validation of these tokens even
    // after the pool they were sent from has been decommissioned at source.
    #[max_len(0)] // used to calculate InitSpace
    pub pool_addresses: Vec<RemoteAddress>,
    pub token_address: RemoteAddress,
    pub decimals: u8, // needed to track decimals from remote to convert properly
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

#[derive(Default, InitSpace, Clone, AnchorDeserialize, AnchorSerialize, PartialEq, Eq)]
pub struct RemoteAddress {
    #[max_len(64)]
    pub address: Vec<u8>,
}

impl RemoteAddress {
    pub const ZERO: Self = Self {
        address: Vec::new(),
    };

    pub fn space(&self) -> usize {
        4 // vector prefix
            + self.address.len() // data length in bytes
    }
}

impl Deref for RemoteAddress {
    type Target = Vec<u8>;

    fn deref(&self) -> &Self::Target {
        &self.address
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnInV1 {
    pub receiver: Vec<u8>, //  The recipient of the tokens on the destination chain
    pub remote_chain_selector: u64, // The chain ID of the destination chain
    pub original_sender: Pubkey, // The original sender of the tx on the source chain
    pub amount: u64, // local solana amount to lock/burn,  The amount of tokens to lock or burn, denominated in the source token's decimals
    pub local_token: Pubkey, //  The address on this chain of the token to lock or burn
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnOutV1 {
    pub dest_token_address: RemoteAddress,
    // Currently encodes the local decimals as there's a chance they differ.
    pub dest_pool_data: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintInV1 {
    pub original_sender: RemoteAddress, //          The original sender of the tx on the source chain
    pub remote_chain_selector: u64,     // ─╮ The chain ID of the source chain
    pub receiver: Pubkey, // ───────────╯ The recipient of the tokens on the destination chain.
    pub amount: [u8; 32], // u256, incoming cross-chain amount - The amount of tokens to release or mint, denominated in the source token's decimals
    pub local_token: Pubkey, //            The address on this chain of the token to release or mint
    /// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
    /// expected pool address for the given remoteChainSelector.
    pub source_pool_address: RemoteAddress, //       The address of the source pool, abi encoded in the case of EVM chains
    pub source_pool_data: Vec<u8>, //          The data received from the source pool to process the release or mint
    /// @dev WARNING: offchainTokenData is untrusted data.
    pub offchain_token_data: Vec<u8>, //       The offchain data to process the release or mint
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintOutV1 {
    // The number of tokens released or minted on the destination chain, denominated in the local token's decimals.
    // This value is expected to be equal to the ReleaseOrMintInV1.amount in the case where the source and destination
    // chain have the same number of decimals.
    pub destination_amount: u64, // token amounts local to solana
}

#[event]
pub struct Burned {
    pub sender: Pubkey,
    pub amount: u64,
    pub mint: Pubkey,
}

#[event]
pub struct Minted {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
    pub mint: Pubkey,
}

#[event]
pub struct Locked {
    pub sender: Pubkey,
    pub amount: u64,
    pub mint: Pubkey,
}

#[event]
pub struct Released {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
    pub mint: Pubkey,
}

// note: configuration events are slightly different than EVM chains because configuration follows different steps
#[event]
pub struct RemoteChainConfigured {
    pub chain_selector: u64,
    pub token: RemoteAddress,
    pub previous_token: RemoteAddress,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
    pub mint: Pubkey,
}

#[event]
pub struct RateLimitConfigured {
    pub chain_selector: u64,
    pub outbound_rate_limit: RateLimitConfig,
    pub inbound_rate_limit: RateLimitConfig,
    pub mint: Pubkey,
}

#[event]
pub struct RemotePoolsAppended {
    pub chain_selector: u64,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
    pub mint: Pubkey,
}

#[event]
pub struct RemoteChainRemoved {
    pub chain_selector: u64,
    pub mint: Pubkey,
}

#[event]
pub struct RouterUpdated {
    pub old_router: Pubkey,
    pub new_router: Pubkey,
    pub mint: Pubkey,
}

#[event]
pub struct OwnershipTransferRequested {
    pub from: Pubkey,
    pub to: Pubkey,
    pub mint: Pubkey,
}

#[event]
pub struct OwnershipTransferred {
    pub from: Pubkey,
    pub to: Pubkey,
    pub mint: Pubkey,
}

#[error_code]
pub enum CcipTokenPoolError {
    #[msg("Pool authority does not match token mint owner")]
    InvalidInitPoolPermissions,
    #[msg("Invalid RMN Remote Address")]
    InvalidRMNRemoteAddress,
    #[msg("Unauthorized")]
    Unauthorized,
    #[msg("Invalid inputs")]
    InvalidInputs,
    #[msg("Invalid state version")]
    InvalidVersion,
    #[msg("Caller is not ramp on router")]
    InvalidPoolCaller,
    #[msg("Sender not allowed")]
    InvalidSender,
    #[msg("Invalid source pool address")]
    InvalidSourcePoolAddress,
    #[msg("Invalid token")]
    InvalidToken,
    #[msg("Invalid token amount conversion")]
    InvalidTokenAmountConversion,
    #[msg("Key already existed in the allowlist")]
    AllowlistKeyAlreadyExisted,
    #[msg("Key did not exist in the allowlist")]
    AllowlistKeyDidNotExist,
    #[msg("Remote pool address already exists")]
    RemotePoolAddressAlreadyExisted,
    #[msg("Expected empty pool addresses during initialization")]
    NonemptyPoolAddressesInit,

    // Rate limit errors
    #[msg("RateLimit: bucket overfilled")]
    RLBucketOverfilled,
    #[msg("RateLimit: max capacity exceeded")]
    RLMaxCapacityExceeded,
    #[msg("RateLimit: rate limit reached")]
    RLRateLimitReached,
    #[msg("RateLimit: invalid rate limit rate")]
    RLInvalidRateLimitRate,
    #[msg("RateLimit: disabled non-zero rate limit")]
    RLDisabledNonZeroRateLimit,

    // Lock/Release errors
    #[msg("Liquidity not accepted")]
    LiquidityNotAccepted,
}

// validate_lock_or_burn checks for correctness on inputs
// - token & pool is correct for chain
// - rate limiting
// - destination chain is not cursed
// - there is no global curse
#[allow(clippy::too_many_arguments)]
pub fn validate_lock_or_burn<'info>(
    lock_or_burn_in: &LockOrBurnInV1,
    config_mint: Pubkey,
    outbound_rate_limit: &mut RateLimitTokenBucket,
    allow_list_enabled: bool,
    allow_list: &[Pubkey],
    rmn_remote: AccountInfo<'info>,
    rmn_remote_curses: AccountInfo<'info>,
    rmn_remote_config: AccountInfo<'info>,
) -> Result<()> {
    // validate token matches configured pool token
    require!(
        config_mint == lock_or_burn_in.local_token,
        CcipTokenPoolError::InvalidToken
    );

    require!(
        !allow_list_enabled || allow_list.contains(&lock_or_burn_in.original_sender),
        CcipTokenPoolError::InvalidSender
    );

    verify_uncursed_cpi(
        rmn_remote,
        rmn_remote_config,
        rmn_remote_curses,
        lock_or_burn_in.remote_chain_selector,
    )?;

    outbound_rate_limit.consume::<Clock>(lock_or_burn_in.amount)
}

// validate_release_or_mint checks for correctness on inputs
// - token & pool is correct for chain
// - rate limiting
// - source chain is not cursed
// - there is no global curse
#[allow(clippy::too_many_arguments)]
pub fn validate_release_or_mint<'info>(
    release_or_mint_in: &ReleaseOrMintInV1,
    parsed_amount: u64,
    config_mint: Pubkey,
    pool_addresses: &[RemoteAddress],
    inbound_rate_limit: &mut RateLimitTokenBucket,
    rmn_remote: AccountInfo<'info>,
    rmn_remote_curses: AccountInfo<'info>,
    rmn_remote_config: AccountInfo<'info>,
) -> Result<()> {
    require_keys_eq!(
        config_mint,
        release_or_mint_in.local_token,
        CcipTokenPoolError::InvalidToken
    );
    require!(
        !pool_addresses.is_empty()
            && !release_or_mint_in.source_pool_address.is_empty()
            && pool_addresses.contains(&release_or_mint_in.source_pool_address),
        CcipTokenPoolError::InvalidSourcePoolAddress
    );

    verify_uncursed_cpi(
        rmn_remote,
        rmn_remote_config,
        rmn_remote_curses,
        release_or_mint_in.remote_chain_selector,
    )?;

    inbound_rate_limit.consume::<Clock>(parsed_amount)
}

pub fn verify_uncursed_cpi<'info>(
    rmn_remote: AccountInfo<'info>,
    rmn_remote_config: AccountInfo<'info>,
    rmn_remote_curses: AccountInfo<'info>,
    chain_selector: u64,
) -> Result<()> {
    let cpi_accounts = rmn_remote::cpi::accounts::InspectCurses {
        config: rmn_remote_config,
        curses: rmn_remote_curses,
    };
    let cpi_context = CpiContext::new(rmn_remote, cpi_accounts);
    rmn_remote::cpi::verify_not_cursed(
        cpi_context,
        CurseSubject::from_chain_selector(chain_selector),
    )
}

pub fn to_svm_token_amount(
    incoming_amount_bytes: [u8; 32], // LE encoded u256
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
