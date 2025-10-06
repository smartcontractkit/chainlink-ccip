package changesets

import (
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

func TestDeriveMCMAddressesEVM(t *testing.T) {
	chains := []chain.BlockChain{
		evm.Chain{Selector: 1},
	}
	mcmAddressRef := &datastore.AddressRef{
		Type:      "ProposerManyChainMultiSig",
		Version:   semver.MustParse("1.0.0"),
		Qualifier: "main",
	}
	timelockAddressRef := &datastore.AddressRef{
		Type:      "RBACTimelock",
		Version:   semver.MustParse("1.0.0"),
		Qualifier: "main",
	}
	input := &MCMSInput{
		OverridePreviousRoot: false,
		ValidUntil:           2756219818,
		TimelockDelay:        mcms_types.NewDuration(3 * time.Hour),
		TimelockAction:       mcms_types.TimelockActionSchedule,
		MCMSAddressRef:       mcmAddressRef,
		TimelockAddressRef:   timelockAddressRef,
	}
	lggr, err := logger.New()
	require.NoError(t, err)
	ds := datastore.NewMemoryDataStore()
	// need to add the addresses to the datastore so they can be found
	timelockAddressRef.Address = common.HexToAddress("0x0000000000000000000000000000000000000001").Hex()
	timelockAddressRef.ChainSelector = 1

	err = ds.Addresses().Add(*timelockAddressRef)
	require.NoError(t, err)

	e := deployment.Environment{DataStore: ds.Seal(), Logger: lggr, BlockChains: chain.NewBlockChainsFromSlice(chains)}
	mcmBuilder := NewEVMMCMBuilder(input)
	timelockAddrs, err := mcmBuilder.DeriveTimelocks(e)
	require.NoError(t, err)
	require.Equal(t, map[mcms_types.ChainSelector]string{
		1: common.HexToAddress("0x0000000000000000000000000000000000000001").Hex(),
	}, timelockAddrs)
}
