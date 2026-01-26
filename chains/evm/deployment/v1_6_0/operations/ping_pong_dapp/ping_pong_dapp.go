package ping_pong_demo

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/ping_pong_demo"
)

var ContractType cldf_deployment.ContractType = "PingPongDemo"
var Version = semver.MustParse("1.5.0")

type ConstructorArgs struct {
	Router   common.Address
	FeeToken common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "ping-pong-demo:deploy",
	Version:          Version,
	Description:      "Deploys the PingPongDemo contract",
	ContractMetadata: ping_pong_demo.PingPongDemoMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(ping_pong_demo.PingPongDemoBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

type SetCounterpartArgs struct {
	CounterpartChainSelector uint64
	CounterpartAddress       common.Address
}

var SetCounterpart = contract.NewWrite(contract.WriteParams[SetCounterpartArgs, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-counterpart",
	Version:         Version,
	Description:     "Calls setCounterpart on the contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, SetCounterpartArgs],
	Validate:        func(SetCounterpartArgs) error { return nil },
	CallContract: func(pingPongDemo *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args SetCounterpartArgs) (*types.Transaction, error) {
		return pingPongDemo.SetCounterpart(opts, args.CounterpartChainSelector, args.CounterpartAddress)
	},
})

var SetCounterpartChainSelector = contract.NewWrite(contract.WriteParams[uint64, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-counterpart-chain-selector",
	Version:         Version,
	Description:     "Calls setCounterpartChainSelector on the contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, uint64],
	Validate:        func(uint64) error { return nil },
	CallContract: func(pingPongDemo *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args uint64) (*types.Transaction, error) {
		return pingPongDemo.SetCounterpartChainSelector(opts, args)
	},
})

var SetCounterpartAddress = contract.NewWrite(contract.WriteParams[common.Address, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-counterpart-address",
	Version:         Version,
	Description:     "Calls setCounterpartAddress on the contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(pingPongDemo *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return pingPongDemo.SetCounterpartAddress(opts, args)
	},
})

var SetPaused = contract.NewWrite(contract.WriteParams[bool, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-paused",
	Version:         Version,
	Description:     "Calls setPaused on the contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, bool],
	Validate:        func(bool) error { return nil },
	CallContract: func(pingPongDemo *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args bool) (*types.Transaction, error) {
		return pingPongDemo.SetPaused(opts, args)
	},
})

var SetOutOfOrderExecution = contract.NewWrite(contract.WriteParams[bool, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:set-out-of-order-execution",
	Version:         Version,
	Description:     "Calls setOutOfOrderExecution on the contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, bool],
	Validate:        func(bool) error { return nil },
	CallContract: func(pingPongDemo *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args bool) (*types.Transaction, error) {
		return pingPongDemo.SetOutOfOrderExecution(opts, args)
	},
})

var StartPingPong = contract.NewWrite(contract.WriteParams[struct{}, *ping_pong_demo.PingPongDemo]{
	Name:            "ping-pong-demo:start-ping-pong",
	Version:         Version,
	Description:     "Calls startPingPong on the contract",
	ContractType:    ContractType,
	ContractABI:     ping_pong_demo.PingPongDemoABI,
	NewContract:     ping_pong_demo.NewPingPongDemo,
	IsAllowedCaller: contract.OnlyOwner[*ping_pong_demo.PingPongDemo, struct{}],
	Validate:        func(struct{}) error { return nil },
	CallContract: func(pingPongDemo *ping_pong_demo.PingPongDemo, opts *bind.TransactOpts, args struct{}) (*types.Transaction, error) {
		return pingPongDemo.StartPingPong(opts)
	},
})
