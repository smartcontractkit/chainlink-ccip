package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	seq1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	onrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
)

const (
	DefaultTxGasLimit        uint32 = 200_000
	DefaultMaxPerMsgGasLimit uint32 = 8_000_000
	DefaultMaxDataBytes      uint32 = 32_000
)

type LaneMigrator struct{}

func (r *LaneMigrator) VerifyPreconditions(e deployment.Environment, cfg deploy.LaneMigratorConfig, _ changesets.MCMSReader) error {
	multipleRefContract := []datastore.AddressRef{
		{
			Type:    datastore.ContractType(committee_verifier.ContractType),
			Version: committee_verifier.Version,
		},
		{
			Type:    datastore.ContractType(executor.ContractType),
			Version: executor.Version,
		},
		{
			Type:    datastore.ContractType(seq1_7.ExecutorProxyType),
			Version: executor.Version,
		},
	}
	singleRefContracts := []datastore.AddressRef{
		// 2.0 contracts
		{
			Type:    datastore.ContractType(onrampops.ContractType),
			Version: onrampops.Version,
		},
		{
			Type:    datastore.ContractType(offrampops.ContractType),
			Version: offrampops.Version,
		},
		{
			Type:    datastore.ContractType(fqops.ContractType),
			Version: fqops.Version,
		},
		// contracts from previous versions relevant to 2.0
		{
			Type:    datastore.ContractType(routerops.ContractType),
			Version: routerops.Version,
		},
		{
			Type:    datastore.ContractType(rmn_remote.ContractType),
			Version: rmn_remote.Version,
		},
		{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: rmn_proxy.Version,
		},
		{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		},
	}
	allContractRefs := append(multipleRefContract, singleRefContracts...)
	for chainSelector, perChainCfg := range cfg.Input {
		err := verifyAllContractsPresent(e, chainSelector, singleRefContracts, multipleRefContract)
		if err != nil {
			return fmt.Errorf("contract verification failed for chain %d: %w", chainSelector, err)
		}
		err = verifyOwnershipOfContracts(e, chainSelector, allContractRefs)
		if err != nil {
			return fmt.Errorf("ownership verification failed for chain %d: %w", chainSelector, err)
		}
		evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
		if !ok {
			return fmt.Errorf("evm chain not found for selector %d", chainSelector)
		}
		err = verifyExistingLaneVersion(e, evmChain, chainSelector, perChainCfg.RemoteChains)
		if err != nil {
			return fmt.Errorf("existing lane version verification failed for chain %d: %w", chainSelector, err)
		}
	}
	return nil
}

// verifyAllContractsPresent checks that all contracts in contractRefs have a corresponding address
// in the datastore for the given chain (i.e. all required contracts are deployed and registered).
// multipleRefContracts are contracts for which there can be multiple addresses of same type and version (e.g. CommitteeVerifier and Executor)
// singleRefContracts are contracts for which there should be only one address for the given type and version in the datastore (e.g. OnRamp, OffRamp, FeeQuoter, Router, RMNProxy, RMNRemote, TokenAdminRegistry)
func verifyAllContractsPresent(e deployment.Environment, chainSelector uint64, singleRefContracts []datastore.AddressRef, multipleContractsOfSameType []datastore.AddressRef) error {
	var missing []string
	for _, ref := range multipleContractsOfSameType {
		refs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSelector),
			datastore.AddressRefByType(ref.Type),
			datastore.AddressRefByVersion(ref.Version),
		)
		if len(refs) == 0 {
			missing = append(missing, fmt.Sprintf("%s@%s", ref.Type, ref.Version))
		}
	}
	for _, ref := range singleRefContracts {
		_, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			missing = append(missing, fmt.Sprintf("%s@%s", ref.Type, ref.Version))
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("not all required contracts are present in datastore for chain %d: missing %v", chainSelector, missing)
	}
	return nil
}

