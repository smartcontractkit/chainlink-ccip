use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use crate::context::UpdateDestChainSelectorConfigNoRealloc;
use crate::events::admin as events;
use crate::state::{CodeVersion, RestoreOnAction};
use crate::{
    AcceptOwnership, AddChainSelector, CcipRouterError, DestChainConfig, DestChainState,
    TransferOwnership, UpdateConfigCCIPRouter, UpdateDestChainSelectorConfig, WithdrawBilledFunds,
};
use crate::{AddOfframp, RemoveOfframp};

use super::super::interfaces::Admin;
use super::fees::do_billing_transfer;

pub struct Impl;
impl Admin for Impl {
    fn transfer_ownership(
        &self,
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(
            proposed_owner != config.owner,
            CcipRouterError::RedundantOwnerProposal
        );
        emit!(events::OwnershipTransferRequested {
            from: config.owner,
            to: proposed_owner,
        });
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()> {
        let config = &mut ctx.accounts.config;
        emit!(events::OwnershipTransferred {
            from: config.owner,
            to: config.proposed_owner,
        });
        // NOTE: take() resets proposed_owner to default
        config.owner = std::mem::take(&mut config.proposed_owner);
        Ok(())
    }

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        code_version: CodeVersion,
    ) -> Result<()> {
        require_neq!(
            code_version,
            CodeVersion::Default,
            CcipRouterError::InvalidCodeVersion
        );
        ctx.accounts.config.default_code_version = code_version;
        Ok(())
    }

    fn set_link_token_mint(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        link_token_mint: Pubkey,
    ) -> Result<()> {
        ctx.accounts.config.link_token_mint = link_token_mint;
        emit!(events::ConfigSet {
            svm_chain_selector: ctx.accounts.config.svm_chain_selector,
            fee_quoter: ctx.accounts.config.fee_quoter,
            rmn_remote: ctx.accounts.config.rmn_remote,
            link_token_mint: ctx.accounts.config.link_token_mint,
            fee_aggregator: ctx.accounts.config.fee_aggregator,
        });
        Ok(())
    }

    fn update_fee_aggregator(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        fee_aggregator: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.fee_aggregator = fee_aggregator;
        emit!(events::ConfigSet {
            svm_chain_selector: config.svm_chain_selector,
            fee_quoter: config.fee_quoter,
            rmn_remote: config.rmn_remote,
            link_token_mint: config.link_token_mint,
            fee_aggregator: config.fee_aggregator,
        });
        Ok(())
    }

    fn update_rmn_remote(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.rmn_remote = rmn_remote;
        emit!(events::ConfigSet {
            svm_chain_selector: config.svm_chain_selector,
            fee_quoter: config.fee_quoter,
            rmn_remote: config.rmn_remote,
            link_token_mint: config.link_token_mint,
            fee_aggregator: config.fee_aggregator,
        });
        Ok(())
    }

    fn add_chain_selector(
        &self,
        ctx: Context<AddChainSelector>,
        new_chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        // Set dest chain config & state
        let dest_chain_state = &mut ctx.accounts.dest_chain_state;
        validate_dest_chain_config(new_chain_selector, &dest_chain_config)?;
        dest_chain_state.version = 1;
        dest_chain_state.chain_selector = new_chain_selector;
        dest_chain_state.config = dest_chain_config.clone();
        dest_chain_state.state = DestChainState {
            sequence_number: 0,
            restore_on_action: RestoreOnAction::None,
            sequence_number_to_restore: 0,
        };

        emit!(events::DestChainAdded {
            dest_chain_selector: new_chain_selector,
            dest_chain_config,
        });

        Ok(())
    }

    fn update_dest_chain_config(
        &self,
        ctx: Context<UpdateDestChainSelectorConfig>,
        dest_chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        validate_dest_chain_config(dest_chain_selector, &dest_chain_config)?;

        ctx.accounts.dest_chain_state.config = dest_chain_config.clone();

        emit!(events::DestChainConfigUpdated {
            dest_chain_selector,
            dest_chain_config,
        });
        Ok(())
    }

