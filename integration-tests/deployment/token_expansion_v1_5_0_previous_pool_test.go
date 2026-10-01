package deployment

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	tpapOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	testsetupV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	v1_2_0_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/burn_mint_token_pool"
	v1_4_0_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/burn_mint_token_pool"
	v1_4_0_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// dummyPreviousPoolRouter and dummyPreviousPoolArmProxy are placeholder addresses for the legacy
// previous pools deployed below. They are never called: the tests only exercise
// applyChainUpdates/applyRampUpdates and rate-limiter reads, never lockOrBurn/releaseOrMint, so the
// router/armProxy wiring itself is irrelevant (mirrors the dummy addresses already used elsewhere
// for v1.5.0 proxy pool deploys in this package).
var (
	dummyPreviousPoolRouter   = common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	dummyPreviousPoolArmProxy = common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
)

// TestTokenExpansionMigration_V1_5_0_PreviousPool_V14 runs the real auto-migrate changeset
// end-to-end for a v1.5.0 *AndProxy pool whose getPreviousPool() points at an enabled v1.4 pool.
// It mirrors the live wxUSD scenario from the ticket, but goes one step further than
// TestEffectiveMigrationRateLimits_V14Previous (which calls EffectiveMigrationRateLimits
// directly): this exercises the TAR active-pool lookup, LegacyPoolAddress population, *AndProxy
// detection, and the real v2 sequence call site, by setting the previous pool tighter on outbound
// and looser on inbound than the proxy so each direction's "winner" differs, proving the
// per-direction min is actually applied through the real migration path.
func TestTokenExpansionMigration_V1_5_0_PreviousPool_V14(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	e, selA, selB := s.env, s.selA, s.selB
	chainA := e.BlockChains.EVMChains()[selA]
	proxyAddr := s.oldPoolAddrA

	proxyOut, proxyIn := readProxyRateLimits(t, chainA, proxyAddr, selB)
	require.True(t, proxyOut.IsEnabled, "proxy outbound should be enabled from setupLegacyConnectedBnMPair")
	require.True(t, proxyIn.IsEnabled, "proxy inbound should be enabled from setupLegacyConnectedBnMPair")

	prevOutCap := new(big.Int).Div(proxyOut.Capacity, big.NewInt(2))
	prevOutRate := new(big.Int).Div(proxyOut.Rate, big.NewInt(2))
	prevInCap := new(big.Int).Mul(proxyIn.Capacity, big.NewInt(10))
	prevInRate := new(big.Int).Mul(proxyIn.Rate, big.NewInt(10))

	previousAddr := deployV14PreviousPool(t, chainA, s.tokAddrA, selB,
		v1_4_0_token_pool.RateLimiterConfig{IsEnabled: true, Capacity: prevOutCap, Rate: prevOutRate},
		v1_4_0_token_pool.RateLimiterConfig{IsEnabled: true, Capacity: prevInCap, Rate: prevInRate},
	)
	setPreviousPool(t, chainA, proxyAddr, previousAddr)

	newPoolA := runV150PreviousPoolMigration(t, e, s, "MIG_NEW_POOL_A_V14_PREV")

	rl, err := newPoolA.GetCurrentRateLimiterState(&bind.CallOpts{Context: t.Context()}, selB, false)
	require.NoError(t, err)
	require.True(t, rl.OutboundRateLimiterState.IsEnabled)
	RequireBigIntsEqual(t, prevOutCap, rl.OutboundRateLimiterState.Capacity, "outbound capacity should come from the tighter v1.4 previous pool")
	RequireBigIntsEqual(t, prevOutRate, rl.OutboundRateLimiterState.Rate, "outbound rate should come from the tighter v1.4 previous pool")
	require.True(t, rl.InboundRateLimiterState.IsEnabled)
	RequireBigIntsEqual(t, proxyIn.Capacity, rl.InboundRateLimiterState.Capacity, "inbound capacity should come from the tighter proxy pool")
	RequireBigIntsEqual(t, proxyIn.Rate, rl.InboundRateLimiterState.Rate, "inbound rate should come from the tighter proxy pool")
}

