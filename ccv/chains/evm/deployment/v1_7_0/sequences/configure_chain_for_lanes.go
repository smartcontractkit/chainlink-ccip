package sequences

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_offchain "github.com/smartcontractkit/chainlink-deployments-framework/offchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
)

var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures an EVM chain as a source & destination for multiple remote chains",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input adapters.ConfigureChainForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Create inputs for each operation
		offRampArgs := make([]offramp.SourceChainConfigArgs, 0, len(input.RemoteChains))
		onRampArgs := make([]onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter.DestChainConfigArgs, 0, len(input.RemoteChains))
		gasPriceUpdates := make([]fee_quoter.GasPriceUpdate, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		destChainSelectorsPerExecutor := make(map[common.Address][]executor.RemoteChainConfigArgs)
		for remoteSelector, remoteConfig := range input.RemoteChains {
			defaultInboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultInboundCCVs))
			for _, ccv := range remoteConfig.DefaultInboundCCVs {
				defaultInboundCCVs = append(defaultInboundCCVs, common.HexToAddress(ccv))
			}
			laneMandatedInboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedInboundCCVs))
			for _, ccv := range remoteConfig.LaneMandatedInboundCCVs {
				laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, common.HexToAddress(ccv))
			}
			// Left-pad remoteConfig.OnRamps with zeros to the left to match the address bytes length
			onRamps := make([][]byte, 0, len(remoteConfig.OnRamps))
			for _, onRamp := range remoteConfig.OnRamps {
				onRamps = append(onRamps, common.LeftPadBytes(onRamp, 32))
			}
			offRampArgs = append(offRampArgs, offramp.SourceChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				SourceChainSelector: remoteSelector,
				IsEnabled:           remoteConfig.AllowTrafficFrom,
				OnRamps:             onRamps,
				DefaultCCVs:         defaultInboundCCVs,
				LaneMandatedCCVs:    laneMandatedInboundCCVs,
			})
			defaultOutboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultOutboundCCVs))
			for _, ccv := range remoteConfig.DefaultOutboundCCVs {
				defaultOutboundCCVs = append(defaultOutboundCCVs, common.HexToAddress(ccv))
			}
			laneMandatedOutboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedOutboundCCVs))
			for _, ccv := range remoteConfig.LaneMandatedOutboundCCVs {
				laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, common.HexToAddress(ccv))
			}
			onRampArgs = append(onRampArgs, onramp.DestChainConfigArgs{
				Router:               common.HexToAddress(input.Router),
				DestChainSelector:    remoteSelector,
				AddressBytesLength:   remoteConfig.AddressBytesLength,
				BaseExecutionGasCost: remoteConfig.BaseExecutionGasCost,
				DefaultCCVs:          defaultOutboundCCVs,
				LaneMandatedCCVs:     laneMandatedOutboundCCVs,
				DefaultExecutor:      common.HexToAddress(remoteConfig.DefaultExecutor), // The proxy address
				OffRamp:              remoteConfig.OffRamp,
			})
			feeQuoterArgs = append(feeQuoterArgs, fee_quoter.DestChainConfigArgs{
				DestChainSelector: remoteSelector,
				DestChainConfig:   remoteConfig.FeeQuoterDestChainConfig,
			})
			gasPriceUpdates = append(gasPriceUpdates, fee_quoter.GasPriceUpdate{
				DestChainSelector: remoteSelector,
				UsdPerUnitGas:     remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas,
			})
			onRampAdds = append(onRampAdds, router.OnRamp{
				DestChainSelector: remoteSelector,
				OnRamp:            common.HexToAddress(input.OnRamp),
			})
			offRampAdds = append(offRampAdds, router.OffRamp{
				SourceChainSelector: remoteSelector,
				OffRamp:             common.HexToAddress(input.OffRamp),
			})
			defaultExecutor := common.HexToAddress(remoteConfig.DefaultExecutor)
			getTargetReport, err := cldf_ops.ExecuteOperation(b, proxy.GetTarget, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       defaultExecutor,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get target address of Executor(%s) on chain %s: %w", defaultExecutor, chain, err)
			}
			if destChainSelectorsPerExecutor[getTargetReport.Output] == nil {
				destChainSelectorsPerExecutor[getTargetReport.Output] = []executor.RemoteChainConfigArgs{}
			}
			destChainSelectorsPerExecutor[getTargetReport.Output] = append(destChainSelectorsPerExecutor[getTargetReport.Output], executor.RemoteChainConfigArgs{
				DestChainSelector: remoteSelector,
				Config:            remoteConfig.ExecutorDestChainConfig,
			})
		}

		// ApplySourceChainConfigUpdates on OffRamp
		offRampReport, err := cldf_ops.ExecuteOperation(b, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.OffRamp),
			Args:          offRampArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to OffRamp(%s) on chain %s: %w", input.OffRamp, chain, err)
		}
		writes = append(writes, offRampReport.Output)

		// ApplyDestChainConfigUpdates on OnRamp
		onRampReport, err := cldf_ops.ExecuteOperation(b, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.OnRamp),
			Args:          onRampArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to OnRamp(%s) on chain %s: %w", input.OnRamp, chain, err)
		}
		writes = append(writes, onRampReport.Output)

		// ApplyDestChainUpdates on each Executor
		for executorAddr, destChainSelectorsToAdd := range destChainSelectorsPerExecutor {
			executorReport, err := cldf_ops.ExecuteOperation(b, executor.ApplyDestChainUpdates, chain, contract.FunctionInput[executor.ApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       executorAddr,
				Args: executor.ApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: destChainSelectorsToAdd,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to Executor(%s) on chain %s: %w", executorAddr, chain, err)
			}
			writes = append(writes, executorReport.Output)
		}

		// ApplyDestChainConfigUpdates on FeeQuoter
		feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.FeeQuoter),
			Args:          feeQuoterArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
		}
		writes = append(writes, feeQuoterReport.Output)

		// UpdatePrices on FeeQuoter (gas prices only, as these are per dest chain)
		feeQuoterUpdatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract.FunctionInput[fee_quoter.PriceUpdates]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.FeeQuoter),
			Args: fee_quoter.PriceUpdates{
				GasPriceUpdates: gasPriceUpdates,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to update gas prices on FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
		}
		writes = append(writes, feeQuoterUpdatePricesReport.Output)

		// ApplyRampUpdates on Router
		routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.Router),
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
		batchOps := []mcms_types.BatchOperation{batchOp}

		for _, committeeVerifier := range input.CommitteeVerifiers {
			committeeVerifierReport, err := cldf_ops.ExecuteSequence(b, ConfigureCommitteeVerifierForLanes, chains, ConfigureCommitteeVerifierForLanesInput{
				ChainSelector:           chain.Selector,
				Router:                  input.Router,
				CommitteeVerifierConfig: committeeVerifier,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure committee verifier for lanes: %w", err)
			}
			batchOps = append(batchOps, committeeVerifierReport.Output.BatchOps...)
		}

		// Initialize contract metadata slice
		contractMetadata := make([]datastore.ContractMetadata, 0)

		// Collect fee token metadata and add to FeeQuoter metadata
		feeTokens, err := collectFeeTokens(b, chain, input.FeeQuoter, chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Add FeeQuoter metadata with fee tokens
		contractMetadata = append(contractMetadata, datastore.ContractMetadata{
			Address:       input.FeeQuoter,
			ChainSelector: chain.Selector,
			Metadata: map[string]interface{}{
				"configured":    true,
				"test_metadata": "fee_quoter_configured",
				"feeTokens":     feeTokens,
			},
		})

		// Collect CommitteeVerifier signature config metadata for each CommitteeVerifier
		for _, committeeVerifier := range input.CommitteeVerifiers {
			var committeeVerifierAddr string
			for _, addr := range committeeVerifier.CommitteeVerifier {
				if addr.Type == datastore.ContractType(committee_verifier.ContractType) {
					committeeVerifierAddr = addr.Address
					break
				}
			}
			if committeeVerifierAddr == "" {
				continue // Skip if we can't find the CommitteeVerifier address
			}

			// Build pending signature configs from input
			// Note: We don't have removals in the current input structure, so SourceChainSelectorsToRemove is empty
			pendingSignatureConfigs := make([]committee_verifier.SignatureConfig, 0, len(committeeVerifier.RemoteChains))
			for remoteSelector, remoteConfig := range committeeVerifier.RemoteChains {
				signers := make([]common.Address, 0, len(remoteConfig.SignatureConfig.Signers))
				for _, signer := range remoteConfig.SignatureConfig.Signers {
					signers = append(signers, common.HexToAddress(signer))
				}
				pendingSignatureConfigs = append(pendingSignatureConfigs, committee_verifier.SignatureConfig{
					SourceChainSelector: remoteSelector,
					Threshold:           remoteConfig.SignatureConfig.Threshold,
					Signers:             signers,
				})
			}

			pendingSignatureConfigArgs := committee_verifier.SignatureConfigArgs{
				SourceChainSelectorsToRemove: []uint64{}, // No removals in current input structure
				SignatureConfigUpdates:       pendingSignatureConfigs,
			}

			committeeVerifierMetadata, err := collectCommitteeVerifierSignatureConfigs(b, chain, committeeVerifierAddr, chain.Selector, []datastore.ContractMetadata{}, input.OffchainClient, pendingSignatureConfigArgs)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to collect CommitteeVerifier signature configs: %w", err)
			}
			contractMetadata = append(contractMetadata, committeeVerifierMetadata...)
		}

		// Collect Router metadata (onRamps and offRamps)
		// Note: OnRamp destination chain config metadata is collected within collectRouterMetadata
		// Project pending changes: merge ramp updates with current state
		routerMetadata, err := collectRouterMetadata(
			b,
			chain,
			input.Router,
			chain.Selector,
			[]datastore.ContractMetadata{},
			router.ApplyRampsUpdatesArgs{
				OnRampUpdates:  onRampAdds,
				OffRampRemoves: []router.OffRamp{}, // removals should be processed by a separate sequence responsible for disconnecting lanes
				OffRampAdds:    offRampAdds,
			},
			input.OnRamp,  // OnRamp address that pending configs apply to
			input.OffRamp, // OffRamp address that pending configs apply to
			onRampArgs,
			offRampArgs,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to collect Router metadata: %w", err)
		}
		contractMetadata = append(contractMetadata, routerMetadata...)

		// Collect TokenAdminRegistry metadata
		// Get TokenAdminRegistry address from OnRamp's static config
		getOnRampStaticConfigReport, err := cldf_ops.ExecuteOperation(b, onramp.GetStaticConfig, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.OnRamp),
			Args:          nil,
		})
		if err != nil {
			// Log error but don't fail - TokenAdminRegistry might not be available
		} else {
			tokenAdminRegistryAddr := getOnRampStaticConfigReport.Output.TokenAdminRegistry.Hex()
			if tokenAdminRegistryAddr != "" && tokenAdminRegistryAddr != "0x0000000000000000000000000000000000000000" {
				tokenAdminRegistryMetadata, err := collectTokenAdminRegistryMetadata(b, chain, chain.Selector, tokenAdminRegistryAddr)
				if err != nil {
					// Log error but don't fail - TokenAdminRegistry metadata collection might fail
				} else {
					contractMetadata = append(contractMetadata, tokenAdminRegistryMetadata)
				}
			}
		}

		return sequences.OnChainOutput{
			Metadata: sequences.Metadata{
				Contracts: contractMetadata,
			},
			BatchOps: batchOps,
		}, nil
	})

