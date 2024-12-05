package scripts

import (
	"testing"

	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDeployCCIPAndAddLanes(t *testing.T) {
	t.Skip()
	t.Parallel()
	tmpDir := "/tmp/ccip-v2"

	env := getTestEnv()

	DeployCCIPAndAddLanes(nil, env, tmpDir)

	t.Fail()
}

func getTestEnv() config.DevspaceEnv {
	env := config.DevspaceEnv{
		Namespace:         "crib-local",
		Provider:          "kind",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		IngressBaseDomain: "main.stage.cldev.sh",
	}
	return env
}

func TestJDConnection(t *testing.T) {
	t.Skip()
	t.Parallel()
	nodeInfos := config.NewCLNodeConfigurer(getTestEnv()).GetNodeInfos()
	jdConfig := devenv.JDConfig{
		// this is for our script connecting to jd
		GRPC: "crib-local-job-distributor-grpc.crib.local:443",
		// this for the node to connect to jd
		// todo: we need different URI here
		// Does it need to be internal or external?
		// ws:// is not needed here
		WSRPC:    "job-distributor-noderpc-lb:80",
		Creds:    insecure.NewCredentials(),
		NodeInfo: nodeInfos,
	}

	connection, err := devenv.NewJDConnection(jdConfig)

	assert.NoError(t, err)
	assert.NotNil(t, connection)
}
