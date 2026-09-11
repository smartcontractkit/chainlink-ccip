package changesets

import (
	"fmt"

	"github.com/gagliardetto/solana-go"

	solanautils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/token_pools"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// TokenMintAuthorityUpdate changes the mint authority of a single token to NewMintAuthority
// (typically a token multisig). The token mint authority transfer is signed by the BnM pool
// program's *upgrade authority* - this is the only account that's authorized to call the on
// chain transfer instruction.
type TokenMintAuthorityUpdate struct {
	TokenPoolRef     cldf_datastore.AddressRef `json:"tokenPoolRef"`
	TokenMint        solana.PublicKey          `json:"tokenMint"`
	NewMintAuthority solana.PublicKey          `json:"newMintAuthority"`
}

// TransferMintAuthoritiesChangesetInput updates the mint authority for a batch of tokens (multiple
// pool/mint/authority triples) in a single proposal.
type TransferMintAuthoritiesChangesetInput struct {
	ChainSelector uint64                     `json:"chainSelector"`
	Updates       []TokenMintAuthorityUpdate `json:"updates"`
	MCMS          mcms.Input                 `json:"mcms"`
}

func TransferMintAuthoritiesChangeset() cldf.ChangeSetV2[TransferMintAuthoritiesChangesetInput] {
	return cldf.CreateChangeSet(transferMintAuthoritiesApply, transferMintAuthoritiesVerify)
}

func transferMintAuthoritiesVerify(env cldf.Environment, input TransferMintAuthoritiesChangesetInput) error {
	if err := cldf.IsValidChainSelector(input.ChainSelector); err != nil {
		return fmt.Errorf("invalid chain selector: %d - %w", input.ChainSelector, err)
	}
	if !env.BlockChains.Exists(input.ChainSelector) {
		return fmt.Errorf("chain with selector %d does not exist", input.ChainSelector)
	}
	if len(input.Updates) == 0 {
		return fmt.Errorf("at least one update is required")
	}
	seen := make(map[solana.PublicKey]struct{}, len(input.Updates))
	for i, u := range input.Updates {
		if datastore_utils.IsAddressRefEmpty(u.TokenPoolRef) {
			return fmt.Errorf("update[%d]: tokenPoolRef must not be empty", i)
		}
		if u.TokenMint.IsZero() {
			return fmt.Errorf("update[%d]: tokenMint must not be zero", i)
		}
		if u.NewMintAuthority.IsZero() {
			return fmt.Errorf("update[%d]: newMintAuthority must not be zero", i)
		}
		fullRef, err := datastore_utils.FindAndFormatRef(env.DataStore, u.TokenPoolRef, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("update[%d]: failed to resolve token pool ref: %w", i, err)
		}
		if fullRef.Type.String() != common_utils.BurnMintTokenPool.String() {
			return fmt.Errorf("update[%d]: mint authority transfer is only supported for BurnMint token pools, but pooled ref has type '%s'", i, fullRef.Type.String())
		}
		if _, ok := seen[u.TokenMint]; ok {
			return fmt.Errorf("update[%d]: duplicate tokenMint %s", i, u.TokenMint)
		}
		seen[u.TokenMint] = struct{}{}
	}
	if err := input.MCMS.Validate(); err != nil {
		return fmt.Errorf("invalid MCMS configuration: %w", err)
	}
	return nil
}

func transferMintAuthoritiesApply(e cldf.Environment, input TransferMintAuthoritiesChangesetInput) (cldf.ChangesetOutput, error) {
	chain, ok := e.BlockChains.SolanaChains()[input.ChainSelector]
	if !ok {
		return cldf.ChangesetOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
	}

	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	ds := cldf_datastore.NewMemoryDataStore()
	if err := ds.Merge(e.DataStore); err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge environment datastore: %w", err)
	}

	// Preflight every update before executing any of them, so the deployer-authority path
	// (which confirms each update immediately) can't partially apply a batch before a later
	// update fails.
	for i, u := range input.Updates {
		poolPubkey, err := datastore_utils.FindAndFormatRef(e.DataStore, u.TokenPoolRef, input.ChainSelector, solanautils.ToAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("update[%d]: failed to resolve token pool ref: %w", i, err)
		}
		poolSigner, _ := tokens.TokenPoolSignerAddress(u.TokenMint, poolPubkey)
		if current := solanautils.GetTokenMintAuthority(chain, u.TokenMint); current != poolSigner && current != u.NewMintAuthority {
			return cldf.ChangesetOutput{}, fmt.Errorf("update[%d]: current mint authority %s is neither the pool signer PDA %s nor the target %s", i, current, poolSigner, u.NewMintAuthority)
		}
	}

	for i, u := range input.Updates {
		poolPubkey, err := datastore_utils.FindAndFormatRef(e.DataStore, u.TokenPoolRef, input.ChainSelector, solanautils.ToAddress)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("update[%d]: failed to resolve token pool ref: %w", i, err)
		}

		tokenProg, err := solanautils.FetchTokenProgramID(e.GetContext(), chain, u.TokenMint)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("update[%d]: failed to resolve token program for mint %s: %w", i, u.TokenMint, err)
		}

		report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pools.TransferMintAuthorityBurnMint, chain, token_pools.Params{
			TokenPool:        poolPubkey,
			TokenMint:        u.TokenMint,
			TokenProgramID:   tokenProg,
			NewMintAuthority: u.NewMintAuthority,
		})
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("update[%d]: failed to transfer mint authority: %w", i, err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ToGenericReport())
		for _, addr := range report.Output.Addresses {
			if err := ds.Addresses().Add(addr); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("update[%d]: failed to add address to datastore: %w", i, err)
			}
		}
	}

	return changesets.NewOutputBuilder(e, changesets.GetRegistry()).
		WithReports(reports).
		WithDataStore(ds).
		WithSingleBatchOpPerChain(batchOps).
		Build(input.MCMS)
}
