use anchor_lang::prelude::*;
use anchor_spl::token::spl_token;
use anchor_spl::token_interface;

use crate::context::ANCHOR_DISCRIMINATOR;
use crate::events::on_ramp as events;
use crate::messages::GetFeeResult;

use ccip_common::seed;
use ccip_common::v1::{validate_and_parse_token_accounts, TokenAccounts};
use fee_quoter::messages::TokenTransferAdditionalData;
use fee_quoter::state::DestChain;

use super::super::interfaces::OnRamp;
use super::fees::{get_fee_cpi, transfer_and_wrap_native_sol, transfer_fee};
use super::messages::pools::{LockOrBurnInV1, LockOrBurnOutV1};
use super::pools::{
    calculate_token_pool_account_indices, interact_with_pool, transfer_token,
    CCIP_LOCK_OR_BURN_V1_RET_BYTES,
};

use crate::GetFee;
use crate::{
    CcipRouterError, CcipSend, Nonce, RampMessageHeader, SVM2AnyMessage, SVM2AnyRampMessage,
    SVM2AnyTokenTransfer, SVMTokenAmount,
};

use helpers::*;

pub struct Impl;
impl OnRamp for Impl {
    fn ccip_send<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, CcipSend<'info>>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
        token_indexes: Vec<u8>,
    ) -> Result<[u8; 32]> {
        helpers::verify_uncursed_cpi(&ctx, dest_chain_selector)?;

        let mut message = message.clone();

        let sender = ctx.accounts.authority.key.to_owned();
        let dest_chain = &mut ctx.accounts.dest_chain_state;

        require!(
            !dest_chain.config.allow_list_enabled
                || dest_chain.config.allowed_senders.contains(&sender),
            CcipRouterError::SenderNotAllowed
        );

        let mut accounts_per_sent_token: Vec<TokenAccounts> = vec![];

        for (i, token_amount) in message.token_amounts.iter().enumerate() {
            require!(
                token_amount.amount != 0,
                CcipRouterError::InvalidInputsTokenAmount
            );

            // Calculate the indexes for the additional accounts of the current token index `i`
            let (start, end) = calculate_token_pool_account_indices(
                i,
                &token_indexes,
                ctx.remaining_accounts.len(),
            )?;

            let current_token_accounts = validate_and_parse_token_accounts(
                ctx.accounts.authority.key(),
                dest_chain_selector,
                ctx.program_id.key(),
                ctx.accounts.config.fee_quoter,
                None,
                &ctx.remaining_accounts[start..end],
            )?;

            accounts_per_sent_token.push(current_token_accounts);
        }

        let billing_token_config_accs: Vec<AccountInfo<'info>> = accounts_per_sent_token
            .iter()
            .map(|a| a.fee_token_config.to_account_info())
            .collect();
        let per_chain_per_token_config_accs: Vec<AccountInfo<'info>> = accounts_per_sent_token
            .iter()
            .map(|a| a.token_billing_config.to_account_info())
            .collect();

        let mut get_fee_remaining_accounts = billing_token_config_accs;
        get_fee_remaining_accounts.extend(per_chain_per_token_config_accs);

        let get_fee_result = get_fee_cpi(
            ctx.accounts.fee_quoter.to_account_info(),
            ctx.accounts.fee_quoter_config.to_account_info(),
            ctx.accounts.fee_quoter_dest_chain.to_account_info(),
            ctx.accounts
                .fee_quoter_billing_token_config
                .to_account_info(),
            ctx.accounts.fee_quoter_link_token_config.to_account_info(),
            dest_chain_selector,
            &message,
            get_fee_remaining_accounts,
        )?;

        let is_paying_with_native_sol = message.fee_token == Pubkey::default();
        if is_paying_with_native_sol {
            transfer_and_wrap_native_sol(
                &ctx.accounts.fee_token_program.to_account_info(),
                &mut ctx.accounts.authority,
                &mut ctx.accounts.fee_token_receiver,
                get_fee_result.amount,
                ctx.bumps.fee_billing_signer,
            )?;
            message.fee_token = spl_token::native_mint::id();
        } else {
            let transfer = token_interface::TransferChecked {
                from: ctx
                    .accounts
                    .fee_token_user_associated_account
                    .to_account_info(),
                to: ctx.accounts.fee_token_receiver.to_account_info(),
                mint: ctx.accounts.fee_token_mint.to_account_info(),
                authority: ctx.accounts.fee_billing_signer.to_account_info(),
            };

            let transferable_fee = SVMTokenAmount {
                token: get_fee_result.token,
                amount: get_fee_result.amount,
            };

            transfer_fee(
                &transferable_fee,
                ctx.accounts.fee_token_program.to_account_info(),
                transfer,
                ctx.accounts.fee_token_mint.decimals,
                ctx.bumps.fee_billing_signer,
            )?;
        }

        let overflow_add = dest_chain.state.sequence_number.checked_add(1);
        require!(
            overflow_add.is_some(),
            CcipRouterError::ReachedMaxSequenceNumber
        );
        dest_chain.state.sequence_number = overflow_add.unwrap();

        let receiver = message.receiver.clone();
        let source_chain_selector = ctx.accounts.config.svm_chain_selector;

        let nonce_counter_account: &mut Account<'info, Nonce> = &mut ctx.accounts.nonce;
        let final_nonce = bump_nonce(
            nonce_counter_account,
            get_fee_result
                .processed_extra_args
                .allow_out_of_order_execution,
        )
        .unwrap();

        let token_count = message.token_amounts.len();
        require!(
            token_indexes.len() == token_count,
            CcipRouterError::InvalidInputsTokenIndices,
        );

        let mut new_message: SVM2AnyRampMessage = SVM2AnyRampMessage {
            sender,
            receiver,
            data: message.data,
            header: RampMessageHeader {
                message_id: [0; 32],
                source_chain_selector,
                dest_chain_selector,
                sequence_number: dest_chain.state.sequence_number,
                nonce: final_nonce,
            },
            extra_args: get_fee_result.processed_extra_args.bytes,
            fee_token: message.fee_token,
            fee_token_amount: get_fee_result.amount.into(),
            fee_value_juels: get_fee_result.juels.into(),
            token_amounts: vec![SVM2AnyTokenTransfer::default(); token_count], // this will be set later
        };

        for (i, (current_token_accounts, token_amount)) in accounts_per_sent_token
            .iter()
            .zip(message.token_amounts.iter())
            .enumerate()
        {
            require_keys_eq!(
                token_amount.token,
                current_token_accounts.mint.key(),
                CcipRouterError::InvalidInputsTokenAccounts,
            );

            let transfer_seeds = &[seed::FEE_BILLING_SIGNER, &[ctx.bumps.fee_billing_signer]];

            // CPI: transfer token amount from user to token pool
            transfer_token(
                token_amount.amount,
                current_token_accounts.token_program,
                current_token_accounts.mint,
                current_token_accounts.user_token_account,
                current_token_accounts.pool_token_account,
                &ctx.accounts.fee_billing_signer.to_account_info(),
                transfer_seeds,
            )?;

            // CPI: call lockOrBurn on token pool
            {
                let lock_or_burn = LockOrBurnInV1 {
                    receiver: get_fee_result
                        .processed_extra_args
                        .token_receiver
                        .as_ref()
                        .unwrap_or(&message.receiver)
                        .clone(),
                    remote_chain_selector: dest_chain_selector,
                    original_sender: ctx.accounts.authority.key(),
                    amount: token_amount.amount,
                    local_token: token_amount.token,
                };

                let mut acc_infos = current_token_accounts
                    .ccip_router_pool_signer
                    .to_account_infos();
                acc_infos.extend_from_slice(&[
                    current_token_accounts.pool_config.to_account_info(),
                    current_token_accounts.token_program.to_account_info(),
                    current_token_accounts.mint.to_account_info(),
                    current_token_accounts.pool_signer.to_account_info(),
                    current_token_accounts.pool_token_account.to_account_info(),
                    ctx.accounts.rmn_remote.to_account_info(),
                    ctx.accounts.rmn_remote_curses.to_account_info(),
                    ctx.accounts.rmn_remote_config.to_account_info(),
                    current_token_accounts.pool_chain_config.to_account_info(),
                ]);
                acc_infos.extend_from_slice(current_token_accounts.remaining_accounts);

                let pool_seeds = &[
                    seed::EXTERNAL_TOKEN_POOLS_SIGNER,
                    current_token_accounts.pool_program.key.as_ref(),
                    &[current_token_accounts.ccip_router_pool_signer_bump],
                ];

                let return_data = interact_with_pool(
                    current_token_accounts.pool_program.key(),
                    current_token_accounts.ccip_router_pool_signer.key(),
                    acc_infos,
                    lock_or_burn,
                    pool_seeds,
                )?;

                let lock_or_burn_out_data = LockOrBurnOutV1::try_from_slice(&return_data)?;

                {
                    // validate the token address based on the destination chain family selector
                    let dest_chain_info = ctx.accounts.fee_quoter_dest_chain.to_account_info();
                    let dest_chain = DestChain::try_from_slice(
                        &dest_chain_info.data.borrow()[ANCHOR_DISCRIMINATOR..],
                    )?;
                    let dest_chain_family_selector = dest_chain.config.chain_family_selector;

                    helpers::validate_transfer_dest_address(
                        dest_chain_family_selector,
                        &lock_or_burn_out_data.dest_token_address,
                    )?;
                }

                new_message.token_amounts[i] = token_transfer(
                    lock_or_burn_out_data,
                    current_token_accounts.pool_config.key(),
                    token_amount,
                    &get_fee_result.token_transfer_additional_data[i],
                )?;
            }
        }

        let message_id = &helpers::hash(&new_message);
        new_message.header.message_id.clone_from(message_id);

        emit!(events::CCIPMessageSent {
            dest_chain_selector,
            sequence_number: new_message.header.sequence_number,
            message: new_message,
        });

        Ok(*message_id)
    }

    fn get_fee<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, GetFee<'info>>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
    ) -> Result<GetFeeResult> {
        let fq_result = get_fee_cpi(
            ctx.accounts.fee_quoter.to_account_info(),
            ctx.accounts.fee_quoter_config.to_account_info(),
            ctx.accounts.fee_quoter_dest_chain.to_account_info(),
            ctx.accounts
                .fee_quoter_billing_token_config
                .to_account_info(),
            ctx.accounts.fee_quoter_link_token_config.to_account_info(),
            dest_chain_selector,
            &message,
            ctx.remaining_accounts.to_vec(),
        )?;

        // not all fields that fee quoter returns are relevant for the user, so just pick the important ones
        Ok(GetFeeResult {
            amount: fq_result.amount,
            juels: fq_result.juels,
            token: fq_result.token,
        })
    }
}

