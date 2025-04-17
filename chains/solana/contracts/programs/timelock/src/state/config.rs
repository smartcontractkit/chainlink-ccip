use anchor_lang::prelude::*;
use static_assertions::const_assert;
use std::mem;

use arrayvec::arrayvec;

use crate::constants::{MAX_SELECTORS, TIMELOCK_ID_PADDED};
use crate::error::TimelockError;

#[derive(AnchorSerialize, AnchorDeserialize, Clone, PartialEq)]
pub enum Role {
    Admin = 0,
    Proposer = 1,
    Executor = 2,
    Canceller = 3,
    Bypasser = 4,
}

#[account(zero_copy)]
#[derive(InitSpace, AnchorSerialize, AnchorDeserialize)]
pub struct Config {
    pub timelock_id: [u8; TIMELOCK_ID_PADDED],

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    pub proposer_role_access_controller: Pubkey,
    pub executor_role_access_controller: Pubkey,
    pub canceller_role_access_controller: Pubkey,
    pub bypasser_role_access_controller: Pubkey,

    pub min_delay: u64, // initial minimum delay for operations

    pub blocked_selectors: BlockedSelectors,
}

impl Config {
    pub fn get_role_controller(&self, role: &Role) -> Pubkey {
        match role {
            Role::Proposer => self.proposer_role_access_controller,
            Role::Executor => self.executor_role_access_controller,
            Role::Canceller => self.canceller_role_access_controller,
            Role::Bypasser => self.bypasser_role_access_controller,
            _ => panic!("invalid role"),
        }
    }
}

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize)]
pub struct BlockedSelectors {
    pub xs: [[u8; 8]; MAX_SELECTORS],
    pub len: u64,
}

arrayvec!(BlockedSelectors, [u8; 8], u64);
const_assert!(
    mem::size_of::<BlockedSelectors>()
        == mem::size_of::<u64>() + mem::size_of::<[u8; 8]>() * MAX_SELECTORS
);

impl Space for BlockedSelectors {
    const INIT_SPACE: usize = 8 + (8 * MAX_SELECTORS);
}

impl BlockedSelectors {
    pub fn is_blocked(&self, selector: &[u8; 8]) -> bool {
        self.as_slice().binary_search(selector).is_ok()
    }

    pub fn block_selector(&mut self, selector: [u8; 8]) -> Result<()> {
        match self.as_slice().binary_search(&selector) {
            Ok(_) => {
                return err!(TimelockError::AlreadyBlocked);
            }
            Err(pos) => {
                require!(
                    self.len() < MAX_SELECTORS,
                    TimelockError::MaxCapacityReached
                );
                self.insert(pos, selector);
            }
        }
        Ok(())
    }

    pub fn unblock_selector(&mut self, selector: [u8; 8]) -> Result<()> {
        match self.as_slice().binary_search(&selector) {
            Ok(pos) => {
                self.remove(pos);
            }
            Err(_) => {
                return err!(TimelockError::SelectorNotFound);
            }
        }
        Ok(())
    }
}
