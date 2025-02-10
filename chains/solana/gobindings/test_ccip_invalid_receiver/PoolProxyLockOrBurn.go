// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_invalid_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PoolProxyLockOrBurn is the `poolProxyLockOrBurn` instruction.
type PoolProxyLockOrBurn struct {
	LockOrBurn *LockOrBurnInV1

	// [0] = [] testPool
	// ··········· CHECK
	//
	// [1] = [] cpiSigner
	// ··········· CHECK
	//
	// [2] = [WRITE] state
	// ··········· CHECK
	//
	// [3] = [] tokenProgram
	// ··········· CHECK
	//
	// [4] = [WRITE] mint
	//
	// [5] = [] poolSigner
	// ··········· CHECK
	//
	// [6] = [WRITE] poolTokenAccount
	//
	// [7] = [WRITE] chainConfig
	// ··········· CHECK
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewPoolProxyLockOrBurnInstructionBuilder creates a new `PoolProxyLockOrBurn` instruction builder.
func NewPoolProxyLockOrBurnInstructionBuilder() *PoolProxyLockOrBurn {
	nd := &PoolProxyLockOrBurn{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetLockOrBurn sets the "lockOrBurn" parameter.
func (inst *PoolProxyLockOrBurn) SetLockOrBurn(lockOrBurn LockOrBurnInV1) *PoolProxyLockOrBurn {
	inst.LockOrBurn = &lockOrBurn
	return inst
}

// SetTestPoolAccount sets the "testPool" account.
// CHECK
func (inst *PoolProxyLockOrBurn) SetTestPoolAccount(testPool ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(testPool)
	return inst
}

// GetTestPoolAccount gets the "testPool" account.
// CHECK
func (inst *PoolProxyLockOrBurn) GetTestPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetCpiSignerAccount sets the "cpiSigner" account.
// CHECK
func (inst *PoolProxyLockOrBurn) SetCpiSignerAccount(cpiSigner ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(cpiSigner)
	return inst
}

// GetCpiSignerAccount gets the "cpiSigner" account.
// CHECK
func (inst *PoolProxyLockOrBurn) GetCpiSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetStateAccount sets the "state" account.
// CHECK
func (inst *PoolProxyLockOrBurn) SetStateAccount(state ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
// CHECK
func (inst *PoolProxyLockOrBurn) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// CHECK
func (inst *PoolProxyLockOrBurn) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// CHECK
func (inst *PoolProxyLockOrBurn) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetMintAccount sets the "mint" account.
func (inst *PoolProxyLockOrBurn) SetMintAccount(mint ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *PoolProxyLockOrBurn) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetPoolSignerAccount sets the "poolSigner" account.
// CHECK
func (inst *PoolProxyLockOrBurn) SetPoolSignerAccount(poolSigner ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(poolSigner)
	return inst
}

// GetPoolSignerAccount gets the "poolSigner" account.
// CHECK
func (inst *PoolProxyLockOrBurn) GetPoolSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetPoolTokenAccountAccount sets the "poolTokenAccount" account.
func (inst *PoolProxyLockOrBurn) SetPoolTokenAccountAccount(poolTokenAccount ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolTokenAccount).WRITE()
	return inst
}

// GetPoolTokenAccountAccount gets the "poolTokenAccount" account.
func (inst *PoolProxyLockOrBurn) GetPoolTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetChainConfigAccount sets the "chainConfig" account.
// CHECK
func (inst *PoolProxyLockOrBurn) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
// CHECK
func (inst *PoolProxyLockOrBurn) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

func (inst PoolProxyLockOrBurn) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PoolProxyLockOrBurn,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PoolProxyLockOrBurn) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PoolProxyLockOrBurn) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LockOrBurn == nil {
			return errors.New("LockOrBurn parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TestPool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CpiSigner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PoolSigner is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ChainConfig is not set")
		}
	}
	return nil
}

func (inst *PoolProxyLockOrBurn) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PoolProxyLockOrBurn")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("LockOrBurn", *inst.LockOrBurn))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    testPool", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("   cpiSigner", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("       state", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("tokenProgram", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("        mint", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("  poolSigner", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("   poolToken", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta(" chainConfig", inst.AccountMetaSlice[7]))
					})
				})
		})
}

func (obj PoolProxyLockOrBurn) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LockOrBurn` param:
	err = encoder.Encode(obj.LockOrBurn)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PoolProxyLockOrBurn) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LockOrBurn`:
	err = decoder.Decode(&obj.LockOrBurn)
	if err != nil {
		return err
	}
	return nil
}

// NewPoolProxyLockOrBurnInstruction declares a new PoolProxyLockOrBurn instruction with the provided parameters and accounts.
func NewPoolProxyLockOrBurnInstruction(
	// Parameters:
	lockOrBurn LockOrBurnInV1,
	// Accounts:
	testPool ag_solanago.PublicKey,
	cpiSigner ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	poolSigner ag_solanago.PublicKey,
	poolTokenAccount ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey) *PoolProxyLockOrBurn {
	return NewPoolProxyLockOrBurnInstructionBuilder().
		SetLockOrBurn(lockOrBurn).
		SetTestPoolAccount(testPool).
		SetCpiSignerAccount(cpiSigner).
		SetStateAccount(state).
		SetTokenProgramAccount(tokenProgram).
		SetMintAccount(mint).
		SetPoolSignerAccount(poolSigner).
		SetPoolTokenAccountAccount(poolTokenAccount).
		SetChainConfigAccount(chainConfig)
}
