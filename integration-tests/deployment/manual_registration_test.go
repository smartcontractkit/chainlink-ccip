package deployment

import (
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"

	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	bnmpool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	solutils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	bnmERC20gen "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

// This test covers all methods on the TokenAdapter interface except
// for `ConfigureTokenForTransfersSequence` and `UpdateAuthorities`.
func TestManualRegistration(t *testing.T) {
	// Define aliases for v1.6.0 and v1.5.1
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)
	v1_5_1, err := semver.NewVersion("1.5.1")
	require.NoError(t, err)

	// Define chains
	solChainSel := chainsel.SOLANA_DEVNET.Selector
	evmChainSel := chainsel.TEST_90000002.Selector

	// Preload Solana programs
	programsPath, ds, err := PreloadSolanaEnvironment(t, solChainSel)
	require.NoError(t, err)

	// Setup test environment
	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solChainSel}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{evmChainSel}),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	// Get chain info
	evmChainInfo, ok := env.BlockChains.EVMChains()[evmChainSel]
	require.True(t, ok)
	solChainInfo, ok := env.BlockChains.SolanaChains()[solChainSel]
	require.True(t, ok)

	// Get deployer info
	evmDeployer := evmChainInfo.DeployerKey.From
	solDeployer := solChainInfo.DeployerKey

	// Initialize v1.6.0 adapters
	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	// Configure deployment registry
	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadapters.EVMDeployer{})
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deploy.MCMSVersion, &solAdapter)

	// Configure MCMS registry
	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})

	// NOTE: registration happens automatically, so the steps below are not needed, but are left here for clarity
	// Configure token registry
	// tokenRegistry := tokensapi.GetTokenAdapterRegistry()
	// tokenRegistry.RegisterTokenAdapter(chainsel.FamilyEVM, v1_6_0, &evmAdapter)
	// tokenRegistry.RegisterTokenAdapter(chainsel.FamilySolana, v1_6_0, &solAdapter)

	// Define SOL token info
	solTokenSupp := big.NewInt(math.MaxInt64)
	solTokenType := solutils.SPLTokens
	solTokenName := "SOLANA Test Token"
	solTokenDeci := uint8(9)
	solTokenSymb := "SOLTEST"

	// NOTE: on solana, the LnR and BnM token pool programs are deployed once
	// using an empty qualifier "", so we should also set the qualifier to ""
	// here. If we don't, the TokenAdapter will try to look up a non-existent
	// token pool address in the datastore and fail.
	solTokenPoolType := cciputils.BurnMintTokenPool
	solTokenPoolQlfr := ""

	// Define EVM token info
	evmTokenSupp := big.NewInt(math.MaxInt64)
	evmTokenType := bnmERC20ops.ContractType
	evmTokenName := "EVM Test Token"
	evmTokenDeci := uint8(18)
	evmTokenSymb := "EVMTEST"

	// NOTE: unlike Solana, EVM token pools can be deployed several times
	// and each token pool is linked to exactly one token. To distinguish
	// between token pools for different tokens, a qualifier is needed.
	evmTokenPoolType := cciputils.BurnMintTokenPool
	evmTokenPoolQlfr := "EVM_TEST_POOL"

	// Deploy TAR + other contracts
	solTokenPrivKey, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			solChainSel: {
				Version:                      v1_6_0,
				MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
				NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
				LinkPremiumMultiplier:        9e17, // 0.9 ETH
				TokenPrivKey:                 solTokenPrivKey.String(),
				TokenDecimals:                9,
			},
			evmChainSel: {
				Version:                                 v1_6_0,
				MaxFeeJuelsPerMsg:                       big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				TokenPriceStalenessThreshold:            uint32(24 * 60 * 60),
				NativeTokenPremiumMultiplier:            1e18, // 1.0 ETH
				LinkPremiumMultiplier:                   9e17, // 0.9 ETH
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
				GasForCallExactCheck:                    uint16(5000),
				TokenDecimals:                           18,
			},
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	MergeAddresses(env, output.DataStore)

	// Setup MCMS
	DeployMCMS(t, env, evmChainSel, []string{cciputils.CLLQualifier})
	DeployMCMS(t, env, solChainSel, []string{cciputils.CLLQualifier})
	SolanaTransferOwnership(t, env, solChainSel)
	EVMTransferOwnership(t, env, evmChainSel)

	// Verify that TAR was deployed
	tarAddress, err := evmAdapter.GetTokenAdminRegistryAddress(env.DataStore, evmChainSel)
	require.NoError(t, err)
	tarContract, err := tarbindings.NewTokenAdminRegistry(tarAddress, evmChainInfo.Client)
	require.NoError(t, err)
	tarOwner, err := tarContract.Owner(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.True(t, evmDeployer.Cmp(tarOwner) == 0, fmt.Sprintf("expected TAR owner %q to be deployer %q", tarOwner.Hex(), evmDeployer.Hex()))

	// This test will call the following methods on the TokenAdapter interface:
	//   - DeployToken
	//   - DeployTokenVerify
	//   - DeployTokenPoolForToken
	//   - RegisterToken
	//   - SetPool
	t.Run("Token Expansion EVM and Solana", func(t *testing.T) {
		// Verify that token and token pool do NOT exist in the datastore yet
		_, err = evmAdapter.FindOneTokenAddress(env.DataStore, evmChainSel, evmTokenSymb)
		require.Error(t, err)
		_, err = evmAdapter.FindLatestTokenPoolAddress(env.DataStore, evmChainSel, evmTokenPoolQlfr, evmTokenPoolType.String())
		require.Error(t, err)

		// Run token expansion
		output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0,
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				solChainSel: {
					TokenPoolQualifier: solTokenPoolQlfr,
					TokenPoolVersion:   v1_6_0,
					PoolType:           solTokenPoolType.String(),
					DeployTokenInput: tokensapi.DeployTokenInput{
						ExistingDataStore:      env.DataStore,
						ChainSelector:          solChainSel,
						Decimals:               solTokenDeci,
						Symbol:                 solTokenSymb,
						Name:                   solTokenName,
						Type:                   solTokenType,
						Supply:                 solTokenSupp,
						PreMint:                big.NewInt(math.MaxInt64 / 2),
						ExternalAdmin:          []string{},
						DisableFreezeAuthority: true,
						Senders:                []string{solDeployer.PublicKey().String()},
						TokenPrivKey:           "", // if empty, a new key will be generated
						CCIPAdmin:              "", // default to timelock admin
					},

					// optional fields left empty, but included here for completeness
					RemoteCounterpartUpdates: map[uint64]tokensapi.RateLimiterConfig{},
					RemoteCounterpartDeletes: []uint64{},
					TokenPoolRateLimitAdmin:  "",
					TokenPoolAdmin:           "",
					TARAdmin:                 "",
				},
				evmChainSel: {
					TokenPoolQualifier: evmTokenPoolQlfr,
					TokenPoolVersion:   v1_5_1,
					PoolType:           evmTokenPoolType.String(),
					DeployTokenInput: tokensapi.DeployTokenInput{
						ExistingDataStore:      env.DataStore,
						ChainSelector:          evmChainSel,
						Decimals:               evmTokenDeci,
						Symbol:                 evmTokenSymb,
						Name:                   evmTokenName,
						Type:                   evmTokenType,
						Supply:                 evmTokenSupp,
						PreMint:                big.NewInt(0),
						ExternalAdmin:          []string{},
						DisableFreezeAuthority: false,      // not needed for EVM
						TokenPrivKey:           "",         // not needed for EVM
						Senders:                []string{}, // not needed for test
						CCIPAdmin:              "",         // default to timelock admin
					},

					// optional fields left empty, but included here for completeness
					RemoteCounterpartUpdates: map[uint64]tokensapi.RateLimiterConfig{},
					RemoteCounterpartDeletes: []uint64{},
					TokenPoolRateLimitAdmin:  "",
					TokenPoolAdmin:           "",
					TARAdmin:                 "",
				},
			},
			MCMS: mcms.Input{
				OverridePreviousRoot: false,
				ValidUntil:           math.MaxUint32,
				TimelockDelay:        mcms_types.MustParseDuration("1s"),
				TimelockAction:       mcms_types.TimelockActionSchedule,
				Qualifier:            cciputils.CLLQualifier,
				Description:          "Token Expansion",
			},
		})
		require.NoError(t, err)
		MergeAddresses(env, output.DataStore)

		t.Run("Check EVM Token Expansion", func(t *testing.T) {
			// Query EVM token info from chain
			tokAddress, err := evmAdapter.FindOneTokenAddress(env.DataStore, evmChainSel, evmTokenSymb)
			require.NoError(t, err)
			tokn, err := bnmERC20gen.NewBurnMintERC20(tokAddress, evmChainInfo.Client)
			require.NoError(t, err)
			supp, err := tokn.MaxSupply(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			deci, err := tokn.Decimals(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			symb, err := tokn.Symbol(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			ownr, err := tokn.Owner(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			name, err := tokn.Name(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)

			// Ensure on-chain token info matches what we provided to the changeset
			require.True(t, evmDeployer.Cmp(ownr) == 0, fmt.Sprintf("expected EVM deployer to be the owner of the deployed token (deployer = %q, token owner = %q", evmDeployer.Hex(), ownr.Hex()))
			require.Equal(t, evmTokenSupp, supp)
			require.Equal(t, evmTokenDeci, deci)
			require.Equal(t, evmTokenSymb, symb)
			require.Equal(t, evmTokenName, name)

			// Query EVM token pool info from chain
			tpAddress, err := evmAdapter.FindLatestTokenPoolAddress(env.DataStore, evmChainSel, evmTokenPoolQlfr, string(evmTokenPoolType))
			require.NoError(t, err)
			tp, err := bnmpool.NewBurnMintTokenPool(tpAddress, evmChainInfo.Client)
			require.NoError(t, err)
			dec, err := tp.GetTokenDecimals(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			tok, err := tp.GetToken(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			own, err := tp.Owner(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)

			// Ensure on-chain token pool info is consistent
			require.True(t, evmDeployer.Cmp(own) == 0, fmt.Sprintf("expected EVM deployer to be the owner of the deployed token pool (deployer = %q, token pool owner = %q", evmDeployer.Hex(), own.Hex()))
			require.Equal(t, evmTokenDeci, dec)
			require.Equal(t, tokAddress, tok)
		})

		t.Run("Check Solana Token Expansion", func(t *testing.T) {
			// TODO: implement this
			t.Skip("Skipping Solana token expansion verification")
		})
	})

	// This test will call the following methods on the EVM TokenAdapter interface:
	//   - DeriveTokenAddress (which also calls AddressRefToBytes)
	//   - ManualRegistration
	t.Run("Manual Registration EVM", func(t *testing.T) {
		// Verify that the token and token pool exist in datastore
		tokAddress, err := evmAdapter.FindOneTokenAddress(env.DataStore, evmChainSel, evmTokenSymb)
		require.NoError(t, err)
		tpAddress, err := evmAdapter.FindLatestTokenPoolAddress(env.DataStore, evmChainSel, evmTokenPoolQlfr, evmTokenPoolType.String())
		require.NoError(t, err)

		// Verify that DeriveTokenAddress works as expected
		derived, err := evmAdapter.DeriveTokenAddress(*env, evmChainSel, datastore.AddressRef{
			ChainSelector: evmChainSel,
			Qualifier:     evmTokenPoolQlfr,
			Type:          datastore.ContractType(evmTokenPoolType),
			Version:       v1_5_1,
		})
		require.NoError(t, err)
		require.True(t, common.BytesToAddress(derived).Cmp(tokAddress) == 0)

		// Verify that no **pending** admin exists for the token at the moment. Also,
		// the TokenExpansion changeset should have set the EVM deployer as the admin
		// for the token.
		tokConfig, err := tarContract.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, tokAddress)
		require.NoError(t, err)
		require.True(t, tokConfig.PendingAdministrator.Cmp(common.Address{}) == 0, fmt.Sprintf("expected pending admin %q to be zero address", tokConfig.PendingAdministrator.Hex()))
		require.True(t, tokConfig.Administrator.Cmp(evmDeployer) == 0, fmt.Sprintf("expected current admin %q to be deployer %q", tokConfig.Administrator.Hex(), evmDeployer.Hex()))
		require.True(t, tokConfig.TokenPool.Cmp(tpAddress) == 0, fmt.Sprintf("expected token pool %q to match registered pool %q", tokConfig.TokenPool.Hex(), tpAddress.Hex()))

		// At this point, the `TokenExpansion` changeset will have already configured an admin
		// for the token, so the EVM manual registration changeset should detect this and call
		// `TransferAdminRole` instead of `ProposeAdministrator`.  Once this changeset is run,
		// the pending admin should be updated to a non-zero address on-chain.
		existingAddresses, err := env.DataStore.Addresses().Fetch()
		require.NoError(t, err)
		output, err = tokensapi.
			ManualRegistration().
			Apply(*env, tokensapi.ManualRegistrationInput{
				ChainAdapterVersion: v1_6_0,
				ExistingAddresses:   existingAddresses,
				ChainSelector:       evmChainSel,
				RegisterTokenConfigs: tokensapi.RegisterTokenConfig{
					ProposedOwner:      evmDeployer.Hex(),
					TokenPoolQualifier: evmTokenPoolQlfr,
					TokenSymbol:        evmTokenSymb,
					PoolType:           evmTokenPoolType.String(),
					SVMExtraArgs:       nil,
				},
				MCMS: mcms.Input{
					OverridePreviousRoot: false,
					ValidUntil:           math.MaxUint32,
					TimelockDelay:        mcms_types.MustParseDuration("1s"),
					TimelockAction:       mcms_types.TimelockActionSchedule,
					Qualifier:            cciputils.CLLQualifier,
					Description:          "Manual Registration",
				},
			})
		require.NoError(t, err)
		MergeAddresses(env, output.DataStore)

		// Verify that a new admin was proposed for the specified token
		tokConfig, err = tarContract.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, tokAddress)
		require.NoError(t, err)
		require.True(t, tokConfig.PendingAdministrator.Cmp(evmDeployer) == 0, fmt.Sprintf("expected pending admin %q to be deployer %q", tokConfig.PendingAdministrator.Hex(), evmDeployer.Hex()))
		require.True(t, tokConfig.Administrator.Cmp(evmDeployer) == 0, fmt.Sprintf("expected current admin %q to be deployer %q", tokConfig.Administrator.Hex(), evmDeployer.Hex()))
		require.True(t, tokConfig.TokenPool.Cmp(tpAddress) == 0, fmt.Sprintf("expected token pool %q to match registered pool %q", tokConfig.TokenPool.Hex(), tpAddress.Hex()))
	})

	// This test will call the following methods on the Solana TokenAdapter interface:
	//   - DeriveTokenAddress (which also calls AddressRefToBytes)
	//   - ManualRegistration
	t.Run("Manual Registration Solana", func(t *testing.T) {
		// TODO: implement this once manual registration is supported for Solana
		t.Skip("Skipping Solana manual registration test - changeset is not implemented yet")
	})
}
