package e2e

import (
	"fmt"
	"testing"

	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
	"github.com/smartcontractkit/chainlink-ccip/devenv/tests"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/stretchr/testify/require"
)

func TestE2ECurse(t *testing.T) {
	in, err := ccip.LoadOutput[ccip.Cfg]("../../env-out.toml")
	require.NoError(t, err)
	if in.ForkedEnvConfig != nil {
		t.Skip("Skipping E2E tests on forked environments, not supported yet")
	}

	selectors, e, err := ccip.NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := framework.SaveContainerLogs(fmt.Sprintf("%s-%s", framework.DefaultCTFLogsDir, t.Name()))
		require.NoError(t, err)
	})

	tests.RunCurseTests(t, e, selectors)
}
