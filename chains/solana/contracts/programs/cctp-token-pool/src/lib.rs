use anchor_lang::prelude::*;
use anchor_lang::solana_program::instruction::{AccountMeta, Instruction};
use solana_program::program::{get_return_data, invoke_signed};

use base_token_pool::common::*;
use base_token_pool::rate_limiter::*;

declare_id!("CCitPr8yZbN8zEBEdwju8bnGgKMYcz6XSTbU61CMedj");

pub const RECEIVE_MESSAGE_DISCRIMINATOR: [u8; 8] = [38, 144, 127, 225, 31, 225, 238, 25]; // global:receive_message
pub const DEPOSIT_FOR_BURN_WITH_CALLER_DISCRIMINATOR: [u8; 8] =
    [167, 222, 19, 114, 85, 21, 14, 118]; // global:deposit_for_burn_with_caller
pub const RECLAIM_EVENT_ACCOUNT_DISCRIMINATOR: [u8; 8] = [94, 198, 180, 159, 131, 236, 15, 174]; // global:reclaim_event_account

pub mod context;
use crate::context::*;

mod cctp_message;
use crate::cctp_message::*;

mod token_pool_extra_data;
use crate::token_pool_extra_data::*;

mod derive;

const SOLANA_DOMAIN_ID: u32 = 5; // Circle's CCTP domain ID for Solana is always 5, see https://developers.circle.com/stablecoins/supported-domains

#[program]
pub mod cctp_token_pool {
    use std::str::FromStr;

    use super::*;

    pub fn initialize(
        ctx: Context<InitializeTokenPool>,
        router: Pubkey,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        ctx.accounts.state.set_inner(State {
            version: 1,
            config: BaseConfig::init(
                &ctx.accounts.mint,
                ctx.program_id.key(),
                ctx.accounts.authority.key(),
                router,
                rmn_remote,
            ),
        });

        Ok(())
    }

    /// Returns the program type (name) and version.
    /// Used by offchain code to easily determine which program & version is being interacted with.
    ///
    /// # Arguments
    /// * `ctx` - The context
    pub fn type_version(_ctx: Context<Empty>) -> Result<String> {
        let response = env!("CCIP_BUILD_TYPE_VERSION").to_string();
        msg!("{}", response);
        Ok(response)
    }

    // TODO remove this method, only used for internal testing
    pub fn set_rmn_remote(ctx: Context<SetConfig>, rmn_remote: Pubkey) -> Result<()> {
        ctx.accounts.state.config.rmn_remote = rmn_remote;
        Ok(())
    }

    pub fn transfer_ownership(ctx: Context<SetConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts.state.config.transfer_ownership(proposed_owner)
    }

