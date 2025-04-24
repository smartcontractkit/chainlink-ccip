/**
 * This program an example of a Invalid CCIP Receiver Program.
 * Used to test CCIP Router execute and check that it fails
 */
use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenAccount};
use example_ccip_receiver::Any2SVMMessage;
use program::TestCcipInvalidReceiver;
use solana_program::pubkey;

declare_id!("FmyF3oW69MSAhyPSiZ69C4RKBdCPv5vAFTScisV7Me2j");

#[program]
pub mod test_ccip_invalid_receiver {
    use anchor_lang::solana_program::instruction::Instruction;
    use anchor_lang::solana_program::program::{get_return_data, invoke_signed};

    use super::*;

    // Actual invalid receiver method
    pub fn ccip_receive(ctx: Context<Initialize>, _message: Any2SVMMessage) -> Result<()> {
        msg!("Not reachable due to uninitialized accounts");

        let counter = &mut ctx.accounts.counter;
        counter.value = 1;

        Ok(())
    }

    ///////////////////////////////////////
    // Mocks of router, onramp & offramp //
    ///////////////////////////////////////
    pub fn add_offramp(
        _ctx: Context<AddOfframp>,
        _source_chain_selector: u64,
        _offramp: Pubkey,
    ) -> Result<()> {
        msg!("Registering offramp as allowed for source chain. This is a mock of the router functionality.");
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
            ctx.accounts.state.to_account_info(),
            ctx.accounts.token_program.to_account_info(),
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.accounts.pool_token_account.to_account_info(),
            ctx.accounts.chain_config.to_account_info(),
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
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

        let seeds: &[&[u8]] = &[
            b"external_token_pools_signer",
            ctx.accounts.test_pool.key.as_ref(),
            &[ctx.bumps.cpi_signer],
        ];

        invoke_signed(&ix, &acc_infos, &[seeds])?;

        let (_, data) = get_return_data().unwrap();

        Ok(data)
    }

    // This is just a dumb proxy towards the test_pool program, but signing the call with a PDA that mimics
    // what the offramp does
    pub fn pool_proxy_lock_or_burn<'info>(
        ctx: Context<'_, '_, 'info, 'info, PoolProxyLockOrBurn<'info>>,
        lock_or_burn: LockOrBurnInV1,
    ) -> Result<Vec<u8>> {
        let mut acc_infos = vec![
            ctx.accounts.cpi_signer.to_account_info(),
            ctx.accounts.state.to_account_info(),
            ctx.accounts.token_program.to_account_info(),
            ctx.accounts.mint.to_account_info(),
            ctx.accounts.pool_signer.to_account_info(),
            ctx.accounts.pool_token_account.to_account_info(),
            ctx.accounts.rmn_remote.to_account_info(),
            ctx.accounts.rmn_remote_curses.to_account_info(),
            ctx.accounts.rmn_remote_config.to_account_info(),
            ctx.accounts.chain_config.to_account_info(),
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
            data: lock_or_burn.to_tx_data(),
        };

        let seeds: &[&[u8]] = &[
            b"external_token_pools_signer",
            ctx.accounts.test_pool.key.as_ref(),
            &[ctx.bumps.cpi_signer],
        ];

        invoke_signed(&ix, &acc_infos, &[seeds])?;

        let (_, data) = get_return_data().unwrap();

        Ok(data)
    }

    pub fn receiver_proxy_execute<'info>(
        ctx: Context<'_, '_, 'info, 'info, ReceiverProxyExecute<'info>>,
        message: Any2SVMMessage,
    ) -> Result<()> {
        let mut acc_infos = vec![
            ctx.accounts.cpi_signer.to_account_info(),
            ctx.accounts.offramp_program.to_account_info(),
            ctx.accounts.allowed_offramp.to_account_info(),
        ];
        acc_infos.extend_from_slice(ctx.remaining_accounts); // these depend on the specific receiver

        let acc_metas: Vec<AccountMeta> = acc_infos
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == ctx.accounts.cpi_signer.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let ix = Instruction {
            program_id: ctx.accounts.test_receiver.key(),
            accounts: acc_metas,
            data: build_receiver_discriminator_and_data(&message)?,
        };

        let seeds: &[&[u8]] = &[
            b"external_execution_config",
            ctx.accounts.test_receiver.key.as_ref(),
            &[ctx.bumps.cpi_signer],
        ];

        invoke_signed(&ix, &acc_infos, &[seeds])?;

        Ok(())
    }
}

const ANCHOR_DISCRIMINATOR: usize = 8;

const TEST_ROUTER: Pubkey = pubkey!("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C");

