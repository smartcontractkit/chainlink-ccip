use anchor_lang::prelude::*;

use crate::config::MultisigConfig;
use crate::constant::*;
use crate::error::*;
use crate::eth_utils::*;
use crate::event::*;
use crate::state::root::*;

pub fn set_root(
    ctx: Context<SetRoot>,
    _multisig_id: [u8; MULTISIG_ID_PADDED],
    root: [u8; 32],
    valid_until: u32,
    metadata: RootMetadataInput,
    metadata_proof: Vec<[u8; 32]>,
) -> Result<()> {
    require!(
        !ctx.accounts.seen_signed_hashes.seen,
        McmError::SignedHashAlreadySeen
    );

    // verify ECDSA signatures on (root, validUntil) and ensure that the root group is successful
    verify_ecdsa_signatures(
        &ctx.accounts.root_signatures.signatures,
        &ctx.accounts.multisig_config,
        &root,
        valid_until,
    )?;

    require!(
        Clock::get()?.unix_timestamp <= valid_until.into(),
        McmError::ValidUntilHasAlreadyPassed
    );

    // verify metadataProof
    {
        let calculated_root = calculate_merkle_root(metadata_proof.clone(), &metadata.hash_leaf());
        require!(root == calculated_root, McmError::ProofCannotBeVerified);
    }

    require_eq!(
        metadata.chain_id,
        ctx.accounts.multisig_config.chain_id,
        McmError::WrongChainId
    );

    require_keys_eq!(
        ctx.accounts.multisig_config.key(),
        metadata.multisig.key(),
        McmError::WrongMultiSig
    );

    let current_op_count = ctx.accounts.expiring_root_and_op_count.op_count;
    // don't allow a new root to be set if there are still outstanding ops that have not been
    // executed, unless override_previous_root is set
    require!(
        current_op_count == ctx.accounts.root_metadata.post_op_count
            || metadata.override_previous_root,
        McmError::PendingOps
    );

    // the signers are responsible for tracking opCount offchain and ensuring that
    // preOpCount equals to opCount
    require_eq!(
        current_op_count,
        metadata.pre_op_count,
        McmError::WrongPreOpCount
    );

    // note: allows preOpCount == postOpCount to support overriding previous root
    // while canceling remaining operations
    require!(
        metadata.pre_op_count <= metadata.post_op_count,
        McmError::WrongPostOpCount
    );

    // After validations, persist all changes
    ctx.accounts
        .expiring_root_and_op_count
        .set_inner(ExpiringRootAndOpCount {
            root,
            valid_until,
            op_count: metadata.pre_op_count,
        });

    ctx.accounts.root_metadata.set_inner(RootMetadata {
        chain_id: metadata.chain_id,
        multisig: metadata.multisig,
        pre_op_count: metadata.pre_op_count,
        post_op_count: metadata.post_op_count,
        override_previous_root: metadata.override_previous_root,
    });

    ctx.accounts.seen_signed_hashes.seen = true; // now it has been seen

    let md = &ctx.accounts.root_metadata;

    emit!(NewRoot {
        root,
        valid_until,
        metadata_chain_id: md.chain_id,
        metadata_multisig: md.multisig,
        metadata_pre_op_count: md.pre_op_count,
        metadata_post_op_count: md.post_op_count,
        metadata_override_previous_root: md.override_previous_root,
    });

    Ok(())
}

fn verify_ecdsa_signatures(
    signatures: &Vec<Signature>,
    multisig_config: &MultisigConfig,
    root: &[u8; 32],
    valid_until: u32,
) -> Result<()> {
    let signed_hash = compute_eth_message_hash(root, valid_until);
    let mut previous_addr: [u8; EVM_ADDRESS_BYTES] = [0; EVM_ADDRESS_BYTES];
    let mut group_vote_counts: [u8; NUM_GROUPS] = [0; NUM_GROUPS];

    for sig in signatures {
        let signer_addr = {
            let recovered = ecdsa_recover_evm_addr(&signed_hash.to_bytes(), sig)?;
            require!(
                recovered.gt(&previous_addr),
                McmError::SignersAddressesMustBeStrictlyIncreasing
            );
            recovered
        };
        previous_addr = signer_addr;

        let signer_idx = multisig_config
            .signers
            .binary_search_by_key(&signer_addr, |s| s.evm_address);

        require!(signer_idx.is_ok(), McmError::InvalidSigner);

        let signer = &multisig_config.signers[signer_idx.unwrap()];

        let mut group = usize::from(signer.group);
        loop {
            group_vote_counts[group] += 1;
            if group_vote_counts[group] != multisig_config.group_quorums[group] {
                // bail out unless we just hit the quorum. we only hit each quorum once,
                // so we never move on to the parent of a group more than once.
                break;
            }
            if group == 0 {
                // reached root
                break;
            }
            group = usize::from(multisig_config.group_parents[group]);
        }
    }

    // the group at the root of the tree (with index 0) determines whether the vote passed,
    // we cannot proceed if it isn't configured with a valid (non-zero) quorum
    require!(
        multisig_config.group_quorums[0] != 0,
        McmError::MissingConfig
    );

    // check if the root group reach its quorum
    require!(
        group_vote_counts[0] >= multisig_config.group_quorums[0],
        McmError::InsufficientSigners
    );

    Ok(())
}

