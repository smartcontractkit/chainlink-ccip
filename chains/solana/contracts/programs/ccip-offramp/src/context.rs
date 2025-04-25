use anchor_lang::prelude::*;
use anchor_spl::token::spl_token::native_mint;
use anchor_spl::token::{Mint, Token, TokenAccount};
use bytemuck::{Pod, Zeroable};
use ccip_common::seed;
use solana_program::sysvar::instructions;

use crate::messages::ExecutionReportSingleChain;
use crate::program::CcipOfframp;
use crate::state::{
    CommitReport, Config, GlobalState, ReferenceAddresses, SourceChain, SourceChainConfig,
};
use crate::CcipOfframpError;

/// Static space allocated to any account: must always be added to space calculations.
pub const ANCHOR_DISCRIMINATOR: usize = 8;

/// Maximum acceptable config version accepted by this module: any accounts with higher
/// version numbers than this will be rejected.
pub const MAX_CONFIG_V: u8 = 1;
const MAX_CHAIN_V: u8 = 1;
const MAX_COMMITREPORT_V: u8 = 1;

pub const fn valid_version(v: u8, max_version: u8) -> bool {
    !uninitialized(v) && v <= max_version
}

pub const fn uninitialized(v: u8) -> bool {
    v == 0
}

#[derive(Accounts)]
pub struct InitializeConfig<'info> {
    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CcipOfframp>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipOfframpError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ReferenceAddresses::INIT_SPACE,
    )]
    pub reference_addresses: AccountLoader<'info, ReferenceAddresses>,

    /// CHECK: Router address, to be persisted in reference_addresses. Just they key is used, no data is read.
    pub router: UncheckedAccount<'info>,
    /// CHECK: FeeQuoter address, to be persisted in reference_addresses. Just they key is used, no data is read.
    pub fee_quoter: UncheckedAccount<'info>,
    /// CHECK: RMNRemote address, to be persisted in reference_addresses. Just they key is used, no data is read.
    pub rmn_remote: UncheckedAccount<'info>,
    /// CHECK: ALT address, to be persisted in reference_addresses. Just they key is used, no data is read.
    pub offramp_lookup_table: UncheckedAccount<'info>,

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
    pub program: Program<'info, CcipOfframp>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipOfframpError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct Empty<'info> {
    // This is unused, but Anchor requires that there is at least one account in the context
    pub clock: Sysvar<'info, Clock>,
}

#[derive(Accounts)]
pub struct UpdateReferenceAddresses<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
    )]
    pub reference_addresses: AccountLoader<'info, ReferenceAddresses>,

    #[account(address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

/// Input from an offchain node, containing the Merkle root and interval for
/// the source chain, and optionally some price updates alongside it
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct CommitInput {
    pub price_updates: PriceUpdates,
    pub merkle_root: Option<MerkleRoot>,
    pub rmn_signatures: Vec<[u8; 64]>, // placeholder for stable interface; r = 32, s = 32; https://github.com/smartcontractkit/chainlink/blob/d1a9f8be2f222ea30bdf7182aaa6428bfa605cf7/contracts/src/v0.8/ccip/interfaces/IRMNRemote.sol#L9
}

impl CommitInput {
    pub fn root(mut raw: &[u8]) -> Result<MerkleRoot> {
        CommitInput::deserialize(&mut raw)?
            .merkle_root
            .ok_or(CcipOfframpError::MissingExpectedMerkleRoot.into())
    }
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
#[derive(Default, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct MerkleRoot {
    pub source_chain_selector: u64, // Remote source chain selector that the Merkle Root is scoped to
    pub on_ramp_address: Vec<u8>,   // Generic onramp address, to support arbitrary sources
    pub min_seq_nr: u64,            // Minimum sequence number, inclusive
    pub max_seq_nr: u64,            // Maximum sequence number, inclusive
    pub merkle_root: [u8; 32],      // Merkle root covering the interval & source chain messages
}

impl MerkleRoot {
    pub fn static_len() -> usize {
        8 + // source chain selector
        4 + // on ramp address, static
        8 + // min msg nr
        8 + // max msg nr
        32 // root
    }

    pub fn len(&self) -> usize {
        Self::static_len() + self.on_ramp_address.len() // on ramp address, dynamic
    }
}

