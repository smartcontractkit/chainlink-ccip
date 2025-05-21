use anchor_lang::prelude::*;

use super::ocr3base::ocr3_set;

use crate::context::{
    AcceptOwnership, AddSourceChain, OcrPluginType, SetOcrConfig, TransferOwnership, UpdateConfig,
    UpdateReferenceAddresses, UpdateSourceChain,
};
use crate::event::admin::{
    ConfigSet, OwnershipTransferRequested, OwnershipTransferred, ReferenceAddressesSet,
    SourceChainAdded, SourceChainConfigUpdated,
};
use crate::instructions::interfaces::Admin;
use crate::state::{
    CodeVersion, Ocr3ConfigInfo, ReferenceAddresses, SourceChain, SourceChainConfig,
    SourceChainState,
};
use crate::CcipOfframpError;

pub struct Impl;
impl Admin for Impl {
    fn transfer_ownership(
        &self,
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        require!(
            proposed_owner != config.owner,
            CcipOfframpError::RedundantOwnerProposal
        );
        emit!(OwnershipTransferRequested {
            from: config.owner,
            to: proposed_owner,
        });
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        emit!(OwnershipTransferred {
            from: config.owner,
            to: config.proposed_owner,
        });
        // NOTE: take() resets proposed_owner to default
        config.owner = std::mem::take(&mut config.proposed_owner);
        Ok(())
    }

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()> {
        require_neq!(
            code_version,
            CodeVersion::Default,
            CcipOfframpError::InvalidCodeVersion
        );
        ctx.accounts.config.load_mut()?.default_code_version = code_version.into();
        Ok(())
    }

    fn update_reference_addresses(
        &self,
        ctx: Context<UpdateReferenceAddresses>,
        router: Pubkey,
        fee_quoter: Pubkey,
        offramp_lookup_table: Pubkey,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        *ctx.accounts.reference_addresses.load_mut()? = ReferenceAddresses {
            version: 1,
            router,
            fee_quoter,
            offramp_lookup_table,
            rmn_remote,
        };

        emit!(ReferenceAddressesSet {
            router,
            fee_quoter,
            offramp_lookup_table,
            rmn_remote
        });

        Ok(())
    }

    fn add_source_chain(
        &self,
        ctx: Context<AddSourceChain>,
        new_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()> {
        // Set source chain config & state
        let source_chain = &mut ctx.accounts.source_chain;
        validate_source_chain_config(new_chain_selector, &source_chain_config)?;
        source_chain.set_inner(SourceChain {
            version: 1,
            chain_selector: new_chain_selector,
            state: SourceChainState { min_seq_nr: 1 },
            config: source_chain_config.clone(),
        });

        emit!(SourceChainAdded {
            source_chain_selector: new_chain_selector,
            source_chain_config,
        });

        Ok(())
    }

    fn disable_source_chain_selector(
        &self,
        ctx: Context<UpdateSourceChain>,
        source_chain_selector: u64,
    ) -> Result<()> {
        let source_chain = &mut ctx.accounts.source_chain;

        source_chain.config.is_enabled = false;

        emit!(SourceChainConfigUpdated {
            source_chain_selector,
            source_chain_config: source_chain.config.clone(),
        });

        Ok(())
    }

    fn update_source_chain_config(
        &self,
        ctx: Context<UpdateSourceChain>,
        source_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()> {
        validate_source_chain_config(source_chain_selector, &source_chain_config)?;

        ctx.accounts.source_chain.config = source_chain_config.clone();

        emit!(SourceChainConfigUpdated {
            source_chain_selector,
            source_chain_config,
        });
        Ok(())
    }

    fn update_svm_chain_selector(
        &self,
        ctx: Context<UpdateConfig>,
        new_chain_selector: u64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;

        config.svm_chain_selector = new_chain_selector;

        emit!(ConfigSet {
            svm_chain_selector: new_chain_selector,
            enable_manual_execution_after: config.enable_manual_execution_after,
        });

        Ok(())
    }

    fn update_enable_manual_execution_after(
        &self,
        ctx: Context<UpdateConfig>,
        new_enable_manual_execution_after: i64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;

        config.enable_manual_execution_after = new_enable_manual_execution_after;

        emit!(ConfigSet {
            svm_chain_selector: config.svm_chain_selector,
            enable_manual_execution_after: new_enable_manual_execution_after,
        });

        Ok(())
    }

    fn set_ocr_config(
        &self,
        ctx: Context<SetOcrConfig>,
        plugin_type: OcrPluginType,
        config_info: Ocr3ConfigInfo,
        signers: Vec<[u8; 20]>,
        transmitters: Vec<Pubkey>,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;

        let is_commit = plugin_type == OcrPluginType::Commit;

        ocr3_set(
            &mut config.ocr3[plugin_type as usize],
            plugin_type,
            Ocr3ConfigInfo {
                config_digest: config_info.config_digest,
                f: config_info.f,
                n: signers.len() as u8,
                is_signature_verification_enabled: if is_commit { 1 } else { 0 },
            },
            signers,
            transmitters,
        )?;

        if is_commit {
            // When the OCR config changes, we reset the sequence number since it is scoped per config digest.
            // Note that s_minSeqNr/roots do not need to be reset as the roots persist
            // across reconfigurations and are de-duplicated separately.
            ctx.accounts.state.latest_price_sequence_number = 0;
        }

        Ok(())
    }
}

fn validate_source_chain_config(
    _source_chain_selector: u64,
    config: &SourceChainConfig,
) -> Result<()> {
    require!(
        !config.on_ramp.is_empty() && !config.on_ramp.is_zero(),
        CcipOfframpError::InvalidOnrampAddress
    );
    // As of now, the config has very few properties and there is little to validate yet
    // (only validity of the remote on_ramp address.) Validations will be added as that config object grows.
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::state::{OnRampAddress, SourceChainConfig};

    #[test]
    fn source_chain_config_validation() {
        let valid_config = SourceChainConfig {
            is_enabled: false,
            is_rmn_verification_disabled: true,
            lane_code_version: crate::state::CodeVersion::Default,
            on_ramp: OnRampAddress::from([0xAB; 64]),
        };

        validate_source_chain_config(1, &valid_config).unwrap();

        let invalid_config_zero_address = SourceChainConfig {
            is_enabled: false,
            is_rmn_verification_disabled: true,
            lane_code_version: crate::state::CodeVersion::Default,
            on_ramp: OnRampAddress::from([0x00; 64]),
        };

        assert_eq!(
            validate_source_chain_config(1, &invalid_config_zero_address).unwrap_err(),
            CcipOfframpError::InvalidOnrampAddress.into()
        );

        let invalid_config_empty_address = SourceChainConfig {
            is_enabled: false,
            is_rmn_verification_disabled: true,
            lane_code_version: crate::state::CodeVersion::Default,
            on_ramp: OnRampAddress::EMPTY,
        };

        assert_eq!(
            validate_source_chain_config(1, &invalid_config_empty_address).unwrap_err(),
            CcipOfframpError::InvalidOnrampAddress.into()
        );
    }
}