    // shared func signature with other programs
    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts.state.config.accept_ownership()
    }

    // set_router changes the expected signers for mint/release + burn/lock method calls
    // this is used to update the router address
    pub fn set_router(ctx: Context<SetConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .config
            .set_router(new_router, ctx.program_id)
    }

    // initialize remote config (with no remote pools as it must be zero sized)
    pub fn init_chain_remote_config(
        ctx: Context<InitializeChainConfig>,
        remote_chain_selector: u64,
        mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        require!(
            cfg.pool_addresses.is_empty(),
            CcipTokenPoolError::NonemptyPoolAddressesInit
        );

        ctx.accounts
            .chain_config
            .base
            .set(remote_chain_selector, mint, cfg)
    }

    // edit remote config
    pub fn edit_chain_remote_config(
        ctx: Context<EditChainConfigDynamicSize>,
        remote_chain_selector: u64,
        mint: Pubkey,
        cfg: RemoteConfig,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .base
            .set(remote_chain_selector, mint, cfg)
    }

    pub fn edit_chain_remote_config_cctp(
        ctx: Context<EditChainConfig>,
        _remote_chain_selector: u64,
        _mint: Pubkey,
        cfg: CctpChain,
    ) -> Result<()> {
        ctx.accounts.chain_config.cctp = cfg;
        emit!(RemoteChainCctpConfigChanged {
            config: ctx.accounts.chain_config.cctp
        });
        Ok(())
    }

    // Add remote pool addresses
    pub fn append_remote_pool_addresses(
        ctx: Context<AppendRemotePoolAddresses>,
        remote_chain_selector: u64,
        _mint: Pubkey,
        addresses: Vec<RemoteAddress>,
    ) -> Result<()> {
        ctx.accounts.chain_config.base.append_remote_pool_addresses(
            remote_chain_selector,
            _mint,
            addresses,
        )
    }

    // set rate limit
    pub fn set_chain_rate_limit(
        ctx: Context<EditChainConfig>,
        remote_chain_selector: u64,
        mint: Pubkey,
        inbound: RateLimitConfig,
        outbound: RateLimitConfig,
    ) -> Result<()> {
        ctx.accounts.chain_config.base.set_chain_rate_limit(
            remote_chain_selector,
            mint,
            inbound,
            outbound,
        )
    }

    // delete chain config
    pub fn delete_chain_config(
        _ctx: Context<DeleteChainConfig>,
        remote_chain_selector: u64,
        mint: Pubkey,
    ) -> Result<()> {
        emit!(RemoteChainRemoved {
            chain_selector: remote_chain_selector,
            mint
        });
        Ok(())
    }

    pub fn configure_allow_list(
        ctx: Context<AddToAllowList>,
        add: Vec<Pubkey>,
        enabled: bool,
    ) -> Result<()> {
        ctx.accounts.state.config.list_enabled = enabled;
        let list = &mut ctx.accounts.state.config.allow_list;
        for key in add {
            require!(
                !list.contains(&key),
                CcipTokenPoolError::AllowlistKeyAlreadyExisted
            );
            list.push(key);
        }

        Ok(())
    }

    pub fn remove_from_allow_list(
        ctx: Context<RemoveFromAllowlist>,
        remove: Vec<Pubkey>,
    ) -> Result<()> {
        let list = &mut ctx.accounts.state.config.allow_list;
        for key in remove {
            require!(
                list.contains(&key),
                CcipTokenPoolError::AllowlistKeyDidNotExist
            );
            list.retain(|k| k != &key);
        }

        Ok(())
    }

    pub fn release_or_mint_tokens<'info>(
        ctx: Context<'_, '_, 'info, 'info, TokenOfframp<'info>>,
        release_or_mint: ReleaseOrMintInV1,
    ) -> Result<ReleaseOrMintOutV1> {
        require!(
            !release_or_mint.offchain_token_data.is_empty(),
            CctpTokenPoolError::InvalidTokenData
        );

        let msg_att = MessageAndAttestation::try_from_slice(&release_or_mint.offchain_token_data)
            .map_err(|_| CctpTokenPoolError::InvalidTokenData)?;

        let remote_token_address_bytes =
            to_solana_pubkey(&ctx.accounts.chain_config.base.remote.token_address).to_bytes();

        let remaining = parse_remaining_release_accounts(
            ctx.remaining_accounts,
            &release_or_mint,
            ctx.accounts.chain_config.cctp.domain_id,
            remote_token_address_bytes,
        )?;

        let parsed_amount = to_svm_token_amount(
            release_or_mint.amount,
            ctx.accounts.chain_config.base.remote.decimals,
            ctx.accounts.state.config.decimals,
        )?;

        let BaseChain {
            remote,
            inbound_rate_limit,
            ..
        } = &mut ctx.accounts.chain_config.base;

        validate_release_or_mint(
            &release_or_mint,
            parsed_amount,
            ctx.accounts.state.config.mint,
            &remote.pool_addresses,
            inbound_rate_limit,
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
        )?;

        validate_cctp_and_ccip_messages(&msg_att.message, &release_or_mint)?;

        let message_and_attestation_bytes = msg_att.try_to_vec().unwrap();
        cctp_receive_message(&ctx, &remaining, &message_and_attestation_bytes)?;
        transfer_cctp_tokens(&ctx, parsed_amount, ctx.accounts.state.config.decimals)?;

        emit!(Minted {
            sender: ctx.accounts.authority.key(),
            recipient: release_or_mint.receiver,
            amount: parsed_amount,
            mint: ctx.accounts.state.config.mint,
        });

        Ok(ReleaseOrMintOutV1 {
            destination_amount: parsed_amount,
        })
    }

    pub fn lock_or_burn_tokens(
        ctx: Context<TokenOnramp>,
        lock_or_burn: LockOrBurnInV1,
    ) -> Result<LockOrBurnOutV1> {
        validate_lock_or_burn(
            &lock_or_burn,
            ctx.accounts.state.config.mint,
            &mut ctx.accounts.chain_config.base.outbound_rate_limit,
            ctx.accounts.state.config.list_enabled,
            &ctx.accounts.state.config.allow_list,
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
        )?;

        require_eq!(
            lock_or_burn.receiver.len(),
            32,
            CctpTokenPoolError::InvalidReceiver
        );

        let cctp_nonce = cctp_deposit_for_burn_with_caller(&ctx, &lock_or_burn)?;

        let cctp_event_data = ctx.accounts.cctp_message_sent_event.try_borrow_data()?;
        // The first 8 bytes are the discriminator, so we skip them
        // Then, the next 32 bytes are the rent payer, which we also skip
        let cctp_message_bytes_len = &cctp_event_data[40..44]; // the next 4 bytes encode the message length
        let cctp_message_bytes = &cctp_event_data[44..]; // then, the rest is the message bytes
        require_eq!(
            cctp_message_bytes.len(),
            u32::from_le_bytes(cctp_message_bytes_len.try_into().unwrap()) as usize,
            CctpTokenPoolError::InvalidMessageSentEventAccount
        );

        emit!(CcipCctpMessageSentEvent {
            original_sender: lock_or_burn.original_sender,
            remote_chain_selector: lock_or_burn.remote_chain_selector,
            msg_total_nonce: lock_or_burn.msg_total_nonce,
            event_address: ctx.accounts.cctp_message_sent_event.key(),
            message_sent_bytes: cctp_message_bytes.to_vec(),
            source_domain: SOLANA_DOMAIN_ID,
            cctp_nonce,
        });

        emit!(Burned {
            sender: ctx.accounts.authority.key(),
            amount: lock_or_burn.amount,
            mint: ctx.accounts.state.config.mint,
        });

        let extra_data = TokenPoolExtraData {
            nonce: cctp_nonce,
            source_domain: SOLANA_DOMAIN_ID,
        };

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.base.remote.token_address.clone(),
            dest_pool_data: extra_data.abi_encode().to_vec(),
        })
    }

    pub fn reclaim_event_account(
        ctx: Context<ReclaimEventAccount>,
        mint: Pubkey,
        _original_sender: Pubkey,
        _remote_chain_selector: u64,
        _msg_nonce: u64,
        attestation: Vec<u8>,
    ) -> Result<()> {
        let attestation_bytes = attestation.try_to_vec()?;
        let mut instruction_data =
            Vec::with_capacity(ANCHOR_DISCRIMINATOR + attestation_bytes.len());
        instruction_data.extend_from_slice(&RECLAIM_EVENT_ACCOUNT_DISCRIMINATOR);
        instruction_data.extend_from_slice(attestation_bytes.as_slice());

        let acc_infos = vec![
            ctx.accounts.pool_signer.to_account_info(),
            ctx.accounts.message_transmitter.to_account_info(),
            ctx.accounts.message_sent_event_account.to_account_info(),
        ];

        let acc_metas: Vec<AccountMeta> = acc_infos
            .iter()
            .flat_map(|acc_info| {
                let is_signer = acc_info.key() == ctx.accounts.pool_signer.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let instruction = Instruction {
            program_id: ctx.accounts.cctp_message_transmitter.key(),
            accounts: acc_metas,
            data: instruction_data,
        };

        let signer_seeds: &[&[u8]] = &[
            POOL_SIGNER_SEED,
            &mint.key().to_bytes(),
            &[ctx.bumps.pool_signer],
        ];

        invoke_signed(&instruction, &acc_infos, &[signer_seeds])?;

        Ok(())
    }

    pub fn derive_accounts_release_or_mint_tokens<'info>(
        ctx: Context<'_, '_, 'info, 'info, Empty>,
        stage: String,
        release_or_mint: ReleaseOrMintInV1,
    ) -> Result<DeriveAccountsResponse> {
        let stage = derive::release_or_mint::OfframpDeriveStage::from_str(&stage)?;

        match stage {
            derive::release_or_mint::OfframpDeriveStage::RetrieveChainConfig => {
                derive::release_or_mint::retrieve_chain_config(&release_or_mint)
            }
            derive::release_or_mint::OfframpDeriveStage::BuildDynamicAccounts => {
                derive::release_or_mint::build_dynamic_accounts(ctx, &release_or_mint)
            }
        }
    }

    pub fn derive_accounts_lock_or_burn_tokens<'info>(
        ctx: Context<'_, '_, 'info, 'info, Empty>,
        stage: String,
        lock_or_burn: LockOrBurnInV1,
    ) -> Result<DeriveAccountsResponse> {
        let stage = derive::lock_or_burn::OnrampDeriveStage::from_str(&stage)?;

        match stage {
            derive::lock_or_burn::OnrampDeriveStage::RetrieveChainConfig => {
                derive::lock_or_burn::retrieve_chain_config(&lock_or_burn)
            }
            derive::lock_or_burn::OnrampDeriveStage::BuildDynamicAccounts => {
                derive::lock_or_burn::build_dynamic_accounts(ctx, &lock_or_burn)
            }
        }
    }
}

