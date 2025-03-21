use anchor_lang::prelude::*;

use crate::{DestChainConfig, SVM2AnyRampMessage};

pub mod events {
    use super::*;

    pub mod on_ramp {
        use super::*;
        #[event]
        pub struct CCIPMessageSent {
            pub dest_chain_selector: u64,
            pub sequence_number: u64,
            pub message: SVM2AnyRampMessage,
        }
    }

    pub mod admin {
        use super::*;

        // NOTE: ownership update events can be found in OwnershipTransferRequested, OwnershipTransferred
        #[event]
        pub struct ConfigSet {
            pub svm_chain_selector: u64,
            pub fee_quoter: Pubkey,
            pub rmn_remote: Pubkey,
            pub link_token_mint: Pubkey,
            pub fee_aggregator: Pubkey,
        }

        #[event]
        pub struct FeeTokenAdded {
            pub fee_token: Pubkey,
            pub enabled: bool,
        }

        #[event]
        pub struct FeeTokenEnabled {
            pub fee_token: Pubkey,
        }

        #[event]
        pub struct FeeTokenDisabled {
            pub fee_token: Pubkey,
        }

        #[event]
        pub struct FeeTokenRemoved {
            pub fee_token: Pubkey,
        }

        #[event]
        pub struct DestChainConfigUpdated {
            pub dest_chain_selector: u64,
            pub dest_chain_config: DestChainConfig,
        }

        #[event]
        pub struct DestChainAdded {
            pub dest_chain_selector: u64,
            pub dest_chain_config: DestChainConfig,
        }

        #[event]
        pub struct OwnershipTransferRequested {
            pub from: Pubkey,
            pub to: Pubkey,
        }

        #[event]
        pub struct OwnershipTransferred {
            pub from: Pubkey,
            pub to: Pubkey,
        }

        #[event]
        pub struct OfframpAdded {
            pub source_chain_selector: u64,
            pub offramp: Pubkey,
        }

        #[event]
        pub struct OfframpRemoved {
            pub source_chain_selector: u64,
            pub offramp: Pubkey,
        }

        #[event]
        pub struct CcipVersionForDestChainVersionBumped {
            pub dest_chain_selector: u64,
            pub previous_sequence_number: u64,
            pub new_sequence_number: u64,
        }

        #[event]
        pub struct CcipVersionForDestChainVersionRolledBack {
            pub dest_chain_selector: u64,
            pub previous_sequence_number: u64,
            pub new_sequence_number: u64,
        }
    }

    pub mod token_admin_registry {
        use super::*;
        #[event]
        pub struct PoolSet {
            pub token: Pubkey,
            pub previous_pool_lookup_table: Pubkey,
            pub new_pool_lookup_table: Pubkey,
        }

        #[event]
        pub struct AdministratorTransferRequested {
            pub token: Pubkey,
            pub current_admin: Pubkey,
            pub new_admin: Pubkey,
        }

        #[event]
        pub struct AdministratorTransferred {
            pub token: Pubkey,
            pub new_admin: Pubkey,
        }
    }
}
