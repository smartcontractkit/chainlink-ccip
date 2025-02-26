use anchor_lang::prelude::*;

declare_id!("BZH9p8zsEgCmYaThPCZ25qZcUdZeGHNxv56STNg7bURr");

#[program]
pub mod ccip_router_dummy_emitter {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize {}