// collectFeeTokens collects metadata for all fee tokens from the FeeQuoter and returns them as a list
func collectFeeTokens(
	b cldf_ops.Bundle,
	chain evm.Chain,
	feeQuoterAddr string,
	chainSelector uint64,
) ([]interface{}, error) {
	// Read fee tokens from the FeeQuoter
	getFeeTokensReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetFeeTokens, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(feeQuoterAddr),
		Args:          nil,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read fee tokens from FeeQuoter(%s) on chain %s: %w", feeQuoterAddr, chain, err)
	}

	feeTokens := make([]interface{}, 0, len(getFeeTokensReport.Output))
	// Read metadata for each fee token
	for _, tokenAddr := range getFeeTokensReport.Output {
		// Read name
		nameReport, err := cldf_ops.ExecuteOperation(b, erc20.Name, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read name of fee token(%s) on chain %s: %w", tokenAddr, chain, err)
		}

		// Read symbol
		symbolReport, err := cldf_ops.ExecuteOperation(b, erc20.Symbol, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read symbol of fee token(%s) on chain %s: %w", tokenAddr, chain, err)
		}

		// Read decimals
		decimalsReport, err := cldf_ops.ExecuteOperation(b, erc20.Decimals, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read decimals of fee token(%s) on chain %s: %w", tokenAddr, chain, err)
		}

		feeTokens = append(feeTokens, map[string]interface{}{
			"address":       tokenAddr.Hex(),
			"chainSelector": chainSelector,
			"name":          nameReport.Output,
			"symbol":        symbolReport.Output,
			"decimals":      decimalsReport.Output,
		})
	}

	return feeTokens, nil
}

