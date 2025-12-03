package tokens

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/gagliardetto/solana-go"
	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
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
	return CreateTokenWith(ctx, program, mint, admin, admin, decimals, client, commitment, false)
}

func CreateTokenFrom(ctx context.Context, newToken TokenPool, admin solana.PublicKey, decimals uint8, client *rpc.Client, commitment rpc.CommitmentType) ([]solana.Instruction, error) {
	return CreateTokenWith(ctx, newToken.Program, newToken.Mint, admin, admin, decimals, client, commitment, newToken.WithTokenExtensions)
}

func CreateTokenWith(ctx context.Context, program, mint, mintAuthority, freezeAuthority solana.PublicKey, decimals uint8, client *rpc.Client, commitment rpc.CommitmentType, createWithExtensions bool) ([]solana.Instruction, error) {
	ixs := []solana.Instruction{}

	// initialize mint account
	mintSize := token.MINT_SIZE
	if program == config.Token2022Program && createWithExtensions {
		mintSize += 1 + 36 + 83 // MintCloseAuthorityExtension overhead: header + ext + padding
	}

	// get stake amount for init
	lamports, err := client.GetMinimumBalanceForRentExemption(ctx, uint64(mintSize), commitment) //nolint:gosec
	if err != nil {
		return nil, err
	}

	initI, err := system.NewCreateAccountInstruction(lamports, uint64(mintSize), program, mintAuthority, mint).ValidateAndBuild() //nolint:gosec
	if err != nil {
		return nil, err
	}

	ixs = append(ixs, initI)

	if program == config.Token2022Program && createWithExtensions {
		closeMintExtensionI, closeMintErr := NewInitializeMintCloseAuthorityIx(mint, &mintAuthority, &program)
		if closeMintErr != nil {
			return nil, closeMintErr
		}
		ixs = append(ixs, closeMintExtensionI)
	}

	// initialize mint
	mintI, err := token.NewInitializeMintInstruction(decimals, mintAuthority, freezeAuthority, mint, solana.SysVarRentPubkey).ValidateAndBuild()
	if err != nil {
		return nil, err
	}

	mintWrap := &TokenInstruction{mintI, program}
	ixs = append(ixs, mintWrap)
	return ixs, nil
}

func NewInitializeMintCloseAuthorityIx(
	mint solana.PublicKey,
	closeAuthority *solana.PublicKey,
	programID *solana.PublicKey,
) (solana.Instruction, error) {
	data := buildIxData(closeAuthority)

	return &solana.GenericInstruction{
		ProgID: *programID,
		AccountValues: solana.AccountMetaSlice{
			{PublicKey: mint, IsSigner: false, IsWritable: true},
		},
		DataBytes: data,
	}, nil
}

func buildIxData(closeAuthority *solana.PublicKey) []byte {
	var buf bytes.Buffer
	buf.WriteByte(25)

	if closeAuthority == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		buf.Write(closeAuthority.Bytes())
	}

	return buf.Bytes()
}

func CreateMultisig(ctx context.Context, payer, program, multisig solana.PublicKey, m uint8, signers []solana.PublicKey, client *rpc.Client, commitment rpc.CommitmentType) ([]solana.Instruction, error) {
	// get stake amount for init
	lamports, err := client.GetMinimumBalanceForRentExemption(ctx, 355, commitment)
	if err != nil {
		return nil, err
	}

	// initialize mint account
	initI, err := system.NewCreateAccountInstruction(lamports, 355, program, payer, multisig).ValidateAndBuild()
	if err != nil {
		return nil, err
	}

	// Manually add the signer metas, as the SDK wrongly tries to set them as transaction signers
	// when they are just meant to be registered as part of the multisig
	raw := token.NewInitializeMultisig2Instruction(m, multisig, []solana.PublicKey{})
	for _, signer := range signers {
		raw.Signers = append(raw.Signers, solana.Meta(signer))
	}
	msigInitIx, err := raw.ValidateAndBuild()
	if err != nil {
		return nil, err
	}

	msigWrap := &TokenInstruction{msigInitIx, program}
	return []solana.Instruction{initI, msigWrap}, nil
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

func NativeTransfer(lamports uint64, from solana.PublicKey, to solana.PublicKey) (solana.Instruction, error) {
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
