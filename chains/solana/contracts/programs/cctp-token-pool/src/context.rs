use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use anchor_spl::token_interface::{Mint, TokenAccount, TokenInterface};

use base_token_pool::common::*;
use ccip_common::seed;

use crate::{CctpTokenPoolError, ChainConfig, MessageAndAttestation, State};

const MAX_POOL_STATE_V: u8 = 1;
const ANCHOR_DISCRIMINATOR: usize = 8;

pub const TOKEN_MESSENGER_MINTER: Pubkey =
    solana_program::pubkey!("CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3");
pub const MESSAGE_TRANSMITTER: Pubkey =
    solana_program::pubkey!("CCTPmbSD7gX1bxKPAmg77w8oFzNFpaQiQUWD43TKaecd");

pub const MESSAGE_SENT_EVENT_SEED: &[u8] = b"cctp_message_sent_event";

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        space = ANCHOR_DISCRIMINATOR + State::INIT_SPACE,
        payer = authority,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(1, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion
    )]
    pub state: Account<'info, State>,

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
    pub state: Account<'info, State>,

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
    pub state: Account<'info, State>,

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
    )]
    pub state: Account<'info, State>,

    #[account(address = *mint.to_account_info().owner)]
    pub token_program: Interface<'info, TokenInterface>,

    #[account(mut)]
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

    #[account(
        mut,
        address = get_associated_token_address_with_program_id(&release_or_mint.receiver, &mint.key(), &token_program.key())
    )]
    pub receiver_token_account: InterfaceAccount<'info, TokenAccount>,
}

// This contains the CCTP-specific accounts that are used in the offramp. As they are too many, they don't fit inside
// a single Context struct (as Anchor would require too much memory to validate all of them), so we use a separate struct
// and do the address validations manually.
pub struct TokenOfframpRemainingAccounts<'info> {
    pub cctp_authority_pda: &'info AccountInfo<'info>,
    pub cctp_message_transmitter_account: &'info AccountInfo<'info>,
    pub cctp_token_messenger_minter: &'info AccountInfo<'info>,
    pub system_program: &'info AccountInfo<'info>,
    pub cctp_event_authority: &'info AccountInfo<'info>,
    pub cctp_message_transmitter: &'info AccountInfo<'info>,
    pub cctp_token_messenger_account: &'info AccountInfo<'info>,
    pub cctp_token_minter_account: &'info AccountInfo<'info>,
    pub cctp_local_token: &'info AccountInfo<'info>,
    pub cctp_custody_token_account: &'info AccountInfo<'info>,
    pub cctp_token_messenger_event_authority: &'info AccountInfo<'info>,
    pub cctp_remote_token_messenger_key: &'info AccountInfo<'info>,
    pub cctp_token_pair: &'info AccountInfo<'info>,
    pub cctp_used_nonces: &'info AccountInfo<'info>,
}