func collectCommitteeVerifierSignatureConfigs(
	b cldf_ops.Bundle,
	chain evm.Chain,
	committeeVerifierAddr string,
	chainSelector uint64,
	contractMetadata []datastore.ContractMetadata,
	offchainClient interface{}, // Optional: cldf_offchain.Client for fetching CSA keys
	pendingSignatureConfigArgs committee_verifier.SignatureConfigArgs, // Pending signature config changes
) ([]datastore.ContractMetadata, error) {
	// Read current signature configs from the contract.
	getAllConfigsReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetAllSignatureConfigs, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(committeeVerifierAddr),
		Args:          nil,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read signature configs from CommitteeVerifier(%s) on chain %s: %w", committeeVerifierAddr, chain, err)
	}

	// Build a map of current signature configs.
	currentConfigs := make(map[uint64]committee_verifier.SignatureConfig)
	for _, cfg := range getAllConfigsReport.Output {
		currentConfigs[cfg.SourceChainSelector] = cfg
	}

	// Apply pending removals: remove selectors that are being removed
	for _, selectorToRemove := range pendingSignatureConfigArgs.SourceChainSelectorsToRemove {
		if _, exists := currentConfigs[selectorToRemove]; exists {
			delete(currentConfigs, selectorToRemove)
		}
		// If selector doesn't exist in current state, it's a NOOP (as per contract behavior)
	}

	// Apply pending updates/adds: override existing or add new configs
	for _, pendingConfig := range pendingSignatureConfigArgs.SignatureConfigUpdates {
		// Override existing or add new
		currentConfigs[pendingConfig.SourceChainSelector] = pendingConfig
	}

	// Convert to metadata
	for selector, cfg := range currentConfigs {
		signersHex := make([]string, len(cfg.Signers))
		csaKeys := make([]string, 0, len(cfg.Signers))

		for i, signer := range cfg.Signers {
			signersHex[i] = signer.Hex()

			// Optionally fetch CSA key from job distributor if offchain client is provided
			if offchainClient != nil {
				if client, ok := offchainClient.(cldf_offchain.Client); ok {
					csaKey, err := getCSAKeyFromOnchainPublicKey(
						b.GetContext(),
						client,
						signer.Hex(),
						chainSelector,
					)
					if err == nil {
						csaKeys = append(csaKeys, csaKey)
						b.Logger.Debugw("Successfully fetched CSA key for signer", "signer", signer.Hex(), "csaKey", csaKey)
					} else {
						// Log warning but continue - CSA key fetch is optional
						b.Logger.Warnw("Failed to fetch CSA key for signer", "signer", signer.Hex(), "error", err)
						csaKeys = append(csaKeys, "")
					}
				} else {
					// Type assertion failed - offchainClient doesn't implement cldf_offchain.Client
					b.Logger.Debugw("Offchain client type assertion failed - CSA keys won't be fetched", "signer", signer.Hex())
				}
			}
		}

		metadata := map[string]interface{}{
			"sourceChainSelector": selector,
			"threshold":           cfg.Threshold,
			"signers":             signersHex,
		}

		// Add CSA keys to metadata if they were fetched
		// Only add if we successfully fetched CSA keys for all signers (no empty strings)
		// Note: CSA keys will only be present when a job distributor (OffchainClient) is configured
		// in the deployment environment. Unit tests typically don't include a job distributor.
		allCSAKeysFetched := len(csaKeys) == len(signersHex) && len(csaKeys) > 0
		if allCSAKeysFetched {
			// Check that all CSA keys are non-empty
			for _, key := range csaKeys {
				if key == "" {
					allCSAKeysFetched = false
					break
				}
			}
		}
		if allCSAKeysFetched {
			metadata["csaKeys"] = csaKeys
		}

		contractMetadata = append(contractMetadata, datastore.ContractMetadata{
			Address:       committeeVerifierAddr,
			ChainSelector: chainSelector,
			Metadata:      metadata,
		})
	}

	return contractMetadata, nil
}

