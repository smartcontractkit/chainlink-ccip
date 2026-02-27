package common

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"maps"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	focr "github.com/smartcontractkit/chainlink-deployments-framework/offchain/ocr"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	// tonSeqs "github.com/smartcontractkit/chainlink-ton/deployment/ccip/1_6_0/sequences"
	// ccip_ton "github.com/smartcontractkit/chainlink-ton/devenv"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/xssnick/tonutils-go/address"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"

	evmseqs "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	solanaSeqs "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type ChainConfig struct {
	Readers              [][32]byte              `json:"readers"`
	FChain               uint8                   `json:"fChain"`
	EncodableChainConfig chainconfig.ChainConfig `json:"encodableChainConfig"`
}

type UpdateChainConfigConfig struct {
	HomeChainSelector  uint64                 `json:"homeChainSelector"`
	RemoteChainRemoves []uint64               `json:"remoteChainRemoves"`
	RemoteChainAdds    map[uint64]ChainConfig `json:"remoteChainAdds"`
	MCMS               mcms.Input             `json:"mcmsConfig"`
}

type AddNodesToCapabilitiesRegistryChangesetConfig struct {
	HomeChainSelector        uint64                `json:"homeChainSelector"`
	NodeP2PIDsPerNodeOpAdmin map[string][][32]byte `json:"nodeP2pIdsPerNodeOpAdmin"`
	MCMS                     mcms.Input            `json:"mcmsConfig"`
}

type AddCapabilityToCapabilitiesRegistryChangesetConfig struct {
	HomeChainSelector uint64     `json:"homeChainSelector"`
	MCMS              mcms.Input `json:"mcmsConfig"`
}

type AddNodeOperatorsToCapabilitiesRegistryChangesetConfig struct {
	HomeChainSelector uint64 `json:"homeChainSelector"`
	Nop               []capabilities_registry.CapabilitiesRegistryNodeOperator
	MCMS              mcms.Input `json:"mcmsConfig"`
}

var UpdateChainConfig = deployment.CreateChangeSet(applyUpdateChainConfig, validateUpdateChainConfig)
var AddDONAndSetCandidate = deployment.CreateChangeSet(applyAddDonAndSetCandidateChangesetConfig, validateAddDonAndSetCandidateChangesetConfig)
var SetCandidate = deployment.CreateChangeSet(applySetCandidateChangesetConfig, validateSetCandidateChangesetConfig)
var PromoteCandidate = deployment.CreateChangeSet(applyPromoteCandidateChangesetConfig, validatePromoteCandidateChangesetConfig)
var AddNodesToCapabilitiesRegistry = deployment.CreateChangeSet(applyAddNodesToCapabilitiesRegistryChangeset, validateAddNodesToCapabilitiesRegistryChangeset)
var AddCapabilityToCapabilitiesRegistry = deployment.CreateChangeSet(applyAddCapabilityToCapabilitiesRegistryChangeset, validateAddCapabilityToCapabilitiesRegistryChangeset)
var AddNodeOperatorsToCapabilitiesRegistry = deployment.CreateChangeSet(applyAddNodeOperatorsToCapabilitiesRegistryChangeset, validateAddNodeOperatorsToCapabilitiesRegistryChangeset)

func validateAddNodeOperatorsToCapabilitiesRegistryChangeset(e deployment.Environment, cfg AddNodeOperatorsToCapabilitiesRegistryChangesetConfig) error {
	return nil
}

func applyAddNodeOperatorsToCapabilitiesRegistryChangeset(e deployment.Environment, cfg AddNodeOperatorsToCapabilitiesRegistryChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.0.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}
	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		SeqAddNodeOperatorsToCapReg,
		e.BlockChains,
		AddNodeOperatorsToCapRegConfig{
			HomeChainSel:  cfg.HomeChainSelector,
			CapReg:        common.HexToAddress(capRegAddr.Address),
			NodeOperators: cfg.Nop,
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("applying chain config updates: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	// since we are only using evm here
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}

func validateAddCapabilityToCapabilitiesRegistryChangeset(e deployment.Environment, cfg AddCapabilityToCapabilitiesRegistryChangesetConfig) error {
	return nil
}

func applyAddCapabilityToCapabilitiesRegistryChangeset(e deployment.Environment, cfg AddCapabilityToCapabilitiesRegistryChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.0.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}
	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}
	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		SeqAddCapabilityToCapReg,
		DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		AddCapabilityToCapRegConfig{
			HomeChainSel: cfg.HomeChainSelector,
			CapReg:       common.HexToAddress(capRegAddr.Address),
			CCIPHome:     common.HexToAddress(ccipHomeAddr.Address),
		})
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("applying chain config updates: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	// since we are only using evm here
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}

func validateAddNodesToCapabilitiesRegistryChangeset(e deployment.Environment, cfg AddNodesToCapabilitiesRegistryChangesetConfig) error {
	return nil
}

