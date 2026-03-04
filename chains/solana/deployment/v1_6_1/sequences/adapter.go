package sequences

import (
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/operations/rmn_remote"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type SolanaAdapter struct {
	timelockAddr map[uint64]solana.PublicKey
}

func (a *SolanaAdapter) GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version, // TODO import 1.6.1 version when it exists
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetRMNRemoteAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(rmnremoteops.ContractType),
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