fn parse_remaining_release_accounts<'info>(
    remaining: &'info [AccountInfo<'info>],
    release_or_mint: &ReleaseOrMintInV1,
    cctp_domain_id: u32,
    remote_token_address: [u8; 32],
) -> Result<TokenOfframpRemainingAccounts<'info>> {
    let mut remaining_accounts = remaining.iter();
    let result = TokenOfframpRemainingAccounts {
        cctp_authority_pda: next_account_info(&mut remaining_accounts)?,
        cctp_message_transmitter_account: next_account_info(&mut remaining_accounts)?,
        cctp_token_messenger_minter: next_account_info(&mut remaining_accounts)?,
        system_program: next_account_info(&mut remaining_accounts)?,
        cctp_event_authority: next_account_info(&mut remaining_accounts)?,
        cctp_message_transmitter: next_account_info(&mut remaining_accounts)?,
        cctp_token_messenger_account: next_account_info(&mut remaining_accounts)?,
        cctp_token_minter_account: next_account_info(&mut remaining_accounts)?,
        cctp_local_token: next_account_info(&mut remaining_accounts)?,
        cctp_custody_token_account: next_account_info(&mut remaining_accounts)?,
        cctp_token_messenger_event_authority: next_account_info(&mut remaining_accounts)?,
        cctp_remote_token_messenger_key: next_account_info(&mut remaining_accounts)?,
        cctp_token_pair: next_account_info(&mut remaining_accounts)?,
        cctp_used_nonces: next_account_info(&mut remaining_accounts)?,
    };

    result.validate(release_or_mint, cctp_domain_id, remote_token_address)?;

    Ok(result)
}

