use anchor_lang::prelude::*;
use solana_program::{instruction::Instruction, program::invoke_signed};

use super::merkle::{calculate_merkle_root, MerkleError};
use super::messages::pools::{ReleaseOrMintInV1, ReleaseOrMintOutV1};
use super::ocr3base::{ocr3_transmit, ReportContext};
use super::ocr3impl::{Ocr3ReportForCommit, Ocr3ReportForExecutionReportSingleChain};
use super::pools::{
    calculate_token_pool_account_indices, get_balance, interact_with_pool,
    validate_and_parse_token_accounts, CCIP_POOL_V1_RET_BYTES,
};

use crate::v1::config::is_on_ramp_configured;
use crate::v1::merkle::LEAF_DOMAIN_SEPARATOR;
use crate::v1::messages::ramps::is_writable;
use crate::{
    Any2SolanaMessage, Any2SolanaRampMessage, BillingTokenConfigWrapper, CcipRouterError,
    CommitInput, CommitReport, CommitReportAccepted, CommitReportContext, DestChain,
    ExecuteReportContext, ExecutionReportSingleChain, ExecutionStateChanged, GasPriceUpdate,
    GlobalState, MessageExecutionState, OcrPluginType, RampMessageHeader,
    SkippedAlreadyExecutedMessage, SolanaTokenAmount, SourceChain, TimestampedPackedU224,
    TokenPriceUpdate, UsdPerTokenUpdated, UsdPerUnitGasUpdated, CCIP_RECEIVE_DISCRIMINATOR,
    DEST_CHAIN_STATE_SEED, EXTERNAL_EXECUTION_CONFIG_SEED, EXTERNAL_TOKEN_POOL_SEED,
    FEE_BILLING_TOKEN_CONFIG, STATE_SEED,
};

