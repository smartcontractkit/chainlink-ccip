package deployment

import (
	"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	cldfenv "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/link_token"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/weth9"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/testadapter"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	onrampbinding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func writeDomainDataStore(t *testing.T, dom domain.Domain, envName string, ds datastore.DataStore) {
	t.Helper()

	envDir := dom.EnvDir(envName)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))

	refs, err := ds.Addresses().Fetch()
	require.NoError(t, err)
	chainMeta, err := ds.ChainMetadata().Fetch()
	require.NoError(t, err)
	contractMeta, err := ds.ContractMetadata().Fetch()
	require.NoError(t, err)

	writeJSON := func(path string, v any) {
		t.Helper()
		b, marshalErr := json.Marshal(v)
		require.NoError(t, marshalErr)
		require.NoError(t, os.WriteFile(path, b, 0o600))
	}
	writeJSON(envDir.AddressRefsFilePath(), refs)
	writeJSON(envDir.ChainMetadataFilePath(), chainMeta)
	writeJSON(envDir.ContractMetadataFilePath(), contractMeta)

	envMeta, err := ds.EnvMetadata().Get()
	if err != nil {
		require.ErrorIs(t, err, datastore.ErrEnvMetadataNotSet)
		require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))
		return
	}
	writeJSON(envDir.EnvMetadataFilePath(), envMeta)
}

func TestProposalHookForCCIPSend(t *testing.T) {
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.AVALANCHE_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	mcmsRegistry := cs_core.GetRegistry()
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	chainInput := make(map[uint64]deployops.ContractDeploymentConfigPerChain)

	for _, chainSel := range chains {
		chainInput[chainSel] = deployops.ContractDeploymentConfigPerChain{
			Version: version,
			// FEE QUOTER CONFIG
			MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
			TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
			LinkPremiumMultiplier:        9e17, // 0.9 ETH
			NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
			// OFFRAMP CONFIG
			PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			GasForCallExactCheck:                    uint16(5000),
		}
	}
	out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	require.NoError(t, out.DataStore.Merge(e.DataStore))
	e.DataStore = out.DataStore.Seal()
	chain1 := lanesapi.ChainDefinition{
		Selector: chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice: big.NewInt(1e9),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector: chain_selectors.AVALANCHE_MAINNET.Selector,
		GasPrice: big.NewInt(1e9),
	}
	_, err = lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{
				Version: version,
				ChainA:  chain1,
				ChainB:  chain2,
			},
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")

	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.AVALANCHE_MAINNET.Selector
	srcChain := e.BlockChains.EVMChains()[srcSel]

	wethRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: srcSel,
		Type:          datastore.ContractType("WETH9"),
	}, srcSel, datastore_utils.FullRef)
	require.NoError(t, err)
	wethAddr := common.HexToAddress(wethRef.Address)
	wethCtr, err := weth9.NewWETH9(wethAddr, srcChain.Client)
	require.NoError(t, err)

	depositOpts := *srcChain.DeployerKey
	depositOpts.Context = t.Context()
	depositOpts.Value = big.NewInt(2e18)
	tx, err := wethCtr.Deposit(&depositOpts)
	require.NoError(t, err)
	_, err = srcChain.Confirm(tx)
	require.NoError(t, err)

	linkRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: srcSel,
		Type:          datastore.ContractType("LinkToken"),
	}, srcSel, datastore_utils.FullRef)
	require.NoError(t, err)
	linkAddr := common.HexToAddress(linkRef.Address)
	linkCtr, err := link_token.NewLinkToken(linkAddr, srcChain.Client)
	require.NoError(t, err)

	const minFunding = int64(1e18)
	linkBal, err := linkCtr.BalanceOf(nil, srcChain.DeployerKey.From)
	require.NoError(t, err)
	if linkBal.Cmp(big.NewInt(minFunding)) < 0 {
		mintOpts := *srcChain.DeployerKey
		mintOpts.Context = t.Context()
		mintTx, mintErr := linkCtr.Mint(&mintOpts, srcChain.DeployerKey.From, big.NewInt(2e18))
		if mintErr != nil {
			grantTx, grantErr := linkCtr.GrantMintRole(&mintOpts, srcChain.DeployerKey.From)
			require.NoError(t, grantErr)
			_, grantConfirmErr := srcChain.Confirm(grantTx)
			require.NoError(t, grantConfirmErr)

			mintTx, mintErr = linkCtr.Mint(&mintOpts, srcChain.DeployerKey.From, big.NewInt(2e18))
			require.NoError(t, mintErr)
		}
		_, err = srcChain.Confirm(mintTx)
		require.NoError(t, err)
	}
	fqAddr := getFeeQuoterFromRamps(t, e, srcSel)
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.FeeQuoterUpdatePricesSequence, e.BlockChains,
		sequences.FeeQuoterUpdatePricesSequenceInput{
			Address:       fqAddr,
			ChainSelector: srcSel,
			UpdatesByChain: fq163ops.PriceUpdates{
				TokenPriceUpdates: []fq163ops.TokenPriceUpdate{
					{SourceToken: linkAddr, UsdPerToken: big.NewInt(1e18)},
					{SourceToken: wethAddr, UsdPerToken: big.NewInt(2e18)},
				},
				GasPriceUpdates: []fq163ops.GasPriceUpdate{
					{DestChainSelector: destSel, UsdPerUnitGas: big.NewInt(1e12)},
				},
			},
		},
	)
	require.NoError(t, err)
	onRampAddrBytes, err := (&sequences.EVMAdapter{}).GetOnRampAddress(e.DataStore, srcSel)
	require.NoError(t, err)
	onRampAddr := common.BytesToAddress(onRampAddrBytes)
	onRamp, err := onrampbinding.NewOnRamp(onRampAddr, srcChain.Client)
	require.NoError(t, err)
	header, err := srcChain.Client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)
	startBlock := header.Number.Uint64()

	dom := domain.NewDomain(t.TempDir(), "test")
	writeDomainDataStore(t, dom, e.Name, e.DataStore)

	hook := hooks.GlobalPostProposalCCIPSendHook(dom)
	hookErr := hook.Func(t.Context(), cldf_changeset.PostProposalHookParams{
		Env: cldf_changeset.ProposalHookEnv{
			Name:   e.Name,
			Logger: e.Logger,
			BlockChains: chain.NewBlockChains(map[uint64]chain.BlockChain{
				chain_selectors.ETHEREUM_MAINNET.Selector: e.BlockChains.EVMChains()[chain_selectors.ETHEREUM_MAINNET.Selector],
			}),
			ForkContext: &cldf_changeset.EVMForkContext{
				ChainConfig: cldfenv.ChainConfig{
					HTTPRPCs: []cldfenv.RPCs{{External: "http://127.0.0.1:8545"}},
				},
			},
		},
		Reports: []cldf_changeset.MCMSTimelockExecuteReport{
			{
				Type:   cldf_changeset.MCMSTimelockExecuteReportType,
				Status: "SUCCESS",
				Input: cldf_changeset.MCMSTimelockExecuteReportInput{
					ChainSelector: srcSel,
				},
			},
		},
	})
	require.NoError(t, hookErr)
	it, err := onRamp.FilterCCIPMessageSent(&bind.FilterOpts{
		Start:   startBlock + 1,
		Context: t.Context(),
	}, []uint64{destSel}, []uint64{})
	require.NoError(t, err)
	count := 0
	for it.Next() {
		count++
	}
	require.NoError(t, it.Error())
	require.Greater(t, count, 0, "expected hook to send at least one CCIP message via default EVM test adapter")
}
