package ccip_home

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"
var Version *semver.Version = semver.MustParse("1.6.3")

type AddDONOpInput struct {
	Nodes                    [][32]byte
	CapabilityConfigurations []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration
	IsPublic                 bool
	AcceptsWorkflows         bool
	F                        uint8
}

var AddDON = contract.NewWrite(contract.WriteParams[AddDONOpInput, *capabilities_registry.CapabilitiesRegistry]{
	Name:            "capabilities-registry:add-don",
	Version:         Version,
	Description:     "Adds a new DON to the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     ccip_home.CCIPHomeABI,
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
	Version:         Version,
	Description:     "Updates an existing DON in the CapabilitiesRegistry",
	ContractType:    utils.CapabilitiesRegistry,
	ContractABI:     ccip_home.CCIPHomeABI,
	NewContract:     capabilities_registry.NewCapabilitiesRegistry,
	IsAllowedCaller: contract.OnlyOwner[*capabilities_registry.CapabilitiesRegistry, UpdateDONOpInput],
	Validate:        func(UpdateDONOpInput) error { return nil },
	CallContract: func(capReg *capabilities_registry.CapabilitiesRegistry, opts *bind.TransactOpts, input UpdateDONOpInput) (*types.Transaction, error) {
		return capReg.UpdateDON(opts, input.ID, input.Nodes, input.CapabilityConfigurations, input.IsPublic, input.F)
	},
})

type ApplyChainConfigUpdatesOpInput struct {
	RemoteChainRemoves []uint64
	RemoteChainAdds    []ccip_home.CCIPHomeChainConfigArgs
}

var ApplyChainConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyChainConfigUpdatesOpInput, *ccip_home.CCIPHome]{
	Name:            "capabilities-registry:update-don",
	Version:         Version,
	Description:     "Updates an existing DON in the CapabilitiesRegistry",
	ContractType:    utils.CCIPHome,
	ContractABI:     ccip_home.CCIPHomeABI,
	NewContract:     ccip_home.NewCCIPHome,
	IsAllowedCaller: contract.OnlyOwner[*ccip_home.CCIPHome, ApplyChainConfigUpdatesOpInput],
	Validate:        func(ApplyChainConfigUpdatesOpInput) error { return nil },
	CallContract: func(ccipHome *ccip_home.CCIPHome, opts *bind.TransactOpts, input ApplyChainConfigUpdatesOpInput) (*types.Transaction, error) {
		return ccipHome.ApplyChainConfigUpdates(opts, input.RemoteChainRemoves, input.RemoteChainAdds)
	},
})
