package sequences

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"maps"
	"math"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	libocr_types "github.com/smartcontractkit/libocr/ragep2p/types"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"golang.org/x/crypto/sha3"
)

var Version = *semver.MustParse("1.6.0")

const (
	CapabilityLabelledName = "ccip"
	CapabilityVersion      = "v1.0.0"
)

var CCIPCapabilityID = Keccak256Fixed(MustABIEncode(`[{"type": "string"}, {"type": "string"}]`, CapabilityLabelledName, CapabilityVersion))

func MustABIEncode(abiString string, args ...any) []byte {
	encoded, err := ABIEncode(abiString, args...)
	if err != nil {
		panic(err)
	}
	return encoded
}

// ABIEncode is the equivalent of abi.encode.
// See a full set of examples https://github.com/ethereum/go-ethereum/blob/420b78659bef661a83c5c442121b13f13288c09f/accounts/abi/packing_test.go#L31
func ABIEncode(abiStr string, values ...interface{}) ([]byte, error) {
	// Create a dummy method with arguments
	inDef := fmt.Sprintf(`[{ "name" : "method", "type": "function", "inputs": %s}]`, abiStr)
	inAbi, err := abi.JSON(strings.NewReader(inDef))
	if err != nil {
		return nil, err
	}
	res, err := inAbi.Pack("method", values...)
	if err != nil {
		return nil, err
	}
	return res[4:], nil
}

func Keccak256Fixed(in []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	// Note this Keccak256 cannot error https://github.com/golang/crypto/blob/master/sha3/sha3.go#L126
	// if we start supporting hashing algos which do, we can change this API to include an error.
	hash.Write(in)
	var h [32]byte
	copy(h[:], hash.Sum(nil))
	return h
}

type DeployHomeChainConfig struct {
	HomeChainSel             uint64
	RMNStaticConfig          rmn_home.RMNHomeStaticConfig
	RMNDynamicConfig         rmn_home.RMNHomeDynamicConfig
	NodeOperators            []capabilities_registry.CapabilitiesRegistryNodeOperator
	NodeP2PIDsPerNodeOpAdmin map[string][][32]byte
}

