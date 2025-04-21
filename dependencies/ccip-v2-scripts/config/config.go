package config

import (
	"crypto/tls"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	"github.com/smartcontractkit/chainlink/deployment/environment/nodeclient"
	"google.golang.org/grpc/credentials"
)

type ChainType int

// Define constants
const (
	EVM ChainType = iota
	SOLANA
	defaultFinalityDepth = 200
	SolStartingPodId     = 1000
)

type ChainConfigurer struct {
	chainID        string
	deployerKey    string
	solDeployerKey solana.PrivateKey
	env            DevspaceEnv
	chainVariant   string
	chainName      string
	chainType      ChainType
	GasPrice       *big.Int
	podID          int
}

type EVMChain struct {
	NetworkId     string
	FinalityDepth int
}

type ChainlinkNodeConfigurer struct {
	env DevspaceEnv
}

func ChainTypeFromString(s string) ChainType {
	switch strings.ToLower(s) {
	case "evm":
		return EVM
	case "solana":
		return SOLANA
	}
	return EVM
}

// String method to convert enum values to readable strings
func (s ChainType) String() string {
	return [...]string{"EVM", "SOLANA"}[s]
}

func GetEnvConfig(env DevspaceEnv) (*devenv.EnvironmentConfig, error) {
	chainConfigurers := getChainConfigurers(env)

	chainConfigs := make([]devenv.ChainConfig, 0)
	for _, chainConfigurer := range chainConfigurers {
		config, err := chainConfigurer.GetDevenvChainConfig()
		if err != nil {
			return nil, fmt.Errorf("error getting chain config: %v", err)
		}

		chainConfigs = append(chainConfigs, *config)
	}
	nodeInfos := NewCLNodeConfigurer(env).GetNodeInfos()

	var grpcUrl string
	if env.CIEnv {
		hostnameSuffix := "job-distributor"
		grpcUrl = gapV2HostName(env, hostnameSuffix)
	} else {
		grpcUrl = fmt.Sprintf("%s-job-distributor-grpc.%s:443", env.Namespace, env.IngressBaseDomain)
	}

	jdConfig := devenv.JDConfig{
		GRPC:  grpcUrl,
		WSRPC: "job-distributor-noderpc-lb:80",
		Creds: credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		}),
		NodeInfo: nodeInfos,
	}

	return &devenv.EnvironmentConfig{
		Chains:   chainConfigs,
		JDConfig: jdConfig,
	}, nil
}

func GetChainConfigBySelector(configs []devenv.ChainConfig, chainSelector uint64) *devenv.ChainConfig {
	chain, ok := chainsel.ChainBySelector(chainSelector)
	if !ok {
		panic("invalid chain selector")
	}
	for _, chainConfig := range configs {
		chainId, err := strconv.ParseUint(chainConfig.ChainID, 10, 0)
		if err != nil {
			panic(err)
		}
		if chainId == chain.EvmChainID {
			return &chainConfig
		}
	}
	return nil
}

func getChainConfigurers(env DevspaceEnv) []ChainConfigurer {
	chainConfigurers := make([]ChainConfigurer, 0, env.BesuChainsCount+env.GethChainsCount+env.SolanaChainsCount)

	// Add Besu chains
	besuChainsCount := env.BesuChainsCount
	besuChains := BuildEVMNetworkConfigs(besuChainsCount)
	for i, chain := range besuChains {
		// Check if the NetworkId is negative before converting to uint64
		nId, err := strconv.Atoi(chain.NetworkId)
		if err != nil {
			panic(fmt.Sprintf("Error: NetworkId %d for chain is not an integer, cannot convert to int", chain.NetworkId))
		}
		if nId < 0 {
			panic(fmt.Sprintf("Error: NetworkId %d for besu chain is negative, cannot convert to uint64", chain.NetworkId))
		}

		var c ChainConfigurer
		switch i {
		case 0:
			c = NewChainConfigurer(env, chain.NetworkId, EVM, "besu", "alpha", 0)
		case 1:
			c = NewChainConfigurer(env, chain.NetworkId, EVM, "besu", "beta", 0)
		default:
			c = NewChainConfigurer(env, chain.NetworkId, EVM, "besu", fmt.Sprintf("besu-%d", nId), 0)
		}

		chainConfigurers = append(chainConfigurers, c)
	}

	// Add Geth chains
	gethChainsCount := env.GethChainsCount
	gethChains := BuildEVMNetworkConfigs(gethChainsCount)
	for _, chain := range gethChains {
		// Check if the NetworkId is negative before converting to uint64
		nId, err := strconv.Atoi(chain.NetworkId)
		if err != nil {
			panic(fmt.Sprintf("Error: NetworkId %d for chain is not an integer, cannot convert to int", chain.NetworkId))
		}
		if nId < 0 {
			panic(fmt.Sprintf("Error: NetworkId %d for geth chain is negative, cannot convert to uint64", chain.NetworkId))
		}
		networkId := nId

		c := NewChainConfigurer(env, chain.NetworkId, EVM, "geth", fmt.Sprintf("geth-%d", networkId), 0)
		chainConfigurers = append(chainConfigurers, c)
	}

	// Add Solana chains
	if env.SolanaChainsCount > 3 {
		panic("Up to 3 sol chains are supported")
	}
	selectors := []string{"22222222222222222222222222222222222222222222", "33333333333333333333333333333333333333333333", "44444444444444444444444444444444444444444444"}
	for i := 0; i <= env.SolanaChainsCount-1; i++ {
		c := NewChainConfigurer(env, selectors[i], SOLANA, "", fmt.Sprintf("solana-local-%d", SolStartingPodId+i), SolStartingPodId+i)
		chainConfigurers = append(chainConfigurers, c)
	}

	return chainConfigurers
}

