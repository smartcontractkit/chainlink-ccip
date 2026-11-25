package tokens_test

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

const multiTokenPoolTestChainSelector = uint64(5009297550715157269)

type mockMultiTokenPoolAdapter struct {
	sequence              *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	authorizationSequence *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	validateErr           error
	validated             []tokens.DeployMultiTokenPoolInput
}

func (m *mockMultiTokenPoolAdapter) ValidateDeployMultiTokenPool(_ cldf.Environment, input tokens.DeployMultiTokenPoolInput) error {
	m.validated = append(m.validated, input)
	return m.validateErr
}

func (m *mockMultiTokenPoolAdapter) DeployMultiTokenPool() *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return m.sequence
}

func (m *mockMultiTokenPoolAdapter) AuthorizeMultiTokenPool() *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return m.authorizationSequence
}

type mockMultiTokenPoolMCMSReader struct {
	timelockRef      datastore.AddressRef
	timelockErr      error
	chainMetadataErr error
	timelockInputs   []mcms.Input
}

func (m *mockMultiTokenPoolMCMSReader) GetChainMetadata(cldf.Environment, uint64, mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{}, m.chainMetadataErr
}

func (m *mockMultiTokenPoolMCMSReader) GetTimelockRef(_ cldf.Environment, _ uint64, input mcms.Input) (datastore.AddressRef, error) {
	m.timelockInputs = append(m.timelockInputs, input)
	return m.timelockRef, m.timelockErr
}

func (m *mockMultiTokenPoolMCMSReader) GetMCMSRef(cldf.Environment, uint64, mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, nil
}

func newMockMultiTokenPoolAdapter(captured *tokens.DeployMultiTokenPoolInput) *mockMultiTokenPoolAdapter {
	return &mockMultiTokenPoolAdapter{
		sequence: cldf_ops.NewSequence(
			"mock-deploy-multi-token-pool",
			semver.MustParse("2.1.0"),
			"Captures multi-token pool deployment input",
			func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
				*captured = input
				return sequences.OnChainOutput{Addresses: []datastore.AddressRef{{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(input.PoolType),
					Version:       input.PoolVersion,
					Qualifier:     input.PoolQualifier,
					Address:       "0x0000000000000000000000000000000000000001",
				}}}, nil
			},
		),
		authorizationSequence: newMockMultiTokenPoolAuthorizationSequence(),
	}
}

func newMockMultiTokenPoolAuthorizationSequence() *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-authorize-multi-token-pool",
		semver.MustParse("2.1.0"),
		"Authorizes the multi-token pool manager",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, nil
		},
	)
}

func newMockMultiTokenPoolEnvironment(t *testing.T) cldf.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return cldf.Environment{
		Logger:           lggr,
		OperationsBundle: cldf_ops.NewBundle(t.Context, lggr, cldf_ops.NewMemoryReporter()),
		DataStore:        datastore.NewMemoryDataStore().Seal(),
		BlockChains:      cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{}),
	}
}

func newMultiTokenPoolTestRegistry(adapter tokens.MultiTokenPoolAdapter) *tokens.MultiTokenPoolAdapterRegistry {
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), adapter)
	return registry
}

func newMultiTokenPoolTestConfig(selectors ...uint64) tokens.DeployMultiTokenPoolsConfig {
	chains := make(map[uint64]tokens.DeployMultiTokenPoolChainConfig, len(selectors))
	for _, selector := range selectors {
		chains[selector] = tokens.DeployMultiTokenPoolChainConfig{
			PoolType:      "BurnMintMultiTokenPool",
			PoolVersion:   semver.MustParse("2.1.0"),
			PoolQualifier: "test-pool",
			Admin:         "0x0000000000000000000000000000000000000002",
		}
	}
	return tokens.DeployMultiTokenPoolsConfig{Chains: chains}
}

func TestMultiTokenPoolAdapterRegistry(t *testing.T) {
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	first := &mockMultiTokenPoolAdapter{}
	second := &mockMultiTokenPoolAdapter{}
	version := semver.MustParse("2.1.0")

	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, version, first)
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, version, second)

	got, ok := registry.GetMultiTokenPoolAdapter(chain_selectors.FamilyEVM, version)
	require.True(t, ok)
	require.Same(t, first, got, "the first registration must win")

	got, ok = registry.GetMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.1"))
	require.True(t, ok)
	require.Same(t, first, got)

	_, ok = registry.GetMultiTokenPoolAdapter(chain_selectors.FamilyEVM, nil)
	require.False(t, ok)
	_, ok = registry.GetMultiTokenPoolAdapter(chain_selectors.FamilySolana, version)
	require.False(t, ok)
}

