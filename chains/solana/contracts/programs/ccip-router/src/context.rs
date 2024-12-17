use anchor_lang::{prelude::*, Ids};
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};
use solana_program::sysvar::instructions;

use crate::ocr3base::Ocr3Report;
use crate::program::CcipRouter;
use crate::state::{ChainState, CommitReport, Config, ExternalExecutionConfig, Nonce};
use crate::{
    BillingTokenConfig, BillingTokenConfigWrapper, CcipRouterError, ExecutionReportSingleChain,
    ReportContext,
};

pub const ANCHOR_DISCRIMINATOR: usize = 8;

// track state versions
pub const MAX_CONFIG_V: u8 = 1;
pub const MAX_CHAINSTATE_V: u8 = 1;
pub const MAX_NONCE_V: u8 = 1;
pub const MAX_COMMITREPORT_V: u8 = 1;
pub const MAX_TOKEN_REGISTRY_V: u8 = 1;
pub const MAX_TOKEN_AND_CHAIN_CONFIG_V: u8 = 1;

// valid_version validates that the passed in version is not 0 (uninitialized)
// and it is within the expected maximum supported version bounds
pub fn valid_version(v: u8, max_v: u8) -> bool {
    v != 0 && v <= max_v
}
pub fn uninitialized(v: u8) -> bool {
    v == 0
}

// Fixed seeds - different contexts must use different PDA seeds
pub const CHAIN_STATE_SEED: &[u8] = b"chain_state";
pub const COMMIT_REPORT_SEED: &[u8] = b"commit_report";
pub const NONCE_SEED: &[u8] = b"nonce";
pub const CONFIG_SEED: &[u8] = b"config";
pub const EXTERNAL_EXECUTION_CONFIG_SEED: &[u8] = b"external_execution_config"; // arbitrary messaging signer
pub const EXTERNAL_TOKEN_POOL_SEED: &[u8] = b"external_token_pools_signer"; // token pool interaction signer
pub const FEE_BILLING_SIGNER_SEEDS: &[u8] = b"fee_billing_signer"; // signer for billing fee token transfer
pub const FEE_BILLING_TOKEN_CONFIG: &[u8] = b"fee_billing_token_config";

// Token
pub const TOKEN_ADMIN_REGISTRY_SEED: &[u8] = b"token_admin_registry";
pub const CCIP_TOKENPOOL_CONFIG: &[u8] = b"ccip_tokenpool_config";
pub const CCIP_TOKENPOOL_SIGNER: &[u8] = b"ccip_tokenpool_signer";
pub const TOKEN_POOL_BILLING_SEED: &[u8] = b"ccip_tokenpool_billing";
pub const TOKEN_POOL_CONFIG_SEED: &[u8] = b"ccip_tokenpool_chainconfig";

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct CommitInput {
    pub price_updates: PriceUpdates,
    pub merkle_root: MerkleRoot,
    // pub rmn_signatures: Vec<[u8; 65]>, // r = 32, s = 32, v = 1; placeholder: RMN not enabled
}

// A collection of token price and gas price updates.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct PriceUpdates {
    pub token_price_updates: Vec<TokenPriceUpdate>,
    pub gas_price_updates: Vec<GasPriceUpdate>,
}

// Token price in USD.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct TokenPriceUpdate {
    pub source_token: Pubkey, // Source token. It is the mint, but called "token" for EVM compatibility.
    pub usd_per_token: [u8; 28], // EVM uses u224, 1e18 USD per 1e18 of the smallest token denomination.
}

// Gas price for a given chain in USD, its value may contain tightly packed fields.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct GasPriceUpdate {
    pub dest_chain_selector: u64,   // Destination chain selector
    pub usd_per_unit_gas: [u8; 28], // EVM uses u224, 1e18 USD per smallest unit (e.g. wei) of destination chain gas
}

impl Ocr3Report for CommitInput {
    fn hash(&self, ctx: &ReportContext) -> [u8; 32] {
        use anchor_lang::solana_program::hash;
        let mut buffer: Vec<u8> = Vec::new();
        self.serialize(&mut buffer).unwrap();
        let report_len = self.len() as u16; // u16 > max tx size, u8 may have overflow
        hash::hashv(&[&report_len.to_le_bytes(), &buffer, &ctx.as_bytes()]).to_bytes()
    }

