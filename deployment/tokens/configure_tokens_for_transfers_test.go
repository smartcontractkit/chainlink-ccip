package tokens_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type MockReader struct{}

const OP_COUNT = 42

func (m *MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		StartingOpCount: OP_COUNT,
	}, nil
}

func (m *MockReader) GetTimelockRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0x01",
		Type:          datastore.ContractType("Timelock"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func (m *MockReader) GetMCMSRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0x02",
		Type:          datastore.ContractType("MCM"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

var transfersTest_NewTokenAdapterRegistry = tokens.GetTokenAdapterRegistry()

type transfersTest_MockTokenAdapter struct {
	deriveTokenErrorMsg  string
	sequenceErrorMsg     string
	deriveFailQualifiers map[string]struct{}
}

func (ma *transfersTest_MockTokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte(ref.Address), nil
}

func (ma *transfersTest_MockTokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-token-sequence",
		semver.MustParse("1.0.0"),
		"Mock sequence for testing token configuration",
		func(bundle cldf_ops.Bundle, deps cldf_chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			if ma.sequenceErrorMsg != "" {
				return sequences.OnChainOutput{}, errors.New(ma.sequenceErrorMsg)
			}
			batchOps := make([]mcms_types.BatchOperation, 0, len(input.RemoteChains))
			for remoteChainSel, remoteChain := range input.RemoteChains {
				batchOps = append(batchOps, mcms_types.BatchOperation{
					ChainSelector: mcms_types.ChainSelector(remoteChainSel),
					Transactions: []mcms_types.Transaction{
						{
							To:               string(remoteChain.RemotePool),
							Data:             remoteChain.RemoteToken,
							AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
						},
					},
				})
			}
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: input.ChainSelector,
						Address:       input.RegistryAddress,
						Type:          datastore.ContractType("Registry"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Type:          datastore.ContractType("TokenPool"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.ExternalAdmin,
						Type:          datastore.ContractType("EOA"),
						Version:       semver.MustParse("1.0.0"),
					},
				},
				BatchOps: batchOps,
			}, nil
		},
	)
}

func (r *transfersTest_MockTokenAdapter) ResolveTokenPoolRef(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, _ uint64, _ string) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, errors.New("unexpected pool resolve in transfers test")
}

func (r *transfersTest_MockTokenAdapter) ResolveTokenRef(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, chainSelector uint64, address string) (datastore.AddressRef, error) {
	const derived = "0x1111111111111111111111111111111111111111"

	if address == derived {
		return datastore.AddressRef{
			ChainSelector: chainSelector,
			Address:       address,
			Type:          datastore.ContractType("Token"),
			Version:       semver.MustParse("1.0.0"),
		}, nil
	}

	return datastore.AddressRef{}, fmt.Errorf("unexpected address: %s", address)
}

func (ma *transfersTest_MockTokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) (string, error) {
	if len(ma.deriveFailQualifiers) > 0 {
		if _, blocked := ma.deriveFailQualifiers[poolRef.Qualifier]; blocked {
			msg := ma.deriveTokenErrorMsg
			if msg == "" {
				msg = "failed to derive remote token address"
			}
			return "", errors.New(msg)
		}
		return "0x1111111111111111111111111111111111111111", nil
	}
	if ma.deriveTokenErrorMsg != "" {
		return "", errors.New(ma.deriveTokenErrorMsg)
	}

	return "0x1111111111111111111111111111111111111111", nil
}

func (ma *transfersTest_MockTokenAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	return 18, nil
}

func (ma *transfersTest_MockTokenAdapter) DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
	return tokenPool, nil
}

func (ma *transfersTest_MockTokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokens.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *transfersTest_MockTokenAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *transfersTest_MockTokenAdapter) DeployToken() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *transfersTest_MockTokenAdapter) DeployTokenVerify(e deployment.Environment, in tokens.DeployTokenInput) error {
	return nil
}

func (ma *transfersTest_MockTokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *transfersTest_MockTokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	return &cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment]{}
}

