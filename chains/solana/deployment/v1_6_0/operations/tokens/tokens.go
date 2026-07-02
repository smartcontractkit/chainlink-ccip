package tokens

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

var (
	LinkContractType cldf_deployment.ContractType = "LinkToken"
	Version          *semver.Version              = semver.MustParse("1.6.0")
)

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
			b.GetContext(),
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
		instructions, err := tokens.CreateTokenWith(
			b.GetContext(),
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
			createATAIx, _, err := tokens.CreateAssociatedTokenAccount(
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
				ata, _, _ := tokens.FindAssociatedTokenAddress(tokenProgramID, mint, sender)
				mintToI, err := tokens.MintTo(input.PreMint, tokenProgramID, mint, ata, chain.DeployerKey.PublicKey())
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

type TokenMultisigParams struct {
	Signers      []solana.PublicKey
	TokenProgram solana.PublicKey
	TokenMint    solana.PublicKey
	TokenSymbol  string
}

var CreateTokenMultisig = operations.NewOperation(
	"solana-token:create-multisig",
	Version,
	"Creates a Token Multisig account for the given token program",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenMultisigParams) (sequences.OnChainOutput, error) {
		ctx := b.GetContext()

		// m is always 1
		const m uint8 = 1

		// --- Validate inputs ---
		if input.TokenProgram.IsZero() {
			return sequences.OnChainOutput{}, fmt.Errorf("token program is zero")
		}
		if !input.TokenProgram.Equals(solana.Token2022ProgramID) && !input.TokenProgram.Equals(solana.TokenProgramID) {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token program: %s", input.TokenProgram.String())
		}
		if len(input.Signers) < 2 {
			return sequences.OnChainOutput{}, fmt.Errorf("signers length must be at least 2, got %d", len(input.Signers))
		}
		if len(input.Signers) > 11 {
			return sequences.OnChainOutput{}, fmt.Errorf("too many signers: %d > %d", len(input.Signers), 11)
		}

		// --- Create multisig keypair (must sign CreateAccount) ---
		multisig, err := solana.NewRandomPrivateKey()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to generate multisig keypair: %w", err)
		}

		// --- Instructions ---
		// get stake amount for init
		lamports, err := chain.Client.GetMinimumBalanceForRentExemption(ctx, tokens.MultisigSize, rpc.CommitmentConfirmed)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		ixs, err := tokens.IxsInitTokenMultisig(input.TokenProgram, lamports, chain.DeployerKey.PublicKey(), multisig.PublicKey(), m, input.Signers)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create multisig instructions: %w", err)
		}
		err = chain.Confirm(ixs, common.AddSigners(multisig))
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm create multisig transaction: %w", err)
		}
		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{
				{
					ChainSelector: chain.Selector,
					Address:       multisig.PublicKey().String(),
					Type:          "TOKEN_MULTISIG",
					Version:       Version,
					Qualifier:     input.TokenSymbol,
					Labels:        datastore.NewLabelSet(input.TokenMint.String()),
				},
			},
		}, nil
	})

type ExtendTokenPoolLookupTableParams struct {
	// AllowDuplicates indicates whether to allow duplicate accounts in the lookup
	// table. When false (the default), then any duplicate account in the Accounts
	// slice will be ignored AND any account that is already present in the lookup
	// table will be ignored. If this is true, then *all* accounts in the Accounts
	// slice will be added to the lookup table, even if they are already present.
	AllowDuplicates bool
	TokenMint       solana.PublicKey
	Router          solana.PublicKey
	Accounts        []solana.PublicKey
}