    fn len(&self) -> usize {
        4 + (32 + 28) * self.price_updates.token_price_updates.len() + // token_price_updates
        4 + (8 + 28) * self.price_updates.gas_price_updates.len() + // gas_price_updates
        self.merkle_root.len()
        // + 4 + 65 * self.rmn_signatures.len()
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
// Struct to hold a merkle root and an interval for a source chain
pub struct MerkleRoot {
    pub source_chain_selector: u64, // Remote source chain selector that the Merkle Root is scoped to
    pub on_ramp_address: Vec<u8>,   // Generic onramp address, to support arbitrary sources
    pub min_seq_nr: u64,            // Minimum sequence number, inclusive
    pub max_seq_nr: u64,            // Maximum sequence number, inclusive
    pub merkle_root: [u8; 32],      // Merkle root covering the interval & source chain messages
}

impl MerkleRoot {
    fn len(&self) -> usize {
        8  + // source chain selector
        4 + self.on_ramp_address.len() + // on ramp address
        8 + // min msg nr
        8 + // max msg nr
        32 // root
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Solana2AnyMessage {
    pub receiver: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
    pub fee_token: Pubkey, // pass zero address if native SOL
    pub extra_args: ExtraArgsInput,

    // solana specific parameter for mapping tokens to set of accounts
    pub token_indexes: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SolanaTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct ExtraArgsInput {
    pub gas_limit: Option<u128>,
    pub allow_out_of_order_execution: Option<bool>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
}

#[derive(Accounts)]
#[instruction(destination_chain_selector: u64, message: Solana2AnyMessage)]
pub struct GetFee<'info> {
    #[account(
        seeds = [CHAIN_STATE_SEED, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub chain_state: Account<'info, ChainState>,

    #[account(
        seeds = [FEE_BILLING_TOKEN_CONFIG, message.fee_token.as_ref()],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,
}

#[derive(Accounts)]
pub struct InitializeCCIPRouter<'info> {
    #[account(
        init,
        seeds = [CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CcipRouter>,
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipRouterError::Unauthorized)]
    // initialization only allowed by program upgrade authority
    pub program_data: Account<'info, ProgramData>,
    #[account(
        init,
        seeds = [EXTERNAL_EXECUTION_CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExternalExecutionConfig::INIT_SPACE,
    )]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>, // messaging CPI signer initialization
    #[account(
        init,
        seeds = [EXTERNAL_TOKEN_POOL_SEED],
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
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(address = config.load()?.proposed_owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
pub struct AddChainSelector<'info> {
    #[account(
        init,
        seeds = [CHAIN_STATE_SEED, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ChainState::INIT_SPACE,
        constraint = uninitialized(chain_state.version) @ CcipRouterError::InvalidInputs, // validate uninitialized
    )]
    pub chain_state: Account<'info, ChainState>,
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
pub struct UpdateChainSelectorConfig<'info> {
    #[account(
        mut,
        seeds = [CHAIN_STATE_SEED, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub chain_state: Account<'info, ChainState>,
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct UpdateConfigCCIPRouter<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
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
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
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
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(token_config: BillingTokenConfig)]
pub struct AddBillingTokenConfig<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        init,
        seeds = [FEE_BILLING_TOKEN_CONFIG, token_config.mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + BillingTokenConfigWrapper::INIT_SPACE,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    /// CHECK: This is the token program OR the token-2022 program. Given that there are 2 options, this can't have the
    /// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
    /// with a constraint enforcing that it is one of the two allowed programs.
    #[account(
        constraint = TokenInterface::ids().contains(&token_program.key()) @ CcipRouterError::InvalidInputs,
    )]
    pub token_program: UncheckedAccount<'info>,

    #[account(
        owner = token_program.key() @ CcipRouterError::InvalidInputs,
        constraint = token_config.mint == fee_token_mint.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        init,
        payer = authority,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer, // use the signer account as the authority // TODO discuss?
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
        seeds = [FEE_BILLING_SIGNER_SEEDS],
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
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        mut,
        seeds = [FEE_BILLING_TOKEN_CONFIG, token_config.mint.key().as_ref()],
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
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        mut,
        close = authority,
        seeds = [FEE_BILLING_TOKEN_CONFIG, fee_token_mint.key().as_ref()],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    /// CHECK: This is the token program OR the token-2022 program. Given that there are 2 options, this can't have the
    /// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
    /// with a constraint enforcing that it is one of the two allowed programs.
    #[account(
        constraint = TokenInterface::ids().contains(&token_program.key()) @ CcipRouterError::InvalidInputs,
    )]
    pub token_program: UncheckedAccount<'info>,

    #[account(
        owner = token_program.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        mut,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer, // use the config account as the authority // TODO discuss?
        associated_token::token_program = token_program,
        constraint = fee_token_receiver.amount == 0 @ CcipRouterError::InvalidInputs, // ensure the account is empty // TODO improve error
    )]
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>,

    /// CHECK: This is the signer for the billing CPIs, used here to close the receiver token account
    #[account(
        mut,
        seeds = [FEE_BILLING_SIGNER_SEEDS],
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
#[instruction(destination_chain_selector: u64, message: Solana2AnyMessage)]
pub struct CcipSend<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        mut,
        seeds = [CHAIN_STATE_SEED, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub chain_state: Account<'info, ChainState>,
    #[account(
        init_if_needed,
        seeds = [NONCE_SEED, destination_chain_selector.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Nonce::INIT_SPACE,
        constraint = uninitialized(nonce.version) || valid_version(nonce.version, MAX_NONCE_V) @ CcipRouterError::InvalidInputs, // if initialized (v != 0), validate state version
    )]
    pub nonce: Account<'info, Nonce>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    ///////////////////
    // billing token //
    ///////////////////
    // TODO improve all usages of CcipRouterError::InvalidInputs to be more specific
    /// CHECK: This is the token program OR the token-2022 program. Given that there are 2 options, this can't have the
    /// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
    /// with a constraint enforcing that it is one of the two allowed programs.
    #[account(
        constraint = TokenInterface::ids().contains(&fee_token_program.key()) @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_program: UncheckedAccount<'info>,

    #[account(
        owner = fee_token_program.key() @ CcipRouterError::InvalidInputs,
        constraint = message.fee_token.key() == fee_token_mint.key() @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        // `message.fee_token` would ideally be named `message.fee_mint` in Solana,
        // but using the `token` nomenclature is more compatible with EVM
        seeds = [FEE_BILLING_TOKEN_CONFIG, message.fee_token.as_ref()], // the arg would ideally be named mint, but message.fee_token was set for EVM consistency
        bump,
    )]
    pub fee_token_config: Account<'info, BillingTokenConfigWrapper>, // pass wSOL config if using native SOL

    #[account(
        mut,
        owner = fee_token_program.key() @ CcipRouterError::InvalidInputs,
        constraint = fee_token_mint.key() == fee_token_user_associated_account.mint.key() @ CcipRouterError::InvalidInputs,
        constraint = authority.key() == fee_token_user_associated_account.owner @ CcipRouterError::InvalidInputs,
    )]
    pub fee_token_user_associated_account: InterfaceAccount<'info, TokenAccount>, // pass zero address is using native SOL // TODO check this is possible or alternatives

    #[account(
        mut,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer, // use the config account as the authority // TODO discuss?
        associated_token::token_program = fee_token_program,
    )]
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>, // pass wSOL config if using native SOL

    /// CHECK: This is the signer for the billing transfer CPI. // TODO improve comment
    #[account(
        seeds = [FEE_BILLING_SIGNER_SEEDS],
        bump
    )]
    pub fee_billing_signer: UncheckedAccount<'info>,

    // CPI signers
    // optional if no tokens are being transferred
    /// CHECK: Using this to sign
    #[account(mut, seeds = [EXTERNAL_TOKEN_POOL_SEED], bump)]
    pub token_pools_signer: Account<'info, ExternalExecutionConfig>,
    // /// CHECK: Using this to sign
    // #[account(mut, seeds = [BILLING_EXECUTION_SEED], bump)]
    // pub billing_signer: Account<'info, ExternalExecutionConfig>,
    // pub fee_token_billing_signer_associated_account: Account<'info, ?>, // pass wSOL associated address if using native SOL

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
#[instruction(report_context: ReportContext, report: CommitInput)]
pub struct CommitReportContext<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        mut,
        seeds = [CHAIN_STATE_SEED, report.merkle_root.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub chain_state: Account<'info, ChainState>,
    #[account(
        init,
        seeds = [COMMIT_REPORT_SEED, report.merkle_root.source_chain_selector.to_le_bytes().as_ref(), report.merkle_root.merkle_root.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + CommitReport::INIT_SPACE,
        constraint = uninitialized(commit_report.version) @ CcipRouterError::InvalidInputs, // validate uninitialized
    )]
    pub commit_report: Account<'info, CommitReport>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This is the sysvar instructions account
    #[account(address = instructions::ID @ CcipRouterError::InvalidInputs)]
    pub sysvar_instructions: UncheckedAccount<'info>,
    // remaining accounts
    // [...billingTokenConfig accounts]
    // [...chainConfig accounts]
}

#[derive(Accounts)]
#[instruction(report: ExecutionReportSingleChain)]
pub struct ExecuteReportContext<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        seeds = [CHAIN_STATE_SEED, report.source_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub chain_state: Account<'info, ChainState>,
    #[account(
        mut,
        seeds = [COMMIT_REPORT_SEED, report.source_chain_selector.to_le_bytes().as_ref(), report.root.as_ref()],
        bump,
        constraint = valid_version(commit_report.version, MAX_COMMITREPORT_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub commit_report: Account<'info, CommitReport>,
    /// CHECK: Using this to sign
    #[account(seeds = [EXTERNAL_EXECUTION_CONFIG_SEED], bump)]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This is a sysvar account
    #[account(address = instructions::ID @ CcipRouterError::InvalidInputs)]
    pub sysvar_instructions: AccountInfo<'info>,
    // CPI signers
    /// CHECK: Using this to sign
    #[account(seeds = [EXTERNAL_TOKEN_POOL_SEED], bump)]
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
