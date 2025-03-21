package scripts

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	cretypes "github.com/smartcontractkit/chainlink/system-tests/lib/cre/types"
	"github.com/smartcontractkit/chainlink/system-tests/lib/crypto"
	"github.com/smartcontractkit/chainlink/system-tests/lib/types"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
)

func GenerateInitialNodeOverrides(env config.DevspaceEnv) {
	envState := model.NewEnvState(logger, env)

	if envState.NodeOverridesExist() {
		logger.Infow("Node config and secret overrides exist, skipping generation")
		return
	}

	keys, err := crypto.GenerateP2PKeys("", env.DonBootNodeCount+env.DonNodeCount)
	if err != nil {
		logger.Fatal("Failed to generate p2p keys: ", err)
		os.Exit(2)
	}

	GenerateSecretsOverrides(envState, env.DonNodeCount+env.DonBootNodeCount, keys)
	GenerateConfigOverrides(env, envState, keys)
}

func NodeConfigOverride(donBootstrapNodePeerID, donBootstrapNodeHost string, donBootstrapNodePort string) string {
	return fmt.Sprintf(`[P2P.V2]
Enabled = true
ListenAddresses = ['0.0.0.0:%s']
DefaultBootstrappers = ['%s@%s:%s']
`,
		donBootstrapNodePort,
		donBootstrapNodePeerID,
		donBootstrapNodeHost,
		donBootstrapNodePort,
	)
}

func GenerateConfigOverrides(env config.DevspaceEnv, envState model.CCIPEnvState, keys *types.P2PKeys) {
	bootNodeInfos := config.NewCLNodeConfigurer(env).GetBootNodeInfos()
	bootstrapHost := bootNodeInfos[0].Name

	peerIDWithPrefix := keys.PeerIDs[0]
	bootstrapPeerID, _ := strings.CutPrefix(peerIDWithPrefix, "p2p_")

	for i := 0; i < env.DonBootNodeCount; i++ {
		nodeInfo := bootNodeInfos[i]
		host := "localhost"
		fileNamePattern := model.BootNodeInitialConfigOverridesFileNamePattern

		override := NodeConfigOverride(bootstrapPeerID, host, nodeInfo.P2PPort)
		envState.SaveStringToFile(override, fmt.Sprintf(fileNamePattern, i), model.CLNodeConfigInputs)
	}

	workerNodeInfos := config.NewCLNodeConfigurer(env).GetWorkerNodeInfos()
	for i := 0; i < env.DonNodeCount; i++ {
		nodeInfo := workerNodeInfos[i]
		host := bootstrapHost
		fileNamePattern := model.WorkerNodeInitialConfigOverridesFileNamePattern
		override := NodeConfigOverride(bootstrapPeerID, host, nodeInfo.P2PPort)
		envState.SaveStringToFile(override, fmt.Sprintf(fileNamePattern, i), model.CLNodeConfigInputs)
	}

	saveNodeDetails(envState, bootstrapPeerID, bootstrapHost, bootNodeInfos[0].P2PPort)
}

func saveNodeDetails(envState model.CCIPEnvState, bootstrapPeerID string, host string, port string) {
	nodesDetails := crib.NodesDetails{
		NodeIDs: make([]string, 0),
		BootstrapNode: crib.BootstrapNode{
			P2PID:        bootstrapPeerID,
			InternalHost: host,
			Port:         port,
		},
	}
	envState.SaveNodeDetails(nodesDetails)
}

func GenerateSecretsOverrides(envState model.CCIPEnvState, nodeCount int, keys *types.P2PKeys) {
	input := &cretypes.GenerateSecretsInput{
		P2PKeys: keys,
		DonMetadata: &cretypes.DonMetadata{
			NodesMetadata: make([]*cretypes.NodeMetadata, nodeCount),
		},
	}
	secretOverrides, err := GenerateSecrets(input)
	logger.Debugw("Generated secret overrides for nodes")
	if err != nil {
		logger.Fatal("Failed to generate secrets override: ", err)
		os.Exit(2)
	}

	for index, override := range secretOverrides {
		if index == 0 {
			envState.SaveStringToFile(override, fmt.Sprintf(model.BootNodeInitialSecretsOverridesFileNamePattern, index), model.CLNodeConfigInputs)
		} else {
			envState.SaveStringToFile(override, fmt.Sprintf(model.NodeInitialSecretsOverridesFileNamePattern, index-1), model.CLNodeConfigInputs)
		}
	}
}

type NodeSecret struct {
	P2PKey types.NodeP2PKey `toml:"P2PKey"`
}

func GenerateSecrets(input *cretypes.GenerateSecretsInput) (cretypes.NodeIndexToSecretsOverride, error) {
	if input == nil {
		return nil, errors.New("input is nil")
	}
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "input validation failed")
	}

	overrides := make(cretypes.NodeIndexToSecretsOverride)

	for i := range input.DonMetadata.NodesMetadata {
		nodeSecret := NodeSecret{}

		if input.P2PKeys != nil {
			nodeSecret.P2PKey = types.NodeP2PKey{
				JSON:     string(input.P2PKeys.EncryptedJSONs[i]),
				Password: input.P2PKeys.Password,
			}
		}

		nodeSecretString, err := toml.Marshal(nodeSecret)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal node secrets")
		}

		overrides[i] = string(nodeSecretString)
	}

	return overrides, nil
}
