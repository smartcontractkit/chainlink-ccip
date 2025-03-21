use anchor_lang::prelude::*;
use rmn_remote::state::CurseSubject;

pub fn verify_uncursed_cpi<'info>(
    rmn_remote: AccountInfo<'info>,
    rmn_remote_config: AccountInfo<'info>,
    rmn_remote_curses: AccountInfo<'info>,
    chain_selector: u64,
) -> Result<()> {
    let cpi_accounts = rmn_remote::cpi::accounts::InspectCurses {
        config: rmn_remote_config,
        curses: rmn_remote_curses,
    };
    let cpi_context = CpiContext::new(rmn_remote, cpi_accounts);
    rmn_remote::cpi::verify_not_cursed(
        cpi_context,
        CurseSubject::from_chain_selector(chain_selector),
    )
}
