package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
)

// mockLombardPoolRemover is a Lombard adapter that also supports removing remote pools.
type mockLombardPoolRemover struct {
	mockLombardAdapter
	removed []tokens.RemoveRemotePoolsSequenceInput
}

var _ tokens.RemotePoolRemover = (*mockLombardPoolRemover)(nil)

func (m *mockLombardPoolRemover) RemoveRemotePools() *cldf_ops.Sequence[tokens.RemoveRemotePoolsSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-remove-remote-pools",
		semver.MustParse("1.0.0"),
		"Mock remove remote pools",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in tokens.RemoveRemotePoolsSequenceInput) (sequences.OnChainOutput, error) {
			m.removed = append(m.removed, in)
			return sequences.OnChainOutput{}, nil
		},
	)
}

func TestRemoveLombardRemotePools_Apply_Success(t *testing.T) {
	env := newLombardTestEnv(t)
	mock := &mockLombardPoolRemover{}
	registry := adapters.NewLombardChainRegistry()
	registry.RegisterLombardChain(chainsel.FamilyEVM, mock)

	entry := tokens.RemoveRemotePoolsSequenceInput{Selector: lombSelA}
	cs := changesets.RemoveLombardRemotePools(registry, lombardRegistry())
	_, err := cs.Apply(env, changesets.LombardTokenPoolsConfig[tokens.RemoveRemotePoolsSequenceInput]{
		Entries: []tokens.RemoveRemotePoolsSequenceInput{entry},
	})
	require.NoError(t, err)
	require.Len(t, mock.removed, 1)
	assert.Equal(t, lombSelA, mock.removed[0].Selector)
}

func TestRemoveLombardRemotePools_Apply_OperationNotSupported(t *testing.T) {
	env := newLombardTestEnv(t)
	registry := adapters.NewLombardChainRegistry()
	registry.RegisterLombardChain(chainsel.FamilyEVM, &mockLombardAdapter{})

	cs := changesets.RemoveLombardRemotePools(registry, lombardRegistry())
	_, err := cs.Apply(env, changesets.LombardTokenPoolsConfig[tokens.RemoveRemotePoolsSequenceInput]{
		Entries: []tokens.RemoveRemotePoolsSequenceInput{{Selector: lombSelA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not support remove remote pools")
}

func TestSetLombardTokenPoolRateLimits_Apply_NoAdapterRegistered(t *testing.T) {
	env := newLombardTestEnv(t)
	cs := changesets.SetLombardTokenPoolRateLimits(adapters.NewLombardChainRegistry(), lombardRegistry())
	_, err := cs.Apply(env, changesets.LombardTokenPoolsConfig[tokens.TPRLRemotes]{
		Entries: []tokens.TPRLRemotes{{ChainSelector: lombSelA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no Lombard adapter registered")
}

func TestSetLombardTokenPoolDynamicConfig_Validate_NoEntries(t *testing.T) {
	env := newLombardTestEnv(t)
	cs := changesets.SetLombardTokenPoolDynamicConfig(adapters.NewLombardChainRegistry(), lombardRegistry())
	err := cs.VerifyPreconditions(env, changesets.LombardTokenPoolsConfig[tokens.SetTokenPoolDynamicConfigSequenceInput]{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one entry")
}

func TestSetCCTPTokenPoolRateLimits_Validate_InvalidUSDCType(t *testing.T) {
	env := newLombardTestEnv(t)
	cs := changesets.SetCCTPTokenPoolRateLimits(adapters.NewCCTPChainRegistry(), lombardRegistry())
	err := cs.VerifyPreconditions(env, changesets.CCTPTokenPoolsConfig[tokens.TPRLRemotes]{
		Entries: []tokens.TPRLRemotes{{ChainSelector: lombSelA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid USDC type")
}

func TestRemoveCCTPRemotePools_Apply_NoAdapterRegistered(t *testing.T) {
	env := newLombardTestEnv(t)
	cs := changesets.RemoveCCTPRemotePools(adapters.NewCCTPChainRegistry(), lombardRegistry())
	_, err := cs.Apply(env, changesets.CCTPTokenPoolsConfig[tokens.RemoveRemotePoolsSequenceInput]{
		USDCType: adapters.Canonical,
		Entries:  []tokens.RemoveRemotePoolsSequenceInput{{Selector: lombSelA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no CCTP adapter registered")
}
