package tokens

import (
	"context"
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/gagliardetto/solana-go"
	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

// parallel test execution can cause race conditions for setting the token program in the solana_go SDK
// the race occurs when multiple programs attempt to utilize the solana_go/programs/token and set different token programs to build transactions from
// the SDK uses a global program ID that is called via ProgramID(): https://github.com/gagliardetto/solana-go/blob/da2193071f56059aa35010a239cece016c4e827f/programs/token/instructions.go#L310
// this is called when the transaction is assembled from multiple instructions (not in ValidateAndBuild) - so it is not static until NewTransaction() is called
var _ solana.Instruction = (*TokenInstruction)(nil)

// TokenInstruction wraps the base token instruction and provides the requested ProgramID rather than depending on the SDK global
type TokenInstruction struct {
	solana.Instruction
	Program solana.PublicKey
}

// ProgramID overrides the default solana.Instruction.ProgramID behavior
func (inst *TokenInstruction) ProgramID() solana.PublicKey {
	return inst.Program
}

// NOTE: functions in this file are mainly wrapped version of the versions that exist in `solana-go` but these allow specifying the token program

func CreateToken(ctx context.Context, program, mint, admin solana.PublicKey, decimals uint8, client *rpc.Client, commitment rpc.CommitmentType) ([]solana.Instruction, error) {
	// get stake amount for init
	lamports, err := client.GetMinimumBalanceForRentExemption(ctx, token.MINT_SIZE, commitment)
	if err != nil {
		return nil, err
	}

	// initialize mint account
	initI, err := system.NewCreateAccountInstruction(lamports, token.MINT_SIZE, program, admin, mint).ValidateAndBuild()
	if err != nil {
		return nil, err
	}

	// initialize mint
	mintI, err := token.NewInitializeMintInstruction(decimals, admin, admin, mint, solana.SysVarRentPubkey).ValidateAndBuild()
	if err != nil {
		return nil, err
	}

	mintWrap := &TokenInstruction{mintI, program}
	return []solana.Instruction{initI, mintWrap}, nil
}

var AssociatedTokenProgramID solana.PublicKey = ata.ProgramID

func CreateAssociatedTokenAccount(tokenProgram, mint, address, payer solana.PublicKey) (ins solana.Instruction, ataAddress solana.PublicKey, err error) {
	base := ata.NewCreateInstruction(payer, address, mint)
	ix, err := base.ValidateAndBuild()
	if err != nil {
		return nil, solana.PublicKey{}, err
	}

	associatedTokenAddress, _, err := FindAssociatedTokenAddress(tokenProgram, mint, address)
	if err != nil {
		return nil, solana.PublicKey{}, err
	}

	accounts := ix.Accounts()
	// replace token program with specified tokenProgram
	accounts[5].PublicKey = tokenProgram
	// replace associated token account with the new one derived with the proper program
	accounts[1].PublicKey = associatedTokenAddress

	// set into instruction
	base.AccountMetaSlice = accounts
	ix.Impl = base
	return ix, associatedTokenAddress, nil
}

func FindAssociatedTokenAddress(tokenProgram, mint, address solana.PublicKey) (addr solana.PublicKey, bump uint8, err error) {
	// NOTE: Do not use `solana.FindAssociatedTokenAddress` as it hardcodes the token program ID to the pre-2022 one.
	// So, instead, manually construct the seeds for associated token addresses.
	return solana.FindProgramAddress([][]byte{
		address[:],
		tokenProgram[:],
		mint[:],
	},
		solana.SPLAssociatedTokenAccountProgramID,
	)
}

func MintTo(amount uint64, program, mint, toAddress, authority solana.PublicKey) (solana.Instruction, error) {
	ix, err := token.NewMintToInstruction(amount, mint, toAddress, authority, nil).ValidateAndBuild()
	return &TokenInstruction{ix, program}, err
}

func TokenTransferChecked(amount uint64, decimals uint8, program, source, mint, destination, owner solana.PublicKey, multisigSigners []solana.PublicKey) (solana.Instruction, error) {
	ix, err := token.NewTransferCheckedInstruction(amount, decimals, source, mint, destination, owner, multisigSigners).ValidateAndBuild()
	return &TokenInstruction{ix, program}, err
}

func SetTokenMintAuthority(program, newAuth, mint, signer solana.PublicKey) (solana.Instruction, error) {
	ix, err := token.NewSetAuthorityInstruction(token.AuthorityMintTokens, newAuth, mint, signer, solana.PublicKeySlice{}).ValidateAndBuild()
	return &TokenInstruction{ix, program}, err
}

func TokenApproveChecked(amount uint64, decimals uint8, program, source, mint, delegate, owner solana.PublicKey, multisigSigners []solana.PublicKey) (solana.Instruction, error) {
	ix, err := token.NewApproveCheckedInstruction(amount, decimals, source, mint, delegate, owner, multisigSigners).ValidateAndBuild()
	return &TokenInstruction{ix, program}, err
}

func TokenSupply(ctx context.Context, client *rpc.Client, mint solana.PublicKey, commitment rpc.CommitmentType) (uint8, int, error) {
	res, err := client.GetTokenSupply(ctx, mint, commitment)
	if err != nil {
		return 0, 0, err
	}
	if res == nil || res.Value == nil {
		return 0, 0, fmt.Errorf("rpc returned nil")
	}
	v, err := strconv.Atoi(res.Value.Amount)
	return res.Value.Decimals, v, err
}

func TokenBalance(ctx context.Context, client *rpc.Client, acc solana.PublicKey, commitment rpc.CommitmentType) (uint8, int, error) {
	res, err := client.GetTokenAccountBalance(ctx, acc, commitment)
	if err != nil {
		return 0, 0, err
	}
	if res == nil || res.Value == nil {
		return 0, 0, fmt.Errorf("rpc returned nil")
	}
	v, err := strconv.Atoi(res.Value.Amount)
	return res.Value.Decimals, v, err
}

func NativeTransfer(program solana.PublicKey, lamports uint64, from solana.PublicKey, to solana.PublicKey) (solana.Instruction, error) {
	return system.NewTransferInstruction(lamports, from, to).ValidateAndBuild()
}

func SyncNative(program solana.PublicKey, tokenAccount solana.PublicKey) (solana.Instruction, error) {
	ix, err := token.NewSyncNativeInstruction(tokenAccount).ValidateAndBuild()
	return &TokenInstruction{ix, program}, err
}

func ToLittleEndianU256(v uint64) [32]byte {
	out := [32]byte{}
	binary.LittleEndian.PutUint64(out[:], v)
	return out
}
