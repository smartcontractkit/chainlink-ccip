use anchor_lang::prelude::*;
use solana_program::instruction::Instruction;
use solana_program::program::invoke_signed;

use crate::context::{seed, ExecuteReportContext, OcrPluginType};
use crate::event::{ExecutionStateChanged, SkippedAlreadyExecutedMessage};
use crate::instructions::interfaces::Execute;
use crate::messages::{
    Any2SVMRampMessage, ExecutionReportSingleChain, RampMessageHeader, SVMTokenAmount,
};
use crate::state::{CommitReport, MessageExecutionState, SourceChain};
use crate::CcipOfframpError;

use super::merkle::{calculate_merkle_root, MerkleError, LEAF_DOMAIN_SEPARATOR};
use super::messages::{is_writable, Any2SVMMessage, ReleaseOrMintInV1, ReleaseOrMintOutV1};
use super::ocr3base::{ocr3_transmit, ReportContext, Signatures};
use super::ocr3impl::Ocr3ReportForExecutionReportSingleChain;
use super::pools::{
    calculate_token_pool_account_indices, get_balance, interact_with_pool,
    validate_and_parse_token_accounts, TokenAccounts, CCIP_POOL_V1_RET_BYTES,
};
use super::rmn::verify_uncursed_cpi;

pub struct Impl;
impl Execute for Impl {
    fn execute<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        report_context_byte_words: [[u8; 32]; 2],
        token_indexes: &[u8],
    ) -> Result<()> {
        let execution_report =
            ExecutionReportSingleChain::deserialize(&mut raw_execution_report.as_ref())
                .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?;
        let report_context = ReportContext::from_byte_words(report_context_byte_words);
        verify_uncursed_cpi(
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            execution_report.source_chain_selector,
        )?;

        // limit borrowing of ctx
        {
            let config = ctx.accounts.config.load()?;
            ocr3_transmit(
                &config.ocr3[OcrPluginType::Execution as usize],
                &ctx.accounts.sysvar_instructions,
                ctx.accounts.authority.key(),
                OcrPluginType::Execution,
                report_context,
                &Ocr3ReportForExecutionReportSingleChain(&execution_report),
                Signatures {
                    rs: vec![],
                    ss: vec![],
                    raw_vs: [0u8; 32],
                },
            )?;
        }

        internal_execute(ctx, execution_report, token_indexes)
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
        let execution_report =
            ExecutionReportSingleChain::deserialize(&mut raw_execution_report.as_ref())
                .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?;
        verify_uncursed_cpi(
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            execution_report.source_chain_selector,
        )?;
        internal_execute(ctx, execution_report, token_indexes)
    }
}

/////////////
// Helpers //
/////////////

