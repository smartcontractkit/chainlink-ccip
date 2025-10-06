package tokens_test

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type MockReader struct{}

const OP_COUNT = 42

func (m *MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		MCMAddress:      input.MCMSAddressRef.Address,
		StartingOpCount: OP_COUNT,
	}, nil
}

var transfersTest_NewTokenAdapterRegistry = tokens.NewTokenAdapterRegistry()

type transfersTest_MockTokenAdapter struct {
	shouldError bool
	errorMsg    string
}

func (ma *transfersTest_MockTokenAdapter) ConvertRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if ma.shouldError {
		return nil, errors.New(ma.errorMsg)
	}
	return []byte("mocked-address-bytes"), nil
}

func (ma *transfersTest_MockTokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-token-sequence",
		semver.MustParse("1.0.0"),
		"Mock sequence for testing token configuration",
		func(bundle cldf_ops.Bundle, deps cldf_chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			if ma.shouldError {
				return sequences.OnChainOutput{}, errors.New(ma.errorMsg)
			}
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{},
				BatchOps: []mcms_types.BatchOperation{{
					ChainSelector: mcms_types.ChainSelector(input.ChainSelector),
					Transactions:  []mcms_types.Transaction{},
				}},
			}, nil
		},
	)
}

func (ma *transfersTest_MockTokenAdapter) DeriveRemoteTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	if ma.shouldError {
		return nil, errors.New(ma.errorMsg)
	}
	return []byte("mocked-remote-token-address"), nil
}

var basicMCMSInput = mcms.Input{
	OverridePreviousRoot: true,
	ValidUntil:           3759765795,
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
}

func makeBaseDataStore() *datastore.MemoryDataStore {
	ds := datastore.NewMemoryDataStore()
	// Chain #1
	ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 5009297550715157269,
		Address:       common.HexToAddress("0x01").Hex(),
		Type:          datastore.ContractType("TokenPool"),
		Version:       semver.MustParse("1.0.0"),
	})
	ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 5009297550715157269,
		Address:       common.HexToAddress("0x02").Hex(),
		Type:          datastore.ContractType("Registry"),
		Version:       semver.MustParse("1.0.0"),
	})
	ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 5009297550715157269,
		Address:       common.HexToAddress("0x03").Hex(),
		Type:          datastore.ContractType("Token"),
		Version:       semver.MustParse("1.0.0"),
	})
	// Chain #2
	ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 15971525489660198786,
		Address:       common.HexToAddress("0x04").Hex(),
		Type:          datastore.ContractType("TokenPool"),
		Version:       semver.MustParse("1.0.0"),
	})
	ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 15971525489660198786,
		Address:       common.HexToAddress("0x05").Hex(),
		Type:          datastore.ContractType("Registry"),
		Version:       semver.MustParse("1.0.0"),
	})
	ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 15971525489660198786,
		Address:       common.HexToAddress("0x06").Hex(),
		Type:          datastore.ContractType("Token"),
		Version:       semver.MustParse("1.0.0"),
	})

	return ds
}

func TestConfigureTokensForTransfers(t *testing.T) {
	tests := []struct {
		desc          string
		makeDataStore func(t *testing.T) *datastore.MemoryDataStore
		cfg           tokens.ConfigureTokensForTransfersConfig
	}{
		{
			desc: "success - inputted remote token",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore()
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{},
				MCMS:   mcms.Input{},
			},
		},
		{
			desc: "success - derived remote token",
		},
		{
			desc: "failure to resolve token pool ref",
		},
		{
			desc: "failure to resolve registry ref",
		},
		{
			desc: "invalid chain selector",
		},
		{
			desc: "no token adapter registered",
		},
		{
			desc: "failure to resolve remote pool ref",
		},
		{
			desc: "failure to resolve remote token ref",
		},
		{
			desc: "failure to derive remote token address",
		},
		{
			desc: "failure to resolve outbound CCV ref",
		},
		{
			desc: "failure to resolve inbound CCV ref",
		},
		{
			desc: "failure to execute sequence",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			// Register MCMS reader
			mcmsRegistry := changesets.NewMCMSReaderRegistry()
			mcmsRegistry.RegisterMCMSReader("evm", &MockReader{})

			// Register token adapter
			tokenAdapterRegistry := tokens.NewTokenAdapterRegistry()
			tokenAdapterRegistry.RegisterTokenAdapter("evm", semver.MustParse("1.0.0"), &transfersTest_MockTokenAdapter{})

			// Create environment with datastore
			ds := tt.makeDataStore(t)
			e := deployment.Environment{
				DataStore: ds.Seal(),
			}

			// Create changeset
			changeset := tokens.ConfigureTokensForTransfers(tokenAdapterRegistry, mcmsRegistry)
			changeset.VerifyPreconditions(e, tokens.ConfigureTokensForTransfersConfig{})
			changeset.Apply(e, tokens.ConfigureTokensForTransfersConfig{})
		})
	}
}
