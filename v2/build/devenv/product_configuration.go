package ccv

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
)

const (
	ConfigureNodesNetwork ConfigPhase = iota
	ConfigureProductContractsJobs
)

var Plog = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel).With().Fields(map[string]any{"component": "ccv"}).Logger()

type CCV struct {
	EAFake              *EAFake       `toml:"ea_fake"`
	Jobs                *Jobs         `toml:"jobs"`
	LinkContractAddress string        `toml:"link_contract_address"`
	CLNodesFundingETH   float64       `toml:"cl_nodes_funding_eth"`
	CLNodesFundingLink  float64       `toml:"cl_nodes_funding_link"`
	ChainFinalityDepth  int64         `toml:"chain_finality_depth"`
	VerificationTimeout time.Duration `toml:"verification_timeout"`
	Verify              bool          `toml:"verify"`

	// Contracts (CLDF)
	AddressesMu *sync.Mutex `toml:"-"`
	Addresses   []string    `toml:"addresses"`

	// These are the settings for CLDF missing functionality we cover with ETHClient, we should remove them later
	GasSettings *GasSettings `toml:"gas_settings"`
}

type DeployedContracts struct {
	// your deployed contract structs here with `toml:''` tags
}

type GasSettings struct {
	FeeCapMultiplier int64 `toml:"fee_cap_multiplier"`
	TipCapMultiplier int64 `toml:"tip_cap_multiplier"`
}

type Jobs struct {
	ConfigPollIntervalSeconds time.Duration `toml:"config_poll_interval_sec"` //nolint:staticcheck
	MaxTaskDurationSec        time.Duration `toml:"max_task_duration_sec"`    //nolint:staticcheck
}

type EAFake struct {
	LowValue  int64 `toml:"low_value"`
	HighValue int64 `toml:"high_value"`
}

type ConfigPhase int

func NewCLDFOperationsEnvironment(bc []*blockchain.Input) ([]uint64, *deployment.Environment, error) {
	providers := make([]cldf_chain.BlockChain, 0)
	selectors := make([]uint64, 0)
	for _, b := range bc {
		chainID := b.Out.ChainID
		rpcWSURL := b.Out.Nodes[0].ExternalWSUrl
		rpcHTTPURL := b.Out.Nodes[0].ExternalHTTPUrl

		d, err := chainsel.GetChainDetailsByChainIDAndFamily(chainID, chainsel.FamilyEVM)
		if err != nil {
			return nil, nil, err
		}
		selectors = append(selectors, d.ChainSelector)

		p, err := cldf_evm_provider.NewRPCChainProvider(
			d.ChainSelector,
			cldf_evm_provider.RPCChainProviderConfig{
				DeployerTransactorGen: cldf_evm_provider.TransactorFromRaw(
					getNetworkPrivateKey(),
				),
				RPCs: []deployment.RPC{
					{
						Name:               "default",
						WSURL:              rpcWSURL,
						HTTPURL:            rpcHTTPURL,
						PreferredURLScheme: deployment.URLSchemePreferenceHTTP,
					},
				},
				ConfirmFunctor: cldf_evm_provider.ConfirmFuncGeth(1 * time.Minute),
			},
		).Initialize(context.Background())
		if err != nil {
			return nil, nil, err
		}
		providers = append(providers, p)
	}

	blockchains := cldf_chain.NewBlockChainsFromSlice(providers)

	lggr, err := logger.NewWith(func(config *zap.Config) {
		config.Development = true
		config.Encoding = "console"
	})
	if err != nil {
		return nil, nil, err
	}

	e := deployment.Environment{
		GetContext:  func() context.Context { return context.Background() },
		Logger:      lggr,
		BlockChains: blockchains,
		DataStore:   datastore.NewMemoryDataStore().Seal(),
	}
	return selectors, &e, nil
}