impl TokenOfframpRemainingAccounts<'_> {
    pub fn validate(
        &self,
        release_or_mint: &ReleaseOrMintInV1,
        remote_token_address: [u8; 32],
    ) -> Result<()> {
        let mint = release_or_mint.local_token;
        let remote_chain_selector = release_or_mint.remote_chain_selector;
        let nonce_seed = Self::first_nonce_seed(&release_or_mint.offchain_token_data)?;

        require_keys_eq!(
            self.cctp_authority_pda.key(),
            Self::get_message_transmitter_pda(&[
                b"message_transmitter_authority",
                TOKEN_MESSENGER_MINTER.as_ref(),
            ])
        );

        require_keys_eq!(
            self.cctp_message_transmitter_account.key(),
            Self::get_message_transmitter_pda(&[b"message_transmitter"])
        );

        require_keys_eq!(
            self.cctp_used_nonces.key(),
            Self::get_message_transmitter_pda(&[
                b"used_nonces",
                to_domain_seed(remote_chain_selector).as_ref(),
                nonce_seed.as_ref()
            ])
        );

        require_keys_eq!(
            self.cctp_token_messenger_minter.key(),
            TOKEN_MESSENGER_MINTER,
        );

        require_keys_eq!(self.system_program.key(), System::id());

        require_keys_eq!(
            self.cctp_event_authority.key(),
            Self::get_message_transmitter_pda(&[b"__event_authority"])
        );

        require_keys_eq!(self.cctp_message_transmitter.key(), MESSAGE_TRANSMITTER);

        require_keys_eq!(
            self.cctp_token_messenger_account.key(),
            Self::get_token_messenger_minter_pda(&[b"token_messenger"])
        );

        require_keys_eq!(
            self.cctp_token_minter_account.key(),
            Self::get_token_messenger_minter_pda(&[b"token_minter"])
        );

        require_keys_eq!(
            self.cctp_local_token.key(),
            Self::get_token_messenger_minter_pda(&[b"local_token", mint.as_ref()])
        );

        require_keys_eq!(
            self.cctp_remote_token_messenger_key.key(),
            Self::get_token_messenger_minter_pda(&[
                b"remote_token_messenger",
                &to_domain_seed(remote_chain_selector)
            ])
        );

        require_keys_eq!(
            self.cctp_token_pair.key(),
            Self::get_token_messenger_minter_pda(&[
                b"token_pair",
                &to_domain_seed(remote_chain_selector),
                remote_token_address.as_ref()
            ])
        );

        require_keys_eq!(
            self.cctp_custody_token_account.key(),
            Self::get_token_messenger_minter_pda(&[b"custody", mint.key().as_ref()])
        );

        require_keys_eq!(
            self.cctp_token_messenger_event_authority.key(),
            Self::get_token_messenger_minter_pda(&[b"__event_authority"])
        );

        Ok(())
    }

    #[inline]
    fn get_message_transmitter_pda(seeds: &[&[u8]]) -> Pubkey {
        Pubkey::find_program_address(seeds, &MESSAGE_TRANSMITTER).0
    }
    #[inline]
    fn get_token_messenger_minter_pda(seeds: &[&[u8]]) -> Pubkey {
        Pubkey::find_program_address(seeds, &TOKEN_MESSENGER_MINTER).0
    }

    fn first_nonce_seed(offchain_token_data: &[u8]) -> Result<Vec<u8>> {
        let message_and_attestation = MessageAndAttestation::try_from_slice(offchain_token_data)
            .expect("Failed to deserialize MessageAndAttestation");

        Ok(message_and_attestation
            .message
            .first_nonce()
            .to_string()
            .as_bytes()
            .to_vec())
    }
}

// TODO alternatively, remove this and store the corresponding DomainID in the ChainConfig
pub fn to_domain_seed(remote_chain_selector: u64) -> Vec<u8> {
    let bytes: &'static [u8] = match remote_chain_selector {
        // TODO update, or better just change it to a PDA
        15 => b"5",
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
    pub state: Account<'info, State>,

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
    )]
    pub chain_config: Account<'info, ChainConfig>,

    // CCTP pool-specific accounts ----------------
    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"sender_authority"],
        bump,
        seeds::program = cctp_token_messenger_minter,
    )]
    pub cctp_authority_pda: UncheckedAccount<'info>,

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        mut,
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

    /// CHECK this is not read by the pool, just forwarded to CCTP
    #[account(
        seeds = [b"remote_token_messenger", &to_domain_seed(lock_or_burn.remote_chain_selector)],
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
            &lock_or_burn.msg_nonce.to_le_bytes(),
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
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, State>,

    #[account(
        init,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
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
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE,
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
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
        ],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE + addresses.len() * RemoteAddress::INIT_SPACE,
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
    pub state: Account<'info, State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
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
    pub state: Account<'info, State>,

    #[account(
        mut,
        close = authority,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            &remote_chain_selector.to_le_bytes(),
            mint.key().as_ref(),
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
        realloc = ANCHOR_DISCRIMINATOR + State::INIT_SPACE + add.len() * 32,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub state: Account<'info, State>,

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
    pub state: Account<'info, State>,

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
