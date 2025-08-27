package ccipv1_7

/*
This code must be migrated to CLDF at some point
*/

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/link_token"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/clnode"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldfchain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldfevmprovider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	csav1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/csa"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

const (
	AnvilKey0                     = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	DefaultNativeTransferGasPrice = 21000
)

const LinkToken cldf.ContractType = "LinkToken"

var _ cldf.ChangeSet[[]uint64] = DeployLinkToken

type JobDistributor struct {
	nodev1.NodeServiceClient
	jobv1.JobServiceClient
	csav1.CSAServiceClient
	WSRPC string
}

type JDConfig struct {
	GRPC  string
	WSRPC string
}

// DeployLinkToken deploys a link token contract to the chain identified by the ChainSelector.
func DeployLinkToken(e cldf.Environment, chains []uint64) (cldf.ChangesetOutput, error) { //nolint:gocritic
	newAddresses := cldf.NewMemoryAddressBook()
	deployGrp := errgroup.Group{}
	for _, chain := range chains {
		family, err := chainsel.GetSelectorFamily(chain)
		if err != nil {
			return cldf.ChangesetOutput{AddressBook: newAddresses}, err
		}
		var deployFn func() error
		switch family {
		case chainsel.FamilyEVM:
			// Deploy EVM LINK token
			deployFn = func() error {
				_, err := deployLinkTokenContractEVM(
					e.Logger, e.BlockChains.EVMChains()[chain], newAddresses,
				)
				return err
			}
		default:
			return cldf.ChangesetOutput{}, fmt.Errorf("unsupported chain family %s", family)
		}
		deployGrp.Go(func() error {
			err := deployFn()
			if err != nil {
				e.Logger.Errorw("Failed to deploy link token", "chain", chain, "err", err)
				return fmt.Errorf("failed to deploy link token for chain %d: %w", chain, err)
			}
			return nil
		})
	}
	return cldf.ChangesetOutput{AddressBook: newAddresses}, deployGrp.Wait()
}

func deployLinkTokenContractEVM(
	lggr logger.Logger,
	chain cldfevm.Chain, //nolint:gocritic
	ab cldf.AddressBook,
) (*cldf.ContractDeploy[*link_token.LinkToken], error) {
	linkToken, err := cldf.DeployContract[*link_token.LinkToken](lggr, chain, ab,
		func(chain cldfevm.Chain) cldf.ContractDeploy[*link_token.LinkToken] {
			var (
				linkTokenAddr common.Address
				tx            *types.Transaction
				linkToken     *link_token.LinkToken
				err2          error
			)
			if !chain.IsZkSyncVM {
				linkTokenAddr, tx, linkToken, err2 = link_token.DeployLinkToken(
					chain.DeployerKey,
					chain.Client,
				)
			} else {
				linkTokenAddr, _, linkToken, err2 = link_token.DeployLinkTokenZk(
					nil,
					chain.ClientZkSyncVM,
					chain.DeployerKeyZkSyncVM,
					chain.Client,
				)
			}
			return cldf.ContractDeploy[*link_token.LinkToken]{
				Address:  linkTokenAddr,
				Contract: linkToken,
				Tx:       tx,
				Tv:       cldf.NewTypeAndVersion(LinkToken, *semver.MustParse("1.0.0")),
				Err:      err2,
			}
		})
	if err != nil {
		lggr.Errorw("Failed to deploy link token", "chain", chain.String(), "err", err)
		return linkToken, err
	}
	return linkToken, nil
}

