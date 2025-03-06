use anchor_lang::prelude::*;

use crate::{
    global_curse_subject, instructions::interfaces::Public, Curses, InspectCurses, RmnRemoteError,
};

pub struct Impl;

impl Public for Impl {
    fn verify_not_cursed<'info>(
        &self,
        ctx: Context<InspectCurses>,
        subject: Vec<u8>,
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

fn is_subject_cursed(curses: &Curses, subject: Vec<u8>) -> bool {
    curses.cursed_subjects.contains(&subject)
}

fn is_chain_globally_cursed(curses: &Curses) -> bool {
    curses.cursed_subjects.contains(&global_curse_subject())
}
