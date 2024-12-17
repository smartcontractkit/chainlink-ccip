use anchor_lang::error_code;

#[error_code]
pub enum PlaceholderError {
    #[msg("Todo generate error type with anchor")]
    Placeholder,
}

#[error_code]
pub enum TimelockError {
    #[msg("The signer is unauthorized")]
    Unauthorized = 0,

    #[msg("Invalid inputs")]
    InvalidInput,

    #[msg("Overflow")]
    Overflow,

    #[msg("Provided ID is invalid")]
    InvalidId,

    #[msg("RBACTimelock: operation not finalized")]
    OperationNotFinalized,

    #[msg("RBACTimelock: operation is already finalized")]
    OperationAlreadyFinalized,

    #[msg("RBACTimelock: too many instructions in the operation")]
    TooManyInstructions,

    // on attempt to create PDA with the same seed(existing operation)
    #[msg("RBACTimelock: operation already scheduled")]
    OperationAlreadyScheduled,

    #[msg("RBACTimelock: insufficient delay")]
    DelayInsufficient,

    // cancel
    #[msg("RBACTimelock: operation cannot be cancelled")]
    OperationNotCancellable,

    #[msg("operation is not ready")]
    OperationNotReady,

    #[msg("Predecessor operation is not found")]
    MissingDependency,

    #[msg("RBACTimelock: selector is blocked")]
    BlockedSelector,

    #[msg("RBACTimelock: Provided access controller is invalid")]
    InvalidAccessController,
}
