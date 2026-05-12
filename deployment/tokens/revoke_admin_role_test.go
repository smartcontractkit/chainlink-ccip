package tokens

import (
	"context"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	revokeTestTokenAddress    = "0x00000000000000000000000000000000000000aa"
	revokeTestTimelockAddress = "0x00000000000000000000000000000000000000bb"
	revokeTestAdminAddress    = "0x00000000000000000000000000000000000000cc"
	revokeTestMCMAddress      = "0x00000000000000000000000000000000000000dd"
)

func TestRevokeTokenAdminRoleVerify(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	env := revokeTestEnv(t, chainSelector, revokeTestDataStore(t, chainSelector))
	registeredRegistry := revokeTestTokenRegistry(&revokeTestTokenAdapter{
		sequence: (&revokeTestSequenceRecorder{}).Sequence(),
	})

	tests := []struct {
		name          string
		registry      *TokenAdapterRegistry
		input         RevokeTokenAdminRoleInput
		errorContains string
	}{
		{
			name:          "requires revocations",
			registry:      registeredRegistry,
			input:         RevokeTokenAdminRoleInput{},
			errorContains: "at least one token admin role revocation is required",
		},
		{
			name:     "requires chain selector",
			registry: registeredRegistry,
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{TokenRef: revokeTestTokenRef(chainSelector)}},
			},
			errorContains: "chain selector is required",
		},
		{
			name:     "requires chain in environment",
			registry: registeredRegistry,
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector + 1,
					TokenRef:      revokeTestTokenRef(chainSelector + 1),
				}},
			},
			errorContains: "not found in environment",
		},
		{
			name:     "requires token ref",
			registry: registeredRegistry,
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{ChainSelector: chainSelector}},
			},
			errorContains: "token ref is required",
		},
		{
			name:     "rejects token ref chain mismatch",
			registry: registeredRegistry,
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef:      revokeTestTokenRef(chainSelector + 1),
				}},
			},
			errorContains: "token ref chain selector mismatch",
		},
		{
			name:     "requires registered adapter",
			registry: newTokenAdapterRegistry(),
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef:      revokeTestTokenRef(chainSelector),
				}},
			},
			errorContains: "no TokenPoolAdapter registered",
		},
		{
			name:     "requires resolvable token ref",
			registry: registeredRegistry,
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Address: "0x00000000000000000000000000000000000000ee",
					},
				}},
			},
			errorContains: "token ref must resolve from datastore or include both address and type",
		},
		{
			name:     "accepts resolvable token ref",
			registry: registeredRegistry,
			input: RevokeTokenAdminRoleInput{
				Revocations: []RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Address: revokeTestTokenAddress,
					},
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := revokeTokenAdminRoleVerify(tt.registry)(env, tt.input)
			if tt.errorContains != "" {
				require.ErrorContains(t, err, tt.errorContains)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestRevokeTokenAdminRoleApplyDefaultsAdminToTimelockAndRerunsSequence(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	recorder := &revokeTestSequenceRecorder{batchOnCalls: map[int]bool{1: true}}
	env := revokeTestEnv(t, chainSelector, revokeTestDataStore(t, chainSelector))
	input := RevokeTokenAdminRoleInput{
		Revocations: []RevokeTokenAdminRoleConfig{{
			ChainSelector: chainSelector,
			TokenRef: datastore.AddressRef{
				Address: revokeTestTokenAddress,
			},
		}},
		MCMS: revokeTestMCMSInput(),
	}
	apply := revokeTokenAdminRoleApply(
		revokeTestTokenRegistry(&revokeTestTokenAdapter{sequence: recorder.Sequence()}),
		revokeTestMCMSRegistry(),
	)

	out, err := apply(env, input)
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	require.Len(t, recorder.inputs, 1)
	require.Equal(t, revokeTestTimelockAddress, recorder.inputs[0].AdminAddress)
	require.Equal(t, revokeTestTimelockAddress, recorder.inputs[0].TimelockAddress)
	require.Equal(t, datastore.ContractType("BurnMintERC20"), recorder.inputs[0].TokenRef.Type)

	out, err = apply(env, input)
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals)
	require.Len(t, recorder.inputs, 2)
}

func TestRevokeTokenAdminRoleApplyUsesExplicitAdmin(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	recorder := &revokeTestSequenceRecorder{}
	env := revokeTestEnv(t, chainSelector, revokeTestDataStore(t, chainSelector))
	input := RevokeTokenAdminRoleInput{
		Revocations: []RevokeTokenAdminRoleConfig{{
			ChainSelector: chainSelector,
			TokenRef:      revokeTestTokenRef(chainSelector),
			AdminAddress:  revokeTestAdminAddress,
		}},
		MCMS: revokeTestMCMSInput(),
	}
	apply := revokeTokenAdminRoleApply(
		revokeTestTokenRegistry(&revokeTestTokenAdapter{sequence: recorder.Sequence()}),
		revokeTestMCMSRegistry(),
	)

	out, err := apply(env, input)
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals)
	require.Len(t, recorder.inputs, 1)
	require.Equal(t, revokeTestAdminAddress, recorder.inputs[0].AdminAddress)
	require.Equal(t, revokeTestTimelockAddress, recorder.inputs[0].TimelockAddress)
}

