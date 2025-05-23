use anchor_lang::prelude::*;

declare_id!("CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3");

#[program]
pub mod cctp_token_messenger_minter_mock {
    use super::*;

    pub fn initialize(
        ctx: Context<InitializeContext>,
        params: InitializeParams,
    ) -> Result<()> {
        msg!("Mock initialize called with token_controller: {}, local_message_transmitter: {}, message_body_version: {}", 
             params.token_controller, params.local_message_transmitter, params.message_body_version);
        
        ctx.accounts.token_messenger.set_inner(TokenMessenger {
            owner: ctx.accounts.upgrade_authority.key(),
            pending_owner: Pubkey::default(),
            local_message_transmitter: params.local_message_transmitter,
            message_body_version: params.message_body_version,
            authority_bump: 0, // Mock value since UncheckedAccount doesn't have bump
        });

        ctx.accounts.token_minter.set_inner(TokenMinter {
            token_controller: params.token_controller,
            pauser: ctx.accounts.upgrade_authority.key(),
            paused: false,
            bump: 0, // This would normally be calculated
        });
        
        emit!(OwnershipTransferred {
            previous_owner: Pubkey::default(),
            new_owner: ctx.accounts.upgrade_authority.key(),
        });
        
        Ok(())
    }

    pub fn transfer_ownership(
        ctx: Context<TransferOwnershipContext>,
        params: TransferOwnershipParams,
    ) -> Result<()> {
        msg!("Mock transfer_ownership called with new_owner: {}", params.new_owner);
        
        let token_messenger = &mut ctx.accounts.token_messenger;
        token_messenger.pending_owner = params.new_owner;
        
        emit!(OwnershipTransferStarted {
            previous_owner: token_messenger.owner,
            new_owner: params.new_owner,
        });
        
        Ok(())
    }

    pub fn accept_ownership(
        ctx: Context<AcceptOwnershipContext>,
        _params: AcceptOwnershipParams,
    ) -> Result<()> {
        msg!("Mock accept_ownership called");
        
        let token_messenger = &mut ctx.accounts.token_messenger;
        let previous_owner = token_messenger.owner;
        token_messenger.owner = ctx.accounts.pending_owner.key();
        token_messenger.pending_owner = Pubkey::default();
        
        emit!(OwnershipTransferred {
            previous_owner,
            new_owner: ctx.accounts.pending_owner.key(),
        });
        
        Ok(())
    }

    pub fn add_remote_token_messenger(
        ctx: Context<AddRemoteTokenMessengerContext>,
        params: AddRemoteTokenMessengerParams,
    ) -> Result<()> {
        msg!("Mock add_remote_token_messenger called with domain: {}, token_messenger: {}", 
             params.domain, params.token_messenger);
        
        ctx.accounts.remote_token_messenger.set_inner(RemoteTokenMessenger {
            domain: params.domain,
            token_messenger: params.token_messenger,
        });
        
        emit!(RemoteTokenMessengerAdded {
            domain: params.domain,
            token_messenger: params.token_messenger,
        });
        
        Ok(())
    }

    pub fn remove_remote_token_messenger(
        ctx: Context<RemoveRemoteTokenMessengerContext>,
        _params: RemoveRemoteTokenMessengerParams,
    ) -> Result<()> {
        msg!("Mock remove_remote_token_messenger called");
        
        let domain = ctx.accounts.remote_token_messenger.domain;
        let token_messenger = ctx.accounts.remote_token_messenger.token_messenger;
        
        emit!(RemoteTokenMessengerRemoved {
            domain,
            token_messenger,
        });
        
        Ok(())
    }

    pub fn deposit_for_burn(
        ctx: Context<DepositForBurnContext>,
        params: DepositForBurnParams,
    ) -> Result<u64> {
        msg!("Mock deposit_for_burn called with amount: {}, destination_domain: {}, mint_recipient: {}", 
             params.amount, params.destination_domain, params.mint_recipient);
        
        emit!(DepositForBurn {
            nonce: 1, // Mock nonce
            burn_token: ctx.accounts.burn_token_mint.key(),
            amount: params.amount,
            depositor: ctx.accounts.owner.key(),
            mint_recipient: params.mint_recipient,
            destination_domain: params.destination_domain,
            destination_token_messenger: Pubkey::default(), // Mock data
            destination_caller: Pubkey::default(), // Mock data
        });
        
        Ok(1) // Return mock nonce
    }

