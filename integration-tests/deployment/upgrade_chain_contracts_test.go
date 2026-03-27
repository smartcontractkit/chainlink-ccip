package deployment

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	solutils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func skipInCI(t *testing.T) {
	t.Helper()
	if os.Getenv("CI") == "true" {
		t.Skip("Skipping in CI — local build is expensive")
	}
}

// TestDeployWithLocalBuild tests the full artifact build flow by building
// Solana programs locally (clone, key replacement, anchor build) and deploying
// them through the unified DeployContracts changeset.
// This is skipped in CI because the local docker build is expensive.
func TestDeployWithLocalBuild(t *testing.T) {
	t.Parallel()
	skipInCI(t)

	solSelector := chain_selectors.SOLANA_MAINNET.Selector
	programsPath := t.TempDir()

	// We don't preload programs — the ArtifactPreparer will build them
	e, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solSelector}, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)

	version := semver.MustParse("1.6.0")
	mint, _ := solana.NewRandomPrivateKey()

	dReg := deployapi.GetRegistry()
	_, err = deployapi.DeployContracts(dReg).Apply(*e, deployapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			solSelector: {
				Version:       version,
				TokenPrivKey:  mint.String(),
				TokenDecimals: 9,
				MaxFeeJuelsPerMsg: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
				ChainSpecific: &solutils.SolanaBuildConfig{
					ContractVersion: solutils.VersionSolanaV1_6_0,
					DestinationDir:  programsPath,
					LocalBuild: &solutils.LocalBuildConfig{
						BuildLocally:        true,
						CleanDestinationDir: true,
						GenerateVanityKeys:  true,
					},
				},
			},
		},
	})
	require.NoError(t, err)

	// Verify deployed contracts exist in datastore
	addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solSelector))
	require.NotEmpty(t, addresses)

	foundRouter := false
	foundFeeQuoter := false
	foundOffRamp := false
	for _, addr := range addresses {
		switch cldf.ContractType(addr.Type) {
		case routerops.ContractType:
			foundRouter = true
		case fqops.ContractType:
			foundFeeQuoter = true
		case offrampops.ContractType:
			foundOffRamp = true
		}
	}
	require.True(t, foundRouter, "Router should be deployed")
	require.True(t, foundFeeQuoter, "FeeQuoter should be deployed")
	require.True(t, foundOffRamp, "OffRamp should be deployed")
}

