use crate::context::Empty;
use anchor_lang::prelude::*;
use base_token_pool::common::POOL_CHAINCONFIG_SEED;
use base_token_pool::common::{
    CcipAccountMeta, CcipTokenPoolError, DeriveAccountsResponse, ReleaseOrMintInV1, ToMeta,
};
use core::fmt;
use std::{
    fmt::{Display, Formatter},
    str::FromStr,
};

use crate::{
    context::{TokenOfframpRemainingAccounts, MESSAGE_TRANSMITTER, TOKEN_MESSENGER_MINTER},
    to_solana_pubkey, ChainConfig,
};

// Local helper to find a readonly CCIP meta for a given seed + program_id combo.
// Short name for compactness.
fn find(seeds: &[&[u8]], program_id: Pubkey) -> CcipAccountMeta {
    Pubkey::find_program_address(seeds, &program_id)
        .0
        .readonly()
}

pub mod release_or_mint {

    use super::*;

    #[derive(Clone, Debug)]
    pub enum OfframpDeriveStage {
        RetrieveChainConfig,
        BuildDynamicAccounts,
    }

    impl Display for OfframpDeriveStage {
        fn fmt(&self, f: &mut Formatter) -> fmt::Result {
            match self {
                OfframpDeriveStage::RetrieveChainConfig => f.write_str("RetrieveChainConfig"),
                OfframpDeriveStage::BuildDynamicAccounts => f.write_str("BuildDynamicAccounts"),
            }
        }
    }

    impl FromStr for OfframpDeriveStage {
        type Err = CcipTokenPoolError;

        fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
            match s {
                "Start" | "RetrieveChainConfig" => Ok(Self::RetrieveChainConfig),
                "BuildDynamicAccounts" => Ok(Self::BuildDynamicAccounts),
                _ => Err(CcipTokenPoolError::InvalidDerivationStage),
            }
        }
    }

    pub fn retrieve_chain_config(
        release_or_mint: &ReleaseOrMintInV1,
    ) -> Result<DeriveAccountsResponse> {
        Ok(DeriveAccountsResponse {
            ask_again_with: vec![find(
                &[
                    POOL_CHAINCONFIG_SEED,
                    &release_or_mint.remote_chain_selector.to_le_bytes(),
                    release_or_mint.local_token.as_ref(),
                ],
                crate::ID,
            )],
            // We don't need the domain for the first few PDAs, so we return them now to keep
            // return sizes balanced.
            accounts_to_save: vec![
                // cctp_authority_pda
                TokenOfframpRemainingAccounts::get_message_transmitter_pda(&[
                    b"message_transmitter_authority",
                    TOKEN_MESSENGER_MINTER.as_ref(),
                ])
                .readonly(),
                // cctp_message_transmitter_account
                TokenOfframpRemainingAccounts::get_message_transmitter_pda(&[
                    b"message_transmitter",
                ])
                .readonly(),
                // cctp_token_messenger_minter
                TOKEN_MESSENGER_MINTER.readonly(),
                // system_program
                System::id().readonly(),
                // cctp_event_authority
                TokenOfframpRemainingAccounts::get_message_transmitter_pda(&[b"__event_authority"])
                    .readonly(),
                // cctp_message_transmitter
                MESSAGE_TRANSMITTER.readonly(),
                // cctp_token_messenger_account
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[
                    b"token_messenger",
                ])
                .readonly(),
                // cctp_token_minter_account
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[b"token_minter"])
                    .writable(),
            ],
            current_stage: OfframpDeriveStage::RetrieveChainConfig.to_string(),
            next_stage: OfframpDeriveStage::BuildDynamicAccounts.to_string(),
            ..Default::default()
        })
    }

    pub fn build_dynamic_accounts<'info>(
        ctx: Context<'_, '_, 'info, 'info, Empty<'info>>,
        release_or_mint: &ReleaseOrMintInV1,
    ) -> Result<DeriveAccountsResponse> {
        let chain_config = Account::<'info, ChainConfig>::try_from(&ctx.remaining_accounts[0])?;
        let domain_id = chain_config.cctp.domain_id;
        let mint = release_or_mint.local_token;
        let nonce_seed =
            TokenOfframpRemainingAccounts::first_nonce_seed(&release_or_mint.offchain_token_data)?;
        let domain_str = domain_id.to_string();
        let domain_seed = domain_str.as_bytes();
        let remote_token_address_bytes =
            to_solana_pubkey(&chain_config.base.remote.token_address).to_bytes();

        Ok(DeriveAccountsResponse {
            accounts_to_save: vec![
                // cctp_local_token
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[
                    b"local_token",
                    mint.as_ref(),
                ])
                .writable(),
                // cctp_custody_token_account
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[
                    b"custody",
                    mint.as_ref(),
                ])
                .writable(),
                // cctp_token_messenger_event_authority
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[
                    b"__event_authority",
                ])
                .readonly(),
                // cctp_remote_token_messenger_key
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[
                    b"remote_token_messenger",
                    domain_seed,
                ])
                .readonly(),
                // cctp_token_pair
                TokenOfframpRemainingAccounts::get_token_messenger_minter_pda(&[
                    b"token_pair",
                    domain_seed,
                    &remote_token_address_bytes,
                ])
                .readonly(),
                // cctp_used_nonces
                TokenOfframpRemainingAccounts::get_message_transmitter_pda(&[
                    b"used_nonces",
                    domain_seed,
                    nonce_seed.as_ref(),
                ])
                .writable(),
            ],
            current_stage: OfframpDeriveStage::BuildDynamicAccounts.to_string(),
            ..Default::default()
        })
    }
}

pub mod lock_or_burn {
    use super::*;

    #[derive(Clone, Debug)]
    pub enum OnrampDeriveStage {
        RetrieveChainConfig,
        BuildDynamicAccounts,
    }

    impl Display for OnrampDeriveStage {
        fn fmt(&self, f: &mut Formatter) -> fmt::Result {
            match self {
                OnrampDeriveStage::RetrieveChainConfig => f.write_str("RetrieveChainConfig"),
                OnrampDeriveStage::BuildDynamicAccounts => f.write_str("BuildDynamicAccounts"),
            }
        }
    }

    impl FromStr for OnrampDeriveStage {
        type Err = CcipTokenPoolError;

        fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
            match s {
                "Start" | "RetrieveChainConfig" => Ok(Self::RetrieveChainConfig),
                "BuildDynamicAccounts" => Ok(Self::BuildDynamicAccounts),
                _ => Err(CcipTokenPoolError::InvalidDerivationStage),
            }
        }
    }
}
