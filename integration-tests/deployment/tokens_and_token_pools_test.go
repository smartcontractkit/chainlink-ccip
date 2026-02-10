package deployment

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_0/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	bnmpool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	solanautils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	solutils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	evmchain "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	solchain "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	bnmERC20gen "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"
)

func TestTokensAndTokenPools(t *testing.T) {
	// Define aliases for v1.6.0 and v1.5.1
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)
	v1_5_1, err := semver.NewVersion("1.5.1")
	require.NoError(t, err)

	// Define chains
	evmChainSelA := chainsel.TEST_90000001.Selector
	evmChainSelB := chainsel.TEST_90000002.Selector
	solChainSel := chainsel.SOLANA_DEVNET.Selector

	// For simplicity, both EVM and Solana will use BurnMint token pools in this test
	evmTokenPoolType := cciputils.BurnMintTokenPool
	solTokenPoolType := cciputils.BurnMintTokenPool

	// Preload Solana programs
	programsPath, ds, err := PreloadSolanaEnvironment(t, solChainSel)
	require.NoError(t, err)

	// Setup test environment
	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solChainSel}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{evmChainSelA, evmChainSelB}),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	// Get chain info
	evmChainA, ok := env.BlockChains.EVMChains()[evmChainSelA]
	require.True(t, ok)
	evmChainB, ok := env.BlockChains.EVMChains()[evmChainSelB]
	require.True(t, ok)
	solChain, ok := env.BlockChains.SolanaChains()[solChainSel]
	require.True(t, ok)

	// Initialize v1.6.0 adapters
	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	// Configure deployment registry
	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmadapters.EVMDeployer{})
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deployapi.MCMSVersion, &solAdapter)

	// Configure MCMS registry
	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})

	// Registration happens automatically, so the `Register...`
	// calls below aren't needed, but are left here for clarity
	tokenRegistry := tokensapi.GetTokenAdapterRegistry()
	// tokenRegistry.RegisterTokenAdapter(chainsel.FamilyEVM, v1_6_0, &evmAdapter)
	// tokenRegistry.RegisterTokenAdapter(chainsel.FamilySolana, v1_6_0, &solAdapter)

	// NOTE: on solana, the LnR and BnM token pool programs are deployed once
	// using an empty qualifier "", so we should also set the qualifier to ""
	// here. If we don't, the TokenAdapter will try to look up a non-existent
	// token pool address in the datastore and fail.
	//
	// Define testing data for Solana
	solTestData := struct {
		TokenPoolQualifier string
		Token              tokensapi.DeployTokenInput
		Deployer           *solana.PrivateKey
		Chain              solchain.Chain
		Deploy             deployapi.ContractDeploymentConfigPerChain
	}{
		TokenPoolQualifier: "",
		Deployer:           solChain.DeployerKey,
		Chain:              solChain,
		Deploy:             NewDefaultDeploymentConfigForSolana(v1_6_0),
		Token: tokensapi.DeployTokenInput{
			Decimals:               uint8(9),
			Symbol:                 "SOL_TEST",
			Name:                   "SOLANA Test Token",
			Type:                   solutils.SPLTokens,
			Supply:                 big.NewInt(math.MaxInt64),
			PreMint:                big.NewInt(math.MaxInt64 / 2),
			ExternalAdmin:          solana.NewWallet().PublicKey().String(),
			DisableFreezeAuthority: true,
			Senders:                []string{solChain.DeployerKey.PublicKey().String()},
			TokenPrivKey:           "", // if empty, a new key will be generated
			CCIPAdmin:              "", // default to timelock admin
		},
	}

	// NOTE: unlike Solana, EVM token pools can be deployed several times
	// and each token pool is linked to exactly one token. To distinguish
	// between token pools for different tokens, a qualifier is needed.
	//
	// Define testing data for EVM
	evmTestData := []struct {
		TokenPoolQualifier string
		Token              tokensapi.DeployTokenInput
		Deployer           common.Address
		TAR                *tarbindings.TokenAdminRegistry
		Chain              evmchain.Chain
		Deploy             deployapi.ContractDeploymentConfigPerChain
	}{
		{
			TokenPoolQualifier: "EVM_TEST_POOL_A",
			Deployer:           evmChainA.DeployerKey.From,
			Chain:              evmChainA,
			TAR:                nil, // populated later
			Deploy:             NewDefaultDeploymentConfigForEVM(v1_6_0),
			Token: tokensapi.DeployTokenInput{
				Decimals:               uint8(18),
				Symbol:                 "EVM_TEST_A",
				Name:                   "EVM Test Token A",
				Type:                   bnmERC20ops.ContractType,
				Supply:                 big.NewInt(0), // unlimited supply
				PreMint:                big.NewInt(0),
				ExternalAdmin:          "",
				DisableFreezeAuthority: false,      // not needed for EVM
				TokenPrivKey:           "",         // not needed for EVM
				Senders:                []string{}, // not needed for test
				CCIPAdmin:              "",         // default to timelock admin
			},
		},
		{
			TokenPoolQualifier: "EVM_TEST_POOL_B",
			Deployer:           evmChainB.DeployerKey.From,
			Chain:              evmChainB,
			TAR:                nil, // populated later
			Deploy:             NewDefaultDeploymentConfigForEVM(v1_6_0),
			Token: tokensapi.DeployTokenInput{
				Decimals:               uint8(18),
				Symbol:                 "EVM_TEST_B",
				Name:                   "EVM Test Token B",
				Type:                   bnmERC20ops.ContractType,
				Supply:                 big.NewInt(0), // unlimited supply
				PreMint:                big.NewInt(0),
				ExternalAdmin:          "",
				DisableFreezeAuthority: false,      // not needed for EVM
				TokenPrivKey:           "",         // not needed for EVM
				Senders:                []string{}, // not needed for test
				CCIPAdmin:              "",         // default to timelock admin
			},
		},
	}

	// Construct deployment input
	deployInput := map[uint64]deployapi.ContractDeploymentConfigPerChain{solChainSel: solTestData.Deploy}
	for _, data := range evmTestData {
		deployInput[data.Chain.Selector] = data.Deploy
	}

	// Deploy TAR + other contracts
	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{Chains: deployInput, MCMS: mcms.Input{}})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	// Deploy MCMS on all chains
	DeployMCMS(t, env, solChainSel, []string{cciputils.CLLQualifier})
	for _, data := range evmTestData {
		DeployMCMS(t, env, data.Chain.Selector, []string{cciputils.CLLQualifier})
	}

	// NOTE: calling TransferOwnership immediately after DeployMCMS for each chain
	// leads to an issue where TransferOwnership cannot find any MCMS contracts in
	// the datastore (possibly because of a non-zero timelock delay?). This causes
	// the ownership transfer step to fail with a "failed to find timelock address
	// for chain" error. It seems that deploying MCMS on all chains first resolves
	// the issue from occurring, which is why it is written this way.

	// Transfer ownership to timelocks on all chains
	SolanaTransferOwnership(t, env, solChainSel)
	for _, data := range evmTestData {
		EVMTransferOwnership(t, env, data.Chain.Selector)
	}

	// Verify that TAR was deployed on EVM chains (and update EVM test data with TAR instances)
	for i, data := range evmTestData {
		tarAddress, err := evmAdapter.GetTokenAdminRegistryAddress(env.DataStore, data.Chain.Selector)
		require.NoError(t, err)
		tarContract, err := tarbindings.NewTokenAdminRegistry(tarAddress, data.Chain.Client)
		require.NoError(t, err)
		tarOwner, err := tarContract.Owner(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, 0, data.Deployer.Cmp(tarOwner), fmt.Sprintf("expected TAR owner %q to be deployer %q", tarOwner.Hex(), data.Deployer.Hex()))
		evmTestData[i].TAR = tarContract
	}

	t.Run("Token Expansion EVM and Solana", func(t *testing.T) {
		// Verify that token and token pool do NOT exist in the datastore yet
		for _, data := range evmTestData {
			_, err = tokensapi.FindOneTokenAddress(&evmAdapter, env.DataStore, data.Chain.Selector, data.Token.Symbol)
			require.Error(t, err)
			_, err = evmAdapter.FindLatestTokenPoolAddress(env.DataStore, data.Chain.Selector, data.TokenPoolQualifier, evmTokenPoolType.String())
			require.Error(t, err)
		}

		// Define token expansion input
		input := map[uint64]tokensapi.TokenExpansionInputPerChain{
			solChainSel: {
				TokenPoolQualifier: solTestData.TokenPoolQualifier,
				TokenPoolVersion:   v1_6_0,
				PoolType:           solTokenPoolType.String(),
				DeployTokenInput:   solTestData.Token,

				// optional fields left empty, but included here for completeness
				RemoteCounterpartUpdates: map[uint64]tokensapi.RateLimiterConfig{},
				RemoteCounterpartDeletes: []uint64{},
				TokenPoolRateLimitAdmin:  "",
				TokenPoolAdmin:           "",
				TARAdmin:                 "",
			},
		}

		// Add EVM chains to the input
		for _, data := range evmTestData {
			input[data.Chain.Selector] = tokensapi.TokenExpansionInputPerChain{
				TokenPoolQualifier: data.TokenPoolQualifier,
				TokenPoolVersion:   v1_5_1,
				PoolType:           evmTokenPoolType.String(),
				DeployTokenInput:   data.Token,
			}
		}

		// Run token expansion
		output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			TokenExpansionInputPerChain: input,
			ChainAdapterVersion:         v1_6_0,
			MCMS:                        NewDefaultInputForMCMS("Token Expansion"),
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
	})

	t.Run("EVM Token Adapter", func(t *testing.T) {
		t.Run("Validate TokenExpansion", func(t *testing.T) {
			for _, data := range evmTestData {
				// Verify that we can find the timelock address in the datastore
				timelockRef, err := datastore_utils.FindAndFormatRef(env.DataStore,
					datastore.AddressRef{Type: datastore.ContractType(common_utils.RBACTimelock)},
					data.Chain.Selector,
					datastore_utils.FullRef,
				)
				require.NoError(t, err)

				// Query EVM token info from the chain
				tokAddress, err := tokensapi.FindOneTokenAddress(&evmAdapter, env.DataStore, data.Chain.Selector, data.Token.Symbol)
				require.NoError(t, err)
				tokn, err := bnmERC20gen.NewBurnMintERC20(tokAddress, data.Chain.Client)
				require.NoError(t, err)
				supp, err := tokn.MaxSupply(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				deci, err := tokn.Decimals(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				symb, err := tokn.Symbol(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				ccipAdmin, err := tokn.GetCCIPAdmin(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				name, err := tokn.Name(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)

				// Verify on-chain token info matches what we provided to the changeset
				require.Equal(t, timelockRef.Address, ccipAdmin.String(), fmt.Sprintf("expected CCIP admin %q to be timelock %q", ccipAdmin.Hex(), timelockRef.Address))
				require.Equal(t, 0, data.Token.Supply.Cmp(supp))
				require.Equal(t, data.Token.Decimals, deci)
				require.Equal(t, data.Token.Symbol, symb)
				require.Equal(t, data.Token.Name, name)

				// Query EVM token pool info from chain
				tpAddress, err := evmAdapter.FindLatestTokenPoolAddress(env.DataStore, data.Chain.Selector, data.TokenPoolQualifier, string(evmTokenPoolType))
				require.NoError(t, err)
				tp, err := bnmpool.NewBurnMintTokenPool(tpAddress, data.Chain.Client)
				require.NoError(t, err)
				dec, err := tp.GetTokenDecimals(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				tok, err := tp.GetToken(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				tpo, err := tp.Owner(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)

				// Verify on-chain token pool info is consistent
				require.Equal(t, 0, data.Deployer.Cmp(tpo), fmt.Sprintf("expected EVM deployer to be the owner of the deployed token pool (deployer = %q, token pool owner = %q", data.Deployer.Hex(), tpo.Hex()))
				require.Equal(t, data.Token.Decimals, dec)
				require.Equal(t, tokAddress, tok)

				// Verify that DeriveTokenAddress works as expected
				derived, err := evmAdapter.DeriveTokenAddress(*env, data.Chain.Selector, datastore.AddressRef{
					ChainSelector: data.Chain.Selector,
					Qualifier:     data.TokenPoolQualifier,
					Type:          datastore.ContractType(evmTokenPoolType),
					Version:       v1_5_1,
				})
				require.NoError(t, err)
				require.Equal(t, 0, common.BytesToAddress(derived).Cmp(tokAddress))
			}
		})

		t.Run("Validate ManualRegistration", func(t *testing.T) {
			for _, data := range evmTestData {
				// Verify that the token and token pool exist in datastore
				tokAddress, err := tokensapi.FindOneTokenAddress(&evmAdapter, env.DataStore, data.Chain.Selector, data.Token.Symbol)
				require.NoError(t, err)
				tpAddress, err := evmAdapter.FindLatestTokenPoolAddress(env.DataStore, data.Chain.Selector, data.TokenPoolQualifier, evmTokenPoolType.String())
				require.NoError(t, err)

				// Verify that no **pending** admin exists for the token at the moment. Also,
				// the TokenExpansion changeset should have set the EVM deployer as the admin
				// for the token.
				tokConfig, err := data.TAR.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, tokAddress)
				require.NoError(t, err)
				require.Equal(t, 0, tokConfig.PendingAdministrator.Cmp(common.Address{}), fmt.Sprintf("expected pending admin %q to be zero address", tokConfig.PendingAdministrator.Hex()))
				require.Equal(t, 0, tokConfig.Administrator.Cmp(data.Deployer), fmt.Sprintf("expected current admin %q to be deployer %q", tokConfig.Administrator.Hex(), data.Deployer.Hex()))
				require.Equal(t, 0, tokConfig.TokenPool.Cmp(tpAddress), fmt.Sprintf("expected token pool %q to match registered pool %q", tokConfig.TokenPool.Hex(), tpAddress.Hex()))

				// At this point, the `TokenExpansion` changeset will have already configured an admin
				// for the token, so the EVM manual registration changeset should detect this and call
				// `TransferAdminRole` instead of `ProposeAdministrator`.  Once this changeset is run,
				// the pending admin should be updated to a non-zero address on-chain.
				output, err = tokensapi.
					ManualRegistration().
					Apply(*env, tokensapi.ManualRegistrationInput{
						ChainAdapterVersion: v1_6_0,
						MCMS:                NewDefaultInputForMCMS("Manual Registration EVM"),
						ExistingAddresses:   env.DataStore.Addresses().Filter(),
						ChainSelector:       data.Chain.Selector,
						RegisterTokenConfigs: tokensapi.RegisterTokenConfig{
							ProposedOwner:      data.Deployer.Hex(),
							TokenPoolQualifier: data.TokenPoolQualifier,
							TokenSymbol:        data.Token.Symbol,
							PoolType:           evmTokenPoolType.String(),
							SVMExtraArgs:       nil,
						},
					})
				require.NoError(t, err)
				testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
				MergeAddresses(t, env, output.DataStore)

				// Verify that a new admin was proposed for the specified token
				tokConfig, err = data.TAR.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, tokAddress)
				require.NoError(t, err)
				require.Equal(t, 0, tokConfig.PendingAdministrator.Cmp(data.Deployer), fmt.Sprintf("expected pending admin %q to be deployer %q", tokConfig.PendingAdministrator.Hex(), data.Deployer.Hex()))
				require.Equal(t, 0, tokConfig.Administrator.Cmp(data.Deployer), fmt.Sprintf("expected current admin %q to be deployer %q", tokConfig.Administrator.Hex(), data.Deployer.Hex()))
				require.Equal(t, 0, tokConfig.TokenPool.Cmp(tpAddress), fmt.Sprintf("expected token pool %q to match registered pool %q", tokConfig.TokenPool.Hex(), tpAddress.Hex()))
			}
		})

		t.Run("Validate ConfigureTokenForTransfers", func(t *testing.T) {
			require.Len(t, evmTestData, 2, "expected exactly two EVM test data entries for this test")
			evmA, evmB := evmTestData[0], evmTestData[1]
			defaultRL := tokensapi.RateLimiterConfig{
				Capacity:  big.NewInt(1_000_000_000),
				Rate:      big.NewInt(100_000_000),
				IsEnabled: true,
			}

			input := tokensapi.ConfigureTokensForTransfersConfig{
				ChainAdapterVersion: v1_6_0,
				MCMS:                NewDefaultInputForMCMS("Configure Tokens For Transfers"),
				Tokens: []tokensapi.TokenTransferConfig{
					{
						ChainSelector: evmA.Chain.Selector,
						ExternalAdmin: "", // inferred
						TokenPoolRef: datastore.AddressRef{
							ChainSelector: evmA.Chain.Selector,
							Qualifier:     evmA.TokenPoolQualifier,
							Type:          datastore.ContractType(evmTokenPoolType),
							Version:       v1_5_1,
						},
						RegistryRef: datastore.AddressRef{
							ChainSelector: evmA.Chain.Selector,
							Address:       evmA.TAR.Address().Hex(),
						},
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							evmB.Chain.Selector: {
								OutboundRateLimiterConfig: defaultRL,
								InboundRateLimiterConfig:  defaultRL,
								OutboundCCVs:              []datastore.AddressRef{},
								InboundCCVs:               []datastore.AddressRef{},
								RemoteToken: &datastore.AddressRef{
									ChainSelector: evmB.Chain.Selector,
									Qualifier:     evmB.Token.Symbol,
									Type:          datastore.ContractType(evmB.Token.Type),
								},
								RemotePool: &datastore.AddressRef{
									ChainSelector: evmB.Chain.Selector,
									Qualifier:     evmB.TokenPoolQualifier,
									Type:          datastore.ContractType(evmTokenPoolType),
								},
							},
						},
					},
				},
			}

			// Query the latest on-chain state for chain A
			poolAddressA, err := evmAdapter.FindLatestTokenPoolAddress(env.DataStore, evmA.Chain.Selector, evmA.TokenPoolQualifier, string(evmTokenPoolType))
			require.NoError(t, err)
			poolA, err := bnmpool.NewBurnMintTokenPool(poolAddressA, evmA.Chain.Client)
			require.NoError(t, err)
			outboundRateLimitAB, err := poolA.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
			require.NoError(t, err)
			inboundRateLimitAB, err := poolA.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
			require.NoError(t, err)
			supportedChainsOnA, err := poolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			remotePoolsAB, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
			require.NoError(t, err)
			remoteTokenAB, err := poolA.GetRemoteToken(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
			require.NoError(t, err)

			// Verify that the token pool on chain A has nothing configured yet for chain B
			require.False(t, outboundRateLimitAB.IsEnabled)
			require.False(t, inboundRateLimitAB.IsEnabled)
			require.Empty(t, supportedChainsOnA)
			require.Empty(t, remotePoolsAB)
			require.Empty(t, remoteTokenAB)

			// For the first iteration, there are no remote chains configured on token pool A so
			// ApplyChainUpdates should be called directly. On the second iteration the "update"
			// path will be taken instead of the "add" path, since chain B will already be fully
			// configured on chain A. Thus, running this twice in a row tests the idempotency of
			// the changeset.
			for range 2 {
				// Run the changeset
				output, err = tokensapi.ConfigureTokensForTransfers(tokenRegistry, mcmsRegistry).Apply(*env, input)
				require.NoError(t, err)
				MergeAddresses(t, env, output.DataStore)

				// Query the latest on-chain state for chain A
				outboundRateLimitAB, err = poolA.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
				require.NoError(t, err)
				inboundRateLimitAB, err = poolA.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
				require.NoError(t, err)
				supportedChainsOnA, err = poolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
				require.NoError(t, err)
				remotePoolsAB, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
				require.NoError(t, err)
				remoteTokenAB, err = poolA.GetRemoteToken(&bind.CallOpts{Context: t.Context()}, evmB.Chain.Selector)
				require.NoError(t, err)

				// Verify that chain B is now supported on chain A
				require.Equal(t, []uint64{evmB.Chain.Selector}, supportedChainsOnA)

				// Verify that the rate limits were set correctly
				require.Equal(t, 0, defaultRL.Capacity.Cmp(outboundRateLimitAB.Capacity))
				require.Equal(t, 0, defaultRL.Rate.Cmp(outboundRateLimitAB.Rate))
				require.True(t, outboundRateLimitAB.IsEnabled)
				require.Equal(t, 0, defaultRL.Capacity.Cmp(inboundRateLimitAB.Capacity))
				require.Equal(t, 0, defaultRL.Rate.Cmp(inboundRateLimitAB.Rate))
				require.True(t, inboundRateLimitAB.IsEnabled)

				// Verify that the remote token pool was set correctly
				poolB, err := evmAdapter.FindLatestTokenPoolAddress(env.DataStore, evmB.Chain.Selector, evmB.TokenPoolQualifier, evmTokenPoolType.String())
				require.NoError(t, err)
				require.Len(t, remotePoolsAB, 1)
				require.True(t, bytes.Equal(remotePoolsAB[0], common.LeftPadBytes(poolB.Bytes(), 32)))

				// Verify that the remote token was set correctly
				tokB, err := tokensapi.FindOneTokenAddress(&evmAdapter, env.DataStore, evmB.Chain.Selector, evmB.Token.Symbol)
				require.NoError(t, err)
				require.True(t, bytes.Equal(remoteTokenAB, common.LeftPadBytes(tokB.Bytes(), 32)))
			}
		})
	})

	t.Run("Solana Token Adapter", func(t *testing.T) {
		t.Run("Validate TokenExpansion ", func(t *testing.T) {
			// TODO: implement this once DeriveTokenAddress is supported for Solana
			t.Skip("Skipping Solana token expansion verification")
		})

		t.Run("Validate ManualRegistration", func(t *testing.T) {
			// TODO: Use here a different token
			// Verify that the token exists in datastore
			tokenAddr, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
				ChainSelector: solTestData.Chain.Selector,
				Qualifier:     solTestData.Token.Symbol,
			}, solTestData.Chain.Selector, datastore_utils.FullRef)
			require.NoError(t, err)
			_, err = solanautils.GetTokenProgramID(deployment.ContractType(tokenAddr.Type))
			require.NoError(t, err)
			tokenPool, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
				ChainSelector: solTestData.Chain.Selector,
				Type:          datastore.ContractType(common_utils.BurnMintTokenPool),
				Version:       common_utils.Version_1_6_0,
			}, solTestData.Chain.Selector, datastore_utils.FullRef)
			require.NoError(t, err)
			tokenPoolProgramId := solana.MustPublicKeyFromBase58(tokenPool.Address)

			// Verify that no **pending** admin exists for the token at the moment. Also,
			// The PDA for TokenAdminRegistry is not initialized
			tokenMint := solana.MustPublicKeyFromBase58(tokenAddr.Address)
			routerAdd, err := solAdapter.GetRouterAddress(env.DataStore, solTestData.Chain.Selector)
			require.NoError(t, err)
			routerProgramId := solana.PublicKeyFromBytes(routerAdd)
			tokenAdminRegistryPDA, _, _ := state.FindTokenAdminRegistryPDA(tokenMint, routerProgramId)

			var tokenAdminRegistryAccount ccip_common.TokenAdminRegistry
			tokenAdminRegistryErr := env.BlockChains.SolanaChains()[solTestData.Chain.Selector].GetAccountDataBorshInto(context.Background(), tokenAdminRegistryPDA, &tokenAdminRegistryAccount)
			require.Error(t, tokenAdminRegistryErr)

			// Verify that the PDA token pool has not been initialized
			tokenPoolStatePDA, _ := tokens.TokenPoolConfigAddress(tokenMint, tokenPoolProgramId)
			// TODO: This will be true when using a new token
			//var tokenPoolStateAccount burnmint_token_pool.State
			//tokenPoolStateErr := env.BlockChains.SolanaChains()[solTestData.Chain.Selector].GetAccountDataBorshInto(context.Background(), tokenPoolStatePDA, &tokenPoolStateAccount)
			//require.Error(t, tokenPoolStateErr)

			// Run the changeset
			output, err = tokensapi.
				ManualRegistration().
				Apply(*env, tokensapi.ManualRegistrationInput{
					ChainAdapterVersion: v1_6_0,
					MCMS:                NewDefaultInputForMCMS("Manual Registration Solana"),
					ExistingDataStore:   env.DataStore,
					ChainSelector:       solTestData.Chain.Selector,
					RegisterTokenConfigs: tokensapi.RegisterTokenConfig{
						ProposedOwner:      solTestData.Token.ExternalAdmin,
						TokenPoolQualifier: solTestData.TokenPoolQualifier,
						TokenSymbol:        solTestData.Token.Symbol,
						PoolType:           solTokenPoolType.String(),
						//SVMExtraArgs: &tokensapi.SVMExtraArgs{
						//	CustomerMintAuthorities: []solana.PublicKey{
						//		solana.MustPublicKeyFromBase58(solTestData.Token.ExternalAdmin),
						//	}},
					},
				})
			require.NoError(t, err)
			testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
			MergeAddresses(t, env, output.DataStore)

			// Verify that a new admin was proposed for the specified token
			var tokenAdminRegistryAccountAfter ccip_common.TokenAdminRegistry
			tarErr := env.BlockChains.SolanaChains()[solTestData.Chain.Selector].GetAccountDataBorshInto(context.Background(), tokenAdminRegistryPDA, &tokenAdminRegistryAccountAfter)
			require.NoError(t, tarErr)
			require.Equal(t, solana.PublicKey{}, tokenAdminRegistryAccountAfter.Administrator)
			require.Equal(t, solana.MustPublicKeyFromBase58(solTestData.Token.ExternalAdmin), tokenAdminRegistryAccountAfter.PendingAdministrator)

			var tokenPoolStateAccountAfter burnmint_token_pool.State
			stateErr := env.BlockChains.SolanaChains()[solTestData.Chain.Selector].GetAccountDataBorshInto(context.Background(), tokenPoolStatePDA, &tokenPoolStateAccountAfter)
			require.NoError(t, stateErr)
			// TODO: This will be true when using a new token
			//require.Equal(t, solana.PublicKey{}, tokenPoolStateAccountAfter.Config.Owner)
			//require.Equal(t, solana.MustPublicKeyFromBase58(solTestData.Token.ExternalAdmin), tokenPoolStateAccountAfter.Config.ProposedOwner)
			require.Equal(t, tokenMint, tokenPoolStateAccountAfter.Config.Mint)
			//require.Equal(t, solana.PublicKey{}, tokenPoolStateAccountAfter.Config.RateLimitAdmin)
		})

		t.Run("Validate ConfigureTokenForTransfers", func(t *testing.T) {
			// TODO: implement this once ConfigureTokenForTransfers is supported for Solana
			t.Skip("Skipping Solana configure token for transfers test - changeset is not implemented yet")
		})
	})
}
