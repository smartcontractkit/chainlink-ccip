package common

import (
	"context"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// BatchGetAccountsArgs configures a batched account fetch. PDAs are fetched in
// chunks of at most AccountMax accounts per RPC call, and OnFound is invoked for
// every account that exists and decodes successfully. Accounts that are missing
// or fail to decode are skipped.
type BatchGetAccountsArgs[T any] struct {
	Commitment rpc.CommitmentType
	AccountMax int
	PDAs       solana.PublicKeySlice
	OnFound    func(i int, accountState T) error
}

// BatchGetAccounts fetches the given PDAs in batches and decodes each account
// into T using a Borsh decoder. It is a thin helper over
// GetMultipleAccountsWithOpts that avoids exceeding the RPC's per-call account
// limit and tolerates missing/undecodable accounts.
func BatchGetAccounts[T any](ctx context.Context, client *rpc.Client, args BatchGetAccountsArgs[T]) error {
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

			var accountState T
			if err := bin.NewBorshDecoder(acct.Data.GetBinary()).Decode(&accountState); err != nil {
				continue
			}

			if err := args.OnFound(start+i, accountState); err != nil {
				return err
			}
		}
	}

	return nil
}
