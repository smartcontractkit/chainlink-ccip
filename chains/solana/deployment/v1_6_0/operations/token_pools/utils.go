package token_pools

import (
	"bytes"
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_0/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
)

func GetAuthorityBurnMint(chain cldf_solana.Chain, program solana.PublicKey, tokenMint solana.PublicKey) solana.PublicKey {
	programData := burnmint_token_pool.State{}
	poolConfigPDA, _ := tokens.TokenPoolConfigAddress(tokenMint, program)
	err := chain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Config.Owner
}

func GetAuthorityLockRelease(chain cldf_solana.Chain, program solana.PublicKey, tokenMint solana.PublicKey) solana.PublicKey {
	programData := lockrelease_token_pool.State{}
	poolConfigPDA, _ := tokens.TokenPoolConfigAddress(tokenMint, program)
	err := chain.GetAccountDataBorshInto(context.Background(), poolConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Config.Owner
}

// get diff of pool addresses
func poolDiff(existingPoolAddresses []base_token_pool.RemoteAddress, newPoolAddresses []base_token_pool.RemoteAddress) []base_token_pool.RemoteAddress {
	var result []base_token_pool.RemoteAddress
	// for every new address, check if it exists in the existing pool addresses
	for _, newAddr := range newPoolAddresses {
		exists := false
		for _, existingAddr := range existingPoolAddresses {
			if bytes.Equal(existingAddr.Address, newAddr.Address) {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, newAddr)
		}
	}
	return result
}
