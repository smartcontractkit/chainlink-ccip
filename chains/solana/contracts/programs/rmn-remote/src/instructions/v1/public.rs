use anchor_lang::prelude::*;

use crate::{instructions::interfaces::Public, CurseSubject, InspectCurses, RmnRemoteError};

pub struct Impl;

impl Public for Impl {
    fn verify_not_cursed<'info>(
        &self,
        ctx: Context<InspectCurses>,
        subject: CurseSubject,
    ) -> Result<()> {
        let curses = &ctx.accounts.config_and_curses;
        require!(
            !curses.is_chain_globally_cursed(),
            RmnRemoteError::GloballyCursed
        );
        require!(
            !curses.is_subject_cursed(subject),
            RmnRemoteError::SubjectCursed
        );
        Ok(())
    }

    fn get_cursed_subjects(&self, ctx: Context<InspectCurses>) -> Result<Vec<CurseSubject>> {
        Ok(ctx.accounts.config_and_curses.cursed_subjects.clone())
    }
}
