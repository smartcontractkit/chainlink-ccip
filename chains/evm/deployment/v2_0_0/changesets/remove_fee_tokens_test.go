package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	fq20binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	fq20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

const (
	removeFeeTokensFQ16ChainSel      = uint64(5009297550715157269)
	removeFeeTokensPRChainSel        = uint64(4356164186791070119)
	removeFeeTokensFQ16MultiChainSel = uint64(6433500567565415381)
)

var (
	legacyLinkToken = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	legacyWethToken = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	extraFeeToken1  = common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")
	extraFeeToken2  = common.HexToAddress("0xDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")
)

func TestRemoveFeeTokens_RemovesExtraFeeTokensFromFeeQuoter20(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{
			removeFeeTokensFQ16ChainSel,
			removeFeeTokensPRChainSel,
			removeFeeTokensFQ16MultiChainSel,
		}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	fq16Chain := e.BlockChains.EVMChains()[removeFeeTokensFQ16ChainSel]
	prChain := e.BlockChains.EVMChains()[removeFeeTokensPRChainSel]
	fq16MultiChain := e.BlockChains.EVMChains()[removeFeeTokensFQ16MultiChainSel]

	fq20FQ16Addr := deployRemoveFeeTokensFixtureWithFeeQuoter16(t, e, ds, removeFeeTokensFQ16ChainSel)
	fq20PRAddr := deployRemoveFeeTokensFixtureWithPriceRegistry(t, e, ds, removeFeeTokensPRChainSel)
	fq20FQ16MultiAddr := deployRemoveFeeTokensFixtureWithMultipleFeeQuoter16Versions(t, e, ds, removeFeeTokensFQ16MultiChainSel)
	e.DataStore = ds.Seal()

	fq20FQ16, err := fq20binding.NewFeeQuoter(fq20FQ16Addr, fq16Chain.Client)
	require.NoError(t, err)
	fq20PR, err := fq20binding.NewFeeQuoter(fq20PRAddr, prChain.Client)
	require.NoError(t, err)
	fq20FQ16Multi, err := fq20binding.NewFeeQuoter(fq20FQ16MultiAddr, fq16MultiChain.Client)
	require.NoError(t, err)

	fq20ByChain := []*fq20binding.FeeQuoter{fq20FQ16, fq20PR, fq20FQ16Multi}

	for _, fq20 := range fq20ByChain {
		beforeTokens, err := fq20.GetFeeTokens(nil)
		require.NoError(t, err)
		require.True(t, containsAllAddresses(beforeTokens, legacyLinkToken, legacyWethToken, extraFeeToken1, extraFeeToken2))
	}

	mcmsRegistry := cs_core.GetRegistry()
	_, err = changesets.RemoveFeeTokens(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.RemoveFeeTokensCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.RemoveFeeTokensCfg{
			ChainSels: []uint64{
				removeFeeTokensFQ16ChainSel,
				removeFeeTokensPRChainSel,
				removeFeeTokensFQ16MultiChainSel,
			},
		},
	})
	require.NoError(t, err)

	for _, fq20 := range fq20ByChain {
		afterTokens, err := fq20.GetFeeTokens(nil)
		require.NoError(t, err)
		require.True(t, containsAllAddresses(afterTokens, legacyLinkToken, legacyWethToken))
		require.False(t, containsAnyAddress(afterTokens, extraFeeToken1, extraFeeToken2))
	}

	hook := changesets.RemoveFeeTokensPostProposalHook()
	require.Equal(t, changesets.RemoveFeeTokensPostProposalHookName, hook.Name)
	err = hook.Func(t.Context(), cldf_changeset.PostProposalHookParams{
		Env: cldf_changeset.ProposalHookEnv{
			Name:        e.Name,
			Logger:      e.Logger,
			BlockChains: e.BlockChains,
			DataStore:   e.DataStore,
		},
		Config: changesets.RemoveFeeTokensCfg{
			ChainSels: []uint64{
				removeFeeTokensFQ16ChainSel,
				removeFeeTokensPRChainSel,
				removeFeeTokensFQ16MultiChainSel,
			},
		},
	})
	require.NoError(t, err)
}

