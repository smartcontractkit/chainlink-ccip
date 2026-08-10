package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	routerops1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	adaptersv2_1_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/adapters"
	rmnops2_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	soladapterv1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	solofframpops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// deployRMN2_1WithProxy deploys RMN 2.1.0 on the given chain together with the RMNProxy that
// fronts it, and registers both in the datastore. The proxy is mandatory for the 2.1.0 fastcurse
// adapter: it resolves the active RMN via utils.ActiveRMNAddress -> RMNProxy.getARM(). It returns
// the RMN address and the RMNProxy address.
func deployRMN2_1WithProxy(
	t *testing.T,
	bundle cldf_ops.Bundle,
	evmChain cldf_evm.Chain,
	ds *datastore.MemoryDataStore,
) (common.Address, common.Address) {
	t.Helper()
	deployRMNOp, err := cldf_ops.ExecuteOperation(bundle, rmnops2_1.Deploy, evmChain, contract.DeployInput[rmnops2_1.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnops2_1.ContractType, *rmnops2_1.Version),
		ChainSelector:  evmChain.Selector,
		Args:           rmnops2_1.ConstructorArgs{CurseAdmins: []common.Address{}},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnops2_1.ContractType),
		Version:       rmnops2_1.Version,
		ChainSelector: evmChain.Selector,
		Address:       deployRMNOp.Output.Address,
	}))
	rmnAddress := common.HexToAddress(deployRMNOp.Output.Address)

	deployRMNProxyOp, err := cldf_ops.ExecuteOperation(bundle, rmnproxyops.Deploy, evmChain, contract.DeployInput[rmnproxyops.ConstructorArgs]{
		ChainSelector:  evmChain.Selector,
		TypeAndVersion: deployment.NewTypeAndVersion(rmnproxyops.ContractType, *semver.MustParse("1.0.0")),
		Args:           rmnproxyops.ConstructorArgs{RMN: rmnAddress},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		ChainSelector: evmChain.Selector,
		Address:       deployRMNProxyOp.Output.Address,
	}))

	return rmnAddress, common.HexToAddress(deployRMNProxyOp.Output.Address)
}

// rmnOwnershipTransferInput builds the transfer-ownership input that hands the RMN 2.1.0
// deployment on each of the given chains over to that chain's timelock.
func rmnOwnershipTransferInput(selectors []uint64, timelockAddrs map[uint64]string) []deploy.TransferOwnershipPerChainInput {
	inputs := make([]deploy.TransferOwnershipPerChainInput, 0, len(selectors))
	for _, sel := range selectors {
		inputs = append(inputs, deploy.TransferOwnershipPerChainInput{
			ChainSelector: sel,
			ContractRef: []datastore.AddressRef{
				{
					Type:    datastore.ContractType(rmnops2_1.ContractType),
					Version: rmnops2_1.Version,
				},
			},
			ProposedOwner: timelockAddrs[sel],
		})
	}
	return inputs
}

