package sequences

import (
	"fmt"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/operations/router"
	tokenpoolops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/operations/token_pools"
	tokensops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/operations/tokens"
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
	return cldf_ops.NewSequence(
		"solana-adapter:configure-token-for-transfers",
		common_utils.Version_1_6_0,
		"Configure a token for cross-chain transfers across multiple chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.SolanaChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			b.Logger.Info("SVM Registering token:", input)

			routerAddr, err := a.GetRouterAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get router address: %w", err)
			}

			tokenAddr, tokenProgramId, err := getTokenMintAndTokenProgram(input.ExistingDataStore, input.TokenRef, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			tpAddr := solana.MustPublicKeyFromBase58(input.TokenPoolAddress)

			// if no token admin provided, ccip admin becomes the admin
			var tokenAdmin solana.PublicKey
			if input.ExternalAdmin != "" {
				tokenAdmin = solana.MustPublicKeyFromBase58(input.ExternalAdmin)
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

			b.Logger.Info("SVM Setting token pool:", input)

			fqAddr, err := a.GetFQAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee quoter address: %w", err)
			}

			spOut, err := operations.ExecuteOperation(b, routerops.SetPool, chains.SolanaChains()[chain.Selector], routerops.PoolParams{
				Router:         solana.PublicKeyFromBytes(routerAddr),
				FeeQuoter:      solana.PublicKeyFromBytes(fqAddr),
				TokenMint:      solana.MustPublicKeyFromBase58(tokenAddr.Address),
				TokenPool:      tpAddr,
				TokenProgramID: tokenProgramId,
				TokenPoolType:  input.PoolType,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set token pool: %w", err)
			}
			result.Addresses = append(result.Addresses, spOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, spOut.Output.BatchOps...)

			tokenMint := solana.MustPublicKeyFromBase58(tokenAddr.Address)
			for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
				op := tokenpoolops.UpsertRemoteChainConfigBurnMint
				tprl := tokenpoolops.UpsertRateLimitsBurnMint
				switch input.PoolType {
				case common_utils.BurnMintTokenPool.String():
					op = tokenpoolops.UpsertRemoteChainConfigBurnMint
					tprl = tokenpoolops.UpsertRateLimitsBurnMint
				case common_utils.LockReleaseTokenPool.String():
					op = tokenpoolops.UpsertRemoteChainConfigLockRelease
					tprl = tokenpoolops.UpsertRateLimitsLockRelease
				default:
					return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", input.PoolType)
				}
				upsertOut, err := operations.ExecuteOperation(b, op, chains.SolanaChains()[chain.Selector],
					tokenpoolops.RemoteChainConfig{
						TokenPool:          tpAddr,
						TokenMint:          tokenMint,
						TokenProgramID:     tokenProgramId,
						RemoteSelector:     remoteChainSelector,
						RemoteTokenAddress: remoteChainConfig.RemoteToken,
						RemoteDecimals:     remoteChainConfig.RemoteDecimals,
						RemotePoolAddress:  remoteChainConfig.RemotePool,
					})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to upsert remote chain config for token pool: %w", err)
				}
				localDecimals, err := utils.GetTokenDecimals(chain, tokenMint)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals for token on chain with selector %d: %w", chain.Selector, err)
				}
				obRL, ibRL := tokenapi.GenerateTPRLConfigs(
					remoteChainConfig.OutboundRateLimiterConfig,
					remoteChainConfig.InboundRateLimiterConfig,
					localDecimals,
					remoteChainConfig.RemoteDecimals,
					chain.Family(),
					common_utils.Version_1_6_0,
				)
				result.Addresses = append(result.Addresses, upsertOut.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, upsertOut.Output.BatchOps...)
				rateLimitOut, err := operations.ExecuteOperation(b, tprl, chains.SolanaChains()[chain.Selector],
					tokenpoolops.RemoteChainConfig{
						TokenPool:                 tpAddr,
						TokenMint:                 tokenMint,
						TokenProgramID:            tokenProgramId,
						RemoteSelector:            remoteChainSelector,
						InboundRateLimiterConfig:  ibRL,
						OutboundRateLimiterConfig: obRL,
					})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limits for token pool: %w", err)
				}
				result.Addresses = append(result.Addresses, rateLimitOut.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, rateLimitOut.Output.BatchOps...)
			}
			return result, nil
		})
}

func (a *SolanaAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return solana.MustPublicKeyFromBase58(ref.Address).Bytes(), nil
}

func (a *SolanaAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	return nil, fmt.Errorf("DeriveTokenAddress not implemented for Solana")
}

func (a *SolanaAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}
	return utils.GetTokenDecimals(chain, solana.PublicKeyFromBytes(token))
}

