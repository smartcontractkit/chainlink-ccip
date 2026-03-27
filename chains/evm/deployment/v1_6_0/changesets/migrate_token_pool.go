package changesets

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	burn_mint_erc20 "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	burn_mint_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/burn_mint_with_external_minter_token_pool"
	hybrid_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"
	evm_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_1_lock_release_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/lock_release_token_pool"
	v1_6_0_burn_mint_with_external_minter_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/burn_mint_with_external_minter_token_pool"
	v1_6_0_hybrid_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool"
	v1_6_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	v1_5_1_lock_release_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var supportedOldPoolTypeAndVersions = map[string]struct{}{
	v1_5_1_lock_release_token_pool_ops.TypeAndVersion.String(): {},
}

var expectedNewPoolTypeAndVersion = v1_6_0_burn_mint_with_external_minter_token_pool_ops.TypeAndVersion.String()

type MigrateTokenPoolConfig struct {
	HubChainSelector     uint64     `json:"hubChainSelector" yaml:"hubChainSelector"`
	HubPoolAddress       string     `json:"hubPoolAddress" yaml:"hubPoolAddress"`
	RemoteChainSelector  uint64     `json:"remoteChainSelector" yaml:"remoteChainSelector"`
	NewRemotePoolAddress string     `json:"newRemotePoolAddress" yaml:"newRemotePoolAddress"`
	OldRemotePoolAddress string     `json:"oldRemotePoolAddress" yaml:"oldRemotePoolAddress"`
	RemoteChainSupply    *big.Int   `json:"remoteChainSupply" yaml:"remoteChainSupply"`
	TargetGroup          uint8      `json:"targetGroup" yaml:"targetGroup"`
	RemoteTARAddress     string     `json:"remoteTARAddress" yaml:"remoteTARAddress"`
	RemoteTokenAddress   string     `json:"remoteTokenAddress" yaml:"remoteTokenAddress"`
	MCMS                 mcms.Input `json:"mcms,omitempty" yaml:"mcms,omitempty"`
}

func MigrateTokenPool(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MigrateTokenPoolConfig] {
	return cldf.CreateChangeSet(
		makeApplyMigrateTokenPool(mcmsRegistry),
		makeVerifyMigrateTokenPool(mcmsRegistry),
	)
}

