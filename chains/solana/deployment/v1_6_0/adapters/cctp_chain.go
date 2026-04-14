package adapters

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/mcms/types"

	sol_utils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/cctp_token_pool"
	sol_token_utils "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ adapters.CCTPChain = &SolanaCCTPChainAdapter{}

type SolanaCCTPChainAdapter struct{}

type solanaCCTPState struct {
	chain        cldf_solana.Chain
	tokenPool    solana.PublicKey
	tokenMint    solana.PublicKey
	tokenProgram solana.PublicKey
	poolSigner   solana.PublicKey
	poolTokenATA solana.PublicKey
}

const solanaCCTPV1Mechanism = "CCTP_V1"

// DeployCCTPChain is a no-op for Solana. The Solana CCTP pool/program already exists and this changeset only wires
// remote-chain configuration to it.
func (c *SolanaCCTPChainAdapter) DeployCCTPChain() *operations.Sequence[adapters.DeployCCTPInput, seq_core.OnChainOutput, adapters.DeployCCTPChainDeps] {
	return operations.NewSequence(
		"solana-cctp-chain:no-op-deploy",
		common_utils.Version_1_6_0,
		"Skips Solana CCTP deployment and only supports lane configuration",
		func(b operations.Bundle, deps adapters.DeployCCTPChainDeps, input adapters.DeployCCTPInput) (seq_core.OnChainOutput, error) {
			return seq_core.OnChainOutput{}, nil
		},
	)
}

