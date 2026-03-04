package common

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
	"github.com/rs/zerolog"

	chainsel "github.com/smartcontractkit/chain-selectors"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/devenv/blockchainutils"
	mcmstypes "github.com/smartcontractkit/mcms/types"
)

var (
	// TestXXXMCMSSigner is a throwaway private key used for signing MCMS proposals.
	// in tests.
	TestXXXMCMSSigner *ecdsa.PrivateKey
)

func init() {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	TestXXXMCMSSigner = key
}

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
		contractVersion = "4f7b7be09c30" // https://github.com/smartcontractkit/chainlink-ton/releases/tag/ton-contracts-build-4f7b7be09c30
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
	// so we can get the LINK address etc.
	tmp := datastore.NewMemoryDataStore()
	tmp.Merge(env.DataStore)
	tmp.Merge(out.DataStore.Seal())
	runningDS.Merge(out.DataStore.Seal())

	// For EVM only, set the timelock admin
	var timelockAdmin common.Address
	chain1, ok := env.BlockChains.EVMChains()[selector]
	if ok {
		timelockAdmin = chain1.DeployerKey.From
	}
	qualifier := "CLLCCIP"
	cs := deployops.DeployMCMS(dReg, nil)
	fcs := deployops.FinalizeDeployMCMS(dReg, nil)
	output, err := cs.Apply(*env, deployops.MCMSDeploymentConfig{
		AdapterVersion: version,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector: {
				Canceller:        SingleGroupMCMS(),
				Bypasser:         SingleGroupMCMS(),
				Proposer:         SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(qualifier),
				TimelockAdmin:    timelockAdmin,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("deploying MCMS: %w", err)
	}
	runningDS.Merge(output.DataStore.Seal())
	tmp.Merge(output.DataStore.Seal())
	env.DataStore = tmp.Seal()

	finalizeOutput, err := fcs.Apply(*env, deployops.MCMSDeploymentConfig{
		AdapterVersion: version,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector: {
				Canceller:        SingleGroupMCMS(),
				Bypasser:         SingleGroupMCMS(),
				Proposer:         SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(qualifier),
				TimelockAdmin:    timelockAdmin,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("finalizing MCMS deployment: %w", err)
	}
	runningDS.Merge(finalizeOutput.DataStore.Seal())

	return runningDS.Seal(), nil
}

func SingleGroupMCMS() mcmstypes.Config {
	publicKey := TestXXXMCMSSigner.Public().(*ecdsa.PublicKey)
	// Convert the public key to an Ethereum address
	address := crypto.PubkeyToAddress(*publicKey)
	c, err := mcmstypes.NewConfig(1, []common.Address{address}, []mcmstypes.Config{})
	if err != nil {
		panic(err)
	}
	return c
}

// getTokenPricesForChain returns the token prices for fee tokens on a chain.
// Uses the optional TokenPriceProvider interface - only chains that implement it (like EVM) provide prices.
// Resolves contract types to addresses using the datastore.
func getTokenPricesForChain(ds datastore.DataStore, selector uint64, version *semver.Version) map[string]*big.Int {
	family, err := chainsel.GetSelectorFamily(selector)
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

func AddNodesToContracts(
	ctx context.Context,
	e *deployment.Environment,
	cls []*simple_node_set.Input,
	nodeKeyBundles map[string]map[string]clclient.NodeKeysBundle,
	ccipHomeSelector uint64, remoteSelectors []uint64,
	homeChainType string,
	bcs []*blockchain.Input,
) error {
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

	// check if ccipHome is owned by a timelock
	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: ccipHomeSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, ccipHomeSelector, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("finding CCIPHome address for selector %d: %w", ccipHomeSelector, err)
	}
	ccipHome, err := ccip_home.NewCCIPHome(
		common.HexToAddress(ccipHomeAddr.Address),
		e.BlockChains.EVMChains()[ccipHomeSelector].Client)
	if err != nil {
		return fmt.Errorf("creating CCIPHome instance for selector %d: %w", ccipHomeSelector, err)
	}
	ccipHomeOwner, err := ccipHome.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("getting CCIPHome owner for selector %d: %w", ccipHomeSelector, err)
	}

	var mcmsInput mcms.Input
	// we consider that if the owner is not the deployer key, then it's timelock owned and we need to use MCMS
	// this is generally true in case of forked chains from mainnet/testnet and only possible with anvil chains in tests
	if ccipHomeOwner != e.BlockChains.EVMChains()[ccipHomeSelector].DeployerKey.From {
		mcmsInput = mcms.Input{
			ValidUntil:     4126214326,
			TimelockAction: mcms_types.TimelockActionBypass,
			Qualifier:      "CLLCCIP",
			Description:    "Home chain CCIP configuration",
		}
	}

	csOut, err := UpdateChainConfig.Apply(*e, UpdateChainConfigConfig{
		HomeChainSelector: ccipHomeSelector,
		RemoteChainAdds:   chainConfigs,
		MCMS:              mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("updating chain config for selector %d: %w", ccipHomeSelector, err)
	}
	if len(csOut.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, csOut.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for updating chain config %d: %w", ccipHomeSelector, err)
		}
		e.Logger.Info("Successfully updated chain config through proposal execution")
	}

	csOut, err = AddDONAndSetCandidate.Apply(*e, AddDonAndSetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: SetCandidatePluginInfo{
			OCRConfigPerRemoteChainSelector: commitOCRConfigs,
			PluginType:                      ccipocr3common.PluginTypeCCIPCommit,
		},
		NonBootstraps:  workerNodes,
		NodeKeyBundles: nodeKeyBundles,
		MCMS:           mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("adding DON and setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	if len(csOut.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, csOut.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for adding DON and setting candidate for selector %d: %w", ccipHomeSelector, err)
		}
		e.Logger.Info("Successfully added DON/set candidate for commit OCR through proposal execution")
	}

	csOut, err = SetCandidate.Apply(*e, SetCandidateChangesetConfig{
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
		MCMS:           mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	if len(csOut.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, csOut.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for setting candidate: %w", err)
		}
		e.Logger.Info("Successfully set candidate for exec OCR through proposal execution")
	}
	csOut, err = PromoteCandidate.Apply(*e, PromoteCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		PluginInfo: []PromoteCandidatePluginInfo{
			{
				PluginType:           ccipocr3common.PluginTypeCCIPExec,
				RemoteChainSelectors: remoteSelectors,
			},
			{
				PluginType:           ccipocr3common.PluginTypeCCIPCommit,
				RemoteChainSelectors: remoteSelectors,
			},
		},
		NonBootstraps: workerNodes,
		MCMS:          mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("promoting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	if len(csOut.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, csOut.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for promoting candidate: %w", err)
		}
		e.Logger.Info("Successfully promoted candidate through proposal execution")
	}
	dReg := deployops.GetRegistry()
	mcmsRegistry := changesetscore.GetRegistry()
	csOut, err = deployops.SetOCR3Config(dReg, mcmsRegistry).Apply(*e, deployops.SetOCR3ConfigArgs{
		HomeChainSel:    ccipHomeSelector,
		RemoteChainSels: remoteSelectors,
		ConfigType:      cciputils.ConfigTypeActive,
		MCMS:            mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("setting OCR3 config for selector %d: %w", ccipHomeSelector, err)
	}
	if len(csOut.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, csOut.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for setting ocr3 config: %w", err)
		}
		e.Logger.Info("Successfully set OCR3 config through proposal execution")
	}
	return nil
}

func AddNodesToCapReg(
	ctx context.Context,
	env *deployment.Environment,
	cls []*simple_node_set.Input,
	bcs []*blockchain.Input,
	ccipHomeSelector uint64,
	needMcms bool,
) error {
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
	homeChain, ok := env.BlockChains.EVMChains()[ccipHomeSelector]
	if !ok {
		return fmt.Errorf("evm chain not found for selector %d", ccipHomeSelector)
	}
	var mcmsInput mcms.Input
	if needMcms {
		mcmsInput = mcms.Input{
			ValidUntil:     4126214326,
			TimelockAction: mcms_types.TimelockActionBypass,
			Qualifier:      "CLLCCIP",
			Description:    "Add node operators to Capabilities Registry",
		}
	}
	// get chain id
	homeChainDetails, ok := chainsel.ChainBySelector(ccipHomeSelector)
	if !ok {
		return fmt.Errorf("could not get chain details for selector %d", ccipHomeSelector)
	}
	var homeChainType string
	for _, bc := range bcs {
		if bc.ChainID == fmt.Sprintf("%d", homeChainDetails.EvmChainID) {
			homeChainType = bc.Type
			break
		}
	}
	// add capability
	out, err := AddCapabilityToCapabilitiesRegistry.Apply(*env, AddCapabilityToCapabilitiesRegistryChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		MCMS:              mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("adding capability to capabilities registry for selector %d: %w", ccipHomeSelector, err)
	}
	if len(out.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, out.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for adding capability to capabilities registry for selector %d: %w", ccipHomeSelector, err)
		}
		env.Logger.Info("Successfully added capability to Capabilities Registry through proposal execution")
	}
	// Add Node Operator
	out, err = AddNodeOperatorsToCapabilitiesRegistry.Apply(*env, AddNodeOperatorsToCapabilitiesRegistryChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		Nop: []capabilities_registry.CapabilitiesRegistryNodeOperator{
			{
				Admin: homeChain.DeployerKey.From,
				Name:  "TestNodeOperator",
			},
		},
		MCMS: mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("adding node operators to capabilities registry for selector %d: %w", ccipHomeSelector, err)
	}
	if len(out.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, out.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for adding node operators to capabilities registry for selector %d: %w", ccipHomeSelector, err)
		}
		env.Logger.Info("Successfully added node operators to Capabilities Registry through proposal execution")
	}
	out, err = AddNodesToCapabilitiesRegistry.Apply(*env, AddNodesToCapabilitiesRegistryChangesetConfig{
		HomeChainSelector:        ccipHomeSelector,
		NodeP2PIDsPerNodeOpAdmin: map[string][][32]byte{"TestNodeOperator": readers},
		MCMS:                     mcmsInput,
	})
	if err != nil {
		return fmt.Errorf("adding nodes to capabilities registry for selector %d: %w", ccipHomeSelector, err)
	}
	if len(out.MCMSTimelockProposals) > 0 {
		if homeChainType != blockchain.TypeAnvil {
			return errors.New("timelock proposals are only supported on Anvil home chains in tests")
		}
		err = blockchainutils.ProcessMCMSProposalsWithTimelockForAnvil(ctx, bcs, out.MCMSTimelockProposals)
		if err != nil {
			return fmt.Errorf("processing MCMS timelock proposals for adding nodes to capabilities registry for selector %d: %w", ccipHomeSelector, err)
		}
		env.Logger.Info("Successfully added node operators to Capabilities Registry through proposal execution")
	}
	return nil
}

