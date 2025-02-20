use anchor_lang::prelude::*;

declare_id!("CPkyVFQmyzmb6HfDE5TQr3NmZAsydcYspBoc3bf6Zo5x");

pub mod context;
use context::*;

pub mod state;
use state::*;

pub mod event;
use event::*;

mod instructions;

#[program]
pub mod rmn_remote {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        ctx.accounts.config.set_inner(Config {
            owner: ctx.accounts.authority.key(),
            version: 1,
            proposed_owner: Pubkey::default(),
            default_code_version: CodeVersion::V1,
        });

        // TODO emit event
        Ok(())
    }
}

#[error_code]
pub enum RmnRemoteError {
    #[msg("The signer is unauthorized")]
    // offset error code so that they don't clash with other programs
    // (Anchor's base custom error code 6000 + offset 3000 = start at 9000)
    Unauthorized = 3000,
    #[msg("Invalid inputs")]
    InvalidInputs,
    #[msg("Proposed owner is the current owner")]
    RedundantOwnerProposal,
    #[msg("Invalid version of the onchain state")]
    InvalidVersion,
    #[msg("Invalid code version")]
    InvalidCodeVersion,
}
