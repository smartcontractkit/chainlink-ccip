package common

import (
	"errors"
	"fmt"
	"math"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cciphomeops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var CCIPHomeABI *abi.ABI

func init() {
	var err error
	CCIPHomeABI, err = ccip_home.CCIPHomeMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
}

type DONSequenceDeps struct {
	HomeChain cldf_evm.Chain
}

type DONAddition struct {
	ExpectedID       uint32
	PluginConfig     ccip_home.CCIPHomeOCR3Config
	PeerIDs          [][32]byte
	F                uint8
	IsPublic         bool
	AcceptsWorkflows bool
}

type AddDONAndSetCandidateSequenceInput struct {
	CapabilitiesRegistry common.Address
	NoSend               bool
	DONs                 []DONAddition
}

var AddDONAndSetCandidateSequence = operations.NewSequence(
	"AddDONAndSetCandidateSequence",
	semver.MustParse("1.0.0"),
	"Adds commit / exec DONs for chains and sets their candidates on CCIPHome",
	func(b operations.Bundle, deps DONSequenceDeps, input AddDONAndSetCandidateSequenceInput) (sequences.OnChainOutput, error) {
		var writes []contract.WriteOutput
		for _, don := range input.DONs {
			encodedSetCandidateCall, err := CCIPHomeABI.Pack(
				"setCandidate",
				don.ExpectedID,
				don.PluginConfig.PluginType,
				don.PluginConfig,
				[32]byte{},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to pack set candidate call: %w", err)
			}
			out, err := operations.ExecuteOperation(
				b,
				cciphomeops.AddDON,
				deps.HomeChain,
				contract.FunctionInput[cciphomeops.AddDONOpInput]{
					ChainSelector: deps.HomeChain.Selector,
					Address:       input.CapabilitiesRegistry,
					Args: cciphomeops.AddDONOpInput{
						Nodes: don.PeerIDs,
						CapabilityConfigurations: []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration{
							{
								CapabilityId: CCIPCapabilityID,
								Config:       encodedSetCandidateCall,
							},
						},
						IsPublic:         don.IsPublic,
						AcceptsWorkflows: don.AcceptsWorkflows,
						F:                don.F,
					},
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AddDON for chain with selector %d and plugin type %s: %w", don.PluginConfig.ChainSelector, don.PluginConfig.PluginType, err)
			}
			writes = append(writes, out.Output)
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

type DONUpdate struct {
	ID             uint32
	PluginConfig   ccip_home.CCIPHomeOCR3Config
	PeerIDs        [][32]byte
	F              uint8
	IsPublic       bool
	ExistingDigest [32]byte
}

type SetCandidateSequenceInput struct {
	CapabilitiesRegistry common.Address
	NoSend               bool
	DONs                 []DONUpdate
}

var SetCandidateSequence = operations.NewSequence(
	"SetCandidateSequence",
	semver.MustParse("1.0.0"),
	"Updates candidates for existing commit / exec DONs across multiple chains",
	func(b operations.Bundle, deps DONSequenceDeps, input SetCandidateSequenceInput) (sequences.OnChainOutput, error) {
		var writes []contract.WriteOutput
		for _, don := range input.DONs {
			encodedSetCandidateCall, err := CCIPHomeABI.Pack(
				"setCandidate",
				don.ID,
				don.PluginConfig.PluginType,
				don.PluginConfig,
				don.ExistingDigest,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to pack set candidate call: %w", err)
			}
			out, err := operations.ExecuteOperation(
				b,
				cciphomeops.UpdateDON,
				deps.HomeChain,
				contract.FunctionInput[cciphomeops.UpdateDONOpInput]{
					Address:       input.CapabilitiesRegistry,
					ChainSelector: deps.HomeChain.Selector,
					Args: cciphomeops.UpdateDONOpInput{
						ID:    don.ID,
						Nodes: don.PeerIDs,
						CapabilityConfigurations: []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration{
							{
								CapabilityId: CCIPCapabilityID,
								Config:       encodedSetCandidateCall,
							},
						},
						IsPublic: don.IsPublic,
						F:        don.F,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute UpdateDON for chain with selector %d and plugin type %s: %w", don.PluginConfig.ChainSelector, don.PluginConfig.PluginType, err)
			}
			writes = append(writes, out.Output)
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

type DONUpdatePromotion struct {
	ID              uint32
	PluginType      uint8
	ChainSelector   uint64
	PeerIDs         [][32]byte
	F               uint8
	IsPublic        bool
	CandidateDigest [32]byte
	ActiveDigest    [32]byte
}

type PromoteCandidateSequenceInput struct {
	CapabilitiesRegistry common.Address
	NoSend               bool
	DONs                 []DONUpdatePromotion
}

var PromoteCandidateSequence = operations.NewSequence(
	"PromoteCandidateSequence",
	semver.MustParse("1.0.0"),
	"Promote candidates for existing commit / exec DONs across multiple chains",
	func(b operations.Bundle, deps DONSequenceDeps, input PromoteCandidateSequenceInput) (sequences.OnChainOutput, error) {
		var writes []contract.WriteOutput
		for _, don := range input.DONs {
			encodedPromoteCandidateCall, err := CCIPHomeABI.Pack(
				"promoteCandidateAndRevokeActive",
				don.ID,
				don.PluginType,
				don.CandidateDigest,
				don.ActiveDigest,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to pack promote candidate call: %w", err)
			}
			out, err := operations.ExecuteOperation(
				b,
				cciphomeops.UpdateDON,
				deps.HomeChain,
				contract.FunctionInput[cciphomeops.UpdateDONOpInput]{
					Address:       input.CapabilitiesRegistry,
					ChainSelector: deps.HomeChain.Selector,
					Args: cciphomeops.UpdateDONOpInput{
						ID:    don.ID,
						Nodes: don.PeerIDs,
						CapabilityConfigurations: []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration{
							{
								CapabilityId: CCIPCapabilityID,
								Config:       encodedPromoteCandidateCall,
							},
						},
						IsPublic: don.IsPublic,
						F:        don.F,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute UpdateDONOp for chain with selector %d and plugin type %s: %w", don.ChainSelector, don.PluginType, err)
			}
			writes = append(writes, out.Output)
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

type ApplyChainConfigUpdatesSequenceInput struct {
	CCIPHome           common.Address
	NoSend             bool
	RemoteChainAdds    []ccip_home.CCIPHomeChainConfigArgs
	RemoteChainRemoves []uint64
	BatchSize          int
}

var ApplyChainConfigUpdatesSequence = operations.NewSequence(
	"ApplyChainConfigUpdatesSequence",
	semver.MustParse("1.0.0"),
	"Updates chain configurations on CCIPHome, using multiple ApplyChainConfigUpdates according to a batch size",
	func(b operations.Bundle, deps DONSequenceDeps, input ApplyChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
		batches := make([]cciphomeops.ApplyChainConfigUpdatesOpInput, 0)
		currentBatch := cciphomeops.ApplyChainConfigUpdatesOpInput{
			RemoteChainRemoves: make([]uint64, 0),
			RemoteChainAdds:    make([]ccip_home.CCIPHomeChainConfigArgs, 0),
		}

		// Track additions for quick lookups. Although we generally process removals first,
		// if an addition for the same chain exists we must batch it with the removal.
		// This is to ensure that there isn't any downtime for the chain in question.
		adds := make(map[uint64]ccip_home.CCIPHomeChainConfigArgs)
		for _, add := range input.RemoteChainAdds {
			adds[add.ChainSelector] = add
		}

		processedAdds := make(map[uint64]struct{})
		for _, removal := range input.RemoteChainRemoves {
			currentBatch.RemoteChainRemoves = append(currentBatch.RemoteChainRemoves, removal)
			// If there's an addition for the same chain, add it to the same batch
			if add, ok := adds[removal]; ok {
				currentBatch.RemoteChainAdds = append(currentBatch.RemoteChainAdds, add)
				processedAdds[removal] = struct{}{}
			}
			batches, currentBatch = maybeSaveCurrentBatch(batches, currentBatch, input.BatchSize)
		}

		// Now, process the remaining additions (those not already processed)
		for _, add := range input.RemoteChainAdds {
			if _, ok := processedAdds[add.ChainSelector]; ok {
				continue
			}
			currentBatch.RemoteChainAdds = append(currentBatch.RemoteChainAdds, add)
			batches, currentBatch = maybeSaveCurrentBatch(batches, currentBatch, input.BatchSize)
		}

		// If any remaining items in the current batch, save it
		if len(currentBatch.RemoteChainRemoves) > 0 || len(currentBatch.RemoteChainAdds) > 0 {
			batches = append(batches, currentBatch)
		}
		writes := make([]contract.WriteOutput, 0)
		for _, batch := range batches {
			out, err := operations.ExecuteOperation(
				b,
				cciphomeops.ApplyChainConfigUpdates,
				deps.HomeChain,
				contract.FunctionInput[cciphomeops.ApplyChainConfigUpdatesOpInput]{
					Address:       input.CCIPHome,
					ChainSelector: deps.HomeChain.Selector,
					Args:          batch,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute ApplyChainConfigUpdatesOp on CCIPHome: %w", err)
			}
			writes = append(writes, out.Output)
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

func maybeSaveCurrentBatch(
	batches []cciphomeops.ApplyChainConfigUpdatesOpInput,
	currentBatch cciphomeops.ApplyChainConfigUpdatesOpInput,
	batchSize int,
) ([]cciphomeops.ApplyChainConfigUpdatesOpInput, cciphomeops.ApplyChainConfigUpdatesOpInput) {
	if len(currentBatch.RemoteChainRemoves)+len(currentBatch.RemoteChainAdds) >= batchSize {
		batches = append(batches, currentBatch)
		currentBatch = cciphomeops.ApplyChainConfigUpdatesOpInput{
			RemoteChainRemoves: make([]uint64, 0),
			RemoteChainAdds:    make([]ccip_home.CCIPHomeChainConfigArgs, 0),
		}
	}
	return batches, currentBatch
}

type AddCapabilityToCapRegConfig struct {
	HomeChainSel uint64
	CapReg       common.Address
	CCIPHome     common.Address
}

var SeqAddCapabilityToCapReg = operations.NewSequence(
	"add-capability-to-cap-reg",
	semver.MustParse("1.0.0"),
	"Adds CCIP capability to Capabilities Registry contract on the home chain.",
	func(b operations.Bundle, deps DONSequenceDeps, input AddCapabilityToCapRegConfig) (output sequences.OnChainOutput, err error) {
		chain := deps.HomeChain

		cr, err := capabilities_registry.NewCapabilitiesRegistry(input.CapReg, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate CapabilitiesRegistry contract: %w", err)
		}
		ccAddr := input.CCIPHome

		capabilities, err := cr.GetCapabilities(nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get capabilities: %w", err)
		}
		capabilityToAdd := capabilities_registry.CapabilitiesRegistryCapability{
			LabelledName:          CapabilityLabelledName,
			Version:               CapabilityVersion,
			CapabilityType:        2, // consensus. not used (?)
			ResponseType:          0, // report. not used (?)
			ConfigurationContract: ccAddr,
		}
		addCapability := true
		for _, cap := range capabilities {
			if cap.LabelledName == capabilityToAdd.LabelledName && cap.Version == capabilityToAdd.Version {
				b.Logger.Infow("Capability already exists, skipping adding capability",
					"labelledName", cap.LabelledName, "version", cap.Version)
				addCapability = false
				break
			}
		}
		var writes []contract.WriteOutput
		// Add the capability to the CapabilitiesRegistry contract only if it does not exist
		if addCapability {
			out, err := operations.ExecuteOperation(
				b,
				cciphomeops.AddCapabilities,
				deps.HomeChain,
				contract.FunctionInput[cciphomeops.AddCapabilitiesOpInput]{
					ChainSelector: chain.Selector,
					Address:       input.CapReg,
					Args: cciphomeops.AddCapabilitiesOpInput{
						Capabilities: []capabilities_registry.CapabilitiesRegistryCapability{
							capabilityToAdd,
						},
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AddCapabilities operation: %w", err)
			}
			writes = append(writes, out.Output)
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)

type AddNodeOperatorsToCapRegConfig struct {
	HomeChainSel  uint64
	CapReg        common.Address
	NodeOperators []capabilities_registry.CapabilitiesRegistryNodeOperator
}

var SeqAddNodeOperatorsToCapReg = operations.NewSequence(
	"add-node-operators-to-cap-reg",
	semver.MustParse("1.0.0"),
	"Adds node operators to Capabilities Registry contract on the home chain.",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input AddNodeOperatorsToCapRegConfig) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		chain := chains.EVMChains()[input.HomeChainSel]

		cr, err := capabilities_registry.NewCapabilitiesRegistry(input.CapReg, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate CapabilitiesRegistry contract: %w", err)
		}

		existingNodeOps, err := cr.GetNodeOperators(nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get existing node operators: %w", err)
		}
		nodeOpsMap := make(map[string]capabilities_registry.CapabilitiesRegistryNodeOperator)
		for _, nop := range input.NodeOperators {
			nodeOpsMap[nop.Admin.String()] = nop
		}
		for _, existingNop := range existingNodeOps {
			if _, ok := nodeOpsMap[existingNop.Admin.String()]; ok {
				b.Logger.Infow("Node operator already exists", "admin", existingNop.Admin.String())
				delete(nodeOpsMap, existingNop.Admin.String())
			}
		}
		nodeOpsToAdd := make([]capabilities_registry.CapabilitiesRegistryNodeOperator, 0, len(nodeOpsMap))
		for _, nop := range nodeOpsMap {
			nodeOpsToAdd = append(nodeOpsToAdd, nop)
		}
		if len(nodeOpsToAdd) > 0 {
			out, err := operations.ExecuteOperation(
				b,
				cciphomeops.AddNodeOperators,
				chain,
				contract.FunctionInput[cciphomeops.AddNodesOperatorsOpInput]{
					ChainSelector: chain.Selector,
					Address:       input.CapReg,
					Args: cciphomeops.AddNodesOperatorsOpInput{
						Nodes: nodeOpsToAdd,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AddNodeOperators operation: %w", err)
			}
			batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{out.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				Addresses: addresses,
				BatchOps:  []mcms_types.BatchOperation{batch},
			}, nil
		}
		b.Logger.Infow("No new node operators to add")
		return sequences.OnChainOutput{}, nil
	},
)

type AddNodesToCapRegConfig struct {
	HomeChainSel             uint64
	CapReg                   common.Address
	NodeP2PIDsPerNodeOpAdmin map[string][][32]byte
}

var SeqAddNodesToCapReg = operations.NewSequence(
	"update-home-chain",
	semver.MustParse("1.0.0"),
	"Updates Capabilities Registry contract on the home chain with new node operators and nodes.",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input AddNodesToCapRegConfig) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		chain := chains.EVMChains()[input.HomeChainSel]

		cr, err := capabilities_registry.NewCapabilitiesRegistry(input.CapReg, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate CapabilitiesRegistry contract: %w", err)
		}
		p2pIDsByNodeOpID, err := getNodesByNodeOpIDs(cr, input.NodeP2PIDsPerNodeOpAdmin)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get nodes by node operator ids: %w", err)
		}
		var nodeParams []capabilities_registry.CapabilitiesRegistryNodeParams

		for nopID, p2pIDs := range p2pIDsByNodeOpID {
			for _, p2pID := range p2pIDs {
				// if any p2pIDs are empty throw error
				if p2pID == ([32]byte{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("p2pID: %x selector: %d: %w", p2pID, chain.Selector, errors.New("empty p2pID"))
				}
				nodeParam := capabilities_registry.CapabilitiesRegistryNodeParams{
					NodeOperatorId:      nopID,
					Signer:              p2pID, // Not used in tests
					P2pId:               p2pID,
					EncryptionPublicKey: p2pID, // Not used in tests
					HashedCapabilityIds: [][32]byte{CCIPCapabilityID},
				}
				nodeParams = append(nodeParams, nodeParam)
			}
		}
		if len(nodeParams) == 0 {
			b.Logger.Infow("No new nodes to add")
			return sequences.OnChainOutput{}, nil
		}
		b.Logger.Infow("Adding nodes", "chain", chain.String(), "nodes", p2pIDsByNodeOpID)
		out, err := operations.ExecuteOperation(
			b,
			cciphomeops.AddNodes,
			chain,
			contract.FunctionInput[cciphomeops.AddNodesOpInput]{
				ChainSelector: chain.Selector,
				Address:       input.CapReg,
				Args: cciphomeops.AddNodesOpInput{
					Nodes: nodeParams,
				},
			},
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AddNodes operation: %w", err)
		}
		batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{out.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batch},
		}, nil
	},
)

func getNodesByNodeOpIDs(capReg *capabilities_registry.CapabilitiesRegistry, nodeP2PIDsPerNodeOpAdmin map[string][][32]byte) (map[uint32][][32]byte, error) {
	existingNodeOps, err := capReg.GetNodeOperators(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing node operators: %w", err)
	}
	// Need to fetch nodeoperators ids to be able to add nodes for corresponding node operators
	p2pIDsByNodeOpID := make(map[uint32][][32]byte)
	foundNopID := make(map[uint32]bool)
	for nopName, p2pID := range nodeP2PIDsPerNodeOpAdmin {
		// this is to find the node operator id for the given node operator name
		// node operator start from id 1, starting from 1 to len(existingNodeOps)
		totalNops := len(existingNodeOps)
		if totalNops >= math.MaxUint32 {
			return nil, errors.New("too many node operators")
		}
		for nopID := uint32(1); nopID <= uint32(totalNops); nopID++ {
			// if we already found the node operator id, skip
			if foundNopID[nopID] {
				continue
			}
			nodeOp, err := capReg.GetNodeOperator(nil, nopID)
			if err != nil {
				return nil, fmt.Errorf("failed to get node operator %d: %w", nopID, err)
			}
			if nodeOp.Name == nopName {
				p2pIDsByNodeOpID[nopID] = p2pID
				foundNopID[nopID] = true
				break
			}
		}
	}
	if len(p2pIDsByNodeOpID) != len(nodeP2PIDsPerNodeOpAdmin) {
		return nil, errors.New("failed to get nodes by nop id all node operators")
	}
	return p2pIDsByNodeOpID, nil
}
