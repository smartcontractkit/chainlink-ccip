use crate::{
    context::ViewConfigOnly,
    instructions::v1::messages::ReleaseOrMintInV1,
    messages::TOKENPOOL_DERIVE_RELEASE_OR_MINT_DISCRIMINATOR,
    state::{
        CcipAccountMeta, DeriveAccountsExecuteParams, DeriveAccountsResponse, ReferenceAddresses,
        ToMeta,
    },
    CcipOfframpError,
};
use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use ccip_common::v1::load_token_admin_registry_checked;
use ccip_common::CommonCcipError;
use ccip_common::{
    seed::{self, EXECUTION_REPORT_BUFFER},
    v1::token_admin_registry_writable,
};
use solana_program::{
    address_lookup_table::state::AddressLookupTable,
    instruction::Instruction,
    program::{get_return_data, invoke},
    sysvar::instructions,
};
use std::{
    fmt::{self, Display, Formatter},
    str::FromStr,
};

const MIN_PAGE_SIZE: usize = 4;

// Local helper to find a readonly CCIP meta for a given seed + program_id combo.
// Short name for compactness.
fn find(seeds: &[&[u8]], program_id: Pubkey) -> CcipAccountMeta {
    Pubkey::find_program_address(seeds, &program_id)
        .0
        .readonly()
}

#[derive(Clone, Debug)]
pub enum DeriveAccountsExecuteStage {
    Start,
    FinishMainAccountList,
    RetrieveTokenLUTs,
    RetrievePoolPrograms,
    // N stages, one per token for the ones below.
    TokenTransferStaticAccounts {
        // Index of the current token being derived.
        token: u32,
        // Might be too many to fit in one response, so the user
        // may be required to request multiple pages.
        page: u32,
    },
    NestedTokenDerive {
        token: u32,
        token_substage: String,
    },
}

impl Display for DeriveAccountsExecuteStage {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        match self {
            DeriveAccountsExecuteStage::Start => f.write_str("Start"),
            DeriveAccountsExecuteStage::FinishMainAccountList => {
                f.write_str("FinishMainAccountList")
            }
            DeriveAccountsExecuteStage::RetrieveTokenLUTs => {
                f.write_str("RetrieveTokenLookupTables")
            }
            DeriveAccountsExecuteStage::RetrievePoolPrograms => f.write_str("RetrievePoolPrograms"),
            DeriveAccountsExecuteStage::TokenTransferStaticAccounts { token, page } => {
                f.write_fmt(format_args!("TokenTransferStaticAccounts/{token}/{page}"))
            }
            DeriveAccountsExecuteStage::NestedTokenDerive {
                token_substage,
                token,
            } => f.write_fmt(format_args!("NestedTokenDerive/{token}/{token_substage}")),
        }
    }
}

impl FromStr for DeriveAccountsExecuteStage {
    type Err = CcipOfframpError;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        let mut s = s.split('/');
        let (a, b, c) = (s.next(), s.next(), s.next());

        match (a, b, c) {
            (Some("Start"), None, None) => Ok(Self::Start),
            (Some("FinishMainAccountList"), None, None) => Ok(Self::FinishMainAccountList),
            (Some("RetrieveTokenLookupTables"), None, None) => Ok(Self::RetrieveTokenLUTs),
            (Some("RetrievePoolPrograms"), None, None) => Ok(Self::RetrievePoolPrograms),
            (Some("TokenTransferStaticAccounts"), Some(token), Some(page)) => {
                Ok(Self::TokenTransferStaticAccounts {
                    page: str::parse::<u32>(page)
                        .or(Err(CcipOfframpError::InvalidDerivationStage))?,
                    token: str::parse::<u32>(token)
                        .or(Err(CcipOfframpError::InvalidDerivationStage))?,
                })
            }
            (Some("NestedTokenDerive"), Some(token), token_substage) => {
                Ok(Self::NestedTokenDerive {
                    token_substage: token_substage.unwrap_or("Start").to_string(),
                    token: str::parse::<u32>(token)
                        .or(Err(CcipOfframpError::InvalidDerivationStage))?,
                })
            }
            _ => Err(CcipOfframpError::InvalidDerivationStage),
        }
    }
}

pub fn derive_execute_accounts_start(source_chain_selector: u64) -> Result<DeriveAccountsResponse> {
    let accounts_to_save = vec![
        find(&[seed::CONFIG], crate::ID),
        find(&[seed::REFERENCE_ADDRESSES], crate::ID),
        find(
            &[seed::SOURCE_CHAIN, &source_chain_selector.to_le_bytes()],
            crate::ID,
        ),
    ];

    let ask_again_with = vec![
        find(
            &[seed::SOURCE_CHAIN, &source_chain_selector.to_le_bytes()],
            crate::ID,
        ),
        find(&[seed::REFERENCE_ADDRESSES], crate::ID),
    ];

    Ok(DeriveAccountsResponse {
        accounts_to_save,
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsExecuteStage::Start.to_string(),
        next_stage: DeriveAccountsExecuteStage::FinishMainAccountList.to_string(),
    })
}

