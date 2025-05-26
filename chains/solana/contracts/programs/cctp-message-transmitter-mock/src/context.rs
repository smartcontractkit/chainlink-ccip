use anchor_lang::prelude::*;

use crate::message::Message;
use crate::program::CctpMessageTransmitterMock;
use crate::state::{MessageSent, MessageTransmitter, UsedNonces};

pub const DISCRIMINATOR_SIZE: usize = 8; // Size of the discriminator for accounts

// Instruction accounts
#[event_cpi]
#[derive(Accounts)]
pub struct InitializeContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,

    #[account()]
    pub upgrade_authority: Signer<'info>,

    // MessageTransmitter state account
    #[account(
        init,
        payer = payer,
        space = DISCRIMINATOR_SIZE + MessageTransmitter::INIT_SPACE,
        seeds = [b"message_transmitter"],
        bump
    )]
    pub message_transmitter: Box<Account<'info, MessageTransmitter>>,

    /// CHECK: ProgramData account, not used for anchor tests
    #[account()]
    pub message_transmitter_program_data: AccountInfo<'info /*, ProgramData*/>,

    pub message_transmitter_program: Program<'info, CctpMessageTransmitterMock>,

    pub system_program: Program<'info, System>,
}

// Instruction accounts
#[derive(Accounts)]
#[instruction(params: SendMessageParams)]
pub struct SendMessageContext<'info> {
    #[account(mut)]
    pub event_rent_payer: Signer<'info>,

    #[account(
        seeds = [b"sender_authority"],
        bump,
        seeds::program = sender_program.key()
    )]
    pub sender_authority_pda: Signer<'info>,

    #[account(mut)]
    pub message_transmitter: Box<Account<'info, MessageTransmitter>>,

    #[account(
        init,
        payer = event_rent_payer,
        space = MessageSent::len(params.message_body.len())?,
    )]
    pub message_sent_event_data: Box<Account<'info, MessageSent>>,

    ///CHECK: Sender program address, e.g. TokenMessenger
    #[account(
        constraint = sender_program.executable
    )]
    pub sender_program: UncheckedAccount<'info>,

    pub system_program: Program<'info, System>,
}

// Instruction accounts
#[event_cpi]
#[derive(Accounts)]
#[instruction(params: ReceiveMessageParams)]
pub struct ReceiveMessageContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,

    #[account()]
    pub caller: Signer<'info>,

    /// CHECK: empty PDA, used to check that handleReceiveMessage was called by MessageTransmitter
    #[account(
        seeds = [b"message_transmitter_authority", receiver.key().as_ref()],
        bump,
    )]
    pub authority_pda: UncheckedAccount<'info>,

    #[account()]
    pub message_transmitter: Box<Account<'info, MessageTransmitter>>,

    // Used nonces state, see UsedNonces struct for more details
    #[account(
        init_if_needed,
        payer = payer,
        space = DISCRIMINATOR_SIZE + UsedNonces::INIT_SPACE,
        seeds = [
            b"used_nonces",
            Message::new(message_transmitter.version, &params.message)?.source_domain()?.to_string().as_bytes(),
            UsedNonces::used_nonces_seed_delimiter(Message::new(message_transmitter.version, &params.message)?.source_domain()?),
            UsedNonces::first_nonce(Message::new(message_transmitter.version, &params.message)?.nonce()?)?.to_string().as_bytes()
        ],
        bump
    )]
    pub used_nonces: Box<Account<'info, UsedNonces>>,

    ///CHECK: Receiver program address, e.g. TokenMessenger
    #[account(
        constraint = receiver.executable,
        constraint = receiver.key() != crate::ID
    )]
    pub receiver: UncheckedAccount<'info>,

    pub system_program: Program<'info, System>,
    // remaining accounts: additional accounts to be passed to the receiver
}

// Instruction accounts
#[derive(Accounts)]
pub struct ReclaimEventAccountContext<'info> {
    /// rent SOL receiver, should match original rent payer
    #[account(mut)]
    pub payee: Signer<'info>,

    #[account(mut)]
    pub message_transmitter: Box<Account<'info, MessageTransmitter>>,

    #[account(
        mut,
        constraint = message_sent_event_data.rent_payer == payee.key(),
        close = payee,
    )]
    pub message_sent_event_data: Box<Account<'info, MessageSent>>,
}

#[derive(Accounts)]
pub struct GetNoncePdaContext<'info> {
    pub message_transmitter: Account<'info, MessageTransmitter>,
}

#[derive(Accounts)]
pub struct IsNonceUsedAccounts<'info> {
    pub used_nonces: Account<'info, UsedNonces>,
}

//////////////////////////
// Instruction parameters
//////////////////////////

#[derive(AnchorSerialize, AnchorDeserialize, Copy, Clone)]
pub struct InitializeParams {
    pub local_domain: u32,
    pub attester: Pubkey,
    pub max_message_body_size: u64,
    pub version: u32,
}

// Instruction parameters
// NOTE: Do not reorder parameters fields. repr(C) is used to fix the layout of the struct
// so SendMessageWithCallerParams can be deserialized as SendMessageParams.
#[repr(C)]
#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct SendMessageParams {
    pub destination_domain: u32,
    pub recipient: Pubkey,
    pub message_body: Vec<u8>,
}

// Instruction parameters
// NOTE: Do not reorder parameters fields. repr(C) is used to fix the layout of the struct
// so SendMessageWithCallerParams can be deserialized as SendMessageParams.
#[repr(C)]
#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct SendMessageWithCallerParams {
    pub destination_domain: u32,
    pub recipient: Pubkey,
    pub message_body: Vec<u8>,
    pub destination_caller: Pubkey,
}

// Instruction parameters
#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct ReceiveMessageParams {
    pub message: Vec<u8>,
    pub attestation: Vec<u8>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct ReclaimEventAccountParams {
    pub attestation: Vec<u8>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct GetNoncePDAParams {
    pub nonce: u64,
    pub source_domain: u32,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct IsNonceUsedParams {
    pub nonce: u64,
}
