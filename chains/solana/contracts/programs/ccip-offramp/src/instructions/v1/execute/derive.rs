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
use ccip_common::{
    router_accounts::TokenAdminRegistry,
    seed::{self, EXECUTION_REPORT_BUFFER},
    CommonCcipError,
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

// Local helper to find a readonly CCIP meta for a given seed + program_id combo.
// Short name for compactness.
fn find(seeds: &[&[u8]], program_id: Pubkey) -> CcipAccountMeta {
    Pubkey::find_program_address(seeds, &program_id)
        .0
        .readonly()
}

#[derive(Clone, Debug)]
pub enum DeriveExecuteAccountsStage {
    GatherBasicInfo,
    BuildMainAccountList,
    RetrieveTokenLUTs,
    // N stages, one per token
    TokenTransferAccounts { token_substage: String },
}

impl Display for DeriveExecuteAccountsStage {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        match self {
            DeriveExecuteAccountsStage::GatherBasicInfo => f.write_str("GatherBasicInfo"),
            DeriveExecuteAccountsStage::BuildMainAccountList => f.write_str("BuildMainAccountList"),
            DeriveExecuteAccountsStage::RetrieveTokenLUTs => {
                f.write_str("RetrieveTokenLookupTables")
            }
            DeriveExecuteAccountsStage::TokenTransferAccounts { token_substage } => {
                f.write_fmt(format_args!("TokenTransferAccounts/{token_substage}"))
            }
        }
    }
}

impl FromStr for DeriveExecuteAccountsStage {
    type Err = CcipOfframpError;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        let mut s = s.split('/');
        let (prefix, suffix) = (s.next(), s.next());

        match (prefix, suffix) {
            (Some("Start" | "GatherBasicInfo"), None) => Ok(Self::GatherBasicInfo),
            (Some("BuildMainAccountList"), None) => Ok(Self::BuildMainAccountList),
            (Some("RetrieveTokenLookupTables"), None) => Ok(Self::RetrieveTokenLUTs),
            (Some("TokenTransferAccounts"), token_substage) => Ok(Self::TokenTransferAccounts {
                token_substage: token_substage.unwrap_or("Start").to_string(),
            }),
            _ => Err(CcipOfframpError::InvalidDerivationStage),
        }
    }
}

pub fn derive_execute_accounts_gather_basic_info(
    source_chain_selector: u64,
) -> Result<DeriveAccountsResponse> {
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
        current_stage: DeriveExecuteAccountsStage::GatherBasicInfo.to_string(),
        next_stage: DeriveExecuteAccountsStage::BuildMainAccountList.to_string(),
    })
}

pub fn derive_execute_accounts_build_main_account_list<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsExecuteParams,
) -> Result<DeriveAccountsResponse> {
    let ReferenceAddresses {
        router, rmn_remote, ..
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

    // If there are no tokens, we're done. If there are tokens, we
    // start by reading the registries on next stage.
    let ask_again_with: Vec<_> = params
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
        .collect();

    let next_stage = if ask_again_with.is_empty() {
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
        "".to_string()
    } else {
        DeriveExecuteAccountsStage::RetrieveTokenLUTs.to_string()
    };

    Ok(DeriveAccountsResponse {
        accounts_to_save,
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveExecuteAccountsStage::BuildMainAccountList.to_string(),
        next_stage,
    })
}

pub fn derive_execute_accounts_retrieve_luts<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
) -> Result<DeriveAccountsResponse> {
    let registries: Vec<_> = ctx
        .remaining_accounts
        .iter()
        .map(|registry| {
            let token_admin_registry_account: Account<TokenAdminRegistry> =
                Account::try_from(registry).expect("parsing token admin registry account");
            token_admin_registry_account
        })
        .collect();

    let lookup_table_metas = registries.iter().map(|r| r.lookup_table.readonly());
    // reference_addresses, followed by all registries, followed by all LUTs
    let mut ask_again_with = vec![find(&[seed::REFERENCE_ADDRESSES], crate::ID)];
    ask_again_with.extend(ctx.remaining_accounts.iter().map(|a| a.key.readonly()));
    ask_again_with.extend(lookup_table_metas);

    Ok(DeriveAccountsResponse {
        accounts_to_save: vec![],
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveExecuteAccountsStage::RetrieveTokenLUTs.to_string(),
        next_stage: DeriveExecuteAccountsStage::TokenTransferAccounts {
            token_substage: "Start".to_string(),
        }
        .to_string(),
    })
}

