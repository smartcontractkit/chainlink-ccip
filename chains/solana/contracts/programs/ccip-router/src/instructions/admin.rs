use anchor_lang::prelude::*;

use crate::context::{
    AcceptOwnership, AddChainSelector, AddOfframp, RemoveOfframp, TransferOwnership,
    UpdateConfigCCIPRouter, UpdateDestChainSelectorConfig, WithdrawBilledFunds,
};
use crate::state::{CodeVersion, DestChainConfig};

use super::v1;

pub fn router(default_code_version: u8) -> Result<&'static dyn Admin> {
    let default: CodeVersion = default_code_version.try_into()?;
    match default {
        CodeVersion::V1 => Ok(&v1::admin::AdminImpl),
        CodeVersion::V2 => Ok(&v1::admin::AdminImpl), // TODO this would use v2 instead (just a PoC here)
    }
}

pub trait Admin {
    fn transfer_ownership(
        &self,
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn update_fee_aggregator(
        &self,
        ctx: Context<UpdateConfigCCIPRouter>,
        fee_aggregator: Pubkey,
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
