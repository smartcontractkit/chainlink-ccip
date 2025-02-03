use anchor_lang::prelude::*;
use anchor_spl::{
    token_interface::{Mint, TokenAccount},
};
use fee_quoter::messages::{GetFeeResult, SVM2AnyMessage, SVMTokenAmount};

declare_id!("CcipSender111111111111111111111111111111111");

pub const TOKEN_ADMIN_SEED: &[u8] = b"sender_token_admin";
pub const CHAIN_CONFIG_SEED: &[u8] = b"remote_chain_config";
pub const CCIP_SENDER: &[u8] = b"ccip_sender";

pub const CCIP_SEND_DISCIRIMINATOR: [u8; 8] = [108, 216, 134, 191, 249, 234, 33, 84]; // ccip_send
pub const CCIP_GET_FEE_DISCRIMINATOR: [u8; 8] = [115, 195, 235, 161, 25, 219, 60, 29]; // get_fee

/// This program an example of a CCIP Receiver Program.
/// Used to test CCIP Router execute.
#[program]
pub mod example_ccip_sender {
    use anchor_spl::{
        token_2022::spl_token_2022::{self, instruction::transfer_checked},
        token_interface,
    };
    use solana_program::{
        instruction::Instruction,
        program::{get_return_data, invoke, invoke_signed},
    };

    use super::*;

