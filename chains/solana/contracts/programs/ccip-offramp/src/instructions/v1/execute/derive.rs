use crate::{
    context::ViewConfigOnly,
    state::{CcipAccountMeta, DeriveAccountsResponse, ReferenceAddresses, ToMeta},
    CcipOfframpError,
};
use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use ccip_common::{
    router_accounts::TokenAdminRegistry,
    seed::{self, EXECUTION_REPORT_BUFFER},
    CommonCcipError,
};
use solana_program::{address_lookup_table::state::AddressLookupTable, sysvar::instructions};
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

#[derive(Copy, Clone, Debug)]
pub enum DeriveExecuteAccountsStage {
    GatherBasicInfo,
    BuildMainAccountList,
    RetrieveTokenLUTs,
    // N stages, one per token.
    TokenTransferAccounts,
}

impl Display for DeriveExecuteAccountsStage {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        f.write_str(match self {
            DeriveExecuteAccountsStage::GatherBasicInfo => "GatherBasicInfo",
            DeriveExecuteAccountsStage::BuildMainAccountList => "BuildMainAccountList",
            DeriveExecuteAccountsStage::RetrieveTokenLUTs => "RetrieveTokenLookupTables",
            DeriveExecuteAccountsStage::TokenTransferAccounts => "TokenTransferAccounts",
        })
    }
}

impl FromStr for DeriveExecuteAccountsStage {
    type Err = CcipOfframpError;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "start" | "gatherbasicinfo" => Ok(Self::GatherBasicInfo),
            "buildmainaccountlist" => Ok(Self::BuildMainAccountList),
            "retrievetokenlookuptables" => Ok(Self::RetrieveTokenLUTs),
            "tokentransferaccounts" => Ok(Self::TokenTransferAccounts),
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
    source_chain_selector: u64,
    merkle_root: &[u8; 32],
    execute_caller: Pubkey,
    mints_of_transferred_tokens: &[Pubkey],
    message_accounts: &[CcipAccountMeta],
    buffer_id: &[u8],
) -> Result<DeriveAccountsResponse> {
    let ReferenceAddresses {
        router, rmn_remote, ..
    } = *AccountLoader::<'info, ReferenceAddresses>::try_from(&ctx.remaining_accounts[1])?
        .load()?;

    let selector = source_chain_selector.to_le_bytes();

    let mut accounts_to_save = vec![
        find(&[seed::COMMIT_REPORT, &selector, merkle_root], crate::ID).writable(),
        crate::ID.readonly(),
        find(
            &[seed::ALLOWED_OFFRAMP, &selector, crate::ID.as_ref()],
            router,
        ),
        execute_caller.writable().signer(),
        solana_program::system_program::ID.readonly(),
        instructions::ID.readonly(),
        rmn_remote.readonly(),
        find(&[seed::CURSES], rmn_remote),
        find(&[seed::CONFIG], rmn_remote),
    ];

    if !message_accounts.is_empty() {
        let find = |seeds, id| Pubkey::find_program_address(seeds, id).0.readonly();
        accounts_to_save.push(message_accounts[0].clone());
        accounts_to_save.push(find(
            &[
                seed::EXTERNAL_EXECUTION_CONFIG,
                message_accounts[0].pubkey.as_ref(),
            ],
            &crate::ID,
        ));
        accounts_to_save.extend_from_slice(&message_accounts[1..]);
    }

    // If there are no tokens, we're done. If there are tokens, we
    // start by reading the registries on next stage.
    let mut ask_again_with: Vec<_> = mints_of_transferred_tokens
        .iter()
        .map(|mint| find(&[seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()], router))
        .collect();

    if !buffer_id.is_empty() {
        // We're in the buffered case, so we need the buffer PDA either on the derived account list
        // or on the next stage
        let buffer_pda = find(
            &[EXECUTION_REPORT_BUFFER, buffer_id, execute_caller.as_ref()],
            crate::ID,
        );
        if ask_again_with.is_empty() {
            accounts_to_save.push(buffer_pda.writable());
        } else {
            ask_again_with.push(buffer_pda.readonly());
        }
    }

    let next_stage = if ask_again_with.is_empty() {
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
    let lookup_table_metas = ctx.remaining_accounts.iter().map(|registry| {
        let token_admin_registry_account: Account<TokenAdminRegistry> =
            Account::try_from(registry).expect("parsing token admin registry account");
        token_admin_registry_account.lookup_table.readonly()
    });

    let mut ask_again_with = vec![find(&[seed::REFERENCE_ADDRESSES], crate::ID)];
    ask_again_with.extend(lookup_table_metas.clone());

    Ok(DeriveAccountsResponse {
        accounts_to_save: vec![],
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveExecuteAccountsStage::RetrieveTokenLUTs.to_string(),
        next_stage: DeriveExecuteAccountsStage::TokenTransferAccounts.to_string(),
    })
}

// We derive accounts for each token in a separate step, to ensure we don't blow up the response size.
pub fn derive_execute_accounts_additional_tokens<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    execute_caller: Pubkey,
    token_receiver: Pubkey,
    source_chain_selector: u64,
    buffer_id: &[u8],
) -> Result<DeriveAccountsResponse> {
    // The reference addresses account and at least one LUT.
    require_gte!(
        ctx.remaining_accounts.len(),
        2,
        CcipOfframpError::InvalidAccountListForPdaDerivation
    );

    let selector = source_chain_selector.to_le_bytes();
    let ReferenceAddresses { fee_quoter, .. } =
        *AccountLoader::<'info, ReferenceAddresses>::try_from(&ctx.remaining_accounts[0])?
            .load()?;

    let lut = &ctx.remaining_accounts[1];
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
        &token_receiver,
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

    let mut accounts_to_save = vec![
        ccip_offramp_pool_signer,
        user_token_account,
        token_billing_config,
        pool_chain_config,
    ];
    accounts_to_save.extend(lookup_table_account.addresses.iter().enumerate().map(
        |(i, a)| match i {
            // PoolConfig, PoolTokenAccount and Mint from the LUT are writable.
            3 | 4 | 7 => a.writable(),
            _ => a.readonly(),
        },
    ));

    let mut ask_again_with = vec![];

    let next_stage = if ctx.remaining_accounts.len() > 2 {
        // We aren't done yet, we need to derive more tokens, so we tell the user
        // to ask again with one fewer token.
        ask_again_with.push(find(&[seed::REFERENCE_ADDRESSES], crate::ID));
        ask_again_with.extend(ctx.remaining_accounts[2..].iter().map(|a| a.key.readonly()));
        DeriveExecuteAccountsStage::TokenTransferAccounts.to_string()
    } else {
        if !buffer_id.is_empty() {
            // We're done and it's the buffered case, so we append the buffer PDA
            let buffer_pda = find(
                &[
                    seed::EXECUTION_REPORT_BUFFER,
                    buffer_id,
                    execute_caller.as_ref(),
                ],
                crate::ID,
            );
            ask_again_with.push(buffer_pda.writable());
        }
        "".to_string()
    };

    Ok(DeriveAccountsResponse {
        ask_again_with,
        accounts_to_save,
        look_up_tables_to_save: vec![*lut.key],
        current_stage: DeriveExecuteAccountsStage::TokenTransferAccounts.to_string(),
        next_stage,
    })
}
