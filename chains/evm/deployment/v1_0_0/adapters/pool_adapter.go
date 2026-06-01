package adapters

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/tokenimpl"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	_ tokensapi.TokenRefResolver = &EVMPoolAdapter{}
	_ tokensapi.TokenAdapter     = &EVMPoolAdapter{}
)

// PoolOps abstracts the version-specific token pool contract calls.
// Each EVM pool version (v1.5.1, v1.6.1) provides an implementation
// that wires into its own bindings/operations.
type PoolOps interface {
	GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error)
	GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error)
	GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error)
	SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error)
	SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) ([]evm_contract.WriteOutput, error)
	// GetCurrentInboundRateLimit reads the on-chain inbound rate limiter state for the given remote
	// chain selector from the token pool at poolAddr. Used by outbound-only TPRL writes to read and
	// pass through the current inbound, and by RateLimitReaderAdapter for cross-chain validation.
	// Returns a zero-value RateLimiterConfig (IsEnabled=false, Capacity=0, Rate=0) when the pool has
	// no inbound configured for the lane.
	GetCurrentInboundRateLimit(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, fastFinality bool) (tokensapi.RateLimiterConfig, error)
	Version() *semver.Version
}

// EVMPoolAdapter provides the shared pool-specific TokenAdapter methods
// for EVM pool versions that follow the same datastore + TAR + BnM pattern
// (currently v1.5.1 and v1.6.1). Version-specific contract calls are
// delegated to the Ops field.
//
// Per-version adapters embed this struct and override only
// ConfigureTokenForTransfersSequence (which differs structurally).
type EVMPoolAdapter struct {
	EVMTokenBase
	Ops PoolOps
	// DeployTokenPoolSeq is injected at construction time to avoid an import
	// cycle between v1_0_0/adapters and v1_6_0/sequences.
	DeployTokenPoolSeq *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

func (a *EVMPoolAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) (string, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return "", fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	// If the ref already has the pool address, then skip the datastore lookup altogether
	// and use it to get the token. If the pool address is NOT in the ref, then fall back
	// to resolving it from the datastore first.
	tokenPoolAddr, err := a.EVMTokenBase.ParseNonZeroAddressRef(e.DataStore, poolRef, chainSelector)
	if err != nil {
		return "", fmt.Errorf("failed to parse token pool address from ref (%s): %w", datastore_utils.SprintRef(poolRef), err)
	}
	tokenAddr, err := a.Ops.GetToken(e.OperationsBundle, chain, tokenPoolAddr)
	if err != nil {
		return "", fmt.Errorf("failed to get token address from token pool ref (%s): %w", datastore_utils.SprintRef(poolRef), err)
	}

	return tokenAddr.Hex(), nil
}

func (a *EVMPoolAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	// Optimization: most tokens are ERC20s, so try to get the decimals directly from
	// the token contract first instead of going through the datastore and token pool
	// contract.
	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle, erc20.GetDecimals, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       common.BytesToAddress(token),
		},
	)
	if err == nil {
		return report.Output, nil
	} else {
		e.Logger.Warnf(
			"failed to get token decimals directly from token contract at address %s - trying token pool at %s: %v",
			common.BytesToAddress(token).Hex(), datastore_utils.SprintRef(poolRef), err,
		)
	}

	// If we can't source the decimals directly from the token then check if the pool
	// address is directly available in the ref. If so, we can skip the datastore and
	// go straight to the pool contract for the decimals. If the ref doesn't have the
	// pool address, then we need to hit the datastore for the full pool ref then get
	// the token decimals from the pool contract.
	poolAddr, err := a.EVMTokenBase.ParseNonZeroAddressRef(e.DataStore, poolRef, chainSelector)
	if err != nil {
		return 0, fmt.Errorf("failed to find token pool address for ref (%s): %w", datastore_utils.SprintRef(poolRef), err)
	} else {
		return a.Ops.GetTokenDecimals(e.OperationsBundle, chain, poolAddr)
	}
}

// GetOnchainInboundRateLimit reads the on-chain inbound rate limiter state on the token pool
// referenced by poolRef on chainSelector, for the given remote selector. fastFinality=true is
// not supported for v1.x EVM pools and returns an error. tokenRef is unused on EVM (pools are
// keyed by pool address alone) and exists for parity with chain families that need the token
// mint to resolve the read.
func (a *EVMPoolAdapter) GetOnchainInboundRateLimit(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, _ datastore.AddressRef, remoteSelector uint64, fastFinality bool) (tokensapi.RateLimiterConfig, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return tokensapi.RateLimiterConfig{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}
	poolAddr, err := a.EVMTokenBase.ParseNonZeroAddressRef(e.DataStore, poolRef, chainSelector)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, fmt.Errorf("failed to find token pool address for ref (%s): %w", datastore_utils.SprintRef(poolRef), err)
	}
	return a.Ops.GetCurrentInboundRateLimit(e.OperationsBundle, chain, poolAddr, remoteSelector, fastFinality)
}

