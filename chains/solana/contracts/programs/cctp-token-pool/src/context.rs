use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

use base_token_pool::common::*;
use burnmint_token_pool::ChainConfig;
use ccip_common::seed;

use crate::{CctpTokenPoolError, MessageAndAttestation};

const MAX_POOL_STATE_V: u8 = 1;
const ANCHOR_DISCRIMINATOR: usize = 8;

pub const TOKEN_MESSENGER_MINTER: Pubkey =
    solana_program::pubkey!("CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3");
pub const MESSAGE_TRANSMITTER: Pubkey =
    solana_program::pubkey!("CCTPmbSD7gX1bxKPAmg77w8oFzNFpaQiQUWD43TKaecd");

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        space = ANCHOR_DISCRIMINATOR + burnmint_token_pool::State::INIT_SPACE,
        payer = authority,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(1, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, crate::program::CctpTokenPool>,

    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()))]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.proposed_owner == authority.key() @ CcipTokenPoolError::Unauthorized
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct TokenOfframp<'info> {
    // CCIP accounts ------------------------
    #[account(
        seeds = [EXTERNAL_TOKEN_POOLS_SIGNER, crate::ID.as_ref()],
        bump,
        seeds::program = offramp_program.key(),
    )]
    pub authority: Signer<'info>,

    /// CHECK offramp program: exists only to derive the allowed offramp PDA
    /// and the authority PDA.
    pub offramp_program: UncheckedAccount<'info>,

    /// CHECK PDA of the router program verifying the signer is an allowed offramp.
    /// If PDA does not exist, the router doesn't allow this offramp
    #[account(
        owner = state.config.router @ CcipTokenPoolError::InvalidPoolCaller, // this guarantees that it was initialized
        seeds = [
            ALLOWED_OFFRAMP,
            release_or_mint.remote_chain_selector.to_le_bytes().as_ref(),
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = state.config.router,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.router != Pubkey::default() @ CcipTokenPoolError::InvalidInputs,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub token_program: Interface<'info, TokenInterface>,

    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK: CPI signer
    #[account(
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        bump,
        address = state.config.pool_signer,
    )]
    pub pool_signer: UncheckedAccount<'info>,

    #[account(
        mut,
        associated_token::mint = mint,
        associated_token::authority = pool_signer,
        associated_token::token_program = token_program,
    )]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    #[account(
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &release_or_mint.remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = state.config.rmn_remote @ CcipTokenPoolError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,

    // CCTP pool-specific accounts ----------------
    #[account(
        mut,
        address = get_associated_token_address_with_program_id(&release_or_mint.receiver, &mint.key(), &token_program.key())
    )]
    pub receiver_token_account: InterfaceAccount<'info, TokenAccount>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"message_transmitter_authority", TOKEN_MESSENGER_MINTER.as_ref()],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_authority_pda: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"message_transmitter"],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_message_transmitter_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"used_nonces", to_domain_seed(release_or_mint.remote_chain_selector).as_ref(), first_nonce_seed(release_or_mint.offchain_token_data)?.as_ref()],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_used_nonces: UncheckedAccount<'info>,

    /// CHECK this is CCTP's TokenMessengerMinter program, which
    /// is invoked by CCTP's MessageTransmitter as it is the receiver of the CCTP message.
    #[account(address = TOKEN_MESSENGER_MINTER @ CctpTokenPoolError::InvalidTokenMessengerMinter)]
    pub cctp_token_messenger_minter: UncheckedAccount<'info>,

    pub system_program: Program<'info, System>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"__event_authority"],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_event_authority: UncheckedAccount<'info>,

    /// CHECK this is the CCTP Message Transmitter program, which
    /// is invoked by this program.
    #[account(address = MESSAGE_TRANSMITTER @ CctpTokenPoolError::InvalidMessageTransmitter)]
    pub cctp_message_transmitter: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"token_messenger"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_token_messenger_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"token_minter"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_token_minter_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"local_token", mint.key().as_ref()],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_local_token: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"remote_token_messenger", &to_domain_seed(release_or_mint.remote_chain_selector)],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_remote_token_messenger_key: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"token_pair", &to_domain_seed(release_or_mint.remote_chain_selector), to_solana_pubkey(&chain_config.base.remote.token_address).as_ref()],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_token_pair: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"custody", mint.key().as_ref()],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_custody_token_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"__event_authority"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_token_messenger_event_authority: UncheckedAccount<'info>,
}

fn first_nonce_seed(offchain_token_data: Vec<u8>) -> Result<Vec<u8>> {
    let message_and_attestation = MessageAndAttestation::try_from_slice(&offchain_token_data)
        .expect("Failed to deserialize MessageAndAttestation");

    let nonce_bytes: [u8; 8] = message_and_attestation.message[12..20]
        .try_into()
        .expect("Failed to convert slice to array");

    let nonce: u64 = u64::from_be_bytes(nonce_bytes);

    // TODO consider depending on the CCTP crate and using their method for this, which already checks for overflows.
    let first_nonce = (nonce - 1) / 6400 * 6400 + 1;

    Ok(first_nonce.to_string().as_bytes().to_vec())
}

