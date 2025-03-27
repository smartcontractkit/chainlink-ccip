use anchor_lang::prelude::*;

declare_id!("5vNJx78mz7KVMjhuipyr9jKBKcMrKYGdjGkgE4LUmjKk");

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

/// A multi-signature contract system that supports signing many transactions targeting
/// multiple chains with a single set of signatures. This program manages multiple multisig
/// configurations through Program Derived Accounts (PDAs), each identified by a unique
/// multisig_id.
///
/// Key Features:
/// - Multiple Configurations: A single deployed program instance can manage
///   multiple independent multisig configurations
/// - Hierarchical Signature Groups: Supports complex approval structures with
///   nested groups and customizable quorum requirements
/// - Merkle Tree Operations: Batches signed operations in Merkle trees for
///   efficient verification and execution
///
/// Usage Flow:
/// 1. Initialize multisig configuration with a unique ID
/// 2. Set up signer hierarchy and group structure
/// 3. Set a Merkle root with authenticated metadata and signatures
/// 4. Execute operations by providing Merkle proofs
#[program]
pub mod mcm {
    #![warn(missing_docs)]
    use super::*;

    /// Initialize a new multisig configuration.
    ///
    /// Creates the foundation for a new multisig instance by initializing the core configuration
    /// PDAs and registering the multisig_id and chain_id. This is the first step in setting up
    /// a new multisig configuration.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the accounts required for initialization:
    ///   - `multisig_config`: PDA that will store the core configuration
    ///   - `root_metadata`: PDA that will store the current root's metadata
    ///   - `expiring_root_and_op_count`: PDA that tracks the current root and operation count
    ///   - `authority`: The deployer who becomes the initial owner
    ///   - `program_data`: Used to validate that the caller is the program's upgrade authority
    /// - `chain_id`: Network identifier for the chain this configuration is targeting
    /// - `multisig_id`: A unique, 32-byte identifier (left-padded) for this multisig instance
    ///
    /// # Access Control
    ///
    /// This instruction can only be called by the program's upgrade authority (typically the deployer).
    ///
    /// # Note
    ///
    /// After initialization, the owner can transfer ownership through the two-step
    /// transfer_ownership/accept_ownership process.
    pub fn initialize(
        ctx: Context<Initialize>,
        chain_id: u64,
        multisig_id: [u8; MULTISIG_ID_PADDED],
    ) -> Result<()> {
        let config = &mut ctx.accounts.multisig_config;
        config.chain_id = chain_id;
        config.multisig_id = multisig_id;
        config.owner = ctx.accounts.authority.key();
        Ok(())
    }

