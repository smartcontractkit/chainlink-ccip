package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	offramp_v1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onramp_v1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	offramp_v2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	onramp_v2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
)

// rampTokenAdminRegistryReaders reads the TokenAdminRegistry configured on each known ramp
// version. The PoolConfigManager constructor does not verify its TAR against the ramps', and a
// mismatch yields a PCM that never validates transfers, so the check is done here.
var rampTokenAdminRegistryReaders = []struct {
	name    string
	matches func(ref datastore.AddressRef) bool
	read    func(b cldf_ops.Bundle, chain evm.Chain, selector uint64, address common.Address) (common.Address, error)
}{
	{"OnRamp 1.6.0", rampMatcher(string(onramp_v1_6_0.ContractType), onramp_v1_6_0.Version.String()), func(b cldf_ops.Bundle, chain evm.Chain, selector uint64, address common.Address) (common.Address, error) {
		report, err := cldf_ops.ExecuteOperation(b, onramp_v1_6_0.GetStaticConfig, chain, contract.FunctionInput[struct{}]{ChainSelector: selector, Address: address}, cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
		return report.Output.TokenAdminRegistry, err
	}},
	{"OnRamp 2.0.0", rampMatcher(string(onramp_v2_0_0.ContractType), onramp_v2_0_0.Version.String()), func(b cldf_ops.Bundle, chain evm.Chain, selector uint64, address common.Address) (common.Address, error) {
		report, err := cldf_ops.ExecuteOperation(b, onramp_v2_0_0.GetStaticConfig, chain, contract.FunctionInput[struct{}]{ChainSelector: selector, Address: address}, cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
		return report.Output.TokenAdminRegistry, err
	}},
	{"OffRamp 1.6.0", rampMatcher(string(offramp_v1_6_0.ContractType), offramp_v1_6_0.Version.String()), func(b cldf_ops.Bundle, chain evm.Chain, selector uint64, address common.Address) (common.Address, error) {
		report, err := cldf_ops.ExecuteOperation(b, offramp_v1_6_0.GetStaticConfig, chain, contract.FunctionInput[struct{}]{ChainSelector: selector, Address: address}, cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
		return report.Output.TokenAdminRegistry, err
	}},
	{"OffRamp 2.0.0", rampMatcher(string(offramp_v2_0_0.ContractType), offramp_v2_0_0.Version.String()), func(b cldf_ops.Bundle, chain evm.Chain, selector uint64, address common.Address) (common.Address, error) {
		report, err := cldf_ops.ExecuteOperation(b, offramp_v2_0_0.GetStaticConfig, chain, contract.FunctionInput[struct{}]{ChainSelector: selector, Address: address}, cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
		return report.Output.TokenAdminRegistry, err
	}},
}

func rampMatcher(contractType, version string) func(ref datastore.AddressRef) bool {
	return func(ref datastore.AddressRef) bool {
		return string(ref.Type) == contractType && ref.Version != nil && ref.Version.String() == version
	}
}

// checkTokenAdminRegistryMatchesRamps requires every OnRamp and OffRamp recorded in the datastore
// for the chain to use the same TokenAdminRegistry the PoolConfigManager will be deployed with.
// Chains with no recorded ramps are not checked: there is nothing to compare against.
func checkTokenAdminRegistryMatchesRamps(
	b cldf_ops.Bundle,
	chain evm.Chain,
	existing []datastore.AddressRef,
	tokenAdminRegistry common.Address,
) error {
	for _, reader := range rampTokenAdminRegistryReaders {
		for _, ref := range existing {
			if !reader.matches(ref) || !common.IsHexAddress(ref.Address) {
				continue
			}
			rampTAR, err := reader.read(b, chain, chain.Selector, common.HexToAddress(ref.Address))
			if err != nil {
				return fmt.Errorf("failed to read the TokenAdminRegistry of %s at %s on chain %d: %w", reader.name, ref.Address, chain.Selector, err)
			}
			if rampTAR != tokenAdminRegistry {
				return fmt.Errorf(
					"%s at %s on chain %d uses TokenAdminRegistry %s but the datastore TokenAdminRegistry is %s; the PoolConfigManager must use the ramps' TokenAdminRegistry",
					reader.name, ref.Address, chain.Selector, rampTAR, tokenAdminRegistry)
			}
		}
	}
	return nil
}