pub fn derive_execute_accounts_additional_tokens<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    params: &DeriveAccountsExecuteParams,
    substage: &str,
) -> Result<DeriveAccountsResponse> {
    // The reference addresses account and at least one registry + LUT.
    require_gte!(
        ctx.remaining_accounts.len(),
        3,
        CcipOfframpError::InvalidAccountListForPdaDerivation
    );

    let max_tokens_left = params
        .token_transfers
        .len()
        .min((ctx.remaining_accounts.len() - 1) / 2);
    let reference_addresses = &ctx.remaining_accounts[0];

    // Token admin registries, one per token transferred.
    let registries = &ctx.remaining_accounts[1..max_tokens_left + 1];
    // Look up tables, one per token transferred.
    let luts = &ctx.remaining_accounts[max_tokens_left + 1..2 * max_tokens_left + 1];
    // Registry of the token we're *currently* deriving
    let first_token_registry = &registries[0];
    // LUT of the token we're *currently* deriving
    let first_token_lut = &luts[0];
    let first_token_admin_registry_account: Account<TokenAdminRegistry> =
        Account::try_from(first_token_registry).expect("parsing token admin registry account");
    let first_lut_data = &mut &first_token_lut.data.borrow()[..];
    let first_lut_account: AddressLookupTable = AddressLookupTable::deserialize(first_lut_data)
        .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let first_token_mint = first_lut_account.addresses[7];
    let first_pool_program = first_lut_account.addresses[2];

    // We find the transfer for this specific token
    let (this_token_index, this_transfer) = params
        .token_transfers
        .iter()
        .enumerate()
        .find(|(_, tt)| tt.transfer.dest_token_address == first_token_mint)
        .ok_or(CcipOfframpError::InvalidAccountListForPdaDerivation)?;
    let tokens_left = max_tokens_left - this_token_index;

    // If we're doing nested derivation (i.e. we are deriving accounts for a
    // dynamic token pool with multiple derivation stages) these are the accounts
    // that the token pool required us to send for this stage.
    let nested_derivation_accounts = &ctx.remaining_accounts[2 * tokens_left + 1..];

    // Reconstruct the `ReleaseOrMint` the same way the offramp will
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

    let mut response = DeriveAccountsResponse::default();
    if substage == "Start" {
        // All tokens derive a static list of accounts. If we're in the first substage
        // of derivation ("Start"), which is common to dynamic and static pools, we derive
        // the list of static accounts now.
        response = response.and(derive_execute_accounts_additional_token_static(
            &release_or_mint,
            params.source_chain_selector,
            reference_addresses,
            first_token_lut,
        )?);
    }

    if first_token_admin_registry_account.supports_auto_derivation {
        // Nested derivation is supported, so we go one level deeper.
        response = response.and(derive_execute_accounts_additional_token_nested(
            &release_or_mint,
            substage,
            nested_derivation_accounts,
            first_pool_program.key(),
        )?);
    }

    if response.next_stage != String::default() {
        // We aren't done yet with this token, so we return as-is. The nested `derive`
        // call has already populated the `ask_again_with` list.
        return Ok(response);
    }

    if tokens_left > 1 {
        // We aren't done yet with all tokens; as need to derive more tokens we tell the user
        // to ask again with one fewer.
        response
            .ask_again_with
            .push(find(&[seed::REFERENCE_ADDRESSES], crate::ID));
        response
            .ask_again_with
            .extend(registries.iter().skip(1).map(|r| r.key.readonly()));
        response
            .ask_again_with
            .extend(luts.iter().skip(1).map(|l| l.key.readonly()));
        response.next_stage = DeriveExecuteAccountsStage::TokenTransferAccounts {
            token_substage: "Start".to_string(),
        }
        .to_string();
    } else if !params.buffer_id.is_empty() {
        // We're done and it's the buffered case, so we append the buffer PDA
        let buffer_pda = find(
            &[
                seed::EXECUTION_REPORT_BUFFER,
                &params.buffer_id,
                params.execute_caller.as_ref(),
            ],
            crate::ID,
        );
        response.accounts_to_save.push(buffer_pda.writable());
    };

    Ok(response)
}