func TestDeployMultiTokenPoolsVerifyPreconditions(t *testing.T) {
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), &mockMultiTokenPoolAdapter{})
	validChainCfg := tokens.DeployMultiTokenPoolChainConfig{
		PoolType:      "BurnMintMultiTokenPool",
		PoolVersion:   semver.MustParse("2.1.0"),
		PoolQualifier: "test-pool",
		Admin:         "0x0000000000000000000000000000000000000002",
	}

	tests := []struct {
		name        string
		cfg         tokens.DeployMultiTokenPoolsConfig
		expectedErr string
	}{
		{name: "valid", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: validChainCfg}}},
		{name: "no chains", cfg: tokens.DeployMultiTokenPoolsConfig{}, expectedErr: "at least one chain is required"},
		{name: "invalid MCMS", cfg: tokens.DeployMultiTokenPoolsConfig{MCMS: mcms.Input{TimelockAction: "invalid"}, Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: validChainCfg}}, expectedErr: "invalid MCMS input"},
		{name: "missing MCMS reader allows deployer fallback", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolVersion: validChainCfg.PoolVersion, PoolQualifier: validChainCfg.PoolQualifier}}}},
		{name: "missing pool type", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolVersion: validChainCfg.PoolVersion, PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "poolType is required"},
		{name: "missing pool version", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "poolVersion is required"},
		{name: "missing qualifier", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolVersion: validChainCfg.PoolVersion}}}, expectedErr: "poolQualifier is required"},
		{name: "invalid selector", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{1: validChainCfg}}, expectedErr: "invalid chain selector 1"},
		{name: "missing adapter", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolVersion: semver.MustParse("3.0.0"), PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "no multi-token pool adapter registered"},
	}

	changeset := tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry())
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := changeset.VerifyPreconditions(cldf.Environment{}, test.cfg)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}

func TestDeployMultiTokenPoolsApply(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolTestChainSelector}))
	require.NoError(t, err)

	var captured tokens.DeployMultiTokenPoolInput
	adapter := newMockMultiTokenPoolAdapter(&captured)
	var authorizationInput tokens.DeployMultiTokenPoolInput
	var authorizationRef datastore.AddressRef
	var authorizationLookupErr error
	adapter.authorizationSequence = cldf_ops.NewSequence(
		"mock-authorize-multi-token-pool-with-deployment-ref",
		semver.MustParse("2.1.0"),
		"Checks that authorization can resolve the deployed stack",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			authorizationInput = input
			authorizationRef, authorizationLookupErr = input.ExistingDataStore.Addresses().Get(datastore.AddressRef{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(input.PoolType),
				Version:       input.PoolVersion,
				Qualifier:     input.PoolQualifier,
			}.Key())
			return sequences.OnChainOutput{}, nil
		},
	)
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), adapter)

	timelockAddress := "0x0000000000000000000000000000000000000002"
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	reader := &mockMultiTokenPoolMCMSReader{timelockRef: datastore.AddressRef{
		ChainSelector: multiTokenPoolTestChainSelector,
		Address:       timelockAddress,
	}}
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, reader)

	changeset := tokens.DeployMultiTokenPools(registry, mcmsRegistry)
	cfg := tokens.DeployMultiTokenPoolsConfig{MCMS: mcms.Input{Qualifier: "custom"}, Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{
		multiTokenPoolTestChainSelector: {
			PoolType:      "BurnMintMultiTokenPool",
			PoolVersion:   semver.MustParse("2.1.0"),
			PoolQualifier: "test-pool",
		},
	}}
	require.NoError(t, changeset.VerifyPreconditions(*e, cfg))
	out, err := changeset.Apply(*e, cfg)
	require.NoError(t, err)

	require.Equal(t, multiTokenPoolTestChainSelector, captured.ChainSelector)
	require.Equal(t, timelockAddress, captured.Admin, "an omitted admin defaults to the selected timelock")
	require.Equal(t, timelockAddress, captured.SelectedTimelockAddress)
	require.Equal(t, e.DataStore, captured.ExistingDataStore)
	require.NoError(t, authorizationLookupErr)
	require.Equal(t, "0x0000000000000000000000000000000000000001", authorizationRef.Address, "authorization receives the just-deployed ref")
	require.Equal(t, timelockAddress, authorizationInput.Admin)
	require.Equal(t, timelockAddress, authorizationInput.SelectedTimelockAddress)
	require.Len(t, reader.timelockInputs, 2, "Verify and Apply must both resolve the selected timelock")
	for _, input := range reader.timelockInputs {
		require.Equal(t, "custom", input.Qualifier)
	}
	refs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, refs, 1)
	require.Equal(t, "test-pool", refs[0].Qualifier)
}

