package glamsterdam

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
)

// applyOnRampDestChainConfigUpdates mirrors onramp.ApplyDestChainConfigUpdates, but always routes
// the write through MCMS regardless of who the deployer key is, per this feature's "always MCMS"
// decision.
var applyOnRampDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]onramp.DestChainConfigArgs, *onramp.OnRampContract]{
	Name:            "glamsterdam:onramp:apply-dest-chain-config-updates",
	Version:         semver.MustParse("2.0.0"),
	Description:     "Calls applyDestChainConfigUpdates on OnRamp, always producing an MCMS proposal",
	ContractType:    onramp.ContractType,
	ContractABI:     onramp.OnRampABI,
	NewContract:     onramp.NewOnRampContract,
	IsAllowedCaller: contract.NoCallersAllowed[*onramp.OnRampContract, []onramp.DestChainConfigArgs],
	Validate:        func([]onramp.DestChainConfigArgs) error { return nil },
	CallContract: func(c *onramp.OnRampContract, opts *bind.TransactOpts, args []onramp.DestChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyDestChainConfigUpdates(opts, args)
	},
})

// applyFeeQuoterDestChainConfigUpdates mirrors fee_quoter.ApplyDestChainConfigUpdates, but always
// routes the write through MCMS.
var applyFeeQuoterDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]fee_quoter.DestChainConfigArgs, *fee_quoter.FeeQuoterContract]{
	Name:            "glamsterdam:fee-quoter:apply-dest-chain-config-updates",
	Version:         semver.MustParse("2.0.0"),
	Description:     "Calls applyDestChainConfigUpdates on FeeQuoter, always producing an MCMS proposal",
	ContractType:    fee_quoter.ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoterContract,
	IsAllowedCaller: contract.NoCallersAllowed[*fee_quoter.FeeQuoterContract, []fee_quoter.DestChainConfigArgs],
	Validate:        func([]fee_quoter.DestChainConfigArgs) error { return nil },
	CallContract: func(c *fee_quoter.FeeQuoterContract, opts *bind.TransactOpts, args []fee_quoter.DestChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyDestChainConfigUpdates(opts, args)
	},
})