pub fn commit<'info>(
    ctx: Context<'_, '_, 'info, 'info, CommitReportContext<'info>>,
    report_context_byte_words: [[u8; 32]; 3],
    report: CommitInput,
    signatures: Vec<[u8; 65]>,
) -> Result<()> {
    let report_context = ReportContext::from_byte_words(report_context_byte_words);

    // The Config Account stores the default values for the Router, the Solana Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
    let config = ctx.accounts.config.load()?;

    // The Config and State for the Source Chain, containing if it is enabled, the on ramp address and the min sequence number expected for future messages
    let source_chain_state = &mut ctx.accounts.source_chain_state;

    require!(
        source_chain_state.config.is_enabled,
        CcipRouterError::UnsupportedSourceChainSelector
    );
    require!(
        is_on_ramp_configured(
            &source_chain_state.config,
            &report.merkle_root.on_ramp_address
        ),
        CcipRouterError::InvalidInputs
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
            CcipRouterError::InvalidInputs
        );
    } else {
        // There are price updates in the report.
        // Remaining accounts represent:
        // - The state account to store the price sequence updates
        // - the accounts to update BillingTokenConfig for token prices
        // - the accounts to update DestChain for gas prices
        // They must be in order:
        // 1. state_account
        // 2. token_accounts[]
        // 3. gas_accounts[]
        // matching the order of the price updates in the CommitInput.
        // They must also all be writable so they can be updated.
        let minimum_remaining_accounts = 1
            + report.price_updates.token_price_updates.len()
            + report.price_updates.gas_price_updates.len();
        require_eq!(
            ctx.remaining_accounts.len(),
            minimum_remaining_accounts,
            CcipRouterError::InvalidInputs
        );

        let ocr_sequence_number = report_context.sequence_number();

        // The Global state PDA is sent as a remaining_account as it is optional to avoid having the lock when not modifying it, so all validations need to be done manually
        let (expected_state_key, _) = Pubkey::find_program_address(&[STATE_SEED], &crate::ID);
        require_keys_eq!(
            ctx.remaining_accounts[0].key(),
            expected_state_key,
            CcipRouterError::InvalidInputs
        );
        require!(
            ctx.remaining_accounts[0].is_writable,
            CcipRouterError::InvalidInputs
        );

        let mut global_state: Account<GlobalState> = Account::try_from(&ctx.remaining_accounts[0])?;

        if global_state.latest_price_sequence_number < ocr_sequence_number {
            // Update the persisted sequence number
            global_state.latest_price_sequence_number = ocr_sequence_number;
            global_state.exit(&crate::ID)?; // as it is manually loaded, it also has to be manually written back

            // For each token price update, unpack the corresponding remaining_account and update the price.
            // Keep in mind that the remaining_accounts are sorted in the same order as tokens and gas price updates in the report.
            for (i, update) in report.price_updates.token_price_updates.iter().enumerate() {
                apply_token_price_update(update, &ctx.remaining_accounts[i + 1])?;
            }

            // Skip the first state account and the ones for token updates
            let offset = report.price_updates.token_price_updates.len() + 1;

            // Do the same for gas price updates
            for (i, update) in report.price_updates.gas_price_updates.iter().enumerate() {
                apply_gas_price_update(update, &ctx.remaining_accounts[i + offset])?;
            }
        } else {
            // TODO check if this is really necessary. EVM has this validation checking that the
            // array of merkle roots in the report is not empty. But here, considering we only have 1 root per report,
            // this check is just validating that the root is not zeroed
            // (which should never happen anyway, so it may be redundant).
            require!(
                report.merkle_root.source_chain_selector > 0,
                CcipRouterError::StaleCommitReport
            );
        }
    }

    // The Commit Report Account stores the information of 1 Commit Report:
    // - Merkle Root
    // - Timestamp of the Commit Report
    // - Interval of Messages: The min and max seq num of the messages in the Merkle Tree
    // - Execution State per each Message: 0 for Untouched, 1 for InProgress, 2 for Success and 3 for Failure
    let commit_report = &mut ctx.accounts.commit_report;
    let root = &report.merkle_root;

    require!(
        root.min_seq_nr <= root.max_seq_nr,
        CcipRouterError::InvalidSequenceInterval
    );
    require!(
        root.max_seq_nr
            .to_owned()
            .checked_sub(root.min_seq_nr)
            .map_or_else(|| false, |seq_size| seq_size <= 64),
        CcipRouterError::InvalidSequenceInterval
    ); // As we have 64 slots to store the execution state
    require!(
        source_chain_state.state.min_seq_nr == root.min_seq_nr,
        CcipRouterError::InvalidSequenceInterval
    );
    require!(root.merkle_root != [0; 32], CcipRouterError::InvalidProof);
    require!(
        commit_report.timestamp == 0,
        CcipRouterError::ExistingMerkleRoot
    );

    let next_seq_nr = root.max_seq_nr.checked_add(1);

    require!(
        next_seq_nr.is_some(),
        CcipRouterError::ReachedMaxSequenceNumber
    );

    source_chain_state.state.min_seq_nr = next_seq_nr.unwrap();

    let clock: Clock = Clock::get()?;
    commit_report.version = 1;
    commit_report.chain_selector = report.merkle_root.source_chain_selector;
    commit_report.merkle_root = report.merkle_root.merkle_root;
    commit_report.timestamp = clock.unix_timestamp;
    commit_report.execution_states = 0;
    commit_report.min_msg_nr = root.min_seq_nr;
    commit_report.max_msg_nr = root.max_seq_nr;

    emit!(CommitReportAccepted {
        merkle_root: root.clone(),
        price_updates: report.price_updates.clone(),
    });

    ocr3_transmit(
        &config.ocr3[OcrPluginType::Commit as usize],
        &ctx.accounts.sysvar_instructions,
        ctx.accounts.authority.key(),
        OcrPluginType::Commit as u8,
        report_context,
        &Ocr3ReportForCommit(&report),
        &signatures,
    )?;

    Ok(())
}

pub fn execute<'info>(
    ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: ExecutionReportSingleChain,
    report_context_byte_words: [[u8; 32]; 3],
) -> Result<()> {
    let report_context = ReportContext::from_byte_words(report_context_byte_words);
    // limit borrowing of ctx
    {
        let config = ctx.accounts.config.load()?;
        ocr3_transmit(
            &config.ocr3[OcrPluginType::Execution as usize],
            &ctx.accounts.sysvar_instructions,
            ctx.accounts.authority.key(),
            OcrPluginType::Execution as u8,
            report_context,
            &Ocr3ReportForExecutionReportSingleChain(&execution_report),
            &[],
        )?;
    }

    internal_execute(ctx, execution_report)
}

pub fn manually_execute<'info>(
    ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: ExecutionReportSingleChain,
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
            CcipRouterError::ManualExecutionNotAllowed
        );
    }
    internal_execute(ctx, execution_report)
}

/////////////
// Helpers //
/////////////

