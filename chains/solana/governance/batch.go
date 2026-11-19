package governance

import (
	"fmt"
	"math/big"

	"github.com/gagliardetto/solana-go"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	"github.com/smartcontractkit/mcms/types"
)

// BuildMCMSTxn converts an instruction into an MCMS transaction. Signer flags are cleared: the
// timelock signs, not the instruction accounts.
func BuildMCMSTxn(ixn solana.Instruction, programID string, contractType string) (*types.Transaction, error) {
	data, err := ixn.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to extract data: %w", err)
	}
	for _, account := range ixn.Accounts() {
		if account.IsSigner {
			account.IsSigner = false
		}
	}
	tx, err := mcms_solana.NewTransaction(
		programID,
		data,
		big.NewInt(0),  // e.g. value
		ixn.Accounts(), // pass along needed accounts
		contractType,   // some string identifying the target
		[]string{},     // any relevant metadata
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
	return &tx, nil
}

// BuildMCMSBatchOperation converts instructions into an MCMS batch operation for one chain.
func BuildMCMSBatchOperation(
	chainSelector uint64,
	ixns []solana.Instruction,
	programID string,
	contractType string) (types.BatchOperation, error) {
	txns := make([]types.Transaction, 0, len(ixns))
	for _, ixn := range ixns {
		tx, err := BuildMCMSTxn(ixn, programID, contractType)
		if err != nil {
			return types.BatchOperation{}, err
		}
		txns = append(txns, *tx)
	}
	return types.BatchOperation{
		ChainSelector: types.ChainSelector(chainSelector),
		Transactions:  txns,
	}, nil
}