func SetupTokensAndTokenPools(env *deployment.Environment, adp []testadapters.TestAdapter) (datastore.DataStore, error) {
	// Get registries and define v1.6.0 alias
	mcmsRegistry := changesetscore.GetRegistry()
	v1_6_0 := semver.MustParse("1.6.0")

	// Ensure all MCMS readers are registered.
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilySolana, &solseq.SolanaAdapter{})

	// To simplify testing, we disable rate limiting on all token transfers.
	disabledRL := tokensapi.RateLimiterConfigFloatInput{Capacity: 0, Rate: 0, IsEnabled: false}

	// This will only store the addresses deployed during this setup.
	outputDS := datastore.NewMemoryDataStore()

	// A helper to create MCMS inputs with long expirations and short timelock delays.
	newInputForMCMS := func(desc string) mcms.Input {
		return mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           math.MaxUint32,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            cciputils.CLLQualifier,
			Description:          desc,
		}
	}

	// A helper to update both the environment datastore and output datastore after each operation.
	mergeDS := func(deltaDS datastore.MutableDataStore) error {
		if deltaDS != nil {
			fullDS := datastore.NewMemoryDataStore()
			if err := outputDS.Merge(deltaDS.Seal()); err != nil {
				return fmt.Errorf("failed to update output datastore: %w", err)
			}
			if err := fullDS.Merge(deltaDS.Seal()); err != nil {
				return fmt.Errorf("failed to merge delta datastore: %w", err)
			}
			if err := fullDS.Merge(env.DataStore); err != nil {
				return fmt.Errorf("failed to merge environment datastore: %w", err)
			}
			env.DataStore = fullDS.Seal()
		}
		return nil
	}

	// Filter out adapters that don't support token transfers (e.g. TON message-passing only).
	var tokenAdapters []testadapters.TestAdapter
	for _, a := range adp {
		if _, err := a.GetRegistryAddress(); errors.Is(err, errors.ErrUnsupported) {
			continue
		}
		tokenAdapters = append(tokenAdapters, a)
	}

	// The deployment map defines the tokens and token pools to deploy
	// and the configurations for their cross-chain interactions. We deploy one token and token pool per chain, and configure them to be transferable to each other.
	dply := map[uint64]tokensapi.TokenExpansionInputPerChain{}
	for _, srcAdapter := range tokenAdapters {
		srcCfg := srcAdapter.GetTokenExpansionConfig()
		srcSel := srcAdapter.ChainSelector()
		srcFamily := srcAdapter.Family()
		if srcFamily != chainsel.FamilyEVM && srcFamily != chainsel.FamilySolana {
			continue // only EVM and Solana are supported for token transfers in 1.6
		}

		for _, dstAdapter := range tokenAdapters {
			// dstCfg := dstAdapter.GetTokenExpansionConfig()
			dstSel := dstAdapter.ChainSelector()

			if srcSel != dstSel {
				srcCfg.TokenTransferConfig.RemoteChains[dstSel] = tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					OutboundRateLimiterConfig: disabledRL,
				}
			}
		}
		dply[srcSel] = srcCfg
	}

	// Deploy one token and its corresponding token pool for all the input chains.
	out, err := tokensapi.TokenExpansion().Apply(*env,
		tokensapi.TokenExpansionInput{
			ChainAdapterVersion:         v1_6_0,
			TokenExpansionInputPerChain: dply,
			MCMS:                        newInputForMCMS("Token Expansion"),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy tokens and token pools: %w", err)
	}
	if err = mergeDS(out.DataStore); err != nil {
		return nil, fmt.Errorf("failed to merge datastore after token expansion: %w", err)
	}

	// Allow the router to withdraw a sensible amount of tokens from the account that will be transferring tokens.
	for _, adapter := range tokenAdapters {
		teConfig := adapter.GetTokenExpansionConfig()
		selector := adapter.ChainSelector()

		oneToken := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(teConfig.DeployTokenInput.Decimals)), nil)
		maxAllow := new(big.Int).Mul(oneToken, big.NewInt(10_000)) // allow 10,000 tokens to be withdrawn
		tokenRef := datastore.AddressRef{
			Type:          datastore.ContractType(teConfig.DeployTokenInput.Type),
			Qualifier:     teConfig.DeployTokenInput.Symbol,
			ChainSelector: selector,
		}
		tokenPoolRef := datastore.AddressRef{
			Type:          datastore.ContractType(teConfig.DeployTokenPoolInput.PoolType),
			Qualifier:     teConfig.DeployTokenPoolInput.TokenPoolQualifier,
			ChainSelector: selector,
		}

		token, err := datastore_utils.FindAndFormatRef(env.DataStore, tokenRef, selector, datastore_utils.FullRef)
		if err != nil {
			return nil, fmt.Errorf("finding token address for selector %d: %w", selector, err)
		}

		err = adapter.AllowRouterToWithdrawTokens(env.GetContext(), token.Address, maxAllow)
		if err != nil {
			return nil, fmt.Errorf("failed to allow router to withdraw tokens for selector %d: %w", selector, err)
		}
		for _, dst := range tokenAdapters {
			if dst.ChainSelector() == selector {
				continue
			}
			dstTeConfig := dst.GetTokenExpansionConfig()
			dstTokenRef := datastore.AddressRef{
				Type:          datastore.ContractType(dstTeConfig.DeployTokenInput.Type),
				Qualifier:     dstTeConfig.DeployTokenInput.Symbol,
				ChainSelector: dst.ChainSelector(),
			}
			dstTokenPoolRef := datastore.AddressRef{
				Type:          datastore.ContractType(dstTeConfig.DeployTokenPoolInput.PoolType),
				Qualifier:     dstTeConfig.DeployTokenPoolInput.TokenPoolQualifier,
				ChainSelector: dst.ChainSelector(),
			}
			// Set some rate limiters for testing - one enabled and one disabled.
			rls := []tokensapi.RateLimiterConfigFloatInput{
				{
					Capacity:  1234567.89, // this is in tokens, not decimal adjusted
					Rate:      123456.789, // this is in tokens per second, not decimal adjusted
					IsEnabled: true,
				},
				{
					IsEnabled: false,
				},
			}
			for _, rl := range rls {
				input := tokensapi.TPRLInput{
					Configs: map[uint64]tokensapi.TPRLConfig{
						selector: {
							ChainAdapterVersion: v1_6_0,
							TokenRef:            tokenRef,
							TokenPoolRef:        tokenPoolRef,
							RemoteOutbounds: map[uint64]tokensapi.RateLimiterConfigFloatInput{
								dst.ChainSelector(): rl,
							},
						},
						dst.ChainSelector(): {
							ChainAdapterVersion: v1_6_0,
							TokenRef:            dstTokenRef,
							TokenPoolRef:        dstTokenPoolRef,
							RemoteOutbounds: map[uint64]tokensapi.RateLimiterConfigFloatInput{
								selector: rl,
							},
						},
					},
				}
				_, err = tokensapi.SetTokenPoolRateLimits().Apply(*env, input)
				if err != nil {
					return nil, fmt.Errorf("setting rate limiter for token %s on selector %d to selector %d: %w",
						teConfig.DeployTokenInput.Symbol, selector, dst.ChainSelector(), err)
				}
			}

		}

	}

	// Return only the newly deployed addresses from this setup.
	return outputDS.Seal(), nil
}