// LoadCLDFEnvironment loads CLDF environment with a memory data store and JD client.
func LoadCLDFEnvironment(in *Cfg) (cldf.Environment, error) {
	ctx := context.Background()

	getCtx := func() context.Context {
		return ctx
	}

	// This only generates a brand new datastore and does not load any existing data.
	// We will need to figure out how data will be persisted and loaded in the future.
	ds := datastore.NewMemoryDataStore().Seal()

	lggr, err := logger.NewWith(func(config *zap.Config) {
		config.Development = true
		config.Encoding = "console"
	})
	if err != nil {
		return cldf.Environment{}, fmt.Errorf("failed to create logger: %w", err)
	}

	blockchains, err := loadCLDFChains(in.Blockchains)
	if err != nil {
		return cldf.Environment{}, fmt.Errorf("failed to load CLDF chains: %w", err)
	}

	jd, err := NewJDClient(ctx, JDConfig{
		GRPC:  in.JD.Out.ExternalGRPCUrl,
		WSRPC: in.JD.Out.ExternalWSRPCUrl,
	})
	if err != nil {
		return cldf.Environment{},
			fmt.Errorf("failed to load offchain client: %w", err)
	}

	opBundle := operations.NewBundle(
		getCtx,
		lggr,
		operations.NewMemoryReporter(),
		operations.WithOperationRegistry(operations.NewOperationRegistry()),
	)

	// This should be configurable somehow
	ocrSecrets, err := ocrSharedSecrets(
		"gesture arch depth song siege potato either tourist boss gadget split crunch sorry right funny term turtle caution another gate garbage happy swarm gossip",
		"want fire right exclude weasel make ankle fade dial impose reflect flower depart during alert taste leopard water wear wrestle assist magnet midnight",
	)
	if err != nil {
		return cldf.Environment{}, fmt.Errorf("failed to create OCR shared secrets: %w", err)
	}

	return cldf.Environment{
		Name:              "local",
		Logger:            lggr,
		ExistingAddresses: cldf.NewMemoryAddressBook(),
		DataStore:         ds,
		Offchain:          jd,
		GetContext:        getCtx,
		OCRSecrets:        ocrSecrets,
		OperationsBundle:  opBundle,
		BlockChains:       cldfchain.NewBlockChainsFromSlice(blockchains),
	}, nil
}

func loadCLDFChains(bcis []*blockchain.Input) ([]cldfchain.BlockChain, error) {
	blockchains := make([]cldfchain.BlockChain, 0)
	for _, bci := range bcis {
		switch bci.Type {
		case "anvil":
			bc, err := loadEVMChain(bci)
			if err != nil {
				return blockchains, fmt.Errorf("failed to load EVM chain %s: %w", bci.ChainID, err)
			}

			blockchains = append(blockchains, bc)
		default:
			return blockchains, fmt.Errorf("unsupported chain type %s", bci.Type)
		}
	}

	return blockchains, nil
}

func loadEVMChain(bci *blockchain.Input) (cldfchain.BlockChain, error) {
	if bci.Out == nil {
		return nil, fmt.Errorf("output configuration for %s blockchain %s is not set", bci.Type, bci.ChainID)
	}

	chainDetails, err := chainsel.GetChainDetailsByChainIDAndFamily(bci.ChainID, chainsel.FamilyEVM)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain details for %s: %w", bci.ChainID, err)
	}

	chain, err := cldfevmprovider.NewRPCChainProvider(
		chainDetails.ChainSelector,
		cldfevmprovider.RPCChainProviderConfig{
			DeployerTransactorGen: cldfevmprovider.TransactorFromRaw(
				getNetworkPrivateKey(),
			),
			RPCs: []cldf.RPC{
				{
					Name:               "default",
					WSURL:              bci.Out.Nodes[0].ExternalWSUrl,
					HTTPURL:            bci.Out.Nodes[0].ExternalHTTPUrl,
					PreferredURLScheme: cldf.URLSchemePreferenceHTTP,
				},
			},
			ConfirmFunctor: cldfevmprovider.ConfirmFuncGeth(1 * time.Minute),
		},
	).Initialize(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize EVM chain %s: %w", bci.ChainID, err)
	}

	return chain, nil
}

func ocrSharedSecrets(xSigners, xProposers string) (cldf.OCRSecrets, error) {
	if xSigners == "" || xProposers == "" {
		return cldf.OCRSecrets{}, errors.New("xsigners or xproposers not set")
	}

	// Lifted from here
	// https://github.com/smartcontractkit/offchain-reporting/blob/14a57d70e50474a2104aa413214e464d6bc69e16/lib/offchainreporting/internal/config/shared_secret_test.go#L32
	// Historically signers (fixed secret) and proposers (ephemeral secret) were
	// combined in this manner. We simply leave that as is.
	xSignersHash := crypto.Keccak256([]byte(xSigners))
	xProposersHash := crypto.Keccak256([]byte(xProposers))
	xSignersHashxProposersHashZero := append(append(append([]byte{}, xSignersHash...), xProposersHash...), 0)
	xSignersHashxProposersHashOne := append(append(append([]byte{}, xSignersHash...), xProposersHash...), 1)
	var sharedSecret [16]byte
	copy(sharedSecret[:], crypto.Keccak256(xSignersHashxProposersHashZero))
	var sk [32]byte
	copy(sk[:], crypto.Keccak256(xSignersHashxProposersHashOne))

	return cldf.OCRSecrets{
		SharedSecret: sharedSecret,
		EphemeralSk:  sk,
	}, nil
}

