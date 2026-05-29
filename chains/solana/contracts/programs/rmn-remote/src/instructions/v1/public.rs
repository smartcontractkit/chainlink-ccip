use std::cell::Ref;

use anchor_lang::prelude::*;
use anchor_lang::system_program;

use crate::config::{load_old_config, ConfigV2};
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
        let account_info = &ctx.accounts.config.to_account_info();
        require_eq!(
            account_info.data_len(),
            ANCHOR_DISCRIMINATOR + ConfigV2::INIT_SPACE, // it's v2 space
            RmnRemoteError::InvalidInputsConfigAccount
        );

        let new_config = old_to_new(account_info.try_borrow_data()?, ctx.bumps.config)?;
        msg!("Computed new config: {:?}", new_config);

        let required_space =
            ANCHOR_DISCRIMINATOR + Config::dynamic_len(new_config.event_authorities.len());
        let minimum_balance = Rent::get()?.minimum_balance(required_space);

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
        account_info.realloc(required_space, false)?;

        // Write back to permanent state
        msg!("Writing migrated RMNRemote Config to account...");
        new_config.try_serialize(&mut &mut ctx.accounts.config.try_borrow_mut_data()?[..])?;

        Ok(())
    }
}

fn old_to_new(bytes: Ref<&mut [u8]>, bump: u8) -> Result<Config> {
    msg!("Loading config V2...");
    let config_v2 = load_old_config::<ConfigV2>(bytes)?;
    msg!("Read config V2: {:?}", config_v2);

    require_eq!(
        config_v2.version,
        2, // confirm it was v2
        RmnRemoteError::InvalidInputsConfigAccount
    );

    let new_config = Config {
        version: Config::LATEST_VERSION,
        owner: config_v2.owner,
        proposed_owner: config_v2.proposed_owner,
        default_code_version: config_v2.default_code_version,
        event_authorities: config_v2.event_authorities,

        // new v3 fields
        curser: config_v2.owner, // initialize to same value as owner
        bump,
    };
    Ok(new_config)
}

fn is_subject_cursed(curses: &Curses, subject: CurseSubject) -> bool {
    curses.cursed_subjects.contains(&subject)
}

fn is_chain_globally_cursed(curses: &Curses) -> bool {
    curses.cursed_subjects.contains(&CurseSubject::GLOBAL)
}

#[cfg(test)]
mod tests {
    use crate::state::CodeVersion;
    use anchor_lang::Discriminator;
    use std::cell::RefCell;

    use super::*;

    #[test]
    fn test_v2_to_v3() {
        let owner = Pubkey::new_unique();
        let proposed_owner = Pubkey::new_unique();
        let event_authority = Pubkey::new_unique();

        let old = ConfigV2 {
            version: 2,
            owner,
            proposed_owner,
            default_code_version: CodeVersion::V1,
            event_authorities: vec![event_authority],
        };

        // Prepend the discriminator: load_old_config_unchecked expects [discriminator | data]
        let mut old_bytes = Config::DISCRIMINATOR.to_vec();
        old_bytes.extend_from_slice(&old.try_to_vec().unwrap());

        // Wrap in RefCell to obtain a Ref<&mut [u8]> (same type as AccountInfo::try_borrow_data)
        let rc: RefCell<&mut [u8]> = RefCell::new(&mut old_bytes);
        let data = rc.borrow();
        let bump = 255;
        let migrated = old_to_new(data, bump).unwrap();

        assert_eq!(
            migrated,
            Config {
                version: 3, // hardcode expected version, so that if LATEST_VERSION is updated without updating this test, the test will fail and alert us to update the migration logic as well
                owner,
                proposed_owner,
                default_code_version: CodeVersion::V1,
                event_authorities: vec![event_authority],
                curser: owner, // curser initialized to same value as owner
                bump,
            }
        );
    }
}
