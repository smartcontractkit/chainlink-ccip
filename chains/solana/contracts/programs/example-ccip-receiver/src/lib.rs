use anchor_lang::prelude::*;
use arrayvec::arrayvec;
declare_id!("CcipReceiver1111111111111111111111111111111");

pub const EXTERNAL_EXECUTION_CONFIG_SEED: &[u8] = b"external_execution_config";

/// This program an example of a CCIP Receiver Program.
/// Used to test CCIP Router execute.
#[program]
pub mod example_ccip_receiver {
    use super::*;

    /// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .init(ctx.accounts.authority.key(), router)
    }

    /// This function is called by the CCIP Router to execute the CCIP message.
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

    pub fn enable_list(
        ctx: Context<UpdateConfig>,
        list_type: ListType,
        enable: bool,
    ) -> Result<()> {
        ctx.accounts
            .state
            .enable_list(ctx.accounts.authority.key(), list_type, enable)
    }

    pub fn add_chain_to(
        ctx: Context<UpdateConfig>,
        list_type: ListType,
        chain_selector: u64,
    ) -> Result<()> {
        ctx.accounts
            .state
            .add_chain_to(ctx.accounts.authority.key(), chain_selector, list_type)
    }

    pub fn remove_chain_from(
        ctx: Context<UpdateConfig>,
        list_type: ListType,
        chain_selector: u64,
    ) -> Result<()> {
        ctx.accounts
            .state
            .add_chain_to(ctx.accounts.authority.key(), chain_selector, list_type)
    }

    pub fn update_router(ctx: Context<UpdateConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .update_router(ctx.accounts.authority.key(), new_router)
    }

    pub fn propose_owner(ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .propose_owner(ctx.accounts.authority.key(), proposed_owner)
    }

    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts
            .state
            .accept_ownership(ctx.accounts.authority.key())
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
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(message: Any2SVMMessage)]
pub struct CcipReceive<'info> {
    // router CPI signer must be first
    #[account(
        constraint = state.is_valid_chain(message.source_chain_selector) @ CcipReceiverError::InvalidChain,
        constraint = state.is_router(authority.key()) @ CcipReceiverError::OnlyRouter,
    )]
    pub authority: Signer<'info>,
    pub state: Account<'info, BaseState>,
}

#[derive(Accounts, Debug)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
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
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.proposed_owner @ CcipReceiverError::OnlyProposedOwner,
    )]
    pub authority: Signer<'info>,
}

// BaseState contains the state for core safety checks that can be leveraged by the implementer
// Base state contains a limited size allow and deny list
// Both are included to handle the size limitations on solana
// If user wants to allow a small number of chains, consider using the allow list (disable deny list)
// If user wants to allow many chains, consider using the deny list (disable allow list)
#[account]
#[derive(InitSpace, Default, Debug)]
pub struct BaseState {
    pub owner: Pubkey,
    pub proposed_owner: Pubkey,
    pub router: Pubkey,
    pub allow: ChainList,
    pub deny: ChainList,
}

#[derive(InitSpace, Debug, Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct ChainList {
    pub is_enabled: bool,
    pub xs: [u64; 20],
    len: u8,
}
arrayvec!(ChainList, u64, u8);

impl BaseState {
    pub fn init(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_eq!(self.owner, Pubkey::default());
        self.owner = owner;
        self.update_router(owner, router)
    }

    pub fn propose_owner(&mut self, owner: Pubkey, proposed_owner: Pubkey) -> Result<()> {
        require_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);
        self.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(&mut self, proposed_owner: Pubkey) -> Result<()> {
        require_eq!(
            self.proposed_owner,
            proposed_owner,
            CcipReceiverError::OnlyProposedOwner
        );
        self.proposed_owner = Pubkey::default();
        self.owner = proposed_owner;
        Ok(())
    }

    pub fn is_router(&self, caller: Pubkey) -> bool {
        Pubkey::find_program_address(&[EXTERNAL_EXECUTION_CONFIG_SEED], &self.router).0 == caller
    }

    pub fn update_router(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_keys_neq!(router, Pubkey::default(), CcipReceiverError::InvalidRouter);
        require_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);
        self.router = router;
        Ok(())
    }

    pub fn is_valid_chain(&self, chain_selector: u64) -> bool {
        if self.allow.is_enabled && self.deny.is_enabled {
            // must be within allow list and not in deny list
            self.allow.binary_search(&chain_selector).is_ok()
                && !self.deny.binary_search(&chain_selector).is_ok()
        } else if self.allow.is_enabled && !self.deny.is_enabled {
            // check allow list only
            self.allow.binary_search(&chain_selector).is_ok()
        } else if !self.allow.is_enabled && self.deny.is_enabled {
            // check deny list only, if present = not valid
            !self.deny.binary_search(&chain_selector).is_ok()
        } else {
            // neither list is enabled, allow everything
            true
        }
    }

    pub fn enable_list(&mut self, owner: Pubkey, list_type: ListType, enable: bool) -> Result<()> {
        require_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);
        let list = self.get_list(list_type);
        list.is_enabled = enable;
        Ok(())
    }

    pub fn add_chain_to(
        &mut self,
        owner: Pubkey,
        chain_selector: u64,
        list_type: ListType,
    ) -> Result<()> {
        require_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);

        let list = self.get_list(list_type);
        require!(list.remaining_capacity() > 0, CcipReceiverError::Full);
        match list.binary_search(&chain_selector) {
            // already present
            Ok(_) => (),
            Err(i) => list.insert(i, chain_selector),
        }

        Ok(())
    }

    pub fn remove_chain_from(
        &mut self,
        owner: Pubkey,
        chain_selector: u64,
        list_type: ListType,
    ) -> Result<()> {
        require_eq!(self.owner, owner, CcipReceiverError::OnlyOwner);
        let list = self.get_list(list_type);
        let index = list.binary_search(&chain_selector);
        if let Ok(index) = index {
            list.remove(index);
        }
        Ok(())
    }

    fn get_list(&mut self, list_type: ListType) -> &mut ChainList {
        match list_type {
            ListType::Allow => &mut self.allow,
            ListType::Deny => &mut self.deny,
        }
    }
}

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