// collectRouterMetadata collects metadata for all onRamps and offRamps from the Router.
// It projects the state after pending ramp updates by merging pending changes with current state.
func collectRouterMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	routerAddr string,
	chainSelector uint64,
	contractMetadata []datastore.ContractMetadata,
	pendingRampUpdates router.ApplyRampsUpdatesArgs, // Pending ramp updates to apply
	onRampAddr string, // OnRamp address that pending configs apply to
	offRampAddr string, // OffRamp address that pending configs apply to
	pendingOnRampConfigs []onramp.DestChainConfigArgs, // Pending OnRamp dest chain configs
	pendingOffRampConfigs []offramp.SourceChainConfigArgs, // Pending OffRamp source chain configs
) ([]datastore.ContractMetadata, error) {
	// Read all offRamps from the Router
	getOffRampsReport, err := cldf_ops.ExecuteOperation(b, router.GetOffRamps, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(routerAddr),
		Args:          nil,
	})
	if err != nil {
		return contractMetadata, fmt.Errorf("failed to read offRamps from Router(%s) on chain %s: %w", routerAddr, chain, err)
	}

	// Group offRamps by sourceChainSelector (current state)
	offRampsByChain := make(map[uint64][]string)
	uniqueChainSelectors := make(map[uint64]struct{})
	for _, offRamp := range getOffRampsReport.Output {
		sourceChainSelector := offRamp.SourceChainSelector
		offRampsByChain[sourceChainSelector] = append(offRampsByChain[sourceChainSelector], offRamp.OffRamp.Hex())
		uniqueChainSelectors[sourceChainSelector] = struct{}{}
	}

	// Get onRamp for each unique sourceChainSelector (current state)
	onRampsByChain := make(map[uint64]string)
	for sourceChainSelector := range uniqueChainSelectors {
		getOnRampReport, err := cldf_ops.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chainSelector,
			Address:       common.HexToAddress(routerAddr),
			Args:          sourceChainSelector,
		})
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to read onRamp for chain selector %d from Router(%s) on chain %s: %w", sourceChainSelector, routerAddr, chain, err)
		}
		onRampsByChain[sourceChainSelector] = getOnRampReport.Output.Hex()
	}

	// Apply pending OnRamp updates: update existing or add new
	for _, onRampUpdate := range pendingRampUpdates.OnRampUpdates {
		destChainSelector := onRampUpdate.DestChainSelector
		onRampAddr := onRampUpdate.OnRamp.Hex()
		onRampsByChain[destChainSelector] = onRampAddr
		uniqueChainSelectors[destChainSelector] = struct{}{}
	}

	// Apply pending OffRamp removes
	for _, offRampRemove := range pendingRampUpdates.OffRampRemoves {
		sourceChainSelector := offRampRemove.SourceChainSelector
		offRampAddr := offRampRemove.OffRamp.Hex()
		// Remove this offRamp from the list for this source chain selector
		if offRamps, exists := offRampsByChain[sourceChainSelector]; exists {
			newOffRamps := make([]string, 0, len(offRamps))
			for _, existing := range offRamps {
				if existing != offRampAddr {
					newOffRamps = append(newOffRamps, existing)
				}
			}
			offRampsByChain[sourceChainSelector] = newOffRamps
		}
	}

	// Apply pending OffRamp adds
	for _, offRampAdd := range pendingRampUpdates.OffRampAdds {
		sourceChainSelector := offRampAdd.SourceChainSelector
		offRampAddr := offRampAdd.OffRamp.Hex()
		// Check if this offRamp is already in the list
		found := false
		if existingOffRamps, exists := offRampsByChain[sourceChainSelector]; exists {
			for _, existing := range existingOffRamps {
				if existing == offRampAddr {
					found = true
					break
				}
			}
		}
		if !found {
			offRampsByChain[sourceChainSelector] = append(offRampsByChain[sourceChainSelector], offRampAddr)
		}
		uniqueChainSelectors[sourceChainSelector] = struct{}{}
	}

	// Build sets of pending ramp addresses (ramps being updated/added)
	pendingOnRampAddrs := make(map[string]struct{})
	for _, onRampUpdate := range pendingRampUpdates.OnRampUpdates {
		pendingOnRampAddrs[onRampUpdate.OnRamp.Hex()] = struct{}{}
	}

	pendingOffRampAddrs := make(map[string]struct{})
	for _, offRampAdd := range pendingRampUpdates.OffRampAdds {
		pendingOffRampAddrs[offRampAdd.OffRamp.Hex()] = struct{}{}
	}

	// Build maps of pending configs by chain selector for the input OnRamp/OffRamp
	// These configs apply to onRampAddr and offRampAddr respectively
	pendingOnRampConfigsBySelector := make(map[uint64]onramp.DestChainConfigArgs)
	for _, config := range pendingOnRampConfigs {
		pendingOnRampConfigsBySelector[config.DestChainSelector] = config
	}

	pendingOffRampConfigsBySelector := make(map[uint64]offramp.SourceChainConfigArgs)
	for _, config := range pendingOffRampConfigs {
		pendingOffRampConfigsBySelector[config.SourceChainSelector] = config
	}

	// Convert to string keys for JSON serialization
	onRampsMap := make(map[string]string)
	for selector, onRampAddr := range onRampsByChain {
		onRampsMap[fmt.Sprintf("%d", selector)] = onRampAddr
	}

	offRampsMap := make(map[string][]string)
	for selector, offRampAddrs := range offRampsByChain {
		offRampsMap[fmt.Sprintf("%d", selector)] = offRampAddrs
	}

	contractMetadata = append(contractMetadata, datastore.ContractMetadata{
		Address:       routerAddr,
		ChainSelector: chainSelector,
		Metadata: map[string]interface{}{
			"onRamps":  onRampsMap,
			"offRamps": offRampsMap,
		},
	})

	// Collect OnRamp metadata for each unique OnRamp
	uniqueOnRamps := make(map[string]struct{})
	for _, onRampAddrLoop := range onRampsByChain {
		uniqueOnRamps[onRampAddrLoop] = struct{}{}
	}

	for onRampAddrLoop := range uniqueOnRamps {
		// Check if this OnRamp is pending (being updated)
		_, isPending := pendingOnRampAddrs[onRampAddrLoop]
		var pendingConfigsForOnRamp map[uint64]onramp.DestChainConfigArgs
		if isPending && onRampAddrLoop == onRampAddr {
			// This is the OnRamp being updated, use its pending configs
			pendingConfigsForOnRamp = pendingOnRampConfigsBySelector
		}

		onRampMetadata, err := collectOnRampDestChainConfigsMetadata(b, chain, onRampAddrLoop, chainSelector, isPending, pendingConfigsForOnRamp)
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to collect OnRamp metadata for %s: %w", onRampAddrLoop, err)
		}
		contractMetadata = append(contractMetadata, onRampMetadata)
	}

	// Collect OffRamp metadata for each unique OffRamp
	uniqueOffRamps := make(map[string]struct{})
	for _, offRampAddrs := range offRampsByChain {
		for _, offRampAddrLoop := range offRampAddrs {
			uniqueOffRamps[offRampAddrLoop] = struct{}{}
		}
	}

	for offRampAddrLoop := range uniqueOffRamps {
		// Check if this OffRamp is pending (being added)
		_, isPending := pendingOffRampAddrs[offRampAddrLoop]
		var pendingConfigsForOffRamp map[uint64]offramp.SourceChainConfigArgs
		if isPending && offRampAddrLoop == offRampAddr {
			// This is the OffRamp being updated, use its pending configs
			pendingConfigsForOffRamp = pendingOffRampConfigsBySelector
		}

		offRampMetadata, err := collectOffRampSourceChainConfigsMetadata(b, chain, offRampAddrLoop, chainSelector, isPending, pendingConfigsForOffRamp)
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to collect OffRamp metadata for %s: %w", offRampAddrLoop, err)
		}
		contractMetadata = append(contractMetadata, offRampMetadata)
	}

	return contractMetadata, nil
}

