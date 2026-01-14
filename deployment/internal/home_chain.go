package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/bytes"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
)

var (
	CCIPHomeABI *abi.ABI
)

func init() {
	var err error
	CCIPHomeABI, err = ccip_home.CCIPHomeMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
}

const (
	CapabilityLabelledName = "ccip"
	CapabilityVersion      = "v1.0.0"
)

var CCIPCapabilityID = hashutil.NewKeccak().Hash(mustABIEncode(`[{"type": "string"}, {"type": "string"}]`, CapabilityLabelledName, CapabilityVersion))

// ABIEncode is the equivalent of abi.encode.
// See a full set of examples https://github.com/ethereum/go-ethereum/blob/420b78659bef661a83c5c442121b13f13288c09f/accounts/abi/packing_test.go#L31
func abiEncode(abiStr string, values ...any) ([]byte, error) {
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

func mustABIEncode(abiString string, args ...any) []byte {
	encoded, err := abiEncode(abiString, args...)
	if err != nil {
		panic(err)
	}
	return encoded
}

// DonIDForChain returns the DON ID for the chain with the given selector
// It looks up with the CCIPHome contract to find the OCR3 configs for the DONs, and returns the DON ID for the chain matching with the given selector from the OCR3 configs
func DonIDForChain(registry *capabilities_registry.CapabilitiesRegistry, ccipHome *ccip_home.CCIPHome, chainSelector uint64) (uint32, error) {
	dons, err := registry.GetDONs(nil)
	if err != nil {
		return 0, fmt.Errorf("get Dons from capability registry: %w", err)
	}
	var donIDs []uint32
	for _, don := range dons {
		if len(don.CapabilityConfigurations) == 1 &&
			don.CapabilityConfigurations[0].CapabilityId == CCIPCapabilityID {
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
		return 0, fmt.Errorf("more than one DON found for (chain selector %d, ccip capability id %x) pair", chainSelector, CCIPCapabilityID[:])
	}

	// no DON found - don ID of 0 indicates that (this is the case in the CR as well).
	if len(donIDs) == 0 {
		return 0, nil
	}

	// DON found - return it.
	return donIDs[0], nil
}

func validateOCR3Config(chainSel uint64, configForOCR3 ccip_home.CCIPHomeOCR3Config, chainConfig *ccip_home.CCIPHomeChainConfig) error {
	if chainConfig != nil {
		// chainConfigs must be set before OCR3 configs due to the added fChain == F validation
		if chainConfig.FChain == 0 || bytes.IsEmpty(chainConfig.Config) || len(chainConfig.Readers) == 0 {
			return fmt.Errorf("chain config is not set for chain selector %d", chainSel)
		}
		for _, reader := range chainConfig.Readers {
			if bytes.IsEmpty(reader[:]) {
				return fmt.Errorf("reader is empty, chain selector %d", chainSel)
			}
		}
		// FRoleDON >= fChain is a requirement
		if configForOCR3.FRoleDON < chainConfig.FChain {
			return fmt.Errorf("OCR3 config FRoleDON is lower than chainConfig FChain, chain %d", chainSel)
		}

		if len(configForOCR3.Nodes) < 3*int(chainConfig.FChain)+1 {
			return fmt.Errorf("number of nodes %d is less than 3 * fChain + 1 %d", len(configForOCR3.Nodes), 3*int(chainConfig.FChain)+1)
		}

		// check that we have enough transmitters for the destination chain.
		// note that this is done onchain, but we'll do it here for good measure to avoid reverts.
		// see https://github.com/smartcontractkit/chainlink-ccip/blob/8529b8c89093d0cd117b73645ea64b2d2a8092f4/chains/evm/contracts/capability/CCIPHome.sol#L511-L514.
		minTransmitterReq := 3*int(chainConfig.FChain) + 1
		var numNonzeroTransmitters int
		for _, node := range configForOCR3.Nodes {
			if len(node.TransmitterKey) > 0 {
				numNonzeroTransmitters++
			}
		}
		if numNonzeroTransmitters < minTransmitterReq {
			return fmt.Errorf("number of transmitters (%d) is less than 3 * fChain + 1 (%d), chain selector %d",
				numNonzeroTransmitters, minTransmitterReq, chainSel)
		}
	}

	// check if there is any zero byte address
	// The reason for this is that the MultiOCR3Base disallows zero addresses and duplicates
	if bytes.IsEmpty(configForOCR3.OfframpAddress) {
		return fmt.Errorf("zero address found in offramp address,  chain %d", chainSel)
	}
	if bytes.IsEmpty(configForOCR3.RmnHomeAddress) {
		return fmt.Errorf("zero address found in rmn home address,  chain %d", chainSel)
	}
	mapSignerKey := make(map[string]struct{})
	mapTransmitterKey := make(map[string]struct{})
	for _, node := range configForOCR3.Nodes {
		if bytes.IsEmpty(node.SignerKey) {
			return fmt.Errorf("zero address found in signer key, chain %d", chainSel)
		}

		// NOTE: We don't check for empty/zero transmitter address because the node can have a zero transmitter address if it does not support the destination chain.

		if bytes.IsEmpty(node.P2pId[:]) {
			return fmt.Errorf("empty p2p id, chain %d", chainSel)
		}

		// Signer and non-zero transmitter duplication must be checked
		if _, ok := mapSignerKey[hexutil.Encode(node.SignerKey)]; ok {
			return fmt.Errorf("duplicate signer key found, chain %d", chainSel)
		}

		// If len(node.TransmitterKey) == 0, the node does not support the destination chain, and we can definitely
		// have more than one node not supporting the destination chain, so the duplicate check doesn't make sense
		// for those.
		if _, ok := mapTransmitterKey[hexutil.Encode(node.TransmitterKey)]; ok && len(node.TransmitterKey) != 0 {
			return fmt.Errorf("duplicate transmitter key found, chain %d", chainSel)
		}
		mapSignerKey[hexutil.Encode(node.SignerKey)] = struct{}{}
		mapTransmitterKey[hexutil.Encode(node.TransmitterKey)] = struct{}{}
	}
	return nil
}

// we can't use the EVM one because we need the 32 byte transmitter address
type MultiOCR3BaseOCRConfigArgs struct {
	ConfigDigest                   [32]byte
	OcrPluginType                  uint8
	F                              uint8
	IsSignatureVerificationEnabled bool
	Signers                        [][]byte
	Transmitters                   [][]byte
}

// BuildSetOCR3ConfigArgs builds the OCR3 config arguments for the OffRamp contract
// using the donID's OCR3 configs from the CCIPHome contract.
func BuildSetOCR3ConfigArgs(
	donID uint32,
	ccipHome *ccip_home.CCIPHome,
	destSelector uint64,
	configType utils.ConfigType,
) ([]MultiOCR3BaseOCRConfigArgs, error) {
	chainCfg, err := ccipHome.GetChainConfig(nil, destSelector)
	if err != nil {
		return nil, fmt.Errorf("error getting chain config for chain selector %d it must be set before OCR3Config set up: %w", destSelector, err)
	}
	var offrampOCR3Configs []MultiOCR3BaseOCRConfigArgs
	for _, pluginType := range []ccipocr3.PluginType{ccipocr3.PluginTypeCCIPCommit, ccipocr3.PluginTypeCCIPExec} {
		ocrConfig, err2 := ccipHome.GetAllConfigs(&bind.CallOpts{
			Context: context.Background(),
		}, donID, uint8(pluginType))
		if err2 != nil {
			return nil, err2
		}

		configForOCR3 := ocrConfig.ActiveConfig
		// we expect only an active config
		switch configType {
		case utils.ConfigTypeActive:
			if ocrConfig.ActiveConfig.ConfigDigest == [32]byte{} {
				return nil, fmt.Errorf("invalid OCR3 config state, expected active config, donID: %d, activeConfig: %v, candidateConfig: %v",
					donID, hexutil.Encode(ocrConfig.ActiveConfig.ConfigDigest[:]), hexutil.Encode(ocrConfig.CandidateConfig.ConfigDigest[:]))
			}
		case utils.ConfigTypeCandidate:
			if ocrConfig.CandidateConfig.ConfigDigest == [32]byte{} {
				return nil, fmt.Errorf("invalid OCR3 config state, expected candidate config, donID: %d, activeConfig: %v, candidateConfig: %v",
					donID, hexutil.Encode(ocrConfig.ActiveConfig.ConfigDigest[:]), hexutil.Encode(ocrConfig.CandidateConfig.ConfigDigest[:]))
			}
			configForOCR3 = ocrConfig.CandidateConfig
		}

		if err := validateOCR3Config(destSelector, configForOCR3.Config, &chainCfg); err != nil {
			return nil, err
		}

		var signerAddresses [][]byte
		var transmitterAddresses [][]byte
		for _, node := range configForOCR3.Config.Nodes {
			signerAddresses = append(signerAddresses, node.SignerKey)
			transmitterAddresses = append(transmitterAddresses, node.TransmitterKey)
		}

		offrampOCR3Configs = append(offrampOCR3Configs, MultiOCR3BaseOCRConfigArgs{
			ConfigDigest:                   configForOCR3.ConfigDigest,
			OcrPluginType:                  uint8(pluginType),
			F:                              configForOCR3.Config.FRoleDON,
			IsSignatureVerificationEnabled: pluginType == ccipocr3.PluginTypeCCIPCommit,
			Signers:                        signerAddresses,
			Transmitters:                   transmitterAddresses,
		})
	}
	return offrampOCR3Configs, nil
}