func (a *SolanaAdapter) DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
	decodedTokenPool := solana.PublicKeyFromBytes(tokenPool)
	decodedToken := solana.PublicKeyFromBytes(token)
	// For Solana, the token pool counterpart is derived using the token mint address and the token pool address as seeds.
	pool, err := tokens.TokenPoolConfigAddress(decodedToken, decodedTokenPool)
	if err != nil {
		return nil, fmt.Errorf("failed to derive token pool counterpart: %w", err)
	}
	return pool.Bytes(), nil
}

// ManualRegistration in Solana registers a token admin registry for a given token and initializes the token pool in CLL Token Pool Program.
// This changeset is used when the owner of the token pool doesn't have the mint authority over the token, but they want to self serve.
// So, this changeset includes the minimum configuration that CCIP Admin needs to do in the Token Admin Registry and in the Token Pool Program
func (a *SolanaAdapter) ManualRegistration() *cldf_ops.Sequence[tokenapi.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"svm-adapter:manual-registration",
		common_utils.Version_1_6_0,
		"Manually register a token and token pool on Solana Chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.ManualRegistrationSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.SolanaChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			tokenRef, tokenProgramId, err := getTokenMintAndTokenProgram(input.ExistingDataStore, input.TokenRef, chain)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token and token program address using the specified reference (%+v): %w", input.TokenRef, err)
			}
			routerAddr, err := a.GetRouterAddress(input.ExistingDataStore, chain.Selector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get router address: %w", err)
			}

			////////////////////////////
			/// Token Admin Registry ///
			////////////////////////////

			// if no token admin provided, ccip admin becomes the admin
			var tokenAdmin solana.PublicKey
			if input.ProposedOwner != "" {
				tokenAdmin = solana.MustPublicKeyFromBase58(input.ProposedOwner)
			}

			rtarOut, err := operations.ExecuteOperation(b, routerops.RegisterTokenAdminRegistry, chains.SolanaChains()[chain.Selector], routerops.TokenAdminRegistryParams{
				Router:            solana.PublicKeyFromBytes(routerAddr),
				TokenMint:         solana.MustPublicKeyFromBase58(tokenRef.Address),
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

			tokenPoolRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, input.TokenPoolRef, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address using the specified reference (%+v): %w", input.TokenPoolRef, err)
			}
			tokenMint := solana.MustPublicKeyFromBase58(tokenRef.Address)
			tokenPool := solana.MustPublicKeyFromBase58(tokenPoolRef.Address)

			if input.SVMExtraArgs == nil || !input.SVMExtraArgs.SkipTokenPoolInit {
				initTPOp := tokenpoolops.InitializeBurnMint
				transferOwnershipTPOp := tokenpoolops.TransferOwnershipBurnMint
				switch tokenPoolRef.Type.String() {
				case common_utils.BurnMintTokenPool.String():
					// Already set to burn mint
				case common_utils.LockReleaseTokenPool.String():
					initTPOp = tokenpoolops.InitializeLockRelease
					transferOwnershipTPOp = tokenpoolops.TransferOwnershipLockRelease
				default:
					return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", tokenPoolRef.Type)
				}
				rmnRemoteAddr, err := a.GetRMNRemoteAddress(input.ExistingDataStore, chain.Selector)
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
					return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize token pool: %w", err)
				}
				result.Addresses = append(result.Addresses, initTPOut.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, initTPOut.Output.BatchOps...)

				transferOwnershipOut, err := operations.ExecuteOperation(b, transferOwnershipTPOp, chains.SolanaChains()[chain.Selector], tokenpoolops.TokenPoolTransferOwnershipInput{
					Program:   tokenPool,
					NewOwner:  solana.MustPublicKeyFromBase58(input.ProposedOwner),
					TokenMint: tokenMint,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership: %w", err)
				}
				result.Addresses = append(result.Addresses, transferOwnershipOut.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, transferOwnershipOut.Output.BatchOps...)
			}
			/////////////////////////////
			/// Create Token Multisig ///
			/////////////////////////////

			if input.SVMExtraArgs != nil && len(input.SVMExtraArgs.CustomerMintAuthorities) > 0 {
				// The multisig will be used as the mint authority (or owner) depending on the pool flow.
				// We include the TokenPoolSigner PDA as one of the multisig signers so the Token Pool Program
				// can "sign" via PDA seeds when it needs to act (PDA signing).

				poolSigner, _ := tokens.TokenPoolSignerAddress(tokenMint, tokenPool)
				signers := make([]solana.PublicKey, 0, 1+len(input.SVMExtraArgs.CustomerMintAuthorities))
				signers = append(signers, poolSigner)

				for _, s := range input.SVMExtraArgs.CustomerMintAuthorities {
					if !s.IsZero() {
						signers = append(signers, s)
					}
				}

				createTokenMultisigOutput, err := operations.ExecuteOperation(b, tokensops.CreateTokenMultisig, chains.SolanaChains()[chain.Selector], tokensops.TokenMultisigParams{
					TokenProgram: tokenProgramId,
					Signers:      signers,
					TokenMint:    tokenMint,
					TokenSymbol:  tokenRef.Qualifier,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create token multisig on-chain: %w", err)
				}
				result.Addresses = append(result.Addresses, createTokenMultisigOutput.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, createTokenMultisigOutput.Output.BatchOps...)
			}

			return result, nil
		})
}

