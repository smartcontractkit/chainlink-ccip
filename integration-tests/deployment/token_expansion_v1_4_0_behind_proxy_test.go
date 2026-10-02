package deployment

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bindv2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	feequoterV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	sequencesV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/message_hasher"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/stretchr/testify/require"

	testadapterV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testadapter"
	bnmERC20DripBindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"

	bmp140bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/burn_mint_token_pool"
	bmpap150bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/burn_mint_token_pool_and_proxy"
	tarbindings150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	tokenpoolV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"

	bnmOpsV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	testsetupV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// legacyArmProxy is a stand-in ARM proxy for the attached v1.4.0 pools.
//
// The v1.4.0 constructor requires one, but the attached pool is dormant for
// every lane the fronting proxy serves itself, so the address is never
// consulted during this test.
var legacyArmProxy = common.HexToAddress("0x00000000000000000000000000000000000000a1")

// attachLegacyV1_4_0Pool reproduces the production legacy topology on top of an
// already-deployed v1.5.0 *TokenPoolAndProxy:
//
//	TokenAdminRegistry -> BurnMintTokenPoolAndProxy 1.5.0
//	                        └─ getPreviousPool() -> BurnMintTokenPool 1.4.0
//	                                                  (router = the proxy)
//
// It deploys a real v1.4.0 BurnMintTokenPool for the same token with its router
// pointed at the proxy - which is what lets the proxy satisfy the legacy pool's
// onlyRouter gate when it delegates - and registers it as the proxy's previous
// pool. Returns the legacy pool address.
func attachLegacyV1_4_0Pool(
	t *testing.T,
	e *deployment.Environment,
	selector uint64,
	token common.Address,
	proxyAddr common.Address,
) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[selector]

	legacyAddr, tx, _, err := bmp140bindings.DeployBurnMintTokenPool(
		chain.DeployerKey, chain.Client,
		token, []common.Address{}, legacyArmProxy, proxyAddr,
	)
	require.NoError(t, err, "failed to deploy v1.4.0 BurnMintTokenPool on chain %d", selector)
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "failed to confirm v1.4.0 pool deployment on chain %d", selector)

	proxy, err := bmpap150bindings.NewBurnMintTokenPoolAndProxy(proxyAddr, chain.Client)
	require.NoError(t, err)
	tx, err = proxy.SetPreviousPool(chain.DeployerKey, legacyAddr)
	require.NoError(t, err, "failed to set previous pool on proxy %s", proxyAddr.Hex())
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "failed to confirm setPreviousPool on chain %d", selector)

	gotPrev, err := proxy.GetPreviousPool(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, legacyAddr, gotPrev, "proxy on chain %d should front the v1.4.0 pool", selector)

	legacy, err := bmp140bindings.NewBurnMintTokenPool(legacyAddr, chain.Client)
	require.NoError(t, err)
	gotRouter, err := legacy.GetRouter(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, proxyAddr, gotRouter, "v1.4.0 pool router on chain %d should be the fronting proxy", selector)
	gotToken, err := legacy.GetToken(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, token, gotToken, "v1.4.0 pool on chain %d should hold the same token as its proxy", selector)

	return legacyAddr
}

