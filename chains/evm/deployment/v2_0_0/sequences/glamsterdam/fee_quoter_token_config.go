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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	tokenadminregistrybindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
)

// tokenAdminRegistryPageSize is how many configured tokens are fetched per
// TokenAdminRegistry.getAllConfiguredTokens call.
const tokenAdminRegistryPageSize = uint64(100)

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

// getAllConfiguredTokensArgs are the pagination args of TokenAdminRegistry.getAllConfiguredTokens.
type getAllConfiguredTokensArgs struct {
	StartIndex uint64
	MaxCount   uint64
}

// getAllConfiguredTokens reads one page of TokenAdminRegistry.getAllConfiguredTokens.
var getAllConfiguredTokens = contract.NewRead(contract.ReadParams[getAllConfiguredTokensArgs, []common.Address, *tokenadminregistrybindings.TokenAdminRegistry]{
	Name:         "glamsterdam:token-admin-registry:get-all-configured-tokens",
	Version:      token_admin_registry.Version,
	Description:  "Calls getAllConfiguredTokens on the TokenAdminRegistry contract",
	ContractType: token_admin_registry.ContractType,
	NewContract:  tokenadminregistrybindings.NewTokenAdminRegistry,
	CallContract: func(c *tokenadminregistrybindings.TokenAdminRegistry, opts *bind.CallOpts, args getAllConfiguredTokensArgs) ([]common.Address, error) {
		return c.GetAllConfiguredTokens(opts, args.StartIndex, args.MaxCount)
	},
})

// collectPaged drains a paginated list: it calls fetch(start, pageSize) until a page comes back
// shorter than pageSize.
func collectPaged(pageSize uint64, fetch func(start, maxCount uint64) ([]common.Address, error)) ([]common.Address, error) {
	var all []common.Address
	for start := uint64(0); ; start += pageSize {
		page, err := fetch(start, pageSize)
		if err != nil {
			return nil, err
		}
		all = append(all, page...)
		if uint64(len(page)) < pageSize {
			return all, nil
		}
	}
}

