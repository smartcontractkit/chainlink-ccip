use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

use base_token_pool::common::*;
use ccip_common::seed;

use crate::cctp_message::CctpMessage;
use crate::program::CctpTokenPool;
use crate::{CctpTokenPoolError, ChainConfig, PoolConfig, State};

const MAX_POOL_STATE_V: u8 = 1;
const MAX_POOL_CONFIG_V: u8 = 1;
const MAX_POOL_CHAIN_CONFIG_V: u8 = 1;

const ANCHOR_DISCRIMINATOR: usize = 8;

pub const TOKEN_MESSENGER_MINTER: Pubkey =
    solana_program::pubkey!("CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3");
pub const MESSAGE_TRANSMITTER: Pubkey =
    solana_program::pubkey!("CCTPmbSD7gX1bxKPAmg77w8oFzNFpaQiQUWD43TKaecd");

pub const MESSAGE_SENT_EVENT_SEED: &[u8] = b"ccip_cctp_message_sent_event";

#[derive(Accounts)]
pub struct InitGlobalConfig<'info> {
    #[account(
        init,
        seeds = [CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + PoolConfig::INIT_SPACE,
    )]
    pub config: Account<'info, PoolConfig>, // Global Config PDA of the Token Pool

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    // Ensures that the provided program is the CctpTokenPool program,
    // and that its associated program data account matches the expected one.
    // This guarantees that only the program's upgrade authority can modify the global config.
    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CctpTokenPool>,
    // Global Config updates only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref(),
        ],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + State::INIT_SPACE,
    )]
    pub state: Account<'info, State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    // Token pool initialization only allowed by program upgrade authority. Initializing token pools managed
    // by the CLL deployment of this program is limited to CLL. Users must deploy their own instance of this program.
    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CctpTokenPool>,

    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()))]
    pub program_data: Account<'info, ProgramData>,

    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_POOL_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub config: Account<'info, PoolConfig>, // Global Config PDA of the Token Pool
}

#[derive(Accounts)]
pub struct AdminUpdateTokenPool<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>, // config PDA for token pool
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(mut)]
    pub authority: Signer<'info>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, CctpTokenPool>,

    // The upgrade authority of the token pool program can update certain values of their token pools only
    // (router and rmn addresses, for testing)
    // This is consistent with the BurnmintTokenPool program, although it is expected that this pool will always
    // be owned by the upgrade authority of the token pool program.
    #[account(constraint = allowed_admin_modify_token_pool(&program_data, &authority, &state) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct Empty {}

#[derive(Accounts)]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized)]
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
        realloc = ANCHOR_DISCRIMINATOR + State::INIT_SPACE + 32 * (state.config.allow_list.len() + add.len()),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(mut, address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
        realloc = ANCHOR_DISCRIMINATOR + State::INIT_SPACE + 32 * (state.config.allow_list.len().saturating_sub(remove.len())),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(mut, address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(address = state.config.proposed_owner @ CcipTokenPoolError::Unauthorized)]
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(address = state.config.token_program)]
    pub token_program: Interface<'info, TokenInterface>,

    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK: CPI signer
    #[account(
        mut, // CCTP's receive_message method requires that this account is mutable,
             // although it does not really mutate it
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
            &release_or_mint.remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        constraint = valid_version(chain_config.version, MAX_POOL_CHAIN_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
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

    #[account(
        mut,
        address = get_associated_token_address_with_program_id(&release_or_mint.receiver, &mint.key(), &token_program.key())
    )]
    pub receiver_token_account: InterfaceAccount<'info, TokenAccount>,
}

#[inline]
pub fn get_message_transmitter_pda(seeds: &[&[u8]]) -> Pubkey {
    Pubkey::find_program_address(seeds, &MESSAGE_TRANSMITTER).0
}
#[inline]
pub fn get_token_messenger_minter_pda(seeds: &[&[u8]]) -> Pubkey {
    Pubkey::find_program_address(seeds, &TOKEN_MESSENGER_MINTER).0
}

