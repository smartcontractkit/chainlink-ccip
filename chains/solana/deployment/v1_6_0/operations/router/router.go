package router

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

var ContractType cldf_deployment.ContractType = "Router"
var DestChainType cldf_deployment.ContractType = "DestChain"
var ProgramName = "ccip_router"
var ProgramSize = 5 * 1024 * 1024
var Version *semver.Version = semver.MustParse("1.6.0")

type ConnectChainsParams struct {
	Router              solana.PublicKey
	OffRamp             solana.PublicKey
	RemoteChainSelector uint64
	AllowlistEnabled    bool
	AllowedSenders      []solana.PublicKey
}

var Deploy = operations.NewOperation(
	"router:deploy",
	Version,
	"Deploys the Router program",
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
	"router:initialize",
	Version,
	"Initializes the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		programData, err := utils.GetSolProgramData(chain, input.Router)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		authority := GetAuthority(chain, input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		instruction, err := ccip_router.NewInitializeInstruction(
			chain.Selector,
			solana.PublicKey{},
			input.FeeQuoter,
			input.LinkToken,
			input.RMNRemote,
			routerConfigPDA,
			authority,
			solana.SystemProgramID,
			input.Router,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm router initialization: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var ConnectChains = operations.NewOperation(
	"router:connect-chains",
	Version,
	"Connects the Router 1.6.0 contract to other chains",
	func(b operations.Bundle, chain cldf_solana.Chain, input ConnectChainsParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		isUpdate := false
		authority := GetAuthority(chain, input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		routerDestChainPDA, _ := state.FindDestChainStatePDA(input.RemoteChainSelector, input.Router)
		var destChainAccount ccip_router.DestChain
		err := chain.GetAccountDataBorshInto(context.Background(), routerDestChainPDA, &destChainAccount)
		if err == nil {
			fmt.Println("Remote chain state account found:", destChainAccount)
			isUpdate = true
		}
		destChainConfig := ccip_router.DestChainConfig{
			AllowedSenders:   input.AllowedSenders,
			AllowListEnabled: input.AllowlistEnabled,
		}
		var ixn solana.Instruction
		batches := make([]types.BatchOperation, 0)
		if isUpdate {
			ixn, err = ccip_router.NewUpdateDestChainConfigInstruction(
				input.RemoteChainSelector,
				destChainConfig,
				routerDestChainPDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build update dest chain instruction: %w", err)
			}
		} else {
			ixn, err = ccip_router.NewAddChainSelectorInstruction(
				input.RemoteChainSelector,
				destChainConfig,
				routerDestChainPDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build add source chain instruction: %w", err)
			}
			err = utils.ExtendLookupTable(chain, input.OffRamp, []solana.PublicKey{routerDestChainPDA})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to extend OffRamp lookup table: %w", err)
			}
		}
		sourceRef := datastore.AddressRef{
			Address:       routerDestChainPDA.String(),
			ChainSelector: chain.Selector,
			Type:          datastore.ContractType(DestChainType),
			Version:       Version,
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
		} else {
			err = chain.Confirm([]solana.Instruction{ixn})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
			}
		}
		return sequences.OnChainOutput{
			BatchOps:  batches,
			Addresses: []datastore.AddressRef{sourceRef},
		}, nil
	},
)

var AddOffRamp = operations.NewOperation(
	"router:add-off-ramp",
	Version,
	"Adds an OffRamp to the Router 1.6.0 contract for a given chain",
	func(b operations.Bundle, chain cldf_solana.Chain, input ConnectChainsParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		authority := GetAuthority(chain, input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		allowedOffRampRemotePDA, _ := state.FindAllowedOfframpPDA(input.RemoteChainSelector, input.OffRamp, input.Router)
		ixn, err := ccip_router.NewAddOfframpInstruction(
			input.RemoteChainSelector,
			input.OffRamp,
			allowedOffRampRemotePDA,
			routerConfigPDA,
			authority,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if authority != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var TransferOwnership = operations.NewOperation(
	"router:transfer-ownership",
	Version,
	"Transfers ownership of the Router 1.6.0 contract to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Program)
		authority := GetAuthority(chain, input.Program)
		if authority != input.CurrentOwner {
			return sequences.OnChainOutput{}, fmt.Errorf("current owner %s does not match on-chain authority %s", input.CurrentOwner.String(), authority.String())
		}
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := ccip_router.NewTransferOwnershipInstruction(
			input.NewOwner,
			configPDA,
			authority,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if authority != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Program.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var AcceptOwnership = operations.NewOperation(
	"router:accept-ownership",
	Version,
	"Accepts ownership of the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Program)
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := ccip_router.NewAcceptOwnershipInstruction(
			configPDA,
			input.NewOwner,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if input.NewOwner != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Program.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

func GetAuthority(chain cldf_solana.Chain, program solana.PublicKey) solana.PublicKey {
	programData := ccip_router.Config{}
	routerConfigPDA, _, _ := state.FindConfigPDA(program)
	err := chain.GetAccountDataBorshInto(context.Background(), routerConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Owner
}

type Params struct {
	FeeQuoter solana.PublicKey
	Router    solana.PublicKey
	LinkToken solana.PublicKey
	RMNRemote solana.PublicKey
}
