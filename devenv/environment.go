package ccip

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/jd"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	devenvcommon "github.com/smartcontractkit/chainlink-ccip/devenv/common"

	chainsel "github.com/smartcontractkit/chain-selectors"
	ns "github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"
)

var Plog = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel).With().Fields(map[string]any{"component": "ccip"}).Logger()

func getCommonNodeConfig(capRegAddr string) string {
	return fmt.Sprintf(`
			[Log]
			JSONConsole = true
			Level = 'debug'
			[Pyroscope]
			ServerAddress = 'http://host.docker.internal:4040'
			Environment = 'local'
			[WebServer]
			SessionTimeout = '999h0m0s'
			HTTPWriteTimeout = '3m'
			SecureCookies = false
			HTTPPort = 6688
			[WebServer.TLS]
			HTTPSPort = 0
			[WebServer.RateLimit]
			Authenticated = 5000
			Unauthenticated = 5000
			[JobPipeline]
			[JobPipeline.HTTPRequest]
			DefaultTimeout = '1m'
			[Log.File]
			MaxSize = '0b'
			[Feature]
			FeedsManager = true
			LogPoller = true
			UICSAKeys = true
			[OCR2]
			Enabled = true
			SimulateTransactions = false
			DefaultTransactionQueueDepth = 1
			[P2P.V2]
			Enabled = true
			ListenAddresses = ['0.0.0.0:6690']
			[Capabilities.ExternalRegistry]
			Address = '%s'
			NetworkID = 'evm'
			ChainID = '1337'
`, capRegAddr)
}

type Cfg struct {
	CLDF               CLDF                `toml:"cldf"                  validate:"required"`
	JD                 *jd.Input           `toml:"jd"`
	Blockchains        []*blockchain.Input `toml:"blockchains"           validate:"required"`
	NodeSets           []*ns.Input         `toml:"nodesets"              validate:"required"`
	CLNodesFundingETH  float64             `toml:"cl_nodes_funding_eth"`
	CLNodesFundingLink float64             `toml:"cl_nodes_funding_link"`
	// NodeConfigOverrides allows users to provide custom TOML configuration that will be
	// appended to the generated node configuration. This is useful for testnets/mainnets
	// where you need to customize settings like FinalityDepth, LogPollInterval, etc.
	// The overrides are applied AFTER the auto-generated config, so they take precedence.
	NodeConfigOverrides string           `toml:"node_config_overrides"`
	ForkedEnvConfig     *ForkedEnvConfig `toml:"forked_env_config"`
}

type ForkedEnvConfig struct {
	ForkURLs          map[string]string `toml:"fork_urls_by_chain_id"`
	ForkBlockNumbers  map[string]uint64 `toml:"fork_block_numbers_by_chain_id"`
	HomeChainSelector uint64            `toml:"home_chain_selector"`
	CLDRootPath       string            `toml:"cld_root_path"`
	CLDEnvironment    string            `toml:"cld_environment"`
}

func checkKeys(in *Cfg) error {
	if getNetworkPrivateKey() != DefaultAnvilKey && in.Blockchains[0].ChainID == "1337" && in.Blockchains[1].ChainID == "2337" {
		return errors.New("you are trying to run simulated chains with a key that do not belong to Anvil, please run 'unset PRIVATE_KEY'")
	}
	if getNetworkPrivateKey() == DefaultAnvilKey && in.Blockchains[0].ChainID != "1337" && in.Blockchains[1].ChainID != "2337" {
		return errors.New("you are trying to run on real networks but is not using the Anvil private key, export your private key 'export PRIVATE_KEY=...'")
	}
	return nil
}