func applyAddNodesToCapabilitiesRegistryChangeset(e deployment.Environment, cfg AddNodesToCapabilitiesRegistryChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.0.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		SeqAddNodesToCapReg,
		e.BlockChains,
		AddNodesToCapRegConfig{
			HomeChainSel:             cfg.HomeChainSelector,
			CapReg:                   common.HexToAddress(capRegAddr.Address),
			NodeP2PIDsPerNodeOpAdmin: cfg.NodeP2PIDsPerNodeOpAdmin,
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("applying chain config updates: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	// since we are only using evm here
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}

func validateUpdateChainConfig(e deployment.Environment, cfg UpdateChainConfigConfig) error {
	return nil
}

func applyUpdateChainConfig(e deployment.Environment, cfg UpdateChainConfigConfig) (deployment.ChangesetOutput, error) {
	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}
	ccipHome, err := ccip_home.NewCCIPHome(
		common.HexToAddress(ccipHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("creating CCIPHome instance: %w", err)
	}

	// Create mapping of all removals to check if we are removing and re-adding chains
	removes := make(map[uint64]struct{}, len(cfg.RemoteChainRemoves))
	for _, chain := range cfg.RemoteChainRemoves {
		removes[chain] = struct{}{}
	}

	var adds []ccip_home.CCIPHomeChainConfigArgs
	for chain, ccfg := range cfg.RemoteChainAdds {
		encodedChainConfig, err := chainconfig.EncodeChainConfig(ccfg.EncodableChainConfig)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("encoding chain config: %w", err)
		}
		chainConfig := ccip_home.CCIPHomeChainConfig{
			Readers: ccfg.Readers,
			FChain:  ccfg.FChain,
			Config:  encodedChainConfig,
		}
		existingCfg, err := ccipHome.GetChainConfig(nil, chain)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("getting existing chain config for chain %d: %w", chain, err)
		}
		// Don't add chain configs again, unless we are removing and re-adding it.
		if _, ok := removes[chain]; !ok && isChainConfigEqual(existingCfg, chainConfig) {
			e.Logger.Infow("Chain config already exists, not applying again",
				"addedChain", chain,
				"chainConfig", chainConfig,
			)
			continue
		}
		adds = append(adds, ccip_home.CCIPHomeChainConfigArgs{
			ChainSelector: chain,
			ChainConfig:   chainConfig,
		})
	}
	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		ApplyChainConfigUpdatesSequence,
		DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		ApplyChainConfigUpdatesSequenceInput{
			CCIPHome:           ccipHome.Address(),
			RemoteChainAdds:    adds,
			RemoteChainRemoves: cfg.RemoteChainRemoves,
			BatchSize:          4, // Conservative batch size to avoid exceeding gas limits (TODO: Make this configurable)
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("applying chain config updates: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	// since we are only using evm here
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}

func isChainConfigEqual(a, b ccip_home.CCIPHomeChainConfig) bool {
	mapReader := make(map[[32]byte]struct{})
	for i := range a.Readers {
		mapReader[a.Readers[i]] = struct{}{}
	}
	for i := range b.Readers {
		if _, ok := mapReader[b.Readers[i]]; !ok {
			return false
		}
	}
	return bytes.Equal(a.Config, b.Config) &&
		a.FChain == b.FChain
}

type SetCandidatePluginInfo struct {
	// OCRConfigPerRemoteChainSelector is the chain selector of the chain where the DON will be added.
	OCRConfigPerRemoteChainSelector map[uint64]CCIPOCRParams `json:"ocrConfigPerRemoteChainSelector"`
	PluginType                      ccipocr3.PluginType      `json:"pluginType"`
}

type AddDonAndSetCandidateChangesetConfig struct {
	HomeChainSelector uint64 `json:"homeChainSelector"`
	FeedChainSelector uint64 `json:"feedChainSelector"`

	// Only set one plugin at a time while you are adding the DON for the first time.
	// For subsequent SetCandidate call use SetCandidateChangeset as that fetches the already added DONID and sets the candidate.
	PluginInfo     SetCandidatePluginInfo                        `json:"pluginInfo"`
	NonBootstraps  []*clclient.ChainlinkClient                   `json:"nonBootstraps"`
	NodeKeyBundles map[string]map[string]clclient.NodeKeysBundle `json:"nodeKeyBundles"`
	MCMS           mcms.Input                                    `json:"mcmsConfig"`
}

func validateAddDonAndSetCandidateChangesetConfig(e deployment.Environment, cfg AddDonAndSetCandidateChangesetConfig) error {
	return nil
}

func applyAddDonAndSetCandidateChangesetConfig(e deployment.Environment, cfg AddDonAndSetCandidateChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.0.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}
	capReg, err := capabilities_registry.NewCapabilitiesRegistry(
		common.HexToAddress(capRegAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}
	ccipHome, err := ccip_home.NewCCIPHome(
		common.HexToAddress(ccipHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}

	rmnHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.RMNHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding RMNHome address: %w", err)
	}
	rmnHome, err := rmn_home.NewRMNHome(
		common.HexToAddress(rmnHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	expectedDonID, err := capReg.GetNextDONId(nil)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("get next don id: %w", err)
	}

	var readers [][32]byte
	for _, node := range cfg.NonBootstraps {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to read p2p keys from node %s: %w", node.URL(), err)
		}
		id := MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
		readers = append(readers, id)
	}

	dons := make([]DONAddition, 0, len(cfg.PluginInfo.OCRConfigPerRemoteChainSelector))
	donUpdates := make([]DONUpdate, 0, len(cfg.PluginInfo.OCRConfigPerRemoteChainSelector))
	i := 0
	for chainSelector, params := range cfg.PluginInfo.OCRConfigPerRemoteChainSelector {
		id, err := DonIDForChain(capReg, ccipHome, chainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("getting don ID for chain selector %d: %w", chainSelector, err)
		}

		family, err := chain_selectors.GetSelectorFamily(chainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("getting chain family for selector %d: %w", chainSelector, err)
		}
		var offRampAddress []byte
		switch family {
		case chain_selectors.FamilyEVM:
			a := &evmseqs.EVMAdapter{}
			offRampAddress, err = a.GetOffRampAddress(e.DataStore, chainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
		case chain_selectors.FamilySolana:
			a := &solanaSeqs.SolanaAdapter{}
			offRampAddress, err = a.GetOffRampAddress(e.DataStore, chainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
		case chain_selectors.FamilyTon:
			// a := &tonSeqs.TonLaneAdapter{}
			// offRampAddress, err = a.GetOffRampAddress(e.DataStore, chainSelector)
			// if err != nil {
			// 	return deployment.ChangesetOutput{}, err
			// }
			panic("TON support temporarily disabled")
		default:
			return deployment.ChangesetOutput{}, fmt.Errorf("unsupported chain family %s for selector %d", family, chainSelector)
		}
		newDONArgs, err := BuildOCR3ConfigForCCIPHome(
			ccipHome,
			e.OCRSecrets,
			offRampAddress,
			chainSelector,
			cfg.NonBootstraps,
			rmnHome.Address(),
			params.OCRParameters,
			params.CommitOffChainConfig,
			params.ExecuteOffChainConfig,
			cfg.NodeKeyBundles,
		)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}

		pluginOCR3Config, ok := newDONArgs[cfg.PluginInfo.PluginType]
		if !ok {
			return deployment.ChangesetOutput{}, fmt.Errorf("missing plugin %s in ocr3Configs",
				cfg.PluginInfo.PluginType.String())
		}
		if id != 0 {
			e.Logger.Infow("DON already exists for chain selector, skipping addition, will just set candidate")
			existingDigest, err := ccipHome.GetCandidateDigest(&bind.CallOpts{
				Context: e.GetContext(),
			}, id, pluginOCR3Config.PluginType)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("get candidate digest from ccipHome: %w", err)
			}
			if existingDigest != [32]byte{} {
				e.Logger.Warnw("Overwriting existing candidate config", "digest", existingDigest, "donID", id, "pluginType", pluginOCR3Config.PluginType)
			}
			donUpdates = append(donUpdates, DONUpdate{
				ID:             id,
				PluginConfig:   pluginOCR3Config,
				PeerIDs:        readers,
				F:              uint8(len(cfg.NonBootstraps) / 3),
				IsPublic:       false,
				ExistingDigest: existingDigest,
			})
			continue
		}
		dons = append(dons, DONAddition{
			ExpectedID:       expectedDonID,
			PluginConfig:     pluginOCR3Config,
			PeerIDs:          readers,
			F:                uint8(len(cfg.NonBootstraps) / 3),
			IsPublic:         false,
			AcceptsWorkflows: false,
		})

		i++
		expectedDonID++
	}
	if len(dons) == 0 {
		e.Logger.Info("No new DONs to add, skipping AddDONAndSetCandidateChangeset, just setting candidate")
		result, err := operations.ExecuteSequence(
			e.OperationsBundle,
			SetCandidateSequence,
			DONSequenceDeps{
				HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
			},
			SetCandidateSequenceInput{
				CapabilitiesRegistry: capReg.Address(),
				DONs:                 donUpdates,
			},
		)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("setting candidate: %w", err)
		}
		mcmsReg := changesets.GetRegistry()
		mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
		return changesets.NewOutputBuilder(e, mcmsReg).
			WithReports(result.ExecutionReports).
			WithBatchOps(result.Output.BatchOps).
			Build(cfg.MCMS)
	}
	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		AddDONAndSetCandidateSequence,
		DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		AddDONAndSetCandidateSequenceInput{
			CapabilitiesRegistry: capReg.Address(),
			DONs:                 dons,
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("adding DON and setting candidate: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}

func BuildOCR3ConfigForCCIPHome(
	ccipHome *ccip_home.CCIPHome,
	ocrSecrets focr.OCRSecrets,
	offRampAddress []byte,
	destSelector uint64,
	nodes []*clclient.ChainlinkClient,
	rmnHomeAddress common.Address,
	ocrParams OCRParameters,
	commitOffchainCfg *pluginconfig.CommitOffchainConfig,
	execOffchainCfg *pluginconfig.ExecuteOffchainConfig,
	nodeKeyBundles map[string]map[string]clclient.NodeKeysBundle,
) (map[ccipocr3.PluginType]ccip_home.CCIPHomeOCR3Config, error) {
	schedule, oracles, err := getOracleIdentities(nodes, nodeKeyBundles, destSelector)
	if err != nil {
		return nil, err
	}

	// Add DON on capability registry contract
	ocr3Configs := make(map[ccipocr3.PluginType]ccip_home.CCIPHomeOCR3Config)
	pluginTypes := make([]ccipocr3.PluginType, 0)
	if commitOffchainCfg != nil {
		pluginTypes = append(pluginTypes, ccipocr3.PluginTypeCCIPCommit)
	}
	if execOffchainCfg != nil {
		pluginTypes = append(pluginTypes, ccipocr3.PluginTypeCCIPExec)
	}
	for _, pluginType := range pluginTypes {
		var encodedOffchainConfig []byte
		var err2 error
		if pluginType == ccipocr3.PluginTypeCCIPCommit {
			if commitOffchainCfg == nil {
				return nil, errors.New("commitOffchainCfg is nil")
			}
			encodedOffchainConfig, err2 = pluginconfig.EncodeCommitOffchainConfig(*commitOffchainCfg)
		} else {
			if execOffchainCfg == nil {
				return nil, errors.New("execOffchainCfg is nil")
			}
			encodedOffchainConfig, err2 = pluginconfig.EncodeExecuteOffchainConfig(*execOffchainCfg)
		}
		if err2 != nil {
			return nil, err2
		}
		signers, transmitters, configF, onchainConfig, offchainConfigVersion, offchainConfig, err2 := ocr3confighelper.ContractSetConfigArgsDeterministic(
			ocrSecrets.EphemeralSk,
			ocrSecrets.SharedSecret,
			ocrParams.DeltaProgress,
			ocrParams.DeltaResend,
			ocrParams.DeltaInitial,
			ocrParams.DeltaRound,
			ocrParams.DeltaGrace,
			ocrParams.DeltaCertifiedCommitRequest,
			ocrParams.DeltaStage,
			ocrParams.Rmax,
			schedule,
			oracles,
			encodedOffchainConfig,
			nil, // maxDurationInitialization
			ocrParams.MaxDurationQuery,
			ocrParams.MaxDurationObservation,
			ocrParams.MaxDurationShouldAcceptAttestedReport,
			ocrParams.MaxDurationShouldTransmitAcceptedReport,
			int(len(nodes)/3),
			[]byte{}, // empty OnChainConfig
		)
		if err2 != nil {
			return nil, err2
		}

		signersBytes := make([][]byte, len(signers))
		for i, signer := range signers {
			signersBytes[i] = signer
		}

		transmittersBytes := make([][]byte, len(transmitters))
		for i, transmitter := range transmitters {
			// TODO: this should just use the addresscodec
			family, err := chain_selectors.GetSelectorFamily(destSelector)
			if err != nil {
				return nil, err
			}
			var parsed []byte

			// if the node does not support the destination chain, the transmitter address is empty.
			if len(transmitter) == 0 {
				transmittersBytes[i] = []byte{}
				continue
			}

			switch family {
			case chain_selectors.FamilyEVM:
				parsed, err2 = common.ParseHexOrString(string(transmitter))
				if err2 != nil {
					return nil, err2
				}
			case chain_selectors.FamilySolana:
				pk, err := solana.PublicKeyFromBase58(string(transmitter))
				if err != nil {
					return nil, fmt.Errorf("failed to decode SVM address '%s': %w", transmitter, err)
				}
				parsed = pk.Bytes()
			case chain_selectors.FamilySui:
				parsed, err = hex.DecodeString(strings.TrimPrefix(string(transmitter), "0x"))
				if err != nil {
					return nil, fmt.Errorf("failed to decode SUI address '%s': %w", transmitter, err)
				}
			case chain_selectors.FamilyTon:
				pk := address.MustParseAddr(string(transmitter))
				if pk == nil || pk.IsAddrNone() {
					return nil, fmt.Errorf("failed to parse TON address '%s'", transmitter)
				}
				// TODO: this reimplements addrCodec's ToRawAddr helper
				parsed = binary.BigEndian.AppendUint32(nil, uint32(pk.Workchain())) //nolint:gosec // G115
				parsed = append(parsed, pk.Data()...)
			case chain_selectors.FamilyAptos:
				parsed, err = hex.DecodeString(strings.TrimPrefix(string(transmitter), "0x"))
				if err != nil {
					return nil, fmt.Errorf("failed to decode Aptos address '%s': %w", transmitter, err)
				}
			}

			transmittersBytes[i] = parsed
		}

		_, err := ocr3confighelper.PublicConfigFromContractConfig(false, ocrtypes.ContractConfig{
			Signers:               signers,
			Transmitters:          transmitters,
			F:                     configF,
			OnchainConfig:         onchainConfig,
			OffchainConfigVersion: offchainConfigVersion,
			OffchainConfig:        offchainConfig,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to validate ocr3 params: %w", err)
		}

		var ocrNodes []ccip_home.CCIPHomeOCR3Node
		for i := range nodes {
			id := MustPeerIDFromString(oracles[i].OracleIdentity.PeerID)
			ocrNodes = append(ocrNodes, ccip_home.CCIPHomeOCR3Node{
				P2pId:          id,
				SignerKey:      signersBytes[i],
				TransmitterKey: transmittersBytes[i],
			})
		}

		_, ok := ocr3Configs[pluginType]
		if ok {
			return nil, fmt.Errorf("pluginType %s already exists in ocr3Configs", pluginType.String())
		}

		ocr3Configs[pluginType] = ccip_home.CCIPHomeOCR3Config{
			PluginType:            uint8(pluginType),
			ChainSelector:         destSelector,
			FRoleDON:              configF,
			OffchainConfigVersion: offchainConfigVersion,
			OfframpAddress:        offRampAddress,
			Nodes:                 ocrNodes,
			OffchainConfig:        offchainConfig,
			RmnHomeAddress:        rmnHomeAddress.Bytes(),
		}
	}

	return ocr3Configs, nil
}

func getOracleIdentities(clClients []*clclient.ChainlinkClient, nodeKeyBundles map[string]map[string]clclient.NodeKeysBundle, destSelector uint64) ([]int, []confighelper.OracleIdentityExtra, error) { //nolint:gocritic
	s := make([]int, len(clClients))
	oracleIdentities := make([]confighelper.OracleIdentityExtra, len(clClients))
	eg := &errgroup.Group{}
	for i, cl := range clClients {
		eg.Go(func() error {
			keys, err := cl.MustReadP2PKeys()
			if err != nil {
				return err
			}
			id := MustPeerIDFromString(keys.Data[0].Attributes.PeerID)

			family, err := chain_selectors.GetSelectorFamily(destSelector)
			if err != nil {
				return err
			}
			var addr, offPrefix, onPrefix, cfgPrefix string
			var ocr2Config clclient.OCR2KeyAttributes
			switch family {
			case chain_selectors.FamilyEVM:
				destId, err := chain_selectors.GetChainIDFromSelector(destSelector)
				if err != nil {
					return err
				}
				addrSrc, err := cl.ReadPrimaryETHKey(destId)
				if err != nil {
					return err
				}
				addr = addrSrc.Attributes.Address
				ocr2Keys, err := cl.MustReadOCR2Keys()
				if err != nil {
					return err
				}

				for _, key := range ocr2Keys.Data {
					if key.Attributes.ChainType == "evm" {
						ocr2Config = key.Attributes
						break
					}
				}
				offPrefix = "ocr2off_evm_"
				onPrefix = "ocr2on_evm_"
				cfgPrefix = "ocr2cfg_evm_"
			case chain_selectors.FamilySolana:
				bundle := nodeKeyBundles[family][id.Raw()]
				addr = bundle.TXKey.Data.Attributes.PublicKey
				ocrKey := bundle.OCR2Key
				ocr2Config = ocrKey.Data.Attributes
				offPrefix = "ocr2off_solana_"
				onPrefix = "ocr2on_solana_"
				cfgPrefix = "ocr2cfg_solana_"
			case chain_selectors.FamilyTon:
				// bundle := nodeKeyBundles[family][id.Raw()]
				// addr, err = ccip_ton.GetNodeAddressFromBundle(&bundle)
				// if err != nil {
				// 	return err
				// }
				// ocrKey := bundle.OCR2Key
				// ocr2Config = ocrKey.Data.Attributes
				// offPrefix = "ocr2off_ton_"
				// onPrefix = "ocr2on_ton_"
				// cfgPrefix = "ocr2cfg_ton_"
				panic("TON support temporarily disabled")
			default:
				return fmt.Errorf("unsupported chain family %s for selector %d", family, destSelector)
			}

			offchainPkBytes, err := hex.DecodeString(strings.TrimPrefix(ocr2Config.OffChainPublicKey, offPrefix))
			if err != nil {
				return err
			}
			offchainPkBytesFixed := [ed25519.PublicKeySize]byte{}
			n := copy(offchainPkBytesFixed[:], offchainPkBytes)
			if n != ed25519.PublicKeySize {
				return fmt.Errorf("wrong number of elements copied (offchainPk, family = %s, %v %v %v)", family, ocr2Config.OffChainPublicKey, len(offchainPkBytes), offchainPkBytes)
			}
			configPkBytes, err := hex.DecodeString(strings.TrimPrefix(ocr2Config.ConfigPublicKey, cfgPrefix))
			if err != nil {
				return err
			}
			configPkBytesFixed := [ed25519.PublicKeySize]byte{}
			n = copy(configPkBytesFixed[:], configPkBytes)
			if n != ed25519.PublicKeySize {
				return fmt.Errorf("wrong number of elements copied (configPk, family = %s)", family)
			}
			onchainPkBytes, err := hex.DecodeString(strings.TrimPrefix(ocr2Config.OnChainPublicKey, onPrefix))
			if err != nil {
				return err
			}
			oracleIdentities[i] = confighelper.OracleIdentityExtra{
				OracleIdentity: confighelper.OracleIdentity{
					OnchainPublicKey:  onchainPkBytes,
					OffchainPublicKey: offchainPkBytesFixed,
					PeerID:            id.Raw(),
					TransmitAccount:   ocrtypes.Account(addr),
				},
				ConfigEncryptionPublicKey: configPkBytesFixed,
			}
			s[i] = 1
			return nil
		})
	}
	return s, oracleIdentities, eg.Wait()
}

func (p SetCandidatePluginInfo) String() string {
	allchains := maps.Keys(p.OCRConfigPerRemoteChainSelector)
	return fmt.Sprintf("PluginType: %s, Chains: %v", p.PluginType.String(), allchains)
}

type SetCandidateChangesetConfig struct {
	HomeChainSelector uint64 `json:"homeChainSelector"`
	FeedChainSelector uint64 `json:"feedChainSelector"`

	PluginInfo     []SetCandidatePluginInfo                      `json:"pluginInfo"`
	NonBootstraps  []*clclient.ChainlinkClient                   `json:"nonBootstraps"`
	NodeKeyBundles map[string]map[string]clclient.NodeKeysBundle `json:"nodeKeyBundles"`
	MCMS           mcms.Input                                    `json:"mcmsConfig"`
}

func validateSetCandidateChangesetConfig(e deployment.Environment, cfg SetCandidateChangesetConfig) error {
	return nil
}

func applySetCandidateChangesetConfig(e deployment.Environment, cfg SetCandidateChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.0.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}
	capReg, err := capabilities_registry.NewCapabilitiesRegistry(
		common.HexToAddress(capRegAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}
	ccipHome, err := ccip_home.NewCCIPHome(
		common.HexToAddress(ccipHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}

	rmnHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.RMNHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding RMNHome address: %w", err)
	}
	rmnHome, err := rmn_home.NewRMNHome(
		common.HexToAddress(rmnHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	pluginInfos := make([]string, 0)
	dons := make([]DONUpdate, 0)
	chainToDonIDs := make(map[uint64]uint32)
	for _, plugin := range cfg.PluginInfo {
		for chainSelector := range plugin.OCRConfigPerRemoteChainSelector {
			donID, err := DonIDForChain(
				capReg,
				ccipHome,
				chainSelector,
			)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("getting donID for chain %d: %w", chainSelector, err)
			}
			// if don doesn't exist use AddDonAndSetCandidateChangeset instead
			if donID == 0 {
				return deployment.ChangesetOutput{}, fmt.Errorf("don doesn't exist in CR for chain %d", chainSelector)
			}
			chainToDonIDs[chainSelector] = donID
		}
	}

	var readers [][32]byte
	for _, node := range cfg.NonBootstraps {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to read p2p keys from node %s: %w", node.URL(), err)
		}
		id := MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
		readers = append(readers, id)
	}

	for _, plugin := range cfg.PluginInfo {
		pluginInfos = append(pluginInfos, plugin.String())
		for chainSelector, params := range plugin.OCRConfigPerRemoteChainSelector {
			family, err := chain_selectors.GetSelectorFamily(chainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("getting chain family for selector %d: %w", chainSelector, err)
			}
			var offRampAddress []byte
			switch family {
			case chain_selectors.FamilyEVM:
				a := &evmseqs.EVMAdapter{}
				offRampAddress, err = a.GetOffRampAddress(e.DataStore, chainSelector)
				if err != nil {
					return deployment.ChangesetOutput{}, err
				}
			case chain_selectors.FamilySolana:
				a := &solanaSeqs.SolanaAdapter{}
				offRampAddress, err = a.GetOffRampAddress(e.DataStore, chainSelector)
				if err != nil {
					return deployment.ChangesetOutput{}, err
				}
			case chain_selectors.FamilyTon:
				// a := &tonSeqs.TonLaneAdapter{}
				// offRampAddress, err = a.GetOffRampAddress(e.DataStore, chainSelector)
				// if err != nil {
				// 	return deployment.ChangesetOutput{}, err
				// }
				panic("TON support temporarily disabled")
			default:
				return deployment.ChangesetOutput{}, fmt.Errorf("unsupported chain family %s for selector %d", family, chainSelector)
			}
			newDONArgs, err := BuildOCR3ConfigForCCIPHome(
				ccipHome,
				e.OCRSecrets,
				offRampAddress,
				chainSelector,
				cfg.NonBootstraps,
				rmnHome.Address(),
				params.OCRParameters,
				params.CommitOffChainConfig,
				params.ExecuteOffChainConfig,
				cfg.NodeKeyBundles,
			)
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}

			config, ok := newDONArgs[plugin.PluginType]
			if !ok {
				return deployment.ChangesetOutput{}, fmt.Errorf("missing %s plugin in ocr3Configs", plugin.PluginType.String())
			}

			// get the current candidate config
			donID := chainToDonIDs[chainSelector]
			existingDigest, err := ccipHome.GetCandidateDigest(&bind.CallOpts{
				Context: e.GetContext(),
			}, donID, config.PluginType)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("get candidate digest from ccipHome: %w", err)
			}
			if existingDigest != [32]byte{} {
				e.Logger.Warnw("Overwriting existing candidate config", "digest", existingDigest, "donID", donID, "pluginType", config.PluginType)
			}

			dons = append(dons, DONUpdate{
				ID:             donID,
				PluginConfig:   config,
				PeerIDs:        readers,
				F:              uint8(len(cfg.NonBootstraps) / 3),
				IsPublic:       false,
				ExistingDigest: existingDigest,
			})
		}
	}
	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		SetCandidateSequence,
		DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		SetCandidateSequenceInput{
			CapabilitiesRegistry: capReg.Address(),
			DONs:                 dons,
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("setting candidate config: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}

func DonIDForChain(registry *capabilities_registry.CapabilitiesRegistry, ccipHome *ccip_home.CCIPHome, chainSelector uint64) (uint32, error) {
	nextDonID, err := registry.GetNextDONId(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return 0, fmt.Errorf("get next don id from capability registry: %w", err)
	}

	var donID uint32
	for donId := uint32(1); donId < nextDonID; donId++ {
		don, err := registry.GetDON(&bind.CallOpts{
			Context: context.Background(),
		}, donId)
		if err != nil {
			return 0, fmt.Errorf("get don %d from capability registry: %w", donId, err)
		}
		if len(don.CapabilityConfigurations) == 1 &&
			don.CapabilityConfigurations[0].CapabilityId == CCIPCapabilityID {
			// Another possible way is to query GetAllConfigs from CCIPHome, but the return data for that would be more expensive.
			// So we first filter DONs by checking if they have only CCIP capability in the capability registry.
			// Then we fetch the config digests from CCIPHome to check the chain selector.
			// This requires multiple calls, but those are cheaper than fetching all configs in one call.
			digests, err := ccipHome.GetConfigDigests(nil, don.Id, uint8(ccipocr3.PluginTypeCCIPCommit))
			if err != nil {
				return 0, fmt.Errorf("get commit config digests from cciphome: %w", err)
			}
			pluginType := uint8(ccipocr3.PluginTypeCCIPCommit)
			if digests.ActiveConfigDigest == [32]byte{} && digests.CandidateConfigDigest == [32]byte{} {
				digests, err = ccipHome.GetConfigDigests(nil, don.Id, uint8(ccipocr3.PluginTypeCCIPExec))
				if err != nil {
					return 0, fmt.Errorf("get exec config digests from cciphome: %w", err)
				}
				pluginType = uint8(ccipocr3.PluginTypeCCIPExec)
			}
			activeConfig, err := ccipHome.GetConfig(&bind.CallOpts{
				Context: context.Background(),
			}, don.Id, pluginType, digests.ActiveConfigDigest)
			if err != nil {
				return 0, fmt.Errorf("get active config from cciphome: %w", err)
			}
			candidateConfig, err := ccipHome.GetConfig(&bind.CallOpts{
				Context: context.Background(),
			}, don.Id, pluginType, digests.CandidateConfigDigest)
			if err != nil {
				return 0, fmt.Errorf("get candidate config from cciphome: %w", err)
			}
			if activeConfig.VersionedConfig.Config.ChainSelector == chainSelector || candidateConfig.VersionedConfig.Config.ChainSelector == chainSelector {
				donID = don.Id
				break
			}
		}
	}

	return donID, nil
}

type PromoteCandidatePluginInfo struct {
	// RemoteChainSelectors is the chain selector of the DONs that we want to promote the candidate config of.
	// Note that each (chain, ccip capability version) pair has a unique DON ID.
	RemoteChainSelectors    []uint64            `json:"remoteChainSelectors"`
	PluginType              ccipocr3.PluginType `json:"pluginType"`
	AllowEmptyConfigPromote bool                `json:"allowEmptyConfigPromote"` // safe guard to prevent promoting empty config to active
}

type PromoteCandidateChangesetConfig struct {
	HomeChainSelector uint64 `json:"homeChainSelector"`

	PluginInfo    []PromoteCandidatePluginInfo `json:"pluginInfo"`
	NonBootstraps []*clclient.ChainlinkClient  `json:"nonBootstraps"`
	MCMS          mcms.Input                   `json:"mcmsConfig"`
}

func validatePromoteCandidateChangesetConfig(e deployment.Environment, cfg PromoteCandidateChangesetConfig) error {
	return nil
}

func applyPromoteCandidateChangesetConfig(e deployment.Environment, cfg PromoteCandidateChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.0.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}
	capReg, err := capabilities_registry.NewCapabilitiesRegistry(
		common.HexToAddress(capRegAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}
	ccipHome, err := ccip_home.NewCCIPHome(
		common.HexToAddress(ccipHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
	}

	donIDs := make(map[uint64]uint32)
	for _, plugin := range cfg.PluginInfo {
		if plugin.PluginType != ccipocr3.PluginTypeCCIPCommit &&
			plugin.PluginType != ccipocr3.PluginTypeCCIPExec {
			return deployment.ChangesetOutput{}, errors.New("PluginType must be set to either CCIPCommit or CCIPExec")
		}
		for _, chainSelector := range plugin.RemoteChainSelectors {
			donID, err := DonIDForChain(
				capReg,
				ccipHome,
				chainSelector,
			)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("fetch don id for chain: %w", err)
			}
			if donID == 0 {
				return deployment.ChangesetOutput{}, fmt.Errorf("don doesn't exist in CR for chain %d", chainSelector)
			}
			// Check that candidate digest and active digest are not both zero - this is enforced onchain.
			candidateDigest, err := ccipHome.GetCandidateDigest(&bind.CallOpts{
				Context: e.GetContext(),
			}, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("fetching %s candidate digest from cciphome: %w", plugin.PluginType.String(), err)
			}
			activeDigest, err := ccipHome.GetActiveDigest(&bind.CallOpts{
				Context: e.GetContext(),
			}, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("fetching %s active digest from cciphome: %w", plugin.PluginType.String(), err)
			}
			// If promoteCandidate is called with AllowEmptyConfigPromote set to false and
			// the CandidateConfig config digest is zero, do not promote the candidate config to active.
			if !plugin.AllowEmptyConfigPromote && candidateDigest == [32]byte{} {
				return deployment.ChangesetOutput{}, fmt.Errorf("%s candidate config digest is empty", plugin.PluginType.String())
			}

			// If the active and candidate config digests are both zero, we should not promote the candidate config to active.
			if activeDigest == [32]byte{} &&
				candidateDigest == [32]byte{} {
				return deployment.ChangesetOutput{}, fmt.Errorf("%s active and candidate config digests are both zero", plugin.PluginType.String())
			}
			donIDs[chainSelector] = donID
		}
	}

	dons := make([]DONUpdatePromotion, 0)
	var readers [][32]byte
	for _, node := range cfg.NonBootstraps {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to read p2p keys from node %s: %w", node.URL(), err)
		}
		id := MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
		readers = append(readers, id)
	}
	for _, plugin := range cfg.PluginInfo {
		for sel, donID := range donIDs {
			candidateDigest, err := ccipHome.GetCandidateDigest(nil, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
			if candidateDigest == [32]byte{} && !plugin.AllowEmptyConfigPromote {
				return deployment.ChangesetOutput{}, errors.New("candidate config digest is zero, promoting empty config is not allowed")
			}
			activeDigest, err := ccipHome.GetActiveDigest(nil, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
			if activeDigest == candidateDigest {
				e.Logger.Infow("Active and candidate config digests are the same, skipping promotion for plugin "+plugin.PluginType.String(), "digest", candidateDigest)
				continue
			}
			e.Logger.Infow("Promoting candidate for plugin "+plugin.PluginType.String(), "digest", candidateDigest)

			dons = append(dons, DONUpdatePromotion{
				ID:              donID,
				PluginType:      uint8(plugin.PluginType),
				ChainSelector:   sel,
				PeerIDs:         readers,
				F:               uint8(len(cfg.NonBootstraps) / 3),
				IsPublic:        false,
				CandidateDigest: candidateDigest,
				ActiveDigest:    activeDigest,
			})
		}
	}

	result, err := operations.ExecuteSequence(
		e.OperationsBundle,
		PromoteCandidateSequence,
		DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		PromoteCandidateSequenceInput{
			CapabilitiesRegistry: capReg.Address(),
			DONs:                 dons,
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("promoting candidate config: %w", err)
	}
	mcmsReg := changesets.GetRegistry()
	mcmsReg.RegisterMCMSReader(chain_selectors.FamilyEVM, &adapters.EVMMCMSReader{})
	return changesets.NewOutputBuilder(e, mcmsReg).
		WithReports(result.ExecutionReports).
		WithBatchOps(result.Output.BatchOps).
		Build(cfg.MCMS)
}
