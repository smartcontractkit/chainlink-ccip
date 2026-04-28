package adapters

import (
	"context"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

// stubOps is a no-op OnRampFeeContractOps used to exercise registration
// behavior without standing up a chain. The returned address is fixed.
type stubOps struct {
	ret common.Address
}

func (s stubOps) GetFeeContractAddress(_ context.Context, _ evm.Chain, _ common.Address) (common.Address, error) {
	return s.ret, nil
}

const fakeOnRampType datastore.ContractType = "FakeOnRampForTest"

func TestRegisterOnRampOps_PanicsOnEmptyType(t *testing.T) {
	t.Parallel()
	r := newEVMFeeContractResolver()
	require.PanicsWithValue(t, "RegisterOnRampOps: empty ContractType", func() {
		r.RegisterOnRampOps("", semver.MustParse("1.0.0"), stubOps{})
	})
}

func TestRegisterOnRampOps_PanicsOnNilVersion(t *testing.T) {
	t.Parallel()
	r := newEVMFeeContractResolver()
	require.Panics(t, func() {
		r.RegisterOnRampOps(fakeOnRampType, nil, stubOps{})
	})
}

func TestRegisterOnRampOps_PanicsOnNilOps(t *testing.T) {
	t.Parallel()
	r := newEVMFeeContractResolver()
	require.Panics(t, func() {
		r.RegisterOnRampOps(fakeOnRampType, semver.MustParse("1.0.0"), nil)
	})
}

func TestRegisterOnRampOps_FirstWins(t *testing.T) {
	t.Parallel()
	r := newEVMFeeContractResolver()
	first := stubOps{ret: common.HexToAddress("0x1111111111111111111111111111111111111111")}
	second := stubOps{ret: common.HexToAddress("0x2222222222222222222222222222222222222222")}

	r.RegisterOnRampOps(fakeOnRampType, semver.MustParse("1.0.0"), first)
	r.RegisterOnRampOps(fakeOnRampType, semver.MustParse("1.0.0"), second)

	got := r.onRamps[newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.0.0"))]
	registered, ok := got.(stubOps)
	require.True(t, ok, "registered Ops must be stubOps")
	require.Equal(t, first.ret, registered.ret, "first registration must win")
}

func TestRegisterOnRampOps_PatchStrippingMergesKeys(t *testing.T) {
	t.Parallel()
	r := newEVMFeeContractResolver()
	v160 := stubOps{ret: common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")}
	v163 := stubOps{ret: common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")}

	r.RegisterOnRampOps(fakeOnRampType, semver.MustParse("1.6.0"), v160)
	// 1.6.3 strips to 1.6.0 → same key, first-wins semantics keep v160.
	r.RegisterOnRampOps(fakeOnRampType, semver.MustParse("1.6.3"), v163)

	require.Len(t, r.onRamps, 1, "patch-version variants must collapse to one map entry")

	got := r.onRamps[newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.6.0"))]
	require.Equal(t, v160.ret, got.(stubOps).ret)

	gotPatch := r.onRamps[newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.6.3"))]
	require.Equal(t, v160.ret, gotPatch.(stubOps).ret, "1.6.3 must resolve to the same Ops as 1.6.0")
}

func TestNewOnRampOpsKey_StripsPatch(t *testing.T) {
	t.Parallel()
	require.Equal(t,
		newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.6.0")),
		newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.6.3")),
	)
	require.Equal(t,
		newOnRampOpsKey(fakeOnRampType, semver.MustParse("2.0.0")),
		newOnRampOpsKey(fakeOnRampType, semver.MustParse("2.0.99")),
	)
	require.NotEqual(t,
		newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.6.0")),
		newOnRampOpsKey(fakeOnRampType, semver.MustParse("1.7.0")),
		"different minors must not collide",
	)
}

func TestEVMFeeContractResolver_Singleton(t *testing.T) {
	a := GetEVMFeeContractResolver()
	b := GetEVMFeeContractResolver()
	require.Same(t, a, b, "GetEVMFeeContractResolver must return the same instance across calls")
	require.NotNil(t, a)
}
