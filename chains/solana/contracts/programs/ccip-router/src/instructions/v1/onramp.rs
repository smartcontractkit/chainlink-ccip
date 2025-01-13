use std::cell::Ref;

use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use super::fee_quoter::{fee_for_msg, transfer_fee, wrap_native_sol};
use super::messages::pools::{LockOrBurnInV1, LockOrBurnOutV1};
use super::pools::{
    calculate_token_pool_account_indices, interact_with_pool, transfer_token, u64_to_le_u256,
    validate_and_parse_token_accounts, TokenAccounts,
};

use crate::{
    AnyExtraArgs, CCIPMessageSent, CcipRouterError, CcipSend, Config, ExtraArgsInput, GetFee,
    Nonce, RampMessageHeader, Solana2AnyMessage, Solana2AnyRampMessage, Solana2AnyTokenTransfer,
    SolanaTokenAmount, EXTERNAL_TOKEN_POOL_SEED,
};

pub fn get_fee<'info>(
    ctx: Context<'_, '_, 'info, 'info, GetFee>,
    dest_chain_selector: u64,
    message: Solana2AnyMessage,
) -> Result<u64> {
    let remaining_accounts = &ctx.remaining_accounts;
    let message = &message;
    require_eq!(
        remaining_accounts.len(),
        2 * message.token_amounts.len(),
        CcipRouterError::InvalidInputsTokenAccounts
    );

    let (token_billing_config_accounts, per_chain_per_token_config_accounts) =
        remaining_accounts.split_at(message.token_amounts.len());

    let token_billing_config_accounts = token_billing_config_accounts
        .iter()
        .zip(message.token_amounts.iter())
        .map(|(a, SolanaTokenAmount { token, .. })| {
            validated_try_to::billing_token_config(a, *token)
        })
        .collect::<Result<Vec<_>>>()?;
    let per_chain_per_token_config_accounts = per_chain_per_token_config_accounts
        .iter()
        .zip(message.token_amounts.iter())
        .map(|(a, SolanaTokenAmount { token, .. })| {
            validated_try_to::per_chain_per_token_config(a, *token, dest_chain_selector)
        })
        .collect::<Result<Vec<_>>>()?;

    Ok(fee_for_msg(
        dest_chain_selector,
        message,
        &ctx.accounts.dest_chain_state,
        &ctx.accounts.billing_token_config.config,
        &token_billing_config_accounts,
        &per_chain_per_token_config_accounts,
    )?
    .amount)
}

pub fn ccip_send<'info>(
    ctx: Context<'_, '_, 'info, 'info, CcipSend<'info>>,
    dest_chain_selector: u64,
    message: Solana2AnyMessage,
) -> Result<()> {
    // The Config Account stores the default values for the Router, the Solana Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
    let config = ctx.accounts.config.load()?;

    let dest_chain = &mut ctx.accounts.dest_chain_state;

    let mut accounts_per_sent_token: Vec<TokenAccounts> = vec![];

    for (i, token_amount) in message.token_amounts.iter().enumerate() {
        require!(
            token_amount.amount != 0,
            CcipRouterError::InvalidInputsTokenAmount
        );

        // Calculate the indexes for the additional accounts of the current token index `i`
        let (start, end) = calculate_token_pool_account_indices(
            i,
            &message.token_indexes,
            ctx.remaining_accounts.len(),
        )?;

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
        dest_chain_selector,
        &message,
        dest_chain,
        &ctx.accounts.fee_token_config.config,
        &token_billing_config_accounts,
        &per_chain_per_token_config_accounts,
    )?;

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
            fee,
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
    let source_chain_selector = config.solana_chain_selector;
    let extra_args = extra_args_or_default(config, message.extra_args);

    let nonce_counter_account: &mut Account<'info, Nonce> = &mut ctx.accounts.nonce;
    let final_nonce = bump_nonce(nonce_counter_account, extra_args).unwrap();

    let token_count = message.token_amounts.len();
    require!(
        message.token_indexes.len() == token_count,
        CcipRouterError::InvalidInputs,
    );

    let mut new_message: Solana2AnyRampMessage = Solana2AnyRampMessage {
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
        token_amounts: vec![Solana2AnyTokenTransfer::default(); token_count],
    };

    let seeds = &[EXTERNAL_TOKEN_POOL_SEED, &[ctx.bumps.token_pools_signer]];
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

            let data = LockOrBurnOutV1::try_from_slice(&return_data)?;
            new_message.token_amounts[i] = Solana2AnyTokenTransfer {
                source_pool_address: current_token_accounts.pool_config.key(),
                dest_token_address: data.dest_token_address,
                extra_data: data.dest_pool_data,
                amount: u64_to_le_u256(token_amount.amount), // pool on receiver chain handles decimals
                dest_exec_data: vec![0; 0],                  // TODO: Implement this
                                                             // Destination chain specific execution data encoded in bytes
                                                             // for an EVM destination, it consists of the amount of gas available for the releaseOrMint
                                                             // and transfer calls made by the offRamp
            };
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

fn hash(msg: &Solana2AnyRampMessage) -> [u8; 32] {
    // TODO: Modify this hash to be similar to the one in EVM
    // https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/libraries/Internal.sol#L129
    // Fixed-size message fields are included in nested hash to reduce stack pressure.
    // - metadata_hash =  sha256("Solana2AnyMessageHashV1", solana_chain_selector, dest_chain_selector, ccip_router_program_id))
    // - first_part = sha256(sender, sequence_number, nonce, fee_token, fee_token_amount)
    // - receiver
    // - message.data
    // - token_amounts
    // - extra_args

    use anchor_lang::solana_program::hash;

    // Push Data Size to ensure that the hash is unique
    let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

    // RampMessageHeader struct
    let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
    let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
    let header_sequence_number = msg.header.sequence_number.to_be_bytes();
    let header_nonce = msg.header.nonce.to_be_bytes();

    // Extra Args struct
    let extra_args_gas_limit = msg.extra_args.gas_limit.to_be_bytes();
    let extra_args_allow_out_of_order_execution =
        [msg.extra_args.allow_out_of_order_execution as u8];

    // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
    let result = hash::hashv(&[
        &msg.sender.to_bytes(),
        &msg.receiver,
        &data_size.to_be_bytes(),
        &msg.data,
        &header_source_chain_selector,
        &header_dest_chain_selector,
        &header_sequence_number,
        &header_nonce,
        &extra_args_gas_limit,
        &extra_args_allow_out_of_order_execution,
    ]);

    result.to_bytes()
}

/// Methods in this module are used to deserialize AccountInfo into the state structs
mod validated_try_to {
    use anchor_lang::prelude::*;

    use crate::{
        BillingTokenConfig, BillingTokenConfigWrapper, CcipRouterError, PerChainPerTokenConfig,
        FEE_BILLING_TOKEN_CONFIG, TOKEN_POOL_BILLING_SEED,
    };

    pub fn per_chain_per_token_config<'info>(
        account: &'info AccountInfo<'info>,
        token: Pubkey,
        dest_chain_selector: u64,
    ) -> Result<PerChainPerTokenConfig> {
        let (expected, _) = Pubkey::find_program_address(
            &[
                TOKEN_POOL_BILLING_SEED,
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

        let (expected, _) =
            Pubkey::find_program_address(&[FEE_BILLING_TOKEN_CONFIG, token.as_ref()], &crate::ID);
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
