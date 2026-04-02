package adapters

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	v1_6_0_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ tokens.TokenAdapter = &TokenAdapter{}

// TokenAdapter is the adapter for EVM tokens using 1.6.1 token pools.
type TokenAdapter struct{}

// ConfigureTokenForTransfersSequence returns the sequence for configuring an EVM token with a 1.6.1 token pool.
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return evm_seq.ConfigureTokenForTransfers
}

// AddressRefToBytes returns an EVM address reference as an EVM address.
func (t *TokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}

// DeriveTokenAddress derives the token address from a token pool reference, returning it as an EVM address.
func (t *TokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	getTokenReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(poolRef.Address),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", poolRef.Address, chain, err)
	}

	return t.AddressRefToBytes(datastore.AddressRef{
		Address: getTokenReport.Output.Hex(),
	})
}

func (t *TokenAdapter) DeployToken() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return v1_6_0_seq.DeployToken
}

func (t *TokenAdapter) DeployTokenVerify(e deployment.Environment, input tokens.DeployTokenInput) error {
	tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
		ChainSelector: input.ChainSelector,
		Type:          datastore.ContractType(input.Type),
		Qualifier:     input.Symbol,
	}, input.ChainSelector, datastore_utils.FullRef)
	if err == nil {
		e.OperationsBundle.Logger.Info("Token already deployed at address:", tokenAddr.Address)
		return nil
	}

	if err := utils.ValidateEVMAddress(input.CCIPAdmin, "CCIPAdmin"); err != nil {
		return err
	}
	if err := utils.ValidateEVMAddress(input.ExternalAdmin, "ExternalAdmin"); err != nil {
		return err
	}

	if input.Decimals > 18 {
		return fmt.Errorf("EVM tokens cannot have more than 18 decimals, got %d", input.Decimals)
	}

	if input.PreMint != nil && input.Supply != nil && *input.Supply != 0 && *input.PreMint > *input.Supply {
		return fmt.Errorf("pre-mint amount cannot be greater than max supply, got pre-mint %d and supply %d", *input.PreMint, *input.Supply)
	}

	return nil
}

func (t *TokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-1.6.1-adapter:deploy-token-pool-for-token",
		cciputils.Version_1_6_1,
		"Deploy a token pool for a token on an EVM chain using the 1.6.0 DeployTokenPool sequence",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			out, err := cldf_ops.ExecuteSequence(b, v1_6_0_seq.DeployTokenPool, chains, input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token pool on chain %d: %w", input.ChainSelector, err)
			}

			var result sequences.OnChainOutput
			result.Addresses = append(result.Addresses, out.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)

			toknFilterDS := datastore.AddressRef{ChainSelector: input.ChainSelector}
			if input.TokenRef.Address != "" {
				toknFilterDS.Address = input.TokenRef.Address
			}
			if input.TokenRef.Qualifier != "" {
				toknFilterDS.Qualifier = input.TokenRef.Qualifier
			}
			if input.TokenRef.Type != "" {
				toknFilterDS.Type = input.TokenRef.Type
			}
			toknRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, toknFilterDS, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token address for symbol %q on chain %d: %w", input.TokenRef.Qualifier, input.ChainSelector, err)
			}

			isToknTypeBnM := toknRef.Type.String() == bnmERC20ops.ContractType.String()
			isPoolTypeBnM := input.PoolType == cciputils.BurnMintTokenPool.String()
			if isPoolTypeBnM && isToknTypeBnM && len(out.Output.Addresses) >= 1 {
				poolRef := out.Output.Addresses[0]

				poolAddrBytes, err := t.AddressRefToBytes(poolRef)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to convert deployed token pool address ref to bytes: %w", err)
				}

				toknAddrBytes, err := t.AddressRefToBytes(toknRef)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token address ref to bytes: %w", err)
				}

				poolAddr := common.BytesToAddress(poolAddrBytes)
				if poolAddr == (common.Address{}) {
					return sequences.OnChainOutput{}, errors.New("deployed token pool address is zero address")
				}

				toknAddr := common.BytesToAddress(toknAddrBytes)
				if toknAddr == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("token address for symbol %q is zero address", input.TokenRef.Qualifier)
				}

				chain, ok := chains.EVMChains()[input.ChainSelector]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
				}

				report, err := cldf_ops.ExecuteOperation(b,
					bnmERC20ops.GrantMintAndBurnRoles,
					chain,
					evm_contract.FunctionInput[common.Address]{
						ChainSelector: input.ChainSelector,
						Address:       toknAddr,
						Args:          poolAddr,
					},
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenRef.Qualifier, input.ChainSelector, err)
				}

				batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{report.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for granting mint and burn roles to token pool %q for token %q on chain %d: %w", poolAddr.Hex(), input.TokenRef.Qualifier, input.ChainSelector, err)
				}

				result.BatchOps = append(result.BatchOps, batchOp)
			}

			return result, nil
		},
	)
}

func (t *TokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokens.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) SetPool() *cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	getTokenDecimalsReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetTokenDecimals, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(poolRef.Address),
	})
	if err != nil {
		return 0, fmt.Errorf("failed to get token decimals from token pool with address %s on %s: %w", poolRef.Address, chain, err)
	}
	return getTokenDecimalsReport.Output, nil
}

func (t *TokenAdapter) DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
	return tokenPool, nil
}

func (t *TokenAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokens.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

