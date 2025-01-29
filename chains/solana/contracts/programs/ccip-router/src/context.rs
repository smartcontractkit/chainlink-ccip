use anchor_lang::prelude::*;
use anchor_spl::associated_token::{get_associated_token_address_with_program_id, AssociatedToken};
use anchor_spl::token::spl_token::native_mint;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};
use solana_program::sysvar::instructions;

use crate::program::CcipRouter;
use crate::state::{CommitReport, Config, Nonce};
use crate::{
    BillingTokenConfig, BillingTokenConfigWrapper, CcipRouterError, DestChain,
    ExecutionReportSingleChain, ExternalExecutionConfig, GlobalState, SVM2AnyMessage, SourceChain,
};

/// Static space allocated to any account: must always be added to space calculations.
pub const ANCHOR_DISCRIMINATOR: usize = 8;

/// Maximum acceptable config version accepted by this module: any accounts with higher
/// version numbers than this will be rejected.
pub const MAX_CONFIG_V: u8 = 1;
const MAX_CHAINSTATE_V: u8 = 1;
const MAX_NONCE_V: u8 = 1;
const MAX_COMMITREPORT_V: u8 = 1;

pub const fn valid_version(v: u8, max_version: u8) -> bool {
    !uninitialized(v) && v <= max_version
}

pub const fn uninitialized(v: u8) -> bool {
    v == 0
}

/// Fixed seeds for PDA derivation: different context must use different seeds.
pub mod seed {
    pub const DEST_CHAIN_STATE: &[u8] = b"dest_chain_state";
    pub const SOURCE_CHAIN_STATE: &[u8] = b"source_chain_state";
    pub const COMMIT_REPORT: &[u8] = b"commit_report";
    pub const NONCE: &[u8] = b"nonce";
    pub const CONFIG: &[u8] = b"config";
    pub const STATE: &[u8] = b"state";
    pub const EXTERNAL_EXECUTION_CONFIG: &[u8] = b"external_execution_config";

    // arbitrary messaging signer
    pub const EXTERNAL_TOKEN_POOL: &[u8] = b"external_token_pools_signer";
    // token pool interaction signer
    pub const FEE_BILLING_SIGNER: &[u8] = b"fee_billing_signer";
    // signer for billing fee token transfer
    pub const FEE_BILLING_TOKEN_CONFIG: &[u8] = b"fee_billing_token_config";

    // token specific
    pub const TOKEN_ADMIN_REGISTRY: &[u8] = b"token_admin_registry";
    pub const CCIP_TOKENPOOL_CONFIG: &[u8] = b"ccip_tokenpool_config";
    pub const CCIP_TOKENPOOL_SIGNER: &[u8] = b"ccip_tokenpool_signer";
    pub const TOKEN_POOL_BILLING: &[u8] = b"ccip_tokenpool_billing";
    pub const TOKEN_POOL_CONFIG: &[u8] = b"ccip_tokenpool_chainconfig";
}

/// Input from an offchain node, containing the Merkle root and interval for
/// the source chain, and optionally some price updates alongside it.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct CommitInput {
    pub price_updates: PriceUpdates,
    pub merkle_root: MerkleRoot,
    pub rmn_signatures: Vec<[u8; 64]>, // placeholder for stable interface; r = 32, s = 32; https://github.com/smartcontractkit/chainlink/blob/d1a9f8be2f222ea30bdf7182aaa6428bfa605cf7/contracts/src/v0.8/ccip/interfaces/IRMNRemote.sol#L9
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct PriceUpdates {
    pub token_price_updates: Vec<TokenPriceUpdate>,
    pub gas_price_updates: Vec<GasPriceUpdate>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct TokenPriceUpdate {
    pub source_token: Pubkey, // It is the mint, but called "token" for EVM compatibility.
    pub usd_per_token: [u8; 28], // EVM uses u224, 1e18 USD per 1e18 of the smallest token denomination.
}

/// Gas price for a given chain in USD; its value may contain tightly packed fields.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct GasPriceUpdate {
    pub dest_chain_selector: u64,
    pub usd_per_unit_gas: [u8; 28], // EVM uses u224, 1e18 USD per smallest unit (e.g. wei) of destination chain gas
}

