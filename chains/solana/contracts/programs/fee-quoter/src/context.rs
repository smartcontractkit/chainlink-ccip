use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::spl_token::native_mint;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

use crate::messages::SVM2AnyMessage;
use crate::program::FeeQuoter;
use crate::state::{
    BillingTokenConfig, BillingTokenConfigWrapper, Config, DestChain, PerChainPerTokenConfig,
};
use crate::FeeQuoterError;

pub const ANCHOR_DISCRIMINATOR: usize = 8; // size in bytes

// Fixed seeds - different contexts must use different PDA seeds
pub mod seed {
    pub const CONFIG: &[u8] = b"config";
    pub const DEST_CHAIN: &[u8] = b"dest_chain";
    pub const FEE_BILLING_SIGNER: &[u8] = b"fee_billing_signer"; // signer for billing fee token transfer
    pub const FEE_BILLING_TOKEN_CONFIG: &[u8] = b"fee_billing_token_config";
    pub const PER_CHAIN_PER_TOKEN_CONFIG: &[u8] = b"per_chain_per_token_config";
}

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
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

    #[account(
        seeds = [seed::DEST_CHAIN, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidInputs,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
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
        owner = token_program.key() @ FeeQuoterError::InvalidInputs,
        constraint = token_config.mint == fee_token_mint.key() @ FeeQuoterError::InvalidInputs,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [seed::FEE_BILLING_TOKEN_CONFIG, token_config.mint.key().as_ref()],
        bump,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [seed::DEST_CHAIN, chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidInputs,
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
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

    #[account(
        init_if_needed,
        seeds = [seed::PER_CHAIN_PER_TOKEN_CONFIG, chain_selector.to_le_bytes().as_ref(), mint.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + PerChainPerTokenConfig::INIT_SPACE,
        constraint = uninitialized(per_chain_per_token_config.version) || valid_version(per_chain_per_token_config.version, MAX_TOKEN_AND_CHAIN_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub per_chain_per_token_config: Account<'info, PerChainPerTokenConfig>,

    // Only the registered admin can perform this action
    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdatePrices<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

    // Only the offramp can update prices
    #[account(address = config.offramp_signer @ FeeQuoterError::Unauthorized)] // TODO change
    pub authority: Signer<'info>,
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
