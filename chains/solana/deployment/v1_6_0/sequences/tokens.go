package sequences

import (
	"fmt"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	tokenpoolops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/token_pools"
	tokensops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *SolanaAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokenapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO: implement me
	return nil
}

func (a *SolanaAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	// TODO: implement me
	return nil, nil
}

func (a *SolanaAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	// TODO: implement me
	return nil, nil
}

// ManualRegistration in Solana registers a token admin registry for a given token and initializes the token pool in CLL Token Pool Program.
// This changeset is used when the owner of the token pool doesn't have the mint authority over the token, but they want to self serve.
// So, this changeset includes the minimum configuration that CCIP Admin needs to do in the Token Admin Registry and in the Token Pool Program
func (a *SolanaAdapter) ManualRegistration() *cldf_ops.Sequence[tokenapi.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"svm-adapter:manual-registration",
		common_utils.Version_1_6_0,
		"Manually register a token and token pool on Solana Chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.ManualRegistrationInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			store := datastore.NewMemoryDataStore()
			for _, addr := range input.ExistingAddresses {
				if err := store.AddressRefStore.Upsert(addr); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to upsert address %v: %w", addr, err)
				}
			}
			chain, ok := chains.SolanaChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			tokenAddr, tokenProgramId, err := getTokenMintAndTokenProgram(input.ExistingDataStore, input.RegisterTokenConfigs.TokenSymbol, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			routerAddr, err := a.GetRouterAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get router address: %w", err)
			}

			////////////////////////////
			/// Token Admin Registry ///
			////////////////////////////

			// if no token admin provided, ccip admin becomes the admin
			var tokenAdmin solana.PublicKey
			if input.RegisterTokenConfigs.ProposedOwner != "" {
				tokenAdmin = solana.MustPublicKeyFromBase58(input.RegisterTokenConfigs.ProposedOwner)
			}

			rtarOut, err := operations.ExecuteOperation(b, routerops.RegisterTokenAdminRegistry, chains.SolanaChains()[chain.Selector], routerops.TokenAdminRegistryParams{
				Router:            solana.PublicKeyFromBytes(routerAddr),
				TokenMint:         solana.MustPublicKeyFromBase58(tokenAddr.Address),
				Admin:             tokenAdmin,
				ExistingAddresses: input.ExistingDataStore.Addresses().Filter(),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token metadata: %w", err)
			}
			result.Addresses = append(result.Addresses, rtarOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, rtarOut.Output.BatchOps...)

			/////////////////////////////
			/// Initialize Token Pool ///
			/////////////////////////////

			tokenPoolAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.RegisterTokenConfigs.TokenPoolQualifier,
				Type:          datastore.ContractType(input.RegisterTokenConfigs.PoolType),
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for symbol '%s' and qualifier '%s': %w", input.RegisterTokenConfigs.TokenSymbol, input.RegisterTokenConfigs.TokenPoolQualifier, err)
			}
			tokenMint := solana.MustPublicKeyFromBase58(tokenAddr.Address)
			tokenPool := solana.MustPublicKeyFromBase58(tokenPoolAddr.Address)

			initTPOp := tokenpoolops.InitializeBurnMint
			transferOwnershipTPOp := tokenpoolops.TransferOwnershipBurnMint
			authority := tokenpoolops.GetAuthorityBurnMint(chain, tokenPool, tokenMint)
			switch input.RegisterTokenConfigs.PoolType {
			case common_utils.BurnMintTokenPool.String():
				// Already set to burn mint
			case common_utils.LockReleaseTokenPool.String():
				initTPOp = tokenpoolops.InitializeLockRelease
				transferOwnershipTPOp = tokenpoolops.TransferOwnershipLockRelease
				authority = tokenpoolops.GetAuthorityLockRelease(chain, tokenPool, tokenMint)
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", input.RegisterTokenConfigs.PoolType)
			}

			rmnRemoteAddr, err := a.GetRMNRemoteAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get RMN remote address: %w", err)
			}

			initTPOut, err := operations.ExecuteOperation(b, initTPOp, chains.SolanaChains()[chain.Selector], tokenpoolops.Params{
				TokenPool:      tokenPool,
				TokenMint:      tokenMint,
				TokenProgramID: tokenProgramId,
				Router:         solana.PublicKeyFromBytes(routerAddr),
				RMNRemote:      solana.PublicKeyFromBytes(rmnRemoteAddr),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token: %w", err)
			}
			result.Addresses = append(result.Addresses, initTPOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, initTPOut.Output.BatchOps...)

			transferOwnershipOut, err := operations.ExecuteOperation(b, transferOwnershipTPOp, chains.SolanaChains()[chain.Selector], tokenpoolops.TokenPoolTransferOwnershipInput{
				Program:      tokenPool,
				CurrentOwner: authority,
				NewOwner:     solana.MustPublicKeyFromBase58(input.RegisterTokenConfigs.ProposedOwner),
				TokenMint:    tokenMint,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token: %w", err)
			}
			result.Addresses = append(result.Addresses, transferOwnershipOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, transferOwnershipOut.Output.BatchOps...)

			/////////////////////////////
			/// Create Token Multisig ///
			/////////////////////////////

			// TODO

			return result, nil
		})
}

