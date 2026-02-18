package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	router_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	onramp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

const testChainSel = uint64(5009297550715157269)

var (
	expectedWETH = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	expectedLINK = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
)

func newTestEnv(t *testing.T) *deployment.Environment {
	t.Helper()
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSel}),
	)
	require.NoError(t, err)
	return e
}

type testContracts struct {
	routerAddr  common.Address
	onRampAddr  common.Address
}

func deployTestContracts(t *testing.T, chain evm.Chain) testContracts {
	t.Helper()

	routerAddr, tx, _, err := router.DeployRouter(
		chain.DeployerKey, chain.Client,
		expectedWETH,
		common.HexToAddress("0x01"),
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	onRampAddr, tx, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		chain.DeployerKey, chain.Client,
		evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
			LinkToken:          expectedLINK,
			ChainSelector:      chain.Selector,
			DestChainSelector:  100,
			DefaultTxGasLimit:  200_000,
			MaxNopFeesJuels:    big.NewInt(0),
			PrevOnRamp:         common.Address{},
			RmnProxy:           common.HexToAddress("0x02"),
			TokenAdminRegistry: common.HexToAddress("0x03"),
		},
		evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
			Router:            routerAddr,
			PriceRegistry:     common.HexToAddress("0x04"),
			MaxDataBytes:      30_000,
			MaxPerMsgGasLimit: 3_000_000,
		},
		evm_2_evm_onramp.RateLimiterConfig{
			IsEnabled: false,
			Capacity:  big.NewInt(0),
			Rate:      big.NewInt(0),
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{},
		[]evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{},
		[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	return testContracts{routerAddr: routerAddr, onRampAddr: onRampAddr}
}

func datastoreWithSourceContracts(t *testing.T, tc testContracts) *datastore.MemoryDataStore {
	t.Helper()
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: testChainSel,
		Type:          datastore.ContractType(router_ops.ContractType),
		Version:       router_ops.Version,
		Address:       tc.routerAddr.Hex(),
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: testChainSel,
		Type:          datastore.ContractType(onramp_ops.ContractType),
		Version:       onramp_ops.Version,
		Address:       tc.onRampAddr.Hex(),
	}))
	return ds
}

func TestDiscoverTokens_ValidatesInput(t *testing.T) {
	e := newTestEnv(t)

	tests := []struct {
		desc        string
		cfg         changesets.DiscoverTokensCfg
		expectedErr string
	}{
		{
			desc:        "rejects empty chain selectors",
			cfg:         changesets.DiscoverTokensCfg{ChainSelectors: nil},
			expectedErr: "at least one chain selector is required",
		},
		{
			desc:        "rejects unknown chain selector",
			cfg:         changesets.DiscoverTokensCfg{ChainSelectors: []uint64{99999}},
			expectedErr: "chain selector 99999 not found in environment EVM chains",
		},
		{
			desc:        "rejects duplicate chain selectors",
			cfg:         changesets.DiscoverTokensCfg{ChainSelectors: []uint64{testChainSel, testChainSel}},
			expectedErr: "duplicate chain selector",
		},
		{
			desc: "accepts valid chain selector",
			cfg:  changesets.DiscoverTokensCfg{ChainSelectors: []uint64{testChainSel}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := changesets.DiscoverTokens.VerifyPreconditions(*e, test.cfg)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDiscoverTokens_SkipsExistingTokens(t *testing.T) {
	e := newTestEnv(t)

	t.Run("skips both when both exist", func(t *testing.T) {
		ds := datastore.NewMemoryDataStore()
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: testChainSel,
			Type:          datastore.ContractType(weth.ContractType),
			Version:       weth.Version,
			Address:       expectedWETH.Hex(),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: testChainSel,
			Type:          datastore.ContractType(link_token.ContractType),
			Version:       link_token.Version,
			Address:       expectedLINK.Hex(),
		}))
		e.DataStore = ds.Seal()

		out, err := changesets.DiscoverTokens.Apply(*e, changesets.DiscoverTokensCfg{
			ChainSelectors: []uint64{testChainSel},
		})
		require.NoError(t, err)

		addrs, err := out.DataStore.Addresses().Fetch()
		require.NoError(t, err)
		require.Empty(t, addrs, "expected no new addresses when tokens already exist")
	})

	t.Run("skips WETH and discovers LINK when only WETH exists", func(t *testing.T) {
		chain := e.BlockChains.EVMChains()[testChainSel]
		tc := deployTestContracts(t, chain)
		ds := datastoreWithSourceContracts(t, tc)

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: testChainSel,
			Type:          datastore.ContractType(weth.ContractType),
			Version:       weth.Version,
			Address:       expectedWETH.Hex(),
		}))
		e.DataStore = ds.Seal()

		out, err := changesets.DiscoverTokens.Apply(*e, changesets.DiscoverTokensCfg{
			ChainSelectors: []uint64{testChainSel},
		})
		require.NoError(t, err)

		addrs, err := out.DataStore.Addresses().Fetch()
		require.NoError(t, err)
		require.Len(t, addrs, 1)
		require.Equal(t, datastore.ContractType(link_token.ContractType), addrs[0].Type)
		require.Equal(t, expectedLINK.Hex(), addrs[0].Address)
	})

	t.Run("skips LINK and discovers WETH when only LINK exists", func(t *testing.T) {
		chain := e.BlockChains.EVMChains()[testChainSel]
		tc := deployTestContracts(t, chain)
		ds := datastoreWithSourceContracts(t, tc)

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: testChainSel,
			Type:          datastore.ContractType(link_token.ContractType),
			Version:       link_token.Version,
			Address:       expectedLINK.Hex(),
		}))
		e.DataStore = ds.Seal()

		out, err := changesets.DiscoverTokens.Apply(*e, changesets.DiscoverTokensCfg{
			ChainSelectors: []uint64{testChainSel},
		})
		require.NoError(t, err)

		addrs, err := out.DataStore.Addresses().Fetch()
		require.NoError(t, err)
		require.Len(t, addrs, 1)
		require.Equal(t, datastore.ContractType(weth.ContractType), addrs[0].Type)
		require.Equal(t, expectedWETH.Hex(), addrs[0].Address)
	})
}

