package config

import (
	"crypto/tls"
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	"github.com/smartcontractkit/chainlink/deployment/environment/nodeclient"
	"google.golang.org/grpc/credentials"
)

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

func getChainConfigurers(env DevspaceEnv) []ChainConfigurer {
	chainConfigurers := []ChainConfigurer{
		NewChainConfigurer(env, uint64(1337), "alpha"),
		NewChainConfigurer(env, uint64(2337), "beta"),
	}

	if env.ChainsCount > 2 {
		//nolint:gosec
		for i := 1; i <= env.ChainsCount-2; i++ {
			const baseChainID uint64 = 90000000
			chainID := baseChainID + uint64(i)
			c := NewChainConfigurer(env, chainID, fmt.Sprintf("nchain-%d", chainID))

			chainConfigurers = append(chainConfigurers, c)
		}
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

type ChainConfigurer struct {
	chainID     uint64
	deployerKey string
	env         DevspaceEnv
	chainName   string
}

func NewChainConfigurer(env DevspaceEnv, chainID uint64, name string) ChainConfigurer {
	// These are generally known private keys used for testing
	testKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	return ChainConfigurer{
		env:         env,
		chainID:     chainID,
		chainName:   name,
		deployerKey: testKey,
	}
}

func (c ChainConfigurer) GetDevenvChainConfig() (*devenv.ChainConfig, error) {
	wsExternalRPC := c.externalWSRPC()
	if wsExternalRPC == nil {
		return nil, fmt.Errorf("wsRPC external url not available")
	}
	wsInternalRPC := c.internalWSRPC()
	if wsInternalRPC == nil {
		return nil, fmt.Errorf("wsRPC internal url not available")
	}
	httpExternalRPC := c.externalHTTPRPC()
	if httpExternalRPC == nil {
		return nil, fmt.Errorf("httpRPC external url not available")
	}
	httpInternalRPC := c.internalWSRPC()
	if httpInternalRPC == nil {
		return nil, fmt.Errorf("httpRPC internal url not available")
	}

	chainConfig := &devenv.ChainConfig{
		ChainID:   c.chainID,
		ChainName: c.chainName,
		ChainType: "EVM",
		WSRPCs: []devenv.CribRPCs{{
			External: *wsExternalRPC,
			Internal: *wsInternalRPC,
		}},
		HTTPRPCs: []devenv.CribRPCs{{
			External: *httpExternalRPC,
			Internal: *httpInternalRPC,
		}},
	}

	err := chainConfig.SetDeployerKey(&c.deployerKey)
	if err != nil {
		return nil, fmt.Errorf("unable to set deployer key, err: %s", err)
	}

	return chainConfig, nil
}

func (c ChainConfigurer) GetTransmittedChainConfigs() crib.ChainConfig {
	chainConfig := crib.ChainConfig{
		ChainID:   c.chainID,
		ChainName: "alpha",
		ChainType: "EVM",
		WSRPCs: []crib.RPC{
			{
				Internal: c.internalWSRPC(),
				External: c.externalWSRPC(),
			},
		},
		HTTPRPCs: []crib.RPC{
			{
				External: c.externalHTTPRPC(),
				Internal: c.internalHTTPRPC(),
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

func (c ChainConfigurer) externalWSRPC() *string {
	if c.env.CIEnv {
		hostName := gapV2HostName(c.env, fmt.Sprintf("geth-%d-ws", c.chainID))
		u := fmt.Sprintf("wss://%s", hostName)
		return &u
	}
	u := fmt.Sprintf("wss://%s-geth-%d-ws.%s", c.env.Namespace, c.chainID, c.env.IngressBaseDomain)
	return &u
}

func (c ChainConfigurer) externalHTTPRPC() *string {
	if c.env.CIEnv {
		hostName := gapV2HostName(c.env, fmt.Sprintf("geth-%d", c.chainID))
		u := fmt.Sprintf("https://%s", hostName)
		return &u
	}
	u := fmt.Sprintf("https://%s-geth-%d-ws.%s:443", c.env.Namespace, c.chainID, c.env.IngressBaseDomain)
	return &u
}

func (c ChainConfigurer) internalWSRPC() *string {
	u := fmt.Sprintf("ws://%s-geth-%d-ws:8546", c.env.Namespace, c.chainID)
	return &u
}

func (c ChainConfigurer) internalHTTPRPC() *string {
	u := fmt.Sprintf("http://%s-geth-%d:8544", c.env.Namespace, c.chainID)
	return &u
}