// This contains the CCTP-specific accounts that are used in the offramp. As they are too many, they don't fit inside
// a single Context struct (as Anchor would require too much memory to validate all of them), so we use a separate struct
// and do the address validations manually.
pub struct TokenOfframpRemainingAccounts<'info> {
    // Accounts that are in the lookup table, as they are static and shared between onramp & offramp
    pub cctp_message_transmitter_account: &'info AccountInfo<'info>,
    pub cctp_token_messenger_minter: &'info AccountInfo<'info>,
    pub system_program: &'info AccountInfo<'info>,
    pub cctp_message_transmitter: &'info AccountInfo<'info>,
    pub cctp_token_messenger_account: &'info AccountInfo<'info>,
    pub cctp_token_minter_account: &'info AccountInfo<'info>,
    pub cctp_local_token: &'info AccountInfo<'info>,
    pub cctp_token_messenger_minter_event_authority: &'info AccountInfo<'info>,

    // Accounts that are not in the lookup table, as they are either dynamic or not shared with onramp
    pub cctp_authority_pda: &'info AccountInfo<'info>,
    pub cctp_event_authority: &'info AccountInfo<'info>,
    pub cctp_custody_token_account: &'info AccountInfo<'info>,
    pub cctp_remote_token_messenger_key: &'info AccountInfo<'info>,
    pub cctp_token_pair: &'info AccountInfo<'info>,
    pub cctp_used_nonces: &'info AccountInfo<'info>,
}