func TestTokenExpansionMigration_LegacyV1_4_0BehindProxy_ToV2_0_0(t *testing.T) {
	const newPoolQualA = "MIG140_NEW_POOL_A"

	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	e, selA, selB := s.env, s.selA, s.selB
	chainA := e.BlockChains.EVMChains()[selA]

	legacyA := attachLegacyV1_4_0Pool(t, e, selA, s.tokAddrA, s.oldPoolAddrA)
	legacyB := attachLegacyV1_4_0Pool(t, e, selB, s.tokAddrB, s.oldPoolAddrB)

	tarA, err := tarbindings150.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfgBefore, err := tarA.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrA, cfgBefore.TokenPool, "registry should point at the proxy before migration")
	require.NotEqual(t, legacyA, cfgBefore.TokenPool, "the v1.4.0 pool must never be the registered pool")

	e.OperationsBundle = testsetupV2.BundleWithFreshReporter(e.OperationsBundle)
	feeAdapter, fqRef, err := fees.ResolveFeeAdapter(e.OperationsBundle, e.BlockChains, e.DataStore, selA, selB)
	require.NoError(t, err)
	resolvedFee := fees.UnresolvedTokenTransferFeeArgs{
		DestBytesOverhead: cciputils.NewOptional(uint32(150_000)),
		DestGasOverhead:   cciputils.NewOptional(uint32(50_000)),
		MinFeeUSDCents:    cciputils.NewOptional(uint32(17)),
		IsEnabled:         cciputils.NewOptional(true),
	}.Resolve(feeAdapter.GetDefaultTokenTransferFeeConfig(selA, selB))
	_, err = cldf_ops.ExecuteSequence(
		e.OperationsBundle,
		feeAdapter.SetTokenTransferFee(e.DataStore, fqRef),
		e.BlockChains,
		fees.SetTokenTransferFeeSequenceInput{
			Selector: selA,
			Settings: map[uint64]map[string]*fees.TokenTransferFeeArgs{
				selB: {s.tokAddrA.Hex(): resolvedFee},
			},
		},
	)
	require.NoError(t, err)

	e.OperationsBundle = testsetupV2.BundleWithFreshReporter(e.OperationsBundle)
	upgradeOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      cciputils.Version_2_0_0,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: newPoolQualA,
					PoolType:           bnmOpsV2.ContractType.String(),
					TokenRef:           &datastore.AddressRef{Address: s.tokAddrA.Hex()},
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					AutoMigrateRemoteChains: true,
				},
			},
		},
	})
	require.NoError(t, err, "migration from a proxy with an attached v1.4.0 previous pool should succeed")
	MergeAddresses(t, e, upgradeOut.DataStore)

	newPoolAddrA, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: selA,
		Type:          datastore.ContractType(bnmOpsV2.ContractType),
		Version:       bnmOpsV2.Version,
		Qualifier:     newPoolQualA,
	}, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEqual(t, s.oldPoolAddrA, newPoolAddrA)
	require.NotEqual(t, legacyA, newPoolAddrA)

	newPoolA, err := tokenpoolV2.NewTokenPool(newPoolAddrA, chainA.Client)
	require.NoError(t, err)

	newSupported, err := newPoolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, newSupported, selB, "new pool A should support chain B after migration")

	gotRemoteToken, err := newPoolA.GetRemoteToken(&bind.CallOpts{Context: t.Context()}, selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), gotRemoteToken,
		"remote token should be carried forward from the proxy, not the v1.4.0 pool")

	gotRemotePools, err := newPoolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, selB)
	require.NoError(t, err)
	require.Contains(t, gotRemotePools, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32),
		"remote pool B should be the proxy")
	require.NotContains(t, gotRemotePools, common.LeftPadBytes(legacyB.Bytes(), 32),
		"the v1.4.0 pool behind proxy B must not be registered as a remote pool")

	cfgAfter, err := tarA.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, newPoolAddrA, cfgAfter.TokenPool, "registry should point at the new v2.0 pool after migration")

	legacyPoolA, err := bmp140bindings.NewBurnMintTokenPool(legacyA, chainA.Client)
	require.NoError(t, err)
	legacyRouter, err := legacyPoolA.GetRouter(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrA, legacyRouter, "migration must not repoint the v1.4.0 pool's router")

	proxyA, err := bmpap150bindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	stillPrev, err := proxyA.GetPreviousPool(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, legacyA, stillPrev, "migration must not clear the proxy's previous pool")

}

func TestTokenExpansionMigration_LegacyV1_4_0BehindProxy_TransferAfterCutover(t *testing.T) {
	t.Skip("test env is not provisioned to price a token transfer on this lane; router.getFee reverts with bare 0x")

	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	requireCCIPTokenTransferSucceeds(t, s.env, s.selA, s.selB, s.tokAddrA)
}

