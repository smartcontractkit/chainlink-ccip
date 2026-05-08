use std::cell::Ref;

use anchor_lang::prelude::*;
use anchor_lang::Discriminator;

use crate::state::{CodeVersion, Config};
use crate::RmnRemoteError;

#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
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

impl From<ConfigV2> for Config {
    fn from(v2: ConfigV2) -> Config {
        Config {
            version: 2,
            owner: v2.owner,
            proposed_owner: v2.proposed_owner,
            default_code_version: v2.default_code_version,
            event_authorities: v2.event_authorities,
            curser: Pubkey::default(), // added in v3, defaults to empty (unset)
            bump: 0,                   // added in v3, defaults to 0
        }
    }
}

pub(super) fn load_config(config: &AccountInfo<'_>) -> Result<Config> {
    let borrowed_data = config.try_borrow_data()?;

    // version is the first byte of the data after the discriminator
    let version_byte = borrowed_data[Config::DISCRIMINATOR.len()];

    match version_byte {
        2 => {
            // this assumes a single entry in the event_authorities vec, which is true at the time of this migration,
            // which will only be executed once.
            const V2_SPACE: usize = Config::DISCRIMINATOR.len() + ConfigV2::INIT_SPACE;

            require_eq!(
                config.data_len(),
                V2_SPACE,
                RmnRemoteError::InvalidInputsConfigAccount
            );
            let old_config = load_old_config::<ConfigV2>(borrowed_data)?;
            return Ok(Config::from(old_config));
        }
        Config::LATEST_VERSION => {
            const MIN_SPACE: usize = Config::DISCRIMINATOR.len() + Config::INIT_SPACE;

            require_gte!(
                config.data_len(),
                MIN_SPACE,
                RmnRemoteError::InvalidInputsConfigAccount
            );
            let mut data: &[u8] = &borrowed_data;
            return Config::try_deserialize(&mut data);
        }
        _ => Err(RmnRemoteError::InvalidInputsConfigAccount.into()),
    }
}

pub(super) fn load_old_config<T>(borrowed_data: Ref<&mut [u8]>) -> Result<T>
where
    T: AnchorDeserialize,
    Config: From<T>,
{
    let (discriminator, data) = borrowed_data.split_at(Config::DISCRIMINATOR.len());
    require!(
        Config::DISCRIMINATOR == discriminator,
        RmnRemoteError::InvalidInputsConfigAccount
    );
    let old =
        T::deserialize(&mut &data[..]).map_err(|_| RmnRemoteError::InvalidInputsConfigAccount)?;
    Ok(old)
}