#[derive(Accounts, Debug)]
#[instruction(message: Any2SVMMessage)]
pub struct Initialize<'info> {
    // router CPI signer must be first
    #[account(mut)]
    pub authority: Signer<'info>,

    /// CHECK offramp program: exists only to derive the allowed offramp PDA
    /// and the authority PDA. Must be second.
    pub offramp_program: UncheckedAccount<'info>,

    // PDA to verify that calling offramp is valid. Must be third. It is left up to the implementer to decide
    // how they want to persist the router address to verify that this is the correct account (e.g. in the top level of
    // a global config/state account for the receiver, which is what this example does, or hard-coded,
    // or stored in any other way in any other account).
    /// CHECK PDA of the router program verifying the signer is an allowed offramp.
    /// If PDA does not exist, the router doesn't allow this offramp
    #[account(
        owner = TEST_ROUTER, // this guarantees that it was initialized
        seeds = [
            ALLOWED_OFFRAMP,
            message.source_chain_selector.to_le_bytes().as_ref(),
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = TEST_ROUTER,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

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

pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";

#[account]
#[derive(Copy, Debug, InitSpace)]
pub struct AllowedOfframp {}

#[derive(Accounts)]
#[instruction(source_chain_selector: u64, offramp: Pubkey)]
pub struct AddOfframp<'info> {
    #[account(
        init,
        seeds = [ALLOWED_OFFRAMP, source_chain_selector.to_le_bytes().as_ref(), offramp.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + AllowedOfframp::INIT_SPACE,
    )]
    pub allowed_offramp: Account<'info, AllowedOfframp>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Counter {
    pub value: u8,
}

#[derive(Accounts)]
pub struct ReceiverProxyExecute<'info> {
    /// CHECK
    pub test_receiver: UncheckedAccount<'info>,

    /// CHECK
    #[account(
        seeds = [b"external_execution_config", test_receiver.key().as_ref()],
        bump,
    )]
    pub cpi_signer: UncheckedAccount<'info>,

    /// CHECK
    pub offramp_program: UncheckedAccount<'info>,

    /// CHECK
    pub allowed_offramp: UncheckedAccount<'info>,
    //
    /*
    Remaining Accounts:
        Example-receiver specific PDAs
            pub approved_sender: UncheckedAccount<'info>,
            pub state: UncheckedAccount<'info>,

        PingPong specific PDAs
            -- see ping pong contract for more details
    */
}

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct PoolProxyReleaseOrMint<'info> {
    /// CHECK
    pub test_pool: UncheckedAccount<'info>,

    /// CHECK
    #[account(
        seeds = [b"external_token_pools_signer", test_pool.key().as_ref()],
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
    pub state: UncheckedAccount<'info>,

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
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK
    pub rmn_remote_config: UncheckedAccount<'info>,

    /// CHECK
    #[account(mut)]
    pub receiver_token_account: UncheckedAccount<'info>,
}

#[derive(Accounts)]
#[instruction(lock_or_burn: LockOrBurnInV1)]
pub struct PoolProxyLockOrBurn<'info> {
    /// CHECK
    pub test_pool: UncheckedAccount<'info>,

    /// CHECK
    #[account(
        seeds = [b"external_token_pools_signer", test_pool.key().as_ref()],
        bump,
    )]
    pub cpi_signer: UncheckedAccount<'info>,

    ///////////////////////////////////
    // Accounts required by Pool CPI //
    ///////////////////////////////////
    /// CHECK
    #[account(mut)]
    pub state: UncheckedAccount<'info>,

    /// CHECK
    pub token_program: AccountInfo<'info>,

    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,

    /// CHECK
    pub pool_signer: UncheckedAccount<'info>,

    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    /// CHECK
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK
    pub rmn_remote_config: UncheckedAccount<'info>,

    /// CHECK
    #[account(mut)]
    pub chain_config: UncheckedAccount<'info>,
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

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnInV1 {
    pub receiver: Vec<u8>, //  The recipient of the tokens on the destination chain
    pub remote_chain_selector: u64, // The chain ID of the destination chain
    pub original_sender: Pubkey, // The original sender of the tx on the source chain
    pub amount: u64, // local solana amount to lock/burn,  The amount of tokens to lock or burn, denominated in the source token's decimals
    pub local_token: Pubkey, //  The address on this chain of the token to lock or burn
}
const TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR: [u8; 8] =
    [0x72, 0xa1, 0x5e, 0x1d, 0x93, 0x19, 0xe8, 0xbf]; // lock_or_burn_tokens

impl LockOrBurnInV1 {
    fn to_tx_data(&self) -> Vec<u8> {
        let mut data = Vec::new();
        data.extend_from_slice(&TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR);
        data.extend_from_slice(&self.try_to_vec().unwrap());
        data
    }
}

pub const CCIP_RECEIVE_DISCRIMINATOR: [u8; 8] = [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]; // ccip_receive

pub fn build_receiver_discriminator_and_data(msg: &Any2SVMMessage) -> Result<Vec<u8>> {
    let message = msg.try_to_vec()?;

    let mut data = Vec::with_capacity(8 + message.len());
    data.extend_from_slice(&CCIP_RECEIVE_DISCRIMINATOR);
    data.extend_from_slice(&message);

    Ok(data)
}
