use anchor_lang::prelude::*;

declare_id!("CCTPmbSD7gX1bxKPAmg77w8oFzNFpaQiQUWD43TKaecd");

#[program]
pub mod cctp_message_transmitter_mock {
    use super::*;

    pub fn initialize(
        ctx: Context<Initialize>,
        params: InitializeParams,
    ) -> Result<()> {
        msg!("Mock initialize called with local_domain: {}", params.local_domain);
        
        ctx.accounts.message_transmitter.set_inner(MessageTransmitter {
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

    pub fn transfer_ownership(
        ctx: Context<OwnershipAccounts>,
        params: TransferOwnershipParams,
    ) -> Result<()> {
        msg!("Mock transfer_ownership called with new_owner: {}", params.new_owner);
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        message_transmitter.pending_owner = params.new_owner;
        
        emit!(OwnershipTransferStarted {
            previous_owner: message_transmitter.owner,
            new_owner: params.new_owner,
        });
        
        Ok(())
    }

    pub fn accept_ownership(
        ctx: Context<AcceptOwnershipAccounts>,
        _params: AcceptOwnershipParams,
    ) -> Result<()> {
        msg!("Mock accept_ownership called");
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let previous_owner = message_transmitter.owner;
        message_transmitter.owner = ctx.accounts.pending_owner.key();
        message_transmitter.pending_owner = Pubkey::default();
        
        emit!(OwnershipTransferred {
            previous_owner,
            new_owner: ctx.accounts.pending_owner.key(),
        });
        
        Ok(())
    }

    pub fn update_pauser(
        ctx: Context<OwnershipAccounts>,
        params: UpdatePauserParams,
    ) -> Result<()> {
        msg!("Mock update_pauser called with new_pauser: {}", params.new_pauser);
        
        ctx.accounts.message_transmitter.pauser = params.new_pauser;
        
        emit!(PauserChanged {
            new_address: params.new_pauser,
        });
        
        Ok(())
    }

    pub fn update_attester_manager(
        ctx: Context<OwnershipAccounts>,
        params: UpdateAttesterManagerParams,
    ) -> Result<()> {
        msg!("Mock update_attester_manager called with new_attester_manager: {}", params.new_attester_manager);
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let previous_attester_manager = message_transmitter.attester_manager;
        message_transmitter.attester_manager = params.new_attester_manager;
        
        emit!(AttesterManagerUpdated {
            previous_attester_manager,
            new_attester_manager: params.new_attester_manager,
        });
        
        Ok(())
    }

    pub fn pause(
        ctx: Context<PauserAccounts>,
        _params: PauseParams,
    ) -> Result<()> {
        msg!("Mock pause called");
        
        ctx.accounts.message_transmitter.paused = true;
        
        emit!(Pause {});
        
        Ok(())
    }

    pub fn unpause(
        ctx: Context<PauserAccounts>,
        _params: UnpauseParams,
    ) -> Result<()> {
        msg!("Mock unpause called");
        
        ctx.accounts.message_transmitter.paused = false;
        
        emit!(Unpause {});
        
        Ok(())
    }

    pub fn set_max_message_body_size(
        ctx: Context<OwnershipAccounts>,
        params: SetMaxMessageBodySizeParams,
    ) -> Result<()> {
        msg!("Mock set_max_message_body_size called with new_max_message_body_size: {}", params.new_max_message_body_size);
        
        ctx.accounts.message_transmitter.max_message_body_size = params.new_max_message_body_size;
        
        emit!(MaxMessageBodySizeUpdated {
            new_max_message_body_size: params.new_max_message_body_size,
        });
        
        Ok(())
    }

    pub fn enable_attester(
        ctx: Context<AttesterManagerAccounts>,
        params: EnableAttesterParams,
    ) -> Result<()> {
        msg!("Mock enable_attester called with new_attester: {}", params.new_attester);
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        if !message_transmitter.enabled_attesters.contains(&params.new_attester) {
            message_transmitter.enabled_attesters.push(params.new_attester);
        }
        
        emit!(AttesterEnabled {
            attester: params.new_attester,
        });
        
        Ok(())
    }

    pub fn disable_attester(
        ctx: Context<AttesterManagerAccounts>,
        params: DisableAttesterParams,
    ) -> Result<()> {
        msg!("Mock disable_attester called with attester: {}", params.attester);
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        message_transmitter.enabled_attesters.retain(|&x| x != params.attester);
        
        emit!(AttesterDisabled {
            attester: params.attester,
        });
        
        Ok(())
    }

    pub fn set_signature_threshold(
        ctx: Context<AttesterManagerAccounts>,
        params: SetSignatureThresholdParams,
    ) -> Result<()> {
        msg!("Mock set_signature_threshold called with new_signature_threshold: {}", params.new_signature_threshold);
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let old_signature_threshold = message_transmitter.signature_threshold;
        message_transmitter.signature_threshold = params.new_signature_threshold;
        
        emit!(SignatureThresholdUpdated {
            old_signature_threshold,
            new_signature_threshold: params.new_signature_threshold,
        });
        
        Ok(())
    }

    pub fn send_message(
        ctx: Context<SendMessageAccounts>,
        _params: SendMessageParams,
    ) -> Result<u64> {
        msg!("Mock send_message called");
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let nonce = message_transmitter.next_available_nonce;
        message_transmitter.next_available_nonce += 1;
        
        Ok(nonce)
    }

    pub fn send_message_with_caller(
        ctx: Context<SendMessageAccounts>,
        _params: SendMessageWithCallerParams,
    ) -> Result<u64> {
        msg!("Mock send_message_with_caller called");
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let nonce = message_transmitter.next_available_nonce;
        message_transmitter.next_available_nonce += 1;
        
        Ok(nonce)
    }

    pub fn replace_message(
        ctx: Context<SendMessageAccounts>,
        _params: ReplaceMessageParams,
    ) -> Result<u64> {
        msg!("Mock replace_message called");
        
        let message_transmitter = &mut ctx.accounts.message_transmitter;
        let nonce = message_transmitter.next_available_nonce;
        message_transmitter.next_available_nonce += 1;
        
        Ok(nonce)
    }

    pub fn receive_message(
        ctx: Context<ReceiveMessageAccounts>,
        params: ReceiveMessageParams,
    ) -> Result<()> {
        msg!("Mock receive_message called with message length: {} and attestation length: {}", 
             params.message.len(), params.attestation.len());
        
        if ctx.accounts.used_nonces.remote_domain == 0 {
            ctx.accounts.used_nonces.remote_domain = 1; // Mock remote domain
            ctx.accounts.used_nonces.first_nonce = 1;   // Start from nonce 1
        }
        
        emit!(MessageReceived {
            caller: ctx.accounts.caller.key(),
            source_domain: 1, // Mock source domain
            nonce: 1,        // Mock nonce
            sender: Pubkey::default(), // Mock sender
            message_body: params.message, // Use the input message
        });
        
        Ok(())
    }

    pub fn reclaim_event_account(
        _ctx: Context<ReclaimEventAccountAccounts>,
        _params: ReclaimEventAccountParams,
    ) -> Result<()> {
        msg!("Mock reclaim_event_account called");
        
        Ok(())
    }

    pub fn get_nonce_pda(
        _ctx: Context<GetNoncePdaAccounts>,
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

#[account]
#[derive(Debug, InitSpace)]
pub struct MessageTransmitter {
    pub owner: Pubkey,
    pub pending_owner: Pubkey,
    pub attester_manager: Pubkey,
    pub pauser: Pubkey,
    pub paused: bool,
    pub local_domain: u32,
    pub version: u32,
    pub signature_threshold: u32,
    #[max_len(1)]
    pub enabled_attesters: Vec<Pubkey>,
    pub max_message_body_size: u64,
    pub next_available_nonce: u64,
}

#[account]
#[derive(Debug, InitSpace)]
pub struct UsedNonces {
    pub remote_domain: u32,
    pub first_nonce: u64,
    pub used_nonces: [u64; 100], // Simplified from the original implementation
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub upgrade_authority: Signer<'info>,
    #[account(
        init,
        payer = payer,
        space = 8 + MessageTransmitter::INIT_SPACE
    )]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    /// CHECK: This account is the program data account for the message transmitter program
    pub message_transmitter_program_data: UncheckedAccount<'info>,
    /// CHECK: This account is the message transmitter program
    pub message_transmitter_program: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This account is used for event emission
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This is the program ID
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct OwnershipAccounts<'info> {
    pub owner: Signer<'info>,
    #[account(mut, has_one = owner @ ProgramError::InvalidAccountData)]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    /// CHECK: This account is used for event emission
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This is the program ID
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnershipAccounts<'info> {
    pub pending_owner: Signer<'info>,
    #[account(mut, constraint = message_transmitter.pending_owner == pending_owner.key() @ ProgramError::InvalidAccountData)]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    /// CHECK: This account is used for event emission
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This is the program ID
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct PauserAccounts<'info> {
    pub pauser: Signer<'info>,
    #[account(mut, has_one = pauser @ ProgramError::InvalidAccountData)]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    /// CHECK: This account is used for event emission
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This is the program ID
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct AttesterManagerAccounts<'info> {
    pub attester_manager: Signer<'info>,
    #[account(mut, has_one = attester_manager @ ProgramError::InvalidAccountData)]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    /// CHECK: This account is used for event emission
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This is the program ID
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct SendMessageAccounts<'info> {
    #[account(mut)]
    pub event_rent_payer: Signer<'info>,
    pub sender_authority_pda: Signer<'info>,
    #[account(mut)]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    #[account(mut)]
    pub message_sent_event_data: Signer<'info>,
    /// CHECK: This account is the sender program
    pub sender_program: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct ReceiveMessageAccounts<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub caller: Signer<'info>,
    /// CHECK: This account is the authority PDA
    pub authority_pda: UncheckedAccount<'info>,
    pub message_transmitter: Account<'info, MessageTransmitter>,
    #[account(
        init_if_needed,
        payer = payer,
        space = 8 + 8 + 4 + 8 + (100 * 8) // Account discriminator + remote_domain + first_nonce + used_nonces array
    )]
    pub used_nonces: Account<'info, UsedNonces>,
    /// CHECK: This account is the receiver program
    pub receiver: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This account is used for event emission
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This is the program ID
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct ReclaimEventAccountAccounts<'info> {
    #[account(mut)]
    pub payee: Signer<'info>,
    #[account(mut)]
    pub message_transmitter: Account<'info, MessageTransmitter>,
    /// CHECK: This account is used for event data
    #[account(mut)]
    pub message_sent_event_data: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct GetNoncePdaAccounts<'info> {
    pub message_transmitter: Account<'info, MessageTransmitter>,
}

#[derive(Accounts)]
pub struct IsNonceUsedAccounts<'info> {
    pub used_nonces: Account<'info, UsedNonces>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct InitializeParams {
    pub local_domain: u32,
    pub attester: Pubkey,
    pub max_message_body_size: u64,
    pub version: u32,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct TransferOwnershipParams {
    pub new_owner: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct AcceptOwnershipParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct UpdatePauserParams {
    pub new_pauser: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct UpdateAttesterManagerParams {
    pub new_attester_manager: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct PauseParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct UnpauseParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct SetMaxMessageBodySizeParams {
    pub new_max_message_body_size: u64,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct EnableAttesterParams {
    pub new_attester: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct DisableAttesterParams {
    pub attester: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct SetSignatureThresholdParams {
    pub new_signature_threshold: u32,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct SendMessageParams {
    pub destination_domain: u32,
    pub recipient: Pubkey,
    pub message_body: Vec<u8>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct SendMessageWithCallerParams {
    pub destination_domain: u32,
    pub recipient: Pubkey,
    pub message_body: Vec<u8>,
    pub destination_caller: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct ReplaceMessageParams {
    pub original_message: Vec<u8>,
    pub original_attestation: Vec<u8>,
    pub new_message_body: Vec<u8>,
    pub new_destination_caller: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct ReceiveMessageParams {
    pub message: Vec<u8>,
    pub attestation: Vec<u8>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct ReclaimEventAccountParams {
    pub attestation: Vec<u8>,  // Not used in the mock, but needed for interface compatibility
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

#[event]
pub struct OwnershipTransferStarted {
    pub previous_owner: Pubkey,
    pub new_owner: Pubkey,
}

#[event]
pub struct OwnershipTransferred {
    pub previous_owner: Pubkey,
    pub new_owner: Pubkey,
}

#[event]
pub struct PauserChanged {
    pub new_address: Pubkey,
}

#[event]
pub struct AttesterManagerUpdated {
    pub previous_attester_manager: Pubkey,
    pub new_attester_manager: Pubkey,
}

#[event]
pub struct MessageReceived {
    pub caller: Pubkey,
    pub source_domain: u32,
    pub nonce: u64,
    pub sender: Pubkey,
    pub message_body: Vec<u8>,
}

#[event]
pub struct SignatureThresholdUpdated {
    pub old_signature_threshold: u32,
    pub new_signature_threshold: u32,
}

#[event]
pub struct AttesterEnabled {
    pub attester: Pubkey,
}

#[event]
pub struct AttesterDisabled {
    pub attester: Pubkey,
}

#[event]
pub struct MaxMessageBodySizeUpdated {
    pub new_max_message_body_size: u64,
}

#[event]
pub struct Pause {}

#[event]
pub struct Unpause {}
