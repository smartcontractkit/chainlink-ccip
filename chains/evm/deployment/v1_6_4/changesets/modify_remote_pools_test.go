package changesets_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	environment "github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	factory_burn_mint_erc20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
)

func TestModifyRemotePoolsChangeset(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	remoteChainSelector := uint64(chain_selectors.TEST_90000002.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)

	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy a mock BurnMintERC20 token contract for use as the pooled token in the BurnMintTokenPool
	burnMintERC20Addr, tx, _, err := factory_burn_mint_erc20.DeployFactoryBurnMintERC20(
		evmChain.DeployerKey,
		evmChain.Client,
		"TestBurnMintToken", // name
		"TST",               // symbol
		18,                  // decimals
		big.NewInt(0),       // max supply
		big.NewInt(0),       // pre mint
		common.Address{1},   // new owner
	)
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm BurnMintERC20 deployment transaction")

	// deploy a mock burn mint token pool
	burnMintTokenPoolAddr, tx, _, err := burn_mint_token_pool.DeployBurnMintTokenPool(
		evmChain.DeployerKey,
		evmChain.Client,
		burnMintERC20Addr,  // token address
		18,                 // decimals
		[]common.Address{}, // allowed callers
		common.Address{1},  // rmnProxy
		common.Address{2},  // router
	)
	require.NoError(t, err, "Failed to deploy BurnMintTokenPool")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm BurnMintTokenPool deployment transaction")
	ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("BurnMintTokenPool"),
		Version:       semver.MustParse("1.6.4"),
		Address:       burnMintTokenPoolAddr.Hex(),
		ChainSelector: chainSelector,
	})

	changesetInput := changesets.ApplyChainUpdatesInput{
		ChainInputs: []changesets.ApplyChainUpdatesPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       burnMintTokenPoolAddr,
				Updates: token_pool_ops.ApplyChainUpdatesArgs{
					RemoteChainSelectorsToRemove: []uint64{},
					ChainsToAdd: []token_pool_ops.ChainUpdate{
						{
							RemoteChainSelector: remoteChainSelector,
							// Remote Pool address is an array of byte-arrays
							RemotePoolAddresses: [][]byte{[]byte{1}},
							RemoteTokenAddress:  []byte{2},
							OutboundRateLimiterConfig: token_pool_bindings.RateLimiterConfig{
								IsEnabled: false,
								Capacity:  big.NewInt(0),
								Rate:      big.NewInt(0),
							},
							InboundRateLimiterConfig: token_pool_bindings.RateLimiterConfig{
								IsEnabled: false,
								Capacity:  big.NewInt(0),
								Rate:      big.NewInt(0),
							},
						},
					},
				},
			},
		},
		MCMS: mcms.Input{},
	}

	changeset := changesets.ApplyChainUpdatesChangeset()
	validate := changeset.VerifyPreconditions(*e, changesetInput)
	require.NoError(t, validate, "Failed to validate ApplyChainUpdatesChangeset")

	changesetOutput, err := changeset.Apply(*e, changesetInput)
	require.NoError(t, err, "Failed to apply ApplyChainUpdatesChangeset")
	require.NotNil(t, changesetOutput, "Changeset output should not be nil")

	require.Greater(t, len(changesetOutput.Reports), 0)

	// check that the token pool is updated
	tokenPool, err := burn_mint_token_pool.NewBurnMintTokenPool(common.HexToAddress(burnMintTokenPoolAddr.Hex()), evmChain.Client)
	require.NoError(t, err, "Failed to get BurnMintTokenPool")

	remotePools, err := tokenPool.GetRemotePools(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err, "Failed to get remote pools")
	require.Equal(t, len(remotePools), 1)
	require.Equal(t, remotePools[0], []byte{1}, "Remote pool should be correct")

	isRemotePool, err := tokenPool.IsRemotePool(&bind.CallOpts{}, remoteChainSelector, remotePools[0])
	require.NoError(t, err, "Failed to check if remote pool is configured")
	require.True(t, isRemotePool, "Remote pool should be configured on the remote chain")

	isSupportedChain, err := tokenPool.IsSupportedChain(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err, "Failed to check if remote chain is supported")
	require.True(t, isSupportedChain, "Remote chain should be supported on the remote chain")

	// Now remove the chain by making the call to the modifyRemotePoolsChangeset with the removal flag.
	modifyRemotePoolsInput := changesets.ModifyRemotePoolsInput{
		ChainInputs: []changesets.ModifyRemotePoolsPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       burnMintTokenPoolAddr,
				Modification: token_pool.RemotePoolModification{
					Operation:           sequences.RemoveRemotePoolOperation,
					RemoteChainSelector: remoteChainSelector,
					RemotePoolAddress:   []byte{1},
				},
			},
		},
	}

	modifyRemotePoolsChangeset := changesets.ModifyRemotePoolsChangeset()
	modifyRemotePoolsValidate := modifyRemotePoolsChangeset.VerifyPreconditions(*e, modifyRemotePoolsInput)
	require.NoError(t, modifyRemotePoolsValidate, "Failed to validate ModifyRemotePoolsChangeset")

	modifyRemotePoolsOutput, err := modifyRemotePoolsChangeset.Apply(*e, modifyRemotePoolsInput)
	require.NoError(t, err, "Failed to apply ModifyRemotePoolsChangeset")
	require.NotNil(t, modifyRemotePoolsOutput, "ModifyRemotePools changeset output should not be nil")
	require.Greater(t, len(modifyRemotePoolsOutput.Reports), 0)

	remotePools, err = tokenPool.GetRemotePools(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err, "Failed to get remote pools after removal")
	require.Equal(t, len(remotePools), 0, "Remote pools should be empty after removal")

	isRemotePool, err = tokenPool.IsRemotePool(&bind.CallOpts{}, remoteChainSelector, []byte{1})
	require.NoError(t, err, "Failed to check if remote pool is configured after removal")
	require.False(t, isRemotePool, "Remote pool should not be configured after removal")

	// Should still be true since removing the pool does not disable the chain intself.
	isSupportedChain, err = tokenPool.IsSupportedChain(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err, "Failed to check if remote chain is supported after removal")
	require.True(t, isSupportedChain, "Remote chain should not be supported after removal")

	// Now Re-Add the chain by making the call to the modifyRemotePoolsChangeset with the addition flag.
	modifyRemotePoolsInput = changesets.ModifyRemotePoolsInput{
		ChainInputs: []changesets.ModifyRemotePoolsPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       burnMintTokenPoolAddr,
				Modification: token_pool.RemotePoolModification{
					Operation:           sequences.AddRemotePoolOperation,
					RemoteChainSelector: remoteChainSelector,
					RemotePoolAddress:   []byte{1},
				},
			},
		},
	}

	modifyRemotePoolsChangeset = changesets.ModifyRemotePoolsChangeset()
	modifyRemotePoolsValidate = modifyRemotePoolsChangeset.VerifyPreconditions(*e, modifyRemotePoolsInput)
	require.NoError(t, modifyRemotePoolsValidate, "Failed to validate ModifyRemotePoolsChangeset")

	modifyRemotePoolsOutput, err = modifyRemotePoolsChangeset.Apply(*e, modifyRemotePoolsInput)
	require.NoError(t, err, "Failed to apply ModifyRemotePoolsChangeset")
	require.NotNil(t, modifyRemotePoolsOutput, "ModifyRemotePools changeset output should not be nil")
	require.Greater(t, len(modifyRemotePoolsOutput.Reports), 0)

	remotePools, err = tokenPool.GetRemotePools(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err, "Failed to get remote pools after re-addition")
	require.Equal(t, len(remotePools), 1, "Remote pools should be re-added")

	isRemotePool, err = tokenPool.IsRemotePool(&bind.CallOpts{}, remoteChainSelector, []byte{1})
	require.NoError(t, err, "Failed to check if remote pool is configured after re-addition")
	require.True(t, isRemotePool, "Remote pool should be configured after re-addition")
}
