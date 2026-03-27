package utils

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/gagliardetto/solana-go"
	solrpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
)

// SetUpgradeAuthority creates a BPF Loader Upgradeable SetAuthority instruction (ID=4).
// isBuffer indicates whether the target is a buffer account (true) or a deployed program (false).
func SetUpgradeAuthority(
	programID solana.PublicKey,
	currentAuthority solana.PublicKey,
	newAuthority solana.PublicKey,
	isBuffer bool,
) solana.Instruction {
	var target solana.PublicKey
	if isBuffer {
		target = programID
	} else {
		target, _, _ = solana.FindProgramAddress([][]byte{programID.Bytes()}, solana.BPFLoaderUpgradeableProgramID)
	}

	keys := solana.AccountMetaSlice{
		solana.NewAccountMeta(target, true, false),
		solana.NewAccountMeta(currentAuthority, false, true),
		solana.NewAccountMeta(newAuthority, false, false),
	}

	return solana.NewInstruction(
		solana.BPFLoaderUpgradeableProgramID,
		keys,
		[]byte{4, 0, 0, 0},
	)
}

// GenerateUpgradeInstruction creates a BPF Loader Upgradeable Upgrade instruction (ID=3).
// This replaces the deployed program's binary with the contents of the buffer.
func GenerateUpgradeInstruction(
	programID solana.PublicKey,
	bufferAddress solana.PublicKey,
	spillAddress solana.PublicKey,
	upgradeAuthority solana.PublicKey,
) solana.Instruction {
	programDataAccount, _, _ := solana.FindProgramAddress([][]byte{programID.Bytes()}, solana.BPFLoaderUpgradeableProgramID)

	keys := solana.AccountMetaSlice{
		solana.NewAccountMeta(programDataAccount, true, false),
		solana.NewAccountMeta(programID, true, false),
		solana.NewAccountMeta(bufferAddress, true, false),
		solana.NewAccountMeta(spillAddress, true, false),
		solana.NewAccountMeta(solana.SysVarRentPubkey, false, false),
		solana.NewAccountMeta(solana.SysVarClockPubkey, false, false),
		solana.NewAccountMeta(upgradeAuthority, false, true),
	}

	return solana.NewInstruction(
		solana.BPFLoaderUpgradeableProgramID,
		keys,
		[]byte{3, 0, 0, 0},
	)
}

// GenerateCloseBufferInstruction creates a BPF Loader Upgradeable Close instruction (ID=5)
// for closing a buffer account and reclaiming its rent.
func GenerateCloseBufferInstruction(
	bufferAddress solana.PublicKey,
	recipient solana.PublicKey,
	authority solana.PublicKey,
) solana.Instruction {
	keys := solana.AccountMetaSlice{
		solana.Meta(bufferAddress).WRITE(),
		solana.Meta(recipient).WRITE(),
		solana.Meta(authority).SIGNER(),
	}

	return solana.NewInstruction(
		solana.BPFLoaderUpgradeableProgramID,
		keys,
		[]byte{5, 0, 0, 0},
	)
}

// GenerateExtendInstruction creates a BPF Loader Upgradeable ExtendProgram instruction (ID=6).
// This extends the program data account to accommodate a larger binary. This is permissionless;
// any payer can extend any program's buffer.
// Returns nil if the program already has enough space.
func GenerateExtendInstruction(
	chain cldf_solana.Chain,
	programID solana.PublicKey,
	payer solana.PublicKey,
	newBinarySize int,
) (*solana.GenericInstruction, error) {
	programDataAccount, _, _ := solana.FindProgramAddress([][]byte{programID.Bytes()}, solana.BPFLoaderUpgradeableProgramID)

	currentSize, err := GetSolProgramSize(chain, programDataAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to get current program size: %w", err)
	}

	if newBinarySize <= currentSize {
		return nil, nil
	}

	extraBytes := newBinarySize - currentSize
	data := binary.LittleEndian.AppendUint32([]byte{}, 6) // ExtendProgram instruction identifier
	//nolint:gosec // G115 - checked above
	data = binary.LittleEndian.AppendUint32(data, uint32(extraBytes+1024)) // padding for future growth

	keys := solana.AccountMetaSlice{
		solana.NewAccountMeta(programDataAccount, true, false),
		solana.NewAccountMeta(programID, true, false),
		solana.NewAccountMeta(solana.SystemProgramID, false, false),
		solana.NewAccountMeta(payer, true, true),
	}

	return solana.NewInstruction(
		solana.BPFLoaderUpgradeableProgramID,
		keys,
		data,
	), nil
}

// DeployToBuffer deploys a program binary to a buffer account using `chain.DeployProgram`
// with isUpgrade=true. Returns the buffer's public key.
func DeployToBuffer(chain cldf_solana.Chain, lggr logger.Logger, programName string) (solana.PublicKey, error) {
	bufferID, err := chain.DeployProgram(lggr, cldf_solana.ProgramInfo{
		Name: programName,
	}, true, false)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to deploy program to buffer: %w", err)
	}
	return solana.MustPublicKeyFromBase58(bufferID), nil
}

// GetBufferSize returns the size of a buffer or program data account.
func GetBufferSize(chain cldf_solana.Chain, account solana.PublicKey) (int, error) {
	accountInfo, err := chain.Client.GetAccountInfoWithOpts(context.Background(), account, &solrpc.GetAccountInfoOpts{
		Commitment: cldf_solana.SolDefaultCommitment,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to get account info: %w", err)
	}
	if accountInfo == nil || accountInfo.Value == nil {
		return 0, fmt.Errorf("account not found: %s", account.String())
	}
	return len(accountInfo.Value.Data.GetBinary()), nil
}
