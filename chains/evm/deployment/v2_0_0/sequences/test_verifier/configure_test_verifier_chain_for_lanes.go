package test_verifier

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/verifier_test_helper"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
)

var ConfigureTestVerifierChainForLanes = cldf_ops.NewSequence(
	"configure-test-verifier-chain-for-lanes",
	semver.MustParse("2.0.0"),
	"Configures the test verifier and token pool for lanes on an EVM chain",
	func(b cldf_ops.Bundle, dep adapters.ConfigureTestVerifierForLanesDeps, input adapters.ConfigureTestVerifierForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Resolve local addresses from datastore.
		verifierRef, err := findRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(verifier_test_helper.ContractType),
			Version: verifier_test_helper.Version,
		}, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("resolve test verifier on chain %d: %w", input.ChainSelector, err)
		}
		verifierAddr := common.HexToAddress(verifierRef.Address)

		resolverRef, err := findRef(dep.DataStore, datastore.AddressRef{
			Type:      datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType),
			Version:   versioned_verifier_resolver.Version,
			Qualifier: TestVerifierResolverQualifier,
		}, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("resolve test verifier resolver on chain %d: %w", input.ChainSelector, err)
		}
		resolverAddr := common.HexToAddress(resolverRef.Address)

		testRouterAddr, err := resolveTestRouter(dep.DataStore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		poolRef, err := findRef(dep.DataStore, datastore.AddressRef{
			Type:      datastore.ContractType("BurnMintTokenPool"),
			Version:   semver.MustParse("2.0.0"),
			Qualifier: changesets.DefaultTokenName,
		}, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("resolve TESTVTR pool on chain %d: %w", input.ChainSelector, err)
		}
		poolAddr := common.HexToAddress(poolRef.Address)

		// Resolve the AdvancedPoolHooks address from the pool.
		hooksReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAdvancedPoolHooks, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       poolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get advanced pool hooks from TESTVTR pool on chain %d: %w", input.ChainSelector, err)
		}
		hooksAddr := hooksReport.Output

		// Parse allowed senders.
		senders := make([]common.Address, len(input.AllowedSenders))
		for i, s := range input.AllowedSenders {
			senders[i] = common.HexToAddress(s)
		}

		// Read which remote chains the pool already supports, so re-runs can skip the
		// ApplyChainUpdates call (which reverts on an already-supported chain).
		supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       poolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains from TESTVTR pool on chain %d: %w", input.ChainSelector, err)
		}
		supportedSet := make(map[uint64]struct{}, len(supportedChainsReport.Output))
		for _, s := range supportedChainsReport.Output {
			supportedSet[s] = struct{}{}
		}

		// Configure each remote chain.
		for remoteChainSelector, remoteCfg := range input.RemoteChains {
			// 1. Configure test verifier remote chain config.
			remoteChainConfigReport, err := cldf_ops.ExecuteOperation(b, verifier_test_helper.ApplyRemoteChainConfigUpdates, chain, contract_utils.FunctionInput[[]verifier_test_helper.RemoteChainConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       verifierAddr,
				Args: []verifier_test_helper.RemoteChainConfigArgs{{
					Router:              testRouterAddr,
					RemoteChainSelector: remoteChainSelector,
					AllowlistEnabled:    true,
					FeeUSDCents:         remoteCfg.FeeUSDCents,
					GasForVerification:  remoteCfg.GasForVerification,
					PayloadSizeBytes:    remoteCfg.PayloadSizeBytes,
				}},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure test verifier remote chain config for chain %d remote %d: %w", input.ChainSelector, remoteChainSelector, err)
			}
			writes = append(writes, remoteChainConfigReport.Output)

			// 2. Configure test verifier allowlist for remote chain.
			getRemoteConfigReports, err := cldf_ops.ExecuteOperation(b, verifier_test_helper.GetRemoteChainConfig, chain, contract_utils.FunctionInput[uint64]{
				ChainSelector: input.ChainSelector,
				Address:       verifierAddr,
				Args:          remoteChainSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get test verifier remote chain config for chain %d remote %d: %w", input.ChainSelector, remoteChainSelector, err)
			}
			currentAllowlist := getRemoteConfigReports.Output.AllowedSendersList
			addedSenders := make([]common.Address, 0)
			removedSenders := make([]common.Address, 0)
			for _, s := range senders {
				found := slices.ContainsFunc(currentAllowlist, func(a common.Address) bool { return a == s })
				if !found {
					addedSenders = append(addedSenders, s)
				}
			}
			for _, s := range currentAllowlist {
				found := slices.ContainsFunc(senders, func(a common.Address) bool { return a == s })
				if !found {
					removedSenders = append(removedSenders, s)
				}
			}

			if len(addedSenders) > 0 || len(removedSenders) > 0 {
				allowlistUpdateReport, err := cldf_ops.ExecuteOperation(b, verifier_test_helper.ApplyAllowlistUpdates, chain, contract_utils.FunctionInput[[]verifier_test_helper.AllowlistConfigArgs]{
					ChainSelector: input.ChainSelector,
					Address:       verifierAddr,
					Args: []verifier_test_helper.AllowlistConfigArgs{{
						DestChainSelector:         remoteChainSelector,
						AllowlistEnabled:          true,
						AddedAllowlistedSenders:   addedSenders,
						RemovedAllowlistedSenders: removedSenders,
					}},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to update test verifier allowlist for chain %d remote %d: %w", input.ChainSelector, remoteChainSelector, err)
				}
				writes = append(writes, allowlistUpdateReport.Output)
			}

			// 3. Wire outbound implementation on resolver.
			outboundReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
				ChainSelector: input.ChainSelector,
				Address:       resolverAddr,
				Args: []versioned_verifier_resolver.OutboundImplementationArgs{{
					DestChainSelector: remoteChainSelector,
					Verifier:          verifierAddr,
				}},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to wire outbound implementation on resolver for chain %d remote %d: %w", input.ChainSelector, remoteChainSelector, err)
			}
			writes = append(writes, outboundReport.Output)

			// 4. Configure token pool for remote chain.
			remoteAdapter, ok := dep.RemoteChains[remoteChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("no remote chain adapter for chain %d", remoteChainSelector)
			}
			remoteTokenBytes, err := remoteAdapter.TokenAddress(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token address for chain %d: %w", remoteChainSelector, err)
			}
			remotePoolBytes, err := remoteAdapter.PoolAddress(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pool address for chain %d: %w", remoteChainSelector, err)
			}

			// ApplyChainUpdates reverts on an already-supported chain, so when the remote is
			// already configured we only ensure the remote pool address is present.
			if _, alreadySupported := supportedSet[remoteChainSelector]; alreadySupported {
				paddedPool := common.LeftPadBytes(remotePoolBytes, 32)
				remotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemotePools, chain, contract_utils.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       poolAddr,
					Args:          remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pools for chain %d on chain %d: %w", remoteChainSelector, input.ChainSelector, err)
				}
				poolPresent := false
				for _, p := range remotePoolsReport.Output {
					if string(p) == string(paddedPool) {
						poolPresent = true
						break
					}
				}
				if !poolPresent {
					addPoolReport, err := cldf_ops.ExecuteOperation(b, token_pool.AddRemotePool, chain, contract_utils.FunctionInput[token_pool.AddRemotePoolArgs]{
						ChainSelector: input.ChainSelector,
						Address:       poolAddr,
						Args: token_pool.AddRemotePoolArgs{
							RemoteChainSelector: remoteChainSelector,
							RemotePoolAddress:   paddedPool,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to add remote pool for chain %d on chain %d: %w", remoteChainSelector, input.ChainSelector, err)
					}
					writes = append(writes, addPoolReport.Output)
				}
			} else {
				poolUpdateReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyChainUpdates, chain, contract_utils.FunctionInput[token_pool.ApplyChainUpdatesArgs]{
					ChainSelector: input.ChainSelector,
					Address:       poolAddr,
					Args: token_pool.ApplyChainUpdatesArgs{
						RemoteChainSelectorsToRemove: []uint64{},
						ChainsToAdd: []token_pool.ChainUpdate{{
							RemoteChainSelector: remoteChainSelector,
							RemotePoolAddresses: [][]byte{common.LeftPadBytes(common.BytesToAddress(remotePoolBytes).Bytes(), 32)},
							RemoteTokenAddress:  common.LeftPadBytes(common.BytesToAddress(remoteTokenBytes).Bytes(), 32),
							OutboundRateLimiterConfig: token_pool.Config{
								IsEnabled: false,
								Capacity:  big.NewInt(0),
								Rate:      big.NewInt(0),
							},
							InboundRateLimiterConfig: token_pool.Config{
								IsEnabled: false,
								Capacity:  big.NewInt(0),
								Rate:      big.NewInt(0),
							},
						}},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chain %d on chain %d: %w", remoteChainSelector, input.ChainSelector, err)
				}
				writes = append(writes, poolUpdateReport.Output)
			}

			// 5. Configure AdvancedPoolHooks to require the test verifier as a CCV
			// for this remote chain. Hooks store the actual verifier contract address (not
			// the resolver) — the resolver is used by the OnRamp/OffRamp for routing,
			// but the hooks directly reference verifiers.
			ccvConfigReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyCCVConfigUpdates, chain, contract_utils.FunctionInput[[]advanced_pool_hooks.CCVConfigArg]{
				ChainSelector: input.ChainSelector,
				Address:       hooksAddr,
				Args: []advanced_pool_hooks.CCVConfigArg{{
					RemoteChainSelector:   remoteChainSelector,
					OutboundCCVs:          []common.Address{resolverAddr},
					ThresholdOutboundCCVs: []common.Address{},
					InboundCCVs:           []common.Address{resolverAddr},
					ThresholdInboundCCVs:  []common.Address{},
				}},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure CCV on advanced pool hooks for chain %d remote %d: %w", input.ChainSelector, remoteChainSelector, err)
			}
			writes = append(writes, ccvConfigReport.Output)
		}

		// Create batch operation from writes.
		if len(writes) > 0 {
			batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			if len(batchOp.Transactions) > 0 {
				batchOps = append(batchOps, batchOp)
			}
		}

		return sequences.OnChainOutput{
			BatchOps: batchOps,
		}, nil
	},
)
