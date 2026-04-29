package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ConfigureRMNCurseAdminsInput holds the parameters for updating authorized callers (curse admins) on an existing RMN contract.
type ConfigureRMNCurseAdminsInput struct {
	ChainSelector uint64
	RMNAddress    common.Address
	Args          rmnops.AuthorizedCallerArgs
}

// ConfigureRMNCurseAdmins applies authorized caller (curse admin) updates to an already-deployed RMN contract.
var ConfigureRMNCurseAdmins = cldf_ops.NewSequence(
	"configure-rmn-curse-admins",
	semver.MustParse("1.0.0"),
	"Applies authorized caller (curse admin) updates to the RMN contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureRMNCurseAdminsInput) (output sequences.OnChainOutput, err error) {
		if input.ChainSelector != chain.Selector {
			return sequences.OnChainOutput{}, fmt.Errorf("input chain selector %d does not match chain %d",
				input.ChainSelector, chain.Selector)
		}
		if len(input.Args.AddedCallers) == 0 && len(input.Args.RemovedCallers) == 0 {
			return sequences.OnChainOutput{}, nil
		}
		report, err := cldf_ops.ExecuteOperation(
			b, rmnops.ApplyAuthorizedCallerUpdates, chain,
			contract.FunctionInput[rmnops.AuthorizedCallerArgs]{
				ChainSelector: chain.Selector,
				Address:       input.RMNAddress,
				Args:          input.Args,
			})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to RMN(%s) on chain %d: %w",
				input.RMNAddress.Hex(), chain.Selector, err)
		}
		batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		output.BatchOps = []mcms_types.BatchOperation{batch}
		return output, nil
	},
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
