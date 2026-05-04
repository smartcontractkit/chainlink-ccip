package cctp_message_transmitter_proxy

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// AllowedCallerConfigArgs is an element of configureAllowedCallers.
type AllowedCallerConfigArgs = gobindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs

var ConfigureAllowedCallers = contract.NewWrite(contract.WriteParams[[]gobindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs, *gobindings.CCTPMessageTransmitterProxy]{
	Name:            "cctp-message-transmitter-proxy:configure-allowed-callers",
	Version:         Version,
	Description:     "Calls configureAllowedCallers on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CCTPMessageTransmitterProxyMetaData.ABI,
	NewContract:     gobindings.NewCCTPMessageTransmitterProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CCTPMessageTransmitterProxy, []gobindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs],
	Validate:        func([]gobindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs) error { return nil },
	CallContract: func(c *gobindings.CCTPMessageTransmitterProxy, opts *bind.TransactOpts, args []gobindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs) (*types.Transaction, error) {
		return c.ConfigureAllowedCallers(opts, args)
	},
})