// ConfigureCCTPChainForLanes appends remote EVM pool addresses to the existing Solana remote-chain config and syncs
// the remote domain/destination caller used for Solana -> remote CCTP burns.
func (c *SolanaCCTPChainAdapter) ConfigureCCTPChainForLanes() *operations.Sequence[adapters.ConfigureCCTPChainForLanesInput, seq_core.OnChainOutput, adapters.ConfigureCCTPChainForLanesDeps] {
	return operations.NewSequence(
		"solana-cctp-chain:configure-cctp-for-lanes",
		common_utils.Version_1_6_0,
		"Configures Solana CCTP remote-chain state for existing lanes",
		func(b operations.Bundle, deps adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) (seq_core.OnChainOutput, error) {
			solChain, ok := deps.BlockChains.SolanaChains()[input.ChainSelector]
			if !ok {
				return seq_core.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}

			tokenPool, err := datastore_utils.FindAndFormatRef(deps.DataStore, input.RegisteredPoolRef, input.ChainSelector, sol_utils.ToAddress)
			if err != nil {
				return seq_core.OnChainOutput{}, fmt.Errorf("failed to resolve Solana CCTP token pool address: %w", err)
			}
			tokenMint, err := solana.PublicKeyFromBase58(input.USDCToken)
			if err != nil {
				return seq_core.OnChainOutput{}, fmt.Errorf("invalid Solana USDC token mint %q: %w", input.USDCToken, err)
			}

			cctp_token_pool.SetProgramID(tokenPool)

			poolConfigPDA, err := sol_token_utils.TokenPoolConfigAddress(tokenMint, tokenPool)
			if err != nil {
				return seq_core.OnChainOutput{}, fmt.Errorf("failed to derive Solana CCTP pool config PDA: %w", err)
			}
			var poolState cctp_token_pool.State
			if err := solChain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &poolState); err != nil {
				return seq_core.OnChainOutput{}, fmt.Errorf("failed to load Solana CCTP pool state: %w", err)
			}

			authority := poolState.Config.Owner
			instructions := make([]solana.Instruction, 0)
			for remoteChainSelector, remoteChainCfg := range input.RemoteChains {
				remotePoolAddress, err := deps.RemoteChains[remoteChainSelector].PoolAddress(
					deps.DataStore,
					deps.BlockChains,
					remoteChainSelector,
					input.RemoteRegisteredPoolRefs[remoteChainSelector],
				)
				if err != nil {
					return seq_core.OnChainOutput{}, fmt.Errorf("failed to resolve remote pool address for chain %d: %w", remoteChainSelector, err)
				}
				var remoteAllowedCallerOnDest []byte
				switch remoteChainCfg.LockOrBurnMechanism {
				case solanaCCTPV1Mechanism:
					remoteAllowedCallerOnDest, err = deps.RemoteChains[remoteChainSelector].CCTPV1AllowedCallerOnDest(
						deps.DataStore,
						deps.BlockChains,
						remoteChainSelector,
					)
					if err != nil {
						return seq_core.OnChainOutput{}, fmt.Errorf("failed to resolve remote CCTP V1 allowed caller for chain %d: %w", remoteChainSelector, err)
					}
				default:
					return seq_core.OnChainOutput{}, fmt.Errorf(
						"solana CCTP only supports %s lanes, got %q for remote chain %d",
						solanaCCTPV1Mechanism,
						remoteChainCfg.LockOrBurnMechanism,
						remoteChainSelector,
					)
				}

				remoteChainConfigPDA, _, err := sol_token_utils.TokenPoolChainConfigPDA(remoteChainSelector, tokenMint, tokenPool)
				if err != nil {
					return seq_core.OnChainOutput{}, fmt.Errorf("failed to derive Solana remote chain config PDA for chain %d: %w", remoteChainSelector, err)
				}

				var remoteChainState cctp_token_pool.ChainConfig
				if err := solChain.GetAccountDataBorshInto(context.Background(), remoteChainConfigPDA, &remoteChainState); err != nil {
					return seq_core.OnChainOutput{}, fmt.Errorf("failed to load existing Solana remote chain config for chain %d: %w", remoteChainSelector, err)
				}

				diff := diffSolanaRemotePoolAddresses(
					remoteChainState.Base.Remote.PoolAddresses,
					[]cctp_token_pool.RemoteAddress{{Address: remotePoolAddress}},
				)
				if len(diff) > 0 {
					appendIx, err := cctp_token_pool.NewAppendRemotePoolAddressesInstruction(
						remoteChainSelector,
						tokenMint,
						diff,
						poolConfigPDA,
						remoteChainConfigPDA,
						authority,
						solana.SystemProgramID,
					).ValidateAndBuild()
					if err != nil {
						return seq_core.OnChainOutput{}, fmt.Errorf("failed to build append-remote-pool instruction for chain %d: %w", remoteChainSelector, err)
					}
					instructions = append(instructions, appendIx)
				}

				desiredDestinationCaller := solana.PublicKeyFromBytes(common.LeftPadBytes(remoteAllowedCallerOnDest, 32))
				if remoteChainState.Cctp.DomainId != remoteChainCfg.DomainIdentifier || remoteChainState.Cctp.DestinationCaller != desiredDestinationCaller {
					editCCTPIx, err := cctp_token_pool.NewEditChainRemoteConfigCctpInstruction(
						remoteChainSelector,
						tokenMint,
						cctp_token_pool.CctpChain{
							DomainId:          remoteChainCfg.DomainIdentifier,
							DestinationCaller: desiredDestinationCaller,
						},
						poolConfigPDA,
						remoteChainConfigPDA,
						authority,
					).ValidateAndBuild()
					if err != nil {
						return seq_core.OnChainOutput{}, fmt.Errorf("failed to build CCTP remote-config instruction for chain %d: %w", remoteChainSelector, err)
					}
					instructions = append(instructions, editCCTPIx)
				}
			}

			if len(instructions) == 0 {
				return seq_core.OnChainOutput{}, nil
			}

			if authority != solChain.DeployerKey.PublicKey() {
				batchOp, err := sol_utils.BuildMCMSBatchOperation(
					input.ChainSelector,
					instructions,
					tokenPool.String(),
					common_utils.CCTPTokenPool.String(),
				)
				if err != nil {
					return seq_core.OnChainOutput{}, fmt.Errorf("failed to build Solana MCMS batch operation: %w", err)
				}
				return seq_core.OnChainOutput{BatchOps: []types.BatchOperation{batchOp}}, nil
			}

			if err := solChain.Confirm(instructions); err != nil {
				return seq_core.OnChainOutput{}, fmt.Errorf("failed to confirm Solana CCTP instructions: %w", err)
			}
			return seq_core.OnChainOutput{}, nil
		},
	)
}

