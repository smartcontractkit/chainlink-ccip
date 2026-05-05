package burn_mint_erc677

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/cross_chain_token"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

func TestPrepareGrantMintAndBurnRolesAddsAdminGrantForDefaultAdminProposalExecutor(t *testing.T) {
	const selector uint64 = 5009297550715157269
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selector}))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[selector]
	timelock := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	pool := common.HexToAddress("0x000000000000000000000000000000000000bEEF")

	tokenAddress, tx, _, err := cross_chain_token.DeployCrossChainToken(
		chain.DeployerKey,
		chain.Client,
		cross_chain_token.BaseERC20ConstructorParams{
			Name:             "Cross Chain Test Token",
			Symbol:           "CCTT",
			MaxSupply:        big.NewInt(0),
			PreMint:          big.NewInt(0),
			PreMintRecipient: common.Address{},
			Decimals:         18,
			CcipAdmin:        timelock,
		},
		common.Address{},
		timelock,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	authority, err := ResolveGrantMintAndBurnRolesAuthority(t.Context(), chain.Client, tokenAddress, timelock)
	require.NoError(t, err)
	require.Equal(t, AuthorityDefaultAdmin, authority.Kind)

	writes, err := PrepareGrantMintAndBurnRoles(e.OperationsBundle, chain, contract.FunctionInput[common.Address]{
		ChainSelector: selector,
		Address:       tokenAddress,
		Args:          pool,
	}, timelock)
	require.NoError(t, err)
	require.Len(t, writes, 2)
	require.False(t, writes[0].Executed())
	require.False(t, writes[1].Executed())

	batchOp, err := contract.NewBatchOperationFromWrites(writes)
	require.NoError(t, err)
	require.Len(t, batchOp.Transactions, 2)
	require.Equal(t, tokenAddress.Hex(), batchOp.Transactions[0].To)
	require.Equal(t, tokenAddress.Hex(), batchOp.Transactions[1].To)

	parsedABI, err := abi.JSON(strings.NewReader(burnMintERC677ABI))
	require.NoError(t, err)
	grantRoleData, err := parsedABI.Pack("grantRole", authority.BurnMintAdminRole, timelock)
	require.NoError(t, err)
	grantMintAndBurnRolesData, err := parsedABI.Pack("grantMintAndBurnRoles", pool)
	require.NoError(t, err)
	require.Equal(t, grantRoleData, batchOp.Transactions[0].Data)
	require.Equal(t, grantMintAndBurnRolesData, batchOp.Transactions[1].Data)

	unauthorizedExecutor := common.HexToAddress("0x000000000000000000000000000000000000CAFE")
	_, err = PrepareGrantMintAndBurnRoles(e.OperationsBundle, chain, contract.FunctionInput[common.Address]{
		ChainSelector: selector,
		Address:       tokenAddress,
		Args:          pool,
	}, unauthorizedExecutor)
	require.ErrorContains(t, err, "cannot grant burn/mint roles")
}