// OffchainClient implementation. This will eventually be moved to CLDF.

// NewJDClient creates a new JobDistributor client.
func NewJDClient(ctx context.Context, cfg JDConfig) (cldf.OffchainClient, error) {
	conn, err := NewJDConnection(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}
	jd := &JobDistributor{
		WSRPC:             cfg.WSRPC,
		NodeServiceClient: nodev1.NewNodeServiceClient(conn),
		JobServiceClient:  jobv1.NewJobServiceClient(conn),
		CSAServiceClient:  csav1.NewCSAServiceClient(conn),
	}

	return jd, err
}

func (jd JobDistributor) GetCSAPublicKey(ctx context.Context) (string, error) {
	keypairs, err := jd.ListKeypairs(ctx, &csav1.ListKeypairsRequest{})
	if err != nil {
		return "", err
	}
	if keypairs == nil || len(keypairs.Keypairs) == 0 {
		return "", errors.New("no keypairs found")
	}
	csakey := keypairs.Keypairs[0].PublicKey
	return csakey, nil
}

// ProposeJob proposes jobs through the jobService and accepts the proposed job on selected node based on ProposeJobRequest.NodeId.
func (jd JobDistributor) ProposeJob(ctx context.Context, in *jobv1.ProposeJobRequest, opts ...grpc.CallOption) (*jobv1.ProposeJobResponse, error) {
	res, err := jd.JobServiceClient.ProposeJob(ctx, in, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to propose job. err: %w", err)
	}
	if res.Proposal == nil {
		return nil, errors.New("failed to propose job. err: proposal is nil")
	}

	return res, nil
}

// NewJDConnection creates new gRPC connection with JobDistributor.
func NewJDConnection(cfg JDConfig) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	interceptors := []grpc.UnaryClientInterceptor{}

	if len(interceptors) > 0 {
		opts = append(opts, grpc.WithChainUnaryInterceptor(interceptors...))
	}

	conn, err := grpc.NewClient(cfg.GRPC, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}

	return conn, nil
}

// FundNodeEIP1559 funds CL node using RPC URL, recipient address and amount of funds to send (ETH).
// Uses EIP-1559 transaction type.
func FundNodeEIP1559(c *ethclient.Client, pkey, recipientAddress string, amountOfFundsInETH float64) error {
	amount := new(big.Float).Mul(big.NewFloat(amountOfFundsInETH), big.NewFloat(1e18))
	amountWei, _ := amount.Int(nil)
	Plog.Info().Str("Addr", recipientAddress).Str("Wei", amountWei.String()).Msg("Funding Node")

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return err
	}
	privateKeyStr := strings.TrimPrefix(pkey, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}
	feeCap, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	tipCap, err := c.SuggestGasTipCap(context.Background())
	if err != nil {
		return err
	}
	recipient := common.HexToAddress(recipientAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &recipient,
		Value:     amountWei,
		Gas:       DefaultNativeTransferGasPrice,
		GasFeeCap: feeCap,
		GasTipCap: tipCap,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return err
	}
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(context.Background(), c, signedTx); err != nil {
		return err
	}
	Plog.Info().Str("Wei", amountWei.String()).Msg("Funded with ETH")
	return nil
}

