use anchor_lang::prelude::*;

declare_id!("6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX");

use program::Mcm;

mod constant;
pub use constant::*;

mod error;
pub use error::*;

mod event;

mod state;
pub use state::root::*;
pub use state::*;

mod eth_utils;
use eth_utils::*;

mod instructions;
use instructions::*;

/// This is mcm program supporting multiple instances of multisig configuration
/// A single deployed program manages multiple multisig states(configurations) identified by multisig_name
#[program]
pub mod mcm {
    use super::*;

    /// initialize a new multisig configuration, store the chain_id and multisig_name
    /// multisig_name is a unique identifier for the multisig configuration(32 bytes, left-padded)
    pub fn initialize(
        ctx: Context<Initialize>,
        chain_id: u64,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
    ) -> Result<()> {
        let config = &mut ctx.accounts.multisig_config;
        config.chain_id = chain_id;
        config.multisig_name = multisig_name;
        config.owner = ctx.accounts.authority.key();
        Ok(())
    }

    // shared func signature with other programs
    pub fn transfer_ownership(
        ctx: Context<TransferOwnership>,
        _multisig_name: [u8; MULTISIG_NAME_PADDED],
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(proposed_owner != config.owner, McmError::InvalidInputs);
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    // shared func signature with other programs
    pub fn accept_ownership(
        ctx: Context<AcceptOwnership>,
        _multisig_name: [u8; MULTISIG_NAME_PADDED],
    ) -> Result<()> {
        ctx.accounts.config.owner = std::mem::take(&mut ctx.accounts.config.proposed_owner);
        ctx.accounts.config.proposed_owner = Pubkey::new_from_array([0; 32]);
        Ok(())
    }

    pub fn set_config(
        ctx: Context<SetConfig>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        signer_groups: Vec<u8>,
        group_quorums: [u8; NUM_GROUPS],
        group_parents: [u8; NUM_GROUPS],
        clear_root: bool,
    ) -> Result<()> {
        instructions::set_config(
            ctx,
            multisig_name,
            signer_groups,
            group_quorums,
            group_parents,
            clear_root,
        )
    }

    pub fn set_root(
        ctx: Context<SetRoot>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        root: [u8; 32],
        valid_until: u32,
        metadata: RootMetadataInput,
        metadata_proof: Vec<[u8; 32]>,
    ) -> Result<()> {
        instructions::set_root(
            ctx,
            multisig_name,
            root,
            valid_until,
            metadata,
            metadata_proof,
        )
    }

    pub fn execute<'info>(
        ctx: Context<'_, '_, '_, 'info, Execute<'info>>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        chain_id: u64,
        nonce: u64,
        data: Vec<u8>, // bytes array
        proof: Vec<[u8; 32]>,
    ) -> Result<()> {
        instructions::execute(ctx, multisig_name, chain_id, nonce, data, proof)
    }

    // batch configuration methods prerequisites for set_config
    pub fn init_signers(
        ctx: Context<InitSigners>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        total_signers: u8,
    ) -> Result<()> {
        instructions::init_signers(ctx, multisig_name, total_signers)
    }

    pub fn append_signers(
        ctx: Context<AppendSigners>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        signers_batch: Vec<[u8; 20]>,
    ) -> Result<()> {
        instructions::append_signers(ctx, multisig_name, signers_batch)
    }

    pub fn clear_signers(
        ctx: Context<ClearSigners>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
    ) -> Result<()> {
        instructions::clear_signers(ctx, multisig_name)
    }

    pub fn finalize_signers(
        ctx: Context<FinalizeSigners>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
    ) -> Result<()> {
        instructions::finalize_signers(ctx, multisig_name)
    }

    // batch configuration methods prerequisites for set_root
    pub fn init_signatures(
        ctx: Context<InitSignatures>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        root: [u8; 32],
        valid_until: u32,
        total_signatures: u8,
    ) -> Result<()> {
        instructions::init_signatures(ctx, multisig_name, root, valid_until, total_signatures)
    }

    pub fn append_signatures(
        ctx: Context<AppendSignatures>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        root: [u8; 32],
        valid_until: u32,
        signatures_batch: Vec<Signature>,
    ) -> Result<()> {
        instructions::append_signatures(ctx, multisig_name, root, valid_until, signatures_batch)
    }
    pub fn clear_signatures(
        ctx: Context<ClearSignatures>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        root: [u8; 32],
        valid_until: u32,
    ) -> Result<()> {
        instructions::clear_signatures(ctx, multisig_name, root, valid_until)
    }

    pub fn finalize_signatures(
        ctx: Context<FinalizeSignatures>,
        multisig_name: [u8; MULTISIG_NAME_PADDED],
        root: [u8; 32],
        valid_until: u32,
    ) -> Result<()> {
        instructions::finalize_signatures(ctx, multisig_name, root, valid_until)
    }
}

#[derive(Accounts)]
#[instruction(chain_id: u64, multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [CONFIG_SEED, multisig_name.as_ref()],
        bump,
        space = ANCHOR_DISCRIMINATOR + config::MultisigConfig::INIT_SPACE,
        payer = authority,
    )]
    pub multisig_config: Account<'info, config::MultisigConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, Mcm>,
    // initialization only allowed by program upgrade authority(in common cases, the initial deployer)
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ AuthError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,

    #[account(
        init,
        seeds = [ROOT_METADATA_SEED, multisig_name.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + RootMetadata::INIT_SPACE
    )]
    pub root_metadata: Account<'info, RootMetadata>,

    #[account(
        init,
        seeds = [EXPIRING_ROOT_AND_OP_COUNT_SEED, multisig_name.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExpiringRootAndOpCount::INIT_SPACE,
    )]
    pub expiring_root_and_op_count: Account<'info, ExpiringRootAndOpCount>,
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct TransferOwnership<'info> {
    #[account(mut, seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub config: Account<'info, config::MultisigConfig>,

    #[account(address = config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct AcceptOwnership<'info> {
    #[account(mut, seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub config: Account<'info, config::MultisigConfig>,

    #[account(address = config.proposed_owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}
