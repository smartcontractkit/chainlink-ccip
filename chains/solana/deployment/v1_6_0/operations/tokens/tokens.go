package tokens

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	soltokens "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

var LinkContractType cldf_deployment.ContractType = "LINK"
var Version *semver.Version = semver.MustParse("1.6.0")

type Params struct {
	ExistingAddresses      []datastore.AddressRef
	TokenProgramName       cldf_deployment.ContractType
	TokenPrivKey           solana.PrivateKey
	TokenDecimals          uint8
	TokenSymbol            string
	ATAList                []solana.PublicKey
	PreMint                uint64
	DisableFreezeAuthority bool
}

var DeployLINK = operations.NewOperation(
	"link:deploy",
	Version,
	"Deploys the LINK token contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (datastore.AddressRef, error) {
		instructions, err := tokens.CreateToken(
			context.Background(),
			solana.TokenProgramID,
			input.TokenPrivKey.PublicKey(),
			chain.DeployerKey.PublicKey(),
			input.TokenDecimals,
			chain.Client,
			cldf_solana.SolDefaultCommitment,
		)
		if err != nil {
			return datastore.AddressRef{}, err
		}
		err = chain.Confirm(instructions, common.AddSigners(input.TokenPrivKey))
		if err != nil {
			return datastore.AddressRef{}, err
		}
		return datastore.AddressRef{
			ChainSelector: chain.Selector,
			Address:       input.TokenPrivKey.PublicKey().String(),
			Version:       Version,
			Type:          datastore.ContractType(LinkContractType),
		}, nil
	},
)

var DeploySolanaToken = operations.NewOperation(
	"solana-token:deploy",
	Version,
	"Deploys and configures an SPL token contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (datastore.AddressRef, error) {
		// CREATE TOKEN
		tokenProgramID, err := utils.GetTokenProgramID(input.TokenProgramName)
		if err != nil {
			return datastore.AddressRef{}, err
		}
		freezeAuthority := utils.GetTimelockSignerPDA(input.ExistingAddresses, chain.Selector, common_utils.CLLQualifier)
		tokenAdminPubKey := chain.DeployerKey.PublicKey()
		// if we're disabling the freeze authority, we first set it to the deployer key so it can
		// immediately revoke it
		if input.DisableFreezeAuthority {
			freezeAuthority = chain.DeployerKey.PublicKey()
		}
		var mint solana.PublicKey
		privKey := input.TokenPrivKey
		if privKey.IsValid() {
			mint = privKey.PublicKey()
		} else {
			privKey, err = solana.NewRandomPrivateKey()
			if err != nil {
				return datastore.AddressRef{}, err
			}
			mint = privKey.PublicKey()
		}
		instructions, err := soltokens.CreateTokenWith(
			context.Background(),
			tokenProgramID,
			mint,
			tokenAdminPubKey,
			freezeAuthority,
			input.TokenDecimals,
			chain.Client,
			cldf_solana.SolDefaultCommitment,
			false,
		)
		if err != nil {
			return datastore.AddressRef{}, err
		}
		err = chain.Confirm(instructions, common.AddSigners(privKey))
		if err != nil {
			return datastore.AddressRef{}, err
		}
		// CREATE ATAs
		for _, sender := range input.ATAList {
			createATAIx, _, err := soltokens.CreateAssociatedTokenAccount(
				tokenProgramID,
				mint,
				sender,
				chain.DeployerKey.PublicKey(),
			)
			if err != nil {
				return datastore.AddressRef{}, err
			}
			if err := chain.Confirm([]solana.Instruction{createATAIx}); err != nil {
				return datastore.AddressRef{}, err
			}
			if input.PreMint > 0 {
				ata, _, _ := soltokens.FindAssociatedTokenAddress(tokenProgramID, mint, sender)
				mintToI, err := soltokens.MintTo(input.PreMint, tokenProgramID, mint, ata, chain.DeployerKey.PublicKey())
				if err != nil {
					return datastore.AddressRef{}, err
				}
				if err := chain.Confirm([]solana.Instruction{mintToI}); err != nil {
					return datastore.AddressRef{}, err
				}
			}
		}
		// DISABLE FREEZE AUTHORITY
		if input.DisableFreezeAuthority {
			err = utils.DisableFreezeAuthority(chain, []solana.PublicKey{mint})
			if err != nil {
				return datastore.AddressRef{}, err
			}
		}
		return datastore.AddressRef{
			ChainSelector: chain.Selector,
			Address:       mint.String(),
			Type:          datastore.ContractType(input.TokenProgramName),
			Version:       Version,
			Qualifier:     input.TokenSymbol,
		}, nil
	},
)

