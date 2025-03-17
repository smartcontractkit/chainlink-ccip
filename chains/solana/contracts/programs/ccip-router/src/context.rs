use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token::spl_token::native_mint;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

use crate::program::CcipRouter;
use crate::state::{Config, Nonce};
use crate::{CcipRouterError, DestChain, DestChainConfig, ExternalExecutionConfig, SVM2AnyMessage};

/// Static space allocated to any account: must always be added to space calculations.
pub const ANCHOR_DISCRIMINATOR: usize = 8;

/// Maximum acceptable config version accepted by this module: any accounts with higher
/// version numbers than this will be rejected.
pub const MAX_CONFIG_V: u8 = 1;
const MAX_CHAINSTATE_V: u8 = 1;
const MAX_NONCE_V: u8 = 1;

pub const fn valid_version(v: u8, max_version: u8) -> bool {
    !uninitialized(v) && v <= max_version
}

pub const fn uninitialized(v: u8) -> bool {
    v == 0
}

/// Fixed seeds for PDA derivation: different context must use different seeds.
pub mod seed {
    pub const DEST_CHAIN_STATE: &[u8] = b"dest_chain_state";
    pub const NONCE: &[u8] = b"nonce";
    pub const CONFIG: &[u8] = b"config";
    pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";

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
pub struct WithdrawBilledFunds<'info> {
    #[account(
        owner = token_program.key() @ CcipRouterError::InvalidInputsMint,
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
            &config.fee_aggregator.key(), &fee_token_mint.key(), &token_program.key()
        ) @ CcipRouterError::InvalidInputsAtaAddress,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
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
    pub config: Account<'info, Config>,

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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(address = config.proposed_owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64, dest_chain_config: DestChainConfig)]
// TODO rename to add only dest
pub struct AddChainSelector<'info> {
    #[account(
        init,
        seeds = [seed::DEST_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + DestChain::INIT_SPACE + dest_chain_config.dynamic_space(),
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64, dest_chain_config: DestChainConfig)]
pub struct UpdateDestChainSelectorConfig<'info> {
    #[account(
        mut,
        seeds = [seed::DEST_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidVersion,
        realloc = ANCHOR_DISCRIMINATOR + DestChain::INIT_SPACE + dest_chain_config.dynamic_space(),
        realloc::payer = authority,
        // `realloc::zero = true` is only necessary in cases where an instruction is capable of reallocating
        // *down* and then *up*, during a single execution. In any other cases (such as this), it's not
        // necessary as the memory will be zero'd automatically on instruction entry.
        realloc::zero = false
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(new_chain_selector: u64)]
// Similar to `UpdateDestChainSelectorConfig` but with no realloc
pub struct UpdateDestChainSelectorConfigNoRealloc<'info> {
    #[account(
        mut,
        seeds = [seed::DEST_CHAIN_STATE, new_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidVersion,
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct UpdateConfigCCIPRouter<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(source_chain_selector: u64, offramp: Pubkey)]
pub struct AddOfframp<'info> {
    #[account(
        init,
        seeds = [seed::ALLOWED_OFFRAMP, source_chain_selector.to_le_bytes().as_ref(), offramp.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + AllowedOfframp::INIT_SPACE,
    )]
    pub allowed_offramp: Account<'info, AllowedOfframp>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(source_chain_selector: u64, offramp: Pubkey)]
pub struct RemoveOfframp<'info> {
    #[account(
        mut,
        seeds = [seed::ALLOWED_OFFRAMP, source_chain_selector.to_le_bytes().as_ref(), offramp.as_ref()],
        bump,
        close = authority,
    )]
    pub allowed_offramp: Account<'info, AllowedOfframp>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[account]
#[derive(Copy, Debug, InitSpace)]
pub struct AllowedOfframp {}

#[derive(Accounts)]
#[instruction(destination_chain_selector: u64, message: SVM2AnyMessage)]
pub struct CcipSend<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [seed::DEST_CHAIN_STATE, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain_state.version, MAX_CHAINSTATE_V) @ CcipRouterError::InvalidVersion,
    )]
    pub dest_chain_state: Account<'info, DestChain>,

    #[account(
        init_if_needed,
        seeds = [seed::NONCE, destination_chain_selector.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Nonce::INIT_SPACE,
        constraint = uninitialized(nonce.version) || valid_version(nonce.version, MAX_NONCE_V) @ CcipRouterError::InvalidVersion,
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
        owner = fee_token_program.key() @ CcipRouterError::InvalidInputsMint,
        constraint = (message.fee_token == Pubkey::default() && fee_token_mint.key() == native_mint::ID)
            || message.fee_token.key() == fee_token_mint.key() @ CcipRouterError::InvalidInputsMint,
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
        address = config.fee_quoter @ CcipRouterError::InvalidVersion,
    )]
    pub fee_quoter: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::CONFIG],
        bump,
        seeds::program = config.fee_quoter,
    )]
    pub fee_quoter_config: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::DEST_CHAIN, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        seeds::program = config.fee_quoter,
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
        seeds::program = config.fee_quoter,
    )]
    pub fee_quoter_billing_token_config: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the Fee Quoter program
    #[account(
        seeds = [fee_quoter::context::seed::FEE_BILLING_TOKEN_CONFIG,
            config.link_token_mint.key().as_ref(),
        ],
        bump,
        seeds::program = config.fee_quoter,
    )]
    pub fee_quoter_link_token_config: UncheckedAccount<'info>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = config.rmn_remote @ CcipRouterError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [rmn_remote::context::seed::CURSES],
        bump,
        seeds::program = config.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [rmn_remote::context::seed::CONFIG],
        bump,
        seeds::program = config.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,

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
#[instruction(token_receiver: Pubkey, chain_selector: u64, router: Pubkey, fee_quoter: Pubkey)]
pub struct TokenAccountsValidationContext<'info> {
    /// CHECK: User Token Account
    #[account(
        constraint = user_token_account.key() == get_associated_token_address_with_program_id(
            &token_receiver.key(),
            &mint.key(),
            &token_program.key()
        ) @ CcipRouterError::InvalidInputsTokenAccounts
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
        owner = router.key() @ CcipRouterError::InvalidInputsTokenAdminRegistryAccounts,
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
        owner = pool_program.key() @ CcipRouterError::InvalidInputsPoolAccounts
    )]
    pub pool_config: AccountInfo<'info>,

    /// CHECK: Pool token account
    #[account(
        address = get_associated_token_address_with_program_id(
            &pool_signer.key(),
            &mint.key(),
            &token_program.key()
        ) @ CcipRouterError::InvalidInputsTokenAccounts
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
    #[account(owner = token_program.key() @ CcipRouterError::InvalidInputsTokenAccounts)]
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