#[derive(Accounts)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
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
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
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
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(address = config.load()?.proposed_owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64, _source_chain_config: SourceChainConfig)]
pub struct AddSourceChain<'info> {
    /// Adding a chain selector implies initializing the state for a new chain
    #[account(
        init,
        seeds = [seed::SOURCE_CHAIN, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + SourceChain::INIT_SPACE,
    )]
    pub source_chain: Account<'info, SourceChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut, address = config.load()?.owner @ CcipOfframpError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(source_chain_selector: u64)]
pub struct UpdateSourceChain<'info> {
    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN, source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain.version, MAX_CHAIN_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub source_chain: Account<'info, SourceChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
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
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
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
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        constraint = valid_version(reference_addresses.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub reference_addresses: AccountLoader<'info, ReferenceAddresses>,

    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN, CommitInput::root(&raw_report)?.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain.version, MAX_CHAIN_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub source_chain: Account<'info, SourceChain>,

    #[account(
        init,
        seeds = [seed::COMMIT_REPORT, CommitInput::root(&raw_report)?.source_chain_selector.to_le_bytes().as_ref(), CommitInput::root(&raw_report)?.merkle_root.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + CommitReport::INIT_SPACE,
    )]
    pub commit_report: Account<'info, CommitReport>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is the sysvar instructions account
    #[account(address = instructions::ID @ CcipOfframpError::InvalidInputsSysvarAccount)]
    pub sysvar_instructions: UncheckedAccount<'info>,

    /// CHECK: This is the signer for the billing CPIs, used here to invoke fee quoter with price updates
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    /// CHECK: fee quoter program account, used to invoke fee quoter with price updates
    #[account(
        address = reference_addresses.load()?.fee_quoter @ CcipOfframpError::InvalidInputsFeeQuoterAccount,
    )]
    pub fee_quoter: UncheckedAccount<'info>,

    /// CHECK: fee quoter allowed price updater account, used to invoke fee quoter with price updates
    /// so that it can authorize the call made by this offramp
    #[account(
        seeds = [seed::ALLOWED_PRICE_UPDATER, fee_billing_signer.key().as_ref()],
        bump,
        seeds::program = fee_quoter.key(),
    )]
    pub fee_quoter_allowed_price_updater: UncheckedAccount<'info>,

    /// CHECK: fee quoter config account, used to invoke fee quoter with price updates
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = reference_addresses.load()?.fee_quoter,
    )]
    pub fee_quoter_config: UncheckedAccount<'info>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = reference_addresses.load()?.rmn_remote @ CcipOfframpError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,
    // remaining accounts
    // global state account (to update the last seen price sequence number)
    // [...billingTokenConfig accounts] fee quoter accounts used to store token prices
    // [...chainConfig accounts] fee quoter accounts used to store gas prices
}

#[derive(Accounts)]
pub struct PriceOnlyCommitReportContext<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        constraint = valid_version(reference_addresses.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub reference_addresses: AccountLoader<'info, ReferenceAddresses>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is the sysvar instructions account
    #[account(address = instructions::ID @ CcipOfframpError::InvalidInputsSysvarAccount)]
    pub sysvar_instructions: UncheckedAccount<'info>,

    /// CHECK: This is the signer for the billing CPIs, used here to invoke fee quoter with price updates
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    /// CHECK: fee quoter program account, used to invoke fee quoter with price updates
    #[account(
        address = reference_addresses.load()?.fee_quoter @ CcipOfframpError::InvalidInputsFeeQuoterAccount,
    )]
    pub fee_quoter: UncheckedAccount<'info>,

    /// CHECK: fee quoter allowed price updater account, used to invoke fee quoter with price updates
    /// so that it can authorize the call made by this offramp
    #[account(
        seeds = [seed::ALLOWED_PRICE_UPDATER, fee_billing_signer.key().as_ref()],
        bump,
        seeds::program = fee_quoter.key(),
    )]
    pub fee_quoter_allowed_price_updater: UncheckedAccount<'info>,

    /// CHECK: fee quoter config account, used to invoke fee quoter with price updates
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = reference_addresses.load()?.fee_quoter,
    )]
    pub fee_quoter_config: UncheckedAccount<'info>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = reference_addresses.load()?.rmn_remote @ CcipOfframpError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,
    // remaining accounts
    // global state account (to update the last seen price sequence number)
    // [...billingTokenConfig accounts] fee quoter accounts used to store token prices
    // [...chainConfig accounts] fee quoter accounts used to store gas prices
}

const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";

