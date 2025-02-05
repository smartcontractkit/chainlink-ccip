use std::ops::Deref;

use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id, token_interface::Mint,
};
use spl_math::uint::U256;

use crate::{validate_token_bucket_config, RateLimitConfig, RateLimitTokenBucket};

pub const ANCHOR_DISCRIMINATOR: usize = 8; // 8-byte anchor discriminator length
pub const POOL_CHAINCONFIG_SEED: &[u8] = b"ccip_tokenpool_chainconfig"; // seed used by CCIP to provide correct chain config to pool
pub const POOL_STATE_SEED: &[u8] = b"ccip_tokenpool_config";
pub const POOL_SIGNER_SEED: &[u8] = b"ccip_tokenpool_signer";

// discriminators
#[allow(dead_code)]
pub const RELEASE_MINT: [u8; 8] = [0x14, 0x94, 0x71, 0xc6, 0xe5, 0xaa, 0x47, 0x30];
#[allow(dead_code)]
pub const LOCK_BURN: [u8; 8] = [0xc8, 0x0e, 0x32, 0x09, 0x2c, 0x5b, 0x79, 0x25];

#[derive(InitSpace, AnchorSerialize, AnchorDeserialize, Clone)]
pub struct Config {
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
    pub ramp_authority: Pubkey, // signer for CCIP calls

    // lockOrBurn allow list
    pub list_enabled: bool,
    #[max_len(0)]
    pub allow_list: Vec<Pubkey>,
}

impl Config {
    pub fn init(
        &mut self,
        mint: &InterfaceAccount<Mint>,
        pool_program: Pubkey,
        owner: Pubkey,
        ramp_authority: Pubkey,
    ) -> Result<()> {
        let token_info = mint.to_account_info();

        self.token_program = *token_info.owner;
        self.mint = mint.key();
        self.decimals = mint.decimals;
        (self.pool_signer, _) = Pubkey::find_program_address(
            &[POOL_SIGNER_SEED, self.mint.key().as_ref()],
            &pool_program,
        );
        self.pool_token_account = get_associated_token_address_with_program_id(
            &self.pool_signer,
            &self.mint,
            &self.token_program,
        );
        self.owner = owner;
        self.rate_limit_admin = owner;
        self.ramp_authority = ramp_authority;

        Ok(())
    }