// collectOnRampDestChainConfigsMetadata collects all destination chain configs from an OnRamp and returns metadata.
// If isPending is true and pendingConfigs is provided, it uses those instead of querying chain state.
func collectOnRampDestChainConfigsMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	onRampAddr string,
	chainSelector uint64,
	isPending bool,
	pendingConfigs map[uint64]onramp.DestChainConfigArgs, // Pending configs by destChainSelector (only used if isPending is true)
) (datastore.ContractMetadata, error) {
	var destChainConfigsList []map[string]interface{}

	// Always query chain state first to get existing configs
	getAllConfigsReport, err := cldf_ops.ExecuteOperation(b, onramp.GetAllDestChainConfigs, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(onRampAddr),
		Args:          nil,
	})
	if err != nil {
		return datastore.ContractMetadata{}, fmt.Errorf("failed to read all dest chain configs from OnRamp(%s) on chain %s: %w", onRampAddr, chain, err)
	}

	// Build a map of existing configs by destChainSelector
	existingConfigsBySelector := make(map[uint64]map[string]interface{})
	for i, destChainSelector := range getAllConfigsReport.Output.DestChainSelectors {
		config := getAllConfigsReport.Output.DestChainConfigs[i]

		// Convert CCV addresses to hex strings
		defaultCCVsHex := make([]string, len(config.DefaultCCVs))
		for j, ccv := range config.DefaultCCVs {
			defaultCCVsHex[j] = ccv.Hex()
		}

		laneMandatedCCVsHex := make([]string, len(config.LaneMandatedCCVs))
		for j, ccv := range config.LaneMandatedCCVs {
			laneMandatedCCVsHex[j] = ccv.Hex()
		}

		existingConfigsBySelector[destChainSelector] = map[string]interface{}{
			"destChainSelector":         destChainSelector,
			"router":                    config.Router.Hex(),
			"messageNumber":             config.MessageNumber,
			"addressBytesLength":        config.AddressBytesLength,
			"tokenReceiverAllowed":      config.TokenReceiverAllowed,
			"messageNetworkFeeUSDCents": config.MessageNetworkFeeUSDCents,
			"tokenNetworkFeeUSDCents":   config.TokenNetworkFeeUSDCents,
			"baseExecutionGasCost":      config.BaseExecutionGasCost,
			"defaultExecutor":           config.DefaultExecutor.Hex(),
			"laneMandatedCCVs":          laneMandatedCCVsHex,
			"defaultCCVs":               defaultCCVsHex,
			"offRamp":                   fmt.Sprintf("0x%x", config.OffRamp),
		}
	}

	if isPending {
		if pendingConfigs != nil && len(pendingConfigs) > 0 {
			// Merge pending configs with existing configs (pending overrides existing)
			for destChainSelector, pendingConfig := range pendingConfigs {
				// Convert pending config to metadata format
				defaultCCVsHex := make([]string, len(pendingConfig.DefaultCCVs))
				for j, ccv := range pendingConfig.DefaultCCVs {
					defaultCCVsHex[j] = ccv.Hex()
				}

				laneMandatedCCVsHex := make([]string, len(pendingConfig.LaneMandatedCCVs))
				for j, ccv := range pendingConfig.LaneMandatedCCVs {
					laneMandatedCCVsHex[j] = ccv.Hex()
				}

				// Override or add pending config
				existingConfigsBySelector[destChainSelector] = map[string]interface{}{
					"destChainSelector":         destChainSelector,
					"router":                    pendingConfig.Router.Hex(),
					"messageNumber":             0, // Not in pending config, default to 0
					"addressBytesLength":        pendingConfig.AddressBytesLength,
					"tokenReceiverAllowed":      false, // Not in pending config, default to false
					"messageNetworkFeeUSDCents": 0,     // Not in pending config, default to 0
					"tokenNetworkFeeUSDCents":   0,     // Not in pending config, default to 0
					"baseExecutionGasCost":      pendingConfig.BaseExecutionGasCost,
					"defaultExecutor":           pendingConfig.DefaultExecutor.Hex(),
					"laneMandatedCCVs":          laneMandatedCCVsHex,
					"defaultCCVs":               defaultCCVsHex,
					"offRamp":                   fmt.Sprintf("0x%x", pendingConfig.OffRamp),
				}
			}
		}
		// Convert map to list
		destChainConfigsList = make([]map[string]interface{}, 0, len(existingConfigsBySelector))
		for _, config := range existingConfigsBySelector {
			destChainConfigsList = append(destChainConfigsList, config)
		}
	} else {
		// Not pending - use existing configs as-is
		destChainConfigsList = make([]map[string]interface{}, 0, len(existingConfigsBySelector))
		for _, config := range existingConfigsBySelector {
			destChainConfigsList = append(destChainConfigsList, config)
		}
	}

	return datastore.ContractMetadata{
		Address:       onRampAddr,
		ChainSelector: chainSelector,
		Metadata: map[string]interface{}{
			"destChainConfigs": destChainConfigsList,
		},
	}, nil
}