impl TokenOfframpRemainingAccounts<'_> {
    pub fn validate(
        &self,
        mint: Pubkey,
        cctp_msg: &CctpMessage,
        domain_id: u32,
        remote_token_address: [u8; 32],
    ) -> Result<()> {
        let nonce_seed = Self::first_nonce_seed(cctp_msg);
        let domain_str = domain_id.to_string();
        let domain_seed = domain_str.as_bytes();

        require_keys_eq!(
            self.cctp_message_transmitter_account.key(),
            get_message_transmitter_pda(&[b"message_transmitter"])
        );

        require_keys_eq!(
            self.cctp_token_messenger_minter.key(),
            TOKEN_MESSENGER_MINTER,
        );

        require_keys_eq!(self.system_program.key(), System::id());

        require_keys_eq!(self.cctp_message_transmitter.key(), MESSAGE_TRANSMITTER);

        require_keys_eq!(
            self.cctp_token_messenger_account.key(),
            get_token_messenger_minter_pda(&[b"token_messenger"])
        );

        require_keys_eq!(
            self.cctp_token_minter_account.key(),
            get_token_messenger_minter_pda(&[b"token_minter"])
        );

        require_keys_eq!(
            self.cctp_local_token.key(),
            get_token_messenger_minter_pda(&[b"local_token", mint.as_ref()])
        );

        require_keys_eq!(
            self.cctp_token_messenger_minter_event_authority.key(),
            get_token_messenger_minter_pda(&[b"__event_authority"])
        );

        require_keys_eq!(
            self.cctp_authority_pda.key(),
            get_message_transmitter_pda(&[
                b"message_transmitter_authority",
                TOKEN_MESSENGER_MINTER.as_ref(),
            ])
        );

        require_keys_eq!(
            self.cctp_event_authority.key(),
            get_message_transmitter_pda(&[b"__event_authority"])
        );

        require_keys_eq!(
            self.cctp_custody_token_account.key(),
            get_token_messenger_minter_pda(&[b"custody", mint.key().as_ref()])
        );

        require_keys_eq!(
            self.cctp_remote_token_messenger_key.key(),
            get_token_messenger_minter_pda(&[b"remote_token_messenger", domain_seed])
        );

        require_keys_eq!(
            self.cctp_token_pair.key(),
            get_token_messenger_minter_pda(&[
                b"token_pair",
                domain_seed,
                remote_token_address.as_ref()
            ])
        );

        require_keys_eq!(
            self.cctp_used_nonces.key(),
            get_message_transmitter_pda(&[b"used_nonces", domain_seed, nonce_seed.as_ref()])
        );

        Ok(())
    }

    pub fn first_nonce_seed(msg: &CctpMessage) -> Vec<u8> {
        msg.first_nonce().to_string().as_bytes().to_vec()
    }
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Box<Account<'info, State>>,

    pub token_program: Interface<'info, TokenInterface>,

    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK: CPI signer. This account is intentionally not initialized, and it will
    /// hold a balance to pay for the rent of initializing the CCTP MessageSentEvent account
    #[account(
        mut,
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
        constraint = valid_version(chain_config.version, MAX_POOL_CHAIN_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    // CCTP pool-specific accounts in the lookup table ----------------
    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        mut,
        seeds = [b"message_transmitter"],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub cctp_message_transmitter_account: UncheckedAccount<'info>,

    /// CHECK this is CCTP's TokenMessengerMinter program, which
    /// is invoked by this program.
    #[account(
        address = TOKEN_MESSENGER_MINTER @ CctpTokenPoolError::InvalidTokenMessengerMinter
    )]
    pub cctp_token_messenger_minter: UncheckedAccount<'info>,

    pub system_program: Program<'info, System>,

    /// CHECK this is CCTP's MessageTransmitter program, which
    /// is invoked transitively by CCTP's TokenMessengerMinter,
    /// which in turn is invoked explicitly by this program.
    #[account(
        address = MESSAGE_TRANSMITTER @ CctpTokenPoolError::InvalidMessageTransmitter
    )]
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
        mut,
        seeds = [b"local_token", mint.key().as_ref()],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_local_token: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"__event_authority"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_event_authority: UncheckedAccount<'info>,

    // CCTP pool-specific accounts NOT in the lookup table ----------------
    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"sender_authority"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_authority_pda: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"remote_token_messenger", chain_config.cctp.domain_id.to_string().as_bytes()],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_remote_token_messenger_key: UncheckedAccount<'info>,

    /// CHECK this is the account in which CCTP will store the event. It is not a PDA of CCTP,
    /// but CCTP will initialize it and become the owner for it.
    #[account(
        mut,
        seeds = [
            MESSAGE_SENT_EVENT_SEED,
            &lock_or_burn.original_sender.to_bytes(),
            &lock_or_burn.remote_chain_selector.to_le_bytes(),
            &lock_or_burn.msg_total_nonce.to_le_bytes(),
        ],
        bump,
        owner = System::id(), // this is not initialized yet, it will later be owned by CCTP
    )]
    pub cctp_message_sent_event: UncheckedAccount<'info>,
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(
        init,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct EditChainConfig<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        constraint = valid_version(chain_config.version, MAX_POOL_CHAIN_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut, constraint = authority.key() == state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        constraint = valid_version(chain_config.version, MAX_POOL_CHAIN_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut, constraint = authority.key() == state.config.owner || authority.key() == state.config.rate_limit_admin @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct SetRateLimitAdmin<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(mut, address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, cfg: RemoteConfig)]
pub struct EditChainConfigDynamicSize<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        constraint = valid_version(chain_config.version, MAX_POOL_CHAIN_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE + cfg.pool_addresses.iter().map(RemoteAddress::space).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut, address = state.config.owner)]
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        constraint = valid_version(chain_config.version, MAX_POOL_CHAIN_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE
            + chain_config.base.remote.pool_addresses.iter().map(RemoteAddress::space).sum::<usize>()
            + addresses.iter().map(RemoteAddress::space).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
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
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        close = authority,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey, original_sender: Pubkey, remote_chain_selector: u64, msg_nonce: u64)]
pub struct ReclaimEventAccount<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    /// CHECK: CPI signer for CCTP operations, validated by seeds constraint
    #[account(
        mut,
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        bump,
    )]
    pub pool_signer: UncheckedAccount<'info>,

    /// CHECK: MessageSent event account to reclaim rent from
    #[account(
        mut,
        seeds = [
            MESSAGE_SENT_EVENT_SEED,
            &original_sender.to_bytes(),
            &remote_chain_selector.to_le_bytes(),
            &msg_nonce.to_le_bytes(),
        ],
        bump,
    )]
    pub message_sent_event_account: UncheckedAccount<'info>,

    /// CHECK: CCTP Message Transmitter program, validated by address constraint
    #[account(
        mut,
        seeds = [b"message_transmitter"],
        bump,
        seeds::program = cctp_message_transmitter,
    )]
    pub message_transmitter: UncheckedAccount<'info>,

    /// CHECK: CCTP Message Transmitter program account
    #[account(address = MESSAGE_TRANSMITTER)]
    pub cctp_message_transmitter: UncheckedAccount<'info>,

    #[account(mut, address = state.funding.manager @ CctpTokenPoolError::InvalidFundManager)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct ReclaimFunds<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,

    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK: CPI signer for SOL recovery, validated by seeds constraint
    #[account(
        mut,
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        address = state.config.pool_signer,
        bump,
    )]
    pub pool_signer: UncheckedAccount<'info>,

    /// CHECK: Destination account to receive SOL. Preconfigured by the owner
    /// to be a particular fund reclaimer
    #[account(
        mut,
        address = state.funding.reclaim_destination @ CctpTokenPoolError::InvalidReclaimDestination
    )]
    pub fund_reclaim_destination: UncheckedAccount<'info>,

    #[account(mut, constraint = authority.key() == state.funding.manager @ CctpTokenPoolError::InvalidFundManager)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

/// Checks that the authority and the token pool owner are the upgrade authority
pub fn allowed_admin_modify_token_pool(
    program_data: &Account<ProgramData>,
    authority: &Signer,
    state: &Account<State>,
) -> bool {
    program_data.upgrade_authority_address == Some(authority.key()) && // Only the upgrade authority of the token pool program can modify certain values of a given token pool
    state.config.owner == authority.key() // if only if the token pool is owned by the upgrade authority
}
