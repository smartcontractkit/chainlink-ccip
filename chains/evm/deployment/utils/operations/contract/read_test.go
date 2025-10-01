package contract

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	tests := []struct {
		desc        string
		input       FunctionInput[int]
		expectedErr string
	}{
		{
			desc: "valid even input",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
		},
		{
			desc: "invalid odd input",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          3,
			},
			expectedErr: "odd value: 3",
		},
		{
			desc: "mismatched chain selector",
			input: FunctionInput[int]{
				ChainSelector: invalidChainSel,
				Address:       address,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "empty address",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Args:          2,
			},
			expectedErr: "address must be specified for test-read",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			read := NewRead(ReadParams[int, string, *testContract]{
				Name:         "test-read",
				Version:      semver.MustParse("1.0.0"),
				Description:  "Test read operation",
				ContractType: testContractType,
				NewContract:  newTestContract,
				CallContract: func(contract *testContract, opts *bind.CallOpts, input int) (string, error) {
					return contract.Read(opts, input)
				},
			})

			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")

			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)

			chain := evm.Chain{
				Selector: validChainSel,
			}

			report, err := operations.ExecuteOperation(bundle, read, chain, test.input)
			if test.expectedErr != "" {
				require.Error(t, err, "Expected ExecuteOperation error but got none")
				require.Contains(t, test.expectedErr, err.Error())
			} else {
				require.NoError(t, err, "Unexpected ExecuteOperation error")
				require.Equal(t, report.Output, "even")
			}
		})
	}
}
