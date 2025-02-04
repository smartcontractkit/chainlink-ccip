use anchor_lang::prelude::*;
use solana_program::sysvar::instructions;

use crate::messages::ExecutionReportSingleChain;
use crate::program::CcipOfframp;
use crate::state::{
    CommitReport, Config, ExternalExecutionConfig, GlobalState, ReferenceAddresses, SourceChain,
    SourceChainConfig,
};
use crate::CcipOfframpError;

/// Static space allocated to any account: must always be added to space calculations.
pub const ANCHOR_DISCRIMINATOR: usize = 8;

/// Maximum acceptable config version accepted by this module: any accounts with higher
/// version numbers than this will be rejected.
pub const MAX_CONFIG_V: u8 = 1;
const MAX_CHAINSTATE_V: u8 = 1;
const MAX_COMMITREPORT_V: u8 = 1;

pub const fn valid_version(v: u8, max_version: u8) -> bool {
    !uninitialized(v) && v <= max_version
}

pub const fn uninitialized(v: u8) -> bool {
    v == 0
}

/// Fixed seeds for PDA derivation: different context must use different seeds.
pub mod seed {
    pub const SOURCE_CHAIN: &[u8] = b"source_chain_state";
    pub const COMMIT_REPORT: &[u8] = b"commit_report";
    pub const CONFIG: &[u8] = b"config";
    pub const REFERENCE_ADDRESSES: &[u8] = b"reference_addresses";
    pub const STATE: &[u8] = b"state";

    // arbitrary messaging signer
    pub const EXTERNAL_EXECUTION_CONFIG: &[u8] = b"external_execution_config";
    // token pool interaction signer
    pub const EXTERNAL_TOKEN_POOL: &[u8] = b"external_token_pools_signer";
    // signer for billing fee token transfer
    pub const FEE_BILLING_SIGNER: &[u8] = b"fee_billing_signer";

    // token specific
    pub const TOKEN_ADMIN_REGISTRY: &[u8] = b"token_admin_registry";
    pub const CCIP_TOKENPOOL_CONFIG: &[u8] = b"ccip_tokenpool_config";
    pub const CCIP_TOKENPOOL_SIGNER: &[u8] = b"ccip_tokenpool_signer";
    pub const TOKEN_POOL_CONFIG: &[u8] = b"ccip_tokenpool_chainconfig";
}

#[derive(Accounts)]
pub struct Initialize<'info> {
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
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ReferenceAddresses::INIT_SPACE,
    )]
    pub reference_addresses: Account<'info, ReferenceAddresses>,

    #[account(
        init,
        seeds = [seed::STATE],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + GlobalState::INIT_SPACE,
    )]
    pub state: Account<'info, GlobalState>,

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

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CcipOfframp>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipOfframpError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
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
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct TransferOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.proposed_owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64, _source_chain_config: SourceChainConfig)]
pub struct AddSourceChain<'info> {
    /// Adding a chain selector implies initializing the state for a new chain,
    /// hence the need to initialize two accounts.
    #[account(
        init,
        seeds = [seed::SOURCE_CHAIN, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + SourceChain::INIT_SPACE,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
pub struct UpdateSourceChain<'info> {
    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain_state.version, MAX_CHAINSTATE_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct SetOcrConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::STATE],
        bump,
    )]
    pub state: Account<'info, GlobalState>,

    #[account(address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(_report_context_byte_words: [[u8; 32]; 2], raw_report: Vec<u8>)]
pub struct CommitReportContext<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        constraint = valid_version(reference_addresses.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub reference_addresses: Account<'info, ReferenceAddresses>,

    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN, CommitInput::deserialize(&mut raw_report.as_ref())?.merkle_root.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain_state.version, MAX_CHAINSTATE_V) @ CcipOfframpError::InvalidInputs,
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
    #[account(address = instructions::ID @ CcipOfframpError::InvalidInputs)]
    pub sysvar_instructions: UncheckedAccount<'info>,

    /// CHECK: This is the signer for the billing CPIs, used here to invoke fee quoter with price updates
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    /// CHECK: fee quoter program account, used to invoke fee quoter with price updates
    #[account(
        address = reference_addresses.fee_quoter @ CcipOfframpError::InvalidInputs,
    )]
    pub fee_quoter: UncheckedAccount<'info>,

    /// CHECK: fee quoter config account, used to invoke fee quoter with price updates
    #[account(
        seeds = [fee_quoter::context::seed::CONFIG],
        bump,
        seeds::program = reference_addresses.fee_quoter,
    )]
    pub fee_quoter_config: UncheckedAccount<'info>,
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
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        constraint = valid_version(reference_addresses.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub reference_addresses: Account<'info, ReferenceAddresses>,

    #[account(
        seeds = [seed::SOURCE_CHAIN, ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain_state.version, MAX_CHAINSTATE_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub source_chain_state: Account<'info, SourceChain>,

    #[account(
        mut,
        seeds = [seed::COMMIT_REPORT, ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.source_chain_selector.to_le_bytes().as_ref(), ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.root.as_ref()],
        bump,
        constraint = valid_version(commit_report.version, MAX_COMMITREPORT_V) @ CcipOfframpError::InvalidInputs,
    )]
    pub commit_report: Account<'info, CommitReport>,

    /// CHECK: Using this to sign
    #[account(seeds = [seed::EXTERNAL_EXECUTION_CONFIG], bump)]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is a sysvar account
    #[account(address = instructions::ID @ CcipOfframpError::InvalidInputs)]
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
