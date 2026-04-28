package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// DeployRMNInput deploys RMN (IRMN onchain implementation used behind RMNProxy / by verifiers).
type DeployRMNInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	Args              rmnops.ConstructorArgs
}

// DeployRMN deploys a new RMN contract or returns the existing address from the datastore.
var DeployRMN = cldf_ops.NewSequence(
	"deploy-rmn",
	semver.MustParse("1.0.0"),
	"Deploys the RMN (curse / IRMN) contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployRMNInput) (output sequences.OnChainOutput, err error) {
		if input.ChainSelector != chain.Selector {
			return sequences.OnChainOutput{}, fmt.Errorf("input chain selector %d does not match chain %d",
				input.ChainSelector, chain.Selector)
		}

		rmnRef, err := contract.MaybeDeployContract(b, rmnops.Deploy, chain, contract.DeployInput[rmnops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmnops.ContractType, *rmnops.Version),
			ChainSelector:  chain.Selector,
			Args:           input.Args,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy RMN: %w", err)
		}
		if rmnRef.Address == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("RMN address is empty after deploy on chain %d", chain.Selector)
		}
		output.Addresses = append(output.Addresses, rmnRef)
		return output, nil
	},
)