// MigrateHybridLockReleaseLiquidity returns a sequence that immediately errors with a clear message.
// Solana CCTP chains do not support hybrid lock-release liquidity migration.
func (c *SolanaCCTPChainAdapter) MigrateHybridLockReleaseLiquidity() *operations.Sequence[adapters.MigrateHybridLockReleaseLiquidityInput, seq_core.OnChainOutput, adapters.MigrateHybridLockReleaseLiquidityDeps] {
	return operations.NewSequence(
		"migrate-hybrid-lock-release-liquidity-unsupported",
		common_utils.Version_1_6_0,
		"Solana CCTP chains do not support hybrid lock-release liquidity migration",
		func(b operations.Bundle, deps adapters.MigrateHybridLockReleaseLiquidityDeps, input adapters.MigrateHybridLockReleaseLiquidityInput) (seq_core.OnChainOutput, error) {
			return seq_core.OnChainOutput{}, fmt.Errorf("liquidity migration is not supported on Solana CCTP chains")
		},
	)
}

// CCTPV1AllowedCallerOnDest returns the Solana pool signer, which is the allowed caller that invokes Circle's
// receive-message path on Solana for the current CCTP v1 integration.
func (c *SolanaCCTPChainAdapter) CCTPV1AllowedCallerOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	state, err := c.resolveState(d, b, chainSelector)
	if err != nil {
		return nil, err
	}
	return state.poolSigner.Bytes(), nil
}

// CCTPV2AllowedCallerOnDest is unsupported for Solana until the Solana adapter supports CCTP v2.
func (c *SolanaCCTPChainAdapter) CCTPV2AllowedCallerOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return nil, fmt.Errorf("chain with selector %d does not support CCTP V2", chainSelector)
}

// AllowedCallerOnSource returns the Solana pool signer, which is the caller that performs deposit-for-burn on source.
func (c *SolanaCCTPChainAdapter) AllowedCallerOnSource(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	state, err := c.resolveState(d, b, chainSelector)
	if err != nil {
		return nil, err
	}
	return state.poolSigner.Bytes(), nil
}

// MintRecipientOnDest returns the Solana pool ATA. Circle mints into the token pool ATA first on Solana.
func (c *SolanaCCTPChainAdapter) MintRecipientOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	state, err := c.resolveState(d, b, chainSelector)
	if err != nil {
		return nil, err
	}
	return state.poolTokenATA.Bytes(), nil
}

// USDCType returns the type of the USDC on the chain.
func (c *SolanaCCTPChainAdapter) USDCType() adapters.USDCType {
	return adapters.Canonical
}

// PoolAddress returns the Solana pool signer bytes. For Solana token pools, the signer PDA identifies the remote pool.
func (c *SolanaCCTPChainAdapter) PoolAddress(d datastore.DataStore, b chain.BlockChains, chainSelector uint64, registeredPoolRef datastore.AddressRef) ([]byte, error) {
	tokenPool, err := datastore_utils.FindAndFormatRef(d, registeredPoolRef, chainSelector, sol_utils.ToAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve Solana token pool address: %w", err)
	}

	tokenMint, err := c.findTokenMintForPool(d, b, chainSelector, tokenPool)
	if err != nil {
		return nil, err
	}
	poolSigner, err := sol_token_utils.TokenPoolSignerAddress(tokenMint, tokenPool)
	if err != nil {
		return nil, fmt.Errorf("failed to derive Solana pool signer: %w", err)
	}
	return poolSigner.Bytes(), nil
}

