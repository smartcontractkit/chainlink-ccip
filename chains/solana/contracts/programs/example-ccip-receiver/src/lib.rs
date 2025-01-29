use anchor_lang::prelude::*;
use arrayvec::arrayvec;
declare_id!("CcipReceiver1111111111111111111111111111111");

pub const EXTERNAL_EXECUTION_CONFIG_SEED: &[u8] = b"external_execution_config";

/// This program an example of a CCIP Receiver Program.
/// Used to test CCIP Router execute.
#[program]
pub mod example_ccip_receiver {
    use super::*;

    /// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        require_keys_neq!(router, Pubkey::default());
        let state = &mut ctx.accounts.state;
        state.router = router;
        state.owner = ctx.accounts.authority.key();
        Ok(())
    }

    /// This function is called by the CCIP Router to execute the CCIP message.
    /// The method name needs to be ccip_receive with Anchor encoding,
    /// if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]
    /// You can send as many accounts as you need, specifying if mutable or not.
    /// But none of them could be an init, realloc or close.
    pub fn ccip_receive(_ctx: Context<CcipReceive>, message: Any2SVMMessage) -> Result<()> {
        // ---------------------------------------
        // implement functionality here
        // ---------------------------------------

        emit!(MessageReceived {
            message_id: message.message_id
        });

        Ok(())
    }

    pub fn enable_chain(ctx: Context<UpdateConfig>, chain_selector: u64) -> Result<()> {
        let state = &mut ctx.accounts.state;
        if state.enabled_chains.remaining_capacity() == 0 {
            state.enabled_chains.extend(&[0]);
        }

        state.enabled_chains.push(chain_selector);
        Ok(())
    }

    pub fn disable_chain(ctx: Context<UpdateConfig>, chain_selector: u64) -> Result<()> {
        let state = &mut ctx.accounts.state;
        let index = state.enabled_chains.binary_search(&chain_selector);
        if let Ok(index) = index {
            state.enabled_chains.remove(index);
        }
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
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(message: Any2SVMMessage)]
pub struct CcipReceive<'info> {
    // router CPI signer must be first
    #[account(
        constraint = valid_chain(state.enabled_chains, message.source_chain_selector),
        address = only_router(state.router),
    )]
    pub authority: Signer<'info>,
    pub state: Account<'info, BaseState>,
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
        address = state.owner, // only owner
    )]
    pub authority: Signer<'info>,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct BaseState {
    pub owner: Pubkey,
    pub router: Pubkey,
    pub enabled_chains: EnabledChains,
}

#[derive(InitSpace, Debug, Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct EnabledChains {
    pub xs: [u64; 0],
    pub len: u8,
}
arrayvec!(EnabledChains, u64, u8);

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SVMMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SVMTokenAmount>,
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // solana local token amount
}

#[event]
pub struct MessageReceived {
    pub message_id: [u8; 32],
}

// only_router returns address of router CPI caller
pub fn only_router(router_program: Pubkey) -> Pubkey {
    Pubkey::find_program_address(&[EXTERNAL_EXECUTION_CONFIG_SEED], &router_program).0
}

// valid_chain checks provided state to ensure it is enabled
pub fn valid_chain(enabled_chains: EnabledChains, chain_selector: u64) -> bool {
    let index = enabled_chains.binary_search(&chain_selector);
    match index {
        Ok(_) => true,
        Err(_) => false,
    }
}
