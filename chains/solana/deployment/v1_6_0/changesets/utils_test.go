package changesets

import (
	"encoding/json"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	solstate "github.com/smartcontractkit/cld-changesets/legacy/pkg/family/solana"
	mcmsSolana "github.com/smartcontractkit/mcms/sdk/solana"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

func TestMcmSeedForAction(t *testing.T) {
	t.Parallel()
	state := &solstate.MCMSWithTimelockState{MCMSWithTimelockPrograms: &solstate.MCMSWithTimelockPrograms{
		ProposerMcmSeed:  solstate.PDASeed{1},
		CancellerMcmSeed: solstate.PDASeed{2},
		BypasserMcmSeed:  solstate.PDASeed{3},
	}}
	require.Equal(t, mcmsSolana.PDASeed{1}, mcmSeedForAction(state, ""))
	require.Equal(t, mcmsSolana.PDASeed{1}, mcmSeedForAction(state, mcmsTypes.TimelockActionSchedule))
	require.Equal(t, mcmsSolana.PDASeed{2}, mcmSeedForAction(state, mcmsTypes.TimelockActionCancel))
	require.Equal(t, mcmsSolana.PDASeed{3}, mcmSeedForAction(state, mcmsTypes.TimelockActionBypass))
}

func TestAppendBatchOperation(t *testing.T) {
	t.Parallel()
	require.Empty(t, appendBatchOperation(nil, 1, nil))

	txns := []mcmsTypes.Transaction{{To: "a"}, {To: "b"}}
	batches := appendBatchOperation(nil, 7, txns)
	require.Equal(t, []mcmsTypes.BatchOperation{{ChainSelector: 7, Transactions: txns}}, batches)
}

func TestBuildMCMSTxn(t *testing.T) {
	t.Parallel()
	programID := solana.NewWallet().PublicKey()
	signer := solana.NewWallet().PublicKey()
	other := solana.NewWallet().PublicKey()
	ixn := solana.NewInstruction(programID, solana.AccountMetaSlice{
		solana.Meta(signer).WRITE().SIGNER(),
		solana.Meta(other),
	}, []byte{1, 2, 3})

	tx, err := BuildMCMSTxn(ixn, programID.String(), shared.Router)
	require.NoError(t, err)
	require.Equal(t, programID.String(), tx.To)
	require.Equal(t, []byte{1, 2, 3}, tx.Data)
	require.Equal(t, string(shared.Router), tx.ContractType)

	var fields mcmsSolana.AdditionalFields
	require.NoError(t, json.Unmarshal(tx.AdditionalFields, &fields))
	require.Len(t, fields.Accounts, 2)
	for _, account := range fields.Accounts {
		require.False(t, account.IsSigner, "the timelock signs, so no account may stay a signer")
	}
	require.True(t, fields.Accounts[0].IsWritable)

	_, err = BuildMCMSTxn(ixn, "not-a-key", shared.Router)
	require.Error(t, err)
}

func TestBuildManyMCMSTxsFrom(t *testing.T) {
	t.Parallel()
	programID := solana.NewWallet().PublicKey()
	ixn := solana.NewInstruction(programID, solana.AccountMetaSlice{}, []byte{9})

	txs, err := BuildManyMCMSTxsFrom([]MCMSTxParams{
		{Ix: ixn, ProgramID: programID.String(), ContractType: shared.Router},
		{Ix: ixn, ProgramID: programID.String(), ContractType: shared.OffRamp},
	})
	require.NoError(t, err)
	require.Len(t, txs, 2)
	require.Equal(t, string(shared.OffRamp), txs[1].ContractType)

	_, err = BuildManyMCMSTxsFrom([]MCMSTxParams{{Ix: ixn, ProgramID: "bad", ContractType: shared.Router}})
	require.Error(t, err)
}

func TestGetTokenProgramID(t *testing.T) {
	t.Parallel()
	id, err := GetTokenProgramID(shared.SPLTokens)
	require.NoError(t, err)
	require.Equal(t, solana.TokenProgramID, id)

	id, err = GetTokenProgramID(shared.SPL2022Tokens)
	require.NoError(t, err)
	require.Equal(t, solana.Token2022ProgramID, id)

	_, err = GetTokenProgramID(shared.Router)
	require.Error(t, err)
}

func TestFetchTimelockSignerAndAuthorityFallback(t *testing.T) {
	t.Parallel()
	const chainSel = uint64(16423721717087811551)
	timelockProgram := solana.NewWallet().PublicKey()
	seed := mcmsSolana.PDASeed{'s', 'e', 'e', 'd'}
	deployer := solana.NewWallet()
	chain := cldf_solana.Chain{Selector: chainSel, DeployerKey: &deployer.PrivateKey}
	chains := cldf_chain.NewBlockChainsFromSlice([]cldf_chain.BlockChain{chain})

	timelockRef := datastore.AddressRef{
		Address:       mcmsSolana.ContractAddress(timelockProgram, seed),
		ChainSelector: chainSel,
		Type:          datastore.ContractType("RBACTimelock"),
		Version:       semver.MustParse("1.6.0"),
		Qualifier:     shared.DefaultMCMSQualifier,
	}

	// no MCMS in the datastore: the authority falls back to the deployer key, even when the
	// deprecated address book still lists a timelock
	ab := cldf.NewMemoryAddressBook()
	require.NoError(t, ab.Save(chainSel, timelockRef.Address, cldf.NewTypeAndVersion("RBACTimelock", *semver.MustParse("1.0.0"))))
	empty := cldf.Environment{ExistingAddresses: ab, DataStore: datastore.NewMemoryDataStore().Seal(), BlockChains: chains}
	_, err := FetchTimelockSigner(empty, chainSel)
	require.Error(t, err)
	require.Equal(t, deployer.PublicKey(), GetAuthorityForIxn(&empty, chain, solanastateview.CCIPChainState{}, shared.Router, solana.PublicKey{}, ""))

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(timelockRef))
	e := cldf.Environment{DataStore: ds.Seal(), BlockChains: chains}
	signer, err := FetchTimelockSigner(e, chainSel)
	require.NoError(t, err)
	require.Equal(t, state.GetTimelockSignerPDA(timelockProgram, state.PDASeed(seed)), signer)
}

func TestValidateMCMSConfigSolanaRejectsInvalidAction(t *testing.T) {
	t.Parallel()
	chain := cldf_solana.Chain{Selector: 1}
	require.NoError(t, ValidateMCMSConfigSolana(cldf.Environment{}, nil, chain, solanastateview.CCIPChainState{}, solana.PublicKey{}, "", nil))

	err := ValidateMCMSConfigSolana(cldf.Environment{}, &cldfproposalutils.TimelockConfig{MCMSAction: "nope"}, chain,
		solanastateview.CCIPChainState{}, solana.PublicKey{}, "", nil)
	require.ErrorContains(t, err, "invalid MCMS action")
}

func TestMcmsQualifier(t *testing.T) {
	t.Parallel()
	require.Equal(t, shared.DefaultMCMSQualifier, mcmsQualifier(cldfproposalutils.TimelockConfig{}, 1))
	cfg := cldfproposalutils.TimelockConfig{TimelockQualifierPerChain: map[uint64]string{1: "RMNMCMS", 2: ""}}
	require.Equal(t, "RMNMCMS", mcmsQualifier(cfg, 1))
	require.Equal(t, shared.DefaultMCMSQualifier, mcmsQualifier(cfg, 2))
	require.Equal(t, shared.DefaultMCMSQualifier, mcmsQualifier(cfg, 3))
}