// collectOffRampSourceChainConfigsMetadata collects all source chain configs from an OffRamp and returns metadata.
// If isPending is true and pendingConfigs is provided, it uses those instead of querying chain state.
func collectOffRampSourceChainConfigsMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	offRampAddr string,
	chainSelector uint64,
	isPending bool,
	pendingConfigs map[uint64]offramp.SourceChainConfigArgs, // Pending configs by sourceChainSelector (only used if isPending is true)
) (datastore.ContractMetadata, error) {
	var sourceChainConfigsList []map[string]interface{}

	// Always query chain state first to get existing configs
	getAllConfigsReport, err := cldf_ops.ExecuteOperation(b, offramp.GetAllSourceChainConfigs, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(offRampAddr),
		Args:          nil,
	})
	if err != nil {
		return datastore.ContractMetadata{}, fmt.Errorf("failed to read all source chain configs from OffRamp(%s) on chain %s: %w", offRampAddr, chain, err)
	}

	// Build a map of existing configs by sourceChainSelector
	existingConfigsBySelector := make(map[uint64]map[string]interface{})
	for i, sourceChainSelector := range getAllConfigsReport.Output.SourceChainSelectors {
		config := getAllConfigsReport.Output.SourceChainConfigs[i]

		// Convert CCV addresses to hex strings
		defaultCCVsHex := make([]string, len(config.DefaultCCVs))
		for j, ccv := range config.DefaultCCVs {
			defaultCCVsHex[j] = ccv.Hex()
		}

		laneMandatedCCVsHex := make([]string, len(config.LaneMandatedCCVs))
		for j, ccv := range config.LaneMandatedCCVs {
			laneMandatedCCVsHex[j] = ccv.Hex()
		}

		// Convert OnRamps bytes to hex strings
		onRampsHex := make([]string, len(config.OnRamps))
		for j, onRampBytes := range config.OnRamps {
			onRampsHex[j] = fmt.Sprintf("0x%x", onRampBytes)
		}

		existingConfigsBySelector[sourceChainSelector] = map[string]interface{}{
			"sourceChainSelector": sourceChainSelector,
			"router":              config.Router.Hex(),
			"isEnabled":           config.IsEnabled,
			"onRamps":             onRampsHex,
			"defaultCCVs":         defaultCCVsHex,
			"laneMandatedCCVs":    laneMandatedCCVsHex,
		}
	}

	if isPending {
		if pendingConfigs != nil && len(pendingConfigs) > 0 {
			// Merge pending configs with existing configs (pending overrides existing)
			for sourceChainSelector, pendingConfig := range pendingConfigs {
				// Convert pending config to metadata format
				defaultCCVsHex := make([]string, len(pendingConfig.DefaultCCVs))
				for j, ccv := range pendingConfig.DefaultCCVs {
					defaultCCVsHex[j] = ccv.Hex()
				}

				laneMandatedCCVsHex := make([]string, len(pendingConfig.LaneMandatedCCVs))
				for j, ccv := range pendingConfig.LaneMandatedCCVs {
					laneMandatedCCVsHex[j] = ccv.Hex()
				}

				// Convert OnRamps bytes to hex strings
				onRampsHex := make([]string, len(pendingConfig.OnRamps))
				for j, onRampBytes := range pendingConfig.OnRamps {
					onRampsHex[j] = fmt.Sprintf("0x%x", onRampBytes)
				}

				// Override or add pending config
				existingConfigsBySelector[sourceChainSelector] = map[string]interface{}{
					"sourceChainSelector": sourceChainSelector,
					"router":              pendingConfig.Router.Hex(),
					"isEnabled":           pendingConfig.IsEnabled,
					"onRamps":             onRampsHex,
					"defaultCCVs":         defaultCCVsHex,
					"laneMandatedCCVs":    laneMandatedCCVsHex,
				}
			}
		}
		// Convert map to list
		sourceChainConfigsList = make([]map[string]interface{}, 0, len(existingConfigsBySelector))
		for _, config := range existingConfigsBySelector {
			sourceChainConfigsList = append(sourceChainConfigsList, config)
		}
	} else {
		// Not pending - use existing configs as-is
		sourceChainConfigsList = make([]map[string]interface{}, 0, len(existingConfigsBySelector))
		for _, config := range existingConfigsBySelector {
			sourceChainConfigsList = append(sourceChainConfigsList, config)
		}
	}

	return datastore.ContractMetadata{
		Address:       offRampAddr,
		ChainSelector: chainSelector,
		Metadata: map[string]interface{}{
			"sourceChainConfigs": sourceChainConfigsList,
		},
	}, nil
}

