use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenAccount};

declare_id!("48LGpn6tPn5SjTtK2wL9uUx48JUWZdZBv11sboy2orCc");

pub const EXTERNAL_EXECUTION_CONFIG_SEED: &[u8] = b"external_execution_config";
pub const APPROVED_SENDER_SEED: &[u8] = b"approved_ccip_sender";
pub const TOKEN_ADMIN_SEED: &[u8] = b"receiver_token_admin";
pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";
pub const STATE: &[u8] = b"state";

/// This program an example of a CCIP Receiver Program.
/// Used to test CCIP Router execute.
#[program]
pub mod example_ccip_receiver {
    use anchor_spl::token_2022::spl_token_2022::{self, instruction::transfer_checked};
    use solana_program::program::invoke_signed;

    use super::*;

    /// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .init(ctx.accounts.authority.key(), router)
    }

    /// This function is called by the CCIP Offramp to execute the CCIP message.
    /// The method name needs to be ccip_receive with Anchor encoding,
    /// if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]
    /// You can send as many accounts as you need, specifying if mutable or not.
    /// But none of them could be an init, realloc or close.
    pub fn ccip_receive(_ctx: Context<CcipReceive>, message: Any2SVMMessage) -> Result<()> {
        // ---------------------------------------
        // implement functionality here
        // ---------------------------------------

        emit!(MessageReceived {
            message_id: message.message_id
        });

        Ok(())
    }

    pub fn update_router(ctx: Context<UpdateConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .update_router(ctx.accounts.authority.key(), new_router)
    }

    // approve_sender creates a PDA to approve the specific source chain + remote address
    pub fn approve_sender(
        _ctx: Context<ApproveSender>,
        _chain_selector: u64,
        _remote_address: Vec<u8>,
    ) -> Result<()> {
        Ok(())
    }

    pub fn unapprove_sender(
        _ctx: Context<UnapproveSender>,
        _chain_selector: u64,
        _remote_address: Vec<u8>,
    ) -> Result<()> {
        Ok(())
    }

    pub fn transfer_ownership(ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .transfer_ownership(ctx.accounts.authority.key(), proposed_owner)
    }

    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts
            .state
            .accept_ownership(ctx.accounts.authority.key())
    }

    pub fn withdraw_tokens(ctx: Context<WithdrawTokens>, amount: u64, decimals: u8) -> Result<()> {
        let mut ix = transfer_checked(
            &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
            &ctx.accounts.program_token_account.key(),
            &ctx.accounts.mint.key(),
            &ctx.accounts.to_token_account.key(),
            &ctx.accounts.token_admin.key(),
            &[],
            amount,
            decimals,
        )?;
        ix.program_id = ctx.accounts.token_program.key(); // set to user specified program

        let seeds = &[TOKEN_ADMIN_SEED, &[ctx.bumps.token_admin]];
        invoke_signed(
            &ix,
            &[
                ctx.accounts.program_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.to_token_account.to_account_info(),
                ctx.accounts.token_admin.to_account_info(),
            ],
            &[&seeds[..]],
        )?;
        Ok(())
    }
}

const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [b"state"],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + BaseState::INIT_SPACE,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        init,
        seeds = [TOKEN_ADMIN_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR,
    )]
    /// CHECK: CPI signer for tokens
    pub token_admin: UncheckedAccount<'info>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(message: Any2SVMMessage)]
pub struct CcipReceive<'info> {
    // Offramp CPI signer PDA must be first.
    // It is not mutable, and thus cannot be used as payer of init/realloc of other PDAs.
    #[account(
        seeds = [EXTERNAL_EXECUTION_CONFIG_SEED, crate::ID.as_ref()],
        bump,
        seeds::program = offramp_program.key(),
    )]
    pub authority: Signer<'info>,

    /// CHECK offramp program: exists only to derive the allowed offramp PDA
    /// and the authority PDA. Must be second.
    pub offramp_program: UncheckedAccount<'info>,

    // PDA to verify that calling offramp is valid. Must be third. It is left up to the implementer to decide
    // how they want to persist the router address to verify that this is the correct account (e.g. in the top level of
    // a global config/state account for the receiver, which is what this example does, or hard-coded,
    // or stored in any other way in any other account).
    /// CHECK PDA of the router program verifying the signer is an allowed offramp.
    /// If PDA does not exist, the router doesn't allow this offramp
    #[account(
        owner = state.router @ CcipReceiverError::InvalidCaller, // this guarantees that it was initialized
        seeds = [
            ALLOWED_OFFRAMP,
            message.source_chain_selector.to_le_bytes().as_ref(),
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = state.router,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

    // -- From here on, these are receiver-specific PDAs.
    #[account(
        seeds = [
            APPROVED_SENDER_SEED,
            message.source_chain_selector.to_le_bytes().as_ref(),
            &[message.sender.len() as u8],
            &message.sender,
        ],
        bump,
    )]
    pub approved_sender: Account<'info, ApprovedSender>, // if PDA does not exist, the message sender and/or source chain are not approved
    #[account(
        seeds = [STATE],
        bump,
    )]
    pub state: Account<'info, BaseState>,
}

