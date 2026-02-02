package cctp

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	erc20_lock_box_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"
	siloed_usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_usdc_token_pool"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"
)

type cctpTestSetup struct {
	Router             common.Address
	RMN                common.Address
	TokenAdminRegistry common.Address
	USDCToken          common.Address
	TokenMessenger     common.Address
	MessageTransmitter common.Address
}

func setupCCTPTestEnvironment(t *testing.T, e *deployment.Environment, chainSelector uint64) cctpTestSetup {
	chain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy chain contracts
	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, chain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		ChainSelector:  chainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{chain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")
	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:  chainSelector,
			ContractParams: testsetup.CreateBasicContractParams(),
			CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
		},
	)
	require.NoError(t, err, "Failed to deploy chain contracts")

	var routerAddr, rmnAddr, tokenAdminRegistryAddr common.Address
	for _, addr := range chainReport.Output.Addresses {
		if addr.Type == datastore.ContractType(router.ContractType) {
			routerAddr = common.HexToAddress(addr.Address)
		}
		if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
			rmnAddr = common.HexToAddress(addr.Address)
		}
		if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
			tokenAdminRegistryAddr = common.HexToAddress(addr.Address)
		}
	}
	require.NotEqual(t, common.Address{}, routerAddr, "Router address should be set")
	require.NotEqual(t, common.Address{}, rmnAddr, "RMN address should be set")
	require.NotEqual(t, common.Address{}, tokenAdminRegistryAddr, "TokenAdminRegistry address should be set")

	// Deploy USDC token (BurnMintERC20)
	usdcTokenAddr, tx, _, err := burn_mint_erc20_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"USD Coin",
		"USDC",
		6,             // decimals
		big.NewInt(0), // maxSupply
		big.NewInt(0), // pre-mint amount
	)
	require.NoError(t, err, "Failed to deploy USDC token")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm USDC token deployment")

	// Deploy MockE2EUSDCTransmitter
	messageTransmitterAddr, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		chain.DeployerKey,
		chain.Client,
		uint32(1),     // version (CCTP V2)
		uint32(1),     // localDomain
		usdcTokenAddr, // token
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTransmitter")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTransmitter deployment")

	// Deploy MockE2EUSDCTokenMessenger
	tokenMessengerAddr, tx, _, err := mock_usdc_token_messenger.DeployMockE2EUSDCTokenMessenger(
		chain.DeployerKey,
		chain.Client,
		uint32(1),              // version (CCTP V2)
		messageTransmitterAddr, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTokenMessenger")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTokenMessenger deployment")

	return cctpTestSetup{
		Router:             routerAddr,
		RMN:                rmnAddr,
		TokenAdminRegistry: tokenAdminRegistryAddr,
		USDCToken:          usdcTokenAddr,
		TokenMessenger:     tokenMessengerAddr,
		MessageTransmitter: messageTransmitterAddr,
	}
}
func TestDeploySiloedUSDCLockRelease(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	setup := setupCCTPTestEnvironment(t, e, chainSelector)
	chain := e.BlockChains.EVMChains()[chainSelector]

	lockReleaseSelectors := []uint64{4949039107694359620, 6433500567565415381}
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		DeploySiloedUSDCLockRelease,
		e.BlockChains,
		DeploySiloedUSDCLockReleaseInput{
			ChainSelector:             chainSelector,
			USDCToken:                 setup.USDCToken.Hex(),
			Router:                    setup.Router.Hex(),
			RMN:                       setup.RMN.Hex(),
			SiloedUSDCTokenPool:       "",
			LockReleaseChainSelectors: lockReleaseSelectors,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")
	require.NotEmpty(t, report.Output.SiloedPoolAddress, "SiloedUSDCTokenPool address should be set")
	require.Len(t, report.Output.LockBoxes, len(lockReleaseSelectors), "Expected lockboxes for each lock-release chain")

	siloedPoolAddr := common.HexToAddress(report.Output.SiloedPoolAddress)
	pool, err := siloed_usdc_token_pool_bindings.NewSiloedUSDCTokenPool(siloedPoolAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate SiloedUSDCTokenPool contract")

	poolCallers, err := pool.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get authorized callers from SiloedUSDCTokenPool")
	require.Empty(t, poolCallers, "SiloedUSDCTokenPool should not have authorized callers before proxy wiring")

	for _, sel := range lockReleaseSelectors {
		lockBoxAddr, ok := report.Output.LockBoxes[sel]
		require.True(t, ok, "Lockbox should be recorded for chain %d", sel)
		lockBoxFromPool, err := pool.GetLockBox(nil, sel)
		require.NoError(t, err, "Failed to get lockbox from SiloedUSDCTokenPool")
		require.Equal(t, common.HexToAddress(lockBoxAddr), lockBoxFromPool, "Lockbox address should match pool config")

		lockBox, err := erc20_lock_box_bindings.NewERC20LockBox(common.HexToAddress(lockBoxAddr), chain.Client)
		require.NoError(t, err, "Failed to instantiate ERC20LockBox contract")
		lockBoxCallers, err := lockBox.GetAllAuthorizedCallers(nil)
		require.NoError(t, err, "Failed to get authorized callers from ERC20LockBox")
		require.Contains(t, lockBoxCallers, siloedPoolAddr, "SiloedUSDCTokenPool should be an authorized caller on lockbox")
	}
}
