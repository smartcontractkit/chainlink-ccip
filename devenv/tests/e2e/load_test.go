package e2e

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/chaos"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/rpc"
	"github.com/smartcontractkit/chainlink-testing-framework/wasp"

	chainsel "github.com/smartcontractkit/chain-selectors"
	f "github.com/smartcontractkit/chainlink-testing-framework/framework"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
)

type ChaosTestCase struct {
	run      func() error
	validate func() error
	name     string
}

type GasTestCase struct {
	increase         *big.Int
	gasFunc          func(t *testing.T, r *rpc.RPCClient, blockPace time.Duration)
	validate         func() error
	name             string
	chainURL         string
	waitBetweenTests time.Duration
}

type CCIPTxGun struct {
	e       *deployment.Environment
	srcImpl ccip.CCIP16ProductConfiguration
	dstImpl ccip.CCIP16ProductConfiguration
}

func NewCCIPTxGun(e *deployment.Environment, srcImpl, dstImpl ccip.CCIP16ProductConfiguration) *CCIPTxGun {
	return &CCIPTxGun{
		e:       e,
		srcImpl: srcImpl,
		dstImpl: dstImpl,
	}
}

// Call sends a CCIP message from srcImpl to dstImpl via the testadapter pattern.
// Family-agnostic: works for any src/dst pair whose impls are registered.
func (m *CCIPTxGun) Call(_ *wasp.Generator) *wasp.Response {
	ctx := context.Background()
	b := ccip.NewDefaultCLDFBundle(m.e)
	m.e.OperationsBundle = b

	receiver := m.dstImpl.CCIPReceiver()
	extraArgs, err := m.dstImpl.GetExtraArgs(receiver, m.srcImpl.Family())
	if err != nil {
		return &wasp.Response{Error: err.Error(), Failed: true}
	}

	msg, err := m.srcImpl.BuildMessage(testadapters.MessageComponents{
		DestChainSelector: m.dstImpl.ChainSelector(),
		Receiver:          receiver,
		Data:              []byte("load test"),
		FeeToken:          "",
		ExtraArgs:         extraArgs,
	})
	if err != nil {
		return &wasp.Response{Error: err.Error(), Failed: true}
	}

	if _, _, err := m.srcImpl.SendMessage(ctx, m.dstImpl.ChainSelector(), msg); err != nil {
		return &wasp.Response{Error: err.Error(), Failed: true}
	}
	return &wasp.Response{Data: "ok"}
}

func gasControlFunc(t *testing.T, r *rpc.RPCClient, blockPace time.Duration) {
	startGasPrice := big.NewInt(2e9)
	// ramp
	for range 10 {
		err := r.PrintBlockBaseFee()
		require.NoError(t, err)
		err = r.AnvilSetNextBlockBaseFeePerGas(startGasPrice)
		require.NoError(t, err)
		startGasPrice = startGasPrice.Add(startGasPrice, big.NewInt(1e9))
		time.Sleep(blockPace)
	}
	// hold
	for range 10 {
		err := r.PrintBlockBaseFee()
		require.NoError(t, err)
		time.Sleep(blockPace)
		err = r.AnvilSetNextBlockBaseFeePerGas(startGasPrice)
		require.NoError(t, err)
	}
	// release
	for range 10 {
		err := r.PrintBlockBaseFee()
		require.NoError(t, err)
		time.Sleep(blockPace)
	}
}

func createLoadProfile(rps int64, testDuration time.Duration, e *deployment.Environment, srcImpl, dstImpl ccip.CCIP16ProductConfiguration) (*wasp.Profile, *CCIPTxGun) {
	gun := NewCCIPTxGun(e, srcImpl, dstImpl)
	profile := wasp.NewProfile().
		Add(wasp.NewGenerator(&wasp.Config{
			LoadType: wasp.RPS,
			GenName:  "src-dst-single-token",
			Schedule: wasp.Combine(
				wasp.Plain(rps, testDuration),
			),
			Gun: gun,
			Labels: map[string]string{
				"go_test_name": "load-clean-src",
				"branch":       "test",
				"commit":       "test",
			},
			LokiConfig: wasp.NewEnvLokiConfig(),
		}))
	return profile, gun
}

