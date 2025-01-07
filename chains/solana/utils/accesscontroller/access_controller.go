package accesscontroller

import (
	"bytes"
	"context"
	"fmt"
	"slices"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func HasAccess(ctx context.Context, client *rpc.Client, accessController solana.PublicKey, address solana.PublicKey, commitment rpc.CommitmentType) (bool, error) {
	var ac access_controller.AccessController
	err := common.GetAccountDataBorshInto(
		ctx,
		client,
		accessController,
		commitment,
		&ac,
	)
	if err != nil {
		return false, fmt.Errorf("failed to get account data: %w", err)
	}
	findInSortedList := func(list []solana.PublicKey, target solana.PublicKey) (int, bool) {
		return slices.BinarySearchFunc(list, target, func(a, b solana.PublicKey) int {
			return bytes.Compare(a.Bytes(), b.Bytes())
		})
	}
	_, found := findInSortedList(ac.AccessList.Xs[:ac.AccessList.Len], address)
	return found, nil
}
