package ccip

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	devenvcommon "github.com/smartcontractkit/chainlink-ccip/devenv/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/jd"

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

// NewEnvironment creates a new CCIP environment either locally in Docker or remotely in K8s.
func NewEnvironment() (*Cfg, error) {
	ctx := context.Background()
	tr := NewTimeTracker(Plog)
	ctx = L.WithContext(ctx)
	if err := framework.DefaultNetwork(nil); err != nil {
		return nil, err
	}

	in, err := Load[Cfg](strings.Split(os.Getenv(EnvVarTestConfigs), ","))
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	if err := checkKeys(in); err != nil {
		return nil, err
	}

	impls := make([]CCIP16ProductConfiguration, 0)
	for _, bc := range in.Blockchains {
		impl, err := NewCCIPImplFromNetwork(bc.Type)
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

	// initialize CLDF framework
	in.CLDF.Init()
	selectors, e, err := NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	if err != nil {
		return nil, fmt.Errorf("creating CLDF operations environment: %w", err)
	}
	L.Info().Any("Selectors", selectors).Msg("Deploying for chain selectors")
	ds := datastore.NewMemoryDataStore()
	ds.Merge(e.DataStore)

	// Deploy Capabilities Registry
	crAddr, tx, _, err := capabilities_registry.DeployCapabilitiesRegistry(
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
	allConfigs := strings.Join(clChainConfigs, "\n")
	for _, nodeSpec := range in.NodeSets[0].NodeSpecs {
		nodeSpec.Node.TestConfigOverrides = allConfigs
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
	// deploy all the contracts
	for i, impl := range impls {
		nkb, err := devenvcommon.CreateNodeKeysBundle(allNodeClients, in.Blockchains[i].Type, in.Blockchains[i].ChainID)
		if err != nil {
			return nil, fmt.Errorf("creating node keys bundle: %w", err)
		}
		err = impl.FundNodes(ctx, in.NodeSets, nkb, in.Blockchains[i], big.NewInt(1), big.NewInt(5))
		if err != nil {
			return nil, fmt.Errorf("funding nodes: %w", err)
		}
		var family string
		switch in.Blockchains[i].Type {
		case "anvil", "geth":
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
		networkInfo, err := chainsel.GetChainDetailsByChainIDAndFamily(in.Blockchains[i].ChainID, family)
		if err != nil {
			return nil, err
		}
		L.Info().Uint64("Selector", networkInfo.ChainSelector).Msg("Deployed chain selector")
		err = impl.PreDeployContractsForSelector(ctx, e, in.NodeSets, networkInfo.ChainSelector, CCIPHomeChain, crAddr.String())
		if err != nil {
			return nil, err
		}
		dsi, err := devenvcommon.DeployContractsForSelector(ctx, e, in.NodeSets, networkInfo.ChainSelector, CCIPHomeChain, crAddr.String())
		if err != nil {
			return nil, err
		}
		err = impl.PostDeployContractsForSelector(ctx, e, in.NodeSets, networkInfo.ChainSelector, CCIPHomeChain, crAddr.String())
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
	e.DataStore = ds.Seal()

	err = devenvcommon.ConfigureContractsForSelectors(ctx, e, in.NodeSets, nodeKeyBundles, CCIPHomeChain, selectors)
	if err != nil {
		return nil, err
	}

	err = CreateJobs(ctx, allNodeClients, nodeKeyBundles)
	if err != nil {
		return nil, fmt.Errorf("creating CCIP jobs: %w", err)
	}

	// connect all the contracts together (on-ramps, off-ramps)
	for i, _ := range impls {
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

	tr.Record("[infra] deployed CL nodes")
	tr.Record("[changeset] deployed product contracts")

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
