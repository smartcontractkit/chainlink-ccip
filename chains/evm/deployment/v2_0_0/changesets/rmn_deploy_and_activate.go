package changesets

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	rmn_proxy_bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	rmn15_bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	rmn_bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_1_0/rmn"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// ActivateRMNCfg configures the ActivateRMN changeset.
type ActivateRMNCfg struct {
	ChainSels []uint64
	// CurseAdmins are optional additional authorized callers (cursers) added at RMN deploy
	// time, keyed by chain selector. The Ultra Fast Curse RBACTimelock is always included.
	CurseAdmins map[uint64][]common.Address
}

var ActivateRMN = func(mcmsRegistry *changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[ActivateRMNCfg]] {
	return cldf_deployment.CreateChangeSet(
		func(e cldf_deployment.Environment, input changesets.WithMCMS[ActivateRMNCfg]) (cldf_deployment.ChangesetOutput, error) {
			return applyDeployAndActivateRMN(e, mcmsRegistry, input)
		},
		validateActivateRMN,
	)
}

func validateActivateRMN(e cldf_deployment.Environment, input changesets.WithMCMS[ActivateRMNCfg]) error {
	if len(input.Cfg.ChainSels) == 0 {
		return fmt.Errorf("at least one chain selector is required")
	}
	if err := validateActivateRMNCurseAdmins(input.Cfg.CurseAdmins, input.Cfg.ChainSels); err != nil {
		return err
	}
	evmChains := e.BlockChains.EVMChains()
	for _, sel := range input.Cfg.ChainSels {
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))
		if err := validateActivateRMNAddresses(addresses, sel); err != nil {
			return err
		}
	}
	return nil
}

func validateActivateRMNCurseAdmins(curseAdmins map[uint64][]common.Address, chainSels []uint64) error {
	if len(curseAdmins) == 0 {
		return nil
	}
	chainSet := make(map[uint64]struct{}, len(chainSels))
	for _, sel := range chainSels {
		chainSet[sel] = struct{}{}
	}
	for sel, addrs := range curseAdmins {
		if _, ok := chainSet[sel]; !ok {
			return fmt.Errorf("curse admins configured for chain %d which is not in ChainSels", sel)
		}
		seen := make(map[common.Address]struct{}, len(addrs))
		for _, addr := range addrs {
			if addr == (common.Address{}) {
				return fmt.Errorf("curse admin address cannot be zero for chain %d", sel)
			}
			if _, dup := seen[addr]; dup {
				return fmt.Errorf("duplicate curse admin %s for chain %d", addr.Hex(), sel)
			}
			seen[addr] = struct{}{}
		}
	}
	return nil
}

func applyDeployAndActivateRMN(
	e cldf_deployment.Environment,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	input changesets.WithMCMS[ActivateRMNCfg],
) (cldf_deployment.ChangesetOutput, error) {
	evmChains := e.BlockChains.EVMChains()
	outputDS := datastore.NewMemoryDataStore()
	rmnBatchOps := make([]mcms_types.BatchOperation, 0, len(input.Cfg.ChainSels))
	cllBatchOps := make([]mcms_types.BatchOperation, 0, len(input.Cfg.ChainSels))

	for _, sel := range input.Cfg.ChainSels {
		chain := evmChains[sel]
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))

		cursedSubjects, err := getCursedSubjectsFromCurrentRMN(e, chain.Client, sel, addresses)
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, err
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployAndActivateRMN, chain, sequences.ActivateRMNInput{
			ChainSelector:     sel,
			ExistingAddresses: addresses,
			CurseAdmins:       input.Cfg.CurseAdmins[sel],
			SubjectsToMigrate: cursedSubjects,
		})
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to activate RMN on chain %d: %w", sel, err)
		}

		for _, ref := range report.Output.Addresses {
			if addErr := outputDS.Addresses().Add(ref); addErr != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add address to datastore: %w", addErr)
			}
		}
		rmnBatchOps = append(rmnBatchOps, report.Output.RMNMCMSBatchOps...)
		cllBatchOps = append(cllBatchOps, report.Output.CLLCCIPBatchOps...)
	}

	var output cldf_deployment.ChangesetOutput
	output.DataStore = outputDS

	if err := input.MCMS.PopulateDefaults(); err != nil {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to populate MCMS defaults: %w", err)
	}
	if len(rmnBatchOps) > 0 {
		rmnOut, err := changesets.NewOutputBuilder(e, mcmsRegistry).
			WithDataStore(outputDS).
			WithBatchOps(rmnBatchOps).
			Build(mcmsInputForActivateRMN(
				input.MCMS,
				common_utils.RMNTimelockQualifier,
				"Accept RMN 2.0 ownership on RMNMCMS timelock",
			))
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to build RMNMCMS proposal: %w", err)
		}
		output.MCMSTimelockProposals = append(output.MCMSTimelockProposals, rmnOut.MCMSTimelockProposals...)
	}

	// CLLCCIPBatchOps is only populated when SetRMN could not run on-chain (proxy not deployer-owned).
	if len(cllBatchOps) > 0 {
		cllChainSels := make([]uint64, 0, len(cllBatchOps))
		for _, op := range cllBatchOps {
			cllChainSels = append(cllChainSels, uint64(op.ChainSelector))
		}
		if err := validateCLLCCIPForProxyProposal(e, cllChainSels); err != nil {
			return cldf_deployment.ChangesetOutput{}, err
		}
		cllOut, err := changesets.NewOutputBuilder(e, mcmsRegistry).
			WithDataStore(outputDS).
			WithBatchOps(cllBatchOps).
			Build(mcmsInputForActivateRMN(
				input.MCMS,
				common_utils.CLLQualifier,
				"Point RMNProxy at RMN 2.0 on CLLCCIP timelock",
			))
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to build CLLCCIP proposal: %w", err)
		}
		output.MCMSTimelockProposals = append(output.MCMSTimelockProposals, cllOut.MCMSTimelockProposals...)
	}

	return output, nil
}

