package changesets

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	focr "github.com/smartcontractkit/chainlink-deployments-framework/offchain/ocr"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/xssnick/tonutils-go/address"
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
}

var UpdateChainConfig = deployment.CreateChangeSet(applyUpdateChainConfig, validateUpdateChainConfig)
var AddDONAndSetCandidate = deployment.CreateChangeSet(applyAddDonAndSetCandidateChangesetConfig, validateAddDonAndSetCandidateChangesetConfig)
var SetCandidate = deployment.CreateChangeSet(applySetCandidateChangesetConfig, validateSetCandidateChangesetConfig)
var PromoteCandidate = deployment.CreateChangeSet(applyUpdateChainConfig, validateUpdateChainConfig)

func validateUpdateChainConfig(e deployment.Environment, cfg UpdateChainConfigConfig) error {
	return nil
}

func applyUpdateChainConfig(e deployment.Environment, cfg UpdateChainConfigConfig) (deployment.ChangesetOutput, error) {
	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
	ccipHome, err := ccip_home.NewCCIPHome(
		common.HexToAddress(ccipHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CCIPHome address: %w", err)
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
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.ApplyChainConfigUpdatesSequence,
		sequences.DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		sequences.ApplyChainConfigUpdatesSequenceInput{
			CCIPHome:           ccipHome.Address(),
			RemoteChainAdds:    adds,
			RemoteChainRemoves: cfg.RemoteChainRemoves,
			BatchSize:          4, // Conservative batch size to avoid exceeding gas limits (TODO: Make this configurable)
		},
	)
	return deployment.ChangesetOutput{}, err
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

type OCRParameters struct {
	DeltaProgress                           time.Duration `json:"deltaProgress"`
	DeltaResend                             time.Duration `json:"deltaResend"`
	DeltaInitial                            time.Duration `json:"deltaInitial"`
	DeltaRound                              time.Duration `json:"deltaRound"`
	DeltaGrace                              time.Duration `json:"deltaGrace"`
	DeltaCertifiedCommitRequest             time.Duration `json:"deltaCertifiedCommitRequest"`
	DeltaStage                              time.Duration `json:"deltaStage"`
	Rmax                                    uint64        `json:"rmax"`
	MaxDurationQuery                        time.Duration `json:"maxDurationQuery"`
	MaxDurationObservation                  time.Duration `json:"maxDurationObservation"`
	MaxDurationShouldAcceptAttestedReport   time.Duration `json:"maxDurationShouldAcceptAttestedReport"`
	MaxDurationShouldTransmitAcceptedReport time.Duration `json:"maxDurationShouldTransmitAcceptedReport"`
}

type CCIPOCRParams struct {
	// OCRParameters contains the parameters for the OCR plugin.
	OCRParameters OCRParameters `json:"ocrParameters"`
	// CommitOffChainConfig contains pointers to Arb feeds for prices.
	CommitOffChainConfig *pluginconfig.CommitOffchainConfig `json:"commitOffChainConfig,omitempty"`
	// ExecuteOffChainConfig contains USDC config.
	ExecuteOffChainConfig *pluginconfig.ExecuteOffchainConfig `json:"executeOffChainConfig,omitempty"`
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
	PluginInfo SetCandidatePluginInfo `json:"pluginInfo"`
}

func validateAddDonAndSetCandidateChangesetConfig(e deployment.Environment, cfg AddDonAndSetCandidateChangesetConfig) error {
	return nil
}

func applyAddDonAndSetCandidateChangesetConfig(e deployment.Environment, cfg AddDonAndSetCandidateChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
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
	rmnHome, err := rmn_home.NewRMNHome(
		common.HexToAddress(rmnHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}
	nodes, err := NodeInfo(e.NodeIDs, e.Offchain)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("get node info: %w", err)
	}
	nonBootstraps := nodes.NonBootstraps()

	expectedDonID, err := capReg.GetNextDONId(nil)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("get next don id: %w", err)
	}

	dons := make([]sequences.DONAddition, len(cfg.PluginInfo.OCRConfigPerRemoteChainSelector))
	i := 0
	a := &sequences.EVMAdapter{}
	for chainSelector, params := range cfg.PluginInfo.OCRConfigPerRemoteChainSelector {
		offRampAddress, err := a.GetOffRampAddress(e.DataStore, chainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}
		newDONArgs, err := BuildOCR3ConfigForCCIPHome(
			ccipHome,
			e.OCRSecrets,
			offRampAddress,
			chainSelector,
			nonBootstraps,
			rmnHome.Address(),
			params.OCRParameters,
			params.CommitOffChainConfig,
			params.ExecuteOffChainConfig,
		)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}

		pluginOCR3Config, ok := newDONArgs[cfg.PluginInfo.PluginType]
		if !ok {
			return deployment.ChangesetOutput{}, fmt.Errorf("missing plugin %s in ocr3Configs",
				cfg.PluginInfo.PluginType.String())
		}

		dons[i] = sequences.DONAddition{
			ExpectedID:       expectedDonID,
			PluginConfig:     pluginOCR3Config,
			PeerIDs:          nonBootstraps.PeerIDs(),
			F:                nonBootstraps.DefaultF(),
			IsPublic:         false,
			AcceptsWorkflows: false,
		}
		i++
		expectedDonID++
	}

	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.AddDONAndSetCandidateSequence,
		sequences.DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		sequences.AddDONAndSetCandidateSequenceInput{
			CapabilitiesRegistry: capReg.Address(),
			DONs:                 dons,
		},
	)
	return deployment.ChangesetOutput{}, err
}