func (a *SolanaAdapter) DeployToken() *cldf_ops.Sequence[tokenapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"DeployToken",
		common_utils.Version_1_6_0,
		"Deploys a token contract on Solana",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			b.Logger.Info("SVM Deploying token:", input)
			chain := chains.SolanaChains()[input.ChainSelector]

			tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.Symbol,
			}, chain.Selector, datastore_utils.FullRef)
			if err == nil {
				b.Logger.Info("Token already deployed at address:", tokenAddr.Address)
				return result, nil
			}

			var privateKey solana.PrivateKey
			if input.TokenPrivKey != "" {
				privateKey = solana.MustPrivateKeyFromBase58(input.TokenPrivKey)
			}
			ataList := []solana.PublicKey{}
			for _, sender := range input.Senders {
				ataList = append(ataList, solana.MustPublicKeyFromBase58(sender))
			}

			deployOut, err := operations.ExecuteOperation(b, tokensops.DeploySolanaToken, chains.SolanaChains()[chain.Selector], tokensops.Params{
				ExistingAddresses:      input.ExistingDataStore.Addresses().Filter(),
				TokenProgramName:       input.Type,
				TokenPrivKey:           privateKey,
				TokenSymbol:            input.Symbol,
				ATAList:                ataList,
				DisableFreezeAuthority: input.DisableFreezeAuthority,
				TokenDecimals:          input.Decimals,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token: %w", err)
			}
			result.Addresses = append(result.Addresses, deployOut.Output)
			return result, nil
		},
	)
}

func (a *SolanaAdapter) DeployTokenVerify(e deployment.Environment, in any) error {
	return nil
}

