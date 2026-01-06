package ccip

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider/rpcclient"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"

	solRpc "github.com/gagliardetto/solana-go/rpc"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	ccipEVM "github.com/smartcontractkit/chainlink-ccip/devenv/chainimpl/ccip-evm"
	ccipSolana "github.com/smartcontractkit/chainlink-ccip/devenv/chainimpl/ccip-solana"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	cldf_solana_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana/provider"
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
	runningDS := datastore.NewMemoryDataStore()
	err := runningDS.Merge(dataStore)
	if err != nil {
		return nil, nil, err
	}
	providers := make([]cldf_chain.BlockChain, 0)
	selectors := make([]uint64, 0)
	for _, b := range bc {
		if b.Type == "anvil" || b.Type == "geth" {
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
					ConfirmFunctor: cldf_evm_provider.ConfirmFuncGeth(1*time.Minute, cldf_evm_provider.WithTickInterval(5*time.Millisecond)),
				},
			).Initialize(context.Background())
			if err != nil {
				return nil, nil, err
			}
			providers = append(providers, p)
		} else if b.Type == "solana" {
			chainID := b.ChainID
			rpcHTTPURL := b.Out.Nodes[0].ExternalHTTPUrl
			rpcWSURL := b.Out.Nodes[0].ExternalWSUrl
			programsPath, err := filepath.Abs(b.ContractsDir)
			if err != nil {
				return nil, nil, err
			}

			if err := os.MkdirAll(programsPath, 0o755); err != nil {
				return nil, nil, err
			}

			d, err := chainsel.GetChainDetailsByChainIDAndFamily(chainID, chainsel.FamilySolana)
			if err != nil {
				return nil, nil, err
			}

			selectors = append(selectors, d.ChainSelector)
			deployerKey := solana.MustPrivateKeyFromBase58("jW5nUtGGFzLA9kfgn6xWG497SdToPLqB8g485HrvFxK727iZNzKJu95JnuRWfNGKTTFsnoXMKcxG1TS76Skab2y")

			p, err := cldf_solana_provider.NewRPCChainProvider(
				d.ChainSelector,
				cldf_solana_provider.RPCChainProviderConfig{
					HTTPURL:        rpcHTTPURL,
					WSURL:          rpcWSURL,
					DeployerKeyGen: cldf_solana_provider.PrivateKeyFromRaw(deployerKey.String()),
					ProgramsPath:   programsPath,
					KeypairDirPath: programsPath, // Use the same path for keypair storage
				},
			).Initialize(context.Background())
			if err != nil {
				return nil, nil, err
			}
			client := solRpc.New(rpcHTTPURL)
			err = utils.FundSolanaAccounts(
				context.Background(),
				[]solana.PublicKey{deployerKey.PublicKey()},
				10,
				client,
			)
			if err != nil {
				return nil, nil, err
			}
			providers = append(providers, p)
		}
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
		DataStore:   runningDS.Seal(),
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
		return ccipEVM.NewEmptyCCIP16EVM(), nil
	case "solana":
		return ccipSolana.NewEmptyCCIP16Solana(), nil
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


