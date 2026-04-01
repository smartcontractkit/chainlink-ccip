package changesets

import (
	"bytes"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"

	burn_mint_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/burn_mint_with_external_minter_token_pool"
	hybrid_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"
	evm_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_1_lock_release_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/lock_release_token_pool"
	v1_6_0_burn_mint_with_external_minter_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/burn_mint_with_external_minter_token_pool"
	v1_6_0_hybrid_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool"
	v1_6_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	v1_5_1_lock_release_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var supportedOldPoolTypeAndVersions = map[string]struct{}{
	v1_5_1_lock_release_token_pool_ops.TypeAndVersion.String(): {},
}

var expectedNewPoolTypeAndVersion = v1_6_0_burn_mint_with_external_minter_token_pool_ops.TypeAndVersion.String()

type MigrateHybridPoolRemoteConfig struct {
	HubChainSelector     uint64         `json:"hubChainSelector" yaml:"hubChainSelector"`
	HubPoolAddress       common.Address `json:"hubPoolAddress" yaml:"hubPoolAddress"`
	RemoteChainSelector  uint64         `json:"remoteChainSelector" yaml:"remoteChainSelector"`
	NewRemotePoolAddress common.Address `json:"newRemotePoolAddress" yaml:"newRemotePoolAddress"`
	OldRemotePoolAddress common.Address `json:"oldRemotePoolAddress" yaml:"oldRemotePoolAddress"`
	TargetGroup          uint8          `json:"targetGroup" yaml:"targetGroup"`
	RemoteTokenAddress   common.Address `json:"remoteTokenAddress" yaml:"remoteTokenAddress"`
	MCMS                 mcms.Input     `json:"mcms,omitempty" yaml:"mcms,omitempty"`
}

func MigrateHybridPoolRemote(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MigrateHybridPoolRemoteConfig] {
	return cldf.CreateChangeSet(
		makeApplyMigrateHybridPoolRemote(mcmsRegistry),
		makeVerifyMigrateHybridPoolRemote(mcmsRegistry),
	)
}

func makeVerifyMigrateHybridPoolRemote(
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, MigrateHybridPoolRemoteConfig) error {
	return func(e cldf.Environment, cfg MigrateHybridPoolRemoteConfig) error {
		if err := cfg.MCMS.Validate(); err != nil {
			return fmt.Errorf("invalid MCMS config: %w", err)
		}
		if cfg.HubChainSelector == cfg.RemoteChainSelector {
			return fmt.Errorf("hub chain selector and remote chain selector must be different")
		}
		if cfg.TargetGroup != 0 && cfg.TargetGroup != 1 {
			return fmt.Errorf("target group must be 0 or 1, got %d", cfg.TargetGroup)
		}
		if cfg.HubPoolAddress == (common.Address{}) {
			return fmt.Errorf("hub pool address cannot be the zero address")
		}
		if cfg.NewRemotePoolAddress == (common.Address{}) {
			return fmt.Errorf("new remote pool address cannot be the zero address")
		}
		if cfg.OldRemotePoolAddress == (common.Address{}) {
			return fmt.Errorf("old remote pool address cannot be the zero address")
		}
		if cfg.OldRemotePoolAddress == cfg.NewRemotePoolAddress {
			return fmt.Errorf("old remote pool address and new remote pool address must be different")
		}
		if cfg.RemoteTokenAddress == (common.Address{}) {
			return fmt.Errorf("remote token address cannot be the zero address")
		}

		hubFamily, err := chain_selectors.GetSelectorFamily(cfg.HubChainSelector)
		if err != nil {
			return fmt.Errorf("invalid hub chain selector %d: %w", cfg.HubChainSelector, err)
		}
		if hubFamily != chain_selectors.FamilyEVM {
			return fmt.Errorf("hub chain selector %d is not an EVM chain", cfg.HubChainSelector)
		}

		remoteFamily, err := chain_selectors.GetSelectorFamily(cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("invalid remote chain selector %d: %w", cfg.RemoteChainSelector, err)
		}
		if remoteFamily != chain_selectors.FamilyEVM {
			return fmt.Errorf("remote chain selector %d is not an EVM chain", cfg.RemoteChainSelector)
		}

		if !e.BlockChains.Exists(cfg.HubChainSelector) {
			return fmt.Errorf("chain with selector %d does not exist", cfg.HubChainSelector)
		}
		if !e.BlockChains.Exists(cfg.RemoteChainSelector) {
			return fmt.Errorf("chain with selector %d does not exist", cfg.RemoteChainSelector)
		}

		hubChain, ok := e.BlockChains.EVMChains()[cfg.HubChainSelector]
		if !ok {
			return fmt.Errorf("hub chain selector %d is not configured as an EVM chain", cfg.HubChainSelector)
		}
		remoteChain, ok := e.BlockChains.EVMChains()[cfg.RemoteChainSelector]
		if !ok {
			return fmt.Errorf("remote chain selector %d is not configured as an EVM chain", cfg.RemoteChainSelector)
		}

		remoteTARAddress, err := resolveRemoteTARAddress(e.DataStore, cfg.RemoteChainSelector)
		if err != nil {
			return err
		}

		if mcmsRegistry == nil {
			return fmt.Errorf("no MCMS reader registry configured")
		}
		mcmsReader, ok := mcmsRegistry.GetMCMSReader(chain_selectors.FamilyEVM)
		if !ok {
			return fmt.Errorf("no MCMS reader registered for chain family '%s'", chain_selectors.FamilyEVM)
		}

		hubTimelockRef, err := mcmsReader.GetTimelockRef(e, cfg.HubChainSelector, cfg.MCMS)
		if err != nil {
			return fmt.Errorf("failed to resolve timelock for hub chain %d with qualifier %s: %w", cfg.HubChainSelector, cfg.MCMS.Qualifier, err)
		}
		if hubTimelockRef.Address == "" {
			return fmt.Errorf("missing timelock for hub chain %d with qualifier %s", cfg.HubChainSelector, cfg.MCMS.Qualifier)
		}
		if !common.IsHexAddress(hubTimelockRef.Address) {
			return fmt.Errorf("invalid timelock address for hub chain %d with qualifier %s: %q", cfg.HubChainSelector, cfg.MCMS.Qualifier, hubTimelockRef.Address)
		}
		hubMCMSRef, err := mcmsReader.GetMCMSRef(e, cfg.HubChainSelector, cfg.MCMS)
		if err != nil {
			return fmt.Errorf("failed to resolve MCMS for hub chain %d with qualifier %s: %w", cfg.HubChainSelector, cfg.MCMS.Qualifier, err)
		}
		if hubMCMSRef.Address == "" {
			return fmt.Errorf("missing MCMS for hub chain %d with qualifier %s", cfg.HubChainSelector, cfg.MCMS.Qualifier)
		}
		if !common.IsHexAddress(hubMCMSRef.Address) {
			return fmt.Errorf("invalid MCMS address for hub chain %d with qualifier %s: %q", cfg.HubChainSelector, cfg.MCMS.Qualifier, hubMCMSRef.Address)
		}

		remoteTimelockRef, err := mcmsReader.GetTimelockRef(e, cfg.RemoteChainSelector, cfg.MCMS)
		if err != nil {
			return fmt.Errorf("failed to resolve timelock for remote chain %d with qualifier %s: %w", cfg.RemoteChainSelector, cfg.MCMS.Qualifier, err)
		}
		if remoteTimelockRef.Address == "" {
			return fmt.Errorf("missing timelock for remote chain %d with qualifier %s", cfg.RemoteChainSelector, cfg.MCMS.Qualifier)
		}
		if !common.IsHexAddress(remoteTimelockRef.Address) {
			return fmt.Errorf("invalid timelock address for remote chain %d with qualifier %s: %q", cfg.RemoteChainSelector, cfg.MCMS.Qualifier, remoteTimelockRef.Address)
		}
		remoteMCMSRef, err := mcmsReader.GetMCMSRef(e, cfg.RemoteChainSelector, cfg.MCMS)
		if err != nil {
			return fmt.Errorf("failed to resolve MCMS for remote chain %d with qualifier %s: %w", cfg.RemoteChainSelector, cfg.MCMS.Qualifier, err)
		}
		if remoteMCMSRef.Address == "" {
			return fmt.Errorf("missing MCMS for remote chain %d with qualifier %s", cfg.RemoteChainSelector, cfg.MCMS.Qualifier)
		}
		if !common.IsHexAddress(remoteMCMSRef.Address) {
			return fmt.Errorf("invalid MCMS address for remote chain %d with qualifier %s: %q", cfg.RemoteChainSelector, cfg.MCMS.Qualifier, remoteMCMSRef.Address)
		}

		if err := verifyTypeAndVersion(
			e.DataStore, cfg.HubChainSelector, cfg.HubPoolAddress,
			cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType), v1_6_0_hybrid_pool_ops.Version,
			hubChain.Client, "hub pool",
		); err != nil {
			return err
		}
		if err := verifyTypeAndVersion(
			e.DataStore, cfg.RemoteChainSelector, cfg.OldRemotePoolAddress,
			cldf_datastore.ContractType(v1_5_1_lock_release_token_pool_ops.ContractType), v1_5_1_lock_release_token_pool_ops.Version,
			remoteChain.Client, "old remote pool",
		); err != nil {
			return err
		}
		if err := verifyTypeAndVersion(
			e.DataStore, cfg.RemoteChainSelector, cfg.NewRemotePoolAddress,
			cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType), v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
			remoteChain.Client, "new remote pool",
		); err != nil {
			return err
		}

		oldPool, err := v1_5_1_lock_release_token_pool_bindings.NewLockReleaseTokenPool(cfg.OldRemotePoolAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind old remote pool %s on chain %d: %w", cfg.OldRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		oldPoolToken, err := oldPool.GetToken(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read token from old remote pool %s on chain %d: %w", cfg.OldRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		if oldPoolToken != cfg.RemoteTokenAddress {
			return fmt.Errorf("old remote pool token %s does not match remote token %s", oldPoolToken, cfg.RemoteTokenAddress)
		}

		newPool, err := burn_mint_with_external_minter_token_pool_bindings.NewBurnMintWithExternalMinterTokenPool(cfg.NewRemotePoolAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind new remote pool %s on chain %d: %w", cfg.NewRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		newPoolToken, err := newPool.GetToken(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read token from new remote pool %s on chain %d: %w", cfg.NewRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		if newPoolToken != cfg.RemoteTokenAddress {
			return fmt.Errorf("new remote pool token %s does not match remote token %s", newPoolToken, cfg.RemoteTokenAddress)
		}

		hubPool, err := hybrid_with_external_minter_token_pool_bindings.NewHybridWithExternalMinterTokenPool(cfg.HubPoolAddress, hubChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind hub pool %s on chain %d: %w", cfg.HubPoolAddress, cfg.HubChainSelector, err)
		}
		isSupportedChain, err := hubPool.IsSupportedChain(&bind.CallOpts{Context: e.GetContext()}, cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("failed to read supported-chain status for remote chain %d from hub pool %s on chain %d: %w", cfg.RemoteChainSelector, cfg.HubPoolAddress, cfg.HubChainSelector, err)
		}
		if !isSupportedChain {
			return fmt.Errorf("remote chain %d is not supported on hub pool %s on chain %d", cfg.RemoteChainSelector, cfg.HubPoolAddress, cfg.HubChainSelector)
		}

		hubRemotePools, err := hubPool.GetRemotePools(&bind.CallOpts{Context: e.GetContext()}, cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("failed to read hub remote pools for remote chain %d from hub pool %s on chain %d: %w", cfg.RemoteChainSelector, cfg.HubPoolAddress, cfg.HubChainSelector, err)
		}
		oldPoolBytes := common.LeftPadBytes(cfg.OldRemotePoolAddress.Bytes(), 32)
		newPoolBytes := common.LeftPadBytes(cfg.NewRemotePoolAddress.Bytes(), 32)
		oldPoolPresent := false
		newPoolPresent := false
		for _, remotePool := range hubRemotePools {
			switch {
			case bytes.Equal(remotePool, oldPoolBytes):
				oldPoolPresent = true
			case bytes.Equal(remotePool, newPoolBytes):
				newPoolPresent = true
			default:
				return fmt.Errorf("unexpected pool %x in hub remote pool set for chain %d", remotePool, cfg.RemoteChainSelector)
			}
		}
		if !oldPoolPresent && !newPoolPresent {
			return fmt.Errorf("neither old pool %s nor new pool %s registered in hub remote pool set for chain %d", cfg.OldRemotePoolAddress, cfg.NewRemotePoolAddress, cfg.RemoteChainSelector)
		}

		hubRemoteToken, err := hubPool.GetRemoteToken(&bind.CallOpts{Context: e.GetContext()}, cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("failed to read hub remote token bytes for remote chain %d from hub pool %s on chain %d: %w", cfg.RemoteChainSelector, cfg.HubPoolAddress, cfg.HubChainSelector, err)
		}
		if !bytesAddressMatches(hubRemoteToken, cfg.RemoteTokenAddress) {
			return fmt.Errorf(
				"hub remote token bytes %x do not match remote token %s for chain %d",
				hubRemoteToken,
				cfg.RemoteTokenAddress,
				cfg.RemoteChainSelector,
			)
		}

		hubPoolOwner, err := hubPool.Owner(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read owner for hub pool %s on chain %d: %w", cfg.HubPoolAddress, cfg.HubChainSelector, err)
		}
		hubTimelockAddress := common.HexToAddress(hubTimelockRef.Address)
		if hubPoolOwner != hubTimelockAddress {
			return fmt.Errorf(
				"hub pool %s owner %s does not match timelock %s on chain %d",
				cfg.HubPoolAddress,
				hubPoolOwner,
				hubTimelockAddress,
				cfg.HubChainSelector,
			)
		}

		tarConfigReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, tar_ops.GetTokenConfig, remoteChain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: cfg.RemoteChainSelector,
			Address:       remoteTARAddress,
			Args:          cfg.RemoteTokenAddress,
		})
		if err != nil {
			return fmt.Errorf("failed to read TAR token config for token %s on chain %d: %w", cfg.RemoteTokenAddress, cfg.RemoteChainSelector, err)
		}
		remoteTimelockAddress := common.HexToAddress(remoteTimelockRef.Address)
		if tarConfigReport.Output.Administrator != remoteTimelockAddress {
			return fmt.Errorf(
				"TAR administrator %s for token %s does not match timelock %s on chain %d",
				tarConfigReport.Output.Administrator,
				cfg.RemoteTokenAddress,
				remoteTimelockAddress,
				cfg.RemoteChainSelector,
			)
		}
		tarPool := tarConfigReport.Output.TokenPool
		if tarPool == (common.Address{}) {
			return fmt.Errorf("TAR has no pool set for token %s on chain %d", cfg.RemoteTokenAddress, cfg.RemoteChainSelector)
		}
		if tarPool != cfg.OldRemotePoolAddress && tarPool != cfg.NewRemotePoolAddress {
			return fmt.Errorf(
				"TAR pool %s for token %s on chain %d is neither old pool %s nor new pool %s",
				tarPool,
				cfg.RemoteTokenAddress,
				cfg.RemoteChainSelector,
				cfg.OldRemotePoolAddress,
				cfg.NewRemotePoolAddress,
			)
		}

		remoteToken, err := erc20.NewERC20(cfg.RemoteTokenAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind remote token %s on chain %d: %w", cfg.RemoteTokenAddress, cfg.RemoteChainSelector, err)
		}
		totalSupply, err := remoteToken.TotalSupply(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read totalSupply for remote token %s on chain %d: %w", cfg.RemoteTokenAddress, cfg.RemoteChainSelector, err)
		}

		if cfg.TargetGroup == 1 {
			lockedTokensReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, v1_6_0_hybrid_pool_ops.GetLockedTokens, hubChain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: cfg.HubChainSelector,
				Address:       cfg.HubPoolAddress,
				Args:          struct{}{},
			})
			if err != nil {
				return fmt.Errorf("failed to read locked token accounting from hub pool %s on chain %d: %w", cfg.HubPoolAddress, cfg.HubChainSelector, err)
			}
			if lockedTokensReport.Output == nil {
				return fmt.Errorf("hub pool %s returned nil locked token accounting on chain %d", cfg.HubPoolAddress, cfg.HubChainSelector)
			}
			if totalSupply.Cmp(lockedTokensReport.Output) > 0 {
				return fmt.Errorf(
					"remote token totalSupply %s exceeds locked token accounting %s for hub pool %s on chain %d",
					totalSupply.String(),
					lockedTokensReport.Output.String(),
					cfg.HubPoolAddress,
					cfg.HubChainSelector,
				)
			}
		}

		return nil
	}
}