func checkForkedEnvIsSet(in *Cfg) error {
	if in.ForkedEnvConfig == nil {
		return errors.New("forked_env_config is not set in the configuration")
	}
	if len(in.ForkedEnvConfig.ForkURLs) == 0 {
		return errors.New("fork_urls are not set in the forked_env_config configuration")
	}
	if in.ForkedEnvConfig.CLDRootPath == "" {
		return errors.New("cld_root_path is not set in the forked_env_config configuration")
	}
	if in.ForkedEnvConfig.CLDEnvironment == "" {
		return errors.New("cld_environment is not set in the forked_env_config configuration")
	}
	if in.ForkedEnvConfig.HomeChainSelector == 0 {
		return errors.New("home_chain_selector is not set in the forked_env_config configuration")
	}
	for i, bc := range in.Blockchains {
		if bc.Type != "anvil" {
			return fmt.Errorf("blockchain %s is not supported in forked environment, only anvil is supported", bc.Type)
		}
		forkURL, ok := in.ForkedEnvConfig.ForkURLs[bc.ChainID]
		if !ok || forkURL == "" {
			return fmt.Errorf("fork_url for chain_id %s is not set in the forked_env_config configuration", bc.ChainID)
		}
		u, err := url.Parse(forkURL)
		if err != nil {
			return fmt.Errorf("invalid fork_url for chain_id %s: %w", bc.ChainID, err)
		}
		if u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("invalid fork_url for chain_id %s: %s", bc.ChainID, forkURL)
		}
		forkedArgs := []string{"--fork-url", in.ForkedEnvConfig.ForkURLs[bc.ChainID]}
		if in.ForkedEnvConfig.ForkBlockNumbers != nil && in.ForkedEnvConfig.ForkBlockNumbers[bc.ChainID] != 0 {
			forkedArgs = append(forkedArgs, "--fork-block-number", fmt.Sprintf("%d", in.ForkedEnvConfig.ForkBlockNumbers[bc.ChainID]))
		}
		in.Blockchains[i].DockerCmdParamsOverrides = append(forkedArgs, in.Blockchains[i].DockerCmdParamsOverrides...)
	}
	return nil
}

