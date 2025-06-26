use std::{
    fmt::{self, Display, Formatter},
    str::FromStr,
};

use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id,
    token::spl_token::{self, native_mint},
};
use ccip_common::seed;
use ccip_common::v1::load_token_admin_registry_checked;
use ccip_common::CommonCcipError;
use solana_program::{
    address_lookup_table::state::AddressLookupTable,
    instruction::Instruction,
    program::{get_return_data, invoke},
};

use crate::{
    context::ViewConfigOnly,
    instructions::v1::{fees::get_fee_cpi, messages::pools::LockOrBurnInV1},
    messages::TOKENPOOL_DERIVE_LOCK_OR_BURN_DISCRIMINATOR,
    state::{CcipAccountMeta, DeriveAccountsCcipSendParams, DeriveAccountsResponse, ToMeta},
    CcipRouterError,
};

use super::helpers::load_nonce;

// Most token pools have four static addresses + the 10 elements of a LUT.
const TOKEN_ACCOUNTS_STATIC_PAGE_SIZE: usize = 14;

// Local helper to find a readonly CCIP meta for a given seed + program_id combo.
// Short name for compactness.
fn find(seeds: &[&[u8]], program_id: Pubkey) -> CcipAccountMeta {
    Pubkey::find_program_address(seeds, &program_id)
        .0
        .readonly()
}

#[derive(Clone, Debug)]
pub enum DeriveAccountsCcipSendStage {
    Start,
    FinishMainAccountList,
    RetrieveTokenLUTs,
    RetrievePoolPrograms,
    // N stages, one per token for the ones below.
    TokenTransferStaticAccounts {
        // Might be too many to fit in one response, so the user
        // may be required to request multiple pages.
        page: u32,
    },
    NestedTokenDerive {
        token_substage: String,
    },
}

impl Display for DeriveAccountsCcipSendStage {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        match self {
            DeriveAccountsCcipSendStage::Start => f.write_str("Start"),
            DeriveAccountsCcipSendStage::FinishMainAccountList => {
                f.write_str("FinishMainAccountList")
            }
            DeriveAccountsCcipSendStage::RetrieveTokenLUTs => {
                f.write_str("RetrieveTokenLookupTables")
            }
            DeriveAccountsCcipSendStage::RetrievePoolPrograms => {
                f.write_str("RetrievePoolPrograms")
            }
            DeriveAccountsCcipSendStage::TokenTransferStaticAccounts { page } => {
                f.write_fmt(format_args!("TokenTransferStaticAccounts/{page}"))
            }
            DeriveAccountsCcipSendStage::NestedTokenDerive { token_substage } => {
                f.write_fmt(format_args!("NestedTokenDerive/{token_substage}"))
            }
        }
    }
}

impl FromStr for DeriveAccountsCcipSendStage {
    type Err = CcipRouterError;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        let mut s = s.split('/');
        let (prefix, suffix) = (s.next(), s.next());

        match (prefix, suffix) {
            (Some("Start"), None) => Ok(Self::Start),
            (Some("FinishMainAccountList"), None) => Ok(Self::FinishMainAccountList),
            (Some("RetrieveTokenLookupTables"), None) => Ok(Self::RetrieveTokenLUTs),
            (Some("RetrievePoolPrograms"), None) => Ok(Self::RetrievePoolPrograms),
            (Some("TokenTransferStaticAccounts"), Some(page)) => {
                Ok(Self::TokenTransferStaticAccounts {
                    page: str::parse::<u32>(page)
                        .or(Err(CcipRouterError::InvalidDerivationStage))?,
                })
            }
            (Some("NestedTokenDerive"), token_substage) => Ok(Self::NestedTokenDerive {
                token_substage: token_substage.unwrap_or("Start").to_string(),
            }),
            _ => Err(CcipRouterError::InvalidDerivationStage),
        }
    }
}