fn apply_token_price_update<'info>(
    token_update: &TokenPriceUpdate,
    token_config_account_info: &'info AccountInfo<'info>,
) -> Result<()> {
    let (expected, _) = Pubkey::find_program_address(
        &[FEE_BILLING_TOKEN_CONFIG, token_update.source_token.as_ref()],
        &crate::ID,
    );
    require_keys_eq!(
        token_config_account_info.key(),
        expected,
        CcipRouterError::InvalidInputs
    );

    require!(
        token_config_account_info.is_writable,
        CcipRouterError::InvalidInputs
    );

    let token_config_account: &mut Account<BillingTokenConfigWrapper> =
        &mut Account::try_from(token_config_account_info)?;

    require!(
        token_config_account.version == 1,
        CcipRouterError::InvalidInputs
    );

    token_config_account.config.usd_per_token = TimestampedPackedU224 {
        value: token_update.usd_per_token,
        timestamp: Clock::get()?.unix_timestamp,
    };

    emit!(UsdPerTokenUpdated {
        token: token_config_account.config.mint,
        value: token_config_account.config.usd_per_token.value,
        timestamp: token_config_account.config.usd_per_token.timestamp,
    });

    // As the account is manually loaded from the AccountInfo, it also needs to be manually
    // written back to so the changes are persisted.
    token_config_account.exit(&crate::ID)
}

fn apply_gas_price_update<'info>(
    gas_update: &GasPriceUpdate,
    dest_chain_state_account_info: &'info AccountInfo<'info>,
) -> Result<()> {
    let (expected, _) = Pubkey::find_program_address(
        &[
            DEST_CHAIN_STATE_SEED,
            gas_update.dest_chain_selector.to_le_bytes().as_ref(),
        ],
        &crate::ID,
    );
    require_keys_eq!(
        dest_chain_state_account_info.key(),
        expected,
        CcipRouterError::InvalidInputs
    );

    require!(
        dest_chain_state_account_info.is_writable,
        CcipRouterError::InvalidInputs
    );

    // The passed-in chain_state account may refer to the same chain but it only corresponds to source.
    // To update the price that values correspond to the destination, which is a different account.
    // As the account is sent as additional accounts, then Anchor won't automatically (de)serialize the account
    // as it is not the one in the context, so we have to do it manually load it and write it back
    let dest_chain_state_account = &mut Account::try_from(dest_chain_state_account_info)?;
    update_chain_state_gas_price(dest_chain_state_account, gas_update)?;
    dest_chain_state_account.exit(&crate::ID)?;
    Ok(())
}

fn update_chain_state_gas_price(
    chain_state_account: &mut Account<DestChain>,
    gas_update: &GasPriceUpdate,
) -> Result<()> {
    require!(
        chain_state_account.version == 1,
        CcipRouterError::InvalidInputs
    );

    chain_state_account.state.usd_per_unit_gas = TimestampedPackedU224 {
        value: gas_update.usd_per_unit_gas,
        timestamp: Clock::get()?.unix_timestamp,
    };

    emit!(UsdPerUnitGasUpdated {
        dest_chain: gas_update.dest_chain_selector,
        value: chain_state_account.state.usd_per_unit_gas.value,
        timestamp: chain_state_account.state.usd_per_unit_gas.timestamp,
    });

    Ok(())
}

