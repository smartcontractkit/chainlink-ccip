use crate::context::{
    get_message_transmitter_pda, get_token_messenger_minter_pda, Empty, MESSAGE_SENT_EVENT_SEED,
};
use anchor_lang::prelude::*;
use base_token_pool::common::{
    CcipAccountMeta, CcipTokenPoolError, DeriveAccountsResponse, LockOrBurnInV1, ReleaseOrMintInV1,
    ToMeta, POOL_CHAINCONFIG_SEED,
};
use core::fmt;
use std::{
    fmt::{Display, Formatter},
    str::FromStr,
};

use crate::{
    context::{TokenOfframpRemainingAccounts, TOKEN_MESSENGER_MINTER},
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
    use crate::MessageAndAttestation;

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
                get_message_transmitter_pda(&[
                    b"message_transmitter_authority",
                    TOKEN_MESSENGER_MINTER.as_ref(),
                ])
                .readonly(),
            ],
            current_stage: OfframpDeriveStage::RetrieveChainConfig.to_string(),
            next_stage: OfframpDeriveStage::BuildDynamicAccounts.to_string(),
            ..Default::default()
        })
    }

    pub fn build_dynamic_accounts<'info>(
        ctx: Context<'_, '_, 'info, 'info, Empty>,
        release_or_mint: &ReleaseOrMintInV1,
    ) -> Result<DeriveAccountsResponse> {
        let chain_config = Account::<'info, ChainConfig>::try_from(&ctx.remaining_accounts[0])?;
        let domain_id = chain_config.cctp.domain_id;
        let mint = release_or_mint.local_token;
        let cctp_msg =
            MessageAndAttestation::try_from_slice(&release_or_mint.offchain_token_data)?.message;
        let nonce_seed = TokenOfframpRemainingAccounts::first_nonce_seed(&cctp_msg);
        let domain_str = domain_id.to_string();
        let domain_seed = domain_str.as_bytes();
        let remote_token_address_bytes =
            to_solana_pubkey(&chain_config.base.remote.token_address).to_bytes();

        Ok(DeriveAccountsResponse {
            accounts_to_save: vec![
                // cctp_event_authority
                get_message_transmitter_pda(&[b"__event_authority"]).readonly(),
                // cctp_custody_token_account
                get_token_messenger_minter_pda(&[b"custody", mint.as_ref()]).writable(),
                // cctp_remote_token_messenger_key
                get_token_messenger_minter_pda(&[b"remote_token_messenger", domain_seed])
                    .readonly(),
                // cctp_token_pair
                get_token_messenger_minter_pda(&[
                    b"token_pair",
                    domain_seed,
                    &remote_token_address_bytes,
                ])
                .readonly(),
                // cctp_used_nonces
                get_message_transmitter_pda(&[b"used_nonces", domain_seed, nonce_seed.as_ref()])
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

    pub fn retrieve_chain_config(lock_or_burn: &LockOrBurnInV1) -> Result<DeriveAccountsResponse> {
        Ok(DeriveAccountsResponse {
            ask_again_with: vec![find(
                &[
                    POOL_CHAINCONFIG_SEED,
                    &lock_or_burn.remote_chain_selector.to_le_bytes(),
                    lock_or_burn.local_token.as_ref(),
                ],
                crate::ID,
            )],
            // The static PDAs have mostly already been returned by CCIP via the LUT, so we just return here the ones not shared with offramp (so not in LUT)
            accounts_to_save: vec![
                get_token_messenger_minter_pda(&[b"sender_authority"]).readonly()
            ],
            current_stage: OnrampDeriveStage::RetrieveChainConfig.to_string(),
            next_stage: OnrampDeriveStage::BuildDynamicAccounts.to_string(),
            ..Default::default()
        })
    }

    pub fn build_dynamic_accounts<'info>(
        ctx: Context<'_, '_, 'info, 'info, Empty>,
        lock_or_burn: &LockOrBurnInV1,
    ) -> Result<DeriveAccountsResponse> {
        let chain_config = Account::<'info, ChainConfig>::try_from(&ctx.remaining_accounts[0])?;
        let domain_id = chain_config.cctp.domain_id;
        let domain_str = domain_id.to_string();
        let domain_seed = domain_str.as_bytes();

        msg!(
            "Sender: {:?}, selector: {:?}, nonce: {:?}",
            lock_or_burn.original_sender,
            lock_or_burn.remote_chain_selector,
            lock_or_burn.msg_total_nonce
        );

        Ok(DeriveAccountsResponse {
            accounts_to_save: vec![
                // cctp_remote_token_messenger_key
                get_token_messenger_minter_pda(&[b"remote_token_messenger", domain_seed])
                    .readonly(),
                // cctp_message_sent_event
                find(
                    &[
                        MESSAGE_SENT_EVENT_SEED,
                        &lock_or_burn.original_sender.to_bytes(),
                        &lock_or_burn.remote_chain_selector.to_le_bytes(),
                        &lock_or_burn.msg_total_nonce.to_le_bytes(),
                    ],
                    crate::ID,
                )
                .writable(),
            ],
            current_stage: OnrampDeriveStage::BuildDynamicAccounts.to_string(),
            next_stage: "".to_string(),
            ..Default::default()
        })
    }
}