// internal_execute is the base execution logic without any additional validation
fn internal_execute<'info>(
    ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: ExecutionReportSingleChain,
    token_indexes: &[u8],
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

    // send tokens any -> SOL
    require!(
        token_indexes.len() == execution_report.message.token_amounts.len()
            && token_indexes.len() == execution_report.offchain_token_data.len(),
        CcipOfframpError::InvalidInputsTokenIndices,
    );
    let seeds = &[seed::EXTERNAL_TOKEN_POOL, &[ctx.bumps.token_pools_signer]];
    let mut token_amounts = vec![SVMTokenAmount::default(); token_indexes.len()];

    // handle tokens
    // note: indexes are used instead of counts in case more accounts need to be passed in remaining_accounts before token accounts
    // token_indexes = [2, 4] where remaining_accounts is [custom_account, custom_account, token1_account1, token1_account2, token2_account1, token2_account2] for example
    for (i, token_amount) in execution_report.message.token_amounts.iter().enumerate() {
        let accs = get_token_accounts_for(
            ctx.accounts.reference_addresses.load()?.router,
            ctx.accounts.reference_addresses.load()?.fee_quoter,
            ctx.remaining_accounts,
            execution_report.message.token_receiver,
            execution_report.message.header.source_chain_selector,
            token_indexes,
            i,
        )?;
        let router_token_pool_signer = &ctx.accounts.token_pools_signer;

        let init_bal = get_balance(accs.user_token_account)?;

        // CPI: call lockOrBurn on token pool
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
            router_token_pool_signer.to_account_info(),
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
        let return_data = interact_with_pool(
            accs.pool_program.key(),
            router_token_pool_signer.key(),
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
    let hashed_leaf = if should_execute_messaging(ctx.remaining_accounts, token_indexes) {
        let (msg_program, msg_accounts) = parse_messaging_accounts(
            token_indexes,
            &execution_report.message.extra_args.is_writable_bitmap,
            ctx.remaining_accounts,
        )?;

        // The External Execution Config Account is used to sign the CPI instruction
        let external_execution_config = &ctx.accounts.external_execution_config;

        // The accounts of the user that will be used in the CPI instruction, none of them are signers
        // They need to specify if mutable or not, but none of them is allowed to init, realloc or close
        // note: CPI signer is always first account
        let mut acc_infos = vec![
            external_execution_config.to_account_info(),
            ctx.accounts.offramp.to_account_info(),
            ctx.accounts.allowed_offramp.to_account_info(),
        ];
        acc_infos.extend_from_slice(msg_accounts);

        let acc_metas: Vec<AccountMeta> = acc_infos
            .to_vec()
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == external_execution_config.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let data = message.build_receiver_discriminator_and_data()?;

        let instruction = Instruction {
            program_id: msg_program.key(), // The receiver Program Id that will handle the ccip_receive message
            accounts: acc_metas,
            data,
        };

        let seeds = &[
            seed::EXTERNAL_EXECUTION_CONFIG,
            &[ctx.bumps.external_execution_config],
        ];
        let signer = &[&seeds[..]];

        invoke_signed(&instruction, &acc_infos, signer)?;
        let recv_and_msg_account_keys = Some(*msg_program.key)
            .into_iter()
            .chain(msg_accounts.iter().map(|a| *a.key));
        verify_merkle_root(
            &execution_report,
            commit_report.merkle_root,
            recv_and_msg_account_keys,
        )?
    } else {
        verify_merkle_root(
            &execution_report,
            commit_report.merkle_root,
            None.into_iter(),
        )?
    };

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

    Ok(())
}

fn get_token_accounts_for<'a>(
    router: Pubkey,
    fee_quoter: Pubkey,
    accounts: &'a [AccountInfo<'a>],
    token_receiver: Pubkey,
    chain_selector: u64,
    token_indexes: &[u8],
    i: usize,
) -> Result<TokenAccounts<'a>> {
    let (start, end) = calculate_token_pool_account_indices(i, token_indexes, accounts.len())?;

    let accs = validate_and_parse_token_accounts(
        token_receiver,
        chain_selector,
        router,
        fee_quoter,
        &accounts[start..end],
    )?;

    Ok(accs)
}

// There is at least one account used for messaging (the first subset of accounts). This is because the first account is the program id to do the CPI
fn should_execute_messaging<'a>(
    remaining_accounts: &'a [AccountInfo<'a>],
    token_indices: &[u8],
) -> bool {
    // The first entry in the accounts corresponds to a token, which means there is no logic receiver
    let only_tokens = token_indices.first().map(|i| *i == 0).unwrap_or_default();
    !remaining_accounts.is_empty() && !only_tokens
}

