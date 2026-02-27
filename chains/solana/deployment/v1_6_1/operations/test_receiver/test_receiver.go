package test_receiver

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	ContractType cldf_deployment.ContractType = "TestReceiver"
	Version      *semver.Version              = semver.MustParse("1.6.1")
	ProgramName                               = "test_ccip_receiver"
)

var Deploy = operations.NewOperation(
	"receiver:deploy",
	Version,
	"Deploys the Receiver program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			ContractType,
			Version,
			"",
			ProgramName)
	},
)

var Initialize = operations.NewOperation(
	"receiver:initialize",
	Version,
	"Initializes the Receiver 1.6.1 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		test_ccip_receiver.SetProgramID(input.Receiver)
		externalExecutionConfigPDA, _, _ := solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, input.Receiver)
		receiverTargetAccount, _, _ := solana.FindProgramAddress([][]byte{[]byte("counter")}, input.Receiver)
		instruction, err := test_ccip_receiver.NewInitializeInstruction(
			input.Router,
			receiverTargetAccount,
			externalExecutionConfigPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		if err := chain.Confirm([]solana.Instruction{instruction}); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm instructions: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

type Params struct {
	Router   solana.PublicKey
	Receiver solana.PublicKey
}
