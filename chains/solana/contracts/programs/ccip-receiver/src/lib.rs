use anchor_lang::prelude::*;
use anchor_spl::{token_interface::Mint, token_interface::TokenAccount};
use solana_program::pubkey;

declare_id!("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH");

pub const EXTERNAL_EXECUTION_CONFIG_SEED: &[u8] = b"external_execution_config";

/// This program an example of a CCIP Receiver Program.
/// Used to test CCIP Router execute.
#[program]
pub mod ccip_receiver {
    use solana_program::{instruction::Instruction, program::invoke_signed};

    use super::*;

    /// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        msg!("Called `initialize` {:?}", ctx);
        ctx.accounts.counter.value = 0;
        Ok(())
    }

    /// This function is called by the CCIP Router to execute the CCIP message.
    /// The method name needs to be ccip_receive with Anchor encoding,
    /// if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]
    /// You can send as many accounts as you need, specifying if mutable or not.
    /// But none of them could be an init, realloc or close.
    /// In this case, it increments the counter value by 1 and logs the parsed message.
    pub fn ccip_receive(ctx: Context<SetData>, message: Any2SolanaMessage) -> Result<()> {
        msg!("Called `ccip_receive` with message {:?}", message);

        let counter = &mut ctx.accounts.counter;
        counter.value += 1;

        // additional accounts trigger additional CPI call
        if !ctx.remaining_accounts.is_empty() {
            let external_execution_config = &ctx.accounts.external_execution_config;
            let (target_program, acc_infos) = ctx.remaining_accounts.split_at(1);

            let acc_metas: Vec<AccountMeta> = acc_infos
                .to_vec()
                .iter()
                .flat_map(|acc_info| {
                    // Check signer from PDA External Execution config
                    let is_signer = acc_info.key() == external_execution_config.key();
                    acc_info.to_account_metas(Some(is_signer))
                })
                .collect();

            let instruction = Instruction {
                program_id: target_program[0].key(),
                accounts: acc_metas,
                data: message.data,
            };

            let seeds = &[
                EXTERNAL_EXECUTION_CONFIG_SEED,
                &[ctx.bumps.external_execution_config],
            ];
            let signer = &[&seeds[..]];

            invoke_signed(&instruction, acc_infos, signer)?;
        };

        Ok(())
    }

    // these functions are called by CCIP token pools that wrap non-token programs
    // if not using Anchor, the discriminator must be:
    // const RELEASE_MINT: [u8; 8] = [0x14, 0x94, 0x71, 0xc6, 0xe5, 0xaa, 0x47, 0x30];
    // const LOCK_BURN: [u8; 8] = [0xc8, 0x0e, 0x32, 0x09, 0x2c, 0x5b, 0x79, 0x25];
    pub fn ccip_token_release_mint(
        _ctx: Context<TokenPool>,
        input: ReleaseOrMintInV1,
    ) -> Result<()> {
        msg!("Called `ccip_token_release_mint` with message {:?}", input);
        Ok(())
    }
    pub fn ccip_token_lock_burn(_ctx: Context<TokenPool>, input: LockOrBurnInV1) -> Result<()> {
        msg!("Called `ccip_token_lock_burn` with message {:?}", input);
        Ok(())
    }

    pub fn tobi_backend<'info>(ctx: Context<'_, '_, 'info, 'info, Tobi>) -> Result<()> {
        msg!("Called `tobi_backend`");
        let acc_info = &ctx.remaining_accounts[0];

        require!(acc_info.is_signer, TobiErr::AccountNotSigner);
        require!(acc_info.is_writable, TobiErr::AccountNotWritable);

        let mut dest_chain_data = deserialize(acc_info)?;

        // dest_chain_data.state.sequence_number += 1;

        persist(acc_info, dest_chain_data)?;

        Ok(())
    }
}

