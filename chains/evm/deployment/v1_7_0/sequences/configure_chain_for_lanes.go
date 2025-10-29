package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type CommitteeVerifierDestChainConfig struct {
	// Whether to allow traffic TO the remote chain
	AllowlistEnabled bool
	// Addresses that are allowed to send messages TO the remote chain
	AddedAllowlistedSenders []common.Address
	// Addresses that are no longer allowed to send messages TO the remote chain
	RemovedAllowlistedSenders []common.Address
	// The fee in USD cents charged for verification on the remote chain
	FeeUSDCents uint16
	// The gas required to execute the verification call on the destination chain. Used in billing.
	GasForVerification uint32
	// The size of the CCV specific payload in bytes. Used in billing.
	PayloadSizeBytes uint32
}

type RemoteChainConfig struct {
	// Whether to allow traffic FROM this remote chain
	AllowTrafficFrom bool
	// The address on the remote chain from which the message is emitted
	// For example, on EVM chains, this is the OnRamp
	CCIPMessageSource []byte
	// The address on the remote chain against which the message gets executed
	CCIPMessageDest []byte
	// The default CCVs that will be applied to messages FROM this remote chain if no receiver is specified
	DefaultInboundCCVs []common.Address
	// Any CCVs that must always be used for messages FROM this remote chain
	LaneMandatedInboundCCVs []common.Address
	// The CCVs that will be used for messages TO this remote chain if none are specified
	DefaultOutboundCCVs []common.Address
	// The CCVs that will always be applied to messages TO this remote chain
	LaneMandatedOutboundCCVs []common.Address
	// The executor that will be used for messages TO this remote chain if none is specified
	// The address corresponds to the Executor contract
	DefaultExecutor common.Address
	// CommitteeVerifierDestChainConfig configures the CommitteeVerifier for this remote chain
	CommitteeVerifierDestChainConfig CommitteeVerifierDestChainConfig
	// FeeQuoterDestChainConfig configures the FeeQuoter for this remote chain
	FeeQuoterDestChainConfig fee_quoter.DestChainConfig
	// ExecutorDestChainConfig configures the Executor for this remote chain
	ExecutorDestChainConfig executor.RemoteChainConfig
	// The length in bytes of addresses on the remote chain
	AddressBytesLength uint8
}

type ConfigureChainForLanesInput struct {
	// The selector of the EVM chain being configured
	ChainSelector uint64
	// The router on the EVM chain being configured
	// We assume that all connections will use the same router, either test or production
	Router common.Address
	// The OnRamp on the EVM chain being configured
	// Similarly, we assume that all connections will use the same OnRamp
	OnRamp common.Address
	// The CommitteeVerifiers on the EVM chain being configured
	CommitteeVerifiers []common.Address
	// The FeeQuoter on the EVM chain being configured
	FeeQuoter common.Address
	// The OffRamp on the EVM chain being configured
	OffRamp common.Address
	// The configuration of each remote chain to configure
	RemoteChains map[uint64]RemoteChainConfig
}

