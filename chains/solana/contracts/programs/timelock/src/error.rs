use anchor_lang::error_code;

// this "AuthError" is separated from the "TimelockError" for error type generation from "anchor-go" tool
// Known issue: only the first error_code block is included in idl.errors field, and go bindings for this first errors not generated.
// anchor-go generates types for error from the second error_code block onwards.
// This might be a bug in anchor-go, should be revisited once program functionality is stable.
// Workaround: keep errors that not likely to change during development in the first error_code block(keeping hardcoded error types for this),
// and other errors in the second block.
#[error_code]
pub enum AuthError {
    #[msg("The signer is unauthorized")]
    Unauthorized = 0,
}

#[error_code]
pub enum TimelockError {
    #[msg("Invalid inputs")]
    InvalidInput,

    #[msg("Overflow")]
    Overflow,

    #[msg("Provided ID is invalid")]
    InvalidId,

    #[msg("operation not finalized")]
    OperationNotFinalized,

    #[msg("operation is already finalized")]
    OperationAlreadyFinalized,

    #[msg("too many instructions in the operation")]
    TooManyInstructions,

    // on attempt to create PDA with the same seed(existing operation)
    #[msg("operation already scheduled")]
    OperationAlreadyScheduled,

    #[msg("insufficient delay")]
    DelayInsufficient,

    #[msg("operation cannot be cancelled")]
    OperationNotCancellable,

    #[msg("operation is not ready")]
    OperationNotReady,

    #[msg("operation is already executed")]
    OperationAlreadyExecuted,

    #[msg("Predecessor operation is not found")]
    MissingDependency,

    #[msg("Provided access controller is invalid")]
    InvalidAccessController,

    #[msg("selector is blocked")]
    BlockedSelector,

    #[msg("selector is already blocked")]
    AlreadyBlocked,

    #[msg("selector not found")]
    SelectorNotFound,

    #[msg("maximum capacity reached for function blocker")]
    MaxCapacityReached,
}