var UpsertTokenMetadata = operations.NewOperation(
	"solana-token:upsert-metadata",
	Version,
	"Upserts metadata for an SPL token",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenMetadataInput) (sequences.OnChainOutput, error) {
		batches := make([]types.BatchOperation, 0)
		out1, err1 := utils.RunCommand("solana", []string{"config", "set", "--url", chain.URL}, chain.ProgramsPath)
		if err1 != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("error setting solana url: %s %w", out1, err1)
		}
		out2, err2 := utils.RunCommand("solana", []string{"config", "set", "--keypair", chain.KeypairPath}, chain.ProgramsPath)
		if err2 != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("error setting solana keypair: %s %w", out2, err2)
		}
		metadata := input.Metadata
		tokenMint := solana.MustPublicKeyFromBase58(metadata.TokenPubkey)
		// initial upload
		if metadata.MetadataJSONPath != "" {
			args := []string{"create", "metadata", "--mint", tokenMint.String(), "--metadata", metadata.MetadataJSONPath}
			output, err := utils.RunCommand("metaboss", args, chain.ProgramsPath)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error uploading token metadata: %s %w", output, err)
			}
			return sequences.OnChainOutput{}, nil
		}
		var mintMetadata token_metadata.Metadata
		metadataPDA, metadataErr := FindMplTokenMetadataPDA(tokenMint)
		if metadataErr != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("error finding metadata account: %w", metadataErr)
		}
		fmt.Println("Metadata", metadataPDA)
		if err := chain.GetAccountDataBorshInto(context.Background(), metadataPDA, &mintMetadata); err != nil {
			fmt.Println("error getting metadata account data, skipping update for", tokenMint.String(), ":", err)
			return sequences.OnChainOutput{}, nil
		}
		newUpdateAuthority := mintMetadata.UpdateAuthority
		newData := token_metadata.DataV2{
			Name:   strings.ReplaceAll(mintMetadata.Data.Name, "\x00", ""),
			Symbol: strings.ReplaceAll(mintMetadata.Data.Symbol, "\x00", ""),
			Uri:    strings.ReplaceAll(mintMetadata.Data.Uri, "\x00", ""),
		}
		if metadata.UpdateAuthority != "" {
			newUpdateAuthority = solana.MustPublicKeyFromBase58(metadata.UpdateAuthority)
		}
		if metadata.UpdateName != "" {
			newData.Name = metadata.UpdateName
		}
		if metadata.UpdateSymbol != "" {
			newData.Symbol = metadata.UpdateSymbol
		}
		if metadata.UpdateURI != "" {
			newData.Uri = metadata.UpdateURI
		}
		instruction, err := modifyTokenMetadataIx(
			metadataPDA, mintMetadata.UpdateAuthority, &newUpdateAuthority, &newData)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("error generating modify metadata ix: %w", err)
		}
		if mintMetadata.UpdateAuthority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{&instruction},
				MplTokenMetadataID.String(),
				MplTokenMetadataProgramName,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
		} else {
			err = chain.Confirm([]solana.Instruction{&instruction})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		}
		return sequences.OnChainOutput{
			BatchOps: batches,
		}, nil
	},
)

type TokenMetadataInput struct {
	ExistingAddresses []datastore.AddressRef
	Metadata          tokenapi.TokenMetadata
}

// PROGRAM ID for Metaplex Metadata Program
var MplTokenMetadataID solana.PublicKey = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")

const MplTokenMetadataProgramName = "MplTokenMetadataProgramName"

// discriminator for update_metadata_account_v2 ix
const UpdateMetadataAccountV2Ix = 15

// discriminator for create_metadata_account
const CreateMetadataAccountV2Ix = 16

func FindMplTokenMetadataPDA(mint solana.PublicKey) (solana.PublicKey, error) {
	seeds := [][]byte{
		[]byte("metadata"),
		MplTokenMetadataID.Bytes(),
		mint.Bytes(),
	}
	pda, _, err := solana.FindProgramAddress(seeds, MplTokenMetadataID)
	return pda, err
}

func modifyTokenMetadataIx(
	metadataPDA, authority solana.PublicKey,
	newAuthority *solana.PublicKey,
	newData *token_metadata.DataV2,
) (solana.GenericInstruction, error) {
	args := token_metadata.UpdateMetadataAccountArgsV2{
		Data:            newData,
		UpdateAuthority: newAuthority,
	}
	ix := token_metadata.NewUpdateMetadataAccountV2Instruction(
		args,
		metadataPDA,
		authority).Build()
	data, err := ix.Data()
	if err != nil {
		return solana.GenericInstruction{}, fmt.Errorf("error building update metadata account data: %w", err)
	}

	instruction := solana.NewInstruction(
		MplTokenMetadataID,
		ix.Accounts(),
		data,
	)
	return *instruction, nil
}
