package hooks

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
)

func TestEVMLaneSanityProvider_ApplySenderPrivateKey_EmptyKey(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	env := cldf.Environment{BlockChains: blockChainsWithSelectors(chain_selectors.ETHEREUM_MAINNET.Selector)}

	err := p.ApplySenderPrivateKey(t.Context(), logger.Test(t), &env, "")
	require.NoError(t, err)
}

func TestEVMLaneSanityProvider_ApplySenderPrivateKey_InvalidKey(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	env := cldf.Environment{BlockChains: blockChainsWithSelectors(chain_selectors.ETHEREUM_MAINNET.Selector)}

	err := p.ApplySenderPrivateKey(t.Context(), logger.Test(t), &env, "not-a-key")
	require.Error(t, err)
	require.ErrorContains(t, err, "parse sender private key")
	require.Contains(t, env.BlockChains.EVMChains(), chain_selectors.ETHEREUM_MAINNET.Selector)
}

func TestEVMLaneSanityProvider_ApplySenderPrivateKey_SetsDeployerKeyOnAllChains(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	ethSel := chain_selectors.ETHEREUM_MAINNET.Selector
	arbSel := chain_selectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector
	env := cldf.Environment{BlockChains: blockChainsWithSelectors(ethSel, arbSel)}

	const anvilKey0 = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	wantAddr := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	err := p.ApplySenderPrivateKey(t.Context(), logger.Test(t), &env, anvilKey0)
	require.NoError(t, err)

	for _, sel := range []uint64{ethSel, arbSel} {
		chain := env.BlockChains.EVMChains()[sel]
		require.NotNil(t, chain.DeployerKey)
		require.Equal(t, wantAddr, chain.DeployerKey.From)
	}
}

func TestEVMLaneSanityProvider_ApplySenderPrivateKey_NoEVMChains(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	env := cldf.Environment{}

	err := p.ApplySenderPrivateKey(t.Context(), logger.Test(t), &env, "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	require.Error(t, err)
	require.ErrorContains(t, err, "no EVM chains")
}

func TestEVMLaneSanityProvider_EncodeReceiverAddress(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	addr := "0x1234567890123456789012345678901234567890"

	got, err := p.EncodeReceiverAddress(cldf.Environment{}, 0, addr)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(common.HexToAddress(addr).Bytes(), 32), got)
}

func TestEVMLaneSanityProvider_EncodeReceiverAddress_Invalid(t *testing.T) {
	p := &EVMLaneSanityProvider{}

	_, err := p.EncodeReceiverAddress(cldf.Environment{}, 0, "not-an-address")
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid EVM receiver address")
}

func TestEVMLaneSanityProvider_MockReceiverAddress_NotDeployed(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	env := cldf.Environment{DataStore: datastore.NewMemoryDataStore().Seal()}

	got, err := p.MockReceiverAddress(env, chain_selectors.ETHEREUM_MAINNET.Selector)
	require.NoError(t, err)
	require.Nil(t, got)
}

func TestEVMLaneSanityProvider_MockReceiverAddress_Deployed(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	sel := chain_selectors.ETHEREUM_MAINNET.Selector
	addr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: sel,
		Type:          datastore.ContractType(mock_receiver.ContractType),
		Version:       mock_receiver.Version,
		Address:       addr,
	}))

	env := cldf.Environment{
		DataStore: ds.Seal(),
		Logger:    logger.Test(t),
	}

	got, err := p.MockReceiverAddress(env, sel)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(common.HexToAddress(addr).Bytes(), 32), got)
}

func TestEVMLaneSanityProvider_FundAndApproveTransferToken_InvalidAddress(t *testing.T) {
	p := &EVMLaneSanityProvider{}
	env := cldf.Environment{BlockChains: blockChainsWithSelectors(chain_selectors.ETHEREUM_MAINNET.Selector)}

	_, err := p.FundAndApproveTransferToken(
		t.Context(), env, chain_selectors.ETHEREUM_MAINNET.Selector, "not-an-address",
	)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid EVM token address")
}

func TestEVMLaneSanityProvider_GetMessageFee_NonEVMMessage(t *testing.T) {
	p := &EVMLaneSanityProvider{}

	fee, err := p.GetMessageFee(t.Context(), cldf.Environment{}, 1, 2, "not-evm-msg")
	require.NoError(t, err)
	require.Empty(t, fee)
}

func TestEVMLaneSanityProvider_GetMessageFee_ChainNotInEnv(t *testing.T) {
	p := &EVMLaneSanityProvider{}

	_, err := p.GetMessageFee(
		t.Context(),
		cldf.Environment{},
		chain_selectors.ETHEREUM_MAINNET.Selector,
		2,
		router.ClientEVM2AnyMessage{},
	)
	require.Error(t, err)
	require.ErrorContains(t, err, "not in environment")
}

func TestFormatFeeAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		raw      string
		decimals uint8
		want     string
	}{
		{name: "zero", raw: "0", decimals: 18, want: "0"},
		{name: "nil", raw: "", decimals: 18, want: "0"},
		{name: "native wei", raw: "1500000000000000", decimals: 18, want: "0.0015"},
		{name: "link juels", raw: "250000000000000000", decimals: 18, want: "0.25"},
		{name: "usdc", raw: "1234567", decimals: 6, want: "1.234567"},
		{name: "whole token", raw: "1000000", decimals: 6, want: "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var raw *big.Int
			if tt.raw != "" {
				raw = new(big.Int)
				raw.SetString(tt.raw, 10)
			}
			require.Equal(t, tt.want, formatFeeAmount(raw, tt.decimals))
		})
	}
}
