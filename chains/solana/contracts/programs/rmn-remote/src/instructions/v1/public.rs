use anchor_lang::prelude::*;
use anchor_lang::system_program;

use crate::config::load_config_v1_unchecked;
use crate::config::ConfigV1;
use crate::context::{MigrateConfigV1ToV2, ANCHOR_DISCRIMINATOR};
use crate::instructions::interfaces::Public;
use crate::state::Config;
use crate::{CurseSubject, Curses, InspectCurses, RmnRemoteError};

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

    fn migrate_config_v1_to_v2(&self, ctx: Context<MigrateConfigV1ToV2>) -> Result<()> {
        let required_v2_space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE;
        let minimum_balance = Rent::get()?.minimum_balance(required_v2_space);

        let account_info = &ctx.accounts.config.to_account_info();

        require_eq!(
            account_info.data_len(),
            ANCHOR_DISCRIMINATOR + ConfigV1::INIT_SPACE, // it's v1 space
            RmnRemoteError::InvalidInputsConfigAccount
        );

        // Extend the account
        msg!("Extending RMNRemote Config account...");
        let current_lamports = account_info.lamports();
        if current_lamports < minimum_balance {
            system_program::transfer(
                CpiContext::new(
                    ctx.accounts.system_program.to_account_info(),
                    system_program::Transfer {
                        from: ctx.accounts.authority.to_account_info(),
                        to: account_info.clone(),
                    },
                ),
                minimum_balance.checked_sub(current_lamports).unwrap(),
            )?;
        }
        account_info.realloc(required_v2_space, false)?;

        // Set the new values
        msg!("Loading config V1...");
        let config_v1 = load_config_v1_unchecked(account_info.try_borrow_data()?)?;
        msg!("Read config V1: {:?}", config_v1);
        require_eq!(
            config_v1.version,
            1, // confirm it was v1
            RmnRemoteError::InvalidInputsConfigAccount
        );

        let mut config: Config = config_v1.try_into()?;

        config.version = 2; // migrate to version 2
        config.event_authorities = vec![]; // backwards-compatible default, so the migration can be permissionless

        // Write back to permanent state
        msg!("Writing migrated RMNRemote Config v2 to account...");
        config.try_serialize(&mut &mut ctx.accounts.config.try_borrow_mut_data()?[..])?;

        Ok(())
    }
}

fn is_subject_cursed(curses: &Curses, subject: CurseSubject) -> bool {
    curses.cursed_subjects.contains(&subject)
}

fn is_chain_globally_cursed(curses: &Curses) -> bool {
    curses.cursed_subjects.contains(&CurseSubject::GLOBAL)
}
