package type_and_version

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	// devSuffix is the suffix for development versions.
	devSuffix = "-dev"
)

// typeAndVersionABI is the ABI for the typeAndVersion function
// function typeAndVersion() external view returns (string)
var typeAndVersionABI = `[{"inputs":[],"name":"typeAndVersion","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

var getTypeAndVersion = cldf_ops.NewOperation(
	"get-type-and-version",
	semver.MustParse("1.0.0"),
	"Fetches the type and version of a contract",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, input contract_utils.FunctionInput[any]) (output string, err error) {
		// Validate input
		if input.ChainSelector != chain.Selector {
			return "", fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
		}
		if input.Address == (common.Address{}) {
			return "", fmt.Errorf("address must be specified for type-and-version")
		}

		// Parse the ABI
		parsedABI, err := abi.JSON(strings.NewReader(typeAndVersionABI))
		if err != nil {
			return "", fmt.Errorf("failed to parse ABI: %w", err)
		}

		// Create a bound contract
		boundContract := bind.NewBoundContract(input.Address, parsedABI, chain.Client, chain.Client, chain.Client)

		// Make the call - boundContract.Call handles packing/unpacking automatically
		var result []interface{}
		err = boundContract.Call(&bind.CallOpts{Context: b.GetContext()}, &result, "typeAndVersion")
		if err != nil {
			return "", fmt.Errorf("failed to call typeAndVersion on contract %s: %w", input.Address, err)
		}

		// Extract the string result
		if len(result) == 0 {
			return "", fmt.Errorf("typeAndVersion returned no results")
		}

		typeAndVersionStr, ok := result[0].(string)
		if !ok {
			return "", fmt.Errorf("typeAndVersion returned unexpected type: %T", result[0])
		}

		return typeAndVersionStr, nil
	},
)