pub fn derive_execute_accounts_build_main_account_list<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsExecuteParams,
) -> Result<DeriveAccountsResponse> {
    let ReferenceAddresses {
        router,
        rmn_remote,
        offramp_lookup_table,
        ..
    } = *AccountLoader::<'info, ReferenceAddresses>::try_from(&ctx.remaining_accounts[1])?
        .load()?;

    let selector = params.source_chain_selector.to_le_bytes();

    let mut accounts_to_save = vec![
        find(
            &[seed::COMMIT_REPORT, &selector, &params.merkle_root],
            crate::ID,
        )
        .writable(),
        crate::ID.readonly(),
        find(
            &[seed::ALLOWED_OFFRAMP, &selector, crate::ID.as_ref()],
            router,
        ),
        params.execute_caller.writable().signer(),
        solana_program::system_program::ID.readonly(),
        instructions::ID.readonly(),
        rmn_remote.readonly(),
        find(&[seed::CURSES], rmn_remote),
        find(&[seed::CONFIG], rmn_remote),
    ];

    if !params.message_accounts.is_empty() {
        let find = |seeds, id| Pubkey::find_program_address(seeds, id).0.readonly();
        accounts_to_save.push(params.message_accounts[0].clone());
        accounts_to_save.push(find(
            &[
                seed::EXTERNAL_EXECUTION_CONFIG,
                params.message_accounts[0].pubkey.as_ref(),
            ],
            &crate::ID,
        ));
        accounts_to_save.extend_from_slice(&params.message_accounts[1..]);
    }

    let (next_stage, ask_again_with) = if params.token_transfers.is_empty() {
        // We're done (no further stages) so we need to consider whether it's the
        // buffered case to append the buffer PDA.
        if !params.buffer_id.is_empty() {
            let buffer_pda = find(
                &[
                    EXECUTION_REPORT_BUFFER,
                    &params.buffer_id,
                    params.execute_caller.as_ref(),
                ],
                crate::ID,
            );
            accounts_to_save.push(buffer_pda.writable());
        }
        ("".to_string(), vec![])
    } else {
        // There are tokens, so we continue by retrieving their lookup tables on next stage.
        (
            DeriveAccountsExecuteStage::RetrieveTokenLUTs.to_string(),
            params
                .token_transfers
                .iter()
                .map(|tt| {
                    find(
                        &[
                            seed::TOKEN_ADMIN_REGISTRY,
                            tt.transfer.dest_token_address.as_ref(),
                        ],
                        router,
                    )
                })
                .collect(),
        )
    };
    Ok(DeriveAccountsResponse {
        accounts_to_save,
        ask_again_with,
        look_up_tables_to_save: vec![offramp_lookup_table],
        current_stage: DeriveAccountsExecuteStage::FinishMainAccountList.to_string(),
        next_stage,
    })
}

pub fn derive_execute_accounts_retrieve_luts<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
) -> Result<DeriveAccountsResponse> {
    // reference_addresses, followed registries and LUTs for each token
    let mut ask_again_with = vec![find(&[seed::REFERENCE_ADDRESSES], crate::ID)];
    ask_again_with.extend(ctx.remaining_accounts.iter().flat_map(|account| {
        let registry = load_token_admin_registry_checked(account)
            .expect("parsing token admin registry account");
        [account.key().readonly(), registry.lookup_table.readonly()].into_iter()
    }));

    Ok(DeriveAccountsResponse {
        accounts_to_save: vec![],
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsExecuteStage::RetrieveTokenLUTs.to_string(),
        next_stage: DeriveAccountsExecuteStage::RetrievePoolPrograms.to_string(),
    })
}

pub fn derive_execute_accounts_retrieve_pool_programs<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
) -> Result<DeriveAccountsResponse> {
    let token_derivation_fixed_accounts_len = 1usize;
    let accounts_per_token_len = 2usize;

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
                let lut = &accs[1];
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
        current_stage: DeriveAccountsExecuteStage::RetrievePoolPrograms.to_string(),
        next_stage: DeriveAccountsExecuteStage::TokenTransferStaticAccounts { page: 0, token: 0 }
            .to_string(),
    })
}

