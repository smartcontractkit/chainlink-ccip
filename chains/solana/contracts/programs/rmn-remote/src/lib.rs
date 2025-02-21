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

    pub fn initialize(ctx: Context<Initialize>, local_chain_selector: u64) -> Result<()> {
        ctx.accounts.config.set_inner(Config {
            owner: ctx.accounts.authority.key(),
            version: 1,
            proposed_owner: Pubkey::default(),
            default_code_version: CodeVersion::V1,
            local_chain_selector,
        });

        emit!(ConfigSet {
            default_code_version: ctx.accounts.config.default_code_version,
            local_chain_selector
        });
        Ok(())
    }
}

#[error_code]
pub enum RmnRemoteError {
    #[msg("The signer is unauthorized")]
    // offset error code so that they don't clash with other programs
    // (Anchor's base custom error code 6000 + offset 3000 = start at 9000)
    Unauthorized = 3000,
    #[msg("Subject is already cursed")]
    SubjectIsAlreadyCursed,
    #[msg("Subject was not cursed")]
    SubjectWasNotCursed,
    #[msg("Proposed owner is the current owner")]
    RedundantOwnerProposal,
    #[msg("Invalid version of the onchain state")]
    InvalidVersion,
    #[msg("The subject is actively cursed")]
    SubjectCursed,
    #[msg("This chain is globally cursed")]
    GloballyCursed,
    #[msg("Invalid code version")]
    InvalidCodeVersion,
}