var DeployHomeChain = cldf_ops.NewSequence(
	"deploy-home-chain",
	semver.MustParse("1.6.0"),
	"Deploys CCIP Home and Capabilities Registry contracts to the home chain.",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input DeployHomeChainConfig) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		chain := chains.EVMChains()[input.HomeChainSel]

		// Deploy Capabilities Registry
		crAddr, tx, capReg, err := capabilities_registry.DeployCapabilitiesRegistry(
			chain.DeployerKey,
			chain.Client,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, datastore.AddressRef{
			ChainSelector: input.HomeChainSel,
			Type:          datastore.ContractType(utils.CapabilitiesRegistry),
			Version:       &Version,
			Address:       crAddr.String(),
		})
		_, err = chain.Confirm(tx)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Deploy CCIPHome
		ccAddr, tx, _, err := ccip_home.DeployCCIPHome(
			chain.DeployerKey,
			chain.Client,
			crAddr,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, datastore.AddressRef{
			ChainSelector: input.HomeChainSel,
			Type:          datastore.ContractType(utils.CCIPHome),
			Version:       &Version,
			Address:       ccAddr.String(),
		})
		_, err = chain.Confirm(tx)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Deploy RMNHome
		rmnAddr, tx, rmnHome, err := rmn_home.DeployRMNHome(
			chain.DeployerKey,
			chain.Client,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, datastore.AddressRef{
			ChainSelector: input.HomeChainSel,
			Type:          datastore.ContractType(utils.RMNHome),
			Version:       &Version,
			Address:       rmnAddr.String(),
		})
		_, err = chain.Confirm(tx)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// considering the RMNHome is recently deployed, there is no digest to overwrite
		configs, err := rmnHome.GetAllConfigs(nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get all configs from RMNHome: %w", err)
		}
		setCandidate := false
		promoteCandidate := false

		// check if the candidate is already set and equal to static and dynamic configs
		if isRMNDynamicConfigEqual(input.RMNDynamicConfig, configs.CandidateConfig.DynamicConfig) &&
			isRMNStaticConfigEqual(input.RMNStaticConfig, configs.CandidateConfig.StaticConfig) {
			b.Logger.Infow("RMNHome candidate is already set and equal to given static and dynamic configs,skip setting candidate")
		} else {
			setCandidate = true
		}
		// check the active config is equal to the static and dynamic configs
		if isRMNDynamicConfigEqual(input.RMNDynamicConfig, configs.ActiveConfig.DynamicConfig) &&
			isRMNStaticConfigEqual(input.RMNStaticConfig, configs.ActiveConfig.StaticConfig) {
			b.Logger.Infow("RMNHome active is already set and equal to given static and dynamic configs," +
				"skip setting and promoting candidate")
			setCandidate = false
			promoteCandidate = false
		} else {
			promoteCandidate = true
		}

		if setCandidate {
			tx, err := rmnHome.SetCandidate(
				chain.DeployerKey, input.RMNStaticConfig, input.RMNDynamicConfig, configs.CandidateConfig.ConfigDigest)
			if _, err := cldf.ConfirmIfNoErrorWithABI(chain, tx, rmn_home.RMNHomeABI, err); err != nil {
				b.Logger.Errorw("Failed to set candidate on RMNHome", "err", err)
				return sequences.OnChainOutput{}, err
			}
			b.Logger.Infow("Set candidate on RMNHome", "chain", chain.String())
		}
		if promoteCandidate {
			rmnCandidateDigest, err := rmnHome.GetCandidateDigest(nil)
			if err != nil {
				b.Logger.Errorw("Failed to get RMNHome candidate digest", "chain", chain.String(), "err", err)
				return sequences.OnChainOutput{}, err
			}

			tx, err := rmnHome.PromoteCandidateAndRevokeActive(chain.DeployerKey, rmnCandidateDigest, [32]byte{})
			if _, err := cldf.ConfirmIfNoErrorWithABI(chain, tx, rmn_home.RMNHomeABI, err); err != nil {
				b.Logger.Errorw("Failed to promote candidate and revoke active on RMNHome", "chain", chain.String(), "err", err)
				return sequences.OnChainOutput{}, err
			}

			rmnActiveDigest, err := rmnHome.GetActiveDigest(nil)
			if err != nil {
				b.Logger.Errorw("Failed to get RMNHome active digest", "chain", chain.String(), "err", err)
				return sequences.OnChainOutput{}, err
			}
			b.Logger.Infow("Got rmn home active digest", "digest", rmnActiveDigest)

			if rmnActiveDigest != rmnCandidateDigest {
				b.Logger.Errorw("RMNHome active digest does not match previously candidate digest",
					"active", rmnActiveDigest, "candidate", rmnCandidateDigest)
				return sequences.OnChainOutput{}, errors.New("RMNHome active digest does not match candidate digest")
			}
			b.Logger.Infow("Promoted candidate and revoked active on RMNHome", "chain", chain.String())
		}
		// check if ccip capability exists in cap reg
		capabilities, err := capReg.GetCapabilities(nil)
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
		// Add the capability to the CapabilitiesRegistry contract only if it does not exist
		if addCapability {
			tx, err := capReg.AddCapabilities(
				chain.DeployerKey, []capabilities_registry.CapabilitiesRegistryCapability{
					capabilityToAdd,
				})
			if _, err := cldf.ConfirmIfNoErrorWithABI(chain, tx, capabilities_registry.CapabilitiesRegistryABI, err); err != nil {
				b.Logger.Errorw("Failed to add capabilities", "chain", chain.String(), "err", err)
				return sequences.OnChainOutput{}, err
			}
			b.Logger.Infow("Added capability to CapabilitiesRegistry",
				"labelledName", capabilityToAdd.LabelledName, "version", capabilityToAdd.Version)
		}

		existingNodeOps, err := capReg.GetNodeOperators(nil)
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
		// Need to fetch nodeoperators ids to be able to add nodes for corresponding node operators
		p2pIDsByNodeOpID := make(map[uint32][][32]byte)
		if len(nodeOpsToAdd) > 0 {
			tx, err := capReg.AddNodeOperators(chain.DeployerKey, input.NodeOperators)
			txBlockNum, err := cldf.ConfirmIfNoErrorWithABI(chain, tx, capabilities_registry.CapabilitiesRegistryABI, err)
			if err != nil {
				b.Logger.Errorw("Failed to add node operators", "chain", chain.String(), "err", err)
				return sequences.OnChainOutput{}, err
			}
			addedEvent, err := capReg.FilterNodeOperatorAdded(&bind.FilterOpts{
				Start:   txBlockNum,
				Context: context.Background(),
			}, nil, nil)
			if err != nil {
				b.Logger.Errorw("Failed to filter NodeOperatorAdded event", "chain", chain.String(), "err", err)
				return sequences.OnChainOutput{}, err
			}

			for addedEvent.Next() {
				for nopName, p2pID := range input.NodeP2PIDsPerNodeOpAdmin {
					if addedEvent.Event.Name == nopName {
						b.Logger.Infow("Added node operator", "admin", addedEvent.Event.Admin, "name", addedEvent.Event.Name)
						p2pIDsByNodeOpID[addedEvent.Event.NodeOperatorId] = p2pID
					}
				}
			}
		} else {
			b.Logger.Infow("No new node operators to add")
			foundNopID := make(map[uint32]bool)
			for nopName, p2pID := range input.NodeP2PIDsPerNodeOpAdmin {
				// this is to find the node operator id for the given node operator name
				// node operator start from id 1, starting from 1 to len(existingNodeOps)
				totalNops := len(existingNodeOps)
				if totalNops >= math.MaxUint32 {
					return sequences.OnChainOutput{}, errors.New("too many node operators")
				}
				for nopID := uint32(1); nopID <= uint32(totalNops); nopID++ {
					// if we already found the node operator id, skip
					if foundNopID[nopID] {
						continue
					}
					nodeOp, err := capReg.GetNodeOperator(nil, nopID)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to get node operator %d: %w", nopID, err)
					}
					if nodeOp.Name == nopName {
						p2pIDsByNodeOpID[nopID] = p2pID
						foundNopID[nopID] = true
						break
					}
				}
			}
		}
		if len(p2pIDsByNodeOpID) != len(input.NodeP2PIDsPerNodeOpAdmin) {
			b.Logger.Errorw("Failed to add all node operators", "added", maps.Keys(p2pIDsByNodeOpID), "expected", maps.Keys(input.NodeP2PIDsPerNodeOpAdmin), "chain", chain.String())
			return sequences.OnChainOutput{}, errors.New("failed to add all node operators")
		}
		// Adds initial set of nodes to CR, who all have the CCIP capability
		if err := addNodes(b.Logger, capReg, chain, p2pIDsByNodeOpID); err != nil {
			return sequences.OnChainOutput{}, err
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{},
		}, nil
	},
)