pub fn derive_execute_accounts_additional_tokens_static<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsExecuteParams,
    page: u32,
    token: u32,
) -> Result<DeriveAccountsResponse> {
    let mut response = DeriveAccountsResponse {
        current_stage: DeriveAccountsExecuteStage::TokenTransferStaticAccounts { token, page }
            .to_string(),
        // It's possible we'll need to return to this function, so we
        // initialize with the same account list
        ask_again_with: ctx
            .remaining_accounts
            .iter()
            .map(|a| a.key().readonly())
            .collect(),
        ..Default::default()
    };

    let token_derivation_fixed_accounts_len = 1usize;
    let accounts_per_token_len = 3usize;
    let [reference_addresses, ..] = ctx.remaining_accounts else {
        return Err(CcipOfframpError::InvalidAccountListForPdaDerivation.into());
    };
    let token_start_index =
        token_derivation_fixed_accounts_len + accounts_per_token_len * token as usize;
    let token_end_index = token_start_index + accounts_per_token_len;

    let [token_registry, token_lut, _] =
        &ctx.remaining_accounts[token_start_index..token_end_index]
    else {
        return Err(CcipOfframpError::InvalidAccountListForPdaDerivation.into());
    };

    if page == 0 {
        response.look_up_tables_to_save = vec![*token_lut.key];
    }

    let lut_data = &mut &token_lut.data.borrow()[..];
    let lut_account: AddressLookupTable = AddressLookupTable::deserialize(lut_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let token_mint = lut_account.addresses[7];
    let pool_program = lut_account.addresses[2];
    let token_program = lut_account.addresses[6];

    let ReferenceAddresses { fee_quoter, .. } =
        *AccountLoader::<'info, ReferenceAddresses>::try_from(reference_addresses)?.load()?;

    let lookup_table_data = &mut &token_lut.data.borrow()[..];
    let lut_account: AddressLookupTable = AddressLookupTable::deserialize(lookup_table_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;

    let ccip_offramp_pool_signer = find(
        &[seed::EXTERNAL_TOKEN_POOLS_SIGNER, pool_program.as_ref()],
        crate::ID,
    )
    .readonly();

    let user_token_account = get_associated_token_address_with_program_id(
        &params.token_receiver,
        &token_mint.key(),
        &token_program.key(),
    )
    .writable();

    let selector = params.source_chain_selector.to_le_bytes();
    let token_billing_config = find(
        &[
            seed::PER_CHAIN_PER_TOKEN_CONFIG,
            &selector,
            token_mint.as_ref(),
        ],
        fee_quoter,
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

    let token_admin_registry = load_token_admin_registry_checked(token_registry)?;

    response.accounts_to_save.extend_from_slice(&[
        ccip_offramp_pool_signer,
        user_token_account,
        token_billing_config,
        pool_chain_config,
    ]);
    response
        .accounts_to_save
        .extend(lut_account.addresses.iter().enumerate().map(|(i, a)| {
            match token_admin_registry_writable::is(&token_admin_registry, i.try_into().unwrap()) {
                true => a.writable(),
                false => a.readonly(),
            }
        }));

    let max_response_accounts = 26;
    if response.ask_again_with.len() + response.accounts_to_save.len() > max_response_accounts {
        let total_accounts_to_save = response.accounts_to_save.len();
        // paging is necessary, because we can't fit everything in one response.
        let max_accounts_per_page = max_response_accounts - response.ask_again_with.len();
        require_gte!(
            max_accounts_per_page,
            MIN_PAGE_SIZE,
            CcipOfframpError::AccountDerivationResponseTooLarge
        );
        let start_of_page = max_accounts_per_page * page as usize;
        let end_of_page = total_accounts_to_save.min(start_of_page + max_accounts_per_page);
        response.accounts_to_save = response.accounts_to_save[start_of_page..end_of_page].to_vec();
        if end_of_page < total_accounts_to_save {
            response.next_stage = DeriveAccountsExecuteStage::TokenTransferStaticAccounts {
                page: page + 1,
                token,
            }
            .to_string();
            // Different pages take the same inputs to derive, so we just reiterate the request:
            return Ok(response);
        }
    }

    let registry = load_token_admin_registry_checked(token_registry)?;
    if registry.supports_auto_derivation {
        // Nested derivation is supported, so we go one level deeper
        response.next_stage = DeriveAccountsExecuteStage::NestedTokenDerive {
            token_substage: "Start".to_string(),
            token,
        }
        .to_string();
    } else if token + 1 < params.token_transfers.len() as u32 {
        // We aren't done yet with all tokens
        response.next_stage = DeriveAccountsExecuteStage::TokenTransferStaticAccounts {
            page: 0,
            token: token + 1,
        }
        .to_string();
    } else {
        // We're done with all tokens
        if !params.buffer_id.is_empty() {
            let buffer_pda = find(
                &[
                    EXECUTION_REPORT_BUFFER,
                    &params.buffer_id,
                    params.execute_caller.as_ref(),
                ],
                crate::ID,
            );
            response.accounts_to_save.push(buffer_pda.writable());
        }
        response.ask_again_with.clear();
        response.next_stage = "".to_string(); // We're done, no more stages.
    }

    Ok(response)
}

pub fn derive_execute_accounts_additional_token_nested<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsExecuteParams,
    substage: &str,
    token: u32,
) -> Result<DeriveAccountsResponse> {
    let token_derivation_fixed_accounts_len = 1usize;
    let accounts_per_token_len = 3usize;
    let token_start_index =
        token_derivation_fixed_accounts_len + accounts_per_token_len * token as usize;
    let token_end_index = token_start_index + accounts_per_token_len;

    let [_, token_lut, _] = &ctx.remaining_accounts[token_start_index..token_end_index] else {
        return Err(CcipOfframpError::InvalidAccountListForPdaDerivation.into());
    };

    let lut_data = &mut &token_lut.data.borrow()[..];
    let lut_account: AddressLookupTable = AddressLookupTable::deserialize(lut_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let pool_program = lut_account.addresses[2];

    let start_of_nested_accounts =
        token_derivation_fixed_accounts_len + params.token_transfers.len() * accounts_per_token_len;
    let mut response = DeriveAccountsResponse {
        ask_again_with: ctx.remaining_accounts[..start_of_nested_accounts]
            .iter()
            .map(|a| a.key().readonly())
            .collect(),
        current_stage: DeriveAccountsExecuteStage::NestedTokenDerive {
            token_substage: substage.to_string(),
            token,
        }
        .to_string(),
        ..Default::default()
    };

    let this_transfer = &params.token_transfers[token as usize];

    let release_or_mint = ReleaseOrMintInV1 {
        original_sender: params.original_sender.clone(),
        receiver: params.token_receiver,
        amount: this_transfer.transfer.amount,
        local_token: this_transfer.transfer.dest_token_address,
        remote_chain_selector: params.source_chain_selector,
        source_pool_address: this_transfer.transfer.source_pool_address.clone(),
        source_pool_data: this_transfer.transfer.extra_data.clone(),
        offchain_token_data: this_transfer.data.clone(),
    };
    let nested_derivation_accounts = &ctx.remaining_accounts[start_of_nested_accounts..];
    let acc_metas: Vec<AccountMeta> = nested_derivation_accounts
        .iter()
        .flat_map(|acc_info| acc_info.to_account_metas(None))
        .collect();

    let mut data = Vec::new();
    data.extend_from_slice(&TOKENPOOL_DERIVE_RELEASE_OR_MINT_DISCRIMINATOR);
    data.extend_from_slice(&substage.try_to_vec().unwrap());
    data.extend_from_slice(&release_or_mint.try_to_vec().unwrap());

    let ix = Instruction {
        program_id: pool_program,
        accounts: acc_metas,
        data,
    };

    invoke(&ix, nested_derivation_accounts)?;
    let (_, data) = get_return_data().unwrap();
    let nested_response = DeriveAccountsResponse::try_from_slice(&data)
        .map_err(|_| CcipOfframpError::InvalidTokenPoolAccountDerivationResponse)?;
    response.accounts_to_save = nested_response.accounts_to_save;
    response.look_up_tables_to_save = nested_response.look_up_tables_to_save;

    // We're coming back from a nested derivation call, so we turn the stage reported
    // by it into our substage.
    if !nested_response.next_stage.is_empty() {
        response
            .ask_again_with
            .extend_from_slice(&nested_response.ask_again_with);
        response.next_stage = DeriveAccountsExecuteStage::NestedTokenDerive {
            token_substage: nested_response.next_stage,
            token,
        }
        .to_string();
    } else if token + 1 < params.token_transfers.len() as u32 {
        response.next_stage = DeriveAccountsExecuteStage::TokenTransferStaticAccounts {
            page: 0,
            token: token + 1,
        }
        .to_string();
    } else {
        // We're done with all tokens
        if !params.buffer_id.is_empty() {
            let buffer_pda = find(
                &[
                    EXECUTION_REPORT_BUFFER,
                    &params.buffer_id,
                    params.execute_caller.as_ref(),
                ],
                crate::ID,
            );
            response.accounts_to_save.push(buffer_pda.writable());
        }
        response.ask_again_with.clear();
        response.next_stage = "".to_string(); // We're done, no more stages.
    }

    Ok(response)
}
