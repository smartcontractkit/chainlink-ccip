use std::cell::Ref;

use anchor_lang::prelude::*;
use anchor_lang::Discriminator;

use crate::state::{CodeVersion, Config};
use crate::RmnRemoteError;

#[derive(AnchorDeserialize, InitSpace, Debug)]
pub struct ConfigV2 {
    // --- v1 fields ---
    pub version: u8,
    pub owner: Pubkey,

    pub proposed_owner: Pubkey,
    pub default_code_version: CodeVersion,

    // --- v2 fields ---
    // Using max_len for INIT_SPACE calculation. At the time of this migration, there is a single entry here
    #[max_len(1)]
    pub event_authorities: Vec<Pubkey>,
}

impl TryFrom<ConfigV2> for Config {
    type Error = anchor_lang::error::Error;

    fn try_from(v2: ConfigV2) -> std::result::Result<Config, Self::Error> {
        require_eq!(
            v2.version,
            2, // this deserialization is only valid for v2
            RmnRemoteError::InvalidInputsConfigAccount
        );

        Ok(Config {
            version: 2,
            owner: v2.owner,
            proposed_owner: v2.proposed_owner,
            default_code_version: v2.default_code_version,
            event_authorities: v2.event_authorities,
            curser: Pubkey::default(), // added in v3, defaults to empty (unset)
        })
    }
}

pub(super) fn load_config(config: &AccountInfo<'_>) -> Result<Config> {
    let borrowed_data = config.try_borrow_data()?;
    let (discriminator, data) = borrowed_data.split_at(Config::DISCRIMINATOR.len());

    require!(
        Config::DISCRIMINATOR == discriminator,
        RmnRemoteError::InvalidInputsConfigAccount
    );

    let version_byte = data[0]; // version is the first byte of the data (after the discriminator)

    // this assumes a single entry in the event_authorities vec, which is true at the time of this migration,
    // which will only be executed once.
    const V2_SPACE: usize = Config::DISCRIMINATOR.len() + ConfigV2::INIT_SPACE;
    if version_byte == 2 && config.data_len() == V2_SPACE {
        let config_v2 = load_config_v2_unchecked(borrowed_data)?;
        return Config::try_from(config_v2);
    }

    const V3_MIN_SPACE: usize = Config::DISCRIMINATOR.len() + Config::INIT_SPACE;
    if version_byte == 3 && config.data_len() >= V3_MIN_SPACE {
        let mut data: &[u8] = &borrowed_data;
        return Config::try_deserialize(&mut data);
    }

    Err(RmnRemoteError::InvalidInputsConfigAccount.into())
}

pub(super) fn load_config_v2_unchecked(borrowed_data: Ref<&mut [u8]>) -> Result<ConfigV2> {
    let (_discriminator, data) = borrowed_data.split_at(Config::DISCRIMINATOR.len());
    let config_v2 = ConfigV2::deserialize(&mut &data[..])
        .map_err(|_| RmnRemoteError::InvalidInputsConfigAccount)?;
    Ok(config_v2)
}