type revokeTestTokenAdapter struct {
	TokenAdapter
	sequence *cldf_ops.Sequence[RevokeTokenAdminRoleSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

func (a *revokeTestTokenAdapter) RevokeTokenAdminRole() *cldf_ops.Sequence[RevokeTokenAdminRoleSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return a.sequence
}

type revokeTestSequenceRecorder struct {
	inputs       []RevokeTokenAdminRoleSequenceInput
	batchOnCalls map[int]bool
	err          error
}

func (r *revokeTestSequenceRecorder) Sequence() *cldf_ops.Sequence[RevokeTokenAdminRoleSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"revoke-test-token-admin-role",
		semver.MustParse("1.0.0"),
		"revoke token admin role test sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input RevokeTokenAdminRoleSequenceInput) (sequences.OnChainOutput, error) {
			r.inputs = append(r.inputs, input)
			if r.err != nil {
				return sequences.OnChainOutput{}, r.err
			}
			if !r.batchOnCalls[len(r.inputs)] {
				return sequences.OnChainOutput{}, nil
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{{
					ChainSelector: mcms_types.ChainSelector(input.ChainSelector),
					Transactions: []mcms_types.Transaction{{
						To:               revokeTestTokenAddress,
						Data:             []byte{0x01},
						AdditionalFields: []byte("{}"),
					}},
				}},
			}, nil
		},
	)
}

type revokeTestMCMSReader struct{}

func (r *revokeTestMCMSReader) GetChainMetadata(_ cldf.Environment, _ uint64, _ mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{StartingOpCount: 1}, nil
}

func (r *revokeTestMCMSReader) GetTimelockRef(_ cldf.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       revokeTestTimelockAddress,
		Type:          datastore.ContractType("Timelock"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func (r *revokeTestMCMSReader) GetMCMSRef(_ cldf.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       revokeTestMCMAddress,
		Type:          datastore.ContractType("MCM"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func revokeTestTokenRegistry(adapter TokenAdapter) *TokenAdapterRegistry {
	registry := newTokenAdapterRegistry()
	registry.RegisterTokenAdapter(chainsel.FamilyEVM, cciputils.Version_1_0_0, adapter)
	return registry
}

func revokeTestMCMSRegistry() *changesets.MCMSReaderRegistry {
	registry := &changesets.MCMSReaderRegistry{}
	registry.RegisterMCMSReader(chainsel.FamilyEVM, &revokeTestMCMSReader{})
	return registry
}

func revokeTestEnv(t *testing.T, chainSelector uint64, ds *datastore.MemoryDataStore) cldf.Environment {
	t.Helper()

	lggr, err := logger.New()
	require.NoError(t, err)
	return cldf.Environment{
		Logger:    lggr,
		DataStore: ds.Seal(),
		GetContext: func() context.Context {
			return context.Background()
		},
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
		BlockChains: cldf_chain.NewBlockChainsFromSlice([]cldf_chain.BlockChain{
			evm.Chain{Selector: chainSelector},
		}),
	}
}

func revokeTestDataStore(t *testing.T, chainSelector uint64) *datastore.MemoryDataStore {
	t.Helper()

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(revokeTestTokenRef(chainSelector)))
	return ds
}

func revokeTestTokenRef(chainSelector uint64) datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       revokeTestTokenAddress,
		Type:          datastore.ContractType("BurnMintERC20"),
		Version:       semver.MustParse("1.0.0"),
	}
}

func revokeTestMCMSInput() mcms.Input {
	return mcms.Input{
		OverridePreviousRoot: true,
		ValidUntil:           3759765795,
		TimelockDelay:        mcms_types.MustParseDuration("1h"),
		TimelockAction:       mcms_types.TimelockActionSchedule,
	}
}