// NewEnvironment creates a new CCIP environment either locally in Docker or remotely in K8s.
func NewEnvironment() (*Cfg, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancelFunc()
	tr := NewTimeTracker(Plog)
	ctx = L.WithContext(ctx)
	if err := framework.DefaultNetwork(nil); err != nil {
		return nil, err
	}

	in, err := Load[Cfg](strings.Split(os.Getenv(EnvVarTestConfigs), ","))
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	if in.ForkedEnvConfig == nil {
		if err := checkKeys(in); err != nil {
			return nil, err
		}
	} else {
		if err := checkForkedEnvIsSet(in); err != nil {
			return nil, err
		}
	}

	impls := make([]CCIP16ProductConfiguration, 0)
	for _, bc := range in.Blockchains {
		var family string
		switch bc.Type {
		case "anvil", "geth":
			family = chainsel.FamilyEVM
		case "solana":
			family = chainsel.FamilySolana
		case "ton":
			family = chainsel.FamilyTon
		default:
			return nil, fmt.Errorf("unsupported blockchain type: %s", bc.Type)
		}
		impl, err := NewCCIPImplFromNetwork(family, bc.ChainID)
		if err != nil {
			return nil, err
		}
		impls = append(impls, impl)
	}
	for i, impl := range impls {
		_, err := impl.DeployLocalNetwork(ctx, in.Blockchains[i])
		if err != nil {
			return nil, fmt.Errorf("failed to deploy local networks: %w", err)
		}
	}

	var initOpts []InitOption
	if in.ForkedEnvConfig != nil {
		// Load DataStore with addresses from config
		cldRootPath := in.ForkedEnvConfig.CLDRootPath
		cldEnvKey := in.ForkedEnvConfig.CLDEnvironment
		L.Info().Str("CLDPath", cldRootPath).Str("CLDEnvKey", cldEnvKey).Msg("Loading CLDF data store from configuration")
		ccipDomain := domain.NewDomain(cldRootPath, CLDDomain)
		envDir := ccipDomain.EnvDir(cldEnvKey)
		envDS, err := envDir.DataStore()
		if err != nil {
			return nil, fmt.Errorf("loading CLD data store from env dir: %w", err)
		}
		initOpts = append(initOpts, WithDataStore(envDS))
	}

	// initialize CLDF framework
	in.CLDF.Init(initOpts...)
	selectors, e, err := NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	if err != nil {
		return nil, fmt.Errorf("creating CLDF operations environment: %w", err)
	}

	L.Info().Any("Selectors", selectors).Msg("Deploying for chain selectors")
	var crAddr common.Address
	ds := datastore.NewMemoryDataStore()
	var homeChainSelector uint64
	if in.ForkedEnvConfig != nil {
		// get Capability Registry Address
		refs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByType(datastore.ContractType(utils.CapabilitiesRegistry)),
			datastore.AddressRefByVersion(semver.MustParse("1.0.0")),
			datastore.AddressRefByChainSelector(in.ForkedEnvConfig.HomeChainSelector),
		)
		if len(refs) != 1 {
			return nil, fmt.Errorf("expected exactly one CapabilitiesRegistry address ref for version 1.1.0, found %d %+v", len(refs), refs)
		}
		crRef := refs[0].Address

		// check if addresses exist
		homeChainSelector = in.ForkedEnvConfig.HomeChainSelector
		if _, ok := e.BlockChains.EVMChains()[homeChainSelector]; !ok {
			return nil, fmt.Errorf("home chain selector %d not found in env evm chains %+v", homeChainSelector, e.BlockChains.EVMChains())
		}
		homechain := e.BlockChains.EVMChains()[homeChainSelector].Client
		capReg, err := capabilities_registry.NewCapabilitiesRegistry(common.HexToAddress(crRef), homechain)
		if err != nil {
			return nil, fmt.Errorf("creating capabilities registry instance: %w", err)
		}
		tv, err := capReg.TypeAndVersion(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("getting capabilities registry type and version: %w", err)
		}
		if !strings.Contains(tv, "CapabilitiesRegistry") {
			return nil, fmt.Errorf("unexpected capabilities registry type and version: %s", tv)
		}

		L.Info().Str("CapabilitiesRegistry", crRef).Str("TypeAndVersion", tv).Msg("Connected to Capabilities Registry")
		crAddr = common.HexToAddress(crRef)
	} else {
		L.Info().Msg("Deploying new Capabilities Registry")
		err = ds.Merge(e.DataStore)
		if err != nil {
			return nil, err
		}
		homeChainSelector = CCIPHomeChain
		var tx *types.Transaction
		// Deploy Capabilities Registry
		crAddr, tx, _, err = capabilities_registry.DeployCapabilitiesRegistry(
			e.BlockChains.EVMChains()[CCIPHomeChain].DeployerKey,
			e.BlockChains.EVMChains()[CCIPHomeChain].Client,
		)
		if err != nil {
			return nil, fmt.Errorf("deploying capabilities registry: %w", err)
		}
		_, err = e.BlockChains.EVMChains()[CCIPHomeChain].Confirm(tx)
		if err != nil {
			return nil, fmt.Errorf("confirming capabilities registry deployment: %w", err)
		}
	}

	tr.Record("[infra] deploying blockchains")

	clChainConfigs := make([]string, 0)
	clChainConfigs = append(clChainConfigs, getCommonNodeConfig(crAddr.String()))
	for i, impl := range impls {
		clChainConfig, err := impl.ConfigureNodes(ctx, in.Blockchains[i])
		if err != nil {
			return nil, fmt.Errorf("failed to deploy local networks: %w", err)
		}
		clChainConfigs = append(clChainConfigs, clChainConfig)
	}
	generatedConfigs := strings.Join(clChainConfigs, "\n")

	// Apply configs to nodes, preserving user overrides
	// Order of precedence (later overrides earlier):
	// 1. Generated configs (common + chain-specific)
	// 2. Top-level NodeConfigOverrides from Cfg
	// 3. Per-node TestConfigOverrides from node specs
	for _, nodeSpec := range in.NodeSets[0].NodeSpecs {
		configParts := []string{generatedConfigs}

		// Add top-level user overrides if provided
		if in.NodeConfigOverrides != "" {
			configParts = append(configParts, in.NodeConfigOverrides)
		}

		// Preserve per-node overrides if they were provided in the config file
		if nodeSpec.Node.TestConfigOverrides != "" {
			configParts = append(configParts, nodeSpec.Node.TestConfigOverrides)
		}

		nodeSpec.Node.TestConfigOverrides = strings.Join(configParts, "\n")
	}
	Plog.Info().Msg("Nodes network configuration is generated")

	prodJDImage := os.Getenv("JD_IMAGE")

	if in.JD != nil {
		if prodJDImage != "" {
			in.JD.Image = prodJDImage
		}
		if len(in.JD.Image) == 0 {
			Plog.Warn().Msg("No JD image provided, skipping JD service startup")
		} else {
			_, err = jd.NewJD(in.JD)
			if err != nil {
				return nil, fmt.Errorf("failed to create JD service: %w", err)
			}
		}
	} else {
		Plog.Warn().Msg("No JD configuration provided, skipping JD service startup")
	}

	// connect JD to nodes here

	tr.Record("[changeset] configured nodes network")
	_, err = ns.NewSharedDBNodeSet(in.NodeSets[0], nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new shared db node set: %w", err)
	}

	nodeKeyBundles := make(map[string]map[string]clclient.NodeKeysBundle, 0)
	allNodeClients, err := clclient.New(in.NodeSets[0].Out.CLNodes)
	if err != nil {
		return nil, fmt.Errorf("connecting to CL nodes: %w", err)
	}

	// Pass funding amounts as-is (in native units like ETH, SOL, TON)
	// Each chain implementation handles its own unit conversion:
	// - EVM: ETH -> wei (×10^18)
	// - Solana: SOL -> lamports (×10^9)
	// - TON: TON -> nanotons (×10^9)
	// Using big.Int with the integer part; fractional units handled by each impl
	nativeAmount := big.NewInt(int64(in.CLNodesFundingETH))
	linkAmount := big.NewInt(int64(in.CLNodesFundingLink))

	// deploy all the contracts
	for i, impl := range impls {
		// Create node key bundles for this chain type
		nkb, err := devenvcommon.CreateNodeKeysBundle(allNodeClients, in.Blockchains[i].Type, in.Blockchains[i].ChainID)
		if err != nil {
			return nil, fmt.Errorf("creating node keys bundle: %w", err)
		}

		// Fund nodes with native token
		err = impl.FundNodes(ctx, in.NodeSets, nkb, in.Blockchains[i], linkAmount, nativeAmount)
		if err != nil {
			return nil, fmt.Errorf("funding nodes: %w", err)
		}
		selector := impl.ChainSelector()

		var family string
		switch in.Blockchains[i].Type {
		case "anvil", "geth":
			// NOTE: this seems like a massive hack, why not for EVM?
			family = chainsel.FamilyEVM
		case "solana":
			family = chainsel.FamilySolana
			nodeKeyBundles[family] = nkb
		case "ton":
			family = chainsel.FamilyTon
			nodeKeyBundles[family] = nkb
		default:
			return nil, fmt.Errorf("unsupported blockchain type: %s", in.Blockchains[i].Type)
		}
		if in.ForkedEnvConfig != nil {
			// Skip deployment on forked environments
			L.Info().Str("ChainID", in.Blockchains[i].ChainID).Msg("Skipping contract deployment on forked environment")
			continue
		}
		L.Info().Uint64("Selector", selector).Msg("Deployed chain selector")
		err = impl.PreDeployContractsForSelector(ctx, e, in.NodeSets, selector, CCIPHomeChain, crAddr.String())
		if err != nil {
			return nil, err
		}
		dsi, err := devenvcommon.DeployContractsForSelector(ctx, e, in.NodeSets, selector, CCIPHomeChain, crAddr.String())
		if err != nil {
			return nil, err
		}
		err = impl.PostDeployContractsForSelector(ctx, e, in.NodeSets, selector, CCIPHomeChain, crAddr.String())
		if err != nil {
			return nil, err
		}
		addresses, err := dsi.Addresses().Fetch()
		if err != nil {
			return nil, err
		}
		a, err := json.Marshal(addresses)
		if err != nil {
			return nil, err
		}
		in.CLDF.AddAddresses(string(a))
		if err := ds.Merge(dsi); err != nil {
			return nil, err
		}
	}

	// Make sure datastore has the latest deployed addresses
	if err = ds.Merge(e.DataStore); err != nil {
		return nil, err
	} else {
		e.DataStore = ds.Seal()
	}

	// Populate test adapters
	adapters := make([]testadapters.TestAdapter, 0, len(impls))
	for _, impl := range impls {
		impl.SetCLDF(e)
		adapters = append(adapters, impl)
	}

	// Deploy a token and its corresponding token pool on each chain
	L.Info().Msg("Deploying tokens and token pools for chains")
	newDS, err := devenvcommon.SetupTokensAndTokenPools(e, adapters)
	if err != nil {
		return nil, fmt.Errorf("failed to setup tokens and token pools: %w", err)
	}
	a, err := json.Marshal(newDS.Addresses().Filter())
	if err != nil {
		return nil, err
	}
	in.CLDF.AddAddresses(string(a))

	var homeChainType string
	if in.ForkedEnvConfig != nil {
		homeChainType = blockchain.TypeAnvil
		err = devenvcommon.AddNodesToCapReg(ctx, e, in.NodeSets, in.Blockchains, homeChainSelector, true)
		if err != nil {
			return nil, err
		}
		// Add addresses from initial CLDF data store
		addresses, err := e.DataStore.Addresses().Fetch()
		if err != nil {
			return nil, err
		}
		a, err := json.Marshal(addresses)
		if err != nil {
			return nil, err
		}
		in.CLDF.AddAddresses(string(a))
	} else {
		// Merge newly deployed addresses into environment data store
		e.DataStore = ds.Seal()
		for _, bc := range in.Blockchains {
			if bc.ChainID == fmt.Sprintf("%d", CCIPHomeChainID) {
				homeChainType = bc.Type
				break
			}
		}
	}

	err = CreateJobs(ctx, allNodeClients, nodeKeyBundles)
	if err != nil {
		return nil, fmt.Errorf("creating CCIP jobs: %w", err)
	}

	err = devenvcommon.AddNodesToContracts(ctx, e, in.NodeSets, nodeKeyBundles, homeChainSelector, selectors, homeChainType, in.Blockchains)
	if err != nil {
		return nil, err
	}

	if in.ForkedEnvConfig != nil {
		L.Info().Msg("Skipping connecting contracts on forked environment")
	} else {
		// connect all the contracts together (on-ramps, off-ramps)
		for i := range impls {
			var family string
			switch in.Blockchains[i].Type {
			case "anvil", "geth":
				family = chainsel.FamilyEVM
			case "solana":
				family = chainsel.FamilySolana
			case "ton":
				family = chainsel.FamilyTon
			default:
				return nil, fmt.Errorf("unsupported blockchain type: %s", in.Blockchains[i].Type)
			}
			networkInfo, err := chainsel.GetChainDetailsByChainIDAndFamily(in.Blockchains[i].ChainID, family)
			if err != nil {
				return nil, err
			}
			selsToConnect := make([]uint64, 0)
			for _, sel := range selectors {
				if sel != networkInfo.ChainSelector {
					selsToConnect = append(selsToConnect, sel)
				}
			}
			err = devenvcommon.ConnectContractsWithSelectors(ctx, e, networkInfo.ChainSelector, selsToConnect)
			if err != nil {
				return nil, err
			}
		}
		tr.Record("[changeset] deployed product contracts")
	}
	tr.Record("[infra] deployed CL nodes")

	Plog.Info().Str("BootstrapNode", in.NodeSets[0].Out.CLNodes[0].Node.ExternalURL).Send()
	for _, n := range in.NodeSets[0].Out.CLNodes[1:] {
		Plog.Info().Str("Node", n.Node.ExternalURL).Send()
	}

	if err := PrintCLDFAddresses(in); err != nil {
		return nil, err
	}
	tr.Print()
	return in, Store(in)
}

