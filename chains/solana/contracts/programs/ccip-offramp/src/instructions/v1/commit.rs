use anchor_lang::prelude::*;
use anchor_spl::token::TokenAccount;
use ccip_common::seed;

use super::ocr3base::{ocr3_transmit, ReportContext, Signatures};
use super::ocr3impl::Ocr3ReportForCommit;

use crate::context::{CloseCommitReportAccount, CommitInput, CommitReportContext, OcrPluginType};
use crate::event::{CommitReportAccepted, CommitReportPDAClosed};
use crate::instructions::interfaces::Commit;
use crate::instructions::v1::rmn::verify_uncursed_cpi;
use crate::state::{CommitReport, GlobalState};
use crate::{CcipOfframpError, PriceOnlyCommitReportContext};

pub struct Impl;
impl Commit for Impl {
    fn commit<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, CommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 2],
        raw_report: Vec<u8>,
        rs: Vec<[u8; 32]>,
        ss: Vec<[u8; 32]>,
        raw_vs: [u8; 32],
    ) -> Result<()> {
        let report = CommitInput::deserialize(&mut raw_report.as_ref())
            .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?;

        // The Config and State for the Source Chain, containing if it is enabled, the on ramp address and the min sequence number expected for future messages
        let source_chain = &mut ctx.accounts.source_chain;

        verify_uncursed_cpi(
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            source_chain.chain_selector,
        )?;

        require!(
            report.merkle_root.is_some(),
            CcipOfframpError::MissingExpectedMerkleRoot
        );

        let report_context = ReportContext::from_byte_words(report_context_byte_words);

        // The Config Account stores the default values for the Router, the SVM Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
        let config = ctx.accounts.config.load()?;

        require!(
            source_chain.config.is_enabled,
            CcipOfframpError::UnsupportedSourceChainSelector
        );
        require!(
            source_chain.config.on_ramp.bytes()
                == report.merkle_root.as_ref().unwrap().on_ramp_address,
            CcipOfframpError::OnrampNotConfigured
        );

        // Check if the report contains price updates
        let empty_token_price_updates = report.price_updates.token_price_updates.is_empty();
        let empty_gas_price_updates = report.price_updates.gas_price_updates.is_empty();

        if empty_token_price_updates && empty_gas_price_updates {
            // If the report does not contain any price updates, then there is nothing to update.
            // Thus, as no price accounts have to be updated, the remaining accounts must be empty.
            require_eq!(
                ctx.remaining_accounts.len(),
                0,
                CcipOfframpError::InvalidInputsNumberOfAccounts
            );
        } else {
            let cpi_seeds = &[seed::FEE_BILLING_SIGNER, &[ctx.bumps.fee_billing_signer]];
            let cpi_signer = &[&cpi_seeds[..]];
            let cpi_program = ctx.accounts.fee_quoter.to_account_info();
            let cpi_accounts = fee_quoter::cpi::accounts::UpdatePrices {
                config: ctx.accounts.fee_quoter_config.to_account_info(),
                authority: ctx.accounts.fee_billing_signer.to_account_info(),
                allowed_price_updater: ctx
                    .accounts
                    .fee_quoter_allowed_price_updater
                    .to_account_info(),
            };

            helpers::update_prices(
                &report,
                &report_context,
                ctx.remaining_accounts,
                cpi_signer,
                cpi_program,
                cpi_accounts,
            )?;
        }

        // The Commit Report Account stores the information of 1 Commit Report:
        // - Merkle Root
        // - Timestamp of the Commit Report
        // - Interval of Messages: The min and max seq num of the messages in the Merkle Tree
        // - Execution State per each Message: 0 for Untouched, 1 for InProgress, 2 for Success and 3 for Failure
        let commit_report = &mut ctx.accounts.commit_report;
        let root = &report.merkle_root.as_ref().unwrap();

        require!(
            root.min_seq_nr <= root.max_seq_nr,
            CcipOfframpError::InvalidSequenceInterval
        );
        require!(
            root.max_seq_nr
                .to_owned()
                .checked_sub(root.min_seq_nr)
                .map_or_else(|| false, |seq_size| seq_size < 64),
            CcipOfframpError::InvalidSequenceInterval
        ); // As we have 64 slots to store the execution state
        require!(
            source_chain.state.min_seq_nr == root.min_seq_nr,
            CcipOfframpError::InvalidSequenceInterval
        );
        require!(root.merkle_root != [0; 32], CcipOfframpError::InvalidProof);
        require!(
            commit_report.timestamp == 0,
            CcipOfframpError::ExistingMerkleRoot
        );

        let next_seq_nr = root.max_seq_nr.checked_add(1);

        require!(
            next_seq_nr.is_some(),
            CcipOfframpError::ReachedMaxSequenceNumber
        );

        source_chain.state.min_seq_nr = next_seq_nr.unwrap();

        let clock: Clock = Clock::get()?;
        commit_report.version = 1;
        commit_report.chain_selector = root.source_chain_selector;
        commit_report.merkle_root = root.merkle_root;
        commit_report.timestamp = clock.unix_timestamp;
        commit_report.execution_states = 0;
        commit_report.min_msg_nr = root.min_seq_nr;
        commit_report.max_msg_nr = root.max_seq_nr;

        emit!(CommitReportAccepted {
            merkle_root: Some((*root).clone()),
            price_updates: report.price_updates.clone(),
        });

        ocr3_transmit(
            &config.ocr3[OcrPluginType::Commit as usize],
            &ctx.accounts.sysvar_instructions,
            ctx.accounts.authority.key(),
            OcrPluginType::Commit,
            report_context,
            &Ocr3ReportForCommit(&report),
            Signatures { rs, ss, raw_vs },
        )?;

        Ok(())
    }

    fn commit_price_only<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, PriceOnlyCommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 2],
        raw_report: Vec<u8>,
        rs: Vec<[u8; 32]>,
        ss: Vec<[u8; 32]>,
        raw_vs: [u8; 32],
    ) -> Result<()> {
        let report = CommitInput::deserialize(&mut raw_report.as_ref())
            .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?;

        verify_uncursed_cpi(
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            // No merkle root, so there's no remote chain selector to check.
            // We pass zero to verify there's no global curse.
            0,
        )?;

        require!(
            report.merkle_root.is_none(),
            CcipOfframpError::UnexpectedMerkleRoot,
        );
        let report_context = ReportContext::from_byte_words(report_context_byte_words);

        // The Config Account stores the default values for the Router, the SVM Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
        let config = ctx.accounts.config.load()?;

        // Check if the report contains price updates. It must, because this is a price-only commit
        require!(
            !report.price_updates.token_price_updates.is_empty()
                || !report.price_updates.gas_price_updates.is_empty(),
            CcipOfframpError::MissingExpectedPriceUpdates
        );

        let cpi_seeds = &[seed::FEE_BILLING_SIGNER, &[ctx.bumps.fee_billing_signer]];
        let cpi_signer = &[&cpi_seeds[..]];
        let cpi_program = ctx.accounts.fee_quoter.to_account_info();
        let cpi_accounts = fee_quoter::cpi::accounts::UpdatePrices {
            config: ctx.accounts.fee_quoter_config.to_account_info(),
            authority: ctx.accounts.fee_billing_signer.to_account_info(),
            allowed_price_updater: ctx
                .accounts
                .fee_quoter_allowed_price_updater
                .to_account_info(),
        };

        helpers::update_prices(
            &report,
            &report_context,
            ctx.remaining_accounts,
            cpi_signer,
            cpi_program,
            cpi_accounts,
        )?;

        emit!(CommitReportAccepted {
            merkle_root: None,
            price_updates: report.price_updates.clone(),
        });

        ocr3_transmit(
            &config.ocr3[OcrPluginType::Commit as usize],
            &ctx.accounts.sysvar_instructions,
            ctx.accounts.authority.key(),
            OcrPluginType::Commit,
            report_context,
            &Ocr3ReportForCommit(&report),
            Signatures { rs, ss, raw_vs },
        )?;

        Ok(())
    }

    fn close_commit_report_account(
        &self,
        ctx: Context<CloseCommitReportAccount>,
        source_chain_selector: u64,
        root: Vec<u8>,
    ) -> Result<()> {
        let commit_report = &ctx.accounts.commit_report;

        // Check if all messages have been executed (must be Success)
        require!(
            all_messages_executed(commit_report),
            CcipOfframpError::CommitReportHasPendingMessages
        );

        // Close the account and convert rent to wrapped SOL
        close_commit_pda(
            ctx.accounts.commit_report.to_account_info(),
            &ctx.accounts.fee_token_receiver,
        )?;

        let merkle_root_array: [u8; 32] = root
            .try_into()
            .map_err(|_| CcipOfframpError::InvalidProof)?;

        emit!(CommitReportPDAClosed {
            source_chain_selector,
            merkle_root: merkle_root_array,
        });

        Ok(())
    }
}

