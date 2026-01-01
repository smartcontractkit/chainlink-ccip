package cctp_message_transmitter_proxy

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCTPMessageTransmitterProxy"

var Version = semver.MustParse("1.6.2")

type ConstructorArgs struct {
	TokenMessenger common.Address
}

type AuthorizedCallerArgs = cctp_message_transmitter_proxy.AuthorizedCallersAuthorizedCallerArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "cctp-message-transmitter-proxy:deploy",
	Version:          Version,
	Description:      "Deploys the CCTPMessageTransmitterProxy contract",
	ContractMetadata: cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy]{
	Name:            "cctp-message-transmitter-proxy:apply-authorized-caller-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies authorized caller updates on the CCTPMessageTransmitterProxy",
	ContractType:    ContractType,
	ContractABI:     cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyABI,
	NewContract:     cctp_message_transmitter_proxy.NewCCTPMessageTransmitterProxy,
	IsAllowedCaller: contract.OnlyOwner[*cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy, AuthorizedCallerArgs],
	Validate: func(proxy *cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy, backend bind.ContractBackend, opts *bind.CallOpts, args AuthorizedCallerArgs) error {
		for _, caller := range args.AddedCallers {
			if caller == (common.Address{}) {
				return errors.New("caller cannot be the zero address")
			}
		}

		return nil
	},
	IsNoop: func(proxy *cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy, opts *bind.CallOpts, args AuthorizedCallerArgs) (bool, error) {
		allowedCallers, err := proxy.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get all authorized callers: %w", err)
		}
		for _, caller := range args.AddedCallers {
			if !slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		for _, caller := range args.RemovedCallers {
			if slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		return true, nil
	},
	CallContract: func(proxy *cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return proxy.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
