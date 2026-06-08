package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	onrampv1_5ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	onrampv1_6ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	fq20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	evm_2_evm_onramp_v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	fq20binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
)

const (
	removeFeeTokensChainSel = uint64(5009297550715157269)
	remoteChainV15          = uint64(111)
	remoteChainV16          = uint64(222)
)

var (
	legacyLinkToken = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	legacyWethToken = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	extraFeeToken1  = common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")
	extraFeeToken2  = common.HexToAddress("0xDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")
)

func TestRemoveFeeTokens_RemovesExtraFeeTokensFromFeeQuoter20(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{removeFeeTokensChainSel}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[removeFeeTokensChainSel]
	fq20Addr := deployRemoveFeeTokensFixture(t, e)

	fq20, err := fq20binding.NewFeeQuoter(fq20Addr, chain.Client)
	require.NoError(t, err)

	beforeTokens, err := fq20.GetFeeTokens(nil)
	require.NoError(t, err)
	require.True(t, containsAllAddresses(beforeTokens, legacyLinkToken, legacyWethToken, extraFeeToken1, extraFeeToken2))

	mcmsRegistry := cs_core.GetRegistry()
	_, err = changesets.RemoveFeeTokens(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.RemoveFeeTokensCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.RemoveFeeTokensCfg{
			ChainSels: []uint64{removeFeeTokensChainSel},
		},
	})
	require.NoError(t, err)

	afterTokens, err := fq20.GetFeeTokens(nil)
	require.NoError(t, err)
	require.True(t, containsAllAddresses(afterTokens, legacyLinkToken, legacyWethToken))
	require.False(t, containsAnyAddress(afterTokens, extraFeeToken1, extraFeeToken2))
}

