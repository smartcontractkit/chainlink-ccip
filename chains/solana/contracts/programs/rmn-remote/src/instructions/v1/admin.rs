use anchor_lang::prelude::*;

use crate::{
    instructions::interfaces::Admin, AcceptOwnership, CodeVersion, ConfigSet, Curse, CurseSubject,
    OwnershipTransferRequested, OwnershipTransferred, RmnRemoteError, SubjectCursed,
    SubjectUncursed, Uncurse, UpdateConfig,
};

pub struct Impl;
impl Admin for Impl {
    fn transfer_ownership(&self, ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(
            proposed_owner != config.owner,
            RmnRemoteError::RedundantOwnerProposal
        );
        emit!(OwnershipTransferRequested {
            from: config.owner,
            to: proposed_owner,
        });
        ctx.accounts.config.proposed_owner = proposed_owner;
        Ok(())
    }

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()> {
        let config = &mut ctx.accounts.config;
        emit!(OwnershipTransferred {
            from: config.owner,
            to: config.proposed_owner,
        });
        // NOTE: take() resets proposed_owner to default
        ctx.accounts.config.owner = std::mem::take(&mut ctx.accounts.config.proposed_owner);
        Ok(())
    }

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()> {
        require_neq!(
            code_version,
            CodeVersion::Default,
            RmnRemoteError::InvalidCodeVersion
        );
        let config = &mut ctx.accounts.config;
        config.default_code_version = code_version;

        emit!(ConfigSet {
            default_code_version: config.default_code_version,
        });
        Ok(())
    }

    fn curse(&self, ctx: Context<Curse>, subject: CurseSubject) -> Result<()> {
        let curses = &mut ctx.accounts.curses;

        require!(
            !curses.cursed_subjects.contains(&subject),
            RmnRemoteError::SubjectIsAlreadyCursed
        );

        curses.cursed_subjects.push(subject);
        emit!(SubjectCursed { subject });
        Ok(())
    }

    fn uncurse(&self, ctx: Context<Uncurse>, subject: CurseSubject) -> Result<()> {
        let curses = &mut ctx.accounts.curses;

        require!(
            curses.cursed_subjects.contains(&subject),
            RmnRemoteError::SubjectWasNotCursed
        );

        curses.cursed_subjects.retain(|c| c != &subject);
        emit!(SubjectUncursed { subject });
        Ok(())
    }
}
