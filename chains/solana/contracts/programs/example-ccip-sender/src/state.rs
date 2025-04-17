use anchor_lang::prelude::*;

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
            CcipSenderError::InvalidProposedOwner
        );
        require_keys_eq!(self.owner, owner, CcipSenderError::OnlyOwner);
        self.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(&mut self, proposed_owner: Pubkey) -> Result<()> {
        require_keys_eq!(
            self.proposed_owner,
            proposed_owner,
            CcipSenderError::OnlyProposedOwner
        );
        self.proposed_owner = Pubkey::default();
        self.owner = proposed_owner;
        Ok(())
    }

    pub fn update_router(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_keys_neq!(router, Pubkey::default(), CcipSenderError::InvalidRouter);
        require_keys_eq!(self.owner, owner, CcipSenderError::OnlyOwner);
        self.router = router;
        Ok(())
    }
}

#[account]
#[derive(InitSpace, Default, Debug)]
pub struct RemoteChainConfig {
    #[max_len(64)]
    pub recipient: Vec<u8>, // the address to send messages to on the destination chain
    #[max_len(0)] // used to calculate InitSpace, total will include number of extra args bytes
    pub extra_args_bytes: Vec<u8>, // specifies the extraARgs to pass into ccip_send, it will be applied to every out going message for a specific chain
}

impl RemoteChainConfig {
    pub fn set_config(&mut self, recipient: Vec<u8>, extra_args_bytes: Vec<u8>) -> Result<()> {
        self.recipient = recipient;
        self.extra_args_bytes = extra_args_bytes;
        Ok(())
    }
}

pub mod builder {
    use anchor_lang::AnchorSerialize;
    use ccip_router::messages::SVM2AnyMessage;

    pub fn instruction(
        msg: &SVM2AnyMessage,
        discriminator: [u8; 8],
        chain_selector: u64,
    ) -> Vec<u8> {
        let message = msg.try_to_vec().unwrap();
        let chain_selector_bytes = chain_selector.to_le_bytes();

        let mut data = discriminator.to_vec();
        data.extend_from_slice(chain_selector_bytes.as_ref());
        data.extend_from_slice(&message);
        data
    }

    pub fn instruction_with_token_indexes(
        msg: &SVM2AnyMessage,
        discriminator: [u8; 8],
        chain_selector: u64,
        token_indexes: &[u8],
    ) -> Vec<u8> {
        let mut data = instruction(msg, discriminator, chain_selector);
        data.extend_from_slice(token_indexes.try_to_vec().unwrap().as_ref());
        data
    }
}

#[event]
pub struct MessageSent {
    pub message_id: [u8; 32],
}

#[error_code]
pub enum CcipSenderError {
    #[msg("Invalid router address")]
    InvalidRouter,
    #[msg("Address is not owner")]
    OnlyOwner,
    #[msg("Address is not proposed_owner")]
    OnlyProposedOwner,
    #[msg("Proposed owner is invalid")]
    InvalidProposedOwner,
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
            CcipSenderError::OnlyOwner.into()
        );
        state.transfer_ownership(state.owner, next_owner).unwrap();

        // only proposed_owner can accept
        assert_eq!(
            state.accept_ownership(Pubkey::new_unique()).unwrap_err(),
            CcipSenderError::OnlyProposedOwner.into(),
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
            CcipSenderError::InvalidRouter.into(),
        );
        assert_eq!(
            state
                .update_router(Pubkey::new_unique(), Pubkey::new_unique())
                .unwrap_err(),
            CcipSenderError::OnlyOwner.into(),
        );
        state
            .update_router(state.owner, Pubkey::new_unique())
            .unwrap();
    }
}
