use anchor_lang::prelude::*;
pub mod common;
pub mod rate_limiter;

// null contract required for IDL + gobinding generation
declare_id!("tZhMpqaxb4znS4RpyzVprQySmXhFpEsrrLPMXy3RCMV");
#[program]
pub mod base_token_pool {}
