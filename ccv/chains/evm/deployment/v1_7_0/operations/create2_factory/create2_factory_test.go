package create2_factory_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	create2_factory_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

// sendFunds sends a percentage of the sender's balance to the destination address
// percentageInt is the percentage as an integer (e.g., 50 for 50%)
func sendFunds(ctx context.Context, chain evm.Chain, to common.Address, percentageInt int64) error {
	// Get the sender's balance
	balance, err := chain.Client.BalanceAt(ctx, chain.DeployerKey.From, nil)
	if err != nil {
		return err
	}

	// Calculate the amount to send as a percentage of the balance
	percentage := big.NewInt(percentageInt)
	hundred := big.NewInt(100)
	amount := new(big.Int).Mul(balance, percentage)
	amount.Div(amount, hundred)

	nonce, err := chain.Client.PendingNonceAt(ctx, chain.DeployerKey.From)
	if err != nil {
		return err
	}

	gasPrice, err := chain.Client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)
	signedTx, err := chain.DeployerKey.Signer(chain.DeployerKey.From, tx)
	if err != nil {
		return err
	}

	err = chain.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	_, err = chain.Confirm(signedTx)
	if err != nil {
		return err
	}

	return nil
}

func TestCREATE2Factory(t *testing.T) {
	chain1Sel := uint64(5009297550715157269)
	chain2Sel := uint64(4949039107694359620)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1Sel, chain2Sel}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	evmChains := e.BlockChains.EVMChains()

	// Deploy CREATE2Factory on chain1
	factory1Report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		create2_factory.Deploy,
		evmChains[chain1Sel],
		contract.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
			ChainSelector:  chain1Sel,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{evmChains[chain1Sel].DeployerKey.From},
			},
		},
	)
	require.NoError(t, err, "Failed to deploy CREATE2Factory on chain1")

	// Adjust the deployer key on chain2 to be the same as chain1
	// + fund the key on chain2
	desiredDeployerKey := evmChains[chain1Sel].DeployerKey
	err = sendFunds(t.Context(), evmChains[chain2Sel], desiredDeployerKey.From, 50)
	require.NoError(t, err, "Failed to send funds to chain2")
	chain2 := evmChains[chain2Sel]
	chain2.DeployerKey = desiredDeployerKey
	evmChains[chain2Sel] = chain2
	// Ensure that the deployer key is funded on chain2
	balance, err := evmChains[chain2Sel].Client.BalanceAt(t.Context(), desiredDeployerKey.From, nil)
	require.NoError(t, err, "Failed to get balance of deployer key on chain2")
	require.Greater(t, balance.Int64(), int64(0), "Deployer key should be funded on chain2")

	// Deploy CREATE2Factory on chain2
	factory2Report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		create2_factory.Deploy,
		evmChains[chain2Sel],
		contract.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
			ChainSelector:  chain2Sel,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{evmChains[chain2Sel].DeployerKey.From},
			},
		},
	)
	require.NoError(t, err, "Failed to deploy CREATE2Factory on chain2")

	// Ensure that the factories addresses are the same on each chain
	require.Equal(t, factory1Report.Output.Address, factory2Report.Output.Address, "Factory addresses should be the same")

	// Now, increment the nonce of the deployer key on chain2 by sending funds to another address
	err = sendFunds(t.Context(), chain2, evmChains[chain1Sel].DeployerKey.From, 1)
	require.NoError(t, err, "Failed to send funds to random address")

	// Ensure that the nonce of the deployer key is different on each chain
	nonce1, err := evmChains[chain1Sel].Client.PendingNonceAt(t.Context(), evmChains[chain1Sel].DeployerKey.From)
	require.NoError(t, err, "Failed to get nonce of deployer key on chain1")
	nonce2, err := evmChains[chain2Sel].Client.PendingNonceAt(t.Context(), evmChains[chain2Sel].DeployerKey.From)
	require.NoError(t, err, "Failed to get nonce of deployer key on chain2")
	require.NotEqual(t, nonce1, nonce2, "Nonce should be different")

	// Now, create a contract via the factory on each chain
	// The resulting addresses should be the same
	_, err = cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		create2_factory.CreateAndTransferOwnership,
		evmChains[chain1Sel],
		contract.FunctionInput[create2_factory.CreateAndTransferOwnershipArgs]{
			ChainSelector: chain1Sel,
			Address:       common.HexToAddress(factory1Report.Output.Address),
			Args: create2_factory.CreateAndTransferOwnershipArgs{
				ComputeAddressArgs: create2_factory.ComputeAddressArgs{
					ABI:             create2_factory_latest.CREATE2FactoryMetaData.ABI,
					Bin:             create2_factory_latest.CREATE2FactoryBin,
					ConstructorArgs: []any{[]common.Address{}},
					Salt:            "salt",
				},
				To: evmChains[chain1Sel].DeployerKey.From,
			},
		})
	require.NoError(t, err, "Failed to create and transfer ownership of contract on chain1")

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, create2_factory.CreateAndTransferOwnership, evmChains[chain2Sel], contract.FunctionInput[create2_factory.CreateAndTransferOwnershipArgs]{
		ChainSelector: chain2Sel,
		Address:       common.HexToAddress(factory2Report.Output.Address),
		Args: create2_factory.CreateAndTransferOwnershipArgs{
			ComputeAddressArgs: create2_factory.ComputeAddressArgs{
				ABI:             create2_factory_latest.CREATE2FactoryMetaData.ABI,
				Bin:             create2_factory_latest.CREATE2FactoryBin,
				ConstructorArgs: []any{[]common.Address{}},
				Salt:            "salt",
			},
		},
	})
	require.NoError(t, err, "Failed to create and transfer ownership of contract on chain2")

	// Since factories should be at the same address, we should filter on both chains
	// using that same factory address
	boundFactory1, err := create2_factory_latest.NewCREATE2Factory(common.HexToAddress(factory1Report.Output.Address), evmChains[chain1Sel].Client)
	require.NoError(t, err, "Failed to bind CREATE2Factory on chain1")

	iter1, err := boundFactory1.FilterContractDeployed(nil, nil)
	require.NoError(t, err, "Failed to filter ContractDeployed events on chain1")
	var contractAddress1 common.Address
	for iter1.Next() {
		contractAddress1 = iter1.Event.ContractAddress
	}
	iter1.Close()

	boundFactory2, err := create2_factory_latest.NewCREATE2Factory(common.HexToAddress(factory2Report.Output.Address), evmChains[chain2Sel].Client)
	require.NoError(t, err, "Failed to bind CREATE2Factory on chain2")

	iter2, err := boundFactory2.FilterContractDeployed(nil, nil)
	require.NoError(t, err, "Failed to filter ContractDeployed events on chain2")
	var contractAddress2 common.Address
	for iter2.Next() {
		contractAddress2 = iter2.Event.ContractAddress
	}
	iter2.Close()

	require.Equal(t, contractAddress1, contractAddress2, "Contract addresses should be the same")
}