#[derive(Accounts)]
#[instruction(raw_report: Vec<u8>)]
pub struct ExecuteReportContext<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        constraint = valid_version(reference_addresses.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub reference_addresses: AccountLoader<'info, ReferenceAddresses>,

    #[account(
        seeds = [seed::SOURCE_CHAIN, ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(source_chain.version, MAX_CHAIN_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub source_chain: Account<'info, SourceChain>,

    #[account(
        mut,
        seeds = [seed::COMMIT_REPORT, ExecutionReportSingleChain::deserialize(&mut raw_report.as_ref())?.source_chain_selector.to_le_bytes().as_ref(), commit_report.merkle_root.as_ref()],
        bump,
        constraint = valid_version(commit_report.version, MAX_COMMITREPORT_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub commit_report: Account<'info, CommitReport>,

    pub offramp: Program<'info, CcipOfframp>,

    /// CHECK PDA of the router program verifying the signer is an allowed offramp.
    /// If PDA does not exist, the router doesn't allow this offramp. This is just used
    /// so that token pools and receivers can then check that the caller is an actual offramp that
    /// has been registered in the router as such for that source chain.
    #[account(
        owner = reference_addresses.load()?.router @ CcipOfframpError::InvalidInputsAllowedOfframpAccount, // this guarantees that it was initialized
        seeds = [ALLOWED_OFFRAMP, source_chain.chain_selector.to_le_bytes().as_ref(), offramp.key().as_ref()],
        bump,
        seeds::program = reference_addresses.load()?.router,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is a sysvar account
    #[account(address = instructions::ID @ CcipOfframpError::InvalidInputsSysvarAccount)]
    pub sysvar_instructions: AccountInfo<'info>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = reference_addresses.load()?.rmn_remote @ CcipOfframpError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,
    // remaining accounts
    // [receiver_program, external_execution_signer, receiver_account, ...user specified accounts from message data for arbitrary messaging]
    // +
    // [
    // ccip_offramp_pools_signer - derivable PDA [seed::EXTERNAL_TOKEN_POOL, pool_program], seeds::program=offramp (not in lookup table)
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
    // ccip_router_pools_signer - derivable PDA [seed::EXTERNAL_TOKEN_POOL, pool_program], seeds::program=router (present in lookup table)
    // ...additional accounts for pool config
    // ] x N tokens
}

#[derive(Accounts)]
#[instruction(source_chain_selector: u64, root: Vec<u8>)]
pub struct CloseCommitReportAccount<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        seeds = [seed::COMMIT_REPORT, source_chain_selector.to_le_bytes().as_ref(), &root],
        bump,
        constraint = valid_version(commit_report.version, MAX_COMMITREPORT_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub commit_report: Account<'info, CommitReport>,

    #[account(
        seeds = [seed::REFERENCE_ADDRESSES],
        bump,
        constraint = valid_version(reference_addresses.load()?.version, MAX_CONFIG_V) @ CcipOfframpError::InvalidVersion,
    )]
    pub reference_addresses: AccountLoader<'info, ReferenceAddresses>,

    #[account(address = native_mint::ID)]
    pub wsol_mint: Account<'info, Mint>,

    #[account(
        mut,
        associated_token::mint = wsol_mint,
        associated_token::authority = fee_billing_signer,
        associated_token::token_program = token_program,
    )]
    pub fee_token_receiver: Account<'info, TokenAccount>,

    /// CHECK: This is the signer for OnRamp billing CPIs, used here to derived OnRamp's fee accum account
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump,
        seeds::program = reference_addresses.load()?.router,
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    pub token_program: Program<'info, Token>,
}

/// It's not possible to store enums in zero_copy accounts, so we wrap the discriminant
/// in a struct to store in config.
#[derive(
    Copy, Clone, AnchorSerialize, AnchorDeserialize, PartialEq, Eq, InitSpace, Pod, Zeroable,
)]
#[repr(C)]
pub struct ConfigOcrPluginType {
    discriminant: u8,
}

impl From<OcrPluginType> for ConfigOcrPluginType {
    fn from(value: OcrPluginType) -> Self {
        Self {
            discriminant: value as u8,
        }
    }
}

impl TryFrom<ConfigOcrPluginType> for OcrPluginType {
    type Error = CcipOfframpError;

    fn try_from(value: ConfigOcrPluginType) -> std::result::Result<Self, Self::Error> {
        match value.discriminant {
            0 => Ok(Self::Commit),
            1 => Ok(Self::Execution),
            _ => Err(CcipOfframpError::InvalidPluginType),
        }
    }
}

#[derive(Copy, Clone, AnchorSerialize, AnchorDeserialize, PartialEq, Eq)]
pub enum OcrPluginType {
    Commit,
    Execution,
}
