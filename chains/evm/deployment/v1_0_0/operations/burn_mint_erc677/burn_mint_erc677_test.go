package burn_mint_erc677

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	wrappers "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
)

func TestResolveGrantMintAndBurnRolesAuthority_owner(t *testing.T) {
	const selector uint64 = 5009297550715157269
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selector}))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[selector]
	pool := common.HexToAddress("0x000000000000000000000000000000000000bEEF")

	tokenAddress, tx, _, err := wrappers.DeployBurnMintERC677(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		"TT",
		18,
		big.NewInt(0),
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	auth, err := ResolveGrantMintAndBurnRolesAuthority(t.Context(), chain.Client, tokenAddress, chain.DeployerKey.From)
	require.NoError(t, err)
	require.Equal(t, AuthorityOwner, auth.Kind)
	require.Equal(t, chain.DeployerKey.From, auth.Owner)

	writes, err := PrepareGrantMintAndBurnRoles(e.OperationsBundle, chain, contract.FunctionInput[common.Address]{
		ChainSelector: selector,
		Address:       tokenAddress,
		Args:          pool,
	}, common.Address{})
	require.NoError(t, err)
	require.Len(t, writes, 1)

	parsedABI, err := abi.JSON(strings.NewReader(wrappers.BurnMintERC677ABI))
	require.NoError(t, err)
	wantData, err := parsedABI.Pack("grantMintAndBurnRoles", pool)
	require.NoError(t, err)
	require.Equal(t, wantData, writes[0].Tx.Data)
}

func TestResolveGrantMintAndBurnRolesAuthority_unauthorizedCaller(t *testing.T) {
	const selector uint64 = 5009297550715157269
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selector}))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[selector]
	tokenAddress, tx, _, err := wrappers.DeployBurnMintERC677(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		"TT",
		18,
		big.NewInt(0),
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	stranger := common.HexToAddress("0x000000000000000000000000000000000000CAFE")
	pool := common.HexToAddress("0x000000000000000000000000000000000000bEEF")

	auth, err := ResolveGrantMintAndBurnRolesAuthority(t.Context(), chain.Client, tokenAddress, stranger)
	require.NoError(t, err)
	require.Equal(t, AuthorityUnauthorized, auth.Kind)

	_, err = PrepareGrantMintAndBurnRoles(e.OperationsBundle, chain, contract.FunctionInput[common.Address]{
		ChainSelector: selector,
		Address:       tokenAddress,
		Args:          pool,
	}, stranger)
	require.ErrorContains(t, err, "not the token owner")
}

func TestPrepareGrantMintAndBurnRoles_timelockMustBeOwnerWhenSet(t *testing.T) {
	const selector uint64 = 5009297550715157269
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selector}))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[selector]
	timelock := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	pool := common.HexToAddress("0x000000000000000000000000000000000000bEEF")

	tokenAddress, tx, _, err := wrappers.DeployBurnMintERC677(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		"TT",
		18,
		big.NewInt(0),
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	_, err = PrepareGrantMintAndBurnRoles(e.OperationsBundle, chain, contract.FunctionInput[common.Address]{
		ChainSelector: selector,
		Address:       tokenAddress,
		Args:          pool,
	}, timelock)
	require.ErrorContains(t, err, "not the token owner")
}
