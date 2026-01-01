package changesets

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployTokenAndPoolCfg struct {
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
	// TokenAddress is the address of the token for which the pool is being deployed.
	TokenAddress common.Address
	// Decimals is the number of decimals used by the token.
	Decimals uint8
	// Router is a reference to the desired router contract.
	// Sometimes we may want to connect to a test router, other times a main router.
	Router datastore.AddressRef
	// Accounts is a map of account addresses to initial mint amounts.
	Accounts map[common.Address]*big.Int
}

func (c DeployTokenAndPoolCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var DeployTokenAndPool = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	tokens.DeployTokenAndPoolInput,
	evm.Chain,
	DeployTokenAndPoolCfg,
]{
	Sequence: tokens.DeployTokenAndPool,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployTokenAndPoolCfg) (tokens.DeployTokenAndPoolInput, error) {
		rmnProxy, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: rmn_proxy.Version,
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return tokens.DeployTokenAndPoolInput{}, fmt.Errorf("failed to resolve rmn proxy ref: %w", err)
		}
		router, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.Router, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return tokens.DeployTokenAndPoolInput{}, fmt.Errorf("failed to resolve router ref: %w", err)
		}

		return tokens.DeployTokenAndPoolInput{
			Accounts: cfg.Accounts,
			DeployTokenPoolInput: tokens.DeployTokenPoolInput{
				ChainSel:         cfg.ChainSel,
				TokenPoolType:    cfg.TokenPoolType,
				TokenPoolVersion: cfg.TokenPoolVersion,
				TokenSymbol:      cfg.TokenSymbol,
				RateLimitAdmin:   cfg.RateLimitAdmin,
				ConstructorArgs: tokens.ConstructorArgs{
					Token:    cfg.TokenAddress,
					Decimals: cfg.Decimals,
					RMNProxy: rmnProxy,
					Router:   router,
				},
			},
		}, nil
	},
	ResolveDep: evm_seq.ResolveEVMChainDep[DeployTokenAndPoolCfg],
})
