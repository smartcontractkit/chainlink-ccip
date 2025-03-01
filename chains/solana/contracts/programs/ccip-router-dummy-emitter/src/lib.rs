use anchor_lang::prelude::*;
pub mod events;
pub mod state;
pub mod messages;

use crate::events::on_ramp::*;
use crate::events::admin::*;
use crate::events::token_admin_registry::*;

use crate::state::DestChainConfig;
use crate::messages::SVM2AnyRampMessage;
use crate::state::CodeVersion;
declare_id!("7sDY5A5S5NZe1zcqEuZybW6ZxAna1NWUZxU4ypdn8UQU");

#[program]
pub mod dummy {
    use super::*;

    /// Emit events gracefully. Will mark transaction as succeded.
    pub fn trigger_all_events(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        //1. Emit CCIPMessageSent
        emit!(CCIPMessageSent {
            dest_chain_selector: 1,
            sequence_number: 1,
            message: SVM2AnyRampMessage::default(),
        });

        // 2. Emit ConfigSet
        emit!(ConfigSet {
            svm_chain_selector: 1,
            fee_quoter: Pubkey::default(),
            link_token_mint: Pubkey::default(),
            fee_aggregator: Pubkey::default(),
        });

        // 3. Emit FeeToken events
        let fee_token = Pubkey::default();
        emit!(FeeTokenAdded {
            fee_token,
            enabled: true,
        });
        emit!(FeeTokenEnabled { fee_token });
        emit!(FeeTokenDisabled { fee_token });
        emit!(FeeTokenRemoved { fee_token });

        // 4. Emit DestChain events
        let dest_chain_config = DestChainConfig {
            lane_code_version: CodeVersion::V1,
            allowed_senders: vec![Pubkey::default()],
            allow_list_enabled: true,
        };
        
        emit!(DestChainConfigUpdated {
            dest_chain_selector: 2,
            dest_chain_config: dest_chain_config.clone(),
        });
        
        emit!(DestChainAdded {
            dest_chain_selector: 3,
            dest_chain_config,
        });

        // 5. Emit Ownership events
        let owner = Pubkey::default();
        let new_owner = Pubkey::default();
        emit!(OwnershipTransferRequested {
            from: owner,
            to: new_owner,
        });
        emit!(OwnershipTransferred {
            from: owner,
            to: new_owner,
        });

        // 6. Emit Offramp events
        let offramp = Pubkey::default();
        emit!(OfframpAdded {
            source_chain_selector: 4,
            offramp,
        });
        emit!(OfframpRemoved {
            source_chain_selector: 4,
            offramp,
        });

        // 7. Emit Version events
        emit!(CcipVersionForDestChainVersionBumped {
            dest_chain_selector: 5,
            previous_sequence_number: 1,
            new_sequence_number: 2,
        });
        emit!(CcipVersionForDestChainVersionRolledBack {
            dest_chain_selector: 5,
            previous_sequence_number: 2,
            new_sequence_number: 1,
        });

        // 8. Emit TokenAdminRegistry events
        let token = Pubkey::default();
        let admin = Pubkey::default();
        let new_admin = Pubkey::default();
        emit!(PoolSet {
            token,
            previous_pool_lookup_table: Pubkey::default(),
            new_pool_lookup_table: Pubkey::default(),
        });
        emit!(AdministratorTransferRequested {
            token,
            current_admin: admin,
            new_admin,
        });
        emit!(AdministratorTransferred {
            token,
            new_admin,
        });
       
        msg!("All events have been emitted!");
        Ok(())
    }

    /// Emit events and reverts. Will mark transaction as failed if client has sent the tx
    /// with SkipPreflight disabled (simulation before sending)
    pub fn trigger_all_events_reverts(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        //1. Emit CCIPMessageSent
        emit!(CCIPMessageSent {
            dest_chain_selector: 1,
            sequence_number: 1,
            message: SVM2AnyRampMessage::default(),
        });

        // 2. Emit ConfigSet
        emit!(ConfigSet {
            svm_chain_selector: 1,
            fee_quoter: Pubkey::default(),
            link_token_mint: Pubkey::default(),
            fee_aggregator: Pubkey::default(),
        });

        // 3. Emit FeeToken events
        let fee_token = Pubkey::default();
        emit!(FeeTokenAdded {
            fee_token,
            enabled: true,
        });
        emit!(FeeTokenEnabled { fee_token });
        emit!(FeeTokenDisabled { fee_token });
        emit!(FeeTokenRemoved { fee_token });

        // 4. Emit DestChain events
        let dest_chain_config = DestChainConfig {
            lane_code_version: CodeVersion::V1,
            allowed_senders: vec![Pubkey::default()],
            allow_list_enabled: true,
        };
        
        emit!(DestChainConfigUpdated {
            dest_chain_selector: 2,
            dest_chain_config: dest_chain_config.clone(),
        });
        
        emit!(DestChainAdded {
            dest_chain_selector: 3,
            dest_chain_config,
        });

        // 5. Emit Ownership events
        let owner = Pubkey::default();
        let new_owner = Pubkey::default();
        emit!(OwnershipTransferRequested {
            from: owner,
            to: new_owner,
        });
        emit!(OwnershipTransferred {
            from: owner,
            to: new_owner,
        });

        // 6. Emit Offramp events
        let offramp = Pubkey::default();
        emit!(OfframpAdded {
            source_chain_selector: 4,
            offramp,
        });
        emit!(OfframpRemoved {
            source_chain_selector: 4,
            offramp,
        });

        // 7. Emit Version events
        emit!(CcipVersionForDestChainVersionBumped {
            dest_chain_selector: 5,
            previous_sequence_number: 1,
            new_sequence_number: 2,
        });
        emit!(CcipVersionForDestChainVersionRolledBack {
            dest_chain_selector: 5,
            previous_sequence_number: 2,
            new_sequence_number: 1,
        });

        // 8. Emit TokenAdminRegistry events
        let token = Pubkey::default();
        let admin = Pubkey::default();
        let new_admin = Pubkey::default();
        emit!(PoolSet {
            token,
            previous_pool_lookup_table: Pubkey::default(),
            new_pool_lookup_table: Pubkey::default(),
        });
        emit!(AdministratorTransferRequested {
            token,
            current_admin: admin,
            new_admin,
        });
        emit!(AdministratorTransferred {
            token,
            new_admin,
        });
       
        msg!("All events have been emitted!");
        return Err(error!(CustomError::TestRevert));
    }
}

/// Context struct for the trigger_all_events instruction
#[derive(Accounts)]
pub struct TriggerAllEvents {}

#[error_code]
pub enum CustomError {
    #[msg("This is a forced revert for testing log persistence.")]
    TestRevert = 4100,
}