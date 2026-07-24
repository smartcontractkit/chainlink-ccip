package glamsterdam_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
)

const (
	discoveryTargetChainSel = uint64(3379446385462418246)
	discoveryLaneChainSelA  = uint64(4949039107694359620)
	discoveryLaneChainSelB  = uint64(5548718428018410741)
	discoveryNoLaneChainSel = uint64(6433500567565415381)
)

// evmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector EVM")) = 0x2812d52c, required
// by FeeQuoter.applyDestChainConfigUpdates for any enabled dest chain config.
var evmFamilySelector = [4]byte{0x28, 0x12, 0xd5, 0x2c}

func deployDiscoveryFeeQuoter(
	t *testing.T,
	e *cldf.Environment,
	chainSel uint64,
	destChainConfigArgs []fq16ops.DestChainConfigArgs,
) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSel]
	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.Deploy, chain, contract.DeployInput[fq16ops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq16ops.ContractType, *fq16ops.Version),
		Args: fq16ops.ConstructorArgs{
			StaticConfig: fq16ops.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1e18),
				LinkToken:                    common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
				TokenPriceStalenessThreshold: 3600,
			},
			DestChainConfigArgs: destChainConfigArgs,
		},
	})
	require.NoError(t, err)

	return common.HexToAddress(out.Output.Address)
}

func TestDiscoverLanesToTarget(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		discoveryLaneChainSelA,
		discoveryLaneChainSelB,
		discoveryNoLaneChainSel,
	}))
	require.NoError(t, err)

	enabledDestChainConfig := fq16ops.DestChainConfig{
		IsEnabled:               true,
		MaxNumberOfTokensPerMsg: 5,
		MaxPerMsgGasLimit:       3_000_000,
		DestGasOverhead:         300_000,
		ChainFamilySelector:     evmFamilySelector,
		DefaultTxGasLimit:       200_000,
		GasMultiplierWeiPerEth:  1e18,
	}
	laneAAddr := deployDiscoveryFeeQuoter(t, e, discoveryLaneChainSelA, []fq16ops.DestChainConfigArgs{
		{DestChainSelector: discoveryTargetChainSel, DestChainConfig: enabledDestChainConfig},
	})
	laneBAddr := deployDiscoveryFeeQuoter(t, e, discoveryLaneChainSelB, []fq16ops.DestChainConfigArgs{
		{DestChainSelector: discoveryTargetChainSel, DestChainConfig: enabledDestChainConfig},
	})
	noLaneAddr := deployDiscoveryFeeQuoter(t, e, discoveryNoLaneChainSel, nil)

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.DiscoverLanesToTarget, e.BlockChains, glamsterdamseq.DiscoverLanesToTargetInput{
		TargetChainSelector: discoveryTargetChainSel,
		FeeQuoterAddressByChain: map[uint64]common.Address{
			discoveryLaneChainSelA:  laneAAddr,
			discoveryLaneChainSelB:  laneBAddr,
			discoveryNoLaneChainSel: noLaneAddr,
		},
	})
	require.NoError(t, err)

	require.Equal(t, []uint64{discoveryLaneChainSelA, discoveryLaneChainSelB}, report.Output.LanesToUpdate)
	require.Equal(t, []uint64{discoveryNoLaneChainSel}, report.Output.NoLane)
}

func TestDiscoverLanesToTarget_UnknownChain(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{discoveryLaneChainSelA}))
	require.NoError(t, err)

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.DiscoverLanesToTarget, e.BlockChains, glamsterdamseq.DiscoverLanesToTargetInput{
		TargetChainSelector: discoveryTargetChainSel,
		FeeQuoterAddressByChain: map[uint64]common.Address{
			discoveryNoLaneChainSel: common.HexToAddress("0x0000000000000000000000000000000000000001"),
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "chain with selector")
}