// internal_execute is the base execution logic without any additional validation
fn internal_execute<'info>(
    ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: ExecutionReportSingleChain,
) -> Result<()> {
    // TODO: Limit send size data to 256

    // The Config Account stores the default values for the Router, the Solana Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
    let config = ctx.accounts.config.load()?;
    let solana_chain_selector = config.solana_chain_selector;

    // The Config and State for the Source Chain, containing if it is enabled, the on ramp address and the min sequence number expected for future messages
    let source_chain_state = &ctx.accounts.source_chain_state;
    require!(
        is_on_ramp_configured(
            &source_chain_state.config,
            &execution_report.message.on_ramp_address
        ),
        CcipRouterError::InvalidInputs
    );

    // The Commit Report Account stores the information of 1 Commit Report:
    // - Merkle Root
    // - Timestamp of the Commit Report
    // - Interval of Messages: The min and max seq num of the messages in the Merkle Tree
    // - Execution State per each Message: 0 for Untouched, 1 for InProgress, 2 for Success and 3 for Failure
    let commit_report = &mut ctx.accounts.commit_report;

    let message_header = execution_report.message.header;

    validate_execution_report(
        &execution_report,
        source_chain_state,
        commit_report,
        &message_header,
        solana_chain_selector,
    )?;

    let original_state = execution_state::get(commit_report, message_header.sequence_number);

    if original_state == MessageExecutionState::Success {
        emit!(SkippedAlreadyExecutedMessage {
            source_chain_selector: message_header.source_chain_selector,
            sequence_number: message_header.sequence_number,
        });
        return Ok(());
    }

    let hashed_leaf = verify_merkle_root(&execution_report)?;

    // send tokens any -> SOL
    require!(
        execution_report.token_indexes.len() == execution_report.message.token_amounts.len()
            && execution_report.token_indexes.len() == execution_report.offchain_token_data.len(),
        CcipRouterError::InvalidInputs,
    );
    let seeds = &[EXTERNAL_TOKEN_POOL_SEED, &[ctx.bumps.token_pools_signer]];
    let mut token_amounts =
        vec![SolanaTokenAmount::default(); execution_report.token_indexes.len()];

    // handle tokens
    // note: indexes are used instead of counts in case more accounts need to be passed in remaining_accounts before token accounts
    // token_indexes = [2, 4] where remaining_accounts is [custom_account, custom_account, token1_account1, token1_account2, token2_account1, token2_account2] for example
    for (i, token_amount) in execution_report.message.token_amounts.iter().enumerate() {
        let (start, end) = calculate_token_pool_account_indices(
            i,
            &execution_report.token_indexes,
            ctx.remaining_accounts.len(),
        )?;
        let acc_list = &ctx.remaining_accounts[start..end];
        let accs = validate_and_parse_token_accounts(
            execution_report.message.receiver,
            execution_report.message.header.source_chain_selector,
            ctx.program_id.key(),
            acc_list,
        )?;
        let router_token_pool_signer = &ctx.accounts.token_pools_signer;

        let init_bal = get_balance(accs.user_token_account)?;

        // CPI: call lockOrBurn on token pool
        let release_or_mint = ReleaseOrMintInV1 {
            original_sender: execution_report.message.sender.clone(),
            receiver: execution_report.message.receiver,
            amount: token_amount.amount,
            local_token: token_amount.dest_token_address,
            remote_chain_selector: execution_report.message.header.source_chain_selector,
            source_pool_address: token_amount.source_pool_address.clone(),
            source_pool_data: token_amount.extra_data.clone(),
            offchain_token_data: execution_report.offchain_token_data[i].clone(),
        };
        let mut acc_infos = router_token_pool_signer.to_account_infos();
        acc_infos.extend_from_slice(&[
            accs.pool_config.to_account_info(),
            accs.token_program.to_account_info(),
            accs.mint.to_account_info(),
            accs.pool_signer.to_account_info(),
            accs.pool_token_account.to_account_info(),
            accs.pool_chain_config.to_account_info(),
            accs.user_token_account.to_account_info(),
        ]);
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
            CcipRouterError::OfframpInvalidDataLength
        );

        token_amounts[i] = SolanaTokenAmount {
            token: accs.mint.key(),
            amount: ReleaseOrMintOutV1::try_from_slice(&return_data)?.destination_amount,
        };

        let post_bal = get_balance(accs.user_token_account)?;
        require!(
            post_bal >= init_bal && post_bal - init_bal == token_amounts[i].amount,
            CcipRouterError::OfframpReleaseMintBalanceMismatch
        );
    }

    let message = Any2SolanaMessage {
        message_id: execution_report.message.header.message_id,
        source_chain_selector: execution_report.source_chain_selector,
        sender: execution_report.message.sender,
        data: execution_report.message.data,
        token_amounts,
    };

    // handle CPI call if there are extra accounts
    // case: no tokens, but there are remaining_accounts passed in
    // case: tokens, but the first token has a non-zero index (indicating extra accounts before token accounts)
    if should_execute_messaging(
        &execution_report.token_indexes,
        ctx.remaining_accounts.is_empty(),
    ) {
        let (msg_program, msg_accounts) = parse_messaging_accounts(
            &execution_report.token_indexes,
            execution_report.message.receiver,
            &execution_report.message.extra_args.accounts,
            &execution_report.message.extra_args.is_writable_bitmap,
            ctx.remaining_accounts,
        )?;

        // The External Execution Config Account is used to sign the CPI instruction
        let external_execution_config = &ctx.accounts.external_execution_config;

        // The accounts of the user that will be used in the CPI instruction, none of them are signers
        // They need to specify if mutable or not, but none of them is allowed to init, realloc or close
        // note: CPI signer is always first account
        let mut acc_infos = external_execution_config.to_account_infos();
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

        let data = build_receiver_discriminator_and_data(message)?;

        let instruction = Instruction {
            program_id: msg_program.key(), // The receiver Program Id that will handle the ccip_receive message
            accounts: acc_metas,
            data,
        };

        let seeds = &[
            EXTERNAL_EXECUTION_CONFIG_SEED,
            &[ctx.bumps.external_execution_config],
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
        message_hash: hashed_leaf,             // Hash of the message using Solana encoding
        state: new_state,
    });

    Ok(())
}

