package ccip_evm

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/devenv/blockchainutils"

	chainsel "github.com/smartcontractkit/chain-selectors"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"

	evmseqs "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

type SourceDestPair struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
}

type AnyMsgSentEvent struct {
	SequenceNumber uint64
	// RawEvent contains the raw event depending on the chain:
	//  EVM:   *onramp.OnRampCCIPMessageSent
	//  Aptos: module_onramp.CCIPMessageSent
	RawEvent any
}

type CCIP16EVM struct {
	e            *deployment.Environment
	chainDetails chainsel.ChainDetails
	ethClients   map[uint64]*ethclient.Client
	testadapters.TestAdapter
}

func NewEmptyCCIP16EVM(chainDetails chainsel.ChainDetails) *CCIP16EVM {
	return &CCIP16EVM{
		chainDetails: chainDetails,
		ethClients:   make(map[uint64]*ethclient.Client),
	}
}

// NewCCIP16EVM creates new smart-contracts wrappers with utility functions for CCIP16EVM implementation.
func NewCCIP16EVM(ctx context.Context, e *deployment.Environment, chainDetails chainsel.ChainDetails) (*CCIP16EVM, error) {
	_ = zerolog.Ctx(ctx)
	out := NewEmptyCCIP16EVM(chainDetails)
	out.SetCLDF(e)
	return out, nil
}

func (m *CCIP16EVM) SetCLDF(e *deployment.Environment) {
	m.e = e
	factory, found := testadapters.GetTestAdapterRegistry().GetTestAdapter(chainsel.FamilyEVM, semver.MustParse("1.6.0"))
	if !found {
		panic(fmt.Sprintf("failed to find testadapter factory for %s", chainsel.FamilyEVM))
	}
	m.TestAdapter = factory(e, m.chainDetails.ChainSelector)
}

func (m *CCIP16EVM) ChainSelector() uint64 {
	return m.chainDetails.ChainSelector
}

func updatePrices(datastore datastore.DataStore, src, dest uint64, srcChain cldf_evm.Chain) error {
	a := &evmseqs.EVMAdapter{}
	fqAddr, err := a.GetFQAddress(datastore, src)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	fq, err := fee_quoter.NewFeeQuoter(
		common.BytesToAddress(fqAddr),
		srcChain.Client)
	if err != nil {
		return fmt.Errorf("failed to create fee quoter instance: %w", err)
	}
	feeTokens, err := fq.GetFeeTokens(nil)
	if err != nil {
		return fmt.Errorf("failed to get fee tokens from fee quoter: %w", err)
	}
	sender := srcChain.DeployerKey
	tx, err := fq.UpdatePrices(sender, fee_quoter.InternalPriceUpdates{
		TokenPriceUpdates: []fee_quoter.InternalTokenPriceUpdate{
			{
				SourceToken: feeTokens[0],
				UsdPerToken: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2000)),
			},
			{
				SourceToken: feeTokens[1],
				UsdPerToken: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2000)),
			},
		},
		GasPriceUpdates: []fee_quoter.InternalGasPriceUpdate{
			{
				DestChainSelector: dest,
				UsdPerUnitGas:     big.NewInt(20000e9),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to update prices: %w", err)
	}
	_, err = srcChain.Confirm(tx)
	if err != nil {
		return fmt.Errorf("failed to confirm update prices transaction: %w", err)
	}
	return nil
}

func (m *CCIP16EVM) GetEOAReceiverAddress(ctx context.Context, chainSelector uint64) ([]byte, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

func (m *CCIP16EVM) ExposeMetrics(
	ctx context.Context,
	source, dest uint64,
	chainIDs []string,
	wsURLs []string,
) ([]string, *prometheus.Registry, error) {
	msgSentTotal.Reset()
	msgExecTotal.Reset()
	srcDstLatency.Reset()

	reg := prometheus.NewRegistry()
	reg.MustRegister(msgSentTotal, msgExecTotal, srcDstLatency)

	lp := NewLokiPusher()
	err := ProcessLaneEvents(ctx, m, lp, &LaneStreamConfig{
		FromSelector:      source,
		ToSelector:        dest,
		AggregatorAddress: "localhost:50051",
		AggregatorSince:   0,
	})
	if err != nil {
		return nil, nil, err
	}
	err = ProcessLaneEvents(ctx, m, lp, &LaneStreamConfig{
		FromSelector:      dest,
		ToSelector:        source,
		AggregatorAddress: "localhost:50051",
		AggregatorSince:   0,
	})
	if err != nil {
		return nil, nil, err
	}
	return []string{}, reg, nil
}

func (m *CCIP16EVM) DeployLocalNetwork(ctx context.Context, bc *blockchain.Input) (*blockchain.Output, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Deploying EVM networks")
	out, err := blockchain.NewBlockchainNetwork(bc)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain network: %w", err)
	}
	return out, nil
}

