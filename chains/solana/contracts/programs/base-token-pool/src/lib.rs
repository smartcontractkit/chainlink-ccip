use anchor_lang::prelude::*;
pub mod common;
pub mod rate_limiter;

// null contract required for IDL + gobinding generation
declare_id!("9vMcoeKAhpSZw1F8PQx8fhiQy1cmLyehomrZFPniPLc1");
#[program]
pub mod base_token_pool {}
