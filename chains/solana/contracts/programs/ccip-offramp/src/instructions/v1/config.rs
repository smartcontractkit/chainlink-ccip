use crate::{state::SourceChainConfig, OnRampAddress};

pub fn get_on_ramps(config: &SourceChainConfig) -> impl Iterator<Item = &OnRampAddress> {
    config.on_ramp.iter().filter(|o| !o.is_empty())
}

pub fn is_on_ramp_configured(config: &SourceChainConfig, on_ramp: &[u8]) -> bool {
    get_on_ramps(config).any(|o| o.bytes() == on_ramp)
}

#[cfg(test)]
mod tests {
    use crate::state::CodeVersion;

    use super::*;

    #[test]
    fn test_is_on_ramp_configured() {
        let sample_addresses: &[OnRampAddress] = &[
            [0x1; 64].into(),
            [0x2, 32].into(),
            vec![0x0; 64].try_into().unwrap(),
        ];
        let mut config = SourceChainConfig {
            is_enabled: true,
            on_ramp: [OnRampAddress::EMPTY, OnRampAddress::EMPTY],
            lane_code_version: CodeVersion::Default,
            is_rmn_verification_disabled: true,
        };

        assert_eq!(get_on_ramps(&config).count(), 0);

        for address in sample_addresses {
            assert!(!is_on_ramp_configured(&config, address.bytes()));
            config.on_ramp[0] = address.clone();
            assert!(is_on_ramp_configured(&config, address.bytes()));
            config.on_ramp[0] = OnRampAddress::EMPTY;
            config.on_ramp[1] = address.clone();
            assert!(is_on_ramp_configured(&config, address.bytes()));
        }
    }
}
