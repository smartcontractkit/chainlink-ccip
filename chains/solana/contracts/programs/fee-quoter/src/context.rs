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
pub const CONFIG_SEED: &[u8] = b"config";
pub const DEST_CHAIN_SEED: &[u8] = b"dest_chain";
pub const FEE_BILLING_SIGNER_SEEDS: &[u8] = b"fee_billing_signer"; // signer for billing fee token transfer
pub const FEE_BILLING_TOKEN_CONFIG_SEED: &[u8] = b"fee_billing_token_config";
pub const TOKEN_POOL_BILLING_SEED: &[u8] = b"ccip_tokenpool_billing";

// valid_version validates that the passed in version is not 0 (uninitialized)
// and it is within the expected maximum supported version bounds
pub fn valid_version(v: u8, max_v: u8) -> bool {
    v != 0 && v <= max_v
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
        seeds = [CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
        constraint = uninitialized(config.version) @ FeeQuoterError::InvalidInputs,
    )]
    pub config: Account<'info, Config>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, FeeQuoter>,

    // initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ FeeQuoterError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    // validate signer is registered admin
    #[account(address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    // validate signer is the new admin, accepting ownership of the contract
    #[account(address = config.proposed_owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(destination_chain_selector: u64, message: SVM2AnyMessage)]
pub struct GetFee<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        seeds = [DEST_CHAIN_SEED, destination_chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub dest_chain: Account<'info, DestChain>,

    #[account(
        seeds = [FEE_BILLING_TOKEN_CONFIG_SEED,
            if message.fee_token == Pubkey::default() {
                native_mint::ID.as_ref() // pre-2022 WSOL
            } else {
                message.fee_token.as_ref()
            }
        ],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,
}

#[derive(Accounts)]
#[instruction(token_config: BillingTokenConfig)]
pub struct AddBillingTokenConfig<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        init,
        seeds = [FEE_BILLING_TOKEN_CONFIG_SEED, token_config.mint.key().as_ref()],
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
        associated_token::authority = fee_billing_signer, // use the signer account as the authority
        associated_token::token_program = token_program,
    )]
    pub fee_token_receiver: InterfaceAccount<'info, TokenAccount>,

    #[account(
        mut,
        // validate signer is registered admin
        address = config.owner @ FeeQuoterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    /// CHECK: This is the signer for the billing CPIs, used here to initialize the receiver token account
    #[account(
        seeds = [FEE_BILLING_SIGNER_SEEDS],
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
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [FEE_BILLING_TOKEN_CONFIG_SEED, token_config.mint.key().as_ref()],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    // validate signer is registered admin
    #[account(address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct RemoveBillingTokenConfig<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        close = authority,
        seeds = [FEE_BILLING_TOKEN_CONFIG_SEED, fee_token_mint.key().as_ref()],
        bump,
    )]
    pub billing_token_config: Account<'info, BillingTokenConfigWrapper>,

    pub token_program: Interface<'info, TokenInterface>,

    #[account(
        owner = token_program.key() @ FeeQuoterError::InvalidInputs,
    )]
    pub fee_token_mint: InterfaceAccount<'info, Mint>,

    #[account(
        mut,
        associated_token::mint = fee_token_mint,
        associated_token::authority = fee_billing_signer, // use the signer account as the authority
        associated_token::token_program = token_program,
        constraint = fee_token_receiver.amount == 0 @ FeeQuoterError::InvalidInputs, // ensure the account is empty // TODO improve error
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
        // validate signer is registered admin
        address = config.owner @ FeeQuoterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(dest_chain_selector: u64)]
pub struct AddDestChain<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        init,
        seeds = [DEST_CHAIN_SEED, dest_chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + DestChain::INIT_SPACE,
    )]
    pub dest_chain: Account<'info, DestChain>,

    #[account(
        mut,
        // validate signer is registered admin
        address = config.owner @ FeeQuoterError::Unauthorized
    )]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(chain_selector: u64)]
pub struct UpdateDestChainConfig<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        mut,
        seeds = [DEST_CHAIN_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        constraint = valid_version(dest_chain.version, MAX_CHAINSTATE_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub dest_chain: Account<'info, DestChain>,

    // validate signer is registered admin
    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(chain_selector: u64, mint: Pubkey)]
pub struct SetTokenBillingConfig<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    #[account(
        init_if_needed,
        seeds = [TOKEN_POOL_BILLING_SEED, chain_selector.to_le_bytes().as_ref(), mint.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + PerChainPerTokenConfig::INIT_SPACE,
        constraint = uninitialized(per_chain_per_token_config.version) || valid_version(per_chain_per_token_config.version, MAX_TOKEN_AND_CHAIN_CONFIG_V) @ FeeQuoterError::InvalidInputs,
    )]
    pub per_chain_per_token_config: Account<'info, PerChainPerTokenConfig>,

    // validate signer is registered admin
    #[account(mut, address = config.owner @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdatePrices<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ FeeQuoterError::InvalidInputs, // validate state version
    )]
    pub config: Account<'info, Config>,

    // validate signer is registered offramp
    #[account(address = config.offramp @ FeeQuoterError::Unauthorized)]
    pub authority: Signer<'info>,
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
