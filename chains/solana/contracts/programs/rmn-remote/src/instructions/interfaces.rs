use anchor_lang::prelude::*;

use crate::{state::CodeVersion, AcceptOwnership, UpdateConfig};

pub trait Public {
    fn verify_uncursed_globally<'info>(&self) -> Result<()>;
    fn verify_uncursed_subject<'info>(&self, s: u64) -> Result<()>;
}

pub trait Admin {
    fn transfer_ownership(&self, ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()>;
}