// FeeQuoterTokenConfigLane describes one source chain with a confirmed lane to the target whose
// FeeQuoter per-token overrides must be migrated.
type FeeQuoterTokenConfigLane struct {
	ChainSelector    uint64
	FeeQuoterAddress common.Address
	// TokenAdminRegistryAddress, if set, is enumerated (getAllConfiguredTokens) to find every token
	// that may carry a per-token override. The FeeQuoter has no on-chain enumeration of overrides,
	// so every registered token is checked and only those with an enabled override are updated.
	TokenAdminRegistryAddress common.Address
	// ExtraCandidateTokens are additional tokens to check (in addition to the registry's).
	ExtraCandidateTokens []common.Address
	// USDCPoolAddresses / LombardPoolAddresses are pools whose underlying token (getToken) must use
	// the USDC / Lombard spec instead of the generic scaling. Their tokens are also candidates.
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

			poolToken := func(pool common.Address) (common.Address, error) {
				out, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, contract.FunctionInput[struct{}]{
					ChainSelector: lane.ChainSelector,
					Address:       pool,
				})
				if err != nil {
					return common.Address{}, fmt.Errorf("failed to read underlying token of pool %s on src %d: %w", pool, lane.ChainSelector, err)
				}
				return out.Output, nil
			}

			usdcTokens := map[common.Address]bool{}
			lombardTokens := map[common.Address]bool{}
			var candidates []common.Address
			seen := map[common.Address]bool{}
			addCandidate := func(t common.Address) {
				if !seen[t] {
					seen[t] = true
					candidates = append(candidates, t)
				}
			}

			// A pool here is only used to recognise its token as USDC/Lombard, so a pool that can't be read
			// (e.g. a stale datastore entry with no code) must not block the batch: warn and carry on. The
			// token, if it has an override, is still found through the TokenAdminRegistry and then gets the
			// generic scaling instead of the USDC/Lombard rule.
			for _, pool := range lane.USDCPoolAddresses {
				t, err := poolToken(pool)
				if err != nil {
					output.Report.AddLine(fmt.Sprintf("chain %d: WARNING - %v; its token will not be recognised as USDC", lane.ChainSelector, err))
					continue
				}
				usdcTokens[t] = true
				addCandidate(t)
			}
			for _, pool := range lane.LombardPoolAddresses {
				t, err := poolToken(pool)
				if err != nil {
					output.Report.AddLine(fmt.Sprintf("chain %d: WARNING - %v; its token will not be recognised as Lombard", lane.ChainSelector, err))
					continue
				}
				lombardTokens[t] = true
				addCandidate(t)
			}
			for _, t := range lane.ExtraCandidateTokens {
				addCandidate(t)
			}
			if lane.TokenAdminRegistryAddress != (common.Address{}) {
				registryTokens, err := collectPaged(tokenAdminRegistryPageSize, func(start, maxCount uint64) ([]common.Address, error) {
					out, err := cldf_ops.ExecuteOperation(b, getAllConfiguredTokens, chain, contract.FunctionInput[getAllConfiguredTokensArgs]{
						ChainSelector: lane.ChainSelector,
						Address:       lane.TokenAdminRegistryAddress,
						Args:          getAllConfiguredTokensArgs{StartIndex: start, MaxCount: maxCount},
					})
					return out.Output, err
				})
				if err != nil {
					return UpdateFeeQuoterTokenTransferFeeConfigOutput{}, fmt.Errorf(
						"failed to list configured tokens of TokenAdminRegistry %s on src %d: %w",
						lane.TokenAdminRegistryAddress, lane.ChainSelector, err,
					)
				}
				for _, t := range registryTokens {
					addCandidate(t)
				}
			}

			var singleTokenArgs []fee_quoter.TokenTransferFeeConfigSingleTokenArgs
			overrides := 0
			for _, token := range candidates {
				cur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetTokenTransferFeeConfig, chain, contract.FunctionInput[fee_quoter.GetTokenTransferFeeConfigArgs]{
					ChainSelector: lane.ChainSelector,
					Address:       lane.FeeQuoterAddress,
					Args: fee_quoter.GetTokenTransferFeeConfigArgs{
						DestChainSelector: input.TargetChainSelector,
						Token:             token,
					},
				})
				if err != nil {
					return UpdateFeeQuoterTokenTransferFeeConfigOutput{}, fmt.Errorf(
						"failed to read FeeQuoter token transfer fee config for src %d, dst %d, token %s: %w",
						lane.ChainSelector, input.TargetChainSelector, token, err,
					)
				}
				if !cur.Output.IsEnabled {
					continue
				}
				overrides++

				var applied uint32
				switch {
				case usdcTokens[token]:
					result := glamsterdamutils.Resolve(FeeQuoterUSDCTokenDestGasOverhead, cur.Output.DestGasOverhead)
					glamsterdamutils.AddField(output.Report, lane.ChainSelector, result)
					applied = result.AppliedValue
				case lombardTokens[token]:
					result := glamsterdamutils.Resolve(FeeQuoterLombardTokenDestGasOverhead, cur.Output.DestGasOverhead)
					glamsterdamutils.AddField(output.Report, lane.ChainSelector, result)
					applied = result.AppliedValue
				default:
					applied = FeeQuoterGenericTokenDestGasOverheadScale(cur.Output.DestGasOverhead)
					output.Report.AddLine(fmt.Sprintf(
						"chain %d: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (token %s) generic scaling %d -> %d "+
							"(no per-token Glamsterdam target)",
						lane.ChainSelector, token, cur.Output.DestGasOverhead, applied,
					))
				}

				if applied == cur.Output.DestGasOverhead {
					continue // nothing to write (e.g. already at the Glamsterdam value)
				}
				newConfig := cur.Output
				newConfig.DestGasOverhead = applied
				singleTokenArgs = append(singleTokenArgs, fee_quoter.TokenTransferFeeConfigSingleTokenArgs{
					Token:                  token,
					TokenTransferFeeConfig: newConfig,
				})
			}

			output.Report.AddLine(fmt.Sprintf(
				"chain %d: FeeQuoter(%s) checked %d candidate tokens, %d with an enabled override for dst %d, %d updated",
				lane.ChainSelector, lane.FeeQuoterAddress, len(candidates), overrides, input.TargetChainSelector, len(singleTokenArgs),
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
