use anchor_lang::prelude::*;

declare_id!("RmnXLft1mSEwDgMKu2okYuHkiazxntFFcZFrrcXxYg7");

pub mod context;
use context::*;

pub mod state;
use state::*;

pub mod event;
use event::*;

mod instructions;

#[program]
pub mod rmn_remote {
    use instructions::router;

    use super::*;

    /// Initializes the Rmn Remote contract.
    ///
    /// The initialization is responsibility of Admin, nothing more than calling this method should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization.
    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        ctx.accounts.config.set_inner(Config {
            owner: ctx.accounts.authority.key(),
            version: 1,
            proposed_owner: Pubkey::default(),
            default_code_version: CodeVersion::V1,
        });

        ctx.accounts.curses.version = 1;

        emit!(ConfigSet {
            default_code_version: ctx.accounts.config.default_code_version,
        });
        Ok(())
    }

    /// Returns the program type (name) and version.
    /// Used by offchain code to easily determine which program & version is being interacted with.
    ///
    /// # Arguments
    /// * `ctx` - The context
    pub fn type_version(_ctx: Context<Empty>) -> Result<String> {
        let response = env!("CCIP_BUILD_TYPE_VERSION").to_string();
        msg!("{}", response);
        Ok(response)
    }

    /// Transfers the ownership of the fee quoter to a new proposed owner.
    ///
    /// Shared func signature with other programs.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the transfer.
    /// * `proposed_owner` - The public key of the new proposed owner.
    pub fn transfer_ownership(ctx: Context<UpdateConfig>, new_owner: Pubkey) -> Result<()> {
        router::admin(ctx.accounts.config.default_code_version).transfer_ownership(ctx, new_owner)
    }

    /// Accepts the ownership of the fee quoter by the proposed owner.
    ///
    /// Shared func signature with other programs.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for accepting ownership.
    /// The new owner must be a signer of the transaction.
    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        router::admin(ctx.accounts.config.default_code_version).accept_ownership(ctx)
    }

    /// Sets the default code version to be used. This is then used by the slim routing layer to determine
    /// which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.
    ///
    /// Shared func signature with other programs.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `code_version` - The new code version to be set as default.
    pub fn set_default_code_version(
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()> {
        router::admin(ctx.accounts.config.default_code_version)
            .set_default_code_version(ctx, code_version)
    }

    /// Curses an abstract subject. If the subject is CurseSubject::GLOBAL,
    /// the entire chain will be cursed.
    ///
    /// Only the CCIP Admin may perform this operation
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for adding a new curse.
    /// * `subject` - The subject to curse.
    pub fn curse(ctx: Context<Curse>, subject: CurseSubject) -> Result<()> {
        router::admin(ctx.accounts.config.default_code_version).curse(ctx, subject)
    }

    /// Uncurses an abstract subject. If the subject is CurseSubject::GLOBAL,
    /// the entire chain curse will be lifted. (note that any other specific
    /// subject curses will remain active.)
    ///
    /// Only the CCIP Admin may perform this operation
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for removing a curse.
    /// * `subject` - The subject to uncurse.
    pub fn uncurse(ctx: Context<Uncurse>, subject: CurseSubject) -> Result<()> {
        router::admin(ctx.accounts.config.default_code_version).uncurse(ctx, subject)
    }

    /// Verifies that the subject is not cursed AND that this chain is not globally cursed.
    /// In case either of those assumptions fail, the instruction reverts.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required to inspect curses.
    /// * `subject` - The subject to verify. Note that this instruction will revert if the chain
    ///   is globally cursed too, even if the provided subject is not explicitly cursed.
    pub fn verify_not_cursed(ctx: Context<InspectCurses>, subject: CurseSubject) -> Result<()> {
        router::public(ctx.accounts.config.default_code_version).verify_not_cursed(ctx, subject)
    }
}

#[error_code]
pub enum RmnRemoteError {
    #[msg("The signer is unauthorized")]
    // offset error code so that they don't clash with other programs
    // (Anchor's base custom error code 6000 + offset 3000 = start at 9000)
    Unauthorized = 3000,
    #[msg("Subject is already cursed")]
    SubjectIsAlreadyCursed,
    #[msg("Subject was not cursed")]
    SubjectWasNotCursed,
    #[msg("Proposed owner is the current owner")]
    RedundantOwnerProposal,
    #[msg("Invalid version of the onchain state")]
    InvalidVersion,
    #[msg("The subject is actively cursed")]
    SubjectCursed,
    #[msg("This chain is globally cursed")]
    GloballyCursed,
    #[msg("Invalid code version")]
    InvalidCodeVersion,
}
