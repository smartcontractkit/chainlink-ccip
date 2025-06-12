use std::{
    fmt::{self, Display, Formatter},
    str::FromStr,
};

use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id, token::spl_token::native_mint,
};
use ccip_common::{router_accounts::TokenAdminRegistry, seed, CommonCcipError};
use solana_program::address_lookup_table::state::AddressLookupTable;

use crate::{
    context::ViewConfigOnly,
    state::{
        CcipAccountMeta, DeriveAccountsCcipSendParams, DeriveAccountsResponse, DerivedLookupTable,
        ToMeta,
    },
    CcipRouterError,
};

// Local helper to find a readonly CCIP meta for a given seed + program_id combo.
// Short name for compactness.
fn find(seeds: &[&[u8]], program_id: Pubkey) -> CcipAccountMeta {
    Pubkey::find_program_address(seeds, &program_id)
        .0
        .readonly()
}

#[derive(Copy, Clone, Debug)]
pub enum DeriveAccountsCcipSendStage {
    Start,
    FinishMainAccountList,
    RetrieveTokenLUTs,
    // N stages, one per token.
    TokenTransferAccounts,
}

impl Display for DeriveAccountsCcipSendStage {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        f.write_str(match self {
            DeriveAccountsCcipSendStage::Start => "Start",
            DeriveAccountsCcipSendStage::FinishMainAccountList => "FinishMainAccountList",
            DeriveAccountsCcipSendStage::RetrieveTokenLUTs => "RetrieveTokenLUTs",
            DeriveAccountsCcipSendStage::TokenTransferAccounts => "TokenTransferAccounts",
        })
    }
}

impl FromStr for DeriveAccountsCcipSendStage {
    type Err = CcipRouterError;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "start" => Ok(Self::Start),
            "finishmainaccountlist" => Ok(Self::FinishMainAccountList),
            "retrievetokenluts" => Ok(Self::RetrieveTokenLUTs),
            "tokentransferaccounts" => Ok(Self::TokenTransferAccounts),
            _ => Err(CcipRouterError::InvalidDerivationStage),
        }
    }
}

pub fn derive_ccip_send_accounts_start(
    DeriveAccountsCcipSendParams {
        dest_chain_selector,
        ccip_send_caller,
        fee_token_mint,
        ..
    }: DeriveAccountsCcipSendParams,
) -> Result<DeriveAccountsResponse> {
    let selector = dest_chain_selector.to_le_bytes();
    let accounts_to_save = vec![
        find(&[seed::CONFIG], crate::ID),
        find(&[seed::DEST_CHAIN_STATE, &selector], crate::ID).writable(),
        find(
            &[seed::NONCE, &selector, ccip_send_caller.key().as_ref()],
            crate::ID,
        )
        .writable(),
        ccip_send_caller.writable().signer(),
        solana_program::system_program::ID.readonly(),
    ];

    let ask_again_with = vec![
        fee_token_mint.readonly(),
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
    DeriveAccountsCcipSendParams {
        dest_chain_selector,
        ccip_send_caller,
        fee_token_mint,
        mints_of_transferred_tokens,
    }: DeriveAccountsCcipSendParams,
) -> Result<DeriveAccountsResponse> {
    let selector = dest_chain_selector.to_le_bytes();
    let fee_token_mint_info = &ctx.remaining_accounts[0];
    let fee_billing_signer = &ctx.remaining_accounts[1];
    let fee_token_program = fee_token_mint_info.owner;

    let is_fee_in_sol = *fee_token_program == Pubkey::default();

    let fee_token_user_ata = if is_fee_in_sol {
        Pubkey::default().readonly()
    } else {
        get_associated_token_address_with_program_id(
            &ccip_send_caller.key(),
            &fee_token_mint,
            &fee_token_program.key(),
        )
        .writable()
    };

    let fee_token_receiver = get_associated_token_address_with_program_id(
        fee_billing_signer.key,
        &fee_token_mint,
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
                    fee_token_mint.as_ref()
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

    // If there are no tokens, we're done. If there are tokens, we
    // start by reading the registries on next stage.
    let ask_again_with: Vec<_> = mints_of_transferred_tokens
        .iter()
        .map(|mint| find(&[seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()], crate::ID))
        .collect();

    let next_stage = if ask_again_with.is_empty() {
        "".to_string()
    } else {
        DeriveAccountsCcipSendStage::RetrieveTokenLUTs.to_string()
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
    let ask_again_with = ctx
        .remaining_accounts
        .iter()
        .map(|registry| {
            let token_admin_registry_account: Account<TokenAdminRegistry> =
                Account::try_from(registry).expect("parsing token admin registry account");
            token_admin_registry_account.lookup_table.readonly()
        })
        .collect();

    Ok(DeriveAccountsResponse {
        accounts_to_save: vec![],
        ask_again_with,
        look_up_tables_to_save: vec![],
        current_stage: DeriveAccountsCcipSendStage::RetrieveTokenLUTs.to_string(),
        next_stage: DeriveAccountsCcipSendStage::TokenTransferAccounts.to_string(),
    })
}

pub fn derive_execute_accounts_additional_tokens<'info>(
    ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
    DeriveAccountsCcipSendParams {
        dest_chain_selector,
        ccip_send_caller,
        ..
    }: DeriveAccountsCcipSendParams,
) -> Result<DeriveAccountsResponse> {
    let selector = dest_chain_selector.to_le_bytes();
    // At least one LUT.
    require_gte!(
        ctx.remaining_accounts.len(),
        1,
        CcipRouterError::InvalidAccountListForPdaDerivation
    );

    let lut = &ctx.remaining_accounts[0];
    let lookup_table_data = &mut &lut.data.borrow()[..];
    let lookup_table_account: AddressLookupTable =
        AddressLookupTable::deserialize(lookup_table_data)
            .map_err(|_| CommonCcipError::InvalidInputsLookupTableAccounts)?;
    let pool_program = lookup_table_account.addresses[2];
    let token_program = lookup_table_account.addresses[6];
    let token_mint = lookup_table_account.addresses[7];

    let sender_token_account = get_associated_token_address_with_program_id(
        &ccip_send_caller,
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
    accounts_to_save.extend(lookup_table_account.addresses.iter().enumerate().map(
        |(i, a)| match i {
            // PoolConfig, PoolTokenAccount and Mint from the LUT are writable.
            3 | 4 | 7 => a.writable(),
            _ => a.readonly(),
        },
    ));

    let mut ask_again_with = vec![];
    let next_stage = if ctx.remaining_accounts.len() > 1 {
        // We aren't done yet, we need to derive more tokens, so we tell the user
        // to ask again with one fewer token.
        ask_again_with.extend(ctx.remaining_accounts[1..].iter().map(|a| a.key.readonly()));
        DeriveAccountsCcipSendStage::TokenTransferAccounts.to_string()
    } else {
        "".to_string()
    };

    Ok(DeriveAccountsResponse {
        ask_again_with,
        accounts_to_save,
        look_up_tables_to_save: vec![DerivedLookupTable {
            address: *lut.key,
            accounts: lookup_table_account.addresses.iter().cloned().collect(),
        }],
        current_stage: DeriveAccountsCcipSendStage::TokenTransferAccounts.to_string(),
        next_stage,
    })
}
