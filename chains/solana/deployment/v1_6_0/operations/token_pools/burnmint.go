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
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (PoolInitializeOut, error) {
		batches := make([]types.BatchOperation, 0)
		out, err := operations.ExecuteOperation(b, InitGlobalConfigBurnMint, chain, input)
		if err != nil {
			return PoolInitializeOut{}, fmt.Errorf("failed to initialize global config: %w", err)
		}
		batches = append(batches, out.Output.BatchOps...)
		burnmint_token_pool.SetProgramID(input.TokenPool)
		programData, err := utils.GetSolProgramData(chain.Client, input.TokenPool)
		if err != nil {
			return PoolInitializeOut{}, err
		}
		upgradeAuthority, err := utils.GetUpgradeAuthority(chain.Client, input.TokenPool)
		if err != nil {
			return PoolInitializeOut{}, err
		}
		poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
		var chainConfig test_token_pool.State
		err = chain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &chainConfig)
		if err == nil {
			b.Logger.Info("BurnMintTokenPool already initialized for token mint:", input.TokenMint.String())
			return PoolInitializeOut{}, nil
		}
		// use the deployer key if we can
		mintAuthority := utils.GetTokenMintAuthority(chain, input.TokenMint)
		signer := upgradeAuthority
		var poolConfig burnmint_token_pool.PoolConfig
		globalConfig, _ := tokens.TokenPoolGlobalConfigPDA(input.TokenPool)
		err = chain.GetAccountDataBorshInto(context.Background(), globalConfig, &poolConfig)
		if err == nil {
			b.Logger.Info("Fetched existing pool config for token mint:", input.TokenMint.String())
			if mintAuthority == chain.DeployerKey.PublicKey() &&
				poolConfig.SelfServedAllowed {
				signer = mintAuthority
			}
		} else {
			b.Logger.Info("No existing pool config found for token mint, defaulting to upgrade authority as signer for initialization:", input.TokenMint.String())
		}
		configPDA, _, _ := state.FindConfigPDA(input.TokenPool)
		ixn, err := burnmint_token_pool.NewInitializeInstruction(
			poolConfigPDA,
			input.TokenMint,
			signer,
			solana.SystemProgramID,
			input.TokenPool,
			programData.Address,
			configPDA,
		).ValidateAndBuild()
		if err != nil {
			return PoolInitializeOut{}, err
		}
		if signer != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.TokenPool.String(),
				common_utils.BurnMintTokenPool.String(),
			)
			if err != nil {
				return PoolInitializeOut{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
			return PoolInitializeOut{
				OnChainOutput: sequences.OnChainOutput{BatchOps: batches}, Initializer: signer}, nil
		} else {
			err = chain.Confirm([]solana.Instruction{ixn})
			if err != nil {
				return PoolInitializeOut{}, err
			}
		}

		return PoolInitializeOut{Initializer: signer}, nil
	})

