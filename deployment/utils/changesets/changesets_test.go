package changesets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

type MockReader struct{}

const OP_COUNT = 42

func (m *MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		MCMAddress:      input.MCMSAddressRef.Address,
		StartingOpCount: OP_COUNT,
	}, nil
}

func init() {
	changesets.RegisterMCMSReader("evm", &MockReader{})
}

var mockSequence = operations.NewSequence(
	"mock-sequence",
	semver.MustParse("1.0.0"),
	"Mock sequence for testing",
	func(b operations.Bundle, deps int, in sequences.OnChainOutput) (sequences.OnChainOutput, error) {
		return in, nil
	},
)

func TestNewFromOnChainSequence(t *testing.T) {
	tests := []struct {
		desc         string
		addresses    []datastore.AddressRef
		writes       []contract.WriteOutput
		resolveInput func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error)
		resolveDep   func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error)
	}{
		{
			desc: "happy path - all executed",
			addresses: []datastore.AddressRef{
				{
					ChainSelector: 4340886533089894000,
					Address:       common.HexToAddress("0x01").String(),
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
				},
			},
			writes: []contract.WriteOutput{
				{
					ChainSelector: 4340886533089894000,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
					ExecInfo: &contract.ExecInfo{
						Hash: common.HexToHash("0x02").Hex(),
					},
				},
			},
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, nil
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
		},
		{
			desc: "happy path - not executed",
			addresses: []datastore.AddressRef{
				{
					ChainSelector: 4340886533089894000,
					Address:       common.HexToAddress("0x01").String(),
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
				},
			},
			writes: []contract.WriteOutput{
				{
					ChainSelector: 4340886533089894000,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
			},
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, nil
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
		},
		{
			desc: "validation error in resolveInput",
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, errors.New("")
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
		},
		{
			desc: "validation error in resolveDep",
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, nil
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, errors.New("")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lggr, err := logger.New()
			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)
			ds := datastore.NewMemoryDataStore()
			err = ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: 4340886533089894000,
				Type:          "Timelock",
				Version:       semver.MustParse("1.0.0"),
				Address:       common.HexToAddress("0x01").Hex(),
			})
			require.NoError(t, err)
			err = ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: 4340886533089894000,
				Type:          "MCM",
				Version:       semver.MustParse("1.0.0"),
				Address:       common.HexToAddress("0x02").Hex(),
			})
			require.NoError(t, err)
			e := deployment.Environment{
				OperationsBundle: bundle,
				DataStore:        ds.Seal(),
			}

			changeset := changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[sequences.OnChainOutput, int, sequences.OnChainOutput]{
				Sequence:     mockSequence,
				ResolveInput: test.resolveInput,
				ResolveDep:   test.resolveDep,
			})

			var expectErr bool
			batchOp, err := contract.NewBatchOperationFromWrites(test.writes)
			require.NoError(t, err)
			input := sequences.OnChainOutput{
				Addresses: test.addresses,
				BatchOps:  []mcms_types.BatchOperation{batchOp},
			}
			// Pre-check that the input can be resolved outside of the changeset flow.
			// This ensures that any errors in resolveInput or resolveDep are from the changeset flow.
			if _, err := test.resolveInput(e, input); err != nil {
				expectErr = true
			}
			if _, err := test.resolveDep(e, input); err != nil {
				expectErr = true
			}

			err = changeset.VerifyPreconditions(e, changesets.WithMCMS[sequences.OnChainOutput]{
				Cfg: input,
			})
			if expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			out, err := changeset.Apply(e, changesets.WithMCMS[sequences.OnChainOutput]{
				Cfg: input,
				MCMS: mcms.Input{
					OverridePreviousRoot: true,
					ValidUntil:           4126214326,
					TimelockDelay:        mcms_types.MustParseDuration("1h"),
					TimelockAction:       mcms_types.TimelockActionSchedule,
					MCMSAddressRef: datastore.AddressRef{
						Type:    "MCM",
						Version: semver.MustParse("1.0.0"),
					},
					TimelockAddressRef: datastore.AddressRef{
						Type:    "Timelock",
						Version: semver.MustParse("1.0.0"),
					},
					Description: "Test Proposal",
				},
			})
			require.NoError(t, err)

			dsAddrs, err := out.DataStore.Addresses().Fetch()
			require.Len(t, dsAddrs, len(input.Addresses))

			if len(batchOp.Transactions) > 0 {
				require.Len(t, out.MCMSTimelockProposals, 1)
				require.Len(t, out.MCMSTimelockProposals[0].Operations, 1)
				require.Equal(t, out.MCMSTimelockProposals[0].OverridePreviousRoot, true)
				require.Equal(t, out.MCMSTimelockProposals[0].ValidUntil, uint32(4126214326))
				require.Equal(t, out.MCMSTimelockProposals[0].Delay, mcms_types.MustParseDuration("1h"))
				require.Equal(t, out.MCMSTimelockProposals[0].Action, mcms_types.TimelockActionSchedule)
				require.Equal(t, out.MCMSTimelockProposals[0].Description, "Test Proposal")
				require.Equal(t, uint64(OP_COUNT), out.MCMSTimelockProposals[0].ChainMetadata[4340886533089894000].StartingOpCount)
			}
		})
	}
}
