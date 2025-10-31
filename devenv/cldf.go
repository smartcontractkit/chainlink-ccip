package ccip

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider/rpcclient"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	ccipEVM "github.com/smartcontractkit/chainlink-ccip/devenv/chainimpl/ccip-evm"
)

type CLDF struct {
	mu        sync.Mutex          `toml:"-"`
	Addresses []string            `toml:"addresses"`
	DataStore datastore.DataStore `toml:"-"`
}

func (c *CLDF) Init() {
	c.DataStore = datastore.NewMemoryDataStore().Seal()
}

func (c *CLDF) AddAddresses(addresses string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Addresses = append(c.Addresses, addresses)
}

func NewCLDFOperationsEnvironment(bc []*blockchain.Input, dataStore datastore.DataStore) ([]uint64, *deployment.Environment, error) {
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
				RPCs: []rpcclient.RPC{
					{
						Name:               "default",
						WSURL:              rpcWSURL,
						HTTPURL:            rpcHTTPURL,
						PreferredURLScheme: rpcclient.URLSchemePreferenceHTTP,
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
		DataStore:   dataStore,
	}
	return selectors, &e, nil
}

// NewDefaultCLDFBundle creates a new default CLDF bundle.
func NewDefaultCLDFBundle(e *deployment.Environment) operations.Bundle {
	return operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
}

func NewCCIPImplFromNetwork(typ string) (CCIP16ProductConfiguration, error) {
	switch typ {
	case "anvil", "geth":
		return &ccipEVM.CCIP16EVM{}, nil
	case "solana":
		panic("implement Solana")
	case "sui":
		panic("implement Sui")
	case "aptos":
		panic("implement Aptos")
	case "ton":
		panic("implement TON")
	default:
		return nil, errors.New("unknown devenv network type " + typ)
	}
}


