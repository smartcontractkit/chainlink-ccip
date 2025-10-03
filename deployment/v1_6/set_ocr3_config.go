package v1_6

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	ccipocr3 "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/globals"
	"github.com/smartcontractkit/chainlink-ccip/deployment/internal"
)

type SetOCR3ConfigInput struct {
	ChainSelector uint64
	Configs       map[ccipocr3.PluginType]OCR3ConfigArgs
}

type OCR3ConfigArgs struct {
	ConfigDigest                   [32]byte
	PluginType                     ccipocr3.PluginType
	F                              uint8
	IsSignatureVerificationEnabled bool
	Signers                        [][]byte
	Transmitters                   [][]byte
}

type SetOCR3Config struct{}

func (cs SetOCR3Config) VerifyPreconditions(env cldf.Environment, config SetOCR3ConfigArgs) error {
	// TODO: do we pass through to the chain Adapter?
	return nil
}

// NOTE: this should become the new standard function that returns generic OCR3ConfigArgs
func ocr3ConfigArgs(e cldf.Environment, homeChainSelector uint64, chainSelector uint64, configType globals.ConfigType) ([]OCR3ConfigArgs, error) {
	state, err := stateview.LoadOnchainState(e)
	if err != nil {
		return nil, err
	}

	donID, err := internal.DonIDForChain(
		state.Chains[homeChainSelector].CapabilityRegistry,
		state.Chains[homeChainSelector].CCIPHome,
		chainSelector,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get DON ID: %w", err)
	}

	// Default to active config if not set
	if configType == "" {
		configType = globals.ConfigTypeActive
	}

	ocr3Args, err := internal.BuildSetOCR3ConfigArgs(
		donID,
		state.Chains[homeChainSelector].CCIPHome,
		chainSelector,
		configType,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build OCR3 config args: %w", err)
	}

	// MAP TO ARGS, this will become unnecessary once BuildSetOCR3Config uses correct types
	var args []OCR3ConfigArgs
	for _, arg := range ocr3Args {
		args = append(args, OCR3ConfigArgs{
			ConfigDigest:                   arg.ConfigDigest,
			PluginType:                     ccipocr3.PluginType(arg.OcrPluginType),
			F:                              arg.F,
			IsSignatureVerificationEnabled: arg.IsSignatureVerificationEnabled,
			Signers:                        arg.Signers,
			Transmitters:                   arg.Transmitters,
		})
	}
	return args, nil
}

type SetOCR3ConfigArgs struct {
	HomeChainSel    uint64
	RemoteChainSels []uint64
	ConfigType      globals.ConfigType
}

func (cs SetOCR3Config) Apply(env cldf.Environment, config SetOCR3ConfigArgs) (cldf.ChangesetOutput, error) {
	finalOutput := cldf.ChangesetOutput{}
	for _, chainSel := range config.RemoteChainSels {
		family, err := chain_selectors.GetSelectorFamily(chainSel)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		adapter, exists := registeredChainAdapters[family]
		if !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", family)
		}
		args, err := ocr3ConfigArgs(env, config.HomeChainSel, chainSel, config.ConfigType)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		configs := make(map[ccipocr3.PluginType]OCR3ConfigArgs, 2)
		for _, arg := range args {
			configs[arg.PluginType] = arg
		}
		output, err := adapter.SetOCR3Config(env, SetOCR3ConfigInput{
			ChainSelector: chainSel,
			Configs:       configs,
		})
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		err = MergeChangesetOutput(env, &finalOutput, output)
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset for chain selector %d: %w", chainSel, err)
		}
	}
	// TODO: aggregate timelock stuff
	return finalOutput, nil
}
