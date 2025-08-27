package ccipv1_7

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
)

const (
	ConfigureNodesNetwork ConfigPhase = iota
	ConfigureProductContractsJobs
)

var Plog = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel).With().Fields(map[string]any{"component": "ccipv1_7"}).Logger()

type CCIPv17 struct {
	EAFake                        *EAFake            `toml:"ea_fake"`
	Jobs                          *Jobs              `toml:"jobs"`
	GasSettings                   *GasSettings       `toml:"gas_settings"`
	DeployedContracts             *DeployedContracts `toml:"deployed_contracts"`
	LinkContractAddress           string             `toml:"link_contract_address"`
	CLNodesFundingETH             float64            `toml:"cl_nodes_funding_eth"`
	CLNodesFundingLink            float64            `toml:"cl_nodes_funding_link"`
	ChainFinalityDepth            int64              `toml:"chain_finality_depth"`
	VerificationTimeout           time.Duration      `toml:"verification_timeout"`
	ContractsConfigurationTimeout time.Duration      `toml:"contracts_configuration_timeout"`
	Verify                        bool               `toml:"verify"`
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

func configureSrcContracts(in *Cfg, c *ethclient.Client, auth *bind.TransactOpts, cl []*clclient.ChainlinkClient, rootAddr string, transmitters []common.Address) error {
	Plog.Info().Msg("Deploying LINK token contract")
	lggr, err := logger.New()
	if err != nil {
		return err
	}

	srcChainID := in.Blockchains[0].Out.ChainID
	//dstChainID := in.Blockchains[1].Out.ChainID
	srcRPCWSURL := in.Blockchains[0].Out.Nodes[0].ExternalWSUrl
	srcRPCHTTPURL := in.Blockchains[0].Out.Nodes[0].ExternalHTTPUrl
	//dstRPCWSURL := in.Blockchains[1].Out.Nodes[0].ExternalWSUrl
	//dstRPCHTTPURL := in.Blockchains[1].Out.Nodes[0].ExternalHTTPUrl

	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)

	srcChainDetails, err := chainsel.GetChainDetailsByChainIDAndFamily(srcChainID, chainsel.FamilyEVM)
	if err != nil {
		return err
	}
	L.Warn().Str("HTTP", srcRPCHTTPURL).Str("WS", srcRPCWSURL).Send()

	srcChainProvider, err := cldf_evm_provider.NewRPCChainProvider(
		srcChainDetails.ChainSelector,
		cldf_evm_provider.RPCChainProviderConfig{
			DeployerTransactorGen: cldf_evm_provider.TransactorFromRaw(
				getNetworkPrivateKey(),
			),
			RPCs: []deployment.RPC{
				{
					Name:               "default",
					WSURL:              srcRPCWSURL,
					HTTPURL:            srcRPCHTTPURL,
					PreferredURLScheme: deployment.URLSchemePreferenceHTTP,
				},
			},
			ConfirmFunctor: cldf_evm_provider.ConfirmFuncGeth(30),
		},
	).Initialize(context.Background())
	if err != nil {
		return err
	}

	chains := cldf_chain.NewBlockChainsFromSlice(
		[]cldf_chain.BlockChain{srcChainProvider},
	)

	e := deployment.Environment{
		GetContext:       func() context.Context { return context.Background() },
		Logger:           lggr,
		OperationsBundle: bundle,
		BlockChains:      chains,
		DataStore:        datastore.NewMemoryDataStore().Seal(),
	}

	out, err := changesets.DeployChainContracts.Apply(e, changesets.DeployChainContractsCfg{
		ChainSelector: srcChainDetails.ChainSelector,
		Params: sequences.ContractParams{
			RMNRemote: sequences.RMNRemoteParams{},
		},
	})
	if err != nil {
		return err
	}

	addresses, err := out.DataStore.Addresses().Fetch()
	spew.Dump(addresses)

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
		_ = in.FakeServer.Out.BaseURLDocker

		// create CCIPv17 jobs here
	}
	return nil
}

