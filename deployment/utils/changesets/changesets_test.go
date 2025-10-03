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

func (m *MockReader) OpCount(e deployment.Environment, chainSelector uint64, mcmAddress string) (uint64, error) {
	return OP_COUNT, nil
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
		input        sequences.OnChainOutput
		resolveInput func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error)
		resolveDep   func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error)
	}{
		{
			desc: "happy path - all executed",
			input: sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: 4340886533089894000,
						Address:       common.HexToAddress("0x01").String(),
						Type:          datastore.ContractType("TestContract"),
						Version:       semver.MustParse("1.0.0"),
					},
				},
				Writes: []contract.WriteOutput{
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
			input: sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: 4340886533089894000,
						Address:       common.HexToAddress("0x01").String(),
						Type:          datastore.ContractType("TestContract"),
						Version:       semver.MustParse("1.0.0"),
					},
				},
				Writes: []contract.WriteOutput{
					{
						ChainSelector: 4340886533089894000,
						Tx: mcms_types.Transaction{
							To:               common.HexToAddress("0x01").Hex(),
							Data:             common.Hex2Bytes("0xdeadbeef"),
							AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
						},
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
			desc:  "validation error in resolveInput",
			input: sequences.OnChainOutput{},
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, errors.New("")
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
		},
		{
			desc:  "validation error in resolveDep",
			input: sequences.OnChainOutput{},
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
			if _, err := test.resolveInput(e, test.input); err != nil {
				expectErr = true
			}
			if _, err := test.resolveDep(e, test.input); err != nil {
				expectErr = true
			}

			err = changeset.VerifyPreconditions(e, changesets.WithMCMS[sequences.OnChainOutput]{
				Cfg: test.input,
			})
			if expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			out, err := changeset.Apply(e, changesets.WithMCMS[sequences.OnChainOutput]{
				Cfg: test.input,
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
			require.Len(t, dsAddrs, len(test.input.Addresses))

			var someNotExecuted bool
			for _, w := range test.input.Writes {
				if !w.Executed() {
					someNotExecuted = true
					break
				}
			}
			if len(test.input.Writes) > 0 && someNotExecuted {
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