func (a *EVMPoolAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokensapi.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-pool-adapter:set-token-pool-rate-limits",
		a.Ops.Version(),
		"Set rate limits for a token pool on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.TPRLRemotes) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			tokenPoolAddr, err := a.EVMTokenBase.ParseNonZeroAddressRef(input.ExistingDataStore, input.TokenPoolRef, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for ref (%s): %w", datastore_utils.SprintRef(input.TokenPoolRef), err)
			}

			if input.SkipIfMissingPermissions {
				timelockFltr := datastore.AddressRef{Type: datastore.ContractType(cciputils.RBACTimelock), ChainSelector: chain.Selector, Qualifier: cciputils.CLLQualifier}
				timelockAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, timelockFltr, chain.Selector, datastore_utils_evm.ToNonZeroEVMAddress)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to find timelock address for chain %d: %w", chain.Selector, err)
				}
				poolOwner, rlAdmin, err := a.Ops.GetPoolAdmins(b.GetContext(), &chain, tokenPoolAddr)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool admins for token pool ref (%+v) on chain %d: %w", input.TokenPoolRef, chain.Selector, err)
				}
				isRateLimitAdmin := rlAdmin == timelockAddr || rlAdmin == chain.DeployerKey.From
				isPoolOwner := poolOwner == timelockAddr || poolOwner == chain.DeployerKey.From
				if !isRateLimitAdmin && !isPoolOwner {
					b.Logger.Warnf(
						"Timelock address %q and deployer address %q are not the owner or rate limit admin for token pool at address %q on chain selector %d. Skipping rate limiter config for this chain.",
						timelockAddr.Hex(), chain.DeployerKey.From.Hex(), tokenPoolAddr.Hex(), chain.Selector,
					)
					return sequences.OnChainOutput{}, nil
				}
			}

			output, err := a.Ops.SetRateLimiterConfig(b, chain, tokenPoolAddr, input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
			}
			if len(output) == 0 {
				return sequences.OnChainOutput{}, nil
			}

			batchOp, err := evm_contract.NewBatchOperationFromWrites(output)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			result.BatchOps = append(result.BatchOps, batchOp)
			return result, nil
		},
	)
}

func (a *EVMPoolAdapter) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-pool-adapter:manual-registration",
		a.Ops.Version(),
		"Manually register a token and token pool on EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ManualRegistrationSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tarAddress, err := a.EVMTokenBase.GetTokenAdminRegistryAddress(input.ExistingDataStore, chain.Selector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", chain.Selector, err)
			}

			// Token address resolution strategy:
			// 1. If TokenRef already has an address, use it directly (skip datastore).
			// 2. Otherwise look up the token in the datastore using TokenRef fields.
			// 3. If step 2 fails (e.g. ambiguous or missing), fall back to reading
			//    the token address from the on-chain token pool contract via TokenPoolRef.
			// 4. If none of the above work, return an error.
			tokenRef := input.TokenRef
			if tokenRef.Address == "" {
				if tokRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, tokenRef, chain.Selector, datastore_utils.FullRef); err != nil {
					b.Logger.Warnf("token address could not be resolved using TokenRef (%s): %v", datastore_utils.SprintRef(tokenRef), err)
					b.Logger.Warnf("attempting to resolve token address using TokenPoolRef instead: (%s)", datastore_utils.SprintRef(input.TokenPoolRef))
					tokenPoolAddr, err := a.EVMTokenBase.ParseNonZeroAddressRef(input.ExistingDataStore, input.TokenPoolRef, chain.Selector)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to find token pool address for ref (%s): %w", datastore_utils.SprintRef(input.TokenPoolRef), err)
					}
					tokenAddr, err := a.Ops.GetToken(b, chain, tokenPoolAddr)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool at ref (%s): %w", datastore_utils.SprintRef(input.TokenPoolRef), err)
					}
					tokenRef = datastore.AddressRef{
						ChainSelector: chain.Selector,
						Address:       tokenAddr.Hex(),
					}
				} else {
					tokenRef = tokRef
				}
			}

			tokenAddrBytes, err := a.AddressRefToBytes(tokenRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token address ref to bytes: %w", err)
			}
			tokenAddr := common.BytesToAddress(tokenAddrBytes)
			if tokenAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token address for ref (%+v) is zero address", tokenRef)
			}

			if !common.IsHexAddress(input.ProposedOwner) {
				return sequences.OnChainOutput{}, fmt.Errorf("proposed owner address %q is not a valid hex address", input.ProposedOwner)
			}
			proposedOwnerAddr := common.HexToAddress(input.ProposedOwner)
			if proposedOwnerAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, errors.New("proposed owner address cannot be the zero address")
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(
				b, chains,
				tarseq.ManualRegistrationSequence,
				tarseq.ManualRegistrationSequenceInput{
					AdminAddress:  proposedOwnerAddr,
					ChainSelector: chain.Selector,
					TokenAddress:  tokenAddr,
					Address:       tarAddress,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to manually register token on chain %d: %w", chain.Selector, err)
			}

			return result, nil
		},
	)
}

