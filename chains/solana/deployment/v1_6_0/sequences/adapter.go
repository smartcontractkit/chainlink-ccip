package sequences

import (
	"encoding/binary"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	laneapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
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
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
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

func (a *SolanaAdapter) GetRMNRemoteAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(rmnremoteops.ContractType),
		Version:       rmnremoteops.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// svmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector SVM")) = 0x1e10bdc4.
var svmFamilySelector = [4]byte{0x1e, 0x10, 0xbd, 0xc4}

func (a *SolanaAdapter) GetFeeQuoterDestChainConfig() laneapi.FeeQuoterDestChainConfig {
	return laneapi.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                1_280,
		MaxPerMsgGasLimit:           400_000,
		DestGasOverhead:             300_000,
		ChainFamilySelector:         binary.BigEndian.Uint32(svmFamilySelector[:]),
		DefaultTokenFeeUSDCents:     35,
		DefaultTokenDestGasOverhead: 150_000,
		DefaultTxGasLimit:           1, // irrelevant for Solana
		NetworkFeeUSDCents:          10,
		V1Params: &laneapi.FeeQuoterV1Params{
			MaxNumberOfTokensPerMsg:    1,
			GasMultiplierWeiPerEth:     11e17,
			GasPriceStalenessThreshold: 90_000,
			EnforceOutOfOrder:          true,
		},
		V2Params: &laneapi.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: 90,
			USDPerUnitGas:            big.NewInt(1e6),
		},
	}
}

func (a *SolanaAdapter) GetDefaultGasPrice() *big.Int {
	return big.NewInt(4e12)
}

func (a *SolanaAdapter) GetChainFamilySelector() [4]byte {
	return svmFamilySelector
}
