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

		// Collect fee token metadata
		contractMetadata, err := collectFeeTokenMetadata(b, chain, input.FeeQuoter, chain.Selector, []datastore.ContractMetadata{})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

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

		return sequences.OnChainOutput{
			Metadata: sequences.Metadata{
				Contracts: contractMetadata,
			},
			BatchOps: batchOps,
		}, nil
	})

// collectFeeTokenMetadata collects metadata for all fee tokens from the FeeQuoter
func collectFeeTokenMetadata(
	b cldf_ops.Bundle,
	chain evm.Chain,
	feeQuoterAddr string,
	chainSelector uint64,
	contractMetadata []datastore.ContractMetadata,
) ([]datastore.ContractMetadata, error) {
	// Read fee tokens from the FeeQuoter
	getFeeTokensReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetFeeTokens, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(feeQuoterAddr),
		Args:          nil,
	})
	if err != nil {
		return contractMetadata, fmt.Errorf("failed to read fee tokens from FeeQuoter(%s) on chain %s: %w", feeQuoterAddr, chain, err)
	}

	// Read metadata for each fee token
	for _, tokenAddr := range getFeeTokensReport.Output {
		// Read name
		nameReport, err := cldf_ops.ExecuteOperation(b, erc20.Name, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to read name of fee token(%s) on chain %s: %w", tokenAddr, chain, err)
		}

		// Read symbol
		symbolReport, err := cldf_ops.ExecuteOperation(b, erc20.Symbol, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to read symbol of fee token(%s) on chain %s: %w", tokenAddr, chain, err)
		}

		// Read decimals
		decimalsReport, err := cldf_ops.ExecuteOperation(b, erc20.Decimals, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       tokenAddr,
			Args:          nil,
		})
		if err != nil {
			return contractMetadata, fmt.Errorf("failed to read decimals of fee token(%s) on chain %s: %w", tokenAddr, chain, err)
		}

		contractMetadata = append(contractMetadata, datastore.ContractMetadata{
			Address:       tokenAddr.Hex(),
			ChainSelector: chainSelector,
			Metadata: map[string]interface{}{
				"name":     nameReport.Output,
				"symbol":   symbolReport.Output,
				"decimals": decimalsReport.Output,
			},
		})
	}

	return contractMetadata, nil
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