// should_execute_messaging checks if there remaining_accounts that are not being used for token pools
// case: no tokens, but there are remaining_accounts passed in
// case: tokens, but the first token has a non-zero index (indicating extra accounts before token accounts)
fn should_execute_messaging(token_indexes: &[u8], remaining_accounts_empty: bool) -> bool {
    (token_indexes.is_empty() && !remaining_accounts_empty)
        || (!token_indexes.is_empty() && token_indexes[0] != 0)
}

/// parse_message_accounts returns all the accounts needed to execute the CPI instruction
/// It also validates that the accounts sent in the message match the ones sent in the source chain
///
/// # Arguments
/// * `token_indexes` - start indexes of token pool accounts, used to determine ending index for arbitrary messaging accounts
/// * `receiver` - receiver address from x-chain message, used to validate `accounts`
/// * `source_accounts` - arbitrary messaging accounts from the x-chain message, used to validate `accounts`. expected order is: [program, ...additional message accounts]
/// * `accounts` - accounts passed via `ctx.remaining_accounts`. expected order is: [program, receiver, ...additional message accounts]
fn parse_messaging_accounts<'info>(
    token_indexes: &[u8],
    receiver: Pubkey,
    source_accounts: &[Pubkey],
    source_bitmap: &u64,
    accounts: &'info [AccountInfo<'info>],
) -> Result<(&'info AccountInfo<'info>, &'info [AccountInfo<'info>])> {
    let end_ind = if token_indexes.is_empty() {
        accounts.len()
    } else {
        token_indexes[0] as usize
    };

    let msg_program = &accounts[0];
    let msg_accounts = &accounts[1..end_ind];

    let source_program = &source_accounts[0];
    let source_msg_accounts = &source_accounts[1..source_accounts.len()];

    require!(
        *source_program == msg_program.key(),
        CcipRouterError::InvalidInputs,
    );

    require!(
        msg_accounts[0].key() == receiver,
        CcipRouterError::InvalidInputs
    );

    // assert same number of accounts passed from message and transaction (not including program)
    // source_msg_accounts + 1 to account for separately passed receiver address
    require!(
        source_msg_accounts.len() + 1 == msg_accounts.len(),
        CcipRouterError::InvalidInputs
    );

    // Validate the addresses of all the accounts match the ones in source chain
    if msg_accounts.len() > 1 {
        // Ignore the first account as it's the receiver
        let accounts_to_validate = &msg_accounts[1..msg_accounts.len()];
        require!(
            accounts_to_validate.len() == source_msg_accounts.len(),
            CcipRouterError::InvalidInputs
        );
        for (i, acc) in source_msg_accounts.iter().enumerate() {
            let current_acc = &msg_accounts[i + 1]; // TODO: remove offset by 1 to skip receiver after receiver refactor
            require!(*acc == current_acc.key(), CcipRouterError::InvalidInputs);
            require!(
                // TODO: remove offset by 1 to skip program after receiver refactor
                is_writable(source_bitmap, (i + 1) as u8) == current_acc.is_writable,
                CcipRouterError::InvalidInputs
            );
        }
    }

    Ok((msg_program, msg_accounts))
}

/// Build the instruction data (discriminator + any other data)
fn build_receiver_discriminator_and_data(ramp_message: Any2SolanaMessage) -> Result<Vec<u8>> {
    let m: std::result::Result<Vec<u8>, std::io::Error> = ramp_message.try_to_vec();
    require!(m.is_ok(), CcipRouterError::InvalidMessage);
    let message = m.unwrap();

    let mut data = Vec::with_capacity(8);
    data.extend_from_slice(&CCIP_RECEIVE_DISCRIMINATOR);
    data.extend_from_slice(&message);

    Ok(data)
}

