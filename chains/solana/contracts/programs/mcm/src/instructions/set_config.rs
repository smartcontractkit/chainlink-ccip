use anchor_lang::prelude::*;

use crate::config::*;
use crate::constant::*;
use crate::error::*;
use crate::eth_utils::*;
use crate::event::*;
use crate::state::root::*;

/// Set the configuration for the multisig instance after validating the input
pub fn set_config(
    ctx: Context<SetConfig>,
    _multisig_name: [u8; MULTISIG_NAME_PADDED], // for pda derivation
    signer_groups: Vec<u8>,
    group_quorums: [u8; NUM_GROUPS],
    group_parents: [u8; NUM_GROUPS],
    clear_root: bool,
) -> Result<()> {
    // signer addresses are preloaded in the ConfigSigners account through InitSigners, AppendSigners and FinalizeSigners instructions
    let signer_addresses = &ctx.accounts.config_signers.signer_addresses;

    require!(
        !signer_addresses.is_empty() && signer_addresses.len() <= MAX_NUM_SIGNERS,
        McmError::OutOfBoundsNumOfSigners
    );

    require!(
        signer_addresses.len() == signer_groups.len(),
        McmError::MismatchedInputSignerVectorsLength
    );
    {
        // count the number of children for each group while validating group structure
        let mut group_children_counts = signer_groups.iter().try_fold(
            [0u8; NUM_GROUPS],
            |mut acc, &group| -> Result<[u8; NUM_GROUPS]> {
                // make sure the specified signer group is in bound
                require!(
                    (group as usize) < NUM_GROUPS,
                    McmError::MismatchedInputGroupArraysLength
                );
                acc[group as usize] = acc[group as usize]
                    .checked_add(1)
                    .ok_or(McmError::Overflow)?;

                Ok(acc)
            },
        )?;

        const ROOT_GROUP: usize = 0;
        // check if the group structure is a tree
        for i in (0..NUM_GROUPS).rev() {
            // validate group structure in backwards(root is 0)

            match i {
                // root should have itself as parent
                ROOT_GROUP => require!(
                    group_parents[ROOT_GROUP] == ROOT_GROUP as u8,
                    McmError::GroupTreeNotWellFormed
                ),
                // make sure the parent group is at a higher level(lower index) than the current group
                _ => require!(group_parents[i] < i as u8, McmError::GroupTreeNotWellFormed),
            }

            let disabled: bool = group_quorums[i] == 0;

            match disabled {
                true => {
                    // validate disabled group has no children
                    require!(
                        group_children_counts[i] == 0,
                        McmError::SignerInDisabledGroup
                    );
                }
                false => {
                    // ensure the group quorum can be met(i.e. have more signers than the quorum)
                    require!(
                        group_children_counts[i] >= group_quorums[i],
                        McmError::OutOfBoundsGroupQuorum
                    );

                    // increase the parent group's children count
                    let parent_index = group_parents[i] as usize;
                    group_children_counts[parent_index] = group_children_counts[parent_index]
                        .checked_add(1)
                        .ok_or(McmError::Overflow)?;
                }
            }
        }
    }

    let config = &mut ctx.accounts.multisig_config;
    config.group_quorums = group_quorums;
    config.group_parents = group_parents;

    {
        let mut signers: Vec<McmSigner> = Vec::with_capacity(signer_addresses.len());
        let mut prev_signer = [0u8; EVM_ADDRESS_BYTES];

        for (index, &evm_addr) in signer_addresses.iter().enumerate() {
            require!(
                evm_addr > prev_signer,
                McmError::SignersAddressesMustBeStrictlyIncreasing
            );

            // update prev signer
            prev_signer = evm_addr;

            signers.push(McmSigner {
                evm_address: evm_addr,
                index: u8::try_from(index).unwrap(), // This is safe due to previous check on signer_addresses length
                group: signer_groups[index],
            })
        }
        config.signers = signers;
    }

    // clear_root is equivalent to overriding with a completely empty root
    if clear_root {
        let expiring_root = &mut ctx.accounts.expiring_root_and_op_count;
        let root_metadata = &mut ctx.accounts.root_metadata;

        // preserve the current op_count
        let current_op_count = expiring_root.op_count;

        // clear the expiring root while preserving op_count
        expiring_root.root = [0u8; 32]; // clear root (equivalent to bytes32(0) in Solidity)
        expiring_root.valid_until = 0; // clear timestamp
        expiring_root.op_count = current_op_count;

        // set root metadata to a cleared state
        root_metadata.chain_id = config.chain_id;
        root_metadata.multisig = config.key();
        root_metadata.pre_op_count = current_op_count;
        root_metadata.post_op_count = current_op_count;
        root_metadata.override_previous_root = true;
    }

    // clear_root is equivalent to overriding with a completely empty root
    if clear_root {
        let expiring_root = &mut ctx.accounts.expiring_root_and_op_count;
        let root_metadata = &mut ctx.accounts.root_metadata;

        // preserve the current op_count
        let current_op_count = expiring_root.op_count;

        // clear the expiring root while preserving op_count
        expiring_root.root = [0u8; 32]; // clear root (equivalent to bytes32(0) in Solidity)
        expiring_root.valid_until = 0; // clear timestamp
        expiring_root.op_count = current_op_count;

        // set root metadata to a cleared state
        root_metadata.chain_id = ctx.accounts.multisig_config.chain_id;
        root_metadata.multisig = ctx.accounts.multisig_config.key();
        root_metadata.pre_op_count = current_op_count;
        root_metadata.post_op_count = current_op_count;
        root_metadata.override_previous_root = true;
    }

    emit!(ConfigSet {
        group_parents,
        group_quorums,
        is_root_cleared: clear_root,
        signers: config.signers.clone(),
    });

    Ok(())
}