    fn bump_ccip_version_for_dest_chain(
        &self,
        ctx: Context<UpdateDestChainSelectorConfigNoRealloc>,
        dest_chain_selector: u64,
    ) -> Result<()> {
        let dest_chain_state = &mut ctx.accounts.dest_chain_state.state;

        let previous_sequence_number = dest_chain_state.sequence_number;

        upgrade(dest_chain_state)?;

        emit!(events::CcipVersionForDestChainVersionBumped {
            dest_chain_selector,
            previous_sequence_number,
            new_sequence_number: dest_chain_state.sequence_number,
        });

        Ok(())
    }

    fn rollback_ccip_version_for_dest_chain(
        &self,
        ctx: Context<UpdateDestChainSelectorConfigNoRealloc>,
        dest_chain_selector: u64,
    ) -> Result<()> {
        let dest_chain_state = &mut ctx.accounts.dest_chain_state.state;

        let previous_sequence_number = dest_chain_state.sequence_number;

        rollback(dest_chain_state)?;

        emit!(events::CcipVersionForDestChainVersionRolledBack {
            dest_chain_selector,
            previous_sequence_number,
            new_sequence_number: dest_chain_state.sequence_number,
        });

        Ok(())
    }

    fn add_offramp(
        &self,
        _ctx: Context<AddOfframp>,
        source_chain_selector: u64,
        offramp: Pubkey,
    ) -> Result<()> {
        emit!(events::OfframpAdded {
            source_chain_selector,
            offramp,
        });
        Ok(())
    }

    fn remove_offramp(
        &self,
        _ctx: Context<RemoveOfframp>,
        source_chain_selector: u64,
        offramp: Pubkey,
    ) -> Result<()> {
        emit!(events::OfframpRemoved {
            source_chain_selector,
            offramp,
        });
        Ok(())
    }

    fn update_svm_chain_selector(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        new_chain_selector: u64,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;

        config.svm_chain_selector = new_chain_selector;

        emit!(events::ConfigSet {
            svm_chain_selector: config.svm_chain_selector,
            fee_quoter: config.fee_quoter,
            rmn_remote: config.rmn_remote,
            link_token_mint: config.link_token_mint,
            fee_aggregator: config.fee_aggregator,
        });

        Ok(())
    }

    fn withdraw_billed_funds(
        &self,
        ctx: Context<WithdrawBilledFunds>,
        transfer_all: bool,
        desired_amount: u64, // if transfer_all is true, this value must be 0
    ) -> Result<()> {
        let transfer = token_interface::TransferChecked {
            from: ctx.accounts.fee_token_accum.to_account_info(),
            to: ctx.accounts.recipient.to_account_info(),
            mint: ctx.accounts.fee_token_mint.to_account_info(),
            authority: ctx.accounts.fee_billing_signer.to_account_info(),
        };

        let amount = if transfer_all {
            require!(
                desired_amount == 0,
                CcipRouterError::InvalidInputsTransferAllAmount
            );
            require!(
                ctx.accounts.fee_token_accum.amount > 0,
                CcipRouterError::InsufficientFunds
            );
            ctx.accounts.fee_token_accum.amount
        } else {
            require!(
                desired_amount > 0,
                CcipRouterError::InvalidInputsTokenAmount
            );
            require!(
                desired_amount <= ctx.accounts.fee_token_accum.amount,
                CcipRouterError::InsufficientFunds
            );
            desired_amount
        };

        do_billing_transfer(
            ctx.accounts.token_program.to_account_info(),
            transfer,
            amount,
            ctx.accounts.fee_token_mint.decimals,
            ctx.bumps.fee_billing_signer,
        )
    }
}

/////////////
// Helpers //
/////////////

fn validate_dest_chain_config(dest_chain_selector: u64, _config: &DestChainConfig) -> Result<()> {
    // As of now, the config has very few properties and there is very little to validate yet.
    // This is mainly a placeholder to add validations as that config object grows.
    require!(
        dest_chain_selector != 0,
        CcipRouterError::InvalidInputsChainSelector
    );
    Ok(())
}

