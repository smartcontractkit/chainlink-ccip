package ccip

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
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
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/wallet"

	chainsel "github.com/smartcontractkit/chain-selectors"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	cldf_solana_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana/provider"
	cldf_ton_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/ton/provider"
	testutils "github.com/smartcontractkit/chainlink-ton/deployment/utils"
	ccipTon "github.com/smartcontractkit/chainlink-ton/devenv-impl"

	ccipEVM "github.com/smartcontractkit/chainlink-ccip/devenv/chainimpl/ccip-evm"
	ccipSolana "github.com/smartcontractkit/chainlink-ccip/devenv/chainimpl/ccip-solana"
)

type initOptions struct {
	DataStore datastore.DataStore
}

type InitOption func(*initOptions)

func WithDataStore(ds datastore.DataStore) InitOption {
	return func(opts *initOptions) {
		opts.DataStore = ds
	}
}

type CLDF struct {
	mu        sync.Mutex          `toml:"-"`
	Addresses []string            `toml:"addresses"`
	DataStore datastore.DataStore `toml:"-"`
}

func (c *CLDF) Init(opts ...InitOption) {
	options := &initOptions{}
	for _, o := range opts {
		o(options)
	}
	if options.DataStore != nil {
		c.DataStore = options.DataStore
	} else {
		c.DataStore = datastore.NewMemoryDataStore().Seal()
	}
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

			// Use default Anvil key for local chain 1337, otherwise use PRIVATE_KEY env var
			privateKey := getNetworkPrivateKey()
			if chainID == "1337" {
				privateKey = DefaultAnvilKey
			}

			p, err := cldf_evm_provider.NewRPCChainProvider(
				d.ChainSelector,
				cldf_evm_provider.RPCChainProviderConfig{
					DeployerTransactorGen: cldf_evm_provider.TransactorFromRaw(
						privateKey,
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
					ClientOpts: []func(client *rpcclient.MultiClient){
						func(client *rpcclient.MultiClient) {
							client.RetryConfig.Timeout = 1 * time.Minute
						},
					},
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
			providers = append(providers, p)
		} else if b.Type == "ton" {
			chainID := b.ChainID
			rpcHTTPURL := b.Out.Nodes[0].ExternalHTTPUrl

			d, err := chainsel.GetChainDetailsByChainIDAndFamily(chainID, chainsel.FamilyTon)
			if err != nil {
				return nil, nil, err
			}
			client, err := testutils.CreateClient(context.Background(), rpcHTTPURL)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to create TON client: %w", err)
			}

			seed := wallet.NewSeed()
			w, err := wallet.FromSeed(client, seed, wallet.ConfigV5R1Final{NetworkGlobalID: wallet.MainnetGlobalID, Workchain: 0})
			if err != nil {
				return nil, nil, fmt.Errorf("failed to create TON wallet: %w", err)
			}
			privateKey, err := wallet.SeedToPrivateKey(seed /*password=*/, "" /*isBIP39=*/, false)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to get private key from seed: %w", err)
			}
			walletVersion := "V5R1"
			deployerSignerGen := cldf_ton_provider.PrivateKeyFromRaw(hex.EncodeToString(privateKey))

			selectors = append(selectors, d.ChainSelector)
			p, err := cldf_ton_provider.NewRPCChainProvider(
				d.ChainSelector,
				cldf_ton_provider.RPCChainProviderConfig{
					HTTPURL:           rpcHTTPURL,
					WalletVersion:     cldf_ton_provider.WalletVersion(walletVersion),
					DeployerSignerGen: deployerSignerGen,
				},
			).Initialize(context.Background())
			if err != nil {
				return nil, nil, err
			}

			err = testutils.FundWalletsNoT(client, []*address.Address{w.Address()}, []tlb.Coins{tlb.MustFromTON("1000")})
			if err != nil {
				return nil, nil, fmt.Errorf("failed to fund TON wallet: %w", err)
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
		return ccipTon.NewEmptyCCIP16TON(), nil
	default:
		return nil, errors.New("unknown devenv network type " + typ)
	}
}
