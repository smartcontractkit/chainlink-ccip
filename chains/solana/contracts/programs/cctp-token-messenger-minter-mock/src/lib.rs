use anchor_lang::prelude::*;

mod context;
use context::*;

mod state;
use state::*;

mod event;
use event::*;

mod burn_message;

declare_id!("CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3");

#[program]
pub mod cctp_token_messenger_minter_mock {
    use super::*;

    pub fn initialize(ctx: Context<InitializeContext>, params: InitializeParams) -> Result<()> {
        msg!("Mock initialize called with token_controller: {}, local_message_transmitter: {}, message_body_version: {}",
             params.token_controller, params.local_message_transmitter, params.message_body_version);

        ctx.accounts.token_messenger.set_inner(TokenMessenger {
            owner: ctx.accounts.upgrade_authority.key(),
            pending_owner: Pubkey::default(),
            local_message_transmitter: params.local_message_transmitter,
            message_body_version: params.message_body_version,
            authority_bump: 0, // Mock value
        });

        ctx.accounts.token_minter.set_inner(TokenMinter {
            token_controller: params.token_controller,
            pauser: ctx.accounts.upgrade_authority.key(),
            paused: false,
            bump: 0, // This would normally be calculated
        });

        Ok(())
    }

    pub fn add_remote_token_messenger(
        ctx: Context<AddRemoteTokenMessengerContext>,
        params: AddRemoteTokenMessengerParams,
    ) -> Result<()> {
        msg!(
            "Mock add_remote_token_messenger called with domain: {}, token_messenger: {}",
            params.domain,
            params.token_messenger
        );

        ctx.accounts
            .remote_token_messenger
            .set_inner(RemoteTokenMessenger {
                domain: params.domain,
                token_messenger: params.token_messenger,
            });

        emit_cpi!(RemoteTokenMessengerAdded {
            domain: params.domain,
            token_messenger: params.token_messenger,
        });

        Ok(())
    }

    pub fn deposit_for_burn(
        ctx: Context<DepositForBurnContext>,
        params: DepositForBurnParams,
    ) -> Result<u64> {
        msg!("Mock deposit_for_burn called with amount: {}, destination_domain: {}, mint_recipient: {}",
             params.amount, params.destination_domain, params.mint_recipient);

        emit_cpi!(DepositForBurn {
            nonce: 1, // Mock nonce
            burn_token: ctx.accounts.burn_token_mint.key(),
            amount: params.amount,
            depositor: ctx.accounts.owner.key(),
            mint_recipient: params.mint_recipient,
            destination_domain: params.destination_domain,
            destination_token_messenger: Pubkey::default(), // Mock data
            destination_caller: Pubkey::default(),          // Mock data
        });

        Ok(1) // Return mock nonce
    }

    pub fn deposit_for_burn_with_caller(
        ctx: Context<DepositForBurnContext>,
        params: DepositForBurnWithCallerParams,
    ) -> Result<u64> {
        msg!("Mock deposit_for_burn_with_caller called with amount: {}, destination_domain: {}, mint_recipient: {}, destination_caller: {}",
             params.amount, params.destination_domain, params.mint_recipient, params.destination_caller);

        emit_cpi!(DepositForBurn {
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

    pub fn handle_receive_message(
        ctx: Context<HandleReceiveMessageContext>,
        params: HandleReceiveMessageParams,
    ) -> Result<()> {
        msg!("Mock handle_receive_message called with remote_domain: {}, sender: {}, message_body length: {}",
             params.remote_domain, params.sender, params.message_body.len());

        emit_cpi!(MintAndWithdraw {
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
        msg!(
            "Mock set_token_controller called with token_controller: {}",
            params.token_controller
        );

        ctx.accounts.token_minter.token_controller = params.token_controller;

        emit_cpi!(SetTokenController {
            token_controller: params.token_controller,
        });

        Ok(())
    }

    pub fn pause(ctx: Context<PauseContext>, _params: PauseParams) -> Result<()> {
        msg!("Mock pause called");

        ctx.accounts.token_minter.paused = true;

        emit_cpi!(Pause {});

        Ok(())
    }

    pub fn unpause(ctx: Context<UnpauseContext>, _params: UnpauseParams) -> Result<()> {
        msg!("Mock unpause called");

        ctx.accounts.token_minter.paused = false;

        emit_cpi!(Unpause {});

        Ok(())
    }

    pub fn update_pauser(
        ctx: Context<UpdatePauserContext>,
        params: UpdatePauserParams,
    ) -> Result<()> {
        msg!(
            "Mock update_pauser called with new_pauser: {}",
            params.new_pauser
        );

        ctx.accounts.token_minter.pauser = params.new_pauser;

        emit_cpi!(PauserChanged {
            new_address: params.new_pauser,
        });

        Ok(())
    }

    pub fn set_max_burn_amount_per_message(
        ctx: Context<SetMaxBurnAmountPerMessageContext>,
        params: SetMaxBurnAmountPerMessageParams,
    ) -> Result<()> {
        msg!(
            "Mock set_max_burn_amount_per_message called with burn_limit_per_message: {}",
            params.burn_limit_per_message
        );

        let local_token = &mut ctx.accounts.local_token;
        local_token.burn_limit_per_message = params.burn_limit_per_message;

        emit_cpi!(SetBurnLimitPerMessage {
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
            bump: 0,         // This would normally be calculated
            custody_bump: 0, // This would normally be calculated
        });

        emit_cpi!(LocalTokenAdded {
            custody: ctx.accounts.custody_token_account.key(),
            mint: ctx.accounts.local_token_mint.key(),
        });

        Ok(())
    }

    pub fn link_token_pair(
        ctx: Context<LinkTokenPairContext>,
        params: LinkTokenPairParams,
    ) -> Result<()> {
        msg!(
            "Mock link_token_pair called with local_token: {}, remote_domain: {}, remote_token: {}",
            params.local_token,
            params.remote_domain,
            params.remote_token
        );

        ctx.accounts.token_pair.set_inner(TokenPair {
            remote_domain: params.remote_domain,
            remote_token: params.remote_token,
            local_token: params.local_token,
            bump: 0, // This would normally be calculated
        });

        emit_cpi!(TokenPairLinked {
            local_token: params.local_token,
            remote_domain: params.remote_domain,
            remote_token: params.remote_token,
        });

        Ok(())
    }
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
    #[msg("Invalid destination domain")]
    InvalidDestinationDomain,
    #[msg("Invalid token pair")]
    InvalidTokenPair,
    #[msg("Malformed message")]
    MalformedMessage,
    #[msg("Invalid message body version")]
    InvalidMessageBodyVersion,
}
