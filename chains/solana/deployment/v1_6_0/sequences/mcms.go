package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	mcmsops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/mcms"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"
)

func (d *SolanaAdapter) DeployMCMS() *operations.Sequence[ccipapi.MCMSDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"deploy-mcms",
		semver.MustParse("1.0.0"),
		"Deploys all MCM contracts with config",
		func(b operations.Bundle, chains cldf_chain.BlockChains, in ccipapi.MCMSDeploymentConfigPerChainWithAddress) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.SolanaChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			// Deploy Access Controller, MCMs and Timelock
			accessControllerRef, err := operations.ExecuteOperation(b, mcmsops.AccessControllerDeploy, chain, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy Access Controller: %w", err)
			}
			output.Addresses = append(output.Addresses, accessControllerRef.Output)

			mcmRef, err := operations.ExecuteOperation(b, mcmsops.McmDeploy, chain, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy MCMs: %w", err)
			}
			output.Addresses = append(output.Addresses, mcmRef.Output)

			timelockRef, err := operations.ExecuteOperation(b, mcmsops.TimelockDeploy, chain, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy Timelock: %w", err)
			}
			output.Addresses = append(output.Addresses, timelockRef.Output)

			accessControllerAddress := solana.MustPublicKeyFromBase58(accessControllerRef.Output.Address)
			mcmAddress := solana.MustPublicKeyFromBase58(mcmRef.Output.Address)
			timelockAddress := solana.MustPublicKeyFromBase58(timelockRef.Output.Address)

			deps := mcmsops.Deps{
				Chain:             chain,
				ExistingAddresses: append(in.ExistingAddresses, output.Addresses...),
				Qualifier:         *in.Qualifier,
			}

			// Initialize Access Controller
			initAccessRef, err := initAccessController(b, deps, accessControllerAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize Access Controller: %w", err)
			}
			output.Addresses = append(output.Addresses, initAccessRef...)
			deps.ExistingAddresses = append(deps.ExistingAddresses, initAccessRef...)

			initMcmRef, err := initMCM(b, deps, in.MCMSDeploymentConfigPerChain, mcmAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize MCMs: %w", err)
			}
			output.Addresses = append(output.Addresses, initMcmRef...)
			deps.ExistingAddresses = append(deps.ExistingAddresses, initMcmRef...)

			initTimelockRef, err := initTimelock(b, deps, in.TimelockMinDelay, timelockAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize Timelock: %w", err)
			}
			output.Addresses = append(output.Addresses, initTimelockRef...)
			deps.ExistingAddresses = append(deps.ExistingAddresses, initTimelockRef...)

			// roles
			err = setupRoles(b, deps, mcmAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to setup roles in Timelock: %w", err)
			}

			return output, err
		},
	)
}

func initAccessController(b operations.Bundle, deps mcmsops.Deps, accessController solana.PublicKey) ([]cldf_datastore.AddressRef, error) {
	roles := []cldf_deployment.ContractType{
		mcmsops.ProposerAccessControllerAccount,
		mcmsops.ExecutorAccessControllerAccount,
		mcmsops.CancellerAccessControllerAccount,
		mcmsops.BypasserAccessControllerAccount,
	}
	var refs []cldf_datastore.AddressRef
	for _, role := range roles {
		ref, err := operations.ExecuteOperation(b, mcmsops.InitAccessControllerOp, deps,
			mcmsops.InitAccessControllerInput{
				ContractType:     role,
				ChainSel:         deps.Chain.ChainSelector(),
				AccessController: accessController,
			})
		if err != nil {
			return nil, fmt.Errorf("failed to init access controller for role %s: %w", role, err)
		}
		refs = append(refs, ref.Output)
	}

	return refs, nil
}

func initMCM(b operations.Bundle, deps mcmsops.Deps, cfg ccipapi.MCMSDeploymentConfigPerChain, mcmAddress solana.PublicKey) ([]cldf_datastore.AddressRef, error) {
	var refs []cldf_datastore.AddressRef
	configs := []struct {
		ctype cldf_deployment.ContractType
		cfg   types.Config
	}{
		{
			utils.BypasserManyChainMultisig,
			cfg.Bypasser,
		},
		{
			utils.CancellerManyChainMultisig,
			cfg.Canceller,
		},
		{
			utils.ProposerManyChainMultisig,
			cfg.Proposer,
		},
	}

	for _, cfg := range configs {
		ref, err := operations.ExecuteOperation(b, mcmsops.InitMCMOp, deps,
			mcmsops.InitMCMInput{
				ContractType: cfg.ctype,
				MCMConfig:    cfg.cfg,
				ChainSel:     deps.Chain.ChainSelector(),
				MCM:          mcmAddress,
			})
		if err != nil {
			return nil, fmt.Errorf("failed to init config type:%q, err:%w", cfg.ctype, err)
		}
		refs = append(refs, ref.Output)
	}
	return refs, nil
}

func initTimelock(b operations.Bundle, deps mcmsops.Deps, minDelay *big.Int, timelockAddress solana.PublicKey) ([]cldf_datastore.AddressRef, error) {
	ref, err := operations.ExecuteOperation(b, mcmsops.InitTimelockOp, deps, mcmsops.InitTimelockInput{
		ContractType: utils.RBACTimelock,
		ChainSel:     deps.Chain.ChainSelector(),
		MinDelay:     minDelay,
		Timelock:     timelockAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init timelock: %w", err)
	}
	return ref.Output, nil
}

func setupRoles(b operations.Bundle, deps mcmsops.Deps, mcmProgram solana.PublicKey) error {
	proposerRef := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.ChainSelector(),
		utils.ProposerManyChainMultisig,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	cancellerRef := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.ChainSelector(),
		utils.CancellerManyChainMultisig,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	bypasserRef := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.ChainSelector(),
		utils.BypasserManyChainMultisig,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	proposerPDA := state.GetMCMSignerPDA(mcmProgram, state.PDASeed([]byte(proposerRef.Address)))
	cancellerPDA := state.GetMCMSignerPDA(mcmProgram, state.PDASeed([]byte(cancellerRef.Address)))
	bypasserPDA := state.GetMCMSignerPDA(mcmProgram, state.PDASeed([]byte(bypasserRef.Address)))
	roles := []struct {
		pdas []solana.PublicKey
		role timelock.Role
	}{
		{
			role: timelock.Proposer_Role,
			pdas: []solana.PublicKey{proposerPDA},
		},
		{
			role: timelock.Executor_Role,
			pdas: []solana.PublicKey{deps.Chain.DeployerKey.PublicKey()},
		},
		{
			role: timelock.Canceller_Role,
			pdas: []solana.PublicKey{cancellerPDA, proposerPDA, bypasserPDA},
		},
		{
			role: timelock.Bypasser_Role,
			pdas: []solana.PublicKey{bypasserPDA},
		},
	}
	for _, role := range roles {
		_, err := operations.ExecuteOperation(b, mcmsops.AddAccessOp, deps, mcmsops.AddAccessInput{
			Role:     role.role,
			Accounts: role.pdas,
		})
		if err != nil {
			return fmt.Errorf("failed to add access for role %d: %w", role.role, err)
		}
	}
	return nil
}