fn to_solana_pubkey(address: &RemoteAddress) -> Pubkey {
    let mut solana_address = [0; 32];
    let remote_bytes = address.len();
    solana_address[32 - remote_bytes..].copy_from_slice(address);
    Pubkey::new_from_array(solana_address)
}

fn cctp_deposit_for_burn_with_caller(
    ctx: &Context<TokenOnramp>,
    lock_or_burn: &LockOrBurnInV1,
) -> Result<u64> {
    let mint_recipient = Pubkey::new_from_array(lock_or_burn.receiver[0..32].try_into().unwrap());

    let destination_domain = ctx.accounts.chain_config.cctp.domain_id;

    let deposit_for_burn_params = DepositForBurnWithCallerParams {
        amount: lock_or_burn.amount,
        destination_domain,
        mint_recipient,
        destination_caller: ctx.accounts.chain_config.cctp.destination_caller,
    };
    let deposit_for_burn_params_bytes = deposit_for_burn_params.try_to_vec().unwrap();

    let mut instruction_data = Vec::with_capacity(8 + deposit_for_burn_params_bytes.len());
    instruction_data.extend_from_slice(&DEPOSIT_FOR_BURN_WITH_CALLER_DISCRIMINATOR);
    instruction_data.extend_from_slice(&deposit_for_burn_params_bytes);

    let acc_infos = vec![
        ctx.accounts.pool_signer.to_account_info(),
        ctx.accounts.pool_signer.to_account_info(),
        ctx.accounts.cctp_authority_pda.to_account_info(),
        ctx.accounts.pool_token_account.to_account_info(),
        ctx.accounts
            .cctp_message_transmitter_account
            .to_account_info(),
        ctx.accounts.cctp_token_messenger_account.to_account_info(),
        ctx.accounts
            .cctp_remote_token_messenger_key
            .to_account_info(),
        ctx.accounts.cctp_token_minter_account.to_account_info(),
        ctx.accounts.cctp_local_token.to_account_info(),
        ctx.accounts.mint.to_account_info(),
        ctx.accounts.cctp_message_sent_event.to_account_info(),
        ctx.accounts.cctp_message_transmitter.to_account_info(),
        ctx.accounts.cctp_token_messenger_minter.to_account_info(),
        ctx.accounts.token_program.to_account_info(),
        ctx.accounts.system_program.to_account_info(),
        ctx.accounts.cctp_event_authority.to_account_info(),
        ctx.accounts.cctp_token_messenger_minter.to_account_info(),
    ];

    let signer_keys = [
        ctx.accounts.pool_signer.key(),
        ctx.accounts.cctp_message_sent_event.key(),
    ];

    let acc_metas: Vec<AccountMeta> = acc_infos
        .iter()
        .flat_map(|acc_info| {
            let is_signer = signer_keys.contains(acc_info.key);
            acc_info.to_account_metas(Some(is_signer))
        })
        .collect();

    let instruction = Instruction {
        program_id: ctx.accounts.cctp_token_messenger_minter.key(),
        accounts: acc_metas,
        data: instruction_data,
    };

    let pool_signer_seeds: &[&[u8]] = &[
        POOL_SIGNER_SEED,
        &ctx.accounts.mint.key().to_bytes(),
        &[ctx.bumps.pool_signer],
    ];

    let message_sent_event_seeds: &[&[u8]] = &[
        MESSAGE_SENT_EVENT_SEED,
        &lock_or_burn.original_sender.to_bytes(),
        &lock_or_burn.remote_chain_selector.to_le_bytes(),
        &lock_or_burn.msg_total_nonce.to_le_bytes(),
        &[ctx.bumps.cctp_message_sent_event],
    ];

    invoke_signed(
        &instruction,
        &acc_infos,
        &[pool_signer_seeds, message_sent_event_seeds],
    )?;

    let (_, cctp_nonce_bytes) = get_return_data().unwrap();
    Ok(u64::from_le_bytes(
        cctp_nonce_bytes.as_slice().try_into().unwrap(),
    ))
}

