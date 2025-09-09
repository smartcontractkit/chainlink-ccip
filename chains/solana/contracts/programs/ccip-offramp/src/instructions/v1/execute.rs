use std::str::FromStr;

use anchor_lang::prelude::*;
use ccip_common::seed;
use ccip_common::v1::{validate_and_parse_token_accounts, TokenAccounts, MIN_TOKEN_POOL_ACCOUNTS};
use solana_program::instruction::Instruction;
use solana_program::program::invoke_signed;

use crate::context::ViewConfigOnly;
use crate::context::{BufferExecutionReportContext, ExecuteReportContext, OcrPluginType};
use crate::event::{ExecutionStateChanged, SkippedAlreadyExecutedMessage};
use crate::instructions::interfaces::Execute;
use crate::messages::{
    Any2SVMRampMessage, ExecutionReportSingleChain, RampMessageHeader, SVMTokenAmount,
};
use crate::state::{
    CommitReport, DeriveAccountsExecuteParams, DeriveAccountsResponse, ExecutionReportBuffer,
    MessageExecutionState, OnRampAddress, SourceChain,
};
use crate::CcipOfframpError;

use super::buffering::{deserialize_from_buffer_account, Buffering};
use super::merkle::{calculate_merkle_root, MerkleError, LEAF_DOMAIN_SEPARATOR};
use super::messages::{is_writable, Any2SVMMessage, ReleaseOrMintInV1, ReleaseOrMintOutV1};
use super::ocr3base::{ocr3_transmit, ReportContext, Signatures};
use super::ocr3impl::Ocr3ReportForExecutionReportSingleChain;
use super::pools::{get_balance, interact_with_pool, CCIP_POOL_V1_RET_BYTES};
use super::rmn::verify_uncursed_cpi;
mod derive;

pub struct Impl;
impl Execute for Impl {
    fn execute<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        report_context_byte_words: [[u8; 32]; 2],
        token_indexes: &[u8],
    ) -> Result<()> {
        let (execution_report, buffered_bytes) = if !raw_execution_report.is_empty() {
            (
                ExecutionReportSingleChain::deserialize(&mut raw_execution_report.as_ref())
                    .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?,
                0,
            )
        } else {
            deserialize_from_buffer_account(
                ctx.remaining_accounts
                    .last()
                    .ok_or(CcipOfframpError::ExecutionReportUnavailable)?,
            )?
        };

        let report_context = ReportContext::from_byte_words(report_context_byte_words);
        verify_uncursed_cpi(
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            execution_report.source_chain_selector,
        )?;

        ocr3_transmit_report(&ctx, &execution_report, report_context, buffered_bytes)?;
        internal_execute(
            ctx,
            execution_report,
            token_indexes,
            raw_execution_report.is_empty(),
        )
    }

    fn manually_execute<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        token_indexes: &[u8],
    ) -> Result<()> {
        // limit borrowing of ctx
        {
            let config = ctx.accounts.config.load()?;

            // validate time has passed
            let clock: Clock = Clock::get()?;
            let current_timestamp = clock.unix_timestamp;
            require!(
                current_timestamp - ctx.accounts.commit_report.timestamp
                    > config.enable_manual_execution_after,
                CcipOfframpError::ManualExecutionNotAllowed
            );
        }
        let execution_report = if !raw_execution_report.is_empty() {
            ExecutionReportSingleChain::deserialize(&mut raw_execution_report.as_ref())
                .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?
        } else {
            deserialize_from_buffer_account(
                ctx.remaining_accounts
                    .last()
                    .ok_or(CcipOfframpError::ExecutionReportUnavailable)?,
            )?
            .0
        };
        verify_uncursed_cpi(
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            execution_report.source_chain_selector,
        )?;
        internal_execute(
            ctx,
            execution_report,
            token_indexes,
            raw_execution_report.is_empty(),
        )
    }

    fn buffer_execution_report(
        &self,
        ctx: Context<BufferExecutionReportContext>,
        buffer_id: Vec<u8>,
        report_length: u32,
        chunk: Vec<u8>,
        chunk_index: u8,
        num_chunks: u8,
    ) -> Result<()> {
        require!(
            buffer_id.len() <= 32 && !buffer_id.is_empty(),
            CcipOfframpError::ExecutionReportBufferInvalidIdSize
        );
        ctx.accounts.execution_report_buffer.add_chunk(
            report_length,
            &chunk,
            chunk_index,
            num_chunks,
        )
    }

    fn derive_accounts_execute<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
        params: DeriveAccountsExecuteParams,
        stage: String,
    ) -> Result<DeriveAccountsResponse> {
        let stage = derive::DeriveAccountsExecuteStage::from_str(stage.as_str())?;

        match stage {
            derive::DeriveAccountsExecuteStage::Start => {
                derive::derive_execute_accounts_start(params.source_chain_selector)
            }
            derive::DeriveAccountsExecuteStage::FinishMainAccountList => {
                derive::derive_execute_accounts_build_main_account_list(ctx, &params)
            }
            derive::DeriveAccountsExecuteStage::RetrieveTokenLUTs => {
                derive::derive_execute_accounts_retrieve_luts(ctx)
            }
            derive::DeriveAccountsExecuteStage::RetrievePoolPrograms => {
                derive::derive_execute_accounts_retrieve_pool_programs(ctx)
            }
            derive::DeriveAccountsExecuteStage::TokenTransferStaticAccounts { token, page } => {
                derive::derive_execute_accounts_additional_tokens_static(ctx, &params, page, token)
            }
            derive::DeriveAccountsExecuteStage::NestedTokenDerive {
                token,
                token_substage,
            } => derive::derive_execute_accounts_additional_token_nested(
                ctx,
                &params,
                &token_substage,
                token,
            ),
        }
    }
}