/// On upgrade, the sequence number is reset to 0 and the previous sequence number is saved for a future rollback.
/// If a rollback has previously occurred, on upgrade the previous sequence number for the current version is restored
/// instead of resetting to 0.
fn upgrade(dest_chain_state: &mut DestChainState) -> Result<()> {
    let current_seq_nr = dest_chain_state.sequence_number;

    dest_chain_state.sequence_number = match dest_chain_state.restore_on_action {
        RestoreOnAction::Upgrade => dest_chain_state.sequence_number_to_restore,
        _ => 0,
    };
    dest_chain_state.sequence_number_to_restore = current_seq_nr;

    // restore on next rollback, as seq nr was of the previous CCIP version
    dest_chain_state.restore_on_action = RestoreOnAction::Rollback;

    Ok(())
}

/// On rollback, the sequence number is restored to the previous version's sequence number. The current sequence number
/// for the version being rolled back to is saved for a future upgrade, so it can resume from where it was.
fn rollback(dest_chain_state: &mut DestChainState) -> Result<()> {
    // If there was loss of information, we can't rollback. We support at most 1 consecutive rollback.
    // So, once a rollback has happened, the admin must bump the CCIP version before another rollback.
    require_eq!(
        dest_chain_state.restore_on_action,
        RestoreOnAction::Rollback,
        CcipRouterError::InvalidCcipVersionRollback
    );

    std::mem::swap(
        &mut dest_chain_state.sequence_number,
        &mut dest_chain_state.sequence_number_to_restore,
    );

    // restore on next upgrade, as seq nr was of the previously-bumped CCIP version
    dest_chain_state.restore_on_action = RestoreOnAction::Upgrade;

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn first_upgrade() {
        let state = &mut DestChainState {
            sequence_number: 5,
            restore_on_action: RestoreOnAction::None, // initial state
            sequence_number_to_restore: 0,
        };
        upgrade(state).unwrap();
        assert_eq!(
            state,
            &DestChainState {
                sequence_number: 0,
                restore_on_action: RestoreOnAction::Rollback,
                sequence_number_to_restore: 5,
            }
        );
    }

    #[test]
    fn second_upgrade() {
        let state = &mut DestChainState {
            sequence_number: 2,
            restore_on_action: RestoreOnAction::Rollback,
            sequence_number_to_restore: 5,
        };
        upgrade(state).unwrap();
        assert_eq!(
            state,
            &DestChainState {
                sequence_number: 0,
                restore_on_action: RestoreOnAction::Rollback,
                sequence_number_to_restore: 2,
            }
        );
    }

    #[test]
    fn first_rollback() {
        let state = &mut DestChainState {
            sequence_number: 3,
            restore_on_action: RestoreOnAction::Rollback,
            sequence_number_to_restore: 2,
        };
        rollback(state).unwrap();
        assert_eq!(
            state,
            &DestChainState {
                sequence_number: 2,
                restore_on_action: RestoreOnAction::Upgrade,
                sequence_number_to_restore: 3,
            }
        );
    }

    #[test]
    fn second_rollback() {
        let state = &mut DestChainState {
            sequence_number: 2,
            restore_on_action: RestoreOnAction::Upgrade,
            sequence_number_to_restore: 3,
        };
        let result = rollback(state);
        assert!(result.is_err()); // can't rollback twice in a row
    }

    #[test]
    fn upgrade_after_rollback() {
        let state = &mut DestChainState {
            sequence_number: 2,
            restore_on_action: RestoreOnAction::Upgrade,
            sequence_number_to_restore: 3,
        };
        upgrade(state).unwrap();
        assert_eq!(
            state,
            &DestChainState {
                sequence_number: 3,
                restore_on_action: RestoreOnAction::Rollback,
                sequence_number_to_restore: 2,
            }
        );
    }
}