    /// Propose a new owner for the multisig instance config.
    ///
    /// Only the current owner (admin) can propose a new owner.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_multisig_id`: The multisig identifier.
    /// - `proposed_owner`: The public key of the proposed new owner.
    pub fn transfer_ownership(
        ctx: Context<TransferOwnership>,
        _multisig_id: [u8; MULTISIG_ID_PADDED],
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(
            proposed_owner != config.owner && proposed_owner != Pubkey::default(),
            McmError::InvalidInputs
        );
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    /// Accept ownership of the multisig config.
    ///
    /// The proposed new owner must call this function to assume ownership.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_multisig_id`: The multisig identifier.
    pub fn accept_ownership(
        ctx: Context<AcceptOwnership>,
        _multisig_id: [u8; MULTISIG_ID_PADDED],
    ) -> Result<()> {
        // NOTE: take() resets proposed_owner to default
        ctx.accounts.config.owner = std::mem::take(&mut ctx.accounts.config.proposed_owner);
        Ok(())
    }

    /// Set up the configuration for the multisig instance.
    ///
    /// Validates and establishes the signer hierarchy, group structure, and quorum requirements.
    /// If `clear_root` is true, it also invalidates the current Merkle root.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the multisig configuration account.
    /// - `multisig_id`: The unique identifier for this multisig instance.
    /// - `signer_groups`: Vector assigning each signer to a specific group (must match signers length).
    /// - `group_quorums`: Array defining the required signatures for each group. A group with quorum=0 is disabled.
    /// - `group_parents`: Array defining the hierarchical relationship between groups, forming a tree structure.
    /// - `clear_root`: If true, invalidates the current root to prevent further operations from being executed.
    ///
    /// # Example
    ///
    /// A group structure like this:
    ///
    /// ```text
    ///                     ┌──────┐
    ///                  ┌─►│2-of-3│◄───────┐
    ///                  │  └──────┘        │
    ///                  │        ▲         │
    ///                  │        │         │
    ///               ┌──┴───┐ ┌──┴───┐ ┌───┴────┐
    ///           ┌──►│1-of-2│ │2-of-2│ │signer A│
    ///           │   └──────┘ └──────┘ └────────┘
    ///           │       ▲      ▲  ▲
    ///           │       │      │  │
    ///   ┌───────┴┐ ┌────┴───┐ ┌┴───────┐
    ///   │signer B│ │signer C│ │signer D│
    ///   └────────┘ └────────┘ └────────┘
    /// ```
    ///
    /// Would be configured with:
    /// - group_quorums = [2, 1, 2, ...] (root: 2-of-3, group1: 1-of-2, group2: 2-of-2)
    /// - group_parents = [0, 0, 0, ...] (all groups under root)
    pub fn set_config(
        ctx: Context<SetConfig>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        signer_groups: Vec<u8>,
        group_quorums: [u8; NUM_GROUPS],
        group_parents: [u8; NUM_GROUPS],
        clear_root: bool,
    ) -> Result<()> {
        instructions::set_config(
            ctx,
            multisig_id,
            signer_groups,
            group_quorums,
            group_parents,
            clear_root,
        )
    }

    /// Set a new Merkle root that defines approved operations.
    ///
    /// This function updates the active Merkle root after verifying ECDSA signatures and validating
    /// the provided metadata against a Merkle proof.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `root`: The new Merkle root to set.
    /// - `valid_until`: timestamp until which the root remains valid.
    /// - `metadata`: Structured input containing chain_id, multisig, and operation counters.
    /// - `metadata_proof`: Merkle proof validating the metadata.
    pub fn set_root(
        ctx: Context<SetRoot>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        root: [u8; 32],
        valid_until: u32,
        metadata: RootMetadataInput,
        metadata_proof: Vec<[u8; 32]>,
    ) -> Result<()> {
        instructions::set_root(
            ctx,
            multisig_id,
            root,
            valid_until,
            metadata,
            metadata_proof,
        )
    }

    /// Executes an operation after verifying it's authorized in the current Merkle root.
    ///
    /// This function:
    /// 1. Performs extensive validation checks on the operation
    ///    - Ensures the operation is within the allowed count range
    ///    - Verifies chain ID matches the configured chain
    ///    - Checks the root has not expired
    ///    - Validates the operation's nonce against current state
    /// 2. Verifies the operation's inclusion in the Merkle tree
    /// 3. Executes the cross-program invocation with the multisig signer PDA
    ///
    /// # Parameters
    ///
    /// - `ctx`: Context containing operation accounts and signer information
    /// - `multisig_id`: Identifier for the multisig instance
    /// - `chain_id`: Network identifier that must match configuration
    /// - `nonce`: Operation counter that must match current state
    /// - `data`: Instruction data to be executed
    /// - `proof`: Merkle proof for operation verification
    ///
    /// # Security Considerations
    ///
    /// This instruction implements secure privilege delegation through PDA signing.
    /// The multisig's signer PDA becomes the authoritative signer for the operation,
    /// allowing controlled execution of privileged actions while maintaining the
    /// security guarantees of the Merkle root validation.
    pub fn execute<'info>(
        ctx: Context<'_, '_, '_, 'info, Execute<'info>>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        chain_id: u64,
        nonce: u64,
        data: Vec<u8>, // bytes array
        proof: Vec<[u8; 32]>,
    ) -> Result<()> {
        instructions::execute(ctx, multisig_id, chain_id, nonce, data, proof)
    }

    /// Initialize the storage for signer addresses.
    ///
    /// Creates a temporary account to hold signer addresses during the multisig configuration process.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `total_signers`: The total number of signers to be added.
    pub fn init_signers(
        ctx: Context<InitSigners>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        total_signers: u8,
    ) -> Result<()> {
        instructions::init_signers(ctx, multisig_id, total_signers)
    }

    /// Append a batch of signer addresses to the temporary storage.
    ///
    /// Allows adding multiple signer addresses in batches to overcome transaction size limits.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `signers_batch`: A batch of Ethereum addresses (20 bytes each) to be added as signers.
    pub fn append_signers(
        ctx: Context<AppendSigners>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        signers_batch: Vec<[u8; 20]>,
    ) -> Result<()> {
        instructions::append_signers(ctx, multisig_id, signers_batch)
    }

