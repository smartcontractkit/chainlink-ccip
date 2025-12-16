package tokens

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type TokenExpansionInput struct {
	TokenExpansionInputPerChain map[uint64]TokenExpansionInputPerChain `yaml:"token-expansion-input-per-chain" json:"tokenExpansionInputPerChain"`
	MCMS                        mcms.Input                             `yaml:"mcms,omitempty" json:"mcms,omitempty"`
}

type TokenExpansionInputPerChain struct {
	DeployTokenInput        DeployTokenInput `yaml:"deploy-token-input" json:"deployTokenInput"`
	TokenPoolQualifier      string           `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType                string           `yaml:"pool-type" json:"poolType"`
	TARAdmin                string           `yaml:"tar-admin" json:"tarAdmin"`
	TokenPoolAdmin          string           `yaml:"token-pool-admin" json:"tokenPoolAdmin"`
	TokenPoolRateLimitAdmin string           `yaml:"token-pool-rate-limit-admin" json:"tokenPoolRateLimitAdmin"`
}

type DeployTokenInput struct {
	Name     string   `yaml:"name" json:"name"`
	Symbol   string   `yaml:"symbol" json:"symbol"`
	Decimals uint8    `yaml:"decimals" json:"decimals"`
	Supply   *big.Int `yaml:"supply" json:"supply"`
	//
	PreMint *big.Int `yaml:"pre-mint" json:"preMint"`
	// Address to be set as the CCIP admin on the token contract
	// who will be allowed to register the token pool for this token in the TokenAdminRegistry
	// if not specified, defaults to the timelock address
	CCIPAdmin common.Address `yaml:"ccip-admin" json:"ccipAdmin"`
	// Customer admin who will be granted admin rights on the token
	ExternalAdmin common.Address `yaml:"external-admin" json:"externalAdmin"`
	// list of addresses who may need special processing in order to send tokens
	// e.g. for Solana, addresses that need associated token accounts created
	Senders []string `yaml:"senders" json:"senders"`
	// SPLToken, ERC20, etc.
	Type cldf.ContractType `yaml:"type" json:"type"`
	// Solana Specific
	// private key in base58 encoding for vanity addresses
	TokenPrivKey string `yaml:"token-priv-key" json:"tokenPrivKey"`
	// if true, the freeze authority will be revoked on token creation
	// and it will be disabled FOREVER
	DisableFreezeAuthority bool `yaml:"disable-freeze-authority" json:"disableFreezeAuthority"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type DeployTokenPoolInput struct {
	TokenSymbol        string `yaml:"token-symbol" json:"tokenSymbol"`
	TokenPoolQualifier string `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string `yaml:"pool-type" json:"poolType"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type RegisterTokenInput struct {
	TokenSymbol string `yaml:"token-symbol" json:"tokenSymbol"`
	TokenAdmin  string `yaml:"token-admin" json:"tokenAdmin"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type SetPoolInput struct {
	TokenSymbol        string `yaml:"token-symbol" json:"tokenSymbol"`
	TokenPoolQualifier string `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string `yaml:"pool-type" json:"poolType"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type UpdateAuthoritiesInput struct {
	RegisterTokenConfig RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}

func TokenExpansion() cldf.ChangeSetV2[TokenExpansionInput] {
	return cldf.CreateChangeSet(tokenExpansionApply(), tokenExpansionVerify())
}

func tokenExpansionVerify() func(cldf.Environment, TokenExpansionInput) error {
	return func(e cldf.Environment, cfg TokenExpansionInput) error {
		// TODO: implement
		return nil
	}
}

func tokenExpansionApply() func(cldf.Environment, TokenExpansionInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg TokenExpansionInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		// ds to collect all addresses created during this changeset
		// this gets passed as output
		ds := datastore.NewMemoryDataStore()
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()
		for selector, input := range cfg.TokenExpansionInputPerChain {
			tmpDatastore := datastore.NewMemoryDataStore()
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, &TokenAdminRegistryVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			// deploy token
			deployTokenInput := input.DeployTokenInput
			deployTokenInput.ExistingDataStore = e.DataStore
			deployTokenInput.ChainSelector = selector
			deployTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployToken(), e.BlockChains, deployTokenInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
			}
			batchOps = append(batchOps, deployTokenReport.Output.BatchOps...)
			reports = append(reports, deployTokenReport.ExecutionReports...)
			for _, r := range deployTokenReport.Output.Addresses {
				if err := tmpDatastore.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			tmpDatastore.Merge(e.DataStore)
			e.DataStore = tmpDatastore.Seal()

			// deploy token pool
			tmpDatastore = datastore.NewMemoryDataStore()
			deployTokenPoolInput := DeployTokenPoolInput{
				TokenSymbol:        deployTokenInput.Symbol,
				TokenPoolQualifier: input.TokenPoolQualifier,
				PoolType:           input.PoolType,
			}
			deployTokenPoolInput.ExistingDataStore = e.DataStore
			deployTokenPoolInput.ChainSelector = selector
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployTokenPoolForToken(), e.BlockChains, deployTokenPoolInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
			}
			batchOps = append(batchOps, deployTokenPoolReport.Output.BatchOps...)
			reports = append(reports, deployTokenPoolReport.ExecutionReports...)
			for _, r := range deployTokenPoolReport.Output.Addresses {
				if err := tmpDatastore.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			tmpDatastore.Merge(e.DataStore)
			e.DataStore = tmpDatastore.Seal()

			// register token
			tmpDatastore = datastore.NewMemoryDataStore()
			registerTokenInput := RegisterTokenInput{
				TokenSymbol: deployTokenInput.Symbol,
				TokenAdmin:  input.TARAdmin,
			}
			registerTokenInput.ExistingDataStore = e.DataStore
			registerTokenInput.ChainSelector = selector
			registerTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.RegisterToken(), e.BlockChains, registerTokenInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
			}
			batchOps = append(batchOps, registerTokenReport.Output.BatchOps...)
			reports = append(reports, registerTokenReport.ExecutionReports...)
			for _, r := range registerTokenReport.Output.Addresses {
				if err := tmpDatastore.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			tmpDatastore.Merge(e.DataStore)
			e.DataStore = tmpDatastore.Seal()

			// set pool
			tmpDatastore = datastore.NewMemoryDataStore()
			setPoolInput := SetPoolInput{
				TokenSymbol:        deployTokenInput.Symbol,
				TokenPoolQualifier: input.TokenPoolQualifier,
				PoolType:           input.PoolType,
			}
			setPoolInput.ExistingDataStore = e.DataStore
			setPoolInput.ChainSelector = selector
			setPoolReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.SetPool(), e.BlockChains, setPoolInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
			}
			batchOps = append(batchOps, setPoolReport.Output.BatchOps...)
			reports = append(reports, setPoolReport.ExecutionReports...)
			for _, r := range setPoolReport.Output.Addresses {
				if err := tmpDatastore.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			tmpDatastore.Merge(e.DataStore)
			e.DataStore = tmpDatastore.Seal()
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