var ExtendTokenPoolLookupTable = operations.NewOperation(
	"solana-token:extend-token-pool-lookup-table",
	Version,
	"Extends the token pool address lookup table with additional accounts",
	func(b operations.Bundle, chain cldf_solana.Chain, input ExtendTokenPoolLookupTableParams) (sequences.OnChainOutput, error) {
		ctx := b.GetContext()

		if input.TokenMint.IsZero() {
			return sequences.OnChainOutput{}, errors.New("token mint is zero")
		}
		if input.Router.IsZero() {
			return sequences.OnChainOutput{}, errors.New("router is zero")
		}
		if len(input.Accounts) == 0 {
			b.Logger.Info("no accounts provided - skipping token pool lookup table extension")
			return sequences.OnChainOutput{}, nil
		}
		if chain.DeployerKey == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("solana deployer key is nil for chain selector '%d'", chain.Selector)
		}

		tokenAdminRegistryPDA, _, err := state.FindTokenAdminRegistryPDA(input.TokenMint, input.Router)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find token admin registry PDA: %w", err)
		}

		var tokenAdminRegistry ccip_common.TokenAdminRegistry
		if err := chain.GetAccountDataBorshInto(ctx, tokenAdminRegistryPDA, &tokenAdminRegistry); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to parse token admin registry account at '%s': %w", tokenAdminRegistryPDA.String(), err)
		}

		lookupTable := tokenAdminRegistry.LookupTable
		if lookupTable.IsZero() {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"token pool lookup table is not set for token '%s' on router '%s'; run configure-for-transfers / set-pool first",
				input.TokenMint.String(), input.Router.String(),
			)
		}

		lutState, err := common.GetAddressLookupTableState(ctx, chain.Client, lookupTable)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to parse addresses from lookup table state at %q: %w", lookupTable.String(), err)
		}
		if !lutState.IsActive() {
			return sequences.OnChainOutput{}, fmt.Errorf("token pool lookup table at %q is not active", lookupTable.String())
		}
		if lutState.Authority == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("token pool lookup table at %q has no authority set", lookupTable.String())
		}
		if !lutState.Authority.Equals(chain.DeployerKey.PublicKey()) {
			return sequences.OnChainOutput{}, fmt.Errorf("expected token pool lookup table at %q to have authority %q (deployer key), but it is %q", lookupTable.String(), chain.DeployerKey.PublicKey().String(), lutState.Authority.String())
		}

		seen := make(map[solana.PublicKey]bool, len(lutState.Addresses))
		if !input.AllowDuplicates {
			for _, entry := range lutState.Addresses {
				seen[entry] = true
			}
		}

		toAdd := make(solana.PublicKeySlice, 0, len(input.Accounts))
		for i, acct := range input.Accounts {
			if acct.IsZero() {
				return sequences.OnChainOutput{}, fmt.Errorf("account at index %d is zero", i)
			}
			if !seen[acct] {
				toAdd.Append(acct)
				if !input.AllowDuplicates {
					seen[acct] = true
				}
			}
		}

		if len(toAdd) == 0 {
			b.Logger.Infof(
				"all %d account(s) already present in token pool lookup table at '%s' - nothing to extend",
				len(input.Accounts), lookupTable.String(),
			)
			return sequences.OnChainOutput{}, nil
		}

		b.Logger.Infof(
			"extending token pool lookup table at '%s' with account(s): %s",
			lookupTable.String(), strings.Join(toAdd.ToBase58(), ", "),
		)

		if err := common.ExtendLookupTable(ctx, chain.Client, lookupTable, *chain.DeployerKey, toAdd); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to extend token pool lookup table at '%s': %w", lookupTable.String(), err)
		}
		if err := common.AwaitSlotChange(ctx, chain.Client); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to await slot change after extending token pool lookup table: %w", err)
		}

		return sequences.OnChainOutput{}, nil
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
		mintMetadata, metadataPDA, err := tokens.GetTokenMetadata(b.GetContext(), chain.Client, tokenMint)
		if err != nil {
			b.Logger.Infof("error getting metadata account data, skipping update for %s: %v", tokenMint.String(), err)
			return sequences.OnChainOutput{}, nil
		}
		b.Logger.Infof("Metadata %s", metadataPDA)
		newData := tokens.GetTokenDataV2(mintMetadata)
		newUpdateAuthority := mintMetadata.UpdateAuthority
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
				tokens.MplTokenMetadataID.String(),
				tokens.MplTokenMetadataProgramName,
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

// discriminator for update_metadata_account_v2 ix
const UpdateMetadataAccountV2Ix = 15

// discriminator for create_metadata_account
const CreateMetadataAccountV2Ix = 16

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
		tokens.MplTokenMetadataID,
		ix.Accounts(),
		data,
	)
	return *instruction, nil
}
