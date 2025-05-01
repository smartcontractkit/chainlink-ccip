use crate::state::BufferedReport;
use anchor_lang::prelude::*;
use ccip_common::seed;

pub const ANCHOR_DISCRIMINATOR: usize = 8; // size in bytes

#[derive(Accounts)]
#[instruction(buffer_id: u64, data: Vec<u8>)]
pub struct AppendExecutionReportData<'info> {
    #[account(
        mut,
        seeds = [seed::EXECUTION_BUFFER, authority.key().as_ref(), &buffer_id.to_le_bytes()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + BufferedReport::INIT_SPACE + buffered_report.raw_report_data.len() + data.len(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub buffered_report: Account<'info, BufferedReport>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(buffer_id: u64)]
pub struct InitializeExecutionReportBuffer<'info> {
    #[account(
        init,
        seeds = [seed::EXECUTION_BUFFER, authority.key().as_ref(), &buffer_id.to_le_bytes()],
        bump,
        space = ANCHOR_DISCRIMINATOR + BufferedReport::INIT_SPACE,
        payer = authority,
    )]
    pub buffered_report: Account<'info, BufferedReport>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(buffer_id: u64)]
pub struct CloseBuffer<'info> {
    #[account(
        mut,
        seeds = [seed::EXECUTION_BUFFER, authority.key().as_ref(), &buffer_id.to_le_bytes()],
        bump,
        close = authority,
    )]
    pub buffered_report: Account<'info, BufferedReport>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(buffer_id: u64, _token_indices: Vec<u8>)]
pub struct ExecuteContext<'info> {
    #[account(
        mut,
        seeds = [seed::EXECUTION_BUFFER, authority.key().as_ref(), &buffer_id.to_le_bytes()],
        bump,
        close = authority,
    )]
    pub buffered_report: Account<'info, BufferedReport>,
    // ------------------------
    // Accounts for offramp CPI: All validations are done by the offramp
    /// CHECK: validated during CPI
    pub config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub reference_addresses: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub source_chain: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    #[account(mut)]
    pub commit_report: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub offramp: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub allowed_offramp: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub rmn_remote: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub rmn_remote_curses: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub rmn_remote_config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub sysvar_instructions: UncheckedAccount<'info>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
    // remaining accounts
    // [receiver_program, external_execution_signer, receiver_account, ...user specified accounts from message data for arbitrary messaging]
    // +
    // [
    // ccip_offramp_pools_signer - derivable PDA [seed::EXTERNAL_TOKEN_POOL, pool_program], seeds::program=offramp (not in lookup table)
    // user/sender token account (must be associated token account - derivable PDA [wallet_addr, token_program, mint])
    // per chain per token config (ccip: billing, ccip admin controlled - derivable PDA [chain_selector, mint])
    // pool chain config (pool: custom configs that may include rate limits & remote chain configs, pool admin controlled - derivable [chain_selector, mint])
    // token pool lookup table
    // token registry PDA
    // pool program
    // pool config
    // pool token account (must be associated token account - derivable PDA [wallet_addr, token_program, mint])
    // pool signer
    // token program
    // token mint
    // ccip_router_pools_signer - derivable PDA [seed::EXTERNAL_TOKEN_POOL, pool_program], seeds::program=router (present in lookup table)
    // ...additional accounts for pool config
    // ] x N tokens
}
