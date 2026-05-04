package lombard_verifier

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// Type aliases for LombardVerifier deploy and token updates.
type (
	RemoteChainConfigArgs = gobindings.BaseVerifierRemoteChainConfigArgs
	RemoteAdapterArgs       = gobindings.LombardVerifierRemoteAdapterArgs
	DynamicConfig           = gobindings.LombardVerifierDynamicConfig
	SupportedTokenArgs      = gobindings.LombardVerifierSupportedTokenArgs
)

var UpdateSupportedTokens = contract.NewWrite(contract.WriteParams[UpdateSupportedTokensArgs, *gobindings.LombardVerifier]{
	Name:            "lombard-verifier:update-supported-tokens",
	Version:         Version,
	Description:     "Calls updateSupportedTokens on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LombardVerifierMetaData.ABI,
	NewContract:     gobindings.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LombardVerifier, UpdateSupportedTokensArgs],
	Validate:        func(UpdateSupportedTokensArgs) error { return nil },
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.TransactOpts, args UpdateSupportedTokensArgs) (*types.Transaction, error) {
		return c.UpdateSupportedTokens(opts, args.TokensToRemove, args.TokensToSet)
	},
})

var VersionTag = contract.NewRead(contract.ReadParams[struct{}, [4]byte, *gobindings.LombardVerifier]{
	Name:         "lombard-verifier:version-tag",
	Version:      Version,
	Description:  "Calls versionTag on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewLombardVerifier,
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.CallOpts, args struct{}) ([4]byte, error) {
		return c.VersionTag(opts)
	},
})

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.BaseVerifierRemoteChainConfigArgs, *gobindings.LombardVerifier]{
	Name:            "lombard-verifier:apply-remote-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyRemoteChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LombardVerifierMetaData.ABI,
	NewContract:     gobindings.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LombardVerifier, []gobindings.BaseVerifierRemoteChainConfigArgs],
	Validate:        func([]gobindings.BaseVerifierRemoteChainConfigArgs) error { return nil },
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.TransactOpts, args []gobindings.BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var SetRemoteAdapters = contract.NewWrite(contract.WriteParams[[]gobindings.LombardVerifierRemoteAdapterArgs, *gobindings.LombardVerifier]{
	Name:            "lombard-verifier:set-remote-adapters",
	Version:         Version,
	Description:     "Calls setRemoteAdapters on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LombardVerifierMetaData.ABI,
	NewContract:     gobindings.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LombardVerifier, []gobindings.LombardVerifierRemoteAdapterArgs],
	Validate:        func([]gobindings.LombardVerifierRemoteAdapterArgs) error { return nil },
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.TransactOpts, args []gobindings.LombardVerifierRemoteAdapterArgs) (*types.Transaction, error) {
		return c.SetRemoteAdapters(opts, args)
	},
})

var SetPath = contract.NewWrite(contract.WriteParams[SetPathArgs, *gobindings.LombardVerifier]{
	Name:            "lombard-verifier:set-path",
	Version:         Version,
	Description:     "Calls setPath on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LombardVerifierMetaData.ABI,
	NewContract:     gobindings.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LombardVerifier, SetPathArgs],
	Validate:        func(SetPathArgs) error { return nil },
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.TransactOpts, args SetPathArgs) (*types.Transaction, error) {
		return c.SetPath(opts, args.RemoteChainSelector, args.LChainId, args.AllowedCaller)
	},
})

var GetPath = contract.NewRead(contract.ReadParams[uint64, gobindings.LombardVerifierPath, *gobindings.LombardVerifier]{
	Name:         "lombard-verifier:get-path",
	Version:      Version,
	Description:  "Calls getPath on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewLombardVerifier,
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.CallOpts, args uint64) (gobindings.LombardVerifierPath, error) {
		return c.GetPath(opts, args)
	},
})

var GetRemoteAdapter = contract.NewRead(contract.ReadParams[GetRemoteAdapterArgs, [32]byte, *gobindings.LombardVerifier]{
	Name:         "lombard-verifier:get-remote-adapter",
	Version:      Version,
	Description:  "Calls getRemoteAdapter on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewLombardVerifier,
	CallContract: func(c *gobindings.LombardVerifier, opts *bind.CallOpts, args GetRemoteAdapterArgs) ([32]byte, error) {
		return c.GetRemoteAdapter(opts, args.RemoteChainSelector, args.Token)
	},
})
