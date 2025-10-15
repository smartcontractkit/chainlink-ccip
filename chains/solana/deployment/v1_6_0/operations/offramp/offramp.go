package offramp

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ContractType cldf_deployment.ContractType = "OffRamp"
var ProgramName = "off_ramp"
var ProgramSize = int(1.5 * 1024 * 1024)
var Version *semver.Version = semver.MustParse("1.6.0")

type Params struct {
	EnableExecutionAfter int64
	FeeQuoter            solana.PublicKey
	Router               solana.PublicKey
	OffRamp              solana.PublicKey
	RMNRemote            solana.PublicKey
}

var Deploy = operations.NewOperation(
	"off-ramp:deploy",
	Version,
	"Deploys the OffRamp program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			ContractType,
			Version,
			"",
			ProgramName,
			ProgramSize)
	},
)

var Initialize = operations.NewOperation(
	"off-ramp:initialize",
	Version,
	"Initializes the OffRamp 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) ([]solana.Instruction, error) {
		programData, err := utils.GetSolProgramData(chain, input.OffRamp)
		if err != nil {
			return nil, fmt.Errorf("failed to get program data: %w", err)
		}
		table, err := common.SetupLookupTable(
			context.Background(),
			chain.Client,
			*chain.DeployerKey,
			[]solana.PublicKey{
				// system
				solana.SystemProgramID,
				solana.ComputeBudget,
				solana.SysVarInstructionsPubkey,
				// token
				solana.Token2022ProgramID,
				solana.TokenProgramID,
				solana.SPLAssociatedTokenAccountProgramID,
			})
		if err != nil {
			return nil, fmt.Errorf("failed to setup lookup table: %w", err)
		}
		offRampReferenceAddressesPDA, _, _ := state.FindOfframpReferenceAddressesPDA(input.OffRamp)
		offRampStatePDA, _, _ := state.FindOfframpStatePDA(input.OffRamp)
		instruction, err := ccip_offramp.NewInitializeInstruction(
			offRampReferenceAddressesPDA,
			input.Router,
			input.FeeQuoter,
			input.RMNRemote,
			table,
			offRampStatePDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
			input.OffRamp,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return nil, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		return []solana.Instruction{instruction}, nil
	},
)

var InitializeConfig = operations.NewOperation(
	"off-ramp:initialize-config",
	Version,
	"Initializes the config of the OffRamp 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) ([]solana.Instruction, error) {
		programData, err := utils.GetSolProgramData(chain, input.OffRamp)
		if err != nil {
			return nil, fmt.Errorf("failed to get program data: %w", err)
		}
		offRampConfigPDA, _, _ := state.FindOfframpConfigPDA(input.OffRamp)
		instruction, err := ccip_offramp.NewInitializeConfigInstruction(
			chain.Selector,
			input.EnableExecutionAfter,
			offRampConfigPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
			input.OffRamp,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return nil, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return nil, fmt.Errorf("failed to confirm initialization: %w", err)
		}
		return []solana.Instruction{instruction}, nil
	},
)

func ExtendLookupTable(chain cldf_solana.Chain, offRampID solana.PublicKey, lookUpTableEntries []solana.PublicKey) error {
	var referenceAddressesAccount ccip_offramp.ReferenceAddresses
	offRampReferenceAddressesPDA, _, _ := state.FindOfframpReferenceAddressesPDA(offRampID)
	err := chain.GetAccountDataBorshInto(context.Background(), offRampReferenceAddressesPDA, &referenceAddressesAccount)
	if err != nil {
		return fmt.Errorf("failed to get offramp reference addresses: %w", err)
	}
	addressLookupTable := referenceAddressesAccount.OfframpLookupTable

	addresses, err := common.GetAddressLookupTable(
		context.Background(),
		chain.Client,
		addressLookupTable)
	if err != nil {
		return fmt.Errorf("failed to get address lookup table: %w", err)
	}

	// calculate diff and add new entries
	seen := make(map[solana.PublicKey]bool)
	toAdd := make([]solana.PublicKey, 0)
	for _, entry := range addresses {
		seen[entry] = true
	}
	for _, entry := range lookUpTableEntries {
		if _, ok := seen[entry]; !ok {
			toAdd = append(toAdd, entry)
		}
	}
	if len(toAdd) == 0 {
		return nil
	}

	if err := common.ExtendLookupTable(
		context.Background(),
		chain.Client,
		addressLookupTable,
		*chain.DeployerKey,
		toAdd,
	); err != nil {
		return fmt.Errorf("failed to extend lookup table: %w", err)
	}
	return nil
}

func DefaultParams() Params {
	return Params{
		EnableExecutionAfter: int64((20 * time.Minute).Seconds()),
	}
}
