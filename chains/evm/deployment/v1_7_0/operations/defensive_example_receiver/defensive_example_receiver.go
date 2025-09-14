package defensive_example_receiver

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/defensive_example_receiver"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "DefensiveExampleReceiver"

type PaymentMethod uint8

const (
	NativeToken PaymentMethod = iota
	FeeToken
)

type ConstructorArgs struct {
	Router   common.Address
	FeeToken common.Address
}

type RemoteChainConfig struct {
	RemoteChainSelector uint64
	ExtraArgs           []byte
	RequiredCCVs        []common.Address
	OptionalCCVs        []common.Address
	OptionalThreshold   uint8
}

type SendDataArgs struct {
	Method            PaymentMethod
	DestChainSelector uint64
	Receiver          []byte
	Data              []byte
}

var Deploy = contract.NewDeploy(
	"defensive-example-receiver:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the DefensiveExampleReceiver contract",
	ContractType,
	ccv_proxy.CCVProxyABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := defensive_example_receiver.DeployDefensiveExample(opts, backend, args.Router, args.FeeToken)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, deployInput ConstructorArgs) (common.Address, error)
	},
)

var EnableRemoteChain = contract.NewWrite(
	"defensive-example-receiver:enable-remote-chain",
	semver.MustParse("1.7.0"),
	"Enables a remote chain on the DefensiveExampleReceiver",
	ContractType,
	defensive_example_receiver.DefensiveExampleABI,
	defensive_example_receiver.NewDefensiveExample,
	contract.OnlyOwner,
	func(RemoteChainConfig) error { return nil },
	func(defensiveExample *defensive_example_receiver.DefensiveExample, opts *bind.TransactOpts, args RemoteChainConfig) (*types.Transaction, error) {
		return defensiveExample.EnableRemoteChain(opts, args.RemoteChainSelector, args.ExtraArgs, args.RequiredCCVs, args.OptionalCCVs, args.OptionalThreshold)
	},
)

var DisableRemoteChain = contract.NewWrite(
	"defensive-example-receiver:disable-remote-chain",
	semver.MustParse("1.7.0"),
	"Disables a remote chain on the DefensiveExampleReceiver",
	ContractType,
	defensive_example_receiver.DefensiveExampleABI,
	defensive_example_receiver.NewDefensiveExample,
	contract.OnlyOwner,
	func(uint64) error { return nil },
	func(defensiveExample *defensive_example_receiver.DefensiveExample, opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
		return defensiveExample.DisableRemoteChain(opts, remoteChainSelector)
	},
)

var ProvideNativeToken = contract.NewWrite(
	"defensive-example-receiver:provide-native-token",
	semver.MustParse("1.7.0"),
	"Provides native token balance for the sender to the DefensiveExampleReceiver contract",
	ContractType,
	defensive_example_receiver.DefensiveExampleABI,
	defensive_example_receiver.NewDefensiveExample,
	func(contract *defensive_example_receiver.DefensiveExample, opts *bind.CallOpts, caller common.Address) (bool, error) {
		return true, nil
	},
	func(*big.Int) error { return nil },
	func(defensiveExample *defensive_example_receiver.DefensiveExample, opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
		opts.Value = value
		defer func() { opts.Value = nil }()
		return defensiveExample.ProvideNativeToken(opts)
	},
)

var SendData = contract.NewWrite(
	"defensive-example-receiver:send-data",
	semver.MustParse("1.7.0"),
	"Sends data to a remote chain",
	ContractType,
	defensive_example_receiver.DefensiveExampleABI,
	defensive_example_receiver.NewDefensiveExample,
	func(contract *defensive_example_receiver.DefensiveExample, opts *bind.CallOpts, caller common.Address) (bool, error) {
		return true, nil
	},
	func(SendDataArgs) error { return nil },
	func(defensiveExample *defensive_example_receiver.DefensiveExample, opts *bind.TransactOpts, args SendDataArgs) (*types.Transaction, error) {
		return defensiveExample.SendData(opts, uint8(args.Method), args.DestChainSelector, args.Receiver, args.Data)
	},
)