func TestDeployMultiTokenPoolsVerifyUsesAdapterValidation(t *testing.T) {
	timelockAddress := "0x0000000000000000000000000000000000000002"
	adapter := &mockMultiTokenPoolAdapter{validateErr: errors.New("unsupported pool type")}
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), adapter)
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, &mockMultiTokenPoolMCMSReader{timelockRef: datastore.AddressRef{
		ChainSelector: multiTokenPoolTestChainSelector,
		Address:       timelockAddress,
	}})
	cfg := tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{
		multiTokenPoolTestChainSelector: {PoolType: "Bogus", PoolVersion: semver.MustParse("2.1.0"), PoolQualifier: "test-pool"},
	}}

	existingDataStore := datastore.NewMemoryDataStore().Seal()
	err := tokens.DeployMultiTokenPools(registry, mcmsRegistry).VerifyPreconditions(cldf.Environment{DataStore: existingDataStore}, cfg)
	require.ErrorContains(t, err, "unsupported pool type")
	require.Len(t, adapter.validated, 1)
	require.Equal(t, timelockAddress, adapter.validated[0].Admin, "validation receives the resolved admin")
	require.Equal(t, timelockAddress, adapter.validated[0].SelectedTimelockAddress)
	require.Equal(t, existingDataStore, adapter.validated[0].ExistingDataStore)
	require.Equal(t, "Bogus", adapter.validated[0].PoolType)
}

func TestDeployMultiTokenPoolsRejectsMissingTimelockWithMCMSQualifier(t *testing.T) {
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), &mockMultiTokenPoolAdapter{})
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, &mockMultiTokenPoolMCMSReader{})
	cfg := tokens.DeployMultiTokenPoolsConfig{
		MCMS: mcms.Input{Qualifier: "custom"},
		Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{
			multiTokenPoolTestChainSelector: {PoolType: "BurnMintMultiTokenPool", PoolVersion: semver.MustParse("2.1.0"), PoolQualifier: "test-pool"},
		},
	}
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolTestChainSelector}))
	require.NoError(t, err)

	err = tokens.DeployMultiTokenPools(registry, mcmsRegistry).VerifyPreconditions(*e, cfg)
	require.ErrorContains(t, err, "refusing to fall back to the deployer")
	err = tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry()).VerifyPreconditions(*e, cfg)
	require.ErrorContains(t, err, "refusing to fall back to the deployer", "a custom qualifier must not fall back when no reader is registered")
}

func TestDeployMultiTokenPoolsApplyRejectsConflictingExistingRef(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolTestChainSelector}))
	require.NoError(t, err)
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: multiTokenPoolTestChainSelector,
		Type:          "BurnMintMultiTokenPool",
		Version:       semver.MustParse("2.1.0"),
		Qualifier:     "test-pool",
		Address:       "0x00000000000000000000000000000000000000ff",
	}))
	e.DataStore = ds.Seal()

	var captured tokens.DeployMultiTokenPoolInput
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), newMockMultiTokenPoolAdapter(&captured))
	cfg := tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{
		multiTokenPoolTestChainSelector: {
			PoolType:      "BurnMintMultiTokenPool",
			PoolVersion:   semver.MustParse("2.1.0"),
			PoolQualifier: "test-pool",
			Admin:         "0x0000000000000000000000000000000000000002",
		},
	}}

	out, err := tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry()).Apply(*e, cfg)
	require.ErrorContains(t, err, "conflicts with the existing datastore address")
	require.Len(t, out.Reports, 1, "the successful sequence report must survive output processing errors")
	require.NotNil(t, out.DataStore)
	refs, fetchErr := out.DataStore.Addresses().Fetch()
	require.NoError(t, fetchErr)
	require.Empty(t, refs)
}

