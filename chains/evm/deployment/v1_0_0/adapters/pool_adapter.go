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
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
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
	GetTokenDecimals(ctx context.Context, chain evm.Chain, poolAddr common.Address) (uint8, error)
	GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error)
	SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error)
	SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) (evm_contract.WriteOutput, error)
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

	// If the ref has the pool address already, then skip the datastore lookup altogether
	if poolRef.Address != "" {
		tokenPoolAddr := poolRef.Address
		if !common.IsHexAddress(tokenPoolAddr) {
			return "", fmt.Errorf("token pool address %q in ref is not a valid hex address", tokenPoolAddr)
		}
		tokenAddr, err := a.Ops.GetToken(e.OperationsBundle, chain, common.HexToAddress(tokenPoolAddr))
		if err != nil {
			return "", fmt.Errorf("failed to get token address from token pool (%s): %w", datastore_utils.SprintRef(poolRef), err)
		}
		return tokenAddr.Hex(), nil
	}

	// If the pool address isn't in the ref, then look it up in the datastore
	tokenPoolAddrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, poolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return "", fmt.Errorf("failed to find token pool in datastore using ref (%+v): %w", poolRef, err)
	}
	tokenPoolAddrBytes, err := a.AddressRefToBytes(tokenPoolAddrRef)
	if err != nil {
		return "", fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}
	tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
	if tokenPoolAddr == (common.Address{}) {
		return "", errors.New("token pool address is zero address")
	}
	tokenAddr, err := a.Ops.GetToken(e.OperationsBundle, chain, tokenPoolAddr)
	if err != nil {
		return "", fmt.Errorf("failed to get token address from token pool ref (%+v): %w", tokenPoolAddrRef, err)
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
	// go straight to the pool contract for decimals.
	if poolRef.Address != "" {
		if !common.IsHexAddress(poolRef.Address) {
			return 0, fmt.Errorf("token pool address %q in ref is not a valid hex address", poolRef.Address)
		} else {
			return a.Ops.GetTokenDecimals(e.GetContext(), chain, common.HexToAddress(poolRef.Address))
		}
	}

	// If the ref doesn't have the pool address, then we need to hit the datastore for
	// the full pool ref, then get the decimals from the pool contract.
	poolAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, poolRef, chainSelector, datastore_utils_evm.ToEVMAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to find token pool address for ref (%s): %w", datastore_utils.SprintRef(poolRef), err)
	} else {
		return a.Ops.GetTokenDecimals(e.GetContext(), chain, poolAddr)
	}
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

			tokenPoolAddrBytes, err := a.AddressRefToBytes(input.TokenPoolRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", err)
			}
			tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
			if tokenPoolAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref (%+v) is zero address", input.TokenPoolRef)
			}

			if input.SkipIfMissingPermissions {
				timelockFltr := datastore.AddressRef{Type: datastore.ContractType(cciputils.RBACTimelock), ChainSelector: chain.Selector, Qualifier: cciputils.CLLQualifier}
				timelockAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, timelockFltr, chain.Selector, datastore_utils_evm.ToEVMAddress)
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
		})
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

			tarAddress, err := GetTokenAdminRegistryAddress(input.ExistingDataStore, chain.Selector, &a.EVMTokenBase)
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
					b.Logger.Warnf("token address could not be resolved using TokenRef (%+v): %v", tokenRef, err)
					b.Logger.Warnf("attempting to resolve token address using TokenPoolRef instead: (%+v)", input.TokenPoolRef)

					tokenPoolRef, poolErr := datastore_utils.FindAndFormatRef(input.ExistingDataStore, input.TokenPoolRef, chain.Selector, datastore_utils.FullRef)
					if poolErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("token pool could not be resolved using TokenPoolRef (%+v): %w", input.TokenPoolRef, poolErr)
					}
					tokenPoolAddrBytes, addrErr := a.AddressRefToBytes(tokenPoolRef)
					if addrErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool address ref to bytes: %w", addrErr)
					}
					tokenPoolAddr := common.BytesToAddress(tokenPoolAddrBytes)
					if tokenPoolAddr == (common.Address{}) {
						return sequences.OnChainOutput{}, fmt.Errorf("token pool address for ref (%+v) is zero address", tokenPoolRef)
					}
					tokenAddr, getErr := a.Ops.GetToken(b, chain, tokenPoolAddr)
					if getErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool ref (%+v): %w", tokenPoolRef, getErr)
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
			result, err = sequences.RunAndMergeSequence(b, chains,
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
		})
}

