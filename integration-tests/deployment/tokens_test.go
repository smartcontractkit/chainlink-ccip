package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func TestTokenExpansion(t *testing.T) {
	t.Parallel()
	programsPath, ds, err := PreloadSolanaEnvironment(t, chain_selectors.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, ds, "Datastore should be created")

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}
	solanaChains := []uint64{
		chain_selectors.SOLANA_MAINNET.Selector,
	}
	allChains := append(evmChains, solanaChains...)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = ds.Seal() // Add preloaded contracts to env datastore

	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	for _, chainSel := range allChains {
		mint, _ := solana.NewRandomPrivateKey()
		out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
				chainSel: {
					Version: version,
					// LINK TOKEN CONFIG
					// token private key used to deploy the LINK token. Solana: base58 encoded private key
					TokenPrivKey: mint.String(),
					// token decimals used to deploy the LINK token
					TokenDecimals: 9,
					// FEE QUOTER CONFIG
					MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
					TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
					LinkPremiumMultiplier:        9e17, // 0.9 ETH
					NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
					// OFFRAMP CONFIG
					PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
					GasForCallExactCheck:                    uint16(5000),
				},
			},
		})
		require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
		out.DataStore.Merge(e.DataStore)
		e.DataStore = out.DataStore.Seal()
	}
	DeployMCMS(t, e, chain_selectors.SOLANA_MAINNET.Selector, []string{cciputils.CLLQualifier})
	SolanaTransferOwnership(t, e, chain_selectors.SOLANA_MAINNET.Selector)
	// TODO: EVM doesn't work with a non-zero timelock delay
	// DeployMCMS(t, e, chain_selectors.ETHEREUM_MAINNET.Selector)
	// EVMTransferOwnership(t, e, chain_selectors.ETHEREUM_MAINNET.Selector)

	sender, _ := solana.NewRandomPrivateKey()
	out, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		DeployTokenInputs: map[uint64]tokensapi.DeployTokenInput{
			chain_selectors.SOLANA_MAINNET.Selector: {
				Name:     "Test Token",
				Symbol:   "TEST",
				Decimals: 9,
				Type:     utils.SPLTokens,
				Senders: []string{
					sender.PublicKey().String(),
				},
				DisableFreezeAuthority: true,
			},
		},
		DeployTokenPoolInputs: map[uint64]tokensapi.DeployTokenPoolInput{
			chain_selectors.SOLANA_MAINNET.Selector: {
				RegisterTokenConfig: tokensapi.RegisterTokenConfig{
					TokenSymbol: "TEST",
					PoolType:    common_utils.BurnMintTokenPool.String(),
				},
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            cciputils.CLLQualifier,
			Description:          "TokensTokensTokens!",
		},
	})
	require.NotNil(t, out, "Changeset output should not be nil")
	require.NoError(t, err, "Failed to apply TokenExpansion changeset")
}
