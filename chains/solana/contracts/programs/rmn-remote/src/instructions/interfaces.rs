use anchor_lang::prelude::*;

use crate::{
    state::CodeVersion, AcceptOwnership, CurseSubject, Subject, UncurseSubject, UpdateConfig,
    VerifyCurse,
};

pub trait Public {
    fn verify_uncursed<'info>(&self, ctx: Context<VerifyCurse>) -> Result<()>;
    fn verify_subject_uncursed<'info>(
        &self,
        ctx: Context<VerifyCurse>,
        subject: Subject,
    ) -> Result<()>;
}

pub trait Admin {
    fn transfer_ownership(&self, ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()>;

    fn curse(&self, ctx: Context<CurseSubject>) -> Result<()>;
    fn uncurse(&self, ctx: Context<CurseSubject>) -> Result<()>;
    fn curse_subject(&self, ctx: Context<CurseSubject>, subject: Subject) -> Result<()>;
    fn uncurse_subject(&self, ctx: Context<UncurseSubject>, subject: Subject) -> Result<()>;
}
