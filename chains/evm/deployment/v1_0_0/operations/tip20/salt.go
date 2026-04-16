package tip20

import (
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const MaxSaltGenerationAttempts = 10

func generateValidSalt(b operations.Bundle, chain evm.Chain, factoryAddr common.Address, deployer common.Address) ([32]byte, error) {
	for range MaxSaltGenerationAttempts {
		var salt [32]byte
		if _, err := rand.Read(salt[:]); err != nil {
			return [32]byte{}, fmt.Errorf("failed to generate random salt: %w", err)
		}

		// NOTE: GetTokenAddress will revert if the salt is already used, so we can
		// use it to check if the salt is valid without actually deploying a token.
		_, err := operations.ExecuteOperation(b, GetTokenAddress, chain, contract.FunctionInput[GetTokenAddressArgs]{
			ChainSelector: chain.Selector,
			Address:       factoryAddr,
			Args: GetTokenAddressArgs{
				Sender: deployer,
				Salt:   salt,
			},
		})
		if err == nil {
			return salt, nil
		}
	}

	return [32]byte{}, fmt.Errorf("could not find a non-reserved salt after %d attempts", MaxSaltGenerationAttempts)
}
