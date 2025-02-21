use anchor_lang::prelude::*;

use crate::{program::RmnRemote, Config, CurseSubject, Cursed, RmnRemoteError};

/// Static space allocated to any account: must always be added to space calculations.
pub const ANCHOR_DISCRIMINATOR: usize = 8;

// valid_version validates that the passed in version is not 0 (uninitialized)
// and it is within the expected maximum supported version bounds
pub fn valid_version(v: u8, max_v: u8) -> bool {
    !uninitialized(v) && v <= max_v
}

pub fn uninitialized(v: u8) -> bool {
    v == 0
}

/// Maximum acceptable config version accepted by this module: any accounts with higher
/// version numbers than this will be rejected.
pub const MAX_CONFIG_V: u8 = 1;

pub mod seed {
    pub const CONFIG: &[u8] = b"config";
    pub const CURSES: &[u8] = b"curses";
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: Account<'info, Config>,

    #[account(
        init,
        seeds = [seed::CURSES],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Cursed::INIT_SPACE,
    )]
    pub curses: Account<'info, Cursed>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, RmnRemote>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ RmnRemoteError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ RmnRemoteError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    // validate signer is registered admin
    #[account(address = config.owner @ RmnRemoteError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ RmnRemoteError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    // validate signer is the new admin, accepting ownership of the contract
    #[account(address = config.proposed_owner @ RmnRemoteError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct Curse<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ RmnRemoteError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.proposed_owner @ RmnRemoteError::Unauthorized)]
    pub authority: Signer<'info>,

    #[account(
        mut,
        seeds = [seed::CURSES],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + curses.dynamic_len() + CurseSubject::INIT_SPACE,
        realloc::payer = authority,
        realloc::zero = false,
    )]
    pub curses: Account<'info, Cursed>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct Uncurse<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ RmnRemoteError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.proposed_owner @ RmnRemoteError::Unauthorized)]
    pub authority: Signer<'info>,

    #[account(
        mut,
        seeds = [seed::CURSES],
        bump,
        realloc = (ANCHOR_DISCRIMINATOR + curses.dynamic_len()).saturating_sub(CurseSubject::INIT_SPACE),
        realloc::payer = authority,
        realloc::zero = false,
    )]
    pub curses: Account<'info, Cursed>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct VerifyUncursed<'info> {
    #[account(
        seeds = [seed::CURSES],
        bump,
    )]
    pub cursed: Account<'info, Cursed>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ RmnRemoteError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
}

#[derive(Accounts)]
pub struct GetCursedSubjects<'info> {
    #[account(
        seeds = [seed::CURSES],
        bump,
    )]
    pub cursed: Account<'info, Cursed>,
}
