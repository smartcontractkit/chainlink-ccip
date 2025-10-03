package mcms

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	owner_contracts "github.com/smartcontractkit/ccip-owner-contracts/pkg/gethwrappers"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type MCMSReader struct{}

func (*MCMSReader) GetCurrentOpCount(e deployment.Environment, chainSelector uint64, mcmAddress string) (uint64, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return 0, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	mcm, err := owner_contracts.NewManyChainMultiSig(common.HexToAddress(mcmAddress), chain.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to instantiate MCM at address %s on chain %s: %w", mcmAddress, chain, err)
	}
	opCount, err := mcm.GetOpCount(&bind.CallOpts{Context: e.GetContext()})
	if err != nil {
		return 0, fmt.Errorf("failed to get op count from MCM at address %s on chain %s: %w", mcmAddress, chain, err)
	}
	return opCount.Uint64(), nil
}