fn to_solana_pubkey(address: &RemoteAddress) -> Pubkey {
    let mut solana_address = [0; 32];
    let remote_bytes = address.len();
    solana_address[32 - remote_bytes..].copy_from_slice(&address);
    Pubkey::new_from_array(solana_address)
}

// TODO alternatively, remove this and store the corresponding DomainID in the ChainConfig
fn to_domain_seed(remote_chain_selector: u64) -> Vec<u8> {
    let bytes: &'static [u8] = match remote_chain_selector {
        // TODO update
        0 => b"0",
        1 => b"1",
        2 => b"2",
        3 => b"3",
        // this should never happen, as we won't even configure unsupported chains for this pool
        _ => panic!("Invalid chain selector for CCTP"),
    };
    bytes.to_vec()
}

#[derive(Accounts)]
#[instruction(lock_or_burn: LockOrBurnInV1)]
pub struct TokenOnramp<'info> {
    // CCIP accounts ------------------------
    #[account(address = state.config.router_onramp_authority @ CcipTokenPoolError::InvalidPoolCaller)]
    pub authority: Signer<'info>,

    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub token_program: Interface<'info, TokenInterface>,

    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK: CPI signer. This account is intentionally not initialized, and it will
    /// hold a balance to pay for the rent of initializing the CCTP MessageSentEvent account
    #[account(
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        bump,
        address = state.config.pool_signer,
    )]
    pub pool_signer: UncheckedAccount<'info>,

    #[account(
        mut,
        associated_token::mint = mint,
        associated_token::authority = pool_signer,
        associated_token::token_program = token_program,
    )]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: RMNRemote program, invoked to check for curses
    #[account(
        address = state.config.rmn_remote @ CcipTokenPoolError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,

    #[account(
        seeds = [
            POOL_CHAINCONFIG_SEED,
            lock_or_burn.remote_chain_selector.to_le_bytes().as_ref(),
            mint.key().as_ref()
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    // CCTP pool-specific accounts ----------------
    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"message_transmitter_authority", TOKEN_MESSENGER_MINTER.as_ref()],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_authority_pda: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"message_transmitter"],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_message_transmitter_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"token_messenger"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_token_messenger_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"remote_token_messenger", &to_domain_seed(lock_or_burn.remote_chain_selector)],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_remote_token_messenger_key: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"token_minter"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_token_minter_account: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"local_token", mint.key().as_ref()],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_local_token: UncheckedAccount<'info>,

    /// CHECK this is the account in which CCTP will store the event. It is not a PDA of CCTP,
    /// but CCTP is the owner for it
    #[account(
        // TODO: this should be a PDA of the pool, that is unique for this message.
        // However, the pool does not have access to the message in order to use its info as validated seeds.
        // We could instead not validate it here and trust that the onramp will derive it correctly.
    )]
    pub cctp_message_sent_event: UncheckedAccount<'info>,

    /// CHECK this is CCTP's MessageTransmitter program, which
    /// is invoked CCTP's TokenMessengerMinter by this program.
    #[account(
        address = MESSAGE_TRANSMITTER @ CctpTokenPoolError::InvalidMessageTransmitter
    )]
    pub cctp_message_transmitter: UncheckedAccount<'info>,

    /// CHECK this is CCTP's TokenMessengerMinter program, which
    /// is invoked by this program.
    #[account(
        address = TOKEN_MESSENGER_MINTER @ CctpTokenPoolError::InvalidTokenMessengerMinter
    )]
    pub cctp_token_messenger_minter: UncheckedAccount<'info>,

    pub system_program: Program<'info, System>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"__event_authority"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_event_authority: UncheckedAccount<'info>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct InitializeChainConfig<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        init,
        payer = authority,
        space = 8 + ChainConfig::INIT_SPACE,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct EditChainConfigDynamicSize<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
        realloc = 8 + ChainConfig::INIT_SPACE,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, addresses: Vec<RemoteAddress>)]
pub struct AppendRemotePoolAddresses<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
        realloc = 8 + ChainConfig::INIT_SPACE + addresses.len() * RemoteAddress::INIT_SPACE,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct SetChainRateLimit<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct DeleteChainConfig<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        close = authority,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(add: Vec<Pubkey>)]
pub struct AddToAllowList<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
        realloc = 8 + burnmint_token_pool::State::INIT_SPACE + add.len() * 32,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remove: Vec<Pubkey>)]
pub struct RemoveFromAllowlist<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct Empty<'info> {
    // This is unused, but Anchor requires that there is at least one account in the context
    pub clock: Sysvar<'info, Clock>,
}
