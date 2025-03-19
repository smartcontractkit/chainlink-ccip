use anchor_lang::prelude::*;
pub mod common;
pub mod rate_limiter;

// null contract required for IDL + gobinding generation
declare_id!("31FHyhZYz2sUc1iA5EnijN2L7bUyMMnvjyX7CiPQoDQq");
#[program]
pub mod base_token_pool {}
