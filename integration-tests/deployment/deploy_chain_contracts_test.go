package deployment

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	mcmsops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/mcms"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	mcmsapi "github.com/smartcontractkit/chainlink-ccip/deployment/v1_0"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts_Apply(t *testing.T) {
	t.Parallel()
	programsPath, ds, err := PreloadSolanaEnvironment(chain_selectors.SOLANA_MAINNET.Selector)
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
	_, err = mcmsapi.DeployContracts(dReg).Apply(*e, mcmsapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]mcmsapi.ContractDeploymentConfigPerChain{
			chain_selectors.SOLANA_MAINNET.Selector: {
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

	cs := mcmsapi.DeployMCMS(dReg)
	output, err := cs.Apply(*e, mcmsapi.MCMSDeploymentConfig{
		Chains: map[uint64]mcmsapi.MCMSDeploymentConfigPerChain{
			chain_selectors.SOLANA_MAINNET.Selector: {
				Canceller:        mcmsapi.SingleGroupMCMSV2(),
				Bypasser:         mcmsapi.SingleGroupMCMSV2(),
				Proposer:         mcmsapi.SingleGroupMCMSV2(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(common_utils.CLLQualifier),
			},
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	e.DataStore = output.DataStore.Seal()
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
	"mcm":               datastore.ContractType(mcmsops.McmProgramType),
	"timelock":          datastore.ContractType(mcmsops.TimelockProgramType),
	"access_controller": datastore.ContractType(mcmsops.AccessControllerProgramType),
}

func PreloadSolanaEnvironment(chainSelector uint64) (string, *datastore.MemoryDataStore, error) {
	programsPath := os.TempDir()
	ds := datastore.NewMemoryDataStore()
	err := utils.DownloadSolanaCCIPProgramArtifacts(context.Background(), programsPath, utils.VersionToShortCommitSHA[utils.VersionSolanaV0_1_1])
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
