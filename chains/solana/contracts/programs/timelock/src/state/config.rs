use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize, Debug, Clone, PartialEq)]
pub enum Role {
    Admin = 0,
    Proposer = 1,
    Executor = 2,
    Canceller = 3,
    Bypasser = 4,
}

#[account]
#[derive(InitSpace)]
pub struct Config {
    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    pub proposer_role_access_controller: Pubkey,
    pub executor_role_access_controller: Pubkey,
    pub canceller_role_access_controller: Pubkey,
    pub bypasser_role_access_controller: Pubkey,

    pub min_delay: u64, // initial minimum delay for operations
}

impl Config {
    pub fn get_role_controller(&self, role: &Role) -> Pubkey {
        match role {
            Role::Proposer => self.proposer_role_access_controller,
            Role::Executor => self.executor_role_access_controller,
            Role::Canceller => self.canceller_role_access_controller,
            Role::Bypasser => self.bypasser_role_access_controller,
            _ => panic!("Invalid role"),
        }
    }
}
