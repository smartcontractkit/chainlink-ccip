/**
 * This program an example of a Invalid CCIP Receiver Program.
 * Used to test CCIP Router execute and check that it fails
 */
use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenAccount};
use example_ccip_receiver::Any2SVMMessage;
use program::TestCcipInvalidReceiver;

declare_id!("9Vjda3WU2gsJgE4VdU6QuDw8rfHLyigfFyWs3XDPNUn8");

#[program]
pub mod test_ccip_invalid_receiver {
    use anchor_lang::solana_program::instruction::Instruction;
    use anchor_lang::solana_program::program::{get_return_data, invoke_signed};

    use super::*;

    pub fn ccip_receive(ctx: Context<Initialize>, _message: Any2SVMMessage) -> Result<()> {
        msg!("Not reachable due to uninitialized accounts");

        let counter = &mut ctx.accounts.counter;
        counter.value = 1;

        Ok(())
    }

    // This is just a dumb proxy towards the test_pool program, but signing the call with a PDA that mimics
    // what the offramp does
    pub fn pool_proxy_release_or_mint<'info>(
        ctx: Context<'_, '_, 'info, 'info, PoolProxyReleaseOrMint<'info>>,
        release_or_mint: ReleaseOrMintInV1,
    ) -> Result<Vec<u8>> {
        let mut acc_infos = vec![
            ctx.accounts.cpi_signer.to_account_info(),
            ctx.accounts.offramp_program.to_account_info(),
            ctx.accounts.allowed_offramp.to_account_info(),
            ctx.accounts.config.to_account_info(),
            ctx.accounts.token_program.to_account_info(),
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.accounts.pool_token_account.to_account_info(),
            ctx.accounts.chain_config.to_account_info(),
            ctx.accounts.receiver_token_account.to_account_info(),
        ];

        acc_infos.extend_from_slice(ctx.remaining_accounts);

        let acc_metas: Vec<AccountMeta> = acc_infos
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == ctx.accounts.cpi_signer.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let ix = Instruction {
            program_id: ctx.accounts.test_pool.key(),
            accounts: acc_metas,
            data: release_or_mint.to_tx_data(),
        };

        let seeds: &[&[u8]] = &[b"external_token_pools_signer", &[ctx.bumps.cpi_signer]];

        invoke_signed(&ix, &acc_infos, &[seeds])?;

        let (_, data) = get_return_data().unwrap();

        Ok(data)
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

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct PoolProxyReleaseOrMint<'info> {
    /// CHECK
    pub test_pool: UncheckedAccount<'info>,

    /// CHECK
    #[account(
        seeds = [b"external_token_pools_signer"],
        bump,
    )]
    pub cpi_signer: UncheckedAccount<'info>,

    ///////////////////////////////////
    // Accounts required by Pool CPI //
    ///////////////////////////////////
    pub offramp_program: Program<'info, TestCcipInvalidReceiver>, // this receiver acts as "dumb" offramp here

    /// CHECK
    pub allowed_offramp: UncheckedAccount<'info>,

    /// CHECK
    #[account(mut)]
    pub config: UncheckedAccount<'info>,

    /// CHECK
    pub token_program: AccountInfo<'info>,

    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK
    pub pool_signer: UncheckedAccount<'info>,

    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    /// CHECK
    #[account(mut)]
    pub chain_config: UncheckedAccount<'info>,

    /// CHECK
    #[account(mut)]
    pub receiver_token_account: UncheckedAccount<'info>,
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintInV1 {
    original_sender: Vec<u8>, //          The original sender of the tx on the source chain
    remote_chain_selector: u64, // ─╮ The chain ID of the source chain
    receiver: Pubkey,         // ───────────╯ The recipient of the tokens on the destination chain.
    amount: [u8; 32], // u256, incoming cross-chain amount - The amount of tokens to release or mint, denominated in the source token's decimals
    local_token: Pubkey, //            The address on this chain of the token to release or mint
    /// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
    /// expected pool address for the given remoteChainSelector.
    source_pool_address: Vec<u8>, //       The address of the source pool, abi encoded in the case of EVM chains
    source_pool_data: Vec<u8>, //          The data received from the source pool to process the release or mint
    /// @dev WARNING: offchainTokenData is untrusted data.
    offchain_token_data: Vec<u8>, //       The offchain data to process the release or mint
}

pub const TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR: [u8; 8] =
    [0x5c, 0x64, 0x96, 0xc6, 0xfc, 0x3f, 0xa4, 0xe4]; // release_or_mint_tokens

impl ReleaseOrMintInV1 {
    pub fn to_tx_data(&self) -> Vec<u8> {
        let mut data = Vec::new();
        data.extend_from_slice(&TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR);
        data.extend_from_slice(&self.try_to_vec().unwrap());
        data
    }
}