/// parse_message_accounts returns all the accounts needed to execute the CPI instruction
/// It also validates that the accounts sent in the message match the ones sent in the source chain
/// Precondition: logic_receiver != 0 && remaining_accounts.len() > 0
///
/// # Arguments
/// * `token_indexes` - start indexes of token pool accounts, used to determine ending index for arbitrary messaging accounts
/// * `remaining_accounts` - accounts passed via `ctx.remaining_accounts`. expected order is: [logic_receiver, ...additional message accounts]
///
/// # Return values
//  * `logic_receiver` is the Program ID of the user's program that will execute the message.
//  * `msg_accounts` the remaining list of accounts used for the arbitrary execution
fn parse_messaging_accounts<'info>(
    token_indexes: &[u8],
    source_bitmap: &u64,
    remaining_accounts: &'info [AccountInfo<'info>],
) -> Result<(&'info AccountInfo<'info>, &'info [AccountInfo<'info>])> {
    let end_index = if token_indexes.is_empty() {
        remaining_accounts.len()
    } else {
        token_indexes[0] as usize
    };

    require!(
        1 <= end_index && end_index <= remaining_accounts.len(), // program id and message accounts need to fit in remaining accounts
        CcipOfframpError::InvalidInputsTokenIndices
    ); // there could be other remaining accounts used for tokens

    let logic_receiver = &remaining_accounts[0];
    let msg_accounts = &remaining_accounts[1..end_index];

    // Validate that the bitmap corresponds to the individual writable flags
    for (i, acc) in msg_accounts.iter().enumerate().skip(1) {
        require!(
            is_writable(source_bitmap, i as u8) == acc.is_writable,
            CcipOfframpError::InvalidWritabilityBitmap
        );
    }

    Ok((logic_receiver, msg_accounts))
}

pub fn verify_merkle_root(
    execution_report: &ExecutionReportSingleChain,
    expected_root: [u8; 32],
    // logic receiver followed by all other message account keys, when they were
    // provided (i.e. when the message isn't a token transfer exclusively)
    recv_and_msg_account_keys: impl Iterator<Item = Pubkey>,
) -> Result<[u8; 32]> {
    let hashed_leaf = hash(&execution_report.message, recv_and_msg_account_keys);
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
    require!(
        execution_report.message.header.nonce == 0,
        CcipOfframpError::InvalidNonce
    );

    require!(
        source_chain_state.config.is_enabled,
        CcipOfframpError::UnsupportedSourceChainSelector
    );

    require!(
        execution_report.message.header.sequence_number >= commit_report.min_msg_nr
            && execution_report.message.header.sequence_number <= commit_report.max_msg_nr,
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

    Ok(())
}

fn hash(
    msg: &Any2SVMRampMessage,
    recv_and_msg_account_keys: impl Iterator<Item = Pubkey>,
) -> [u8; 32] {
    use anchor_lang::solana_program::keccak;

    // Calculate vectors size to ensure that the hash is unique
    let sender_size = msg.sender.len() as u16; // it should fit in a u8, but it's safer to use u16
    let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

    // RampMessageHeader struct
    let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
    let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
    let header_sequence_number = msg.header.sequence_number.to_be_bytes();
    let header_nonce = msg.header.nonce.to_be_bytes();

    let remaining_account_bytes: Vec<u8> = recv_and_msg_account_keys
        .flat_map(|k| k.try_to_vec().unwrap())
        .collect();

    // As similar as https://github.com/smartcontractkit/chainlink/blob/d1a9f8be2f222ea30bdf7182aaa6428bfa605cf7/contracts/src/v0.8/ccip/libraries/Internal.sol#L111
    let result = keccak::hashv(&[
        LEAF_DOMAIN_SEPARATOR.as_slice(),
        // metadata hash
        "Any2SVMMessageHashV1".as_bytes(),
        &header_source_chain_selector,
        &header_dest_chain_selector,
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
            Pubkey::try_from("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8").unwrap(),
            Pubkey::try_from("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH").unwrap(),
        ]
        .into_iter();

        let hash_result = hash(&message, remaining_account_keys);

        assert_eq!(
            "c82035cdc1d1e58606afeaf137b71de280e1e2cafdfdc621944eecccb105d730",
            hex::encode(hash_result)
        );
    }
}
