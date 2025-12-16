package sequences

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	tokenpoolops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/token_pools"
	tokensops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/tokens"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *SolanaAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokenapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *SolanaAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return nil, nil
}
func (a *SolanaAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	return nil, nil
}

func (a *SolanaAdapter) ManualRegistration() *cldf_ops.Sequence[tokenapi.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
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
			switch input.RegisterTokenConfig.PoolType {
			case common_utils.BurnMintTokenPool.String():
				op = tokenpoolops.InitializeBurnMint
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type '%s' for Solana", input.RegisterTokenConfig.PoolType)
			}

			tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     input.RegisterTokenConfig.TokenSymbol,
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token address for symbol '%s': %w", input.RegisterTokenConfig.TokenSymbol, err)
			}
			tokenProgramId, err := utils.GetTokenProgramID(deployment.ContractType(tokenAddr.Type))
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token program ID for token type '%s': %w", tokenAddr.Type, err)
			}

			tokenpoolQualifier := common_utils.CLLQualifier
			if input.RegisterTokenConfig.TokenPoolQualifier != "" {
				tokenpoolQualifier = input.RegisterTokenConfig.TokenPoolQualifier
			}
			tokenPoolAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: chain.Selector,
				Qualifier:     tokenpoolQualifier,
				Type:          datastore.ContractType(input.RegisterTokenConfig.PoolType),
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for symbol '%s' and qualifier '%s': %w", input.RegisterTokenConfig.TokenSymbol, input.RegisterTokenConfig.TokenPoolQualifier, err)
			}

			deployOut, err := operations.ExecuteOperation(b, op, chains.SolanaChains()[chain.Selector], tokenpoolops.Params{
				TokenPool:      solana.MustPublicKeyFromBase58(tokenPoolAddr.Address),
				TokenMint:      solana.MustPublicKeyFromBase58(tokenAddr.Address),
				TokenProgramID: tokenProgramId,
				// // SPLToken or SPLToken2022
				// TokenMint solana.PublicKey
				// // SPLToken or SPLToken2022
				// TokenProgramID solana.PublicKey
				// // Only used for certain ops
				// RMNRemote        solana.PublicKey
				// Router           solana.PublicKey
				// NewMintAuthority solana.PublicKey
				// OldMintAuthority solana.PublicKey
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token: %w", err)
			}
			result.Addresses = append(result.Addresses, deployOut.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, deployOut.Output.BatchOps...)
			return result, nil
		},
	)
}
func (a *SolanaAdapter) RegisterToken() *cldf_ops.Sequence[tokenapi.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
func (a *SolanaAdapter) SetPool() *cldf_ops.Sequence[tokenapi.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
func (a *SolanaAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokenapi.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