pub fn derive_ccip_send_accounts_start(
    params: DeriveAccountsCcipSendParams,
) -> Result<DeriveAccountsResponse> {
    let selector = params.dest_chain_selector.to_le_bytes();
    let accounts_to_save = vec![
        find(&[seed::CONFIG], crate::ID),
        find(&[seed::DEST_CHAIN_STATE, &selector], crate::ID).writable(),
        find(
            &[
                seed::NONCE,
                &selector,
                params.ccip_send_caller.key().as_ref(),
            ],
            crate::ID,
        )
        .writable(),
        params.ccip_send_caller.writable().signer(),
        solana_program::system_program::ID.readonly(),
    ];

    let is_fee_in_sol = params.message.fee_token == Pubkey::default();

    let ask_again_with = vec![
        if is_fee_in_sol {
            spl_token::native_mint::id()
        } else {
            params.message.fee_token
        }
        .readonly(),
        find(&[seed::FEE_BILLING_SIGNER], crate::ID),
    ];

    Ok(DeriveAccountsResponse {
        accounts_to_save,
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsCcipSendStage::Start.to_string(),
        next_stage: DeriveAccountsCcipSendStage::FinishMainAccountList.to_string(),
    })
}

pub fn derive_ccip_send_accounts_finish_main_account_list<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: DeriveAccountsCcipSendParams,
) -> Result<DeriveAccountsResponse> {
    let selector = params.dest_chain_selector.to_le_bytes();
    let fee_token_mint_info = &ctx.remaining_accounts[0];
    let fee_billing_signer = &ctx.remaining_accounts[1];

    let is_fee_in_sol = params.message.fee_token == Pubkey::default();

    let (fee_token_user_ata, fee_token_program) = if is_fee_in_sol {
        (Pubkey::default().readonly(), spl_token::ID.readonly())
    } else {
        let fee_token_program = fee_token_mint_info.owner;
        (
            get_associated_token_address_with_program_id(
                &params.ccip_send_caller.key(),
                &params.message.fee_token,
                &fee_token_program.key(),
            )
            .writable(),
            fee_token_program.readonly(),
        )
    };

    let fee_token_receiver = get_associated_token_address_with_program_id(
        fee_billing_signer.key,
        fee_token_mint_info.key,
        &fee_token_program.pubkey,
    );

    let config = &ctx.accounts.config;

    let accounts_to_save = vec![
        fee_token_program.readonly(),
        fee_token_mint_info.key.readonly(),
        fee_token_user_ata,
        fee_token_receiver.writable(),
        fee_billing_signer.key.readonly(),
        config.fee_quoter.readonly(),
        find(&[seed::CONFIG], config.fee_quoter).readonly(),
        find(&[seed::DEST_CHAIN, &selector], config.fee_quoter).readonly(),
        find(
            &[
                seed::FEE_BILLING_TOKEN_CONFIG,
                fee_token_mint_info.key().as_ref(),
            ],
            config.fee_quoter,
        ),
        find(
            &[
                seed::FEE_BILLING_TOKEN_CONFIG,
                config.link_token_mint.key().as_ref(),
            ],
            config.fee_quoter,
        ),
        config.rmn_remote.readonly(),
        find(&[seed::CURSES], config.rmn_remote),
        find(&[seed::CONFIG], config.rmn_remote),
    ];

    // If there are no tokens, we're done.
    let (next_stage, ask_again_with) = if params.message.token_amounts.is_empty() {
        ("".to_string(), vec![])
    } else {
        let mints_of_transferred_tokens = params.message.token_amounts.iter().map(|ta| ta.token);
        let next_stage_accounts_per_transferred_token =
            mints_of_transferred_tokens.flat_map(|mint| {
                [
                    find(&[seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()], crate::ID),
                    find(
                        &[seed::FEE_BILLING_TOKEN_CONFIG, mint.as_ref()],
                        config.fee_quoter,
                    ),
                    find(
                        &[seed::PER_CHAIN_PER_TOKEN_CONFIG, &selector, mint.as_ref()],
                        config.fee_quoter,
                    ),
                ]
                .into_iter()
            });

        let mut ask_again_with = vec![
            config.fee_quoter.readonly(),
            find(&[seed::CONFIG], config.fee_quoter).readonly(),
            find(&[seed::DEST_CHAIN, &selector], config.fee_quoter).readonly(),
            find(
                &[
                    seed::FEE_BILLING_TOKEN_CONFIG,
                    if is_fee_in_sol {
                        native_mint::ID.as_ref() // pre-2022 WSOL
                    } else {
                        params.message.fee_token.as_ref()
                    },
                ],
                config.fee_quoter,
            ),
            find(
                &[
                    seed::FEE_BILLING_TOKEN_CONFIG,
                    config.link_token_mint.key().as_ref(),
                ],
                config.fee_quoter,
            ),
        ];
        ask_again_with.extend(next_stage_accounts_per_transferred_token);
        (
            DeriveAccountsCcipSendStage::RetrieveTokenLUTs.to_string(),
            ask_again_with,
        )
    };

    Ok(DeriveAccountsResponse {
        ask_again_with,
        accounts_to_save,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsCcipSendStage::FinishMainAccountList.to_string(),
        next_stage,
    })
}