func (a *SolanaAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokenapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"DeployTokenPoolForToken",
		common_utils.Version_1_6_0,
		"Configures a token pool for a given token on Solana. Doesn't actually deploy a new token pool contract, as Solana token pools are just accounts.",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			b.Logger.Info("SVM Deploying token:", input)
			chain := chains.SolanaChains()[input.ChainSelector]

			op := tokenpoolops.InitializeBurnMint
			switch input.PoolType {
			case common_utils.BurnMintTokenPool.String():
				op = tokenpoolops.InitializeBurnMint
			case common_utils.LockReleaseTokenPool.String():
				op = tokenpoolops.InitializeLockRelease
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", input.PoolType)
			}
			tokenAddr, tokenProgramId, err := getTokenMintAndTokenProgram(input.ExistingDataStore, input.TokenSymbol, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			tokenPoolAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.TokenPoolQualifier,
				Type:          datastore.ContractType(input.PoolType),
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for symbol '%s' and qualifier '%s': %w", input.TokenSymbol, input.TokenPoolQualifier, err)
			}
			routerAddr, err := a.GetRouterAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get router address: %w", err)
			}
			rmnRemoteAddr, err := a.GetRMNRemoteAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get RMN remote address: %w", err)
			}

			tokenMint := solana.MustPublicKeyFromBase58(tokenAddr.Address)
			tokenPool := solana.MustPublicKeyFromBase58(tokenPoolAddr.Address)

			deployOut, err := operations.ExecuteOperation(b, op, chains.SolanaChains()[chain.Selector], tokenpoolops.Params{
				TokenPool:      tokenPool,
				TokenMint:      tokenMint,
				TokenProgramID: tokenProgramId,
				Router:         solana.PublicKeyFromBytes(routerAddr),
				RMNRemote:      solana.PublicKeyFromBytes(rmnRemoteAddr),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token: %w", err)
			}
			result.Addresses = append(result.Addresses, deployOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, deployOut.Output.BatchOps...)

			poolSigner, _ := tokens.TokenPoolSignerAddress(tokenMint, tokenPool)

			// ATA for token pool
			tokenPoolATA, _, err := tokens.FindAssociatedTokenAddress(tokenProgramId, tokenMint, poolSigner)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find associated token address for token pool: %w", err)
			}

			// Check if the ATA is already initialized
			_, err = chain.Client.GetAccountInfo(b.GetContext(), tokenPoolATA)
			if err != nil { // it means that the ATA does not exist
				ixn, _, err := tokens.CreateAssociatedTokenAccount(
					tokenProgramId,
					tokenMint,
					poolSigner,
					chain.DeployerKey.PublicKey(),
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create the instruction to create associated token account for tokenpool (mint: %s, pool: %s): %w", tokenMint.String(), tokenPool.String(), err)
				}
				if err := chain.Confirm([]solana.Instruction{ixn}); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create associated token account for tokenpool (mint: %s, pool: %s): %w", tokenMint.String(), tokenPool.String(), err)
				}
			}

			mintAuthority := utils.GetTokenMintAuthority(chain, tokenMint)
			// make pool mint_authority for token, if necessary
			if input.PoolType == common_utils.BurnMintTokenPool.String() && tokenMint != solana.SolMint {
				if mintAuthority.String() == chain.DeployerKey.PublicKey().String() {
					ixn, err := tokens.SetTokenMintAuthority(
						tokenProgramId,
						poolSigner,
						tokenMint,
						chain.DeployerKey.PublicKey(),
					)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to generate instructions: %w", err)
					}
					if err := chain.Confirm([]solana.Instruction{ixn}); err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set mint authority: %w", err)
					}
					b.Logger.Infow("Setting mint authority", "poolSigner", poolSigner.String())
				} else {
					b.Logger.Warnw("Token's mint authority is not with deployer key, skipping setting poolSigner as mint authority",
						"poolType", input.PoolType, "mintAuthority", tokenMint,
						"deployer", chain.DeployerKey.PublicKey().String(), "poolSigner", poolSigner.String())
				}
			} else {
				b.Logger.Warnw("PoolType is not a BurnMintTokenPool, skipping setting poolSigner as mint authority",
					"poolType", input.PoolType, "mintAuthority", tokenMint,
					"deployer", chain.DeployerKey.PublicKey().String(), "poolSigner", poolSigner.String())
			}

			return result, nil
		},
	)
}

func getTokenMintAndTokenProgram(store datastore.DataStore, tokenSymbol string, chain cldf_solana.Chain) (datastore.AddressRef, solana.PublicKey, error) {
	tokenAddr, err := datastore_utils.FindAndFormatRef(store, datastore.AddressRef{
		ChainSelector: chain.Selector,
		Qualifier:     tokenSymbol,
	}, chain.Selector, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, solana.PublicKey{}, fmt.Errorf("failed to find token address for symbol '%s': %w", tokenSymbol, err)
	}
	tokenProgramId, err := utils.GetTokenProgramID(deployment.ContractType(tokenAddr.Type))
	if err != nil {
		return datastore.AddressRef{}, solana.PublicKey{}, fmt.Errorf("failed to get token program ID for token type '%s': %w", tokenAddr.Type, err)
	}
	return tokenAddr, tokenProgramId, nil
}

