package adapters_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnm_drip_v1_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	v2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	evm_tokens_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// TestTokenExpansion_SiloedLockReleaseTokenPool covers deploying a SiloedLockReleaseTokenPool through
// TokenExpansion - the entry point external repos use. Before lockBoxGroups existed the siloed type
// fell through to DeployLockReleaseTokenPool, whose bytecode map only holds LockReleaseTokenPool
// 2.0.0, so this failed with "no bytecode defined for SiloedLockReleaseTokenPool 2.0.0".
func TestTokenExpansion_SiloedLockReleaseTokenPool(t *testing.T) {
	const (
		symbol      = "TSILO"
		remoteChain = uint64(4949039107694359620)
	)
	chainSel := uint64(5009297550715157269)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	ds := datastore.NewMemoryDataStore()
	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSel,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err)

	deployChainOut, err := v2_0_0.DeployChainContracts(changesets.GetRegistry()).Apply(*e, changesets.WithMCMS[v2_0_0.DeployChainContractsCfg]{
		Cfg: v2_0_0.DeployChainContractsCfg{
			ChainSel:         chainSel,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			Params:           testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Merge(deployChainOut.DataStore.Seal()))
	e.DataStore = ds.Seal()

	deployer := e.BlockChains.EVMChains()[chainSel].DeployerKey.From
	preMint := uint64(100_000)

	output, err := tokens.TokenExpansion().Apply(*e, tokens.TokenExpansionInput{
		ChainAdapterVersion: semver.MustParse("2.0.0"),
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokens.TokenExpansionInputPerChain{
			chainSel: {
				TokenPoolVersion:      siloed_lock_release_token_pool.Version,
				SkipOwnershipTransfer: true,
				DeployTokenInput: &tokens.DeployTokenInput{
					Name:          "Test Token " + symbol,
					Symbol:        symbol,
					Decimals:      18,
					PreMint:       &preMint,
					ExternalAdmin: deployer.Hex(),
					CCIPAdmin:     deployer.Hex(),
					Type:          bnm_drip_v1_0.ContractType,
				},
				DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
					PoolType:           string(siloed_lock_release_token_pool.ContractType),
					TokenPoolQualifier: symbol,
					LockBoxGroups:      [][]uint64{{remoteChain}},
				},
			},
		},
	})
	require.NoError(t, err, "TokenExpansion should succeed for a siloed lock release pool")

	require.NoError(t, ds.Merge(output.DataStore.Seal()))
	e.DataStore = ds.Seal()

	poolAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(siloed_lock_release_token_pool.ContractType),
		Version:       siloed_lock_release_token_pool.Version,
		Qualifier:     symbol,
	}, chainSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "Siloed token pool should exist in datastore")

	// The lockbox is recorded under the group qualifier so re-runs and ownership transfer can find it.
	lockBoxAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(erc20_lock_box.ContractType),
		Version:       erc20_lock_box.Version,
		Qualifier:     evm_tokens_sequences.LockBoxQualifier(symbol, []uint64{remoteChain}),
	}, chainSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "Per-silo lock box should exist in datastore")

	poolContract, err := siloed_lock_release_token_pool.NewSiloedLockReleaseTokenPoolContract(
		poolAddr, e.BlockChains.EVMChains()[chainSel].Client,
	)
	require.NoError(t, err)

	onChainLockBox, err := poolContract.GetLockBox(&bind.CallOpts{Context: t.Context()}, remoteChain)
	require.NoError(t, err, "getLockBox should resolve for the configured remote chain")
	require.Equal(t, lockBoxAddr, onChainLockBox, "Expected the pool to point at the deployed lock box")
}
