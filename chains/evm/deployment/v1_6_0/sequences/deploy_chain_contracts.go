package sequences

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	pingpongdappops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/ping_pong_dapp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/nonce_manager"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

func (a *EVMAdapter) DeployChainContracts() *cldf_ops.Sequence[deployops.ContractDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DeployChainContracts
}

// just a wrapper around the v1.0.0 deployer for now
func (a *EVMAdapter) DeployMCMS() *cldf_ops.Sequence[deployops.MCMSDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMDeployer{}
	return evmDeployer.DeployMCMS()
}

// FinalizeDeployMCMS finalizes the deployment of MCM contracts, e.g., by initializing timelock ownership
func (a *EVMAdapter) FinalizeDeployMCMS() *cldf_ops.Sequence[deployops.MCMSDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMDeployer{}
	return evmDeployer.FinalizeDeployMCMS()
}

// Sets a timelock as admin of a newly deployed timelock
func (a *EVMAdapter) GrantAdminRoleToTimelock() *operations.Sequence[deployops.GrantAdminRoleToTimelockConfigPerChainWithSelector, sequences.OnChainOutput, chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMDeployer{}
	return evmDeployer.GrantAdminRoleToTimelock()
}

func (a *EVMAdapter) UpdateMCMSConfig() *operations.Sequence[deployops.UpdateMCMSConfigInputPerChainWithSelector, sequences.OnChainOutput, chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMDeployer{}
	return evmDeployer.UpdateMCMSConfig()
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("1.6.0"),
	"Deploys all required contracts for CCIP 1.6.0 to an EVM chain",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input deployops.ContractDeploymentConfigPerChainWithAddress) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		chain := chains.EVMChains()[input.ChainSelector]

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

		// Deploy WETH
		wethRef, err := contract.MaybeDeployContract(b, weth.Deploy, chain, contract.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(weth.ContractType, *weth.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := contract.MaybeDeployContract(b, link.Deploy, chain, contract.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link.ContractType, *link.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMN 2.1.0, reusing the chain's existing one when the datastore already has it (for
		// example from the v2.0.0 sequence or from DeployAndActivateRMN). The legacy RMNRemote 1.6.0
		// is no longer deployed here, so input.LegacyRMN is unused on this path.
		//
		// Use the Ultra Fast Curse MCMS timelock as the RMN's curse admin when it is available.
		//
		// Unlike the v2.0.0 sequence, a missing timelock is not fatal here: this sequence does not
		// deploy MCMS (see the TODOs above) and the 1.6 flow runs it before MCMS exists, so requiring
		// the timelock would make the ordering unsatisfiable. When it is absent the RMN deploys with
		// no curse admin and can only be cursed by its owner -- there is no fast-curse path until one
		// is added via ApplyAuthorizedCallerUpdates.
		curseAdmins := []common.Address{}
		if ultraFastCurseTimelock, ok := resolveUltraFastCurseTimelock(input.ExistingAddresses, chain.Selector); ok {
			curseAdmins = append(curseAdmins, ultraFastCurseTimelock)
		}

		rmnRef, err := contract.MaybeDeployContract(b, rmnops.Deploy, chain, contract.DeployInput[rmnops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmnops.ContractType, *rmnops.Version),
			ChainSelector:  chain.Selector,
			Args:           rmnops.ConstructorArgs{CurseAdmins: curseAdmins},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnRef)

		// Deploy RMNProxy
		rmnProxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
			ChainSelector:  chain.Selector,
			Args: rmn_proxy.ConstructorArgs{
				RMN: common.HexToAddress(rmnRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnProxyRef)

		// Point the RMNProxy at the RMN, for the case where the RMN was deployed but the RMNProxy
		// already existed and so did not get it via the constructor. Guarded by a read: repointing a
		// live proxy switches every RMN consumer on the chain at once and drops all active curses,
		// since curse state does not carry across RMN instances, so it must not fire on a no-op run.
		activeRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.GetRMN, chain, contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(rmnProxyRef.Address),
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if activeRMNReport.Output != common.HexToAddress(rmnRef.Address) {
			setRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.SetRMN, chain, contract.FunctionInput[rmn_proxy.SetRMNArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(rmnProxyRef.Address),
				Args: rmn_proxy.SetRMNArgs{
					RMN: common.HexToAddress(rmnRef.Address),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			writes = append(writes, setRMNReport.Output)
		}

		// Deploy Router
		routerRef, err := contract.MaybeDeployContract(b, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.ContractType, *router.Version),
			ChainSelector:  chain.Selector,
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, routerRef)

		// Deploy Test Router
		testRouterRef, err := contract.MaybeDeployContract(b, router.DeployTestRouter, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.TestRouterContractType, *router.Version),
			ChainSelector:  chain.Selector,
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, testRouterRef)

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := contract.MaybeDeployContract(b, token_admin_registry.Deploy, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy NonceManager
		nonceManagerRef, err := contract.MaybeDeployContract(b, nonce_manager.Deploy, chain, contract.DeployInput[nonce_manager.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(nonce_manager.ContractType, *nonce_manager.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, nonceManagerRef)

		// Deploy FeeQuoter
		feeQuoterRef, err := contract.MaybeDeployContract(b, fq163ops.Deploy, chain, contract.DeployInput[fq163ops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(fq163ops.ContractType, *fq163ops.Version),
			ChainSelector:  chain.Selector,
			Args: fq163ops.ConstructorArgs{
				StaticConfig: fq163ops.StaticConfig{
					MaxFeeJuelsPerMsg:            input.MaxFeeJuelsPerMsg,
					LinkToken:                    common.HexToAddress(linkRef.Address),
					TokenPriceStalenessThreshold: input.TokenPriceStalenessThreshold,
				},
				PriceUpdaters: []common.Address{
					// TODO: Add Timelock here when MCMS support is needed.
					chain.DeployerKey.From,
				},
				FeeTokens: []common.Address{
					common.HexToAddress(linkRef.Address),
					common.HexToAddress(wethRef.Address),
				},
				TokenPriceFeeds:            []fq163ops.TokenPriceFeedUpdate{},
				TokenTransferFeeConfigArgs: []fq163ops.TokenTransferFeeConfigArgs{},
				PremiumMultiplierWeiPerEthArgs: []fq163ops.PremiumMultiplierWeiPerEthArgs{
					{
						PremiumMultiplierWeiPerEth: input.LinkPremiumMultiplier,
						Token:                      common.HexToAddress(linkRef.Address),
					},
					{
						PremiumMultiplierWeiPerEth: input.NativeTokenPremiumMultiplier,
						Token:                      common.HexToAddress(wethRef.Address),
					},
				},
				DestChainConfigArgs: []fq163ops.DestChainConfigArgs{},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, feeQuoterRef)

		// Deploy OffRamp
		offRampRef, err := contract.MaybeDeployContract(b, offrampops.Deploy, chain, contract.DeployInput[offrampops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(offrampops.ContractType, *offrampops.Version),
			ChainSelector:  chain.Selector,
			Args: offrampops.ConstructorArgs{
				StaticConfig: offrampops.StaticConfig{
					ChainSelector:        chain.Selector,
					GasForCallExactCheck: input.GasForCallExactCheck,
					RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
					NonceManager:         common.HexToAddress(nonceManagerRef.Address),
					TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
				},
				DynamicConfig: offrampops.DynamicConfig{
					FeeQuoter:                               common.HexToAddress(feeQuoterRef.Address),
					PermissionLessExecutionThresholdSeconds: input.PermissionLessExecutionThresholdSeconds,
					MessageInterceptor:                      common.HexToAddress(input.MessageInterceptor),
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OffRamp: %w", err)
		}
		addresses = append(addresses, offRampRef)

		// Deploy OnRamp
		onRampRef, err := contract.MaybeDeployContract(b, onrampops.Deploy, chain, contract.DeployInput[onrampops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(onrampops.ContractType, *onrampops.Version),
			ChainSelector:  chain.Selector,
			Args: onrampops.ConstructorArgs{
				StaticConfig: onrampops.StaticConfig{
					ChainSelector:      chain.Selector,
					RmnRemote:          common.HexToAddress(rmnProxyRef.Address),
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
					NonceManager:       common.HexToAddress(nonceManagerRef.Address),
				},
				DynamicConfig: onrampops.DynamicConfig{
					FeeQuoter:     common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator: chain.DeployerKey.From,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OnRamp: %w", err)
		}
		addresses = append(addresses, onRampRef)

		// Deploy Ping Pong Dapp (optional - only when DeployPingPongDapp is true)
		if input.DeployPingPongDapp {
			pingPongDappRef, err := contract.MaybeDeployContract(b, pingpongdappops.Deploy, chain, contract.DeployInput[pingpongdappops.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(pingpongdappops.ContractType, *pingpongdappops.Version),
				ChainSelector:  chain.Selector,
				Args: pingpongdappops.ConstructorArgs{
					Router:   common.HexToAddress(routerRef.Address),
					FeeToken: common.HexToAddress(linkRef.Address),
				},
			}, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy Ping Pong Dapp: %w", err)
			}
			addresses = append(addresses, pingPongDappRef)

			// Fund Ping Pong Dapp with LINK tokens for cross-chain message fees
			// First, grant mint role to the deployer
			_, err = cldf_ops.ExecuteOperation(b, link.GrantMintRole, chain, contract.FunctionInput[link.GrantMintRoleArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(linkRef.Address),
				Args: link.GrantMintRoleArgs{
					Minter: chain.DeployerKey.From,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint role for LINK: %w", err)
			}

			// Mint 20 LINK (20 * 10^18 wei) directly to the PingPongDemo contract
			// Retry with backoff for external networks where grant confirmation may take time
			pingPongFundingAmount := new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
			mintInput := contract.FunctionInput[link.MintArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(linkRef.Address),
				Args: link.MintArgs{
					To:     common.HexToAddress(pingPongDappRef.Address),
					Amount: pingPongFundingAmount,
				},
			}

			maxRetries := 5
			retryDelay := 3 * time.Second
			var mintErr error
			for attempt := 1; attempt <= maxRetries; attempt++ {
				_, mintErr = cldf_ops.ExecuteOperation(b, link.Mint, chain, mintInput)
				if mintErr == nil {
					break
				}
				// Only retry on SenderNotMinter error (minter role not yet confirmed)
				if !strings.Contains(mintErr.Error(), "SenderNotMinter") {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to mint LINK to Ping Pong Dapp: %w", mintErr)
				}
				if attempt < maxRetries {
					b.Logger.Warnf("Mint failed with SenderNotMinter (attempt %d/%d), retrying in %v...", attempt, maxRetries, retryDelay)
					time.Sleep(retryDelay)
				}
			}
			if mintErr != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to mint LINK to Ping Pong Dapp after %d retries: %w", maxRetries, mintErr)
			}
		}

		// Add Authorized Caller to NonceManager
		_, err = cldf_ops.ExecuteOperation(b, nonce_manager.ApplyAuthorizedCallerUpdates, chain, contract.FunctionInput[nonce_manager.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(nonceManagerRef.Address),
			Args: nonce_manager.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(offRampRef.Address),
					common.HexToAddress(onRampRef.Address),
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Add Authorized Caller to FQ
		_, err = cldf_ops.ExecuteOperation(b, fqops.ApplyAuthorizedCallerUpdates, chain, contract.FunctionInput[fqops.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(feeQuoterRef.Address),
			Args: fqops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(offRampRef.Address),
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

// resolveUltraFastCurseTimelock returns the Ultra Fast Curse RBACTimelock for the chain, reporting
// false when it is not present or not a usable EVM address. Callers treat absence as "no curse
// admin" rather than an error, since this sequence can run before MCMS is deployed.
func resolveUltraFastCurseTimelock(refs []datastore.AddressRef, chainSelector uint64) (common.Address, bool) {
	ref := datastore_utils.GetAddressRef(
		refs, chainSelector, common_utils.RBACTimelock, mcms_ops.MCMSVersion, common_utils.UltraFastCurseMCMSQualifier)
	if ref.Address == "" {
		return common.Address{}, false
	}
	addr, err := evm_datastore_utils.ToEVMAddress(ref)
	if err != nil {
		return common.Address{}, false
	}

	return addr, true
}