func TestE2ELoad(t *testing.T) {
	in, err := ccip.LoadOutput[ccip.Cfg]("../../env-out.toml")
	require.NoError(t, err)
	if in.ForkedEnvConfig != nil {
		t.Skip("Skipping E2E tests on forked environments, not supported yet")
	}
	if os.Getenv("LOKI_URL") == "" {
		_ = os.Setenv("LOKI_URL", ccip.DefaultLokiURL)
	}
	srcRPCURL := in.Blockchains[0].Out.Nodes[0].ExternalHTTPUrl
	dstRPCURL := in.Blockchains[1].Out.Nodes[0].ExternalHTTPUrl

	selectors, e, err := ccip.NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	require.NoError(t, err)
	b := ccip.NewDefaultCLDFBundle(e)
	e.OperationsBundle = b

	impls := make([]ccip.CCIP16ProductConfiguration, 0)
	for _, selector := range selectors {
		family, err := chainsel.GetSelectorFamily(selector)
		require.NoError(t, err)
		chainID, err := chainsel.GetChainIDFromSelector(selector)
		require.NoError(t, err)
		i, err := ccip.NewCCIPImplFromNetwork(family, chainID)
		require.NoError(t, err)
		i.SetCLDF(e)
		impls = append(impls, i)
	}

	t.Run("clean", func(t *testing.T) {
		// just a clean load test to measure performance
		rps := int64(1)
		loadDuration := 15 * time.Minute

		p, _ := createLoadProfile(rps, loadDuration, e, impls[0], impls[1])

		_, err = p.Run(true)
		require.NoError(t, err)

		assertLoki(t, in, time.Now())
		assertPrometheus(t, in, time.Now())
	})

	t.Run("rpc latency", func(t *testing.T) {
		// 400ms latency for any RPC node
		_, err = chaos.ExecPumba("netem --tc-image=ghcr.io/alexei-led/pumba-debian-nettools --duration=30s delay --time=400 re2:blockchain-node-.*", 0*time.Second)
		require.NoError(t, err)

		rps := int64(1)
		loadDuration := 30 * time.Second

		p, _ := createLoadProfile(rps, loadDuration, e, impls[0], impls[1])

		_, err = p.Run(true)
		require.NoError(t, err)

		assertLoki(t, in, time.Now())
		assertPrometheus(t, in, time.Now())
	})

	t.Run("gas", func(t *testing.T) {
		rps := int64(1)
		loadDuration := 30 * time.Second

		p, _ := createLoadProfile(rps, loadDuration, e, impls[0], impls[1])

		_, err = p.Run(false)
		require.NoError(t, err)

		waitBetweenTests := 5 * time.Second

		tcs := []GasTestCase{
			{
				name:             "Slow spike src",
				chainURL:         srcRPCURL,
				waitBetweenTests: waitBetweenTests,
				increase:         big.NewInt(1e9),
				gasFunc:          gasControlFunc,
				validate:         func() error { return nil },
			},
			{
				name:             "Fast spike src",
				chainURL:         srcRPCURL,
				waitBetweenTests: waitBetweenTests,
				increase:         big.NewInt(5e9),
				gasFunc:          gasControlFunc,
				validate:         func() error { return nil },
			},
			{
				name:             "Slow spike dst",
				chainURL:         dstRPCURL,
				waitBetweenTests: waitBetweenTests,
				increase:         big.NewInt(1e9),
				gasFunc:          gasControlFunc,
				validate:         func() error { return nil },
			},
			{
				name:             "Fast spike dst",
				chainURL:         dstRPCURL,
				waitBetweenTests: waitBetweenTests,
				increase:         big.NewInt(5e9),
				gasFunc:          gasControlFunc,
				validate:         func() error { return nil },
			},
		}
		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				t.Log(tc.name)
				r := rpc.New(tc.chainURL, nil)
				tc.gasFunc(t, r, 1*time.Second)
				err = tc.validate()
				require.NoError(t, err)
				time.Sleep(tc.waitBetweenTests)
			})
		}
		p.Wait()

		assertLoki(t, in, time.Now())
		assertPrometheus(t, in, time.Now())
	})

	t.Run("reorgs", func(t *testing.T) {
		// env-geth.toml is required for the environment!
		rps := int64(1)
		loadDuration := 120 * time.Second

		p, _ := createLoadProfile(rps, loadDuration, e, impls[0], impls[1])

		_, err = p.Run(false)
		require.NoError(t, err)

		tcs := []struct {
			validate   func() error
			name       string
			chainURL   string
			wait       time.Duration
			reorgDepth int
		}{
			{
				name:       "Reorg src with depth: 1",
				wait:       30 * time.Second,
				chainURL:   srcRPCURL,
				reorgDepth: 1,
				validate: func() error {
					// add clients and validate
					return nil
				},
			},
			{
				name:       "Reorg dst with depth: 1",
				wait:       30 * time.Second,
				chainURL:   dstRPCURL,
				reorgDepth: 1,
				validate: func() error {
					return nil
				},
			},
			{
				name:       "Reorg src with depth: 5",
				wait:       30 * time.Second,
				chainURL:   srcRPCURL,
				reorgDepth: 5,
				validate: func() error {
					return nil
				},
			},
			{
				name:       "Reorg dst with depth: 5",
				wait:       30 * time.Second,
				chainURL:   dstRPCURL,
				reorgDepth: 5,
				validate: func() error {
					return nil
				},
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				r := rpc.New(tc.chainURL, nil)
				err := r.GethSetHead(tc.reorgDepth)
				require.NoError(t, err)
				time.Sleep(tc.wait)
				err = tc.validate()
				require.NoError(t, err)
			})
		}
		p.Wait()

		assertLoki(t, in, time.Now())
		assertPrometheus(t, in, time.Now())
	})

	t.Run("services_chaos", func(t *testing.T) {
		rps := int64(1)
		loadDuration := 90 * time.Second

		tcs := []ChaosTestCase{
			{
				name: "Reboot a single node",
				run: func() error {
					_, err = chaos.ExecPumba(
						"stop --duration=20s --restart re2:don-node1",
						30*time.Second,
					)
					return nil
				},
				validate: func() error { return nil },
			},
			{
				name: "Reboot two nodes",
				run: func() error {
					_, err = chaos.ExecPumba(
						"stop --duration=20s --restart re2:don-node1",
						0*time.Second,
					)
					_, err = chaos.ExecPumba(
						"stop --duration=20s --restart re2:don-node2",
						30*time.Second,
					)
					return err
				},
				validate: func() error { return nil },
			},
			{
				name: "One slow CL node",
				run: func() error {
					_, err = chaos.ExecPumba(
						"netem --tc-image=ghcr.io/alexei-led/pumba-debian-nettools --duration=30s delay --time=1000 re2:don-node1",
						30*time.Second,
					)
					return err
				},
				validate: func() error { return nil },
			},
		}

		p, _ := createLoadProfile(rps, loadDuration, e, impls[0], impls[1])

		_, err = p.Run(false)
		require.NoError(t, err)

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				t.Log(tc.name)
				err = tc.run()
				require.NoError(t, err)
				err = tc.validate()
				require.NoError(t, err)
			})
		}
		p.Wait()

		assertLoki(t, in, time.Now())
		assertPrometheus(t, in, time.Now())
	})
}

