package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

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

			committeeVerifierMetadata, err := collectCommitteeVerifierSignatureConfigs(b, chain, committeeVerifierAddr, chain.Selector, []datastore.ContractMetadata{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to collect CommitteeVerifier signature configs: %w", err)
			}
			contractMetadata = append(contractMetadata, committeeVerifierMetadata...)
		}

		// Collect OnRamp destination chain config metadata
		onRampDestChainMetadata, err := collectOnRampDestChainConfigs(b, chain, input.OnRamp, chain.Selector, []datastore.ContractMetadata{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to collect OnRamp destination chain configs: %w", err)
		}
		contractMetadata = append(contractMetadata, onRampDestChainMetadata...)

		// Collect Router metadata (onRamps and offRamps)
		routerMetadata, err := collectRouterMetadata(b, chain, input.Router, chain.Selector, []datastore.ContractMetadata{})
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

	// Convert to metadata
	for selector, cfg := range currentConfigs {
		signersHex := make([]string, len(cfg.Signers))
		for i, signer := range cfg.Signers {
			signersHex[i] = signer.Hex()
		}
		contractMetadata = append(contractMetadata, datastore.ContractMetadata{
			Address:       committeeVerifierAddr,
			ChainSelector: chainSelector,
			Metadata: map[string]interface{}{
				"sourceChainSelector": selector,
				"threshold":           cfg.Threshold,
				"signers":             signersHex,
			},
		})
	}

	return contractMetadata, nil
}

// collectOnRampDestChainConfigs collects metadata for all destination chain configurations from the OnRamp.
// For each destination chain, it saves: router, defaultExecutor, laneMandatedCCVs, defaultCCVs, and offRamp bytes.
// Note: To determine if a Router/OffRamp/OnRamp is test, check the AddressRef's Type field in the datastore
// (e.g., Type == "TestRouter" vs Type == "Router").
func collectOnRampDestChainConfigs(
	b cldf_ops.Bundle,
	chain evm.Chain,
	onRampAddr string,
	chainSelector uint64,
	contractMetadata []datastore.ContractMetadata,
) ([]datastore.ContractMetadata, error) {
	// TODO: Uncomment once GetAllDestChainConfigs is available in gobindings
	// Read all destination chain configs from the OnRamp
	// getAllConfigsReport, err := cldf_ops.ExecuteOperation(b, onramp.GetAllDestChainConfigs, chain, contract.FunctionInput[any]{
	// 	ChainSelector: chainSelector,
	// 	Address:       common.HexToAddress(onRampAddr),
	// 	Args:          nil,
	// })
	// if err != nil {
	// 	return contractMetadata, fmt.Errorf("failed to read all dest chain configs from OnRamp(%s) on chain %s: %w", onRampAddr, chain, err)
	// }

	// TODO: Uncomment once GetAllDestChainConfigs is available in gobindings
	// for i, destChainSelector := range getAllConfigsReport.Output.DestChainSelectors {
	// 	destChainConfig := getAllConfigsReport.Output.DestChainConfigs[i]
	//
	// 	offRampBytes := destChainConfig.OffRamp
	//
	// 	// Convert CCV addresses to hex strings
	// 	defaultCCVsHex := make([]string, len(destChainConfig.DefaultCCVs))
	// 	for j, ccv := range destChainConfig.DefaultCCVs {
	// 		defaultCCVsHex[j] = ccv.Hex()
	// 	}
	//
	// 	laneMandatedCCVsHex := make([]string, len(destChainConfig.LaneMandatedCCVs))
	// 	for j, ccv := range destChainConfig.LaneMandatedCCVs {
	// 		laneMandatedCCVsHex[j] = ccv.Hex()
	// 	}
	//
	// 	contractMetadata = append(contractMetadata, datastore.ContractMetadata{
	// 		Address:       onRampAddr,
	// 		ChainSelector: chainSelector,
	// 		Metadata: map[string]interface{}{
	// 			"destChainSelector": destChainSelector,
	// 			"router":            destChainConfig.Router.Hex(),
	// 			"defaultExecutor":   destChainConfig.DefaultExecutor.Hex(),
	// 			"laneMandatedCCVs":  laneMandatedCCVsHex,
	// 			"defaultCCVs":       defaultCCVsHex,
	// 			"offRamp":           fmt.Sprintf("0x%x", offRampBytes),
	// 		},
	// 	})
	// }

	return contractMetadata, nil
}