func (a *SolanaAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokenapi.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenPoolRateLimits",
		common_utils.Version_1_6_0,
		"Sets rate limits for a token pool on Solana",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input tokenapi.TPRLRemotes) (sequences.OnChainOutput, error) {
			chain := chains.SolanaChains()[input.ChainSelector]

			op := tokenpoolops.UpsertRateLimitsBurnMint
			switch input.TokenPoolRef.Type.String() {
			case common_utils.BurnMintTokenPool.String():
				op = tokenpoolops.UpsertRateLimitsBurnMint
			case common_utils.LockReleaseTokenPool.String():
				op = tokenpoolops.UpsertRateLimitsLockRelease
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", input.TokenPoolRef.Type.String())
			}
			tokenAddr, tokenProgramId, err := getTokenMintAndTokenProgram(input.ExistingDataStore, input.TokenRef, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			tokenMint := solana.MustPublicKeyFromBase58(tokenAddr.Address)
			tokenPool := solana.MustPublicKeyFromBase58(input.TokenPoolRef.Address)
			rateLimitOut, err := operations.ExecuteOperation(b, op, chains.SolanaChains()[chain.Selector],
				tokenpoolops.RemoteChainConfig{
					TokenPool:                 tokenPool,
					TokenMint:                 tokenMint,
					TokenProgramID:            tokenProgramId,
					RemoteSelector:            input.RemoteChainSelector,
					InboundRateLimiterConfig:  input.InboundRateLimiterConfig,
					OutboundRateLimiterConfig: input.OutboundRateLimiterConfig,
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limits for token pool: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: rateLimitOut.Output.BatchOps,
			}, nil
		},
	)
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
			needsTokenDeploy := true
			var rawTokenAddr string

			tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Type:          datastore.ContractType(input.Type),
				Qualifier:     input.Symbol,
			}, chain.Selector, datastore_utils.FullRef)
			if err == nil {
				b.Logger.Info("Token already deployed at address:", tokenAddr.Address)
				needsTokenDeploy = false
				rawTokenAddr = tokenAddr.Address
			}
			if needsTokenDeploy {
				var privateKey solana.PrivateKey
				if input.TokenPrivKey != "" {
					privateKey = solana.MustPrivateKeyFromBase58(input.TokenPrivKey)
				}
				ataList := []solana.PublicKey{}
				for _, sender := range input.Senders {
					ataList = append(ataList, solana.MustPublicKeyFromBase58(sender))
				}
				var premint uint64 = 0
				if input.PreMint != nil {
					premint = input.PreMint.Uint64()
				}
				deployOut, err := operations.ExecuteOperation(b, tokensops.DeploySolanaToken, chains.SolanaChains()[chain.Selector], tokensops.Params{
					ExistingAddresses:      input.ExistingDataStore.Addresses().Filter(),
					TokenProgramName:       input.Type,
					TokenPrivKey:           privateKey,
					TokenSymbol:            input.Symbol,
					ATAList:                ataList,
					PreMint:                premint,
					DisableFreezeAuthority: input.DisableFreezeAuthority,
					TokenDecimals:          input.Decimals,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token: %w", err)
				}
				result.Addresses = append(result.Addresses, deployOut.Output)
				rawTokenAddr = deployOut.Output.Address
			}
			// irrespective of whether the token was just deployed or already existed, we attempt to upload metadata if it was provided, since the metadata might not have been uploaded in a previous deployment
			if input.TokenMetadata != nil {
				input.TokenMetadata.TokenPubkey = rawTokenAddr
				_, err = operations.ExecuteOperation(b, tokensops.UpsertTokenMetadata, chains.SolanaChains()[chain.Selector], tokensops.TokenMetadataInput{
					ExistingAddresses: input.ExistingDataStore.Addresses().Filter(),
					Metadata:          *input.TokenMetadata,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to upload token metadata: %w", err)
				}
			}
			return result, nil
		},
	)
}

