package e2e_test

import (
	"testing"

	ccv "github.com/smartcontractkit/chainlink-ccip/v2/devenv"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
)

func TestE2E(t *testing.T) {
	in, err := ccv.LoadOutput[ccv.Cfg]("../../env-out.toml")
	require.NoError(t, err)
	c, _, _, err := ccv.ETHClient(in.Blockchains[0].Out.Nodes[0].ExternalWSUrl, in.CCV.GasSettings)
	require.NoError(t, err)
	clNodes, err := clclient.New(in.NodeSets[0].Out.CLNodes)
	require.NoError(t, err)
	_ = clNodes
	_ = c
	// connect your contracts with CLD here and assert CCV lanes are working
}
