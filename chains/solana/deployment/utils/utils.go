package utils

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	solbinary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	solrpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	soltokens "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// Run a command in a specific directory
func RunCommand(command string, args []string, workDir string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = workDir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}

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

func GetSolProgramData(client *solrpc.Client, programID solana.PublicKey) (struct {
	DataType uint32
	Address  solana.PublicKey
}, error) {
	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	data, err := client.GetAccountInfoWithOpts(context.Background(), programID, &solrpc.GetAccountInfoOpts{
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
	err := common.GetAccountDataBorshIntoAnchor(context.Background(), chain.Client, offRampReferenceAddressesPDA, cldf_solana.SolDefaultCommitment, &referenceAddressesAccount)
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

// GetTokenProgramID returns the program ID for the given token program name
func GetTokenProgramID(programName cldf_deployment.ContractType) (solana.PublicKey, error) {
	tokenPrograms := map[cldf_deployment.ContractType]solana.PublicKey{
		SPLTokens:     solana.TokenProgramID,
		SPL2022Tokens: solana.Token2022ProgramID,
	}

	programID, ok := tokenPrograms[programName]
	if !ok {
		return solana.PublicKey{}, fmt.Errorf("invalid token program: %s. Must be one of: %s, %s", programName, SPLTokens, SPL2022Tokens)
	}
	return programID, nil
}

func MintTokens(chain cldf_solana.Chain, tokenProgramID, mint solana.PublicKey, amountToAddress map[string]uint64) error {
	for toAddress, amount := range amountToAddress {
		toAddressBase58 := solana.MustPublicKeyFromBase58(toAddress)
		// get associated token account for toAddress
		ata, _, _ := soltokens.FindAssociatedTokenAddress(tokenProgramID, mint, toAddressBase58)
		mintToI, err := soltokens.MintTo(amount, tokenProgramID, mint, ata, chain.DeployerKey.PublicKey())
		if err != nil {
			return err
		}
		if err := chain.Confirm([]solana.Instruction{mintToI}); err != nil {
			return err
		}
	}
	return nil
}

func DisableFreezeAuthority(chain cldf_solana.Chain, tokenMints []solana.PublicKey) error {
	_, err1 := RunCommand("solana", []string{"config", "set", "--url", chain.URL}, chain.ProgramsPath)
	if err1 != nil {
		return fmt.Errorf("error setting solana url: %w", err1)
	}
	_, err2 := RunCommand("solana", []string{"config", "set", "--keypair", chain.KeypairPath}, chain.ProgramsPath)
	if err2 != nil {
		return fmt.Errorf("error setting solana keypair: %w", err2)
	}

	for _, tokenPubkey := range tokenMints {
		args := []string{"authorize", tokenPubkey.String(), "freeze", "--disable"}
		_, err := RunCommand("spl-token", args, chain.ProgramsPath)
		if err != nil {
			return fmt.Errorf("error disabling freeze authority: %w", err)
		}
	}
	return nil
}

func GetTokenMintAuthority(chain cldf_solana.Chain, tokenMint solana.PublicKey) solana.PublicKey {
	var mintData token.Mint
	var mintAuthority solana.PublicKey
	err := chain.GetAccountDataBorshInto(context.Background(), tokenMint, &mintData)
	if err != nil {
		return solana.PublicKey{}
	}
	if mintData.MintAuthority != nil {
		mintAuthority = *mintData.MintAuthority
	}
	return mintAuthority
}