// Checks consistency in the data between the CCTP and CCIP messages
fn validate_cctp_and_ccip_messages(
    cctp_msg: &CctpMessage,
    release_or_mint: &ReleaseOrMintInV1,
) -> Result<()> {
    let source_extra_data = TokenPoolExtraData::abi_decode(&release_or_mint.source_pool_data)?;

    require_gte!(
        cctp_msg.data.len(),
        CctpMessage::MINIMUM_LENGTH,
        CctpTokenPoolError::MalformedCctpMessage
    );

    require_eq!(
        cctp_msg.source_domain(),
        source_extra_data.source_domain,
        CctpTokenPoolError::InvalidSourceDomain
    );

    require_eq!(
        cctp_msg.destination_domain(),
        SOLANA_DOMAIN_ID,
        CctpTokenPoolError::InvalidDestDomain
    );

    require_eq!(
        cctp_msg.nonce(),
        source_extra_data.nonce,
        CctpTokenPoolError::InvalidNonce
    );

    Ok(())
}

fn cctp_receive_message<'info>(
    ctx: &Context<'_, '_, 'info, 'info, TokenOfframp<'info>>,
    remaining: &TokenOfframpRemainingAccounts<'info>,
    message_and_attestation_bytes: &[u8],
) -> Result<()> {
    let mut instruction_data = Vec::with_capacity(
        RECEIVE_MESSAGE_DISCRIMINATOR.len() + message_and_attestation_bytes.len(),
    );
    instruction_data.extend_from_slice(&RECEIVE_MESSAGE_DISCRIMINATOR);
    instruction_data.extend_from_slice(message_and_attestation_bytes);

    let acc_infos = vec![
        ctx.accounts.pool_signer.clone().to_account_info(),
        ctx.accounts.pool_signer.clone().to_account_info(),
        remaining.cctp_authority_pda.to_account_info(),
        remaining.cctp_message_transmitter_account.to_account_info(),
        remaining.cctp_used_nonces.to_account_info(),
        remaining.cctp_token_messenger_minter.to_account_info(),
        remaining.system_program.to_account_info(),
        remaining.cctp_event_authority.to_account_info(),
        remaining.cctp_message_transmitter.to_account_info(),
        remaining.cctp_token_messenger_account.to_account_info(),
        remaining.cctp_remote_token_messenger_key.to_account_info(),
        remaining.cctp_token_minter_account.to_account_info(),
        remaining.cctp_local_token.to_account_info(),
        remaining.cctp_token_pair.to_account_info(),
        ctx.accounts.pool_token_account.clone().to_account_info(), // always mint to the tp ATA first
        remaining.cctp_custody_token_account.to_account_info(),
        ctx.accounts.token_program.clone().to_account_info(),
        remaining
            .cctp_token_messenger_event_authority
            .to_account_info(),
        remaining.cctp_token_messenger_minter.to_account_info(),
    ];

    let acc_metas: Vec<AccountMeta> = acc_infos
        .iter()
        .flat_map(|acc_info| {
            let is_signer = acc_info.key() == ctx.accounts.pool_signer.key();
            acc_info.to_account_metas(Some(is_signer))
        })
        .collect();

    let instruction = Instruction {
        program_id: remaining.cctp_message_transmitter.key(),
        accounts: acc_metas,
        data: instruction_data,
    };

    let signer_seeds: &[&[u8]] = &[
        POOL_SIGNER_SEED,
        &ctx.accounts.mint.key().to_bytes(),
        &[ctx.bumps.pool_signer],
    ];

    invoke_signed(&instruction, &acc_infos, &[signer_seeds])
        .map_err(|_| CctpTokenPoolError::FailedCctpCpi.into())
}

