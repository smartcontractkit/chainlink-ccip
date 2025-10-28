package offramp

import (
	"context"
	"fmt"
	"math"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

var ContractType cldf_deployment.ContractType = "OffRamp"
var SourceChainType cldf_deployment.ContractType = "SourceChain"
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

type ConnectChainsParams struct {
	OffRamp             solana.PublicKey
	RemoteChainSelector uint64
	SourceOnRamp        []byte
	EnabledAsSource     bool
}

type SetOcr3Params struct {
	OffRamp            solana.PublicKey
	SetOCR3ConfigInput deployops.SetOCR3ConfigInput
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
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		ccip_offramp.SetProgramID(input.OffRamp)
		programData, err := utils.GetSolProgramData(chain, input.OffRamp)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		authority := GetAuthority(chain, input.OffRamp)
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
			return sequences.OnChainOutput{}, fmt.Errorf("failed to setup lookup table: %w", err)
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
			authority,
			solana.SystemProgramID,
			input.OffRamp,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm offramp initialization: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var InitializeConfig = operations.NewOperation(
	"off-ramp:initialize-config",
	Version,
	"Initializes the config of the OffRamp 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		ccip_offramp.SetProgramID(input.OffRamp)
		programData, err := utils.GetSolProgramData(chain, input.OffRamp)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		authority := GetAuthority(chain, input.OffRamp)
		offRampConfigPDA, _, _ := state.FindOfframpConfigPDA(input.OffRamp)
		instruction, err := ccip_offramp.NewInitializeConfigInstruction(
			chain.Selector,
			input.EnableExecutionAfter,
			offRampConfigPDA,
			authority,
			solana.SystemProgramID,
			input.OffRamp,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm initialization: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var ConnectChains = operations.NewOperation(
	"off-ramp:connect-chains",
	Version,
	"Connects the OffRamp 1.6.0 contract to other chains",
	func(b operations.Bundle, chain cldf_solana.Chain, input ConnectChainsParams) (sequences.OnChainOutput, error) {
		ccip_offramp.SetProgramID(input.OffRamp)
		isUpdate := false
		authority := GetAuthority(chain, input.OffRamp)
		offRampConfigPDA, _, _ := state.FindOfframpConfigPDA(input.OffRamp)
		offRampSourceChainPDA, _, _ := state.FindOfframpSourceChainPDA(input.RemoteChainSelector, input.OffRamp)
		var sourceChainAccount ccip_offramp.SourceChain
		err := chain.GetAccountDataBorshInto(context.Background(), offRampSourceChainPDA, &sourceChainAccount)
		if err == nil {
			fmt.Println("Remote chain state account found:", sourceChainAccount)
			isUpdate = true
		}
		var onRampAddress ccip_offramp.OnRampAddress
		copy(onRampAddress.Bytes[:], input.SourceOnRamp)
		addressBytesLen := len(onRampAddress.Bytes)
		if addressBytesLen < 0 || addressBytesLen > math.MaxUint32 {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid on ramp address length: %d", addressBytesLen)
		}
		onRampAddress.Len = uint32(addressBytesLen)
		validSourceChainConfig := ccip_offramp.SourceChainConfig{
			OnRamp:    onRampAddress,
			IsEnabled: input.EnabledAsSource,
		}
		var ixn solana.Instruction
		batches := make([]types.BatchOperation, 0)
		if isUpdate {
			ixn, err = ccip_offramp.NewUpdateSourceChainConfigInstruction(
				input.RemoteChainSelector,
				validSourceChainConfig,
				offRampSourceChainPDA,
				offRampConfigPDA,
				authority,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build update dest chain instruction: %w", err)
			}
		} else {
			ixn, err = ccip_offramp.NewAddSourceChainInstruction(
				input.RemoteChainSelector,
				validSourceChainConfig,
				offRampSourceChainPDA,
				offRampConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build add source chain instruction: %w", err)
			}
			err = utils.ExtendLookupTable(chain, input.OffRamp, []solana.PublicKey{offRampSourceChainPDA})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to extend OffRamp lookup table: %w", err)
			}
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.OffRamp.String(),
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
		sourceRef := datastore.AddressRef{
			Address:       offRampSourceChainPDA.String(),
			ChainSelector: chain.Selector,
			Type:          datastore.ContractType(SourceChainType),
			Version:       Version,
		}
		return sequences.OnChainOutput{
			BatchOps:  batches,
			Addresses: []datastore.AddressRef{sourceRef},
		}, nil
	},
)

var SetOcr3 = operations.NewOperation(
	"off-ramp:set-ocr3",
	Version,
	"Sets the OCR3 configuration for the OffRamp 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input SetOcr3Params) (sequences.OnChainOutput, error) {
		ccip_offramp.SetProgramID(input.OffRamp)
		authority := GetAuthority(chain, input.OffRamp)
		batches := make([]types.BatchOperation, 0)
		offRampConfigPDA, _, _ := state.FindOfframpConfigPDA(input.OffRamp)
		offRampStatePDA, _, _ := state.FindOfframpStatePDA(input.OffRamp)
		for _, arg := range input.SetOCR3ConfigInput.Configs {
			var ocrType ccip_offramp.OcrPluginType
			switch arg.PluginType {
			case ccipocr3.PluginTypeCCIPCommit:
				ocrType = ccip_offramp.Commit_OcrPluginType
			case ccipocr3.PluginTypeCCIPExec:
				ocrType = ccip_offramp.Execution_OcrPluginType
			default:
				return sequences.OnChainOutput{}, fmt.Errorf("unsupported OCR plugin type: %d", arg.PluginType)
			}

			var signerAddresses [][20]byte
			var transmitterAddresses []solana.PublicKey
			for _, signer := range arg.Signers {
				var solanaSigner [20]uint8
				if len(signer) != 20 {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid signer length: %d", len(signer))
				}
				copy(solanaSigner[:], signer)
				signerAddresses = append(signerAddresses, solanaSigner)
			}
			for _, transmitter := range arg.Transmitters {
				solanaTransmitter := solana.PublicKeyFromBytes(transmitter)
				transmitterAddresses = append(transmitterAddresses, solanaTransmitter)
			}

			instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
				ocrType,
				ccip_offramp.Ocr3ConfigInfo{
					ConfigDigest:                   arg.ConfigDigest,
					F:                              arg.F,
					IsSignatureVerificationEnabled: btoi(arg.IsSignatureVerificationEnabled),
				},
				signerAddresses,
				transmitterAddresses,
				offRampConfigPDA,
				offRampStatePDA,
				authority,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build set OCR3 config instruction: %w", err)
			}
			if authority != chain.DeployerKey.PublicKey() {
				b, err := utils.BuildMCMSBatchOperation(
					chain.Selector,
					[]solana.Instruction{instruction},
					input.OffRamp.String(),
					ContractType.String(),
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
				}
				batches = append(batches, b)
			} else {
				err = chain.Confirm([]solana.Instruction{instruction})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm set OCR3 config: %w", err)
				}
			}
		}

		return sequences.OnChainOutput{BatchOps: batches}, nil
	},
)

var TransferOwnership = operations.NewOperation(
	"off-ramp:transfer-ownership",
	Version,
	"Transfers ownership of the OffRamp 1.6.0 contract to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		ccip_offramp.SetProgramID(input.Program)
		authority := GetAuthority(chain, input.Program)
		if authority != input.CurrentOwner {
			return sequences.OnChainOutput{}, fmt.Errorf("current owner %s does not match on-chain authority %s", input.CurrentOwner.String(), authority.String())
		}
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := ccip_offramp.NewTransferOwnershipInstruction(
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
	"off-ramp:accept-ownership",
	Version,
	"Accepts ownership of the OffRamp 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		ccip_offramp.SetProgramID(input.Program)
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := ccip_offramp.NewAcceptOwnershipInstruction(
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

func btoi(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func GetAuthority(chain cldf_solana.Chain, program solana.PublicKey) solana.PublicKey {
	programData := ccip_offramp.Config{}
	offRampConfigPDA, _, _ := state.FindOfframpConfigPDA(program)
	err := chain.GetAccountDataBorshInto(context.Background(), offRampConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Owner
}