/// Struct to hold a merkle root and an interval for a source chain
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct MerkleRoot {
    pub source_chain_selector: u64, // Remote source chain selector that the Merkle Root is scoped to
    pub on_ramp_address: Vec<u8>,   // Generic onramp address, to support arbitrary sources
    pub min_seq_nr: u64,            // Minimum sequence number, inclusive
    pub max_seq_nr: u64,            // Maximum sequence number, inclusive
    pub merkle_root: [u8; 32],      // Merkle root covering the interval & source chain messages
}

impl MerkleRoot {
    pub fn len(&self) -> usize {
        8  + // source chain selector
        4 + self.on_ramp_address.len() + // on ramp address
        8 + // min msg nr
        8 + // max msg nr
        32 // root
    }
}

#[derive(Accounts)]
pub struct WithdrawBilledFunds<'info> {
    #[account(
        owner = token_program.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        mut,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer,
        associated_token::token_program = token_program,
    )]
    pub fee_token_accum: InterfaceAccount<'info, TokenAccount>,

    #[account(
        mut,
        constraint = recipient.key() == get_associated_token_address_with_program_id(
            &config.load()?.fee_aggregator.key(), &fee_token_mint.key(), &token_program.key()
        ) @ CcipRouterError::InvalidInputs,
    )]
    pub recipient: InterfaceAccount<'info, TokenAccount>,

    pub token_program: Interface<'info, TokenInterface>,

    /// CHECK: This is the signer for the billing CPIs, used here to close the receiver token account
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct InitializeCCIPRouter<'info> {
    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        init,
        seeds = [seed::STATE],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + GlobalState::INIT_SPACE,
    )]
    pub state: Account<'info, GlobalState>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CcipRouter>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipRouterError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,

    #[account(
        init,
        seeds = [seed::EXTERNAL_EXECUTION_CONFIG],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExternalExecutionConfig::INIT_SPACE,
    )]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>, // messaging CPI signer initialization

    #[account(
        init,
        seeds = [seed::EXTERNAL_TOKEN_POOL],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExternalExecutionConfig::INIT_SPACE,
    )]
    pub token_pools_signer: Account<'info, ExternalExecutionConfig>, // token pool CPI signer initialization
}

#[derive(Accounts)]
pub struct TransferOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.proposed_owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
pub struct AddChainSelector<'info> {
    /// Adding a chain selector implies initializing the state for a new chain,
    /// hence the need to initialize two accounts.

    #[account(
        init,
        seeds = [seed::SOURCE_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + SourceChain::INIT_SPACE,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        init,
        seeds = [seed::DEST_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + DestChain::INIT_SPACE,
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
pub struct UpdateSourceChainSelectorConfig<'info> {
    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
pub struct UpdateDestChainSelectorConfig<'info> {
    #[account(
        mut,
        seeds = [seed::DEST_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs,
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct UpdateConfigCCIPRouter<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdateSupportedChainsConfigCCIPRouter<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct SetOcrConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::STATE],
        bump,
    )]
    pub state: Account<'info, GlobalState>,

    #[account(address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(token_config: BillingTokenConfig)]