fn deserialize(info: &AccountInfo) -> Result<DestChain> {
    let mut data: &[u8] = &info.try_borrow_data()?;
    DestChain::try_deserialize(&mut data)
}

fn persist(info: &AccountInfo, dest_chain: DestChain) -> Result<()> {
    let mut data = info.try_borrow_mut_data()?;
    let dst: &mut [u8] = &mut data;
    let mut writer = BpfWriter::new(dst);
    dest_chain.try_serialize(&mut writer)
}

const ANCHOR_DISCRIMINATOR: usize = 8;

#[error_code]
pub enum TobiErr {
    #[msg("Account is not signer")]
    AccountNotSigner,
    #[msg("Account is not writable")]
    AccountNotWritable,
}

#[derive(Accounts)]
pub struct Tobi {}

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [b"counter"],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Counter::INIT_SPACE,
    )]
    pub counter: Account<'info, Counter>,

    #[account(
        init,
        seeds = [EXTERNAL_EXECUTION_CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExternalExecutionConfig::INIT_SPACE,
    )]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct SetData<'info> {
    // router CPI signer must be first
    #[account(owner = CCIP_ROUTER)]
    pub authority: Signer<'info>,
    // ccip router expects "receiver" to be second
    /// CHECK: Using this to sign
    #[account(mut, seeds = [EXTERNAL_EXECUTION_CONFIG_SEED], bump)]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,
    #[account(
        mut,
        seeds = [b"counter"],
        bump,
    )]
    pub counter: Account<'info, Counter>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct TokenPool<'info> {
    // token pool CPI signer
    pub authority: Signer<'info>,
    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    pub mint: InterfaceAccount<'info, Mint>,
    /// CHECK: CPI to underlying token program
    pub token_program: UncheckedAccount<'info>,
    // remaining accounts
    // [0] receiver_token_account // only included for release_mint
    // [...] extra passed accounts
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Counter {
    pub value: u8,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct ExternalExecutionConfig {}

// TODO: Import types and constants from CCIP Router, it should be an imported crate
// But for now, we are copying the structs here as the final design of the messages is not done.

const CCIP_ROUTER: Pubkey = pubkey!("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8");

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SolanaTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // solana local token amount
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnInV1 {
    receiver: Vec<u8>, //  The recipient of the tokens on the destination chain, abi encoded
    remote_chain_selector: u64, // The chain ID of the destination chain
    original_sender: Pubkey, // The original sender of the tx on the source chain
    amount: u64, // local solana amount to lock/burn,  The amount of tokens to lock or burn, denominated in the source token's decimals
    local_token: Pubkey, //  The address on this chain of the token to lock or burn
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintInV1 {
    original_sender: Vec<u8>, //          The original sender of the tx on the source chain
    remote_chain_selector: u64, // ─╮ The chain ID of the source chain
    receiver: Pubkey,         // ───────────╯ The recipient of the tokens on the destination chain.
    amount: [u8; 32], // u256, incoming cross-chain amount - The amount of tokens to release or mint, denominated in the source token's decimals
    local_token: Pubkey, //            The address on this chain of the token to release or mint
    /// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
    /// expected pool address for the given remoteChainSelector.
    source_pool_address: Vec<u8>, //       The address of the source pool, abi encoded in the case of EVM chains
    source_pool_data: Vec<u8>, //          The data received from the source pool to process the release or mint
    /// @dev WARNING: offchainTokenData is untrusted data.
    offchain_token_data: Vec<u8>, //       The offchain data to process the release or mint
}

#[account]
#[derive(InitSpace, Debug)]
pub struct DestChain {
    // Config for Solana2Any
    pub version: u8,
    pub state: DestChainState,   // values that are updated automatically
    pub config: DestChainConfig, // values configured by an admin
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct SourceChainConfig {
    pub is_enabled: bool, // Flag whether the source chain is enabled or not
    #[max_len(64)]
    pub on_ramp: Vec<u8>, // OnRamp address on the source chain
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct SourceChainState {
    pub min_seq_nr: u64, // The min sequence number expected for future messages
}

#[account]
#[derive(InitSpace, Debug)]
pub struct SourceChain {
    // Config for Any2Solana
    pub version: u8,
    pub state: SourceChainState, // values that are updated automatically
    pub config: SourceChainConfig, // values configured by an admin
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct DestChainState {
    pub sequence_number: u64, // The last used sequence number
    pub usd_per_unit_gas: TimestampedPackedU224,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct DestChainConfig {
    pub is_enabled: bool, // Whether this destination chain is enabled

    pub max_number_of_tokens_per_msg: u16, // Maximum number of distinct ERC20 token transferred per message
    pub max_data_bytes: u32,               // Maximum payload data size in bytes
    pub max_per_msg_gas_limit: u32,        // Maximum gas limit for messages targeting EVMs
    pub dest_gas_overhead: u32, //  Gas charged on top of the gasLimit to cover destination chain costs
    pub dest_gas_per_payload_byte: u16, // Destination chain gas charged for passing each byte of `data` payload to receiver
    pub dest_data_availability_overhead_gas: u32, // Extra data availability gas charged on top of the message, e.g. for OCR
    pub dest_gas_per_data_availability_byte: u16, // Amount of gas to charge per byte of message data that needs availability
    pub dest_data_availability_multiplier_bps: u16, // Multiplier for data availability gas, multiples of bps, or 0.0001

    // The following three properties are defaults, they can be overridden by setting the TokenTransferFeeConfig for a token
    pub default_token_fee_usdcents: u16, // Default token fee charged per token transfer
    pub default_token_dest_gas_overhead: u32, //  Default gas charged to execute the token transfer on the destination chain
    pub default_tx_gas_limit: u32,            // Default gas limit for a tx
    pub gas_multiplier_wei_per_eth: u64, // Multiplier for gas costs, 1e18 based so 11e17 = 10% extra cost.
    pub network_fee_usdcents: u32, // Flat network fee to charge for messages, multiples of 0.01 USD
    pub gas_price_staleness_threshold: u32, // The amount of time a gas price can be stale before it is considered invalid (0 means disabled)
    pub enforce_out_of_order: bool, // Whether to enforce the allowOutOfOrderExecution extraArg value to be true.
    pub chain_family_selector: [u8; 4], // Selector that identifies the destination chain's family. Used to determine the correct validations to perform for the dest chain.
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize, Debug)]
pub struct TimestampedPackedU224 {
    pub value: [u8; 28],
    pub timestamp: i64, // maintaining the type that Solana returns for the time (solana_program::clock::UnixTimestamp = i64)
}

use solana_program::program_memory::sol_memcpy;
use std::cmp;
use std::io::{self, Write};

#[derive(Debug, Default)]
pub struct BpfWriter<T> {
    inner: T,
    pos: u64,
}

impl<T> BpfWriter<T> {
    pub fn new(inner: T) -> Self {
        Self { inner, pos: 0 }
    }
}

impl Write for BpfWriter<&mut [u8]> {
    fn write(&mut self, buf: &[u8]) -> io::Result<usize> {
        if self.pos >= self.inner.len() as u64 {
            return Ok(0);
        }

        let amt = cmp::min(
            self.inner.len().saturating_sub(self.pos as usize),
            buf.len(),
        );
        sol_memcpy(&mut self.inner[(self.pos as usize)..], buf, amt);
        self.pos += amt as u64;
        Ok(amt)
    }

    fn write_all(&mut self, buf: &[u8]) -> io::Result<()> {
        if self.write(buf)? == buf.len() {
            Ok(())
        } else {
            Err(io::Error::new(
                io::ErrorKind::WriteZero,
                "failed to write whole buffer",
            ))
        }
    }

    fn flush(&mut self) -> io::Result<()> {
        Ok(())
    }
}