/////////////
// Helpers //
/////////////
fn ocr3_transmit_report<'info>(
    ctx: &Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: &ExecutionReportSingleChain,
    report_context: ReportContext,
    buffered_bytes: usize,
) -> Result<()> {
    let config = ctx.accounts.config.load()?;

    ocr3_transmit(
        &config.ocr3[OcrPluginType::Execution as usize],
        &ctx.accounts.sysvar_instructions,
        ctx.accounts.authority.key(),
        OcrPluginType::Execution,
        report_context,
        &Ocr3ReportForExecutionReportSingleChain(execution_report),
        Signatures {
            rs: vec![],
            ss: vec![],
            raw_vs: [0u8; 32],
        },
        buffered_bytes,
    )?;
    Ok(())
}

// internal_execute is the base execution logic without any additional validation
fn internal_execute<'info>(
    ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: ExecutionReportSingleChain,
    token_indexes: &[u8],
    report_is_buffered: bool,
) -> Result<()> {
    // TODO: Limit send size data to 256

    // The Config Account stores the default values for the Router, the SVM Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
    let config = ctx.accounts.config.load()?;
    let svm_chain_selector = config.svm_chain_selector;

    // The Config and State for the Source Chain, containing if it is enabled, the on ramp address and the min sequence number expected for future messages
    let source_chain = &ctx.accounts.source_chain;

    // The Commit Report Account stores the information of 1 Commit Report:
    // - Merkle Root
    // - Timestamp of the Commit Report
    // - Interval of Messages: The min and max seq num of the messages in the Merkle Tree
    // - Execution State per each Message: 0 for Untouched, 1 for InProgress, 2 for Success and 3 for Failure
    let commit_report = &mut ctx.accounts.commit_report;

    let message_header = execution_report.message.header;

    validate_execution_report(
        &execution_report,
        source_chain,
        commit_report,
        &message_header,
        svm_chain_selector,
    )?;

    let original_state = execution_state::get(commit_report, message_header.sequence_number);

    if original_state == MessageExecutionState::Success {
        emit!(SkippedAlreadyExecutedMessage {
            source_chain_selector: message_header.source_chain_selector,
            sequence_number: message_header.sequence_number,
        });
        return Ok(());
    }

    let remaining_accounts_layout = ExecuteReportContextRemainingAccountsLayout::new(
        ctx.remaining_accounts,
        token_indexes,
        report_is_buffered,
        &execution_report.message.extra_args.is_writable_bitmap,
    )?;

    let hashed_leaf: [u8; 32] = if remaining_accounts_layout.should_execute_messaging() {
        // Verify merkle root before doing any token operations or CPI calls
        let keys = remaining_accounts_layout.receiver_and_message_account_keys();
        verify_merkle_root(
            &execution_report,
            commit_report.merkle_root,
            keys,
            &source_chain.config.on_ramp,
        )?
    } else {
        verify_merkle_root(
            &execution_report,
            commit_report.merkle_root,
            None,
            &source_chain.config.on_ramp,
        )?
    };

    // Mark as InProgress and emit ExecutionStateChanged event; the event will be used by offchain to recognize that
    // the message has been attempted in order to avoid more attempts.
    // An attempt is considered valid only if it passes necessary validation, such as merkle proof check.
    // Since Solana keeps logs even if the transaction errors, this approach works regardless of attempt outcome.
    // This event should be emitted before any operations that calls 3rd party programs.
    let in_progress_state = MessageExecutionState::InProgress;
    execution_state::set(
        commit_report,
        message_header.sequence_number,
        in_progress_state.to_owned(),
    );

    emit!(ExecutionStateChanged {
        source_chain_selector: message_header.source_chain_selector,
        sequence_number: message_header.sequence_number,
        message_id: message_header.message_id,
        message_hash: hashed_leaf,
        state: in_progress_state,
    });

    // send tokens any -> SOL
    require!(
        token_indexes.len() == execution_report.message.token_amounts.len()
            && token_indexes.len() == execution_report.offchain_token_data.len(),
        CcipOfframpError::InvalidInputsTokenIndices,
    );
    let mut token_amounts = vec![SVMTokenAmount::default(); token_indexes.len()];

    // handle tokens
    // note: indexes are used instead of counts in case more accounts need to be passed in remaining_accounts before token accounts
    // token_indexes = [2, 4] where remaining_accounts is [custom_account, custom_account, token1_account1, token1_account2, token2_account1, token2_account2] for example
    for (i, token_amount) in execution_report.message.token_amounts.iter().enumerate() {
        let accs = remaining_accounts_layout.get_token_accounts_for(
            ctx.accounts.reference_addresses.load()?.router,
            ctx.accounts.reference_addresses.load()?.fee_quoter,
            execution_report.message.token_receiver,
            execution_report.message.header.source_chain_selector,
            i,
        )?;
        let offramp_token_pool_signer = accs
            .ccip_offramp_pool_signer
            .ok_or(CcipOfframpError::InvalidInputsPoolAccounts)?;

        let init_bal = get_balance(accs.user_token_account)?;

        // CPI: call releaseOrMint on token pool
        let release_or_mint = ReleaseOrMintInV1 {
            original_sender: execution_report.message.sender.clone(),
            receiver: execution_report.message.token_receiver,
            amount: token_amount.amount,
            local_token: token_amount.dest_token_address,
            remote_chain_selector: execution_report.message.header.source_chain_selector,
            source_pool_address: token_amount.source_pool_address.clone(),
            source_pool_data: token_amount.extra_data.clone(),
            offchain_token_data: execution_report.offchain_token_data[i].clone(),
        };
        let mut acc_infos = vec![
            offramp_token_pool_signer.to_account_info(),
            ctx.accounts.offramp.to_account_info(),
            ctx.accounts.allowed_offramp.to_account_info(),
            accs.pool_config.to_account_info(),
            accs.token_program.to_account_info(),
            accs.mint.to_account_info(),
            accs.pool_signer.to_account_info(),
            accs.pool_token_account.to_account_info(),
            accs.pool_chain_config.to_account_info(),
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            accs.user_token_account.to_account_info(),
        ];
        acc_infos.extend_from_slice(accs.remaining_accounts);

        let seeds = &[
            seed::EXTERNAL_TOKEN_POOLS_SIGNER,
            accs.pool_program.key.as_ref(),
            &[accs.ccip_offramp_pool_signer_bump],
        ];

        let return_data = interact_with_pool(
            accs.pool_program.key(),
            offramp_token_pool_signer.key(),
            acc_infos,
            release_or_mint,
            seeds,
        )?;

        require!(
            return_data.len() == CCIP_POOL_V1_RET_BYTES,
            CcipOfframpError::OfframpInvalidDataLength
        );

        // parse pool return data into SVMTokenAmount
        token_amounts[i] = SVMTokenAmount {
            token: accs.mint.key(),
            amount: ReleaseOrMintOutV1::try_from_slice(&return_data)?.destination_amount,
        };

        // validate user received tokens according to the amount returned by the token pool
        let post_bal = get_balance(accs.user_token_account)?;
        require!(
            post_bal >= init_bal && post_bal - init_bal == token_amounts[i].amount,
            CcipOfframpError::OfframpReleaseMintBalanceMismatch
        );
    }

    let message = Any2SVMMessage {
        message_id: execution_report.message.header.message_id,
        source_chain_selector: execution_report.source_chain_selector,
        sender: execution_report.message.sender.clone(),
        data: execution_report.message.data.clone(),
        token_amounts,
    };

    // handle CPI call if there are message accounts in the remaining_accounts
    // case: no tokens, but there are remaining_accounts passed in
    // case: tokens and messages, so the first token has a non-zero index (indicating extra accounts before token accounts)
    if remaining_accounts_layout.should_execute_messaging() {
        let msg_accs = remaining_accounts_layout
            .messaging_accounts
            .as_ref()
            .unwrap(); // as there is messaging, the option is guaranteed to be Some

        // The accounts of the user that will be used in the CPI instruction, none of them are signers
        // They need to specify if mutable or not, but none of them is allowed to init, realloc or close
        // note: CPI signer is always first account
        let mut acc_infos = vec![
            msg_accs.external_execution_signer.to_account_info(),
            ctx.accounts.offramp.to_account_info(),
            ctx.accounts.allowed_offramp.to_account_info(),
        ];

        let mut acc_metas: Vec<AccountMeta> = acc_infos
            .iter()
            .flat_map(|acc_info| {
                let is_signer = acc_info.key() == msg_accs.external_execution_signer.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let remaining_metas: Vec<AccountMeta> = msg_accs
            .remaining_messaging_accounts
            .iter()
            .enumerate()
            .map(|(i, acc_info)| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == msg_accs.external_execution_signer.key();
                let is_writable = is_writable(
                    &execution_report.message.extra_args.is_writable_bitmap,
                    i as u8,
                );

                if is_writable {
                    AccountMeta::new(*acc_info.key, is_signer)
                } else {
                    AccountMeta::new_readonly(*acc_info.key, is_signer)
                }
            })
            .collect();

        acc_infos.extend_from_slice(msg_accs.remaining_messaging_accounts);
        acc_metas.extend_from_slice(&remaining_metas);

        let data = message.build_receiver_discriminator_and_data()?;

        let instruction = Instruction {
            program_id: msg_accs.logic_receiver.key(), // The receiver Program Id that will handle the ccip_receive message
            accounts: acc_metas,
            data,
        };

        let seeds = &[
            seed::EXTERNAL_EXECUTION_CONFIG,
            msg_accs.logic_receiver.key.as_ref(),
            &[msg_accs.external_execution_signer_bump],
        ];
        let signer = &[&seeds[..]];

        invoke_signed(&instruction, &acc_infos, signer)?;
    }

    let new_state = MessageExecutionState::Success;
    execution_state::set(
        commit_report,
        message_header.sequence_number,
        new_state.to_owned(),
    );

    emit!(ExecutionStateChanged {
        source_chain_selector: message_header.source_chain_selector,
        sequence_number: message_header.sequence_number,
        message_id: message_header.message_id, // Unique identifier for the message, generated with the source chain's encoding scheme
        message_hash: hashed_leaf,             // Hash of the message using SVM encoding
        state: new_state,
    });

    if report_is_buffered {
        let buffer =
            Account::<ExecutionReportBuffer>::try_from(ctx.remaining_accounts.last().unwrap())?;
        buffer.close(ctx.accounts.authority.to_account_info())?;
    }

    Ok(())
}

// Encodes the structured layout of the Execute Report Context remaining accounts, so they can
// be accurately used where needed instead of passed as a raw array.
pub struct ExecuteReportContextRemainingAccountsLayout<'a> {
    pub messaging_accounts: Option<ExecuteContextRemainingMessagingAccounts<'a>>,
    pub token_accounts_per_token: Vec<ExecuteContextRemainingTokenAccounts<'a>>,
    pub _buffering_account: Option<ExecuteContextRemainingBufferingAccount<'a>>,
}