// collectTokenAdminRegistryMetadata collects metadata for all tokens from the TokenAdminRegistry
// tokenAdminRegistryAddr should be the address of the TokenAdminRegistry contract
func collectTokenAdminRegistryMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	tokenAdminRegistryAddr string,
) (datastore.ContractMetadata, error) {
	if tokenAdminRegistryAddr == "" {
		return datastore.ContractMetadata{}, fmt.Errorf("TokenAdminRegistry address is required")
	}

	// Get all configured tokens (using max uint64 for maxCount to get all tokens)
	maxUint64 := uint64(18446744073709551615) // 2^64 - 1
	getAllTokensReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetAllConfiguredTokens, chain, contract.FunctionInput[token_admin_registry.GetAllConfiguredTokensArgs]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(tokenAdminRegistryAddr),
		Args: token_admin_registry.GetAllConfiguredTokensArgs{
			StartIndex: 0,
			MaxCount:   maxUint64,
		},
	})
	if err != nil {
		return datastore.ContractMetadata{}, fmt.Errorf("failed to get all configured tokens from TokenAdminRegistry(%s) on chain %s: %w", tokenAdminRegistryAddr, chain, err)
	}

	tokens := getAllTokensReport.Output
	tokensList := make([]map[string]interface{}, 0, len(tokens))

	// For each token, collect its metadata
	for _, tokenAddr := range tokens {
		tokenAddrStr := tokenAddr.Hex()

		// Get token config from TokenAdminRegistry
		getTokenConfigReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetTokenConfig, chain, contract.FunctionInput[common.Address]{
			ChainSelector: chainSelector,
			Address:       common.HexToAddress(tokenAdminRegistryAddr),
			Args:          tokenAddr,
		})
		if err != nil {
			return datastore.ContractMetadata{}, fmt.Errorf("failed to get token config for token %s: %w", tokenAddrStr, err)
		}
		tokenConfig := getTokenConfigReport.Output

		// Get token name, symbol, decimals (like we do for feeTokens)
		nameReport, err := cldf_ops.ExecuteOperation(b, erc20.Name, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return datastore.ContractMetadata{}, fmt.Errorf("failed to get token name for %s: %w", tokenAddrStr, err)
		}

		symbolReport, err := cldf_ops.ExecuteOperation(b, erc20.Symbol, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return datastore.ContractMetadata{}, fmt.Errorf("failed to get token symbol for %s: %w", tokenAddrStr, err)
		}

		decimalsReport, err := cldf_ops.ExecuteOperation(b, erc20.Decimals, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return datastore.ContractMetadata{}, fmt.Errorf("failed to get token decimals for %s: %w", tokenAddrStr, err)
		}

		// Build token metadata object
		tokenMetadata := map[string]interface{}{
			"name":                 nameReport.Output,
			"symbol":               symbolReport.Output,
			"decimals":             decimalsReport.Output,
			"admin":                tokenConfig.Administrator.Hex(),
			"pendingAdministrator": tokenConfig.PendingAdministrator.Hex(),
		}

		// If token has a pool, get pool metadata
		if tokenConfig.TokenPool != (common.Address{}) {
			tokenPoolAddr := tokenConfig.TokenPool.Hex()

			// Get RMN proxy from TokenPool
			getRmnProxyReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRMNProxy, chain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       tokenConfig.TokenPool,
				Args:          nil,
			})
			if err != nil {
				return datastore.ContractMetadata{}, fmt.Errorf("failed to get RMN proxy from TokenPool %s: %w", tokenPoolAddr, err)
			}
			rmnProxyAddr := getRmnProxyReport.Output

			// Get ARM and Owner from RMNProxy
			getARMReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.GetARM, chain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       rmnProxyAddr,
				Args:          nil,
			})
			if err != nil {
				return datastore.ContractMetadata{}, fmt.Errorf("failed to get ARM from RMNProxy %s: %w", rmnProxyAddr.Hex(), err)
			}

			getOwnerReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.Owner, chain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       rmnProxyAddr,
				Args:          nil,
			})
			if err != nil {
				return datastore.ContractMetadata{}, fmt.Errorf("failed to get owner from RMNProxy %s: %w", rmnProxyAddr.Hex(), err)
			}

			// Build nested structure: tokenPool -> rmnProxy -> {owner, arm}
			tokenMetadata["tokenPool"] = map[string]interface{}{
				"address": tokenPoolAddr,
				"rmnProxy": map[string]interface{}{
					"address": rmnProxyAddr.Hex(),
					"owner":   getOwnerReport.Output.Hex(),
					"arm":     getARMReport.Output.Hex(),
				},
			}
		} else {
			tokenMetadata["tokenPool"] = nil
		}

		// Add token address to the metadata object
		tokenMetadata["address"] = tokenAddrStr
		tokensList = append(tokensList, tokenMetadata)
	}

	return datastore.ContractMetadata{
		Address:       tokenAdminRegistryAddr,
		ChainSelector: chainSelector,
		Metadata: map[string]interface{}{
			"tokens": tokensList,
		},
	}, nil
}

