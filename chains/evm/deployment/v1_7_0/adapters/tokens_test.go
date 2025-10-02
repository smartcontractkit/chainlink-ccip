package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	evm_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func init() {
	tokens.RegisterTokenAdapter("evm", semver.MustParse("1.7.0"), &adapters.TokenAdapter{})
}

func TestTokenAdapter(t *testing.T) {
	chainA := uint64(5009297550715157269)
	chainB := uint64(4949039107694359620)
	e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
		chainA: {NumAdditionalAccounts: 1},
		chainB: {NumAdditionalAccounts: 1},
	})
	require.NoError(t, err, "Failed to create test environment with 2 chains")

	// On each chain, deploy chain contracts & a token + token pool
	for _, chainSel := range []uint64{chainA, chainB} {
		deployChainOut, err := changesets.DeployChainContracts.Apply(e, changesets.DeployChainContractsCfg{
			ChainSel: chainSel,
			Params:   testsetup.CreateBasicContractParams(),
		})
		require.NoError(t, err, "Failed to apply DeployChainContracts changeset")

		refs, err := datastore_utils.FindAndFormatEachRef(deployChainOut.DataStore.Seal(), []datastore.AddressRef{
			{ChainSelector: chainSel, Type: datastore.ContractType(router.ContractType)},
			{ChainSelector: chainSel, Type: datastore.ContractType(rmn_proxy.ContractType)},
		}, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Failed to find deployed contract refs in datastore after DeployChainContracts changeset")
		router := refs[0]
		rmnProxy := refs[1]

		_, err = changesets.DeployBurnMintTokenAndPool.Apply(e, evm_tokens.DeployBurnMintTokenAndPoolInput{
			Accounts: map[common.Address]*big.Int{
				e.BlockChains.EVMChains()[chainSel].DeployerKey.From: big.NewInt(1_000_000),
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
		})
		require.NoError(t, err, "Failed to apply DeployBurnMintTokenAndPool changeset")
	}
}