pub struct ExecuteContextRemainingMessagingAccounts<'a> {
    pub logic_receiver: &'a AccountInfo<'a>,
    pub external_execution_signer: &'a AccountInfo<'a>,
    pub external_execution_signer_bump: u8,
    pub remaining_messaging_accounts: &'a [AccountInfo<'a>],
}

pub struct ExecuteContextRemainingTokenAccounts<'a> {
    pub accounts: &'a [AccountInfo<'a>],
}

pub struct ExecuteContextRemainingBufferingAccount<'a> {
    pub _execution_buffer_account: &'a AccountInfo<'a>,
}

impl<'a> ExecuteReportContextRemainingAccountsLayout<'a> {
    pub fn should_execute_messaging(&self) -> bool {
        self.messaging_accounts.is_some()
    }

    pub fn receiver_and_message_account_keys(&self) -> Vec<Pubkey> {
        self.messaging_accounts
            .iter()
            .flat_map(|accs| {
                Some(accs.logic_receiver.key())
                    .into_iter()
                    .chain(accs.remaining_messaging_accounts.iter().map(|a| a.key()))
            })
            .collect()
    }

    pub fn get_token_accounts_for<'b>(
        &'b self,
        router: Pubkey,
        fee_quoter: Pubkey,
        token_receiver: Pubkey,
        chain_selector: u64,
        i: usize,
    ) -> Result<TokenAccounts<'a>> {
        let accs = self
            .token_accounts_per_token
            .get(i)
            .map(|accounts_per_token| accounts_per_token.accounts)
            .ok_or(CcipOfframpError::InvalidInputsTokenIndices)?;

        validate_and_parse_token_accounts(
            token_receiver,
            chain_selector,
            router,
            fee_quoter,
            Some(crate::ID),
            accs,
        )
    }

    pub fn new(
        remaining_accounts: &'a [AccountInfo<'a>],
        token_indices: &[u8],
        report_is_buffered: bool,
        source_bitmap: &u64,
    ) -> Result<Self> {
        ////////////////////////
        // Buffering subslice //
        ////////////////////////

        // First, if the report is buffered, it means the last account is the buffering account.
        // We can simply remove it from consideration from now on.
        let (remaining_accounts_without_buffering, buffering_account) = if report_is_buffered {
            (
                &remaining_accounts[0..(remaining_accounts.len() - 1)],
                Some(ExecuteContextRemainingBufferingAccount {
                    _execution_buffer_account: remaining_accounts
                        .last()
                        .ok_or(CcipOfframpError::ExecutionReportUnavailable)?,
                }),
            )
        } else {
            (remaining_accounts, None)
        };

        ////////////////////////
        // Messaging subslice //
        ////////////////////////
        let only_tokens = token_indices.first().map(|i| *i == 0).unwrap_or_default();
        let messaging_accounts_exist =
            !remaining_accounts_without_buffering.is_empty() && !only_tokens;
        let messaging_accounts = if messaging_accounts_exist {
            let end_index = token_indices
                .first()
                .map(|i| *i as usize)
                .unwrap_or(remaining_accounts_without_buffering.len());

            require!(
                1 <= end_index && end_index <= remaining_accounts_without_buffering.len(), // program id and message accounts need to fit in remaining accounts
                CcipOfframpError::InvalidInputsTokenIndices
            ); // there could be other remaining accounts used for tokens

            let logic_receiver = &remaining_accounts_without_buffering[0];
            let external_execution_signer = &remaining_accounts_without_buffering[1];
            let msg_accounts = &remaining_accounts_without_buffering[2..end_index];

            // Validate the derivation of the external_execution_signer and calculate its bump
            let (expected_signer_key, signer_bump) = Pubkey::find_program_address(
                &[
                    seed::EXTERNAL_EXECUTION_CONFIG,
                    logic_receiver.key().as_ref(),
                ],
                &crate::ID,
            );
            require_keys_eq!(
                external_execution_signer.key(),
                expected_signer_key,
                CcipOfframpError::InvalidInputsExternalExecutionSignerAccount
            );

            // Validate that the bitmap corresponds to the individual writable flags
            for (i, acc) in msg_accounts.iter().enumerate() {
                require!(
                    !is_writable(source_bitmap, i as u8) || acc.is_writable,
                    CcipOfframpError::InvalidWritabilityBitmap
                );
            }

            let messaging_accounts = ExecuteContextRemainingMessagingAccounts {
                logic_receiver,
                external_execution_signer,
                external_execution_signer_bump: signer_bump,
                remaining_messaging_accounts: msg_accounts,
            };

            Some(messaging_accounts)
        } else {
            None
        };

        ////////////////////
        // Token subslice //
        ////////////////////
        let indices_with_sentinel = token_indices
            .iter()
            .copied()
            .map(|i| i as usize)
            // We add a sentinel for the last token range ending at the end of the account list.
            .chain(Some(remaining_accounts_without_buffering.len()));
        let pairwise_indices = indices_with_sentinel
            .clone()
            .zip(indices_with_sentinel.skip(1));

        let token_accounts_per_token: Vec<ExecuteContextRemainingTokenAccounts> = pairwise_indices
            .map(|(start, end)| {
                assert!(
                    start < end
                        && end <= remaining_accounts_without_buffering.len()
                        && end - start >= MIN_TOKEN_POOL_ACCOUNTS,
                );

                ExecuteContextRemainingTokenAccounts {
                    accounts: &remaining_accounts_without_buffering[start..end],
                }
            })
            .collect();

        Ok(Self {
            messaging_accounts,
            token_accounts_per_token,
            _buffering_account: buffering_account,
        })
    }
}

