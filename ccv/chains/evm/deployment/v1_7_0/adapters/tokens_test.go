package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnm_drip_v1_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
	v1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/changesets"
	burn_mint_token_pool_v1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	_ "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	v1_7_0 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	bnm_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

const (
	outbound = 0
	inbound  = 1
)

// requireRateLimiterScaled asserts that on-chain rate limiter state matches the expected token amounts
// after scaling by decimals and (for inbound) the 1.1x factor used in GenerateTPRLConfigs.
func requireRateLimiterScaled(t *testing.T, rate, capacity float64, actualRate, actualCapacity *big.Int, decimals int, isInbound bool) {
	extraPercent := 0.0
	if isInbound {
		extraPercent = 0.10
	}
	expectedRate := tokens.ScaleFloatToBigInt(rate, decimals, extraPercent)
	expectedCapacity := tokens.ScaleFloatToBigInt(capacity, decimals, extraPercent)
	if actualRate == nil {
		actualRate = big.NewInt(0)
	}
	if actualCapacity == nil {
		actualCapacity = big.NewInt(0)
	}
	require.Zero(t, expectedRate.Cmp(actualRate), "Rate limiter rate should match (scaled)")
	require.Zero(t, expectedCapacity.Cmp(actualCapacity), "Rate limiter capacity should match (scaled)")
}