func TestDeployMultiTokenPoolsRejectsResolvedTimelockWithBlankAddress(t *testing.T) {
	tests := []struct {
		name  string
		ref   datastore.AddressRef
		admin string
	}{
		{
			name: "typed ref with default admin",
			ref: datastore.AddressRef{
				ChainSelector: multiTokenPoolTestChainSelector,
				Type:          "Timelock",
				Version:       semver.MustParse("1.0.0"),
			},
		},
		{
			name: "qualified ref with explicit admin",
			ref: datastore.AddressRef{
				ChainSelector: multiTokenPoolTestChainSelector,
				Qualifier:     "custom",
			},
			admin: "0x0000000000000000000000000000000000000003",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var captured tokens.DeployMultiTokenPoolInput
			registry := newMultiTokenPoolTestRegistry(newMockMultiTokenPoolAdapter(&captured))
			mcmsRegistry := changesets.NewMCMSReaderRegistry()
			mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, &mockMultiTokenPoolMCMSReader{timelockRef: test.ref})
			changeset := tokens.DeployMultiTokenPools(registry, mcmsRegistry)
			cfg := newMultiTokenPoolTestConfig(multiTokenPoolTestChainSelector)
			chainCfg := cfg.Chains[multiTokenPoolTestChainSelector]
			chainCfg.Admin = test.admin
			cfg.Chains[multiTokenPoolTestChainSelector] = chainCfg
			e := newMockMultiTokenPoolEnvironment(t)

			err := changeset.VerifyPreconditions(e, cfg)
			require.ErrorContains(t, err, "has an empty address")
			out, err := changeset.Apply(e, cfg)
			require.ErrorContains(t, err, "has an empty address")
			require.Empty(t, out.Reports)
		})
	}
}

func TestDeployMultiTokenPoolsPropagatesTimelockReaderErrors(t *testing.T) {
	var captured tokens.DeployMultiTokenPoolInput
	registry := newMultiTokenPoolTestRegistry(newMockMultiTokenPoolAdapter(&captured))
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, &mockMultiTokenPoolMCMSReader{timelockErr: errors.New("reader unavailable")})
	changeset := tokens.DeployMultiTokenPools(registry, mcmsRegistry)
	cfg := newMultiTokenPoolTestConfig(multiTokenPoolTestChainSelector)
	e := newMockMultiTokenPoolEnvironment(t)

	err := changeset.VerifyPreconditions(e, cfg)
	require.ErrorContains(t, err, "failed to resolve timelock")
	require.ErrorContains(t, err, "reader unavailable")
	out, err := changeset.Apply(e, cfg)
	require.ErrorContains(t, err, "failed to resolve timelock")
	require.ErrorContains(t, err, "reader unavailable")
	require.Empty(t, out.Reports)
}