pub fn verify_merkle_root(
    execution_report: &ExecutionReportSingleChain,
    expected_root: [u8; 32],
    // logic receiver followed by all other message account keys, when they were
    // provided (i.e. when the message isn't a token transfer exclusively)
    recv_and_msg_account_keys: impl IntoIterator<Item = Pubkey>,
    on_ramp_address: &OnRampAddress,
) -> Result<[u8; 32]> {
    let hashed_leaf = hash(
        &execution_report.message,
        recv_and_msg_account_keys,
        on_ramp_address,
    );
    let verified_root: std::result::Result<[u8; 32], MerkleError> =
        calculate_merkle_root(hashed_leaf, &execution_report.proofs);
    require!(
        verified_root.is_ok() && verified_root.unwrap() == expected_root,
        CcipOfframpError::InvalidProof
    );
    Ok(hashed_leaf)
}

pub fn validate_execution_report<'info>(
    execution_report: &ExecutionReportSingleChain,
    source_chain_state: &Account<'info, SourceChain>,
    commit_report: &Account<'info, CommitReport>,
    message_header: &RampMessageHeader,
    svm_chain_selector: u64,
) -> Result<()> {
    require_eq!(message_header.nonce, 0, CcipOfframpError::InvalidNonce);

    require!(
        source_chain_state.config.is_enabled,
        CcipOfframpError::UnsupportedSourceChainSelector
    );

    require!(
        message_header.sequence_number >= commit_report.min_msg_nr
            && message_header.sequence_number <= commit_report.max_msg_nr,
        CcipOfframpError::InvalidSequenceInterval
    );

    require!(
        message_header.source_chain_selector == execution_report.source_chain_selector,
        CcipOfframpError::UnsupportedSourceChainSelector
    );
    require!(
        message_header.dest_chain_selector == svm_chain_selector,
        CcipOfframpError::UnsupportedDestinationChainSelector
    );
    require!(
        commit_report.timestamp != 0,
        CcipOfframpError::RootNotCommitted
    );
    require!(
        // This can't be easily done directly in the context, as the execution report
        // may have been retrieved in different ways (via parameter or buffering account.)
        commit_report.chain_selector == execution_report.source_chain_selector,
        CcipOfframpError::RootNotCommitted
    );

    Ok(())
}

