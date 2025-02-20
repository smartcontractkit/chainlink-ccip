use anchor_lang::prelude::*;

use crate::{instructions::interfaces::Public, Subject};

pub struct Impl;
impl Public for Impl {
    fn verify_uncursed_globally<'info>(&self, ctx: Context<crate::VerifyCurse>) -> Result<()> {
        todo!()
    }

    fn verify_uncursed_subject<'info>(&self, subject: Subject) -> Result<()> {
        todo!()
    }
}
