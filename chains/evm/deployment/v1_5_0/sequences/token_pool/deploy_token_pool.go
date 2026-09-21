package token_pool

import (
	"errors"
	"fmt"

	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	adaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// DeployTokenPool deploys a v1.5.0 token pool.
//
// NOTE: this lives in the token_pool subpackage rather than alongside the other v1.5.0
// sequences because it needs v1_0_0/adapters.EVMTokenBase for the datastore helpers, and
// v1_0_0/adapters already imports v1_5_0/sequences for the TokenAdminRegistry flow. Importing
// the adapters package from v1_5_0/sequences would close that cycle. (v1.5.1 has no such
// constraint: nothing upstream imports v1_5_1/sequences.)
//
// Unlike the v1.5.1 sequence, only BurnMintTokenPoolAndProxy is supported. It is the sole
// v1.5.0 pool type with a deployed footprint, and it is the only one the v1.5.0 TokenAdapter
// claims to configure — the lock-release and rebasing *AndProxy variants have bindings but no
// adapter support, so deploying one here would produce a pool no changeset could then wire up.
var DeployTokenPool = cldf_ops.NewSequence(
	"deploy-token-pool",
	utils.Version_1_5_0,
	"Deploy v1.5.0 token pool contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
		}

		// Validate required deployment inputs
		poolutil := adaptersV1_0_0.EVMTokenBase{}
		if input.TokenPoolVersion == nil {
			return sequences.OnChainOutput{}, errors.New("TokenPoolVersion is required")
		}
		if input.TokenRef == nil {
			return sequences.OnChainOutput{}, errors.New("TokenRef is required")
		}

		// Parse the token ref as an EVM address
		tokenAddress, err := poolutil.ParseNonZeroAddressRef(input.ExistingDataStore, input.TokenRef.Clone(), chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve token address from ref: %w", err)
		}

		// If no pool qualifier is provided, then fall back to using the token address
		poolQualifier := input.TokenPoolQualifier
		if poolQualifier == "" {
			poolQualifier = tokenAddress.Hex()
		}

		// NOTE: the datastore uses the type, selector, qualifier, and version of an address
		// ref to uniquely identify records, so the query below should only match one record
		// at most. If multiple records are returned, then this would indicate an issue with
		// the datastore's data integrity. If no matches are returned, then the ref does not
		// exist and we proceed with the deployment.
		matches := input.ExistingDataStore.Addresses().Filter(
			datastore.AddressRefByType(datastore.ContractType(input.PoolType)),
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByQualifier(poolQualifier),
			datastore.AddressRefByVersion(input.TokenPoolVersion),
		)
		if len(matches) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"multiple token pools found in datastore with type '%s', version '%s', qualifier '%s' on chain with selector %d",
				input.PoolType, input.TokenPoolVersion.String(), poolQualifier, chain.Selector,
			)
		}
		if len(matches) == 1 {
			b.Logger.Infof("Token pool already deployed: %s", datastore_utils.SprintRef(matches[0]))
			return sequences.OnChainOutput{Addresses: matches}, nil
		}

		// Infer pool deployment inputs. NOTE: unlike v1.5.1, the v1.5.0 constructor takes no
		// localTokenDecimals argument, so the token decimals are not read here.
		rmnProxyAddr, err := poolutil.GetRMNProxyAddress(input.ExistingDataStore, chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve rmn proxy address for chain selector %d: %w", chain.Selector, err)
		}
		routerAddr, err := poolutil.ResolveRouterAddress(input.ExistingDataStore, chain.Selector, input.RouterRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve router address for chain selector %d: %w", chain.Selector, err)
		}
		allowlist, err := poolutil.ParseAddressStrings(input.Allowlist)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to parse allowlist: %w", err)
		}

		// Build type and version struct
		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.PoolType),
			*input.TokenPoolVersion,
		)

		// Deploy the desired pool contract
		var poolRef datastore.AddressRef
		switch typeAndVersion.String() {
		case bmtpap.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, bmtpap.Deploy, chain, contract.DeployInput[bmtpap.ConstructorArgs]{
				TypeAndVersion: bmtpap.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: bmtpap.ConstructorArgs{
					Token:     tokenAddress,
					Allowlist: allowlist,
					RmnProxy:  rmnProxyAddr,
					Router:    routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintTokenPoolAndProxy v1.5.0: %w", err)
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported v1.5.0 token pool type and version: %s", typeAndVersion)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{poolRef},
			BatchOps:  []mcms_types.BatchOperation{},
		}, nil
	},
)
