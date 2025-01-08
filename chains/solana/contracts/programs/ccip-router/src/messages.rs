use crate::{BillingTokenConfig, CcipRouterError, DestChain, ReportContext};
use anchor_lang::prelude::*;
use ethnum::U256;

// TODO this file has to be broken up

use crate::ocr3base::Ocr3Report;

pub const CHAIN_FAMILY_SELECTOR_EVM: u32 = 0x2812d52c;

const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
// Family-agnostic header for OnRamp & OffRamp messages.
// The messageId is not expected to match hash(message), since it may originate from another ramp family
pub struct RampMessageHeader {
    pub message_id: [u8; 32], // Unique identifier for the message, generated with the source chain's encoding scheme
    pub source_chain_selector: u64, // the chain selector of the source chain, note: not chainId
    pub dest_chain_selector: u64, // the chain selector of the destination chain, note: not chainId
    pub sequence_number: u64, // sequence number, not unique across lanes
    pub nonce: u64, // nonce for this lane for this sender, not unique across senders/lanes
}

impl RampMessageHeader {
    pub fn len(&self) -> usize {
        32 // message_id
        + 8 // source_chain
        + 8 // dest_chain
        + 8 // sequence
        + 8 // nonce
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
/// Report that is submitted by the execution DON at the execution phase. (including chain selector data)
pub struct ExecutionReportSingleChain {
    pub source_chain_selector: u64,
    pub message: Any2SolanaRampMessage,
    pub offchain_token_data: Vec<Vec<u8>>, // https://github.com/smartcontractkit/chainlink/blob/885baff9479e935e0fc34d9f52214a32c158eac5/contracts/src/v0.8/ccip/libraries/Internal.sol#L72
    pub root: [u8; 32],
    pub proofs: Vec<[u8; 32]>,

    // NOT HASHED
    pub token_indexes: Vec<u8>, // outside of message because this is not available during commit stage
}

impl Ocr3Report for ExecutionReportSingleChain {
    fn hash(&self, _: &ReportContext) -> [u8; 32] {
        [0; 32] // not needed, this report is not hashed for signing
    }
    fn len(&self) -> usize {
        let offchain_token_data_len = self
            .offchain_token_data
            .iter()
            .fold(0, |acc, e| acc + 4 + e.len());

        8 // source chain selector
        + self.message.len() // ccip message
        + 4 + offchain_token_data_len// offchain_token_data
        + 32 // root
        + 4 + self.proofs.len() * 32 // count + proofs
        + 4 + self.token_indexes.len() // token_indexes
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace)]
pub struct SolanaAccountMeta {
    pub pubkey: Pubkey,
    pub is_writable: bool,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct SolanaExtraArgs {
    pub compute_units: u32,
    pub accounts: Vec<SolanaAccountMeta>,
}

impl SolanaExtraArgs {
    pub fn len(&self) -> usize {
        4 // compute units
        + 4 + self.accounts.len() * SolanaAccountMeta::INIT_SPACE // additional accounts
    }
}

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct AnyExtraArgs {
    pub gas_limit: u128,
    pub allow_out_of_order_execution: bool,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaRampMessage {
    pub header: RampMessageHeader,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    // receiver is used as the target for the two main functionalities
    // token transfers: recipient of token transfers (associated token addresses are validated against this address)
    // arbitrary messaging: expected account in the declared arbitrary messaging accounts (2nd in the list of the accounts)
    pub receiver: Pubkey,
    pub token_amounts: Vec<Any2SolanaTokenTransfer>,
    pub extra_args: SolanaExtraArgs,
}

impl Any2SolanaRampMessage {
    pub fn hash(&self, on_ramp_address: &[u8]) -> [u8; 32] {
        use anchor_lang::solana_program::hash;

        // Calculate vectors size to ensure that the hash is unique
        let sender_size = [self.sender.len() as u8];
        let on_ramp_address_size = [on_ramp_address.len() as u8];
        let data_size = self.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

        // RampMessageHeader struct
        let header_source_chain_selector = self.header.source_chain_selector.to_be_bytes();
        let header_dest_chain_selector = self.header.dest_chain_selector.to_be_bytes();
        let header_sequence_number = self.header.sequence_number.to_be_bytes();
        let header_nonce = self.header.nonce.to_be_bytes();

        // Extra Args struct
        let extra_args_compute_units = self.extra_args.compute_units.to_be_bytes();
        let extra_args_accounts_len = [self.extra_args.accounts.len() as u8];
        let extra_args_accounts = self.extra_args.accounts.try_to_vec().unwrap();

        // TODO: Hash token amounts

        // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
        // As similar as https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L402
        let result = hash::hashv(&[
            "Any2SolanaMessageHashV1".as_bytes(),
            &header_source_chain_selector,
            &header_dest_chain_selector,
            &on_ramp_address_size,
            on_ramp_address,
            &self.header.message_id,
            &self.receiver.to_bytes(),
            &header_sequence_number,
            &extra_args_compute_units,
            &extra_args_accounts_len,
            &extra_args_accounts,
            &header_nonce,
            &sender_size,
            &self.sender,
            &data_size.to_be_bytes(),
            &self.data,
        ]);

        result.to_bytes()
    }

    pub fn len(&self) -> usize {
        let token_len = self.token_amounts.iter().fold(0, |acc, e| acc + e.len());

        self.header.len() // header
        + 4 + self.sender.len() // sender
        + 4 + self.data.len() // data
        + 32 // receiver
        + 4 + token_len // token_amount
        + self.extra_args.len() // extra_args
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
// Family-agnostic message emitted from the OnRamp
// Note: hash(Any2SolanaRampMessage) != hash(Solana2AnyRampMessage) due to encoding & parameter differences
// messageId = hash(Solana2AnyRampMessage) using the source EVM chain's encoding format
pub struct Solana2AnyRampMessage {
    pub header: RampMessageHeader, // Message header
    pub sender: Pubkey,            // sender address on the source chain
    pub data: Vec<u8>,             // arbitrary data payload supplied by the message sender
    pub receiver: Vec<u8>,         // receiver address on the destination chain
    pub extra_args: AnyExtraArgs, // destination-chain specific extra args, such as the gasLimit for EVM chains
    pub fee_token: Pubkey,
    pub token_amounts: Vec<Solana2AnyTokenTransfer>,
}

impl Solana2AnyRampMessage {
    pub fn hash(&self) -> [u8; 32] {
        // TODO: Modify this hash to be similar to the one in EVM
        // https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/libraries/Internal.sol#L129
        // Fixed-size message fields are included in nested hash to reduce stack pressure.
        // - metadata_hash =  sha256("Solana2AnyMessageHashV1", solana_chain_selector, dest_chain_selector, ccip_router_program_id))
        // - first_part = sha256(sender, sequence_number, nonce, fee_token, fee_token_amount)
        // - receiver
        // - message.data
        // - token_amounts
        // - extra_args

        use anchor_lang::solana_program::hash;

        // Push Data Size to ensure that the hash is unique
        let data_size = self.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

        // RampMessageHeader struct
        let header_source_chain_selector = self.header.source_chain_selector.to_be_bytes();
        let header_dest_chain_selector = self.header.dest_chain_selector.to_be_bytes();
        let header_sequence_number = self.header.sequence_number.to_be_bytes();
        let header_nonce = self.header.nonce.to_be_bytes();

        // Extra Args struct
        let extra_args_gas_limit = self.extra_args.gas_limit.to_be_bytes();
        let extra_args_allow_out_of_order_execution =
            [self.extra_args.allow_out_of_order_execution as u8];

        // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
        let result = hash::hashv(&[
            &self.sender.to_bytes(),
            &self.receiver,
            &data_size.to_be_bytes(),
            &self.data,
            &header_source_chain_selector,
            &header_dest_chain_selector,
            &header_sequence_number,
            &header_nonce,
            &extra_args_gas_limit,
            &extra_args_allow_out_of_order_execution,
        ]);

        result.to_bytes()
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct Solana2AnyTokenTransfer {
    // The source pool address. This value is trusted as it was obtained through the onRamp. It can be relied
    // upon by the destination pool to validate the source pool.
    pub source_pool_address: Pubkey,
    // The address of the destination token.
    // This value is UNTRUSTED as any pool owner can return whatever value they want.
    pub dest_token_address: Vec<u8>,
    // Optional pool data to be transferred to the destination chain. By default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    pub extra_data: Vec<u8>,
    pub amount: [u8; 32], // LE encoded u256 -  cross-chain token amount is always u256
    // Destination chain data used to execute the token transfer on the destination chain. For an EVM destination, it
    // consists of the amount of gas available for the releaseOrMint and transfer calls made by the offRamp.
    pub dest_exec_data: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaTokenTransfer {
    // The source pool address encoded to bytes. This value is trusted as it is obtained through the onRamp. It can be
    // relied upon by the destination pool to validate the source pool.
    pub source_pool_address: Vec<u8>,
    pub dest_token_address: Pubkey, // Address of destination token
    pub dest_gas_amount: u32, // The amount of gas available for the releaseOrMint and transfer calls on the offRamp.
    // Optional pool data to be transferred to the destination chain. Be default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    pub extra_data: Vec<u8>,
    pub amount: [u8; 32], // LE encoded u256, any cross-chain token amounts are u256
}

impl Any2SolanaTokenTransfer {
    pub fn len(&self) -> usize {
        4 + self.source_pool_address.len() // source_pool
        + 32 // token_address
        + 4  // gas_amount
        + 4 + self.extra_data.len()  // extra_data
        + 32 // amount
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Solana2AnyMessage {
    pub receiver: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
    pub fee_token: Pubkey, // pass zero address if native SOL
    pub extra_args: ExtraArgsInput,

    // solana specific parameter for mapping tokens to set of accounts
    pub token_indexes: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default, Debug, PartialEq, Eq)]
pub struct SolanaTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct ExtraArgsInput {
    pub gas_limit: Option<u128>,
    pub allow_out_of_order_execution: Option<bool>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
}

impl Solana2AnyMessage {
    pub fn validate(
        &self,
        dest_chain: &DestChain,
        token_config: &BillingTokenConfig,
    ) -> Result<()> {
        require!(
            dest_chain.config.is_enabled,
            CcipRouterError::DestinationChainDisabled
        );

        require!(token_config.enabled, CcipRouterError::FeeTokenDisabled);

        require_gte!(
            dest_chain.config.max_data_bytes,
            self.data.len() as u32,
            CcipRouterError::MessageTooLarge
        );

        require_gte!(
            dest_chain.config.max_number_of_tokens_per_msg as usize,
            self.token_amounts.len(),
            CcipRouterError::UnsupportedNumberOfTokens
        );

        self.validate_dest_family_address(dest_chain.config.chain_family_selector)
    }

    pub fn validate_dest_family_address(&self, chain_family_selector: [u8; 4]) -> Result<()> {
        const PRECOMPILE_SPACE: u32 = 1024;

        let selector = u32::from_be_bytes(chain_family_selector);
        // Only EVM is supported as a destination family.
        require_eq!(
            selector,
            CHAIN_FAMILY_SELECTOR_EVM,
            CcipRouterError::UnsupportedChainFamilySelector
        );

        require_eq!(self.receiver.len(), 32, CcipRouterError::InvalidEVMAddress);

        let address: U256 = U256::from_be_bytes(
            self.receiver
                .clone()
                .try_into()
                .map_err(|_| CcipRouterError::InvalidEncoding)?,
        );

        require!(address <= U160_MAX, CcipRouterError::InvalidEVMAddress);

        if let Ok(small_address) = TryInto::<u32>::try_into(address) {
            require_gte!(
                small_address,
                PRECOMPILE_SPACE,
                CcipRouterError::InvalidEVMAddress
            )
        };

        Ok(())
    }
}

#[cfg(test)]
pub(crate) mod tests {
    use crate::utils::Exponential;

    use super::*;
    use anchor_lang::solana_program::pubkey::Pubkey;
    use anchor_spl::token::spl_token::native_mint;
    use bytemuck::Zeroable;

    /// Builds a message and hash it, it's compared with a known hash
    #[test]
    fn test_hash() {
        let message = Any2SolanaRampMessage {
            sender: [
                1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0,
            ]
            .to_vec(),
            receiver: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb").unwrap(),
            data: vec![4, 5, 6],
            header: RampMessageHeader {
                message_id: [
                    8, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                    0, 0, 0, 0, 0, 0,
                ],
                source_chain_selector: 67,
                dest_chain_selector: 78,
                sequence_number: 89,
                nonce: 90,
            },
            token_amounts: [].to_vec(), // TODO: hash token amounts
            extra_args: SolanaExtraArgs {
                compute_units: 1000,
                accounts: vec![SolanaAccountMeta {
                    pubkey: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
                        .unwrap(),
                    is_writable: true,
                }],
            },
        };

        let on_ramp_address = &[1, 2, 3].to_vec();
        let hash_result = message.hash(on_ramp_address);

        assert_eq!(
            "03da97f96c82237d8a8ab0f68d4f7ba02afe188b4a876f348278fbf2226312ed",
            hex::encode(hash_result)
        );
    }

    #[test]
    fn message_not_validated_for_disabled_destination_chain() {
        let mut chain = sample_dest_chain();
        chain.config.is_enabled = false;

        assert_eq!(
            sample_message()
                .validate(&chain, &sample_billing_config())
                .unwrap_err(),
            CcipRouterError::DestinationChainDisabled.into()
        );
    }

    #[test]
    fn message_not_validated_for_disabled_token() {
        let mut billing_config = sample_billing_config();
        billing_config.enabled = false;

        assert_eq!(
            sample_message()
                .validate(&sample_dest_chain(), &billing_config)
                .unwrap_err(),
            CcipRouterError::FeeTokenDisabled.into()
        );
    }

    #[test]
    fn large_message_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.data = vec![0; dest_chain.config.max_data_bytes as usize + 1];
        assert_eq!(
            message
                .validate(&sample_dest_chain(), &sample_billing_config())
                .unwrap_err(),
            CcipRouterError::MessageTooLarge.into()
        );
    }

    #[test]
    fn invalid_addresses_fail_to_validate() {
        let mut address_bigger_than_u160_max = vec![0u8; 32];
        address_bigger_than_u160_max[11] = 1;
        let mut address_in_precompile_space = vec![0u8; 32];
        address_in_precompile_space[30] = 1;
        let incorrect_length_address = vec![1u8, 12];

        let invalid_addresses = [
            address_bigger_than_u160_max,
            address_in_precompile_space,
            incorrect_length_address,
        ];

        let mut message = sample_message();
        for address in invalid_addresses {
            message.receiver = address;
            assert_eq!(
                message
                    .validate(&sample_dest_chain(), &sample_billing_config())
                    .unwrap_err(),
                CcipRouterError::InvalidEVMAddress.into()
            );
        }
    }

    #[test]
    fn message_with_too_many_tokens_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.token_amounts = vec![
            SolanaTokenAmount {
                token: Pubkey::new_unique(),
                amount: 1
            };
            dest_chain.config.max_number_of_tokens_per_msg as usize + 1
        ];
        assert_eq!(
            message
                .validate(&sample_dest_chain(), &sample_billing_config())
                .unwrap_err(),
            CcipRouterError::UnsupportedNumberOfTokens.into()
        );
    }

    pub fn sample_message() -> Solana2AnyMessage {
        let mut receiver = vec![0u8; 32];

        // Arbitrary value that pushes the address to the right EVM range
        // (above precompile space, under u160::max)
        receiver[20] = 0xA;

        Solana2AnyMessage {
            receiver,
            data: vec![],
            token_amounts: vec![],
            fee_token: Pubkey::zeroed(),
            extra_args: crate::ExtraArgsInput {
                gas_limit: None,
                allow_out_of_order_execution: None,
            },
            token_indexes: vec![],
        }
    }

    pub fn sample_billing_config() -> BillingTokenConfig {
        let mut value = [0; 28];
        value.clone_from_slice(&3u32.e(18).to_be_bytes()[4..]);
        BillingTokenConfig {
            enabled: true,
            mint: native_mint::ID,
            usd_per_token: crate::TimestampedPackedU224 {
                value,
                timestamp: 100,
            },
            premium_multiplier_wei_per_eth: 1,
        }
    }

    pub fn sample_dest_chain() -> DestChain {
        let mut value = [0; 28];
        // L1 gas price
        value[0..14].clone_from_slice(&1u32.e(18).to_be_bytes()[18..]);
        // L2 gas price
        value[14..].clone_from_slice(&U256::new(22u128).to_be_bytes()[18..]);
        DestChain {
            version: 1,
            chain_selector: 1,
            state: crate::DestChainState {
                sequence_number: 0,
                usd_per_unit_gas: crate::TimestampedPackedU224 {
                    value,
                    timestamp: 100,
                },
            },
            config: crate::DestChainConfig {
                is_enabled: true,
                max_number_of_tokens_per_msg: 5,
                max_data_bytes: 200,
                max_per_msg_gas_limit: 0,
                dest_gas_overhead: 1,
                dest_gas_per_payload_byte: 0,
                dest_data_availability_overhead_gas: 0,
                dest_gas_per_data_availability_byte: 1,
                dest_data_availability_multiplier_bps: 1,
                default_token_fee_usdcents: 100,
                default_token_dest_gas_overhead: 0,
                default_tx_gas_limit: 0,
                gas_multiplier_wei_per_eth: 1,
                network_fee_usdcents: 100,
                gas_price_staleness_threshold: 10,
                enforce_out_of_order: false,
                chain_family_selector: CHAIN_FAMILY_SELECTOR_EVM.to_be_bytes(),
            },
        }
    }
}
