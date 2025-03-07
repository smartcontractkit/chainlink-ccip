use anchor_lang::prelude::*;

use crate::{
    state::CodeVersion, AcceptOwnership, Curse, CurseSubject, InspectCurses, Uncurse, UpdateConfig,
};

pub trait Public {
    fn verify_not_cursed(&self, ctx: Context<InspectCurses>, subject: CurseSubject) -> Result<()>;
}

pub trait Admin {
    fn transfer_ownership(&self, ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()>;

    fn curse(&self, ctx: Context<Curse>, subject: CurseSubject) -> Result<()>;
    fn uncurse(&self, ctx: Context<Uncurse>, subject: CurseSubject) -> Result<()>;
}