// TestUpgradeChainContracts tests the full upgrade flow:
// 1. Build and deploy original contracts with vanity keys
// 2. Deploy MCMS and transfer ownership
// 3. Build upgraded contracts with matching keys (declare_id replacement)
// 4. Upgrade programs in place via the UpgradeContracts changeset
// 5. Verify programs were upgraded in place (same addresses)
//
// Skipped in CI because the two local docker builds are expensive.
func TestUpgradeChainContracts(t *testing.T) {
	t.Parallel()
	skipInCI(t)

	solSelector := chain_selectors.SOLANA_MAINNET.Selector
	programsPath := t.TempDir()

	e, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solSelector}, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)

	version := semver.MustParse("1.6.0")
	mint, _ := solana.NewRandomPrivateKey()

	// Step 1: Initial deploy with local build + vanity keys
	dReg := deployapi.GetRegistry()
	_, err = deployapi.DeployContracts(dReg).Apply(*e, deployapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			solSelector: {
				Version:       version,
				TokenPrivKey:  mint.String(),
				TokenDecimals: 9,
				MaxFeeJuelsPerMsg: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
				ChainSpecific: &solutils.SolanaBuildConfig{
					ContractVersion: solutils.VersionSolanaV1_6_0,
					DestinationDir:  programsPath,
					LocalBuild: &solutils.LocalBuildConfig{
						BuildLocally:        true,
						CleanDestinationDir: true,
						GenerateVanityKeys:  true,
					},
				},
			},
		},
	})
	require.NoError(t, err)

	// Step 2: Deploy MCMS and transfer ownership
	DeployMCMS(t, e, solSelector, []string{common_utils.CLLQualifier})
	SolanaTransferMCMSContracts(t, e, solSelector, common_utils.CLLQualifier, true)
	SolanaTransferOwnership(t, e, solSelector)

	// Capture deployed program addresses before upgrade
	chain := e.BlockChains.SolanaChains()[solSelector]
	allAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solSelector))
	preUpgradeAddresses := make(map[cldf.ContractType]string)
	for _, addr := range allAddresses {
		preUpgradeAddresses[cldf.ContractType(addr.Type)] = addr.Address
	}

	routerAddr := preUpgradeAddresses[routerops.ContractType]
	fqAddr := preUpgradeAddresses[fqops.ContractType]
	offRampAddr := preUpgradeAddresses[offrampops.ContractType]
	require.NotEmpty(t, routerAddr, "Router address should exist")
	require.NotEmpty(t, fqAddr, "FeeQuoter address should exist")
	require.NotEmpty(t, offRampAddr, "OffRamp address should exist")

	// Step 3: Build upgrade artifacts with keys matching deployed programs
	upgradeVersion := semver.MustParse("1.6.1")
	upgradeAuthority := chain.DeployerKey.PublicKey()

	uReg := deployapi.GetUpgraderRegistry()
	_, err = deployapi.UpgradeContracts(uReg, dReg).Apply(*e, deployapi.ContractUpgradeConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractUpgradeConfigPerChain{
			solSelector: {
				Version: version,
				Upgrades: map[cldf.ContractType]*semver.Version{
					routerops.ContractType:  upgradeVersion,
					fqops.ContractType:      upgradeVersion,
					offrampops.ContractType: upgradeVersion,
				},
				UpgradeAuthority: upgradeAuthority.String(),
				ChainSpecific: &solutils.SolanaBuildConfig{
					ContractVersion: solutils.VersionSolanaV1_6_1,
					DestinationDir:  programsPath,
					LocalBuild: &solutils.LocalBuildConfig{
						BuildLocally:        true,
						CleanDestinationDir: true,
						CleanGitDir:         true,
						UpgradeKeys: map[cldf.ContractType]string{
							routerops.ContractType:  routerAddr,
							fqops.ContractType:      fqAddr,
							offrampops.ContractType: offRampAddr,
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)

	// Step 4: Verify programs were upgraded in place (same addresses)
	postUpgradeAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solSelector))

	routerCount := 0
	fqCount := 0
	offRampCount := 0
	for _, addr := range postUpgradeAddresses {
		switch cldf.ContractType(addr.Type) {
		case routerops.ContractType:
			routerCount++
			require.Equal(t, routerAddr, addr.Address, "Router should be upgraded in place")
		case fqops.ContractType:
			fqCount++
			require.Equal(t, fqAddr, addr.Address, "FeeQuoter should be upgraded in place")
		case offrampops.ContractType:
			offRampCount++
		}
	}
	require.Equal(t, 1, routerCount, "Should have exactly one Router")
	require.Equal(t, 1, fqCount, "Should have exactly one FeeQuoter")
	require.GreaterOrEqual(t, offRampCount, 1, "Should have at least one OffRamp")
}

// TestDeployWithDownloadedArtifacts tests the preloaded artifact path
// (the default in-memory test flow) through the unified API.
// This does NOT use skipInCI — it works with preloaded .so files.
func TestDeployWithDownloadedArtifacts(t *testing.T) {
	t.Parallel()

	solSelector := chain_selectors.SOLANA_MAINNET.Selector
	programsPath, ds, err := PreloadSolanaEnvironment(t, solSelector)
	require.NoError(t, err)

	e, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solSelector}, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	version := semver.MustParse("1.6.0")
	mint, _ := solana.NewRandomPrivateKey()

	dReg := deployapi.GetRegistry()
	_, err = deployapi.DeployContracts(dReg).Apply(*e, deployapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			solSelector: {
				Version:       version,
				TokenPrivKey:  mint.String(),
				TokenDecimals: 9,
				MaxFeeJuelsPerMsg: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
				// No ChainSpecific — artifacts are already preloaded
			},
		},
	})
	require.NoError(t, err)

	addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solSelector))
	require.NotEmpty(t, addresses)
}
