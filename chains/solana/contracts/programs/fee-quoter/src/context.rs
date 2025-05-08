use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::spl_token::native_mint;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};
use ccip_common::seed;

use crate::messages::SVM2AnyMessage;
use crate::program::FeeQuoter;
use crate::state::{
    AllowedPriceUpdater, BillingTokenConfig, BillingTokenConfigWrapper, Config, DestChain,
    PerChainPerTokenConfig,
};
use crate::FeeQuoterError;

pub const ANCHOR_DISCRIMINATOR: usize = 8; // size in bytes

// valid_version validates that the passed in version is not 0 (uninitialized)
// and it is within the expected maximum supported version bounds
pub fn valid_version(v: u8, max_v: u8) -> bool {
    !uninitialized(v) && v <= max_v
}
pub fn uninitialized(v: u8) -> bool {
    v == 0
}

const MAX_CONFIG_V: u8 = 1;
const MAX_CHAINSTATE_V: u8 = 1;
const MAX_TOKEN_AND_CHAIN_CONFIG_V: u8 = 1;
const MAX_BILLING_TOKEN_CONFIG_V: u8 = 1;

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: Account<'info, Config>,

    pub link_token_mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, FeeQuoter>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ FeeQuoterError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct Empty<'info> {
    // This is unused, but Anchor requires that there is at least one account in the context
    pub clock: Sysvar<'info, Clock>,
}

#[derive(Accounts)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    // validate signer is registered admin
    #[account(address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct UpdateConfigLinkMint<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    pub link_token_mint: InterfaceAccount<'info, Mint>,

    // validate signer is registered admin
    #[account(address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    // validate signer is the new admin, accepting ownership of the contract
    #[account(address = config.proposed_owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(destination_chain_selector: u64, message: SVM2AnyMessage)]
pub struct GetFee<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(
        seeds = [seed::DEST_CHAIN, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub dest_chain: Account<'info, DestChain>,

    #[account(
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG,
            if message.fee_token == Pubkey::default() {
                native_mint::ID.as_ref() // pre-2022 WSOL
            } else {
                message.fee_token.as_ref()
            }
        ],
        bump,
        constraint = valid_version(billing_token_config.version, MAX_BILLING_TOKEN_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    #[account(
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG, config.link_token_mint.as_ref()],
        bump,
    )]
    pub link_token_config: Account<'info, BillingTokenConfigWrapper>,
    //
    // remaining_accounts:
    // - First all BillingTokenConfigWrapper accounts (one per token transferred)
    // - Then all PerChainPerTokenConfig accounts (one per token transferred)
}

#[derive(Accounts)]
#[instruction(token_config: BillingTokenConfig)]
pub struct AddBillingTokenConfig<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

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
        owner = token_program.key() @ FeeQuoterError::InvalidInputsMintOwner,
        constraint = token_config.mint == fee_token_mint.key() @ FeeQuoterError::InvalidInputsMint,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        init,
        payer = authority,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer,
        associated_token::token_program = token_program,
    )]
    // TODO figure out how to set the close authority to FQ
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>,

    #[account(
        mut,
        // Only the registered admin can add a new billing token config
        address = config.owner @ FeeQuoterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    /// CHECK: This is the signer for the billing CPIs, used here to initialize the receiver token account
    #[account(
        seeds = [seed::FEE_BILLING_SIGNER],
        bump,
        seeds::program = config.onramp,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG, token_config.mint.key().as_ref()],
        bump,
        constraint = valid_version(billing_token_config.version, MAX_BILLING_TOKEN_CONFIG_V) @ FeeQuoterError::InvalidVersion,
        constraint = billing_token_config.config.mint == token_config.mint @ FeeQuoterError::InvalidInputsMint,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    // Only the registered admin can update a billing token config
    #[account(address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(dest_chain_selector: u64)]
pub struct AddDestChain<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(
        init,
        seeds = [seed::DEST_CHAIN, dest_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + DestChain::INIT_SPACE,
    )]
    pub dest_chain: Account<'info, DestChain>,

    #[account(
        mut,
        // Only the registered admin can perform this action
        address = config.owner @ FeeQuoterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(chain_selector: u64)]
pub struct UpdateDestChainConfig<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [seed::DEST_CHAIN, chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub dest_chain: Account<'info, DestChain>,

    // Only the registered admin can perform this action
    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(chain_selector: u64, mint: Pubkey)]
pub struct SetTokenTransferFeeConfig<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(
        init_if_needed,
        seeds = [seed::PER_CHAIN_PER_TOKEN_CONFIG, chain_selector.to_le_bytes().as_ref(), mint.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + PerChainPerTokenConfig::INIT_SPACE,
        constraint = uninitialized(per_chain_per_token_config.version) || valid_version(per_chain_per_token_config.version, MAX_TOKEN_AND_CHAIN_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub per_chain_per_token_config: Account<'info, PerChainPerTokenConfig>,

    // Only the registered admin can perform this action
    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(price_updater: Pubkey)]
pub struct AddPriceUpdater<'info> {
    #[account(
        init,
        seeds = [seed::ALLOWED_PRICE_UPDATER,  price_updater.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + AllowedPriceUpdater::INIT_SPACE,
    )]
    pub allowed_price_updater: Account<'info, AllowedPriceUpdater>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(price_updater: Pubkey)]
pub struct RemovePriceUpdater<'info> {
    #[account(
        mut,
        seeds = [seed::ALLOWED_PRICE_UPDATER,  price_updater.as_ref()],
        bump,
        close = authority,
    )]
    pub allowed_price_updater: Account<'info, AllowedPriceUpdater>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,

    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdatePrices<'info> {
    pub authority: Signer<'info>,

    /// CHECK: This is the AllowedPriceUpdater account for the calling authority. It is used to verify that the caller
    /// was added by the owner as an allowed price updater. The constraints enforced guarantee that it is the right PDA
    /// and that it was initialized.
    #[account(
        owner = crate::ID @ FeeQuoterError::UnauthorizedPriceUpdater, // this guarantees that it was initialized
        seeds = [seed::ALLOWED_PRICE_UPDATER, authority.key().as_ref()],
        bump,
    )]
    pub allowed_price_updater: UncheckedAccount<'info>,

    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    // Remaining accounts represent:
    // - the accounts to update BillingTokenConfig for token prices
    // - the accounts to update DestChain for gas prices
    // They must be in order:
    // 1. token_accounts[]
    // 2. gas_accounts[]
    // matching the order of the price updates.
    // They must also all be writable so they can be updated.
}

// Token price in USD.
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