func deployRemoveFeeTokensFixtureWithFeeQuoter16(
	t *testing.T,
	e *cldf.Environment,
	ds datastore.MutableDataStore,
	chainSel uint64,
) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSel]
	deployFeeQuoter160(t, e, ds, chain, chainSel, []common.Address{legacyLinkToken, legacyWethToken})

	return deployFeeQuoter20WithExtraTokens(t, e, ds, chain, chainSel)
}

func deployRemoveFeeTokensFixtureWithMultipleFeeQuoter16Versions(
	t *testing.T,
	e *cldf.Environment,
	ds datastore.MutableDataStore,
	chainSel uint64,
) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSel]

	// FQ 1.6.0 only knows about LINK. If the changeset picked this version, WETH would be
	// treated as an extra FQ 2.0 token and removed.
	deployFeeQuoter160(t, e, ds, chain, chainSel, []common.Address{legacyLinkToken})

	// FQ 1.6.3 is the latest 1.6.x deployment and should be selected by GetFeeQuoterAddress.
	fq163Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq163ops.Deploy, chain, contract.DeployInput[fq163ops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq163ops.ContractType, *fq163ops.Version),
		Args: fq163ops.ConstructorArgs{
			StaticConfig: fq163ops.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1e18),
				LinkToken:                    legacyLinkToken,
				TokenPriceStalenessThreshold: 3600,
			},
			PriceUpdaters: []common.Address{chain.DeployerKey.From},
			FeeTokens:     []common.Address{legacyLinkToken, legacyWethToken},
			PremiumMultiplierWeiPerEthArgs: []fq163ops.PremiumMultiplierWeiPerEthArgs{
				{Token: legacyLinkToken, PremiumMultiplierWeiPerEth: 9e17},
				{Token: legacyWethToken, PremiumMultiplierWeiPerEth: 1e18},
			},
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(fq163Out.Output))

	return deployFeeQuoter20WithExtraTokens(t, e, ds, chain, chainSel)
}

func deployFeeQuoter160(
	t *testing.T,
	e *cldf.Environment,
	ds datastore.MutableDataStore,
	chain cldf_evm.Chain,
	chainSel uint64,
	feeTokens []common.Address,
) {
	t.Helper()

	premiumArgs := make([]fq16ops.FeeTokenArgs, 0, len(feeTokens))
	for _, token := range feeTokens {
		multiplier := uint64(1e18)
		if token == legacyLinkToken {
			multiplier = 9e17
		}
		premiumArgs = append(premiumArgs, fq16ops.FeeTokenArgs{
			Token:                      token,
			PremiumMultiplierWeiPerEth: multiplier,
		})
	}

	fq16Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.Deploy, chain, contract.DeployInput[fq16ops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq16ops.ContractType, *fq16ops.Version),
		Args: fq16ops.ConstructorArgs{
			StaticConfig: fq16ops.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1e18),
				LinkToken:                    legacyLinkToken,
				TokenPriceStalenessThreshold: 3600,
			},
			PriceUpdaters:                  []common.Address{chain.DeployerKey.From},
			FeeTokens:                      feeTokens,
			PremiumMultiplierWeiPerEthArgs: premiumArgs,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(fq16Out.Output))
}

func deployRemoveFeeTokensFixtureWithPriceRegistry(
	t *testing.T,
	e *cldf.Environment,
	ds datastore.MutableDataStore,
	chainSel uint64,
) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSel]

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
		ChainSelector: chainSel,
		Type:          datastore.ContractType(priceregistryops.ContractType),
		Version:       priceregistryops.Version,
		Address:       prAddr.Hex(),
	}))

	return deployFeeQuoter20WithExtraTokens(t, e, ds, chain, chainSel)
}

func deployFeeQuoter20WithExtraTokens(
	t *testing.T,
	e *cldf.Environment,
	ds datastore.MutableDataStore,
	chain cldf_evm.Chain,
	chainSel uint64,
) common.Address {
	t.Helper()

	fq20Out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq20ops.Deploy, chain, contract.DeployInput[fq20ops.ConstructorArgs]{
		ChainSelector:  chainSel,
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
		ChainSelector: chainSel,
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
