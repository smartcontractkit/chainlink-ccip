package router

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
)

var ContractType cldf_deployment.ContractType = "Router"
var TestRouterContractType cldf_deployment.ContractType = "TestRouter"
var Version *semver.Version = semver.MustParse("1.2.0")

type ConstructorArgs struct {
	WrappedNative common.Address
	RMNProxy      common.Address
}

type OnRamp = router.RouterOnRamp

type OffRamp = router.RouterOffRamp

type ApplyRampsUpdatesArgs struct {
	OnRampUpdates  []OnRamp
	OffRampRemoves []OffRamp
	OffRampAdds    []OffRamp
}

type EVMTokenAmount = router.ClientEVMTokenAmount

type EVM2AnyMessage = router.ClientEVM2AnyMessage

type CCIPSendArgs struct {
	Value             *big.Int
	DestChainSelector uint64
	EVM2AnyMessage    EVM2AnyMessage
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "router:deploy",
	Version:          Version,
	Description:      "Deploys the Router contract",
	ContractMetadata: router.RouterMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(router.RouterBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var DeployTestRouter = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "test-router:deploy",
	Version:          Version,
	Description:      "Deploys the Test Router contract",
	ContractMetadata: router.RouterMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(TestRouterContractType, *Version).String(): {
			EVM: common.FromHex(router.RouterBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

func NewWriteApplyRampUpdates(c *router.Router) *cld_ops.Operation[contract.FunctionInput[ApplyRampsUpdatesArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ApplyRampsUpdatesArgs, *router.Router]{
		Name:            "router:apply-ramp-updates",
		Version:         Version,
		Description:     "Applies ramp updates to the Router",
		ContractType:    ContractType,
		ContractABI:     router.RouterABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*router.Router, ApplyRampsUpdatesArgs],
		Validate:        func(ApplyRampsUpdatesArgs) error { return nil },
		CallContract: func(r *router.Router, opts *bind.TransactOpts, args ApplyRampsUpdatesArgs) (*types.Transaction, error) {
			return r.ApplyRampUpdates(opts, args.OnRampUpdates, args.OffRampRemoves, args.OffRampAdds)
		},
	})
}

func NewWriteCCIPSend(c *router.Router) *cld_ops.Operation[contract.FunctionInput[CCIPSendArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[CCIPSendArgs, *router.Router]{
		Name:            "router:ccip-send",
		Version:         Version,
		Description:     "Sends a CCIP message via the Router",
		ContractType:    ContractType,
		ContractABI:     router.RouterABI,
		Contract:        c,
		IsAllowedCaller: contract.AllCallersAllowed[*router.Router, CCIPSendArgs],
		Validate:        func(CCIPSendArgs) error { return nil },
		CallContract: func(r *router.Router, opts *bind.TransactOpts, args CCIPSendArgs) (*types.Transaction, error) {
			opts.Value = args.Value
			defer func() { opts.Value = nil }()
			return r.CcipSend(opts, args.DestChainSelector, args.EVM2AnyMessage)
		},
	})
}

func NewReadGetOffRamps(c *router.Router) *cld_ops.Operation[contract.FunctionInput[struct{}], []OffRamp, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, []OffRamp, *router.Router]{
		Name:         "router:get-off-ramps",
		Version:      Version,
		Description:  "Gets all off ramps on the router",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(r *router.Router, opts *bind.CallOpts, args struct{}) ([]OffRamp, error) {
			return r.GetOffRamps(opts)
		},
	})
}

func NewReadGetOnRamp(c router.RouterInterface) *cld_ops.Operation[contract.FunctionInput[uint64], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, common.Address, router.RouterInterface]{
		Name:         "router:get-on-ramp",
		Version:      Version,
		Description:  "Gets the on ramp for a given destination chain selector",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(r router.RouterInterface, opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
			return r.GetOnRamp(opts, destChainSelector)
		},
	})
}

func NewReadGetFee(c *router.Router) *cld_ops.Operation[contract.FunctionInput[CCIPSendArgs], *big.Int, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[CCIPSendArgs, *big.Int, *router.Router]{
		Name:         "router:get-fee",
		Version:      Version,
		Description:  "Gets the fee for a message",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(r *router.Router, opts *bind.CallOpts, args CCIPSendArgs) (*big.Int, error) {
			return r.GetFee(opts, args.DestChainSelector, args.EVM2AnyMessage)
		},
	})
}

func NewReadIsChainSupported(c *router.Router) *cld_ops.Operation[contract.FunctionInput[uint64], bool, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, bool, *router.Router]{
		Name:         "router:isChainSupported",
		Version:      Version,
		Description:  "If the router supports a given destination chain selector",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(r *router.Router, opts *bind.CallOpts, args uint64) (bool, error) {
			return r.IsChainSupported(opts, args)
		},
	})
}

func NewReadGetWrappedNative(c *router.Router) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, *router.Router]{
		Name:         "router:get-wrapped-native",
		Version:      Version,
		Description:  "Gets the wrapped native address",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(r *router.Router, opts *bind.CallOpts, args struct{}) (common.Address, error) {
			return r.GetWrappedNative(opts)
		},
	})
}

func NewWriteSetWrappedNative(c *router.Router) *cld_ops.Operation[contract.FunctionInput[common.Address], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[common.Address, *router.Router]{
		Name:            "router:set-wrapped-native",
		Version:         Version,
		Description:     "Sets the wrapped native address",
		ContractType:    ContractType,
		ContractABI:     router.RouterABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*router.Router, common.Address],
		Validate:        func(common.Address) error { return nil },
		CallContract: func(r *router.Router, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
			return r.SetWrappedNative(opts, args)
		},
	})
}
