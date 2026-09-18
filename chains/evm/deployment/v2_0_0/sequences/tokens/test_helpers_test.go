package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnm_drip_v1_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	latest_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	tokenscore "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var thresholdAmountForAdditionalCCVs = big.NewInt(1_000_000_000_000_000_000)

type tokenExpansionResult struct {
	TokenAddress         common.Address
	TokenPoolAddress     common.Address
	AdvancedHooksAddress common.Address
}

// deployTokenAndPoolViaExpansion uses TokenExpansion to deploy a test token and burn-mint token pool.
// It populates the environment's DataStore with the chain report addresses, calls TokenExpansion,
// and merges the output back into the environment's DataStore.
func deployTokenAndPoolViaExpansion(
	t *testing.T,
	e *deployment.Environment,
	chainSel uint64,
	chainReportAddresses []datastore.AddressRef,
) tokenExpansionResult {
	return deployTokenAndPoolViaExpansionWithDecimals(t, e, chainSel, chainReportAddresses, 18)
}

func deployTokenAndPoolViaExpansionWithDecimals(
	t *testing.T,
	e *deployment.Environment,
	chainSel uint64,
	chainReportAddresses []datastore.AddressRef,
	decimals uint8,
) tokenExpansionResult {
	t.Helper()

	ds := datastore.NewMemoryDataStore()
	for _, addr := range chainReportAddresses {
		require.NoError(t, ds.Addresses().Add(addr))
	}
	e.DataStore = ds.Seal()

	preMint := uint64(1_000)
	deployer := e.BlockChains.EVMChains()[chainSel].DeployerKey.From

	out, err := tokenscore.TokenExpansion().Apply(*e, tokenscore.TokenExpansionInput{
		ChainAdapterVersion: semver.MustParse("2.0.0"),
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokenscore.TokenExpansionInputPerChain{
			chainSel: {
				TokenPoolVersion:      burn_mint_token_pool.Version,
				SkipOwnershipTransfer: true,
				DeployTokenInput: &tokenscore.DeployTokenInput{
					Name:          "TEST",
					Symbol:        "TEST",
					Decimals:      decimals,
					PreMint:       &preMint,
					ExternalAdmin: deployer.Hex(),
					CCIPAdmin:     deployer.Hex(),
					Type:          bnm_drip_v1_0.ContractType,
				},
				DeployTokenPoolInput: &tokenscore.DeployTokenPoolInput{
					PoolType:                         string(burn_mint_token_pool.ContractType),
					TokenPoolQualifier:               "TEST",
					ThresholdAmountForAdditionalCCVs: thresholdAmountForAdditionalCCVs.String(),
				},
			},
		},
	})
	require.NoError(t, err, "TokenExpansion should succeed")

	require.NoError(t, ds.Merge(out.DataStore.Seal()))
	e.DataStore = ds.Seal()

	tokenAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(bnm_drip_v1_0.ContractType),
		Qualifier:     "TEST",
	}, chainSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "Token should exist in datastore after TokenExpansion")

	poolAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(burn_mint_token_pool.ContractType),
		Version:       burn_mint_token_pool.Version,
		Qualifier:     "TEST",
	}, chainSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "Token pool should exist in datastore after TokenExpansion")

	// TokenExpansion no longer deploys AdvancedPoolHooks. Deploy one separately and wire it into
	// the pool so tests can continue to exercise hook-dependent configuration.
	chain := e.BlockChains.EVMChains()[chainSel]
	hooksQualifier := "TEST"
	hooksRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, advanced_pool_hooks.Deploy, chain, contract_utils.DeployInput[advanced_pool_hooks.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *advanced_pool_hooks.Version),
		Args: advanced_pool_hooks.ConstructorArgs{
			ThresholdAmountForAdditionalCCVs: thresholdAmountForAdditionalCCVs,
		},
		Qualifier: &hooksQualifier,
	}, nil)
	require.NoError(t, err, "Failed to deploy AdvancedPoolHooks")

	hooksAddr := common.HexToAddress(hooksRef.Address)

	postDS := datastore.NewMemoryDataStore()
	postDS.Merge(e.DataStore)
	require.NoError(t, postDS.Addresses().Add(hooksRef), "Failed to add AdvancedPoolHooks to datastore")
	e.DataStore = postDS.Seal()

	poolContract, err := latest_token_pool.NewTokenPool(poolAddr, chain.Client)
	require.NoError(t, err, "Failed to bind token pool contract")
	auth := *chain.DeployerKey
	auth.Nonce = nil
	tx, err := poolContract.UpdateAdvancedPoolHooks(&auth, hooksAddr)
	require.NoError(t, err, "Failed to wire AdvancedPoolHooks into the pool")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm AdvancedPoolHooks wiring")

	return tokenExpansionResult{
		TokenAddress:         tokenAddr,
		TokenPoolAddress:     poolAddr,
		AdvancedHooksAddress: hooksAddr,
	}
}
