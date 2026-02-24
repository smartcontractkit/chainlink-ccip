use std::cell::Ref;

use anchor_lang::prelude::*;
use anchor_lang::Discriminator;

use crate::context::ANCHOR_DISCRIMINATOR;
use crate::state::{CodeVersion, Config};
use crate::RmnRemoteError;

#[derive(AnchorDeserialize, InitSpace, Debug)]
pub(super) struct ConfigV1 {
    pub version: u8,
    pub owner: Pubkey,

    pub proposed_owner: Pubkey,
    pub default_code_version: CodeVersion,
}

impl TryFrom<ConfigV1> for Config {
    type Error = anchor_lang::error::Error;

    fn try_from(v1: ConfigV1) -> std::result::Result<Config, Self::Error> {
        require_eq!(
            v1.version,
            1, // this deserialization is only valid for v1
            RmnRemoteError::InvalidInputsConfigAccount
        );

        Ok(Config {
            version: 1,
            owner: v1.owner,
            proposed_owner: v1.proposed_owner,
            default_code_version: v1.default_code_version,
            event_authorities: vec![], // this is not part of the v1 data, it defaults to empty
        })
    }
}

pub(super) fn load_config(config: &AccountInfo<'_>) -> Result<Config> {
    let borrowed_data = config.try_borrow_data()?;
    let (discriminator, data) = borrowed_data.split_at(ANCHOR_DISCRIMINATOR);

    require!(
        Config::DISCRIMINATOR == discriminator,
        RmnRemoteError::InvalidInputsConfigAccount
    );

    let version_byte = data[0]; // version is the first byte of the data (after the discriminator)

    const V1_SPACE: usize = ANCHOR_DISCRIMINATOR + ConfigV1::INIT_SPACE;
    if config.data_len() == V1_SPACE && version_byte == 1 {
        let config_v1 = load_config_v1_unchecked(borrowed_data)?;
        return Config::try_from(config_v1);
    }

    // Use >= for the size because the event_authorities vec makes the size variable
    const V2_MIN_SPACE: usize = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE;
    if config.data_len() >= V2_MIN_SPACE && version_byte == 2 {
        let mut data: &[u8] = &borrowed_data;
        return Config::try_deserialize(&mut data);
    }

    Err(RmnRemoteError::InvalidInputsConfigAccount.into())
}

pub(super) fn load_config_v1_unchecked(borrowed_data: Ref<&mut [u8]>) -> Result<ConfigV1> {
    let (_discriminator, data) = borrowed_data.split_at(ANCHOR_DISCRIMINATOR);
    let config_v1 = ConfigV1::deserialize(&mut &data[..])
        .map_err(|_| RmnRemoteError::InvalidInputsConfigAccount)?;
    Ok(config_v1)
}
