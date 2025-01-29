use std::cell::Ref;

use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use super::fee_quoter::{fee_for_msg, transfer_fee, wrap_native_sol};
use super::merkle::LEAF_DOMAIN_SEPARATOR;
use super::messages::pools::{LockOrBurnInV1, LockOrBurnOutV1};
use super::pools::{
    calculate_token_pool_account_indices, interact_with_pool, transfer_token,
    validate_and_parse_token_accounts, TokenAccounts, CCIP_LOCK_OR_BURN_V1_RET_BYTES,
};
use super::price_math::get_validated_token_price;

use crate::seed;
use crate::{
    AnyExtraArgs, BillingTokenConfig, CCIPMessageSent, CcipRouterError, CcipSend, Config,
    DestChainConfig, ExtraArgsInput, Nonce, PerChainPerTokenConfig, RampMessageHeader,
    SVM2AnyMessage, SVM2AnyRampMessage, SVM2AnyTokenTransfer, SVMTokenAmount,
};

pub fn ccip_send<'info>(
    ctx: Context<'_, '_, 'info, 'info, CcipSend<'info>>,
    dest_chain_selector: u64,
    message: SVM2AnyMessage,
    token_indexes: Vec<u8>,
) -> Result<()> {
    // The Config Account stores the default values for the Router, the SVM Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
    let config = ctx.accounts.config.load()?;

    let dest_chain = &mut ctx.accounts.dest_chain_state;

    let mut accounts_per_sent_token: Vec<TokenAccounts> = vec![];

    for (i, token_amount) in message.token_amounts.iter().enumerate() {
        require!(
            token_amount.amount != 0,
            CcipRouterError::InvalidInputsTokenAmount
        );

        // Calculate the indexes for the additional accounts of the current token index `i`
        let (start, end) =
            calculate_token_pool_account_indices(i, &token_indexes, ctx.remaining_accounts.len())?;

        let current_token_accounts = validate_and_parse_token_accounts(
            ctx.accounts.authority.key(),
            dest_chain_selector,
            ctx.program_id.key(),
            &ctx.remaining_accounts[start..end],
        )?;

        accounts_per_sent_token.push(current_token_accounts);
    }

    let token_billing_config_accounts = accounts_per_sent_token
        .iter()
        .map(|accs| validated_try_to::billing_token_config(accs.fee_token_config, accs.mint.key()))
        .collect::<Result<Vec<_>>>()?;

    let per_chain_per_token_config_accounts = accounts_per_sent_token
        .iter()
        .map(|accs| {
            validated_try_to::per_chain_per_token_config(
                accs.token_billing_config,
                accs.mint.key(),
                dest_chain_selector,
            )
        })
        .collect::<Result<Vec<_>>>()?;

    let fee = fee_for_msg(
        &message,
        dest_chain,
        &ctx.accounts.fee_token_config.config,
        &token_billing_config_accounts,
        &per_chain_per_token_config_accounts,
    )?;

    let link_fee = convert(
        &fee,
        &ctx.accounts.fee_token_config.config,
        &ctx.accounts.link_token_config.config,
    )?;

    require_gte!(
        config.max_fee_juels_per_msg,
        link_fee.amount as u128,
        CcipRouterError::MessageFeeTooHigh,
    );

    let is_paying_with_native_sol = message.fee_token == Pubkey::default();
    if is_paying_with_native_sol {
        wrap_native_sol(
            &ctx.accounts.fee_token_program.to_account_info(),
            &mut ctx.accounts.authority,
            &mut ctx.accounts.fee_token_receiver,
            fee.amount,
            ctx.bumps.fee_billing_signer,
        )?;
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

        transfer_fee(
            &fee,
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

    let sender = ctx.accounts.authority.key.to_owned();
    let receiver = message.receiver.clone();
    let source_chain_selector = config.svm_chain_selector;
    let extra_args = extra_args_or_default(config, message.extra_args);

    let nonce_counter_account: &mut Account<'info, Nonce> = &mut ctx.accounts.nonce;
    let final_nonce = bump_nonce(nonce_counter_account, extra_args).unwrap();

    let token_count = message.token_amounts.len();
    require!(
        token_indexes.len() == token_count,
        CcipRouterError::InvalidInputs,
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
        extra_args,
        fee_token: message.fee_token,
        fee_token_amount: fee.amount.into(),
        fee_value_juels: link_fee.amount.into(),
        token_amounts: vec![SVM2AnyTokenTransfer::default(); token_count],
    };

    let seeds = &[seed::EXTERNAL_TOKEN_POOL, &[ctx.bumps.token_pools_signer]];
    for (i, (current_token_accounts, token_amount)) in accounts_per_sent_token
        .iter()
        .zip(message.token_amounts.iter())
        .enumerate()
    {
        let router_token_pool_signer = &ctx.accounts.token_pools_signer;

        // CPI: transfer token amount from user to token pool
        transfer_token(
            token_amount.amount,
            current_token_accounts.token_program,
            current_token_accounts.mint,
            current_token_accounts.user_token_account,
            current_token_accounts.pool_token_account,
            router_token_pool_signer,
            seeds,
        )?;

        // CPI: call lockOrBurn on token pool
        {
            let lock_or_burn = LockOrBurnInV1 {
                receiver: message.receiver.clone(),
                remote_chain_selector: dest_chain_selector,
                original_sender: ctx.accounts.authority.key(),
                amount: token_amount.amount,
                local_token: token_amount.token,
            };

            let mut acc_infos = router_token_pool_signer.to_account_infos();
            acc_infos.extend_from_slice(&[
                current_token_accounts.pool_config.to_account_info(),
                current_token_accounts.token_program.to_account_info(),
                current_token_accounts.mint.to_account_info(),
                current_token_accounts.pool_signer.to_account_info(),
                current_token_accounts.pool_token_account.to_account_info(),
                current_token_accounts.pool_chain_config.to_account_info(),
            ]);
            acc_infos.extend_from_slice(current_token_accounts.remaining_accounts);

            let return_data = interact_with_pool(
                current_token_accounts.pool_program.key(),
                router_token_pool_signer.key(),
                acc_infos,
                lock_or_burn,
                seeds,
            )?;

            let lock_or_burn_out_data = LockOrBurnOutV1::try_from_slice(&return_data)?;
            new_message.token_amounts[i] = token_transfer(
                lock_or_burn_out_data,
                current_token_accounts.pool_config.key(),
                token_amount,
                &dest_chain.config,
                &per_chain_per_token_config_accounts[i],
            )?;
        }
    }

    let message_id = &hash(&new_message);
    new_message.header.message_id.clone_from(message_id);

    emit!(CCIPMessageSent {
        dest_chain_selector,
        sequence_number: new_message.header.sequence_number,
        message: new_message,
    });

    Ok(())
}

fn token_transfer(
    lock_or_burn_out_data: LockOrBurnOutV1,
    source_pool_address: Pubkey,
    token_amount: &SVMTokenAmount,
    dest_chain_config: &DestChainConfig,
    token_config: &PerChainPerTokenConfig,
) -> Result<SVM2AnyTokenTransfer> {
    let dest_gas_amount = if token_config.billing.is_enabled {
        token_config.billing.dest_gas_overhead
    } else {
        dest_chain_config.default_token_dest_gas_overhead
    };

    let extra_data = lock_or_burn_out_data.dest_pool_data;
    let extra_data_length = extra_data.len() as u32;

    require!(
        extra_data_length <= CCIP_LOCK_OR_BURN_V1_RET_BYTES
            || extra_data_length <= token_config.billing.dest_bytes_overhead,
        CcipRouterError::SourceTokenDataTooLarge
    );

    // TODO: Revisit when/if non-EVM destinations from SVM become supported.
    // for an EVM destination, exec data it consists of the amount of gas available for the releaseOrMint
    // and transfer calls made by the offRamp
    let dest_exec_data = ethnum::U256::new(dest_gas_amount.into())
        .to_be_bytes()
        .to_vec();

    Ok(SVM2AnyTokenTransfer {
        source_pool_address,
        dest_token_address: lock_or_burn_out_data.dest_token_address,
        extra_data,
        amount: token_amount.amount.into(), // pool on receiver chain handles decimals
        dest_exec_data,
    })
}

/////////////
// Helpers //
/////////////
fn extra_args_or_default(default_config: Ref<Config>, extra_args: ExtraArgsInput) -> AnyExtraArgs {
    let mut result_args = AnyExtraArgs {
        gas_limit: default_config.default_gas_limit.to_owned(),
        allow_out_of_order_execution: default_config.default_allow_out_of_order_execution != 0,
    };

    if let Some(gas_limit) = extra_args.gas_limit {
        gas_limit.clone_into(&mut result_args.gas_limit)
    }

    if let Some(allow_out_of_order_execution) = extra_args.allow_out_of_order_execution {
        allow_out_of_order_execution.clone_into(&mut result_args.allow_out_of_order_execution)
    }

    result_args
}

fn bump_nonce(nonce_counter_account: &mut Account<Nonce>, extra_args: AnyExtraArgs) -> Result<u64> {
    // Avoid Re-initialization attack as the account is init_if_needed
    // https://solana.com/developers/courses/program-security/reinitialization-attacks#add-is-initialized-check
    if nonce_counter_account.version == 0 {
        // The authority must be the owner of the PDA, as it's their Public Key in the seed
        // If the account is not initialized, initialize it
        nonce_counter_account.version = 1;
        nonce_counter_account.counter = 0;
    }

    // TODO: Require config.enforce_out_of_order => extra_args.allow_out_of_order_execution
    let mut final_nonce = 0;
    if !extra_args.allow_out_of_order_execution {
        nonce_counter_account.counter = nonce_counter_account.counter.checked_add(1).unwrap();
        final_nonce = nonce_counter_account.counter;
    }
    Ok(final_nonce)
}

fn hash(msg: &SVM2AnyRampMessage) -> [u8; 32] {
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

// Converts a token amount to one denominated in another token (e.g. from WSOL to LINK)
pub fn convert(
    source_token_amount: &SVMTokenAmount,
    source_config: &BillingTokenConfig,
    target_config: &BillingTokenConfig,
) -> Result<SVMTokenAmount> {
    assert!(source_config.mint == source_token_amount.token);
    let source_price = get_validated_token_price(source_config)?;
    let target_price = get_validated_token_price(target_config)?;

    Ok(SVMTokenAmount {
        token: target_config.mint,
        amount: ((source_price * source_token_amount.amount).0 / target_price.0)
            .try_into()
            .map_err(|_| CcipRouterError::InvalidTokenPrice)?,
    })
}

/// Methods in this module are used to deserialize AccountInfo into the state structs
mod validated_try_to {
    use anchor_lang::prelude::*;

    use crate::{
        seed::{self},
        BillingTokenConfig, BillingTokenConfigWrapper, CcipRouterError, PerChainPerTokenConfig,
    };

    pub fn per_chain_per_token_config<'info>(
        account: &'info AccountInfo<'info>,
        token: Pubkey,
        dest_chain_selector: u64,
    ) -> Result<PerChainPerTokenConfig> {
        let (expected, _) = Pubkey::find_program_address(
            &[
                seed::TOKEN_POOL_BILLING,
                dest_chain_selector.to_le_bytes().as_ref(),
                token.key().as_ref(),
            ],
            &crate::ID,
        );
        require_keys_eq!(account.key(), expected, CcipRouterError::InvalidInputs);
        let account = Account::<PerChainPerTokenConfig>::try_from(account)?;
        require_eq!(
            account.version,
            1, // the v1 version of the onramp will always be tied to version 1 of the state
            CcipRouterError::InvalidInputs
        );
        Ok(account.into_inner())
    }

    // Returns Ok(None) when parsing the ZERO address, which is a valid input from users
    // specifying a token that has no Billing config.
    pub fn billing_token_config<'info>(
        account: &'info AccountInfo<'info>,
        token: Pubkey,
    ) -> Result<Option<BillingTokenConfig>> {
        if account.key() == Pubkey::default() {
            return Ok(None);
        }

        let (expected, _) = Pubkey::find_program_address(
            &[seed::FEE_BILLING_TOKEN_CONFIG, token.as_ref()],
            &crate::ID,
        );
        require_keys_eq!(account.key(), expected, CcipRouterError::InvalidInputs);
        let account = Account::<BillingTokenConfigWrapper>::try_from(account)?;
        require_eq!(
            account.version,
            1, // the v1 version of the onramp will always be tied to version 1 of the state
            CcipRouterError::InvalidInputs
        );
        Ok(Some(account.into_inner().config))
    }
}

#[cfg(test)]
mod tests {
    use ethnum::U256;

    use crate::state::TokenBilling;

    use super::super::messages::ramps::tests::{sample_billing_config, sample_dest_chain};

    pub fn sample_additional_token() -> (BillingTokenConfig, PerChainPerTokenConfig) {
        let mint = Pubkey::new_unique();
        (
            sample_billing_config(),
            PerChainPerTokenConfig {
                version: 1,
                chain_selector: 0,
                mint,
                billing: TokenBilling {
                    min_fee_usdcents: 50,
                    max_fee_usdcents: 4294967295,
                    deci_bps: 0,
                    dest_gas_overhead: 180000,
                    dest_bytes_overhead: 640,
                    is_enabled: true,
                },
            },
        )
    }

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
                1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0,
            ]
            .to_vec(),
            extra_args: AnyExtraArgs {
                gas_limit: 1,
                allow_out_of_order_execution: true,
            },
            fee_token: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb").unwrap(),
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
            "009bc51872fe41ea096bd881bf52e3daf07c80e112ffeeba6aa503d8281b6bfd",
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

        let dest_chain = sample_dest_chain();
        let (token_billing_config, mut per_chain_per_token_config) = sample_additional_token();
        let token_amount = &SVMTokenAmount {
            token: token_billing_config.mint,
            amount: 100,
        };

        let transfer = token_transfer(
            lock_or_burn_out_data.clone(),
            source_pool_address,
            token_amount,
            &dest_chain.config,
            &per_chain_per_token_config,
        )
        .unwrap();

        let expected_exec_data =
            ethnum::U256::new(per_chain_per_token_config.billing.dest_gas_overhead.into())
                .to_be_bytes();

        assert!(transfer.extra_data.is_empty());
        assert_eq!(transfer.dest_exec_data, expected_exec_data);

        // If we now disable billing overrides, the gas overhead will change
        per_chain_per_token_config.billing.is_enabled = false;
        let expected_exec_data =
            ethnum::U256::new(dest_chain.config.default_token_dest_gas_overhead.into())
                .to_be_bytes();

        let transfer = token_transfer(
            lock_or_burn_out_data,
            source_pool_address,
            token_amount,
            &dest_chain.config,
            &per_chain_per_token_config,
        )
        .unwrap();
        assert!(transfer.extra_data.is_empty());
        assert_eq!(transfer.dest_exec_data, expected_exec_data);
    }

    #[test]
    fn token_transfer_validates_data_length() {
        let dest_chain = sample_dest_chain();
        let (token_billing_config, per_chain_per_token_config) = sample_additional_token();
        let token_amount = &SVMTokenAmount {
            token: token_billing_config.mint,
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
            &dest_chain.config,
            &per_chain_per_token_config,
        )
        .unwrap();

        let unreasonable_size = (CCIP_LOCK_OR_BURN_V1_RET_BYTES
            .max(per_chain_per_token_config.billing.dest_gas_overhead)
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
                &dest_chain.config,
                &per_chain_per_token_config,
            )
            .unwrap_err(),
            CcipRouterError::SourceTokenDataTooLarge.into()
        );
    }
}