// Returns gap v2 hostname for the given env
// hostnameSuffix should be the same as exposed serviceName
func gapV2HostName(env DevspaceEnv, hostnameSuffix string) string {
	//${crib-namespace}-${hostnameSuffix}.${dnsZone}.${baseDomain}
	// example gap-crib-ci-123123-job-distributor.public main.stage.cldev.sh
	return fmt.Sprintf("gap-%s-%s.public.%s:443", env.Namespace, hostnameSuffix, env.IngressBaseDomain)
}

func GetTransmittedChainConfigs(env DevspaceEnv) []crib.ChainConfig {
	chainConfigurers := getChainConfigurers(env)

	chainConfigs := make([]crib.ChainConfig, 0)
	for _, chainConfigurer := range chainConfigurers {
		config := chainConfigurer.GetTransmittedChainConfigs()

		chainConfigs = append(chainConfigs, config)
	}

	return chainConfigs
}

func NewCLNodeConfigurer(env DevspaceEnv) ChainlinkNodeConfigurer {
	return ChainlinkNodeConfigurer{
		env: env,
	}
}

func (c ChainlinkNodeConfigurer) GetNodeInfos() []devenv.NodeInfo {
	var nodes []devenv.NodeInfo
	nodes = append(nodes, c.GetBootNodeInfos()...)
	nodes = append(nodes, c.GetWorkerNodeInfos()...)

	return nodes
}

func (c ChainlinkNodeConfigurer) GetWorkerNodeInfos() []devenv.NodeInfo {
	var nodes []devenv.NodeInfo
	for i := 0; i < c.env.DonNodeCount; i++ {
		bootNode := c.getNodeInfo(fmt.Sprintf("ccip-%d", i), false)
		nodes = append(nodes, bootNode)
	}

	return nodes
}

func (c ChainlinkNodeConfigurer) GetBootNodeInfos() []devenv.NodeInfo {
	var nodes []devenv.NodeInfo
	for i := 0; i < c.env.DonBootNodeCount; i++ {
		bootNode := c.getNodeInfo(fmt.Sprintf("ccip-bt-%d", i), true)
		nodes = append(nodes, bootNode)
	}
	return nodes
}