func TestDeployMultiTokenPoolsApplyReplaysDeploymentAndReauthorizesWithSameReporter(t *testing.T) {
	const (
		deploymentSequenceID    = "mock-deploy-multi-token-pool-replay"
		deploymentOperationID   = "mock-deploy-multi-token-pool-contract"
		authorizationSequenceID = "mock-authorize-multi-token-pool-replay"
	)
	var deploymentExecutions int
	var deploymentOperationExecutions int
	var authorizationExecutions int
	deploymentOperation := cldf_ops.NewOperation(
		deploymentOperationID,
		semver.MustParse("2.1.0"),
		"Deploys a contract",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ uint64) (string, error) {
			deploymentOperationExecutions++
			return "deployed", nil
		},
	)
	deploymentSequence := cldf_ops.NewSequence(
		deploymentSequenceID,
		semver.MustParse("2.1.0"),
		"Deploys a multi-token pool",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			deploymentExecutions++
			if _, err := cldf_ops.ExecuteOperation(b, deploymentOperation, chains, input.ChainSelector); err != nil {
				return sequences.OnChainOutput{}, err
			}
			return sequences.OnChainOutput{Addresses: []datastore.AddressRef{{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(input.PoolType),
				Version:       input.PoolVersion,
				Qualifier:     input.PoolQualifier,
				Address:       "0x0000000000000000000000000000000000000001",
			}}}, nil
		},
	)
	authorizationSequence := cldf_ops.NewSequence(
		authorizationSequenceID,
		semver.MustParse("2.1.0"),
		"Authorizes the multi-token pool stores",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			authorizationExecutions++
			return sequences.OnChainOutput{}, nil
		},
	)
	registry := newMultiTokenPoolTestRegistry(&mockMultiTokenPoolAdapter{
		sequence:              deploymentSequence,
		authorizationSequence: authorizationSequence,
	})
	changeset := tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry())
	cfg := newMultiTokenPoolTestConfig(multiTokenPoolTestChainSelector)
	chainCfg := cfg.Chains[multiTokenPoolTestChainSelector]
	chainCfg.Admin = ""
	cfg.Chains[multiTokenPoolTestChainSelector] = chainCfg
	e := newMockMultiTokenPoolEnvironment(t)
	reporter := cldf_ops.NewMemoryReporter()
	e.OperationsBundle = cldf_ops.NewBundle(t.Context, e.Logger, reporter)

	require.NoError(t, changeset.VerifyPreconditions(e, cfg), "a missing reader and timelock allow the direct deployer-owned flow")
	first, err := changeset.Apply(e, cfg)
	require.NoError(t, err)
	second, err := changeset.Apply(e, cfg)
	require.NoError(t, err)

	require.Equal(t, 1, deploymentExecutions, "the successful deployment sequence must be replayed from its report")
	require.Equal(t, 1, deploymentOperationExecutions, "the deployment child operation must not execute again")
	require.Equal(t, 2, authorizationExecutions, "authorization must execute on every Apply")
	require.Len(t, first.Reports, 3)
	require.Len(t, second.Reports, 3, "replayed deployment reports must be retained alongside fresh authorization reports")
	secondReportCounts := make(map[string]int)
	var firstDeploymentReportID string
	var secondDeploymentReportID string
	var firstAuthorizationKey string
	var secondAuthorizationKey string
	for _, report := range first.Reports {
		switch report.Def.ID {
		case deploymentSequenceID:
			firstDeploymentReportID = report.ID
		case authorizationSequenceID:
			firstAuthorizationKey = report.IdempotencyKey
		}
	}
	for _, report := range second.Reports {
		secondReportCounts[report.Def.ID]++
		switch report.Def.ID {
		case deploymentSequenceID:
			secondDeploymentReportID = report.ID
		case authorizationSequenceID:
			secondAuthorizationKey = report.IdempotencyKey
		}
	}
	require.Equal(t, 1, secondReportCounts[deploymentSequenceID])
	require.Equal(t, 1, secondReportCounts[deploymentOperationID], "the original child deployment report must survive replay")
	require.Equal(t, 1, secondReportCounts[authorizationSequenceID])
	require.NotEmpty(t, firstDeploymentReportID)
	require.Equal(t, firstDeploymentReportID, secondDeploymentReportID)
	require.NotEmpty(t, firstAuthorizationKey)
	require.NotEmpty(t, secondAuthorizationKey)
	require.NotEqual(t, firstAuthorizationKey, secondAuthorizationKey)
	refs, err := second.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, refs, 1)
}

func TestDeployMultiTokenPoolsApplyPreservesReportsOnLaterSequenceError(t *testing.T) {
	const (
		sequenceID  = "mock-deploy-multi-token-pool-error"
		operationID = "mock-multi-token-pool-child-before-error"
	)
	successSelector := chain_selectors.TEST_90000001.Selector
	failingSelector := chain_selectors.TEST_90000002.Selector
	operation := cldf_ops.NewOperation(
		operationID,
		semver.MustParse("2.1.0"),
		"Completes a child operation",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ uint64) (string, error) {
			return "done", nil
		},
	)
	sequence := cldf_ops.NewSequence(
		sequenceID,
		semver.MustParse("2.1.0"),
		"Fails after a child operation on the requested chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			if input.ChainSelector == failingSelector {
				if _, err := cldf_ops.ExecuteOperation(b, operation, chains, input.ChainSelector); err != nil {
					return sequences.OnChainOutput{}, err
				}
				return sequences.OnChainOutput{}, errors.New("mock sequence failure")
			}
			return sequences.OnChainOutput{Addresses: []datastore.AddressRef{{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(input.PoolType),
				Version:       input.PoolVersion,
				Qualifier:     input.PoolQualifier,
				Address:       "0x0000000000000000000000000000000000000001",
			}}}, nil
		},
	)
	registry := newMultiTokenPoolTestRegistry(&mockMultiTokenPoolAdapter{
		sequence:              sequence,
		authorizationSequence: newMockMultiTokenPoolAuthorizationSequence(),
	})
	changeset := tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry())
	cfg := newMultiTokenPoolTestConfig(successSelector, failingSelector)
	e := newMockMultiTokenPoolEnvironment(t)

	out, err := changeset.Apply(e, cfg)
	require.ErrorContains(t, err, "mock sequence failure")
	require.Len(t, out.Reports, 4, "the earlier deployment and authorization, failing deployment, and its child must each appear once")
	reportCounts := make(map[string]int)
	for _, report := range out.Reports {
		reportCounts[report.Def.ID]++
	}
	require.Equal(t, 2, reportCounts[sequenceID])
	require.Equal(t, 1, reportCounts[operationID])
	require.Equal(t, 1, reportCounts["mock-authorize-multi-token-pool"])
	require.NotNil(t, out.DataStore)
	refs, fetchErr := out.DataStore.Addresses().Fetch()
	require.NoError(t, fetchErr)
	require.Len(t, refs, 1)
	require.Equal(t, successSelector, refs[0].ChainSelector)
}

