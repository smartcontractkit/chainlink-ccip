package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/sdk/evm"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type SeqMCMSDeploymentCfg struct {
	ContractType      cldf.ContractType
	MCMConfig         *types.Config
	Label             *string
	Qualifier         *string
	ExistingAddresses []datastore.AddressRef
}

var SeqDeployMCMWithConfig = cldf_ops.NewSequence(
	"seq-deploy-mcm-with-config",
	semver.MustParse("1.0.0"),
	"Deploys MCM contract & sets config",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqMCMSDeploymentCfg) (output sequences.OnChainOutput, err error) {
		// Deploy MCM contracts
		var mcmAddr datastore.AddressRef
		switch in.ContractType {
		case utils.ProposerManyChainMultisig:
			mcmAddr, err = contract.MaybeDeployContract(b, ops.OpDeployProposerMCM, chain, contract.DeployInput[struct{}]{
				ChainSelector:  chain.Selector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.ProposerManyChainMultisig, *semver.MustParse("1.0.0")),
			}, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		case utils.BypasserManyChainMultisig:
			mcmAddr, err = contract.MaybeDeployContract(b, ops.OpDeployBypasserMCM, chain, contract.DeployInput[struct{}]{
				ChainSelector:  chain.Selector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.BypasserManyChainMultisig, *semver.MustParse("1.0.0")),
			}, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		case utils.CancellerManyChainMultisig:
			mcmAddr, err = contract.MaybeDeployContract(b, ops.OpDeployCancellerMCM, chain, contract.DeployInput[struct{}]{
				ChainSelector:  chain.Selector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.CancellerManyChainMultisig, *semver.MustParse("1.0.0")),
			}, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported contract type for seq-deploy-mcm-with-config: %s", in.ContractType)
		}

		// Set config
		groupQuorums, groupParents, signerAddresses, signerGroups, err := evm.ExtractSetConfigInputs(in.MCMConfig)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		_, err = cldf_ops.ExecuteOperation(b, ops.OpEVMSetConfigMCM, chain,
			contract.FunctionInput[ops.OpSetConfigMCMInput]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(mcmAddr.Address),
				Args: ops.OpSetConfigMCMInput{
					SignerAddresses: signerAddresses,
					SignerGroups:    signerGroups,
					GroupQuorums:    groupQuorums,
					GroupParents:    groupParents,
				},
			})

		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		b.Logger.Infof("Deployed %s at address %s on chain %s", in.ContractType, mcmAddr.Address, chain.Name)
		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{mcmAddr},
		}, nil
	},
)
