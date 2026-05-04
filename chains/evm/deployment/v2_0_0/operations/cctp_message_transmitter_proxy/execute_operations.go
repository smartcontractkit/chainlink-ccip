package cctp_message_transmitter_proxy

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// AuthorizedCallerArgs matches applyAuthorizedCallerUpdates input.
type AuthorizedCallerArgs = gobindings.AuthorizedCallersAuthorizedCallerArgs

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[gobindings.AuthorizedCallersAuthorizedCallerArgs, *gobindings.CCTPMessageTransmitterProxy]{
	Name:            "cctp-message-transmitter-proxy:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CCTPMessageTransmitterProxyMetaData.ABI,
	NewContract:     gobindings.NewCCTPMessageTransmitterProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CCTPMessageTransmitterProxy, gobindings.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(gobindings.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(c *gobindings.CCTPMessageTransmitterProxy, opts *bind.TransactOpts, args gobindings.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *gobindings.CCTPMessageTransmitterProxy]{
	Name:         "cctp-message-transmitter-proxy:get-all-authorized-callers",
	Version:      Version,
	Description:  "Calls getAllAuthorizedCallers on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCCTPMessageTransmitterProxy,
	CallContract: func(c *gobindings.CCTPMessageTransmitterProxy, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
		return c.GetAllAuthorizedCallers(opts)
	},
})
