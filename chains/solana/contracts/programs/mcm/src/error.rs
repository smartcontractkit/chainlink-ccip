use anchor_lang::error_code;

// error range
// note: custom numeric error codes start from 6000 unless specified like #[error_code(offset = 1000)]
// https://github.com/coral-xyz/anchor/blob/c25bd7b7ebbcaf12f6b8cbd3e6f34ae4e2833cb2/lang/syn/src/codegen/error.rs#L72
// Anchor built-in errors: https://anchor.so/errors
//
// [0:100]   Global errors
// [100:N]   Function errors
// todo: anchor-go error generation support update?

// todo: align message with EVM

#[error_code]
pub enum McmError {
    #[msg("Invalid multisig")]
    WrongMultiSig = 0, // 6000

    #[msg("Invalid chainID")]
    WrongChainId,

    #[msg("The signer is unauthorized")]
    Unauthorized,

    #[msg("Invalid inputs")]
    InvalidInputs,

    #[msg("overflow occurred.")]
    Overflow,

    #[msg("Invalid signature")]
    InvalidSignature,

    #[msg("Failed ECDSA recover")]
    FailedEcdsaRecover,

    #[msg("Invalid root length")]
    InvalidRootLen,

    #[msg("Config signers not finalized")]
    SignersNotFinalized,

    #[msg("Config signers already finalized")]
    SignersAlreadyFinalized,

    #[msg("Signatures already finalized")]
    SignaturesAlreadyFinalized,

    #[msg("Uploaded signatures count mismatch")]
    SignatureCountMismatch,

    #[msg("Too many signatures")]
    TooManySignatures,

    #[msg("Signatures not finalized")]
    SignaturesNotFinalized,

    #[msg("Signatures root mismatch")]
    SignaturesRootMismatch,

    #[msg("Signatures valid until mismatch")]
    SignaturesValidUntilMismatch,

    #[msg("The input vectors for signer addresses and signer groups must have the same length")]
    MismatchedInputSignerVectorsLength = 200,

    #[msg("The number of signers is 0 or greater than MAX_NUM_SIGNERS")]
    OutOfBoundsNumOfSigners,

    #[msg("The input arrays for group parents and group quorums must be of length NUM_GROUPS")]
    MismatchedInputGroupArraysLength,

    #[msg("the group tree isn't well-formed.")]
    GroupTreeNotWellFormed,

    #[msg("a disabled group contains a signer.")]
    SignerInDisabledGroup,

    #[msg("the quorum of some group is larger than the number of signers in it.")]
    OutOfBoundsGroupQuorum,

    // Prevents signers from including more than one signature
    #[msg("the signers' addresses are not a strictly increasing monotone sequence.")]
    SignersAddressesMustBeStrictlyIncreasing,

    #[msg("The combination of signature and valid_until has already been seen")]
    SignedHashAlreadySeen,

    #[msg("Invalid signer")]
    InvalidSigner,

    #[msg("Missing configuration")]
    MissingConfig,

    #[msg("Insufficient signers")]
    InsufficientSigners,

    #[msg("Valid until has already passed")]
    ValidUntilHasAlreadyPassed,

    #[msg("Proof cannot be verified")]
    ProofCannotBeVerified,

    #[msg("Pending operations")]
    PendingOps,

    #[msg("Wrong pre-operation count")]
    WrongPreOpCount,

    #[msg("Wrong post-operation count")]
    WrongPostOpCount,

    #[msg("Post-operation count reached")]
    PostOpCountReached,

    #[msg("Root expired")]
    RootExpired,

    #[msg("Wrong nonce")]
    WrongNonce,
}
