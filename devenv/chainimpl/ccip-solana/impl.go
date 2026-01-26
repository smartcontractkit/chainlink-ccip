package ccip_solana

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip/consts"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
	chainsel "github.com/smartcontractkit/chain-selectors"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	solanaseqs "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	solCommonUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
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

type CCIP16Solana struct {
	e                   *deployment.Environment
	chainDetails        chainsel.ChainDetails
	ExpectedSeqNumRange map[SourceDestPair]ccipocr3common.SeqNumRange
	ExpectedSeqNumExec  map[SourceDestPair][]uint64
	MsgSentEvents       []*AnyMsgSentEvent
	testadapters.TestAdapter
}

func NewEmptyCCIP16Solana(chainDetails chainsel.ChainDetails) *CCIP16Solana {
	return &CCIP16Solana{
		chainDetails:        chainDetails,
		ExpectedSeqNumRange: make(map[SourceDestPair]ccipocr3common.SeqNumRange),
		ExpectedSeqNumExec:  make(map[SourceDestPair][]uint64),
		MsgSentEvents:       make([]*AnyMsgSentEvent, 0),
	}
}

// NewCCIP16Solana creates new smart-contracts wrappers with utility functions for CCIP16Solana implementation.
func NewCCIP16Solana(ctx context.Context, e *deployment.Environment, chainDetails chainsel.ChainDetails) (*CCIP16Solana, error) {
	_ = zerolog.Ctx(ctx)
	out := NewEmptyCCIP16Solana(chainDetails)
	out.SetCLDF(e)
	return out, nil
}

func (m *CCIP16Solana) SetCLDF(e *deployment.Environment) {
	m.e = e
	factory, found := testadapters.GetTestAdapterRegistry().GetTestAdapter(chain_selectors.FamilySolana, semver.MustParse("1.6.0"))
	if !found {
		panic(fmt.Sprintf("failed to find testadapter factory for %s", chain_selectors.FamilySolana))
	}
	m.TestAdapter = factory(e, m.chainDetails.ChainSelector)
}

func (m *CCIP16Solana) ChainSelector() uint64 {
	return m.chainDetails.ChainSelector
}

func updatePrices(ds datastore.DataStore, src, dest uint64, srcChain cldf_solana.Chain) error {
	a := &solanaseqs.SolanaAdapter{}
	fqAddr, err := a.GetFQAddress(ds, src)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	linkAddr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: src,
		Type:          datastore.ContractType("LINK"),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to get LINK address: %w", err)
	}
	fqID := solana.PublicKeyFromBytes(fqAddr)
	fee_quoter.SetProgramID(solana.PublicKeyFromBytes(fqAddr))
	fqAllowedPriceUpdaterPDA, _, _ := state.FindFqAllowedPriceUpdaterPDA(srcChain.DeployerKey.PublicKey(), fqID)
	feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(fqID)
	ixn := fee_quoter.NewUpdatePricesInstruction(
		[]fee_quoter.TokenPriceUpdate{
			{
				SourceToken: solana.WrappedSol,
				UsdPerToken: [28]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 117, 223, 115, 252, 225, 110, 70, 208, 64, 0, 0},
			},
			{
				SourceToken: solana.MustPublicKeyFromBase58(linkAddr.Address),
				UsdPerToken: [28]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 117, 223, 115, 252, 225, 110, 70, 208, 64, 0, 0},
			},
		},
		[]fee_quoter.GasPriceUpdate{
			{
				DestChainSelector: dest,
				UsdPerUnitGas:     solCommonUtil.To28BytesBE(1),
			},
		},
		srcChain.DeployerKey.PublicKey(),
		fqAllowedPriceUpdaterPDA,
		feeQuoterConfigPDA,
	)
	billingTokenConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(solana.WrappedSol, fqID)
	ixn.Append(solana.Meta(billingTokenConfigPDA).WRITE())
	billingTokenConfigPDA, _, _ = state.FindFqBillingTokenConfigPDA(solana.MustPublicKeyFromBase58(linkAddr.Address), fqID)
	ixn.Append(solana.Meta(billingTokenConfigPDA).WRITE())
	fqDestPDA, _, _ := state.FindFqDestChainPDA(dest, fqID)
	ixn.Append(solana.Meta(fqDestPDA).WRITE())
	ix, err := ixn.ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to validate and build instruction: %w", err)
	}
	err = srcChain.Confirm([]solana.Instruction{ix})
	if err != nil {
		return fmt.Errorf("failed to confirm get fee tokens transaction: %w", err)
	}
	return nil
}

