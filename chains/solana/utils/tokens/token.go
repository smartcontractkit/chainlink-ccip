package tokens

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
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

// PROGRAM ID for Metaplex Metadata Program
var MplTokenMetadataID solana.PublicKey = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")

// MplTokenMetadataProgramName is the name used in the Metaplex Metadata Program for the token metadata account
const MplTokenMetadataProgramName = "MplTokenMetadataProgramName"

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

const MultisigSize = 355

func CreateMultisig(ctx context.Context, payer, tokenProgram, multisig solana.PublicKey, m uint8, signers []solana.PublicKey, client *rpc.Client, commitment rpc.CommitmentType) ([]solana.Instruction, error) {
	// get stake amount for init
	lamports, err := client.GetMinimumBalanceForRentExemption(ctx, MultisigSize, commitment)
	if err != nil {
		return nil, err
	}

	return IxsInitTokenMultisig(tokenProgram, lamports, payer, multisig, m, signers)
}

func IxsInitTokenMultisig(tokenProgram solana.PublicKey, lamports uint64, payer solana.PublicKey, multisig solana.PublicKey, m uint8, signers []solana.PublicKey) ([]solana.Instruction, error) {
	if tokenProgram.IsZero() {
		return nil, fmt.Errorf("token program is zero")
	}
	if !tokenProgram.Equals(solana.Token2022ProgramID) && !tokenProgram.Equals(solana.TokenProgramID) {
		return nil, fmt.Errorf("unsupported token program: %s", tokenProgram.String())
	}
	// initialize Multisig account
	initI, err := system.NewCreateAccountInstruction(lamports, MultisigSize, tokenProgram, payer, multisig).ValidateAndBuild()
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

	msigWrap := &TokenInstruction{msigInitIx, tokenProgram}
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

func FindMplTokenMetadataPDA(mint solana.PublicKey) (solana.PublicKey, error) {
	seeds := [][]byte{
		[]byte("metadata"),
		MplTokenMetadataID.Bytes(),
		mint.Bytes(),
	}
	pda, _, err := solana.FindProgramAddress(seeds, MplTokenMetadataID)
	return pda, err
}

func ToLittleEndianU256(v uint64) [32]byte {
	out := [32]byte{}
	binary.LittleEndian.PutUint64(out[:], v)
	return out
}

func GetTokenMetadata(ctx context.Context, client *rpc.Client, mint solana.PublicKey) (token_metadata.Metadata, solana.PublicKey, error) {
	metadataPDA, err := FindMplTokenMetadataPDA(mint)
	if err != nil {
		return token_metadata.Metadata{}, solana.PublicKey{}, fmt.Errorf("failed to find metadata PDA: %w", err)
	}

	var mintMetadata token_metadata.Metadata
	if err := client.GetAccountDataBorshInto(ctx, metadataPDA, &mintMetadata); err != nil {
		return token_metadata.Metadata{}, solana.PublicKey{}, fmt.Errorf("failed to get account data for metadata PDA: %w", err)
	}

	return mintMetadata, metadataPDA, nil
}

func GetTokenDataV2(metadata token_metadata.Metadata) token_metadata.DataV2 {
	return token_metadata.DataV2{
		Symbol:               strings.ReplaceAll(metadata.Data.Symbol, "\x00", ""),
		Name:                 strings.ReplaceAll(metadata.Data.Name, "\x00", ""),
		Uri:                  strings.ReplaceAll(metadata.Data.Uri, "\x00", ""),
		SellerFeeBasisPoints: metadata.Data.SellerFeeBasisPoints,
		Collection:           metadata.Collection,
		Creators:             metadata.Data.Creators,
		Uses:                 metadata.Uses,
	}
}
