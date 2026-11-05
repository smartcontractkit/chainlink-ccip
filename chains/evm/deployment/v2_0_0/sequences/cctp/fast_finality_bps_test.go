package cctp

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/cctp_verifier"
)

// Dynamic config values the verifier already has are not written again, and unset ones are left alone. The write
// itself needs a chain, so it is covered by the deploy and configure tests on a simulated one.
func TestDesiredDynamicConfigWrites_NoWriteWhenUnchanged(t *testing.T) {
	admin := common.HexToAddress("0x2222222222222222222222222222222222222222")
	current := cctp_verifier.DynamicConfig{
		FeeAggregator:   common.HexToAddress("0x14eaEc5b6d2E59db8b680c5A06326504956a25f3"),
		AllowlistAdmin:  admin,
		FastFinalityBps: 50,
	}
	bps := uint16(50)

	for name, tc := range map[string]struct {
		bps   *uint16
		admin *common.Address
	}{
		"same values":  {bps: &bps, admin: &admin},
		"only the fee": {bps: &bps},
		"only admin":   {admin: &admin},
	} {
		t.Run(name, func(t *testing.T) {
			writes, err := desiredDynamicConfigWrites(cldf_ops.Bundle{}, evm.Chain{}, common.Address{}, current, tc.bps, tc.admin)
			require.NoError(t, err)
			require.Empty(t, writes)
		})
	}
}
