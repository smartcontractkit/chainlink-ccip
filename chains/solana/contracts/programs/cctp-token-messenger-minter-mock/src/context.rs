use anchor_lang::prelude::*;
use anchor_spl::token::{Mint, Token, TokenAccount};
use cctp_message_transmitter_mock::{program::CctpMessageTransmitterMock, MessageTransmitter};

use crate::burn_message::BurnMessage;
use crate::program::CctpTokenMessengerMinterMock;
use crate::state::*;
use crate::TokenMessengerMinterError;

const DISCRIMINATOR_SIZE: usize = 8;

///////////////////////////////
// Parameters for instructions
///////////////////////////////

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct InitializeParams {
    pub token_controller: Pubkey,
    pub local_message_transmitter: Pubkey,
    pub message_body_version: u32,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct AddRemoteTokenMessengerParams {
    pub domain: u32,
    pub token_messenger: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct DepositForBurnParams {
    pub amount: u64,
    pub destination_domain: u32,
    pub mint_recipient: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct DepositForBurnWithCallerParams {
    pub amount: u64,
    pub destination_domain: u32,
    pub mint_recipient: Pubkey,
    pub destination_caller: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct HandleReceiveMessageParams {
    pub remote_domain: u32,
    pub sender: Pubkey,
    pub message_body: Vec<u8>,
    pub authority_bump: u8,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct SetTokenControllerParams {
    pub token_controller: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct PauseParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct UnpauseParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct UpdatePauserParams {
    pub new_pauser: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct SetMaxBurnAmountPerMessageParams {
    pub burn_limit_per_message: u64,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct AddLocalTokenParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct LinkTokenPairParams {
    pub local_token: Pubkey,
    pub remote_domain: u32,
    pub remote_token: Pubkey,
}

/////////////////////////////
// Contexts for instructions
/////////////////////////////

#[derive(Accounts)]
pub struct InitializeContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub upgrade_authority: Signer<'info>,
    #[account(mut)]
    /// CHECK: empty PDA, used to check that sendMessage was called by TokenMessenger
    #[account(
        seeds = [b"sender_authority"],
        bump
    )]
    pub authority_pda: UncheckedAccount<'info>,
    // TokenMessenger state account
    #[account(
        init,
        payer = payer,
        space = DISCRIMINATOR_SIZE + TokenMessenger::INIT_SPACE,
        seeds = [b"token_messenger"],
        bump
    )]
    pub token_messenger: Box<Account<'info, TokenMessenger>>,

    // TokenMinter state account
    #[account(
        init,
        payer = payer,
        space = DISCRIMINATOR_SIZE + TokenMinter::INIT_SPACE,
        seeds = [b"token_minter"],
        bump
    )]
    pub token_minter: Box<Account<'info, TokenMinter>>,

    /// CHECK: ProgramData account, not used for anchor tests
    #[account()]
    pub token_messenger_minter_program_data: AccountInfo<'info /*, ProgramData*/>,

    pub token_messenger_minter_program:
        Program<'info, crate::program::CctpTokenMessengerMinterMock>,

    pub system_program: Program<'info, System>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct AddRemoteTokenMessengerContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub owner: Signer<'info>,
    #[account(has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(init, payer = payer, space = 8 + std::mem::size_of::<RemoteTokenMessenger>())]
    pub remote_token_messenger: Account<'info, RemoteTokenMessenger>,
    pub system_program: Program<'info, System>,
}

// Instruction accounts
#[event_cpi]
#[derive(Accounts)]
#[instruction(params: DepositForBurnParams)]
pub struct DepositForBurnContext<'info> {
    #[account()]
    pub owner: Signer<'info>,

    #[account(mut)]
    pub event_rent_payer: Signer<'info>,

    /// CHECK: empty PDA, used to check that sendMessage was called by TokenMessenger
    #[account(
        seeds = [b"sender_authority"],
        bump = token_messenger.authority_bump,
    )]
    pub sender_authority_pda: UncheckedAccount<'info>,

    #[account(
        mut,
        constraint = burn_token_account.mint == burn_token_mint.key(),
        has_one = owner
    )]
    pub burn_token_account: Box<Account<'info, TokenAccount>>,

    #[account(mut)]
    pub message_transmitter: Box<Account<'info, MessageTransmitter>>,

    #[account()]
    pub token_messenger: Box<Account<'info, TokenMessenger>>,

    #[account(
        constraint = params.destination_domain == remote_token_messenger.domain @ TokenMessengerMinterError::InvalidDestinationDomain
    )]
    pub remote_token_messenger: Box<Account<'info, RemoteTokenMessenger>>,

    #[account()]
    pub token_minter: Box<Account<'info, TokenMinter>>,

    #[account(
        mut,
        seeds = [
            b"local_token",
            burn_token_mint.key().as_ref(),
        ],
        bump = local_token.bump,
    )]
    pub local_token: Box<Account<'info, LocalToken>>,

    #[account(mut)]
    pub burn_token_mint: Box<Account<'info, Mint>>,

    /// CHECK: Account to store MessageSent event data in. Any non-PDA uninitialized address.
    #[account(mut)]
    pub message_sent_event_data: Signer<'info>,

    pub message_transmitter_program: Program<'info, CctpMessageTransmitterMock>,

    pub token_messenger_minter_program: Program<'info, CctpTokenMessengerMinterMock>,

    pub token_program: Program<'info, Token>,

    pub system_program: Program<'info, System>,
}