var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures an EVM chain as a source & destination for multiple remote chains",
	func(b operations.Bundle, chain evm.Chain, input ConfigureChainForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)

		// Create inputs for each operation
		offRampArgs := make([]offramp.SourceChainConfigArgs, 0, len(input.RemoteChains))
		onRampArgs := make([]onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		committeeVerifierDestConfigArgs := make([]committee_verifier.DestChainConfigArgs, 0, len(input.RemoteChains))
		committeeVerifierAllowlistArgs := make([]committee_verifier.AllowlistConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter.DestChainConfigArgs, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		destChainSelectorsPerExecutor := make(map[common.Address][]executor.RemoteChainConfigArgs)
		for remoteSelector, remoteConfig := range input.RemoteChains {
			offRampArgs = append(offRampArgs, offramp.SourceChainConfigArgs{
				Router:              input.Router,
				SourceChainSelector: remoteSelector,
				IsEnabled:           remoteConfig.AllowTrafficFrom,
				OnRamp:              remoteConfig.CCIPMessageSource,
				DefaultCCV:          remoteConfig.DefaultInboundCCVs,
				LaneMandatedCCVs:    remoteConfig.LaneMandatedInboundCCVs,
			})
			onRampArgs = append(onRampArgs, onramp.DestChainConfigArgs{
				Router:             input.Router,
				AddressBytesLength: remoteConfig.AddressBytesLength,
				DestChainSelector:  remoteSelector,
				DefaultCCVs:        remoteConfig.DefaultOutboundCCVs,
				LaneMandatedCCVs:   remoteConfig.LaneMandatedOutboundCCVs,
				DefaultExecutor:    remoteConfig.DefaultExecutor,
				OffRamp:            remoteConfig.CCIPMessageDest,
			})
			committeeVerifierDestConfigArgs = append(committeeVerifierDestConfigArgs, committee_verifier.DestChainConfigArgs{
				Router:             input.Router,
				DestChainSelector:  remoteSelector,
				AllowlistEnabled:   remoteConfig.CommitteeVerifierDestChainConfig.AllowlistEnabled,
				FeeUSDCents:        remoteConfig.CommitteeVerifierDestChainConfig.FeeUSDCents,
				GasForVerification: remoteConfig.CommitteeVerifierDestChainConfig.GasForVerification,
				PayloadSizeBytes:   remoteConfig.CommitteeVerifierDestChainConfig.PayloadSizeBytes,
			})
			committeeVerifierAllowlistArgs = append(committeeVerifierAllowlistArgs, committee_verifier.AllowlistConfigArgs{
				AllowlistEnabled:          remoteConfig.CommitteeVerifierDestChainConfig.AllowlistEnabled,
				AddedAllowlistedSenders:   remoteConfig.CommitteeVerifierDestChainConfig.AddedAllowlistedSenders,
				RemovedAllowlistedSenders: remoteConfig.CommitteeVerifierDestChainConfig.RemovedAllowlistedSenders,
				DestChainSelector:         remoteSelector,
			})
			feeQuoterArgs = append(feeQuoterArgs, fee_quoter.DestChainConfigArgs{
				DestChainSelector: remoteSelector,
				DestChainConfig:   remoteConfig.FeeQuoterDestChainConfig,
			})
			onRampAdds = append(onRampAdds, router.OnRamp{
				DestChainSelector: remoteSelector,
				OnRamp:            input.OnRamp,
			})
			offRampAdds = append(offRampAdds, router.OffRamp{
				SourceChainSelector: remoteSelector,
				OffRamp:             input.OffRamp,
			})
			if destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor] == nil {
				destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor] = []executor.RemoteChainConfigArgs{}
			}
			destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor] = append(destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor], executor.RemoteChainConfigArgs{
				DestChainSelector: remoteSelector,
				Config:            remoteConfig.ExecutorDestChainConfig,
			})
		}

		// ApplySourceChainConfigUpdates on OffRamp
		offRampReport, err := cldf_ops.ExecuteOperation(b, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.OffRamp,
			Args:          offRampArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to OffRamp(%s) on chain %s: %w", input.OffRamp, chain, err)
		}
		writes = append(writes, offRampReport.Output)

		// ApplyDestChainConfigUpdates on OnRamp
		onRampReport, err := cldf_ops.ExecuteOperation(b, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.OnRamp,
			Args:          onRampArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to OnRamp(%s) on chain %s: %w", input.OnRamp, chain, err)
		}
		writes = append(writes, onRampReport.Output)

		// ApplyDestChainConfigUpdates on CommitteeVerifier
		for _, committeeVerifier := range input.CommitteeVerifiers {
			committeeVerifierReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       committeeVerifier,
				Args:          committeeVerifierDestConfigArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to CommitteeVerifier(%s) on chain %s: %w", committeeVerifier, chain, err)
			}
			writes = append(writes, committeeVerifierReport.Output)
		}

		// ApplyDestChainUpdates on each Executor
		for ExecutorAddr, destChainSelectorsToAdd := range destChainSelectorsPerExecutor {
			ExecutorReport, err := cldf_ops.ExecuteOperation(b, executor.ApplyDestChainUpdates, chain, contract.FunctionInput[executor.ApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       ExecutorAddr,
				Args: executor.ApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: destChainSelectorsToAdd,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to Executor(%s) on chain %s: %w", ExecutorAddr, chain, err)
			}
			writes = append(writes, ExecutorReport.Output)
		}

		// ApplyAllowlistUpdates on CommitteeVerifier
		for _, committeeVerifier := range input.CommitteeVerifiers {
			committeeVerifierAllowlistReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyAllowlistUpdates, chain, contract.FunctionInput[[]committee_verifier.AllowlistConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       committeeVerifier,
				Args:          committeeVerifierAllowlistArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allowlist updates to CommitteeVerifier(%s) on chain %s: %w", committeeVerifier, chain, err)
			}
			writes = append(writes, committeeVerifierAllowlistReport.Output)
		}

		// ApplyDestChainConfigUpdates on FeeQuoter
		feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.FeeQuoter,
			Args:          feeQuoterArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
		}
		writes = append(writes, feeQuoterReport.Output)

		// ApplyRampUpdates on Router
		routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.Router,
			Args: router.ApplyRampsUpdatesArgs{
				OnRampUpdates:  onRampAdds,
				OffRampRemoves: []router.OffRamp{}, // removals should be processed by a separate sequence responsible for disconnecting lanes
				OffRampAdds:    offRampAdds,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", input.Router, chain, err)
		}
		writes = append(writes, routerReport.Output)

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
