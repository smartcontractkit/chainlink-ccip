/**
 * This program is meant to only be used in integration tests on localnet.
 * Used to test CPIs made by other programs (with actual business logic).
 */
use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak;

declare_id!("2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm");

#[program]
pub mod external_program_cpi_stub {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        msg!("Called `initialize` {:?}", ctx);
        ctx.accounts.u8_value.value = 1;
        Ok(())
    }

    pub fn empty(ctx: Context<Empty>) -> Result<()> {
        msg!("Called `empty` {:?}", ctx);
        Ok(())
    }

    pub fn u8_instruction_data(ctx: Context<Empty>, data: u8) -> Result<()> {
        msg!("Called `u8_instruction_data` {:?} and data {data}", ctx);
        Ok(())
    }

    pub fn struct_instruction_data(ctx: Context<Empty>, data: Value) -> Result<()> {
        msg!(
            "Called `struct_instruction_data` {:?} and data {:?}",
            ctx,
            data
        );
        Ok(())
    }

    pub fn account_read(ctx: Context<AccountRead>) -> Result<()> {
        msg!("Called `account_read` {:?}", ctx);
        Ok(())
    }

    pub fn account_mut(ctx: Context<AccountMut>) -> Result<()> {
        msg!("Called `account_mut` {:?}", ctx);
        let u8_value = &mut ctx.accounts.u8_value;
        u8_value.value += 1;
        Ok(())
    }

    ///instruction that accepts arbitrarily large instruction data.
    pub fn big_instruction_data(_ctx: Context<Empty>, data: Vec<u8>) -> Result<()> {
        msg!(
            "Called `big_instruction_data` with data length: {}",
            data.len()
        );
        Ok(())
    }

    /// no-op instruction that does nothing, also can be used to test maximum account references(remaining_accounts)
    pub fn no_op(_ctx: Context<Empty>) -> Result<()> {
        Ok(())
    }

    pub fn compute_heavy(_ctx: Context<Empty>, iterations: u32) -> Result<()> {
        let mut hash = [0u8; 32];

        // Initialize with some data
        hash.iter_mut().enumerate().for_each(|(i, byte)| {
            *byte = i as u8;
        });

        // Perform multiple hash operations
        for _ in 0..iterations {
            hash = keccak::hash(&hash).to_bytes();
        }

        // Don't log with msg!() to avoid variable CU consumption
        Ok(())
    }
}

const VALUE_SEED: &[u8] = b"u8_value";
const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts, Debug)]
pub struct Empty {}

#[account]
#[derive(InitSpace, Debug)]
pub struct Value {
    pub value: u8,
}

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [VALUE_SEED],
        bump,
        payer = stub_caller,
        space = ANCHOR_DISCRIMINATOR + Value::INIT_SPACE,
    )]
    pub u8_value: Account<'info, Value>,

    #[account(mut)]
    pub stub_caller: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct AccountRead<'info> {
    #[account(
        seeds = [VALUE_SEED],
        bump,
    )]
    pub u8_value: Account<'info, Value>,
}

#[derive(Accounts, Debug)]
pub struct AccountMut<'info> {
    #[account(
        mut,
        seeds = [VALUE_SEED],
        bump,
    )]
    pub u8_value: Account<'info, Value>,

    pub stub_caller: Signer<'info>,

    pub system_program: Program<'info, System>,
}