// getCSAKeyFromOnchainPublicKey queries the job distributor to find the CSA key
// associated with the given onchain public key for a specific chain selector.
func getCSAKeyFromOnchainPublicKey(
	ctx context.Context,
	offchainClient cldf_offchain.Client,
	onchainPublicKey string,
	chainSelector uint64,
) (string, error) {
	if offchainClient == nil {
		return "", errors.New("offchain client is nil, ensure it is configured")
	}

	if chainSelector == 0 {
		return "", errors.New("chain selector is required and cannot be 0")
	}

	// Normalize the onchain public key (remove 0x prefix if present, convert to lowercase)
	onchainKey := strings.TrimPrefix(strings.ToLower(onchainPublicKey), "0x")

	// Decode to bytes for comparison
	onchainKeyBytes, err := hex.DecodeString(onchainKey)
	if err != nil {
		return "", fmt.Errorf("invalid onchain public key format: %w", err)
	}

	// List all nodes from the job distributor
	resp, err := offchainClient.ListNodes(ctx, &nodev1.ListNodesRequest{
		Filter: &nodev1.ListNodesRequest_Filter{},
	})
	if err != nil {
		return "", fmt.Errorf("failed to list nodes: %w", err)
	}

	if len(resp.GetNodes()) == 0 {
		return "", errors.New("no nodes found in job distributor")
	}

	// Collect all node IDs to fetch detailed information
	var nodeIds []string
	for _, node := range resp.GetNodes() {
		nodeIds = append(nodeIds, node.GetId())
	}

	// Get detailed node information (includes OCR configs with onchain public keys)
	allNodes, err := deployment.NodeInfo(nodeIds, offchainClient)
	if err != nil {
		return "", fmt.Errorf("failed to get nodes info (requested %d nodes): %w", len(nodeIds), err)
	}

	if len(allNodes) == 0 {
		return "", fmt.Errorf("deployment.NodeInfo returned no nodes (requested %d node IDs: %v)", len(nodeIds), nodeIds)
	}

	// Search through all nodes for the matching onchain public key
	for _, node := range allNodes {
		ocrConfig, exists := node.OCRConfigForChainSelector(chainSelector)
		if !exists {
			continue
		}

		// Compare onchain public keys (they're stored as bytes)
		if bytesEqual(ocrConfig.OnchainPublicKey, onchainKeyBytes) {
			return node.CSAKey, nil
		}
	}

	return "", fmt.Errorf("node with onchain public key %s not found for chain selector %d (searched %d nodes, onchain key bytes: %x)", onchainPublicKey, chainSelector, len(allNodes), onchainKeyBytes)
}

// bytesEqual compares two byte slices for equality
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