// TestTokenExpansionMigration_V1_5_0_PreviousPool_V12 is the v1.2 counterpart: the previous pool's
// buckets are keyed by the proxy's address (onRamp/offRamp) rather than the remote chain selector,
// exercising the readV12Limits path instead of readV14Limits.
func TestTokenExpansionMigration_V1_5_0_PreviousPool_V12(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	e, selA, selB := s.env, s.selA, s.selB
	chainA := e.BlockChains.EVMChains()[selA]
	proxyAddr := s.oldPoolAddrA

	proxyOut, proxyIn := readProxyRateLimits(t, chainA, proxyAddr, selB)
	require.True(t, proxyOut.IsEnabled, "proxy outbound should be enabled from setupLegacyConnectedBnMPair")
	require.True(t, proxyIn.IsEnabled, "proxy inbound should be enabled from setupLegacyConnectedBnMPair")

	prevOutCap := new(big.Int).Div(proxyOut.Capacity, big.NewInt(2))
	prevOutRate := new(big.Int).Div(proxyOut.Rate, big.NewInt(2))
	prevInCap := new(big.Int).Mul(proxyIn.Capacity, big.NewInt(10))
	prevInRate := new(big.Int).Mul(proxyIn.Rate, big.NewInt(10))

	previousAddr := deployV12PreviousPool(t, chainA, s.tokAddrA, proxyAddr,
		v1_2_0_burn_mint_token_pool.RateLimiterConfig{IsEnabled: true, Capacity: prevOutCap, Rate: prevOutRate},
		v1_2_0_burn_mint_token_pool.RateLimiterConfig{IsEnabled: true, Capacity: prevInCap, Rate: prevInRate},
	)
	setPreviousPool(t, chainA, proxyAddr, previousAddr)

	newPoolA := runV150PreviousPoolMigration(t, e, s, "MIG_NEW_POOL_A_V12_PREV")

	rl, err := newPoolA.GetCurrentRateLimiterState(&bind.CallOpts{Context: t.Context()}, selB, false)
	require.NoError(t, err)
	require.True(t, rl.OutboundRateLimiterState.IsEnabled)
	RequireBigIntsEqual(t, prevOutCap, rl.OutboundRateLimiterState.Capacity, "outbound capacity should come from the tighter v1.2 previous pool")
	RequireBigIntsEqual(t, prevOutRate, rl.OutboundRateLimiterState.Rate, "outbound rate should come from the tighter v1.2 previous pool")
	require.True(t, rl.InboundRateLimiterState.IsEnabled)
	RequireBigIntsEqual(t, proxyIn.Capacity, rl.InboundRateLimiterState.Capacity, "inbound capacity should come from the tighter proxy pool")
	RequireBigIntsEqual(t, proxyIn.Rate, rl.InboundRateLimiterState.Rate, "inbound rate should come from the tighter proxy pool")
}

// readProxyRateLimits reads the actual on-chain outbound/inbound buckets off a v1.5.0 *AndProxy
// pool for remoteSelector, used as the ground truth for building a previous pool that is
// deliberately tighter in one direction and looser in the other.
func readProxyRateLimits(t *testing.T, chain evm.Chain, proxyAddr common.Address, remoteSelector uint64) (outbound, inbound tokensapi.RateLimiterConfig) {
	t.Helper()

	contract, err := tpapOps.NewTokenPoolAndProxyContract(proxyAddr, chain.Client)
	require.NoError(t, err)
	opts := &bind.CallOpts{Context: t.Context()}

	out, err := contract.GetCurrentOutboundRateLimiterState(opts, remoteSelector)
	require.NoError(t, err)
	in, err := contract.GetCurrentInboundRateLimiterState(opts, remoteSelector)
	require.NoError(t, err)

	return tokensapi.RateLimiterConfig{IsEnabled: out.IsEnabled, Capacity: out.Capacity, Rate: out.Rate},
		tokensapi.RateLimiterConfig{IsEnabled: in.IsEnabled, Capacity: in.Capacity, Rate: in.Rate}
}

// setPreviousPool wires proxy.setPreviousPool(previous). No generated wrapper exists for this
// setter (only the ABI constant), so the call is built from the raw ABI, mirroring the pattern the
// generated TokenPoolAndProxyContract wrapper itself uses.
func setPreviousPool(t *testing.T, chain evm.Chain, proxyAddr, previousAddr common.Address) {
	t.Helper()

	parsedABI, err := abi.JSON(strings.NewReader(tpapOps.TokenPoolAndProxyABI))
	require.NoError(t, err)
	bound := bind.NewBoundContract(proxyAddr, parsedABI, chain.Client, chain.Client, chain.Client)
	tx, err := bound.Transact(chain.DeployerKey, "setPreviousPool", previousAddr)
	require.NoError(t, err)
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)
}

