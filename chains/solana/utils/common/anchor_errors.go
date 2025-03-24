package common

import (
	ag_binary "github.com/gagliardetto/binary"
)

// represents standard Anchor program errors
// Note: This is not an exhaustive list of Anchor errors.
// For a complete list, please refer to https://anchor.so/errors

//nolint:all
type AnchorError ag_binary.BorshEnum

//nolint:all
const (
	ConstraintSeeds_AnchorError AnchorError = iota
	AccountNotInitialized_AnchorError
	InstructionDidNotDeserialize_AnchorError
	ConstraintHasOne_AnchorError
	AccountOwnedByWrongProgram_AnchorError
	ConstraintTokenOwner_AnchorError
	InvalidProgramId_AnchorError
)

func (value AnchorError) String() string {
	switch value {
	case ConstraintSeeds_AnchorError:
		return "ConstraintSeeds"
	case AccountNotInitialized_AnchorError:
		return "AccountNotInitialized"
	case InstructionDidNotDeserialize_AnchorError:
		return "InstructionDidNotDeserialize"
	case ConstraintHasOne_AnchorError:
		return "ConstraintHasOne"
	case AccountOwnedByWrongProgram_AnchorError:
		return "AccountOwnedByWrongProgram"
	case ConstraintTokenOwner_AnchorError:
		return "ConstraintTokenOwner"
	case InvalidProgramId_AnchorError:
		return "InvalidProgramId"
	default:
		return ""
	}
}
