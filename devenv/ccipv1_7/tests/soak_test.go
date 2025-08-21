package ccipv2_tests

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"

	ccipv17 "github.com/smartcontractkit/devenv/ccipv17"

	f "github.com/smartcontractkit/chainlink-testing-framework/framework"
)

var L = ccipv17.Plog

func TestSoak(t *testing.T) {
	in, err := ccipv17.LoadOutput[ccipv17.Cfg]("../env-out.toml")
	require.NoError(t, err)
	c, _, _, err := ccipv17.ETHClient(in.Blockchains[0].Out.Nodes[0].ExternalWSUrl, in.CCIPv17.GasSettings)
	require.NoError(t, err)
	clNodes, err := clclient.New(in.NodeSets[0].Out.CLNodes)
	require.NoError(t, err)
	_ = clNodes
	_ = c

	// connect your contracts with CLD here and assert CCIPv17 lanes are working

	// assert resources, memory, CPU consumption
	start := time.Now()
	end := time.Now()
	checkResourceConsumption(t, in, start, end)
}

func checkResourceConsumption(t *testing.T, in *ccipv17.Cfg, start, end time.Time) {
	pc := f.NewPrometheusQueryClient(f.LocalPrometheusBaseURL)
	// example Prometheus query, assert resources, CPU, Memory, etc
	// no more than 10% CPU at the end of the test
	maxCPU := 10.0
	cpuResp, err := pc.Query("sum(rate(container_cpu_usage_seconds_total{name=~\".*don.*\"}[5m])) by (name) *100", end)
	require.NoError(t, err)
	cpu := f.ToLabelsMap(cpuResp)
	for i := 0; i < in.NodeSets[0].Nodes; i++ {
		nodeLabel := fmt.Sprintf("name:don-node%d", i)
		nodeCpu, err := strconv.ParseFloat(cpu[nodeLabel][0].(string), 64)
		L.Info().Int("Node", i).Float64("CPU", nodeCpu).Msg("CPU usage percentage")
		require.NoError(t, err)
		require.LessOrEqual(t, nodeCpu, maxCPU)
	}
	// no more than 200mb for this test
	maxMem := int(200e6) // 200mb
	memoryResp, err := pc.Query("sum(container_memory_rss{name=~\".*don.*\"}) by (name)", end)
	require.NoError(t, err)
	mem := f.ToLabelsMap(memoryResp)
	for i := 0; i < in.NodeSets[0].Nodes; i++ {
		nodeLabel := fmt.Sprintf("name:don-node%d", i)
		nodeMem, err := strconv.Atoi(mem[nodeLabel][0].(string))
		L.Info().Int("Node", i).Int("Memory", nodeMem).Msg("Total memory")
		require.NoError(t, err)
		require.LessOrEqual(t, nodeMem, maxMem)
	}
}
