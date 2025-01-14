use anchor_lang::{error::Error, require};

use crate::{CcipRouterError, SourceChainConfig};

pub fn is_valid_on_ramp(config: &SourceChainConfig) -> bool {
    let len = config.on_ramp.len();

    len == 0 || len == 64 || len == 128
}

pub fn get_on_ramps(config: &SourceChainConfig) -> Result<Vec<[u8; 64]>, Error> {
    require!(is_valid_on_ramp(config), CcipRouterError::InvalidInputs);

    let valid_on_ramps: Vec<[u8; 64]> = match config.on_ramp.len() {
        0 => Vec::new(),
        64 => vec![config.on_ramp[0..64]
            .try_into()
            .map_err(|_| CcipRouterError::InvalidInputs)?],
        128 => vec![
            config.on_ramp[0..64]
                .try_into()
                .map_err(|_| CcipRouterError::InvalidInputs)?,
            config.on_ramp[64..128]
                .try_into()
                .map_err(|_| CcipRouterError::InvalidInputs)?,
        ],
        _ => return Err(CcipRouterError::InvalidInputs.into()),
    };
    Ok(valid_on_ramps)
}

pub fn is_on_ramp_configured(config: &SourceChainConfig, on_ramp: &[u8]) -> bool {
    let mut on_ramp_bytes = [0u8; 64];
    let len = on_ramp.len().min(64);
    on_ramp_bytes[..len].copy_from_slice(&on_ramp[..len]);

    let valid_on_ramps = get_on_ramps(config);
    if valid_on_ramps.is_err() {
        return false;
    }
    valid_on_ramps.unwrap().contains(&on_ramp_bytes)
}

#[allow(dead_code)]
fn trim_leading_zeros(input: &[u8]) -> &[u8] {
    let start = input
        .iter()
        .position(|&byte| byte != 0)
        .unwrap_or(input.len());
    &input[start..]
}

#[cfg(test)]
mod tests {
    use crate::CcipRouterError;

    use super::*;

    #[test]
    fn test_is_valid_on_ramp() {
        struct SourceChainTestCase {
            config: SourceChainConfig,
            expected: bool,
        }

        let cases = vec![
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![],
                    is_enabled: true,
                },
                expected: true,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 64],
                    is_enabled: true,
                },
                expected: true,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 128],
                    is_enabled: true,
                },
                expected: true,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 1],
                    is_enabled: true,
                },
                expected: false,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 63],
                    is_enabled: true,
                },
                expected: false,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 65],
                    is_enabled: true,
                },
                expected: false,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 127],
                    is_enabled: true,
                },
                expected: false,
            },
            SourceChainTestCase {
                config: SourceChainConfig {
                    on_ramp: vec![0; 129],
                    is_enabled: true,
                },
                expected: false,
            },
        ];

        for case in cases {
            let result = is_valid_on_ramp(&case.config);
            assert_eq!(
                result,
                case.expected,
                "Failed for on_ramp length: {}",
                case.config.on_ramp.len()
            );
        }
    }

    #[test]
    fn test_get_on_ramps() {
        struct TestCase {
            on_ramp: Vec<u8>,
            expected: Result<Vec<[u8; 64]>, Error>,
        }

        let cases = vec![
            TestCase {
                on_ramp: vec![],
                expected: Ok(Vec::new()),
            },
            TestCase {
                on_ramp: vec![1; 64],
                expected: Ok(vec![[1; 64]]),
            },
            TestCase {
                on_ramp: [vec![2; 64], vec![3; 64]].concat(),
                expected: Ok(vec![[2; 64], [3; 64]]),
            },
            TestCase {
                on_ramp: vec![3; 1],
                expected: Err(CcipRouterError::InvalidInputs.into()),
            },
            TestCase {
                on_ramp: vec![4; 63],
                expected: Err(CcipRouterError::InvalidInputs.into()),
            },
            TestCase {
                on_ramp: vec![5; 65],
                expected: Err(CcipRouterError::InvalidInputs.into()),
            },
            TestCase {
                on_ramp: vec![6; 127],
                expected: Err(CcipRouterError::InvalidInputs.into()),
            },
            TestCase {
                on_ramp: vec![7; 129],
                expected: Err(CcipRouterError::InvalidInputs.into()),
            },
        ];

        for case in cases {
            let config = SourceChainConfig {
                is_enabled: true,
                on_ramp: case.on_ramp,
            };
            let result = get_on_ramps(&config);
            assert_eq!(
                result,
                case.expected,
                "Failed for on_ramp length: {}",
                config.on_ramp.len()
            );
        }
    }

    #[test]
    fn test_is_on_ramp_configured() {
        struct ConfigTestCase {
            config_on_ramps: Vec<u8>,
            given_on_ramp: Vec<u8>,
            expected: bool,
        }

        let empty_config = vec![];
        let one_address_config = vec![1; 64];
        let two_addresses_config = [vec![2; 64], vec![3; 64]].concat();
        let two_smaller_addresses_config =
            [vec![4; 32], vec![0; 32], vec![5; 32], vec![0; 32]].concat();

        let cases = vec![
            // No address configured
            ConfigTestCase {
                config_on_ramps: empty_config.clone(),
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: empty_config.clone(),
                given_on_ramp: vec![1; 64],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: empty_config.clone(),
                given_on_ramp: vec![4; 64],
                expected: false,
            },
            // One address configured
            ConfigTestCase {
                config_on_ramps: one_address_config.clone(),
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: one_address_config.clone(),
                given_on_ramp: vec![1; 64],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: one_address_config.clone(),
                given_on_ramp: vec![4; 64],
                expected: false,
            },
            // Two addresses configured
            ConfigTestCase {
                config_on_ramps: two_addresses_config.clone(),
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: two_addresses_config.clone(),
                given_on_ramp: vec![2; 64],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: two_addresses_config.clone(),
                given_on_ramp: vec![3; 64],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: two_addresses_config.clone(),
                given_on_ramp: [vec![2; 32], vec![3; 32]].concat(),
                expected: false,
            },
            // Two addresses configured with smaller addresses
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config.clone(),
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config.clone(),
                given_on_ramp: vec![0; 32],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config.clone(),
                given_on_ramp: vec![4; 32],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config.clone(),
                given_on_ramp: vec![5; 32],
                expected: true,
            },
        ];

        for case in cases {
            let config = SourceChainConfig {
                is_enabled: true,
                on_ramp: case.config_on_ramps,
            };
            let result = is_on_ramp_configured(&config, &case.given_on_ramp);
            assert_eq!(
                result, case.expected,
                "Failed for on_ramp: {:?}",
                case.given_on_ramp
            );
        }
    }

    #[test]
    fn test_trim_leading_zeros() {
        let input = vec![0, 0, 0, 1, 2, 3];
        let expected_output = &[1, 2, 3];
        assert_eq!(trim_leading_zeros(&input), expected_output);

        let input = vec![1, 2, 3];
        let expected_output = &[1, 2, 3];
        assert_eq!(trim_leading_zeros(&input), expected_output);

        let input = vec![0, 0, 0];
        let expected_output: &[u8] = &[];
        assert_eq!(trim_leading_zeros(&input), expected_output);

        let input = vec![];
        let expected_output: &[u8] = &[];
        assert_eq!(trim_leading_zeros(&input), expected_output);
    }
}
