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
	TokenPoolVersion *semver.Version `yaml:"token-pool-version" json:"tokenPoolVersion"`
	// will deploy a token if DeployTokenInput is not nil,
	// otherwise assumes token is already deployed and will look up the token address from the datastore based on the input parameters
	DeployTokenInput *DeployTokenInput `yaml:"deploy-token-input" json:"deployTokenInput"`
	// will deploy a token pool for the token if DeployTokenPoolInput is not nil
	DeployTokenPoolInput *DeployTokenPoolInput `yaml:"deploy-token-pool-input" json:"deployTokenPoolInput"`
	// if not nil, will try to fully configure the token for transfers, including registering the token and token pool on-chain and setting the pool on the token
	TokenTransferConfig *TokenTransferConfig `yaml:"token-transfer-config" json:"tokenTransferConfig"`
}

type DeployTokenInput struct {
	Name     string   `yaml:"name" json:"name"`
	Symbol   string   `yaml:"symbol" json:"symbol"`
	Decimals uint8    `yaml:"decimals" json:"decimals"`
	Supply   *big.Int `yaml:"supply" json:"supply"`
	PreMint  *big.Int `yaml:"pre-mint" json:"preMint"`
	// Customer admin who will be granted admin rights on the token
	// Use string to keep this struct chain-agnostic (EVM uses hex, Solana uses base58, etc.)
	ExternalAdmin string `yaml:"external-admin" json:"externalAdmin"`
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
	// Token metadata to be uploaded
	TokenMetadata *TokenMetadata `yaml:"token-metadata,omitempty" json:"tokenMetadata,omitempty"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}

// Right now this is only used for Solana tokens but we can extend this to other VMs if needed in the future
type TokenMetadata struct {
	// not specified by the user. Overwritten by the deployment system to pass
	// the token address to chain operations for metadata upload and updates
	TokenPubkey string `yaml:"token-pubkey" json:"tokenPubkey"`
	// https://metaboss.dev/create.html#metadata
	// only to be provided on initial upload, it takes in name, symbol, uri
	// after initial upload, those fields can be updated using the update inputs
	// put the json in ccip/env/input dir in CLD
	MetadataJSONPath string `yaml:"metadata-json-path" json:"metadataJsonPath"`
	UpdateAuthority  string `yaml:"update-authority" json:"updateAuthority"` // used to set update authority of the token metadata PDA after initial upload
	// https://metaboss.dev/update.html#update-name
	UpdateName string `yaml:"update-name" json:"updateName"` // used to update the name of the token metadata PDA after initial upload
	// https://metaboss.dev/update.html#update-symbol
	UpdateSymbol string `yaml:"update-symbol" json:"updateSymbol"` // used to update the symbol of the token metadata PDA after initial upload
	// https://metaboss.dev/update.html#update-uri
	UpdateURI string `yaml:"update-uri" json:"updateUri"` // used to update the uri of the token metadata PDA after initial upload
}

type DeployTokenPoolInput struct {
	// TokenRef is a reference to the token in the datastore.
	// If this is provided, it will be cross checked against the deployed token
	TokenRef           *datastore.AddressRef `yaml:"token-ref" json:"tokenRef"`
	TokenPoolQualifier string                `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string                `yaml:"pool-type" json:"poolType"`
	TokenPoolVersion   *semver.Version       `yaml:"token-pool-version" json:"tokenPoolVersion"`
	Allowlist          []string              `yaml:"allowlist" json:"allowlist"`
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
	TokenRef   datastore.AddressRef `yaml:"token-ref" json:"tokenRef"`
	TokenAdmin string               `yaml:"token-admin" json:"tokenAdmin"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type SetPoolInput struct {
	TokenRef           datastore.AddressRef `yaml:"token-ref" json:"tokenRef"`
	TokenPoolQualifier string               `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string               `yaml:"pool-type" json:"poolType"`
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
		allRemotes := make(map[uint64]RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef])
		allTokenConfigs := make([]TokenTransferConfig, 0)
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
			var tokenRef *datastore.AddressRef
			var tokenPool *datastore.AddressRef
			if deployTokenInput != nil {
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
				deployTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployToken(), e.BlockChains, *deployTokenInput)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy token on chain %d: %w", selector, err)
				}
				batchOps = append(batchOps, deployTokenReport.Output.BatchOps...)
				reports = append(reports, deployTokenReport.ExecutionReports...)
				if len(deployTokenReport.Output.Addresses) != 0 {
					tokenRef = &deployTokenReport.Output.Addresses[0]
				}
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
			}

			if input.DeployTokenPoolInput != nil {
				refToConnect := tokenRef
				providedRef := input.DeployTokenPoolInput.TokenRef
				if refToConnect == nil && providedRef == nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("no token deployed or provided for chain selector %d, cannot deploy token pool without token address", selector)
				} else if refToConnect != nil && providedRef != nil {
					// cross check the deployed token address with the provided token ref address
					if refToConnect.Address != providedRef.Address {
						return cldf.ChangesetOutput{}, fmt.Errorf("token address deployed does not match the provided token ref address for chain selector %d: deployed token address %s, provided token ref address %s", selector, refToConnect.Address, providedRef.Address)
					}
					if refToConnect.Qualifier != providedRef.Qualifier {
						return cldf.ChangesetOutput{}, fmt.Errorf("token qualifier deployed does not match the provided token ref qualifier for chain selector %d: deployed token qualifier %s, provided token ref qualifier %s", selector, refToConnect.Qualifier, providedRef.Qualifier)
					}
				} else if refToConnect == nil {
					// if token is not deployed by this changeset but token ref is provided, use the provided token ref
					refToConnect = input.DeployTokenPoolInput.TokenRef
				}
				// deploy token pool
				tmpDatastore = datastore.NewMemoryDataStore()
				deployTokenPoolInput := DeployTokenPoolInput{
					TokenRef:           refToConnect,
					TokenPoolVersion:   input.TokenPoolVersion,
					TokenPoolQualifier: input.DeployTokenPoolInput.TokenPoolQualifier,
					PoolType:           input.DeployTokenPoolInput.PoolType,
				}
				deployTokenPoolInput.ExistingDataStore = e.DataStore
				deployTokenPoolInput.ChainSelector = selector
				deployTokenPoolReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployTokenPoolForToken(), e.BlockChains, deployTokenPoolInput)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy token pool for token on chain %d: %w", selector, err)
				}
				batchOps = append(batchOps, deployTokenPoolReport.Output.BatchOps...)
				reports = append(reports, deployTokenPoolReport.ExecutionReports...)
				if len(deployTokenPoolReport.Output.Addresses) != 0 {
					tokenPool = &deployTokenPoolReport.Output.Addresses[0]
				}
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
			}
			allRemotes[selector] = RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
				RemoteToken: tokenRef,
				RemotePool:  tokenPool,
			}
			// if token transfer config is provided, we will update the remote chain config with the token and token pool addresses and
			// save the token transfer config for processing after all tokens and token pools have been deployed
			if input.TokenTransferConfig != nil {
				actualPool := tokenPool
				if actualPool == nil {
					// if token pool is not deployed by this changeset, we expect the user to provide the token pool address in the TokenTransferConfig
					actualPool = &input.TokenTransferConfig.TokenPoolRef
					if actualPool == nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("no token pool deployed or provided for chain selector %d, cannot configure token for transfers without token pool address", selector)
					}
				}
				actualToken := tokenRef
				if actualToken == nil {
					// if token is not deployed by this changeset, we expect the user to provide the token address in the TokenTransferConfig
					actualToken = &input.TokenTransferConfig.TokenRef
					if actualToken == nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("no token deployed or provided for chain selector %d, cannot configure token for transfers without token address", selector)
					}
				}
				allRemotes[selector] = RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					RemoteToken: actualToken,
					RemotePool:  actualPool,
				}
			}
		}

		// now that we have all the token and token pools, we can loop through the token configs again
		// and update the remote chain configs with the correct token and token pool addresses before configuring the tokens for transfers
		for selector, input := range cfg.TokenExpansionInputPerChain {
			if input.TokenTransferConfig != nil {
				for remoteSelector, remoteConfig := range input.TokenTransferConfig.RemoteChains {
					if _, exists := allRemotes[remoteSelector]; exists {
						if remoteConfig.RemoteToken == nil {
							remoteConfig.RemoteToken = allRemotes[remoteSelector].RemoteToken
						}
						if remoteConfig.RemotePool == nil {
							remoteConfig.RemotePool = allRemotes[remoteSelector].RemotePool
						}
						cfg.TokenExpansionInputPerChain[selector].TokenTransferConfig.RemoteChains[remoteSelector] = remoteConfig
					} else {
						allRemotes[remoteSelector] = remoteConfig
					}
				}
				if len(input.TokenTransferConfig.RemoteChains) != 0 {
					allTokenConfigs = append(allTokenConfigs, *input.TokenTransferConfig)
				}
			}
		}

		fmt.Printf("all remote configs: %+v", allRemotes)
		// finally, we process the token configs for transfers, which will register the tokens and token pools on-chain and set the pool on the token if necessary
		transferOps, transferReports, err := processTokenConfigForChain(e, allTokenConfigs, cfg.ChainAdapterVersion)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to process token configs for transfers: %w", err)
		}
		batchOps = append(batchOps, transferOps...)
		reports = append(reports, transferReports...)

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
