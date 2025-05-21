use anchor_lang::prelude::*;
use anchor_lang::solana_program::{
    instruction::{AccountMeta, Instruction},
    program::invoke_signed,
};
use base_token_pool::{common::*, rate_limiter::*};

declare_id!("UsDcjWP1QYak64nKofkmA5ZWPrceoXsyZKiXZMvjk5V");

pub const RECEIVE_MESSAGE_DISCRIMINATOR: [u8; 8] = [0xd2, 0xf3, 0x33, 0xf6, 0x02, 0xf2, 0xea, 0x35]; // global:receive_message
pub const DEPOSIT_FOR_BURN_WITH_CALLER_DISCRIMINATOR: [u8; 8] = [0x90, 0x88, 0x2a, 0x0c, 0xa4, 0x58, 0x3a, 0x2a]; // global:deposit_for_burn_with_caller

pub const TOKEN_MESSENGER_MINTER: Pubkey = solana_program::pubkey!("CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3");
pub const MESSAGE_TRANSMITTER: Pubkey = solana_program::pubkey!("CCTPmbSD7gX1bxKPAmg77w8oFzNFpaQiQUWD43TKaecd");

pub mod context;
use crate::context::*;

#[program]
pub mod usdc_token_pool {
    use super::*;

    pub fn initialize(
        ctx: Context<InitializeTokenPool>,
        router: Pubkey,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        ctx.accounts.state.set_inner(burnmint_token_pool::State {
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

    pub fn type_version(_ctx: Context<TypeVersion>) -> Result<String> {
        let response = env!("CCIP_BUILD_TYPE_VERSION").to_string();
        msg!("{}", response);
        Ok(response)
    }

    pub fn transfer_ownership(ctx: Context<SetConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts.state.config.transfer_ownership(proposed_owner)
    }

    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts.state.config.accept_ownership()
    }

    pub fn set_router(ctx: Context<SetConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .config
            .set_router(new_router, ctx.program_id)
    }


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

    pub fn set_chain_rate_limit(
        ctx: Context<SetChainRateLimit>,
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

    pub fn release_or_mint_tokens(
        ctx: Context<TokenOfframp>,
        release_or_mint: ReleaseOrMintInV1,
    ) -> Result<ReleaseOrMintOutV1> {
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

        require!(
            !release_or_mint.offchain_token_data.is_empty(),
            UsdcTokenPoolError::InvalidTokenData
        );

        let message_and_attestation: MessageAndAttestation =
            try_from_slice_unchecked(&release_or_mint.offchain_token_data)
                .map_err(|_| UsdcTokenPoolError::InvalidTokenData)?;


        let message_and_attestation_bytes = message_and_attestation.try_to_vec().unwrap();
        let mut instruction_data = Vec::with_capacity(8 + message_and_attestation_bytes.len());
        instruction_data.extend_from_slice(&RECEIVE_MESSAGE_DISCRIMINATOR);
        instruction_data.extend_from_slice(&message_and_attestation_bytes);

        let accounts = vec![
            AccountMeta::new(ctx.accounts.cctp_payer.key(), true),
            AccountMeta::new_readonly(ctx.accounts.cctp_payer.key(), true),
            AccountMeta::new_readonly(ctx.accounts.authority.key(), false),
            AccountMeta::new_readonly(ctx.accounts.message_transmitter.key(), false),
            AccountMeta::new(ctx.accounts.used_nonces.key(), false),
            AccountMeta::new_readonly(ctx.accounts.receiver.key(), false),
            AccountMeta::new_readonly(ctx.accounts.system_program.key(), false),
        ];

        let instruction = Instruction {
            program_id: ctx.accounts.message_transmitter.key(),
            accounts,
            data: instruction_data,
        };

        invoke_signed(
            &instruction,
            &[
                ctx.accounts.cctp_payer.to_account_info(),
                ctx.accounts.cctp_payer.to_account_info(),
                ctx.accounts.authority.to_account_info(),
                ctx.accounts.message_transmitter.to_account_info(),
                ctx.accounts.used_nonces.to_account_info(),
                ctx.accounts.receiver.to_account_info(),
                ctx.accounts.system_program.to_account_info(),
            ],
            &[],
        )?;

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
            UsdcTokenPoolError::InvalidReceiver
        );
        let mint_recipient = Pubkey::new_from_array(lock_or_burn.receiver[0..32].try_into().unwrap());

        let destination_domain = match lock_or_burn.destination_chain_selector {
        // TODO update this mapping to be from ChainSelector -> DomainID
            1 => 0, // Ethereum Mainnet
            2 => 1, // Optimism
            3 => 2, // Arbitrum
            4 => 3, // Avalanche
            _ => return Err(UsdcTokenPoolError::InvalidDomain.into()),
        };

        let deposit_for_burn_params = DepositForBurnWithCallerParams {
            amount: lock_or_burn.amount,
            destination_domain,
            mint_recipient,
            destination_caller: Pubkey::default(), // TODO: Default for now, this should be the destination token minter contract
        };

        let deposit_for_burn_params_bytes = deposit_for_burn_params.try_to_vec().unwrap();
        let mut instruction_data = Vec::with_capacity(8 + deposit_for_burn_params_bytes.len());
        instruction_data.extend_from_slice(&DEPOSIT_FOR_BURN_WITH_CALLER_DISCRIMINATOR);
        instruction_data.extend_from_slice(&deposit_for_burn_params_bytes);

        let accounts = vec![
            AccountMeta::new(ctx.accounts.cctp_event_rent_payer.key(), true),
            AccountMeta::new_readonly(ctx.accounts.authority.key(), true),
            AccountMeta::new_readonly(ctx.accounts.token_messenger.key(), false),
            AccountMeta::new_readonly(ctx.accounts.message_transmitter.key(), false),
            AccountMeta::new_readonly(ctx.accounts.mint.key(), false),
            AccountMeta::new_readonly(ctx.accounts.pool_token_account.key(), false),
            AccountMeta::new_readonly(ctx.accounts.system_program.key(), false),
        ];

        let instruction = Instruction {
            program_id: ctx.accounts.token_messenger.key(),
            accounts,
            data: instruction_data,
        };

        let pool_signer_seeds = &[
            POOL_SIGNER_SEED,
            ctx.accounts.mint.key().as_ref(),
            &[ctx.bumps.pool_signer],
        ];

        invoke_signed(
            &instruction,
            &[
                ctx.accounts.cctp_event_rent_payer.to_account_info(),
                ctx.accounts.authority.to_account_info(),
                ctx.accounts.token_messenger.to_account_info(),
                ctx.accounts.message_transmitter.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.pool_token_account.to_account_info(),
                ctx.accounts.system_program.to_account_info(),
            ],
            &[&pool_signer_seeds],
        )?;

        emit!(Burned {
            sender: ctx.accounts.authority.key(),
            amount: lock_or_burn.amount,
            mint: ctx.accounts.state.config.mint,
        });

        Ok(LockOrBurnOutV1 {
            dest_token_address: ctx.accounts.chain_config.base.remote.token_address.clone(),
            dest_pool_data: {
                let mut abi_encoded_decimals = vec![0u8; 32];
                abi_encoded_decimals[31] = ctx.accounts.state.config.decimals;
                abi_encoded_decimals
            },
        })
    }


}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct MessageAndAttestation {
    pub message: Vec<u8>,
    pub attestation: Vec<u8>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct DepositForBurnWithCallerParams {
    pub amount: u64,
    pub destination_domain: u32,
    pub mint_recipient: Pubkey,
    pub destination_caller: Pubkey,
}



#[error_code]
pub enum UsdcTokenPoolError {
    #[msg("Invalid token data")]
    InvalidTokenData,
    #[msg("Invalid receiver")]
    InvalidReceiver,
    #[msg("Invalid domain")]
    InvalidDomain,
    #[msg("Invalid token messenger")]
    InvalidTokenMessenger,
    #[msg("Invalid message transmitter")]
    InvalidMessageTransmitter,
}