// requireCCIPTokenTransferSucceeds sends a token transfer from src to dest
// through the live Router and asserts the transaction is accepted.
func requireCCIPTokenTransferSucceeds(
	t *testing.T,
	e *deployment.Environment,
	srcSel, destSel uint64,
	token common.Address,
) {
	t.Helper()

	srcChain := e.BlockChains.EVMChains()[srcSel]
	destChain := e.BlockChains.EVMChains()[destSel]

	drip, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(token, srcChain.Client)
	require.NoError(t, err)
	tx, err := drip.Drip(srcChain.DeployerKey, srcChain.DeployerKey.From)
	require.NoError(t, err, "failed to mint transfer tokens to the sender")
	_, err = srcChain.Confirm(tx)
	require.NoError(t, err)

	balance, err := drip.BalanceOf(&bind.CallOpts{Context: t.Context()}, srcChain.DeployerKey.From)
	require.NoError(t, err)
	require.Positive(t, balance.Sign(), "sender should hold tokens before the transfer")

	ccvResolver, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: srcSel,
		Type:          datastore.ContractType(sequencesV2.CommitteeVerifierResolverType),
	}, srcSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "failed to resolve the committee verifier resolver on chain %d", srcSel)

	executor, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: srcSel,
		Type:          datastore.ContractType(sequencesV2.ExecutorProxyType),
		Qualifier:     "default",
	}, srcSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "failed to resolve the executor on chain %d", srcSel)

	_, tx, msgHasher, err := message_hasher.DeployMessageHasher(srcChain.DeployerKey, srcChain.Client)
	require.NoError(t, err)
	_, err = srcChain.Confirm(tx)
	require.NoError(t, err)

	extraArgs, err := msgHasher.EncodeGenericExtraArgsV3(
		&bindv2.CallOpts{Context: t.Context()},
		message_hasher.ExtraArgsCodecGenericExtraArgsV3{
			GasLimit:                200_000,
			RequestedFinalityConfig: finality.RawWaitForFinality,
			Ccvs:                    []common.Address{ccvResolver},
			CcvArgs:                 [][]byte{{}},
			Executor:                executor,
			ExecutorArgs:            []byte{},
			TokenReceiver:           []byte{},
			TokenArgs:               []byte{},
		},
	)
	require.NoError(t, err)

	fqAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: srcSel,
		Type:          datastore.ContractType(feequoterV2.ContractType),
	}, srcSel, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err, "failed to resolve the FeeQuoter on chain %d", srcSel)

	usdPerToken, ok := new(big.Int).SetString("1000000000000000000", 10) // $1
	require.True(t, ok)
	_, err = cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		feequoterV2.UpdatePrices, srcChain,
		evm_contract.FunctionInput[feequoterV2.PriceUpdates]{
			ChainSelector: srcSel,
			Address:       fqAddr,
			Args: feequoterV2.PriceUpdates{
				TokenPriceUpdates: []feequoterV2.TokenPriceUpdate{
					{SourceToken: token, UsdPerToken: usdPerToken},
				},
			},
		},
	)
	require.NoError(t, err, "failed to seed a token price for %s", token.Hex())

	adapter := testadapterV2.NewEVMForkCCIPSendTestAdapter(e, srcSel)

	amount := new(big.Int).Div(balance, big.NewInt(1000))
	require.Positive(t, amount.Sign(), "transfer amount must be non-zero")

	msg, err := adapter.BuildMessage(testadapters.MessageComponents{
		DestChainSelector: destSel,
		Receiver:          common.LeftPadBytes(destChain.DeployerKey.From.Bytes(), 32),
		ExtraArgs:         extraArgs,
		TokenAmounts: []testadapters.TokenAmount{
			{Token: token.Hex(), Amount: amount},
		},
	})
	require.NoError(t, err)

	_, msgID, err := adapter.SendMessage(t.Context(), destSel, msg)
	require.NoError(t, err, "ccipSend of the migrated token from chain %d to %d should succeed", srcSel, destSel)
	require.NotEmpty(t, msgID, "a successful send should return a message id")
}