func (a *SolanaAdapter) RegisterToken() *cldf_ops.Sequence[tokenapi.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"RegisterToken",
		common_utils.Version_1_6_0,
		"Registers a token on Solana",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input tokenapi.RegisterTokenInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			b.Logger.Info("SVM Registering token:", input)
			chain := chains.SolanaChains()[input.ChainSelector]

			routerAddr, err := a.GetRouterAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get router address: %w", err)
			}

			tokenAddr, _, err := getTokenMintAndTokenProgram(input.ExistingDataStore, input.TokenSymbol, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			// if no token admin provided, ccip admin becomes the admin
			var tokenAdmin solana.PublicKey
			if input.TokenAdmin != "" {
				tokenAdmin = solana.MustPublicKeyFromBase58(input.TokenAdmin)
			}

			rtarOut, err := operations.ExecuteOperation(b, routerops.RegisterTokenAdminRegistry, chains.SolanaChains()[chain.Selector], routerops.TokenAdminRegistryParams{
				Router:            solana.PublicKeyFromBytes(routerAddr),
				TokenMint:         solana.MustPublicKeyFromBase58(tokenAddr.Address),
				Admin:             tokenAdmin,
				ExistingAddresses: input.ExistingDataStore.Addresses().Filter(),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token metadata: %w", err)
			}
			result.Addresses = append(result.Addresses, rtarOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, rtarOut.Output.BatchOps...)
			pendingSigner := rtarOut.Output.PendingSigner

			// accept if we can
			atarOut, err := operations.ExecuteOperation(b, routerops.AcceptTokenAdminRegistry, chains.SolanaChains()[chain.Selector], routerops.TokenAdminRegistryParams{
				Router:            solana.PublicKeyFromBytes(routerAddr),
				TokenMint:         solana.MustPublicKeyFromBase58(tokenAddr.Address),
				Admin:             pendingSigner,
				ExistingAddresses: input.ExistingDataStore.Addresses().Filter(),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token metadata: %w", err)
			}
			result.Addresses = append(result.Addresses, atarOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, atarOut.Output.BatchOps...)
			return result, nil
		},
	)
}

func (a *SolanaAdapter) SetPool() *cldf_ops.Sequence[tokenapi.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetPool",
		common_utils.Version_1_6_0,
		"Sets the token pool for a given token on Solana",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input tokenapi.SetPoolInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			b.Logger.Info("SVM Setting token pool:", input)
			chain := chains.SolanaChains()[input.ChainSelector]

			tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.TokenSymbol,
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token address for symbol '%s': %w", input.TokenSymbol, err)
			}
			tokenProgramId, err := utils.GetTokenProgramID(deployment.ContractType(tokenAddr.Type))
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token program ID for token type '%s': %w", tokenAddr.Type, err)
			}

			tokenPoolAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.TokenPoolQualifier,
				Type:          datastore.ContractType(input.PoolType),
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for symbol '%s' and qualifier '%s': %w", input.TokenSymbol, input.TokenPoolQualifier, err)
			}
			routerAddr, err := a.GetRouterAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get router address: %w", err)
			}

			fqAddr, err := a.GetFQAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee quoter address: %w", err)
			}

			spOut, err := operations.ExecuteOperation(b, routerops.SetPool, chains.SolanaChains()[chain.Selector], routerops.PoolParams{
				Router:         solana.PublicKeyFromBytes(routerAddr),
				FeeQuoter:      solana.PublicKeyFromBytes(fqAddr),
				TokenMint:      solana.MustPublicKeyFromBase58(tokenAddr.Address),
				TokenPool:      solana.MustPublicKeyFromBase58(tokenPoolAddr.Address),
				TokenProgramID: tokenProgramId,
				TokenPoolType:  input.PoolType,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set token pool: %w", err)
			}
			result.Addresses = append(result.Addresses, spOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, spOut.Output.BatchOps...)
			return result, nil
		},
	)
}

func (a *SolanaAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokenapi.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