pub fn init_signers(
    ctx: Context<InitSigners>,
    _multisig_name: [u8; MULTISIG_NAME_PADDED],
    total_signers: u8,
) -> Result<()> {
    require!(
        total_signers > 0 && total_signers <= MAX_NUM_SIGNERS as u8,
        McmError::OutOfBoundsNumOfSigners
    );
    let config_signers = &mut ctx.accounts.config_signers;
    config_signers.total_signers = total_signers;

    // Note: is_finalized stays false until finalization
    Ok(())
}

pub fn append_signers(
    ctx: Context<AppendSigners>,
    _multisig_name: [u8; MULTISIG_NAME_PADDED],
    signers_batch: Vec<[u8; 20]>,
) -> Result<()> {
    let config_signers = &mut ctx.accounts.config_signers;

    // check bounds
    require!(
        config_signers.signer_addresses.len() + signers_batch.len()
            <= config_signers.total_signers as usize,
        McmError::OutOfBoundsNumOfSigners
    );

    // check if the signers are strictly increasing from the last signer
    let mut prev_signer = config_signers
        .signer_addresses
        .last()
        .copied()
        .unwrap_or([0u8; EVM_ADDRESS_BYTES]);

    for sig in signers_batch {
        require!(
            sig > prev_signer,
            McmError::SignersAddressesMustBeStrictlyIncreasing
        );
        prev_signer = sig;
        config_signers.signer_addresses.push(sig);
    }
    Ok(())
}

pub fn clear_signers(
    _ctx: Context<ClearSigners>,
    _multisig_name: [u8; MULTISIG_NAME_PADDED],
) -> Result<()> {
    // NOTE: ctx.accounts.config_signers is closed to be able to re-initialized,
    // also allow finalized config_signers to be cleared
    Ok(())
}

pub fn finalize_signers(
    ctx: Context<FinalizeSigners>,
    _multisig_name: [u8; MULTISIG_NAME_PADDED],
) -> Result<()> {
    let config_signers = &mut ctx.accounts.config_signers;

    require!(
        !config_signers.signer_addresses.is_empty()
            && config_signers.signer_addresses.len() == config_signers.total_signers as usize,
        McmError::OutOfBoundsNumOfSigners
    );

    config_signers.is_finalized = true;
    Ok(())
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED, multisig_name.as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + MultisigConfig::space_with_signers(
            config_signers.signer_addresses.len()
        ),
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(
        mut,
        seeds = [CONFIG_SIGNERS_SEED, multisig_name.as_ref()],
        bump,
        constraint = config_signers.is_finalized @ McmError::SignersNotFinalized,
        close = authority // close after config set
    )]
    pub config_signers: Account<'info, ConfigSigners>, // preloaded signers account

    #[account(mut, seeds = [ROOT_METADATA_SEED, multisig_name.as_ref()], bump)]
    pub root_metadata: Account<'info, RootMetadata>,

    #[account(mut, seeds = [EXPIRING_ROOT_AND_OP_COUNT_SEED, multisig_name.as_ref()], bump)]
    pub expiring_root_and_op_count: Account<'info, ExpiringRootAndOpCount>,

    #[account(mut, address = multisig_config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED], total_signers: u8)]
pub struct InitSigners<'info> {
    #[account(seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(
        init,
        payer = authority,
        space = ConfigSigners::space(total_signers as usize),
        seeds = [CONFIG_SIGNERS_SEED, multisig_name.as_ref()],
        bump
    )]
    pub config_signers: Account<'info, ConfigSigners>,

    #[account(mut, address = multisig_config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct AppendSigners<'info> {
    #[account(seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(
        mut,
        seeds = [CONFIG_SIGNERS_SEED, multisig_name.as_ref()],
        bump,
        constraint = !config_signers.is_finalized @ McmError::SignersAlreadyFinalized
    )]
    pub config_signers: Account<'info, ConfigSigners>,

    #[account(mut, address = multisig_config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct ClearSigners<'info> {
    #[account(seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(
        mut,
        seeds = [CONFIG_SIGNERS_SEED, multisig_name.as_ref()],
        bump,
        close = authority // close so that it can be re-initialized
    )]
    pub config_signers: Account<'info, ConfigSigners>,

    #[account(mut, address = multisig_config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct FinalizeSigners<'info> {
    #[account(seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(
        mut,
        seeds = [CONFIG_SIGNERS_SEED, multisig_name.as_ref()],
        bump,
        constraint = !config_signers.is_finalized @ McmError::SignersAlreadyFinalized
    )]
    pub config_signers: Account<'info, ConfigSigners>,

    #[account(mut, address = multisig_config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}
