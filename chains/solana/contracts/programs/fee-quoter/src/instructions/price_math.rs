use anchor_lang::prelude::*;
use std::ops::{Add, AddAssign, Mul};

use ethnum::U256;

use crate::messages::SVMTokenAmount;
use crate::state::{BillingTokenConfig, TimestampedPackedU224};
use crate::FeeQuoterError;

pub trait Exponential {
    fn e(self, exponent: u8) -> U256;
}

impl<T: Into<u32>> Exponential for T {
    fn e(self, exponent: u8) -> U256 {
        U256::from(self.into()) * U256::new(10).pow(exponent as u32)
    }
}

// USD with 18 decimals (i.e. $8 -> 8e18)
#[derive(Debug, Clone, Default, PartialEq, Eq, PartialOrd, Ord)]
pub(super) struct Usd18Decimals(pub U256);

pub(super) fn get_validated_token_price(
    token_config: &BillingTokenConfig,
) -> Result<Usd18Decimals> {
    let timestamp = token_config.usd_per_token.timestamp;
    let price: Usd18Decimals = (&token_config.usd_per_token).into();

    // NOTE: There's no validation done with respect to token price staleness since data feeds are not
    // supported in solana. Only the existence of `any` timestamp is checked, to ensure the price
    // was set at least once.
    require!(
        price.0 != 0 && timestamp != 0,
        FeeQuoterError::InvalidTokenPrice
    );

    Ok(price)
}

impl Usd18Decimals {
    pub const ZERO: Self = Self(U256::ZERO);

    pub fn from_usd_cents(cents: u32) -> Self {
        Self(U256::new(cents.into()) * 1u32.e(16))
    }

    pub fn from_token_amount(sta: &SVMTokenAmount, price: &Usd18Decimals) -> Self {
        Usd18Decimals(U256::new(sta.amount.into()) * price.0 / 1u32.e(18))
    }
}

impl Add for Usd18Decimals {
    type Output = Self;

    fn add(self, rhs: Self) -> Self::Output {
        Self(self.0 + rhs.0)
    }
}

impl AddAssign for Usd18Decimals {
    fn add_assign(&mut self, rhs: Self) {
        self.0 += rhs.0
    }
}

impl<T: Into<U256>> Mul<T> for Usd18Decimals {
    type Output = Self;

    fn mul(mut self, rhs: T) -> Self::Output {
        self.0 *= rhs.into();
        self
    }
}

impl From<&TimestampedPackedU224> for Usd18Decimals {
    fn from(tpu: &TimestampedPackedU224) -> Self {
        let mut u256_buffer = [0u8; 32];
        u256_buffer[4..32].clone_from_slice(&tpu.value);
        Usd18Decimals(U256::from_be_bytes(u256_buffer))
    }
}