fn hash(
    msg: &Any2SVMRampMessage,
    recv_and_msg_account_keys: impl IntoIterator<Item = Pubkey>,
    on_ramp_address: &OnRampAddress,
) -> [u8; 32] {
    use anchor_lang::solana_program::keccak;

    // Calculate vectors size to ensure that the hash is unique
    let sender_size = msg.sender.len() as u16; // it should fit in a u8, but it's safer to use u16
    let on_ramp_address_size = on_ramp_address.bytes().len() as u16; // it should fit in a u8, but it's safer to use u16
    let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

    // RampMessageHeader struct
    let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
    let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
    let header_sequence_number = msg.header.sequence_number.to_be_bytes();
    let header_nonce = msg.header.nonce.to_be_bytes();

    let remaining_account_bytes: Vec<u8> = recv_and_msg_account_keys
        .into_iter()
        .flat_map(|k| k.try_to_vec().unwrap())
        .collect();

    // As similar as https://github.com/smartcontractkit/chainlink/blob/d1a9f8be2f222ea30bdf7182aaa6428bfa605cf7/contracts/src/v0.8/ccip/libraries/Internal.sol#L111
    let result = keccak::hashv(&[
        LEAF_DOMAIN_SEPARATOR.as_slice(),
        // metadata hash
        "Any2SVMMessageHashV1".as_bytes(),
        &header_source_chain_selector,
        &header_dest_chain_selector,
        &on_ramp_address_size.to_be_bytes(),
        on_ramp_address.bytes(),
        // message header
        &msg.header.message_id,
        &msg.token_receiver.to_bytes(),
        &header_sequence_number,
        msg.extra_args.try_to_vec().unwrap().as_ref(), // borsh serialized
        &header_nonce,
        // message
        &sender_size.to_be_bytes(),
        &msg.sender,
        &data_size.to_be_bytes(),
        &msg.data,
        // token transfers
        msg.token_amounts.try_to_vec().unwrap().as_ref(), // borsh serialized
        // Remaining accounts (passed outside `Any2SVMRampMessage`)
        &remaining_account_bytes,
    ]);

    result.to_bytes()
}

