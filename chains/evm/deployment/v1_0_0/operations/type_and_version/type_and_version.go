package type_and_version

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

var ContractType cldf_deployment.ContractType = "ITypeAndVersion"
var Version = utils.Version_1_0_0

const TypeAndVersionABI = `[ { "type": "function", "name": "typeAndVersion", "inputs": [], "outputs": [ { "name": "", "type": "string", "internalType": "string" } ], "stateMutability": "view" } ]`

type TypeAndVersionContract struct {
	address  common.Address
	abi      abi.ABI
	backend  bind.ContractBackend
	contract *bind.BoundContract
}

func NewTypeAndVersionContract(
	address common.Address,
	backend bind.ContractBackend,
) (*TypeAndVersionContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypeAndVersionABI))
	if err != nil {
		return nil, err
	}
	return &TypeAndVersionContract{
		address:  address,
		abi:      parsed,
		backend:  backend,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (c *TypeAndVersionContract) Address() common.Address {
	return c.address
}

func (c *TypeAndVersionContract) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []any
	err := c.contract.Call(opts, &out, "typeAndVersion")
	if err != nil {
		return "", err
	}
	return *abi.ConvertType(out[0], new(string)).(*string), nil
}

type TypeAndVersion struct {
	Type    cldf_deployment.ContractType
	Version *semver.Version
}

var GetTypeAndVersion = contract.NewRead(contract.ReadParams[struct{}, TypeAndVersion, *TypeAndVersionContract]{
	Name:         "type-and-version:get-type-and-version",
	Version:      Version,
	Description:  "Gets the type and version of the contract",
	ContractType: ContractType,
	NewContract:  NewTypeAndVersionContract,
	CallContract: func(c *TypeAndVersionContract, opts *bind.CallOpts, args struct{}) (TypeAndVersion, error) {
		typeAndVersion, err := c.TypeAndVersion(opts)
		if err != nil {
			return TypeAndVersion{}, err
		}
		typeAndVersionValues := strings.Split(typeAndVersion, " ")
		if len(typeAndVersionValues) < 2 {
			return TypeAndVersion{}, fmt.Errorf("invalid type and version %s, expected format: <type> <version>", typeAndVersion)
		}
		version, err := semver.NewVersion(typeAndVersionValues[1])
		if err != nil {
			return TypeAndVersion{}, fmt.Errorf("failed parsing version %s: %w", typeAndVersionValues[1], err)
		}
		return TypeAndVersion{
			Type:    cldf_deployment.ContractType(typeAndVersionValues[0]),
			Version: version,
		}, nil
	},
})
