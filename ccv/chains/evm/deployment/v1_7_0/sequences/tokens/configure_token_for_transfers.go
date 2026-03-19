package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ConfigureTokenForTransfers = cldf_ops.NewSequence(
	"configure-token-for-transfers",
	semver.MustParse("2.0.0"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (output sequences.OnChainOutput, err error) {
		ops := make([]mcms_types.BatchOperation, 0)
		evmChain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		var tokenAddress common.Address
		if input.TokenAddress != "" {
			tokenAddress = common.HexToAddress(input.TokenAddress)
		} else {
			tokenAddrReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, evmChain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.TokenPoolAddress),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
			}
			tokenAddress = tokenAddrReport.Output
		}
		tokenPoolAddress := common.HexToAddress(input.TokenPoolAddress)
		registryTokenPoolAddress := tokenPoolAddress
		if input.RegistryTokenPoolAddress != "" {
			registryTokenPoolAddress = common.HexToAddress(input.RegistryTokenPoolAddress)
		}

		// If a liquidity migration is requested, derive the old pool from the TAR and run the migration sequence.
		if input.LiquidityMigrationAmount != nil || input.LiquidityMigrationBasisPoints != nil {
			if input.TimelockAddress == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("TimelockAddress is required when liquidity migration is requested")
			}

			registryAddr := common.HexToAddress(input.RegistryAddress)
			tokenConfigReport, err := cldf_ops.ExecuteOperation(b, tar_ops.GetTokenConfig, evmChain, evm_contract.FunctionInput[common.Address]{
				ChainSelector: input.ChainSelector,
				Address:       registryAddr,
				Args:          tokenAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config from registry for token %s: %w", tokenAddress, err)
			}
			oldPoolAddress := tokenConfigReport.Output.TokenPool
			if oldPoolAddress == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("no pool currently registered for token %s on chain %d", tokenAddress, input.ChainSelector)
			}

			migrationReport, err := cldf_ops.ExecuteSequence(b, MigrateLockReleasePoolLiquidity, chains, tokens.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   input.ChainSelector,
				OldPoolAddress:  oldPoolAddress.Hex(),
				NewPoolAddress:  input.TokenPoolAddress,
				TimelockAddress: input.TimelockAddress,
				Amount:          input.LiquidityMigrationAmount,
				BasisPoints:     input.LiquidityMigrationBasisPoints,
				SetPoolConfig:   nil,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute liquidity migration on chain %d: %w", input.ChainSelector, err)
			}
			ops = append(ops, migrationReport.Output.BatchOps...)
		}

		// Validate the pool supports the token
		isSupported, err := cldf_ops.ExecuteOperation(b, token_pool.IsSupportedToken, evmChain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
			Args:          tokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if token %s is supported by token pool %s on %s: %w", tokenAddress, tokenPoolAddress, evmChain, err)
		}
		if !isSupported.Output {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s is not supported by token pool %s", tokenAddress, tokenPoolAddress)
		}

		// Configure minimum block confirmation (skip if already set to desired value).
		currentMinBlockReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetMinBlockConfirmations, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get min block confirmation from token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
		}
		if currentMinBlockReport.Output != input.MinFinalityValue {
			configureMinBlockConfirmationReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetMinBlockConfirmations, evmChain, evm_contract.FunctionInput[uint16]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddress,
				Args:          input.MinFinalityValue,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure minimum block confirmation for token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
			}
			configureMinBlockConfirmationOps, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{configureMinBlockConfirmationReport.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from write outputs: %w", err)
			}
			ops = append(ops, configureMinBlockConfirmationOps)
		}

		// Get the advanced pool hooks address
		advancedPoolHooksAddress, err := cldf_ops.ExecuteOperation(b, token_pool.GetAdvancedPoolHooks, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get advanced pool hooks address from token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
		}

		// Configure token pool for all remote chains (validates active pool supported chains once when upgrading).
		configureTokenPoolForRemoteChainsReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChains, evmChain, ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:     input.ChainSelector,
			TokenPoolAddress:  tokenPoolAddress,
			AdvancedPoolHooks: advancedPoolHooksAddress.Output,
			RemoteChains:      input.RemoteChains,
			RegistryAddress:   common.HexToAddress(input.RegistryAddress),
			TokenAddress:      tokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chains: %w", err)
		}
		ops = append(ops, configureTokenPoolForRemoteChainsReport.Output.BatchOps...)

		// Register the token with the token admin registry
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, v1_5_0.RegisterToken, evmChain, v1_5_0.RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenAddress:              tokenAddress,
			TokenPoolAddress:          registryTokenPoolAddress,
			ExternalAdmin:             common.HexToAddress(input.ExternalAdmin),
			TokenAdminRegistryAddress: common.HexToAddress(input.RegistryAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to register token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
		}
		ops = append(ops, registerTokenReport.Output.BatchOps...)

		return sequences.OnChainOutput{
			BatchOps: ops,
		}, nil
	},
)
