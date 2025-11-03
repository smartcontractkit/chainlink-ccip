package sequences

import (
	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	laneapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func init() {
	v, err := semver.NewVersion("1.6.0")
	if err != nil {
		panic(err)
	}
	laneapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
	deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilySolana, v, &SolanaAdapter{})
	deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
	mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilySolana, &SolanaAdapter{})
}

type SolanaAdapter struct {
	timelockAddr map[uint64]solana.PublicKey
}

func (a *SolanaAdapter) GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return a.GetRouterAddress(ds, chainSelector)
}

func (a *SolanaAdapter) GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(offramp.ContractType),
		Version:       offramp.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       fee_quoter.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
