package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	erc20_lock_box_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"

	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestERC20LockBoxDeployChangeset(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)

	// Add a TokenAdminRegistry contract reference to the datastore 	so that the changeset can find it
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          "TokenAdminRegistry",
		Version:       semver.MustParse("1.5.0"),
		Address:       common.Address{2}.Hex(),
		ChainSelector: chainSelector,
	}))

	e.DataStore = ds.Seal()
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	changesetInput := changesets.ERC20LockboxDeployInput{
		ChainInputs: []changesets.ERC20LockboxDeployInputPerChain{
			{
				ChainSelector: chainSelector,
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
	require.NotNil(t, erc20Lockbox, "ERC20LockBox should not be nil")
}