func (ma *transfersTest_MockTokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokens.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

var basicMCMSInput = mcms.Input{
	OverridePreviousRoot: true,
	ValidUntil:           3759765795,
	TimelockDelay:        mcms_types.MustParseDuration("1h"),
	TimelockAction:       mcms_types.TimelockActionSchedule,
}

func makeBaseDataStore(t *testing.T, chains []uint64) *datastore.MemoryDataStore {
	ds := datastore.NewMemoryDataStore()

	for _, chain := range chains {
		err := ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-token-pool", chain),
			Type:          datastore.ContractType("TokenPool"),
			Version:       semver.MustParse("1.0.0"),
			Qualifier:     "default",
		})
		require.NoError(t, err)
		ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-registry", chain),
			Type:          datastore.ContractType("Registry"),
			Version:       semver.MustParse("1.0.0"),
		})
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-token", chain),
			Type:          datastore.ContractType("Token"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-mcm", chain),
			Type:          datastore.ContractType("MCM"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-timelock", chain),
			Type:          datastore.ContractType("Timelock"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
	}

	return ds
}

// makeTwoChainDSWithRemoteEdgePools is like makeBaseDataStore for the two standard test chains but adds
// a second TokenPool per chain (Address "%d-remote-token-pool", Qualifier "remote-edge").
func makeTwoChainDSWithRemoteEdgePools(t *testing.T) *datastore.MemoryDataStore {
	t.Helper()
	const (
		chainA = 5009297550715157269
		chainB = 15971525489660198786
	)
	ds := makeBaseDataStore(t, []uint64{chainA, chainB})
	for _, chain := range []uint64{chainA, chainB} {
		err := ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-remote-token-pool", chain),
			Type:          datastore.ContractType("TokenPool"),
			Version:       semver.MustParse("1.0.0"),
			Qualifier:     "remote-edge",
		})
		require.NoError(t, err)
	}
	return ds
}

