package changesets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

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
		describe     func(in sequences.OnChainOutput, dep int) string
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
						Executed: true,
					},
				},
			},
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, nil
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
			describe: func(in sequences.OnChainOutput, dep int) string {
				return "mock description"
			},
		},
		/* --- IGNORE ---
		MCMS is not supported yet, attempting to use it will result in errors
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
				Writes: []call.WriteOutput{
					{
						ChainSelector: 4340886533089894000,
						Tx: mcms_types.Transaction{
							To:               common.HexToAddress("0x01").Hex(),
							Data:             common.Hex2Bytes("0xdeadbeef"),
							AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
						},
						Executed: false,
					},
				},
			},
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, nil
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
			describe: func(in sequences.OnChainOutput, dep int) string {
				return "mock description"
			},
		},
		*/
		{
			desc:  "validation error in resolveInput",
			input: sequences.OnChainOutput{},
			resolveInput: func(e deployment.Environment, cfg sequences.OnChainOutput) (sequences.OnChainOutput, error) {
				return cfg, errors.New("")
			},
			resolveDep: func(e deployment.Environment, cfg sequences.OnChainOutput) (int, error) {
				return 0, nil
			},
			describe: func(in sequences.OnChainOutput, dep int) string {
				return "mock description"
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
			describe: func(in sequences.OnChainOutput, dep int) string {
				return "mock description"
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
			e := deployment.Environment{OperationsBundle: bundle}

			changeset := changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[sequences.OnChainOutput, int, sequences.OnChainOutput]{
				Sequence:     mockSequence,
				ResolveInput: test.resolveInput,
				ResolveDep:   test.resolveDep,
				Describe:     test.describe,
			})

			var expectErr bool
			if _, err := test.resolveInput(e, test.input); err != nil {
				expectErr = true
			}
			if _, err := test.resolveDep(e, test.input); err != nil {
				expectErr = true
			}

			err = changeset.VerifyPreconditions(e, test.input)
			if expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			out, err := changeset.Apply(e, test.input)
			require.NoError(t, err)

			dsAddrs, err := out.DataStore.Addresses().Fetch()
			require.Len(t, dsAddrs, len(test.input.Addresses))

			var someNotExecuted bool
			for _, w := range test.input.Writes {
				if !w.Executed {
					someNotExecuted = true
					break
				}
			}
			if len(test.input.Writes) > 0 && someNotExecuted {
				require.Len(t, out.MCMSTimelockProposals, 1)
				require.Len(t, out.MCMSTimelockProposals[0].Operations, 1)
				require.Equal(t, test.describe(test.input, 0), out.MCMSTimelockProposals[0].Description)
			}
		})
	}
}
