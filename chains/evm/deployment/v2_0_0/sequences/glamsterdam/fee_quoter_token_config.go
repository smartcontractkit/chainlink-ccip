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

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	feequoterbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
)

// applyFeeQuoterTokenTransferFeeConfigUpdates mirrors fee_quoter.ApplyTokenTransferFeeConfigUpdates,
// but always routes the write through MCMS regardless of who the deployer key is.
var applyFeeQuoterTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs, *fee_quoter.FeeQuoterContract]{
	Name:            "glamsterdam:fee-quoter:apply-token-transfer-fee-config-updates",
	Version:         semver.MustParse("2.0.0"),
	Description:     "Calls applyTokenTransferFeeConfigUpdates on FeeQuoter, always producing an MCMS proposal",
	ContractType:    fee_quoter.ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoterContract,
	IsAllowedCaller: contract.NoCallersAllowed[*fee_quoter.FeeQuoterContract, fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(c *fee_quoter.FeeQuoterContract, opts *bind.TransactOpts, args fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return c.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

// getAllTokenTransferFeeConfigs reads every per-token override on a FeeQuoter, for every destination
// chain, in a single call. (The generated ops package doesn't wrap this getter, so it is bound here.)
var getAllTokenTransferFeeConfigs = contract.NewRead(contract.ReadParams[struct{}, feequoterbindings.GetAllTokenTransferFeeConfigs, *feequoterbindings.FeeQuoter]{
	Name:         "glamsterdam:fee-quoter:get-all-token-transfer-fee-configs",
	Version:      semver.MustParse("2.0.0"),
	Description:  "Calls getAllTokenTransferFeeConfigs on the FeeQuoter contract",
	ContractType: fee_quoter.ContractType,
	NewContract:  feequoterbindings.NewFeeQuoter,
	CallContract: func(c *feequoterbindings.FeeQuoter, opts *bind.CallOpts, _ struct{}) (feequoterbindings.GetAllTokenTransferFeeConfigs, error) {
		return c.GetAllTokenTransferFeeConfigs(opts)
	},
})

// FeeQuoterTokenConfigLane describes one source chain with a confirmed lane to the target whose
// FeeQuoter per-token overrides must be migrated.
type FeeQuoterTokenConfigLane struct {
	ChainSelector    uint64
	FeeQuoterAddress common.Address
	// USDCPoolAddresses / LombardPoolAddresses are pools whose underlying token (getToken) must use
	// the USDC / Lombard spec instead of the generic scaling. A pool that can't be read is skipped
	// with a warning (its token then gets the generic scaling).
	USDCPoolAddresses    []common.Address
	LombardPoolAddresses []common.Address
}

// UpdateFeeQuoterTokenTransferFeeConfigInput is the input to UpdateFeeQuoterTokenTransferFeeConfig.
type UpdateFeeQuoterTokenTransferFeeConfigInput struct {
	// TargetChainSelector is the chain selector moving to Glamsterdam.
	TargetChainSelector uint64
	Lanes               []FeeQuoterTokenConfigLane
}

// UpdateFeeQuoterTokenTransferFeeConfigOutput is the output of UpdateFeeQuoterTokenTransferFeeConfig.
type UpdateFeeQuoterTokenTransferFeeConfigOutput struct {
	// BatchOps contains at most one MCMS batch operation per lane (one FeeQuoter write covering
	// every updated token).
	BatchOps []mcms_types.BatchOperation
	Report   *glamsterdamutils.Report
}

// UpdateFeeQuoterTokenTransferFeeConfig migrates the FeeQuoter per-token TokenTransferFeeConfig
// overrides for the target destination. The v2.0 OnRamp charges a token's destination gas from the
// pool's own fee config when the pool is IPoolV2 and that config is enabled, and from this
// FeeQuoter override (or the default) otherwise — so both must be migrated, or the Glamsterdam
// change is lost whenever the active source flips.
//
// Overrides are read with one getAllTokenTransferFeeConfigs call per FeeQuoter, so every token that
// has an enabled override for the target is covered regardless of what kind of pool it uses (on
// mainnet the large majority are ordinary BurnMint/LockRelease tokens, not USDC/Lombard):
//
//   - USDC tokens use FeeQuoterUSDCTokenDestGasOverhead and Lombard tokens
//     FeeQuoterLombardTokenDestGasOverhead: the literal Glamsterdam value when the current override
//     equals the expected Prague baseline, a no-op when it already equals the Glamsterdam value, and
//     the ratio fallback otherwise (logged as MISMATCH).
//   - Every other token with an enabled override is scaled by
//     FeeQuoterGenericTokenDestGasOverheadScale (see its caveat about idempotency).
//   - Tokens without an enabled override are left alone: they use the destination default, which the
//     main gas-config sequence already migrates.
var UpdateFeeQuoterTokenTransferFeeConfig = cldf_ops.NewSequence(
	"UpdateFeeQuoterTokenTransferFeeConfigV2",
	semver.MustParse("2.0.0"),
	"Updates v2.0 FeeQuoter per-token transfer fee overrides for lanes pointed at the Glamsterdam target chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input UpdateFeeQuoterTokenTransferFeeConfigInput) (UpdateFeeQuoterTokenTransferFeeConfigOutput, error) {
		output := UpdateFeeQuoterTokenTransferFeeConfigOutput{Report: glamsterdamutils.NewReport()}

		for _, lane := range input.Lanes {
			chain, ok := chains.EVMChains()[lane.ChainSelector]
			if !ok {
				return UpdateFeeQuoterTokenTransferFeeConfigOutput{}, fmt.Errorf("chain with selector %d not found", lane.ChainSelector)
			}

			// Pools are only used to recognise their token as USDC/Lombard, so a pool that can't be read
			// (e.g. a stale datastore entry with no code) must not block the batch: warn and carry on.
			recognise := func(pools []common.Address, into map[common.Address]bool, kind string) {
				for _, pool := range pools {
					out, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, contract.FunctionInput[struct{}]{
						ChainSelector: lane.ChainSelector,
						Address:       pool,
					})
					if err != nil {
						output.Report.AddLine(fmt.Sprintf(
							"chain %d: WARNING - failed to read underlying token of pool %s: %v; its token will not be recognised as %s",
							lane.ChainSelector, pool, err, kind,
						))
						continue
					}
					into[out.Output] = true
				}
			}
			usdcTokens := map[common.Address]bool{}
			lombardTokens := map[common.Address]bool{}
			recognise(lane.USDCPoolAddresses, usdcTokens, "USDC")
			recognise(lane.LombardPoolAddresses, lombardTokens, "Lombard")

			all, err := cldf_ops.ExecuteOperation(b, getAllTokenTransferFeeConfigs, chain, contract.FunctionInput[struct{}]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
			})
			if err != nil {
				return UpdateFeeQuoterTokenTransferFeeConfigOutput{}, fmt.Errorf(
					"failed to read FeeQuoter %s token transfer fee configs on src %d: %w", lane.FeeQuoterAddress, lane.ChainSelector, err,
				)
			}

			var singleTokenArgs []fee_quoter.TokenTransferFeeConfigSingleTokenArgs
			overrides := 0
			for i, dest := range all.Output.DestChainSelectors {
				if dest != input.TargetChainSelector {
					continue
				}
				for j, token := range all.Output.TransferTokens[i] {
					cur := all.Output.TokenTransferFeeConfigs[i][j]
					if !cur.IsEnabled {
						continue
					}
					overrides++

					var applied uint32
					switch {
					case usdcTokens[token]:
						result := glamsterdamutils.Resolve(FeeQuoterUSDCTokenDestGasOverhead, cur.DestGasOverhead)
						glamsterdamutils.AddField(output.Report, lane.ChainSelector, result)
						applied = result.AppliedValue
					case lombardTokens[token]:
						result := glamsterdamutils.Resolve(FeeQuoterLombardTokenDestGasOverhead, cur.DestGasOverhead)
						glamsterdamutils.AddField(output.Report, lane.ChainSelector, result)
						applied = result.AppliedValue
					default:
						applied = FeeQuoterGenericTokenDestGasOverheadScale(cur.DestGasOverhead)
						output.Report.AddLine(fmt.Sprintf(
							"chain %d: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (token %s) generic scaling %d -> %d "+
								"(no per-token Glamsterdam target)",
							lane.ChainSelector, token, cur.DestGasOverhead, applied,
						))
					}

					if applied == cur.DestGasOverhead {
						continue // nothing to write (e.g. already at the Glamsterdam value)
					}
					singleTokenArgs = append(singleTokenArgs, fee_quoter.TokenTransferFeeConfigSingleTokenArgs{
						Token: token,
						TokenTransferFeeConfig: fee_quoter.TokenTransferFeeConfig{
							FeeUSDCents:       cur.FeeUSDCents,
							DestGasOverhead:   applied,
							DestBytesOverhead: cur.DestBytesOverhead,
							IsEnabled:         cur.IsEnabled,
						},
					})
				}
			}

			output.Report.AddLine(fmt.Sprintf(
				"chain %d: FeeQuoter(%s) has %d enabled token overrides for dst %d, %d updated",
				lane.ChainSelector, lane.FeeQuoterAddress, overrides, input.TargetChainSelector, len(singleTokenArgs),
			))

			if len(singleTokenArgs) == 0 {
				continue
			}

			write, err := cldf_ops.ExecuteOperation(b, applyFeeQuoterTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
				Args: fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs{
					TokenTransferFeeConfigArgs: []fee_quoter.TokenTransferFeeConfigArgs{
						{DestChainSelector: input.TargetChainSelector, TokenTransferFeeConfigs: singleTokenArgs},
					},
				},
			})
			if err != nil {
				return UpdateFeeQuoterTokenTransferFeeConfigOutput{}, fmt.Errorf("failed to apply FeeQuoter token transfer fee config update for src %d: %w", lane.ChainSelector, err)
			}
			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{write.Output})
			if err != nil {
				return UpdateFeeQuoterTokenTransferFeeConfigOutput{}, fmt.Errorf("failed to build batch operation for FeeQuoter %s on src %d: %w", lane.FeeQuoterAddress, lane.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}

		return output, nil
	},
)
