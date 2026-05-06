package deploy

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	ccipocr3 "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/internal"
	deploycore "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"

	mcms_types "github.com/smartcontractkit/mcms/types"
)

func SetOCR3Config(deployerReg *deploycore.DeployerRegistry, mcmsReg *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[deploycore.SetOCR3ConfigArgs] {
	return cldf.CreateChangeSet(setOCR3ConfigApply(deployerReg, mcmsReg), setOCR3ConfigVerify(deployerReg, mcmsReg))
}

func setOCR3ConfigVerify(_ *deploycore.DeployerRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, deploycore.SetOCR3ConfigArgs) error {
	return func(e cldf.Environment, cfg deploycore.SetOCR3ConfigArgs) error {
		// TODO: implement
		return nil
	}
}

func ocr3ConfigArgs(e cldf.Environment, homeChainSelector uint64, chainSelector uint64, configType utils.ConfigType) ([]deploycore.OCR3ConfigArgs, error) {
	ccipHomeAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: homeChainSelector,
		Type:          datastore.ContractType(utils.CCIPHome),
		Version:       &deploycore.OCR3Version,
	}, homeChainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find CCIPHome address in datastore: %w", err)
	}

	ccipHome, err := ccip_home.NewCCIPHome(common.HexToAddress(ccipHomeAddr.Address), e.BlockChains.EVMChains()[homeChainSelector].Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate CCIPHome contract: %w", err)
	}

	crAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: homeChainSelector,
		Type:          datastore.ContractType(utils.CapabilitiesRegistry),
		Version:       &deploycore.CapabilitiesRegistryVersion,
	}, homeChainSelector, datastore_utils.FullRef)
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

	var args []deploycore.OCR3ConfigArgs
	for _, arg := range ocr3Args {
		args = append(args, deploycore.OCR3ConfigArgs{
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

func setOCR3ConfigApply(d *deploycore.DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, deploycore.SetOCR3ConfigArgs) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg deploycore.SetOCR3ConfigArgs) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		batchOps := make([]mcms_types.BatchOperation, 0)
		for _, selector := range cfg.RemoteChainSels {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			adapter, exists := d.GetDeployer(family, &deploycore.OCR3Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, deploycore.OCR3Version.String())
			}
			args, err := ocr3ConfigArgs(e, cfg.HomeChainSel, selector, cfg.ConfigType)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			configs := make(map[ccipocr3.PluginType]deploycore.OCR3ConfigArgs, 2)
			for _, arg := range args {
				configs[arg.PluginType] = arg
			}
			ocrReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SetOCR3Config(), e.BlockChains,
				deploycore.SetOCR3ConfigInput{
					ChainSelector: selector,
					Datastore:     e.DataStore,
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
