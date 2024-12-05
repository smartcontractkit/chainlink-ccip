package config

import (
	"crypto/tls"
	"fmt"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	"github.com/smartcontractkit/chainlink/deployment/environment/nodeclient"
	"google.golang.org/grpc/credentials"
)

func NewEnvConfig(env DevspaceEnv) devenv.EnvironmentConfig {
	alphaConfigurer := NewChainConfigurer(env.Namespace, env.IngressBaseDomain, uint64(1337))
	betaConfigurer := NewChainConfigurer(env.Namespace, env.IngressBaseDomain, uint64(2337))
	chainsConfig := []devenv.ChainConfig{
		alphaConfigurer.configureChain(),
		betaConfigurer.configureChain(),
	}
	nodeInfos := NewCLNodeConfigurer(env).GetNodeInfos()
	jdConfig := devenv.JDConfig{
		GRPC:  fmt.Sprintf("%s-job-distributor-grpc.%s:443", env.Namespace, env.IngressBaseDomain),
		WSRPC: "job-distributor-noderpc-lb:80",
		Creds: credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		}),
		NodeInfo: nodeInfos,
	}

	return devenv.EnvironmentConfig{
		Chains:   chainsConfig,
		JDConfig: jdConfig,
	}
}

type ChainlinkNodeConfigurer struct {
	env DevspaceEnv
}

func NewCLNodeConfigurer(env DevspaceEnv) ChainlinkNodeConfigurer {
	return ChainlinkNodeConfigurer{
		env: env,
	}
}

func (c ChainlinkNodeConfigurer) GetNodeInfos() []devenv.NodeInfo {
	var nodes []devenv.NodeInfo
	for i := 0; i < c.env.DonBootNodeCount; i++ {
		bootNode := c.getNodeInfo(fmt.Sprintf("ccip-bt-%d", i), true)
		nodes = append(nodes, bootNode)
	}

	for i := 0; i < c.env.DonNodeCount; i++ {
		bootNode := c.getNodeInfo(fmt.Sprintf("ccip-%d", i), false)
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
			InternalIP: fmt.Sprintf("%s-%s", c.env.Namespace, nodeName),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		P2PPort:     "6690",
		IsBootstrap: isBootstrap,
		Name:        nodeName,
		AdminAddr:   "",
		MultiAddr:   "",
	}
}

type ChainConfigurer struct {
	cribNamespace string
	ingressDomain string
	chainID       uint64
	deployerKey   string
}

func NewChainConfigurer(cribNamespace string, ingressDomain string, chainID uint64) ChainConfigurer {
	// These are generally known private keys used for testing
	testKeys := map[string]string{
		"1337": "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"2337": "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	}

	chainIDStr := strconv.FormatUint(chainID, 10)

	return ChainConfigurer{
		cribNamespace: cribNamespace,
		ingressDomain: ingressDomain,
		chainID:       chainID,
		deployerKey:   testKeys[chainIDStr],
	}
}

func (c ChainConfigurer) configureChain() devenv.ChainConfig {
	chainConfig := devenv.ChainConfig{
		ChainID:   c.chainID,
		ChainName: "alpha",
		ChainType: "EVM",
		WSRPCs:    []string{c.chainWSRPC()},
		HTTPRPCs:  []string{c.chainHTTPRPC()},
	}

	err := chainConfig.SetDeployerKey(&c.deployerKey)
	if err != nil {
		panic(err)
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

func (c ChainConfigurer) chainWSRPC() string {
	return fmt.Sprintf("wss://%s-geth-%d-ws.%s:443", c.cribNamespace, c.chainID, c.ingressDomain)
}

func (c ChainConfigurer) chainHTTPRPC() string {
	return fmt.Sprintf("https://%s-geth-%d-ws.%s:443", c.cribNamespace, c.chainID, c.ingressDomain)
}