// Check if all messages have been executed
fn all_messages_executed(report: &CommitReport) -> bool {
    let num_messages = report.max_msg_nr.saturating_sub(report.min_msg_nr) + 1;

    // execution_states follow geometric series 2^1 + 2^3 + 2^5 + ... + 2 * 2^(2 * (num_messages - 1))
    // it can be converted to 2 * (4^0 + 4^1 + 4^2 + ... + 4^(num_messages - 1))
    // sum is calculated as 2 * (4^num_messages - 1) / 3
    let fully_executed = (4u128.pow(num_messages as u32) - 1) * 2 / 3;
    report.execution_states == fully_executed
}

// Close the commit report PDA and transfer the SOL to the fee token receiver
fn close_commit_pda<'info>(
    commit_report: AccountInfo<'info>,
    fee_token_receiver: &Account<'info, TokenAccount>,
) -> Result<()> {
    // Get lamports before closing
    let lamports = commit_report.lamports();

    // Close the commit report PDA by setting its balance to 0
    **commit_report.try_borrow_mut_lamports()? = 0;

    **fee_token_receiver
        .to_account_info()
        .try_borrow_mut_lamports()? += lamports;

    // Due to https://github.com/solana-labs/solana/issues/9711
    // we will not call SyncNative on the sol receiver here.
    // SyncNative will be called when ccipSend is called by user,
    // or it can be called at time of OnRamp fee withdrawal.

    Ok(())
}

