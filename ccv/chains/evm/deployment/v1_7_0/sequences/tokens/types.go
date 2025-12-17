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
	// LocalTokenDecimals is the number of decimals used by the token.
	LocalTokenDecimals uint8
	// Allowlist is the list of addresses allowed to transfer the token.
	Allowlist []common.Address
	// RMNProxy is the RMN proxy contract.
	RMNProxy common.Address
	// Router is the router contract.
	Router common.Address
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
	TokenSymbol string
	// RateLimitAdmin is an additional address allowed to set rate limiters.
	// If left empty, setRateLimitAdmin will not be attempted.
	RateLimitAdmin common.Address
	// ThresholdAmountForAdditionalCCVs is the transfer amount above which additional CCVs are required.
	ThresholdAmountForAdditionalCCVs *big.Int
	// ConstructorArgs are the constructor arguments for the token pool.
	ConstructorArgs ConstructorArgs
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

	return nil
}