pub fn derive_ccip_send_accounts_retrieve_luts<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsCcipSendParams,
) -> Result<DeriveAccountsResponse> {
    let fee_quoter_fixed_accounts_len = 5usize;
    let accounts_per_token_len = 3usize;

    // Accounts needed for get_fee cpi
    // + [registry, billing_config, per_chain_per_token_config, lookup_table]
    let mut ask_again_with = vec![];
    ask_again_with.extend(
        ctx.remaining_accounts[0..fee_quoter_fixed_accounts_len]
            .iter()
            .map(|a| a.key.readonly()),
    );
    ask_again_with.push(find(
        &[
            seed::NONCE,
            params.dest_chain_selector.to_le_bytes().as_ref(),
            params.ccip_send_caller.key().as_ref(),
        ],
        crate::ID,
    ));

    ask_again_with.extend(
        ctx.remaining_accounts[fee_quoter_fixed_accounts_len..]
            .chunks(accounts_per_token_len)
            .flat_map(|accs| {
                let registry = &accs[0];
                let token_admin_registry_account =
                    load_token_admin_registry_checked(registry).unwrap();
                let lut = token_admin_registry_account.lookup_table;
                accs.iter()
                    .map(|a| a.key.readonly())
                    .chain(Some(lut.readonly()))
            }),
    );

    Ok(DeriveAccountsResponse {
        accounts_to_save: vec![],
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsCcipSendStage::RetrieveTokenLUTs.to_string(),
        next_stage: DeriveAccountsCcipSendStage::RetrievePoolPrograms.to_string(),
    })
}

pub fn derive_ccip_send_accounts_retrieve_pool_programs<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
) -> Result<DeriveAccountsResponse> {
    let token_derivation_fixed_accounts_len = 6usize;
    let accounts_per_token_len = 4usize;

    // Accounts needed for get_fee cpi
    // + [registry, billing_config, per_chain_per_token_config, lookup_table, pool_program]
    let mut ask_again_with = vec![];
    ask_again_with.extend(
        ctx.remaining_accounts[0..token_derivation_fixed_accounts_len]
            .iter()
            .map(|a| a.key.readonly()),
    );

    ask_again_with.extend(
        ctx.remaining_accounts[token_derivation_fixed_accounts_len..]
            .chunks(accounts_per_token_len)
            .flat_map(|accs| {
                let lut = &accs[3];
                let lut_data = &mut &lut.data.borrow()[..];
                let lut_account: AddressLookupTable =
                    AddressLookupTable::deserialize(lut_data).expect("Deserialize LUT");
                let pool_program = lut_account.addresses[2];
                accs.iter()
                    .map(|a| a.key.readonly())
                    .chain(Some(pool_program.readonly()))
            }),
    );

    Ok(DeriveAccountsResponse {
        accounts_to_save: vec![],
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsCcipSendStage::RetrievePoolPrograms.to_string(),
        next_stage: DeriveAccountsCcipSendStage::TokenTransferStaticAccounts { page: 0 }
            .to_string(),
    })
}

