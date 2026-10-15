package token_pool

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	drip_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"

	bmtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/operations/burn_mint_token_pool"
	lrtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/operations/lock_release_token_pool"
	bmbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/burn_mint_token_pool"
	lrbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/lock_release_token_pool"
)

var testChainSelector = chain_selectors.ETHEREUM_MAINNET.Selector

const testTokenSymbol = "TEST"

// Stand-in addresses for the router and ARM proxy. The configure sequence never
// calls either, so they only need to be non-zero and stable.
var (
	testRouterAddress   = common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	testArmProxyAddress = common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
)

// setupPoolEnv builds a simulated chain with a deployed BurnMintERC20WithDrip
// token, and returns the environment alongside the token address.
//
// Unlike the v1.5.x equivalent this seeds no datastore refs: the v1.4.0 adapter
// deploys nothing, so the pools below are deployed straight from the generated
// bindings and no router / RMN proxy lookup ever happens.
func setupPoolEnv(t *testing.T) (*cldf.Environment, common.Address) {
	t.Helper()

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{testChainSelector}))
	require.NoError(t, err, "Failed to create test environment")

	chain := e.BlockChains.EVMChains()[testChainSelector]

	return e, deployTestDripToken(t, chain, testTokenSymbol)
}

func deployTestDripToken(t *testing.T, chain evm.Chain, symbol string) common.Address {
	t.Helper()

	tokenAddr, tx, _, err := drip_bindings.DeployBurnMintERC20WithDrip(
		chain.DeployerKey,
		chain.Client,
		"Test Drip Token",
		symbol,
	)
	require.NoError(t, err, "Failed to deploy test drip token")

	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm test drip token deployment")

	return tokenAddr
}

// deployConfigurablePool deploys a v1.4.0 BurnMintTokenPool to configure against.
func deployConfigurablePool(t *testing.T) (*cldf.Environment, common.Address) {
	t.Helper()

	return deployConfigurablePoolOfType(t, string(bmtp.ContractType))
}

// deployConfigurablePoolOfType deploys a v1.4.0 pool of the given type directly
// from the generated bindings.
func deployConfigurablePoolOfType(t *testing.T, poolType string) (*cldf.Environment, common.Address) {
	t.Helper()

	e, tokenAddr := setupPoolEnv(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]

	var (
		poolAddr common.Address
		tx       *types.Transaction
		err      error
	)

	switch poolType {
	case string(bmtp.ContractType):
		poolAddr, tx, _, err = bmbindings.DeployBurnMintTokenPool(
			chain.DeployerKey, chain.Client,
			tokenAddr, []common.Address{}, testArmProxyAddress, testRouterAddress,
		)
	case string(lrtp.ContractType):
		poolAddr, tx, _, err = lrbindings.DeployLockReleaseTokenPool(
			chain.DeployerKey, chain.Client,
			tokenAddr, []common.Address{}, testArmProxyAddress, true, testRouterAddress,
		)
	default:
		t.Fatalf("unsupported v1.4.0 pool type %q", poolType)
	}
	require.NoError(t, err, "Failed to deploy v1.4.0 %s", poolType)

	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm v1.4.0 %s deployment", poolType)

	return e, poolAddr
}