#[event_cpi]
#[derive(Accounts)]
#[instruction(params: HandleReceiveMessageParams)]
pub struct HandleReceiveMessageContext<'info> {
    // authority_pda is a Signer to ensure that this instruction
    // can only be called by Message Transmitter
    #[account(
        seeds = [b"message_transmitter_authority", crate::ID.as_ref()],
        bump = params.authority_bump,
        seeds::program = cctp_message_transmitter_mock::ID
    )]
    pub authority_pda: Signer<'info>,

    #[account()]
    pub token_messenger: Box<Account<'info, TokenMessenger>>,

    #[account(
        constraint = params.remote_domain == remote_token_messenger.domain @ TokenMessengerMinterError::InvalidDestinationDomain
    )]
    pub remote_token_messenger: Box<Account<'info, RemoteTokenMessenger>>,

    #[account()]
    pub token_minter: Box<Account<'info, TokenMinter>>,

    #[account(
        mut,
        seeds = [
            b"local_token",
            local_token.mint.key().as_ref()
        ],
        bump = local_token.bump
    )]
    pub local_token: Box<Account<'info, LocalToken>>,

    #[account(
        constraint = token_pair.local_token == local_token.key() @ TokenMessengerMinterError::InvalidTokenPair,
        seeds = [
            b"token_pair",
            params.remote_domain.to_string().as_bytes(),
            BurnMessage::new(token_messenger.message_body_version, &params.message_body)?.burn_token()?.as_ref()
        ],
        bump = token_pair.bump,
    )]
    pub token_pair: Box<Account<'info, TokenPair>>,

    // Recipient's token account
    #[account(
        mut,
        constraint = recipient_token_account.mint == local_token.mint
    )]
    pub recipient_token_account: Box<Account<'info, TokenAccount>>,

    // Custody token account (could be changed to token mint)
    #[account(
        mut,
        constraint = custody_token_account.mint == local_token.mint,
        seeds = [
            b"custody",
            local_token.mint.as_ref()
        ],
        bump = local_token.custody_bump
    )]
    pub custody_token_account: Box<Account<'info, TokenAccount>>,

    // Token program to be used by the Receiver
    pub token_program: Program<'info, Token>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct SetTokenControllerContext<'info> {
    pub owner: Signer<'info>,
    #[account(has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(mut)]
    pub token_minter: Account<'info, TokenMinter>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct PauseContext<'info> {
    pub pauser: Signer<'info>,
    #[account(mut, has_one = pauser @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct UnpauseContext<'info> {
    pub pauser: Signer<'info>,
    #[account(mut, has_one = pauser @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct UpdatePauserContext<'info> {
    pub owner: Signer<'info>,
    #[account(has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(mut)]
    pub token_minter: Account<'info, TokenMinter>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct SetMaxBurnAmountPerMessageContext<'info> {
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    #[account(mut)]
    pub local_token: Account<'info, LocalToken>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct AddLocalTokenContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    #[account(init, payer = payer, space = 8 + std::mem::size_of::<LocalToken>())]
    pub local_token: Account<'info, LocalToken>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub custody_token_account: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub local_token_mint: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_program: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
}

#[event_cpi]
#[derive(Accounts)]
pub struct LinkTokenPairContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    #[account(init, payer = payer, space = 8 + std::mem::size_of::<TokenPair>())]
    pub token_pair: Account<'info, TokenPair>,
    pub system_program: Program<'info, System>,
}
