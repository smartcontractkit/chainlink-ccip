package token_pools

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/test_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

var LockReleaseProgramName = "lockrelease_token_pool"

var DeployLockRelease = operations.NewOperation(
	"lockrelease:deploy",
	common_utils.Version_1_6_0,
	"Deploys the LockReleaseTokenPool program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			common_utils.LockReleaseTokenPool,
			common_utils.Version_1_6_0,
			"",
			LockReleaseProgramName)
	},
)

var InitializeLockRelease = operations.NewOperation(
	"lockrelease:initialize",
	common_utils.Version_1_6_0,
	"Initializes the LockReleaseTokenPool program",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		batches := make([]types.BatchOperation, 0)
		out, err := operations.ExecuteOperation(b, InitGlobalConfigLockRelease, chain, input)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize global config: %w", err)
		}
		batches = append(batches, out.Output.BatchOps...)
		lockrelease_token_pool.SetProgramID(input.TokenPool)
		programData, err := utils.GetSolProgramData(chain.Client, input.TokenPool)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		upgradeAuthority, err := utils.GetUpgradeAuthority(chain.Client, input.TokenPool)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
		var chainConfig test_token_pool.State
		err = chain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &chainConfig)
		if err == nil {
			b.Logger.Info("LockReleaseTokenPool already initialized for token mint:", input.TokenMint.String())
			return sequences.OnChainOutput{}, nil
		}
		configPDA, _, _ := state.FindConfigPDA(input.TokenPool)
		ixn, err := lockrelease_token_pool.NewInitializeInstruction(
			poolConfigPDA,
			input.TokenMint,
			upgradeAuthority,
			solana.SystemProgramID,
			input.TokenPool,
			programData.Address,
			configPDA,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if upgradeAuthority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.TokenPool.String(),
				common_utils.LockReleaseTokenPool.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
			return sequences.OnChainOutput{BatchOps: batches}, nil
		} else {
			err = chain.Confirm([]solana.Instruction{ixn})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		}

		return sequences.OnChainOutput{}, nil
	})

var InitGlobalConfigLockRelease = operations.NewOperation(
	"lockrelease:global_config",
	common_utils.Version_1_6_0,
	"Initializes the LockReleaseTokenPool global config",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		return initGlobalConfigTokenPool(b, chain, input, initGlobalCfgParams{
			PoolTypeLabel: common_utils.LockReleaseTokenPool.String(),
			LogName:       "LockReleaseTokenPool",
			SetProgramID:  lockrelease_token_pool.SetProgramID,
			BuildInitIx: func(configPDA solana.PublicKey, upgradeAuthority solana.PublicKey, programData solana.PublicKey) (solana.Instruction, error) {
				return lockrelease_token_pool.NewInitGlobalConfigInstruction(
					input.Router,
					input.RMNRemote,
					configPDA,
					upgradeAuthority,
					solana.SystemProgramID,
					input.TokenPool,
					programData,
				).ValidateAndBuild()
			},
		})
	},
)

var TransferOwnershipLockRelease = operations.NewOperation(
	"lockrelease:transfer-ownership",
	common_utils.Version_1_6_0,
	"Transfers ownership of the LockReleaseTokenPool token mint PDA to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenPoolTransferOwnershipInput) (sequences.OnChainOutput, error) {
		lockrelease_token_pool.SetProgramID(input.Program)
		authority := GetAuthorityLockRelease(chain, input.Program, input.TokenMint)
		if authority != input.CurrentOwner {
			return sequences.OnChainOutput{}, fmt.Errorf("current owner %s does not match on-chain authority %s", input.CurrentOwner.String(), authority.String())
		}
		tokenPoolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.Program)
		ixn, err := lockrelease_token_pool.NewTransferOwnershipInstruction(
			input.NewOwner,
			tokenPoolConfigPDA,
			input.TokenMint,
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
				common_utils.LockReleaseTokenPool.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm transfer ownership: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

func GetAuthorityLockRelease(chain cldf_solana.Chain, program solana.PublicKey, tokenMint solana.PublicKey) solana.PublicKey {
	programData := lockrelease_token_pool.State{}
	poolConfigPDA, _ := tokens.TokenPoolConfigAddress(tokenMint, program)
	err := chain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Config.Owner
}
