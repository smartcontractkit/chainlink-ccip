use anchor_lang::prelude::*;
use anchor_lang::system_program;

use crate::config::{load_config_v2_unchecked, ConfigV2};
use crate::context::{MigrateConfigV2ToV3, ANCHOR_DISCRIMINATOR};
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

    fn migrate_config_v2_to_v3(&self, ctx: Context<MigrateConfigV2ToV3>) -> Result<()> {
        let required_v3_space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE;
        let minimum_balance = Rent::get()?.minimum_balance(required_v3_space);

        let account_info = &ctx.accounts.config.to_account_info();

        require_eq!(
            account_info.data_len(),
            ANCHOR_DISCRIMINATOR + ConfigV2::INIT_SPACE, // it's v2 space
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
        account_info.realloc(required_v3_space, false)?;

        // Set the new values
        msg!("Loading config V2...");
        let config_v2 = load_config_v2_unchecked(account_info.try_borrow_data()?)?;
        msg!("Read config V2: {:?}", config_v2);
        require_eq!(
            config_v2.version,
            2, // confirm it was v2
            RmnRemoteError::InvalidInputsConfigAccount
        );

        let new_config = Config {
            version: 3,
            owner: config_v2.owner,
            proposed_owner: config_v2.proposed_owner,
            default_code_version: config_v2.default_code_version,
            event_authorities: config_v2.event_authorities,

            // new v3 fields
            curser: config_v2.owner, // initialize to same value as owner, so there is no downtime on cursing
        };

        // Write back to permanent state
        msg!("Writing migrated RMNRemote Config to account...");
        new_config.try_serialize(&mut &mut ctx.accounts.config.try_borrow_mut_data()?[..])?;

        Ok(())
    }
}

fn is_subject_cursed(curses: &Curses, subject: CurseSubject) -> bool {
    curses.cursed_subjects.contains(&subject)
}

fn is_chain_globally_cursed(curses: &Curses) -> bool {
    curses.cursed_subjects.contains(&CurseSubject::GLOBAL)
}
