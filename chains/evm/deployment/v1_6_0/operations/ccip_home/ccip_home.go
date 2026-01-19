package ccip_home

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/core/types"

	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var (
	CCIPHomeVersion             = semver.MustParse("1.6.0")
	CapabilitiesRegistryVersion = semver.MustParse("1.0.0")
)

type AddDONOpInput struct {
	Nodes                    [][32]byte
	CapabilityConfigurations []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration
	IsPublic                 bool
	AcceptsWorkflows         bool
	F                        uint8
}

var AddDON = contract.NewWrite(contract.WriteParams[AddDONOpInput, *capabilities_registry.CapabilitiesRegistry]{
	Name:            "capabilities-registry:add-don",
	Version:         CapabilitiesRegistryVersion,
	Description:     "Adds a new DON to the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     capabilities_registry.CapabilitiesRegistryABI,
	NewContract:     capabilities_registry.NewCapabilitiesRegistry,
	IsAllowedCaller: contract.OnlyOwner[*capabilities_registry.CapabilitiesRegistry, AddDONOpInput],
	Validate:        func(AddDONOpInput) error { return nil },
	CallContract: func(capReg *capabilities_registry.CapabilitiesRegistry, opts *bind.TransactOpts, input AddDONOpInput) (*types.Transaction, error) {
		return capReg.AddDON(opts, input.Nodes, input.CapabilityConfigurations, input.IsPublic, input.AcceptsWorkflows, input.F)
	},
})

type UpdateDONOpInput struct {
	ID                       uint32
	Nodes                    [][32]byte
	CapabilityConfigurations []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration
	IsPublic                 bool
	F                        uint8
}

var UpdateDON = contract.NewWrite(contract.WriteParams[UpdateDONOpInput, *capabilities_registry.CapabilitiesRegistry]{
	Name:            "capabilities-registry:update-don",
	Version:         CapabilitiesRegistryVersion,
	Description:     "Updates an existing DON in the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     capabilities_registry.CapabilitiesRegistryABI,
	NewContract:     capabilities_registry.NewCapabilitiesRegistry,
	IsAllowedCaller: contract.OnlyOwner[*capabilities_registry.CapabilitiesRegistry, UpdateDONOpInput],
	Validate:        func(UpdateDONOpInput) error { return nil },
	CallContract: func(capReg *capabilities_registry.CapabilitiesRegistry, opts *bind.TransactOpts, input UpdateDONOpInput) (*types.Transaction, error) {
		return capReg.UpdateDON(opts, input.ID, input.Nodes, input.CapabilityConfigurations, input.IsPublic, input.F)
	},
})

type AddNodesOpInput struct {
	Nodes []capabilities_registry.CapabilitiesRegistryNodeParams
}

var AddNodes = contract.NewWrite(contract.WriteParams[AddNodesOpInput, *capabilities_registry.CapabilitiesRegistry]{
	Name:            "capabilities-registry:add-nodes",
	Version:         CapabilitiesRegistryVersion,
	Description:     "Adds nodes to an existing node operator in the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     capabilities_registry.CapabilitiesRegistryABI,
	NewContract:     capabilities_registry.NewCapabilitiesRegistry,
	IsAllowedCaller: contract.OnlyOwner[*capabilities_registry.CapabilitiesRegistry, AddNodesOpInput],
	Validate:        func(AddNodesOpInput) error { return nil },
	CallContract: func(capReg *capabilities_registry.CapabilitiesRegistry, opts *bind.TransactOpts, input AddNodesOpInput) (*types.Transaction, error) {
		return capReg.AddNodes(opts, input.Nodes)
	},
})

type AddNodesOperatorsOpInput struct {
	Nodes []capabilities_registry.CapabilitiesRegistryNodeOperator
}

var AddNodeOperators = contract.NewWrite(contract.WriteParams[AddNodesOperatorsOpInput, *capabilities_registry.CapabilitiesRegistry]{
	Name:            "capabilities-registry:add-node-operators",
	Version:         CapabilitiesRegistryVersion,
	Description:     "Adds new node operators to the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     capabilities_registry.CapabilitiesRegistryABI,
	NewContract:     capabilities_registry.NewCapabilitiesRegistry,
	IsAllowedCaller: contract.OnlyOwner[*capabilities_registry.CapabilitiesRegistry, AddNodesOperatorsOpInput],
	Validate:        func(AddNodesOperatorsOpInput) error { return nil },
	CallContract: func(capReg *capabilities_registry.CapabilitiesRegistry, opts *bind.TransactOpts, input AddNodesOperatorsOpInput) (*types.Transaction, error) {
		return capReg.AddNodeOperators(opts, input.Nodes)
	},
})

type AddCapabilitiesOpInput struct {
	Capabilities []capabilities_registry.CapabilitiesRegistryCapability
}

var AddCapabilities = contract.NewWrite(contract.WriteParams[AddCapabilitiesOpInput, *capabilities_registry.CapabilitiesRegistry]{
	Name:            "capabilities-registry:add-capability",
	Version:         CapabilitiesRegistryVersion,
	Description:     "Adds a new capability to the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     capabilities_registry.CapabilitiesRegistryABI,
	NewContract:     capabilities_registry.NewCapabilitiesRegistry,
	IsAllowedCaller: contract.OnlyOwner[*capabilities_registry.CapabilitiesRegistry, AddCapabilitiesOpInput],
	Validate:        func(AddCapabilitiesOpInput) error { return nil },
	CallContract: func(capReg *capabilities_registry.CapabilitiesRegistry, opts *bind.TransactOpts, input AddCapabilitiesOpInput) (*types.Transaction, error) {
		return capReg.AddCapabilities(opts, input.Capabilities)
	},
})

type ApplyChainConfigUpdatesOpInput struct {
	RemoteChainRemoves []uint64
	RemoteChainAdds    []ccip_home.CCIPHomeChainConfigArgs
}

var ApplyChainConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyChainConfigUpdatesOpInput, *ccip_home.CCIPHome]{
	Name:            "ccip-home:apply-chain-config-updates",
	Version:         CCIPHomeVersion,
	Description:     "Applies chain config updates to the CCIPHome contract",
	ContractType:    utils.CCIPHome,
	ContractABI:     ccip_home.CCIPHomeABI,
	NewContract:     ccip_home.NewCCIPHome,
	IsAllowedCaller: contract.OnlyOwner[*ccip_home.CCIPHome, ApplyChainConfigUpdatesOpInput],
	Validate:        func(ApplyChainConfigUpdatesOpInput) error { return nil },
	CallContract: func(ccipHome *ccip_home.CCIPHome, opts *bind.TransactOpts, input ApplyChainConfigUpdatesOpInput) (*types.Transaction, error) {
		return ccipHome.ApplyChainConfigUpdates(opts, input.RemoteChainRemoves, input.RemoteChainAdds)
	},
})