func (m *CCIP16Solana) GetEOAReceiverAddress(ctx context.Context, chainSelector uint64) ([]byte, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

func (m *CCIP16Solana) GetTokenBalance(ctx context.Context, chainSelector uint64, address, tokenAddress []byte) (*big.Int, error) {
	_ = zerolog.Ctx(ctx)
	return big.NewInt(0), nil
}

func (m *CCIP16Solana) ExposeMetrics(
	ctx context.Context,
	source, dest uint64,
	chainIDs []string,
	wsURLs []string,
) ([]string, *prometheus.Registry, error) {
	return nil, nil, nil
}

func (m *CCIP16Solana) DeployLocalNetwork(ctx context.Context, bc *blockchain.Input) (*blockchain.Output, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Deploying Solana networks")
	d, err := chainsel.GetChainDetailsByChainIDAndFamily(bc.ChainID, chainsel.FamilySolana)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain details: %w", err)
	}
	err = PreloadSolanaEnvironment(bc.ContractsDir, d.ChainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to preload Solana environment: %w", err)
	}
	bc.SolanaPrograms = solanaProgramIDs
	out, err := blockchain.NewBlockchainNetwork(bc)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain network: %w", err)
	}
	return out, nil
}

var solanaProgramIDs = map[string]string{
	"ccip_router":               "Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C",
	"test_token_pool":           "JuCcZ4smxAYv9QHJ36jshA7pA3FuQ3vQeWLUeAtZduJ",
	"burnmint_token_pool":       "41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB",
	"lockrelease_token_pool":    "8eqh8wppT9c5rw4ERqNCffvU6cNFJWff9WmkcYtmGiqC",
	"fee_quoter":                "FeeQPGkKDeRV1MgoYfMH6L8o3KeuYjwUZrgn4LRKfjHi",
	"test_ccip_receiver":        "EvhgrPhTDt4LcSPS2kfJgH6T6XWZ6wT3X9ncDGLT1vui",
	"ccip_offramp":              "offqSMQWgQud6WJz694LRzkeN5kMYpCHTpXQr3Rkcjm",
	"mcm":                       "5vNJx78mz7KVMjhuipyr9jKBKcMrKYGdjGkgE4LUmjKk",
	"timelock":                  "DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA",
	"access_controller":         "6KsN58MTnRQ8FfPaXHiFPPFGDRioikj9CdPvPxZJdCjb",
	"external_program_cpi_stub": "2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm",
	"rmn_remote":                "RmnXLft1mSEwDgMKu2okYuHkiazxntFFcZFrrcXxYg7",
	"cctp_token_pool":           "CCiTPESGEevd7TBU8EGBKrcxuRq7jx3YtW6tPidnscaZ",
}

var solanaContracts = map[string]datastore.ContractType{
	"ccip_router":            datastore.ContractType(routerops.ContractType),
	"fee_quoter":             datastore.ContractType(fqops.ContractType),
	"ccip_offramp":           datastore.ContractType(offrampops.ContractType),
	"rmn_remote":             datastore.ContractType(rmnremoteops.ContractType),
	"mcm":                    datastore.ContractType(utils.McmProgramType),
	"timelock":               datastore.ContractType(utils.TimelockProgramType),
	"access_controller":      datastore.ContractType(utils.AccessControllerProgramType),
	"burnmint_token_pool":    datastore.ContractType(common_utils.BurnMintTokenPool),
	"lockrelease_token_pool": datastore.ContractType(common_utils.LockReleaseTokenPool),
	"test_ccip_receiver":     datastore.ContractType("TestReceiver"),
}

func PreloadSolanaEnvironment(programsPath string, chainSelector uint64) error {
	err := utils.DownloadSolanaCCIPProgramArtifacts(context.Background(), programsPath, utils.VersionToShortCommitSHA[utils.VersionSolanaV0_1_1])
	if err != nil {
		return err
	}
	return nil
}

