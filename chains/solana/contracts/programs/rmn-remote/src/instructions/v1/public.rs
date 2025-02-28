use anchor_lang::prelude::*;

use crate::{
    instructions::interfaces::Public, CurseSubject, Curses, InspectCurses, RmnRemoteError,
};

pub struct Impl;

impl Public for Impl {
    fn verify_not_cursed<'info>(
        &self,
        ctx: Context<InspectCurses>,
        subject: CurseSubject,
    ) -> Result<()> {
        let curses = &ctx.accounts.curses;
        require!(
            !is_chain_globally_cursed(curses),
            RmnRemoteError::GloballyCursed
        );
        require!(
            !is_subject_cursed(curses, subject),
            RmnRemoteError::SubjectCursed
        );
        Ok(())
    }
}

fn is_subject_cursed(curses: &Curses, subject: CurseSubject) -> bool {
    curses.cursed_subjects.contains(&subject)
}

fn is_chain_globally_cursed(curses: &Curses) -> bool {
    curses.cursed_subjects.contains(&CurseSubject::GLOBAL)
}