// applyCommitteeVerifierRemoteChainConfigUpdates mirrors
// committee_verifier.ApplyRemoteChainConfigUpdates, but always routes the write through MCMS.
var applyCommitteeVerifierRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]committee_verifier.RemoteChainConfigArgs, *committee_verifier.CommitteeVerifierContract]{
	Name:            "glamsterdam:committee-verifier:apply-remote-chain-config-updates",
	Version:         semver.MustParse("2.0.0"),
	Description:     "Calls applyRemoteChainConfigUpdates on CommitteeVerifier, always producing an MCMS proposal",
	ContractType:    committee_verifier.ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifierContract,
	IsAllowedCaller: contract.NoCallersAllowed[*committee_verifier.CommitteeVerifierContract, []committee_verifier.RemoteChainConfigArgs],
	Validate:        func([]committee_verifier.RemoteChainConfigArgs) error { return nil },
	CallContract: func(c *committee_verifier.CommitteeVerifierContract, opts *bind.TransactOpts, args []committee_verifier.RemoteChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

// LaneAddresses is the set of contract addresses to update on one source chain that has a
// confirmed lane pointed at the Glamsterdam target chain.
type LaneAddresses struct {
	ChainSelector            uint64
	OnRampAddress            common.Address
	FeeQuoterAddress         common.Address
	CommitteeVerifierAddress common.Address
	// OffRampAddress is optional. If set, OffRamp's immutable gas fields are read and sanity
	// checked against their expected Prague baseline (no write is ever made — there is no
	// setter for either field). Leave as the zero address to skip this chain's check.
	OffRampAddress common.Address
}

// UpdateGasConfigInput is the input to UpdateGasConfig.
type UpdateGasConfigInput struct {
	// TargetChainSelector is the chain selector moving to Glamsterdam.
	TargetChainSelector uint64
	// Lanes are the source chains with a confirmed lane pointed at the target chain.
	Lanes []LaneAddresses
}

// UpdateGasConfigOutput is the output of UpdateGasConfig.
type UpdateGasConfigOutput struct {
	// BatchOps contains one MCMS batch operation per chain in Lanes.
	BatchOps []mcms_types.BatchOperation
	// Report is a human-readable summary of every field resolved on every lane, for inclusion in
	// the changeset's MCMS proposal description.
	Report *glamsterdamutils.Report
}

// UpdateGasConfig reads the current on-chain gas config for every confirmed lane, resolves each
// field in the v2.0 Glamsterdam mapping table against its expected Prague baseline (applying the
// literal Glamsterdam value on a match, or the field's fallback rule on a mismatch), and packages
// the resulting writes into one MCMS batch operation per chain. Every write is routed through
// MCMS regardless of who the deployer key is.
var UpdateGasConfig = cldf_ops.NewSequence(
	"UpdateGasConfigV2",
	semver.MustParse("2.0.0"),
	"Updates v2.0 source-side gas config for lanes pointed at the Glamsterdam target chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input UpdateGasConfigInput) (UpdateGasConfigOutput, error) {
		output := UpdateGasConfigOutput{Report: glamsterdamutils.NewReport()}

		for _, lane := range input.Lanes {
			chain, ok := chains.EVMChains()[lane.ChainSelector]
			if !ok {
				return UpdateGasConfigOutput{}, fmt.Errorf("chain with selector %d not found", lane.ChainSelector)
			}

			var writes []contract.WriteOutput

			// --- OnRamp: BaseExecutionGasCost (row 1) ---
			onRampCur, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.OnRampAddress,
				Args:          input.TargetChainSelector,
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf(
					"failed to read OnRamp dest chain config for src %d, dst %d: %w", lane.ChainSelector, input.TargetChainSelector, err,
				)
			}
			baseExecResult := glamsterdamutils.Resolve(OnRampBaseExecutionGasCost, onRampCur.Output.BaseExecutionGasCost)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, baseExecResult)

			onRampWrite, err := cldf_ops.ExecuteOperation(b, applyOnRampDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.OnRampAddress,
				Args: []onramp.DestChainConfigArgs{
					{
						DestChainSelector:         input.TargetChainSelector,
						Router:                    onRampCur.Output.Router,
						AddressBytesLength:        onRampCur.Output.AddressBytesLength,
						TokenReceiverAllowed:      onRampCur.Output.TokenReceiverAllowed,
						MessageNetworkFeeUSDCents: onRampCur.Output.MessageNetworkFeeUSDCents,
						TokenNetworkFeeUSDCents:   onRampCur.Output.TokenNetworkFeeUSDCents,
						BaseExecutionGasCost:      baseExecResult.AppliedValue,
						DefaultCCVs:               onRampCur.Output.DefaultCCVs,
						LaneMandatedCCVs:          onRampCur.Output.LaneMandatedCCVs,
						DefaultExecutor:           onRampCur.Output.DefaultExecutor,
						OffRamp:                   onRampCur.Output.OffRamp,
					},
				},
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf("failed to apply OnRamp dest chain config update for src %d: %w", lane.ChainSelector, err)
			}
			writes = append(writes, onRampWrite.Output)

			// --- FeeQuoter: DefaultTokenDestGasOverhead, MaxPerMsgGasLimit,
			// DestGasPerPayloadByteBase, DefaultTxGasLimit (rows 2-5) ---
			fqCur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
				Args:          input.TargetChainSelector,
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf(
					"failed to read FeeQuoter dest chain config for src %d, dst %d: %w", lane.ChainSelector, input.TargetChainSelector, err,
				)
			}

			defaultTokenDestGasOverheadResult := glamsterdamutils.Resolve(FeeQuoterDefaultTokenDestGasOverhead, fqCur.Output.DefaultTokenDestGasOverhead)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, defaultTokenDestGasOverheadResult)

			maxPerMsgGasLimitResult := glamsterdamutils.Resolve(FeeQuoterMaxPerMsgGasLimit, fqCur.Output.MaxPerMsgGasLimit)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, maxPerMsgGasLimitResult)

			destGasPerPayloadByteBaseResult := glamsterdamutils.Resolve(FeeQuoterDestGasPerPayloadByteBase, fqCur.Output.DestGasPerPayloadByteBase)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, destGasPerPayloadByteBaseResult)

			defaultTxGasLimitResult := glamsterdamutils.Resolve(FeeQuoterDefaultTxGasLimit, fqCur.Output.DefaultTxGasLimit)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, defaultTxGasLimitResult)

			newFQConfig := fqCur.Output
			newFQConfig.DefaultTokenDestGasOverhead = defaultTokenDestGasOverheadResult.AppliedValue
			newFQConfig.MaxPerMsgGasLimit = maxPerMsgGasLimitResult.AppliedValue
			newFQConfig.DestGasPerPayloadByteBase = destGasPerPayloadByteBaseResult.AppliedValue
			newFQConfig.DefaultTxGasLimit = defaultTxGasLimitResult.AppliedValue

			fqWrite, err := cldf_ops.ExecuteOperation(b, applyFeeQuoterDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
				Args: []fee_quoter.DestChainConfigArgs{
					{DestChainSelector: input.TargetChainSelector, DestChainConfig: newFQConfig},
				},
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf("failed to apply FeeQuoter dest chain config update for src %d: %w", lane.ChainSelector, err)
			}
			writes = append(writes, fqWrite.Output)

			// --- CommitteeVerifier: GasForVerification (row 8) ---
			cvCur, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.CommitteeVerifierAddress,
				Args:          input.TargetChainSelector,
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf(
					"failed to read CommitteeVerifier remote chain config for src %d, dst %d: %w", lane.ChainSelector, input.TargetChainSelector, err,
				)
			}
			gasForVerificationResult := glamsterdamutils.Resolve(CommitteeVerifierGasForVerification, cvCur.Output.RemoteChainConfig.GasForVerification)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, gasForVerificationResult)

			newCVConfig := cvCur.Output.RemoteChainConfig
			newCVConfig.GasForVerification = gasForVerificationResult.AppliedValue

			cvWrite, err := cldf_ops.ExecuteOperation(b, applyCommitteeVerifierRemoteChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.RemoteChainConfigArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.CommitteeVerifierAddress,
				Args:          []committee_verifier.RemoteChainConfigArgs{newCVConfig},
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf("failed to apply CommitteeVerifier remote chain config update for src %d: %w", lane.ChainSelector, err)
			}
			writes = append(writes, cvWrite.Output)

			// --- OffRamp: GasForCallExactCheck, MaxGasBufferToUpdateState (rows 6-7) ---
			// Read-only sanity check: both fields are immutable with no setter, so no write is
			// ever produced here, but an unexpected on-chain value is still worth surfacing.
			if lane.OffRampAddress != (common.Address{}) {
				offRampCur, err := cldf_ops.ExecuteOperation(b, offramp.GetStaticConfig, chain, contract.FunctionInput[struct{}]{
					ChainSelector: lane.ChainSelector,
					Address:       lane.OffRampAddress,
				})
				if err != nil {
					return UpdateGasConfigOutput{}, fmt.Errorf("failed to read OffRamp static config for src %d: %w", lane.ChainSelector, err)
				}

				if offRampCur.Output.GasForCallExactCheck != OffRampExpectedGasForCallExactCheck {
					line := fmt.Sprintf(
						"chain %d: OffRamp.StaticConfig.GasForCallExactCheck is immutable with no setter, but current "+
							"value %d does not match expected Prague baseline %d - unexpected deployment, investigate manually",
						lane.ChainSelector, offRampCur.Output.GasForCallExactCheck, OffRampExpectedGasForCallExactCheck,
					)
					b.Logger.Warn(line)
					output.Report.AddLine(line)
				}
				if offRampCur.Output.MaxGasBufferToUpdateState != OffRampExpectedMaxGasBufferToUpdateState {
					line := fmt.Sprintf(
						"chain %d: OffRamp.StaticConfig.MaxGasBufferToUpdateState is immutable with no setter, but current "+
							"value %d does not match expected Prague baseline %d - unexpected deployment, investigate manually",
						lane.ChainSelector, offRampCur.Output.MaxGasBufferToUpdateState, OffRampExpectedMaxGasBufferToUpdateState,
					)
					b.Logger.Warn(line)
					output.Report.AddLine(line)
				}
			}

			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf("failed to build batch operation for src %d: %w", lane.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}

		return output, nil
	},
)