func (m *CCIP16EVM) ConfigureNodes(ctx context.Context, bc *blockchain.Input) (string, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring CL nodes for evm")
	name := fmt.Sprintf("node-evm-%s", uuid.New().String()[0:5])

	// Check if this is an external chain (user pre-configured the Out section in TOML)
	// External chains are detected by checking if the URLs are external (not localhost/docker)
	isExternalChain := bc.Out != nil && len(bc.Out.Nodes) > 0 &&
		!strings.Contains(bc.Out.Nodes[0].ExternalHTTPUrl, "host.docker.internal") &&
		!strings.Contains(bc.Out.Nodes[0].ExternalHTTPUrl, "localhost") &&
		!strings.Contains(bc.Out.Nodes[0].ExternalHTTPUrl, "127.0.0.1")

	if isExternalChain {
		// For external chains (testnets/mainnets), don't generate any EVM config.
		// The user must provide the full [[EVM]] config including [[EVM.Nodes]] via node_config_overrides.
		// This avoids duplicate ChainID errors when both auto-generated and user configs exist.
		l.Info().Str("ChainID", bc.ChainID).Msg("External chain detected - skipping auto-generated EVM config (user provides via node_config_overrides)")
		return "", nil
	}

	// For local chains, generate full EVM config
	finality := 1
	return fmt.Sprintf(`
[[EVM]]
LogPollInterval = '1s'
BlockBackfillDepth = 100
ChainID = '%s'
MinIncomingConfirmations = 1
MinContractPayment = '0.0000001 link'
FinalityDepth = %d

[[EVM.Nodes]]
Name = '%s'
WSURL = '%s'
HTTPURL = '%s'`,
		bc.ChainID,
		finality,
		name,
		bc.Out.Nodes[0].InternalWSUrl,
		bc.Out.Nodes[0].InternalHTTPUrl,
	), nil
}

func (m *CCIP16EVM) PreDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	return nil
}

func (m *CCIP16EVM) PostDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	return nil
}

func (m *CCIP16EVM) FundNodes(ctx context.Context, ns []*simple_node_set.Input, nodeKeyBundles map[string]clclient.NodeKeysBundle, bc *blockchain.Input, linkAmount, nativeAmount *big.Int) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Funding CL nodes with ETH and LINK")
	nodeClients, err := clclient.New(ns[0].Out.CLNodes)
	if err != nil {
		return fmt.Errorf("connecting to CL nodes: %w", err)
	}
	ethKeyAddressesSrc := make([]string, 0)
	for i, nc := range nodeClients {
		addrSrc, err := nc.ReadPrimaryETHKey(bc.ChainID)
		if err != nil {
			return fmt.Errorf("getting primary ETH key from OCR node %d (src chain): %w", i, err)
		}
		ethKeyAddressesSrc = append(ethKeyAddressesSrc, addrSrc.Attributes.Address)
		l.Info().
			Int("Idx", i).
			Str("ETHKeySrc", addrSrc.Attributes.Address).
			Msg("Node info")
	}
	// Use WS URL if available, otherwise fallback to HTTP URL for HTTP-only mode
	rpcURL := bc.Out.Nodes[0].ExternalWSUrl
	if rpcURL == "" {
		rpcURL = bc.Out.Nodes[0].ExternalHTTPUrl
		l.Info().Str("URL", rpcURL).Msg("Using HTTP URL for ETH client (HTTP-only mode)")
	}
	clientSrc, _, _, err := blockchainutils.ETHClient(ctx, rpcURL, &blockchainutils.GasSettings{
		FeeCapMultiplier: 2,
		TipCapMultiplier: 2,
	})
	if err != nil {
		return fmt.Errorf("could not create basic eth client: %w", err)
	}
	// Use default Anvil key for local chain 1337, otherwise use PRIVATE_KEY env var
	privateKey := blockchainutils.GetNetworkPrivateKey()
	if bc.ChainID == "1337" {
		privateKey = blockchainutils.DefaultAnvilKey
	}

	// nativeAmount is in ETH units (integer) - use directly for FundNodeEIP1559
	// EVM-specific conversion: FundNodeEIP1559 expects ETH as float64
	nativeAmountETH := float64(nativeAmount.Int64())

	for _, addr := range ethKeyAddressesSrc {
		if err := blockchainutils.FundNodeEIP1559(ctx, clientSrc, privateKey, addr, nativeAmountETH); err != nil {
			return fmt.Errorf("failed to fund CL nodes on src chain: %w", err)
		}
	}
	return nil
}
