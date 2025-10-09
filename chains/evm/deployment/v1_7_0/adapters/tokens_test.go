package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/adapters"
	v1_7_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	evm_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestTokenAdapter(t *testing.T) {
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

		router, err := datastore_utils.FindAndFormatRef(deployChainOut.DataStore.Seal(), datastore.AddressRef{
			ChainSelector: chainSel,
			Type:          datastore.ContractType(router.ContractType),
		}, chainSel, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Failed to find deployed router ref in datastore after DeployChainContracts changeset")
		rmnProxy, err := datastore_utils.FindAndFormatRef(deployChainOut.DataStore.Seal(), datastore.AddressRef{
			ChainSelector: chainSel,
			Type:          datastore.ContractType(rmn_proxy.ContractType),
		}, chainSel, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Failed to find deployed rmn proxy ref in datastore after DeployChainContracts changeset")

		deployTokenAndPoolOut, err := v1_7_0.DeployBurnMintTokenAndPool(mcmsRegistry).Apply(*e, changesets.WithMCMS[evm_tokens.DeployBurnMintTokenAndPoolInput]{
			Cfg: evm_tokens.DeployBurnMintTokenAndPoolInput{
				Accounts: map[common.Address]*big.Int{
					e.BlockChains.EVMChains()[chainSel].DeployerKey.From: big.NewInt(1_000_000),
				},
				TokenInfo: evm_tokens.TokenInfo{
					Decimals:  18,
					MaxSupply: big.NewInt(10_000_000),
					Name:      "TEST",
				},
				DeployTokenPoolInput: evm_tokens.DeployTokenPoolInput{
					ChainSel:         chainSel,
					TokenPoolType:    datastore.ContractType(burn_mint_token_pool.ContractType),
					TokenPoolVersion: semver.MustParse("1.7.0"),
					TokenSymbol:      "TEST",
					ConstructorArgs: token_pool.ConstructorArgs{
						LocalTokenDecimals: 18,
						Router:             router,
						RMNProxy:           rmnProxy,
					},
				},
			},
		})
		require.NoError(t, err, "Failed to apply DeployBurnMintTokenAndPool changeset")
		err = ds.Merge(deployTokenAndPoolOut.DataStore.Seal())
	}

	// Overwrite datastore in the environment
	e.DataStore = ds.Seal()

	remoteChainConfig := tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
		RemoteToken: &datastore.AddressRef{
			Type:      datastore.ContractType(burn_mint_erc677.ContractType),
			Version:   semver.MustParse("1.0.0"),
			Qualifier: "TEST",
		},
		RemotePool: &datastore.AddressRef{
			Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
			Version:   semver.MustParse("1.7.0"),
			Qualifier: "TEST",
		},
		InboundRateLimiterConfig: tokens.RateLimiterConfig{
			IsEnabled: true,
			Rate:      big.NewInt(10),
			Capacity:  big.NewInt(100),
		},
		OutboundRateLimiterConfig: tokens.RateLimiterConfig{
			IsEnabled: true,
			Rate:      big.NewInt(10),
			Capacity:  big.NewInt(100),
		},
		OutboundCCVs: []datastore.AddressRef{
			{
				Type:    datastore.ContractType(committee_verifier.ContractType),
				Version: semver.MustParse("1.7.0"),
			},
		},
		InboundCCVs: []datastore.AddressRef{
			{
				Type:    datastore.ContractType(committee_verifier.ContractType),
				Version: semver.MustParse("1.7.0"),
			},
		},
	}

	_, err = tokens.ConfigureTokensForTransfers(tokenAdapterRegistry, mcmsRegistry).Apply(*e, tokens.ConfigureTokensForTransfersConfig{
		Tokens: []tokens.TokenTransferConfig{
			{
				ChainSelector: chainA,
				TokenPoolRef: datastore.AddressRef{
					Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
					Version:   semver.MustParse("1.7.0"),
					Qualifier: "TEST",
				},
				RegistryRef: datastore.AddressRef{
					Type:    datastore.ContractType(token_admin_registry.ContractType),
					Version: semver.MustParse("1.5.0"),
				},
				RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					chainB: remoteChainConfig,
				},
			},
			{
				ChainSelector: chainB,
				TokenPoolRef: datastore.AddressRef{
					Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
					Version:   semver.MustParse("1.7.0"),
					Qualifier: "TEST",
				},
				RegistryRef: datastore.AddressRef{
					Type:    datastore.ContractType(token_admin_registry.ContractType),
					Version: semver.MustParse("1.5.0"),
				},
				RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					chainA: remoteChainConfig,
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
			Type:          datastore.ContractType(burn_mint_token_pool.ContractType),
			Version:       semver.MustParse("1.7.0"),
			Qualifier:     "TEST",
		}, chainSel, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Failed to find deployed token pool ref in datastore")
		tokenAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			ChainSelector: chainSel,
			Type:          datastore.ContractType(burn_mint_erc677.ContractType),
			Version:       semver.MustParse("1.0.0"),
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
			Version:       semver.MustParse("1.7.0"),
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

		inboundRateLimiterReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetCurrentInboundRateLimiterState, evmChain, contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       tokenPoolAddr,
			Args:          remoteChainSel,
		})
		require.NoError(t, err, "Failed to get inbound rate limiter config from token pool")
		require.Equal(t, remoteChainConfig.InboundRateLimiterConfig.IsEnabled, inboundRateLimiterReport.Output.IsEnabled, "Inbound rate limiter enabled state should match")
		require.Equal(t, remoteChainConfig.InboundRateLimiterConfig.Rate, inboundRateLimiterReport.Output.Rate, "Inbound rate limiter rate should match")
		require.Equal(t, remoteChainConfig.InboundRateLimiterConfig.Capacity, inboundRateLimiterReport.Output.Capacity, "Inbound rate limiter capacity should match")

		outboundRateLimiterReport, err := operations.ExecuteOperation(e.OperationsBundle, token_pool.GetCurrentOutboundRateLimiterState, evmChain, contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       tokenPoolAddr,
			Args:          remoteChainSel,
		})
		require.NoError(t, err, "Failed to get outbound rate limiter config from token pool")
		require.Equal(t, remoteChainConfig.OutboundRateLimiterConfig.IsEnabled, outboundRateLimiterReport.Output.IsEnabled, "Outbound rate limiter enabled state should match")
		require.Equal(t, remoteChainConfig.OutboundRateLimiterConfig.Rate, outboundRateLimiterReport.Output.Rate, "Outbound rate limiter rate should match")
		require.Equal(t, remoteChainConfig.OutboundRateLimiterConfig.Capacity, outboundRateLimiterReport.Output.Capacity, "Outbound rate limiter capacity should match")

		boundTokenPool, err := tp_bindings.NewTokenPool(tokenPoolAddr, evmChain.Client)
		require.NoError(t, err, "Failed to instantiate token pool contract")
		inboundCCVs, err := boundTokenPool.GetRequiredInboundCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{})
		require.NoError(t, err, "Failed to get inbound CCVs from token pool")
		require.Len(t, inboundCCVs, 1, "Number of inbound CCVs should match")
		require.Equal(t, verifierAddr, inboundCCVs[0], "Inbound CCV address should match")

		outboundCCVs, err := boundTokenPool.GetRequiredOutboundCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{})
		require.NoError(t, err, "Failed to get outbound CCVs from token pool")
		require.Len(t, outboundCCVs, 1, "Number of outbound CCVs should match")
		require.Equal(t, verifierAddr, outboundCCVs[0], "Outbound CCV address should match")
	}
}
