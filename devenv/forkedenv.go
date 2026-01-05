package ccip

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/jd"
	ns "github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	devenvcommon "github.com/smartcontractkit/chainlink-ccip/devenv/common"
)

func checkForkedEnvIsSet(in *Cfg) error {
	if in.ForkedEnvConfig == nil {
		return errors.New("forked_env_config is not set in the configuration")
	}
	if in.ForkedEnvConfig.ForkURL == "" {
		return errors.New("fork_url is not set in the forked_env_config configuration")
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
		in.Blockchains[i].DockerCmdParamsOverrides = append(in.Blockchains[i].DockerCmdParamsOverrides,
			"--fork-url "+in.ForkedEnvConfig.ForkURL)
	}
	return nil
}

func NewForkedEnvironment() (*Cfg, error) {
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
	if err := checkForkedEnvIsSet(in); err != nil {
		return nil, err
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

	// Load DataStore with addresses from config
	cldRootPath := in.ForkedEnvConfig.CLDRootPath
	cldEnvKey := in.ForkedEnvConfig.CLDEnvironment
	L.Info().Str("CLDPath", cldRootPath).Str("CLDEnvKey", cldEnvKey).Msg("Loading CLDF data store from configuration")
	envDir := domain.NewEnvDir(cldRootPath, CLDDomain, cldEnvKey)
	ds, err := envDir.DataStore()
	if err != nil {
		return nil, fmt.Errorf("loading CLD data store from env dir: %w", err)
	}
	// initialize CLDF framework
	in.CLDF.Init(WithDataStore(ds))
	selectors, e, err := NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	if err != nil {
		return nil, fmt.Errorf("creating CLDF operations environment: %w", err)
	}
	L.Info().Any("Selectors", selectors).Msg("Deploying for chain selectors")

	// get Capability Registry Address
	refs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByType("CapabilitiesRegistry"),
		datastore.AddressRefByVersion(semver.MustParse("1.1.0")))
	if len(refs) != 1 {
		return nil, fmt.Errorf("expected exactly one CapabilitiesRegistry address ref for version 1.1.0, found %d %+v", len(refs), refs)
	}
	crRef := refs[0].Address

	tr.Record("[infra] deploying blockchains")

	clChainConfigs := make([]string, 0)
	clChainConfigs = append(clChainConfigs, getCommonNodeConfig(crRef))
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
	for i, impl := range impls {
		nkb, err := impl.FundNodes(ctx, in.NodeSets, in.Blockchains[i], big.NewInt(1), big.NewInt(5))
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
	}
	err = impls[0].ConfigureContractsForSelectors(ctx, e, in.NodeSets, nodeKeyBundles, CCIPHomeChain, selectors)
	if err != nil {
		return nil, err
	}
	nodeClients, err := clclient.New(in.NodeSets[0].Out.CLNodes)
	if err != nil {
		return nil, fmt.Errorf("connecting to CL nodes: %w", err)
	}
	bootstrapNode := nodeClients[0]
	bootstrapKeys, err := bootstrapNode.MustReadOCR2Keys()
	if err != nil {
		return nil, fmt.Errorf("reading bootstrap node OCR keys: %w", err)
	}
	// bootstrap is 0
	workerNodes := nodeClients[1:]

	// create jobs post-deployment for home chain
	bootstrapP2PKeys, err := bootstrapNode.MustReadP2PKeys()
	if err != nil {
		return nil, fmt.Errorf("reading worker node P2P keys: %w", err)
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
		return nil, fmt.Errorf("creating CCIP job spec: %w", err)
	}
	_, _, err = bootstrapNode.CreateJobRaw(raw)
	if err != nil {
		return nil, fmt.Errorf("creating CCIP job: %w", err)
	}
	for _, node := range workerNodes {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return nil, fmt.Errorf("reading worker node P2P keys: %w", err)
		}
		L.Info().Str("Node", node.Config.URL).Any("PeerIDs", nodeP2PIds).Msg("Adding worker peer ID")
		ocrKeys, err := node.MustReadOCR2Keys()
		if err != nil {
			return nil, fmt.Errorf("reading worker node OCR keys: %w", err)
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
			return nil, fmt.Errorf("creating CCIP job spec: %w", err)
		}
		L.Info().Str("RawSpec", raw).Msg("Creating CCIP job on worker node")
		_, _, err = node.CreateJobRaw(raw)
		if err != nil {
			return nil, fmt.Errorf("creating CCIP job: %w", err)
		}
	}

	tr.Record("[infra] deployed CL nodes")
	tr.Record("[changeset] configured product contracts")

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
