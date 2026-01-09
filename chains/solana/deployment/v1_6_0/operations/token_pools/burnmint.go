package token_pools

import (
	"bytes"
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/test_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	token_deployments "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

var BurnMintProgramName = "burnmint_token_pool"

var DeployBurnMint = operations.NewOperation(
	"burnmint:deploy",
	common_utils.Version_1_6_0,
	"Deploys the BurnMintTokenPool program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			common_utils.BurnMintTokenPool,
			common_utils.Version_1_6_0,
			"",
			BurnMintProgramName)
	},
)

var InitializeBurnMint = operations.NewOperation(
	"burnmint:initialize",
	common_utils.Version_1_6_0,
	"Initializes the BurnMintTokenPool program",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		batches := make([]types.BatchOperation, 0)
		out, err := operations.ExecuteOperation(b, InitGlobalConfigBurnMint, chain, input)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize global config: %w", err)
		}
		batches = append(batches, out.Output.BatchOps...)
		burnmint_token_pool.SetProgramID(input.TokenPool)
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
			b.Logger.Info("BurnMintTokenPool already initialized for token mint:", input.TokenMint.String())
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
				common_utils.BurnMintTokenPool.String(),
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

var InitGlobalConfigBurnMint = operations.NewOperation(
	"burnmint:global_config",
	common_utils.Version_1_6_0,
	"Initializes the BurnMintTokenPool global config",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.TokenPool)
		programData, err := utils.GetSolProgramData(chain.Client, input.TokenPool)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		upgradeAuthority, err := utils.GetUpgradeAuthority(chain.Client, input.TokenPool)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		configPDA, _, _ := state.FindConfigPDA(input.TokenPool)
		var chainConfig base_token_pool.BaseConfig
		_ = chain.GetAccountDataBorshInto(context.Background(), configPDA, &chainConfig)
		// already initialized
		if !chainConfig.TokenProgram.IsZero() {
			b.Logger.Info("BurnMintTokenPool global config already initialized for token pool:", input.TokenPool.String())
			return sequences.OnChainOutput{}, nil
		}
		batches := make([]types.BatchOperation, 0)
		ixn, err := burnmint_token_pool.NewInitGlobalConfigInstruction(
			input.Router,
			input.RMNRemote,
			configPDA,
			upgradeAuthority,
			solana.SystemProgramID,
			input.TokenPool,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if upgradeAuthority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.TokenPool.String(),
				common_utils.BurnMintTokenPool.String(),
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

var TransferMintAuthorityBurnMint = operations.NewOperation(
	"burnmint:transfer_mint_authority",
	common_utils.Version_1_6_0,
	"Transfers the mint authority of the token pool's mint",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.TokenPool)
		programData, err := utils.GetSolProgramData(chain.Client, input.TokenPool)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		authority := GetAuthorityBurnMint(chain, input.TokenPool, input.TokenMint)
		poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
		poolSignerPDA, _ := tokens.TokenPoolSignerAddress(input.TokenMint, input.TokenPool)
		batches := make([]types.BatchOperation, 0)
		ixn, err := burnmint_token_pool.NewTransferMintAuthorityToMultisigInstruction(
			poolConfigPDA,
			input.TokenMint,
			input.TokenProgramID,
			poolSignerPDA,
			authority,
			input.NewMintAuthority,
			input.TokenPool,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.TokenPool.String(),
				common_utils.BurnMintTokenPool.String(),
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

var UpsertRemoteChainConfigBurnMint = operations.NewOperation(
	"burnmint:init_chain_remote_config",
	common_utils.Version_1_6_0,
	"Initializes the BurnMintTokenPool chain remote config",
	func(b operations.Bundle, chain cldf_solana.Chain, input RemoteChainConfig) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.TokenPool)
		remoteConfig := base_token_pool.RemoteConfig{
			PoolAddresses: []base_token_pool.RemoteAddress{
				{
					Address: input.RemotePoolAddress,
				},
			},
			TokenAddress: base_token_pool.RemoteAddress{
				Address: input.RemoteTokenAddress,
			},
			Decimals: input.RemoteDecimals,
		}
		authority := GetAuthorityBurnMint(chain, input.TokenPool, input.TokenMint)
		poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
		// check if remote chain config already exists
		remoteChainConfigPDA, _, _ := tokens.TokenPoolChainConfigPDA(input.RemoteSelector, input.TokenMint, input.TokenPool)
		isSuportedChain := false
		existingConfig := base_token_pool.BaseChain{}
		var remoteChainConfigAccount base_token_pool.BaseChain
		err := chain.GetAccountDataBorshInto(context.Background(), remoteChainConfigPDA, &remoteChainConfigAccount)
		if err == nil {
			isSuportedChain = true
			existingConfig = remoteChainConfigAccount
		}
		batches := make([]types.BatchOperation, 0)
		var ixn solana.Instruction
		if isSuportedChain {
			// if the token address has changed or if the override config flag is set, edit the remote config (just overwrite the existing remote config)
			if !bytes.Equal(existingConfig.Remote.TokenAddress.Address, input.RemoteTokenAddress) || input.ForceOverrideRemoteConfig {
				ixn, err = burnmint_token_pool.NewEditChainRemoteConfigInstruction(
					input.RemoteSelector,
					input.TokenMint,
					remoteConfig,
					poolConfigPDA,
					remoteChainConfigPDA,
					authority,
					solana.SystemProgramID,
				).ValidateAndBuild()
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
			} else {
				// diff between [existing remote pool addresses on solana chain] vs [what was just derived from evm chain]
				poolAddresses := existingConfig.Remote.PoolAddresses
				// translate to base
				baseAddresses := make([]base_token_pool.RemoteAddress, len(poolAddresses))
				for i, cfg := range poolAddresses {
					baseAddresses[i] = base_token_pool.RemoteAddress{
						Address: cfg.Address,
					}
				}
				diff := poolDiff(baseAddresses, remoteConfig.PoolAddresses)
				if len(diff) > 0 {
					ixn, err = burnmint_token_pool.NewAppendRemotePoolAddressesInstruction(
						input.RemoteSelector,
						input.TokenMint,
						diff, // evm supports multiple remote pools per token
						poolConfigPDA,
						remoteChainConfigPDA,
						authority,
						solana.SystemProgramID,
					).ValidateAndBuild()
					if err != nil {
						return sequences.OnChainOutput{}, err
					}
				}
			}
		} else {
			ixn, err = burnmint_token_pool.NewInitChainRemoteConfigInstruction(
				input.RemoteSelector,
				input.TokenMint,
				remoteConfig,
				poolConfigPDA,
				remoteChainConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.TokenPool.String(),
				common_utils.BurnMintTokenPool.String(),
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

var UpsertRateLimitsBurnMint = operations.NewOperation(
	"burnmint:rate_limits",
	common_utils.Version_1_6_0,
	"Initializes the BurnMintTokenPool rate limits for a remote chain",
	func(b operations.Bundle, chain cldf_solana.Chain, input RemoteChainConfig) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.TokenPool)
		inbound := base_token_pool.RateLimitConfig{
			Enabled:  input.InboundRateLimiterConfig.IsEnabled,
			Capacity: input.InboundRateLimiterConfig.Capacity.Uint64(),
			Rate:     input.InboundRateLimiterConfig.Rate.Uint64(),
		}
		outbound := base_token_pool.RateLimitConfig{
			Enabled:  input.OutboundRateLimiterConfig.IsEnabled,
			Capacity: input.OutboundRateLimiterConfig.Capacity.Uint64(),
			Rate:     input.OutboundRateLimiterConfig.Rate.Uint64(),
		}
		authority := GetAuthorityBurnMint(chain, input.TokenPool, input.TokenMint)
		poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
		// check if remote chain config already exists
		remoteChainConfigPDA, _, _ := tokens.TokenPoolChainConfigPDA(input.RemoteSelector, input.TokenMint, input.TokenPool)
		batches := make([]types.BatchOperation, 0)
		ixn, err := burnmint_token_pool.NewSetChainRateLimitInstruction(
			input.RemoteSelector,
			input.TokenMint,
			inbound,
			outbound,
			poolConfigPDA,
			remoteChainConfigPDA,
			authority,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.TokenPool.String(),
				common_utils.BurnMintTokenPool.String(),
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

var TransferOwnership = operations.NewOperation(
	"burnmint:transfer-ownership",
	common_utils.Version_1_6_0,
	"Transfers ownership of the BurnMintTokenPool token mint PDA to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenPoolTransferOwnershipInput) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.Program)
		authority := GetAuthorityBurnMint(chain, input.Program, input.TokenMint)
		if authority != input.CurrentOwner {
			return sequences.OnChainOutput{}, fmt.Errorf("current owner %s does not match on-chain authority %s", input.CurrentOwner.String(), authority.String())
		}
		tokenPoolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.Program)
		ixn, err := burnmint_token_pool.NewTransferOwnershipInstruction(
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
				common_utils.BurnMintTokenPool.String(),
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

var AcceptOwnership = operations.NewOperation(
	"burnmint:accept-ownership",
	common_utils.Version_1_6_0,
	"Accepts ownership of the BurnMintTokenPool token mint PDA",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenPoolTransferOwnershipInput) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.Program)
		tokenPoolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.Program)
		ixn, err := burnmint_token_pool.NewAcceptOwnershipInstruction(
			tokenPoolConfigPDA,
			input.TokenMint,
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
				common_utils.BurnMintTokenPool.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm accept ownership: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

type Params struct {
	TokenPool solana.PublicKey
	TokenMint solana.PublicKey
	// SPLToken or SPLToken2022
	TokenProgramID solana.PublicKey
	// Only used for certain ops
	RMNRemote        solana.PublicKey
	Router           solana.PublicKey
	NewMintAuthority solana.PublicKey
	OldMintAuthority solana.PublicKey
}

type RemoteChainConfig struct {
	TokenPool solana.PublicKey
	TokenMint solana.PublicKey
	// SPLToken or SPLToken2022
	TokenProgramID            solana.PublicKey
	RemoteSelector            uint64
	RemoteTokenAddress        []byte
	RemotePoolAddress         []byte
	RemoteDecimals            uint8
	ForceOverrideRemoteConfig bool
	InboundRateLimiterConfig  token_deployments.RateLimiterConfig
	OutboundRateLimiterConfig token_deployments.RateLimiterConfig
}
