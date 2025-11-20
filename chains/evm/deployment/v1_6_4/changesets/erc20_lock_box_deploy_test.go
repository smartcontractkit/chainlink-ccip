package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	erc20_lock_box_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func TestERC20LockBoxDeployChangeset(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)

	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	tokenAdminRegistryAddress := common.Address{1}

	changesetInput := changesets.ERC20LockboxDeployInput{
		ChainInputs: []changesets.ERC20LockboxDeployInputPerChain{
			{
				ChainSelector:      chainSelector,
				TokenAdminRegistry: tokenAdminRegistryAddress,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy ERC20Lockbox",
		},
	}

	mcmsRegistry := changesets_utils.GetRegistry()
	changesetOutput, err := changesets.ERC20LockboxDeployChangeset(mcmsRegistry).Apply(*e, changesetInput)
	require.NoError(t, err, "Failed to apply ERC20LockboxDeployChangeset")
	require.NotNil(t, changesetOutput, "Changeset output should not be nil")
	require.Greater(t, len(changesetOutput.Reports), 0)

	erc20LockboxAddress, err := changesetOutput.DataStore.Addresses().Get(datastore.NewAddressRefKey(chainSelector, "ERC20Lockbox", semver.MustParse("1.6.4"), ""))
	require.NoError(t, err, "Failed to get ERC20Lockbox address")

	erc20Lockbox, err := erc20_lock_box_bindings.NewERC20LockBox(common.HexToAddress(erc20LockboxAddress.Address), evmChain.Client)
	require.NoError(t, err, "Failed to create ERC20LockBox")
	tokenAdminRegistry, err := erc20Lockbox.ITokenAdminRegistry(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to get token admin registry")
	require.Equal(t, tokenAdminRegistryAddress, tokenAdminRegistry, "Expected token admin registry address to be in ERC20LockBox")
}