func isRMNStaticConfigEqual(a, b rmn_home.RMNHomeStaticConfig) bool {
	if len(a.Nodes) != len(b.Nodes) {
		return false
	}
	nodesByPeerID := make(map[libocr_types.PeerID]rmn_home.RMNHomeNode)
	for i := range a.Nodes {
		nodesByPeerID[a.Nodes[i].PeerId] = a.Nodes[i]
	}
	for i := range b.Nodes {
		node, exists := nodesByPeerID[b.Nodes[i].PeerId]
		if !exists {
			return false
		}
		if !bytes.Equal(node.OffchainPublicKey[:], b.Nodes[i].OffchainPublicKey[:]) {
			return false
		}
	}

	return bytes.Equal(a.OffchainConfig, b.OffchainConfig)
}

func isRMNDynamicConfigEqual(a, b rmn_home.RMNHomeDynamicConfig) bool {
	if len(a.SourceChains) != len(b.SourceChains) {
		return false
	}
	sourceChainBySelector := make(map[uint64]rmn_home.RMNHomeSourceChain)
	for i := range a.SourceChains {
		sourceChainBySelector[a.SourceChains[i].ChainSelector] = a.SourceChains[i]
	}
	for i := range b.SourceChains {
		sourceChain, exists := sourceChainBySelector[b.SourceChains[i].ChainSelector]
		if !exists {
			return false
		}
		if sourceChain.FObserve != b.SourceChains[i].FObserve {
			return false
		}
		if sourceChain.ObserverNodesBitmap.Cmp(b.SourceChains[i].ObserverNodesBitmap) != 0 {
			return false
		}
	}
	return bytes.Equal(a.OffchainConfig, b.OffchainConfig)
}