// Populates datastore with the predeployed program addresses
// pass map [programName]:ContractType of contracts to populate datastore with
func populateDatastore(ds *datastore.MemoryAddressRefStore, contracts map[string]datastore.ContractType, version *semver.Version, qualifier string, chainSel uint64) error {
	for programName, programID := range solanaProgramIDs {
		ct, ok := contracts[programName]
		if !ok {
			continue
		}

		err := ds.Add(datastore.AddressRef{
			Address:       programID,
			ChainSelector: chainSel,
			Qualifier:     qualifier,
			Type:          ct,
			Version:       version,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *CCIP16Solana) ConfigureNodes(ctx context.Context, bc *blockchain.Input) (string, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring CL nodes for solana")
	name := fmt.Sprintf("node-solana-%s", uuid.New().String()[0:5])
	return fmt.Sprintf(`
	[[Solana]]
	ChainID = '%s'
	Enabled = true
	TxRetryTimeout = '90s'
	TxConfirmTimeout = '5m'
	TxExpirationRebroadcast = true
	TxRetentionTimeout = '8h'
	FeeBumpPeriod = '0s'
	FeeEstimatorMode = 'blockhistory'
	ComputeUnitPriceMin = 1
	ComputeUnitPriceMax = 500000
	BlockHistorySize = 150
	LogPollerStartingLookback = '24h'

	[Solana.MultiNode]
	VerifyChainID = false
	Enabled = true
	SyncThreshold = 170
	PollFailureThreshold = 5
	PollInterval = "15s"
	NewHeadsPollInterval = "5s"
	SelectionMode = "PriorityLevel"
	LeaseDuration = "1m"
	NodeIsSyncingEnabled = false
	FinalizedBlockPollInterval = "5s"
	EnforceRepeatableRead = true
	DeathDeclarationDelay = "20s"
	NodeNoNewHeadsThreshold = "20s"
	NoNewFinalizedHeadsThreshold = "20s"
	FinalityTagEnabled = true
	FinalityDepth = 0
	FinalizedBlockOffset = 50

	[[Solana.Nodes]]
	Name = '%s'
	URL = '%s'`,
		bc.ChainID,
		name,
		bc.Out.Nodes[0].InternalHTTPUrl,
	), nil
}

func (m *CCIP16Solana) PreDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	ds := datastore.NewMemoryDataStore()
	ds.Merge(env.DataStore)
	err := populateDatastore(ds.AddressRefStore, solanaContracts, semver.MustParse("1.6.0"), "", selector)
	if err != nil {
		return err
	}
	env.DataStore = ds.Seal()
	return nil
}

func (m *CCIP16Solana) PostDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	// post contract setup for testing
	a := &solanaseqs.SolanaAdapter{}
	linkAddr, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType("LINK"),
	}, selector, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to get LINK address: %w", err)
	}
	fqAddr, err := a.GetFQAddress(env.DataStore, selector)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	routerAddr, err := a.GetRouterAddress(env.DataStore, selector)
	if err != nil {
		return fmt.Errorf("failed to get router address: %w", err)
	}
	chain := env.BlockChains.SolanaChains()[selector]
	fqID := solana.PublicKeyFromBytes(fqAddr)
	routerID := solana.PublicKeyFromBytes(routerAddr)
	fee_quoter.SetProgramID(solana.PublicKeyFromBytes(fqAddr))
	fqAllowedPriceUpdaterPDA, _, _ := state.FindFqAllowedPriceUpdaterPDA(chain.DeployerKey.PublicKey(), fqID)
	feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(fqID)
	ixn, err := fee_quoter.NewAddPriceUpdaterInstruction(
		chain.DeployerKey.PublicKey(),
		fqAllowedPriceUpdaterPDA,
		feeQuoterConfigPDA,
		chain.DeployerKey.PublicKey(),
		solana.SystemProgramID,
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to get fee tokens from fee quoter: %w", err)
	}
	err = chain.Confirm([]solana.Instruction{ixn})
	if err != nil {
		return fmt.Errorf("failed to add price updater: %w", err)
	}
	for _, tokenPubKey := range []solana.PublicKey{
		solana.WrappedSol,
		solana.MustPublicKeyFromBase58(linkAddr.Address),
	} {
		billingTokenConfig := fee_quoter.BillingTokenConfig{
			Enabled: true,
			Mint:    tokenPubKey,
			UsdPerToken: fee_quoter.TimestampedPackedU224{
				Timestamp: time.Now().Unix(),
				Value:     [28]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 117, 223, 115, 252, 225, 110, 70, 208, 64, 0, 0},
			},
			PremiumMultiplierWeiPerEth: 1e18,
		}
		tokenBillingPDA, _, _ := state.FindFqBillingTokenConfigPDA(tokenPubKey, fqID)
		// we dont need to handle test router here because we explicitly create this and token Receiver for test router
		billingSignerPDA, _, _ := state.FindFeeBillingSignerPDA(routerID)
		tokenProgramID := solana.TokenProgramID
		tokenReceiver, _, _ := tokens.FindAssociatedTokenAddress(tokenProgramID, tokenPubKey, billingSignerPDA)
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(fqID)
		ixn, err := fee_quoter.NewAddBillingTokenConfigInstruction(
			billingTokenConfig,
			feeQuoterConfigPDA,
			tokenBillingPDA,
			tokenProgramID,
			tokenPubKey,
			tokenReceiver,
			chain.DeployerKey.PublicKey(),
			billingSignerPDA,
			ata.ProgramID,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return fmt.Errorf("failed to build add billing token config instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return fmt.Errorf("failed to add price updater: %w", err)
		}
	}
	return nil
}

func (m *CCIP16Solana) FundNodes(ctx context.Context, ns []*simple_node_set.Input, nodeKeyBundles map[string]clclient.NodeKeysBundle, bc *blockchain.Input, linkAmount, nativeAmount *big.Int) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Funding CL nodes with ETH and LINK")
	var keys []solana.PublicKey
	for _, nk := range nodeKeyBundles {
		keys = append(keys, solana.MustPublicKeyFromBase58(nk.TXKey.Data.Attributes.PublicKey))
	}
	client := solRpc.New(bc.Out.Nodes[0].ExternalHTTPUrl)
	return utils.FundSolanaAccounts(
		ctx,
		keys,
		10,
		client,
	)
}