func deployContractsForSelector(in *Cfg, e *deployment.Environment, selector uint64) error {
	L.Info().Uint64("Selector", selector).Msg("Configuring per-chain contracts bundle")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle
	out, err := changesets.DeployChainContracts.Apply(*e, changesets.DeployChainContractsCfg{
		ChainSelector: selector,
		Params: sequences.ContractParams{
			// TODO: Router contract implementation is missing
			RMNRemote:     sequences.RMNRemoteParams{},
			CCVAggregator: sequences.CCVAggregatorParams{},
			CommitOnRamp: sequences.CommitOnRampParams{
				// TODO: add mocked contract here
				FeeAggregator: common.HexToAddress("0x01"),
			},
			CCVProxy: sequences.CCVProxyParams{
				FeeAggregator: common.HexToAddress("0x01"),
			},
			FeeQuoter: sequences.FeeQuoterParams{
				// expose in TOML config
				MaxFeeJuelsPerMsg:              big.NewInt(2e18),
				TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
				LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
				WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
			},
			CommitOffRamp: sequences.CommitOffRampParams{
				SignatureConfigArgs: commit_offramp.SignatureConfigArgs{
					{
						// OCR3 or something else?
						ConfigDigest: [32]byte{0x01},
						F:            1,
						Signers: []common.Address{
							common.HexToAddress("0x02"),
							common.HexToAddress("0x03"),
							common.HexToAddress("0x04"),
							common.HexToAddress("0x05"),
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	addresses, err := out.DataStore.Addresses().Fetch()
	in.CCV.AddressesMu.Lock()
	defer in.CCV.AddressesMu.Unlock()
	a, err := json.Marshal(addresses)
	if err != nil {
		return err
	}
	in.CCV.Addresses = append(in.CCV.Addresses, string(a))
	return nil
}

func configureJobs(in *Cfg, clNodes []*clclient.ChainlinkClient) error {
	bootstrapNode := clNodes[0]
	workerNodes := clNodes[1:]
	// example bootstrap job, use JD here?
	_ = bootstrapNode

	for _, chainlinkNode := range workerNodes {
		_, err := chainlinkNode.PrimaryEthAddress()
		if err != nil {
			return fmt.Errorf("getting primary ETH address from OCR node have failed: %w", err)
		}
		_, err = chainlinkNode.MustReadOCR2Keys()
		if err != nil {
			return fmt.Errorf("getting OCR keys from OCR node have failed: %w", err)
		}
		_ = in.Fake.Out.ExternalHTTPURL
		_ = in.Fake.Out.InternalHTTPURL

		// create CCV jobs here
	}
	return nil
}

func setupFakes(fakeServiceURL string) error {
	//  example fake service configuration (mocking external adapter responses)
	//  r := resty.New().SetBaseURL(fakeServiceURL)
	//  _, err = r.R().Post(fmt.Sprintf(`/set_ea?low=%d&high=%d`, in.CCV.EAFake.LowValue, in.CCV.EAFake.HighValue))
	//  if err != nil {
	//  	return fmt.Errorf("could not set ea fake values: %w", err)
	//  }
	//  Plog.Info().
	//	  Int64("FeedAnswerLow", in.CCV.EAFake.LowValue).
	//	  Int64("FeedAnswerHigh", in.CCV.EAFake.HighValue).
	//	  Msg("Setting fake external adapter (data feed) values")
	return nil
}

// DefaultProductConfiguration is default product configuration that includes:
// - CL nodes config generation
// - On-chain part deployment using CLDF
func DefaultProductConfiguration(in *Cfg, phase ConfigPhase) error {
	Plog.Info().Msg("Generating CL nodes config")
	pkey := getNetworkPrivateKey()
	if pkey == "" {
		return fmt.Errorf("PRIVATE_KEY environment variable not set")
	}
	switch phase {
	case ConfigureNodesNetwork:
		Plog.Info().Msg("Applying default CL nodes configuration")
		srcBlockchain := in.Blockchains[0].Out.Nodes[0]
		dstBlockchain := in.Blockchains[1].Out.Nodes[0]
		srcChainID := in.Blockchains[0].ChainID
		dstChainID := in.Blockchains[1].ChainID
		// configure node set and generate CL nodes configs
		netConfig := fmt.Sprintf(`
       [[EVM]]
       LogPollInterval = '1s'
       BlockBackfillDepth = 100
       LinkContractAddress = '%s'
       ChainID = '%s'
       MinIncomingConfirmations = 1
       MinContractPayment = '0.0000001 link'
       FinalityDepth = %d

       [[EVM.Nodes]]
       Name = 'src'
       WsUrl = '%s'
       HttpUrl = '%s'

       [[EVM]]
       LogPollInterval = '1s'
       BlockBackfillDepth = 100
       LinkContractAddress = '%s'
       ChainID = '%s'
       MinIncomingConfirmations = 1
       MinContractPayment = '0.0000001 link'
       FinalityDepth = %d

       [[EVM.Nodes]]
       Name = 'dst'
       WsUrl = '%s'
       HttpUrl = '%s'
`,
			in.CCV.LinkContractAddress,
			srcChainID,
			in.CCV.ChainFinalityDepth,
			srcBlockchain.InternalWSUrl,
			srcBlockchain.InternalHTTPUrl,

			in.CCV.LinkContractAddress,
			dstChainID,
			in.CCV.ChainFinalityDepth,
			dstBlockchain.InternalWSUrl,
			dstBlockchain.InternalHTTPUrl,
		) + `
	   [Log]
	   JSONConsole = true
	   Level = 'debug'
	   [Pyroscope]
	   ServerAddress = 'http://host.docker.internal:4040'
	   Environment = 'local'
	   [WebServer]
       SessionTimeout = '999h0m0s'
       HTTPWriteTimeout = '3m'
	   SecureCookies = false
	   HTTPPort = 6688
	   [WebServer.TLS]
	   HTTPSPort = 0
       [WebServer.RateLimit]
       Authenticated = 5000
       Unauthenticated = 5000
	   [JobPipeline]
	   [JobPipeline.HTTPRequest]
	   DefaultTimeout = '1m'
       [Log.File]
       MaxSize = '0b'
	` + `
       [Feature]
       FeedsManager = true
       LogPoller = true
       UICSAKeys = true
       [OCR2]
       Enabled = true
       SimulateTransactions = false
       DefaultTransactionQueueDepth = 1
       [P2P.V2]
       Enabled = true
       ListenAddresses = ['0.0.0.0:6690']
`
		for _, nodeSpec := range in.NodeSets[0].NodeSpecs {
			nodeSpec.Node.TestConfigOverrides = netConfig
		}
		Plog.Info().Msg("Nodes network configuration is generated")
		return nil
	case ConfigureProductContractsJobs:

		//* Funding all CL nodes with ETH *//

		Plog.Info().Msg("Connecting to CL nodes")
		nodeClients, err := clclient.New(in.NodeSets[0].Out.CLNodes)
		if err != nil {
			return fmt.Errorf("connecting to CL nodes: %w", err)
		}
		transmittersSrc, transmittersDst := make([]common.Address, 0), make([]common.Address, 0)
		ethKeyAddressesSrc, ethKeyAddressesDst := make([]string, 0), make([]string, 0)
		for i, nc := range nodeClients {
			addrSrc, err := nc.ReadPrimaryETHKey(in.Blockchains[0].ChainID)
			if err != nil {
				return fmt.Errorf("getting primary ETH key from OCR node %d (src chain): %w", i, err)
			}
			ethKeyAddressesSrc = append(ethKeyAddressesSrc, addrSrc.Attributes.Address)
			transmittersSrc = append(transmittersSrc, common.HexToAddress(addrSrc.Attributes.Address))
			addrDst, err := nc.ReadPrimaryETHKey(in.Blockchains[1].ChainID)
			if err != nil {
				return fmt.Errorf("getting primary ETH key from OCR node %d (dst chain): %w", i, err)
			}
			ethKeyAddressesDst = append(ethKeyAddressesDst, addrDst.Attributes.Address)
			transmittersDst = append(transmittersDst, common.HexToAddress(addrDst.Attributes.Address))
			Plog.Info().
				Int("Idx", i).
				Str("ETHKeySrc", addrSrc.Attributes.Address).
				Str("ETHKeyDst", addrDst.Attributes.Address).
				Msg("Node info")
		}
		clientSrc, _, _, err := ETHClient(in.Blockchains[0].Out.Nodes[0].ExternalWSUrl, in.CCV.GasSettings)
		if err != nil {
			return fmt.Errorf("could not create basic eth client: %w", err)
		}
		clientDst, _, _, err := ETHClient(in.Blockchains[1].Out.Nodes[0].ExternalWSUrl, in.CCV.GasSettings)
		if err != nil {
			return fmt.Errorf("could not create basic eth client: %w", err)
		}
		for _, addr := range ethKeyAddressesSrc {
			if err := FundNodeEIP1559(clientSrc, pkey, addr, in.CCV.CLNodesFundingETH); err != nil {
				return fmt.Errorf("failed to fund CL nodes on dst chain: %w", err)
			}
		}
		for _, addr := range ethKeyAddressesDst {
			if err := FundNodeEIP1559(clientDst, pkey, addr, in.CCV.CLNodesFundingETH); err != nil {
				return fmt.Errorf("failed to fund CL nodes on dst chain: %w", err)
			}
		}

		// * Configuring src and dst contracts * //
		selectors, e, err := NewCLDFOperationsEnvironment(in.Blockchains)
		if err != nil {
			return fmt.Errorf("creating CLDF operations environment: %w", err)
		}
		L.Info().Any("Selectors", selectors).Msg("Deploying for chain selectors")
		eg := &errgroup.Group{}
		in.CCV.AddressesMu = &sync.Mutex{}
		for _, sel := range selectors {
			eg.Go(func() error {
				err = deployContractsForSelector(in, e, sel)
				if err != nil {
					return fmt.Errorf("could not configure contracts for chain selector %d: %w", sel, err)
				}
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return err
		}
		if err := configureJobs(in, nodeClients); err != nil {
			return fmt.Errorf("could not configure jobs: %w", err)
		}
		if err := setupFakes(in.Fake.Out.ExternalHTTPURL); err != nil {
			return fmt.Errorf("could not setup fake server: %w", err)
		}

		Plog.Info().Str("BootstrapNode", in.NodeSets[0].Out.CLNodes[0].Node.ExternalURL).Send()
		for _, n := range in.NodeSets[0].Out.CLNodes[1:] {
			Plog.Info().Str("Node", n.Node.ExternalURL).Send()
		}
		if err := verifyEnvironment(in); err != nil {
			return err
		}
		return nil
	}
	return nil
}
