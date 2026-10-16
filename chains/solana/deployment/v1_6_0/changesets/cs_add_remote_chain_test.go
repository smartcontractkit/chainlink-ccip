package changesets

import (
	"bytes"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/xssnick/tonutils-go/address"

	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

func TestValidateRemoteOnRamp(t *testing.T) {
	t.Parallel()
	evmSel := chainsel.ETHEREUM_MAINNET.Selector
	solSel := chainsel.SOLANA_MAINNET.Selector
	aptosSel := chainsel.APTOS_MAINNET.Selector
	suiSel := chainsel.SUI_MAINNET.Selector
	tonSel := chainsel.TON_MAINNET.Selector

	e := cldf.Environment{DataStore: sealedDataStore(t,
		testRef(evmSel, "0x1", string(shared.OnRamp), "1.6.0", ""),
		testRef(solSel, "router", string(shared.Router), "1.6.0", ""),
		testRef(aptosSel, "0x2", string(shared.AptosCCIPType), "1.6.0", ""),
		testRef(suiSel, "0x3", "SuiCCIP", "1.6.0", ""),
		testRef(tonSel, "EQ", string(shared.Router), "1.6.0", ""),
	)}
	for _, sel := range []uint64{evmSel, solSel, aptosSel, suiSel, tonSel} {
		require.NoError(t, validateRemoteOnRamp(e, sel), "chain %d", sel)
	}

	// the EVM onramp is only loaded at 1.6.0
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(evmSel, "0x1", string(shared.OnRamp), "1.5.0", ""))}
	require.ErrorContains(t, validateRemoteOnRamp(e, evmSel), "onramp contract does not exist on evm chain")

	// a superseded ref doesn't count
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(solSel, "router", string(shared.Router), "1.6.0", "", shared.SupersededLabel))}
	require.ErrorContains(t, validateRemoteOnRamp(e, solSel), "router contract does not exist on solana chain")

	// another chain's onramp doesn't count
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(evmSel, "0x1", string(shared.OnRamp), "1.6.0", ""))}
	require.Error(t, validateRemoteOnRamp(e, chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector))

	require.Error(t, validateRemoteOnRamp(cldf.Environment{}, evmSel))
	require.Error(t, validateRemoteOnRamp(e, 1))
}

func TestRemoteOnRampAddressBytes(t *testing.T) {
	t.Parallel()
	evmSel := chainsel.ETHEREUM_MAINNET.Selector
	solSel := chainsel.SOLANA_MAINNET.Selector
	aptosSel := chainsel.APTOS_MAINNET.Selector
	suiSel := chainsel.SUI_MAINNET.Selector
	tonSel := chainsel.TON_MAINNET.Selector

	router := solana.NewWallet().PublicKey()
	tonAddr := address.NewAddress(0, 255, bytes.Repeat([]byte{7}, 32))
	e := cldf.Environment{DataStore: sealedDataStore(t,
		testRef(evmSel, "0x00000000000000000000000000000000000000aa", string(shared.OnRamp), "1.6.0", ""),
		testRef(evmSel, "0x00000000000000000000000000000000000000bb", string(shared.OnRamp), "1.5.0", ""),
		testRef(solSel, router.String(), string(shared.Router), "1.6.0", ""),
		testRef(aptosSel, "0xabc", string(shared.AptosCCIPType), "1.6.0", ""),
		testRef(suiSel, "0xdef", "SuiOnRamp", "1.6.0", ""),
		testRef(tonSel, tonAddr.String(), string(shared.OnRamp), "1.6.0", ""),
	)}

	got, err := remoteOnRampAddressBytes(e, evmSel)
	require.NoError(t, err)
	require.Equal(t, append(make([]byte, 19), 0xaa), got)

	got, err = remoteOnRampAddressBytes(e, solSel)
	require.NoError(t, err)
	require.Equal(t, router.Bytes(), got)

	got, err = remoteOnRampAddressBytes(e, aptosSel)
	require.NoError(t, err)
	require.Equal(t, append(make([]byte, 30), 0x0a, 0xbc), got)

	got, err = remoteOnRampAddressBytes(e, suiSel)
	require.NoError(t, err)
	require.Equal(t, append(make([]byte, 30), 0x0d, 0xef), got)

	// workchain -1 as a big-endian int32, then the account data
	got, err = remoteOnRampAddressBytes(e, tonSel)
	require.NoError(t, err)
	require.Equal(t, append([]byte{0xff, 0xff, 0xff, 0xff}, bytes.Repeat([]byte{7}, 32)...), got)

	// the latest active version wins
	newer := solana.NewWallet().PublicKey()
	e = cldf.Environment{DataStore: sealedDataStore(t,
		testRef(solSel, router.String(), string(shared.Router), "1.6.0", ""),
		testRef(solSel, newer.String(), string(shared.Router), "1.6.1", ""),
		testRef(solSel, solana.NewWallet().PublicKey().String(), string(shared.Router), "1.6.2", "", shared.SupersededLabel),
	)}
	got, err = remoteOnRampAddressBytes(e, solSel)
	require.NoError(t, err)
	require.Equal(t, newer.Bytes(), got)

	// two active refs at the same version are ambiguous
	e = cldf.Environment{DataStore: sealedDataStore(t,
		testRef(solSel, router.String(), string(shared.Router), "1.6.0", ""),
		testRef(solSel, newer.String(), string(shared.Router), "1.6.0", "other"),
	)}
	_, err = remoteOnRampAddressBytes(e, solSel)
	require.ErrorContains(t, err, "ambiguous")

	// the EVM onramp is only read at 1.6.0
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(evmSel, "0x00000000000000000000000000000000000000bb", string(shared.OnRamp), "1.5.0", ""))}
	_, err = remoteOnRampAddressBytes(e, evmSel)
	require.ErrorContains(t, err, "no onramp found")

	_, err = remoteOnRampAddressBytes(cldf.Environment{}, evmSel)
	require.Error(t, err)
}