// deployV14PreviousPool deploys a raw v1.4 BurnMintTokenPool over tokenAddr and configures its
// per-remote-chain outbound/inbound buckets for remoteSelector via applyChainUpdates.
func deployV14PreviousPool(
	t *testing.T, chain evm.Chain, tokenAddr common.Address, remoteSelector uint64,
	outbound, inbound v1_4_0_token_pool.RateLimiterConfig,
) common.Address {
	t.Helper()

	previousAddr, tx, _, err := v1_4_0_burn_mint_token_pool.DeployBurnMintTokenPool(
		chain.DeployerKey, chain.Client, tokenAddr, nil, dummyPreviousPoolArmProxy, dummyPreviousPoolRouter,
	)
	require.NoError(t, err)
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	transactor, err := v1_4_0_token_pool.NewTokenPoolTransactor(previousAddr, chain.Client)
	require.NoError(t, err)
	tx, err = transactor.ApplyChainUpdates(chain.DeployerKey, []v1_4_0_token_pool.TokenPoolChainUpdate{{
		RemoteChainSelector:       remoteSelector,
		Allowed:                   true,
		OutboundRateLimiterConfig: outbound,
		InboundRateLimiterConfig:  inbound,
	}})
	require.NoError(t, err)
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	return previousAddr
}

// deployV12PreviousPool deploys a raw v1.2 BurnMintTokenPool over tokenAddr and configures its
// onRamp/offRamp buckets keyed by proxyAddr via applyRampUpdates - v1.2 pools share one bucket per
// ramp address across every lane the proxy serves, rather than per remote chain selector.
func deployV12PreviousPool(
	t *testing.T, chain evm.Chain, tokenAddr, proxyAddr common.Address,
	outbound, inbound v1_2_0_burn_mint_token_pool.RateLimiterConfig,
) common.Address {
	t.Helper()

	previousAddr, tx, _, err := v1_2_0_burn_mint_token_pool.DeployBurnMintTokenPool(
		chain.DeployerKey, chain.Client, tokenAddr, nil, dummyPreviousPoolArmProxy,
	)
	require.NoError(t, err)
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	transactor, err := v1_2_0_burn_mint_token_pool.NewBurnMintTokenPoolTransactor(previousAddr, chain.Client)
	require.NoError(t, err)
	tx, err = transactor.ApplyRampUpdates(chain.DeployerKey,
		[]v1_2_0_burn_mint_token_pool.TokenPoolRampUpdate{{Ramp: proxyAddr, Allowed: true, RateLimiterConfig: outbound}},
		[]v1_2_0_burn_mint_token_pool.TokenPoolRampUpdate{{Ramp: proxyAddr, Allowed: true, RateLimiterConfig: inbound}},
	)
	require.NoError(t, err)
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	return previousAddr
}

// runV150PreviousPoolMigration drives the real auto-migrate changeset for chain A's v1.5.0 proxy
// pool and returns the resulting v2.0 pool, ready for rate-limit assertions.
func runV150PreviousPoolMigration(t *testing.T, e *deployment.Environment, s legacyBnMPair, newPoolQual string) *tokenpoolV2_0_0.TokenPool {
	t.Helper()
	selA := s.selA

	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	upgradeOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      cciputils.Version_2_0_0,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: newPoolQual,
					PoolType:           bnmOpsV2_0_0.ContractType.String(),
					TokenRef:           &datastore.AddressRef{Address: s.tokAddrA.Hex()},
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					AutoMigrateRemoteChains: true,
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, upgradeOut.DataStore)

	newPoolRef := datastore.AddressRef{
		ChainSelector: selA,
		Type:          datastore.ContractType(bnmOpsV2_0_0.ContractType),
		Version:       bnmOpsV2_0_0.Version,
		Qualifier:     newPoolQual,
	}
	newPoolAddrA, err := datastore_utils.FindAndFormatRef(e.DataStore, newPoolRef, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEqual(t, s.oldPoolAddrA, newPoolAddrA, "new pool must be a distinct contract from the old proxy pool")

	chainA := e.BlockChains.EVMChains()[selA]
	newPoolA, err := tokenpoolV2_0_0.NewTokenPool(newPoolAddrA, chainA.Client)
	require.NoError(t, err)
	return newPoolA
}
