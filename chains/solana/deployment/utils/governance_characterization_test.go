package utils

import (
	"encoding/json"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/stretchr/testify/require"

	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// Characterization tests: they pin the current behavior of the governance helpers so the move
// into the shared governance module is provably behavior-preserving.

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
	const sel = uint64(124615329519749607)
	mk := func(typ, addr, q string) cldf_datastore.AddressRef {
		return cldf_datastore.AddressRef{
			ChainSelector: sel, Type: cldf_datastore.ContractType(typ), Version: common_utils.Version_1_6_0,
			Qualifier: q, Address: addr,
		}
	}
	refs := []cldf_datastore.AddressRef{
		mk("AccessControllerProgram", "ac", ""),
		mk("ProposerAccessControllerAccount", "p", "q"),
		mk("ExecutorAccessControllerAccount", "e", "q"),
		mk("CancellerAccessControllerAccount", "c", "q"),
		mk("BypasserAccessControllerAccount", "b", "q"),
		mk("RBACTimelock", "t", "q"),
		mk("ProposerManyChainMultiSig", "pm", "q"),
		mk("CancellerManyChainMultiSig", "cm", "q"),
		mk("BypasserManyChainMultiSig", "bm", "q"),
		mk("RBACTimelock", "other-qualifier", "z"),
	}
	chain := cldf_solana.Chain{Selector: sel}

	got := GetAllMCMS(chain, "q", refs)
	want := []string{"ac", "p", "e", "c", "b", "t", "pm", "cm", "bm"}
	require.Len(t, got, len(want))
	for i, w := range want {
		require.Equal(t, w, got[i].Address, "index %d", i)
	}

	// refs of another qualifier or a missing ref come back as empty refs, in place
	got = GetAllMCMS(chain, "missing", refs)
	require.Len(t, got, 9)
	require.Equal(t, "ac", got[0].Address)
	for _, r := range got[1:] {
		require.Empty(t, r.Address)
	}
}

func TestGovernanceConstants_Golden(t *testing.T) {
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
	require.Equal(t, "SPLTokens", string(SPLTokens))
	require.Equal(t, "SPL2022Tokens", string(SPL2022Tokens))
}

const goldenBatchOperation = `{"chainSelector":124615329519749607,"transactions":[{"contractType":"Router","tags":[],"to":"11111111111111111111111111111111","data":"AgAAAAUAAAAAAAAA","additionalFields":{"accounts":[{"PublicKey":"DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA","IsWritable":true,"IsSigner":false},{"PublicKey":"11111111111111111111111111111112","IsWritable":true,"IsSigner":false}],"value":0}},{"contractType":"Router","tags":[],"to":"11111111111111111111111111111111","data":"AgAAAAUAAAAAAAAA","additionalFields":{"accounts":[{"PublicKey":"DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA","IsWritable":true,"IsSigner":false},{"PublicKey":"11111111111111111111111111111112","IsWritable":true,"IsSigner":false}],"value":0}}]}`
