use anchor_lang::prelude::*;

use crate::{instructions::interfaces::Public, CurseSubject, InspectCurses, RmnRemoteError};

pub struct Impl;

impl Public for Impl {
    fn verify_not_cursed<'info>(
        &self,
        ctx: Context<InspectCurses>,
        subject: CurseSubject,
    ) -> Result<()> {
        let cursed_subjects = &ctx.accounts.cursed.subjects;
        let local_chain_selector_subject =
            CurseSubject::from_chain_selector(ctx.accounts.config.local_chain_selector);

        require!(
            !cursed_subjects.iter().any(|cs| *cs == subject),
            RmnRemoteError::SubjectCursed
        );
        require!(
            !cursed_subjects
                .iter()
                .any(|cs| *cs == local_chain_selector_subject),
            RmnRemoteError::GloballyCursed
        );
        Ok(())
    }

    fn get_cursed_subjects(&self, ctx: Context<InspectCurses>) -> Result<Vec<CurseSubject>> {
        Ok(ctx.accounts.cursed.subjects.clone())
    }
}