func TestDeployMultiTokenPoolsApplyPreservesOutputOnAuthorizationError(t *testing.T) {
	var captured tokens.DeployMultiTokenPoolInput
	adapter := newMockMultiTokenPoolAdapter(&captured)
	adapter.authorizationSequence = cldf_ops.NewSequence(
		"mock-authorize-multi-token-pool-error",
		semver.MustParse("2.1.0"),
		"Fails authorization",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, errors.New("mock authorization failure")
		},
	)
	registry := newMultiTokenPoolTestRegistry(adapter)
	changeset := tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry())
	cfg := newMultiTokenPoolTestConfig(multiTokenPoolTestChainSelector)
	e := newMockMultiTokenPoolEnvironment(t)

	out, err := changeset.Apply(e, cfg)
	require.ErrorContains(t, err, "failed to authorize multi-token pool")
	require.ErrorContains(t, err, "mock authorization failure")
	require.Len(t, out.Reports, 2, "deployment and failed authorization reports must be returned")
	require.Equal(t, "mock-deploy-multi-token-pool", out.Reports[0].Def.ID)
	require.Equal(t, "mock-authorize-multi-token-pool-error", out.Reports[1].Def.ID)
	require.NotNil(t, out.DataStore)
	refs, fetchErr := out.DataStore.Addresses().Fetch()
	require.NoError(t, fetchErr)
	require.Len(t, refs, 1)
}

func TestDeployMultiTokenPoolsApplyPreservesOutputOnProposalBuildError(t *testing.T) {
	var captured tokens.DeployMultiTokenPoolInput
	sequence := cldf_ops.NewSequence(
		"mock-deploy-multi-token-pool-proposal-error",
		semver.MustParse("2.1.0"),
		"Returns an address and a proposal operation",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			captured = input
			return sequences.OnChainOutput{Addresses: []datastore.AddressRef{{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(input.PoolType),
				Version:       input.PoolVersion,
				Qualifier:     input.PoolQualifier,
				Address:       "0x0000000000000000000000000000000000000001",
			}}}, nil
		},
	)
	authorizationSequence := cldf_ops.NewSequence(
		"mock-authorize-multi-token-pool-proposal-error",
		semver.MustParse("2.1.0"),
		"Returns a proposal operation",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{{
				ChainSelector: mcms_types.ChainSelector(input.ChainSelector),
				Transactions:  []mcms_types.Transaction{{To: "0x0000000000000000000000000000000000000001"}},
			}}}, nil
		},
	)
	registry := newMultiTokenPoolTestRegistry(&mockMultiTokenPoolAdapter{
		sequence:              sequence,
		authorizationSequence: authorizationSequence,
	})
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, &mockMultiTokenPoolMCMSReader{
		timelockRef: datastore.AddressRef{
			ChainSelector: multiTokenPoolTestChainSelector,
			Address:       "0x0000000000000000000000000000000000000002",
		},
		chainMetadataErr: errors.New("metadata unavailable"),
	})
	changeset := tokens.DeployMultiTokenPools(registry, mcmsRegistry)
	cfg := newMultiTokenPoolTestConfig(multiTokenPoolTestChainSelector)
	e := newMockMultiTokenPoolEnvironment(t)

	out, err := changeset.Apply(e, cfg)
	require.ErrorContains(t, err, "failed to get chain metadata")
	require.ErrorContains(t, err, "metadata unavailable")
	require.Len(t, out.Reports, 2, "deployment and authorization reports must survive proposal build errors")
	require.NotNil(t, out.DataStore)
	refs, fetchErr := out.DataStore.Addresses().Fetch()
	require.NoError(t, fetchErr)
	require.Len(t, refs, 1)
	require.Equal(t, "0x0000000000000000000000000000000000000002", captured.SelectedTimelockAddress)
}
