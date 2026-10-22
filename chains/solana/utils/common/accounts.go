//revive:disable:var-naming // legacy package name
package common

import (
	"context"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	ccip_common_v0_1_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_0/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/ccip_common"
)

// BatchGetAccountsArgs configures a batched account fetch. PDAs are fetched in
// chunks of at most AccountMax accounts per RPC call, and OnFound is invoked for
// every account that exists, with its index in PDAs. Missing accounts are skipped.
// Accounts are passed through raw (not decoded).
type BatchGetAccountsArgs struct {
	Commitment rpc.CommitmentType
	AccountMax int
	PDAs       solana.PublicKeySlice
	OnFound    func(i int, account *rpc.Account) error
}

// BatchGetAccounts fetches the given PDAs in batches. It is a thin helper over
// GetMultipleAccountsWithOpts that avoids exceeding the RPC's per-call account
// limit.
func BatchGetAccounts(ctx context.Context, client *rpc.Client, args BatchGetAccountsArgs) error {
	if args.AccountMax <= 0 {
		return fmt.Errorf("account max must be greater than zero, got %d", args.AccountMax)
	}

	for start := 0; start < len(args.PDAs); start += args.AccountMax {
		end := min(start+args.AccountMax, len(args.PDAs))

		res, err := client.GetMultipleAccountsWithOpts(ctx, args.PDAs[start:end], &rpc.GetMultipleAccountsOpts{
			Commitment: args.Commitment,
		})
		if err != nil {
			return fmt.Errorf("failed to batch-fetch accounts: %w", err)
		}

		for i, acct := range res.Value {
			if acct == nil {
				continue
			}
			if err := args.OnFound(start+i, acct); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetTokenAdminRegistry fetches and decodes a ccip-common TokenAdminRegistry account.
//
// Anchor does not rewrite account bytes on program upgrade, so accounts created before the
// trailing supports_auto_derivation field was added remain the 169-byte V1 layout. The
// on-chain program reads both layouts (load_token_admin_registry_checked in
// ccip-common/src/v1.rs), but the generated Go bindings only describe V2, so we fall back to
// the v0.1.0 binding for V1 accounts. Returns rpc.ErrNotFound when the account does not exist.
func GetTokenAdminRegistry(ctx context.Context, client *rpc.Client, pda solana.PublicKey, commitment rpc.CommitmentType) (ccip_common.TokenAdminRegistry, error) {
	resp, err := client.GetAccountInfoWithOpts(ctx, pda, &rpc.GetAccountInfoOpts{
		Commitment: commitment,
		DataSlice:  nil,
	})
	if err != nil {
		return ccip_common.TokenAdminRegistry{}, err
	}
	if resp == nil || resp.Value == nil {
		return ccip_common.TokenAdminRegistry{}, rpc.ErrNotFound
	}

	return decodeTokenAdminRegistry(resp.Value.Data.GetBinary())
}

// decodeTokenAdminRegistry decodes the current V2 layout, falling back to the legacy V1
// layout. V1 is a strict prefix of V2 (only the trailing supports_auto_derivation bool is
// missing), which the on-chain TryFrom<TokenAdminRegistryV1> defaults to false.
func decodeTokenAdminRegistry(data []byte) (ccip_common.TokenAdminRegistry, error) {
	var v2 ccip_common.TokenAdminRegistry
	v2Err := v2.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	if v2Err == nil {
		return v2, nil
	}

	var v1 ccip_common_v0_1_0.TokenAdminRegistry
	if v1Err := v1.UnmarshalWithDecoder(bin.NewBorshDecoder(data)); v1Err != nil {
		// Surface the V2 error: it describes the current layout callers expect.
		return ccip_common.TokenAdminRegistry{}, v2Err
	}

	return ccip_common.TokenAdminRegistry{
		Version:                v1.Version,
		Administrator:          v1.Administrator,
		PendingAdministrator:   v1.PendingAdministrator,
		LookupTable:            v1.LookupTable,
		WritableIndexes:        v1.WritableIndexes,
		Mint:                   v1.Mint,
		SupportsAutoDerivation: false, // not present in V1; matches on-chain TryFrom
	}, nil
}
