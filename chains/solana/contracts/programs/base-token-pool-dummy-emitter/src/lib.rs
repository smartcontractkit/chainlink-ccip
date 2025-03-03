use anchor_lang::prelude::*;

pub mod events;
pub mod state;

use events::*;
use state::*;

declare_id!("CTw4kTVDnSrBohARHWWPwnvRjNYbBx79rDHMR7XcLYDa");

#[program]
pub mod base_token_pool_dummy_emitter {
    use super::*;

    /// Emit events gracefully. Will mark transaction as succeeded.
    pub fn trigger_all_events(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        // 1. Emit TokensConsumed event
        emit!(TokensConsumed {
            tokens: 1000,
        });

        // 2. Emit ConfigChanged event
        emit!(ConfigChanged {
            config: RateLimitConfig {
                enabled: true,
                capacity: 1000000,
                rate: 1000,
            },
        });

        // 3. Emit token events
        emit!(Burned {
            sender: Pubkey::default(),
            amount: 1000,
        });

        emit!(Minted {
            sender: Pubkey::default(),
            recipient: Pubkey::default(),
            amount: 1000,
        });

        emit!(Locked {
            sender: Pubkey::default(),
            amount: 1000,
        });

        emit!(Released {
            sender: Pubkey::default(),
            recipient: Pubkey::default(),
            amount: 1000,
        });

        // 4. Emit chain config events
        let remote_address = RemoteAddress { address: vec![1, 2, 3] };
        emit!(RemoteChainConfigured {
            chain_selector: 1,
            token: remote_address.clone(),
            previous_token: RemoteAddress::default(),
            pool_addresses: vec![remote_address.clone()],
            previous_pool_addresses: vec![],
        });

        // 5. Emit rate limit events
        let rate_limit_config = RateLimitConfig {
            enabled: true,
            capacity: 1000000,
            rate: 1000,
        };
        emit!(RateLimitConfigured {
            chain_selector: 1,
            outbound_rate_limit: rate_limit_config.clone(),
            inbound_rate_limit: rate_limit_config,
        });

        // 6. Emit pool events
        emit!(RemotePoolsAppended {
            chain_selector: 1,
            pool_addresses: vec![remote_address.clone()],
            previous_pool_addresses: vec![],
        });

        emit!(RemoteChainRemoved {
            chain_selector: 1,
        });

        // 7. Emit router events
        emit!(RouterUpdated {
            old_router: Pubkey::default(),
            new_router: Pubkey::default(),
        });

        // 8. Emit ownership events
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

        msg!("All events have been emitted!");
        Ok(())
    }

    /// Emit events and reverts. Will mark transaction as failed.
    pub fn trigger_all_events_reverts(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        trigger_all_events(_ctx)?;
        Err(error!(CustomError::IntentionalRevert))
    }
}

#[derive(Accounts)]
pub struct TriggerAllEvents {}

#[error_code]
pub enum CustomError {
    #[msg("Transaction intentionally reverted for testing")]
    IntentionalRevert,
}