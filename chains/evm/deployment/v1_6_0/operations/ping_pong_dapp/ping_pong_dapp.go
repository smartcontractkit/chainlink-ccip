package pingpongdapp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/ping_pong_demo"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "PingPongDemo"
var Version *semver.Version = semver.MustParse("1.5.0")

// ConstructorArgs contains the arguments for deploying the PingPongDemo contract.
type ConstructorArgs struct {
	Router   common.Address
	FeeToken common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "ping-pong-demo:deploy",
	Version:          Version,
	Description:      "Deploys the PingPongDemo contract for cross-chain ping pong messaging",
	ContractMetadata: ping_pong_demo.PingPongDemoMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(ping_pong_demo.PingPongDemoBin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})

// SetCounterpartArgs contains the arguments for setting the counterpart chain and address.
type SetCounterpartArgs struct {
	CounterpartChainSelector uint64
	CounterpartAddress       []byte
}

var SetCounterpart = contract.NewWrite(contract.WriteParams[SetCounterpartArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-counterpart",
	Version:         Version,
	Description:     "Sets the counterpart chain selector and address for the PingPongDemo contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, SetCounterpartArgs],
	Validate:        func(SetCounterpartArgs) error { return nil },
	CallContract: func(pingPong *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args SetCounterpartArgs) (*types.Transaction, error) {
		return pingPong.SetCounterpart(opts, args.CounterpartChainSelector, args.CounterpartAddress)
	},
})

// SetCounterpartChainSelectorArgs contains the arguments for setting just the counterpart chain selector.
type SetCounterpartChainSelectorArgs struct {
	ChainSelector uint64
}

var SetCounterpartChainSelector = contract.NewWrite(contract.WriteParams[SetCounterpartChainSelectorArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-counterpart-chain-selector",
	Version:         Version,
	Description:     "Sets the counterpart chain selector for the PingPongDemo contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, SetCounterpartChainSelectorArgs],
	Validate:        func(SetCounterpartChainSelectorArgs) error { return nil },
	CallContract: func(pingPong *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args SetCounterpartChainSelectorArgs) (*types.Transaction, error) {
		return pingPong.SetCounterpartChainSelector(opts, args.ChainSelector)
	},
})

// SetCounterpartAddressArgs contains the arguments for setting just the counterpart address.
type SetCounterpartAddressArgs struct {
	Address []byte
}

var SetCounterpartAddress = contract.NewWrite(contract.WriteParams[SetCounterpartAddressArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-counterpart-address",
	Version:         Version,
	Description:     "Sets the counterpart address for the PingPongDemo contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, SetCounterpartAddressArgs],
	Validate:        func(SetCounterpartAddressArgs) error { return nil },
	CallContract: func(pingPong *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args SetCounterpartAddressArgs) (*types.Transaction, error) {
		return pingPong.SetCounterpartAddress(opts, args.Address)
	},
})

// SetPausedArgs contains the arguments for pausing/unpausing the contract.
type SetPausedArgs struct {
	Paused bool
}

var SetPaused = contract.NewWrite(contract.WriteParams[SetPausedArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-paused",
	Version:         Version,
	Description:     "Pauses or unpauses the PingPongDemo contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, SetPausedArgs],
	Validate:        func(SetPausedArgs) error { return nil },
	CallContract: func(pingPong *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args SetPausedArgs) (*types.Transaction, error) {
		return pingPong.SetPaused(opts, args.Paused)
	},
})

// SetOutOfOrderExecutionArgs contains the arguments for enabling/disabling out of order execution.
type SetOutOfOrderExecutionArgs struct {
	OutOfOrderExecution bool
}

var SetOutOfOrderExecution = contract.NewWrite(contract.WriteParams[SetOutOfOrderExecutionArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-out-of-order-execution",
	Version:         Version,
	Description:     "Enables or disables out of order execution for the PingPongDemo contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, SetOutOfOrderExecutionArgs],
	Validate:        func(SetOutOfOrderExecutionArgs) error { return nil },
	CallContract: func(pingPong *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args SetOutOfOrderExecutionArgs) (*types.Transaction, error) {
		return pingPong.SetOutOfOrderExecution(opts, args.OutOfOrderExecution)
	},
})

// StartPingPongArgs is an empty struct since startPingPong takes no arguments.
type StartPingPongArgs struct{}

var StartPingPong = contract.NewWrite(contract.WriteParams[StartPingPongArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:start-ping-pong",
	Version:         Version,
	Description:     "Starts the ping pong messaging loop",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, StartPingPongArgs],
	Validate:        func(StartPingPongArgs) error { return nil },
	CallContract: func(pingPong *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, _ StartPingPongArgs) (*types.Transaction, error) {
		return pingPong.StartPingPong(opts)
	},
})
