package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestDeployTokenPool(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput
		expectedErr string
	}{
		{
			desc: "happy path",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				var rmnProxyAddress common.Address
				var routerAddress common.Address
				for _, addr := range chainReport.Output.Addresses {
					if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
						rmnProxyAddress = common.HexToAddress(addr.Address)
					}
					if addr.Type == datastore.ContractType(router.ContractType) {
						routerAddress = common.HexToAddress(addr.Address)
					}
				}
				return tokens.DeployTokenPoolInput{
					ChainSel:         chainReport.Input.ChainSelector,
					TokenPoolType:    datastore.ContractType(burn_mint_token_pool.ContractType),
					TokenPoolVersion: semver.MustParse("1.7.0"),
					TokenSymbol:      tokenReport.Input.Args.Symbol,
					RateLimitAdmin:   common.HexToAddress("0x01"),
					ConstructorArgs: token_pool.ConstructorArgs{
						Token:              common.HexToAddress(tokenReport.Output.Address),
						LocalTokenDecimals: 18,
						Allowlist: []common.Address{
							common.HexToAddress("0x02"),
						},
						RMNProxy: rmnProxyAddress,
						Router:   routerAddress,
					},
				}
			},
			expectedErr: "",
		},
		{
			desc: "incorrect chain selector",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				return tokens.DeployTokenPoolInput{
					ChainSel: 1,
				}
			},
			expectedErr: "chain selector 1 does not match chain",
		},
		{
			desc: "token symbol not defined",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				return tokens.DeployTokenPoolInput{
					ChainSel: chainReport.Input.ChainSelector,
				}
			},
			expectedErr: "token symbol must be defined",
		},
		{
			desc: "token pool type not defined",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				return tokens.DeployTokenPoolInput{
					ChainSel:    chainReport.Input.ChainSelector,
					TokenSymbol: tokenReport.Input.Args.Symbol,
				}
			},
			expectedErr: "token pool type must be defined",
		},
		{
			desc: "token pool version not defined",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				return tokens.DeployTokenPoolInput{
					ChainSel:      chainReport.Input.ChainSelector,
					TokenSymbol:   tokenReport.Input.Args.Symbol,
					TokenPoolType: datastore.ContractType(burn_mint_token_pool.ContractType),
				}
			},
			expectedErr: "token pool version must be defined",
		},
		{
			desc: "token address not defined",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				return tokens.DeployTokenPoolInput{
					ChainSel:         chainReport.Input.ChainSelector,
					TokenSymbol:      tokenReport.Input.Args.Symbol,
					TokenPoolType:    datastore.ContractType(burn_mint_token_pool.ContractType),
					TokenPoolVersion: semver.MustParse("1.7.0"),
				}
			},
			expectedErr: "token address must be defined",
		},
		{
			desc: "rmn proxy address not defined",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				return tokens.DeployTokenPoolInput{
					ChainSel:         chainReport.Input.ChainSelector,
					TokenSymbol:      tokenReport.Input.Args.Symbol,
					TokenPoolType:    datastore.ContractType(burn_mint_token_pool.ContractType),
					TokenPoolVersion: semver.MustParse("1.7.0"),
					ConstructorArgs: token_pool.ConstructorArgs{
						Token: common.HexToAddress(tokenReport.Output.Address),
					},
				}
			},
			expectedErr: "rmn proxy address must be defined",
		},
		{
			desc: "router address not defined",
			makeInput: func(tokenReport operations.Report[contract.DeployInput[burn_mint_erc677.ConstructorArgs], datastore.AddressRef], chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployTokenPoolInput {
				var rmnProxyAddress common.Address
				for _, addr := range chainReport.Output.Addresses {
					if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
						rmnProxyAddress = common.HexToAddress(addr.Address)
					}
				}
				return tokens.DeployTokenPoolInput{
					ChainSel:         chainReport.Input.ChainSelector,
					TokenSymbol:      tokenReport.Input.Args.Symbol,
					TokenPoolType:    datastore.ContractType(burn_mint_token_pool.ContractType),
					TokenPoolVersion: semver.MustParse("1.7.0"),
					ConstructorArgs: token_pool.ConstructorArgs{
						Token:    common.HexToAddress(tokenReport.Output.Address),
						RMNProxy: rmnProxyAddress,
					},
				}
			},
			expectedErr: "router address must be defined",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSel := uint64(5009297550715157269)
			e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
				chainSel: {NumAdditionalAccounts: 1},
			})
			require.NoError(t, err, "Failed to create test environment")

			// Deploy chain
			chainReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSel],
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSel,
					ContractParams: testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Deploy token
			tokenReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				burn_mint_erc677.Deploy,
				e.BlockChains.EVMChains()[chainSel],
				contract.DeployInput[burn_mint_erc677.ConstructorArgs]{
					ChainSelector:  chainSel,
					TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc677.ContractType, *semver.MustParse("1.0.0")),
					Args: burn_mint_erc677.ConstructorArgs{
						Name:      "Test Token",
						Symbol:    "TEST",
						Decimals:  18,
						MaxSupply: big.NewInt(1_000_000),
					},
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")

			// Deploy token pool
			input := test.makeInput(tokenReport, chainReport)
			poolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenPool,
				e.BlockChains.EVMChains()[chainSel],
				input,
			)
			if test.expectedErr != "" {
				require.Error(t, err, "ExecuteSequence should error")
				require.Contains(t, err.Error(), test.expectedErr)
				return
			}
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, poolReport.Output.Addresses, 1, "Expected 1 address in output")

			// Check token
			getTokenReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				token_pool.GetToken,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       common.HexToAddress(poolReport.Output.Addresses[0].Address),
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.ConstructorArgs.Token, getTokenReport.Output, "Expected token address to be the same as the deployed token")

			// Check router
			getRouterReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				token_pool.GetRouter,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       common.HexToAddress(poolReport.Output.Addresses[0].Address),
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.ConstructorArgs.Router, getRouterReport.Output, "Expected router address to be the same as the deployed router")

			// Check rmn proxy
			getRmnProxyReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				token_pool.GetRmnProxy,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       common.HexToAddress(poolReport.Output.Addresses[0].Address),
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.ConstructorArgs.RMNProxy, getRmnProxyReport.Output, "Expected rmn proxy address to be the same as the deployed rmn proxy")

			// Check allowlist
			getAllowlistReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				token_pool.GetAllowList,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       common.HexToAddress(poolReport.Output.Addresses[0].Address),
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.ConstructorArgs.Allowlist, getAllowlistReport.Output, "Expected allowlist address to be the same as the deployed allowlist")

			// Check rate limit admin
			/* TODO: This is broken, rateLimitAdmin is not getting set for some reason. Need to investigate.
			getRateLimitAdminReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				token_pool.GetRateLimitAdmin,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       common.HexToAddress(poolReport.Output.Addresses[0].Address),
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.RateLimitAdmin, getRateLimitAdminReport.Output, "Expected rate limit admin address to be the same as the deployed rate limit admin")
			*/
		})
	}
}