var InitGlobalConfigBurnMint = operations.NewOperation(
	"burnmint:global_config",
	common_utils.Version_1_6_0,
	"Initializes the BurnMintTokenPool global config",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		return initGlobalConfigTokenPool(b, chain, input, initGlobalCfgParams{
			PoolTypeLabel: common_utils.BurnMintTokenPool.String(),
			LogName:       "BurnMintTokenPool",
			SetProgramID:  burnmint_token_pool.SetProgramID,
			BuildInitIx: func(configPDA solana.PublicKey, upgradeAuthority solana.PublicKey, programData solana.PublicKey) (solana.Instruction, error) {
				return burnmint_token_pool.NewInitGlobalConfigInstruction(
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
			PoolAddresses: []base_token_pool.RemoteAddress{},
			TokenAddress: base_token_pool.RemoteAddress{
				Address: input.RemoteTokenAddress,
			},
			Decimals: input.RemoteDecimals,
		}
		authority := GetAuthorityBurnMint(chain, input.TokenPool, input.TokenMint)
		poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
		// check if remote chain config already exists
		remoteChainConfigPDA, _, _ := tokens.TokenPoolChainConfigPDA(input.RemoteSelector, input.TokenMint, input.TokenPool)
		isSupportedChain := false
		existingConfig := base_token_pool.BaseChain{}
		var remoteChainConfigAccount base_token_pool.BaseChain
		err := chain.GetAccountDataBorshInto(context.Background(), remoteChainConfigPDA, &remoteChainConfigAccount)
		if err == nil {
			isSupportedChain = true
			existingConfig = remoteChainConfigAccount
		}
		batches := make([]types.BatchOperation, 0)
		var ixns []solana.Instruction
		if isSupportedChain {
			remoteConfig.PoolAddresses = append(remoteConfig.PoolAddresses,
				base_token_pool.RemoteAddress{
					Address: input.RemotePoolAddress,
				})
			// if the token address has changed or if the override config flag is set, edit the remote config (just overwrite the existing remote config)
			if !bytes.Equal(existingConfig.Remote.TokenAddress.Address, input.RemoteTokenAddress) || input.ForceOverrideRemoteConfig {
				ixn, err := burnmint_token_pool.NewEditChainRemoteConfigInstruction(
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
				ixns = append(ixns, ixn)
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
					ixn, err := burnmint_token_pool.NewAppendRemotePoolAddressesInstruction(
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
					ixns = append(ixns, ixn)
				}
			}
		} else {
			ixn, err := burnmint_token_pool.NewInitChainRemoteConfigInstruction(
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
			ixns = append(ixns, ixn)
			appendIxn, err := burnmint_token_pool.NewAppendRemotePoolAddressesInstruction(
				input.RemoteSelector,
				input.TokenMint,
				[]base_token_pool.RemoteAddress{
					{
						Address: input.RemotePoolAddress,
					},
				},
				poolConfigPDA,
				remoteChainConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			ixns = append(ixns, appendIxn)
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				ixns,
				input.TokenPool.String(),
				common_utils.BurnMintTokenPool.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
			return sequences.OnChainOutput{BatchOps: batches}, nil
		} else {
			err = chain.Confirm(ixns)
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
		var inboundCapacity uint64 = 0
		if input.InboundRateLimiterConfig.Capacity != nil {
			inboundCapacity = input.InboundRateLimiterConfig.Capacity.Uint64()
		}
		var inboundRate uint64 = 0
		if input.InboundRateLimiterConfig.Rate != nil {
			inboundRate = input.InboundRateLimiterConfig.Rate.Uint64()
		}
		inbound := base_token_pool.RateLimitConfig{
			Enabled:  input.InboundRateLimiterConfig.IsEnabled,
			Capacity: inboundCapacity,
			Rate:     inboundRate,
		}
		var outboundCapacity uint64 = 0
		if input.OutboundRateLimiterConfig.Capacity != nil {
			outboundCapacity = input.OutboundRateLimiterConfig.Capacity.Uint64()
		}
		var outboundRate uint64 = 0
		if input.OutboundRateLimiterConfig.Rate != nil {
			outboundRate = input.OutboundRateLimiterConfig.Rate.Uint64()
		}
		outbound := base_token_pool.RateLimitConfig{
			Enabled:  input.OutboundRateLimiterConfig.IsEnabled,
			Capacity: outboundCapacity,
			Rate:     outboundRate,
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

var TransferOwnershipBurnMint = operations.NewOperation(
	"burnmint:transfer-ownership",
	common_utils.Version_1_6_0,
	"Transfers ownership of the BurnMintTokenPool token mint PDA to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenPoolTransferOwnershipInput) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.Program)
		// there is a chance we perform an initialize and transfer ownership in the same sequence
		// so we have to assume the input owner is correct, even if it doesn't match the current on-chain authority (since the initialize might be pending a proposal)
		authority := input.CurrentOwner
		if authority.IsZero() {
			b.Logger.Info("Current owner not provided for burn and mint token pool with token mint:", input.TokenMint.String())
			return sequences.OnChainOutput{}, fmt.Errorf("current owner must be provided for burn and mint token pool")
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

func GetAuthorityBurnMint(chain cldf_solana.Chain, program solana.PublicKey, tokenMint solana.PublicKey) solana.PublicKey {
	programData := burnmint_token_pool.State{}
	poolConfigPDA, _ := tokens.TokenPoolConfigAddress(tokenMint, program)
	err := chain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &programData)
	if err != nil {
		// if there is no pool config, default to upgrade authority as the signer for initialization and ownership transfers
		upgradeAuthority, err := utils.GetUpgradeAuthority(chain.Client, program)
		if err != nil {
			return solana.PublicKey{}
		}
		return upgradeAuthority
	}
	return programData.Config.Owner
}