    /// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .init(ctx.accounts.authority.key(), router)
    }

    pub fn ccip_send(
        ctx: Context<CcipSend>,
        dest_chain_selector: u64,
        token_amounts: Vec<SVMTokenAmount>,
        data: Vec<u8>,
        fee_token: Pubkey,
    ) -> Result<()> {
        let message = SVM2AnyMessage {
            receiver: ctx.accounts.chain_config.extra_args_bytes.clone(),
            data,
            token_amounts: token_amounts.clone(),
            fee_token,
            extra_args: ctx.accounts.chain_config.extra_args_bytes.clone(),
        };

        // TODO: process tokens
        for _amount in token_amounts {
            // transfer tokens to this contract and approve the router to take possession during message processing
        }

        // TODO: get fee from router
        let fee: GetFeeResult = {
            let acc_infos: Vec<AccountInfo> = [
                ctx.accounts.ccip_fee_quoter_config.to_account_info(),
                ctx.accounts.ccip_fee_quoter_dest_chain.to_account_info(),
                ctx.accounts
                    .ccip_fee_quoter_billing_token_config
                    .to_account_info(),
                ctx.accounts
                    .ccip_fee_quoter_link_token_config
                    .to_account_info(),
            ]
            .to_vec();

            let acc_metas: Vec<AccountMeta> = acc_infos
                .to_vec()
                .iter()
                .flat_map(|acc_info| acc_info.to_account_metas(None))
                .collect();

            let instruction = Instruction {
                program_id: ctx.accounts.ccip_fee_quoter.key(),
                accounts: acc_metas,
                data: builder::instruction(
                    &message,
                    CCIP_GET_FEE_DISCRIMINATOR,
                    dest_chain_selector,
                ),
            };

            invoke(&instruction, &acc_infos)?;
            let (_, data) = get_return_data().unwrap();
            GetFeeResult::try_from_slice(&data)?
        };

        // TODO: if fee token is not native, transfer the fee amounts from sender to the program and approve the router
        if fee_token != Pubkey::default() {
            token_interface::transfer_checked(
                CpiContext::new_with_signer(
                    ctx.accounts.ccip_fee_token_program.to_account_info(),
                    token_interface::TransferChecked {
                        from: ctx.accounts.authority_fee_token_ata.to_account_info(),
                        mint: ctx.accounts.ccip_fee_token_mint.to_account_info(),
                        to: ctx.accounts.ccip_fee_token_user_ata.to_account_info(),
                        authority: ctx.accounts.ccip_sender.to_account_info(),
                    },
                    &[&[CCIP_SENDER], &[&[ctx.bumps.ccip_sender]]],
                ),
                fee.amount,
                ctx.accounts.ccip_fee_token_mint.decimals,
            )?;
        }

        let acc_infos: Vec<AccountInfo> = [
            ctx.accounts.ccip_config.to_account_info(),
            ctx.accounts.ccip_dest_chain_state.to_account_info(),
            ctx.accounts.ccip_sender_nonce.to_account_info(),
            ctx.accounts.ccip_sender.to_account_info(),
            ctx.accounts.system_program.to_account_info(),
            ctx.accounts.ccip_fee_token_program.to_account_info(),
            ctx.accounts.ccip_fee_token_mint.to_account_info(),
            ctx.accounts.ccip_fee_token_user_ata.to_account_info(),
            ctx.accounts.ccip_fee_token_receiver.to_account_info(),
            ctx.accounts.ccip_fee_billing_signer.to_account_info(),
            ctx.accounts.ccip_fee_quoter.to_account_info(),
            ctx.accounts.ccip_fee_quoter_config.to_account_info(),
            ctx.accounts.ccip_fee_quoter_dest_chain.to_account_info(),
            ctx.accounts
                .ccip_fee_quoter_billing_token_config
                .to_account_info(),
            ctx.accounts
                .ccip_fee_quoter_link_token_config
                .to_account_info(),
            ctx.accounts.ccip_token_pools_signer.to_account_info(),
        ]
        .to_vec();

        let acc_metas: Vec<AccountMeta> = acc_infos
            .to_vec()
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == ctx.accounts.ccip_sender.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let instruction = Instruction {
            program_id: ctx.accounts.ccip_router.key(),
            accounts: acc_metas,
            data: builder::instruction_with_token_indexes(
                &message,
                CCIP_SEND_DISCIRIMINATOR,
                dest_chain_selector,
                &[],
            ),
        };

        let seeds = &[CCIP_SENDER, &[ctx.bumps.ccip_sender]];
        let signer = &[&seeds[..]];

        invoke_signed(&instruction, &acc_infos, signer)?;

        // TODO: emit message sent event

        Ok(())
    }

    pub fn update_router(ctx: Context<UpdateConfig>, new_router: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .update_router(ctx.accounts.authority.key(), new_router)
    }

    pub fn transfer_ownership(ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
        ctx.accounts
            .state
            .transfer_ownership(ctx.accounts.authority.key(), proposed_owner)
    }

    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts
            .state
            .accept_ownership(ctx.accounts.authority.key())
    }

    pub fn withdraw_tokens(ctx: Context<WithdrawTokens>, amount: u64, decimals: u8) -> Result<()> {
        let mut ix = transfer_checked(
            &spl_token_2022::ID, // use spl-token-2022 to compile instruction - change program later
            &ctx.accounts.program_token_account.key(),
            &ctx.accounts.mint.key(),
            &ctx.accounts.to_token_account.key(),
            &ctx.accounts.token_admin.key(),
            &[],
            amount,
            decimals,
        )?;
        ix.program_id = ctx.accounts.token_program.key(); // set to user specified program

        let seeds = &[TOKEN_ADMIN_SEED, &[ctx.bumps.token_admin]];
        invoke_signed(
            &ix,
            &[
                ctx.accounts.program_token_account.to_account_info(),
                ctx.accounts.mint.to_account_info(),
                ctx.accounts.to_token_account.to_account_info(),
                ctx.accounts.token_admin.to_account_info(),
            ],
            &[&seeds[..]],
        )?;
        Ok(())
    }

    // creates initial chain config
    // only can be called for allowing a chain
    pub fn init_chain_config(
        ctx: Context<InitChainConfig>,
        _chain_selector: u64,
        recipient: Vec<u8>,
        extra_args_bytes: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .set_config(recipient, extra_args_bytes)
    }

    // updates the chain config parameters
    // configs cannot be changed through init_chain_config due to realloc needs
    pub fn update_chain_config(
        ctx: Context<UpdateChainConfig>,
        _chain_selector: u64,
        recipient: Vec<u8>,
        extra_args_bytes: Vec<u8>,
    ) -> Result<()> {
        ctx.accounts
            .chain_config
            .set_config(recipient, extra_args_bytes)
    }

    // disables remote chain by closing chain config account
    // can be re-enabled with init_chain_config
    pub fn remove_chain_config(
        _ctx: Context<RemoveChainConfig>,
        _chain_selector: u64,
    ) -> Result<()> {
        Ok(())
    }
}

