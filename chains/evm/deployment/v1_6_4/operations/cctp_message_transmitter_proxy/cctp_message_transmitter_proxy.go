package cctp_message_transmitter_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/cctp_message_transmitter_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCTPMessageTransmitterProxy"
var Version *semver.Version = semver.MustParse("1.6.4")

type ConstructorArgs struct {
	TokenMessenger common.Address
}

type AllowedCallerConfigArgs = cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyAllowedCallerConfigArgs

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

var CCTPMessageTransmitterProxyConfigureAllowedCallers = contract.NewWrite(contract.WriteParams[[]AllowedCallerConfigArgs, *cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy]{
	Name:            "cctp-message-transmitter-proxy:configure-allowed-callers",
	Version:         Version,
	Description:     "Configures the allowed callers for the CCTPMessageTransmitterProxy contract",
	ContractType:    ContractType,
	ContractABI:     cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyABI,
	NewContract:     cctp_message_transmitter_proxy.NewCCTPMessageTransmitterProxy,
	IsAllowedCaller: contract.OnlyOwner[*cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy, []AllowedCallerConfigArgs],
	Validate:        func(args []AllowedCallerConfigArgs) error { return nil },
	CallContract: func(cctpMessageTransmitterProxy *cctp_message_transmitter_proxy.CCTPMessageTransmitterProxy, opts *bind.TransactOpts, args []AllowedCallerConfigArgs) (*types.Transaction, error) {
		return cctpMessageTransmitterProxy.ConfigureAllowedCallers(opts, []cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyAllowedCallerConfigArgs(args))
	},
})
