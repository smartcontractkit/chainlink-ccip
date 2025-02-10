// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ReleaseOrMintTokens is the `releaseOrMintTokens` instruction.
type ReleaseOrMintTokens struct {
	ReleaseOrMint *ReleaseOrMintInV1

	// [0] = [SIGNER] authority
	//
	// [1] = [] offrampProgram
	// ··········· CHECK offramp program: exists only to derive the allowed offramp PDA
	// ··········· and the authority PDA.
	//
	// [2] = [] allowedOfframp
	// ··········· CHECK PDA of the router program verifying the signer is an allowed offramp.
	// ··········· If PDA does not exist, the router doesn't allow this offramp
	//
	// [3] = [WRITE] state
	//
	// [4] = [] tokenProgram
	//
	// [5] = [WRITE] mint
	//
	// [6] = [] poolSigner
	//
	// [7] = [WRITE] poolTokenAccount
	//
	// [8] = [WRITE] chainConfig
	//
	// [9] = [WRITE] receiverTokenAccount
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewReleaseOrMintTokensInstructionBuilder creates a new `ReleaseOrMintTokens` instruction builder.
func NewReleaseOrMintTokensInstructionBuilder() *ReleaseOrMintTokens {
	nd := &ReleaseOrMintTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetReleaseOrMint sets the "releaseOrMint" parameter.
func (inst *ReleaseOrMintTokens) SetReleaseOrMint(releaseOrMint ReleaseOrMintInV1) *ReleaseOrMintTokens {
	inst.ReleaseOrMint = &releaseOrMint
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ReleaseOrMintTokens) SetAuthorityAccount(authority ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ReleaseOrMintTokens) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOfframpProgramAccount sets the "offrampProgram" account.
// CHECK offramp program: exists only to derive the allowed offramp PDA
// and the authority PDA.
func (inst *ReleaseOrMintTokens) SetOfframpProgramAccount(offrampProgram ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(offrampProgram)
	return inst
}

// GetOfframpProgramAccount gets the "offrampProgram" account.
// CHECK offramp program: exists only to derive the allowed offramp PDA
// and the authority PDA.
func (inst *ReleaseOrMintTokens) GetOfframpProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAllowedOfframpAccount sets the "allowedOfframp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp
func (inst *ReleaseOrMintTokens) SetAllowedOfframpAccount(allowedOfframp ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(allowedOfframp)
	return inst
}

// GetAllowedOfframpAccount gets the "allowedOfframp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp
func (inst *ReleaseOrMintTokens) GetAllowedOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetStateAccount sets the "state" account.
func (inst *ReleaseOrMintTokens) SetStateAccount(state ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ReleaseOrMintTokens) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *ReleaseOrMintTokens) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *ReleaseOrMintTokens) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetMintAccount sets the "mint" account.
func (inst *ReleaseOrMintTokens) SetMintAccount(mint ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *ReleaseOrMintTokens) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetPoolSignerAccount sets the "poolSigner" account.
func (inst *ReleaseOrMintTokens) SetPoolSignerAccount(poolSigner ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolSigner)
	return inst
}

// GetPoolSignerAccount gets the "poolSigner" account.
func (inst *ReleaseOrMintTokens) GetPoolSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetPoolTokenAccountAccount sets the "poolTokenAccount" account.
func (inst *ReleaseOrMintTokens) SetPoolTokenAccountAccount(poolTokenAccount ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(poolTokenAccount).WRITE()
	return inst
}

// GetPoolTokenAccountAccount gets the "poolTokenAccount" account.
func (inst *ReleaseOrMintTokens) GetPoolTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetChainConfigAccount sets the "chainConfig" account.
func (inst *ReleaseOrMintTokens) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
func (inst *ReleaseOrMintTokens) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetReceiverTokenAccountAccount sets the "receiverTokenAccount" account.
func (inst *ReleaseOrMintTokens) SetReceiverTokenAccountAccount(receiverTokenAccount ag_solanago.PublicKey) *ReleaseOrMintTokens {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(receiverTokenAccount).WRITE()
	return inst
}

// GetReceiverTokenAccountAccount gets the "receiverTokenAccount" account.
func (inst *ReleaseOrMintTokens) GetReceiverTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

func (inst ReleaseOrMintTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ReleaseOrMintTokens,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ReleaseOrMintTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ReleaseOrMintTokens) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ReleaseOrMint == nil {
			return errors.New("ReleaseOrMint parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OfframpProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AllowedOfframp is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolSigner is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PoolTokenAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.ChainConfig is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ReceiverTokenAccount is not set")
		}
	}
	return nil
}

func (inst *ReleaseOrMintTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ReleaseOrMintTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ReleaseOrMint", *inst.ReleaseOrMint))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("offrampProgram", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("allowedOfframp", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("          mint", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("    poolSigner", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("     poolToken", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("   chainConfig", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta(" receiverToken", inst.AccountMetaSlice[9]))
					})
				})
		})
}

func (obj ReleaseOrMintTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ReleaseOrMint` param:
	err = encoder.Encode(obj.ReleaseOrMint)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ReleaseOrMintTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ReleaseOrMint`:
	err = decoder.Decode(&obj.ReleaseOrMint)
	if err != nil {
		return err
	}
	return nil
}

// NewReleaseOrMintTokensInstruction declares a new ReleaseOrMintTokens instruction with the provided parameters and accounts.
func NewReleaseOrMintTokensInstruction(
	// Parameters:
	releaseOrMint ReleaseOrMintInV1,
	// Accounts:
	authority ag_solanago.PublicKey,
	offrampProgram ag_solanago.PublicKey,
	allowedOfframp ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	poolSigner ag_solanago.PublicKey,
	poolTokenAccount ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey,
	receiverTokenAccount ag_solanago.PublicKey) *ReleaseOrMintTokens {
	return NewReleaseOrMintTokensInstructionBuilder().
		SetReleaseOrMint(releaseOrMint).
		SetAuthorityAccount(authority).
		SetOfframpProgramAccount(offrampProgram).
		SetAllowedOfframpAccount(allowedOfframp).
		SetStateAccount(state).
		SetTokenProgramAccount(tokenProgram).
		SetMintAccount(mint).
		SetPoolSignerAccount(poolSigner).
		SetPoolTokenAccountAccount(poolTokenAccount).
		SetChainConfigAccount(chainConfig).
		SetReceiverTokenAccountAccount(receiverTokenAccount)
}
