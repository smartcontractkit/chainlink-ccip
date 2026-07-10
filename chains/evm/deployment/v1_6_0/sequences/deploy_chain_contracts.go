package sequences

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/link_token"

	fqbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	offrampbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	onrampbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	noncebindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/nonce_manager"
	rmnproxybindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	pingpongdappops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/ping_pong_dapp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/nonce_manager"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
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
		wethRef, err := evmops.MaybeDeployContract(b, weth.Deploy, chain, contract.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(weth.ContractType, *weth.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := evmops.MaybeDeployContract(b, link.Deploy, chain, contract.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link.ContractType, *link.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := evmops.MaybeDeployContract(b, rmn_remote.Deploy, chain, contract.DeployInput[rmn_remote.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_remote.ContractType, *rmn_remote.Version),
			Args: rmn_remote.ConstructorArgs{
				LocalChainSelector: chain.Selector,
				LegacyRMN:          common.HexToAddress(input.LegacyRMN),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnRemoteRef)

		// Deploy RMNProxy
		rmnProxyRef, err := evmops.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
			Args: rmn_proxy.ConstructorArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnProxyRef)

		// Set the RMNRemote on the RMNProxy
		// Included in case the RMNRemote got deployed but the RMNProxy already existed.
		// In this case, we would not have set the RMNRemote in the constructor.
		// We would need to update the RMN on the existing RMNProxy.
		setRMNReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(rmnProxyRef.Address), rmnproxybindings.NewRMNProxy, rmn_proxy.NewWriteSetRMN, rmn_proxy.SetRMNArgs{
			RMN: common.HexToAddress(rmnRemoteRef.Address),
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, setRMNReport.Output)

		// Deploy Router
		routerRef, err := evmops.MaybeDeployContract(b, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.ContractType, *router.Version),
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
		testRouterRef, err := evmops.MaybeDeployContract(b, router.DeployTestRouter, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.TestRouterContractType, *router.Version),
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
		tokenAdminRegistryRef, err := evmops.MaybeDeployContract(b, token_admin_registry.Deploy, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy NonceManager
		nonceManagerRef, err := evmops.MaybeDeployContract(b, nonce_manager.Deploy, chain, contract.DeployInput[nonce_manager.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(nonce_manager.ContractType, *nonce_manager.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, nonceManagerRef)

		// Deploy FeeQuoter
		feeQuoterRef, err := evmops.MaybeDeployContract(b, fq163ops.Deploy, chain, contract.DeployInput[fq163ops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(fq163ops.ContractType, *fq163ops.Version),
			Args: fq163ops.ConstructorArgs{
				StaticConfig: fqbindings.FeeQuoterStaticConfig{
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
				TokenPriceFeeds:            []fqbindings.FeeQuoterTokenPriceFeedUpdate{},
				TokenTransferFeeConfigArgs: []fqbindings.FeeQuoterTokenTransferFeeConfigArgs{},
				PremiumMultiplierWeiPerEthArgs: []fqbindings.FeeQuoterPremiumMultiplierWeiPerEthArgs{
					{
						PremiumMultiplierWeiPerEth: input.LinkPremiumMultiplier,
						Token:                      common.HexToAddress(linkRef.Address),
					},
					{
						PremiumMultiplierWeiPerEth: input.NativeTokenPremiumMultiplier,
						Token:                      common.HexToAddress(wethRef.Address),
					},
				},
				DestChainConfigArgs: []fqbindings.FeeQuoterDestChainConfigArgs{},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, feeQuoterRef)

		// Deploy OffRamp
		offRampRef, err := evmops.MaybeDeployContract(b, offrampops.Deploy, chain, contract.DeployInput[offrampops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(offrampops.ContractType, *offrampops.Version),
			Args: offrampops.ConstructorArgs{
				StaticConfig: offrampbindings.OffRampStaticConfig{
					ChainSelector:        chain.Selector,
					GasForCallExactCheck: input.GasForCallExactCheck,
					RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
					NonceManager:         common.HexToAddress(nonceManagerRef.Address),
					TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
				},
				DynamicConfig: offrampbindings.OffRampDynamicConfig{
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
		onRampRef, err := evmops.MaybeDeployContract(b, onrampops.Deploy, chain, contract.DeployInput[onrampops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(onrampops.ContractType, *onrampops.Version),
			Args: onrampops.ConstructorArgs{
				StaticConfig: onrampbindings.OnRampStaticConfig{
					ChainSelector:      chain.Selector,
					RmnRemote:          common.HexToAddress(rmnProxyRef.Address),
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
					NonceManager:       common.HexToAddress(nonceManagerRef.Address),
				},
				DynamicConfig: onrampbindings.OnRampDynamicConfig{
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
			pingPongDappRef, err := evmops.MaybeDeployContract(b, pingpongdappops.Deploy, chain, contract.DeployInput[pingpongdappops.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(pingpongdappops.ContractType, *pingpongdappops.Version),
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
			_, err = evmops.ExecuteWrite(b, chain, common.HexToAddress(linkRef.Address), link_token.NewLinkToken, link.NewWriteGrantMintRole, link.GrantMintRoleArgs{
				Minter: chain.DeployerKey.From,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint role for LINK: %w", err)
			}

			// Mint 20 LINK (20 * 10^18 wei) directly to the PingPongDemo contract
			// Retry with backoff for external networks where grant confirmation may take time
			pingPongFundingAmount := new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
			mintArgs := link.MintArgs{
				To:     common.HexToAddress(pingPongDappRef.Address),
				Amount: pingPongFundingAmount,
			}

			maxRetries := 5
			retryDelay := 3 * time.Second
			var mintErr error
			for attempt := 1; attempt <= maxRetries; attempt++ {
				_, mintErr = evmops.ExecuteWrite(b, chain, common.HexToAddress(linkRef.Address), link_token.NewLinkToken, link.NewWriteMint, mintArgs)
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
		_, err = evmops.ExecuteWrite(b, chain, common.HexToAddress(nonceManagerRef.Address), noncebindings.NewNonceManager, nonce_manager.NewWriteApplyAuthorizedCallerUpdates, nonce_manager.AuthorizedCallerArgs{
			AddedCallers: []common.Address{
				common.HexToAddress(offRampRef.Address),
				common.HexToAddress(onRampRef.Address),
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Add Authorized Caller to FQ
		_, err = evmops.ExecuteWrite(b, chain, common.HexToAddress(feeQuoterRef.Address), evmops.BindAs[fqbindings.FeeQuoterInterface](fqbindings.NewFeeQuoter), fq163ops.NewWriteApplyAuthorizedCallerUpdates, fqbindings.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{
				common.HexToAddress(offRampRef.Address),
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