func (a *SolanaAdapter) DeployTokenVerify(e deployment.Environment, in tokenapi.DeployTokenInput) error {
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
			tokenAddr, tokenProgramId, err := getTokenMintAndTokenProgram(input.ExistingDataStore, *input.TokenRef, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			tokenPoolAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.TokenPoolQualifier,
				Type:          datastore.ContractType(input.PoolType),
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for symbol '%s' and qualifier '%s': %w", input.TokenRef.Qualifier, input.TokenPoolQualifier, err)
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

func getTokenMintAndTokenProgram(store datastore.DataStore, tokenRef datastore.AddressRef, chain cldf_solana.Chain) (datastore.AddressRef, solana.PublicKey, error) {
	ref := datastore.AddressRef{
		ChainSelector: chain.Selector,
	}
	if tokenRef.Address != "" {
		ref.Address = tokenRef.Address
	}
	if tokenRef.Qualifier != "" {
		ref.Qualifier = tokenRef.Qualifier
	}
	tokenAddr, err := datastore_utils.FindAndFormatRef(store, ref, chain.Selector, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, solana.PublicKey{}, fmt.Errorf("failed to find token address for ref '%+v': %w", ref, err)
	}
	tokenProgramId, err := utils.GetTokenProgramID(deployment.ContractType(tokenAddr.Type))
	if err != nil {
		return datastore.AddressRef{}, solana.PublicKey{}, fmt.Errorf("failed to get token program ID for token type '%s': %w", tokenAddr.Type, err)
	}
	return tokenAddr, tokenProgramId, nil
}

func (a *SolanaAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokenapi.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	return cldf_ops.NewSequence(
		"svm-adapter:update-authorities",
		common_utils.Version_1_6_0,
		"Update authorities for a token and token pool on Solana Chain",
		func(b cldf_ops.Bundle, e *deployment.Environment, input tokenapi.UpdateAuthoritiesInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := e.BlockChains.SolanaChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			ds := e.DataStore
			tokenRef, _, err := getTokenMintAndTokenProgram(ds, input.TokenRef, chain)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token and token program address using the specified reference (%+v): %w", input.TokenRef, err)
			}

			timelockSigner := utils.GetTimelockSignerPDA(
				ds.Addresses().Filter(),
				chain.Selector,
				common_utils.CLLQualifier,
			)

			tokenPoolRef, err := datastore_utils.FindAndFormatRef(ds, input.TokenPoolRef, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address using the specified reference (%+v): %w", input.TokenPoolRef, err)
			}
			tokenMint := solana.MustPublicKeyFromBase58(tokenRef.Address)
			tokenPool := solana.MustPublicKeyFromBase58(tokenPoolRef.Address)

			transferOwnershipTPOp := tokenpoolops.TransferOwnershipBurnMint
			acceptOwnershipTPOp := tokenpoolops.AcceptOwnershipBurnMint
			tprlAdminTPOp := tokenpoolops.UpdateRateLimitAdminBurnMint
			switch tokenPoolRef.Type.String() {
			case common_utils.BurnMintTokenPool.String():
				// Already set to burn mint
			case common_utils.LockReleaseTokenPool.String():
				transferOwnershipTPOp = tokenpoolops.TransferOwnershipLockRelease
				acceptOwnershipTPOp = tokenpoolops.AcceptOwnershipLockRelease
				tprlAdminTPOp = tokenpoolops.UpdateRateLimitAdminLockRelease
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", tokenPoolRef.Type)
			}

			tprlAdminOut, err := operations.ExecuteOperation(b, tprlAdminTPOp, e.BlockChains.SolanaChains()[chain.Selector], tokenpoolops.TokenPoolTransferOwnershipInput{
				Program:   tokenPool,
				TokenMint: tokenMint,
				NewOwner:  timelockSigner,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update TPRL admin: %w", err)
			}
			result.Addresses = append(result.Addresses, tprlAdminOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, tprlAdminOut.Output.BatchOps...)

			b.Logger.Infof("Transferring ownership of token pool %s to timelock signer %s", tokenPool.String(), timelockSigner.String())
			transferOwnershipOut, err := operations.ExecuteOperation(b, transferOwnershipTPOp, e.BlockChains.SolanaChains()[chain.Selector], tokenpoolops.TokenPoolTransferOwnershipInput{
				Program:   tokenPool,
				NewOwner:  timelockSigner,
				TokenMint: tokenMint,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership: %w", err)
			}
			result.Addresses = append(result.Addresses, transferOwnershipOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, transferOwnershipOut.Output.BatchOps...)

			b.Logger.Infof("Accepting ownership of token pool %s by timelock signer %s", tokenPool.String(), timelockSigner.String())
			acceptOwnershipOut, err := operations.ExecuteOperation(b, acceptOwnershipTPOp, e.BlockChains.SolanaChains()[chain.Selector], tokenpoolops.TokenPoolTransferOwnershipInput{
				Program:   tokenPool,
				NewOwner:  timelockSigner,
				TokenMint: tokenMint,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership: %w", err)
			}
			result.Addresses = append(result.Addresses, acceptOwnershipOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, acceptOwnershipOut.Output.BatchOps...)

			return result, nil
		})
}