type SpecArgs struct {
	P2PV2Bootstrappers     []string          `toml:"p2pV2Bootstrappers"`
	CapabilityVersion      string            `toml:"capabilityVersion"`
	CapabilityLabelledName string            `toml:"capabilityLabelledName"`
	OCRKeyBundleIDs        map[string]string `toml:"ocrKeyBundleIDs"`
	P2PKeyID               string            `toml:"p2pKeyID"`
	RelayConfigs           map[string]any    `toml:"relayConfigs"`
	PluginConfig           map[string]any    `toml:"pluginConfig"`
}

// NewCCIPSpecToml creates a new CCIP spec in toml format from the given spec args.
func NewCCIPSpecToml(spec SpecArgs) (string, error) {
	type fullSpec struct {
		SpecArgs
		Type          string `toml:"type"`
		SchemaVersion uint64 `toml:"schemaVersion"`
		Name          string `toml:"name"`
		ExternalJobID string `toml:"externalJobID"`
	}
	extJobID, err := ExternalJobID(spec)
	if err != nil {
		return "", fmt.Errorf("failed to generate external job id: %w", err)
	}
	marshaled, err := toml.Marshal(fullSpec{
		SpecArgs:      spec,
		Type:          "ccip",
		SchemaVersion: 1,
		Name:          fmt.Sprintf("%s-%s", "ccip", extJobID),
		ExternalJobID: extJobID,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal spec into toml: %w", err)
	}

	return string(marshaled), nil
}

func CreateJobs(ctx context.Context, nodeClients []*clclient.ChainlinkClient, nodeKeyBundles map[string]map[string]clclient.NodeKeysBundle) error {
	bootstrapNode := nodeClients[0]
	bootstrapKeys, err := bootstrapNode.MustReadOCR2Keys()
	if err != nil {
		return fmt.Errorf("reading bootstrap node OCR keys: %w", err)
	}
	// bootstrap is 0
	workerNodes := nodeClients[1:]

	// create jobs post-deployment for home chain
	bootstrapP2PKeys, err := bootstrapNode.MustReadP2PKeys()
	if err != nil {
		return fmt.Errorf("reading worker node P2P keys: %w", err)
	}
	bootstrapId := devenvcommon.MustPeerIDFromString(bootstrapP2PKeys.Data[0].Attributes.PeerID)
	ocrKeyBundleIDs := map[string]string{
		"evm": bootstrapKeys.Data[0].ID,
	}
	for family, nkb := range nodeKeyBundles {
		ocrKeyBundleIDs[family] = nkb[bootstrapId.Raw()].OCR2Key.Data.ID
	}
	L.Info().Str("ocrKeyBundleIDs", fmt.Sprintf("%+v", ocrKeyBundleIDs)).Msg("Read OCR keys for bootstrap node")
	raw, err := NewCCIPSpecToml(SpecArgs{
		P2PV2Bootstrappers:     []string{},
		CapabilityVersion:      "v1.0.0",
		CapabilityLabelledName: "ccip",
		OCRKeyBundleIDs:        ocrKeyBundleIDs,
		P2PKeyID:               bootstrapId.String(),
		RelayConfigs:           nil,
		PluginConfig:           map[string]any{},
	})
	if err != nil {
		return fmt.Errorf("creating CCIP job spec: %w", err)
	}
	_, _, err = bootstrapNode.CreateJobRaw(raw)
	if err != nil {
		return fmt.Errorf("creating CCIP job: %w", err)
	}
	for _, node := range workerNodes {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return fmt.Errorf("reading worker node P2P keys: %w", err)
		}
		L.Info().Str("Node", node.Config.URL).Any("PeerIDs", nodeP2PIds).Msg("Adding worker peer ID")
		ocrKeys, err := node.MustReadOCR2Keys()
		if err != nil {
			return fmt.Errorf("reading worker node OCR keys: %w", err)
		}
		L.Info().Str("Node", node.Config.URL).Any("OCRKeys", ocrKeys).Msg("Adding worker OCR keys")
		ocrKeyBundleIDs := map[string]string{
			"evm": ocrKeys.Data[0].ID,
		}
		id := devenvcommon.MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
		for family, nkb := range nodeKeyBundles {
			ocrKeyBundleIDs[family] = nkb[id.Raw()].OCR2Key.Data.ID
		}
		L.Info().Str("ocrKeyBundleIDs", fmt.Sprintf("%+v", ocrKeyBundleIDs)).Msg("Read OCR keys for worker node")
		raw, err := NewCCIPSpecToml(SpecArgs{
			P2PV2Bootstrappers: []string{
				fmt.Sprintf("%s@%s", strings.TrimPrefix(bootstrapId.String(), "p2p_"), "don-node0:6690"),
			},
			CapabilityVersion:      "v1.0.0",
			CapabilityLabelledName: "ccip",
			OCRKeyBundleIDs:        ocrKeyBundleIDs,
			P2PKeyID:               id.String(),
			RelayConfigs:           nil,
			PluginConfig:           map[string]any{},
		})
		if err != nil {
			return fmt.Errorf("creating CCIP job spec: %w", err)
		}
		L.Info().Str("RawSpec", raw).Msg("Creating CCIP job on worker node")
		_, _, err = node.CreateJobRaw(raw)
		if err != nil {
			return fmt.Errorf("creating CCIP job: %w", err)
		}
	}
	return nil
}

func ExternalJobID(spec SpecArgs) (string, error) {
	in := fmt.Appendf(nil, "%s%s%s", spec.CapabilityLabelledName, spec.CapabilityVersion, spec.P2PKeyID)
	sha256Hash := sha256.New()
	sha256Hash.Write(in)
	in = sha256Hash.Sum(nil)[:16]
	// tag as valid UUID v4 https://github.com/google/uuid/blob/0f11ee6918f41a04c201eceeadf612a377bc7fbc/version4.go#L53-L54
	in[6] = (in[6] & 0x0f) | 0x40 // Version 4
	in[8] = (in[8] & 0x3f) | 0x80 // Variant is 10
	id, err := uuid.FromBytes(in)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
