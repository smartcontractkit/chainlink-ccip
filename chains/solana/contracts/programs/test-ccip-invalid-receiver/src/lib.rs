/**
 * This program an example of a Invalid CCIP Receiver Program.
 * Used to test CCIP Router execute and check that it fails
 */
use anchor_lang::prelude::*;
use example_ccip_receiver::Any2SVMMessage;

declare_id!("9Vjda3WU2gsJgE4VdU6QuDw8rfHLyigfFyWs3XDPNUn8");

#[program]
pub mod test_ccip_invalid_receiver {
    use super::*;

    pub fn ccip_receive(ctx: Context<Initialize>, _message: Any2SVMMessage) -> Result<()> {
        msg!("Not reachable due to uninitialized accounts");

        let counter = &mut ctx.accounts.counter;
        counter.value = 1;

        Ok(())
    }
}

const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    // router CPI signer must be first
    #[account(mut)]
    pub authority: Signer<'info>,

    #[account(
        init,
        seeds = [b"counter"],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Counter::INIT_SPACE,
    )]
    pub counter: Account<'info, Counter>,

    pub system_program: Program<'info, System>,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Counter {
    pub value: u8,
}
