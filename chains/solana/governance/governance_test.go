package governance

import (
	"encoding/json"
	"os/exec"
	"strings"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/stretchr/testify/require"

	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestBuildMCMSBatchOperation_Golden(t *testing.T) {
	t.Parallel()
	from := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	to := solana.MustPublicKeyFromBase58("11111111111111111111111111111112")
	ix, err := system.NewTransferInstruction(5, from, to).ValidateAndBuild()
	require.NoError(t, err)

	op, err := BuildMCMSBatchOperation(
		124615329519749607,
		[]solana.Instruction{ix, ix},
		"11111111111111111111111111111111",
		"Router",
	)
	require.NoError(t, err)
	got, err := json.Marshal(op)
	require.NoError(t, err)
	require.JSONEq(t, goldenBatchOperation, string(got))

	empty, err := BuildMCMSBatchOperation(1, nil, "x", "y")
	require.NoError(t, err)
	require.Equal(t, uint64(1), uint64(empty.ChainSelector))
	require.NotNil(t, empty.Transactions)
	require.Empty(t, empty.Transactions)
}

func TestGetAllMCMS_Golden(t *testing.T) {
	t.Parallel()
	types := []struct{ typ, addr, qualifier string }{
		{"AccessControllerProgram", "ac", ""},
		{"ProposerAccessControllerAccount", "p", "q"},
		{"ExecutorAccessControllerAccount", "e", "q"},
		{"CancellerAccessControllerAccount", "c", "q"},
		{"BypasserAccessControllerAccount", "b", "q"},
		{"RBACTimelock", "t", "q"},
		{"ProposerManyChainMultiSig", "pm", "q"},
		{"CancellerManyChainMultiSig", "cm", "q"},
		{"BypasserManyChainMultiSig", "bm", "q"},
	}
	for _, version := range []string{"1.6.0", "2.0.0"} {
		v := version160
		if version == "2.0.0" {
			v = version200
		}
		var refs []cldf_datastore.AddressRef
		for _, r := range types {
			refs = append(refs, testRef(v, r.typ, r.addr, r.qualifier))
		}
		refs = append(refs, testRef(v, "RBACTimelock", "other-qualifier", "z"))
		for i := range refs {
			refs[i].ChainSelector = testChainSel
		}

		got := GetAllMCMS(v, testChainSel, "q", refs)
		want := []string{"ac", "p", "e", "c", "b", "t", "pm", "cm", "bm"}
		require.Len(t, got, len(want))
		for i, w := range want {
			require.Equal(t, w, got[i].Address, "version %s index %d", version, i)
		}

		// a missing ref, or one of another qualifier, comes back as an empty ref in place
		got = GetAllMCMS(v, testChainSel, "missing", refs)
		require.Len(t, got, 9)
		require.Equal(t, "ac", got[0].Address)
		for _, r := range got[1:] {
			require.Empty(t, r.Address)
		}

		// refs of the other version are not read
		other := version200
		if version == "2.0.0" {
			other = version160
		}
		for _, r := range GetAllMCMS(other, testChainSel, "q", refs) {
			require.Empty(t, r.Address)
		}
	}
}

// The strings are written to datastores, so they must never change.
func TestContractTypes_Golden(t *testing.T) {
	t.Parallel()
	require.Equal(t, "RBACTimelockProgram", string(TimelockProgramType))
	require.Equal(t, "ManyChainMultiSigProgram", string(McmProgramType))
	require.Equal(t, "AccessControllerProgram", string(AccessControllerProgramType))
	require.Equal(t, "ProposerSeed", string(ProposerSeed))
	require.Equal(t, "CancellerSeed", string(CancellerSeed))
	require.Equal(t, "BypasserSeed", string(BypasserSeed))
	require.Equal(t, "RBACTimelockSeed", string(RBACTimelockSeed))
	require.Equal(t, "ProposerAccessControllerAccount", string(ProposerAccessControllerAccount))
	require.Equal(t, "ExecutorAccessControllerAccount", string(ExecutorAccessControllerAccount))
	require.Equal(t, "CancellerAccessControllerAccount", string(CancellerAccessControllerAccount))
	require.Equal(t, "BypasserAccessControllerAccount", string(BypasserAccessControllerAccount))
	require.Equal(t, "RBACTimelock", string(RBACTimelock))
	require.Equal(t, "ProposerManyChainMultiSig", string(ProposerManyChainMultisig))
	require.Equal(t, "CancellerManyChainMultiSig", string(CancellerManyChainMultisig))
	require.Equal(t, "BypasserManyChainMultiSig", string(BypasserManyChainMultisig))
}

// The module must stay importable from the CCIP 2.0 deployment module without dragging chainlink-ccip's
// deployment module (and the Go and dependency versions it forces) along.
func TestDoesNotDependOnDeploymentModule(t *testing.T) {
	t.Parallel()
	out, err := exec.CommandContext(t.Context(), "go", "list", "-deps", ".").CombinedOutput()
	require.NoError(t, err, string(out))
	for _, pkg := range strings.Fields(string(out)) {
		require.False(t, strings.HasPrefix(pkg, "github.com/smartcontractkit/chainlink-ccip/deployment"),
			"unexpected dependency %s", pkg)
		require.False(t, strings.HasPrefix(pkg, "github.com/smartcontractkit/chainlink-deployments-framework/deployment"),
			"unexpected dependency %s", pkg)
	}
}

const goldenBatchOperation = `{"chainSelector":124615329519749607,"transactions":[{"contractType":"Router","tags":[],"to":"11111111111111111111111111111111","data":"AgAAAAUAAAAAAAAA","additionalFields":{"accounts":[{"PublicKey":"DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA","IsWritable":true,"IsSigner":false},{"PublicKey":"11111111111111111111111111111112","IsWritable":true,"IsSigner":false}],"value":0}},{"contractType":"Router","tags":[],"to":"11111111111111111111111111111111","data":"AgAAAAUAAAAAAAAA","additionalFields":{"accounts":[{"PublicKey":"DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA","IsWritable":true,"IsSigner":false},{"PublicKey":"11111111111111111111111111111112","IsWritable":true,"IsSigner":false}],"value":0}}]}`
