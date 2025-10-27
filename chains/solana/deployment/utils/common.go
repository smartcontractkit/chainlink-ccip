package utils

import (
	"fmt"
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	"github.com/smartcontractkit/mcms/types"
)

const TimelockProgramType cldf_deployment.ContractType = "RBACTimelockProgram"

// Common parameters for transferring ownership of a program
type TransferOwnershipParams struct {
	Program      solana.PublicKey
	CurrentOwner solana.PublicKey
	NewOwner     solana.PublicKey
}

func BuildMCMSBatchOperation(
	chainSelector uint64,
	ixns []solana.Instruction,
	programID string,
	contractType string) (types.BatchOperation, error) {
	txns := make([]types.Transaction, 0, len(ixns))
	for _, ixn := range ixns {
		data, err := ixn.Data()
		if err != nil {
			return types.BatchOperation{}, fmt.Errorf("failed to extract data: %w", err)
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
			return types.BatchOperation{}, fmt.Errorf("failed to create transaction: %w", err)
		}
		txns = append(txns, tx)
	}
	return types.BatchOperation{
		ChainSelector: types.ChainSelector(chainSelector),
		Transactions:  txns,
	}, nil
}

func GetTimelockSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	qualifier string) solana.PublicKey {
	timelockProgram := datastore.GetAddressRef(
		existingAddresses,
		TimelockProgramType,
		common_utils.Version_1_6_0,
		"",
	)
	// timelock seeds stored as a separate program type
	// qualifier will identify the correct timelock instance
	timelockSeed := datastore.GetAddressRef(
		existingAddresses,
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0,
		qualifier,
	)
	return state.GetTimelockSignerPDA(
		solana.MustPublicKeyFromBase58(timelockProgram.Address),
		state.PDASeed([]byte(timelockSeed.Address)),
	)
}