    pub fn deposit_for_burn_with_caller(
        ctx: Context<DepositForBurnContext>,
        params: DepositForBurnWithCallerParams,
    ) -> Result<u64> {
        msg!("Mock deposit_for_burn_with_caller called with amount: {}, destination_domain: {}, mint_recipient: {}, destination_caller: {}", 
             params.amount, params.destination_domain, params.mint_recipient, params.destination_caller);
        
        emit!(DepositForBurn {
            nonce: 1, // Mock nonce
            burn_token: ctx.accounts.burn_token_mint.key(),
            amount: params.amount,
            depositor: ctx.accounts.owner.key(),
            mint_recipient: params.mint_recipient,
            destination_domain: params.destination_domain,
            destination_token_messenger: Pubkey::default(), // Mock data
            destination_caller: params.destination_caller,
        });
        
        Ok(1) // Return mock nonce
    }

    pub fn replace_deposit_for_burn(
        ctx: Context<ReplaceDepositForBurnContext>,
        params: ReplaceDepositForBurnParams,
    ) -> Result<u64> {
        msg!("Mock replace_deposit_for_burn called with original_message length: {}, original_attestation length: {}", 
             params.original_message.len(), params.original_attestation.len());
        
        emit!(DepositForBurn {
            nonce: 1, // Mock nonce
            burn_token: Pubkey::default(), // Not available in this context
            amount: 0, // Not available in this context
            depositor: ctx.accounts.owner.key(),
            mint_recipient: params.new_mint_recipient,
            destination_domain: 0, // Not available in this context
            destination_token_messenger: Pubkey::default(),
            destination_caller: params.new_destination_caller,
        });
        
        Ok(1) // Return mock nonce
    }

    pub fn handle_receive_message(
        ctx: Context<HandleReceiveMessageContext>,
        params: HandleReceiveMessageParams,
    ) -> Result<()> {
        msg!("Mock handle_receive_message called with remote_domain: {}, sender: {}, message_body length: {}", 
             params.remote_domain, params.sender, params.message_body.len());
        
        emit!(MintAndWithdraw {
            mint_recipient: ctx.accounts.recipient_token_account.key(),
            amount: 100, // Mock amount
            mint_token: ctx.accounts.custody_token_account.key(),
        });
        
        Ok(())
    }

    pub fn set_token_controller(
        ctx: Context<SetTokenControllerContext>,
        params: SetTokenControllerParams,
    ) -> Result<()> {
        msg!("Mock set_token_controller called with token_controller: {}", params.token_controller);
        
        ctx.accounts.token_minter.token_controller = params.token_controller;
        
        emit!(SetTokenController {
            token_controller: params.token_controller,
        });
        
        Ok(())
    }

    pub fn pause(
        ctx: Context<PauseContext>,
        _params: PauseParams,
    ) -> Result<()> {
        msg!("Mock pause called");
        
        ctx.accounts.token_minter.paused = true;
        
        emit!(Pause {});
        
        Ok(())
    }

    pub fn unpause(
        ctx: Context<UnpauseContext>,
        _params: UnpauseParams,
    ) -> Result<()> {
        msg!("Mock unpause called");
        
        ctx.accounts.token_minter.paused = false;
        
        emit!(Unpause {});
        
        Ok(())
    }

    pub fn update_pauser(
        ctx: Context<UpdatePauserContext>,
        params: UpdatePauserParams,
    ) -> Result<()> {
        msg!("Mock update_pauser called with new_pauser: {}", params.new_pauser);
        
        ctx.accounts.token_minter.pauser = params.new_pauser;
        
        emit!(PauserChanged {
            new_address: params.new_pauser,
        });
        
        Ok(())
    }