fn derive_execute_accounts_additional_token_nested(
    release_or_mint_in: &ReleaseOrMintInV1,
    substage: &str,
    nested_derivation_accounts: &[AccountInfo<'_>],
    pool_program: Pubkey,
) -> Result<DeriveAccountsResponse> {
    let acc_metas: Vec<AccountMeta> = nested_derivation_accounts
        .iter()
        .flat_map(|acc_info| acc_info.to_account_metas(None))
        .collect();

    let mut data = Vec::new();
    data.extend_from_slice(&TOKENPOOL_DERIVE_RELEASE_OR_MINT_DISCRIMINATOR);
    data.extend_from_slice(&substage.try_to_vec().unwrap());
    data.extend_from_slice(&release_or_mint_in.try_to_vec().unwrap());

    let ix = Instruction {
        program_id: pool_program,
        accounts: acc_metas,
        data,
    };

    invoke(&ix, nested_derivation_accounts)?;

    let (_, data) = get_return_data().unwrap();
    DeriveAccountsResponse::try_from_slice(&data)
        .map_err(|_| CcipOfframpError::InvalidTokenPoolAccountDerivationResponse.into())
}

// Derives all static accounts for one token: those that don't change per invocation
// or can be derived from its fixed LUT. Tokens with dynamic account lists require
// further steps.
fn derive_execute_accounts_additional_token_static<'info>(
    release_or_mint_in: &ReleaseOrMintInV1,
    source_chain_selector: u64,
    reference_addresses: &'info AccountInfo<'info>,
    lut: &AccountInfo,
) -> Result<DeriveAccountsResponse> {
    let mut response = DeriveAccountsResponse::default();
    let selector = source_chain_selector.to_le_bytes();
    let ReferenceAddresses { fee_quoter, .. } =
        *AccountLoader::<'info, ReferenceAddresses>::try_from(reference_addresses)?.load()?;

    let lookup_table_data = &mut &lut.data.borrow()[..];
    let lookup_table_account: AddressLookupTable =
        AddressLookupTable::deserialize(lookup_table_data)
            .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;

    let pool_program = lookup_table_account.addresses[2];
    let token_program = lookup_table_account.addresses[6];
    let token_mint = lookup_table_account.addresses[7];
    let ccip_offramp_pool_signer = find(
        &[seed::EXTERNAL_TOKEN_POOLS_SIGNER, pool_program.as_ref()],
        crate::ID,
    )
    .readonly();

    let user_token_account = get_associated_token_address_with_program_id(
        &release_or_mint_in.receiver,
        &token_mint.key(),
        &token_program.key(),
    )
    .writable();

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

    response.accounts_to_save.extend_from_slice(&[
        ccip_offramp_pool_signer,
        user_token_account,
        token_billing_config,
        pool_chain_config,
    ]);
    response
        .accounts_to_save
        .extend(
            lookup_table_account
                .addresses
                .iter()
                .enumerate()
                .map(|(i, a)| match i {
                    // PoolConfig, PoolTokenAccount and Mint from the LUT are writable.
                    3 | 4 | 7 => a.writable(),
                    _ => a.readonly(),
                }),
        );

    response.look_up_tables_to_save.push(*lut.key);
    response.current_stage = DeriveExecuteAccountsStage::TokenTransferAccounts {
        token_substage: "Start".to_string(),
    }
    .to_string();

    Ok(response)
}
