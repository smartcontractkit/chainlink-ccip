use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
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
const MAX_CHAIN_V: u8 = 1;
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
/// the source chain, and optionally some price updates alongside it.
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
    /// Adding a chain selector implies initializing the state for a new chain,
    /// hence the need to initialize two accounts.
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
#[instruction(new_chain_selector: u64)]
pub struct UpdateSourceChain<'info> {
    #[account(
        mut,
        seeds = [seed::SOURCE_CHAIN, new_chain_selector.to_le_bytes().as_ref()],
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
        seeds = [fee_quoter::context::seed::ALLOWED_PRICE_UPDATER, fee_billing_signer.key().as_ref()],
        bump,
        seeds::program = fee_quoter.key(),
    )]
    pub fee_quoter_allowed_price_updater: UncheckedAccount<'info>,

    /// CHECK: fee quoter config account, used to invoke fee quoter with price updates
    #[account(
        seeds = [fee_quoter::context::seed::CONFIG],
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
        seeds = [rmn_remote::context::seed::CURSES],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [rmn_remote::context::seed::CONFIG],
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
        seeds = [fee_quoter::context::seed::ALLOWED_PRICE_UPDATER, fee_billing_signer.key().as_ref()],
        bump,
        seeds::program = fee_quoter.key(),
    )]
    pub fee_quoter_allowed_price_updater: UncheckedAccount<'info>,

    /// CHECK: fee quoter config account, used to invoke fee quoter with price updates
    #[account(
        seeds = [fee_quoter::context::seed::CONFIG],
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
        seeds = [rmn_remote::context::seed::CURSES],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [rmn_remote::context::seed::CONFIG],
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

    /// CHECK: Using this to sign
    #[account(seeds = [seed::EXTERNAL_EXECUTION_CONFIG], bump)]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    /// CHECK: This is a sysvar account
    #[account(address = instructions::ID @ CcipOfframpError::InvalidInputsSysvarAccount)]
    pub sysvar_instructions: AccountInfo<'info>,

    #[account(seeds = [seed::EXTERNAL_TOKEN_POOL], bump)]
    pub token_pools_signer: Account<'info, ExternalExecutionConfig>,

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
        seeds = [rmn_remote::context::seed::CURSES],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [rmn_remote::context::seed::CONFIG],
        bump,
        seeds::program = reference_addresses.load()?.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,
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

#[derive(Accounts)]
#[instruction(token_receiver: Pubkey, chain_selector: u64, router: Pubkey, fee_quoter: Pubkey)]
pub struct TokenAccountsValidationContext<'info> {
    /// CHECK: User Token Account
    #[account(
        constraint = user_token_account.key() == get_associated_token_address_with_program_id(
            &token_receiver.key(),
            &mint.key(),
            &token_program.key()
        ) @ CcipOfframpError::InvalidInputsTokenAccounts
    )]
    pub user_token_account: AccountInfo<'info>,

    // TODO: determine if this can be zero key for optional billing config?
    /// CHECK: Per chain token billing config
    // billing: configured via CCIP fee quoter
    // chain config: configured via pool
    #[account(
        seeds = [
            fee_quoter::context::seed::PER_CHAIN_PER_TOKEN_CONFIG,
            chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref(),
        ],
        seeds::program = fee_quoter.key(),
        bump
    )]
    pub token_billing_config: AccountInfo<'info>,

    /// CHECK: Pool chain config
    #[account(
        seeds = [
            seed::TOKEN_POOL_CONFIG,
            chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref(),
        ],
        seeds::program = pool_program.key(),
        bump
    )]
    pub pool_chain_config: AccountInfo<'info>,

    /// CHECK: Lookup table
    pub lookup_table: AccountInfo<'info>,

    /// CHECK: Token admin registry
    #[account(
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        seeds::program = router.key(),
        bump,
        owner = router.key() @ CcipOfframpError::InvalidInputsTokenAdminRegistryAccounts,
    )]
    pub token_admin_registry: AccountInfo<'info>,

    /// CHECK: Pool program
    pub pool_program: AccountInfo<'info>,

    // todo: PDA constraint violation will emit AccountConstraintViolation error instead of InvalidInputsPoolAccounts
    /// CHECK: Pool config
    #[account(
        seeds = [seed::CCIP_TOKENPOOL_CONFIG, mint.key().as_ref()],
        seeds::program = pool_program.key(),
        bump,
        owner = pool_program.key() @ CcipOfframpError::InvalidInputsPoolAccounts
    )]
    pub pool_config: AccountInfo<'info>,

    /// CHECK: Pool token account
    #[account(
        address = get_associated_token_address_with_program_id(
            &pool_signer.key(),
            &mint.key(),
            &token_program.key()
        ) @ CcipOfframpError::InvalidInputsTokenAccounts
    )]
    pub pool_token_account: AccountInfo<'info>,

    // todo: PDA constraint violation will emit AccountConstraintViolation error instead of InvalidInputsPoolAccounts
    /// CHECK: Pool signer
    #[account(
        seeds = [seed::CCIP_TOKENPOOL_SIGNER, mint.key().as_ref()],
        seeds::program = pool_program.key(),
        bump
    )]
    pub pool_signer: AccountInfo<'info>,

    /// CHECK: Token program
    pub token_program: AccountInfo<'info>,

    /// CHECK: Mint
    #[account(owner = token_program.key() @ CcipOfframpError::InvalidInputsTokenAccounts)]
    pub mint: AccountInfo<'info>,

    // todo: PDA constraint violation will emit AccountConstraintViolation error instead of InvalidInputsConfigAccounts
    /// CHECK: Fee token config
    #[account(
        seeds = [
            fee_quoter::context::seed::FEE_BILLING_TOKEN_CONFIG,
            mint.key().as_ref()
        ],
        seeds::program = fee_quoter.key(),
        bump
    )]
    pub fee_token_config: AccountInfo<'info>,
}
