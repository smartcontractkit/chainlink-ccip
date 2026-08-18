package test_verifier

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/verifier_test_helper"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// TestVerifierVersionTag must match the VersionTag passed to the test verifier's constructor.
var TestVerifierVersionTag = [4]byte{0, 0, 0, 1}

// TestVerifierResolverQualifier is the fixed CREATE2 salt for the test verifier resolver.
const TestVerifierResolverQualifier = "redeploy-test-verifier-resolver"

var DeployTestVerifierChain = cldf_ops.NewSequence(
	"deploy-test-verifier-chain",
	semver.MustParse("2.0.0"),
	"Deploys TESTVTR drip token, matching BurnMintTokenPool, test verifier, and resolver on an EVM chain",
	func(b cldf_ops.Bundle, dep adapters.DeployTestVerifierChainDeps, input adapters.DeployTestVerifierChainInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Existing addresses feed MaybeDeployContract so re-runs skip already-deployed contracts.
		existingAddresses := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
		)

		// Resolve required addresses from datastore.
		rmnProxyAddr, err := resolveRMNProxy(dep.DataStore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		testRouterAddr, err := resolveTestRouter(dep.DataStore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// 1. Deploy BurnMintERC20WithDrip token.
		tokenRef, err := contract_utils.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain, contract_utils.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
			Qualifier:      &input.TokenSymbol,
			Args: burn_mint_erc20_with_drip.ConstructorArgs{
				Name:   input.TokenName,
				Symbol: input.TokenSymbol,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s token on chain %d: %w", input.TokenSymbol, input.ChainSelector, err)
		}
		addresses = append(addresses, tokenRef)
		tokenAddr := common.HexToAddress(tokenRef.Address)

		// 2. Deploy AdvancedPoolHooks (no allowlist — allowlist is on the verifier).
		hooksRef, err := contract_utils.MaybeDeployContract(b, advanced_pool_hooks.Deploy, chain, contract_utils.DeployInput[advanced_pool_hooks.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *advanced_pool_hooks.Version),
			Qualifier:      &input.TokenSymbol,
			Args: advanced_pool_hooks.ConstructorArgs{
				Allowlist:                        []common.Address{},
				ThresholdAmountForAdditionalCCVs: big.NewInt(0),
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy advanced pool hooks on chain %d: %w", input.ChainSelector, err)
		}
		addresses = append(addresses, hooksRef)
		hooksAddr := common.HexToAddress(hooksRef.Address)

		// 3. Deploy BurnMintTokenPool v2.0.0 wired to TestRouter.
		poolRef, err := contract_utils.MaybeDeployContract(b, burn_mint_token_pool.Deploy, chain, contract_utils.DeployInput[burn_mint_token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_token_pool.ContractType, *burn_mint_token_pool.Version),
			Qualifier:      &input.TokenSymbol,
			Args: burn_mint_token_pool.ConstructorArgs{
				Token:              tokenAddr,
				LocalTokenDecimals: input.TokenDecimals,
				AdvancedPoolHooks:  hooksAddr,
				RmnProxy:           rmnProxyAddr,
				Router:             testRouterAddr,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s pool on chain %d: %w", input.TokenSymbol, input.ChainSelector, err)
		}
		addresses = append(addresses, poolRef)
		poolAddr := common.HexToAddress(poolRef.Address)

		// 3b. Add the pool as an authorized caller on the hooks so pool.lockOrBurn can
		// pass the preflightCheck (empty AuthorizedCallers = nobody allowed).
		authCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetAllAuthorizedCallers, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       hooksAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks on chain %d: %w", input.ChainSelector, err)
		}
		if !slices.Contains(authCallersReport.Output, poolAddr) {
			authCallerReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[advanced_pool_hooks.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSelector,
				Address:       hooksAddr,
				Args: advanced_pool_hooks.AuthorizedCallerArgs{
					AddedCallers: []common.Address{poolAddr},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize pool on advanced pool hooks on chain %d: %w", input.ChainSelector, err)
			}
			writes = append(writes, authCallerReport.Output)
		}

		// 4. Grant mint+burn roles on token to pool.
		grantReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain, contract_utils.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       tokenAddr,
			Args:          poolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint+burn roles to pool on chain %d: %w", input.ChainSelector, err)
		}
		writes = append(writes, grantReport.Output)

		// 5. Grant mint+burn roles on token to deployer (for minting).
		grantDeployerReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain, contract_utils.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       tokenAddr,
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint+burn roles to deployer on chain %d: %w", input.ChainSelector, err)
		}
		writes = append(writes, grantDeployerReport.Output)

		// 6. Mint initial supply to specified accounts. Idempotent: only mints the shortfall
		// between the desired amount and the account's current balance, so re-runs do not
		// inflate supply.
		for accountStr, amountStr := range input.PreMintAccounts {
			account := common.HexToAddress(accountStr)
			amount, ok := new(big.Int).SetString(amountStr, 10)
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid pre-mint amount %q for account %s on chain %d", amountStr, accountStr, input.ChainSelector)
			}
			balanceReport, err := cldf_ops.ExecuteOperation(b, erc20.BalanceOf, chain, contract_utils.FunctionInput[common.Address]{
				ChainSelector: input.ChainSelector,
				Address:       tokenAddr,
				Args:          account,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to read balance of %s on chain %d: %w", accountStr, input.ChainSelector, err)
			}
			shortfall := new(big.Int).Sub(amount, balanceReport.Output)
			if shortfall.Sign() <= 0 {
				b.Logger.Infof("Account %s already has %s TESTVTR on chain %d; skipping mint", accountStr, balanceReport.Output.String(), input.ChainSelector)
				continue
			}
			mintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.Mint, chain, contract_utils.FunctionInput[burn_mint_erc20_with_drip.MintArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenAddr,
				Args: burn_mint_erc20_with_drip.MintArgs{
					Account: account,
					Amount:  shortfall,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to mint %s tokens to %s on chain %d: %w", shortfall.String(), accountStr, input.ChainSelector, err)
			}
			writes = append(writes, mintReport.Output)
		}

		// 7. Revoke mint+burn roles from deployer for safety.
		revokeMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.RevokeMintRole, chain, contract_utils.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       tokenAddr,
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke mint role from deployer on chain %d: %w", input.ChainSelector, err)
		}
		writes = append(writes, revokeMintReport.Output)

		revokeBurnReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.RevokeBurnRole, chain, contract_utils.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       tokenAddr,
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke burn role from deployer on chain %d: %w", input.ChainSelector, err)
		}
		writes = append(writes, revokeBurnReport.Output)

		// 8. Deploy VerifierTestHelper (wired to TestRouter).
		verifierRef, err := contract_utils.MaybeDeployContract(b, verifier_test_helper.Deploy, chain, contract_utils.DeployInput[verifier_test_helper.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(verifier_test_helper.ContractType, *verifier_test_helper.Version),
			Args: verifier_test_helper.ConstructorArgs{
				TestRouter: testRouterAddr,
				Rmn:        rmnProxyAddr,
				VersionTag: TestVerifierVersionTag,
				TestToken:  tokenAddr,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy test verifier on chain %d: %w", input.ChainSelector, err)
		}
		addresses = append(addresses, verifierRef)

		// 9. Deploy VersionedVerifierResolver via CREATE2, or reuse the existing one. The
		// CREATE2 salt is fixed, so a second run would revert with FailedDeployment — check
		// the datastore first.
		resolverRefs := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByType(datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType)),
			datastore.AddressRefByVersion(versioned_verifier_resolver.Version),
			datastore.AddressRefByQualifier(TestVerifierResolverQualifier),
		)
		if len(resolverRefs) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("expected at most 1 test verifier resolver on chain %d, got %d", input.ChainSelector, len(resolverRefs))
		}
		var resolverRef datastore.AddressRef
		if len(resolverRefs) == 1 {
			resolverRef = resolverRefs[0]
		} else {
			create2FactoryAddr, err := resolveCREATE2Factory(dep.DataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("resolve CREATE2Factory on chain %d: %w", input.ChainSelector, err)
			}
			resolverReport, err := cldf_ops.ExecuteSequence(b, evm_sequences.DeployVerifierResolverViaCREATE2, chain, evm_sequences.DeployVerifierResolverViaCREATE2Input{
				CREATE2Factory: create2FactoryAddr,
				ChainSelector:  input.ChainSelector,
				Qualifier:      TestVerifierResolverQualifier,
				Type:           datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType),
				Version:        versioned_verifier_resolver.Version,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy test verifier resolver on chain %d: %w", input.ChainSelector, err)
			}
			if len(resolverReport.Output.Addresses) != 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("expected exactly one resolver address on chain %d, got %d", input.ChainSelector, len(resolverReport.Output.Addresses))
			}
			resolverRef = resolverReport.Output.Addresses[0]
			addresses = append(addresses, resolverRef)
			writes = append(writes, resolverReport.Output.Writes...)
		}

		// 10. Wire inbound implementation on resolver, skipping if already set.
		verifierAddr := common.HexToAddress(verifierRef.Address)
		resolverAddr := common.HexToAddress(resolverRef.Address)
		inboundWrites, err := ensureInboundImplementation(b, chain, input.ChainSelector, resolverAddr, verifierAddr)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, inboundWrites...)

		// 11. Register token on TokenAdminRegistry (set pool).
		registryAddr, err := resolveTokenAdminRegistry(dep.DataStore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve token admin registry on chain %d: %w", input.ChainSelector, err)
		}
		registerReport, err := cldf_ops.ExecuteSequence(b, v1_5_0.RegisterToken, chain, v1_5_0.RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenAddress:              tokenAddr,
			TokenPoolAddress:          poolAddr,
			TokenAdminRegistryAddress: registryAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to register token on chain %d: %w", input.ChainSelector, err)
		}
		batchOps = append(batchOps, registerReport.Output.BatchOps...)

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
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	},
)