func TestTokenAdapter(t *testing.T) {
	tokenAdapterRegistry := tokens.GetTokenAdapterRegistry()
	tokenAdapterRegistry.RegisterTokenAdapter("evm", semver.MustParse("1.6.1"), &v1_6_1_adapters.TokenAdapter{})

	tests := []struct {
		desc               string
		deriveTokenAddress bool
	}{
		{
			desc:               "derive remote token address",
			deriveTokenAddress: true,
		},
		{
			desc: "input remote token address",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			chainA := uint64(5009297550715157269)
			chainB := uint64(4949039107694359620)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainA, chainB}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			mcmsRegistry := changesets.GetRegistry()

			// On each chain, deploy chain contracts & a token + token pool
			ds := datastore.NewMemoryDataStore()
			for _, chainSel := range []uint64{chainA, chainB} {
				create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
					ChainSelector:  chainSel,
					Args: create2_factory.ConstructorArgs{
						AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
					},
				}, nil)
				require.NoError(t, err, "Failed to deploy CREATE2Factory")

				deployChainOut, err := v1_7_0.DeployChainContracts(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_7_0.DeployChainContractsCfg]{
					Cfg: v1_7_0.DeployChainContractsCfg{
						ChainSel:         chainSel,
						CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
						Params:           testsetup.CreateBasicContractParams(),
						DeployerKeyOwned: true,
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = ds.Merge(deployChainOut.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts changeset")

				e.DataStore = ds.Seal()

				// Deploy a legacy 1.6.1 on chain B (chain A uses TokenExpansion below)
				if chainSel == chainB {
					deployTokenAndPoolOut, err := v1_6_1.DeployTokenAndPool(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_6_1.DeployTokenAndPoolCfg]{
						Cfg: v1_6_1.DeployTokenAndPoolCfg{
							Accounts: map[common.Address]*big.Int{
								e.BlockChains.EVMChains()[chainSel].DeployerKey.From: big.NewInt(1_000_000),
							},
							ChainSel:         chainSel,
							TokenPoolType:    datastore.ContractType(burn_mint_token_pool_v1_6_1.ContractType),
							TokenPoolVersion: burn_mint_token_pool_v1_6_1.Version,
							TokenSymbol:      "TEST",
							Decimals:         18,
							Router: datastore.AddressRef{
								ChainSelector: chainSel,
								Type:          datastore.ContractType(router.ContractType),
								Version:       semver.MustParse("1.2.0"),
							},
						},
					})
					require.NoError(t, err, "Failed to apply DeployTokenAndPool changeset")
					err = ds.Merge(deployTokenAndPoolOut.DataStore.Seal())
					e.DataStore = ds.Seal()
				}
			}

			// Deploy 2.0.0 token + pool on chain A via TokenExpansion
			preMint := uint64(1_000)
			expansionOut, err := tokens.TokenExpansion().Apply(*e, tokens.TokenExpansionInput{
				ChainAdapterVersion: semver.MustParse("2.0.0"),
				MCMS:                mcms.Input{},
				TokenExpansionInputPerChain: map[uint64]tokens.TokenExpansionInputPerChain{
					chainA: {
						TokenPoolVersion:      burn_mint_token_pool.Version,
						SkipOwnershipTransfer: true,
						DeployTokenInput: &tokens.DeployTokenInput{
							Name:          "TEST",
							Symbol:        "TEST",
							Decimals:      18,
							PreMint:       &preMint,
							ExternalAdmin: e.BlockChains.EVMChains()[chainA].DeployerKey.From.Hex(),
							CCIPAdmin:     e.BlockChains.EVMChains()[chainA].DeployerKey.From.Hex(),
							Type:          bnm_drip_v1_0.ContractType,
						},
						DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
							PoolType:           string(burn_mint_token_pool.ContractType),
							TokenPoolQualifier: "TEST",
						},
					},
				},
			})
			require.NoError(t, err, "Failed to apply TokenExpansion changeset")
			err = ds.Merge(expansionOut.DataStore.Seal())
			require.NoError(t, err, "Failed to merge TokenExpansion datastore")
			e.DataStore = ds.Seal()

			// Remote token refs differ by chain: chain A (2.0.0) stores tokens with bnm_drip_v1_0.ContractType,
			// chain B (1.6.1) stores them with burn_mint_erc20_with_drip.ContractType.
			var remoteTokenForChainA *datastore.AddressRef // looking up chain A's token (from chain B's perspective)
			var remoteTokenForChainB *datastore.AddressRef // looking up chain B's token (from chain A's perspective)
			if test.deriveTokenAddress {
				remoteTokenForChainA = &datastore.AddressRef{
					Type:      datastore.ContractType(bnm_drip_v1_0.ContractType),
					Qualifier: "TEST",
				}
				remoteTokenForChainB = &datastore.AddressRef{
					Type:      datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
					Version:   burn_mint_erc20_with_drip.Version,
					Qualifier: "TEST",
				}
			}

			getRemoteChainConfig := func(
				remoteToken *datastore.AddressRef,
				remotePoolVersion *semver.Version,
				ccvs []datastore.AddressRef,
			) tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef] {
				return tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					RemoteToken: remoteToken,
					RemotePool: &datastore.AddressRef{
						Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
						Version:   remotePoolVersion,
						Qualifier: "TEST",
					},
					DefaultFinalityInboundRateLimiterConfig: tokens.RateLimiterConfigFloatInput{
						IsEnabled: true,
						Rate:      10,
						Capacity:  100,
					},
					DefaultFinalityOutboundRateLimiterConfig: tokens.RateLimiterConfigFloatInput{
						IsEnabled: true,
						Rate:      10,
						Capacity:  100,
					},
					CustomFinalityInboundRateLimiterConfig: tokens.RateLimiterConfigFloatInput{
						IsEnabled: true,
						Rate:      10,
						Capacity:  100,
					},
					CustomFinalityOutboundRateLimiterConfig: tokens.RateLimiterConfigFloatInput{
						IsEnabled: true,
						Rate:      10,
						Capacity:  100,
					},
					OutboundCCVs: ccvs,
					InboundCCVs:  ccvs,
				}
			}

			_, err = tokens.ConfigureTokensForTransfers(tokenAdapterRegistry, mcmsRegistry).Apply(*e, tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: chainA,
						TokenPoolRef: datastore.AddressRef{
							Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
							Version:   burn_mint_token_pool.Version,
							Qualifier: "TEST",
						},
						RegistryRef: datastore.AddressRef{
							Type:    datastore.ContractType(token_admin_registry.ContractType),
							Version: semver.MustParse("1.5.0"),
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							chainB: getRemoteChainConfig(remoteTokenForChainB, burn_mint_token_pool_v1_6_1.Version, []datastore.AddressRef{
								{
									Type:    datastore.ContractType(committee_verifier.ContractType),
									Version: committee_verifier.Version,
								},
							}),
						},
					},
					{
						ChainSelector: chainB,
						TokenPoolRef: datastore.AddressRef{
							Type:      datastore.ContractType(burn_mint_token_pool_v1_6_1.ContractType),
							Version:   burn_mint_token_pool_v1_6_1.Version,
							Qualifier: "TEST",
						},
						RegistryRef: datastore.AddressRef{
							Type:    datastore.ContractType(token_admin_registry.ContractType),
							Version: semver.MustParse("1.5.0"),
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							chainA: getRemoteChainConfig(remoteTokenForChainA, burn_mint_token_pool.Version, nil),
						},
					},
				},
			})
			require.NoError(t, err, "Failed to apply ConfigureTokensForTransfers changeset")

			// Clear bundle for checks, otherwise the operations framework will skip duplicate calls
			e.OperationsBundle = operations.NewBundle(
				e.GetContext,
				e.Logger,
				operations.NewMemoryReporter(),
			)
			for _, chainSel := range []uint64{chainA, chainB} {
				evmChain := e.BlockChains.EVMChains()[chainSel]

				var tokenPoolType datastore.ContractType
				var version *semver.Version
				var tokenType datastore.ContractType
				if chainSel == chainA {
					tokenPoolType = datastore.ContractType(burn_mint_token_pool.ContractType)
					version = burn_mint_token_pool.Version
					tokenType = datastore.ContractType(bnm_drip_v1_0.ContractType)
				} else {
					tokenPoolType = datastore.ContractType(burn_mint_token_pool_v1_6_1.ContractType)
					version = burn_mint_token_pool_v1_6_1.Version
					tokenType = datastore.ContractType(burn_mint_erc20_with_drip.ContractType)
				}

				tokenPoolAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          tokenPoolType,
					Version:       version,
					Qualifier:     "TEST",
				}, chainSel, evm_datastore_utils.ToEVMAddress)
				require.NoError(t, err, "Failed to find deployed token pool ref in datastore")
				tokenAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          tokenType,
					Qualifier:     "TEST",
				}, chainSel, evm_datastore_utils.ToEVMAddress)
				require.NoError(t, err, "Failed to find deployed token ref in datastore")
				registryAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          datastore.ContractType(token_admin_registry.ContractType),
					Version:       semver.MustParse("1.5.0"),
				}, chainSel, evm_datastore_utils.ToEVMAddress)
				require.NoError(t, err, "Failed to find deployed registry ref in datastore")
				verifierAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          datastore.ContractType(committee_verifier.ContractType),
					Version:       committee_verifier.Version,
				}, chainSel, evm_datastore_utils.ToEVMAddress)
				require.NoError(t, err, "Failed to find deployed verifier ref in datastore")

				tokenConfigReport, err := operations.ExecuteOperation(e.OperationsBundle, token_admin_registry.GetTokenConfig, evmChain, contract.FunctionInput[common.Address]{
					ChainSelector: chainSel,
					Address:       registryAddr,
					Args:          tokenAddr,
				})
				require.NoError(t, err, "Failed to get token config from token admin registry")
				require.Equal(t, tokenPoolAddr, tokenConfigReport.Output.TokenPool, "Token pool address in registry should match deployed token pool address")
				require.Equal(t, evmChain.DeployerKey.From, tokenConfigReport.Output.Administrator, "Deployer should be the admin of the token in the registry")

				chainSupportReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetSupportedChains, evmChain, contract.FunctionInput[struct{}]{
					ChainSelector: chainSel,
					Address:       tokenPoolAddr,
				})
				require.NoError(t, err, "Failed to get supported chains from token pool")
				require.Len(t, chainSupportReport.Output, 1, "There should be 1 supported remote chain in the token pool")
				var remoteChainSel uint64
				if chainSel == chainA {
					remoteChainSel = chainB
				} else {
					remoteChainSel = chainA
				}
				require.Equal(t, remoteChainSel, chainSupportReport.Output[0], "Remote chain in token pool should match expected")

				// GetCurrentRateLimiterState is only available in version 2.0.0+
				if version.GreaterThan(semver.MustParse("1.6.9")) || version.Equal(semver.MustParse("2.0.0")) {
					rateLimiterStateReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetCurrentRateLimiterState, evmChain, contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
						ChainSelector: chainSel,
						Address:       tokenPoolAddr,
						Args: token_pool.GetCurrentRateLimiterStateArgs{
							RemoteChainSelector: remoteChainSel,
						},
					})
					require.NoError(t, err, "Failed to get rate limiter config from token pool")
					currentStates := rateLimiterStateReport.Output
					cfg := getRemoteChainConfig(nil, nil, nil)
					const decimals = 18
					require.Equal(t, cfg.DefaultFinalityInboundRateLimiterConfig.IsEnabled, currentStates.InboundRateLimiterState.IsEnabled, "Inbound rate limiter enabled state should match")
					requireRateLimiterScaled(t, cfg.DefaultFinalityInboundRateLimiterConfig.Rate, cfg.DefaultFinalityInboundRateLimiterConfig.Capacity, currentStates.InboundRateLimiterState.Rate, currentStates.InboundRateLimiterState.Capacity, decimals, true)
					require.Equal(t, cfg.DefaultFinalityOutboundRateLimiterConfig.IsEnabled, currentStates.OutboundRateLimiterState.IsEnabled, "Outbound rate limiter enabled state should match")
					requireRateLimiterScaled(t, cfg.DefaultFinalityOutboundRateLimiterConfig.Rate, cfg.DefaultFinalityOutboundRateLimiterConfig.Capacity, currentStates.OutboundRateLimiterState.Rate, currentStates.OutboundRateLimiterState.Capacity, decimals, false)
				}

				// Chain A has a 2.0.0 token pool so should have set CCVs
				if chainSel == chainA {
					boundTokenPool, err := tp_bindings.NewTokenPool(tokenPoolAddr, evmChain.Client)
					require.NoError(t, err, "Failed to instantiate token pool contract")
					inboundCCVs, err := boundTokenPool.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{}, inbound)
					require.NoError(t, err, "Failed to get inbound CCVs from token pool")
					require.Len(t, inboundCCVs, 1, "Number of inbound CCVs should match")
					require.Equal(t, verifierAddr, inboundCCVs[0], "Inbound CCV address should match")

					outboundCCVs, err := boundTokenPool.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{}, outbound)
					require.NoError(t, err, "Failed to get outbound CCVs from token pool")
					require.Len(t, outboundCCVs, 1, "Number of outbound CCVs should match")
					require.Equal(t, verifierAddr, outboundCCVs[0], "Outbound CCV address should match")
				}
			}
		})
	}
}

