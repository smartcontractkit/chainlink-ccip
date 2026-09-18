package changesets

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
)

var emptyBlockChains = cldf_chain.NewBlockChains(nil)

func TestApplyCCTPDefaults_FillsEmptyCanonicalChain(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Equal(t, "0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5", got.Chains[chainSel].TokenMessengerV1)
	require.Equal(t, "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA", got.Chains[chainSel].TokenMessengerV2)
	require.Equal(t, "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238", got.Chains[chainSel].USDCToken)
	// TokenDecimals is no longer defaulted by applyCCTPDefaults: it is only consumed by the
	// non-canonical deploy and resolved on-chain (see makeApplyDeployCCTPChains).
	require.Zero(t, got.Chains[chainSel].TokenDecimals)
}

func TestApplyCCTPDefaults_ExplicitTokenDecimalsTakePrecedence(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical, TokenDecimals: 7},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Equal(t, uint8(7), got.Chains[chainSel].TokenDecimals)
}

func TestApplyCCTPDefaults_ExplicitValuesTakePrecedence(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {
				USDCType:         adapters.Canonical,
				TokenMessengerV1: "0x1111111111111111111111111111111111111111",
				TokenMessengerV2: "0x2222222222222222222222222222222222222222",
				USDCToken:        "0x3333333333333333333333333333333333333333",
			},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Equal(t, "0x1111111111111111111111111111111111111111", got.Chains[chainSel].TokenMessengerV1)
	require.Equal(t, "0x2222222222222222222222222222222222222222", got.Chains[chainSel].TokenMessengerV2)
	require.Equal(t, "0x3333333333333333333333333333333333333333", got.Chains[chainSel].USDCToken)
}

func TestApplyCCTPDefaults_NonCanonicalIsUntouched(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.NonCanonical},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Empty(t, got.Chains[chainSel].TokenMessengerV1)
	require.Empty(t, got.Chains[chainSel].TokenMessengerV2)
	require.Empty(t, got.Chains[chainSel].USDCToken)
	require.Empty(t, got.Chains[chainSel].RemoteChains)
	require.Zero(t, got.Chains[chainSel].TokenDecimals)
}

func TestApplyCCTPDefaults_NonCanonicalRemoteDomainIsDefaulted(t *testing.T) {
	localSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	remoteSel := chain_selectors.AVALANCHE_TESTNET_FUJI.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			localSel: {
				USDCType: adapters.NonCanonical,
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					remoteSel: {},
				},
			},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	// Address defaults must not be applied for non-canonical chains, but the
	// remote domain identifier must still be resolved.
	require.Empty(t, got.Chains[localSel].USDCToken)
	require.Equal(t, uint32(1), got.Chains[localSel].RemoteChains[remoteSel].DomainIdentifier)
}

func TestApplyCCTPDefaults_UnknownChainIsUntouched(t *testing.T) {
	const chainSel uint64 = 1234567890
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Empty(t, got.Chains[chainSel].TokenMessengerV1)
	require.Empty(t, got.Chains[chainSel].TokenMessengerV2)
	require.Empty(t, got.Chains[chainSel].USDCToken)
}

func TestApplyCCTPDefaults_FillsRemoteDomainIdentifier(t *testing.T) {
	localSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	remoteSel := chain_selectors.AVALANCHE_TESTNET_FUJI.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			localSel: {
				USDCType: adapters.Canonical,
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					remoteSel: {},
				},
			},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Equal(t, uint32(1), got.Chains[localSel].RemoteChains[remoteSel].DomainIdentifier)
}

func TestApplyCCTPDefaults_ResolvesDeployerContractFromDatastore(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType("CREATE2Factory"),
		Version:       semver.MustParse("2.0.0"),
		Address:       "0x000000000000000000000000000000000000c2f2",
	}))
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, ds.Seal(), cfg)

	require.Equal(t, "0x000000000000000000000000000000000000c2f2", got.Chains[chainSel].DeployerContract)
}

func TestApplyCCTPDefaults_ExplicitDeployerContractTakesPrecedence(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType("CREATE2Factory"),
		Version:       semver.MustParse("2.0.0"),
		Address:       "0x000000000000000000000000000000000000c2f2",
	}))
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical, DeployerContract: "0x1111111111111111111111111111111111111111"},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, ds.Seal(), cfg)

	require.Equal(t, "0x1111111111111111111111111111111111111111", got.Chains[chainSel].DeployerContract)
}

func TestApplyCCTPDefaults_MissingDeployerContractLeavesItEmpty(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType("CREATE2Factory"),
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x000000000000000000000000000000000000c2f2",
	}))
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, ds.Seal(), cfg)

	require.Empty(t, got.Chains[chainSel].DeployerContract)
}

func TestApplyCCTPDefaults_ExplicitRemoteDomainIdentifierTakesPrecedence(t *testing.T) {
	localSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	remoteSel := chain_selectors.AVALANCHE_TESTNET_FUJI.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			localSel: {
				USDCType: adapters.Canonical,
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					remoteSel: {DomainIdentifier: 99},
				},
			},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	require.Equal(t, uint32(99), got.Chains[localSel].RemoteChains[remoteSel].DomainIdentifier)
}

func TestApplyCCTPDefaults_DefaultFeeAggregatorFromDeployerKey(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	deployer := common.HexToAddress("0x1111111111111111111111111111111111111111")
	blockChains := cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{
		chainSel: &evm.Chain{Selector: chainSel, DeployerKey: &bind.TransactOpts{From: deployer}},
	})
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical},
		},
	}

	got := applyCCTPDefaults(blockChains, nil, cfg)

	require.Equal(t, deployer.Hex(), got.Chains[chainSel].FeeAggregator)
}

func TestApplyCCTPDefaults_FeeAggregatorUnsetWithoutDeployerKey(t *testing.T) {
	chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	cfg := DeployCCTPChainsConfig{
		Chains: map[uint64]CCTPChainConfig{
			chainSel: {USDCType: adapters.Canonical},
		},
	}

	got := applyCCTPDefaults(emptyBlockChains, nil, cfg)

	// No deployer key in the environment => FeeAggregator stays unset, no error.
	require.Empty(t, got.Chains[chainSel].FeeAggregator)
}
