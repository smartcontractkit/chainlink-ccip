package common

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"
)

func DeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) (datastore.DataStore, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring contracts for selector")
	l.Info().Any("Selector", selector).Msg("Deploying for chain selectors")
	runningDS := datastore.NewMemoryDataStore()

	l.Info().Uint64("Selector", selector).Msg("Configuring per-chain contracts bundle")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		env.Logger,
		operations.NewMemoryReporter(),
	)
	env.OperationsBundle = bundle

	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	mint, _ := solana.NewRandomPrivateKey()
	out, err := deployops.DeployContracts(dReg).Apply(*env, deployops.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
			selector: {
				Version: version,
				// LINK TOKEN CONFIG
				// token private key used to deploy the LINK token. Solana: base58 encoded private key
				TokenPrivKey: mint.String(),
				// token decimals used to deploy the LINK token
				TokenDecimals: 9,
				// FEE QUOTER CONFIG
				MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
				LinkPremiumMultiplier:        9e17, // 0.9 ETH
				NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
				// OFFRAMP CONFIG
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
				GasForCallExactCheck:                    uint16(5000),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contracts: %w", err)
	}
	nodeClients, err := clclient.New(cls[0].Out.CLNodes)
	if err != nil {
		return nil, fmt.Errorf("connecting to CL nodes: %w", err)
	}
	// bootstrap is 0
	workerNodes := nodeClients[1:]
	if selector == ccipHomeSelector {
		var readers [][32]byte
		for _, node := range workerNodes {
			nodeP2PIds, err := node.MustReadP2PKeys()
			if err != nil {
				return nil, fmt.Errorf("reading worker node P2P keys: %w", err)
			}
			l.Info().Str("Node", node.Config.URL).Str("PeerID", nodeP2PIds.Data[0].Attributes.PeerID).Msg("Adding reader peer ID")
			id := MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
			readers = append(readers, id)
			l.Info().Msgf("peerID: %+v", id)
			l.Info().Msgf("peer ID from bytes: %s", id.Raw())
		}
		// safe to get EVM chain as CCIP home is only deployed on EVM
		chain, ok := env.BlockChains.EVMChains()[selector]
		if !ok {
			return nil, fmt.Errorf("evm chain not found for selector %d", selector)
		}
		ccipHomeOut, err := DeployHomeChain.Apply(*env, DeployHomeChainConfig{
			HomeChainSel: selector,
			CapReg:       common.HexToAddress(crAddr),
			RMNStaticConfig: rmn_home.RMNHomeStaticConfig{
				Nodes:          []rmn_home.RMNHomeNode{},
				OffchainConfig: []byte("static config"),
			},
			RMNDynamicConfig: rmn_home.RMNHomeDynamicConfig{
				SourceChains:   []rmn_home.RMNHomeSourceChain{},
				OffchainConfig: []byte("dynamic config"),
			},
			NodeOperators: []capabilities_registry.CapabilitiesRegistryNodeOperator{
				{
					Admin: chain.DeployerKey.From,
					Name:  "NodeOperator",
				},
			},
			NodeP2PIDsPerNodeOpAdmin: map[string][][32]byte{"NodeOperator": readers},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to deploy home chain contracts: %w", err)
		}
		out.DataStore.Merge(ccipHomeOut.DataStore.Seal())
		out.DataStore.Addresses().Add(
			datastore.AddressRef{
				ChainSelector: selector,
				Type:          datastore.ContractType(utils.CapabilitiesRegistry),
				Version:       semver.MustParse("1.0.0"),
				Address:       crAddr,
			},
		)
	}

	env.DataStore = out.DataStore.Seal()
	runningDS.Merge(env.DataStore)

	return runningDS.Seal(), nil
}

func ConnectContractsWithSelectors(ctx context.Context, e *deployment.Environment, selector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("FromSelector", selector).Any("ToSelectors", remoteSelectors).Msg("Connecting contracts with selectors")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	mcmsRegistry := changesetscore.GetRegistry()
	version := semver.MustParse("1.6.0")
	chainA := lanesapi.ChainDefinition{
		Selector:                 selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, cciputils.GetSelectorHex(selector)),
	}
	for _, destSelector := range remoteSelectors {
		chainB := lanesapi.ChainDefinition{
			Selector:                 destSelector,
			GasPrice:                 big.NewInt(1e9),
			FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, cciputils.GetSelectorHex(destSelector)),
		}
		l.Info().Uint64("ChainASelector", chainA.Selector).Uint64("ChainBSelector", chainB.Selector).Msg("Connecting chain pairs")
		_, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
			Lanes: []lanesapi.LaneConfig{
				{
					Version: version,
					ChainA:  chainA,
					ChainB:  chainB,
				},
			},
		})
		if err != nil {
			return fmt.Errorf("connecting chains %d and %d: %w", chainA.Selector, chainB.Selector, err)
		}
	}

	return nil
}

