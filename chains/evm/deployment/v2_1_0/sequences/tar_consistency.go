package sequences

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
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

// rampReader reads the TokenAdminRegistry configured on one ramp contract version.
type rampReader struct {
	contractType string
	version      *semver.Version
	read         func(b cldf_ops.Bundle, chain evm.Chain, address common.Address) (common.Address, error)
}

func (r rampReader) name() string { return fmt.Sprintf("%s %s", r.contractType, r.version) }

func (r rampReader) matches(ref datastore.AddressRef) bool {
	return string(ref.Type) == r.contractType && ref.Version != nil && ref.Version.Equal(r.version)
}

// newRampReader adapts a ramp's getStaticConfig read. The read bypasses the operations cache.
func newRampReader[CFG any](
	contractType string,
	version *semver.Version,
	op *cldf_ops.Operation[contract.FunctionInput[struct{}], CFG, evm.Chain],
	tokenAdminRegistry func(CFG) common.Address,
) rampReader {
	return rampReader{
		contractType: contractType,
		version:      version,
		read: func(b cldf_ops.Bundle, chain evm.Chain, address common.Address) (common.Address, error) {
			report, err := cldf_ops.ExecuteOperation(b, op, chain,
				contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: address},
				cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
			if err != nil {
				return common.Address{}, err
			}
			return tokenAdminRegistry(report.Output), nil
		},
	}
}

var rampContractTypes = []string{string(onramp_v1_6_0.ContractType), string(offramp_v1_6_0.ContractType)}

// rampReaders lists the ramp versions whose TokenAdminRegistry can be read. The PoolConfigManager
// constructor does not check its TAR against the ramps'.
var rampReaders = []rampReader{
	newRampReader(string(onramp_v1_6_0.ContractType), onramp_v1_6_0.Version, onramp_v1_6_0.GetStaticConfig,
		func(c onramp_v1_6_0.StaticConfig) common.Address { return c.TokenAdminRegistry }),
	newRampReader(string(onramp_v2_0_0.ContractType), onramp_v2_0_0.Version, onramp_v2_0_0.GetStaticConfig,
		func(c onramp_v2_0_0.StaticConfig) common.Address { return c.TokenAdminRegistry }),
	newRampReader(string(offramp_v1_6_0.ContractType), offramp_v1_6_0.Version, offramp_v1_6_0.GetStaticConfig,
		func(c offramp_v1_6_0.StaticConfig) common.Address { return c.TokenAdminRegistry }),
	newRampReader(string(offramp_v2_0_0.ContractType), offramp_v2_0_0.Version, offramp_v2_0_0.GetStaticConfig,
		func(c offramp_v2_0_0.StaticConfig) common.Address { return c.TokenAdminRegistry }),
}

// checkTokenAdminRegistryMatchesRamps requires the newest recorded version of each ramp type
// (OnRamp, OffRamp) on the chain to use the TokenAdminRegistry the PoolConfigManager will be
// deployed with. Older recorded versions are ignored because retired ramps stay in the datastore.
// Chains with no recorded ramps are not checked.
func checkTokenAdminRegistryMatchesRamps(
	b cldf_ops.Bundle,
	chain evm.Chain,
	existing []datastore.AddressRef,
	tokenAdminRegistry common.Address,
) error {
	for _, contractType := range rampContractTypes {
		reader, ok := newestRecordedRampReader(contractType, existing)
		if !ok {
			continue
		}
		if b.GetContext == nil {
			return errors.New("an operations bundle is required to read the ramps' TokenAdminRegistry")
		}
		for _, ref := range existing {
			if !reader.matches(ref) || !common.IsHexAddress(ref.Address) {
				continue
			}
			rampTAR, err := reader.read(b, chain, common.HexToAddress(ref.Address))
			if err != nil {
				return fmt.Errorf("failed to read the TokenAdminRegistry of %s at %s on chain %d: %w", reader.name(), ref.Address, chain.Selector, err)
			}
			if rampTAR != tokenAdminRegistry {
				return fmt.Errorf(
					"%s at %s on chain %d uses TokenAdminRegistry %s but the datastore TokenAdminRegistry is %s; the PoolConfigManager must use the ramps' TokenAdminRegistry",
					reader.name(), ref.Address, chain.Selector, rampTAR, tokenAdminRegistry)
			}
		}
	}
	return nil
}

// newestRecordedRampReader returns the reader for the highest version of contractType that has
// a ref in existing.
func newestRecordedRampReader(contractType string, existing []datastore.AddressRef) (rampReader, bool) {
	var newest rampReader
	found := false
	for _, reader := range rampReaders {
		if reader.contractType != contractType {
			continue
		}
		recorded := false
		for _, ref := range existing {
			if reader.matches(ref) {
				recorded = true
				break
			}
		}
		if recorded && (!found || reader.version.GreaterThan(newest.version)) {
			newest, found = reader, true
		}
	}
	return newest, found
}
