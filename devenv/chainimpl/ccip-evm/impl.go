package ccip_evm

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

var ccipMessageSentTopic = onramp.OnRampCCIPMessageSent{}.Topic()

type CCIP16EVM struct {
	e                      *deployment.Environment
	chainDetailsBySelector map[uint64]chainsel.ChainDetails
	ethClients             map[uint64]*ethclient.Client
}

// NewCCIP16EVM creates new smart-contracts wrappers with utility functions for CCIP16EVM implementation.
func NewCCIP16EVM(ctx context.Context, e *deployment.Environment) (*CCIP16EVM, error) {
	_ = zerolog.Ctx(ctx)
	return &CCIP16EVM{
		e: e,
	}, nil
}

func (m *CCIP16EVM) SetCLDF(e *deployment.Environment) {
	m.e = e
}

func (m *CCIP16EVM) SendMessage(ctx context.Context, src, dest uint64, fields any, opts any) error {
	_ = zerolog.Ctx(ctx)
	return nil
}

func (m *CCIP16EVM) GetExpectedNextSequenceNumber(ctx context.Context, from, to uint64) (uint64, error) {
	_ = zerolog.Ctx(ctx)
	return 0, nil
}

// WaitOneSentEventBySeqNo wait and fetch strictly one CCIPMessageSent event by selector and sequence number and selector.
func (m *CCIP16EVM) WaitOneSentEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

// WaitOneExecEventBySeqNo wait and fetch strictly one ExecutionStateChanged event by sequence number and selector.
func (m *CCIP16EVM) WaitOneExecEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

func (m *CCIP16EVM) GetEOAReceiverAddress(ctx context.Context, chainSelector uint64) ([]byte, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

func (m *CCIP16EVM) GetTokenBalance(ctx context.Context, chainSelector uint64, address, tokenAddress []byte) (*big.Int, error) {
	_ = zerolog.Ctx(ctx)
	return big.NewInt(0), nil
}

func (m *CCIP16EVM) ExposeMetrics(
	ctx context.Context,
	source, dest uint64,
	chainIDs []string,
	wsURLs []string,
) ([]string, *prometheus.Registry, error) {
	msgSentTotal.Reset()
	msgExecTotal.Reset()
	srcDstLatency.Reset()

	reg := prometheus.NewRegistry()
	reg.MustRegister(msgSentTotal, msgExecTotal, srcDstLatency)

	lp := NewLokiPusher()
	err := ProcessLaneEvents(ctx, m, lp, &LaneStreamConfig{
		FromSelector:      source,
		ToSelector:        dest,
		AggregatorAddress: "localhost:50051",
		AggregatorSince:   0,
	})
	if err != nil {
		return nil, nil, err
	}
	err = ProcessLaneEvents(ctx, m, lp, &LaneStreamConfig{
		FromSelector:      dest,
		ToSelector:        source,
		AggregatorAddress: "localhost:50051",
		AggregatorSince:   0,
	})
	if err != nil {
		return nil, nil, err
	}
	return []string{}, reg, nil
}

func (m *CCIP16EVM) DeployLocalNetwork(ctx context.Context, bc *blockchain.Input) (*blockchain.Output, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Deploying EVM networks")
	out, err := blockchain.NewBlockchainNetwork(bc)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain network: %w", err)
	}
	return out, nil
}

func (m *CCIP16EVM) ConfigureNodes(ctx context.Context, bc *blockchain.Input) (string, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring CL nodes")
	name := fmt.Sprintf("node-evm-%s", uuid.New().String()[0:5])
	finality := 1
	return fmt.Sprintf(`
       [[EVM]]
       LogPollInterval = '1s'
       BlockBackfillDepth = 100
       ChainID = '%s'
       MinIncomingConfirmations = 1
       MinContractPayment = '0.0000001 link'
       FinalityDepth = %d

       [[EVM.Nodes]]
       Name = '%s'
       WsUrl = '%s'
       HttpUrl = '%s'`,
		bc.ChainID,
		finality,
		name,
		bc.Out.Nodes[0].InternalWSUrl,
		bc.Out.Nodes[0].InternalHTTPUrl,
	), nil
}

