use anchor_lang::prelude::*;
pub mod common;
pub mod rate_limiter;

// null contract required for IDL + gobinding generation
declare_id!("TokenPooL1111111111111111111111111111111111");
#[program]
pub mod base_token_pool {}
