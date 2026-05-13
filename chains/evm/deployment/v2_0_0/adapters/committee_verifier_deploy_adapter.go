package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

// EVMCommitteeVerifierDeployAdapter implements
// ccvdeploymentadapters.CommitteeVerifierDeployAdapter for EVM chains by
// proxying to the existing sequences.DeployCommitteeVerifier sequence. It
// converts the chain-agnostic input shape into the EVM-specific input,
// resolving the RMNProxy address from the caller-supplied existing
// addresses.
type EVMCommitteeVerifierDeployAdapter struct{}

var _ ccvdeploymentadapters.CommitteeVerifierDeployAdapter = (*EVMCommitteeVerifierDeployAdapter)(nil)

var evmDeployCommitteeVerifier = cldf_ops.NewSequence(
	"evm-deploy-committee-verifier",
	semver.MustParse("2.0.0"),
	"Chain-agnostic wrapper around the EVM DeployCommitteeVerifier sequence",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input ccvdeploymentadapters.DeployCommitteeVerifierInput) (ccvdeploymentadapters.DeployCommitteeVerifierOutput, error) {
		evmChains := chains.EVMChains()
		evmChain, ok := evmChains[input.ChainSelector]
		if !ok {
			return ccvdeploymentadapters.DeployCommitteeVerifierOutput{},
				fmt.Errorf("EVM chain not found for selector %d", input.ChainSelector)
		}

		create2Factory, err := parseRequiredHexAddress(input.DeployerContract, "DeployerContract")
		if err != nil {
			return ccvdeploymentadapters.DeployCommitteeVerifierOutput{}, err
		}

		rmnAddr, err := resolveRMNProxyAddress(input.ExistingAddresses, input.ChainSelector)
		if err != nil {
			return ccvdeploymentadapters.DeployCommitteeVerifierOutput{}, err
		}

		feeAgg, err := parseRequiredNonZeroHexAddress(
			input.Params.FeeAggregator,
			fmt.Sprintf("committee %q FeeAggregator", input.Params.Qualifier),
		)
		if err != nil {
			return ccvdeploymentadapters.DeployCommitteeVerifierOutput{}, err
		}

		var allowlistAdmin common.Address
		if input.Params.AllowlistAdmin != "" {
			allowlistAdmin, err = parseHexAddress(
				input.Params.AllowlistAdmin,
				fmt.Sprintf("committee %q AllowlistAdmin", input.Params.Qualifier),
			)
			if err != nil {
				return ccvdeploymentadapters.DeployCommitteeVerifierOutput{}, err
			}
		}

		evmInput := sequences.DeployCommitteeVerifierInput{
			ChainSelector:     input.ChainSelector,
			ExistingAddresses: input.ExistingAddresses,
			CREATE2Factory:    create2Factory,
			RMN:               rmnAddr,
			Params: sequences.CommitteeVerifierParams{
				Version:          input.Params.Version,
				AllowlistAdmin:   allowlistAdmin,
				FeeAggregator:    feeAgg,
				StorageLocations: input.Params.StorageLocations,
				Qualifier:        input.Params.Qualifier,
			},
		}

		report, err := cldf_ops.ExecuteSequence(b, sequences.DeployCommitteeVerifier, evmChain, evmInput)
		if err != nil {
			return ccvdeploymentadapters.DeployCommitteeVerifierOutput{},
				fmt.Errorf("EVM DeployCommitteeVerifier failed: %w", err)
		}

		out := ccvdeploymentadapters.DeployCommitteeVerifierOutput{
			Addresses: report.Output.Addresses,
			BatchOps:  report.Output.BatchOps,
		}
		if !input.DeployerKeyOwned {
			// Ownership transfer is not yet wired through this chain-agnostic
			// flow. Surface every deployed contract so a future caller can
			// build the transfer step, mirroring how the full DeployChainContracts
			// sequence enumerates ownableContracts.
			out.RefsToTransferOwnership = report.Output.Addresses
		}
		return out, nil
	},
)

func (a *EVMCommitteeVerifierDeployAdapter) DeployCommitteeVerifier() *cldf_ops.Sequence[ccvdeploymentadapters.DeployCommitteeVerifierInput, ccvdeploymentadapters.DeployCommitteeVerifierOutput, cldf_chain.BlockChains] {
	return evmDeployCommitteeVerifier
}

// resolveRMNProxyAddress finds the RMNProxy (ARMProxy contract type) on the
// given chain in existingAddresses. The EVM CommitteeVerifier constructor
// requires this address, so callers must ensure the proxy is already
// deployed when invoking the standalone CCV deploy.
func resolveRMNProxyAddress(existingAddresses []datastore.AddressRef, chainSelector uint64) (common.Address, error) {
	for _, ref := range existingAddresses {
		if ref.ChainSelector != chainSelector {
			continue
		}
		if ref.Type != datastore.ContractType(rmn_proxy.ContractType) {
			continue
		}
		if ref.Version == nil || !ref.Version.Equal(rmn_proxy.Version) {
			continue
		}
		if !common.IsHexAddress(ref.Address) {
			return common.Address{}, fmt.Errorf("RMNProxy address %q on chain %d is not a valid hex address", ref.Address, chainSelector)
		}
		return common.HexToAddress(ref.Address), nil
	}
	return common.Address{}, fmt.Errorf("RMNProxy (type=%s version=%s) not found in ExistingAddresses for chain %d", rmn_proxy.ContractType, rmn_proxy.Version, chainSelector)
}