func TestDiscoverTokens_ErrorsWhenSourceContractsMissing(t *testing.T) {
	t.Run("errors when Router missing for WETH discovery", func(t *testing.T) {
		e := newTestEnv(t)
		e.DataStore = datastore.NewMemoryDataStore().Seal()

		_, err := changesets.DiscoverTokens.Apply(*e, changesets.DiscoverTokensCfg{
			ChainSelectors: []uint64{testChainSel},
		})
		require.ErrorContains(t, err, "cannot resolve Router")
	})

	t.Run("errors when EVM2EVMOnRamp missing for LINK discovery", func(t *testing.T) {
		e := newTestEnv(t)
		chain := e.BlockChains.EVMChains()[testChainSel]

		routerAddr, tx, _, err := router.DeployRouter(
			chain.DeployerKey, chain.Client,
			expectedWETH,
			common.HexToAddress("0x01"),
		)
		require.NoError(t, err)
		_, err = chain.Confirm(tx)
		require.NoError(t, err)

		ds := datastore.NewMemoryDataStore()
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: testChainSel,
			Type:          datastore.ContractType(router_ops.ContractType),
			Version:       router_ops.Version,
			Address:       routerAddr.Hex(),
		}))
		e.DataStore = ds.Seal()

		_, err = changesets.DiscoverTokens.Apply(*e, changesets.DiscoverTokensCfg{
			ChainSelectors: []uint64{testChainSel},
		})
		require.ErrorContains(t, err, "cannot resolve EVM2EVMOnRamp")
	})
}

func TestDiscoverTokens_DiscoversTokensFromOnChainContracts(t *testing.T) {
	e := newTestEnv(t)
	chain := e.BlockChains.EVMChains()[testChainSel]
	tc := deployTestContracts(t, chain)
	ds := datastoreWithSourceContracts(t, tc)
	e.DataStore = ds.Seal()

	out, err := changesets.DiscoverTokens.Apply(*e, changesets.DiscoverTokensCfg{
		ChainSelectors: []uint64{testChainSel},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, addrs, 2, "expected exactly WETH and LINK addresses")

	addrByType := make(map[datastore.ContractType]datastore.AddressRef)
	for _, ref := range addrs {
		addrByType[ref.Type] = ref
	}

	wethRef, ok := addrByType[datastore.ContractType(weth.ContractType)]
	require.True(t, ok, "WETH9 should be discovered")
	require.Equal(t, expectedWETH.Hex(), wethRef.Address)
	require.Equal(t, weth.Version.String(), wethRef.Version.String())
	require.Equal(t, testChainSel, wethRef.ChainSelector)

	linkRef, ok := addrByType[datastore.ContractType(link_token.ContractType)]
	require.True(t, ok, "LinkToken should be discovered")
	require.Equal(t, expectedLINK.Hex(), linkRef.Address)
	require.Equal(t, link_token.Version.String(), linkRef.Version.String())
	require.Equal(t, testChainSel, linkRef.ChainSelector)
}