// ensureInboundImplementation sets the test verifier as the inbound implementation on the
// resolver, skipping the write if the on-chain state already matches. Required for
// idempotency: re-running the sequence must not fail or emit redundant transactions.
func ensureInboundImplementation(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	chainSelector uint64,
	resolverAddr common.Address,
	verifierAddr common.Address,
) ([]contract_utils.WriteOutput, error) {
	desired := versioned_verifier_resolver.InboundImplementationArgs{
		Version:  TestVerifierVersionTag,
		Verifier: verifierAddr,
	}
	currentReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.GetAllInboundImplementations, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       resolverAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inbound implementations from resolver on chain %d: %w", chainSelector, err)
	}
	for _, cur := range currentReport.Output {
		if cur.Version == desired.Version && cur.Verifier == desired.Verifier {
			return nil, nil
		}
	}
	report, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
		ChainSelector: chainSelector,
		Address:       resolverAddr,
		Args:          []versioned_verifier_resolver.InboundImplementationArgs{desired},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to wire inbound implementation on resolver for chain %d: %w", chainSelector, err)
	}
	return []contract_utils.WriteOutput{report.Output}, nil
}

func resolveRMNProxy(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmn_proxy.ContractType),
		Version: rmn_proxy.Version,
	}, chainSelector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve RMNProxy on chain %d: %w", chainSelector, err)
	}
	return addr, nil
}

func resolveTestRouter(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(router.TestRouterContractType),
		Version: router.Version,
	}, chainSelector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve TestRouter on chain %d: %w", chainSelector, err)
	}
	return addr, nil
}

func resolveCREATE2Factory(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType("CREATE2Factory"),
		Version: semver.MustParse("2.0.0"),
	}, chainSelector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve CREATE2Factory on chain %d: %w", chainSelector, err)
	}
	return addr, nil
}

func resolveTokenAdminRegistry(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(token_admin_registry.ContractType),
		Version: token_admin_registry.Version,
	}, chainSelector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve TokenAdminRegistry on chain %d: %w", chainSelector, err)
	}
	return addr, nil
}

// findRef finds a single address ref matching the given criteria.
// The qualifier in filter is always applied (empty matches only canonical refs).
func findRef(ds datastore.DataStore, filter datastore.AddressRef, chainSelector uint64) (datastore.AddressRef, error) {
	return datastore_utils.FindAndFormatRef(ds, filter, chainSelector, datastore_utils.FullRef)
}