fn transfer_cctp_tokens(ctx: &Context<TokenOfframp>, amount: u64, decimals: u8) -> Result<()> {
    let seeds = &[
        POOL_SIGNER_SEED,
        &ctx.accounts.mint.key().to_bytes(),
        &[ctx.bumps.pool_signer],
    ];
    let signer_seeds: &[&[&[u8]]] = &[seeds];
    let cpi_ctx = CpiContext::new_with_signer(
        ctx.accounts.token_program.to_account_info(),
        anchor_spl::token::TransferChecked {
            from: ctx.accounts.pool_token_account.to_account_info(),
            to: ctx.accounts.receiver_token_account.to_account_info(),
            authority: ctx.accounts.pool_signer.to_account_info(),
            mint: ctx.accounts.mint.to_account_info(),
        },
        signer_seeds,
    );

    anchor_spl::token::transfer_checked(cpi_ctx, amount, decimals)
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct MessageAndAttestation {
    pub message: CctpMessage,
    pub attestation: Vec<u8>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct DepositForBurnWithCallerParams {
    pub amount: u64,
    pub destination_domain: u32,
    pub mint_recipient: Pubkey,
    pub destination_caller: Pubkey,
}

// NOTE: accounts derivations must be native to program - will cause ownership check issues if imported
// wraps functionality from shared Config and Chain types
#[account]
#[derive(InitSpace)]
pub struct State {
    pub version: u8,
    pub config: BaseConfig,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy, InitSpace)]
pub struct CctpChain {
    // Domain ID for CCTP, used to identify the chain. This is a sequential number starting from 0.
    // Using u32 here because it's what CCTP uses in its Params structs.
    pub domain_id: u32,
    pub destination_caller: Pubkey, // The allowed caller to invoke CCTP's receive_message on the dest chain
}

#[account]
#[derive(InitSpace)]
pub struct ChainConfig {
    pub base: BaseChain,
    pub cctp: CctpChain,
}

#[event]
pub struct RemoteChainCctpConfigChanged {
    pub config: CctpChain,
}

#[event]
pub struct CcipCctpMessageSentEvent {
    // Seeds for the CCTP message sent event account
    pub original_sender: Pubkey,
    pub remote_chain_selector: u64,
    pub msg_total_nonce: u64,

    // Actual event account address, derived from the seeds above
    pub event_address: Pubkey,

    // CCTP values identifying the message
    pub source_domain: u32, // The source chain domain ID, which for Solana is always 5
    pub cctp_nonce: u64,

    // CCTP message bytes, used to get the attestation offchain and receive the message on dest
    pub message_sent_bytes: Vec<u8>,
}

#[error_code]
pub enum CctpTokenPoolError {
    #[msg("Invalid token data")]
    InvalidTokenData = 6000, // offset for CctpTokenPoolErrors, so they don't overlap with errors of other CCIP programs
    #[msg("Invalid receiver")]
    InvalidReceiver,
    #[msg("Invalid source domain")]
    InvalidSourceDomain,
    #[msg("Invalid destination domain")]
    InvalidDestDomain,
    #[msg("Invalid nonce")]
    InvalidNonce,
    #[msg("CCTP message is malformed or too short")]
    MalformedCctpMessage,
    #[msg("Invalid Token Messenger Minter")]
    InvalidTokenMessengerMinter,
    #[msg("Invalid Message Transmitter")]
    InvalidMessageTransmitter,
    #[msg("Invalid Message Sent Event Account")]
    InvalidMessageSentEventAccount,
    #[msg("Invalid Token Pool Extra Data")]
    InvalidTokenPoolExtraData,
    #[msg("Failed CCTP CPI")]
    FailedCctpCpi,
}