mod helpers {
    use ccip_common::{
        v1::{validate_evm_address, validate_svm_address},
        CommonCcipError, CHAIN_FAMILY_SELECTOR_EVM, CHAIN_FAMILY_SELECTOR_SVM,
    };
    use rmn_remote::state::CurseSubject;

    use super::*;

    pub const LEAF_DOMAIN_SEPARATOR: [u8; 32] = [0; 32];

    pub fn verify_uncursed_cpi<'info>(
        ctx: &Context<'_, '_, 'info, 'info, CcipSend<'_>>,
        dest_chain_selector: u64,
    ) -> Result<()> {
        let cpi_program = ctx.accounts.rmn_remote.to_account_info();
        let cpi_accounts = rmn_remote::cpi::accounts::InspectCurses {
            config: ctx.accounts.rmn_remote_config.to_account_info(),
            curses: ctx.accounts.rmn_remote_curses.to_account_info(),
        };
        let cpi_context = CpiContext::new(cpi_program, cpi_accounts);
        rmn_remote::cpi::verify_not_cursed(
            cpi_context,
            CurseSubject::from_chain_selector(dest_chain_selector),
        )
    }

    pub(super) fn token_transfer(
        lock_or_burn_out_data: LockOrBurnOutV1,
        source_pool_address: Pubkey,
        token_amount: &SVMTokenAmount,
        additional_data: &TokenTransferAdditionalData,
    ) -> Result<SVM2AnyTokenTransfer> {
        let dest_gas_amount = additional_data.dest_gas_overhead;

        let extra_data = lock_or_burn_out_data.dest_pool_data;
        let extra_data_length = extra_data.len() as u32;

        require!(
            extra_data_length <= CCIP_LOCK_OR_BURN_V1_RET_BYTES
                || extra_data_length <= additional_data.dest_bytes_overhead,
            CcipRouterError::SourceTokenDataTooLarge
        );

        let dest_exec_data = dest_gas_amount.to_be_bytes().to_vec();

        Ok(SVM2AnyTokenTransfer {
            source_pool_address,
            dest_token_address: lock_or_burn_out_data.dest_token_address,
            extra_data,
            amount: token_amount.amount.into(), // pool on receiver chain handles decimals
            dest_exec_data,
        })
    }

    pub(super) fn bump_nonce(
        nonce_counter_account: &mut Account<Nonce>,
        allow_out_of_order_execution: bool,
    ) -> Result<u64> {
        // Avoid Re-initialization attack as the account is init_if_needed
        // https://solana.com/developers/courses/program-security/reinitialization-attacks#add-is-initialized-check
        if nonce_counter_account.version == 0 {
            // The authority must be the owner of the PDA, as it's their Public Key in the seed
            // If the account is not initialized, initialize it
            nonce_counter_account.version = 1;
            nonce_counter_account.counter = 0;
        }

        // config.enforce_out_of_order checked in `validate_svm2any`
        let mut final_nonce = 0;
        if !allow_out_of_order_execution {
            nonce_counter_account.counter = nonce_counter_account.counter.checked_add(1).unwrap();
            final_nonce = nonce_counter_account.counter;
        }
        Ok(final_nonce)
    }

    pub(super) fn hash(msg: &SVM2AnyRampMessage) -> [u8; 32] {
        use anchor_lang::solana_program::keccak;

        // Push Data Size to ensure that the hash is unique
        let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

        // RampMessageHeader struct
        let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
        let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
        let header_sequence_number = msg.header.sequence_number.to_be_bytes();
        let header_nonce = msg.header.nonce.to_be_bytes();

        // similar to: https://github.com/smartcontractkit/chainlink/blob/d1a9f8be2f222ea30bdf7182aaa6428bfa605cf7/contracts/src/v0.8/ccip/libraries/Internal.sol#L134
        let result = keccak::hashv(&[
            LEAF_DOMAIN_SEPARATOR.as_slice(),
            // metadata
            "SVM2AnyMessageHashV1".as_bytes(),
            &header_source_chain_selector,
            &header_dest_chain_selector,
            &crate::ID.to_bytes(), // onramp: ccip_router program
            // message header
            &msg.sender.to_bytes(),
            &header_sequence_number,
            &header_nonce,
            &msg.fee_token.to_bytes(),
            // The cross-chain amounts are encoded in little endian, but
            // this is irrelevant to the hashing function as long as both
            // sides agree.
            &msg.fee_token_amount.to_bytes(),
            &msg.fee_value_juels.to_bytes(),
            // messaging
            &[msg.receiver.len() as u8],
            &msg.receiver,
            &data_size.to_be_bytes(),
            &msg.data,
            // tokens
            &msg.token_amounts.try_to_vec().unwrap(),
            // extra args
            msg.extra_args.try_to_vec().unwrap().as_ref(), // borsh serialize
        ]);

        result.to_bytes()
    }

    pub(super) fn validate_transfer_dest_address(
        chain_family_selector: [u8; 4],
        dest_token_address: &[u8],
    ) -> Result<()> {
        let selector = u32::from_be_bytes(chain_family_selector);
        match selector {
            CHAIN_FAMILY_SELECTOR_EVM => validate_evm_address(dest_token_address),
            CHAIN_FAMILY_SELECTOR_SVM => validate_svm_address(dest_token_address, true),
            _ => Err(CommonCcipError::InvalidChainFamilySelector.into()),
        }
    }

    #[cfg(test)]
    mod tests {
        use ethnum::U256;

        use super::*;

        /// Builds a message and hash it, it's compared with a known hash
        #[test]
        fn test_hash() {
            let message = SVM2AnyRampMessage {
                header: RampMessageHeader {
                    message_id: [0; 32],
                    source_chain_selector: 10,
                    dest_chain_selector: 20,
                    sequence_number: 30,
                    nonce: 40,
                },
                sender: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTa").unwrap(),
                data: vec![4, 5, 6],
                receiver: [
                    1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                    0, 0, 0, 0, 0, 0,
                ]
                .to_vec(),
                extra_args: fee_quoter::extra_args::GenericExtraArgsV2 {
                    gas_limit: 1,
                    allow_out_of_order_execution: true,
                }
                .serialize_with_tag(),
                fee_token: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
                    .unwrap(),
                fee_token_amount: 50u32.into(),
                token_amounts: [SVM2AnyTokenTransfer {
                    source_pool_address: Pubkey::try_from(
                        "DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc",
                    )
                    .unwrap(),
                    dest_token_address: vec![0, 1, 2, 3],
                    extra_data: vec![4, 5, 6],
                    amount: U256::from_le_bytes([1; 32]).into(),
                    dest_exec_data: vec![4, 5, 6],
                }]
                .to_vec(),
                fee_value_juels: 500u32.into(),
            };

            let hash_result = hash(&message);

            assert_eq!(
                "46abd733e950b0b0e05b1b4b040cf2df6c38899af71aad058ff08bfa96d4b532",
                hex::encode(hash_result)
            );
        }

        #[test]
        fn token_transfer_with_no_pool_data() {
            let source_pool_address = Pubkey::new_unique();
            let lock_or_burn_out_data = LockOrBurnOutV1 {
                dest_token_address: Pubkey::new_unique().to_bytes().to_vec(),
                dest_pool_data: vec![],
            };

            let additional_token_transfer_data = TokenTransferAdditionalData {
                dest_bytes_overhead: 640,
                dest_gas_overhead: 180000,
            };
            let token_amount = &SVMTokenAmount {
                token: Pubkey::new_unique(),
                amount: 100,
            };

            let transfer = token_transfer(
                lock_or_burn_out_data.clone(),
                source_pool_address,
                token_amount,
                &additional_token_transfer_data,
            )
            .unwrap();

            let expected_exec_data = additional_token_transfer_data
                .dest_gas_overhead
                .to_be_bytes();

            assert!(transfer.extra_data.is_empty());
            assert_eq!(transfer.dest_exec_data, expected_exec_data);
        }

        #[test]
        fn token_transfer_validates_data_length() {
            let additional_token_transfer_data = TokenTransferAdditionalData {
                dest_bytes_overhead: 640,
                dest_gas_overhead: 180000,
            };
            let token_amount = &SVMTokenAmount {
                token: Pubkey::new_unique(),
                amount: 100,
            };

            let reasonable_size = CCIP_LOCK_OR_BURN_V1_RET_BYTES as usize;
            let source_pool_address = Pubkey::new_unique();
            let lock_or_burn_out_data = LockOrBurnOutV1 {
                dest_token_address: Pubkey::new_unique().to_bytes().to_vec(),
                dest_pool_data: vec![b'A'; reasonable_size],
            };

            token_transfer(
                lock_or_burn_out_data,
                source_pool_address,
                token_amount,
                &additional_token_transfer_data,
            )
            .unwrap();

            let unreasonable_size = (CCIP_LOCK_OR_BURN_V1_RET_BYTES
                .max(additional_token_transfer_data.dest_gas_overhead)
                + 1) as usize;
            let lock_or_burn_out_data = LockOrBurnOutV1 {
                dest_token_address: Pubkey::new_unique().to_bytes().to_vec(),
                dest_pool_data: vec![b'A'; unreasonable_size],
            };

            assert_eq!(
                token_transfer(
                    lock_or_burn_out_data,
                    source_pool_address,
                    token_amount,
                    &additional_token_transfer_data,
                )
                .unwrap_err(),
                CcipRouterError::SourceTokenDataTooLarge.into()
            );
        }
    }
}