func deployRemoveFeeTokensFixture(t *testing.T, e *cldf.Environment) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[removeFeeTokensChainSel]
	ds := datastore.NewMemoryDataStore()

	rmnProxy := common.HexToAddress("0x7777777777777777777777777777777777777777")
	placeholder := common.HexToAddress("0x8888888888888888888888888888888888888888")

	routerOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, routerops.Deploy, chain, contract.DeployInput[routerops.ConstructorArgs]{
		ChainSelector:  removeFeeTokensChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(routerops.ContractType, *routerops.Version),
		Args: routerops.ConstructorArgs{
			WrappedNative: legacyWethToken,
			RMNProxy:      rmnProxy,
		},
	})
	require.NoError(t, err)
	routerAddr := common.HexToAddress(routerOut.Output.Address)
	require.NoError(t, ds.Addresses().Add(routerOut.Output))

	prAddr, tx, _, err := price_registry.DeployPriceRegistry(
		chain.DeployerKey, chain.Client,
		[]common.Address{chain.DeployerKey.From},
		[]common.Address{legacyLinkToken, legacyWethToken},
		3600,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: removeFeeTokensChainSel,
		Type:          datastore.ContractType(priceregistryops.ContractType),
		Version:       priceregistryops.Version,
		Address:       prAddr.Hex(),
	}))

	fq16Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.Deploy, chain, contract.DeployInput[fq16ops.ConstructorArgs]{
		ChainSelector:  removeFeeTokensChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq16ops.ContractType, *fq16ops.Version),
		Args: fq16ops.ConstructorArgs{
			StaticConfig: fq16ops.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1e18),
				LinkToken:                    legacyLinkToken,
				TokenPriceStalenessThreshold: 3600,
			},
			PriceUpdaters: []common.Address{chain.DeployerKey.From},
			FeeTokens:     []common.Address{legacyLinkToken, legacyWethToken},
			PremiumMultiplierWeiPerEthArgs: []fq16ops.FeeTokenArgs{
				{Token: legacyLinkToken, PremiumMultiplierWeiPerEth: 9e17},
				{Token: legacyWethToken, PremiumMultiplierWeiPerEth: 1e18},
			},
		},
	})
	require.NoError(t, err)
	fq16Addr := common.HexToAddress(fq16Out.Output.Address)
	require.NoError(t, ds.Addresses().Add(fq16Out.Output))

	onRamp15Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onrampv1_5ops.DeployOnRamp, chain, contract.DeployInput[onrampv1_5ops.ConstructorArgs]{
		ChainSelector:  removeFeeTokensChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(onrampv1_5ops.ContractType, *onrampv1_5ops.Version),
		Args: onrampv1_5ops.ConstructorArgs{
			StaticConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampStaticConfig{
				LinkToken:          legacyLinkToken,
				ChainSelector:      removeFeeTokensChainSel,
				DestChainSelector:  remoteChainV15,
				DefaultTxGasLimit:  200_000,
				MaxNopFeesJuels:    big.NewInt(1e18),
				RmnProxy:           rmnProxy,
				TokenAdminRegistry: placeholder,
			},
			DynamicConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampDynamicConfig{
				Router:            routerAddr,
				PriceRegistry:     prAddr,
				MaxDataBytes:      30_000,
				MaxPerMsgGasLimit: 3_000_000,
			},
			RateLimiterConfig: evm_2_evm_onramp_v1_5_0.RateLimiterConfig{
				IsEnabled: false,
				Capacity:  big.NewInt(0),
				Rate:      big.NewInt(0),
			},
		},
	})
	require.NoError(t, err)
	onRamp15Addr := common.HexToAddress(onRamp15Out.Output.Address)

	onRamp16Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onrampv1_6ops.Deploy, chain, contract.DeployInput[onrampv1_6ops.ConstructorArgs]{
		ChainSelector:  removeFeeTokensChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(onrampv1_6ops.ContractType, *onrampv1_6ops.Version),
		Args: onrampv1_6ops.ConstructorArgs{
			StaticConfig: onrampv1_6ops.StaticConfig{
				ChainSelector:      removeFeeTokensChainSel,
				RmnRemote:          rmnProxy,
				NonceManager:       utils.RandomAddress(),
				TokenAdminRegistry: placeholder,
			},
			DynamicConfig: onrampv1_6ops.DynamicConfig{
				FeeQuoter:     fq16Addr,
				FeeAggregator: chain.DeployerKey.From,
				AllowlistAdmin: chain.DeployerKey.From,
			},
			DestChainConfigArgs: []onrampv1_6ops.DestChainConfigArgs{
				{
					DestChainSelector: remoteChainV16,
					Router:            routerAddr,
					AllowlistEnabled:  false,
				},
			},
		},
	})
	require.NoError(t, err)
	onRamp16Addr := common.HexToAddress(onRamp16Out.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, routerops.ApplyRampUpdates, chain, contract.FunctionInput[routerops.ApplyRampsUpdatesArgs]{
		ChainSelector: removeFeeTokensChainSel,
		Address:       routerAddr,
		Args: routerops.ApplyRampsUpdatesArgs{
			OnRampUpdates: []router.RouterOnRamp{
				{DestChainSelector: remoteChainV15, OnRamp: onRamp15Addr},
				{DestChainSelector: remoteChainV16, OnRamp: onRamp16Addr},
			},
			OffRampAdds: []router.RouterOffRamp{
				{SourceChainSelector: remoteChainV15, OffRamp: utils.RandomAddress()},
				{SourceChainSelector: remoteChainV16, OffRamp: utils.RandomAddress()},
			},
		},
	})
	require.NoError(t, err)

	fq20Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq20ops.Deploy, chain, contract.DeployInput[fq20ops.ConstructorArgs]{
		ChainSelector:  removeFeeTokensChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq20ops.ContractType, *fq20ops.Version),
		Args: fq20ops.ConstructorArgs{
			StaticConfig: fq20ops.StaticConfig{
				MaxFeeJuelsPerMsg: big.NewInt(1e18),
				LinkToken:         legacyLinkToken,
			},
			PriceUpdaters: []common.Address{chain.DeployerKey.From},
		},
	})
	require.NoError(t, err)
	fq20Addr := common.HexToAddress(fq20Out.Output.Address)
	require.NoError(t, ds.Addresses().Add(fq20Out.Output))

	tokenPrice := big.NewInt(1e18)
	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, fq20ops.UpdatePrices, chain, contract.FunctionInput[fq20ops.PriceUpdates]{
		ChainSelector: removeFeeTokensChainSel,
		Address:       fq20Addr,
		Args: fq20ops.PriceUpdates{
			TokenPriceUpdates: []fq20ops.TokenPriceUpdate{
				{SourceToken: legacyLinkToken, UsdPerToken: tokenPrice},
				{SourceToken: legacyWethToken, UsdPerToken: tokenPrice},
				{SourceToken: extraFeeToken1, UsdPerToken: tokenPrice},
				{SourceToken: extraFeeToken2, UsdPerToken: tokenPrice},
			},
		},
	})
	require.NoError(t, err)

	e.DataStore = ds.Seal()
	return fq20Addr
}

func containsAllAddresses(tokens []common.Address, expected ...common.Address) bool {
	set := make(map[common.Address]struct{}, len(tokens))
	for _, token := range tokens {
		set[token] = struct{}{}
	}
	for _, addr := range expected {
		if _, ok := set[addr]; !ok {
			return false
		}
	}
	return true
}

func containsAnyAddress(tokens []common.Address, candidates ...common.Address) bool {
	set := make(map[common.Address]struct{}, len(tokens))
	for _, token := range tokens {
		set[token] = struct{}{}
	}
	for _, addr := range candidates {
		if _, ok := set[addr]; ok {
			return true
		}
	}
	return false
}
