package v1_6

import (
	ccipocr3 "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
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
	// NOTE: this is a workaround and it only validates RemoteChainSelectors
	return ops.SetOCR3Config{}.VerifyPreconditions(env, ops.SetOCR3OffRampConfig{RemoteChainSels: config.RemoteChainSels})
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

	ocr3Args, err := internal.BuildSetOCR3ConfigArgsAptos(
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
	// TODO: loop over tonChains
	args, err := ocr3ConfigArgs(env, config.HomeChainSel, config.RemoteChainSels[0], config.ConfigType)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}

	configs := make(map[ccipocr3.PluginType]OCR3ConfigArgs, 2)
	for _, arg := range args {
		configs[arg.PluginType] = arg
	}

	// TODO: don't only wrap TON
	return ops.SetOCR3Config{}.Apply(env, ops.SetOCR3OffRampConfig{
		// TODO: map[remoteChainSels => configs]
		RemoteChainSels: config.RemoteChainSels,
		Configs:         configs,
	})
}