const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [b"state"],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + BaseState::INIT_SPACE,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        init,
        seeds = [TOKEN_ADMIN_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR,
    )]
    /// CHECK: CPI signer for tokens
    pub token_admin: UncheckedAccount<'info>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64)]
pub struct CcipSend<'info> {
    #[account(
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        seeds = [CCIP_SENDER],
        bump,
    )]
    #[account(mut)]
    /// CHECK: sender CPI caller
    pub ccip_sender: UncheckedAccount<'info>,
    #[account(
        mut,
        associated_token::authority = authority.key(),
        associated_token::mint = ccip_fee_token_mint.key(),
        associated_token::token_program = ccip_fee_token_program.key(),
    )]
    pub authority_fee_token_ata: InterfaceAccount<'info, TokenAccount>,
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    // ------------------------
    // required ccip accounts - accounts are checked by the router
    #[account(
        address = state.router @ CcipSenderError::InvalidRouter,
    )]
    /// CHECK: used as router entry point for CPI
    pub ccip_router: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_config: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: validated during CPI
    pub ccip_dest_chain_state: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: validated during CPI
    pub ccip_sender_nonce: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_token_program: UncheckedAccount<'info>,
    #[account(
        owner = ccip_fee_token_program.key(),
    )]
    pub ccip_fee_token_mint: InterfaceAccount<'info, Mint>,
    #[account(
        mut,
        associated_token::authority = ccip_sender.key(),
        associated_token::mint = ccip_fee_token_mint.key(),
        associated_token::token_program = ccip_fee_token_program.key(),
    )]
    pub ccip_fee_token_user_ata: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        associated_token::authority = ccip_fee_billing_signer.key(),
        associated_token::mint = ccip_fee_token_mint.key(),
        associated_token::token_program = ccip_fee_token_program.key(),
    )]
    pub ccip_fee_token_receiver: InterfaceAccount<'info, TokenAccount>,
    /// CHECK: validated during CPI
    pub ccip_fee_billing_signer: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_dest_chain: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_billing_token_config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_link_token_config: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: validated during CPI
    pub ccip_token_pools_signer: UncheckedAccount<'info>,
}

#[derive(Accounts, Debug)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64, _recipient: Vec<u8>, extra_args_bytes: Vec<u8>)]
pub struct InitChainConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        init,
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        space = ANCHOR_DISCRIMINATOR + RemoteChainConfig::INIT_SPACE + extra_args_bytes.len(),
        payer = authority,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        mut,
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64, _recipient: Vec<u8>, extra_args_bytes: Vec<u8>)]
pub struct UpdateChainConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + RemoteChainConfig::INIT_SPACE + extra_args_bytes.len(),
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        mut,
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64)]
pub struct RemoveChainConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        close = authority,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        mut,
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.proposed_owner @ CcipSenderError::OnlyProposedOwner,
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts, Debug)]
pub struct WithdrawTokens<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        token::mint = mint,
        token::authority = token_admin,
        token::token_program = token_program,
    )]
    pub program_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        token::mint = mint,
        token::token_program = token_program,
    )]
    pub to_token_account: InterfaceAccount<'info, TokenAccount>,
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(address = *mint.to_account_info().owner)]
    /// CHECK: CPI to token program
    pub token_program: AccountInfo<'info>,
    #[account(
        seeds = [TOKEN_ADMIN_SEED],
        bump,
    )]
    /// CHECK: CPI signer for tokens
    pub token_admin: UncheckedAccount<'info>,
    #[account(
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
}

// BaseState contains the state for core safety checks that can be leveraged by the implementer
// Base state contains a limited size allow and deny list
// Both are included to handle the size limitations on solana
// If user wants to allow a small number of chains, consider using the allow list (disable deny list)
// If user wants to allow many chains, consider using the deny list (disable allow list)
#[account]
#[derive(InitSpace, Default, Debug)]
pub struct BaseState {
    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    pub router: Pubkey,
}

impl BaseState {
    pub fn init(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_eq!(self.owner, Pubkey::default());
        self.owner = owner;
        self.update_router(owner, router)
    }

