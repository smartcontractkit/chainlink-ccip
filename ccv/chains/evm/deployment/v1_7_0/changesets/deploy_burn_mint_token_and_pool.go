package changesets

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployBurnMintTokenAndPoolCfg struct {
	DeployTokenPoolCfg
	// Accounts is a map of account addresses to initial mint amounts.
	Accounts map[common.Address]*big.Int
	// TokenInfo is the information about the token to be deployed.
	// Token symbol will be taken from DeployTokenPoolInput.TokenSymbol.
	TokenInfo tokens.TokenInfo
}

var DeployBurnMintTokenAndPool = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	tokens.DeployBurnMintTokenAndPoolInput,
	evm.Chain,
	DeployBurnMintTokenAndPoolCfg,
]{
	Sequence: tokens.DeployBurnMintTokenAndPool,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployBurnMintTokenAndPoolCfg) (tokens.DeployBurnMintTokenAndPoolInput, error) {
		rmnProxy, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: semver.MustParse("1.0.0"),
		}, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return tokens.DeployBurnMintTokenAndPoolInput{}, fmt.Errorf("failed to resolve rmn proxy ref: %w", err)
		}
		router, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.Router, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return tokens.DeployBurnMintTokenAndPoolInput{}, fmt.Errorf("failed to resolve router ref: %w", err)
		}

		return tokens.DeployBurnMintTokenAndPoolInput{
			Accounts:  cfg.Accounts,
			TokenInfo: cfg.TokenInfo,
			DeployTokenPoolInput: tokens.DeployTokenPoolInput{
				ChainSel:         cfg.ChainSel,
				TokenPoolType:    cfg.TokenPoolType,
				TokenPoolVersion: cfg.TokenPoolVersion,
				TokenSymbol:      cfg.TokenSymbol,
				RateLimitAdmin:   cfg.RateLimitAdmin,
				ConstructorArgs: token_pool.ConstructorArgs{
					Token:              cfg.TokenAddress,
					LocalTokenDecimals: cfg.LocalTokenDecimals,
					Allowlist:          cfg.Allowlist,
					RMNProxy:           rmnProxy,
					Router:             router,
				},
			},
		}, nil
	},
	ResolveDep: evm_seq.ResolveEVMChainDep[DeployBurnMintTokenAndPoolCfg],
})
