package deploy

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	ccipocr3 "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/internal"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"

	mcms_types "github.com/smartcontractkit/mcms/types"
)

// OCR3 is specifically a 1.6.0 feature
var OCR3Version = *semver.MustParse("1.6.0")

type SetOCR3ConfigInput struct {
	ChainSelector uint64
	Datastore     datastore.DataStore
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

func SetOCR3Config(deployerReg *DeployerRegistry, mcmsReg *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[SetOCR3ConfigArgs] {
	return cldf.CreateChangeSet(setOCR3ConfigApply(deployerReg, mcmsReg), setOCR3ConfigVerify(deployerReg, mcmsReg))
}

func setOCR3ConfigVerify(_ *DeployerRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, SetOCR3ConfigArgs) error {
	return func(e cldf.Environment, cfg SetOCR3ConfigArgs) error {
		// TODO: implement
		return nil
	}
}

func ocr3ConfigArgs(e cldf.Environment, homeChainSelector uint64, chainSelector uint64, configType utils.ConfigType) ([]OCR3ConfigArgs, error) {
	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       &OCR3Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find CCIPHome address in datastore: %w", err)
	}

	ccipHome, err := ccip_home.NewCCIPHome(common.HexToAddress(ccipHomeAddr.Address), e.BlockChains.EVMChains()[homeChainSelector].Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate CCIPHome contract: %w", err)
	}

	crAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       &OCR3Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find CapabilitiesRegistry address in datastore: %w", err)
	}

	cr, err := capabilities_registry.NewCapabilitiesRegistry(common.HexToAddress(crAddr.Address), e.BlockChains.EVMChains()[homeChainSelector].Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate CapabilitiesRegistry contract: %w", err)
	}
	donID, err := internal.DonIDForChain(
		cr,
		ccipHome,
		chainSelector,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get DON ID: %w", err)
	}

	// Default to active config if not set
	if configType == "" {
		configType = utils.ConfigTypeActive
	}

	ocr3Args, err := internal.BuildSetOCR3ConfigArgs(
		donID,
		ccipHome,
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
	ConfigType      utils.ConfigType
	MCMS            mcms.Input
}

func setOCR3ConfigApply(d *DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetOCR3ConfigArgs) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg SetOCR3ConfigArgs) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		batchOps := make([]mcms_types.BatchOperation, 0)
		for _, selector := range cfg.RemoteChainSels {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			adapter, exists := d.GetDeployer(family, &OCR3Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, OCR3Version.String())
			}
			args, err := ocr3ConfigArgs(e, cfg.HomeChainSel, selector, cfg.ConfigType)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			configs := make(map[ccipocr3.PluginType]OCR3ConfigArgs, 2)
			for _, arg := range args {
				configs[arg.PluginType] = arg
			}
			ocrReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SetOCR3Config(), e.BlockChains,
				SetOCR3ConfigInput{
					ChainSelector: selector,
					Configs:       configs,
				})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy Contract on chain with selector %d: %w", selector, err)
			}
			reports = append(reports, ocrReport.ExecutionReports...)
			batchOps = append(batchOps, ocrReport.Output.BatchOps...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
