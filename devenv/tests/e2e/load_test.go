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

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/chaos"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/rpc"
	"github.com/smartcontractkit/chainlink-testing-framework/wasp"

	chainsel "github.com/smartcontractkit/chain-selectors"
	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
	f "github.com/smartcontractkit/chainlink-testing-framework/framework"
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

// SentMessage represents a message that was sent and needs verification.
type SentMessage struct {
	SeqNo     uint64
	MessageID [32]byte
	SentTime  time.Time
}

type EVMTXGun struct {
	cfg  *ccip.Cfg
	e    *deployment.Environment
	impl ccip.CCIP16ProductConfiguration
}

func NewEVMTransactionGun(cfg *ccip.Cfg, e *deployment.Environment, selectors []uint64, impl ccip.CCIP16ProductConfiguration, s, d evm.Chain) *EVMTXGun {
	return &EVMTXGun{
		cfg:  cfg,
		e:    e,
		impl: impl,
	}
}

// Call implements example gun call, assertions on response bodies should be done here.
func (m *EVMTXGun) Call(_ *wasp.Generator) *wasp.Response {
	b := ccip.NewDefaultCLDFBundle(m.e)
	m.e.OperationsBundle = b

	chainIDs := make([]string, 0)
	for _, bc := range m.cfg.Blockchains {
		chainIDs = append(chainIDs, bc.ChainID)
	}

	srcChain, err := chainsel.GetChainDetailsByChainIDAndFamily(chainIDs[0], chainsel.FamilyEVM)
	if err != nil {
		return &wasp.Response{Error: err.Error(), Failed: true}
	}
	dstChain, err := chainsel.GetChainDetailsByChainIDAndFamily(chainIDs[1], chainsel.FamilyEVM)
	if err != nil {
		return &wasp.Response{Error: err.Error(), Failed: true}
	}

	_, _ = srcChain, dstChain

	// err := m.impl.SendMessage(...)
	// if err != nil {
	// 	return &wasp.Response{Error: error.New("something"), Failed: true}
	// }
	return &wasp.Response{Data: "ok"}
}

func gasControlFunc(t *testing.T, r *rpc.RPCClient, blockPace time.Duration) {
	startGasPrice := big.NewInt(2e9)
	// ramp
	for i := 0; i < 10; i++ {
		err := r.PrintBlockBaseFee()
		require.NoError(t, err)
		err = r.AnvilSetNextBlockBaseFeePerGas(startGasPrice)
		require.NoError(t, err)
		startGasPrice = startGasPrice.Add(startGasPrice, big.NewInt(1e9))
		time.Sleep(blockPace)
	}
	// hold
	for i := 0; i < 10; i++ {
		err := r.PrintBlockBaseFee()
		require.NoError(t, err)
		time.Sleep(blockPace)
		err = r.AnvilSetNextBlockBaseFeePerGas(startGasPrice)
		require.NoError(t, err)
	}
	// release
	for i := 0; i < 10; i++ {
		err := r.PrintBlockBaseFee()
		require.NoError(t, err)
		time.Sleep(blockPace)
	}
}

func createLoadProfile(in *ccip.Cfg, rps int64, testDuration time.Duration, e *deployment.Environment, selectors []uint64, impl ccip.CCIP16ProductConfiguration, s, d evm.Chain) (*wasp.Profile, *EVMTXGun) {
	gun := NewEVMTransactionGun(in, e, selectors, impl, s, d)
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
	if os.Getenv("LOKI_URL") == "" {
		_ = os.Setenv("LOKI_URL", ccip.DefaultLokiURL)
	}
	srcRPCURL := in.Blockchains[0].Out.Nodes[0].ExternalHTTPUrl
	dstRPCURL := in.Blockchains[1].Out.Nodes[0].ExternalHTTPUrl

	selectors, e, err := ccip.NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	require.NoError(t, err)
	chains := e.BlockChains.EVMChains()
	require.NotNil(t, chains)
	srcChain := chains[selectors[0]]
	dstChain := chains[selectors[1]]
	b := ccip.NewDefaultCLDFBundle(e)
	e.OperationsBundle = b

	// ctx := ccip.Plog.WithContext(context.Background())

	chainIDs, wsURLs := make([]string, 0), make([]string, 0)
	for _, bc := range in.Blockchains {
		chainIDs = append(chainIDs, bc.ChainID)
		wsURLs = append(wsURLs, bc.Out.Nodes[0].ExternalWSUrl)
	}

	impls := make([]ccip.CCIP16ProductConfiguration, 0)
	for _, bc := range in.Blockchains {
		i, err := ccip.NewCCIPImplFromNetwork(bc.Out.Type)
		require.NoError(t, err)
		i.SetCLDF(e)
		impls = append(impls, i)
	}

	t.Run("clean", func(t *testing.T) {
		// just a clean load test to measure performance
		rps := int64(5)
		loadDuration := 30 * time.Second

		p, _ := createLoadProfile(in, rps, loadDuration, e, selectors, impls[0], srcChain, dstChain)

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

		p, _ := createLoadProfile(in, rps, loadDuration, e, selectors, impls[0], srcChain, dstChain)

		_, err = p.Run(true)
		require.NoError(t, err)

		assertLoki(t, in, time.Now())
		assertPrometheus(t, in, time.Now())
	})

	t.Run("gas", func(t *testing.T) {
		rps := int64(1)
		loadDuration := 30 * time.Second

		p, _ := createLoadProfile(in, rps, loadDuration, e, selectors, impls[0], srcChain, dstChain)

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

		p, _ := createLoadProfile(in, rps, loadDuration, e, selectors, impls[0], srcChain, dstChain)

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

		p, _ := createLoadProfile(in, rps, loadDuration, e, selectors, impls[0], srcChain, dstChain)

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
	// no more than 200mb for this test
	maxMem := int(200e6) // 200mb
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