    /// Clear the temporary signer storage.
    ///
    /// Closes the account storing signer addresses, allowing it to be reinitialized if needed.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    pub fn clear_signers(
        ctx: Context<ClearSigners>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
    ) -> Result<()> {
        instructions::clear_signers(ctx, multisig_id)
    }

    /// Finalize the signer configuration.
    ///
    /// Marks the signer list as complete and ready for incorporation into the multisig configuration.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    pub fn finalize_signers(
        ctx: Context<FinalizeSigners>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
    ) -> Result<()> {
        instructions::finalize_signers(ctx, multisig_id)
    }

    /// Initialize storage for ECDSA signatures.
    ///
    /// Creates a temporary account to hold signatures that will validate a new Merkle root.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `root`: The new Merkle root these signatures will approve.
    /// - `valid_until`: Timestamp until which the root will remain valid.
    /// - `total_signatures`: The total number of signatures to be added.
    pub fn init_signatures(
        ctx: Context<InitSignatures>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        root: [u8; 32],
        valid_until: u32,
        total_signatures: u8,
    ) -> Result<()> {
        instructions::init_signatures(ctx, multisig_id, root, valid_until, total_signatures)
    }

    /// Append a batch of ECDSA signatures to the temporary storage.
    ///
    /// Allows adding multiple signatures in batches to overcome transaction size limits.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `root`: The Merkle root being approved.
    /// - `valid_until`: Timestamp until which the root will remain valid.
    /// - `signatures_batch`: A batch of ECDSA signatures to be verified.
    pub fn append_signatures(
        ctx: Context<AppendSignatures>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        root: [u8; 32],
        valid_until: u32,
        signatures_batch: Vec<Signature>,
    ) -> Result<()> {
        instructions::append_signatures(ctx, multisig_id, root, valid_until, signatures_batch)
    }

    /// Clear the temporary signature storage.
    ///
    /// Closes the account storing signatures, allowing it to be reinitialized if needed.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `root`: The Merkle root associated with the signatures.
    /// - `valid_until`: Timestamp until which the root would remain valid.
    pub fn clear_signatures(
        ctx: Context<ClearSignatures>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        root: [u8; 32],
        valid_until: u32,
    ) -> Result<()> {
        instructions::clear_signatures(ctx, multisig_id, root, valid_until)
    }

    /// Finalize the signature configuration.
    ///
    /// Marks the signature list as finalized and ready for verification when setting a new root.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing required accounts.
    /// - `multisig_id`: The multisig instance identifier.
    /// - `root`: The Merkle root associated with the signatures.
    /// - `valid_until`: Timestamp until which the root will remain valid.
    pub fn finalize_signatures(
        ctx: Context<FinalizeSignatures>,
        multisig_id: [u8; MULTISIG_ID_PADDED],
        root: [u8; 32],
        valid_until: u32,
    ) -> Result<()> {
        instructions::finalize_signatures(ctx, multisig_id, root, valid_until)
    }
}

#[derive(Accounts)]
#[instruction(chain_id: u64, multisig_id: [u8; MULTISIG_ID_PADDED])]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [CONFIG_SEED, multisig_id.as_ref()],
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
        seeds = [ROOT_METADATA_SEED, multisig_id.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + RootMetadata::INIT_SPACE
    )]
    pub root_metadata: Account<'info, RootMetadata>,

    #[account(
        init,
        seeds = [EXPIRING_ROOT_AND_OP_COUNT_SEED, multisig_id.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExpiringRootAndOpCount::INIT_SPACE,
    )]
    pub expiring_root_and_op_count: Account<'info, ExpiringRootAndOpCount>,
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED])]
pub struct TransferOwnership<'info> {
    #[account(mut, seeds = [CONFIG_SEED, multisig_id.as_ref()], bump)]
    pub config: Account<'info, config::MultisigConfig>,

    #[account(address = config.owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED])]
pub struct AcceptOwnership<'info> {
    #[account(mut, seeds = [CONFIG_SEED, multisig_id.as_ref()], bump)]
    pub config: Account<'info, config::MultisigConfig>,

    #[account(address = config.proposed_owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}
