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
            require_neq!(rate, 0, CcipTokenPoolError::RLRateLimitReached);
            // Wait required until the bucket is refilled enough to accept this value, round up to next higher second
            // Consume is not guaranteed to succeed after wait time passes if there is competing traffic.
            // This acts as a lower bound of wait time.
            let min_wait_sec = request_tokens
                .saturating_sub(self.tokens)
                .saturating_add(rate.saturating_sub(1))
                .checked_div(rate)
                .ok_or(CcipTokenPoolError::RLRateLimitReached)?;

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
        // Saturating math: a refill above u64::MAX is capped at capacity below anyway.
        let time_diff = current_timestamp.saturating_sub(self.last_updated);
        let increase = time_diff.saturating_mul(self.cfg.rate);
        let refill = self.tokens.saturating_add(increase);

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
            cfg.rate <= cfg.capacity, // allow setting rate and capacity to 0
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

    fn now() -> u64 {
        TIMESTAMP.with(|ts| *ts.borrow())
    }

    fn bucket(tokens: u64, last_updated: u64, capacity: u64, rate: u64) -> RateLimitTokenBucket {
        RateLimitTokenBucket {
            tokens,
            last_updated,
            cfg: RateLimitConfig {
                enabled: true,
                capacity,
                rate,
            },
        }
    }

    #[test]
    fn refill_after_long_idle_fills_to_capacity() {
        elapse(1_800_000_000);
        // 9 decimals, 5,000 tokens/s, 1M capacity: overflows u64 after ~42.7 idle days.
        let capacity = 1_000_000 * 1_000_000_000;
        let rate = 5_000 * 1_000_000_000;
        let mut b = bucket(capacity, now(), capacity, rate);

        elapse(43 * 24 * 60 * 60);
        b.consume::<MockTimestamper>(1).unwrap();
        assert_eq!(b.tokens, capacity - 1);
        assert_eq!(b.last_updated, now());
    }

    #[test]
    fn first_refill_of_new_bucket_fills_to_capacity() {
        elapse(1_800_000_000);
        // A new bucket has last_updated = 0, so time_diff is the whole unix epoch.
        let capacity = u64::MAX / 2;
        let rate = 1_000_000_000_000;
        let mut b = bucket(0, 0, capacity, rate);

        b.consume::<MockTimestamper>(capacity).unwrap();
        assert_eq!(b.tokens, 0);
        assert_eq!(b.last_updated, now());
    }

    #[test]
    fn refill_with_max_rate_fills_to_capacity() {
        let mut b = bucket(0, now(), u64::MAX, u64::MAX);

        elapse(2);
        b.consume::<MockTimestamper>(1).unwrap();
        assert_eq!(b.tokens, u64::MAX - 1);
    }

    #[test]
    fn refill_near_max_tokens_fills_to_capacity() {
        let mut b = bucket(u64::MAX - 1, now(), u64::MAX, 10);

        elapse(10);
        b.consume::<MockTimestamper>(1).unwrap();
        assert_eq!(b.tokens, u64::MAX - 1);
    }

    #[test]
    fn refill_with_last_updated_in_future_does_not_refill() {
        elapse(1_000);
        let mut b = bucket(0, now() + 100, 100, 1);

        assert_eq!(
            b.consume::<MockTimestamper>(1),
            Err(CcipTokenPoolError::RLRateLimitReached.into())
        );
        assert_eq!(b.tokens, 0);
        assert_eq!(b.last_updated, now());
    }

    #[test]
    fn rate_limit_reached_with_zero_rate() {
        let mut b = bucket(0, now(), 10, 0);

        elapse(10);
        assert_eq!(
            b.consume::<MockTimestamper>(1),
            Err(CcipTokenPoolError::RLRateLimitReached.into())
        );
        assert_eq!(b.tokens, 0);
    }

    #[test]
    fn rate_limit_reached_with_max_values_does_not_overflow_wait_time() {
        // No time elapses, so the bucket stays empty and the wait time is computed.
        let mut b = bucket(0, now(), u64::MAX, u64::MAX);

        assert_eq!(
            b.consume::<MockTimestamper>(u64::MAX),
            Err(CcipTokenPoolError::RLRateLimitReached.into())
        );
        assert_eq!(b.tokens, 0);
    }
}
