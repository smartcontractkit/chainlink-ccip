use anchor_lang::prelude::*;

use crate::context::{
    AcceptOwnership, AddBillingTokenConfig, AddDestChain, AddPriceUpdater, GasPriceUpdate, GetFee,
    RemovePriceUpdater, SetTokenTransferFeeConfig, TokenPriceUpdate, UpdateBillingTokenConfig,
    UpdateConfig, UpdateConfigLinkMint, UpdateDestChainConfig, UpdatePrices,
};
use crate::messages::{GetFeeResult, SVM2AnyMessage};
use crate::state::{BillingTokenConfig, CodeVersion, DestChainConfig, TokenTransferFeeConfig};

pub trait Public {
    fn get_fee<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, GetFee>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
    ) -> Result<GetFeeResult>;
}

pub trait Prices {
    fn update_prices<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, UpdatePrices<'info>>,
        token_updates: Vec<TokenPriceUpdate>,
        gas_updates: Vec<GasPriceUpdate>,
    ) -> Result<()>;
}

pub trait Admin {
    fn transfer_ownership(&self, ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()>;

    fn set_max_fee_juels_per_msg(
        &self,
        ctx: Context<UpdateConfig>,
        max_fee_juels_per_msg: u128,
    ) -> Result<()>;

    fn set_link_token_mint(&self, ctx: Context<UpdateConfigLinkMint>) -> Result<()>;

    fn add_billing_token_config(
        &self,
        ctx: Context<AddBillingTokenConfig>,
        config: BillingTokenConfig,
    ) -> Result<()>;

    fn update_billing_token_config(
        &self,
        ctx: Context<UpdateBillingTokenConfig>,
        config: BillingTokenConfig,
    ) -> Result<()>;

    fn add_dest_chain(
        &self,
        ctx: Context<AddDestChain>,
        chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()>;

    fn disable_dest_chain(
        &self,
        ctx: Context<UpdateDestChainConfig>,
        chain_selector: u64,
    ) -> Result<()>;

    fn update_dest_chain_config(
        &self,
        ctx: Context<UpdateDestChainConfig>,
        chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()>;

    fn add_price_updater(
        &self,
        _ctx: Context<AddPriceUpdater>,
        price_updater: Pubkey,
    ) -> Result<()>;

    fn remove_price_updater(
        &self,
        _ctx: Context<RemovePriceUpdater>,
        price_updater: Pubkey,
    ) -> Result<()>;

    fn set_token_transfer_fee_config(
        &self,
        ctx: Context<SetTokenTransferFeeConfig>,
        chain_selector: u64,
        mint: Pubkey,
        cfg: TokenTransferFeeConfig,
    ) -> Result<()>;
}
