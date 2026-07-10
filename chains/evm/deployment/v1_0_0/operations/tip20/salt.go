package tip20

import (
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const MaxSaltGenerationAttempts = 10

func generateValidSalt(b cld_ops.Bundle, chain cldf_evm.Chain, factoryAddr common.Address, deployer common.Address) ([32]byte, error) {
	for range MaxSaltGenerationAttempts {
		var salt [32]byte
		if _, err := rand.Read(salt[:]); err != nil {
			return [32]byte{}, fmt.Errorf("failed to generate random salt: %w", err)
		}

		// NOTE: GetTokenAddress will revert if the salt is already used, so we can
		// use it to check if the salt is valid without actually deploying a token.
		_, err := evmops.ExecuteRead(b, chain, factoryAddr, NewTIP20Factory, NewReadGetTokenAddress, GetTokenAddressArgs{
			Sender: deployer,
			Salt:   salt,
		})
		if err == nil {
			return salt, nil
		}
	}

	return [32]byte{}, fmt.Errorf("could not find a non-reserved salt after %d attempts", MaxSaltGenerationAttempts)
}