func (m *CCIP16EVM) DeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64) (datastore.DataStore, error) {
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

	chain, ok := env.BlockChains.EVMChains()[selector]
	if !ok {
		return nil, fmt.Errorf("evm chain not found for selector %d", selector)
	}
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	out, err := deployops.DeployContracts(dReg).Apply(*env, deployops.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
			chain.Selector: {
				Version: version,
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
		var nodeOperators []capabilities_registry.CapabilitiesRegistryNodeOperator
		var nodeP2PIDsPerNodeOpAdmin = make(map[string][][32]byte)
		for _, node := range workerNodes {
			nodeP2PIds, err := node.MustReadP2PKeys()
			if err != nil {
				return nil, fmt.Errorf("reading worker node P2P keys: %w", err)
			}
			nodeTransmitterAddress, err := node.PrimaryEthAddress()
			if err != nil {
				return nil, fmt.Errorf("reading worker node transmitter address: %w", err)
			}
			nodeP2PIDsPerNodeOpAdmin[node.Config.URL] = make([][32]byte, 0)
			for _, id := range nodeP2PIds.Data {
				nodeOperators = append(nodeOperators, capabilities_registry.CapabilitiesRegistryNodeOperator{
					Admin: common.HexToAddress(nodeTransmitterAddress),
					Name:  string(node.Config.URL),
				})
				var peerID [32]byte
				copy(peerID[:], []byte(id.Attributes.PeerID))
				nodeP2PIDsPerNodeOpAdmin[node.Config.URL] = append(
					nodeP2PIDsPerNodeOpAdmin[node.Config.URL], peerID,
				)
			}
		}
		ccipHomeOut, err := changesets.DeployHomeChain.Apply(*env, sequences.DeployHomeChainConfig{
			HomeChainSel: selector,
			RMNStaticConfig: rmn_home.RMNHomeStaticConfig{
				Nodes:          []rmn_home.RMNHomeNode{},
				OffchainConfig: []byte("static config"),
			},
			RMNDynamicConfig: rmn_home.RMNHomeDynamicConfig{
				SourceChains:   []rmn_home.RMNHomeSourceChain{},
				OffchainConfig: []byte("dynamic config"),
			},
			NodeOperators:            nodeOperators,
			NodeP2PIDsPerNodeOpAdmin: nodeP2PIDsPerNodeOpAdmin,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to deploy home chain contracts: %w", err)
		}
		out.DataStore.Merge(ccipHomeOut.DataStore.Seal())
	}
	env.DataStore = out.DataStore.Seal()
	runningDS.Merge(env.DataStore)

	return runningDS.Seal(), nil
}

func (m *CCIP16EVM) ConnectContractsWithSelectors(ctx context.Context, e *deployment.Environment, selector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("FromSelector", selector).Any("ToSelectors", remoteSelectors).Msg("Connecting contracts with selectors")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	// we're assuming all dest chains are EVM for this implementation
	evmEncoded, err := hex.DecodeString(cciputils.EVMFamilySelector)
	if err != nil {
		return fmt.Errorf("encoding EVM family selector: %w", err)
	}
	mcmsRegistry := changesetscore.GetRegistry()
	version := semver.MustParse("1.6.0")
	chainA := lanesapi.ChainDefinition{
		Selector:                 selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, evmEncoded),
	}
	for _, destSelector := range remoteSelectors {
		chainB := lanesapi.ChainDefinition{
			Selector:                 destSelector,
			GasPrice:                 big.NewInt(1e9),
			FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, evmEncoded),
		}
		l.Info().Uint64("ChainASelector", chainA.Selector).Uint64("ChainBSelector", chainB.Selector).Msg("Connecting chain pairs")
		_, err = lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
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

func (m *CCIP16EVM) ConfigureContractsForSelectors(ctx context.Context, e *deployment.Environment, cls []*simple_node_set.Input, ccipHomeSelector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("HomeChainSelector", ccipHomeSelector).Msg("Configuring contracts for home chain selector")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	// Build the CCIPHome chain configs.
	chainConfigs := make(map[uint64]changesets.ChainConfig)
	commitOCRConfigs := make(map[uint64]changesets.CCIPOCRParams)
	execOCRConfigs := make(map[uint64]changesets.CCIPOCRParams)
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
		for _, id := range nodeP2PIds.Data {
			var peerID [32]byte
			copy(peerID[:], []byte(id.Attributes.PeerID))
			readers = append(readers, peerID)
		}
	}
	for _, chain := range remoteSelectors {
		ocrOverride := func(ocrParams changesets.CCIPOCRParams) changesets.CCIPOCRParams {
			if ocrParams.CommitOffChainConfig != nil {
				ocrParams.CommitOffChainConfig.RMNEnabled = false
			}
			return ocrParams
		}
		commitOCRConfigs[chain] = changesets.DeriveOCRParamsForCommit(changesets.SimulationTest, ccipHomeSelector, nil, ocrOverride)
		execOCRConfigs[chain] = changesets.DeriveOCRParamsForExec(changesets.SimulationTest, nil, ocrOverride)

		l.Info().Msgf("setting readers for chain %d to %v due to no topology", chain, len(readers))
		chainConfigs[chain] = changesets.ChainConfig{
			Readers: readers,
			FChain: uint8(len(readers) / 3),
			EncodableChainConfig: chainconfig.ChainConfig{
				GasPriceDeviationPPB:      ccipocr3common.BigInt{Int: big.NewInt(1000)},
				DAGasPriceDeviationPPB:    ccipocr3common.BigInt{Int: big.NewInt(1000)},
				OptimisticConfirmations:   changesets.OptimisticConfirmations,
				ChainFeeDeviationDisabled: false,
			},
		}
	}

	_, err = changesets.UpdateChainConfig.Apply(*e, changesets.UpdateChainConfigConfig{
		HomeChainSelector: ccipHomeSelector,
		RemoteChainAdds:   chainConfigs,
	})
	if err != nil {
		return fmt.Errorf("updating chain config for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = changesets.AddDONAndSetCandidate.Apply(*e, changesets.AddDonAndSetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: changesets.SetCandidatePluginInfo{
			OCRConfigPerRemoteChainSelector: commitOCRConfigs,
			PluginType:                      ccipocr3common.PluginTypeCCIPCommit,
		},
		NonBootstraps:    workerNodes,
	})
	if err != nil {
		return fmt.Errorf("adding DON and setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = changesets.SetCandidate.Apply(*e, changesets.SetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: []changesets.SetCandidatePluginInfo{
			{
				OCRConfigPerRemoteChainSelector: execOCRConfigs,
				PluginType:                      ccipocr3common.PluginTypeCCIPExec,
			},
		},
		NonBootstraps:    workerNodes,
	})
	if err != nil {
		return fmt.Errorf("setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = changesets.PromoteCandidate.Apply(*e, changesets.PromoteCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		PluginInfo: []changesets.PromoteCandidatePluginInfo{
			{
				PluginType:           ccipocr3common.PluginTypeCCIPCommit,
				RemoteChainSelectors: remoteSelectors,
			},
			{
				PluginType:           ccipocr3common.PluginTypeCCIPExec,
				RemoteChainSelectors: remoteSelectors,
			},
		},
		NonBootstraps:    workerNodes,
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

func (m *CCIP16EVM) FundNodes(ctx context.Context, ns []*simple_node_set.Input, bc *blockchain.Input, linkAmount, nativeAmount *big.Int) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Funding CL nodes with ETH and LINK")
	nodeClients, err := clclient.New(ns[0].Out.CLNodes)
	if err != nil {
		return fmt.Errorf("connecting to CL nodes: %w", err)
	}
	ethKeyAddressesSrc := make([]string, 0)
	for i, nc := range nodeClients {
		addrSrc, err := nc.ReadPrimaryETHKey(bc.ChainID)
		if err != nil {
			return fmt.Errorf("getting primary ETH key from OCR node %d (src chain): %w", i, err)
		}
		ethKeyAddressesSrc = append(ethKeyAddressesSrc, addrSrc.Attributes.Address)
		l.Info().
			Int("Idx", i).
			Str("ETHKeySrc", addrSrc.Attributes.Address).
			Msg("Node info")
	}
	clientSrc, _, _, err := ETHClient(ctx, bc.Out.Nodes[0].ExternalWSUrl, &GasSettings{
		FeeCapMultiplier: 2,
		TipCapMultiplier: 2,
	})
	if err != nil {
		return fmt.Errorf("could not create basic eth client: %w", err)
	}
	for _, addr := range ethKeyAddressesSrc {
		a, _ := nativeAmount.Float64()
		if err := FundNodeEIP1559(ctx, clientSrc, getNetworkPrivateKey(), addr, a); err != nil {
			return fmt.Errorf("failed to fund CL nodes on src chain: %w", err)
		}
	}
	return nil
}

// GetContractAddrForSelector get contract address by type and chain selector.
func GetContractAddrForSelector(addresses []string, selector uint64, contractType datastore.ContractType) (common.Address, error) {
	var contractAddr common.Address
	for _, addr := range addresses {
		var refs []datastore.AddressRef
		err := json.Unmarshal([]byte(addr), &refs)
		if err != nil {
			return common.Address{}, err
		}
		for _, ref := range refs {
			if ref.ChainSelector == selector && ref.Type == contractType {
				contractAddr = common.HexToAddress(ref.Address)
			}
		}
	}
	return contractAddr, nil
}