// ExampleCLDFDeployment demonstrates how to use CLDF with a single LINK token deployment.
func ExampleCLDFDeployment(in *Cfg) error {
	Plog.Info().Msg("[Data Feeds] Loading CLDF environment")
	env, err := LoadCLDFEnvironment(in)
	if err != nil {
		return fmt.Errorf("failed to load CLDF environment: %w", err)
	}

	Plog.Info().Msg("[Data Feeds] Applying an example product configuration 🚀")
	Plog.Info().Msg("[Data Feeds] Deploying contracts..")
	chainDetails, err := chainsel.GetChainDetailsByChainIDAndFamily(
		in.Blockchains[0].ChainID, chainsel.FamilyEVM,
	)
	if err != nil {
		return fmt.Errorf("failed to get chain details for: %w", err)
	}

	_, err = DeployLinkToken(env, []uint64{
		chainDetails.ChainSelector,
	})
	if err != nil {
		return fmt.Errorf("failed to deploy link token: %w", err)
	}
	Plog.Info().Msg("[Data Feeds] Generating CL nodes network configuration")
	Plog.Info().Msg("[Data Feeds] Creating config for Source chain")
	srcNetCfg, err := clnode.NewNetworkCfg(&clnode.EVMNetworkConfig{
		MinIncomingConfirmations: 1,
		MinContractPayment:       "0.00001 link",
		ChainID:                  in.Blockchains[0].Out.ChainID,
		EVMNodes: []*clnode.EVMNode{
			{
				Name:     "src",
				SendOnly: false,
				Order:    100,
			},
		},
	}, in.Blockchains[0].Out)
	if err != nil {
		return fmt.Errorf("failed to configure network 1337: %w", err)
	}
	Plog.Info().Msg("[Data Feeds] Creating config for Destination chain")
	dstNetCfg, err := clnode.NewNetworkCfg(&clnode.EVMNetworkConfig{
		MinIncomingConfirmations: 1,
		MinContractPayment:       "0.00001 link",
		ChainID:                  in.Blockchains[1].Out.ChainID,
		EVMNodes: []*clnode.EVMNode{
			{
				Name:     "dst",
				SendOnly: false,
				Order:    100,
			},
		},
	}, in.Blockchains[1].Out)
	if err != nil {
		return fmt.Errorf("failed to configure network 2337: %w", err)
	}
	Plog.Info().Msg("[Data Feeds] Applying default CL nodes configuration")
	// configure node set and generate CL nodes configs
	netConfig := srcNetCfg + dstNetCfg + `
	   [Log]
	   JSONConsole = true
	   Level = 'debug'
	   [Pyroscope]
	   ServerAddress = 'http://host.docker.internal:4040'
	   Environment = 'local'
	   [WebServer]
	   HTTPWriteTimeout = '30s'
	   SecureCookies = false
	   Port = 6688
	   [WebServer.TLS]
	   HTTPSPort = 0
	   [JobPipeline]
	   [JobPipeline.HTTPRequest]
	   DefaultTimeout = '30s'
	`
	for _, nodeSpec := range in.NodeSets[0].NodeSpecs {
		nodeSpec.Node.TestConfigOverrides = netConfig
	}
	Plog.Info().Msg("[Data Feeds] Product configuration is finished 🎉")
	return nil
}

/*
This is just a basic ETH client, CLDF should provide something like this
*/

// ETHClient creates a basic Ethereum client using PRIVATE_KEY env var and tip/cap gas settings.
func ETHClient(wsURL string, gasSettings *GasSettings) (*ethclient.Client, *bind.TransactOpts, string, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not connect to eth client: %w", err)
	}
	privateKey, err := crypto.HexToECDSA(getNetworkPrivateKey())
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not parse private key: %w", err)
	}
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey).String()
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not get chain ID: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not create transactor: %w", err)
	}
	fc, tc, err := MultiplyEIP1559GasPrices(client, gasSettings.FeeCapMultiplier, gasSettings.TipCapMultiplier)
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not get bumped gas price: %w", err)
	}
	auth.GasFeeCap = fc
	auth.GasTipCap = tc
	Plog.Info().
		Str("GasFeeCap", fc.String()).
		Str("GasTipCap", tc.String()).
		Msg("Default gas prices set")
	return client, auth, address, nil
}

// MultiplyEIP1559GasPrices returns bumped EIP1159 gas prices increased by multiplier.
func MultiplyEIP1559GasPrices(client *ethclient.Client, fcMult, tcMult int64) (*big.Int, *big.Int, error) {
	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}
	tipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, err
	}

	return new(big.Int).Mul(feeCap, big.NewInt(fcMult)), new(big.Int).Mul(tipCap, big.NewInt(tcMult)), nil
}