mod execution_state {
    use crate::state::{CommitReport, MessageExecutionState};

    pub fn set(
        report: &mut CommitReport,
        sequence_number: u64,
        execution_state: MessageExecutionState,
    ) {
        let packed = &mut report.execution_states;
        let dif = sequence_number.checked_sub(report.min_msg_nr);
        assert!(dif.is_some(), "Sequence number out of bounds");
        let i = dif.unwrap();
        assert!(i < 64, "Sequence number out of bounds");

        // Clear the 2 bits at position 'i'
        *packed &= !(0b11 << (i * 2));
        // Set the new value in the cleared bits
        *packed |= (execution_state as u128) << (i * 2);
    }

    pub fn get(report: &CommitReport, sequence_number: u64) -> MessageExecutionState {
        let packed = report.execution_states;
        let dif = sequence_number.checked_sub(report.min_msg_nr);
        assert!(dif.is_some(), "Sequence number out of bounds");
        let i = dif.unwrap();
        assert!(i < 64, "Sequence number out of bounds");

        let mask = 0b11 << (i * 2);
        let state = (packed & mask) >> (i * 2);
        MessageExecutionState::try_from(state).unwrap()
    }

    #[cfg(test)]
    mod tests {
        use super::*;

        #[test]
        fn test_set_state() {
            let mut commit_report = CommitReport {
                version: 1,
                chain_selector: 0,
                merkle_root: [0; 32],
                timestamp: 0,
                min_msg_nr: 0,
                max_msg_nr: 64,
                execution_states: 0,
            };

            set(&mut commit_report, 0, MessageExecutionState::Success);
            assert_eq!(get(&commit_report, 0), MessageExecutionState::Success);

            set(&mut commit_report, 1, MessageExecutionState::Failure);
            assert_eq!(get(&commit_report, 1), MessageExecutionState::Failure);

            set(&mut commit_report, 2, MessageExecutionState::Untouched);
            assert_eq!(get(&commit_report, 2), MessageExecutionState::Untouched);

            set(&mut commit_report, 3, MessageExecutionState::InProgress);
            assert_eq!(get(&commit_report, 3), MessageExecutionState::InProgress);
        }