func mcmsInputForActivateRMN(
	base mcms.Input,
	qualifier, description string,
) mcms.Input {
	return mcms.Input{
		OverridePreviousRoot: base.OverridePreviousRoot,
		ValidUntil:           base.ValidUntil,
		TimelockAction:       mcms_types.TimelockActionSchedule,
		Qualifier:            qualifier,
		Description:          description,
	}
}

func validateActivateRMNAddresses(addresses []datastore.AddressRef, chainSelector uint64) error {
	if ref := datastore_utils.GetAddressRef(
		addresses,
		chainSelector,
		common_utils.RBACTimelock,
		mcms_ops.MCMSVersion,
		common_utils.RMNTimelockQualifier,
	); ref.Address == "" {
		return fmt.Errorf(
			"ownership transfer requires RMNMCMS RBACTimelock (qualifier %q) in datastore for chain %d",
			common_utils.RMNTimelockQualifier, chainSelector,
		)
	}
	return nil
}

func getCursedSubjectsFromCurrentRMN(
	e cldf_deployment.Environment,
	client bind.ContractBackend,
	chainSelector uint64,
	addresses []datastore.AddressRef,
) ([][16]byte, error) {
	proxyRef := datastore_utils.GetAddressRef(addresses, chainSelector, rmnproxyops.ContractType, rmnproxyops.Version, "")
	if proxyRef.Address == "" {
		return nil, nil
	}

	ctx := context.Background()
	if e.GetContext != nil {
		ctx = e.GetContext()
	}

	proxyC, err := rmn_proxy_bind.NewRMNProxy(common.HexToAddress(proxyRef.Address), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate RMNProxy on chain %d: %w", chainSelector, err)
	}

	currentRMNAddr, err := proxyC.GetARM(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to read active RMN from RMNProxy on chain %d: %w", chainSelector, err)
	}
	if currentRMNAddr == (common.Address{}) {
		return nil, nil
	}

	code, err := client.CodeAt(ctx, currentRMNAddr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect active RMN code on chain %d: %w", chainSelector, err)
	}
	if len(code) == 0 {
		return nil, nil
	}

	currentRMNC, err := rmn_bind.NewRMN(currentRMNAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate current RMN on chain %d: %w", chainSelector, err)
	}

	cursedSubjects, err := currentRMNC.GetCursedSubjects(&bind.CallOpts{Context: ctx})
	if err == nil {
		return cursedSubjects, nil
	}

	legacyCursedSubjects, legacyErr := getCursedSubjectsFromRMN15(ctx, client, currentRMNAddr)
	if legacyErr == nil {
		return legacyCursedSubjects, nil
	}

	return nil, fmt.Errorf(
		"failed to read cursed subjects from current RMN on chain %d (v2/v1.6 read: %w, v1.5 fallback: %v)",
		chainSelector,
		err,
		legacyErr,
	)
}

func getCursedSubjectsFromRMN15(
	ctx context.Context,
	client bind.ContractBackend,
	rmnAddr common.Address,
) ([][16]byte, error) {
	rmn15C, err := rmn15_bind.NewRMNContract(rmnAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate RMN 1.5 contract: %w", err)
	}

	count, err := rmn15C.GetCursedSubjectsCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to read RMN 1.5 cursed subject count: %w", err)
	}
	if count.Sign() == 0 {
		return nil, nil
	}

	const pageSize int64 = 64
	offset := big.NewInt(0)
	limit := big.NewInt(pageSize)
	knownSubjects := make(map[[16]byte]struct{})

	for {
		recordedOps, readErr := rmn15C.GetRecordedCurseRelatedOps(&bind.CallOpts{Context: ctx}, offset, limit)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read RMN 1.5 curse operation history: %w", readErr)
		}
		if len(recordedOps) == 0 {
			break
		}
		for _, op := range recordedOps {
			knownSubjects[op.Subject] = struct{}{}
		}
		offset.Add(offset, big.NewInt(int64(len(recordedOps))))
		if len(recordedOps) < int(pageSize) {
			break
		}
	}

	activeCurses := make([][16]byte, 0, count.Int64())
	for subject := range knownSubjects {
		isCursed, cursedErr := rmn15C.IsCursed(&bind.CallOpts{Context: ctx}, subject)
		if cursedErr != nil {
			return nil, fmt.Errorf("failed to check RMN 1.5 curse state for subject %x: %w", subject, cursedErr)
		}
		if isCursed {
			activeCurses = append(activeCurses, subject)
		}
	}

	return activeCurses, nil
}

func validateCLLCCIPForProxyProposal(e cldf_deployment.Environment, chainSels []uint64) error {
	for _, sel := range chainSels {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))
		if ref := datastore_utils.GetAddressRef(
			addresses,
			sel,
			common_utils.RBACTimelock,
			mcms_ops.MCMSVersion,
			common_utils.CLLQualifier,
		); ref.Address == "" {
			return fmt.Errorf(
				"RMNProxy.SetRMN requires CLLCCIP RBACTimelock (qualifier %q) in datastore for chain %d",
				common_utils.CLLQualifier, sel,
			)
		}
	}
	return nil
}
