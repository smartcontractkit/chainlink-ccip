package utils

import (
	"context"
	"fmt"

	solbinary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	solrpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
)

func GetSolProgramSize(chain cldf_solana.Chain, programID solana.PublicKey) (int, error) {
	accountInfo, err := chain.Client.GetAccountInfoWithOpts(context.Background(), programID, &solrpc.GetAccountInfoOpts{
		Commitment: cldf_solana.SolDefaultCommitment,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to get account info: %w", err)
	}
	if accountInfo == nil {
		return 0, fmt.Errorf("program account not found: %w", err)
	}
	programBytes := len(accountInfo.Value.Data.GetBinary())
	return programBytes, nil
}

func GetSolProgramData(chain cldf_solana.Chain, programID solana.PublicKey) (struct {
	DataType uint32
	Address  solana.PublicKey
}, error) {
	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	data, err := chain.Client.GetAccountInfoWithOpts(context.Background(), programID, &solrpc.GetAccountInfoOpts{
		Commitment: solrpc.CommitmentConfirmed,
	})
	if err != nil {
		return programData, fmt.Errorf("failed to deploy program: %w", err)
	}

	err = solbinary.UnmarshalBorsh(&programData, data.Bytes())
	if err != nil {
		return programData, fmt.Errorf("failed to unmarshal program data: %w", err)
	}
	return programData, nil
}

func ExtendLookupTable(chain cldf_solana.Chain, offRampID solana.PublicKey, lookUpTableEntries []solana.PublicKey) error {
	var referenceAddressesAccount ccip_offramp.ReferenceAddresses
	offRampReferenceAddressesPDA, _, _ := state.FindOfframpReferenceAddressesPDA(offRampID)
	err := chain.GetAccountDataBorshInto(context.Background(), offRampReferenceAddressesPDA, &referenceAddressesAccount)
	if err != nil {
		return fmt.Errorf("failed to get offramp reference addresses: %w", err)
	}
	addressLookupTable := referenceAddressesAccount.OfframpLookupTable

	addresses, err := common.GetAddressLookupTable(
		context.Background(),
		chain.Client,
		addressLookupTable)
	if err != nil {
		return fmt.Errorf("failed to get address lookup table: %w", err)
	}

	// calculate diff and add new entries
	seen := make(map[solana.PublicKey]bool)
	toAdd := make([]solana.PublicKey, 0)
	for _, entry := range addresses {
		seen[entry] = true
	}
	for _, entry := range lookUpTableEntries {
		if _, ok := seen[entry]; !ok {
			toAdd = append(toAdd, entry)
		}
	}
	if len(toAdd) == 0 {
		return nil
	}

	if err := common.ExtendLookupTable(
		context.Background(),
		chain.Client,
		addressLookupTable,
		*chain.DeployerKey,
		toAdd,
	); err != nil {
		return fmt.Errorf("failed to extend lookup table: %w", err)
	}
	return nil
}