        #[test]
        #[should_panic(expected = "Sequence number out of bounds")]
        fn test_set_state_out_of_bounds() {
            let mut commit_report = CommitReport {
                version: 1,
                chain_selector: 1,
                merkle_root: [0; 32],
                timestamp: 1,
                min_msg_nr: 1500,
                max_msg_nr: 1530,
                execution_states: 0,
            };

            set(&mut commit_report, 65, MessageExecutionState::Success);
        }

        #[test]
        fn test_get_state() {
            let mut commit_report = CommitReport {
                version: 1,
                chain_selector: 1,
                merkle_root: [0; 32],
                timestamp: 1,
                min_msg_nr: 1500,
                max_msg_nr: 1530,
                execution_states: 0,
            };

            set(&mut commit_report, 1501, MessageExecutionState::Success);
            set(&mut commit_report, 1505, MessageExecutionState::Failure);
            set(&mut commit_report, 1520, MessageExecutionState::Untouched);
            set(&mut commit_report, 1523, MessageExecutionState::InProgress);

            assert_eq!(get(&commit_report, 1501), MessageExecutionState::Success);
            assert_eq!(get(&commit_report, 1505), MessageExecutionState::Failure);
            assert_eq!(get(&commit_report, 1520), MessageExecutionState::Untouched);
            assert_eq!(get(&commit_report, 1523), MessageExecutionState::InProgress);
        }

