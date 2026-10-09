package changesets

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	mcmsSolana "github.com/smartcontractkit/mcms/sdk/solana"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

func testRef(chainSel uint64, address string, contractType string, version string, qualifier string, labels ...string) datastore.AddressRef {
	ref := datastore.AddressRef{
		ChainSelector: chainSel,
		Address:       address,
		Type:          datastore.ContractType(contractType),
		Version:       semver.MustParse(version),
		Qualifier:     qualifier,
	}
	if len(labels) > 0 {
		ref.Labels = datastore.NewLabelSet(labels...)
	}
	return ref
}

func sealedDataStore(t *testing.T, refs ...datastore.AddressRef) datastore.DataStore {
	t.Helper()
	ds := datastore.NewMemoryDataStore()
	for _, ref := range refs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	return ds.Seal()
}

func TestValidateHomeChainRefs(t *testing.T) {
	t.Parallel()
	const home = uint64(5009297550715157269)
	complete := []datastore.AddressRef{
		testRef(home, "0x1", "CapabilitiesRegistry", "1.0.0", ""),
		testRef(home, "0x2", "CCIPHome", "1.6.0", ""),
		testRef(home, "0x3", "RMNHome", "1.6.0", ""),
	}
	require.NoError(t, validateHomeChainRefs(cldf.Environment{DataStore: sealedDataStore(t, complete...)}, home))

	// another chain's refs don't count
	require.ErrorContains(t, validateHomeChainRefs(cldf.Environment{DataStore: sealedDataStore(t, complete...)}, 1), "CapabilitiesRegistry not found")

	// a superseded ref doesn't count
	superseded := append(complete[:2:2], testRef(home, "0x3", "RMNHome", "1.6.0", "", shared.SupersededLabel))
	require.ErrorContains(t, validateHomeChainRefs(cldf.Environment{DataStore: sealedDataStore(t, superseded...)}, home), "RMNHome not found")

	require.Error(t, validateHomeChainRefs(cldf.Environment{}, home))
}

func TestFindOffRampAtVersion(t *testing.T) {
	t.Parallel()
	const sel = uint64(124615329519749607)
	current := solana.NewWallet().PublicKey()
	redeployed := solana.NewWallet().PublicKey()
	e := cldf.Environment{DataStore: sealedDataStore(t,
		testRef(sel, current.String(), string(shared.OffRamp), "1.6.0", ""),
		testRef(sel, redeployed.String(), string(shared.OffRamp), "1.6.4", ""),
		testRef(sel, solana.NewWallet().PublicKey().String(), string(shared.OffRamp), "1.6.9", "", shared.SupersededLabel),
	)}

	got, err := findOffRampAtVersion(e, sel, semver.MustParse("1.6.4"))
	require.NoError(t, err)
	require.Equal(t, redeployed, got)

	// no offramp at that exact version: the caller redeploys
	got, err = findOffRampAtVersion(e, sel, semver.MustParse("1.6.5"))
	require.NoError(t, err)
	require.True(t, got.IsZero())

	// superseded refs are ignored
	got, err = findOffRampAtVersion(e, sel, semver.MustParse("1.6.9"))
	require.NoError(t, err)
	require.True(t, got.IsZero())
}

func TestLoadMCMSStateIfDeployed(t *testing.T) {
	t.Parallel()
	const sel = uint64(124615329519749607)
	timelockProgram := solana.NewWallet().PublicKey()
	seed := mcmsSolana.PDASeed{'t', 'l'}

	// no MCMS on the chain: an empty state, not an error
	state, err := loadMCMSStateIfDeployed(cldf.Environment{DataStore: sealedDataStore(t)}, sel)
	require.NoError(t, err)
	require.True(t, state.TimelockProgram.IsZero())

	e := cldf.Environment{DataStore: sealedDataStore(t,
		testRef(sel, mcmsSolana.ContractAddress(timelockProgram, seed), "RBACTimelock", "1.6.0", shared.DefaultMCMSQualifier),
		// a second bundle under its own qualifier is not picked up
		testRef(sel, mcmsSolana.ContractAddress(solana.NewWallet().PublicKey(), seed), "RBACTimelock", "1.6.0", "RMNMCMS"),
	)}
	state, err = loadMCMSStateIfDeployed(e, sel)
	require.NoError(t, err)
	require.Equal(t, timelockProgram, state.TimelockProgram)
	require.Equal(t, seed[:], state.TimelockSeed[:])

	_, err = loadMCMSStateIfDeployed(cldf.Environment{}, sel)
	require.Error(t, err)
}

func TestDeployChainContractsConfigRejectsMCMSDeploy(t *testing.T) {
	t.Parallel()
	sel := chainsel.SOLANA_MAINNET.Selector
	state := solanastateview.CCIPOnChainState{SolChains: map[uint64]solanastateview.CCIPChainState{
		sel: {Router: solana.NewWallet().PublicKey()},
	}}
	cfg := DeployChainContractsConfig{
		HomeChainSelector:      chainsel.ETHEREUM_MAINNET.Selector,
		ChainSelector:          sel,
		MCMSWithTimelockConfig: &cldfproposalutils.MCMSWithTimelockConfig{},
	}
	require.ErrorContains(t, cfg.Validate(cldf.Environment{}, state), "MCMSWithTimelockConfig is no longer supported")

	cfg.MCMSWithTimelockConfig = nil
	require.NoError(t, cfg.Validate(cldf.Environment{}, state))
}