mod helpers {
    use fee_quoter::cpi::accounts::UpdatePrices;

    use super::*;
    pub(super) fn update_prices<'info>(
        report: &CommitInput,
        report_context: &ReportContext,
        remaining_accounts: &'info [AccountInfo<'info>],
        cpi_signer: &[&[&[u8]]; 1],
        cpi_program: AccountInfo<'info>,
        cpi_accounts: UpdatePrices<'info>,
    ) -> Result<()> {
        // Remaining accounts represent:
        // - The state account to store the price sequence updates
        // - the accounts to update BillingTokenConfig for token prices
        // - the accounts to update DestChain for gas prices
        // They must be in order:
        // 1. state_account
        // 2. fee quoter token_accounts[]
        // 3. fee quoter gas_accounts[]
        // matching the order of the price updates in the CommitInput.
        // They must also all be writable so they can be updated.

        let minimum_remaining_accounts = 1
            + report.price_updates.token_price_updates.len()
            + report.price_updates.gas_price_updates.len();
        require_eq!(
            remaining_accounts.len(),
            minimum_remaining_accounts,
            CcipOfframpError::InvalidInputsNumberOfAccounts
        );

        let ocr_sequence_number = report_context.sequence_number();

        // The Global state PDA is sent as a remaining_account as it is optional to avoid having the lock when not modifying it, so all validations need to be done manually
        let (expected_state_key, _) = Pubkey::find_program_address(&[seed::STATE], &crate::ID);
        require_keys_eq!(
            remaining_accounts[0].key(),
            expected_state_key,
            CcipOfframpError::InvalidInputsGlobalStateAccount
        );
        require!(
            remaining_accounts[0].is_writable,
            CcipOfframpError::InvalidInputsMissingWritable
        );

        let mut global_state: Account<'info, GlobalState> =
            Account::try_from(&remaining_accounts[0])?;

        if global_state.latest_price_sequence_number <= ocr_sequence_number {
            // Update the persisted sequence number
            global_state.latest_price_sequence_number = ocr_sequence_number;
            global_state.exit(&crate::ID)?; // as it is manually loaded, it also has to be manually written back

            let cpi_remaining_accounts = remaining_accounts[1..].to_vec();
            let cpi_ctx = CpiContext::new_with_signer(cpi_program, cpi_accounts, cpi_signer)
                .with_remaining_accounts(cpi_remaining_accounts);

            let token_price_updates = report
                .price_updates
                .token_price_updates
                .iter()
                .map(|u| fee_quoter::context::TokenPriceUpdate {
                    source_token: u.source_token,
                    usd_per_token: u.usd_per_token,
                })
                .collect();
            let gas_price_update = report
                .price_updates
                .gas_price_updates
                .iter()
                .map(|u| fee_quoter::context::GasPriceUpdate {
                    dest_chain_selector: u.dest_chain_selector,
                    usd_per_unit_gas: u.usd_per_unit_gas,
                })
                .collect();

            fee_quoter::cpi::update_prices(cpi_ctx, token_price_updates, gas_price_update)?;
        }

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn commit_report_has_pending_messages() {
        let mut report = CommitReport {
            version: 1,
            chain_selector: 0,
            merkle_root: [0; 32],
            timestamp: 0,
            min_msg_nr: 42,
            max_msg_nr: 42,
            execution_states: 0,
        };
        assert!(
            !all_messages_executed(&report),
            "Message should still be pending"
        );
        report.execution_states = 3;
        assert!(
            !all_messages_executed(&report),
            "Failed message does not count as executed"
        );
        report.execution_states = 2;
        assert!(all_messages_executed(&report), "Single successful message");

        // Add 2 more messages
        report.max_msg_nr += 2;
        report.execution_states = 0b100011;
        assert!(
            !all_messages_executed(&report),
            "Mix - failed, untouched, failed"
        );
        report.execution_states = 0b101010;
        assert!(all_messages_executed(&report), "All messages executed");
    }
}
