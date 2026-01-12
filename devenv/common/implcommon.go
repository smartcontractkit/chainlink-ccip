package common

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/rs/zerolog"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func DeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) (datastore.DataStore, error) {
	l := zerolog.Ctx(ctx)
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
	// Usually set in the GH workflow
	// If not set, defaults to a known working commit hash
	// If set to a specific version, fetches from https://github.com/smartcontractkit/chainlink-ton/releases
	// Directory needs to exist at ../contracts/build relative to chainlink-ccip/devenv for TON
	contractVersion := os.Getenv("DEPLOY_CONTRACT_VERSION")
	if contractVersion == "" {
		contractVersion = "a60d19e33dc8" // Jan 5, 2026 commit hash
	}
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
				// TON SPECIFIC CONFIG
				ContractVersion: contractVersion,
				// PING PONG DEMO - deploy for cross-chain testing
				DeployPingPongDapp: true,
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
			id := MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
			readers = append(readers, id)
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

// getTokenPricesForChain returns the token prices for fee tokens on a chain.
// Uses the optional TokenPriceProvider interface - only chains that implement it (like EVM) provide prices.
// Resolves contract types to addresses using the datastore.
func getTokenPricesForChain(ds datastore.DataStore, selector uint64, version *semver.Version) map[string]*big.Int {
	family, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		return make(map[string]*big.Int)
	}

	adapter, exists := lanesapi.GetLaneAdapterRegistry().GetLaneAdapter(family, version)
	if !exists {
		return make(map[string]*big.Int)
	}

	// Check if adapter implements TokenPriceProvider (optional interface)
	priceProvider, ok := adapter.(lanesapi.TokenPriceProvider)
	if !ok {
		return make(map[string]*big.Int)
	}

	// Get prices keyed by contract type
	typePrices := priceProvider.GetDefaultTokenPrices()

	// Resolve contract types to addresses
	addressPrices := make(map[string]*big.Int)
	for contractType, price := range typePrices {
		refs := ds.Addresses().Filter(
			datastore.AddressRefByType(contractType),
			datastore.AddressRefByChainSelector(selector),
		)
		for _, ref := range refs {
			addressPrices[ref.Address] = price
		}
	}

	return addressPrices
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

	// Get token prices for the source chain - uses the LaneAdapter for chain-specific logic
	chainATokenPrices := getTokenPricesForChain(e.DataStore, selector, version)

	chainA := lanesapi.ChainDefinition{
		Selector:                 selector,
		GasPrice:                 lanesapi.DefaultGasPrice(selector),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, selector),
		TokenPrices:              chainATokenPrices,
	}
	for _, destSelector := range remoteSelectors {
		// Get token prices for the destination chain - uses the LaneAdapter for chain-specific logic
		chainBTokenPrices := getTokenPricesForChain(e.DataStore, destSelector, version)

		chainB := lanesapi.ChainDefinition{
			Selector:                 destSelector,
			GasPrice:                 lanesapi.DefaultGasPrice(destSelector),
			FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, destSelector),
			TokenPrices:              chainBTokenPrices,
		}
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

	// Configure PingPong contracts (silently skips chains that don't support it)
	err := lanesapi.ConfigurePingPongForLanes(*e, lanesapi.GetPingPongAdapterRegistry(), version, selector, remoteSelectors)
	if err != nil {
		return fmt.Errorf("configuring PingPong for selector %d: %w", selector, err)
	}

	return nil
}

func AddNodesToContracts(ctx context.Context, e *deployment.Environment, cls []*simple_node_set.Input, nodeKeyBundles map[string]map[string]clclient.NodeKeysBundle, ccipHomeSelector uint64, remoteSelectors []uint64) error {
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