func addNodes(
	lggr logger.Logger,
	capReg *capabilities_registry.CapabilitiesRegistry,
	chain cldf_evm.Chain,
	p2pIDsByNodeOpId map[uint32][][32]byte,
) error {
	var nodeParams []capabilities_registry.CapabilitiesRegistryNodeParams
	nodes, err := capReg.GetNodes(nil)
	if err != nil {
		return err
	}
	existingNodeParams := make(map[libocr_types.PeerID]capabilities_registry.CapabilitiesRegistryNodeParams)
	for _, node := range nodes {
		existingNodeParams[node.P2pId] = capabilities_registry.CapabilitiesRegistryNodeParams{
			NodeOperatorId:      node.NodeOperatorId,
			Signer:              node.Signer,
			P2pId:               node.P2pId,
			EncryptionPublicKey: node.EncryptionPublicKey,
			HashedCapabilityIds: node.HashedCapabilityIds,
		}
	}
	for nopID, p2pIDs := range p2pIDsByNodeOpId {
		for _, p2pID := range p2pIDs {
			// if any p2pIDs are empty throw error
			if p2pID == ([32]byte{}) {
				return fmt.Errorf("p2pID: %x selector: %d: %w", p2pID, chain.Selector, errors.New("empty p2pID"))
			}
			nodeParam := capabilities_registry.CapabilitiesRegistryNodeParams{
				NodeOperatorId:      nopID,
				Signer:              p2pID, // Not used in tests
				P2pId:               p2pID,
				EncryptionPublicKey: p2pID, // Not used in tests
				HashedCapabilityIds: [][32]byte{CCIPCapabilityID},
			}
			if existing, ok := existingNodeParams[p2pID]; ok {
				if isEqualCapabilitiesRegistryNodeParams(existing, nodeParam) {
					lggr.Infow("Node already exists", "p2pID", p2pID)
					continue
				}
			}

			nodeParams = append(nodeParams, nodeParam)
		}
	}
	if len(nodeParams) == 0 {
		lggr.Infow("No new nodes to add")
		return nil
	}
	lggr.Infow("Adding nodes", "chain", chain.String(), "nodes", p2pIDsByNodeOpId)
	tx, err := capReg.AddNodes(chain.DeployerKey, nodeParams)
	if err != nil {
		lggr.Errorw("Failed to add nodes", "chain", chain.String(),
			"err", cldf.DecodedErrFromABIIfDataErr(err, capabilities_registry.CapabilitiesRegistryABI))
		return err
	}
	_, err = chain.Confirm(tx)
	return err
}

func isEqualCapabilitiesRegistryNodeParams(a, b capabilities_registry.CapabilitiesRegistryNodeParams) bool {
	if len(a.HashedCapabilityIds) != len(b.HashedCapabilityIds) {
		return false
	}
	for i := range a.HashedCapabilityIds {
		if !bytes.Equal(a.HashedCapabilityIds[i][:], b.HashedCapabilityIds[i][:]) {
			return false
		}
	}
	return a.NodeOperatorId == b.NodeOperatorId &&
		bytes.Equal(a.Signer[:], b.Signer[:]) &&
		bytes.Equal(a.P2pId[:], b.P2pId[:]) &&
		bytes.Equal(a.EncryptionPublicKey[:], b.EncryptionPublicKey[:])
}