func TestConfigureTokensForTransfers_Apply(t *testing.T) {
	tests := []struct {
		desc                            string
		makeDataStore                   func(t *testing.T) *datastore.MemoryDataStore
		cfg                             tokens.ConfigureTokensForTransfersConfig
		shouldDeriveToken               bool
		expectedSequenceErrorMsg        string
		expectedTokenDerivationErrorMsg string
		tokenPoolDeriveFailQualifiers   map[string]struct{}
	}{
		{
			desc: "success - inputted remote token",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeTwoChainDSWithRemoteEdgePools(t)
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
							Qualifier:     "default",
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
									Address:       "15971525489660198786-remote-token-pool",
									Qualifier:     "remote-edge",
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{
									IsEnabled: true,
									Capacity:  1000,
									Rate:      100,
								},
							},
						},
					},
					{
						ChainSelector: 15971525489660198786,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
							Qualifier:     "default",
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							5009297550715157269: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
									Address:       "5009297550715157269-remote-token-pool",
									Qualifier:     "remote-edge",
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{
									IsEnabled: true,
									Capacity:  1000,
									Rate:      100,
								},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			tokenPoolDeriveFailQualifiers: map[string]struct{}{"remote-edge": {}},
		},
		{
			desc: "success - derived remote token",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: nil, // This will trigger derivation
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{
									IsEnabled: true,
									Capacity:  1000,
									Rate:      100,
								},
							},
						},
					},
					{
						ChainSelector: 15971525489660198786,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							5009297550715157269: {
								RemoteToken: nil, // This will trigger derivation
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{
									IsEnabled: true,
									Capacity:  1000,
									Rate:      100,
								},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			shouldDeriveToken: true,
		},
		{
			desc: "failure to resolve token pool ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Don't add token pool address - only add registry
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-registry",
					Type:          datastore.ContractType("Registry"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve token pool ref",
		},
		{
			desc: "failure to resolve registry ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Only add token pool, not registry
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-token-pool",
					Type:          datastore.ContractType("TokenPool"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve registry ref",
		},
		{
			desc: "invalid chain selector",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 0, // Invalid chain selector
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "unknown chain selector",
		},
		{
			desc: "failure to resolve remote pool ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseDataStore(t, []uint64{5009297550715157269})
				// Don't add remote chain data
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
					{
						ChainSelector: 15971525489660198786,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							5009297550715157269: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			// Datastore has no addresses for chain 15971525489660198786; error may be "failed to resolve token pool ref" or "failed to resolve remote pool ref" depending on map iteration order.
			expectedSequenceErrorMsg: "failed to resolve",
		},
		{
			desc: "failure to resolve remote token ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
				// Remove remote token
				err := ds.Addresses().Delete(datastore.NewAddressRefKey(15971525489660198786, "Token", semver.MustParse("1.0.0"), ""))
				require.NoError(t, err)
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
					{
						ChainSelector: 15971525489660198786,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							5009297550715157269: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve remote token after derivation",
		},
		{
			desc: "failure to derive remote token address",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeTwoChainDSWithRemoteEdgePools(t)
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
							Qualifier:     "default",
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: nil, // Will trigger derivation which should fail
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
									Address:       "15971525489660198786-remote-token-pool",
									Qualifier:     "remote-edge",
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
					{
						ChainSelector: 15971525489660198786,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
							Qualifier:     "default",
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							5009297550715157269: {
								RemoteToken: nil, // Will trigger derivation which should fail
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
									Address:       "5009297550715157269-remote-token-pool",
									Qualifier:     "remote-edge",
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedTokenDerivationErrorMsg: "failed to derive remote token address",
			tokenPoolDeriveFailQualifiers:   map[string]struct{}{"remote-edge": {}},
		},
		{
			desc: "failure to execute sequence",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
					{
						ChainSelector: 15971525489660198786,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 15971525489660198786,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							5009297550715157269: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								OutboundRateLimiterConfig: &tokens.RateLimiterConfigFloatInput{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "sequence execution failed",
		},
	}
	// Register MCMS reader
	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader("evm", &MockReader{})
	tokenAdapterRegistry := tokens.GetTokenAdapterRegistry()
	mockAdapter := &transfersTest_MockTokenAdapter{}
	tokenAdapterRegistry.RegisterTokenAdapter("evm", semver.MustParse("1.0.0"), mockAdapter)
	tokenAdapterRegistry.RegisterTokenRefResolver("evm", mockAdapter)

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			// Register token adapter with appropriate error condition

			// Set error conditions for specific test cases
			mockAdapter.deriveTokenErrorMsg = tt.expectedTokenDerivationErrorMsg
			mockAdapter.sequenceErrorMsg = tt.expectedSequenceErrorMsg
			mockAdapter.deriveFailQualifiers = tt.tokenPoolDeriveFailQualifiers

			// Create environment with datastore
			ds := tt.makeDataStore(t)
			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")
			bundle := cldf_ops.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				cldf_ops.NewMemoryReporter(),
			)
			e := deployment.Environment{
				OperationsBundle: bundle,
				Logger:           lggr,
				DataStore:        ds.Seal(),
			}

			// Run changeset
			changeset := tokens.ConfigureTokensForTransfers(tokenAdapterRegistry, mcmsRegistry)
			out, err := changeset.Apply(e, tt.cfg)
			if tt.expectedSequenceErrorMsg != "" {
				require.ErrorContains(t, err, tt.expectedSequenceErrorMsg)
				return
			}
			if tt.expectedTokenDerivationErrorMsg != "" {
				require.ErrorContains(t, err, tt.expectedTokenDerivationErrorMsg)
				return
			}
			require.NoError(t, err)

			// Validate output
			require.Len(t, out.MCMSTimelockProposals, 1)
			for _, proposal := range out.MCMSTimelockProposals {
				require.Equal(t, tt.cfg.MCMS.Description, proposal.Description)
				require.Equal(t, tt.cfg.MCMS.OverridePreviousRoot, proposal.OverridePreviousRoot)
				require.Equal(t, tt.cfg.MCMS.ValidUntil, proposal.ValidUntil)
				require.Equal(t, tt.cfg.MCMS.TimelockDelay, proposal.Delay)
				require.Equal(t, tt.cfg.MCMS.TimelockAction, proposal.Action)

				require.Len(t, proposal.Operations, len(tt.cfg.Tokens))
				for _, op := range proposal.Operations {
					// For derived remote token test, expect mocked address
					if tt.shouldDeriveToken {
						require.Equal(t, []byte("0x1111111111111111111111111111111111111111"), op.Transactions[0].Data)
					} else {
						require.Equal(t, common.LeftPadBytes([]byte(fmt.Sprintf("%d-token", op.ChainSelector)), 32), op.Transactions[0].Data)
					}
					if len(tt.tokenPoolDeriveFailQualifiers) > 0 {
						require.Equal(t, fmt.Sprintf("%d-remote-token-pool", op.ChainSelector), op.Transactions[0].To)
					} else {
						require.Equal(t, fmt.Sprintf("%d-token-pool", op.ChainSelector), op.Transactions[0].To)
					}
				}
			}
		})
	}
}

func TestRemoteOutbounds_DefaultBucket_legacyRateLimitAlias(t *testing.T) {
	rl := tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}
	ro := tokens.RemoteOutbounds{RateLimit: &rl}
	require.NoError(t, ro.Validate())

	d, dOk := ro.DefaultBucket()
	require.True(t, dOk)
	_, fOk := ro.FastFinalityBucket()
	require.False(t, fOk)
	require.Equal(t, rl, d.RateLimit)
	require.False(t, d.FastFinality)
}

func TestRemoteOutbounds_DefaultAndFastFinality_slices(t *testing.T) {
	ro := tokens.RemoteOutbounds{
		Outbounds: []tokens.RateLimitConfig{
			{RateLimit: tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 50, Rate: 5}, FastFinality: false},
			{RateLimit: tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 200, Rate: 20}, FastFinality: true},
		},
	}
	require.NoError(t, ro.Validate())

	d, dOk := ro.DefaultBucket()
	f, fOk := ro.FastFinalityBucket()
	require.True(t, dOk)
	require.True(t, fOk)
	require.NoError(t, d.RateLimit.Validate())
	require.NoError(t, f.RateLimit.Validate())
	require.Equal(t, ro.Outbounds[0].RateLimit, d.RateLimit)
	require.Equal(t, ro.Outbounds[1].RateLimit, f.RateLimit)
	require.False(t, d.FastFinality)
	require.True(t, f.FastFinality)
}

