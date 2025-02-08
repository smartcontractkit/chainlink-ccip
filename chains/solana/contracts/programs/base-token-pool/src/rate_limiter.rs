use std::cmp::min;

use anchor_lang::prelude::*;
use solana_program::log::sol_log;

use crate::common::CcipTokenPoolError;

// NOTE: spl-token or spl-token-2022 use u64 for amounts

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RateLimitTokenBucket {
    pub tokens: u64, // Current number of tokens that are in the bucket. Represents how many tokens a single transaction can consume in a single transaction - different than total number of tokens within a pool
    pub last_updated: u64, // Timestamp in seconds of the last token refill, good for 100+ years.
    pub cfg: RateLimitConfig,
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RateLimitConfig {
    pub enabled: bool, // Indication whether the rate limiting is enabled or not
    pub capacity: u64, // Maximum number of tokens that can be in the bucket.
    pub rate: u64,     // Number of tokens per second that the bucket is refilled.
}

impl RateLimitTokenBucket {
    // consume calculates the new amount of tokens in a bucket (up to the capacity)
    // and validates that the bucket has enough to fulfill the request
    pub fn consume(&mut self, request_tokens: u64) -> Result<()> {
        if !self.cfg.enabled || request_tokens == 0 {
            return Ok(());
        }

        let clock: Clock = Clock::get()?;
        let current_timestamp = clock.unix_timestamp as u64; // positive part of i64 will always fit in u64

        let mut tokens = self.tokens;
        let capacity = self.cfg.capacity;
        let time_diff = current_timestamp.checked_sub(self.last_updated).unwrap();

        if time_diff != 0 {
            // this should never happen - requires manual resetting of config if this occurs
            require!(tokens <= capacity, CcipTokenPoolError::RLBucketOverfilled);

            // Refill tokens when arriving at a new block time
            tokens = self.calculate_refill(capacity, tokens, time_diff, self.cfg.rate);

            self.last_updated = current_timestamp;
        }

        require!(
            capacity >= request_tokens,
            CcipTokenPoolError::RLMaxCapacityExceeded,
        );

        if tokens < request_tokens {
            let rate = self.cfg.rate;
            // Wait required until the bucket is refilled enough to accept this value, round up to next higher second
            // Consume is not guaranteed to succeed after wait time passes if there is competing traffic.
            // This acts as a lower bound of wait time.
            let min_wait_sec = (request_tokens.checked_sub(tokens).unwrap()
                + rate.checked_sub(1).unwrap())
            .checked_div(self.cfg.rate)
            .unwrap();

            // print human readable message before reverting with total wait time
            let msg = "min wait (s): ".to_owned() + &min_wait_sec.to_string();
            sol_log(&msg);
        }
        require!(
            tokens >= request_tokens,
            CcipTokenPoolError::RLRateLimitReached
        );

        self.tokens = tokens.checked_sub(request_tokens).unwrap();
        emit!(TokensConsumed {
            tokens: request_tokens
        });

        Ok(())
    }

    // set_token_bucket_config sets + validates a new config and updates the number of tokens in the bucket
    pub fn set_token_bucket_config(&mut self, config: RateLimitConfig) -> Result<()> {
        let clock: Clock = Clock::get()?;
        let current_timestamp = clock.unix_timestamp as u64; // positive part of i64 will always fit in u64
        let time_diff = current_timestamp.checked_sub(self.last_updated).unwrap();
        if time_diff != 0 {
            // Refill tokens when arriving at a new block time
            self.tokens =
                self.calculate_refill(self.cfg.capacity, self.tokens, time_diff, self.cfg.rate);

            self.last_updated = current_timestamp;
        }

        self.tokens = min(config.capacity, self.tokens);
        self.cfg = config.clone();

        emit!(ConfigChanged { config });

        Ok(())
    }

    // calculate_refill returns either the maximum capacity of the bucket or the current + refill amount
    fn calculate_refill(&self, capacity: u64, tokens: u64, time_diff: u64, rate: u64) -> u64 {
        min(capacity, tokens + time_diff * rate)
    }
}

pub fn validate_token_bucket_config(cfg: &RateLimitConfig) -> Result<()> {
    if cfg.enabled {
        require!(
            cfg.rate < cfg.capacity && cfg.rate != 0,
            CcipTokenPoolError::RLInvalidRateLimitRate
        );
    } else {
        require!(
            cfg.rate == 0 && cfg.capacity == 0,
            CcipTokenPoolError::RLDisabledNonZeroRateLimit
        );
    }
    Ok(())
}

#[event]
pub struct TokensConsumed {
    pub tokens: u64,
}

#[event]
pub struct ConfigChanged {
    pub config: RateLimitConfig,
}
