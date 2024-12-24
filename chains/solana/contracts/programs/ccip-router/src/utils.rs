use std::ops::{Add, AddAssign, Mul};

use anchor_lang::prelude::*;
use ethnum::U256;

use crate::{CcipRouterError, SolanaTokenAmount};

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
pub struct Usd18Decimals(pub U256);

impl Usd18Decimals {
    pub const ZERO: Self = Self(U256::ZERO);

    pub fn from_usd_cents(cents: u32) -> Self {
        Self(U256::new(cents.into()) * 1u32.e(16))
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

impl Mul<U256> for Usd18Decimals {
    type Output = Self;

    fn mul(mut self, rhs: U256) -> Self::Output {
        self.0 *= rhs;
        self
    }
}

impl SolanaTokenAmount {
    pub fn value(&self, price: &Usd18Decimals) -> Usd18Decimals {
        Usd18Decimals((U256::new(self.amount.into()) * price.0) / 1u32.e(18))
    }

    pub fn amount(token: Pubkey, value: Usd18Decimals, price: Usd18Decimals) -> Result<Self> {
        Ok(Self {
            token,
            amount: ((value.0 * 1u32.e(18)) / price.0)
                .try_into()
                .map_err(|_| CcipRouterError::InvalidTokenPrice)?,
        })
    }
}
