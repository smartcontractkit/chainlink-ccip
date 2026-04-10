package tip20

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	ContractType   deployment.ContractType = "TIP20Token"
	Version                                = semver.MustParse("1.0.0")
	TypeAndVersion                         = deployment.NewTypeAndVersion(ContractType, *Version)
)

// NOTE: these addresses are the same on both Tempo testnet and mainnet, so we don't need separate constants per environment.
const (
	TokenFactoryAddress = "0x20Fc000000000000000000000000000000000000"
	TokenThetaUSD       = "0x20c0000000000000000000000000000000000003"
	TokenBetaUSD        = "0x20c0000000000000000000000000000000000002"
	TokenAlphaUSD       = "0x20c0000000000000000000000000000000000001"
	TokenPathUSD        = "0x20c0000000000000000000000000000000000000"
)

// NOTE: the gas station app in slack uses PathUSD, so we use it as the default for ease of use.
// If we want to deploy a token with currency USD, then its quote token must also be USD.
const DefaultQuoteToken = TokenPathUSD

// NOTE: we chose USD as the default currency since most of the well-known TIP20 tokens on Tempo
// use USD as their currency (e.g. ThetaUSD, BetaUSD, AlphaUSD, PathUSD).
const DefaultCurrency = "USD"

type FactoryDeployArgs struct {
	Currency   string         // The token currency. Defaults to USD if not provided.
	Symbol     string         // The token symbol. This is a required input.
	Name       string         // The token name. This is a required input.
	QuoteToken common.Address // Address of a pre-existing TIP20 token to use as the quote token. Defaults to PathUSD if not provided.
	Admin      common.Address // The token admin. Defaults to the deployer address if not provided.
	Salt       [32]byte       // Optional salt for deterministic deployment. Defaults to a random salt if not provided.
}

// Deploy deploys the TIP20 token contract with the provided deploy arguments. The TIP20 token is ERC20 compliant and includes additional
// features as defined in the TIP20 standard: https://www.mintlify.com/tempoxyz/tempo/protocol/tip20/overview#erc-20-compatibility. This
// sequence is only applicable for Tempo testnet / mainnet. The token is deployed via the factory contract as recommended in the docs. We
// use sensible defaults for QuoteToken, Currency, Admin, and Salt to reduce the configuration burden on the user when deploying a TIP20
// token.
//
//	Factory Contract: https://github.com/tempoxyz/tempo/blob/a20e2e46c7cba6164ef95c91bf83d5fc614750f3/tips/ref-impls/src/TIP20Factory.sol#L1
//	Token Contract: https://github.com/tempoxyz/tempo/blob/a20e2e46c7cba6164ef95c91bf83d5fc614750f3/tips/ref-impls/src/TIP20.sol#L1
//	Docs: https://www.mintlify.com/tempoxyz/tempo/protocol/tip20/overview
var Deploy = operations.NewSequence(
	"tip20:deploy",
	Version,
	"Deploys a TIP20 token via the TIP20 factory. Only applicable for Tempo testnet / mainnet.",
	func(b operations.Bundle, chain evm.Chain, input FactoryDeployArgs) (sequences.OnChainOutput, error) {
		isTempoTestnet := chainsel.TEMPO_TESTNET_MODERATO.Selector == chain.Selector || chainsel.TEMPO_TESTNET.Selector == chain.Selector
		isTempoMainnet := chainsel.TEMPO_MAINNET.Selector == chain.Selector
		if !isTempoTestnet && !isTempoMainnet {
			return sequences.OnChainOutput{}, errors.New("TIP20 token deployment is only supported on Tempo testnet and mainnet")
		}

		factoryAddr := common.HexToAddress(TokenFactoryAddress)
		deployerKey := chain.DeployerKey.From
		if input.Symbol == "" {
			return sequences.OnChainOutput{}, errors.New("symbol is required")
		}
		if input.Name == "" {
			return sequences.OnChainOutput{}, errors.New("name is required")
		}
		if input.QuoteToken == (common.Address{}) {
			input.QuoteToken = common.HexToAddress(DefaultQuoteToken)
		}
		if input.Currency == "" {
			input.Currency = DefaultCurrency
		}
		if input.Admin == (common.Address{}) {
			input.Admin = deployerKey
		}
		if input.Salt == [32]byte{} {
			if salt, err := generateValidSalt(b, chain, factoryAddr, deployerKey); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to produce a valid salt for token deployment: %w", err)
			} else {
				input.Salt = salt
			}
		}

		isQuoteTokenValid, err := operations.ExecuteOperation(b, IsTIP20, chain, contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       factoryAddr,
			Args:          input.QuoteToken,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("isTIP20 quote token: %w", err)
		}
		if !isQuoteTokenValid.Output {
			return sequences.OnChainOutput{}, errors.New("quoteToken must be a valid TIP-20 token address")
		}

		createTokenReport, err := operations.ExecuteOperation(b, CreateToken, chain, contract.FunctionInput[CreateTokenArgs]{
			ChainSelector: chain.Selector,
			Address:       factoryAddr,
			Args: CreateTokenArgs{
				QuoteToken: input.QuoteToken,
				Currency:   input.Currency,
				Symbol:     input.Symbol,
				Admin:      input.Admin,
				Name:       input.Name,
				Salt:       input.Salt,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("createToken: %w", err)
		}

		tokenAddrReport, err := operations.ExecuteOperation(b, GetTokenAddress, chain, contract.FunctionInput[GetTokenAddressArgs]{
			ChainSelector: chain.Selector,
			Address:       factoryAddr,
			Args: GetTokenAddressArgs{
				Sender: deployerKey,
				Salt:   input.Salt,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("getTokenAddress after deploy: %w", err)
		}

		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{createTokenReport.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("batch operation: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{
				{
					ChainSelector: chain.Selector,
					Address:       tokenAddrReport.Output.Hex(),
					Qualifier:     input.Symbol,
					Type:          datastore.ContractType(ContractType),
					Version:       Version,
				},
			},
			BatchOps: []mcms_types.BatchOperation{
				batchOp,
			},
		}, nil
	},
)
