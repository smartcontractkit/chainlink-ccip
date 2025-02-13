use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenAccount};
use example_ccip_receiver::{
    Any2SVMMessage, BaseState, CcipReceiverError, ALLOWED_OFFRAMP, EXTERNAL_EXECUTION_CONFIG_SEED,
};

declare_id!("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH");

/// This program an example of a CCIP Receiver Program.
/// Used to test CCIP Router execute.
#[program]
pub mod test_ccip_receiver {
    use solana_program::instruction::Instruction;
    use solana_program::program::invoke_signed;

    use super::*;

    /// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
    pub fn initialize(ctx: Context<Initialize>, router: Pubkey) -> Result<()> {
        msg!("Called `initialize` {:?}", ctx);
        ctx.accounts.counter.value = 0;
        ctx.accounts
            .counter
            .state
            .init(ctx.accounts.authority.key(), router)
    }

    /// This function is called by the CCIP Router to execute the CCIP message.
    /// The method name needs to be ccip_receive with Anchor encoding,
    /// if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]
    /// You can send as many accounts as you need, specifying if mutable or not.
    /// But none of them could be an init, realloc or close.
    /// In this case, it increments the counter value by 1 and logs the parsed message.
    pub fn ccip_receive(ctx: Context<CcipReceive>, message: Any2SVMMessage) -> Result<()> {
        msg!("Called `ccip_receive` with message {:?}", message);

        let counter = &mut ctx.accounts.counter;
        counter.value += 1;

        // additional accounts trigger additional CPI call
        if !ctx.remaining_accounts.is_empty() {
            let external_execution_config = &ctx.accounts.external_execution_config;
            let (target_program, acc_infos) = ctx.remaining_accounts.split_at(1);

            let acc_metas: Vec<AccountMeta> = acc_infos
                .to_vec()
                .iter()
                .flat_map(|acc_info| {
                    // Check signer from PDA External Execution config
                    let is_signer = acc_info.key() == external_execution_config.key();
                    acc_info.to_account_metas(Some(is_signer))
                })
                .collect();

            let instruction = Instruction {
                program_id: target_program[0].key(),
                accounts: acc_metas,
                data: message.data,
            };

            let seeds = &[
                EXTERNAL_EXECUTION_CONFIG_SEED,
                &[ctx.bumps.external_execution_config],
            ];
            let signer = &[&seeds[..]];

            invoke_signed(&instruction, acc_infos, signer)?;
        };

        Ok(())
    }

    // these functions are called by CCIP token pools that wrap non-token programs
    // if not using Anchor, the discriminator must be:
    // const RELEASE_MINT: [u8; 8] = [0x14, 0x94, 0x71, 0xc6, 0xe5, 0xaa, 0x47, 0x30];
    // const LOCK_BURN: [u8; 8] = [0xc8, 0x0e, 0x32, 0x09, 0x2c, 0x5b, 0x79, 0x25];
    pub fn ccip_token_release_mint(
        _ctx: Context<TokenPool>,
        input: ReleaseOrMintInV1,
    ) -> Result<()> {
        msg!("Called `ccip_token_release_mint` with message {:?}", input);
        Ok(())
    }
    pub fn ccip_token_lock_burn(_ctx: Context<TokenPool>, input: LockOrBurnInV1) -> Result<()> {
        msg!("Called `ccip_token_lock_burn` with message {:?}", input);
        Ok(())
    }
}

const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [b"counter"],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Counter::INIT_SPACE,
    )]
    pub counter: Account<'info, Counter>,

    #[account(
        init,
        seeds = [EXTERNAL_EXECUTION_CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExternalExecutionConfig::INIT_SPACE,
    )]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(message: Any2SVMMessage)]
pub struct CcipReceive<'info> {
    // Offramp CPI signer PDA must be first
    // It is not mutable, and thus cannot be used as payer of init/realloc of other PDAs.
    #[account(
        seeds = [EXTERNAL_EXECUTION_CONFIG_SEED],
        bump,
        seeds::program = offramp_program.key(),
    )]
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
        owner = counter.state.router @ CcipReceiverError::InvalidCaller, // this guarantees that it was initialized
        seeds = [
            ALLOWED_OFFRAMP,
            message.source_chain_selector.to_le_bytes().as_ref(),
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = counter.state.router,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,
    // ccip router expects "receiver" to be second
    /// CHECK: Using this to sign
    #[account(mut, seeds = [EXTERNAL_EXECUTION_CONFIG_SEED], bump)]
    pub external_execution_config: Account<'info, ExternalExecutionConfig>,
    #[account(
        mut,
        seeds = [b"counter"],
        bump,
    )]
    pub counter: Account<'info, Counter>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct TokenPool<'info> {
    // token pool CPI signer
    pub authority: Signer<'info>,
    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    pub mint: InterfaceAccount<'info, Mint>,
    /// CHECK: CPI to underlying token program
    pub token_program: UncheckedAccount<'info>,
    // remaining accounts
    // [0] receiver_token_account // only included for release_mint
    // [...] extra passed accounts
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Counter {
    pub value: u8,
    pub state: BaseState,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct ExternalExecutionConfig {}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnInV1 {
    receiver: Vec<u8>, //  The recipient of the tokens on the destination chain, abi encoded
    remote_chain_selector: u64, // The chain ID of the destination chain
    original_sender: Pubkey, // The original sender of the tx on the source chain
    amount: u64, // local solana amount to lock/burn,  The amount of tokens to lock or burn, denominated in the source token's decimals
    local_token: Pubkey, //  The address on this chain of the token to lock or burn
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