pub struct AddBillingTokenConfig<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        init,
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG, token_config.mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + BillingTokenConfigWrapper::INIT_SPACE,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    pub token_program: Interface<'info, TokenInterface>,

    #[account(
        owner = token_program.key() @ CcipRouterError::InvalidInputs,
        constraint = token_config.mint == fee_token_mint.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        init,
        payer = authority,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer,
        associated_token::token_program = token_program,
    )]
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>,

    #[account(
        mut,
        address = config.load()?.owner @ CcipRouterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    /// CHECK: This is the signer for the billing CPIs, used here to close the receiver token account
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    pub associated_token_program: Program<'info, AssociatedToken>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(token_config: BillingTokenConfig)]
pub struct UpdateBillingTokenConfig<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG, token_config.mint.key().as_ref()],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    #[account(
        address = config.load()?.owner @ CcipRouterError::Unauthorized
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct RemoveBillingTokenConfig<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        close = authority,
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG, fee_token_mint.key().as_ref()],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    pub token_program: Interface<'info, TokenInterface>,

    #[account(
        owner = token_program.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        mut,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer, // use the signer account as the authority
        associated_token::token_program = token_program,
        constraint = fee_token_receiver.amount == 0 @ CcipRouterError::InvalidInputs, // ensure the account is empty // TODO improve error
    )]
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>,

    /// CHECK: This is the signer for the billing CPIs, used here to close the receiver token account

    #[account(
        mut,
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    #[account(
        mut,
        address = config.load()?.owner @ CcipRouterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(destination_chain_selector: u64, message: SVM2AnyMessage)]
pub struct CcipSend<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::DEST_CHAIN_STATE, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs,
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        init_if_needed,
        seeds = [seed::NONCE, destination_chain_selector.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Nonce::INIT_SPACE,
        constraint = uninitialized(nonce.version) || valid_version(nonce.version, MAX_NONCE_V) @ CcipRouterError::InvalidInputs,
    )]
    pub nonce: Account<'info, Nonce>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /////////////
    // billing //
    /////////////
    pub fee_token_program: Interface<'info, TokenInterface>,

    #[account(
        owner = fee_token_program.key() @ CcipRouterError::InvalidInputs,
        constraint = (message.fee_token == Pubkey::default() && fee_token_mint.key() == native_mint::ID)
            || message.fee_token.key() == fee_token_mint.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>, // pass pre-2022 wSOL if using native SOL

    /// CHECK: This is the associated token account for the user paying the fee.
    /// If paying with native SOL, this must be the zero address.
    #[account(
        // address must be either zero (paying with native SOL) or must be a WRITABLE associated token account
        constraint = (message.fee_token == Pubkey::default() && fee_token_user_associated_account.key() == Pubkey::default())
            || fee_token_user_associated_account.is_writable @ CcipRouterError::InvalidInputsAtaWritable,
        // address must be either zero (paying with native SOL) or
        // the associated token account address for the caller and fee token used
        constraint = (message.fee_token == Pubkey::default() && fee_token_user_associated_account.key() == Pubkey::default())
            || fee_token_user_associated_account.key() == get_associated_token_address_with_program_id(
                &authority.key(),
                &fee_token_mint.key(),
                &fee_token_program.key(),
            ) @ CcipRouterError::InvalidInputsAtaAddress,
    )]
    pub fee_token_user_associated_account: UncheckedAccount<'info>,

    #[account(
        mut,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer,
        associated_token::token_program = fee_token_program,
    )]
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>, // pass pre-2022 wSOL receiver if using native SOL

    /// CHECK: This is the signer for the billing transfer CPI.
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    ////////////////////
    // fee quoter CPI //
    ////////////////////
    /// CHECK: This is the account for the Fee Quoter program
    #[account(
        address = config.load()?.fee_quoter @ CcipRouterError::InvalidInputs,
    )]
    pub fee_quoter: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::CONFIG],
        bump,
        seeds::program = config.load()?.fee_quoter,
    )]
    pub fee_quoter_config: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::DEST_CHAIN, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        seeds::program = config.load()?.fee_quoter,
    )]
    pub fee_quoter_dest_chain: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::FEE_BILLING_TOKEN_CONFIG,
            if message.fee_token == Pubkey::default() {
                native_mint::ID.as_ref() // pre-2022 WSOL
            } else {
                message.fee_token.as_ref()
            },
        ],
        bump,
        seeds::program = config.load()?.fee_quoter,
    )]
    pub fee_quoter_billing_token_config: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::FEE_BILLING_TOKEN_CONFIG,
            config.load()?.link_token_mint.key().as_ref(),
        ],
        bump,
        seeds::program = config.load()?.fee_quoter,
    )]
    pub fee_quoter_link_token_config: UncheckedAccount<'info>,

    /// CPI signers, optional if no tokens are being transferred.
    /// CHECK: Using this to sign.
    #[account(mut, seeds = [seed::EXTERNAL_TOKEN_POOL], bump)]
    pub token_pools_signer: Account<'info, ExternalExecutionConfig>,
    // remaining accounts (not explicitly listed)
    // [
    // user/sender token account (must be associated token account - derivable PDA [wallet_addr, token_program, mint])
    // ccip pool chain config (ccip: billing, ccip admin controlled - derivable PDA [chain_selector, mint])
    // pool chain config (pool: custom configs that may include rate limits & remote chain configs, pool admin controlled - derivable [chain_selector, mint])
    // token pool lookup table
    // token registry PDA
    // pool program
    // pool config
    // pool token account (must be associated token account - derivable PDA [wallet_addr, token_program, mint])
    // pool signer
    // token program
    // token mint
    // ...additional accounts for pool config
    // ] x N tokens
}

