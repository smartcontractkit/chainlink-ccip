package adapters

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
)

const testRMNProxyAddr = "0x1111111111111111111111111111111111111111"

func proxyRef(chainSelector uint64, addr string, ver *semver.Version) datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(rmn_proxy.ContractType),
		Version:       ver,
		Address:       addr,
	}
}

func TestResolveRMNProxyAddress_Found(t *testing.T) {
	sel := uint64(5009297550715157269)
	got, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{proxyRef(sel, testRMNProxyAddr, rmn_proxy.Version)},
		sel,
	)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(testRMNProxyAddr), got)
}

func TestResolveRMNProxyAddress_PicksMatchingChain(t *testing.T) {
	sel := uint64(5009297550715157269)
	otherSel := uint64(4949039107694359620)
	otherAddr := "0x2222222222222222222222222222222222222222"

	got, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{
			proxyRef(otherSel, otherAddr, rmn_proxy.Version),
			proxyRef(sel, testRMNProxyAddr, rmn_proxy.Version),
		},
		sel,
	)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(testRMNProxyAddr), got)
}

func TestResolveRMNProxyAddress_NotFoundEmpty(t *testing.T) {
	_, err := resolveRMNProxyAddress(nil, 1)
	require.ErrorContains(t, err, "RMNProxy")
	require.ErrorContains(t, err, "not found")
}

func TestResolveRMNProxyAddress_NotFoundWrongChain(t *testing.T) {
	_, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{proxyRef(42, testRMNProxyAddr, rmn_proxy.Version)},
		1,
	)
	require.ErrorContains(t, err, "not found")
}

func TestResolveRMNProxyAddress_NotFoundWrongType(t *testing.T) {
	sel := uint64(1)
	_, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{{
			ChainSelector: sel,
			Type:          datastore.ContractType("SomeOtherContract"),
			Version:       rmn_proxy.Version,
			Address:       testRMNProxyAddr,
		}},
		sel,
	)
	require.ErrorContains(t, err, "not found")
}

func TestResolveRMNProxyAddress_NotFoundWrongVersion(t *testing.T) {
	sel := uint64(1)
	wrongVer := semver.MustParse("9.9.9")
	_, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{proxyRef(sel, testRMNProxyAddr, wrongVer)},
		sel,
	)
	require.ErrorContains(t, err, "not found")
}

func TestResolveRMNProxyAddress_NotFoundNilVersion(t *testing.T) {
	sel := uint64(1)
	_, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{proxyRef(sel, testRMNProxyAddr, nil)},
		sel,
	)
	require.ErrorContains(t, err, "not found")
}

func TestResolveRMNProxyAddress_MalformedAddress(t *testing.T) {
	sel := uint64(1)
	_, err := resolveRMNProxyAddress(
		[]datastore.AddressRef{proxyRef(sel, "not-a-hex-address", rmn_proxy.Version)},
		sel,
	)
	require.ErrorContains(t, err, "not a valid hex address")
}
