package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	onramp_changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/changesets"
	utils_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// MockReader implements mcms.Reader for test purposes
type MockReader struct{}

func (m *MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{StartingOpCount: 1}, nil
}

func (m *MockReader) GetTimelockRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0x0000000000000000000000000000000000000001",
		Type:          datastore.ContractType("Timelock"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func (m *MockReader) GetMCMSRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0x0000000000000000000000000000000000000002",
		Type:          datastore.ContractType("MCM"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func init() {
	registry := utils_changesets.GetRegistry()
	registry.RegisterMCMSReader("evm", &MockReader{})
}

// =============================================================================
// ConfigureFeeSweep Tests
// =============================================================================

func TestConfigureFeeSweep_Validation(t *testing.T) {
	t.Parallel()

	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("Timelock"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000001",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("MCM"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000002",
	})
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	t.Run("missing chain in allowed recipients", func(t *testing.T) {
		allowedRecipients := map[uint64]common.Address{
			12345: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		}

		registry := utils_changesets.GetRegistry()
		cs := onramp_changesets.ConfigureFeeSweepV150(allowedRecipients)(registry)

		input := utils_changesets.WithMCMS[onramp_changesets.ConfigureFeeSweepV150Cfg]{
			Cfg: onramp_changesets.ConfigureFeeSweepV150Cfg{
				ChainSel:      chainSel,
				OnRampAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
				Treasury:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: true, ValidUntil: 4126214326,
				TimelockDelay:  mcms_types.MustParseDuration("1h"),
				TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
			},
		}

		err := cs.VerifyPreconditions(*e, input)
		require.NoError(t, err)

		_, err = cs.Apply(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not in the allowed recipients list")
	})

	t.Run("zero treasury address", func(t *testing.T) {
		allowedRecipients := map[uint64]common.Address{
			chainSel: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		}

		registry := utils_changesets.GetRegistry()
		cs := onramp_changesets.ConfigureFeeSweepV150(allowedRecipients)(registry)

		input := utils_changesets.WithMCMS[onramp_changesets.ConfigureFeeSweepV150Cfg]{
			Cfg: onramp_changesets.ConfigureFeeSweepV150Cfg{
				ChainSel:      chainSel,
				OnRampAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
				Treasury:      common.Address{},
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: true, ValidUntil: 4126214326,
				TimelockDelay:  mcms_types.MustParseDuration("1h"),
				TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
			},
		}

		err := cs.VerifyPreconditions(*e, input)
		require.NoError(t, err)

		_, err = cs.Apply(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "zero address")
	})
}

// =============================================================================
// SweepLinkFees Tests
// =============================================================================

func TestSweepLinkFees_Validation(t *testing.T) {
	t.Parallel()

	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("Timelock"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000001",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("MCM"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000002",
	})
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	t.Run("treasury mismatch with allowed list", func(t *testing.T) {
		allowedRecipients := map[uint64]common.Address{
			chainSel: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		}

		registry := utils_changesets.GetRegistry()
		cs := onramp_changesets.SweepLinkFeesV150(allowedRecipients)(registry)

		input := utils_changesets.WithMCMS[onramp_changesets.SweepLinkFeesV150Cfg]{
			Cfg: onramp_changesets.SweepLinkFeesV150Cfg{
				ChainSel:         chainSel,
				OnRampAddress:    common.HexToAddress("0x3333333333333333333333333333333333333333"),
				ExpectedTreasury: common.HexToAddress("0x5555555555555555555555555555555555555555"),
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: true, ValidUntil: 4126214326,
				TimelockDelay:  mcms_types.MustParseDuration("1h"),
				TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
			},
		}

		err := cs.VerifyPreconditions(*e, input)
		require.NoError(t, err)

		_, err = cs.Apply(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not the approved treasury")
	})
}

// =============================================================================
// SweepNonLinkFees Tests
// =============================================================================

func TestSweepNonLinkFees_Validation(t *testing.T) {
	t.Parallel()

	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("Timelock"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000001",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("MCM"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000002",
	})
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	t.Run("empty fee tokens list is valid", func(t *testing.T) {
		allowedRecipients := map[uint64]common.Address{
			chainSel: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		}

		registry := utils_changesets.GetRegistry()
		cs := onramp_changesets.SweepNonLinkFeesV150(allowedRecipients)(registry)

		input := utils_changesets.WithMCMS[onramp_changesets.SweepNonLinkFeesV150Cfg]{
			Cfg: onramp_changesets.SweepNonLinkFeesV150Cfg{
				ChainSel:      chainSel,
				OnRampAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
				FeeTokens:     []common.Address{},
				Treasury:      allowedRecipients[chainSel],
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: true, ValidUntil: 4126214326,
				TimelockDelay:  mcms_types.MustParseDuration("1h"),
				TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
			},
		}

		err := cs.VerifyPreconditions(*e, input)
		require.NoError(t, err)
	})

	t.Run("treasury validation failure", func(t *testing.T) {
		allowedRecipients := map[uint64]common.Address{
			chainSel: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		}

		registry := utils_changesets.GetRegistry()
		cs := onramp_changesets.SweepNonLinkFeesV150(allowedRecipients)(registry)

		input := utils_changesets.WithMCMS[onramp_changesets.SweepNonLinkFeesV150Cfg]{
			Cfg: onramp_changesets.SweepNonLinkFeesV150Cfg{
				ChainSel:      chainSel,
				OnRampAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
				FeeTokens:     []common.Address{common.HexToAddress("0xAAAA")},
				Treasury:      common.HexToAddress("0x9999999999999999999999999999999999999999"),
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: true, ValidUntil: 4126214326,
				TimelockDelay:  mcms_types.MustParseDuration("1h"),
				TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
			},
		}

		err := cs.VerifyPreconditions(*e, input)
		require.NoError(t, err)

		_, err = cs.Apply(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not the approved treasury")
	})
}

// =============================================================================
// SweepAllOnRampsV150 (Mega Flow) Tests
// =============================================================================

func TestSweepAllOnRampsV150_NoFeeQuoterInDatastore(t *testing.T) {
	t.Parallel()

	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("WETH9"),
		Version: semver.MustParse("1.0.0"), Address: "0x4200000000000000000000000000000000000006",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("RBACTimelock"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000099",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("Timelock"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000001",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel, Type: datastore.ContractType("MCM"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000002",
	})
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	allowedRecipients := map[uint64]common.Address{
		chainSel: common.HexToAddress("0x1111111111111111111111111111111111111111"),
	}

	registry := utils_changesets.GetRegistry()
	cs := onramp_changesets.SweepAllOnRampsV150(allowedRecipients)(registry)

	input := utils_changesets.WithMCMS[onramp_changesets.SweepAllOnRampsV150Cfg]{
		Cfg: onramp_changesets.SweepAllOnRampsV150Cfg{
			ChainSel: chainSel,
			Treasury: allowedRecipients[chainSel],
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: true, ValidUntil: 4126214326,
			TimelockDelay:  mcms_types.MustParseDuration("1h"),
			TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
		},
	}

	// FeeQuoter resolution now happens before OnRamp discovery, so with no
	// FeeQuoter in the datastore the changeset fails at that point.
	err = cs.VerifyPreconditions(*e, input)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no FeeQuoter found")
}

// =============================================================================
// Chain Selector Validation Tests
// =============================================================================

func TestChangesets_NonexistentChain(t *testing.T) {
	t.Parallel()

	validChainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{validChainSel}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: validChainSel, Type: datastore.ContractType("Timelock"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000001",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: validChainSel, Type: datastore.ContractType("MCM"),
		Version: semver.MustParse("1.0.0"), Address: "0x0000000000000000000000000000000000000002",
	})
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	nonexistentChainSel := uint64(999999)

	t.Run("ConfigureFeeSweep with nonexistent chain", func(t *testing.T) {
		allowedRecipients := map[uint64]common.Address{
			nonexistentChainSel: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		}

		registry := utils_changesets.GetRegistry()
		cs := onramp_changesets.ConfigureFeeSweepV150(allowedRecipients)(registry)

		input := utils_changesets.WithMCMS[onramp_changesets.ConfigureFeeSweepV150Cfg]{
			Cfg: onramp_changesets.ConfigureFeeSweepV150Cfg{
				ChainSel:      nonexistentChainSel,
				OnRampAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
				Treasury:      allowedRecipients[nonexistentChainSel],
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: true, ValidUntil: 4126214326,
				TimelockDelay:  mcms_types.MustParseDuration("1h"),
				TimelockAction: mcms_types.TimelockActionSchedule, Description: "Test",
			},
		}

		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
	})
}