func ConfigureContractsForSelectors(ctx context.Context, e *deployment.Environment, cls []*simple_node_set.Input, nodeKeyBundles map[string]map[string]clclient.NodeKeysBundle, ccipHomeSelector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("HomeChainSelector", ccipHomeSelector).Msg("Configuring contracts for home chain selector")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	// Build the CCIPHome chain configs.
	chainConfigs := make(map[uint64]ChainConfig)
	commitOCRConfigs := make(map[uint64]CCIPOCRParams)
	execOCRConfigs := make(map[uint64]CCIPOCRParams)
	nodeClients, err := clclient.New(cls[0].Out.CLNodes)
	if err != nil {
		return fmt.Errorf("connecting to CL nodes: %w", err)
	}
	// bootstrap is 0
	workerNodes := nodeClients[1:]
	var readers [][32]byte
	for _, node := range workerNodes {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return fmt.Errorf("reading worker node P2P keys: %w", err)
		}
		id := MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
		readers = append(readers, id)
	}
	for _, chain := range remoteSelectors {
		ocrOverride := func(ocrParams CCIPOCRParams) CCIPOCRParams {
			if ocrParams.CommitOffChainConfig != nil {
				ocrParams.CommitOffChainConfig.RMNEnabled = false
			}
			return ocrParams
		}
		commitOCRConfigs[chain] = DeriveOCRParamsForCommit(SimulationTest, ccipHomeSelector, nil, ocrOverride)
		execOCRConfigs[chain] = DeriveOCRParamsForExec(SimulationTest, nil, ocrOverride)

		l.Info().Msgf("setting readers for chain %d to %v due to no topology", chain, len(readers))
		chainConfigs[chain] = ChainConfig{
			Readers: readers,
			FChain:  uint8(len(readers) / 3),
			EncodableChainConfig: chainconfig.ChainConfig{
				GasPriceDeviationPPB:      ccipocr3common.BigInt{Int: big.NewInt(1000)},
				DAGasPriceDeviationPPB:    ccipocr3common.BigInt{Int: big.NewInt(1000)},
				OptimisticConfirmations:   OptimisticConfirmations,
				ChainFeeDeviationDisabled: false,
			},
		}
	}

	_, err = UpdateChainConfig.Apply(*e, UpdateChainConfigConfig{
		HomeChainSelector: ccipHomeSelector,
		RemoteChainAdds:   chainConfigs,
	})
	if err != nil {
		return fmt.Errorf("updating chain config for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = AddDONAndSetCandidate.Apply(*e, AddDonAndSetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: SetCandidatePluginInfo{
			OCRConfigPerRemoteChainSelector: commitOCRConfigs,
			PluginType:                      ccipocr3common.PluginTypeCCIPCommit,
		},
		NonBootstraps:  workerNodes,
		NodeKeyBundles: nodeKeyBundles,
	})
	if err != nil {
		return fmt.Errorf("adding DON and setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = SetCandidate.Apply(*e, SetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: []SetCandidatePluginInfo{
			{
				OCRConfigPerRemoteChainSelector: execOCRConfigs,
				PluginType:                      ccipocr3common.PluginTypeCCIPExec,
			},
		},
		NonBootstraps:  workerNodes,
		NodeKeyBundles: nodeKeyBundles,
	})
	if err != nil {
		return fmt.Errorf("setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = PromoteCandidate.Apply(*e, PromoteCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		PluginInfo: []PromoteCandidatePluginInfo{
			{
				PluginType:           ccipocr3common.PluginTypeCCIPCommit,
				RemoteChainSelectors: remoteSelectors,
			},
			{
				PluginType:           ccipocr3common.PluginTypeCCIPExec,
				RemoteChainSelectors: remoteSelectors,
			},
		},
		NonBootstraps: workerNodes,
	})
	if err != nil {
		return fmt.Errorf("promoting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	dReg := deployops.GetRegistry()
	mcmsRegistry := changesetscore.GetRegistry()
	_, err = deployops.SetOCR3Config(dReg, mcmsRegistry).Apply(*e, deployops.SetOCR3ConfigArgs{
		HomeChainSel:    ccipHomeSelector,
		RemoteChainSels: remoteSelectors,
		ConfigType:      cciputils.ConfigTypeActive,
	})
	if err != nil {
		return fmt.Errorf("setting OCR3 config for selector %d: %w", ccipHomeSelector, err)
	}
	return nil
}