func TestTokenExpansion(t *testing.T) {

	chainA := uint64(5009297550715157269)
	chainB := uint64(4949039107694359620)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainA, chainB}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	mcmsRegistry := changesets.GetRegistry()

	ds := datastore.NewMemoryDataStore()
	for _, chainSel := range []uint64{chainA, chainB} {
		create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSel,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
			},
		}, nil)
		require.NoError(t, err)

		deployChainOut, err := v1_7_0.DeployChainContracts(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_7_0.DeployChainContractsCfg]{
			Cfg: v1_7_0.DeployChainContractsCfg{
				ChainSel:         chainSel,
				CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
				Params:           testsetup.CreateBasicContractParams(),
				DeployerKeyOwned: true,
			},
		})
		require.NoError(t, err)
		err = ds.Merge(deployChainOut.DataStore.Seal())
		require.NoError(t, err)
		e.DataStore = ds.Seal()
	}

	deployerA := e.BlockChains.EVMChains()[chainA].DeployerKey.From
	deployerB := e.BlockChains.EVMChains()[chainB].DeployerKey.From

	maxSupply := uint64(1_000_000)
	preMint := uint64(100_000)

	type chainTestData struct {
		symbol   string
		deployer common.Address
	}
	chainData := map[uint64]chainTestData{
		chainA: {symbol: "TSTA", deployer: deployerA},
		chainB: {symbol: "TSTB", deployer: deployerB},
	}

	expansionInput := make(map[uint64]tokens.TokenExpansionInputPerChain)
	for chainSel, data := range chainData {
		expansionInput[chainSel] = tokens.TokenExpansionInputPerChain{
			TokenPoolVersion:      burn_mint_token_pool.Version,
			SkipOwnershipTransfer: true,
			DeployTokenInput: &tokens.DeployTokenInput{
				Name:          "Test Token " + data.symbol,
				Symbol:        data.symbol,
				Decimals:      18,
				Supply:        &maxSupply,
				PreMint:       &preMint,
				ExternalAdmin: data.deployer.Hex(),
				CCIPAdmin:     data.deployer.Hex(),
				Type:          bnm_drip_v1_0.ContractType,
			},
			DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
				PoolType:           string(burn_mint_token_pool.ContractType),
				TokenPoolQualifier: data.symbol,
			},
		}
	}

	output, err := tokens.TokenExpansion().Apply(*e, tokens.TokenExpansionInput{
		ChainAdapterVersion:        semver.MustParse("2.0.0"),
		MCMS:                       mcms.Input{},
		TokenExpansionInputPerChain: expansionInput,
	})
	require.NoError(t, err, "TokenExpansion should succeed")

	err = ds.Merge(output.DataStore.Seal())
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	// Fresh operations bundle to avoid duplicate-call skipping
	e.OperationsBundle = operations.NewBundle(e.GetContext, e.Logger, operations.NewMemoryReporter())

	for _, chainSel := range []uint64{chainA, chainB} {
		data := chainData[chainSel]
		evmChain := e.BlockChains.EVMChains()[chainSel]

		// Verify token exists in datastore and on-chain
		tokenAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: chainSel,
			Type:          datastore.ContractType(bnm_drip_v1_0.ContractType),
			Qualifier:     data.symbol,
		}, chainSel, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Token should exist in datastore")

		tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(tokenAddr, evmChain.Client)
		require.NoError(t, err)

		onChainSymbol, err := tokenContract.Symbol(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, data.symbol, onChainSymbol)

		onChainDecimals, err := tokenContract.Decimals(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, uint8(18), onChainDecimals)

		expectedPreMint := tokens.ScaleTokenAmount(new(big.Int).SetUint64(preMint), 18)
		balance, err := tokenContract.BalanceOf(&bind.CallOpts{Context: t.Context()}, data.deployer)
		require.NoError(t, err)
		require.Equal(t, expectedPreMint.String(), balance.String(), "Deployer should hold pre-minted tokens")

		// Verify token pool exists in datastore
		poolAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: chainSel,
			Type:          datastore.ContractType(burn_mint_token_pool.ContractType),
			Version:       burn_mint_token_pool.Version,
			Qualifier:     data.symbol,
		}, chainSel, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Token pool should exist in datastore")

		// Verify token pool points to the correct token
		getTokenReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, evmChain, contract.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       poolAddr,
		})
		require.NoError(t, err)
		require.Equal(t, tokenAddr, getTokenReport.Output, "Token pool should point to the deployed token")

		// Verify token pool decimals
		getDecimalsReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetTokenDecimals, evmChain, contract.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       poolAddr,
		})
		require.NoError(t, err)
		require.Equal(t, uint8(18), getDecimalsReport.Output, "Token pool decimals should match token decimals")

		// Verify mint/burn roles were granted to the pool on the token
		minterRole, err := tokenContract.MINTERROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		hasMinterRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, poolAddr)
		require.NoError(t, err)
		require.True(t, hasMinterRole, "Token pool should have minter role on the token")

		burnerRole, err := tokenContract.BURNERROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		hasBurnerRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, burnerRole, poolAddr)
		require.NoError(t, err)
		require.True(t, hasBurnerRole, "Token pool should have burner role on the token")
	}
}