func TestRemoteOutbounds_FastFinalitySliceWithLegacyAlias(t *testing.T) {
	ro := tokens.RemoteOutbounds{
		RateLimit: &tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 7010, Rate: 701},
		Outbounds: []tokens.RateLimitConfig{
			{
				RateLimit:    tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 6060, Rate: 606},
				FastFinality: true,
			},
		},
	}
	require.NoError(t, ro.Validate())

	d, dOk := ro.DefaultBucket()
	f, fOk := ro.FastFinalityBucket()
	require.True(t, dOk)
	require.True(t, fOk)
	require.NoError(t, d.RateLimit.Validate())
	require.NoError(t, f.RateLimit.Validate())
	require.Equal(t, *ro.RateLimit, d.RateLimit)
	require.Equal(t, ro.Outbounds[0].RateLimit, f.RateLimit)
	require.False(t, d.FastFinality)
	require.True(t, f.FastFinality)
}

func TestRemoteChainConfig_GetOutboundInboundRateLimitBuckets(t *testing.T) {
	cfg := tokens.RemoteChainConfig[[]byte, string]{
		OutboundRateLimits: []tokens.RateLimitConfig{
			{RateLimit: tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1, Rate: 1}, FastFinality: false},
			{RateLimit: tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2, Rate: 2}, FastFinality: true},
		},
		InboundRateLimits: []tokens.RateLimitConfig{
			{RateLimit: tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3, Rate: 3}, FastFinality: false},
			{RateLimit: tokens.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 4, Rate: 4}, FastFinality: true},
		},
	}

	ffOB, ok := cfg.GetOutboundRateLimitBuckets().FastFinalityBucket()
	require.True(t, ok)
	require.Equal(t, 2.0, ffOB.RateLimit.Capacity)

	ffIB, ok := cfg.GetInboundRateLimitBuckets().FastFinalityBucket()
	require.True(t, ok)
	require.Equal(t, 4.0, ffIB.RateLimit.Capacity)
}
