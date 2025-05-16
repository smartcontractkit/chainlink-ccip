use anchor_lang::prelude::*;

use crate::context::{
    AcceptOwnership, AddChainSelector, AddOfframp, CcipSend, RemoveOfframp, TransferOwnership,
    UpdateConfigCCIPRouter, UpdateDestChainSelectorConfig, UpdateDestChainSelectorConfigNoRealloc,
    WithdrawBilledFunds,
};
use crate::messages::{GetFeeResult, SVM2AnyMessage};
use crate::state::{CodeVersion, DestChainConfig};
use crate::token_context::{
    AcceptAdminRoleTokenAdminRegistry, ModifyTokenAdminRegistry,
    OverridePendingTokenAdminRegistryByCCIPAdmin, OverridePendingTokenAdminRegistryByOwner,
    RegisterTokenAdminRegistryByCCIPAdmin, RegisterTokenAdminRegistryByOwner,
    SetPoolTokenAdminRegistry,
};
use crate::GetFee;

pub trait Admin {
    fn transfer_ownership(
        &self,
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        code_version: CodeVersion,
    ) -> Result<()>;

    fn set_link_token_mint(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        link_token_mint: Pubkey,
    ) -> Result<()>;

    fn update_fee_aggregator(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        fee_aggregator: Pubkey,
    ) -> Result<()>;

    fn update_rmn_remote(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        rmn_remote: Pubkey,
    ) -> Result<()>;

    fn add_chain_selector(
        &self,
        ctx: Context<AddChainSelector>,
        new_chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()>;

    fn update_dest_chain_config(
        &self,
        ctx: Context<UpdateDestChainSelectorConfig>,
        dest_chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()>;

    fn bump_ccip_version_for_dest_chain(
        &self,
        ctx: Context<UpdateDestChainSelectorConfigNoRealloc>,
        dest_chain_selector: u64,
    ) -> Result<()>;

    fn rollback_ccip_version_for_dest_chain(
        &self,
        ctx: Context<UpdateDestChainSelectorConfigNoRealloc>,
        dest_chain_selector: u64,
    ) -> Result<()>;

    fn add_offramp(
        &self,
        _ctx: Context<AddOfframp>,
        source_chain_selector: u64,
        offramp: Pubkey,
    ) -> Result<()>;

    fn remove_offramp(
        &self,
        _ctx: Context<RemoveOfframp>,
        source_chain_selector: u64,
        offramp: Pubkey,
    ) -> Result<()>;

    fn update_svm_chain_selector(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        new_chain_selector: u64,
    ) -> Result<()>;

    fn withdraw_billed_funds(
        &self,
        ctx: Context<WithdrawBilledFunds>,
        transfer_all: bool,
        desired_amount: u64, // if transfer_all is false, this value must be 0
    ) -> Result<()>;
}

pub trait OnRamp {
    fn ccip_send<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, CcipSend<'info>>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
        token_indexes: Vec<u8>,
    ) -> Result<[u8; 32]>;

    fn get_fee<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, GetFee<'info>>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
    ) -> Result<GetFeeResult>;
}

pub trait TokenAdminRegistry {
    /// Proposes an administrator for the given token as pending administrator
    fn ccip_admin_propose_administrator(
        &self,
        ctx: Context<RegisterTokenAdminRegistryByCCIPAdmin>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()>;

    /// Overrides the pending administrator for the given token
    fn ccip_admin_override_pending_administrator(
        &self,
        ctx: Context<OverridePendingTokenAdminRegistryByCCIPAdmin>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()>;

    /// Proposes an administrator for the given token as pending administrator
    fn owner_propose_administrator(
        &self,
        ctx: Context<RegisterTokenAdminRegistryByOwner>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()>;

    /// Overrides the pending administrator for the given token
    fn owner_override_pending_administrator(
        &self,
        ctx: Context<OverridePendingTokenAdminRegistryByOwner>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()>;

    /// Transfers the administrator role for a token to a new address with a 2-step process
    fn transfer_admin_role_token_admin_registry(
        &self,
        ctx: Context<ModifyTokenAdminRegistry>,
        new_admin: Pubkey,
    ) -> Result<()>;

    /// Accepts the admin role for a token
    fn accept_admin_role_token_admin_registry(
        &self,
        ctx: Context<AcceptAdminRoleTokenAdminRegistry>,
    ) -> Result<()>;

    /// Sets the lookup table for pool for a token.
    /// Setting the lookup table to address(0) effectively delists the token from CCIP.
    /// Setting the lookup table to any other address enables the token on CCIP.
    fn set_pool(
        &self,
        ctx: Context<SetPoolTokenAdminRegistry>,
        writable_indexes: Vec<u8>,
    ) -> Result<()>;
}