func (a *EVMPoolAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-pool-adapter:deploy-token-pool-for-token",
		a.Ops.Version(),
		"Deploy a token pool for a token on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			if a.DeployTokenPoolSeq == nil {
				return sequences.OnChainOutput{}, errors.New("DeployTokenPoolSeq is not set on EVMPoolAdapter")
			}

			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			out, err := cldf_ops.ExecuteSequence(b, a.DeployTokenPoolSeq, chains, input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token pool on chain %d: %w", input.ChainSelector, err)
			}

			result.Addresses = append(result.Addresses, out.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)
			if input.TokenRef == nil {
				return sequences.OnChainOutput{}, errors.New("token ref must be provided in input to DeployTokenPoolForToken sequence for EVM pools")
			}

			tokenRef := input.TokenRef.Clone()
			if !datastore_utils.IsAddressRefFullyPopulated(tokenRef) {
				if ref, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, tokenRef, input.ChainSelector, datastore_utils.FullRef); err == nil {
					tokenRef = ref
				} else {
					b.Logger.Warnf("token ref (%s) is not fully populated and could not be resolved in datastore - attempting to resolve ref from on-chain data: %v", datastore_utils.SprintRef(tokenRef), err)
					if tokenRef.Address == "" {
						return sequences.OnChainOutput{}, fmt.Errorf("token ref (%s) is missing address field so on-chain resolution cannot be attempted", datastore_utils.SprintRef(tokenRef))
					}
					fallbackRef, err := a.ResolveTokenRef(b, chains, input.ExistingDataStore, input.ChainSelector, tokenRef.Address)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve token ref from on-chain data for token address %s: %w", tokenRef.Address, err)
					}
					tokenRef = fallbackRef
				}
			}

			var poolRef datastore.AddressRef
			if len(out.Output.Addresses) >= 1 {
				poolRef = out.Output.Addresses[0]
			}

			var writes []evm_contract.WriteOutput
			if !datastore_utils.IsAddressRefEmpty(poolRef) {
				if tokenPoolRolesWrites, err := a.TidyTokenPoolRoles(b, chain, input, poolRef, tokenRef); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to tidy token pool roles: %w", err)
				} else {
					writes = append(writes, tokenPoolRolesWrites...)
				}
				if input.RateLimitAdmin != "" {
					rlAdminHex := input.RateLimitAdmin
					if !common.IsHexAddress(rlAdminHex) {
						return sequences.OnChainOutput{}, fmt.Errorf("rate limit admin address %q is not a valid hex address", input.RateLimitAdmin)
					}
					poolAddr, err := datastore_utils_evm.ToNonZeroEVMAddress(poolRef)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool ref to EVM address for chain %d: %w", input.ChainSelector, err)
					}
					output, err := a.Ops.SetRateLimitAdmin(b, chain, poolAddr, common.HexToAddress(rlAdminHex))
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limit admin: %w", err)
					}
					if len(output) > 0 {
						writes = append(writes, output...)
					}
				}
			}

			if tokenRolesWrites, err := a.TidyTokenRoles(b, chain, input, tokenRef); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to tidy token roles: %w", err)
			} else {
				writes = append(writes, tokenRolesWrites...)
			}

			if len(writes) > 0 {
				batchOp, bErr := evm_contract.NewBatchOperationFromWrites(writes)
				if bErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for token role adjustments: %w", bErr)
				}
				result.BatchOps = append(result.BatchOps, batchOp)
			}

			return result, nil
		},
	)
}