// collectRouterMetadata collects metadata for all onRamps and offRamps from the Router
func collectRouterMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	routerAddr string,
	chainSelector uint64,
	contractMetadata []datastore.ContractMetadata,
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

	// Group offRamps by sourceChainSelector
	offRampsByChain := make(map[uint64][]string)
	uniqueChainSelectors := make(map[uint64]struct{})
	for _, offRamp := range getOffRampsReport.Output {
		sourceChainSelector := offRamp.SourceChainSelector
		offRampsByChain[sourceChainSelector] = append(offRampsByChain[sourceChainSelector], offRamp.OffRamp.Hex())
		uniqueChainSelectors[sourceChainSelector] = struct{}{}
	}

	// Get onRamp for each unique sourceChainSelector
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
	for _, onRampAddr := range onRampsByChain {
		uniqueOnRamps[onRampAddr] = struct{}{}
	}

	for onRampAddr := range uniqueOnRamps {
		onRampMetadata, err := collectOnRampDestChainConfigsMetadata(b, chain, onRampAddr, chainSelector)
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to collect OnRamp metadata for %s: %w", onRampAddr, err)
		}
		contractMetadata = append(contractMetadata, onRampMetadata)
	}

	// Collect OffRamp metadata for each unique OffRamp
	uniqueOffRamps := make(map[string]struct{})
	for _, offRampAddrs := range offRampsByChain {
		for _, offRampAddr := range offRampAddrs {
			uniqueOffRamps[offRampAddr] = struct{}{}
		}
	}

	for offRampAddr := range uniqueOffRamps {
		offRampMetadata, err := collectOffRampSourceChainConfigsMetadata(b, chain, offRampAddr, chainSelector)
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to collect OffRamp metadata for %s: %w", offRampAddr, err)
		}
		contractMetadata = append(contractMetadata, offRampMetadata)
	}

	return contractMetadata, nil
}

// collectOnRampDestChainConfigsMetadata collects all destination chain configs from an OnRamp and returns metadata
func collectOnRampDestChainConfigsMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	onRampAddr string,
	chainSelector uint64,
) (datastore.ContractMetadata, error) {
	// Read all destination chain configs from the OnRamp
	getAllConfigsReport, err := cldf_ops.ExecuteOperation(b, onramp.GetAllDestChainConfigs, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(onRampAddr),
		Args:          nil,
	})
	if err != nil {
		return datastore.ContractMetadata{}, fmt.Errorf("failed to read all dest chain configs from OnRamp(%s) on chain %s: %w", onRampAddr, chain, err)
	}

	// Convert dest chain configs to JSON-serializable format
	destChainConfigsList := make([]map[string]interface{}, 0, len(getAllConfigsReport.Output.DestChainSelectors))
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

		destChainConfigsList = append(destChainConfigsList, map[string]interface{}{
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
		})
	}

	return datastore.ContractMetadata{
		Address:       onRampAddr,
		ChainSelector: chainSelector,
		Metadata: map[string]interface{}{
			"destChainConfigs": destChainConfigsList,
		},
	}, nil
}

// collectOffRampSourceChainConfigsMetadata collects all source chain configs from an OffRamp and returns metadata
func collectOffRampSourceChainConfigsMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	offRampAddr string,
	chainSelector uint64,
) (datastore.ContractMetadata, error) {
	// Read all source chain configs from the OffRamp
	getAllConfigsReport, err := cldf_ops.ExecuteOperation(b, offramp.GetAllSourceChainConfigs, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(offRampAddr),
		Args:          nil,
	})
	if err != nil {
		return datastore.ContractMetadata{}, fmt.Errorf("failed to read all source chain configs from OffRamp(%s) on chain %s: %w", offRampAddr, chain, err)
	}

	// Convert source chain configs to JSON-serializable format
	sourceChainConfigsList := make([]map[string]interface{}, 0, len(getAllConfigsReport.Output.SourceChainSelectors))
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

		sourceChainConfigsList = append(sourceChainConfigsList, map[string]interface{}{
			"sourceChainSelector": sourceChainSelector,
			"router":              config.Router.Hex(),
			"isEnabled":           config.IsEnabled,
			"onRamps":             onRampsHex,
			"defaultCCVs":         defaultCCVsHex,
			"laneMandatedCCVs":    laneMandatedCCVsHex,
		})
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