pub fn derive_ccip_send_accounts_additional_tokens_static<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: DeriveAccountsCcipSendParams,
    page: u32,
) -> Result<DeriveAccountsResponse> {
    let mut response = DeriveAccountsResponse {
        current_stage: DeriveAccountsCcipSendStage::TokenTransferStaticAccounts { page }
            .to_string(),
        ..Default::default()
    };
    let token_derivation_fixed_accounts_len = 6usize;
    let accounts_per_token_len = 5usize;

    let token_derivation_fixed_accounts =
        &ctx.remaining_accounts[..token_derivation_fixed_accounts_len];
    // We extract the accounts for the first token
    let [token_registry, _billing_config, _per_chain_per_token_config, token_lut, _pool_program] =
        &ctx.remaining_accounts[token_derivation_fixed_accounts_len
            ..token_derivation_fixed_accounts_len + accounts_per_token_len]
    else {
        return Err(CcipRouterError::InvalidAccountListForPdaDerivation.into());
    };

    let lut_data = &mut &token_lut.data.borrow()[..];
    let lut_account: AddressLookupTable = AddressLookupTable::deserialize(lut_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let token_mint = lut_account.addresses[7];
    let pool_program = lut_account.addresses[2];
    let token_program = lut_account.addresses[6];
    let selector = params.dest_chain_selector.to_le_bytes();
    let sender_token_account = get_associated_token_address_with_program_id(
        &params.ccip_send_caller,
        &token_mint,
        &token_program.key(),
    )
    .writable();

    let token_billing_config = find(
        &[
            seed::PER_CHAIN_PER_TOKEN_CONFIG,
            &selector,
            token_mint.as_ref(),
        ],
        ctx.accounts.config.fee_quoter,
    )
    .readonly();

    let pool_chain_config = find(
        &[
            seed::TOKEN_POOL_CHAIN_CONFIG,
            &selector,
            token_mint.as_ref(),
        ],
        pool_program,
    )
    .writable();

    if page == 0 {
        response.look_up_tables_to_save = vec![*token_lut.key];
    }

    let mut all_accounts_to_save = vec![
        sender_token_account,
        token_billing_config,
        pool_chain_config,
    ];
    all_accounts_to_save.extend(
        lut_account
            .addresses
            .iter()
            .enumerate()
            .map(|(i, a)| match i {
                // PoolConfig, PoolTokenAccount and Mint from the LUT are writable.
                3 | 4 | 7 => a.writable(),
                _ => a.readonly(),
            }),
    );
    let start_of_page = TOKEN_ACCOUNTS_STATIC_PAGE_SIZE * page as usize;
    let num_accounts_in_page =
        TOKEN_ACCOUNTS_STATIC_PAGE_SIZE.min(all_accounts_to_save.len() - start_of_page);
    let end_of_page = start_of_page + num_accounts_in_page;
    all_accounts_to_save[start_of_page..end_of_page].clone_into(&mut response.accounts_to_save);

    if end_of_page != all_accounts_to_save.len() {
        response.next_stage =
            DeriveAccountsCcipSendStage::TokenTransferStaticAccounts { page: page + 1 }.to_string();
        // Different pages take the same inputs to derive, so we just reiterate the request:
        response.ask_again_with = ctx
            .remaining_accounts
            .iter()
            .map(|a| a.key().readonly())
            .collect();
        return Ok(response);
    }

    let (this_token_index, _) = params
        .message
        .token_amounts
        .iter()
        .enumerate()
        .find(|(_, ta)| ta.token == token_mint)
        .ok_or(CcipRouterError::InvalidAccountListForPdaDerivation)?;

    let number_of_tokens_left = params.message.token_amounts.len() - this_token_index;
    let token_admin_registry = load_token_admin_registry_checked(token_registry)?;

    if token_admin_registry.supports_auto_derivation {
        // Nested derivation is supported, so we go one level deeper
        response.next_stage = DeriveAccountsCcipSendStage::NestedTokenDerive {
            token_substage: "Start".to_string(),
        }
        .to_string();
        response.ask_again_with = ctx
            .remaining_accounts
            .iter()
            .map(|a| a.key.readonly())
            .collect();
    } else if number_of_tokens_left > 1 {
        // We aren't done yet with all tokens; as need to derive more tokens we tell the user
        // to ask again with one fewer.
        response.ask_again_with.clear();
        response.ask_again_with.extend(
            token_derivation_fixed_accounts
                .iter()
                .chain(
                    ctx.remaining_accounts
                        .iter()
                        .skip(token_derivation_fixed_accounts_len + accounts_per_token_len),
                )
                .map(|a| a.key.readonly()),
        );

        response.next_stage =
            DeriveAccountsCcipSendStage::TokenTransferStaticAccounts { page: 0 }.to_string();
    }

    Ok(response)
}

#[allow(clippy::too_many_arguments)]
pub fn derive_ccip_send_accounts_additional_token_nested<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsCcipSendParams,
    substage: &str,
) -> Result<DeriveAccountsResponse> {
    let token_derivation_fixed_accounts_len = 6usize;
    let accounts_per_token_len = 5usize;

    let token_derivation_fixed_accounts =
        &ctx.remaining_accounts[..token_derivation_fixed_accounts_len];
    // We extract the accounts for the first token
    let [_token_registry, billing_config, per_chain_per_token_config, token_lut, pool_program] =
        &ctx.remaining_accounts[token_derivation_fixed_accounts_len
            ..token_derivation_fixed_accounts_len + accounts_per_token_len]
    else {
        return Err(CcipRouterError::InvalidAccountListForPdaDerivation.into());
    };

    let get_fee_result = get_fee_cpi(
        token_derivation_fixed_accounts[0].clone(),
        token_derivation_fixed_accounts[1].clone(),
        token_derivation_fixed_accounts[2].clone(),
        token_derivation_fixed_accounts[3].clone(),
        token_derivation_fixed_accounts[4].clone(),
        params.dest_chain_selector,
        &params.message,
        vec![billing_config.clone(), per_chain_per_token_config.clone()],
    )?;

    let lut_data = &mut &token_lut.data.borrow()[..];
    let lut_account: AddressLookupTable = AddressLookupTable::deserialize(lut_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let token_mint = lut_account.addresses[7];

    let (this_token_index, transfer) = params
        .message
        .token_amounts
        .iter()
        .enumerate()
        .find(|(_, ta)| ta.token == token_mint)
        .ok_or(CcipRouterError::InvalidAccountListForPdaDerivation)?;
    let number_of_tokens_left = params.message.token_amounts.len() - this_token_index;
    let start_of_nested_accounts =
        token_derivation_fixed_accounts_len + accounts_per_token_len * number_of_tokens_left;

    let nested_derivation_accounts = &ctx.remaining_accounts[start_of_nested_accounts..];

    // +1 because in `ccip_send` it will be bumped before `LockOrBurnInV1` reaches
    // the pool.
    let total_nonce = load_nonce(&token_derivation_fixed_accounts[5])
        .map(|n| n.total_nonce)
        .unwrap_or_default()
        + 1;

    let lock_or_burn = LockOrBurnInV1 {
        receiver: get_fee_result
            .processed_extra_args
            .token_receiver
            .as_ref()
            .unwrap_or(&params.message.receiver)
            .clone(),
        remote_chain_selector: params.dest_chain_selector,
        original_sender: params.ccip_send_caller,
        amount: transfer.amount,
        local_token: transfer.token,
        msg_total_nonce: total_nonce,
    };

    let acc_metas: Vec<AccountMeta> = nested_derivation_accounts
        .iter()
        .flat_map(|acc_info| acc_info.to_account_metas(None))
        .collect();

    let mut data = Vec::new();
    data.extend_from_slice(&TOKENPOOL_DERIVE_LOCK_OR_BURN_DISCRIMINATOR);
    data.extend_from_slice(&substage.try_to_vec().unwrap());
    data.extend_from_slice(&lock_or_burn.try_to_vec().unwrap());

    let ix = Instruction {
        program_id: pool_program.key(),
        accounts: acc_metas,
        data,
    };

    invoke(&ix, nested_derivation_accounts)?;

    let (_, data) = get_return_data().unwrap();
    let mut response = DeriveAccountsResponse {
        ask_again_with: ctx.remaining_accounts[..start_of_nested_accounts]
            .iter()
            .map(|a| a.key().readonly())
            .collect(),
        ..Default::default()
    };
    response = response.and(
        DeriveAccountsResponse::try_from_slice(&data)
            .map_err(|_| CcipRouterError::InvalidTokenPoolAccountDerivationResponse)?,
    );

    // We're coming back from a nested derivation call, so we turn the stage reported
    // by it into our substage.
    if !response.next_stage.is_empty() {
        response.next_stage = DeriveAccountsCcipSendStage::NestedTokenDerive {
            token_substage: response.next_stage,
        }
        .to_string();
    } else if number_of_tokens_left > 1 {
        // We're done with the nested calls (next stage is empty) and there are more
        // tokens to do, so we go back to the static stage with one fewer token.
        response.ask_again_with.clear();
        response.ask_again_with.extend(
            token_derivation_fixed_accounts
                .iter()
                .chain(
                    ctx.remaining_accounts
                        .iter()
                        .skip(token_derivation_fixed_accounts_len + accounts_per_token_len),
                )
                .map(|a| a.key.readonly()),
        );

        response.next_stage =
            DeriveAccountsCcipSendStage::TokenTransferStaticAccounts { page: 0 }.to_string();
    }

    Ok(response)
}