    pub fn transfer_ownership(&mut self, proposed_owner: Pubkey) -> Result<()> {
        require!(
            proposed_owner != self.owner && proposed_owner != Pubkey::default(),
            CcipTokenPoolError::InvalidInputs
        );
        self.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(&mut self) -> Result<()> {
        self.owner = std::mem::take(&mut self.proposed_owner);
        self.proposed_owner = Pubkey::default();

        Ok(())
    }

    pub fn set_ramp_authority(&mut self, new_authority: Pubkey) -> Result<()> {
        require!(
            new_authority != Pubkey::default(),
            CcipTokenPoolError::InvalidInputs
        );

        let old_authority = self.ramp_authority;
        self.ramp_authority = new_authority;
        emit!(RouterUpdated {
            old_authority,
            new_authority
        });
        Ok(())
    }

    pub fn update_allow_list(
        &mut self,
        enabled: Option<bool>,
        add: Vec<Pubkey>,
        remove: Vec<Pubkey>,
    ) -> Result<()> {
        if enabled.is_some() {
            self.list_enabled = enabled.unwrap();
        };

        for k in remove {
            match self.allow_list.binary_search(&k) {
                Ok(v) => {
                    self.allow_list.remove(v);
                }
                Err(_) => {}
            }
        }

        for k in add {
            match self.allow_list.binary_search(&k) {
                Ok(_) => {}
                Err(v) => {
                    self.allow_list.insert(v, k);
                }
            }
        }

        Ok(())
    }
}

#[account]
#[derive(InitSpace)]
pub struct ChainConfig {
    pub remote: RemoteConfig,
    pub inbound_rate_limit: RateLimitTokenBucket,
    pub outbound_rate_limit: RateLimitTokenBucket,
}

impl ChainConfig {
    pub fn set(&mut self, remote_chain_selector: u64, new_cfg: RemoteConfig) -> Result<()> {
        let old_mint = self.remote.token_address.clone();
        let old_pools = self.remote.pool_addresses.clone();

        self.remote = new_cfg;

        emit!(RemoteChainConfigured {
            chain_selector: remote_chain_selector,
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
        addresses: Vec<RemoteAddress>,
    ) -> Result<()> {
        let old_pools = self.remote.pool_addresses.clone();

        self.remote.pool_addresses.append(&mut addresses.clone());

        emit!(RemotePoolsAppended {
            chain_selector: remote_chain_selector,
            pool_addresses: self.remote.pool_addresses.clone(),
            previous_pool_addresses: old_pools,
        });
        Ok(())
    }

    pub fn set_chain_rate_limit(
        &mut self,
        remote_chain_selector: u64,
        inbound: RateLimitConfig,
        outbound: RateLimitConfig,
    ) -> Result<()> {
        validate_token_bucket_config(&inbound)?;
        validate_token_bucket_config(&outbound)?;

        self.inbound_rate_limit.cfg = inbound.clone();
        self.outbound_rate_limit.cfg = outbound.clone();

        emit!(RateLimitConfigured {
            chain_selector: remote_chain_selector,
            outbound_rate_limit: outbound,
            inbound_rate_limit: inbound,
        });

        Ok(())
    }
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RemoteConfig {
    #[max_len(0)]
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
    pub dest_pool_data: RemoteAddress,
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
    pub source_pool_data: RemoteAddress, //          The data received from the source pool to process the release or mint
    /// @dev WARNING: offchainTokenData is untrusted data.
    pub offchain_token_data: RemoteAddress, //       The offchain data to process the release or mint
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
}

#[event]
pub struct Minted {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Locked {
    pub sender: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Released {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
}

// note: configuration events are slightly different than EVM chains because configuration follows different steps
#[event]
pub struct RemoteChainConfigured {
    pub chain_selector: u64,
    pub token: RemoteAddress,
    pub previous_token: RemoteAddress,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
}

#[event]
pub struct RateLimitConfigured {
    pub chain_selector: u64,
    pub outbound_rate_limit: RateLimitConfig,
    pub inbound_rate_limit: RateLimitConfig,
}

#[event]
pub struct RemotePoolsAppended {
    pub chain_selector: u64,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
}

#[event]
pub struct RemoteChainRemoved {
    pub chain_selector: u64,
}

#[event]
pub struct RouterUpdated {
    pub old_authority: Pubkey,
    pub new_authority: Pubkey,
}

#[error_code]
pub enum CcipTokenPoolError {
    #[msg("Pool authority does not match token mint owner")]
    InvalidInitPoolPermissions,
    #[msg("Unauthorized")]
    Unauthorized,
    #[msg("Invalid inputs")]
    InvalidInputs,
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
}

// validate_lock_or_burn checks for correctness on inputs
// - token & pool is correct for chain
// - rate limiting
pub fn validate_lock_or_burn(
    lock_or_burn_in: &LockOrBurnInV1,
    config_mint: Pubkey,
    outbound_rate_limit: &mut RateLimitTokenBucket,
    allow_list: &[Pubkey],
) -> Result<()> {
    // validate token matches configured pool token
    require!(
        config_mint == lock_or_burn_in.local_token,
        CcipTokenPoolError::InvalidToken
    );

    require!(
        allow_list.contains(&lock_or_burn_in.original_sender),
        CcipTokenPoolError::InvalidSender
    );

    outbound_rate_limit.consume(lock_or_burn_in.amount)
}

// validate_lock_or_burn checks for correctness on inputs
// - token & pool is correct for chain
// - rate limiting
pub fn validate_release_or_mint(
    release_or_mint_in: &ReleaseOrMintInV1,
    parsed_amount: u64,
    config_mint: Pubkey,
    pool_addresses: &[RemoteAddress],
    inbound_rate_limit: &mut RateLimitTokenBucket,
) -> Result<()> {
    require_eq!(
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

    inbound_rate_limit.consume(parsed_amount)
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