// TokenAddress returns the Solana USDC mint bytes for the configured CCTP pool.
func (c *SolanaCCTPChainAdapter) TokenAddress(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	state, err := c.resolveState(d, b, chainSelector)
	if err != nil {
		return nil, err
	}
	return state.tokenMint.Bytes(), nil
}

func (c *SolanaCCTPChainAdapter) resolveState(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) (solanaCCTPState, error) {
	solChain, ok := b.SolanaChains()[chainSelector]
	if !ok {
		return solanaCCTPState{}, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	cctpPoolRefs := d.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(datastore.ContractType(common_utils.CCTPTokenPool)),
	)
	if len(cctpPoolRefs) != 1 {
		return solanaCCTPState{}, fmt.Errorf("expected exactly 1 Solana CCTP token pool on chain %d, found %d", chainSelector, len(cctpPoolRefs))
	}

	tokenPool, err := sol_utils.ToAddress(cctpPoolRefs[0])
	if err != nil {
		return solanaCCTPState{}, fmt.Errorf("failed to decode Solana CCTP token pool address: %w", err)
	}
	tokenMint, err := c.findTokenMintForPool(d, b, chainSelector, tokenPool)
	if err != nil {
		return solanaCCTPState{}, err
	}
	poolSigner, err := sol_token_utils.TokenPoolSignerAddress(tokenMint, tokenPool)
	if err != nil {
		return solanaCCTPState{}, fmt.Errorf("failed to derive Solana pool signer: %w", err)
	}
	tokenProgram, err := sol_utils.FetchTokenProgramID(context.Background(), solChain, tokenMint)
	if err != nil {
		return solanaCCTPState{}, fmt.Errorf("failed to resolve Solana token program for mint %s: %w", tokenMint, err)
	}
	poolTokenATA, _, err := sol_token_utils.FindAssociatedTokenAddress(tokenProgram, tokenMint, poolSigner)
	if err != nil {
		return solanaCCTPState{}, fmt.Errorf("failed to derive Solana pool ATA: %w", err)
	}

	return solanaCCTPState{
		chain:        solChain,
		tokenPool:    tokenPool,
		tokenMint:    tokenMint,
		tokenProgram: tokenProgram,
		poolSigner:   poolSigner,
		poolTokenATA: poolTokenATA,
	}, nil
}

func (c *SolanaCCTPChainAdapter) findTokenMintForPool(d datastore.DataStore, b chain.BlockChains, chainSelector uint64, tokenPool solana.PublicKey) (solana.PublicKey, error) {
	solChain, ok := b.SolanaChains()[chainSelector]
	if !ok {
		return solana.PublicKey{}, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	candidateTokenRefs := d.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector))
	for _, ref := range candidateTokenRefs {
		if ref.Type != datastore.ContractType(sol_utils.SPLTokens) && ref.Type != datastore.ContractType(sol_utils.SPL2022Tokens) {
			continue
		}

		tokenMint, err := sol_utils.ToAddress(ref)
		if err != nil {
			continue
		}
		poolConfigPDA, err := sol_token_utils.TokenPoolConfigAddress(tokenMint, tokenPool)
		if err != nil {
			return solana.PublicKey{}, fmt.Errorf("failed to derive Solana pool config PDA for mint %s: %w", tokenMint, err)
		}

		var poolState cctp_token_pool.State
		if err := solChain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &poolState); err == nil {
			return tokenMint, nil
		}
	}

	return solana.PublicKey{}, fmt.Errorf("failed to find Solana token mint backing CCTP token pool %s on chain %d", tokenPool, chainSelector)
}

func diffSolanaRemotePoolAddresses(existing, desired []cctp_token_pool.RemoteAddress) []cctp_token_pool.RemoteAddress {
	diff := make([]cctp_token_pool.RemoteAddress, 0)
	for _, desiredAddr := range desired {
		found := false
		for _, existingAddr := range existing {
			if bytes.Equal(existingAddr.Address, desiredAddr.Address) {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, desiredAddr)
		}
	}
	return diff
}
