use std::cmp::min;

use anchor_lang::prelude::*;
use solana_program::log::sol_log;

use crate::common::CcipTokenPoolError;

// NOTE: spl-token or spl-token-2022 use u64 for amounts

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RateLimitTokenBucket {
    pub tokens: u64, // Current number of tokens that are in the bucket. Represents how many tokens a single transaction can consume in a single transaction - different than total number of tokens within a pool
    pub last_updated: u64, // Timestamp in seconds of the last token refill, good for 100+ years.
    cfg: RateLimitConfig,
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RateLimitConfig {
    pub enabled: bool, // Indication whether the rate limiting is enabled or not
    pub capacity: u64, // Maximum number of tokens that can be in the bucket.
    pub rate: u64,     // Number of tokens per second that the bucket is refilled.
}

pub trait Timestamper: Sized {
    fn unix_timestamp(&self) -> u64;
    fn instance() -> std::result::Result<Self, ProgramError>;
}

impl Timestamper for Clock {
    fn unix_timestamp(&self) -> u64 {
        self.unix_timestamp as u64
    }

    fn instance() -> std::result::Result<Self, ProgramError> {
        Self::get()
    }
}

impl RateLimitTokenBucket {
    // consume calculates the new amount of tokens in a bucket (up to the capacity)
    // and validates that the bucket has enough to fulfill the request
    pub fn consume<C: Timestamper>(&mut self, request_tokens: u64) -> Result<()> {
        if !self.cfg.enabled || request_tokens == 0 {
            return Ok(());
        }

        let capacity = self.cfg.capacity;
        require!(
            capacity >= request_tokens,
            CcipTokenPoolError::RLMaxCapacityExceeded,
        );

        self.refill::<C>()?;

        if self.tokens < request_tokens {
            let rate = self.cfg.rate;
            // Wait required until the bucket is refilled enough to accept this value, round up to next higher second
            // Consume is not guaranteed to succeed after wait time passes if there is competing traffic.
            // This acts as a lower bound of wait time.
            let min_wait_sec = (request_tokens
                .checked_sub(self.tokens)
                .unwrap()
                .checked_add(rate.checked_sub(1).unwrap()))
            .unwrap()
            .checked_div(rate)
            .unwrap();

            // print human readable message before reverting with total wait time
            let msg = "min wait (s): ".to_owned() + &min_wait_sec.to_string();
            sol_log(&msg);

            return Err(CcipTokenPoolError::RLRateLimitReached.into());
        }

        self.tokens = self.tokens.checked_sub(request_tokens).unwrap();
        emit!(TokensConsumed {
            tokens: request_tokens
        });

        Ok(())
    }

    fn refill<C: Timestamper>(&mut self) -> Result<()> {
        let current_timestamp = C::instance()?.unix_timestamp();
        let time_diff = current_timestamp.checked_sub(self.last_updated).unwrap();
        let increase = time_diff.checked_mul(self.cfg.rate).unwrap();
        let refill = self.tokens.checked_add(increase).unwrap();

        self.tokens = min(self.cfg.capacity, refill);
        self.last_updated = current_timestamp;
        Ok(())
    }

    // set_token_bucket_config sets + validates a new config and updates the number of tokens in the bucket
    pub fn set_token_bucket_config(&mut self, config: RateLimitConfig) -> Result<()> {
        validate_token_bucket_config(&config)?;
        self.cfg = config.clone();
        self.refill::<Clock>()?;

        emit!(ConfigChanged { config });

        Ok(())
    }
}

fn validate_token_bucket_config(cfg: &RateLimitConfig) -> Result<()> {
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

#[cfg(test)]
mod tests {
    use std::cell::RefCell;

    use super::*;

    #[derive(Default)]
    struct MockTimestamper;

    // Persistent state per test thread.
    thread_local! {
        static TIMESTAMP: RefCell<u64> = Default::default();
    }

    impl Timestamper for MockTimestamper {
        fn unix_timestamp(&self) -> u64 {
            TIMESTAMP.with(|ts| *ts.borrow())
        }

        fn instance() -> std::result::Result<Self, ProgramError> {
            Ok(Self)
        }
    }

    fn elapse(seconds: u64) {
        TIMESTAMP.with(|ts| *ts.borrow_mut() += seconds)
    }

    #[test]
    fn bucket_grows_even_during_failing_attempts() {
        let mut bucket = RateLimitTokenBucket {
            tokens: 0,
            last_updated: 0,
            cfg: RateLimitConfig {
                enabled: true,
                capacity: 50,
                // Refills one token per second
                rate: 1,
            },
        };

        for _ in 0..5 {
            assert_eq!(
                bucket.consume::<MockTimestamper>(5),
                Err(CcipTokenPoolError::RLRateLimitReached.into())
            );
            elapse(1);
        }
        // Finally enough time has passed and the bucket is full.
        bucket.consume::<MockTimestamper>(5).unwrap();
    }
}