func setupFakes(fakeServiceURL string) error {
	//  example fake service configuration (mocking external adapter responses)
	//  r := resty.New().SetBaseURL(fakeServiceURL)
	//  _, err = r.R().Post(fmt.Sprintf(`/set_ea?low=%d&high=%d`, in.CCIPv17.EAFake.LowValue, in.CCIPv17.EAFake.HighValue))
	//  if err != nil {
	//  	return fmt.Errorf("could not set ea fake values: %w", err)
	//  }
	//  Plog.Info().
	//	  Int64("FeedAnswerLow", in.CCIPv17.EAFake.LowValue).
	//	  Int64("FeedAnswerHigh", in.CCIPv17.EAFake.HighValue).
	//	  Msg("Setting fake external adapter (data feed) values")
	return nil
}

// DefaultProductConfiguration is default product configuration that includes:
// - Deploying required prerequisites (LINK token, shared contracts)
// - Applying product-specific changesets
// - Creating cldf.Environment, connecting to components, see *Cfg fields
// - Generating CL nodes configs
// All the data can be added *Cfg struct like and is synced between local machine and remote environment
// so later both local and remote tests can use it.
func DefaultProductConfiguration(in *Cfg, phase ConfigPhase) error {
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
			in.CCIPv17.LinkContractAddress,
			srcChainID,
			in.CCIPv17.ChainFinalityDepth,
			srcBlockchain.InternalWSUrl,
			srcBlockchain.InternalHTTPUrl,

			in.CCIPv17.LinkContractAddress,
			dstChainID,
			in.CCIPv17.ChainFinalityDepth,
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
		Plog.Info().Msg("Nodes network configuration is finished")
	case ConfigureProductContractsJobs:
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
		clientSrc, authSrc, rootAddr, err := ETHClient(in.Blockchains[0].Out.Nodes[0].ExternalWSUrl, in.CCIPv17.GasSettings)
		if err != nil {
			return fmt.Errorf("could not create basic eth client: %w", err)
		}
		clientDst, authDst, _, err := ETHClient(in.Blockchains[1].Out.Nodes[0].ExternalWSUrl, in.CCIPv17.GasSettings)
		if err != nil {
			return fmt.Errorf("could not create basic eth client: %w", err)
		}
		for _, addr := range ethKeyAddressesSrc {
			if err := FundNodeEIP1559(clientSrc, pkey, addr, in.CCIPv17.CLNodesFundingETH); err != nil {
				return fmt.Errorf("failed to fund CL nodes on dst chain: %w", err)
			}
		}
		for _, addr := range ethKeyAddressesDst {
			if err := FundNodeEIP1559(clientDst, pkey, addr, in.CCIPv17.CLNodesFundingETH); err != nil {
				return fmt.Errorf("failed to fund CL nodes on dst chain: %w", err)
			}
		}
		err = configureSrcContracts(in, clientSrc, authSrc, nodeClients, rootAddr, transmittersSrc)
		if err != nil {
			return fmt.Errorf("could not configure contracts (src chain): %w", err)
		}
		err = configureSrcContracts(in, clientDst, authDst, nodeClients, rootAddr, transmittersDst)
		if err != nil {
			return fmt.Errorf("could not configure contracts (dst chain): %w", err)
		}
		if err := configureJobs(in, nodeClients); err != nil {
			return fmt.Errorf("could not configure jobs: %w", err)
		}
		if err := setupFakes(in.FakeServer.Out.BaseURLHost); err != nil {
			return fmt.Errorf("could not setup fake server: %w", err)
		}

		Plog.Info().Str("BootstrapNode", in.NodeSets[0].Out.CLNodes[0].Node.ExternalURL).Send()
		for _, n := range in.NodeSets[0].Out.CLNodes[1:] {
			Plog.Info().Str("Node", n.Node.ExternalURL).Send()
		}
		if err := verifyEnvironment(in); err != nil {
			return err
		}
		// store contract addresses or use CLD address book here
		in.CCIPv17.DeployedContracts = &DeployedContracts{}
	}
	return nil
}
