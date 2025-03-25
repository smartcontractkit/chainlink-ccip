package rmn

import (
	"fmt"
	"strconv"

	"github.com/pelletier/go-toml/v2"
	"github.com/smartcontractkit/chainlink/deployment/ccip/changeset"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
)

type Config struct {
	Networking   Networking    `toml:"networking"`
	HomeChain    HomeChain     `toml:"home_chain"`
	ChainParams  []ChainParam  `toml:"chain_params"`
	RemoteChains []RemoteChain `toml:"remote_chains"`
}

type Networking struct {
	Bootstrappers []string `toml:"bootstrappers"`
}

type HomeChain struct {
	Name                 string `toml:"name"`
	CapabilitiesRegistry string `toml:"capabilities_registry"`
	CCIPHome             string `toml:"ccip_home"`
	RMNHome              string `toml:"rmn_home"`
}

type ChainParam struct {
	Name      string          `toml:"name"`
	Stability StabilityConfig `toml:"stability"`
}

type StabilityConfig struct {
	// One Of: ["ConfirmationDepth", "FinalityTag" ]
	Type              string `toml:"type"`
	SoftConfirmations int    `toml:"soft_confirmations"`
	HardConfirmations int    `toml:"hard_confirmations"`
}

type RemoteChain struct {
	Name                   string `toml:"name"`
	OnRampStartBlockNumber int    `toml:"on_ramp_start_block_number"`
	OnRamp                 string `toml:"on_ramp"`
	OffRamp                string `toml:"off_ramp"`
	RMNRemote              string `toml:"rmn_remote"`
}

func GenerateSharedConfigTOML(ccipOnChainState changeset.CCIPOnChainState, envConfig *devenv.EnvironmentConfig, envState model.CCIPEnvState, clBootstrapNode crib.BootstrapNode) error {
	var homeChain HomeChain
	remoteChains := make([]RemoteChain, 0)
	for chainSelector, chain := range ccipOnChainState.Chains {
		chainConfig := config.GetChainConfigBySelector(envConfig.Chains, chainSelector)
		if chainConfig == nil {
			return fmt.Errorf("chain '%d' does not exist in chainConfig", chainSelector)
		}
		chainID := chainConfig.ChainID

		if chain.CCIPHome != nil {
			if chain.CapabilityRegistry == nil {
				return fmt.Errorf("chain capability registry is required for home chain")
			}
			if chain.CCIPHome == nil {
				return fmt.Errorf("chain ccip home is required for home chain")
			}
			if chain.RMNHome == nil {
				return fmt.Errorf("chain rmn home is required for home chain")
			}

			homeChain = HomeChain{
				Name:                 chainName(chainConfig),
				CapabilitiesRegistry: chain.CapabilityRegistry.Address().String(),
				CCIPHome:             chain.CCIPHome.Address().String(),
				RMNHome:              chain.RMNHome.Address().String(),
			}
		}

		if chain.OnRamp == nil {
			return fmt.Errorf("chain '%d' does not have a OnRamp set", chainID)
		}
		if chain.OffRamp == nil {
			return fmt.Errorf("chain '%d' does not have an OffRamp set", chainID)
		}
		if chain.RMNRemote == nil {
			return fmt.Errorf("chain '%d' does not have a RMNRemote set", chainID)
		}
		remoteChains = append(remoteChains, RemoteChain{
			Name:                   chainName(chainConfig),
			OnRampStartBlockNumber: 0,
			OnRamp:                 chain.OnRamp.Address().String(),
			OffRamp:                chain.OffRamp.Address().String(),
			RMNRemote:              chain.RMNRemote.Address().String(),
		})
	}

	chainParams := make([]ChainParam, 0)

	for _, chain := range envConfig.Chains {
		chainParams = append(chainParams, ChainParam{
			Name: chainName(&chain),
			// todo: replace with Stability: StabilityConfig{Type: "FinalityTag"},
			// once we have newer geth eth2 blockchain working
			Stability: StabilityConfig{
				Type:              "ConfirmationDepth",
				SoftConfirmations: 0,
				HardConfirmations: 100,
			},
		})
	}

	sharedConfigToml := Config{
		Networking: Networking{
			Bootstrappers: []string{fmt.Sprintf("%s@%s:%s", clBootstrapNode.P2PID, clBootstrapNode.InternalHost, clBootstrapNode.Port)},
		},
		HomeChain:    homeChain,
		ChainParams:  chainParams,
		RemoteChains: remoteChains,
	}

	tomlData, err := toml.Marshal(sharedConfigToml)
	if err != nil {
		return fmt.Errorf("failed to marshal shared config: %w", err)
	}

	envState.SaveRMNSharedToml(tomlData)
	return nil
}

type LocalTomlConfig struct {
	Networking LocalTomlNetworking `toml:"networking"`
	Chains     []LocalTomlChain    `toml:"chains"`
}

type LocalTomlNetworking struct {
	RageProxy string `toml:"rageproxy"`
}

type LocalTomlChain struct {
	Name string `toml:"name"`
	RPC  string `toml:"rpc"`
}

func GenerateLocalToml(env config.DevspaceEnv, chains []devenv.ChainConfig, envState model.CCIPEnvState) error {
	localTomlChains := make([]LocalTomlChain, 0)

	for _, chain := range chains {
		configurer := config.NewChainConfigurerFromChainConfig(env, chain, "geth")
		httpRPC := configurer.InternalHTTPRPC()
		localTomlChains = append(localTomlChains, LocalTomlChain{
			Name: chainName(&chain),
			RPC:  *httpRPC,
		})
	}

	localToml := LocalTomlConfig{
		Networking: LocalTomlNetworking{
			RageProxy: "127.0.0.1:23456",
		},
		Chains: localTomlChains,
	}
	tomlData, err := toml.Marshal(localToml)
	if err != nil {
		return fmt.Errorf("failed to marshal local toml config: %w", err)
	}
	envState.SaveRMNLocalToml(tomlData)

	return nil
}

// RMN requires some custom chainName logic, chain names in RMN are different than in chain-selectors repo
func chainName(chainConfig *devenv.ChainConfig) string {
	chainIDString := strconv.FormatUint(chainConfig.ChainID, 10)

	switch chainIDString {
	case "1337":
		return "DevnetAlpha"
	case "2337":
		return "DevnetBeta"
	}

	return fmt.Sprintf("test-%s", chainIDString)
}
