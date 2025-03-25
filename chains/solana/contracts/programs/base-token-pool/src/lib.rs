use anchor_lang::prelude::*;
pub mod common;
pub mod rate_limiter;

// null contract required for IDL + gobinding generation
declare_id!("6REgveBio1N4kcy1NYMJusPVadMZwkfrezzbzr5xhPqb");
#[program]
pub mod base_token_pool {}