// verifyExistingLaneVersion checks that for the given chain selector, the existing Router contract has the
// onRamp configured for the remote chains and that the onRamp and fee quoter versions are correct.
// This ensures that the existing lanes are using the expected versions of contracts before we attempt to update them to use the prod Router.
func verifyExistingLaneVersion(e deployment.Environment, evmChain evm.Chain, chainSelector uint64, remoteChains []uint64) error {
	// fetch router address from datastore
	routerRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(routerops.ContractType),
		Version:       routerops.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("error fetching router address ref for chain %d: %w", chainSelector, err)
	}
	// get onRamp for remote chains
	for _, remoteChainSelector := range remoteChains {
		onRampOnRouterOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, routerops.GetOnRamp,
			evmChain, contract.FunctionInput[uint64]{
				ChainSelector: chainSelector,
				Address:       routerRef,
				Args:          remoteChainSelector,
			})
		if err != nil {
			return fmt.Errorf("error fetching onRamp from router for chain %d and remote chain %d: %w", chainSelector, remoteChainSelector, err)
		}
		if onRampOnRouterOut.Output == (common.Address{}) {
			return fmt.Errorf("precondition failed for chain %d and remote chain %d: "+
				"expected to find an onRamp configured on the Router for the remote chain, but got zero address. "+
				"Please configure lanes with the prod Router first before migrating", chainSelector, remoteChainSelector)
		}
		_, onRampVersion, err := utils.TypeAndVersion(onRampOnRouterOut.Output, evmChain.Client)
		if err != nil {
			return fmt.Errorf("error fetching onRamp version for chain %d and remote chain %d: %w", chainSelector, remoteChainSelector, err)
		}

		if !onRampVersion.Equal(onrampops_v160.Version) {
			return fmt.Errorf(
				"precondition failed for chain %d and remote chain %d: expected onRamp version on Router to be %s, but got version %s. ",
				chainSelector, remoteChainSelector, onrampops_v160.Version.String(), onRampVersion.String(),
			)
		}

		// get the fee quoter from onRamp
		feeQuoterOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onrampops_v160.GetDynamicConfig,
			evmChain, contract.FunctionInput[struct{}]{
				ChainSelector: chainSelector,
				Address:       onRampOnRouterOut.Output,
			})
		if err != nil {
			return fmt.Errorf("error fetching fee quoter from onRamp for chain %d and remote chain %d: %w", chainSelector, remoteChainSelector, err)
		}
		if feeQuoterOut.Output.FeeQuoter == (common.Address{}) {
			return fmt.Errorf("precondition failed for chain %d and remote chain %d: "+
				"expected to find a fee quoter configured on the onRamp for the remote chain, but got zero address. "+
				"Please configure lanes with the prod Router first before migrating", chainSelector, remoteChainSelector)
		}
		// check the fee quoter version
		_, feeQuoterVersion, err := utils.TypeAndVersion(feeQuoterOut.Output.FeeQuoter, evmChain.Client)
		if err != nil {
			return fmt.Errorf("error fetching fee quoter version from feequoter %s for chain %d and remote chain %d: %w",
				feeQuoterOut.Output.FeeQuoter.String(), chainSelector, remoteChainSelector, err)
		}
		if !feeQuoterVersion.Equal(fqops.Version) {
			return fmt.Errorf("precondition failed for chain %d and remote chain %d: "+
				"expected fee quoter version on onRamp to be %s, but got version %s. ",
				chainSelector, remoteChainSelector, fqops.Version.String(), feeQuoterVersion.String())
		}
	}
	return nil
}

func verifyOwnershipOfContracts(e deployment.Environment, chainSelector uint64, contractRefs []datastore.AddressRef) error {
	cllCCIPTimelock, rmnTimelock, _, err := seq1_7.ResolveOwnershipDeps(e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
	), chainSelector)

	if err != nil {
		return fmt.Errorf("error resolving ownership dependencies for chain %d: %w", chainSelector, err)
	}
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return fmt.Errorf("chain with selector %d not found in environment", chainSelector)
	}
	for _, ref := range contractRefs {
		refs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSelector),
			datastore.AddressRefByType(ref.Type),
			datastore.AddressRefByVersion(ref.Version),
		)
		if len(refs) == 0 {
			return fmt.Errorf("no address ref found for contract type %s version %s on chain %d", ref.Type, ref.Version, chainSelector)
		}
		for _, addrRef := range refs {
			addr, err := evm_datastore_utils.ToEVMAddress(addrRef)
			if err != nil {
				return fmt.Errorf("error formatting address ref %s for contract type %s version %s on chain %d: %w",
					addrRef.Address, ref.Type, ref.Version, chainSelector, err)
			}
			currentOwner, _, err := mcms_seq.LoadOwnableContract(addr, evmChain.Client)
			if err != nil {
				return fmt.Errorf("failed to load ownable contract %s (%s): %w", addr, ref.Type, err)
			}
			expectedTimelockAddr := cllCCIPTimelock
			if ref.Type == datastore.ContractType(rmn_remote.ContractType) {
				expectedTimelockAddr = rmnTimelock
			}
			if currentOwner != expectedTimelockAddr {
				return fmt.Errorf("precondition failed for chain %d: expected owner of "+
					"contract type %s version %s at address %s is timelock %s, but got %s",
					chainSelector, ref.Type, ref.Version, addr, expectedTimelockAddr, currentOwner)
			}
		}
	}
	return nil
}