func makeVerifyMigrateTokenPool(
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, MigrateTokenPoolConfig) error {
	return func(e cldf.Environment, cfg MigrateTokenPoolConfig) error {
		if err := cfg.MCMS.Validate(); err != nil {
			return fmt.Errorf("invalid MCMS config: %w", err)
		}
		if cfg.HubChainSelector == cfg.RemoteChainSelector {
			return fmt.Errorf("hub chain selector and remote chain selector must be different")
		}
		if cfg.TargetGroup != 0 && cfg.TargetGroup != 1 {
			return fmt.Errorf("target group must be 0 or 1, got %d", cfg.TargetGroup)
		}
		if cfg.RemoteChainSupply == nil {
			return fmt.Errorf("remote chain supply must be provided")
		}
		if cfg.RemoteChainSupply.Sign() < 0 {
			return fmt.Errorf("remote chain supply must be non-negative")
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

		hubPoolAddress, err := parseRequiredAddress("hubPoolAddress", cfg.HubPoolAddress)
		if err != nil {
			return err
		}
		newRemotePoolAddress, err := parseRequiredAddress("newRemotePoolAddress", cfg.NewRemotePoolAddress)
		if err != nil {
			return err
		}
		oldRemotePoolAddress, err := parseRequiredAddress("oldRemotePoolAddress", cfg.OldRemotePoolAddress)
		if err != nil {
			return err
		}
		if oldRemotePoolAddress == newRemotePoolAddress {
			return fmt.Errorf("old remote pool address and new remote pool address must be different")
		}
		remoteTARAddress, err := parseRequiredAddress("remoteTARAddress", cfg.RemoteTARAddress)
		if err != nil {
			return err
		}
		remoteTokenAddress, err := parseRequiredAddress("remoteTokenAddress", cfg.RemoteTokenAddress)
		if err != nil {
			return err
		}

		hubChain, ok := e.BlockChains.EVMChains()[cfg.HubChainSelector]
		if !ok {
			return fmt.Errorf("hub chain selector %d is not configured as an EVM chain", cfg.HubChainSelector)
		}
		remoteChain, ok := e.BlockChains.EVMChains()[cfg.RemoteChainSelector]
		if !ok {
			return fmt.Errorf("remote chain selector %d is not configured as an EVM chain", cfg.RemoteChainSelector)
		}

		mcmsReader, err := getEVMMCMSReader(mcmsRegistry)
		if err != nil {
			return err
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

		hubPoolTypeAndVersion, err := getTypeAndVersionString(hubPoolAddress, hubChain.Client)
		if err != nil {
			return fmt.Errorf("failed to read typeAndVersion for hub pool %s on chain %d: %w", hubPoolAddress, cfg.HubChainSelector, err)
		}
		if hubPoolTypeAndVersion != v1_6_0_hybrid_pool_ops.TypeAndVersion.String() {
			return fmt.Errorf(
				"unexpected hub pool typeAndVersion %q for hub pool %s, expected %q",
				hubPoolTypeAndVersion,
				hubPoolAddress,
				v1_6_0_hybrid_pool_ops.TypeAndVersion.String(),
			)
		}

		oldPoolTypeAndVersion, err := getTypeAndVersionString(oldRemotePoolAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to read typeAndVersion for old remote pool %s on chain %d: %w", oldRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		if _, ok := supportedOldPoolTypeAndVersions[oldPoolTypeAndVersion]; !ok {
			return fmt.Errorf("unsupported old pool typeAndVersion %q for old remote pool %s", oldPoolTypeAndVersion, oldRemotePoolAddress)
		}

		newPoolTypeAndVersion, err := getTypeAndVersionString(newRemotePoolAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to read typeAndVersion for new remote pool %s on chain %d: %w", newRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		if newPoolTypeAndVersion != expectedNewPoolTypeAndVersion {
			return fmt.Errorf(
				"unexpected new pool typeAndVersion %q for new remote pool %s, expected %q",
				newPoolTypeAndVersion,
				newRemotePoolAddress,
				expectedNewPoolTypeAndVersion,
			)
		}

		oldPool, err := v1_5_1_lock_release_token_pool_bindings.NewLockReleaseTokenPool(oldRemotePoolAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind old remote pool %s on chain %d: %w", oldRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		oldPoolToken, err := oldPool.GetToken(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read token from old remote pool %s on chain %d: %w", oldRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		if oldPoolToken != remoteTokenAddress {
			return fmt.Errorf("old remote pool token %s does not match remote token %s", oldPoolToken, remoteTokenAddress)
		}

		newPool, err := burn_mint_with_external_minter_token_pool_bindings.NewBurnMintWithExternalMinterTokenPool(newRemotePoolAddress, remoteChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind new remote pool %s on chain %d: %w", newRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		newPoolToken, err := newPool.GetToken(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read token from new remote pool %s on chain %d: %w", newRemotePoolAddress, cfg.RemoteChainSelector, err)
		}
		if newPoolToken != remoteTokenAddress {
			return fmt.Errorf("new remote pool token %s does not match remote token %s", newPoolToken, remoteTokenAddress)
		}

		hubPool, err := hybrid_with_external_minter_token_pool_bindings.NewHybridWithExternalMinterTokenPool(hubPoolAddress, hubChain.Client)
		if err != nil {
			return fmt.Errorf("failed to bind hub pool %s on chain %d: %w", hubPoolAddress, cfg.HubChainSelector, err)
		}
		isSupportedChain, err := hubPool.IsSupportedChain(&bind.CallOpts{Context: e.GetContext()}, cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("failed to read supported-chain status for remote chain %d from hub pool %s on chain %d: %w", cfg.RemoteChainSelector, hubPoolAddress, cfg.HubChainSelector, err)
		}
		if !isSupportedChain {
			return fmt.Errorf("remote chain %d is not supported on hub pool %s on chain %d", cfg.RemoteChainSelector, hubPoolAddress, cfg.HubChainSelector)
		}

		hubRemotePools, err := hubPool.GetRemotePools(&bind.CallOpts{Context: e.GetContext()}, cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("failed to read hub remote pools for remote chain %d from hub pool %s on chain %d: %w", cfg.RemoteChainSelector, hubPoolAddress, cfg.HubChainSelector, err)
		}
		oldPoolBytes := common.LeftPadBytes(oldRemotePoolAddress.Bytes(), 32)
		newPoolBytes := common.LeftPadBytes(newRemotePoolAddress.Bytes(), 32)
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
			return fmt.Errorf("neither old pool %s nor new pool %s registered in hub remote pool set for chain %d", oldRemotePoolAddress, newRemotePoolAddress, cfg.RemoteChainSelector)
		}

		hubRemoteToken, err := hubPool.GetRemoteToken(&bind.CallOpts{Context: e.GetContext()}, cfg.RemoteChainSelector)
		if err != nil {
			return fmt.Errorf("failed to read hub remote token bytes for remote chain %d from hub pool %s on chain %d: %w", cfg.RemoteChainSelector, hubPoolAddress, cfg.HubChainSelector, err)
		}
		if !bytesAddressMatches(hubRemoteToken, remoteTokenAddress) {
			return fmt.Errorf(
				"hub remote token bytes %x do not match remote token %s for chain %d",
				hubRemoteToken,
				remoteTokenAddress,
				cfg.RemoteChainSelector,
			)
		}

		hubPoolOwner, err := hubPool.Owner(&bind.CallOpts{Context: e.GetContext()})
		if err != nil {
			return fmt.Errorf("failed to read owner for hub pool %s on chain %d: %w", hubPoolAddress, cfg.HubChainSelector, err)
		}
		hubTimelockAddress := common.HexToAddress(hubTimelockRef.Address)
		if hubPoolOwner != hubTimelockAddress {
			return fmt.Errorf(
				"hub pool %s owner %s does not match timelock %s on chain %d",
				hubPoolAddress,
				hubPoolOwner,
				hubTimelockAddress,
				cfg.HubChainSelector,
			)
		}

		tarConfigReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, tar_ops.GetTokenConfig, remoteChain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: cfg.RemoteChainSelector,
			Address:       remoteTARAddress,
			Args:          remoteTokenAddress,
		})
		if err != nil {
			return fmt.Errorf("failed to read TAR token config for token %s on chain %d: %w", remoteTokenAddress, cfg.RemoteChainSelector, err)
		}
		remoteTimelockAddress := common.HexToAddress(remoteTimelockRef.Address)
		if tarConfigReport.Output.Administrator != remoteTimelockAddress {
			return fmt.Errorf(
				"TAR administrator %s for token %s does not match timelock %s on chain %d",
				tarConfigReport.Output.Administrator,
				remoteTokenAddress,
				remoteTimelockAddress,
				cfg.RemoteChainSelector,
			)
		}

		groupReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, v1_6_0_hybrid_pool_ops.GetGroup, hubChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: cfg.HubChainSelector,
			Address:       hubPoolAddress,
			Args:          cfg.RemoteChainSelector,
		})
		if err != nil {
			return fmt.Errorf("failed to read group for remote chain %d on hub chain %d: %w", cfg.RemoteChainSelector, cfg.HubChainSelector, err)
		}
		if groupReport.Output != cfg.TargetGroup {
			remoteToken, err := burn_mint_erc20.NewBurnMintERC20(remoteTokenAddress, remoteChain.Client)
			if err != nil {
				return fmt.Errorf("failed to bind remote token %s on chain %d: %w", remoteTokenAddress, cfg.RemoteChainSelector, err)
			}
			totalSupply, err := remoteToken.TotalSupply(&bind.CallOpts{Context: e.GetContext()})
			if err != nil {
				return fmt.Errorf("failed to read totalSupply for remote token %s on chain %d: %w", remoteTokenAddress, cfg.RemoteChainSelector, err)
			}
			if totalSupply.Cmp(cfg.RemoteChainSupply) != 0 {
				return fmt.Errorf(
					"remote chain supply %s does not match remote token totalSupply %s for token %s on chain %d",
					cfg.RemoteChainSupply.String(),
					totalSupply.String(),
					remoteTokenAddress,
					cfg.RemoteChainSelector,
				)
			}

			if cfg.TargetGroup == 1 {
				lockedTokensReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, v1_6_0_hybrid_pool_ops.GetLockedTokens, hubChain, evm_contract.FunctionInput[struct{}]{
					ChainSelector: cfg.HubChainSelector,
					Address:       hubPoolAddress,
					Args:          struct{}{},
				})
				if err != nil {
					return fmt.Errorf("failed to read locked token accounting from hub pool %s on chain %d: %w", hubPoolAddress, cfg.HubChainSelector, err)
				}
				if lockedTokensReport.Output == nil {
					return fmt.Errorf("hub pool %s returned nil locked token accounting on chain %d", hubPoolAddress, cfg.HubChainSelector)
				}
				if cfg.RemoteChainSupply.Cmp(lockedTokensReport.Output) > 0 {
					return fmt.Errorf(
						"remote chain supply %s exceeds locked token accounting %s for hub pool %s on chain %d",
						cfg.RemoteChainSupply.String(),
						lockedTokensReport.Output.String(),
						hubPoolAddress,
						cfg.HubChainSelector,
					)
				}
			}
		}

		if err := validateAddressRefIfPresent(
			e.DataStore,
			cfg.HubChainSelector,
			hubPoolAddress,
			cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
			v1_6_0_hybrid_pool_ops.Version,
			"hub pool",
		); err != nil {
			return err
		}
		if err := validateAddressRefIfPresent(
			e.DataStore,
			cfg.RemoteChainSelector,
			oldRemotePoolAddress,
			cldf_datastore.ContractType(v1_5_1_lock_release_token_pool_ops.ContractType),
			v1_5_1_lock_release_token_pool_ops.Version,
			"old remote pool",
		); err != nil {
			return err
		}
		if err := validateAddressRefIfPresent(
			e.DataStore,
			cfg.RemoteChainSelector,
			newRemotePoolAddress,
			cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
			v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
			"new remote pool",
		); err != nil {
			return err
		}
		if err := validateAddressRefIfPresent(
			e.DataStore,
			cfg.RemoteChainSelector,
			remoteTARAddress,
			cldf_datastore.ContractType(tar_ops.ContractType),
			tar_ops.Version,
			"remote token admin registry",
		); err != nil {
			return err
		}

		return nil
	}
}

