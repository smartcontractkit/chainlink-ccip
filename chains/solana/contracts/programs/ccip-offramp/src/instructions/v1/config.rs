use crate::state::SourceChainConfig;

pub fn get_on_ramps(config: &SourceChainConfig) -> Vec<&[u8]> {
    let first_item = trim_trailing_zeros(&config.on_ramp[0]);
    let second_item = trim_trailing_zeros(&config.on_ramp[1]);

    let mut on_ramps: Vec<&[u8]> = Vec::new();

    // filter out empty values
    if !first_item.is_empty() {
        on_ramps.push(first_item);
    }
    if !second_item.is_empty() {
        on_ramps.push(second_item);
    }

    on_ramps
}

pub fn is_on_ramp_configured(config: &SourceChainConfig, on_ramp: &[u8]) -> bool {
    let valid_on_ramps = get_on_ramps(config);

    valid_on_ramps.contains(&on_ramp)
}

fn trim_trailing_zeros(input: &[u8]) -> &[u8] {
    let end = input
        .iter()
        .rposition(|&byte| byte != 0)
        .map_or(0, |pos| pos + 1);
    &input[..end]
}

#[cfg(test)]
mod tests {
    use crate::state::CodeVersion;

    use super::*;

    #[test]
    fn test_get_on_ramps() {
        struct TestCase<'a> {
            on_ramp: [[u8; 64]; 2],
            expected: Vec<&'a [u8]>,
        }

        let address_4: [u8; 64] = {
            let mut array = [0; 64];
            array[..32].copy_from_slice(&[4; 32]);
            array[32..].copy_from_slice(&[0; 32]);
            array
        };

        let cases = vec![
            TestCase {
                on_ramp: [[0; 64], [0; 64]],
                expected: vec![],
            },
            TestCase {
                on_ramp: [[1; 64], [0; 64]],
                expected: vec![&[1; 64]],
            },
            TestCase {
                on_ramp: [[0; 64], [1; 64]],
                expected: vec![&[1; 64]],
            },
            TestCase {
                on_ramp: [[2; 64], [3; 64]],
                expected: vec![&[2; 64], &[3; 64]],
            },
            TestCase {
                on_ramp: [address_4, [0; 64]],
                expected: vec![&[4; 32]],
            },
            TestCase {
                on_ramp: [[0; 64], address_4],
                expected: vec![&[4; 32]],
            },
        ];

        for case in cases {
            let config = SourceChainConfig {
                is_enabled: true,
                on_ramp: case.on_ramp,
                lane_code_version: CodeVersion::Default,
            };
            let result = get_on_ramps(&config);
            assert_eq!(
                result, case.expected,
                "Failed on get_on_ramps for config: {:?}",
                config.on_ramp
            );
        }
    }

    #[test]
    fn test_is_on_ramp_configured() {
        struct ConfigTestCase {
            config_on_ramps: [[u8; 64]; 2],
            given_on_ramp: Vec<u8>,
            expected: bool,
        }

        let address_4: [u8; 64] = {
            let mut array = [0; 64];
            array[..32].copy_from_slice(&[4; 32]);
            array[32..].copy_from_slice(&[0; 32]);
            array
        };

        let address_5: [u8; 64] = {
            let mut array = [0; 64];
            array[..32].copy_from_slice(&[5; 32]);
            array[32..].copy_from_slice(&[0; 32]);
            array
        };

        let empty_config = [[0; 64], [0; 64]];
        let one_address_config = [[1; 64], [0; 64]];
        let two_addresses_config = [[2; 64], [3; 64]];
        let two_smaller_addresses_config = [address_4, address_5];

        let cases = vec![
            // No address configured
            ConfigTestCase {
                config_on_ramps: empty_config,
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: empty_config,
                given_on_ramp: vec![1; 64],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: empty_config,
                given_on_ramp: vec![4; 64],
                expected: false,
            },
            // One address configured
            ConfigTestCase {
                config_on_ramps: one_address_config,
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: one_address_config,
                given_on_ramp: vec![1; 64],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: one_address_config,
                given_on_ramp: vec![4; 64],
                expected: false,
            },
            // Two addresses configured
            ConfigTestCase {
                config_on_ramps: two_addresses_config,
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: two_addresses_config,
                given_on_ramp: vec![2; 64],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: two_addresses_config,
                given_on_ramp: vec![3; 64],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: two_addresses_config,
                given_on_ramp: [vec![2; 32], vec![3; 32]].concat(),
                expected: false,
            },
            // Two addresses configured with smaller addresses
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config,
                given_on_ramp: vec![],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config,
                given_on_ramp: vec![0; 32],
                expected: false,
            },
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config,
                given_on_ramp: vec![4; 32],
                expected: true,
            },
            ConfigTestCase {
                config_on_ramps: two_smaller_addresses_config,
                given_on_ramp: vec![5; 32],
                expected: true,
            },
        ];

        for case in cases {
            let config = SourceChainConfig {
                is_enabled: true,
                on_ramp: case.config_on_ramps,
                lane_code_version: CodeVersion::Default,
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
    fn test_trim_trailing_zeros() {
        // only right padding is removed
        let input = vec![0, 0, 0, 1, 2, 3];
        let expected_output = &[0, 0, 0, 1, 2, 3];
        assert_eq!(trim_trailing_zeros(&input), expected_output);

        let input = vec![1, 2, 3, 0, 0, 0];
        let expected_output = &[1, 2, 3];
        assert_eq!(trim_trailing_zeros(&input), expected_output);

        let input = vec![1, 2, 3];
        let expected_output = &[1, 2, 3];
        assert_eq!(trim_trailing_zeros(&input), expected_output);

        let input = vec![0, 0, 0];
        let expected_output: &[u8] = &[];
        assert_eq!(trim_trailing_zeros(&input), expected_output);

        let input = vec![];
        let expected_output: &[u8] = &[];
        assert_eq!(trim_trailing_zeros(&input), expected_output);
    }
}
