use std::{
    fmt::{self, Display, Formatter},
    str::FromStr,
};

use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id, token::spl_token::native_mint,
};
use ccip_common::{router_accounts::TokenAdminRegistry, seed, CommonCcipError};
use solana_program::{
    address_lookup_table::state::AddressLookupTable,
    instruction::Instruction,
    program::{get_return_data, invoke},
};

use crate::{
    context::ViewConfigOnly,
    instructions::v1::{fees::get_fee_cpi, messages::pools::LockOrBurnInV1},
    messages::{SVMTokenAmount, TOKENPOOL_DERIVE_LOCK_OR_BURN_DISCRIMINATOR},
    state::{CcipAccountMeta, DeriveAccountsCcipSendParams, DeriveAccountsResponse, ToMeta},
    CcipRouterError,
};

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
    // N stages, one per token
    TokenTransferAccounts { token_substage: String },
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
            DeriveAccountsCcipSendStage::TokenTransferAccounts { token_substage } => {
                f.write_fmt(format_args!("TokenTransferAccounts/{token_substage}"))
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
            (Some("TokenTransferAccounts"), token_substage) => Ok(Self::TokenTransferAccounts {
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

    let ask_again_with = vec![
        params.message.fee_token.readonly(),
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
    let fee_token_program = fee_token_mint_info.owner;

    let is_fee_in_sol = *fee_token_program == Pubkey::default();

    let fee_token_user_ata = if is_fee_in_sol {
        Pubkey::default().readonly()
    } else {
        get_associated_token_address_with_program_id(
            &params.ccip_send_caller.key(),
            &params.message.fee_token,
            &fee_token_program.key(),
        )
        .writable()
    };

    let fee_token_receiver = get_associated_token_address_with_program_id(
        fee_billing_signer.key,
        &params.message.fee_token,
        &fee_token_program.key(),
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
    ask_again_with.extend(
        ctx.remaining_accounts[fee_quoter_fixed_accounts_len..]
            .chunks(accounts_per_token_len)
            .flat_map(|accs| {
                let registry = &accs[0];
                let token_admin_registry_account: Account<TokenAdminRegistry> =
                    Account::try_from(registry).expect("parsing token admin registry account");
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
        next_stage: DeriveAccountsCcipSendStage::TokenTransferAccounts {
            token_substage: "Start".to_string(),
        }
        .to_string(),
    })
}

pub fn derive_execute_accounts_additional_tokens<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: DeriveAccountsCcipSendParams,
    substage: &str,
) -> Result<DeriveAccountsResponse> {
    let fee_quoter_fixed_accounts_len = 5usize;
    let accounts_per_token_len = 4usize;

    let fee_quoter_fixed_accounts = &ctx.remaining_accounts[..fee_quoter_fixed_accounts_len];
    // We extract the accounts for the first token
    let [token_registry, billing_config, per_chain_per_token_config, token_lut] = &ctx
        .remaining_accounts
        [fee_quoter_fixed_accounts_len..fee_quoter_fixed_accounts_len + accounts_per_token_len]
    else {
        return Err(CcipRouterError::InvalidAccountListForPdaDerivation.into());
    };

    let lut_data = &mut &token_lut.data.borrow()[..];
    let lut_account: AddressLookupTable = AddressLookupTable::deserialize(lut_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let token_mint = lut_account.addresses[7];
    let pool_program = lut_account.addresses[2];
    let token_program = lut_account.addresses[6];
    let mut response = DeriveAccountsResponse::default();
    if substage == "Start" {
        // All tokens derive a static list of accounts. If we're in the first substage
        // of derivation ("Start"), which is common to dynamic and static pools, we derive
        // the list of static accounts now.
        response = response.and(derive_ccip_send_additional_token_static(
            &ctx,
            &params,
            token_lut,
            lut_account,
            token_mint,
            pool_program,
            token_program,
        )?);
    }

    let token_registry_account: Account<TokenAdminRegistry> =
        Account::try_from(token_registry).expect("parsing token admin registry account");
    if token_registry_account.supports_auto_derivation {
        let (this_token_index, this_transfer) = params
            .message
            .token_amounts
            .iter()
            .enumerate()
            .find(|(_, ta)| ta.token == token_mint)
            .ok_or(CcipRouterError::InvalidAccountListForPdaDerivation)?;

        let number_of_tokens_left = params.message.token_amounts.len() - this_token_index;
        let nested_derivation_accounts = &ctx.remaining_accounts
            [fee_quoter_fixed_accounts_len + accounts_per_token_len * number_of_tokens_left..];

        // Nested derivation is supported, so we go one level deeper.
        response = response.and(derive_ccip_send_accounts_additional_token_nested(
            &params,
            substage,
            fee_quoter_fixed_accounts,
            nested_derivation_accounts,
            billing_config,
            per_chain_per_token_config,
            pool_program.key(),
            this_transfer,
        )?);
    }

    Ok(response)
}

fn derive_ccip_send_additional_token_static<'info>(
    ctx: &Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsCcipSendParams,
    token_lut: &AccountInfo<'info>,
    lut_account: AddressLookupTable<'_>,
    token_mint: Pubkey,
    pool_program: Pubkey,
    token_program: Pubkey,
) -> std::result::Result<DeriveAccountsResponse, Error> {
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

    let mut accounts_to_save = vec![
        sender_token_account,
        token_billing_config,
        pool_chain_config,
    ];
    accounts_to_save.extend(
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

    let mut ask_again_with = vec![];
    let next_stage = if ctx.remaining_accounts.len() > 1 {
        // We aren't done yet, we need to derive more tokens, so we tell the user
        // to ask again with one fewer token.
        ask_again_with.extend(ctx.remaining_accounts[1..].iter().map(|a| a.key.readonly()));
        DeriveAccountsCcipSendStage::TokenTransferAccounts {
            token_substage: "Start".to_string(),
        }
        .to_string()
    } else {
        "".to_string()
    };

    Ok(DeriveAccountsResponse {
        ask_again_with,
        accounts_to_save,
        look_up_tables_to_save: vec![*token_lut.key],
        current_stage: DeriveAccountsCcipSendStage::TokenTransferAccounts {
            token_substage: "Start".to_string(),
        }
        .to_string(),
        next_stage,
    })
}

#[allow(clippy::too_many_arguments)]
fn derive_ccip_send_accounts_additional_token_nested<'info>(
    params: &DeriveAccountsCcipSendParams,
    substage: &str,
    fee_quoter_fixed_accounts: &[AccountInfo<'info>],
    nested_derivation_accounts: &[AccountInfo<'info>],
    billing_config: &AccountInfo<'info>,
    per_chain_per_token_config: &AccountInfo<'info>,
    pool_program: Pubkey,
    transfer: &SVMTokenAmount,
) -> Result<DeriveAccountsResponse> {
    let get_fee_result = get_fee_cpi(
        fee_quoter_fixed_accounts[0].clone(),
        fee_quoter_fixed_accounts[1].clone(),
        fee_quoter_fixed_accounts[2].clone(),
        fee_quoter_fixed_accounts[3].clone(),
        fee_quoter_fixed_accounts[4].clone(),
        params.dest_chain_selector,
        &params.message,
        vec![billing_config.clone(), per_chain_per_token_config.clone()],
    )?;
    let nonce = 0u64; // TODO

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
        msg_total_nonce: nonce,
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
        program_id: pool_program,
        accounts: acc_metas,
        data,
    };

    invoke(&ix, nested_derivation_accounts)?;

    let (_, data) = get_return_data().unwrap();
    DeriveAccountsResponse::try_from_slice(&data)
        .map_err(|_| CcipRouterError::InvalidTokenPoolAccountDerivationResponse.into())
}
