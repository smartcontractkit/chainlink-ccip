use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};
use solana_program::address_lookup_table;

use crate::seed;
use crate::CommonCcipError;

#[derive(Accounts)]
#[instruction(token_receiver: Pubkey, chain_selector: u64, router: Pubkey, fee_quoter: Pubkey)]
pub struct TokenAccountsValidationContext<'info> {
    #[account(
        constraint = user_token_account.key() == get_associated_token_address_with_program_id(
            &token_receiver.key(),
            &mint.key(),
            &token_program.key()
        ) @ CommonCcipError::InvalidInputsTokenAccounts,
    )]
    pub user_token_account: InterfaceAccount<'info, TokenAccount>,

    /// CHECK: Per chain token billing config PDA
    // billing: configured via CCIP fee quoter
    // chain config: configured via pool
    #[account(
        seeds = [
            seed::PER_CHAIN_PER_TOKEN_CONFIG,
            chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref(),
        ],
        seeds::program = fee_quoter.key(),
        bump
    )]
    pub token_billing_config: UncheckedAccount<'info>,

    /// CHECK: Pool chain config PDA
    #[account(
        seeds = [
            seed::TOKEN_POOL_CONFIG,
            chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref(),
        ],
        seeds::program = pool_program.key(),
        owner = pool_program.key() @ CommonCcipError::InvalidInputsPoolAccounts,
        bump
    )]
    pub pool_chain_config: UncheckedAccount<'info>,

    /// CHECK: Lookup table
    #[account(owner = address_lookup_table::program::id() @ CommonCcipError::InvalidInputsLookupTableAccounts)]
    pub lookup_table: UncheckedAccount<'info>,

    /// CHECK: Token admin registry PDA
    #[account(
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        seeds::program = router.key(),
        bump,
        owner = router.key() @ CommonCcipError::InvalidInputsTokenAdminRegistryAccounts,
    )]
    pub token_admin_registry: UncheckedAccount<'info>,

    /// CHECK: Pool program
    #[account(executable)]
    pub pool_program: UncheckedAccount<'info>,

    /// CHECK: Pool config PDA
    #[account(
        seeds = [seed::CCIP_TOKENPOOL_CONFIG, mint.key().as_ref()],
        seeds::program = pool_program.key(),
        bump,
        owner = pool_program.key() @ CommonCcipError::InvalidInputsPoolAccounts
    )]
    pub pool_config: UncheckedAccount<'info>,

    #[account(
        address = get_associated_token_address_with_program_id(
            &pool_signer.key(),
            &mint.key(),
            &token_program.key()
        ) @ CommonCcipError::InvalidInputsTokenAccounts
    )]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    /// CHECK: Pool signer PDA
    #[account(
        seeds = [seed::CCIP_TOKENPOOL_SIGNER, mint.key().as_ref()],
        seeds::program = pool_program.key(),
        bump
    )]
    pub pool_signer: UncheckedAccount<'info>,

    pub token_program: Interface<'info, TokenInterface>,

    #[account(owner = token_program.key() @ CommonCcipError::InvalidInputsTokenAccounts)]
    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK: Fee token config PDA
    #[account(
        seeds = [
            seed::FEE_BILLING_TOKEN_CONFIG,
            mint.key().as_ref()
        ],
        seeds::program = fee_quoter.key(),
        bump
    )]
    pub fee_token_config: UncheckedAccount<'info>,

    /// CHECK: The signer to be used by the router program to invoke the pool program
    #[account(
        seeds = [seed::EXTERNAL_TOKEN_POOLS_SIGNER, pool_program.key().as_ref()],
        bump,
        seeds::program = router.key(),
    )]
    pub ccip_router_pool_signer: UncheckedAccount<'info>,
}
