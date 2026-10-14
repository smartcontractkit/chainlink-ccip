//revive:disable:var-naming // legacy package name
package common

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
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