func makeApplyMigrateTokenPool(
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, MigrateTokenPoolConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MigrateTokenPoolConfig) (cldf.ChangesetOutput, error) {
		hubPoolAddress, err := parseRequiredAddress("hubPoolAddress", cfg.HubPoolAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		newRemotePoolAddress, err := parseRequiredAddress("newRemotePoolAddress", cfg.NewRemotePoolAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		oldRemotePoolAddress, err := parseRequiredAddress("oldRemotePoolAddress", cfg.OldRemotePoolAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		remoteTARAddress, err := parseRequiredAddress("remoteTARAddress", cfg.RemoteTARAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		remoteTokenAddress, err := parseRequiredAddress("remoteTokenAddress", cfg.RemoteTokenAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, v1_6_0_sequences.MigrateTokenPool, e.BlockChains, v1_6_0_sequences.MigrateTokenPoolInput{
			HubChainSelector:     cfg.HubChainSelector,
			HubPoolAddress:       hubPoolAddress,
			RemoteChainSelector:  cfg.RemoteChainSelector,
			NewRemotePoolAddress: newRemotePoolAddress,
			OldRemotePoolAddress: oldRemotePoolAddress,
			RemoteChainSupply:    cfg.RemoteChainSupply,
			TargetGroup:          cfg.TargetGroup,
			RemoteTARAddress:     remoteTARAddress,
			RemoteTokenAddress:   remoteTokenAddress,
		})
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to migrate token pool: %w", err)
		}

		ds := cldf_datastore.NewMemoryDataStore()
		if !addressRefExistsWithTypeVersion(
			e.DataStore,
			cfg.HubChainSelector,
			hubPoolAddress,
			cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
			v1_6_0_hybrid_pool_ops.Version,
		) {
			if err := ds.Addresses().Add(cldf_datastore.AddressRef{
				ChainSelector: cfg.HubChainSelector,
				Type:          cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
				Version:       v1_6_0_hybrid_pool_ops.Version,
				Address:       hubPoolAddress.Hex(),
			}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to persist hub pool ref %s on chain %d: %w", hubPoolAddress, cfg.HubChainSelector, err)
			}
		}
		if !addressRefExistsWithTypeVersion(
			e.DataStore,
			cfg.RemoteChainSelector,
			newRemotePoolAddress,
			cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
			v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
		) {
			if err := ds.Addresses().Add(cldf_datastore.AddressRef{
				ChainSelector: cfg.RemoteChainSelector,
				Type:          cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
				Version:       v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
				Address:       newRemotePoolAddress.Hex(),
			}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to persist new remote pool ref %s on chain %d: %w", newRemotePoolAddress, cfg.RemoteChainSelector, err)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(report.ExecutionReports).
			WithBatchOps(report.Output.BatchOps).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}

func parseRequiredAddress(fieldName string, value string) (common.Address, error) {
	if !common.IsHexAddress(value) {
		return common.Address{}, fmt.Errorf("%s is not a valid hex address: %s", fieldName, value)
	}
	address := common.HexToAddress(value)
	if address == (common.Address{}) {
		return common.Address{}, fmt.Errorf("%s cannot be the zero address", fieldName)
	}
	return address, nil
}

func getEVMMCMSReader(mcmsRegistry *changesets.MCMSReaderRegistry) (changesets.MCMSReader, error) {
	if mcmsRegistry == nil {
		return nil, fmt.Errorf("no MCMS reader registry configured")
	}
	reader, ok := mcmsRegistry.GetMCMSReader(chain_selectors.FamilyEVM)
	if !ok {
		return nil, fmt.Errorf("no MCMS reader registered for chain family '%s'", chain_selectors.FamilyEVM)
	}
	return reader, nil
}

func getTypeAndVersionString(address common.Address, backend bind.ContractBackend) (string, error) {
	contractType, version, err := evm_utils.TypeAndVersion(address, backend)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", contractType, version.String()), nil
}

func validateAddressRefIfPresent(
	ds cldf_datastore.DataStore,
	chainSelector uint64,
	address common.Address,
	expectedType cldf_datastore.ContractType,
	expectedVersion *semver.Version,
	label string,
) error {
	refs := ds.Addresses().Filter(
		cldf_datastore.AddressRefByChainSelector(chainSelector),
		cldf_datastore.AddressRefByAddress(address.Hex()),
	)
	if len(refs) == 0 {
		return nil
	}
	for _, ref := range refs {
		if ref.Type == expectedType && ref.Version != nil && ref.Version.Equal(expectedVersion) {
			return nil
		}
	}
	return fmt.Errorf(
		"%s %s on chain %d has datastore refs but none match expected type=%s version=%s",
		label,
		address,
		chainSelector,
		expectedType,
		expectedVersion.String(),
	)
}

func addressRefExistsWithTypeVersion(
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
