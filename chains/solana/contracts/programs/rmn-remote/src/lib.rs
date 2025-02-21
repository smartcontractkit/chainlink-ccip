use anchor_lang::prelude::*;

declare_id!("CPkyVFQmyzmb6HfDE5TQr3NmZAsydcYspBoc3bf6Zo5x");

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
    /// * `local_chain_selector` This chain selector will be used to verify if this chain is globally
    ///    cursed, by transforming it into a curse subject (EVM equivalent: `bytes16(u128(selector))`)
    pub fn initialize(ctx: Context<Initialize>, local_chain_selector: u64) -> Result<()> {
        ctx.accounts.config.set_inner(Config {
            owner: ctx.accounts.authority.key(),
            version: 1,
            proposed_owner: Pubkey::default(),
            default_code_version: CodeVersion::V1,
            local_chain_selector,
        });

        emit!(ConfigSet {
            default_code_version: ctx.accounts.config.default_code_version,
            local_chain_selector
        });
        Ok(())
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

    /// Sets the local chain selector. This chain selector will be used to verify if this chain is globally
    /// cursed, by transforming it into a curse subject (EVM equivalent: `bytes16(u128(selector))`)
    ///
    /// This should only be called in case of a previous misconfiguration during initialization.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `local_chain_selector` - The local chain selector.
    pub fn set_local_chain_selector(
        ctx: Context<UpdateConfig>,
        local_chain_selector: u64,
    ) -> Result<()> {
        router::admin(ctx.accounts.config.default_code_version)
            .set_local_chain_selector(ctx, local_chain_selector)
    }

    /// Curses an abstract subject. If the subject is CurseSubject::from_chain_selector(local_chain_selector),
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

    /// Uncurses an abstract subject. If the subject is CurseSubject::from_chain_selector(local_chain_selector),
    /// the entire chain curse will be lifted.
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

    /// Retrieves a list of cursed subjects. Note this function will not revert if there's an active
    /// curse: It is to be used to retrieve information only.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required to inspect curses.
    pub fn get_cursed_subjects(ctx: Context<InspectCurses>) -> Result<Vec<CurseSubject>> {
        router::public(ctx.accounts.config.default_code_version).get_cursed_subjects(ctx)
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