func TestFastCurseSolanaAndEVM(t *testing.T) {
	chain1 := chainsel.TEST_90000001.Selector
	chain2 := chainsel.TEST_90000002.Selector
	programsPath, dstr, err := PreloadSolanaEnvironment(t, chainsel.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, dstr, "Datastore should be created")
	solanaChains := []uint64{
		chainsel.SOLANA_MAINNET.Selector,
	}
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2}),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)
	require.NotNil(t, env, "Environment should be created")
	env.DataStore = dstr.Seal() // Add preloaded contracts to env datastore
	mint, _ := solana.NewRandomPrivateKey()

	dReg := deploy.GetRegistry()
	version := semver.MustParse("1.6.0")
	SeedUltraFastCurseMCMS(t, env)
	_, err = deploy.DeployContracts(dReg).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			chainsel.SOLANA_MAINNET.Selector: {
				Version: version,
				// LINK TOKEN CONFIG
				// token private key used to deploy the LINK token. Solana: base58 encoded private key
				TokenPrivKey: mint.String(),
				// token decimals used to deploy the LINK token
				TokenDecimals: 9,
				// FEE QUOTER CONFIG
				MaxFeeJuelsPerMsg: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				// OFFRAMP CONFIG
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			},
		},
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")

	// Ensure the Solana OffRamp is configured with chain2 as an enabled source chain.
	// The Solana fastcurse adapter's connectivity check for lane actions reads the OffRamp
	// source-chain state (PDA) and will treat the chains as disconnected unless this is set.
	_, err = cldf_ops.ExecuteOperation(
		env.OperationsBundle,
		solofframpops.ConnectChains,
		env.BlockChains.SolanaChains()[chainsel.SOLANA_MAINNET.Selector],
		solofframpops.ConnectChainsParams{
			OffRamp:                   solana.MustPublicKeyFromBase58(solanaProgramIDs["ccip_offramp"]),
			RemoteChainSelector:       chain2,
			SourceOnRamp:              common.HexToAddress("0x0000000000000000000000000000000000000001").Bytes(),
			EnabledAsSource:           true,
			IsRMNVerificationDisabled: true,
		},
	)
	require.NoError(t, err, "Failed to connect Solana OffRamp to EVM chain2")

	DeployMCMS(t, env, chainsel.SOLANA_MAINNET.Selector, []string{deploymentutils.CLLQualifier})
	SolanaTransferOwnership(t, env, chainsel.SOLANA_MAINNET.Selector)
	ds := datastore.NewMemoryDataStore()
	bundle := env.OperationsBundle
	rmnAddresses := make(map[uint64]common.Address)
	// deploy RMN 2.1 (behind its RMNProxy) and a router on both EVM chains
	for _, sel := range []uint64{chain1, chain2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		rmnAddr, rmnProxy := deployRMN2_1WithProxy(t, bundle, evmChain, ds)
		rmnAddresses[sel] = rmnAddr
		// mock wrapped native
		wNative := utils.RandomAddress()

		deployRouterOp, err := cldf_ops.ExecuteOperation(bundle, routerops1_2.Deploy, evmChain, contract.DeployInput[routerops1_2.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
			Args: routerops1_2.ConstructorArgs{
				WrappedNative: wNative,
				RMNProxy:      rmnProxy,
			},
		})
		require.NoError(t, err)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			Type:          datastore.ContractType(routerops1_2.ContractType),
			Version:       semver.MustParse("1.2.0"),
			ChainSelector: sel,
			Address:       deployRouterOp.Output.Address,
		}))
		routerAddr := deployRouterOp.Output.Address
		// add some dummy onramps to the router so that chain is supported,
		// on chain1, add chain2 as supported dest chain and vice versa
		onRamp := utils.RandomAddress()
		offRamp := utils.RandomAddress()
		var destChainSelector uint64
		if sel == chain1 {
			destChainSelector = chain2
		} else {
			destChainSelector = chain1
		}
		_, err = cldf_ops.ExecuteOperation(bundle, routerops1_2.ApplyRampUpdates, evmChain, contract.FunctionInput[routerops1_2.ApplyRampsUpdatesArgs]{
			Address:       common.HexToAddress(routerAddr),
			ChainSelector: evmChain.Selector,
			Args: routerops1_2.ApplyRampsUpdatesArgs{
				OnRampUpdates: []routerops1_2.OnRamp{
					{
						DestChainSelector: destChainSelector,
						OnRamp:            onRamp,
					},
					{
						DestChainSelector: chainsel.SOLANA_MAINNET.Selector,
						OnRamp:            common.BytesToAddress([]byte(solanaProgramIDs["ccip_router"])),
					},
				},
				OffRampAdds: []routerops1_2.OffRamp{
					{
						SourceChainSelector: destChainSelector,
						OffRamp:             offRamp,
					},
					{
						SourceChainSelector: chainsel.SOLANA_MAINNET.Selector,
						OffRamp:             common.BytesToAddress([]byte(solanaProgramIDs["ccip_offramp"])),
					},
				},
			},
		})
		require.NoError(t, err)
	}

	// deploy mcms
	cs := deploy.DeployMCMS(dReg, nil)
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS:           testhelpers.MCMSInputForQualifier(deploymentutils.CLLQualifier),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chain1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        ptr.String(deploymentutils.CLLQualifier),
			},
			chain2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        ptr.String(deploymentutils.CLLQualifier),
			},
		},
	})
	require.NoError(t, err)
	// store addresses in ds
	allAddrRefs, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range allAddrRefs {
		require.NoError(t, ds.Addresses().Add(addrRef))
		// Qualifier-scoped on purpose: the datastore also holds the UltraFastCurse RBACTimelock, and
		// an unqualified match would let it win this loop and become the proposed owner.
		if addrRef.Type == datastore.ContractType(deploymentutils.RBACTimelock) &&
			addrRef.Qualifier == deploymentutils.CLLQualifier {
			timelockAddrs[addrRef.ChainSelector] = addrRef.Address
		}
	}
	// update env datastore
	require.NoError(t, ds.Merge(env.DataStore))
	env.DataStore = ds.Seal()
	// transfer ownership of the RMN deployments to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs:    rmnOwnershipTransferInput([]uint64{chain1, chain2}, timelockAddrs),
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Transfer ownership to timelock for fast curse test",
		},
	}

	// register chain adapter
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(deploy.GetTransferOwnershipRegistry(), changesets.GetRegistry())
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN to respective MCMS")
	// now generate a curse proposal
	curseCfg := fastcurse.RMNCurseConfig{
		CurseActions: []fastcurse.CurseActionInput{
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain1,
				SubjectChainSelector: chain2,
				Version:              rmnops2_1.Version,
			},
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain2,
				SubjectChainSelector: chain1,
				Version:              rmnops2_1.Version,
			},
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain2,
				SubjectChainSelector: chainsel.SOLANA_MAINNET.Selector,
				Version:              rmnops2_1.Version,
			},
			{
				// Lane curse actions must be represented bidirectionally; reverse direction can be any version.
				IsGlobalCurse:        false,
				ChainSelector:        chainsel.SOLANA_MAINNET.Selector,
				SubjectChainSelector: chain2,
				Version:              semver.MustParse("1.6.0"),
			},
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Curse proposal for fast curse test",
		},
	}
	adv2_1_0 := adaptersv2_1_0.NewCurseAdapter()

	curseChangeset := fastcurse.CurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually cursed
	evmChain1 := env.BlockChains.EVMChains()[chain1]
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	rmnC, err := rmnops2_1.NewRMNContract(rmnAddresses[chain1], evmChain1.Client)
	require.NoError(t, err)
	isCursed, err := rmnC.IsCursed(nil, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain2 should be cursed on rmn in chain1")

	rmnC2, err := rmnops2_1.NewRMNContract(rmnAddresses[chain2], evmChain2.Client)
	require.NoError(t, err)
	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain1 should be cursed on rmn in chain2")
	t.Logf("Subjects successfully cursed %x on chain1 %d and %x on chain2 %d", adv2_1_0.SelectorToSubject(chain2), chain1, adv2_1_0.SelectorToSubject(chain1), chain2)

	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chainsel.SOLANA_MAINNET.Selector))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on solana chain should be cursed on rmn in chain2")
	t.Logf("Subject successfully cursed %x on solana chain %d in rmnremote on chain2 %d", soladapterv1_6_0.NewCurseAdapter().SelectorToSubject(chainsel.SOLANA_MAINNET.Selector), chainsel.SOLANA_MAINNET.Selector, chain2)

	/*
		 Enable Solana checks later
		 // find rmn_remote address for solana chain
		solanaRMNRemoteAddrRef, err := datastore_utils.FindAndFormatRef(
			env.DataStore,
			datastore.AddressRef{
				ChainSelector: chainsel.SOLANA_MAINNET.Selector,
				Type:          datastore.ContractType(solrmnremoteops.ContractType),
				Version:       solrmnremoteops.Version,
			},
			chain2,
			solutils.ToAddress,
		)
		require.NoError(t, err)

		isCursed, err = solrmnremoteops.IsSubjectCursed(
			env.BlockChains.SolanaChains()[chainsel.SOLANA_MAINNET.Selector],
			solanaRMNRemoteAddrRef,
			solrmn_remote.CurseSubject{
				Value: soladapterv1_6_0.NewCurseAdapter().SelectorToSubject(chain2),
			},
		)
		require.NoError(t, err)
		require.True(t, isCursed, "subject on chain1 should be cursed on solana rmnremote")
		t.Logf("Subject successfully cursed %x on chain1 %d in solana rmnremote on solana chain %d", adv2_1_0.SelectorToSubject(chain2), chain1, chainsel.SOLANA_MAINNET.Selector)
	*/
	// Now uncurse the subjects
	// reset the operation bundle to clear any cached values
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	uncurseChangeset := fastcurse.UncurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = uncurseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually uncursed
	isCursed, err = rmnC.IsCursed(nil, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be uncursed on rmn in chain1")

	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be uncursed on rmn in chain2")
	t.Logf("Subjects successfully uncursed %x on chain1 %d and %x on chain2 %d", adv2_1_0.SelectorToSubject(chain2), chain1, adv2_1_0.SelectorToSubject(chain1), chain2)
	/*
			 Enable Solana checks later
		isCursed, err = solrmnremoteops.IsSubjectCursed(
			env.BlockChains.SolanaChains()[chainsel.SOLANA_MAINNET.Selector],
			solanaRMNRemoteAddrRef,
			solrmn_remote.CurseSubject{
				Value: soladapterv1_6_0.NewCurseAdapter().SelectorToSubject(chain2),
			},
		)
		require.NoError(t, err)
		require.False(t, isCursed, "subject on chain1 should be cursed on solana rmnremote")
		t.Logf("Subject successfully uncursed %x on chain1 %d in solana rmnremote on solana chain %d", adv2_1_0.SelectorToSubject(chain2), chain1, chainsel.SOLANA_MAINNET.Selector)
	*/
}
