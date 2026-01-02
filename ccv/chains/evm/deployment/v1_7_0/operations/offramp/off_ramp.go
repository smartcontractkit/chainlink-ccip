package offramp

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OffRamp"

var Version = semver.MustParse("1.7.0")

type StaticConfig = offramp.OffRampStaticConfig

type ConstructorArgs struct {
	StaticConfig StaticConfig
}

type SourceChainConfigArgs = offramp.OffRampSourceChainConfigArgs

type SourceChainConfig = offramp.OffRampSourceChainConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "off-ramp:deploy",
	Version:          Version,
	Description:      "Deploys the OffRamp contract",
	ContractMetadata: offramp.OffRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(offramp.OffRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]SourceChainConfigArgs, *offramp.OffRamp]{
	Name:            "off-ramp:apply-source-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to source chain configurations on the OffRamp",
	ContractType:    ContractType,
	ContractABI:     offramp.OffRampABI,
	NewContract:     offramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*offramp.OffRamp, []SourceChainConfigArgs],
	Validate: func(offRamp *offramp.OffRamp, backend bind.ContractBackend, opts *bind.CallOpts, args []SourceChainConfigArgs) error {
		return nil
	},
	IsNoop: func(offRamp *offramp.OffRamp, opts *bind.CallOpts, args []SourceChainConfigArgs) (bool, error) {
		for _, arg := range args {
			actualSourceChainConfig, err := offRamp.GetSourceChainConfig(opts, arg.SourceChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get source chain config: %w", err)
			}
			if actualSourceChainConfig.IsEnabled != arg.IsEnabled ||
				actualSourceChainConfig.Router != arg.Router {
				return false, nil
			}
			if len(actualSourceChainConfig.DefaultCCVs) != len(arg.DefaultCCVs) ||
				len(actualSourceChainConfig.LaneMandatedCCVs) != len(arg.LaneMandatedCCVs) {
				return false, nil
			}
			slices.SortFunc(actualSourceChainConfig.DefaultCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(actualSourceChainConfig.LaneMandatedCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(arg.DefaultCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(arg.LaneMandatedCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			if !slices.Equal(actualSourceChainConfig.DefaultCCVs, arg.DefaultCCVs) ||
				!slices.Equal(actualSourceChainConfig.LaneMandatedCCVs, arg.LaneMandatedCCVs) {
				return false, nil
			}

			if len(actualSourceChainConfig.OnRamps) != len(arg.OnRamps) {
				return false, nil
			}
			for _, argOnRamp := range arg.OnRamps {
				found := false
				for _, actualOnRamp := range actualSourceChainConfig.OnRamps {
					if bytes.Equal(actualOnRamp, argOnRamp) {
						found = true
						break
					}
				}
				if !found {
					return false, nil
				}
			}
		}

		return true, nil
	},
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.TransactOpts, args []SourceChainConfigArgs) (*types.Transaction, error) {
		return offRamp.ApplySourceChainConfigUpdates(opts, args)
	},
})

var GetSourceChainConfig = contract.NewRead(contract.ReadParams[uint64, SourceChainConfig, *offramp.OffRamp]{
	Name:         "off-ramp:get-source-chain-config",
	Version:      Version,
	Description:  "Gets the source chain configuration for a given source chain selector",
	ContractType: ContractType,
	NewContract:  offramp.NewOffRamp,
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.CallOpts, args uint64) (SourceChainConfig, error) {
		return offRamp.GetSourceChainConfig(opts, args)
	},
})
