package tokens

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type ConstructorArgs struct {
	// Token is the token for which the pool is being deployed.
	Token common.Address
	// Decimals is the number of decimals used by the token.
	Decimals uint8
	// RMNProxy is the RMN proxy contract.
	RMNProxy common.Address
	// Router is the router contract.
	Router common.Address
	// BurnAddress is the burn address for BurnToAddressMintTokenPool. Optional for other pool types.
	BurnAddress common.Address
}

// AdvancedPoolHooksConfig contains optional configuration for AdvancedPoolHooks.
type AdvancedPoolHooksConfig struct {
	// Allowlist is a list of addresses allowed to trigger lockOrBurn.
	// Empty list means allowlist is disabled.
	Allowlist []common.Address
	// PolicyEngine is the policy engine address. Zero address disables policy checks.
	PolicyEngine common.Address
	// AuthorizedCallers specifies the set of callers authorized to invoke preflightCheck/postflightCheck.
	// AuthorizedCallers check is always enforced, regardless of whether the list is empty.
	AuthorizedCallers []common.Address
}

// DeployTokenPoolInput is the input for the DeployTokenPool sequence.
type DeployTokenPoolInput struct {
	// ChainSel is the chain selector for the chain being configured.
	ChainSel uint64
	// TokenPoolType is the type of the token pool to deploy.
	TokenPoolType datastore.ContractType
	// TokenPoolVersion is the version of the token pool to deploy.
	TokenPoolVersion *semver.Version

	// TokenSymbol is the symbol of the token to be configured.
	// This symbol will be stored in the returned AddressRef.
	//
	// TODO: this field is not named correctly - it should be renamed to `TokenPoolQualifier`.
	//
	TokenSymbol string

	// RateLimitAdmin is an additional address allowed to set rate limiters.
	// If left empty, setRateLimitAdmin will not be attempted.
	RateLimitAdmin common.Address
	// ThresholdAmountForAdditionalCCVs is the transfer amount above which additional CCVs are required.
	ThresholdAmountForAdditionalCCVs *big.Int
	// FeeAdmin is an additional address (besides the pool owner) allowed to call
	// withdrawFeeTokens on the pool.
	FeeAdmin common.Address
	// ConstructorArgs are the constructor arguments for the token pool.
	ConstructorArgs ConstructorArgs
	// AdvancedPoolHooksConfig contains optional configuration for AdvancedPoolHooks.
	AdvancedPoolHooksConfig AdvancedPoolHooksConfig
	// LockBoxGroups declares the liquidity topology for a SiloedLockReleaseTokenPool: each group is a
	// set of remote chain selectors sharing one ERC20LockBox, so chains in different groups have
	// isolated liquidity. Required for SiloedLockReleaseTokenPool, ignored by all other pool types.
	LockBoxGroups [][]uint64
}

// ValidateLockBoxGroups checks that the declared silo topology is well formed: at least one group,
// no empty groups, and no chain in more than one group (which would make the lockbox mapping
// ambiguous, since configureLockBoxes is last-write-wins per chain).
func (c DeployTokenPoolInput) ValidateLockBoxGroups() error {
	if len(c.LockBoxGroups) == 0 {
		return fmt.Errorf("lock box groups must be defined for pool type %s", c.TokenPoolType)
	}
	seen := make(map[uint64]int, len(c.LockBoxGroups))
	for i, group := range c.LockBoxGroups {
		if len(group) == 0 {
			return fmt.Errorf("lock box group %d is empty", i)
		}
		for _, sel := range group {
			if prev, dup := seen[sel]; dup {
				return fmt.Errorf("remote chain selector %d appears in lock box groups %d and %d", sel, prev, i)
			}
			seen[sel] = i
		}
	}

	return nil
}

func (c DeployTokenPoolInput) ChainSelector() uint64 {
	return c.ChainSel
}

func (c DeployTokenPoolInput) Validate(chain evm.Chain) error {
	if c.ChainSel != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSel, chain)
	}
	if c.TokenSymbol == "" {
		return errors.New("token symbol must be defined")
	}
	if c.TokenPoolType == "" {
		return errors.New("token pool type must be defined")
	}
	if c.TokenPoolVersion == nil {
		return errors.New("token pool version must be defined")
	}
	if c.ConstructorArgs.Token == (common.Address{}) {
		return errors.New("token address must be defined")
	}
	if c.ConstructorArgs.RMNProxy == (common.Address{}) {
		return errors.New("rmn proxy address must be defined")
	}
	if c.ConstructorArgs.Router == (common.Address{}) {
		return errors.New("router address must be defined")
	}
	if c.ThresholdAmountForAdditionalCCVs == nil {
		return errors.New("threshold amount for additional ccvs must be defined")
	}
	// Fee aggregator can be zero address; it's optional

	return nil
}
