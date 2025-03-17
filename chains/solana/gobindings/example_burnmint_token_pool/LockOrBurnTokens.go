// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_burnmint_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// LockOrBurnTokens is the `lockOrBurnTokens` instruction.
type LockOrBurnTokens struct {
	LockOrBurn *LockOrBurnInV1

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] state
	//
	// [2] = [] tokenProgram
	//
	// [3] = [WRITE] mint
	//
	// [4] = [] poolSigner
	//
	// [5] = [WRITE] poolTokenAccount
	//
	// [6] = [] rmnRemote
	//
	// [7] = [] rmnRemoteCurses
	//
	// [8] = [] rmnRemoteConfig
	//
	// [9] = [WRITE] chainConfig
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewLockOrBurnTokensInstructionBuilder creates a new `LockOrBurnTokens` instruction builder.
func NewLockOrBurnTokensInstructionBuilder() *LockOrBurnTokens {
	nd := &LockOrBurnTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetLockOrBurn sets the "lockOrBurn" parameter.
func (inst *LockOrBurnTokens) SetLockOrBurn(lockOrBurn LockOrBurnInV1) *LockOrBurnTokens {
	inst.LockOrBurn = &lockOrBurn
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *LockOrBurnTokens) SetAuthorityAccount(authority ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *LockOrBurnTokens) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetStateAccount sets the "state" account.
func (inst *LockOrBurnTokens) SetStateAccount(state ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *LockOrBurnTokens) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *LockOrBurnTokens) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *LockOrBurnTokens) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetMintAccount sets the "mint" account.
func (inst *LockOrBurnTokens) SetMintAccount(mint ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *LockOrBurnTokens) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetPoolSignerAccount sets the "poolSigner" account.
func (inst *LockOrBurnTokens) SetPoolSignerAccount(poolSigner ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(poolSigner)
	return inst
}

// GetPoolSignerAccount gets the "poolSigner" account.
func (inst *LockOrBurnTokens) GetPoolSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetPoolTokenAccountAccount sets the "poolTokenAccount" account.
func (inst *LockOrBurnTokens) SetPoolTokenAccountAccount(poolTokenAccount ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(poolTokenAccount).WRITE()
	return inst
}

// GetPoolTokenAccountAccount gets the "poolTokenAccount" account.
func (inst *LockOrBurnTokens) GetPoolTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetRmnRemoteAccount sets the "rmnRemote" account.
func (inst *LockOrBurnTokens) SetRmnRemoteAccount(rmnRemote ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rmnRemote)
	return inst
}

// GetRmnRemoteAccount gets the "rmnRemote" account.
func (inst *LockOrBurnTokens) GetRmnRemoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetRmnRemoteCursesAccount sets the "rmnRemoteCurses" account.
func (inst *LockOrBurnTokens) SetRmnRemoteCursesAccount(rmnRemoteCurses ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rmnRemoteCurses)
	return inst
}

// GetRmnRemoteCursesAccount gets the "rmnRemoteCurses" account.
func (inst *LockOrBurnTokens) GetRmnRemoteCursesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetRmnRemoteConfigAccount sets the "rmnRemoteConfig" account.
func (inst *LockOrBurnTokens) SetRmnRemoteConfigAccount(rmnRemoteConfig ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rmnRemoteConfig)
	return inst
}

// GetRmnRemoteConfigAccount gets the "rmnRemoteConfig" account.
func (inst *LockOrBurnTokens) GetRmnRemoteConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetChainConfigAccount sets the "chainConfig" account.
func (inst *LockOrBurnTokens) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *LockOrBurnTokens {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
func (inst *LockOrBurnTokens) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

func (inst LockOrBurnTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_LockOrBurnTokens,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst LockOrBurnTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *LockOrBurnTokens) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LockOrBurn == nil {
			return errors.New("LockOrBurn parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PoolSigner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PoolTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.RmnRemote is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.RmnRemoteCurses is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.RmnRemoteConfig is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ChainConfig is not set")
		}
	}
	return nil
}

func (inst *LockOrBurnTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("LockOrBurnTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("LockOrBurn", *inst.LockOrBurn))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("          state", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("           mint", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("     poolSigner", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("      poolToken", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("      rmnRemote", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("rmnRemoteCurses", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("rmnRemoteConfig", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("    chainConfig", inst.AccountMetaSlice[9]))
					})
				})
		})
}

func (obj LockOrBurnTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LockOrBurn` param:
	err = encoder.Encode(obj.LockOrBurn)
	if err != nil {
		return err
	}
	return nil
}
func (obj *LockOrBurnTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LockOrBurn`:
	err = decoder.Decode(&obj.LockOrBurn)
	if err != nil {
		return err
	}
	return nil
}

// NewLockOrBurnTokensInstruction declares a new LockOrBurnTokens instruction with the provided parameters and accounts.
func NewLockOrBurnTokensInstruction(
	// Parameters:
	lockOrBurn LockOrBurnInV1,
	// Accounts:
	authority ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	poolSigner ag_solanago.PublicKey,
	poolTokenAccount ag_solanago.PublicKey,
	rmnRemote ag_solanago.PublicKey,
	rmnRemoteCurses ag_solanago.PublicKey,
	rmnRemoteConfig ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey) *LockOrBurnTokens {
	return NewLockOrBurnTokensInstructionBuilder().
		SetLockOrBurn(lockOrBurn).
		SetAuthorityAccount(authority).
		SetStateAccount(state).
		SetTokenProgramAccount(tokenProgram).
		SetMintAccount(mint).
		SetPoolSignerAccount(poolSigner).
		SetPoolTokenAccountAccount(poolTokenAccount).
		SetRmnRemoteAccount(rmnRemote).
		SetRmnRemoteCursesAccount(rmnRemoteCurses).
		SetRmnRemoteConfigAccount(rmnRemoteConfig).
		SetChainConfigAccount(chainConfig)
}
