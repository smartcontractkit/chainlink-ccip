package rmn_remote

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/rmn_remote"
	rmn161 "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_1/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	ContractType cldf_deployment.ContractType = "RMNRemote"
	ProgramName                               = "rmn_remote"
	Version      *semver.Version              = semver.MustParse("1.6.1")
)

// TODO include 1.6.0 code here when fully migrating to 1.6.1

type EventAuthoritiesInput struct {
	EventAuthorities   []solana.PublicKey
	RMNRemote          solana.PublicKey
	RMNRemoteConfigPDA solana.PublicKey
}

var SetEventAuthorities = operations.NewOperation(
	"rmn-remote:set-event-authorities",
	Version,
	"Sets the event authorities list on the RMNRemote contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input EventAuthoritiesInput) (sequences.OnChainOutput, error) {
		authority := GetAuthority(chain, input.RMNRemote)

		ixn, err := rmn161.NewSetEventAuthoritiesInstruction(
			input.EventAuthorities,
			input.RMNRemoteConfigPDA,
			authority,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build set event authorities instruction: %w", err)
		}

		batches := make([]types.BatchOperation, 0)
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.RMNRemote.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
		} else {
			err := chain.Confirm([]solana.Instruction{ixn})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm set-event-authorities instruction: %w", err)
			}
		}
		return sequences.OnChainOutput{BatchOps: batches}, nil
	},
)

func GetAuthority(chain cldf_solana.Chain, program solana.PublicKey) solana.PublicKey {
	programData := rmn_remote.Config{}
	rmnRemoteConfigPDA, _, _ := state.FindRMNRemoteConfigPDA(program)
	err := chain.GetAccountDataBorshInto(context.Background(), rmnRemoteConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Owner
}
