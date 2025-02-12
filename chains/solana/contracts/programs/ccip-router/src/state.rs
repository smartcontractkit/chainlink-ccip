use std::fmt::Display;

use anchor_lang::prelude::*;

// zero_copy is used to prevent hitting stack/heap memory limits
#[account(zero_copy)] // TODO this is no longer needed as zero_copy
#[derive(InitSpace, AnchorSerialize, AnchorDeserialize)]
pub struct Config {
    pub version: u8,
    _padding0: [u8; 7],
    pub svm_chain_selector: u64,

    _padding1: [u8; 8],

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    _padding2: [u8; 8],
    pub fee_quoter: Pubkey,

    pub link_token_mint: Pubkey,
    pub fee_aggregator: Pubkey, // Allowed address to withdraw billed fees to (will use ATAs derived from it)
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug, PartialEq, Eq, Copy)]
pub enum RestoreOnAction {
    None,
    Upgrade,
    Rollback,
}

impl Display for RestoreOnAction {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            RestoreOnAction::None => write!(f, "None"),
            RestoreOnAction::Upgrade => write!(f, "Upgrade"),
            RestoreOnAction::Rollback => write!(f, "Rollback"),
        }
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug, PartialEq, Eq, Copy)]
pub struct DestChainState {
    pub sequence_number: u64, // The last used sequence number. On upgrades, this is reset to 0.

    // The following property is used to support one rollback operation, in which the sequence number of
    // the previous OnRamp version is restored. The upgrade/rollback is done per-lane (i.e. per dest chain).
    // If it's 0, that means that it is not possible to rollback to the previous version. As version upgrades
    // are not often, we won't need to rollback multiple versions.
    pub sequence_number_to_restore: u64, // The last used sequence number in the previous onramp version
    pub restore_on_action: RestoreOnAction,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct DestChainConfig {
    // list of senders authorized to send messages to this destination chain.
    // Note: The attribute name `max_len` is slightly misleading: it is not in any
    // way limiting the actual length of the vector during initialization; it just
    // helps the InitSpace derive macro work out the initial space. We can leave it at
    // zero and calculate the actual length in the instruction context.
    #[max_len(0)]
    pub allowed_senders: Vec<Pubkey>,
    pub allow_list_enabled: bool,
}

impl DestChainConfig {
    pub fn space(&self) -> usize {
        Self::INIT_SPACE + self.dynamic_space()
    }

    pub fn dynamic_space(&self) -> usize {
        self.allowed_senders.len() * std::mem::size_of::<Pubkey>()
    }
}

#[account]
#[derive(InitSpace, Debug)]
pub struct DestChain {
    // Config for SVM2Any
    pub version: u8,
    pub chain_selector: u64,     // Chain selector used for the seed
    pub state: DestChainState,   // values that are updated automatically
    pub config: DestChainConfig, // values configured by an admin
}

#[account]
#[derive(InitSpace)]
pub struct Nonce {
    pub version: u8,  // version to check if nonce account is already initialized
    pub counter: u64, // Counter per user and per lane to use as nonce for all the messages to be executed in order
}

#[account]
#[derive(InitSpace)]
pub struct ExternalExecutionConfig {}
