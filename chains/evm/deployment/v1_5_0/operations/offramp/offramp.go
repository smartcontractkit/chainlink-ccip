package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_offramp"
)

var (
	ContractType cldf_deployment.ContractType = "EVM2EVMOffRamp"
	Version      *semver.Version              = semver.MustParse("1.5.0")
)

func NewReadOffRampStaticConfig(c *evm_2_evm_offramp.EVM2EVMOffRamp) *cld_ops.Operation[contract.FunctionInput[struct{}], evm_2_evm_offramp.EVM2EVMOffRampStaticConfig, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, evm_2_evm_offramp.EVM2EVMOffRampStaticConfig, *evm_2_evm_offramp.EVM2EVMOffRamp]{
		Name:         "offramp:static-config",
		Version:      Version,
		Description:  "Reads the static config from the OffRamp 1.5.0 contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(offRamp *evm_2_evm_offramp.EVM2EVMOffRamp, opts *bind.CallOpts, args struct{}) (evm_2_evm_offramp.EVM2EVMOffRampStaticConfig, error) {
			return offRamp.GetStaticConfig(opts)
		},
	})
}

func NewReadOffRampDynamicConfig(c *evm_2_evm_offramp.EVM2EVMOffRamp) *cld_ops.Operation[contract.FunctionInput[struct{}], evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, *evm_2_evm_offramp.EVM2EVMOffRamp]{
		Name:         "offramp:dynamic-config",
		Version:      Version,
		Description:  "Reads the dynamic config from the OffRamp 1.5.0 contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(offRamp *evm_2_evm_offramp.EVM2EVMOffRamp, opts *bind.CallOpts, args struct{}) (evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, error) {
			return offRamp.GetDynamicConfig(opts)
		},
	})
}
