package sequences

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	burn_mint_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_to_address_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_with_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/lock_release_token_pool"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func TestDeployTokenPool(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "TEST"
	tokenDecimals := uint8(18)

	testCases := []struct {
		name                string
		poolType            cldf.ContractType
		poolVersion         *semver.Version
		acceptLiquidity     *bool
		burnAddress         string
		allowlist           []string
		expectedTypeVersion string
	}{
		{
			name:                "BurnMintTokenPool_v1_5_1",
			poolType:            burn_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			expectedTypeVersion: "BurnMintTokenPool 1.5.1",
		},
		{
			name:                "BurnFromMintTokenPool_v1_5_1",
			poolType:            burn_from_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			expectedTypeVersion: "BurnFromMintTokenPool 1.5.1",
		},
		{
			name:                "BurnToAddressMintTokenPool_v1_5_1",
			poolType:            burn_to_address_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			burnAddress:         "0x000000000000000000000000000000000000dead",
			expectedTypeVersion: "BurnToAddressTokenPool 1.5.1",
		},
		{
			name:                "BurnWithFromMintTokenPool_v1_5_1",
			poolType:            burn_with_from_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			expectedTypeVersion: "BurnWithFromMintTokenPool 1.5.1",
		},
		{
			name:                "LockReleaseTokenPool_v1_5_1_accept_liquidity",
			poolType:            lock_release_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			acceptLiquidity:     boolPtr(true),
			expectedTypeVersion: "LockReleaseTokenPool 1.5.1",
		},
		{
			name:                "LockReleaseTokenPool_v1_5_1_no_liquidity",
			poolType:            lock_release_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			acceptLiquidity:     boolPtr(false),
			expectedTypeVersion: "LockReleaseTokenPool 1.5.1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			chain := e.BlockChains.EVMChains()[chainSelector]
			ds := datastore.NewMemoryDataStore()

			tokenAddr := deployTestToken(t, chain, tokenSymbol, tokenDecimals)
			ref := datastore.AddressRef{
				Type:          datastore.ContractType(burn_mint_erc20.ContractType),
				Version:       semver.MustParse("1.0.0"),
				Address:       tokenAddr.Hex(),
				ChainSelector: chainSelector,
				Qualifier:     tokenSymbol,
			}
			err = ds.Addresses().Add(ref)
			require.NoError(t, err, "Failed to add token address to datastore")

			routerAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
			err = ds.Addresses().Add(datastore.AddressRef{
				Type:          datastore.ContractType(router.ContractType),
				Version:       router.Version,
				Address:       routerAddress.Hex(),
				ChainSelector: chainSelector,
			})
			require.NoError(t, err, "Failed to add router address to datastore")

			rmnProxyAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
			err = ds.Addresses().Add(datastore.AddressRef{
				Type:          datastore.ContractType(rmnproxyops.ContractType),
				Version:       semver.MustParse("1.0.0"),
				Address:       rmnProxyAddress.Hex(),
				ChainSelector: chainSelector,
			})
			require.NoError(t, err, "Failed to add RMN proxy address to datastore")

			e.DataStore = ds.Seal()

			input := tokenapi.DeployTokenPoolInput{
				TokenRef:           &ref,
				TokenPoolQualifier: "",
				PoolType:           string(tc.poolType),
				TokenPoolVersion:   tc.poolVersion,
				Allowlist:          tc.allowlist,
				AcceptLiquidity:    tc.acceptLiquidity,
				BurnAddress:        tc.burnAddress,
				ChainSelector:      chainSelector,
				ExistingDataStore:  e.DataStore,
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
			require.NoError(t, err, "Failed to execute DeployTokenPool sequence for %s", tc.name)
			require.NotNil(t, report, "DeployTokenPool report should not be nil for %s", tc.name)

			require.GreaterOrEqual(t, len(report.Output.Addresses), 1, "Should have at least one deployed address")

			poolRef := report.Output.Addresses[0]
			require.Equal(t, input.TokenRef.Address, poolRef.Qualifier)
			t.Logf("Deployed %s at address: %s", tc.name, poolRef.Address)

			poolAddr := common.HexToAddress(poolRef.Address)
			poolContract, err := burn_mint_token_pool_bindings.NewBurnMintTokenPool(poolAddr, chain.Client)
			require.NoError(t, err, "Failed to create pool contract binding")

			typeAndVersion, err := poolContract.TypeAndVersion(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get TypeAndVersion from pool")
			require.Equal(t, tc.expectedTypeVersion, typeAndVersion, "TypeAndVersion mismatch for %s", tc.name)

			onChainToken, err := poolContract.GetToken(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get token address from pool")
			require.Equal(t, tokenAddr, onChainToken, "Token address mismatch for %s", tc.name)

			onChainDecimals, err := poolContract.GetTokenDecimals(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get token decimals from pool")
			require.Equal(t, tokenDecimals, onChainDecimals, "Token decimals mismatch for %s", tc.name)

			onChainRouter, err := poolContract.GetRouter(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get router address from pool")
			require.Equal(t, routerAddress, onChainRouter, "Router address mismatch for %s", tc.name)

			onChainRmnProxy, err := poolContract.GetRmnProxy(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get RMN proxy address from pool")
			require.Equal(t, rmnProxyAddress, onChainRmnProxy, "RMN proxy address mismatch for %s", tc.name)
		})
	}
}

func deployTestToken(t *testing.T, chain evm.Chain, symbol string, decimals uint8) common.Address {
	tokenAddr, tx, _, err := bnm_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		symbol,
		decimals,
		big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)),
		big.NewInt(0),
	)
	require.NoError(t, err, "Failed to deploy test token")

	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm test token deployment")

	return tokenAddr
}

func boolPtr(b bool) *bool {
	return &b
}