pub fn verify_merkle_root(execution_report: &ExecutionReportSingleChain) -> Result<[u8; 32]> {
    let hashed_leaf = hash(&execution_report.message);
    let verified_root: std::result::Result<[u8; 32], MerkleError> =
        calculate_merkle_root(hashed_leaf, execution_report.proofs.clone());
    require!(
        verified_root.is_ok() && verified_root.unwrap() == execution_report.root,
        CcipRouterError::InvalidProof
    );
    Ok(hashed_leaf)
}

// TODO: Refactor this to use the same structure as messages: execution_report.validate(..)
pub fn validate_execution_report<'info>(
    execution_report: &ExecutionReportSingleChain,
    source_chain_state: &Account<'info, SourceChain>,
    commit_report: &Account<'info, CommitReport>,
    message_header: &RampMessageHeader,
    solana_chain_selector: u64,
) -> Result<()> {
    require!(
        execution_report.message.header.nonce == 0,
        CcipRouterError::InvalidInputs
    );

    require!(
        source_chain_state.config.is_enabled,
        CcipRouterError::UnsupportedSourceChainSelector
    );

    require!(
        execution_report.message.header.sequence_number >= commit_report.min_msg_nr
            && execution_report.message.header.sequence_number <= commit_report.max_msg_nr,
        CcipRouterError::InvalidSequenceInterval
    );

    require!(
        message_header.source_chain_selector == execution_report.source_chain_selector,
        CcipRouterError::UnsupportedSourceChainSelector
    );
    require!(
        message_header.dest_chain_selector == solana_chain_selector,
        CcipRouterError::UnsupportedDestinationChainSelector
    );
    require!(
        commit_report.timestamp != 0,
        CcipRouterError::RootNotCommitted
    );

    Ok(())
}

fn hash(msg: &Any2SolanaRampMessage) -> [u8; 32] {
    use anchor_lang::solana_program::hash;

    // Calculate vectors size to ensure that the hash is unique
    let sender_size = [msg.sender.len() as u8];
    let on_ramp_address_size = [msg.on_ramp_address.len() as u8];
    let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

    // RampMessageHeader struct
    let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
    let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
    let header_sequence_number = msg.header.sequence_number.to_be_bytes();
    let header_nonce = msg.header.nonce.to_be_bytes();

    // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
    // As similar as https://github.com/smartcontractkit/chainlink/blob/d1a9f8be2f222ea30bdf7182aaa6428bfa605cf7/contracts/src/v0.8/ccip/libraries/Internal.sol#L111
    let result = hash::hashv(&[
        LEAF_DOMAIN_SEPARATOR.as_slice(),
        // metadata hash
        "Any2SolanaMessageHashV1".as_bytes(),
        &header_source_chain_selector,
        &header_dest_chain_selector,
        &on_ramp_address_size,
        &msg.on_ramp_address,
        // message header
        &msg.header.message_id,
        &msg.receiver.to_bytes(),
        &header_sequence_number,
        msg.extra_args.try_to_vec().unwrap().as_ref(), // borsh serialized
        &header_nonce,
        // message
        &sender_size,
        &msg.sender,
        &data_size.to_be_bytes(),
        &msg.data,
        // token transfers
        msg.token_amounts.try_to_vec().unwrap().as_ref(), // borsh serialized
    ]);

    result.to_bytes()
}

mod execution_state {
    use crate::{CommitReport, MessageExecutionState};

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
    use super::*;
    use crate::{Any2SolanaRampMessage, Any2SolanaTokenTransfer, SolanaExtraArgs};

    /// Builds a message and hash it, it's compared with a known hash
    #[test]
    fn test_hash() {
        let on_ramp_address = &[1, 2, 3].to_vec();

        let message = Any2SolanaRampMessage {
            sender: [
                1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0,
            ]
            .to_vec(),
            receiver: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb").unwrap(),
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
            token_amounts: [Any2SolanaTokenTransfer {
                source_pool_address: vec![0, 1, 2, 3],
                dest_token_address: Pubkey::try_from(
                    "DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc",
                )
                .unwrap(),
                dest_gas_amount: 100,
                extra_data: vec![4, 5, 6],
                amount: [1; 32],
            }]
            .to_vec(),
            extra_args: SolanaExtraArgs {
                compute_units: 1000,
                is_writable_bitmap: 1,
                accounts: vec![
                    Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb").unwrap(),
                ],
            },
            on_ramp_address: on_ramp_address.clone(),
        };
        let hash_result = hash(&message);

        assert_eq!(
            "1636e87682c7622432edefeccae977d0e64f30251eee1b02e02b7156a58dfebf",
            hex::encode(hash_result)
        );
    }
}
