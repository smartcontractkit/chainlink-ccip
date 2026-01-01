package router

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
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

var ApplyRampUpdates = contract.NewWrite(contract.WriteParams[ApplyRampsUpdatesArgs, *router.Router]{
	Name:            "router:apply-ramp-updates",
	Version:         Version,
	Description:     "Applies ramp updates to the Router",
	ContractType:    ContractType,
	ContractABI:     router.RouterABI,
	NewContract:     router.NewRouter,
	IsAllowedCaller: contract.OnlyOwner[*router.Router, ApplyRampsUpdatesArgs],
	Validate: func(router *router.Router, backend bind.ContractBackend, opts *bind.CallOpts, args ApplyRampsUpdatesArgs) error {
		offRamps, err := router.GetOffRamps(opts)
		if err != nil {
			return fmt.Errorf("failed to get off ramps on router with address %s: %w", router.Address(), err)
		}
		for _, offRampRemoval := range args.OffRampRemoves {
			exists := false
			for _, offRamp := range offRamps {
				if offRampRemoval.SourceChainSelector == offRamp.SourceChainSelector && offRampRemoval.OffRamp == offRamp.OffRamp {
					exists = true
					break
				}
			}
			if !exists {
				return fmt.Errorf("off ramp %s does not exist on router with address %s for source chain selector %d", offRampRemoval.OffRamp, router.Address(), offRampRemoval.SourceChainSelector)
			}
		}

		return nil
	},
	IsNoop: func(router *router.Router, opts *bind.CallOpts, args ApplyRampsUpdatesArgs) (bool, error) {
		offRamps, err := router.GetOffRamps(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get off ramps on router with address %s: %w", router.Address(), err)
		}

		for _, offRampAdd := range args.OffRampAdds {
			exists := false
			for _, offRamp := range offRamps {
				if offRampAdd.SourceChainSelector == offRamp.SourceChainSelector && offRampAdd.OffRamp == offRamp.OffRamp {
					exists = true
					break
				}
			}
			if !exists {
				return false, nil
			}
		}

		for _, onRampUpdate := range args.OnRampUpdates {
			onRamp, err := router.GetOnRamp(opts, onRampUpdate.DestChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get on ramp on router with address %s for dest chain selector %d: %w", router.Address(), onRampUpdate.DestChainSelector, err)
			}
			if onRamp != onRampUpdate.OnRamp {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(router *router.Router, opts *bind.TransactOpts, args ApplyRampsUpdatesArgs) (*types.Transaction, error) {
		return router.ApplyRampUpdates(opts, args.OnRampUpdates, args.OffRampRemoves, args.OffRampAdds)
	},
})

var CCIPSend = contract.NewWrite(contract.WriteParams[CCIPSendArgs, *router.Router]{
	Name:            "router:ccip-send",
	Version:         Version,
	Description:     "Sends a CCIP message via the Router",
	ContractType:    ContractType,
	ContractABI:     router.RouterABI,
	NewContract:     router.NewRouter,
	IsAllowedCaller: contract.AllCallersAllowed[*router.Router, CCIPSendArgs],
	Validate: func(router *router.Router, backend bind.ContractBackend, opts *bind.CallOpts, args CCIPSendArgs) error {
		return nil
	},
	IsNoop: func(router *router.Router, opts *bind.CallOpts, args CCIPSendArgs) (bool, error) {
		return false, nil
	},
	CallContract: func(router *router.Router, opts *bind.TransactOpts, args CCIPSendArgs) (*types.Transaction, error) {
		opts.Value = args.Value
		defer func() { opts.Value = nil }()
		return router.CcipSend(opts, args.DestChainSelector, args.EVM2AnyMessage)
	},
})

var GetOffRamps = contract.NewRead(contract.ReadParams[any, []OffRamp, *router.Router]{
	Name:         "router:get-off-ramps",
	Version:      Version,
	Description:  "Gets all off ramps on the router",
	ContractType: ContractType,
	NewContract:  router.NewRouter,
	CallContract: func(router *router.Router, opts *bind.CallOpts, args any) ([]OffRamp, error) {
		return router.GetOffRamps(opts)
	},
})

var GetOnRamp = contract.NewRead(contract.ReadParams[uint64, common.Address, *router.Router]{
	Name:         "router:get-on-ramp",
	Version:      Version,
	Description:  "Gets the on ramp for a given destination chain selector",
	ContractType: ContractType,
	NewContract:  router.NewRouter,
	CallContract: func(router *router.Router, opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
		return router.GetOnRamp(opts, destChainSelector)
	},
})

var GetFee = contract.NewRead(contract.ReadParams[CCIPSendArgs, *big.Int, *router.Router]{
	Name:         "router:get-fee",
	Version:      Version,
	Description:  "Gets the fee for a message",
	ContractType: ContractType,
	NewContract:  router.NewRouter,
	CallContract: func(router *router.Router, opts *bind.CallOpts, args CCIPSendArgs) (*big.Int, error) {
		return router.GetFee(opts, args.DestChainSelector, args.EVM2AnyMessage)
	},
})

var IsChainSupported = contract.NewRead(contract.ReadParams[uint64, bool, *router.Router]{
	Name:         "router:isChainSupported",
	Version:      Version,
	Description:  "If the router supports a given destination chain selector",
	ContractType: ContractType,
	NewContract:  router.NewRouter,
	CallContract: func(router *router.Router, opts *bind.CallOpts, args uint64) (bool, error) {
		return router.IsChainSupported(opts, args)
	},
})
