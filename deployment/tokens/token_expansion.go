package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type TokenExpansionInput struct {
	// per-chain configuration for token expansion
	TokenExpansionInputPerChain map[uint64]TokenExpansionInputPerChain `yaml:"token-expansion-input-per-chain" json:"tokenExpansionInputPerChain"`
	ChainAdapterVersion         *semver.Version                        `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
	MCMS                        mcms.Input                             `yaml:"mcms,omitempty" json:"mcms"`
}

type TokenExpansionInputPerChain struct {
	TokenPoolVersion *semver.Version  `yaml:"token-pool-version" json:"tokenPoolVersion"`
	DeployTokenInput DeployTokenInput `yaml:"deploy-token-input" json:"deployTokenInput"`
	// only necessary if we want to specifically query for a token pool with a given type + qualifier
	TokenPoolQualifier string `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string `yaml:"pool-type" json:"poolType"`
	// only necessary if we want to set specific admin authorities. Will default to timelock admin otherwise
	TARAdmin                string `yaml:"tar-admin" json:"tarAdmin"`
	TokenPoolAdmin          string `yaml:"token-pool-admin" json:"tokenPoolAdmin"`
	TokenPoolRateLimitAdmin string `yaml:"token-pool-rate-limit-admin" json:"tokenPoolRateLimitAdmin"`
	// rate lmiter config per remote chain
	// we will look up the remote token from the top level token expansion config
	RemoteCounterpartUpdates map[uint64]RateLimiterConfig `yaml:"remote-counterpart-updates" json:"remoteCounterpartUpdates"`
	// if true, will delete the remote counterpart token pool on the specified chains
	RemoteCounterpartDeletes []uint64 `yaml:"remote-counterpart-deletes" json:"remoteCounterpartDeletes"`
}

type DeployTokenInput struct {
	Name     string   `yaml:"name" json:"name"`
	Symbol   string   `yaml:"symbol" json:"symbol"`
	Decimals uint8    `yaml:"decimals" json:"decimals"`
	Supply   *big.Int `yaml:"supply" json:"supply"`
	PreMint  *big.Int `yaml:"pre-mint" json:"preMint"`
	// Customer admin who will be granted admin rights on the token
	// For EVM, expect to have only one Admin address to be passed on whereas Solana may have multiple multisig signers.
	// Use string to keep this struct chain-agnostic (EVM uses hex, Solana uses base58, etc.)
	ExternalAdmin []string `yaml:"external-admin" json:"externalAdmin"`
	// Address to be set as the CCIP admin on the token contract, defaults to the timelock address
	CCIPAdmin string
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
	TokenSymbol        string          `yaml:"token-symbol" json:"tokenSymbol"`
	TokenPoolQualifier string          `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string          `yaml:"pool-type" json:"poolType"`
	TokenPoolVersion   *semver.Version `yaml:"token-pool-version" json:"tokenPoolVersion"`
	Allowlist          []string        `yaml:"allowlist" json:"allowlist"`
	// AcceptLiquidity is used by LockReleaseTokenPool (v1.5.1 only) to indicate
	// whether the pool should accept liquidity from liquidity providers
	AcceptLiquidity *bool `yaml:"accept-liquidity" json:"acceptLiquidity"`
	// BurnAddress is used by BurnToAddressMintTokenPool to specify the address
	// where tokens will be burned to
	BurnAddress string `yaml:"burn-address" json:"burnAddress"`
	// TokenGovernor is used by BurnMintWithExternalMinterTokenPool kind of pools to specify the token governor contract address
	// if it is not provided, the token governor will be fetched from the datastore based on the token symbol
	TokenGovernor string `yaml:"token-governor,omitempty" json:"tokenGovernor,omitempty"`
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
	TokenPoolAdmin          string `yaml:"token-pool-admin" json:"tokenPoolAdmin"`
	TokenPoolRateLimitAdmin string `yaml:"token-pool-rate-limit-admin" json:"tokenPoolRateLimitAdmin"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}

func TokenExpansion() cldf.ChangeSetV2[TokenExpansionInput] {
	return cldf.CreateChangeSet(tokenExpansionApply(), tokenExpansionVerify())
}

func tokenExpansionVerify() func(cldf.Environment, TokenExpansionInput) error {
	return func(e cldf.Environment, cfg TokenExpansionInput) error {
		tokenPoolRegistry := GetTokenAdapterRegistry()
		for selector, input := range cfg.TokenExpansionInputPerChain {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return fmt.Errorf("not a valid selector: %v", err)
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, cfg.ChainAdapterVersion)
			if !exists {
				return fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			// deploy token
			deployTokenInput := input.DeployTokenInput
			deployTokenInput.ExistingDataStore = e.DataStore
			deployTokenInput.ChainSelector = selector
			err = tokenPoolAdapter.DeployTokenVerify(e, input)
			if err != nil {
				return fmt.Errorf("failed to verify deploy token input for chain selector %d: %w", selector, err)
			}
		}
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
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, cfg.ChainAdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}

			// deploy token
			deployTokenInput := input.DeployTokenInput
			deployTokenInput.ExistingDataStore = e.DataStore
			deployTokenInput.ChainSelector = selector

			// if token is deployed by CLL, set CCIP admin as RBACTimelock by default.
			// If input has CCIPAdmin and which is external address, set that address as CCIPAdmin
			// and we may not be able to register the token by CLL in that case.
			if deployTokenInput.CCIPAdmin == "" {
				filter := datastore.AddressRef{
					Type:          datastore.ContractType(common_utils.RBACTimelock),
					ChainSelector: deployTokenInput.ChainSelector,
					Qualifier:     cfg.MCMS.Qualifier,
				}

				timelockAddr, err := datastore_utils.FindAndFormatRef(
					deployTokenInput.ExistingDataStore,
					filter,
					deployTokenInput.ChainSelector,
					datastore_utils.FullRef,
				)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf(
						"couldn't find the RBACTimelock address in datastore for selector %d and qualifier %s: %w",
						deployTokenInput.ChainSelector, cfg.MCMS.Qualifier, err,
					)
				}

				deployTokenInput.CCIPAdmin = timelockAddr.Address
			}
			deployTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployToken(), e.BlockChains, deployTokenInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy token on chain %d: %w", selector, err)
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
				TokenPoolVersion:   input.TokenPoolVersion,
				PoolType:           input.PoolType,
			}
			deployTokenPoolInput.ExistingDataStore = e.DataStore
			deployTokenPoolInput.ChainSelector = selector
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployTokenPoolForToken(), e.BlockChains, deployTokenPoolInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy token pool for token on chain %d: %w", selector, err)
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