pub fn init_signatures(
    ctx: Context<InitSignatures>,
    _multisig_id: [u8; MULTISIG_ID_PADDED],
    _root: [u8; 32],
    _valid_until: u32,
    total_signatures: u8,
) -> Result<()> {
    let signatures_account = &mut ctx.accounts.signatures;
    signatures_account.total_signatures = total_signatures;

    // Note: is_finalized stays false until finalization
    Ok(())
}

pub fn append_signatures(
    ctx: Context<AppendSignatures>,
    _multisig_id: [u8; MULTISIG_ID_PADDED],
    _root: [u8; 32],
    _valid_until: u32,
    signatures_batch: Vec<Signature>,
) -> Result<()> {
    let signatures_account = &mut ctx.accounts.signatures;

    require!(
        signatures_account.signatures.len() + signatures_batch.len()
            <= signatures_account.total_signatures as usize,
        McmError::TooManySignatures
    );

    for sig in signatures_batch {
        signatures_account.signatures.push(sig);
    }
    Ok(())
}

pub fn clear_signatures(
    _ctx: Context<ClearSignatures>,
    _multisig_id: [u8; MULTISIG_ID_PADDED],
    _root: [u8; 32],
    _valid_until: u32,
) -> Result<()> {
    // NOTE: ctx.accounts.signatures is closed to be able to re-initialized,
    // also allow finalized signatures_account to be cleared
    Ok(())
}

pub fn finalize_signatures(
    ctx: Context<FinalizeSignatures>,
    _multisig_id: [u8; MULTISIG_ID_PADDED],
    _root: [u8; 32],
    _valid_until: u32,
) -> Result<()> {
    let signatures_account = &mut ctx.accounts.signatures;

    require!(
        !signatures_account.signatures.is_empty()
            && signatures_account.signatures.len() == signatures_account.total_signatures as usize,
        McmError::SignatureCountMismatch
    );

    signatures_account.is_finalized = true;
    Ok(())
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED], root: [u8; 32], valid_until: u32)]
pub struct SetRoot<'info> {
    #[account(
        mut,
        seeds = [ROOT_SIGNATURES_SEED, multisig_id.as_ref(), root.as_ref(), valid_until.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        constraint = root_signatures.is_finalized @ McmError::SignaturesNotFinalized,
        close = authority
    )]
    pub root_signatures: Account<'info, RootSignatures>, // preloaded signatures account

    #[account(mut, seeds = [ROOT_METADATA_SEED, multisig_id.as_ref()], bump)]
    pub root_metadata: Account<'info, RootMetadata>,

    #[account(
        init,
        seeds = [SEEN_SIGNED_HASHES_SEED, multisig_id.as_ref(), root.as_ref(), valid_until.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + SeenSignedHash::INIT_SPACE,
    )]
    pub seen_signed_hashes: Account<'info, SeenSignedHash>,

    #[account(mut, seeds = [EXPIRING_ROOT_AND_OP_COUNT_SEED, multisig_id.as_ref()], bump)]
    pub expiring_root_and_op_count: Account<'info, ExpiringRootAndOpCount>,

    #[account(seeds = [CONFIG_SEED, multisig_id.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(
    multisig_id: [u8; MULTISIG_ID_PADDED],
    root: [u8; 32],
    valid_until: u32,
    total_signatures: u8,
)]
pub struct InitSignatures<'info> {
    #[account(
        init,
        payer = authority,
        space = RootSignatures::space(total_signatures as usize),
        seeds = [ROOT_SIGNATURES_SEED, multisig_id.as_ref(), root.as_ref(), valid_until.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump
    )]
    pub signatures: Account<'info, RootSignatures>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED], root: [u8; 32], valid_until: u32)]
pub struct AppendSignatures<'info> {
    #[account(
        mut,
        seeds = [ROOT_SIGNATURES_SEED, multisig_id.as_ref(), root.as_ref(), valid_until.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        constraint = !signatures.is_finalized @ McmError::SignaturesAlreadyFinalized,
    )]
    pub signatures: Account<'info, RootSignatures>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED], root: [u8; 32], valid_until: u32)]
pub struct ClearSignatures<'info> {
    #[account(
        mut,
        seeds = [ROOT_SIGNATURES_SEED, multisig_id.as_ref(), root.as_ref(), valid_until.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        close = authority, // close so that it can be re-initialized
    )]
    pub signatures: Account<'info, RootSignatures>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED], root: [u8; 32], valid_until: u32)]
pub struct FinalizeSignatures<'info> {
    #[account(
        mut,
        seeds = [ROOT_SIGNATURES_SEED, multisig_id.as_ref(), root.as_ref(), valid_until.to_le_bytes().as_ref(), authority.key().as_ref()],
        bump,
        constraint = !signatures.is_finalized @ McmError::SignaturesAlreadyFinalized,
    )]
    pub signatures: Account<'info, RootSignatures>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