func makeApplyMigrateHybridPoolRemote(
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, MigrateHybridPoolRemoteConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MigrateHybridPoolRemoteConfig) (cldf.ChangesetOutput, error) {
		remoteTARAddress, err := resolveRemoteTARAddress(e.DataStore, cfg.RemoteChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		remoteChain, ok := e.BlockChains.EVMChains()[cfg.RemoteChainSelector]
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("remote chain selector %d is not configured as an EVM chain", cfg.RemoteChainSelector)
		}

		remoteToken, err := erc20.NewERC20(cfg.RemoteTokenAddress, remoteChain.Client)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to bind remote token %s on chain %d: %w", cfg.RemoteTokenAddress, cfg.RemoteChainSelector, err)
		}
		totalSupply, err := remoteToken.TotalSupply(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to read totalSupply for remote token %s on chain %d: %w", cfg.RemoteTokenAddress, cfg.RemoteChainSelector, err)
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, v1_6_0_sequences.MigrateHybridPoolRemote, e.BlockChains, v1_6_0_sequences.MigrateHybridPoolRemoteInput{
			HubChainSelector:     cfg.HubChainSelector,
			HubPoolAddress:       cfg.HubPoolAddress,
			RemoteChainSelector:  cfg.RemoteChainSelector,
			NewRemotePoolAddress: cfg.NewRemotePoolAddress,
			OldRemotePoolAddress: cfg.OldRemotePoolAddress,
			RemoteChainSupply:    totalSupply,
			TargetGroup:          cfg.TargetGroup,
			RemoteTARAddress:     remoteTARAddress,
			RemoteTokenAddress:   cfg.RemoteTokenAddress,
		})
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to migrate token pool: %w", err)
		}

		ds := cldf_datastore.NewMemoryDataStore()
		if !AddressRefExistsWithTypeVersion(
			e.DataStore,
			cfg.HubChainSelector,
			cfg.HubPoolAddress,
			cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
			v1_6_0_hybrid_pool_ops.Version,
		) {
			if err := ds.Addresses().Add(cldf_datastore.AddressRef{
				ChainSelector: cfg.HubChainSelector,
				Type:          cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
				Version:       v1_6_0_hybrid_pool_ops.Version,
				Address:       cfg.HubPoolAddress.Hex(),
			}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to persist hub pool ref %s on chain %d: %w", cfg.HubPoolAddress, cfg.HubChainSelector, err)
			}
		}
		if !AddressRefExistsWithTypeVersion(
			e.DataStore,
			cfg.RemoteChainSelector,
			cfg.NewRemotePoolAddress,
			cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
			v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
		) {
			if err := ds.Addresses().Add(cldf_datastore.AddressRef{
				ChainSelector: cfg.RemoteChainSelector,
				Type:          cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
				Version:       v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
				Address:       cfg.NewRemotePoolAddress.Hex(),
			}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to persist new remote pool ref %s on chain %d: %w", cfg.NewRemotePoolAddress, cfg.RemoteChainSelector, err)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(report.ExecutionReports).
			WithBatchOps(report.Output.BatchOps).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}

func resolveRemoteTARAddress(ds cldf_datastore.DataStore, remoteChainSelector uint64) (common.Address, error) {
	tarRef := cldf_datastore.AddressRef{
		Type:          cldf_datastore.ContractType(tar_ops.ContractType),
		ChainSelector: remoteChainSelector,
		Version:       tar_ops.Version,
	}
	addr, err := datastore_utils.FindAndFormatRef(ds, tarRef, remoteChainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to resolve TokenAdminRegistry address on chain %d: %w", remoteChainSelector, err)
	}
	return addr, nil
}

func verifyTypeAndVersion(
	ds cldf_datastore.DataStore,
	chainSelector uint64,
	address common.Address,
	expectedType cldf_datastore.ContractType,
	expectedVersion *semver.Version,
	backend bind.ContractBackend,
	label string,
) error {
	refs := ds.Addresses().Filter(
		cldf_datastore.AddressRefByChainSelector(chainSelector),
		cldf_datastore.AddressRefByAddress(address.Hex()),
	)
	for _, ref := range refs {
		if ref.Type == expectedType && ref.Version != nil && ref.Version.Equal(expectedVersion) {
			return nil
		}
		if ref.Type != "" || ref.Version != nil {
			return fmt.Errorf(
				"%s %s on chain %d has datastore ref type=%s version=%s, expected type=%s version=%s",
				label, address, chainSelector,
				ref.Type, ref.Version,
				expectedType, expectedVersion,
			)
		}
	}
	if len(refs) == 0 {
		contractType, version, err := evm_utils.TypeAndVersion(address, backend)
		if err != nil {
			return fmt.Errorf("failed to read typeAndVersion for %s %s on chain %d: %w", label, address, chainSelector, err)
		}
		actual := fmt.Sprintf("%s %s", contractType, version.String())
		expected := fmt.Sprintf("%s %s", expectedType, expectedVersion.String())
		if actual != expected {
			return fmt.Errorf(
				"unexpected typeAndVersion %q for %s %s on chain %d, expected %q",
				actual, label, address, chainSelector, expected,
			)
		}
	}
	return nil
}

func AddressRefExistsWithTypeVersion(
	ds cldf_datastore.DataStore,
	chainSelector uint64,
	address common.Address,
	expectedType cldf_datastore.ContractType,
	expectedVersion *semver.Version,
) bool {
	refs := ds.Addresses().Filter(
		cldf_datastore.AddressRefByChainSelector(chainSelector),
		cldf_datastore.AddressRefByAddress(address.Hex()),
	)
	for _, ref := range refs {
		if ref.Type == expectedType && ref.Version != nil && ref.Version.Equal(expectedVersion) {
			return true
		}
	}
	return false
}

func bytesAddressMatches(encoded []byte, expectedAddress common.Address) bool {
	paddedExpected := common.LeftPadBytes(expectedAddress.Bytes(), 32)
	return bytes.Equal(encoded, paddedExpected) || bytes.Equal(encoded, expectedAddress.Bytes())
}
