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
		newestVersion, err := newestRecordedRampVersion(contractType, existing)
		if err != nil {
			return fmt.Errorf("cannot determine the newest recorded ramp on chain %d: %w", chain.Selector, err)
		}
		if newestVersion == nil {
			continue
		}
		rampName := fmt.Sprintf("%s %s", contractType, newestVersion)
		reader, err := compatibleRampReader(contractType, newestVersion)
		if err != nil {
			return fmt.Errorf("unsupported newest recorded ramp %s on chain %d: %w", rampName, chain.Selector, err)
		}
		if b.GetContext == nil {
			return errors.New("an operations bundle is required to read the ramps' TokenAdminRegistry")
		}
		for _, ref := range existing {
			if string(ref.Type) != contractType || ref.Version == nil || !ref.Version.Equal(newestVersion) {
				continue
			}
			if !common.IsHexAddress(ref.Address) {
				return fmt.Errorf("newest recorded ramp %s on chain %d has invalid address %q", rampName, chain.Selector, ref.Address)
			}
			rampTAR, err := reader.read(b, chain, common.HexToAddress(ref.Address))
			if err != nil {
				return fmt.Errorf("failed to read the TokenAdminRegistry of %s at %s on chain %d: %w", rampName, ref.Address, chain.Selector, err)
			}
			if rampTAR != tokenAdminRegistry {
				return fmt.Errorf(
					"%s at %s on chain %d uses TokenAdminRegistry %s but the datastore TokenAdminRegistry is %s; the PoolConfigManager must use the ramps' TokenAdminRegistry",
					rampName, ref.Address, chain.Selector, rampTAR, tokenAdminRegistry)
			}
		}
	}
	return nil
}

// newestRecordedRampVersion considers every recorded ref, including versions without an exact
// registered reader. A nil version is malformed and prevents a safe newest-version decision.
func newestRecordedRampVersion(contractType string, existing []datastore.AddressRef) (*semver.Version, error) {
	var newest *semver.Version
	for _, ref := range existing {
		if string(ref.Type) != contractType {
			continue
		}
		if ref.Version == nil {
			return nil, fmt.Errorf("recorded %s at %q has no version", contractType, ref.Address)
		}
		if newest == nil || ref.Version.GreaterThan(newest) {
			newest = ref.Version
		}
	}
	return newest, nil
}

// compatibleRampReader uses the newest registered reader from the same major version. The
// static-config ABI used to read the TAR is compatible across minor and patch releases.
func compatibleRampReader(contractType string, version *semver.Version) (rampReader, error) {
	var compatible rampReader
	found := false
	for _, reader := range rampReaders {
		if reader.contractType != contractType || reader.version.Major() != version.Major() {
			continue
		}
		if !found || reader.version.GreaterThan(compatible.version) {
			compatible = reader
			found = true
		}
	}
	if !found {
		return rampReader{}, fmt.Errorf("no compatible static-config reader is registered for %s %s", contractType, version)
	}
	return compatible, nil
}
