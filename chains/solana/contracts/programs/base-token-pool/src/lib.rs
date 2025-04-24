use anchor_lang::prelude::*;
pub mod common;
pub mod rate_limiter;

// null contract required for IDL + gobinding generation
declare_id!("4YTxvoJZB5NUD4jtPPAttTtZfKLwfQmXzbQ6vFSXch5k");
#[program]
pub mod base_token_pool {}
