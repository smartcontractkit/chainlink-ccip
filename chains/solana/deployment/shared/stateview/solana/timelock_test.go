package solana

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	mcmsolana "github.com/smartcontractkit/mcms/sdk/solana"
	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcmscontracts "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/contracts/mcms"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

func testRef(chainSel uint64, addr string, t cldf.ContractType, ver *semver.Version, qualifier string, labels ...string) datastore.AddressRef {
	ref := datastore.AddressRef{
		ChainSelector: chainSel,
		Address:       addr,
		Type:          datastore.ContractType(t),
		Qualifier:     qualifier,
	}
	if ver != nil {
		v := *ver
		ref.Version = &v
	}
	if len(labels) > 0 {
		ref.Labels = datastore.NewLabelSet(labels...)
	}
	return ref
}

func TestValidateSolanaTimelockConfig(t *testing.T) {
	program := solana.MustPublicKeyFromBase58("9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM")
	valid := mcmsolana.ContractAddress(program, mcmsolana.PDASeed{1, 2, 3})
	v16 := shared.Version1_6_0

	// configured qualifier resolves its own bundle, ignoring the default-qualified decoy
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(testRef(1, valid, mcmscontracts.RBACTimelock, &v16, "RMNMCMS")))
	require.NoError(t, ds.Addresses().Add(testRef(1, valid, mcmscontracts.ProposerManyChainMultisig, &v16, "RMNMCMS")))
	require.NoError(t, ds.Addresses().Add(testRef(1, "not-a-program-address", mcmscontracts.RBACTimelock, &v16, shared.DefaultMCMSQualifier)))
	tc := &cldfproposalutils.TimelockConfig{TimelockQualifierPerChain: map[uint64]string{1: "RMNMCMS"}}
	require.NoError(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: ds.Seal()}, 1, tc))

	// default qualifier, action-specific contract; empty action defaults in place
	// (framework validateCommon contract)
	ds2 := datastore.NewMemoryDataStore()
	require.NoError(t, ds2.Addresses().Add(testRef(1, valid, mcmscontracts.RBACTimelock, &v16, shared.DefaultMCMSQualifier)))
	require.NoError(t, ds2.Addresses().Add(testRef(1, valid, mcmscontracts.BypasserManyChainMultisig, &v16, shared.DefaultMCMSQualifier)))
	require.NoError(t, ds2.Addresses().Add(testRef(1, valid, mcmscontracts.ProposerManyChainMultisig, &v16, shared.DefaultMCMSQualifier)))
	tc2 := &cldfproposalutils.TimelockConfig{MCMSAction: mcmstypes.TimelockActionBypass}
	require.NoError(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: ds2.Seal()}, 1, tc2))

	tcDefault := &cldfproposalutils.TimelockConfig{}
	require.NoError(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: ds2.Seal()}, 1, tcDefault))
	require.Equal(t, mcmstypes.TimelockActionSchedule, tcDefault.MCMSAction)

	// The default bundle may be stored without a qualifier by legacy deployments.
	dsEmptyDefault := datastore.NewMemoryDataStore()
	require.NoError(t, dsEmptyDefault.Addresses().Add(testRef(1, valid, mcmscontracts.RBACTimelock, &v16, "")))
	require.NoError(t, dsEmptyDefault.Addresses().Add(testRef(1, valid, mcmscontracts.ProposerManyChainMultisig, &v16, "")))
	require.NoError(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: dsEmptyDefault.Seal()}, 1, &cldfproposalutils.TimelockConfig{}))

	// Custom qualifiers remain strict and do not fall back to the unqualified bundle.
	tcEmptyCustom := &cldfproposalutils.TimelockConfig{TimelockQualifierPerChain: map[uint64]string{1: "RMNMCMS"}}
	require.Error(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: dsEmptyDefault.Seal()}, 1, tcEmptyCustom))

	// missing contract
	_, err := dataStoreSolanaContractAddress(cldf.Environment{DataStore: ds2.Seal()}, 1, mcmscontracts.CancellerManyChainMultisig, shared.DefaultMCMSQualifier)
	require.ErrorContains(t, err, "no CancellerManyChainMultiSig ref")

	// superseded refs are invisible
	ds3 := datastore.NewMemoryDataStore()
	require.NoError(t, ds3.Addresses().Add(testRef(1, valid, mcmscontracts.RBACTimelock, &v16, shared.DefaultMCMSQualifier, shared.SupersededLabel)))
	_, err = dataStoreSolanaContractAddress(cldf.Environment{DataStore: ds3.Seal()}, 1, mcmscontracts.RBACTimelock, shared.DefaultMCMSQualifier)
	require.ErrorContains(t, err, "no RBACTimelock ref")

	// nil config and invalid action
	require.Error(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: ds2.Seal()}, 1, nil))
	tcBad := &cldfproposalutils.TimelockConfig{MCMSAction: "nope"}
	require.ErrorContains(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: ds2.Seal()}, 1, tcBad), "invalid MCMS action")

	// configured qualifier that holds nothing errors (strict, no fallback)
	tc4 := &cldfproposalutils.TimelockConfig{TimelockQualifierPerChain: map[uint64]string{1: "OTHER"}}
	require.Error(t, ValidateSolanaTimelockConfig(cldf.Environment{DataStore: ds2.Seal()}, 1, tc4))
}
