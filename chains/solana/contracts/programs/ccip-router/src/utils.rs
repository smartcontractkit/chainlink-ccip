use ethnum::U256;

pub trait Exponential {
    fn e(self, exponent: u8) -> U256;
}

impl<T: Into<u32>> Exponential for T {
    fn e(self, exponent: u8) -> U256 {
        U256::from(self.into()) * U256::new(10).pow(exponent as u32)
    }
}