func (c ChainlinkNodeConfigurer) getNodeInfo(nodeName string, isBootstrap bool) devenv.NodeInfo {
	protocol := "https"
	port := 443
	// Workaround for https://smartcontract-it.atlassian.net/browse/CRIB-547
	if c.env.Provider == "kind" {
		protocol = "http"
		port = 80
	}

	return devenv.NodeInfo{
		CLConfig: nodeclient.ChainlinkConfig{
			URL:        fmt.Sprintf("%s://%s-%s.%s:%d", protocol, c.env.Namespace, nodeName, c.env.IngressBaseDomain, port),
			Email:      "admin@chain.link",
			Password:   "hWDmgcub2gUhyrG6cxriqt7T",
			InternalIP: nodeName,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		P2PPort:     "5001",
		IsBootstrap: isBootstrap,
		Name:        nodeName,
		AdminAddr:   "",
		MultiAddr:   "",
	}
}

func NewChainConfigurer(env DevspaceEnv, chainID string, chainType ChainType, chainVariant, name string, podId int) ChainConfigurer {
	// These are generally known private keys used for testing
	testKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	var gasPrice *big.Int

	if chainVariant == "besu" {
		// the same key as used in the FWOG env alpha chain
		testKey = "8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63"
		gasPrice = big.NewInt(2000000)
	}
	// Generate sol key for the deployment
	solTestKey, err := solana.PrivateKeyFromBase58("57qbvFjTChfNwQxqkFZwjHp7xYoPZa7f9ow6GA59msfCH1g6onSjKUTrrLp4w1nAwbwQuit8YgJJ2AwT9BSwownC")
	if err != nil {
		panic("Could not create localnet private key")
	}

	return ChainConfigurer{
		env:            env,
		chainID:        chainID,
		chainName:      name,
		chainVariant:   chainVariant,
		deployerKey:    testKey,
		solDeployerKey: solTestKey,
		chainType:      chainType,
		GasPrice:       gasPrice,
		podID:          podId,
	}
}

func NewChainConfigurerFromChainConfig(env DevspaceEnv, chainConfig devenv.ChainConfig, chainVariant string, podId int) ChainConfigurer {
	return NewChainConfigurer(env, chainConfig.ChainID, ChainTypeFromString(chainConfig.ChainType), chainVariant, chainConfig.ChainName, podId)
}

func (c ChainConfigurer) GetDevenvChainConfig() (*devenv.ChainConfig, error) {
	wsExternalRPC := c.ExternalWSRPC()
	if wsExternalRPC == nil {
		return nil, fmt.Errorf("wsRPC external url not available")
	}
	wsInternalRPC := c.InternalWSRPC()
	if wsInternalRPC == nil {
		return nil, fmt.Errorf("wsRPC internal url not available")
	}
	httpExternalRPC := c.ExternalHTTPRPC()
	if httpExternalRPC == nil {
		return nil, fmt.Errorf("httpRPC external url not available")
	}
	httpInternalRPC := c.InternalWSRPC()
	if httpInternalRPC == nil {
		return nil, fmt.Errorf("httpRPC internal url not available")
	}

	chainConfig := &devenv.ChainConfig{
		ChainID:   c.chainID,
		ChainName: c.chainName,
		ChainType: c.chainType.String(),
		WSRPCs: []devenv.CribRPCs{{
			External: *wsExternalRPC,
			Internal: *wsInternalRPC,
		}},
		HTTPRPCs: []devenv.CribRPCs{{
			External: *httpExternalRPC,
			Internal: *httpInternalRPC,
		}},
		SolDeployerKey: c.solDeployerKey,
	}

	err := chainConfig.SetDeployerKey(&c.deployerKey)
	chainConfig.DeployerKey.GasPrice = c.GasPrice

	if err != nil {
		return nil, fmt.Errorf("unable to set deployer key, err: %s", err)
	}

	return chainConfig, nil
}

func (c ChainConfigurer) GetTransmittedChainConfigs() crib.ChainConfig {
	chainConfig := crib.ChainConfig{
		ChainID:   c.chainID,
		ChainName: c.GetChainName(),
		ChainType: c.chainType.String(),
		WSRPCs: []crib.RPC{
			{
				Internal: c.InternalWSRPC(),
				External: c.ExternalWSRPC(),
			},
		},
		HTTPRPCs: []crib.RPC{
			{
				External: c.ExternalHTTPRPC(),
				Internal: c.InternalHTTPRPC(),
			},
		},
	}

	return chainConfig
}

func ChainSelector(chainID uint64) uint64 {
	chain, ok := chainsel.ChainByEvmChainID(chainID)

	if !ok {
		panic("chain not found")
	}

	homeChainSelector := chain.Selector
	return homeChainSelector
}

func (c ChainConfigurer) ExternalWSRPC() *string {
	var u string
	if c.env.CIEnv {
		hostName := gapV2HostName(c.env, fmt.Sprintf("%s-%s-ws", strings.ToLower(c.chainName), c.chainID))
		u = fmt.Sprintf("wss://%s", hostName)
	} else if c.chainType == EVM && c.chainVariant == "besu" {
		u = fmt.Sprintf("wss://chain-%s-rpc.%s/ws/", c.chainName, c.env.IngressBaseDomain)
	} else if c.chainType == SOLANA {
		u = fmt.Sprintf("wss://%s-%s-%d-ws.%s", c.env.Namespace, c.chainTypeHostNamePart(), c.podID, c.env.IngressBaseDomain)
	} else {
		u = fmt.Sprintf("wss://%s-%s-%s-ws.%s", c.env.Namespace, c.chainTypeHostNamePart(), c.chainID, c.env.IngressBaseDomain)
	}
	return &u
}

func (c ChainConfigurer) ExternalHTTPRPC() *string {
	var u string
	if c.env.CIEnv {
		hostName := gapV2HostName(c.env, fmt.Sprintf("%s-%s", strings.ToLower(c.chainVariant), c.chainID))
		u = fmt.Sprintf("https://%s", hostName)
	} else if c.chainType == EVM && c.chainVariant == "besu" {
		u = fmt.Sprintf("https://chain-%s-rpc.%s", c.chainName, c.env.IngressBaseDomain)
	} else if c.chainType == SOLANA {
		u = fmt.Sprintf("https://%s-%s-%d-http.%s:443", c.env.Namespace, c.chainTypeHostNamePart(), c.podID, c.env.IngressBaseDomain)
	} else {
		u = fmt.Sprintf("https://%s-%s-%s-http.%s:443", c.env.Namespace, c.chainTypeHostNamePart(), c.chainID, c.env.IngressBaseDomain)
	}
	return &u
}

func (c ChainConfigurer) chainTypeHostNamePart() string {
	var chainType string
	if c.chainType == EVM && c.chainVariant == "besu" {
		chainType = "besu"
	} else if c.chainType == EVM && c.chainVariant == "geth" {
		chainType = "geth"
	} else if c.chainType == SOLANA {
		chainType = "solana"
	}
	return chainType
}

func (c ChainConfigurer) InternalWSRPC() *string {
	var u string

	switch {
	case c.chainType == EVM && c.chainVariant == "besu":
		u = fmt.Sprintf("ws://%s-node-rpc-1.chain-%s.svc.cluster.local:8546", strings.ToLower(c.chainVariant), c.chainName)
	case c.chainType == EVM && c.chainVariant == "geth":
		u = fmt.Sprintf("ws://%s-%s-ws:8546", strings.ToLower(c.chainVariant), c.chainID)
	case c.chainType == SOLANA:
		u = fmt.Sprintf("ws://%s-%d:8545", strings.ToLower(c.chainType.String()), c.podID)
	default:
		return nil
	}
	return &u
}

func (c ChainConfigurer) InternalHTTPRPC() *string {
	var u string
	switch {
	case c.chainType == EVM && c.chainVariant == "besu":
		u = fmt.Sprintf("http://%s-node-rpc-1.chain-%s.svc.cluster.local:8545", c.chainVariant, c.chainName)
	case c.chainType == EVM && c.chainVariant == "geth":
		u = fmt.Sprintf("http://%s-%s:8544", strings.ToLower(c.chainVariant), c.chainID)
	case c.chainType == SOLANA:
		u = fmt.Sprintf("http://%s-%d:8544", strings.ToLower(c.chainType.String()), c.podID)
	default:
		return nil
	}
	return &u
}

func (c ChainConfigurer) GetChainName() string {
	return fmt.Sprintf("%s-simulated-%s", strings.ToLower(c.chainType.String()), c.chainID)
}

func BuildEVMNetworkConfigs(chainsCount int) []EVMChain {
	// If chainsCount is 0, return an empty slice
	if chainsCount == 0 {
		return []EVMChain{}
	}

	// Initialize the chains slice
	chains := []EVMChain{
		{NetworkId: "1337", FinalityDepth: defaultFinalityDepth},
	}

	// Add the second chain if chainsCount > 1
	if chainsCount > 1 {
		chains = append(chains, EVMChain{NetworkId: "2337", FinalityDepth: defaultFinalityDepth})
	}

	// Add subsequent chains starting from 90000000 if chainsCount > 2
	for i := 2; i < chainsCount; i++ {
		networkId := int64(90000000 + i - 1)
		chains = append(chains, EVMChain{NetworkId: strconv.FormatUint(uint64(networkId), 10), FinalityDepth: defaultFinalityDepth})
	}

	return chains
}
