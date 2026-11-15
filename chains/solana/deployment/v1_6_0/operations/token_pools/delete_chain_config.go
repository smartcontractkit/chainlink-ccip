package token_pools

import (
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

// DeleteChainConfigInput is the input for deleting a Solana token pool's remote chain config for one
// remote chain. This is the Solana equivalent of EVM's applyChainUpdates([selector], []): it removes
// the whole config (remote pools, remote token, decimals and rate limits), not just a pool entry.
type DeleteChainConfigInput struct {
	TokenPool      solana.PublicKey
	TokenMint      solana.PublicKey
	RemoteSelector uint64
}

// deleteChainConfig reads the pool's remote chain config for one selector and deletes it. A missing
// config account is the normal "lane already retired" state, so it is skipped (with a warning)
// rather than erroring, keeping re-runs idempotent. The authority and instruction builders are
// supplied by the caller so the burnmint and lockrelease programs share the logic.
func deleteChainConfig(
	b operations.Bundle,
	chain cldf_solana.Chain,
	input DeleteChainConfigInput,
	poolTypeLabel string,
	getAuthority func(cldf_solana.Chain, solana.PublicKey, solana.PublicKey) (solana.PublicKey, error),
	buildDeleteIx func(remoteSelector uint64, mint, poolConfigPDA, remoteChainConfigPDA, authority solana.PublicKey) (solana.Instruction, error),
) (sequences.OnChainOutput, error) {
	authority, err := getAuthority(chain, input.TokenPool, input.TokenMint)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get authority for token pool: %w", err)
	}

	remoteChainConfigPDA, _, _ := tokens.TokenPoolChainConfigPDA(input.RemoteSelector, input.TokenMint, input.TokenPool)
	poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)

	// The on-chain account is a ChainConfig (8-byte discriminator + BaseChain), not a bare
	// BaseChain. ChainConfig has the same discriminator and layout across burnmint/lockrelease
	// pools, so either binding decodes any pool's account. Only existence matters here.
	var remoteChainConfigAccount burnmint_token_pool.ChainConfig
	if err = chain.GetAccountDataBorshInto(b.GetContext(), remoteChainConfigPDA, &remoteChainConfigAccount); err != nil {
		if errors.Is(err, rpc.ErrNotFound) {
			b.Logger.Warnf("skipping deletion of remote chain config for remote chain %d from pool %s on chain %d: remote chain not configured on this pool", input.RemoteSelector, input.TokenPool.String(), chain.Selector)
			return sequences.OnChainOutput{}, nil
		}
		return sequences.OnChainOutput{}, fmt.Errorf("failed to decode remote chain config at PDA %s on chain %d for remote %d: %w", remoteChainConfigPDA, chain.Selector, input.RemoteSelector, err)
	}

	ixn, err := buildDeleteIx(
		input.RemoteSelector,
		input.TokenMint,
		poolConfigPDA,
		remoteChainConfigPDA,
		authority,
	)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	if authority != chain.DeployerKey.PublicKey() {
		batch, err := utils.BuildMCMSBatchOperation(
			chain.Selector,
			[]solana.Instruction{ixn},
			input.TokenPool.String(),
			poolTypeLabel,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batch}}, nil
	}

	if err := chain.Confirm([]solana.Instruction{ixn}); err != nil {
		return sequences.OnChainOutput{}, err
	}
	return sequences.OnChainOutput{}, nil
}

var DeleteChainConfigBurnMint = operations.NewOperation(
	"burnmint:delete_chain_config",
	common_utils.Version_1_6_0,
	"Deletes the BurnMintTokenPool remote chain config",
	func(b operations.Bundle, chain cldf_solana.Chain, input DeleteChainConfigInput) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.TokenPool)
		return deleteChainConfig(b, chain, input, common_utils.BurnMintTokenPool.String(), GetAuthorityBurnMint,
			func(remoteSelector uint64, mint, poolConfigPDA, remoteChainConfigPDA, authority solana.PublicKey) (solana.Instruction, error) {
				return burnmint_token_pool.NewDeleteChainConfigInstruction(remoteSelector, mint, poolConfigPDA, remoteChainConfigPDA, authority).ValidateAndBuild()
			},
		)
	},
)

var DeleteChainConfigLockRelease = operations.NewOperation(
	"lockrelease:delete_chain_config",
	common_utils.Version_1_6_0,
	"Deletes the LockReleaseTokenPool remote chain config",
	func(b operations.Bundle, chain cldf_solana.Chain, input DeleteChainConfigInput) (sequences.OnChainOutput, error) {
		lockrelease_token_pool.SetProgramID(input.TokenPool)
		return deleteChainConfig(b, chain, input, common_utils.LockReleaseTokenPool.String(), GetAuthorityLockRelease,
			func(remoteSelector uint64, mint, poolConfigPDA, remoteChainConfigPDA, authority solana.PublicKey) (solana.Instruction, error) {
				return lockrelease_token_pool.NewDeleteChainConfigInstruction(remoteSelector, mint, poolConfigPDA, remoteChainConfigPDA, authority).ValidateAndBuild()
			},
		)
	},
)
