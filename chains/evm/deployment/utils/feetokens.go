package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

// EVMFeeTokenAddresses converts the family-agnostic fee token withdrawal list into the address
// slice that EVM's withdrawFeeTokens(address[]) expects.
//
// The on-chain call always sweeps the contract's entire balance of each token, so a caller-supplied
// Amount cannot be honoured. Rather than withdraw more than was asked for, this rejects any entry
// carrying an amount.
func EVMFeeTokenAddresses(feeTokens []fees.FeeTokenWithdrawal, chainSelector uint64) ([]common.Address, error) {
	if len(feeTokens) == 0 {
		return nil, fmt.Errorf("no fee tokens specified for chain %d", chainSelector)
	}

	addrs := make([]common.Address, 0, len(feeTokens))
	for i, tok := range feeTokens {
		if !common.IsHexAddress(tok.Token) {
			return nil, fmt.Errorf("invalid fee token address %q at feeTokens[%d] on chain %d", tok.Token, i, chainSelector)
		}
		if tok.Amount != nil {
			return nil, fmt.Errorf(
				"partial withdrawals are not supported on EVM: fee token %s at feeTokens[%d] on chain %d specifies an amount, "+
					"but withdrawFeeTokens always sweeps the full balance; omit the amount to proceed",
				tok.Token, i, chainSelector)
		}
		addrs = append(addrs, common.HexToAddress(tok.Token))
	}

	return addrs, nil
}