#[derive(Accounts, Debug)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [STATE],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.owner @ CcipReceiverError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts, Debug)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [STATE],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.proposed_owner @ CcipReceiverError::OnlyProposedOwner,
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64, remote_sender: Vec<u8>)]
pub struct ApproveSender<'info> {
    #[account(
        seeds = [STATE],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        init,
        seeds = [
            APPROVED_SENDER_SEED,
            chain_selector.to_le_bytes().as_ref(),
            &[remote_sender.len() as u8],
            &remote_sender,
        ],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ApprovedSender::INIT_SPACE,
    )]
    pub approved_sender: Account<'info, ApprovedSender>,
    #[account(
        mut,
        address = state.owner @ CcipReceiverError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64, remote_sender: Vec<u8>)]
pub struct UnapproveSender<'info> {
    #[account(
        mut,
        seeds = [STATE],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        seeds = [
            APPROVED_SENDER_SEED,
            chain_selector.to_le_bytes().as_ref(),
            &[remote_sender.len() as u8],
            &remote_sender,
        ],
        bump,
        close = authority,
    )]
    pub approved_sender: Account<'info, ApprovedSender>,
    #[account(
        mut,
        address = state.owner @ CcipReceiverError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct WithdrawTokens<'info> {
    #[account(
        mut,
        seeds = [STATE],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        token::mint = mint,
        token::authority = token_admin,
        token::token_program = token_program,
    )]
    pub program_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        token::mint = mint,
        token::token_program = token_program,
    )]
    pub to_token_account: InterfaceAccount<'info, TokenAccount>,
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(address = *mint.to_account_info().owner)]
    /// CHECK: CPI to token program
    pub token_program: AccountInfo<'info>,
    #[account(
        seeds = [TOKEN_ADMIN_SEED],
        bump,
    )]
    /// CHECK: CPI signer for tokens
    pub token_admin: UncheckedAccount<'info>,
    #[account(
        address = state.owner @ CcipReceiverError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
}

// BaseState contains the state for core safety checks that can be leveraged by the implementer
#[account]
#[derive(InitSpace, Default, Debug)]
pub struct BaseState {
    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    pub router: Pubkey,
}

impl BaseState {
    pub fn init(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_keys_eq!(self.owner, Pubkey::default());
        self.owner = owner;
        self.update_router(owner, router)
    }

    pub fn transfer_ownership(&mut self, owner: Pubkey, proposed_owner: Pubkey) -> Result<()> {
        require!(
            proposed_owner != self.owner && proposed_owner != Pubkey::default(),
            CcipReceiverError::InvalidProposedOwner
        );
        require_keys_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);
        self.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(&mut self, proposed_owner: Pubkey) -> Result<()> {
        require_keys_eq!(
            self.proposed_owner,
            proposed_owner,
            CcipReceiverError::OnlyProposedOwner
        );
        // NOTE: take() resets proposed_owner to default
        self.owner = std::mem::take(&mut self.proposed_owner);
        Ok(())
    }

    pub fn is_router(&self, caller: Pubkey) -> bool {
        Pubkey::find_program_address(&[EXTERNAL_EXECUTION_CONFIG_SEED], &self.router).0 == caller
    }

    pub fn update_router(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_keys_neq!(router, Pubkey::default(), CcipReceiverError::InvalidRouter);
        require_keys_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);
        self.router = router;
        Ok(())
    }
}

#[account]
#[derive(InitSpace, Default, Debug)]
pub struct ApprovedSender {}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SVMMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SVMTokenAmount>,
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // solana local token amount
}

#[error_code]
pub enum CcipReceiverError {
    #[msg("Address is not router external execution PDA")]
    OnlyRouter,
    #[msg("Invalid router address")]
    InvalidRouter,
    #[msg("Invalid combination of chain and sender")]
    InvalidChainAndSender,
    #[msg("Address is not owner")]
    OnlyOwner,
    #[msg("Address is not proposed_owner")]
    OnlyProposedOwner,
    #[msg("Caller is not allowed")]
    InvalidCaller,
    #[msg("Proposed owner is invalid")]
    InvalidProposedOwner,
}

#[event]
pub struct MessageReceived {
    pub message_id: [u8; 32],
}

#[cfg(test)]
mod tests {
    use super::*;

    fn create_state() -> BaseState {
        BaseState {
            owner: Pubkey::new_unique(),
            ..BaseState::default()
        }
    }

    #[test]
    fn ownership() {
        let mut state = create_state();
        let next_owner = Pubkey::new_unique();

        // only owner can propose
        assert_eq!(
            state
                .transfer_ownership(Pubkey::new_unique(), Pubkey::new_unique())
                .unwrap_err(),
            CcipReceiverError::OnlyOwner.into()
        );
        state.transfer_ownership(state.owner, next_owner).unwrap();

        // only proposed_owner can accept
        assert_eq!(
            state.accept_ownership(Pubkey::new_unique()).unwrap_err(),
            CcipReceiverError::OnlyProposedOwner.into(),
        );
        state.accept_ownership(next_owner).unwrap();
    }

    #[test]
    fn router() {
        let mut state = create_state();

        assert_eq!(
            state
                .update_router(state.owner, Pubkey::default())
                .unwrap_err(),
            CcipReceiverError::InvalidRouter.into(),
        );
        assert_eq!(
            state
                .update_router(Pubkey::new_unique(), Pubkey::new_unique())
                .unwrap_err(),
            CcipReceiverError::OnlyOwner.into(),
        );
        state
            .update_router(state.owner, Pubkey::new_unique())
            .unwrap();
    }
}