#[derive(Accounts)]
#[instruction(_report_context_byte_words: [[u8; 32]; 2], raw_report: Vec<u8>)]
pub struct CommitReportContext<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN_STATE, CommitInput::deserialize(&mut raw_report.as_ref())?.merkle_root.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        init,
        seeds = [seed::COMMIT_REPORT, CommitInput::deserialize(&mut raw_report.as_ref())?.merkle_root.source_chain_selector.to_le_bytes().as_ref(), CommitInput::deserialize(&mut raw_report.as_ref())?.merkle_root.merkle_root.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + CommitReport::INIT_SPACE,
    )]
    pub commit_report: Account<'info, CommitReport>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is the sysvar instructions account
    #[account(address = instructions::ID @ CcipRouterError::InvalidInputs)]
    pub sysvar_instructions: UncheckedAccount<'info>,
    // remaining accounts
    // global state account (to update the price sequence number)
    // [...billingTokenConfig accounts]
    // [...chainConfig accounts]
}

#[derive(Accounts)]
#[instruction(raw_report: Vec<u8>)]
pub struct ExecuteReportContext<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        seeds = [seed::SOURCE_CHAIN_STATE, ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        mut,
        seeds = [seed::COMMIT_REPORT, ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.source_chain_selector.to_le_bytes().as_ref(), ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.root.as_ref()],
        bump,
        constraint = valid_version(commit_report.version, MAX_COMMITREPORT_V) @ CcipRouterError::InvalidInputs,
    )]
    pub commit_report: Account<'info, CommitReport>,

    /// CHECK: Using this to sign
    #[account(seeds = [seed::EXTERNAL_EXECUTION_CONFIG], bump)]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is a sysvar account
    #[account(address = instructions::ID @ CcipRouterError::InvalidInputs)]
    pub sysvar_instructions: AccountInfo<'info>,

    #[account(seeds = [seed::EXTERNAL_TOKEN_POOL], bump)]
    pub token_pools_signer: Account<'info, ExternalExecutionConfig>,
    // remaining accounts
    // [receiver_program, receiver_account, ...user specified accounts from message data for arbitrary messaging]
    // +
    // [
    // user/sender token account (must be associated token account - derivable PDA [wallet_addr, token_program, mint])
    // per chain per token config (ccip: billing, ccip admin controlled - derivable PDA [chain_selector, mint])
    // pool chain config (pool: custom configs that may include rate limits & remote chain configs, pool admin controlled - derivable [chain_selector, mint])
    // token pool lookup table
    // token registry PDA
    // pool program
    // pool config
    // pool token account (must be associated token account - derivable PDA [wallet_addr, token_program, mint])
    // pool signer
    // token program
    // token mint
    // ...additional accounts for pool config
    // ] x N tokens
}

#[derive(Copy, Clone, AnchorSerialize, AnchorDeserialize)]
pub enum OcrPluginType {
    Commit,
    Execution,
}
