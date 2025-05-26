use anchor_lang::prelude::*;

mod context;
use context::*;

mod event;
use event::*;

mod message;

pub mod state;
use state::*;

declare_id!("CCTPmbSD7gX1bxKPAmg77w8oFzNFpaQiQUWD43TKaecd");

#[program]
pub mod cctp_message_transmitter_mock {
    use super::*;

    pub fn initialize(ctx: Context<InitializeContext>, params: InitializeParams) -> Result<()> {
        msg!(
            "Mock initialize called with local_domain: {}",
            params.local_domain
        );

        ctx.accounts
            .message_transmitter
            .set_inner(MessageTransmitter {
                owner: ctx.accounts.upgrade_authority.key(),
                pending_owner: Pubkey::default(),
                attester_manager: ctx.accounts.upgrade_authority.key(),
                pauser: ctx.accounts.upgrade_authority.key(),
                paused: false,
                local_domain: params.local_domain,
                version: params.version,
                signature_threshold: 1,
                enabled_attesters: vec![params.attester],
                max_message_body_size: params.max_message_body_size,
                next_available_nonce: 1,
            });

        emit!(OwnershipTransferred {
            previous_owner: Pubkey::default(),
            new_owner: ctx.accounts.upgrade_authority.key(),
        });

        Ok(())
    }

    pub fn send_message(
        ctx: Context<SendMessageContext>,
        _params: SendMessageParams,
    ) -> Result<u64> {
        msg!("Mock send_message called");

        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let nonce = message_transmitter.next_available_nonce;
        message_transmitter.next_available_nonce += 1;

        Ok(nonce)
    }

    pub fn send_message_with_caller(
        ctx: Context<SendMessageContext>,
        _params: SendMessageWithCallerParams,
    ) -> Result<u64> {
        msg!("Mock send_message_with_caller called");

        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let nonce = message_transmitter.next_available_nonce;
        message_transmitter.next_available_nonce += 1;

        Ok(nonce)
    }

    pub fn receive_message(
        ctx: Context<ReceiveMessageContext>,
        params: ReceiveMessageParams,
    ) -> Result<()> {
        msg!(
            "Mock receive_message called with message length: {} and attestation length: {}",
            params.message.len(),
            params.attestation.len()
        );

        if ctx.accounts.used_nonces.remote_domain == 0 {
            ctx.accounts.used_nonces.remote_domain = 1; // Mock remote domain
            ctx.accounts.used_nonces.first_nonce = 1; // Start from nonce 1
        }

        emit!(MessageReceived {
            caller: ctx.accounts.caller.key(),
            source_domain: 1,             // Mock source domain
            nonce: 1,                     // Mock nonce
            sender: Pubkey::default(),    // Mock sender
            message_body: params.message, // Use the input message
        });

        Ok(())
    }

    pub fn reclaim_event_account(
        _ctx: Context<ReclaimEventAccountContext>,
        _params: ReclaimEventAccountParams,
    ) -> Result<()> {
        msg!("Mock reclaim_event_account called");

        Ok(())
    }

    pub fn get_nonce_pda(
        _ctx: Context<GetNoncePdaContext>,
        _params: GetNoncePDAParams,
    ) -> Result<Pubkey> {
        msg!("Mock get_nonce_pda called");

        Ok(Pubkey::default())
    }

    pub fn is_nonce_used(
        _ctx: Context<IsNonceUsedAccounts>,
        params: IsNonceUsedParams,
    ) -> Result<bool> {
        msg!("Mock is_nonce_used called for nonce: {}", params.nonce);

        Ok(false)
    }
}

#[error_code]
pub enum CctpMessageTransmitterMockError {
    #[msg("Invalid authority")]
    InvalidAuthority,
    #[msg("Instruction is not allowed at this time")]
    ProgramPaused,
    #[msg("Invalid message transmitter state")]
    InvalidMessageTransmitterState,
    #[msg("Invalid signature threshold")]
    InvalidSignatureThreshold,
    #[msg("Signature threshold already set")]
    SignatureThresholdAlreadySet,
    #[msg("Invalid owner")]
    InvalidOwner,
    #[msg("Invalid pauser")]
    InvalidPauser,
    #[msg("Invalid attester manager")]
    InvalidAttesterManager,
    #[msg("Invalid attester")]
    InvalidAttester,
    #[msg("Attester already enabled")]
    AttesterAlreadyEnabled,
    #[msg("Too few enabled attesters")]
    TooFewEnabledAttesters,
    #[msg("Signature threshold is too low")]
    SignatureThresholdTooLow,
    #[msg("Attester already disabled")]
    AttesterAlreadyDisabled,
    #[msg("Message body exceeds max size")]
    MessageBodyLimitExceeded,
    #[msg("Invalid destination caller")]
    InvalidDestinationCaller,
    #[msg("Invalid message recipient")]
    InvalidRecipient,
    #[msg("Sender is not permitted")]
    SenderNotPermitted,
    #[msg("Invalid source domain")]
    InvalidSourceDomain,
    #[msg("Invalid destination domain")]
    InvalidDestinationDomain,
    #[msg("Invalid message version")]
    InvalidMessageVersion,
    #[msg("Invalid used nonces account")]
    InvalidUsedNoncesAccount,
    #[msg("Invalid recipient program")]
    InvalidRecipientProgram,
    #[msg("Invalid nonce")]
    InvalidNonce,
    #[msg("Nonce already used")]
    NonceAlreadyUsed,
    #[msg("Message is too short")]
    MessageTooShort,
    #[msg("Malformed message")]
    MalformedMessage,
    #[msg("Invalid signature order or dupe")]
    InvalidSignatureOrderOrDupe,
    #[msg("Invalid attester signature")]
    InvalidAttesterSignature,
    #[msg("Invalid attestation length")]
    InvalidAttestationLength,
    #[msg("Invalid signature recovery ID")]
    InvalidSignatureRecoveryId,
    #[msg("Invalid signature S value")]
    InvalidSignatureSValue,
    #[msg("Invalid message hash")]
    InvalidMessageHash,
    #[msg("MathError")]
    MathError,
}