    pub fn transfer_ownership(&mut self, owner: Pubkey, proposed_owner: Pubkey) -> Result<()> {
        require_eq!(self.owner, owner, CcipSenderError::OnlyOwner);
        self.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(&mut self, proposed_owner: Pubkey) -> Result<()> {
        require_eq!(
            self.proposed_owner,
            proposed_owner,
            CcipSenderError::OnlyProposedOwner
        );
        self.proposed_owner = Pubkey::default();
        self.owner = proposed_owner;
        Ok(())
    }

    pub fn update_router(&mut self, owner: Pubkey, router: Pubkey) -> Result<()> {
        require_keys_neq!(router, Pubkey::default(), CcipSenderError::InvalidRouter);
        require_eq!(self.owner, owner, CcipSenderError::OnlyOwner);
        self.router = router;
        Ok(())
    }
}

#[account]
#[derive(InitSpace, Default, Debug)]
pub struct RemoteChainConfig {
    #[max_len(64)]
    pub recipient: Vec<u8>, // the address to send messages to on the destination chain
    #[max_len(0)] // used to calculate InitSpace, total will include number of extra args bytes
    pub extra_args_bytes: Vec<u8>, // specifies the extraARgs to pass into ccip_send, it will be applied to every out going message for a specific chain
}

impl RemoteChainConfig {
    pub fn set_config(&mut self, recipient: Vec<u8>, extra_args_bytes: Vec<u8>) -> Result<()> {
        self.recipient = recipient;
        self.extra_args_bytes = extra_args_bytes;
        Ok(())
    }
}

pub mod builder {
    use anchor_lang::AnchorSerialize;
    use fee_quoter::messages::SVM2AnyMessage;

    pub fn instruction(
        msg: &SVM2AnyMessage,
        discriminator: [u8; 8],
        chain_selector: u64,
    ) -> Vec<u8> {
        let message = msg.try_to_vec().unwrap();
        let chain_selector_bytes = chain_selector.to_le_bytes();

        let mut data = discriminator.to_vec();
        data.extend_from_slice(&chain_selector_bytes.to_vec());
        data.extend_from_slice(&message);
        data
    }

    pub fn instruction_with_token_indexes(
        msg: &SVM2AnyMessage,
        discriminator: [u8; 8],
        chain_selector: u64,
        token_indexes: &[u8],
    ) -> Vec<u8> {
        let mut data = instruction(msg, discriminator, chain_selector);
        data.extend_from_slice(token_indexes.try_to_vec().unwrap().as_ref());
        data
    }
}

#[error_code]
pub enum CcipSenderError {
    #[msg("Invalid router address")]
    InvalidRouter,
    #[msg("Address is not owner")]
    OnlyOwner,
    #[msg("Address is not proposed_owner")]
    OnlyProposedOwner,
}

#[cfg(test)]
mod tests {
    use super::*;

    fn create_state() -> BaseState {
        BaseState {
            owner: Pubkey::new_unique(),
            ..BaseState::default()
        }
    }

    #[test]
    fn ownership() {
        let mut state = create_state();
        let next_owner = Pubkey::new_unique();

        // only owner can propose
        assert_eq!(
            state
                .transfer_ownership(Pubkey::new_unique(), Pubkey::new_unique())
                .unwrap_err(),
            CcipSenderError::OnlyOwner.into()
        );
        state.transfer_ownership(state.owner, next_owner).unwrap();

        // only proposed_owner can accept
        assert_eq!(
            state.accept_ownership(Pubkey::new_unique()).unwrap_err(),
            CcipSenderError::OnlyProposedOwner.into(),
        );
        state.accept_ownership(next_owner).unwrap();
    }

    #[test]
    fn router() {
        let mut state = create_state();

        assert_eq!(
            state
                .update_router(state.owner, Pubkey::default())
                .unwrap_err(),
            CcipSenderError::InvalidRouter.into(),
        );
        assert_eq!(
            state
                .update_router(Pubkey::new_unique(), Pubkey::new_unique())
                .unwrap_err(),
            CcipSenderError::OnlyOwner.into(),
        );
        state
            .update_router(state.owner, Pubkey::new_unique())
            .unwrap();
    }
}