func (a *EVMPoolAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-pool-adapter:deploy-token-pool-for-token",
		a.Ops.Version(),
		"Deploy a token pool for a token on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			var writes []evm_contract.WriteOutput

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

			var result sequences.OnChainOutput
			result.Addresses = append(result.Addresses, out.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)

			toknFilterDS := datastore.AddressRef{ChainSelector: input.ChainSelector}
			if input.TokenRef.Address != "" {
				toknFilterDS.Address = input.TokenRef.Address
			}
			if input.TokenRef.Qualifier != "" {
				toknFilterDS.Qualifier = input.TokenRef.Qualifier
			}
			if input.TokenRef.Type != "" {
				toknFilterDS.Type = input.TokenRef.Type
			}

			toknRef, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, toknFilterDS, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find token address for symbol %q on chain %d: %w", input.TokenRef.Qualifier, input.ChainSelector, err)
			}
			toknAddr, err := datastore_utils_evm.ToEVMAddress(toknRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token ref to EVM address for chain %d: %w", input.ChainSelector, err)
			}
			if toknAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token address for symbol %q is zero address", input.TokenRef.Qualifier)
			}

			var poolRef datastore.AddressRef
			if len(out.Output.Addresses) >= 1 {
				poolRef = out.Output.Addresses[0]
			}

			if !datastore_utils.IsAddressRefEmpty(poolRef) {
				if tokenPoolRolesWrites, err := tidyTokenPoolRoles(b, chain, input, poolRef, toknRef); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to tidy token pool roles: %w", err)
				} else {
					writes = append(writes, tokenPoolRolesWrites...)
				}
				if input.RateLimitAdmin != "" {
					rlAdminHex := input.RateLimitAdmin
					if !common.IsHexAddress(rlAdminHex) {
						return sequences.OnChainOutput{}, fmt.Errorf("rate limit admin address %q is not a valid hex address", input.RateLimitAdmin)
					}
					rlAdminAddr := common.HexToAddress(rlAdminHex)
					if rlAdminAddr == (common.Address{}) {
						return sequences.OnChainOutput{}, errors.New("rate limit admin address cannot be the zero address")
					}
					poolAddr, err := datastore_utils_evm.ToEVMAddress(poolRef)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to convert token pool ref to EVM address for chain %d: %w", input.ChainSelector, err)
					}
					if poolAddr == (common.Address{}) {
						return sequences.OnChainOutput{}, errors.New("deployed token pool address cannot be the zero address")
					}
					output, err := a.Ops.SetRateLimitAdmin(b, chain, poolAddr, rlAdminAddr)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limit admin: %w", err)
					}
					writes = append(writes, output)
				}
			}

			if tokenRolesWrites, err := tidyTokenRoles(b, chain, input, toknRef); err != nil {
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
func tidyTokenPoolRoles(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input tokensapi.DeployTokenPoolInput,
	poolRef datastore.AddressRef,
	tokenRef datastore.AddressRef,
) ([]evm_contract.WriteOutput, error) {
	tokenAddr, err := datastore_utils_evm.ToEVMAddress(tokenRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert token ref to EVM address for chain %d: %w", input.ChainSelector, err)
	}
	poolAddress, err := datastore_utils_evm.ToEVMAddress(poolRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert token pool ref to EVM address for chain %d: %w", input.ChainSelector, err)
	}

	if input.PoolType == cciputils.BurnMintTokenPool.String() {
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
				tokenRef.Type.String(), poolAddress.Hex(), input.TokenRef.Qualifier, input.ChainSelector,
			)
			return nil, nil
		}

		if grantWrites, grantErr := tokenImpl.GrantPoolRoles(b, chain, tokenAddr, poolAddress, common.HexToAddress(input.TimelockAddress)); grantErr != nil {
			return nil, fmt.Errorf("failed to grant pool roles for token type %q (token %q, pool %q) on chain %d: %w", tokenRef.Type, input.TokenRef.Qualifier, poolAddress.Hex(), input.ChainSelector, grantErr)
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
func tidyTokenRoles(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input tokensapi.DeployTokenPoolInput,
	tokenRef datastore.AddressRef,
) ([]evm_contract.WriteOutput, error) {
	tokenAddr, err := datastore_utils_evm.ToEVMAddress(tokenRef)
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

	timelockRef := datastore_utils.GetAddressRef(
		input.ExistingDataStore.Addresses().Filter(),
		input.ChainSelector,
		cciputils.RBACTimelock,
		cciputils.Version_1_0_0,
		cciputils.CLLQualifier,
	)
	if datastore_utils.IsAddressRefEmpty(timelockRef) {
		b.Logger.Infof("CLL timelock not found for chain %d; keeping deployer as token admin", input.ChainSelector)
		return nil, nil
	}
	timelockAddr, err := datastore_utils_evm.ToEVMAddress(timelockRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert timelock ref to EVM address for chain %d: %w", input.ChainSelector, err)
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

// GetTokenAdminRegistryAddress looks up the TAR (v1.5.0) address from the datastore.
func GetTokenAdminRegistryAddress(ds datastore.DataStore, selector uint64, base *EVMTokenBase) (common.Address, error) {
	filters := datastore.AddressRef{
		Type:          datastore.ContractType(tarops.ContractType),
		ChainSelector: selector,
		Version:       tarops.Version,
	}
	ref, err := datastore_utils.FindAndFormatRef(ds, filters, selector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token admin registry address on chain %d: %w", selector, err)
	}
	addr, err := base.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}
	return common.BytesToAddress(addr), nil
}
