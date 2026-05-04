package cctp_verifier

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// Type aliases for CCTP lane configuration and deploy inputs.
type (
	SetDomainArgs         = gobindings.CCTPVerifierSetDomainArgs
	RemoteChainConfigArgs = gobindings.BaseVerifierRemoteChainConfigArgs
	DynamicConfig         = gobindings.CCTPVerifierDynamicConfig
	BaseVerifierArgs      = gobindings.CCTPVerifierBaseVerifierArgs
)

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.BaseVerifierRemoteChainConfigArgs, *gobindings.CCTPVerifier]{
	Name:            "cctp-verifier:apply-remote-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyRemoteChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CCTPVerifierMetaData.ABI,
	NewContract:     gobindings.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CCTPVerifier, []gobindings.BaseVerifierRemoteChainConfigArgs],
	Validate:        func([]gobindings.BaseVerifierRemoteChainConfigArgs) error { return nil },
	CallContract: func(c *gobindings.CCTPVerifier, opts *bind.TransactOpts, args []gobindings.BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var SetDomains = contract.NewWrite(contract.WriteParams[[]gobindings.CCTPVerifierSetDomainArgs, *gobindings.CCTPVerifier]{
	Name:            "cctp-verifier:set-domains",
	Version:         Version,
	Description:     "Calls setDomains on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CCTPVerifierMetaData.ABI,
	NewContract:     gobindings.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CCTPVerifier, []gobindings.CCTPVerifierSetDomainArgs],
	Validate:        func([]gobindings.CCTPVerifierSetDomainArgs) error { return nil },
	CallContract: func(c *gobindings.CCTPVerifier, opts *bind.TransactOpts, args []gobindings.CCTPVerifierSetDomainArgs) (*types.Transaction, error) {
		return c.SetDomains(opts, args)
	},
})

var GetAllowedFinalityConfig = contract.NewRead(contract.ReadParams[struct{}, [4]byte, *gobindings.CCTPVerifier]{
	Name:         "cctp-verifier:get-allowed-finality-config",
	Version:      Version,
	Description:  "Calls getAllowedFinalityConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCCTPVerifier,
	CallContract: func(c *gobindings.CCTPVerifier, opts *bind.CallOpts, args struct{}) ([4]byte, error) {
		return c.GetAllowedFinalityConfig(opts)
	},
})

var SetAllowedFinalityConfig = contract.NewWrite(contract.WriteParams[[4]byte, *gobindings.CCTPVerifier]{
	Name:            "cctp-verifier:set-allowed-finality-config",
	Version:         Version,
	Description:     "Calls setAllowedFinalityConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CCTPVerifierMetaData.ABI,
	NewContract:     gobindings.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CCTPVerifier, [4]byte],
	Validate:        func([4]byte) error { return nil },
	CallContract: func(c *gobindings.CCTPVerifier, opts *bind.TransactOpts, args [4]byte) (*types.Transaction, error) {
		return c.SetAllowedFinalityConfig(opts, args)
	},
})

var VersionTag = contract.NewRead(contract.ReadParams[struct{}, [4]byte, *gobindings.CCTPVerifier]{
	Name:         "cctp-verifier:version-tag",
	Version:      Version,
	Description:  "Calls versionTag on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCCTPVerifier,
	CallContract: func(c *gobindings.CCTPVerifier, opts *bind.CallOpts, args struct{}) ([4]byte, error) {
		return c.VersionTag(opts)
	},
})