    pub fn set_max_burn_amount_per_message(
        ctx: Context<SetMaxBurnAmountPerMessageContext>,
        params: SetMaxBurnAmountPerMessageParams,
    ) -> Result<()> {
        msg!("Mock set_max_burn_amount_per_message called with burn_limit_per_message: {}", 
             params.burn_limit_per_message);
        
        let local_token = &mut ctx.accounts.local_token;
        local_token.burn_limit_per_message = params.burn_limit_per_message;
        
        emit!(SetBurnLimitPerMessage {
            token: ctx.accounts.local_token.key(),
            burn_limit_per_message: params.burn_limit_per_message,
        });
        
        Ok(())
    }

    pub fn add_local_token(
        ctx: Context<AddLocalTokenContext>,
        _params: AddLocalTokenParams,
    ) -> Result<()> {
        msg!("Mock add_local_token called");
        
        ctx.accounts.local_token.set_inner(LocalToken {
            custody: ctx.accounts.custody_token_account.key(),
            mint: ctx.accounts.local_token_mint.key(),
            burn_limit_per_message: u64::MAX, // Default to max
            messages_sent: 0,
            messages_received: 0,
            amount_sent: 0,
            amount_received: 0,
            bump: 0, // This would normally be calculated
            custody_bump: 0, // This would normally be calculated
        });
        
        emit!(LocalTokenAdded {
            custody: ctx.accounts.custody_token_account.key(),
            mint: ctx.accounts.local_token_mint.key(),
        });
        
        Ok(())
    }

    pub fn remove_local_token(
        ctx: Context<RemoveLocalTokenContext>,
        _params: RemoveLocalTokenParams,
    ) -> Result<()> {
        msg!("Mock remove_local_token called");
        
        emit!(LocalTokenRemoved {
            custody: ctx.accounts.custody_token_account.key(),
            mint: ctx.accounts.local_token.mint,
        });
        
        Ok(())
    }

    pub fn link_token_pair(
        ctx: Context<LinkTokenPairContext>,
        params: LinkTokenPairParams,
    ) -> Result<()> {
        msg!("Mock link_token_pair called with local_token: {}, remote_domain: {}, remote_token: {}", 
             params.local_token, params.remote_domain, params.remote_token);
        
        ctx.accounts.token_pair.set_inner(TokenPair {
            remote_domain: params.remote_domain,
            remote_token: params.remote_token,
            local_token: params.local_token,
            bump: 0, // This would normally be calculated
        });
        
        emit!(TokenPairLinked {
            local_token: params.local_token,
            remote_domain: params.remote_domain,
            remote_token: params.remote_token,
        });
        
        Ok(())
    }

    pub fn unlink_token_pair(
        ctx: Context<UnlinkTokenPairContext>,
        _params: UninkTokenPairParams,
    ) -> Result<()> {
        msg!("Mock unlink_token_pair called");
        
        emit!(TokenPairUnlinked {
            local_token: ctx.accounts.token_pair.local_token,
            remote_domain: ctx.accounts.token_pair.remote_domain,
            remote_token: ctx.accounts.token_pair.remote_token,
        });
        
        Ok(())
    }

    pub fn burn_token_custody(
        ctx: Context<BurnTokenCustodyContext>,
        params: BurnTokenCustodyParams,
    ) -> Result<()> {
        msg!("Mock burn_token_custody called with amount: {}", params.amount);
        
        emit!(TokenCustodyBurned {
            custody_token_account: ctx.accounts.custody_token_account.key(),
            amount: params.amount,
        });
        
        Ok(())
    }
}

#[account]
#[derive(Debug)]
pub struct TokenMessenger {
    pub owner: Pubkey,
    pub pending_owner: Pubkey,
    pub local_message_transmitter: Pubkey,
    pub message_body_version: u32,
    pub authority_bump: u8,
}

#[account]
#[derive(Debug)]
pub struct RemoteTokenMessenger {
    pub domain: u32,
    pub token_messenger: Pubkey,
}

#[account]
#[derive(Debug)]
pub struct TokenMinter {
    pub token_controller: Pubkey,
    pub pauser: Pubkey,
    pub paused: bool,
    pub bump: u8,
}

#[account]
#[derive(Debug)]
pub struct TokenPair {
    pub remote_domain: u32,
    pub remote_token: Pubkey,
    pub local_token: Pubkey,
    pub bump: u8,
}

