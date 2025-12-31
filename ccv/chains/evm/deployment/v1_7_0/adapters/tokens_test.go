package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	v1_7_0 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	outbound = 0
	inbound  = 1
)

func TestTokenAdapter(t *testing.T) {
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

			mcmsRegistry := changesets.NewMCMSReaderRegistry()
			tokenAdapterRegistry := tokens.NewTokenAdapterRegistry()
			tokenAdapterRegistry.RegisterTokenAdapter("evm", semver.MustParse("1.7.0"), &adapters.TokenAdapter{})

			// On each chain, deploy chain contracts & a token + token pool
			ds := datastore.NewMemoryDataStore()
			for _, chainSel := range []uint64{chainA, chainB} {
				deployChainOut, err := v1_7_0.DeployChainContracts(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_7_0.DeployChainContractsCfg]{
					Cfg: v1_7_0.DeployChainContractsCfg{
						ChainSel: chainSel,
						Params:   testsetup.CreateBasicContractParams(),
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = ds.Merge(deployChainOut.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts changeset")

				e.DataStore = ds.Seal()
				deployTokenAndPoolOut, err := v1_7_0.DeployTokenAndPool(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_7_0.DeployTokenAndPoolCfg]{
					Cfg: v1_7_0.DeployTokenAndPoolCfg{
						Accounts: map[common.Address]*big.Int{
							e.BlockChains.EVMChains()[chainSel].DeployerKey.From: big.NewInt(1_000_000),
						},
						ChainSel:                         chainSel,
						TokenPoolType:                    datastore.ContractType(burn_mint_token_pool.BurnMintContractType),
						TokenPoolVersion:                 burn_mint_token_pool.Version,
						TokenSymbol:                      "TEST",
						Decimals:                         18,
						ThresholdAmountForAdditionalCCVs: big.NewInt(1e18),
						Router: datastore.AddressRef{
							ChainSelector: chainSel,
							Type:          datastore.ContractType(router.ContractType),
							Version:       router.Version,
						},
					},
				})
				require.NoError(t, err, "Failed to apply DeployBurnMintTokenAndPool changeset")
				err = ds.Merge(deployTokenAndPoolOut.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployTokenAndPool changeset")
			}

			// Overwrite datastore in the environment
			e.DataStore = ds.Seal()

			var remoteToken *datastore.AddressRef
			if test.deriveTokenAddress {
				remoteToken = &datastore.AddressRef{
					Type:      datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
					Version:   burn_mint_erc20_with_drip.Version,
					Qualifier: "TEST",
				}
			}

			getRemoteChainConfig := func(
				remotePoolVersion *semver.Version,
				ccvs []datastore.AddressRef,
			) tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef] {
				return tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					RemoteToken: remoteToken,
					RemotePool: &datastore.AddressRef{
						Type:      datastore.ContractType(burn_mint_token_pool.BurnMintContractType),
						Version:   remotePoolVersion,
						Qualifier: "TEST",
					},
					DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(10, 100),
					DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfig(20, 200),
					CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfig(30, 300),
					CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(40, 400),
					TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
					OutboundCCVs:                             ccvs,
					InboundCCVs:                              ccvs,
				}
			}

			_, err = tokens.ConfigureTokensForTransfers(tokenAdapterRegistry, mcmsRegistry).Apply(*e, tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: chainA,
						TokenPoolRef: datastore.AddressRef{
							Type:      datastore.ContractType(burn_mint_token_pool.BurnMintContractType),
							Version:   burn_mint_token_pool.Version,
							Qualifier: "TEST",
						},
						RegistryRef: datastore.AddressRef{
							Type:    datastore.ContractType(token_admin_registry.ContractType),
							Version: token_admin_registry.Version,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							chainB: getRemoteChainConfig(burn_mint_token_pool.Version, []datastore.AddressRef{
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
							Type:      datastore.ContractType(burn_mint_token_pool.BurnMintContractType),
							Version:   burn_mint_token_pool.Version,
							Qualifier: "TEST",
						},
						RegistryRef: datastore.AddressRef{
							Type:    datastore.ContractType(token_admin_registry.ContractType),
							Version: token_admin_registry.Version,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							chainA: getRemoteChainConfig(burn_mint_token_pool.Version, []datastore.AddressRef{
								{
									Type:    datastore.ContractType(committee_verifier.ContractType),
									Version: committee_verifier.Version,
								},
							}),
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

				tokenPoolAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          datastore.ContractType(burn_mint_token_pool.BurnMintContractType),
					Version:       burn_mint_token_pool.Version,
					Qualifier:     "TEST",
				}, chainSel, evm_datastore_utils.ToEVMAddress)
				require.NoError(t, err, "Failed to find deployed token pool ref in datastore")
				tokenAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
					Version:       burn_mint_erc20_with_drip.Version,
					Qualifier:     "TEST",
				}, chainSel, evm_datastore_utils.ToEVMAddress)
				require.NoError(t, err, "Failed to find deployed token ref in datastore")
				registryAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: chainSel,
					Type:          datastore.ContractType(token_admin_registry.ContractType),
					Version:       token_admin_registry.Version,
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

				chainSupportReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetSupportedChains, evmChain, contract.FunctionInput[any]{
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

				rateLimiterStateReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetCurrentRateLimiterState, evmChain, contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
					ChainSelector: chainSel,
					Address:       tokenPoolAddr,
					Args: token_pool.GetCurrentRateLimiterStateArgs{
						RemoteChainSelector:     remoteChainSel,
						CustomBlockConfirmation: false,
					},
				})
				currentStates := rateLimiterStateReport.Output
				require.NoError(t, err, "Failed to get rate limiter config from token pool")
				require.Equal(t, getRemoteChainConfig(nil, nil).DefaultFinalityInboundRateLimiterConfig.IsEnabled, currentStates.InboundRateLimiterState.IsEnabled, "Inbound rate limiter enabled state should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).DefaultFinalityInboundRateLimiterConfig.Rate, currentStates.InboundRateLimiterState.Rate, "Inbound rate limiter rate should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).DefaultFinalityInboundRateLimiterConfig.Capacity, currentStates.InboundRateLimiterState.Capacity, "Inbound rate limiter capacity should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).DefaultFinalityOutboundRateLimiterConfig.IsEnabled, currentStates.OutboundRateLimiterState.IsEnabled, "Outbound rate limiter enabled state should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).DefaultFinalityOutboundRateLimiterConfig.Rate, currentStates.OutboundRateLimiterState.Rate, "Outbound rate limiter rate should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).DefaultFinalityOutboundRateLimiterConfig.Capacity, currentStates.OutboundRateLimiterState.Capacity, "Outbound rate limiter capacity should match")

				rateLimiterStateReport, err = operations.ExecuteOperation(e.OperationsBundle, token_pool.GetCurrentRateLimiterState, evmChain, contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
					ChainSelector: chainSel,
					Address:       tokenPoolAddr,
					Args: token_pool.GetCurrentRateLimiterStateArgs{
						RemoteChainSelector:     remoteChainSel,
						CustomBlockConfirmation: true,
					},
				})
				currentStates = rateLimiterStateReport.Output
				require.NoError(t, err, "Failed to get rate limiter config from token pool")
				require.Equal(t, getRemoteChainConfig(nil, nil).CustomFinalityInboundRateLimiterConfig.IsEnabled, currentStates.InboundRateLimiterState.IsEnabled, "Inbound rate limiter enabled state should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).CustomFinalityInboundRateLimiterConfig.Rate, currentStates.InboundRateLimiterState.Rate, "Inbound rate limiter rate should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).CustomFinalityInboundRateLimiterConfig.Capacity, currentStates.InboundRateLimiterState.Capacity, "Inbound rate limiter capacity should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).CustomFinalityOutboundRateLimiterConfig.IsEnabled, currentStates.OutboundRateLimiterState.IsEnabled, "Outbound rate limiter enabled state should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).CustomFinalityOutboundRateLimiterConfig.Rate, currentStates.OutboundRateLimiterState.Rate, "Outbound rate limiter rate should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).CustomFinalityOutboundRateLimiterConfig.Capacity, currentStates.OutboundRateLimiterState.Capacity, "Outbound rate limiter capacity should match")

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

				tokenTransferFeeConfig, err := boundTokenPool.GetTokenTransferFeeConfig(nil, common.Address{}, remoteChainSel, 0, []byte{})
				require.NoError(t, err, "Failed to get token transfer fee config from token pool")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.IsEnabled, tokenTransferFeeConfig.IsEnabled, "Token transfer fee config enabled state should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.DestGasOverhead, tokenTransferFeeConfig.DestGasOverhead, "Token transfer fee config dest gas overhead should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.DestBytesOverhead, tokenTransferFeeConfig.DestBytesOverhead, "Token transfer fee config dest bytes overhead should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.DefaultFinalityFeeUSDCents, tokenTransferFeeConfig.DefaultBlockConfirmationFeeUSDCents, "Token transfer fee config default finality fee USDCents should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.CustomFinalityFeeUSDCents, tokenTransferFeeConfig.CustomBlockConfirmationFeeUSDCents, "Token transfer fee config custom finality fee USDCents should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.DefaultFinalityTransferFeeBps, tokenTransferFeeConfig.DefaultBlockConfirmationTransferFeeBps, "Token transfer fee config default finality transfer fee BPS should match")
				require.Equal(t, getRemoteChainConfig(nil, nil).TokenTransferFeeConfig.CustomFinalityTransferFeeBps, tokenTransferFeeConfig.CustomBlockConfirmationTransferFeeBps, "Token transfer fee config custom finality transfer fee BPS should match")
			}
		})
	}
}
