package deployment

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	mcmsapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestDeployChainContracts_Apply(t *testing.T) {
	t.Parallel()
	programsPath, ds, err := PreloadSolanaEnvironment(t, chain_selectors.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, ds, "Datastore should be created")

	e, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{chain_selectors.SOLANA_MAINNET.Selector}, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	e.DataStore = ds.Seal() // Add preloaded contracts to env datastore
	mint, _ := solana.NewRandomPrivateKey()

	dReg := mcmsapi.GetRegistry()
	version := semver.MustParse("1.6.0")
	_, err = mcmsapi.DeployContracts(dReg).Apply(*e, mcmsapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]mcmsapi.ContractDeploymentConfigPerChain{
			chain_selectors.SOLANA_MAINNET.Selector: {
				Version: version,
				// LINK TOKEN CONFIG
				// token private key used to deploy the LINK token. Solana: base58 encoded private key
				TokenPrivKey: mint.String(),
				// token decimals used to deploy the LINK token
				TokenDecimals: 9,
				// FEE QUOTER CONFIG
				MaxFeeJuelsPerMsg: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				// OFFRAMP CONFIG
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			},
		},
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")

	DeployMCMS(t, e, chain_selectors.SOLANA_MAINNET.Selector, []string{common_utils.CLLQualifier, common_utils.RMNTimelockQualifier})
	SolanaTransferMCMSContracts(t, e, chain_selectors.SOLANA_MAINNET.Selector, common_utils.CLLQualifier)
	SolanaTransferOwnership(t, e, chain_selectors.SOLANA_MAINNET.Selector)
	// Transfer from one qualifier to another
	fromQualifier := common_utils.CLLQualifier
	toQualifier := common_utils.RMNTimelockQualifier
	chain := e.BlockChains.SolanaChains()[chain_selectors.SOLANA_MAINNET.Selector]
	fromTimelockSigner := utils.GetTimelockSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		fromQualifier)
	toTimelockSigner := utils.GetTimelockSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		toQualifier)
	toMcmSigner := utils.GetMCMSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.ProposerManyChainMultisig,
		toQualifier,
	)
	err = utils.FundSolanaAccounts(
		context.Background(),
		[]solana.PublicKey{toTimelockSigner, toMcmSigner},
		100,
		chain.Client,
	)
	require.NoError(t, err)
	mcmsInput := mcmsapi.TransferOwnershipInput{
		ChainInputs: []mcmsapi.TransferOwnershipPerChainInput{
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(routerops.ContractType),
						Version: routerops.Version,
					},
				},
				CurrentOwner:  fromTimelockSigner.String(),
				ProposedOwner: toTimelockSigner.String(),
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(offrampops.ContractType),
						Version: offrampops.Version,
					},
				},
				CurrentOwner:  fromTimelockSigner.String(),
				ProposedOwner: toTimelockSigner.String(),
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(fqops.ContractType),
						Version: fqops.Version,
					},
				},
				CurrentOwner:  fromTimelockSigner.String(),
				ProposedOwner: toTimelockSigner.String(),
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(rmnremoteops.ContractType),
						Version: rmnremoteops.Version,
					},
				},
				CurrentOwner:  fromTimelockSigner.String(),
				ProposedOwner: toTimelockSigner.String(),
			},
		},
		AdapterVersion: semver.MustParse("1.6.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            fromQualifier,
			Description:          "Transfer ownership test",
		},
	}

	transferOutput, err := mcmsapi.TransferOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInput)
	require.NoError(t, err)
	require.Greater(t, len(transferOutput.Reports), 0)
	require.Equal(t, 1, len(transferOutput.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, transferOutput.MCMSTimelockProposals, false)

	mcmsInput.MCMS.Qualifier = toQualifier

	acceptOutput, err := mcmsapi.AcceptOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInput)
	require.NoError(t, err)
	require.Greater(t, len(acceptOutput.Reports), 0)
	require.Equal(t, 1, len(acceptOutput.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *e, acceptOutput.MCMSTimelockProposals, false)
	// router verify
	program := datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		routerops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, toTimelockSigner, routerops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
	// offramp verify
	program = datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		offrampops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, toTimelockSigner, offrampops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
	// fee quoter verify
	program = datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		fqops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, toTimelockSigner, fqops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
	// rmn remote verify
	program = datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		rmnremoteops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, toTimelockSigner, rmnremoteops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
}

var solanaProgramIDs = map[string]string{
	"ccip_router": "Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C",
	// "test_token_pool":           "JuCcZ4smxAYv9QHJ36jshA7pA3FuQ3vQeWLUeAtZduJ",
	// "burnmint_token_pool":       "41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB",
	// "lockrelease_token_pool":    "8eqh8wppT9c5rw4ERqNCffvU6cNFJWff9WmkcYtmGiqC",
	"fee_quoter": "FeeQPGkKDeRV1MgoYfMH6L8o3KeuYjwUZrgn4LRKfjHi",
	// "test_ccip_receiver":        "EvhgrPhTDt4LcSPS2kfJgH6T6XWZ6wT3X9ncDGLT1vui",
	"ccip_offramp":      "offqSMQWgQud6WJz694LRzkeN5kMYpCHTpXQr3Rkcjm",
	"mcm":               "5vNJx78mz7KVMjhuipyr9jKBKcMrKYGdjGkgE4LUmjKk",
	"timelock":          "DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA",
	"access_controller": "6KsN58MTnRQ8FfPaXHiFPPFGDRioikj9CdPvPxZJdCjb",
	// "external_program_cpi_stub": "2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm",
	"rmn_remote": "RmnXLft1mSEwDgMKu2okYuHkiazxntFFcZFrrcXxYg7",
	// "cctp_token_pool":           "CCiTPESGEevd7TBU8EGBKrcxuRq7jx3YtW6tPidnscaZ",
}

var solanaContracts = map[string]datastore.ContractType{
	"ccip_router":       datastore.ContractType(routerops.ContractType),
	"fee_quoter":        datastore.ContractType(fqops.ContractType),
	"ccip_offramp":      datastore.ContractType(offrampops.ContractType),
	"rmn_remote":        datastore.ContractType(rmnremoteops.ContractType),
	"mcm":               datastore.ContractType(utils.McmProgramType),
	"timelock":          datastore.ContractType(utils.TimelockProgramType),
	"access_controller": datastore.ContractType(utils.AccessControllerProgramType),
}

func PreloadSolanaEnvironment(t *testing.T, chainSelector uint64) (string, *datastore.MemoryDataStore, error) {
	programsPath := t.TempDir()
	ds := datastore.NewMemoryDataStore()
	err := utils.DownloadSolanaCCIPProgramArtifacts(t.Context(), programsPath, utils.VersionToShortCommitSHA[utils.VersionSolanaV0_1_1])
	if err != nil {
		return "", nil, err
	}
	err = populateDatastore(ds.AddressRefStore, solanaContracts, semver.MustParse("1.6.0"), "", chainSelector)
	if err != nil {
		return "", nil, err
	}
	return programsPath, ds, nil
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