// tidyTokenPoolRoles grants a token pool the token-side roles required for its
// pool type. Burn/mint pools delegate role selection to the registered token
// strategy because token contracts expose different role APIs.
func (a *EVMPoolAdapter) TidyTokenPoolRoles(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input tokensapi.DeployTokenPoolInput,
	tokenPoolRef datastore.AddressRef,
	tokenRef datastore.AddressRef,
) ([]evm_contract.WriteOutput, error) {
	tokenPoolAddr, err := datastore_utils_evm.ToNonZeroEVMAddress(tokenPoolRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert token pool ref to EVM address for chain %d: %w", input.ChainSelector, err)
	}
	tokenAddr, err := datastore_utils_evm.ToNonZeroEVMAddress(tokenRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert token ref to EVM address for chain %d: %w", input.ChainSelector, err)
	}

	if a.IsBurnMintPoolType(input.PoolType) {
		tokenImpl, ok := tokenimpl.Get(deployment.ContractType(tokenRef.Type))
		if !ok {
			b.Logger.Warnf(
				"unsupported token type %q for token at ref (%s); skipping pool role grants for this token on chain %d",
				tokenRef.Type.String(), datastore_utils.SprintRef(tokenRef), input.ChainSelector,
			)
			return nil, nil
		}

		tokenCaps := tokenImpl.Capabilities()
		if !tokenCaps.ParticipatesInPoolRoleGrant {
			b.Logger.Warnf(
				"token type %q has no pool role grant strategy registered, skipping grant for token pool %q on token %q on chain %d",
				tokenRef.Type.String(), tokenPoolAddr.Hex(), input.TokenRef.Qualifier, input.ChainSelector,
			)
			return nil, nil
		}

		if grantWrites, grantErr := tokenImpl.GrantPoolRoles(b, chain, tokenAddr, tokenPoolAddr, common.HexToAddress(input.TimelockAddress)); grantErr != nil {
			return nil, fmt.Errorf("failed to grant pool roles for token type %q (token %q, pool %q) on chain %d: %w", tokenRef.Type, input.TokenRef.Qualifier, tokenPoolAddr.Hex(), input.ChainSelector, grantErr)
		} else {
			return grantWrites, nil
		}
	}

	return nil, nil
}

// tidyTokenRoles will grant timelock admin rights on the token and remove
// the deployer EOA as an admin. If timelock is not found in the datastore
// (i.e. not deployed/not applicable which can be the case in test cases),
// then it leaves the deployer account as an admin so the token isn't left
// without an operator.
func (a *EVMPoolAdapter) TidyTokenRoles(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input tokensapi.DeployTokenPoolInput,
	tokenRef datastore.AddressRef,
) ([]evm_contract.WriteOutput, error) {
	tokenAddr, err := datastore_utils_evm.ToNonZeroEVMAddress(tokenRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert token ref to EVM address for chain %d: %w", input.ChainSelector, err)
	}

	tokenImpl, ok := tokenimpl.Get(deployment.ContractType(tokenRef.Type))
	if !ok {
		b.Logger.Warnf(
			"unsupported token type %q for token at ref (%s); skipping admin role tidy for this token on chain %d",
			tokenRef.Type.String(), datastore_utils.SprintRef(tokenRef), input.ChainSelector,
		)
		return nil, nil
	}

	tokenCaps := tokenImpl.Capabilities()
	if !tokenCaps.SupportsAdminRole {
		b.Logger.Warnf(
			"token type %q does not support admin role management; skipping tidy of token admin roles for token at ref (%s) on chain %d",
			tokenRef.Type.String(), datastore_utils.SprintRef(tokenRef), input.ChainSelector,
		)
		return nil, nil
	}

	timelockAddr, err := a.GetTimelockAddressCLL(input.ExistingDataStore, input.ChainSelector)
	if err != nil {
		b.Logger.Infof("CLL timelock not found for chain %d; keeping deployer as token admin: %s", input.ChainSelector, err.Error())
		return nil, nil
	}

	grantWrites, err := tokenImpl.GrantAdminRole(b, chain, tokenAddr, timelockAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to grant timelock admin role for token %q on chain %d: %w", tokenAddr.Hex(), input.ChainSelector, err)
	}
	revokeWrites, err := tokenImpl.RevokeAdminRole(b, chain, tokenAddr, chain.DeployerKey.From)
	if err != nil {
		return nil, fmt.Errorf("failed to revoke deployer admin role for token %q on chain %d: %w", tokenAddr.Hex(), input.ChainSelector, err)
	}

	return append(grantWrites, revokeWrites...), nil
}
