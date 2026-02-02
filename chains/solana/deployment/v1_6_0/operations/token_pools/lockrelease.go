package token_pools

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/burnmint_token_pool"
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
		ixn, err := burnmint_token_pool.NewInitializeInstruction(
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