// UpdateVersionWithRouter is a sequence that updates Ramps to use the new Router and also updates the fee quoter dest chain config with default tx gas limit as 8M
//
// It fetches the existing onRamp and offRamp addresses from the provided ExistingAddresses, then calls the necessary functions to update the onRamp and offRamp to use the new Router.
//
// This sequence assumes that the destChainConfig on OnRamp and SourceChainConfig on OffRamp do not need to be updated, and only updates the Router address used by the Ramps.
// If you need to update the destChainConfig or sourceChainConfig, please use the ConnectChains sequence instead.
func (r *LaneMigrator) UpdateVersionWithRouter() *cldf_ops.Sequence[deploy.RampUpdaterConfig, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-updater:sequence-update-ramps-with-router",
		semver.MustParse("2.0.0"),
		"Updates Ramps contracts to use the updated Router contract",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.RampUpdaterConfig) (output sequences.OnChainOutput, err error) {
			var writes []contract.WriteOutput
			c, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("c with selector %d not found in environment", input.ChainSelector)
			}
			ds := datastore.NewMemoryDataStore()
			for _, addrRef := range input.ExistingAddresses {
				if err := ds.Addresses().Add(addrRef); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error adding address ref to temp datastore: %w", err)
				}
			}
			tempDS := ds.Seal()
			// fetch onRamp and offRamp from the existing addresses
			onRampAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(onrampops.ContractType),
					Version:       onrampops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding onRamp address ref: %w", err)
			}
			offRampAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(offrampops.ContractType),
					Version:       offrampops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding offRamp address ref: %w", err)
			}
			feequoterAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(fqops.ContractType),
					Version:       fqops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding feequoter address ref: %w", err)
			}
			feeQuoterContract, err := fee_quoter.NewFeeQuoter(feequoterAddr, c.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error creating fee quoter contract instance: %w", err)
			}
			routerRef := input.RouterAddr
			routerAddr, err := evm_datastore_utils.ToEVMAddress(routerRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error formatting router address ref: %w", err)
			}
			var onRampArgs []onrampops.DestChainConfigArgs
			var offRampArgs []offrampops.SourceChainConfigArgs
			var fqArgs []fqops.DestChainConfigArgs
			for _, remoteChainSelector := range input.RemoteChainSelectors {
				// get existing destChainConfig for the onRamp
				existingDestChainCfgOut, err := cldf_ops.ExecuteOperation(b, onrampops.GetDestChainConfig, c, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       onRampAddr,
					Args:          remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing destChainConfig for onRamp: %w", err)
				}
				existingDestChainCfg := existingDestChainCfgOut.Output
				if existingDestChainCfg.AddressBytesLength == 0 {
					return sequences.OnChainOutput{}, fmt.Errorf("no destchain config is set for remote chain %d on chain %d on onRamp."+
						" configure lanes with test router first before migrating", remoteChainSelector, input.ChainSelector)
				}
				// update router on onRamp for the remote c
				existingDestChainCfg.Router = routerAddr

				// get the sourceChainConfig for the offRamp
				srcChainCfgOut, err := cldf_ops.ExecuteOperation(b, offrampops.GetSourceChainConfig, c, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       offRampAddr,
					Args:          remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing sourceChainConfig for offRamp: %w", err)
				}
				existingSrcChainCfg := srcChainCfgOut.Output
				// update router on offRamp for the remote c
				existingSrcChainCfg.Router = routerAddr
				onRampArgs = append(onRampArgs, onrampops.DestChainConfigArgs{
					DestChainSelector:         remoteChainSelector,
					Router:                    routerAddr,
					AddressBytesLength:        existingDestChainCfg.AddressBytesLength,
					TokenReceiverAllowed:      existingDestChainCfg.TokenReceiverAllowed,
					MessageNetworkFeeUSDCents: existingDestChainCfg.MessageNetworkFeeUSDCents,
					TokenNetworkFeeUSDCents:   existingDestChainCfg.TokenNetworkFeeUSDCents,
					BaseExecutionGasCost:      existingDestChainCfg.BaseExecutionGasCost,
					DefaultCCVs:               existingDestChainCfg.DefaultCCVs,
					LaneMandatedCCVs:          existingDestChainCfg.LaneMandatedCCVs,
					DefaultExecutor:           existingDestChainCfg.DefaultExecutor,
					OffRamp:                   existingDestChainCfg.OffRamp,
				})

				offRampArgs = append(offRampArgs, offrampops.SourceChainConfigArgs{
					SourceChainSelector: remoteChainSelector,
					Router:              routerAddr,
					IsEnabled:           existingSrcChainCfg.IsEnabled,
					OnRamps:             existingSrcChainCfg.OnRamps,
					DefaultCCVs:         existingSrcChainCfg.DefaultCCVs,
					LaneMandatedCCVs:    existingSrcChainCfg.LaneMandatedCCVs,
				})
				// this also needs feequoter update for dest chain config,
				// fetch existing destChainConfig for feequoter
				dstChainCfg, err := feeQuoterContract.GetDestChainConfig(&bind.CallOpts{
					Context: b.GetContext(),
				}, remoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing destChainConfig for fee quoter: %w", err)
				}
				fqArgs = append(fqArgs, fqops.DestChainConfigArgs{
					DestChainSelector: remoteChainSelector,
					DestChainConfig: fqops.DestChainConfig{
						IsEnabled:                   dstChainCfg.IsEnabled,
						MaxDataBytes:                DefaultMaxDataBytes,
						MaxPerMsgGasLimit:           DefaultMaxPerMsgGasLimit,
						DestGasOverhead:             dstChainCfg.DestGasOverhead,
						DestGasPerPayloadByteBase:   dstChainCfg.DestGasPerPayloadByteBase,
						ChainFamilySelector:         dstChainCfg.ChainFamilySelector,
						DefaultTokenFeeUSDCents:     dstChainCfg.DefaultTokenFeeUSDCents,
						DefaultTokenDestGasOverhead: dstChainCfg.DefaultTokenDestGasOverhead,
						DefaultTxGasLimit:           DefaultTxGasLimit,
						NetworkFeeUSDCents:          dstChainCfg.NetworkFeeUSDCents,
						LinkFeeMultiplierPercent:    dstChainCfg.LinkFeeMultiplierPercent,
					},
				})
			}
			//  set the destChainConfig with the updated router
			writeOutputOnRamp, err := cldf_ops.ExecuteOperation(b, onrampops.ApplyDestChainConfigUpdates, c, contract.FunctionInput[[]onrampops.DestChainConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       onRampAddr,
				Args:          onRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying destChainConfig update to onRamp: %w", err)
			}
			writes = append(writes, writeOutputOnRamp.Output)
			// now set the sourceChainConfig with the updated router
			writeOutputOffRamp, err := cldf_ops.ExecuteOperation(
				b, offrampops.ApplySourceChainConfigUpdates, c,
				contract.FunctionInput[[]offrampops.SourceChainConfigArgs]{
					ChainSelector: input.ChainSelector,
					Address:       offRampAddr,
					Args:          offRampArgs,
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying sourceChainConfig update to offRamp: %w", err)
			}
			writes = append(writes, writeOutputOffRamp.Output)
			// update fq 2.0 to have defaultTxLimit set to 8M
			fqDestChainUpdateRep, err := cldf_ops.ExecuteOperation(b, fqops.ApplyDestChainConfigUpdates, c, contract.FunctionInput[[]fqops.DestChainConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       feequoterAddr,
				Args:          fqArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying destChainConfig update to fee quoter: %w", err)
			}
			writes = append(writes, fqDestChainUpdateRep.Output)
			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batchOp},
			}, nil
		})
}