#[repr(u8)]
#[derive(AnchorSerialize, AnchorDeserialize)]
pub enum ListType {
    Allow,
    Deny,
}

#[error_code]
pub enum CcipReceiverError {
    #[msg("Address is not router external execution PDA")]
    OnlyRouter,
    #[msg("Invalid router address")]
    InvalidRouter,
    #[msg("Invalid chain")]
    InvalidChain,
    #[msg("Address is not owner")]
    OnlyOwner,
    #[msg("Address is not proposed_owner")]
    OnlyProposedOwner,
    #[msg("List is full")]
    Full,
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
                .propose_owner(Pubkey::new_unique(), Pubkey::new_unique())
                .unwrap_err(),
            CcipReceiverError::OnlyOwner.into()
        );
        state.propose_owner(state.owner, next_owner).unwrap();

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

    #[test]
    fn chains() {
        let mut state = create_state();

        assert_eq!(
            state
                .enable_list(Pubkey::new_unique(), ListType::Allow, true)
                .unwrap_err(),
            CcipReceiverError::OnlyOwner.into(),
        );
        assert_eq!(
            state
                .add_chain_to(Pubkey::new_unique(), 1, ListType::Deny)
                .unwrap_err(),
            CcipReceiverError::OnlyOwner.into(),
        );
        assert_eq!(
            state
                .remove_chain_from(Pubkey::new_unique(), 1, ListType::Deny)
                .unwrap_err(),
            CcipReceiverError::OnlyOwner.into(),
        );

        assert!(state.allow.is_empty());
        assert!(state.deny.is_empty());

        // add 2 chain selector to allow
        state
            .add_chain_to(state.owner, 10, ListType::Allow)
            .unwrap();
        assert_eq!(state.allow.len(), 1);
        assert_eq!(state.allow.xs[0], 10);
        state
            .add_chain_to(state.owner, 40, ListType::Allow)
            .unwrap();
        assert_eq!(state.allow.len(), 2);
        assert_eq!(state.allow.xs[0], 10);
        assert_eq!(state.allow.xs[1], 40);

        // add 3 chain selectors to deny
        state.add_chain_to(state.owner, 20, ListType::Deny).unwrap();
        assert_eq!(state.deny.len(), 1);
        assert_eq!(state.deny.xs[0], 20);
        state.add_chain_to(state.owner, 40, ListType::Deny).unwrap();
        assert_eq!(state.deny.len(), 2);
        assert_eq!(state.deny.xs[0], 20);
        assert_eq!(state.deny.xs[1], 40);
        state.add_chain_to(state.owner, 21, ListType::Deny).unwrap();
        assert_eq!(state.deny.len(), 3);
        assert_eq!(state.deny.xs[0], 20);
        assert_eq!(state.deny.xs[1], 21);
        assert_eq!(state.deny.xs[2], 40);

        // remove chain selector from deny
        state
            .remove_chain_from(state.owner, 21, ListType::Deny)
            .unwrap();
        assert_eq!(state.deny.len(), 2);
        assert_eq!(state.deny.xs[0], 20);
        assert_eq!(state.deny.xs[1], 40);
        // remove same chain selector from deny
        state
            .remove_chain_from(state.owner, 21, ListType::Deny)
            .unwrap();
        assert_eq!(state.deny.len(), 2);
        assert_eq!(state.deny.xs[0], 20);
        assert_eq!(state.deny.xs[1], 40);

        // no lists enabled
        assert!(state.is_valid_chain(10)); // in allow list
        assert!(state.is_valid_chain(20)); // in deny list
        assert!(state.is_valid_chain(30)); // in neither list
        assert!(state.is_valid_chain(40)); // in both lists

        // only allow list enabled
        state
            .enable_list(state.owner, ListType::Allow, true)
            .unwrap();
        assert!(state.is_valid_chain(10)); // in allow list
        assert!(!state.is_valid_chain(20)); // in deny list
        assert!(!state.is_valid_chain(30)); // in neither list
        assert!(state.is_valid_chain(40)); // in both lists

        // both lists enabled
        state
            .enable_list(state.owner, ListType::Deny, true)
            .unwrap();
        assert!(state.is_valid_chain(10)); // in allow list
        assert!(!state.is_valid_chain(20)); // in deny list
        assert!(!state.is_valid_chain(30)); // in neither list
        assert!(!state.is_valid_chain(40)); // in both lists

        // only deny list enabled
        state
            .enable_list(state.owner, ListType::Allow, false)
            .unwrap();
        assert!(state.is_valid_chain(10)); // in allow list
        assert!(!state.is_valid_chain(20)); // in deny list
        assert!(state.is_valid_chain(30)); // in neither list
        assert!(!state.is_valid_chain(40)); // in both lists
    }
}
