use anchor_lang::prelude::*;

mod context;
use crate::context::*;

mod state;

declare_id!("buffGjr75PtEtV6D3pbKhaq59CR2qp4mwZDYGspggxZ");

#[program]
pub mod execution_buffer {
    use super::*;

    pub fn manually_execute_buffered<'info>(
        ctx: Context<'_, '_, 'info, 'info, ExecuteContext<'info>>,
        _buffer_id: u64,
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        let cpi_accounts = ccip_offramp::cpi::accounts::ExecuteReportContext {
            config: ctx.accounts.config.to_account_info(),
            reference_addresses: ctx.accounts.reference_addresses.to_account_info(),
            source_chain: ctx.accounts.source_chain.to_account_info(),
            commit_report: ctx.accounts.commit_report.to_account_info(),
            offramp: ctx.accounts.offramp.to_account_info(),
            allowed_offramp: ctx.accounts.allowed_offramp.to_account_info(),
            authority: ctx.accounts.authority.to_account_info(),
            system_program: ctx.accounts.system_program.to_account_info(),
            sysvar_instructions: ctx.accounts.sysvar_instructions.to_account_info(),
            rmn_remote: ctx.accounts.rmn_remote.to_account_info(),
            rmn_remote_curses: ctx.accounts.rmn_remote_curses.to_account_info(),
            rmn_remote_config: ctx.accounts.rmn_remote_config.to_account_info(),
        };
        let cpi_remaining_accounts = ctx.remaining_accounts.to_vec();
        let cpi_context = CpiContext::new(ctx.accounts.offramp.to_account_info(), cpi_accounts)
            .with_remaining_accounts(cpi_remaining_accounts);
        ccip_offramp::cpi::manually_execute(
            cpi_context,
            ctx.accounts.buffered_report.raw_report_data.clone(),
            token_indexes,
        )
    }

    pub fn append_execution_report_data<'info>(
        ctx: Context<'_, '_, 'info, 'info, AppendExecutionReportData>,
        _buffer_id: u64,
        data: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts
            .buffered_report
            .raw_report_data
            .extend_from_slice(&data);
        Ok(())
    }

    pub fn initialize_execution_report_buffer<'info>(
        _ctx: Context<'_, '_, 'info, 'info, InitializeExecutionReportBuffer>,
        _buffer_id: u64,
    ) -> Result<()> {
        Ok(())
    }

    pub fn close_buffer<'info>(
        _ctx: Context<'_, '_, 'info, 'info, CloseBuffer>,
        _buffer_id: u64,
    ) -> Result<()> {
        Ok(())
    }
}
