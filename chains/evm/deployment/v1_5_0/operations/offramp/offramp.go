package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_offramp"
)

var (
	ContractType cldf_deployment.ContractType = "EVM2EVMOffRamp"
	Version      *semver.Version              = semver.MustParse("1.5.0")

	OffRampStaticConfig = contract.NewRead(contract.ReadParams[any, evm_2_evm_offramp.EVM2EVMOffRampStaticConfig, *evm_2_evm_offramp.EVM2EVMOffRamp]{
		Name:         "offramp:static-config",
		Version:      Version,
		Description:  "Reads the static config from the OffRamp 1.5.0 contract",
		ContractType: ContractType,
		NewContract:  evm_2_evm_offramp.NewEVM2EVMOffRamp,
		CallContract: func(offRamp *evm_2_evm_offramp.EVM2EVMOffRamp, opts *bind.CallOpts, args any) (evm_2_evm_offramp.EVM2EVMOffRampStaticConfig, error) {
			return offRamp.GetStaticConfig(opts)
		},
	})

	OffRampDynamicConfig = contract.NewRead(contract.ReadParams[any, evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, *evm_2_evm_offramp.EVM2EVMOffRamp]{
		Name:         "offramp:dynamic-config",
		Version:      Version,
		Description:  "Reads the dynamic config from the OffRamp 1.5.0 contract",
		ContractType: ContractType,
		NewContract:  evm_2_evm_offramp.NewEVM2EVMOffRamp,
		CallContract: func(offRamp *evm_2_evm_offramp.EVM2EVMOffRamp, opts *bind.CallOpts, args any) (evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, error) {
			return offRamp.GetDynamicConfig(opts)
		},
	})
)