// assertLoki this is an example method demonstrating how we can grep logs to assert
func assertLoki(t *testing.T, in *ccip.Cfg, end time.Time) {
	logs, err := f.NewLokiQueryClient(f.LocalLokiBaseURL, "", f.BasicAuth{}, f.QueryParams{
		Query:     "{job=\"ctf\",container=\"don-node1\"}",
		StartTime: end.Add(-time.Minute),
		EndTime:   end,
		Limit:     100,
	}).QueryRange(context.Background())
	require.NoError(t, err)
	fmt.Println(logs)
}

// assertPrometheus is an example method demonstrating how we can assert Prometheus metrics
func assertPrometheus(t *testing.T, in *ccip.Cfg, end time.Time) {
	pc := f.NewPrometheusQueryClient(f.LocalPrometheusBaseURL)
	// no more than 10% CPU for this test
	maxCPU := 10.0
	cpuResp, err := pc.Query("sum(rate(container_cpu_usage_seconds_total{name=~\".*don.*\"}[5m])) by (name) *100", end)
	require.NoError(t, err)
	cpu := f.ToLabelsMap(cpuResp)
	for i := 0; i < in.NodeSets[0].Nodes; i++ {
		nodeLabel := fmt.Sprintf("name:don-node%d", i)
		nodeCpu, err := strconv.ParseFloat(cpu[nodeLabel][0].(string), 64)
		ccip.Plog.Info().Int("Node", i).Float64("CPU", nodeCpu).Msg("CPU usage percentage")
		require.NoError(t, err)
		require.LessOrEqual(t, nodeCpu, maxCPU)
	}
	// no more than 400mb for this test
	maxMem := int(400e6) // 200mb
	memoryResp, err := pc.Query("sum(container_memory_rss{name=~\".*don.*\"}) by (name)", end)
	require.NoError(t, err)
	mem := f.ToLabelsMap(memoryResp)
	for i := 0; i < in.NodeSets[0].Nodes; i++ {
		nodeLabel := fmt.Sprintf("name:don-node%d", i)
		nodeMem, err := strconv.Atoi(mem[nodeLabel][0].(string))
		ccip.Plog.Info().Int("Node", i).Int("Memory", nodeMem).Msg("Total memory")
		require.NoError(t, err)
		require.LessOrEqual(t, nodeMem, maxMem)
	}
}