func BuildOCR3ConfigForCCIPHome(
	ccipHome *ccip_home.CCIPHome,
	ocrSecrets focr.OCRSecrets,
	offRampAddress []byte,
	destSelector uint64,
	nodes Nodes,
	rmnHomeAddress common.Address,
	ocrParams OCRParameters,
	commitOffchainCfg *pluginconfig.CommitOffchainConfig,
	execOffchainCfg *pluginconfig.ExecuteOffchainConfig,
) (map[ccipocr3.PluginType]ccip_home.CCIPHomeOCR3Config, error) {
	// check if we have info from this node for another chain in the same destFamily
	destFamily, err := chain_selectors.GetSelectorFamily(destSelector)
	if err != nil {
		return nil, err
	}

	var p2pIDs [][32]byte
	// Get OCR3 Config from helper
	var schedule []int
	var oracles []confighelper.OracleIdentityExtra
	for _, node := range nodes {
		schedule = append(schedule, 1)

		// TODO: not every node supports the destination chain, but nodes must have an OCR identity for the
		// destination chain, in order to be able to participate in the OCR protocol, sign reports, etc.
		// However, JD currently only returns the "OCRConfig" for chains that are explicitly supported by the node,
		// presumably in the TOML config.
		// JD should instead give us the OCR identity for the destination chain, and, if the node does NOT
		// actually support the chain (in terms of TOML config), then return an empty transmitter address,
		// which is what we're supposed to set anyway if that particular node doesn't support the destination chain.
		// The current workaround is to check if we have the OCR identity for the destination chain based off of
		// the node's OCR identity for another chain in the same family.
		// This is a HACK, because it is entirely possible that the destination chain is a unique family,
		// and no other supported chain by the node has the same family, e.g. Solana.
		cfg, exists := node.OCRConfigForChainSelector(destSelector)
		if !exists {
			// check if we have an oracle identity for another chain in the same family as destFamily.
			allOCRConfigs := node.AllOCRConfigs()
			for chainDetails, ocrConfig := range allOCRConfigs {
				chainFamily, err := chain_selectors.GetSelectorFamily(chainDetails.ChainSelector)
				if err != nil {
					return nil, err
				}

				if chainFamily == destFamily {
					cfg = ocrConfig
					break
				}
			}

			if cfg.OffchainPublicKey == [32]byte{} {
				return nil, fmt.Errorf(
					"no OCR config for chain %d (family %s) from node %s (peer id %s) and no other OCR config for another chain in the same family",
					destSelector, destFamily, node.Name, node.PeerID.String(),
				)
			}
		}

		var transmitAccount ocrtypes.Account
		if !exists {
			// empty account means that the node cannot transmit for this chain
			// we replace this with a canonical address with the oracle ID as the address when doing the ocr config validation below, but it should remain empty
			// in the CCIPHome OCR config and it should not be included in the destination chain transmitters whitelist.
			transmitAccount = ocrtypes.Account("")
		} else {
			transmitAccount = cfg.TransmitAccount
		}

		p2pIDs = append(p2pIDs, node.PeerID)
		oracles = append(oracles, confighelper.OracleIdentityExtra{
			OracleIdentity: confighelper.OracleIdentity{
				OnchainPublicKey:  cfg.OnchainPublicKey,    // should be the same for all chains within the same family
				TransmitAccount:   transmitAccount,         // different per chain (!) can be empty if the node does not support the destination chain
				OffchainPublicKey: cfg.OffchainPublicKey,   // should be the same for all chains within the same family
				PeerID:            cfg.PeerID.String()[4:], // should be the same for all oracle identities
			},
			ConfigEncryptionPublicKey: cfg.ConfigEncryptionPublicKey, // should be the same for all chains within the same family
		})
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
			int(nodes.DefaultF()),
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

		_, err = ocr3confighelper.PublicConfigFromContractConfig(false, ocrtypes.ContractConfig{
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
			ocrNodes = append(ocrNodes, ccip_home.CCIPHomeOCR3Node{
				P2pId:          p2pIDs[i],
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

		// chainConfig, err := ccipHome.GetChainConfig(nil, destSelector)
		// if err != nil {
		// 	return nil, fmt.Errorf("can't get chain config for %d: %w", destSelector, err)
		// }
		// if err := validateOCR3Config(destSelector, ocr3Configs[pluginType], &chainConfig); err != nil {
		// 	return nil, fmt.Errorf("failed to validate ocr3 config: %w", err)
		// }
	}

	return ocr3Configs, nil
}

func (p SetCandidatePluginInfo) String() string {
	allchains := maps.Keys(p.OCRConfigPerRemoteChainSelector)
	return fmt.Sprintf("PluginType: %s, Chains: %v", p.PluginType.String(), allchains)
}

type SetCandidateChangesetConfig struct {
	HomeChainSelector uint64 `json:"homeChainSelector"`
	FeedChainSelector uint64 `json:"feedChainSelector"`

	PluginInfo []SetCandidatePluginInfo `json:"pluginInfo"`
}

func validateSetCandidateChangesetConfig(e deployment.Environment, cfg SetCandidateChangesetConfig) error {
	return nil
}

func applySetCandidateChangesetConfig(e deployment.Environment, cfg SetCandidateChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
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
	rmnHome, err := rmn_home.NewRMNHome(
		common.HexToAddress(rmnHomeAddr.Address),
		e.BlockChains.EVMChains()[cfg.HomeChainSelector].Client)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("finding CapabilitiesRegistry address: %w", err)
	}

	nodes, err := NodeInfo(e.NodeIDs, e.Offchain)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("get node info: %w", err)
	}
	nonBootstraps := nodes.NonBootstraps()

	pluginInfos := make([]string, 0)
	dons := make([]sequences.DONUpdate, 0)
	a := &sequences.EVMAdapter{}
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

	for _, plugin := range cfg.PluginInfo {
		pluginInfos = append(pluginInfos, plugin.String())
		for chainSelector, params := range plugin.OCRConfigPerRemoteChainSelector {
			offRampAddress, err := a.GetOffRampAddress(e.DataStore, chainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
			newDONArgs, err := BuildOCR3ConfigForCCIPHome(
				ccipHome,
				e.OCRSecrets,
				offRampAddress,
				chainSelector,
				nodes.NonBootstraps(),
				rmnHome.Address(),
				params.OCRParameters,
				params.CommitOffChainConfig,
				params.ExecuteOffChainConfig,
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

			dons = append(dons, sequences.DONUpdate{
				ID:             donID,
				PluginConfig:   config,
				PeerIDs:        nonBootstraps.PeerIDs(),
				F:              nonBootstraps.DefaultF(),
				IsPublic:       false,
				ExistingDigest: existingDigest,
			})
		}
	}
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.SetCandidateSequence,
		sequences.DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		sequences.SetCandidateSequenceInput{
			CapabilitiesRegistry: capReg.Address(),
			DONs:                 dons,
		},
	)
	return deployment.ChangesetOutput{}, err
}

func DonIDForChain(registry *capabilities_registry.CapabilitiesRegistry, ccipHome *ccip_home.CCIPHome, chainSelector uint64) (uint32, error) {
	dons, err := registry.GetDONs(nil)
	if err != nil {
		return 0, fmt.Errorf("get Dons from capability registry: %w", err)
	}
	var donIDs []uint32
	for _, don := range dons {
		if len(don.CapabilityConfigurations) == 1 &&
			don.CapabilityConfigurations[0].CapabilityId == sequences.CCIPCapabilityID {
			configs, err := ccipHome.GetAllConfigs(nil, don.Id, uint8(ccipocr3.PluginTypeCCIPCommit))
			if err != nil {
				return 0, fmt.Errorf("get all commit configs from cciphome: %w", err)
			}
			if configs.ActiveConfig.ConfigDigest == [32]byte{} && configs.CandidateConfig.ConfigDigest == [32]byte{} {
				configs, err = ccipHome.GetAllConfigs(nil, don.Id, uint8(ccipocr3.PluginTypeCCIPExec))
				if err != nil {
					return 0, fmt.Errorf("get all exec configs from cciphome: %w", err)
				}
			}
			if configs.ActiveConfig.Config.ChainSelector == chainSelector || configs.CandidateConfig.Config.ChainSelector == chainSelector {
				donIDs = append(donIDs, don.Id)
			}
		}
	}

	// more than one DON is an error
	if len(donIDs) > 1 {
		return 0, fmt.Errorf("more than one DON found for (chain selector %d, ccip capability id %x) pair", chainSelector, sequences.CCIPCapabilityID[:])
	}

	// no DON found - don ID of 0 indicates that (this is the case in the CR as well).
	if len(donIDs) == 0 {
		return 0, nil
	}

	// DON found - return it.
	return donIDs[0], nil
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

	PluginInfo []PromoteCandidatePluginInfo `json:"pluginInfo"`
}

func validatePromoteCandidateChangesetConfig(e deployment.Environment, cfg PromoteCandidateChangesetConfig) error {
	return nil
}

func applyPromoteCandidateChangesetConfig(e deployment.Environment, cfg PromoteCandidateChangesetConfig) (deployment.ChangesetOutput, error) {
	capRegAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: cfg.HomeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       semver.MustParse("1.6.0"),
	}, cfg.HomeChainSelector, datastore_utils.FullRef)
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
			pluginConfigs, err := ccipHome.GetAllConfigs(&bind.CallOpts{
				Context: e.GetContext(),
			}, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("fetching %s configs from cciphome: %w", plugin.PluginType.String(), err)
			}
			// If promoteCandidate is called with AllowEmptyConfigPromote set to false and
			// the CandidateConfig config digest is zero, do not promote the candidate config to active.
			if !plugin.AllowEmptyConfigPromote && pluginConfigs.CandidateConfig.ConfigDigest == [32]byte{} {
				return deployment.ChangesetOutput{}, fmt.Errorf("%s candidate config digest is empty", plugin.PluginType.String())
			}

			// If the active and candidate config digests are both zero, we should not promote the candidate config to active.
			if pluginConfigs.ActiveConfig.ConfigDigest == [32]byte{} &&
				pluginConfigs.CandidateConfig.ConfigDigest == [32]byte{} {
				return deployment.ChangesetOutput{}, fmt.Errorf("%s active and candidate config digests are both zero", plugin.PluginType.String())
			}
			donIDs[chainSelector] = donID
		}
	}

	nodes, err := NodeInfo(e.NodeIDs, e.Offchain)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("fetch node info: %w", err)
	}
	nonBootstraps := nodes.NonBootstraps()

	dons := make([]sequences.DONUpdatePromotion, 0)
	for _, plugin := range cfg.PluginInfo {
		for _, donID := range donIDs {
			digest, err := ccipHome.GetCandidateDigest(nil, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
			if digest == [32]byte{} && !plugin.AllowEmptyConfigPromote {
				return deployment.ChangesetOutput{}, errors.New("candidate config digest is zero, promoting empty config is not allowed")
			}
			allConfigs, err := ccipHome.GetAllConfigs(nil, donID, uint8(plugin.PluginType))
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
			e.Logger.Infow("Promoting candidate for plugin "+plugin.PluginType.String(), "digest", digest)

			dons = append(dons, sequences.DONUpdatePromotion{
				ID:              donID,
				PluginType:      uint8(plugin.PluginType),
				ChainSelector:   allConfigs.CandidateConfig.Config.ChainSelector,
				PeerIDs:         nonBootstraps.PeerIDs(),
				F:               nonBootstraps.DefaultF(),
				IsPublic:        false,
				CandidateDigest: allConfigs.CandidateConfig.ConfigDigest,
				ActiveDigest:    allConfigs.ActiveConfig.ConfigDigest,
			})
		}
	}

	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.PromoteCandidateSequence,
		sequences.DONSequenceDeps{
			HomeChain: e.BlockChains.EVMChains()[cfg.HomeChainSelector],
		},
		sequences.PromoteCandidateSequenceInput{
			CapabilitiesRegistry: capReg.Address(),
			DONs:                 dons,
		},
	)
	return deployment.ChangesetOutput{}, err
}