        #[test]
        #[should_panic(expected = "Sequence number out of bounds")]
        fn test_get_state_out_of_bounds() {
            let commit_report = CommitReport {
                version: 1,
                chain_selector: 1,
                merkle_root: [0; 32],
                timestamp: 1,
                min_msg_nr: 1500,
                max_msg_nr: 1530,
                execution_states: 0,
            };

            get(&commit_report, 65);
        }
    }
}

#[cfg(test)]
mod tests {
    use ethnum::U256;

    use super::*;
    use crate::messages::{Any2SVMRampExtraArgs, Any2SVMRampMessage, Any2SVMTokenTransfer};

    /// Builds a message and hash it, it's compared with a known hash
    #[test]
    fn test_hash() {
        let on_ramp_address: OnRampAddress = [1, 2, 3].into();

        let message = Any2SVMRampMessage {
            sender: [
                1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0,
            ]
            .to_vec(),
            token_receiver: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
                .unwrap(),
            data: vec![4, 5, 6],
            header: RampMessageHeader {
                message_id: [
                    8, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                    0, 0, 0, 0, 0, 0,
                ],
                source_chain_selector: 67,
                dest_chain_selector: 78,
                sequence_number: 89,
                nonce: 90,
            },
            token_amounts: [Any2SVMTokenTransfer {
                source_pool_address: vec![0, 1, 2, 3],
                dest_token_address: Pubkey::try_from(
                    "DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc",
                )
                .unwrap(),
                dest_gas_amount: 100,
                extra_data: vec![4, 5, 6],
                amount: U256::from_le_bytes([1; 32]).into(),
            }]
            .to_vec(),
            extra_args: Any2SVMRampExtraArgs {
                compute_units: 1000,
                is_writable_bitmap: 1,
            },
        };
        let remaining_account_keys = [
            Pubkey::try_from("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C").unwrap(),
            Pubkey::try_from("EvhgrPhTDt4LcSPS2kfJgH6T6XWZ6wT3X9ncDGLT1vui").unwrap(),
        ]
        .into_iter();

        let hash_result = hash(&message, remaining_account_keys, &on_ramp_address);

        assert_eq!(
            "5ddb3c9fccb01abee926ec6112afa075dc81fdfe1e2902595d9c1d1d1de4f1d1",
            hex::encode(hash_result)
        );
    }
}
