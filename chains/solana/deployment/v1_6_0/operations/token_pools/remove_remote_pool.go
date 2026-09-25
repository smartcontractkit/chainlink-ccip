package token_pools

import (
	"bytes"
	"fmt"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

// RemoveRemotePoolInput is the input for removing a single remote pool address from a Solana
// token pool's remote chain config. The remote pool address is the raw 32-byte public key of
// the remote pool, in the remote chain family's native byte encoding.
type RemoveRemotePoolInput struct {
	TokenPool         solana.PublicKey
	TokenMint         solana.PublicKey
	RemoteSelector    uint64
	RemotePoolAddress []byte
}

// removeRemotePoolTokenPool reads the existing remote chain config, removes the target remote
// pool address, and rewrites the config via EditChainRemoteConfig. It returns a clear error
// when the target remote pool is not currently configured. The authority and instruction
// builders are supplied by the caller so the burnmint and lockrelease programs share the logic.
func removeRemotePoolTokenPool(
	b operations.Bundle,
	chain cldf_solana.Chain,
	input RemoveRemotePoolInput,
	poolTypeLabel string,
	getAuthority func(cldf_solana.Chain, solana.PublicKey, solana.PublicKey) (solana.PublicKey, error),
	buildEditIx func(remoteSelector uint64, mint solana.PublicKey, cfg base_token_pool.RemoteConfig, poolConfigPDA, remoteChainConfigPDA, authority, systemProgram solana.PublicKey) (solana.Instruction, error),
) (sequences.OnChainOutput, error) {
	authority, err := getAuthority(chain, input.TokenPool, input.TokenMint)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get authority for token pool: %w", err)
	}

	poolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
	remoteChainConfigPDA, _, _ := tokens.TokenPoolChainConfigPDA(input.RemoteSelector, input.TokenMint, input.TokenPool)

	// The on-chain account is a ChainConfig (8-byte discriminator + BaseChain), not a bare
	// BaseChain. Decoding into BaseChain consumes the discriminator as the start of the
	// payload and corrupts the parse. ChainConfig has the same discriminator and layout
	// across burnmint/lockrelease pools, so either binding decodes any pool's account.
	var remoteChainConfigAccount burnmint_token_pool.ChainConfig
	err = chain.GetAccountDataBorshInto(b.GetContext(), remoteChainConfigPDA, &remoteChainConfigAccount)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to decode remote chain config at PDA %s on chain %d for remote %d: %w", remoteChainConfigPDA, chain.Selector, input.RemoteSelector, err)
	}

	existing := remoteChainConfigAccount.Base.Remote.PoolAddresses
	target := input.RemotePoolAddress
	remaining := make([]base_token_pool.RemoteAddress, 0, len(existing))
	found := false
	for _, addr := range existing {
		if bytes.Equal(addr.Address, target) {
			found = true
			continue
		}
		remaining = append(remaining, base_token_pool.RemoteAddress{Address: addr.Address})
	}
	if !found {
		return sequences.OnChainOutput{}, fmt.Errorf(
			"remote pool %x is not configured for remote chain %d on pool %s (chain %d)",
			target, input.RemoteSelector, input.TokenPool.String(), chain.Selector,
		)
	}

	remoteConfig := base_token_pool.RemoteConfig{
		PoolAddresses: remaining,
		TokenAddress:  remoteChainConfigAccount.Base.Remote.TokenAddress,
		Decimals:      remoteChainConfigAccount.Base.Remote.Decimals,
	}

	ixn, err := buildEditIx(
		input.RemoteSelector,
		input.TokenMint,
		remoteConfig,
		poolConfigPDA,
		remoteChainConfigPDA,
		authority,
		solana.SystemProgramID,
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

var RemoveRemotePoolBurnMint = operations.NewOperation(
	"burnmint:remove_remote_pool",
	common_utils.Version_1_6_0,
	"Removes a remote pool address from the BurnMintTokenPool remote chain config",
	func(b operations.Bundle, chain cldf_solana.Chain, input RemoveRemotePoolInput) (sequences.OnChainOutput, error) {
		burnmint_token_pool.SetProgramID(input.TokenPool)
		return removeRemotePoolTokenPool(b, chain, input, common_utils.BurnMintTokenPool.String(), GetAuthorityBurnMint,
			func(remoteSelector uint64, mint solana.PublicKey, cfg base_token_pool.RemoteConfig, poolConfigPDA, remoteChainConfigPDA, authority, systemProgram solana.PublicKey) (solana.Instruction, error) {
				return burnmint_token_pool.NewEditChainRemoteConfigInstruction(
					remoteSelector, mint, cfg, poolConfigPDA, remoteChainConfigPDA, authority, systemProgram,
				).ValidateAndBuild()
			},
		)
	},
)

var RemoveRemotePoolLockRelease = operations.NewOperation(
	"lockrelease:remove_remote_pool",
	common_utils.Version_1_6_0,
	"Removes a remote pool address from the LockReleaseTokenPool remote chain config",
	func(b operations.Bundle, chain cldf_solana.Chain, input RemoveRemotePoolInput) (sequences.OnChainOutput, error) {
		lockrelease_token_pool.SetProgramID(input.TokenPool)
		return removeRemotePoolTokenPool(b, chain, input, common_utils.LockReleaseTokenPool.String(), GetAuthorityLockRelease,
			func(remoteSelector uint64, mint solana.PublicKey, cfg base_token_pool.RemoteConfig, poolConfigPDA, remoteChainConfigPDA, authority, systemProgram solana.PublicKey) (solana.Instruction, error) {
				return lockrelease_token_pool.NewEditChainRemoteConfigInstruction(
					remoteSelector, mint, cfg, poolConfigPDA, remoteChainConfigPDA, authority, systemProgram,
				).ValidateAndBuild()
			},
		)
	},
)