#[account]
#[derive(Debug)]
pub struct LocalToken {
    pub custody: Pubkey,
    pub mint: Pubkey,
    pub burn_limit_per_message: u64,
    pub messages_sent: u64,
    pub messages_received: u64,
    pub amount_sent: u128,
    pub amount_received: u128,
    pub bump: u8,
    pub custody_bump: u8,
}

#[derive(Accounts)]
pub struct InitializeContext<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,
    pub upgrade_authority: Signer<'info>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub authority_pda: UncheckedAccount<'info>,
    #[account(init, payer = payer, space = 8 + std::mem::size_of::<TokenMessenger>())]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(init, payer = payer, space = 8 + std::mem::size_of::<TokenMinter>())]
    pub token_minter: Account<'info, TokenMinter>,
    /// CHECK: This account is used for the protocol
    pub token_messenger_minter_program_data: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_messenger_minter_program: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct TransferOwnershipContext<'info> {
    pub owner: Signer<'info>,
    #[account(mut, has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnershipContext<'info> {
    pub pending_owner: Signer<'info>,
    #[account(mut, constraint = token_messenger.pending_owner == pending_owner.key() @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

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
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct RemoveRemoteTokenMessengerContext<'info> {
    #[account(mut)]
    pub payee: Signer<'info>,
    pub owner: Signer<'info>,
    #[account(has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(mut)]
    pub remote_token_messenger: Account<'info, RemoteTokenMessenger>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct DepositForBurnContext<'info> {
    pub owner: Signer<'info>,
    #[account(mut)]
    pub event_rent_payer: Signer<'info>,
    /// CHECK: This account is used for the protocol
    pub sender_authority_pda: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub burn_token_account: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub message_transmitter: UncheckedAccount<'info>,
    pub token_messenger: Account<'info, TokenMessenger>,
    pub remote_token_messenger: Account<'info, RemoteTokenMessenger>,
    pub token_minter: Account<'info, TokenMinter>,
    #[account(mut)]
    pub local_token: Account<'info, LocalToken>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub burn_token_mint: UncheckedAccount<'info>,
    #[account(mut)]
    pub message_sent_event_data: Signer<'info>,
    /// CHECK: This account is used for the protocol
    pub message_transmitter_program: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_messenger_minter_program: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_program: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct ReplaceDepositForBurnContext<'info> {
    pub owner: Signer<'info>,
    #[account(mut)]
    pub event_rent_payer: Signer<'info>,
    /// CHECK: This account is used for the protocol
    pub sender_authority_pda: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub message_transmitter: UncheckedAccount<'info>,
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(mut)]
    pub message_sent_event_data: Signer<'info>,
    /// CHECK: This account is used for the protocol
    pub message_transmitter_program: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_messenger_minter_program: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct HandleReceiveMessageContext<'info> {
    pub authority_pda: Signer<'info>,
    pub token_messenger: Account<'info, TokenMessenger>,
    pub remote_token_messenger: Account<'info, RemoteTokenMessenger>,
    pub token_minter: Account<'info, TokenMinter>,
    #[account(mut)]
    pub local_token: Account<'info, LocalToken>,
    pub token_pair: Account<'info, TokenPair>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub recipient_token_account: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub custody_token_account: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_program: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct SetTokenControllerContext<'info> {
    pub owner: Signer<'info>,
    #[account(has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(mut)]
    pub token_minter: Account<'info, TokenMinter>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct PauseContext<'info> {
    pub pauser: Signer<'info>,
    #[account(mut, has_one = pauser @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct UnpauseContext<'info> {
    pub pauser: Signer<'info>,
    #[account(mut, has_one = pauser @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct UpdatePauserContext<'info> {
    pub owner: Signer<'info>,
    #[account(has_one = owner @ ProgramError::InvalidAccountData)]
    pub token_messenger: Account<'info, TokenMessenger>,
    #[account(mut)]
    pub token_minter: Account<'info, TokenMinter>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct SetMaxBurnAmountPerMessageContext<'info> {
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    #[account(mut)]
    pub local_token: Account<'info, LocalToken>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

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
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct RemoveLocalTokenContext<'info> {
    #[account(mut)]
    pub payee: Signer<'info>,
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    #[account(mut)]
    pub local_token: Account<'info, LocalToken>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub custody_token_account: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_program: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

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
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct UnlinkTokenPairContext<'info> {
    #[account(mut)]
    pub payee: Signer<'info>,
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    #[account(mut)]
    pub token_pair: Account<'info, TokenPair>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct BurnTokenCustodyContext<'info> {
    #[account(mut)]
    pub payee: Signer<'info>,
    pub token_controller: Signer<'info>,
    #[account(has_one = token_controller @ ProgramError::InvalidAccountData)]
    pub token_minter: Account<'info, TokenMinter>,
    pub local_token: Account<'info, LocalToken>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub custody_token_account: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: This account is used for the protocol
    pub custody_token_mint: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub token_program: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub event_authority: UncheckedAccount<'info>,
    /// CHECK: This account is used for the protocol
    pub program: UncheckedAccount<'info>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct InitializeParams {
    pub token_controller: Pubkey,
    pub local_message_transmitter: Pubkey,
    pub message_body_version: u32,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct TransferOwnershipParams {
    pub new_owner: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct AcceptOwnershipParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct AddRemoteTokenMessengerParams {
    pub domain: u32,
    pub token_messenger: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct RemoveRemoteTokenMessengerParams {}

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
pub struct ReplaceDepositForBurnParams {
    pub original_message: Vec<u8>,
    pub original_attestation: Vec<u8>,
    pub new_destination_caller: Pubkey,
    pub new_mint_recipient: Pubkey,
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
pub struct RemoveLocalTokenParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct LinkTokenPairParams {
    pub local_token: Pubkey,
    pub remote_domain: u32,
    pub remote_token: Pubkey,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct UninkTokenPairParams {}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct BurnTokenCustodyParams {
    pub amount: u64,
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
pub struct DepositForBurn {
    pub nonce: u64,
    pub burn_token: Pubkey,
    pub amount: u64,
    pub depositor: Pubkey,
    pub mint_recipient: Pubkey,
    pub destination_domain: u32,
    pub destination_token_messenger: Pubkey,
    pub destination_caller: Pubkey,
}

#[event]
pub struct MintAndWithdraw {
    pub mint_recipient: Pubkey,
    pub amount: u64,
    pub mint_token: Pubkey,
}

#[event]
pub struct RemoteTokenMessengerAdded {
    pub domain: u32,
    pub token_messenger: Pubkey,
}

#[event]
pub struct RemoteTokenMessengerRemoved {
    pub domain: u32,
    pub token_messenger: Pubkey,
}

#[event]
pub struct SetTokenController {
    pub token_controller: Pubkey,
}

#[event]
pub struct PauserChanged {
    pub new_address: Pubkey,
}

#[event]
pub struct SetBurnLimitPerMessage {
    pub token: Pubkey,
    pub burn_limit_per_message: u64,
}

#[event]
pub struct LocalTokenAdded {
    pub custody: Pubkey,
    pub mint: Pubkey,
}

#[event]
pub struct LocalTokenRemoved {
    pub custody: Pubkey,
    pub mint: Pubkey,
}

#[event]
pub struct TokenPairLinked {
    pub local_token: Pubkey,
    pub remote_domain: u32,
    pub remote_token: Pubkey,
}

#[event]
pub struct TokenPairUnlinked {
    pub local_token: Pubkey,
    pub remote_domain: u32,
    pub remote_token: Pubkey,
}

#[event]
pub struct Pause {}

#[event]
pub struct Unpause {}

#[event]
pub struct TokenCustodyBurned {
    pub custody_token_account: Pubkey,
    pub amount: u64,
}

#[error_code]
pub enum TokenMessengerMinterError {
    #[msg("Invalid authority")]
    InvalidAuthority,
    #[msg("Invalid token minter state")]
    InvalidTokenMinterState,
    #[msg("Program paused")]
    ProgramPaused,
    #[msg("Invalid token pair state")]
    InvalidTokenPairState,
    #[msg("Invalid local token state")]
    InvalidLocalTokenState,
    #[msg("Invalid pauser")]
    InvalidPauser,
    #[msg("Invalid token controller")]
    InvalidTokenController,
    #[msg("Burn amount exceeded")]
    BurnAmountExceeded,
    #[msg("Invalid amount")]
    InvalidAmount,
}
