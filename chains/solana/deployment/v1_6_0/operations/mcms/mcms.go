package mcms

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	"github.com/smartcontractkit/mcms/types"
)

var (
	AccessControllerProgramType cldf_deployment.ContractType = "AccessControllerProgram"
	AccessControllerProgramName                              = "access_controller"
	AccessControllerProgramSize                              = 1 * 1024 * 1024

	TimelockProgramName = "timelock"
	TimelockProgramSize = 1 * 1024 * 1024

	McmProgramName                              = "mcm"
	McmProgramSize                              = 1 * 1024 * 1024

	ProposerAccessControllerAccount  cldf_deployment.ContractType = "ProposerAccessControllerAccount"
	ExecutorAccessControllerAccount  cldf_deployment.ContractType = "ExecutorAccessControllerAccount"
	CancellerAccessControllerAccount cldf_deployment.ContractType = "CancellerAccessControllerAccount"
	BypasserAccessControllerAccount  cldf_deployment.ContractType = "BypasserAccessControllerAccount"
)

var AccessControllerDeploy = operations.NewOperation(
	"access-controller:deploy",
	common_utils.Version_1_6_0,
	"Deploys the Access Controller program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []cldf_datastore.AddressRef) (cldf_datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			AccessControllerProgramType,
			common_utils.Version_1_6_0,
			"",
			AccessControllerProgramName,
			AccessControllerProgramSize)
	},
)

var TimelockDeploy = operations.NewOperation(
	"timelock:deploy",
	common_utils.Version_1_6_0,
	"Deploys the Timelock program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []cldf_datastore.AddressRef) (cldf_datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			utils.TimelockProgramType,
			common_utils.Version_1_6_0,
			"",
			TimelockProgramName,
			TimelockProgramSize)
	},
)

var McmDeploy = operations.NewOperation(
	"mcm:deploy",
	common_utils.Version_1_6_0,
	"Deploys the Many Chain Multi Sig program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []cldf_datastore.AddressRef) (cldf_datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			utils.McmProgramType,
			common_utils.Version_1_6_0,
			"",
			McmProgramName,
			McmProgramSize)
	},
)

var InitAccessControllerOp = operations.NewOperation(
	"init-access-controller",
	common_utils.Version_1_6_0,
	"Initializes access controller for solana",
	initAccessController,
)

var InitMCMOp = operations.NewOperation(
	"init-mcm-program",
	common_utils.Version_1_6_0,
	"Initializes MCMProgram for solana",
	initMCM,
)

var InitTimelockOp = operations.NewOperation(
	"init-timelock-program",
	common_utils.Version_1_6_0,
	"Initializes timelock for solana",
	initTimelock,
)

var AddAccessOp = operations.NewOperation(
	"add-access-op",
	common_utils.Version_1_6_0,
	"Adds access to provided role for timelock",
	addAccess,
)

type (
	Deps struct {
		Chain             cldf_solana.Chain
		ExistingAddresses []cldf_datastore.AddressRef
		Qualifier         string
	}

	InitAccessControllerInput struct {
		ContractType     cldf_deployment.ContractType
		ChainSel         uint64
		AccessController solana.PublicKey
	}

	InitMCMInput struct {
		ContractType cldf_deployment.ContractType
		MCMConfig    types.Config
		ChainSel     uint64
		MCM          solana.PublicKey
	}

	InitTimelockInput struct {
		ContractType cldf_deployment.ContractType
		ChainSel     uint64
		MinDelay     *big.Int
		Timelock     solana.PublicKey
	}

	AddAccessInput struct {
		Qualifier string
		Role      timelock.Role
		Accounts  []solana.PublicKey
		ChainSel  uint64
	}
)

func initAccessController(b operations.Bundle, deps Deps, in InitAccessControllerInput) (cldf_datastore.AddressRef, error) {
	access_controller.SetProgramID(in.AccessController)
	// Should be one of the AccessControllerAccount types
	ref := datastore.GetAddressRef(
		deps.ExistingAddresses,
		in.ChainSel,
		in.ContractType,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)

	var accessControllerAccount solana.PublicKey
	if ref.Address != "" {
		accessControllerAccount = solana.PublicKeyFromBytes([]byte(ref.Address))
		var data access_controller.AccessController
		err := common.GetAccountDataBorshInto(b.GetContext(), deps.Chain.Client, accessControllerAccount, rpc.CommitmentConfirmed, &data)
		if err == nil {
			b.Logger.Infow("access controller already initialized, skipping initialization", "chain", deps.Chain.String())
			return cldf_datastore.AddressRef{}, nil
		}

		return cldf_datastore.AddressRef{}, fmt.Errorf("unable to read access controller account config %s", accessControllerAccount.String())
	}

	b.Logger.Infow("access controller not initialized, initializing", "chain", deps.Chain.String())

	account, err := solana.NewRandomPrivateKey()
	if err != nil {
		return cldf_datastore.AddressRef{}, fmt.Errorf("failed to generate new random private key for access controller account: %w", err)
	}

	err = initializeAccessController(b, deps.Chain, in.AccessController, account)
	if err != nil {
		return cldf_datastore.AddressRef{}, fmt.Errorf("failed to initialize access controller: %w", err)
	}

	b.Logger.Infow("initialized access controller", "account", account.PublicKey())

	return cldf_datastore.AddressRef{
		Address:       account.PublicKey().String(),
		ChainSelector: deps.Chain.Selector,
		Type:          cldf_datastore.ContractType(in.ContractType),
		Qualifier:     deps.Qualifier,
		Version:       common_utils.Version_1_6_0,
	}, nil
}

// discriminator + owner + proposed owner + access_list (64 max addresses + length)
const accessControllerAccountSize = uint64(8 + 32 + 32 + ((32 * 64) + 8))

func initializeAccessController(
	b operations.Bundle, chain cldf_solana.Chain, programID solana.PublicKey, roleAccount solana.PrivateKey,
) error {
	rentExemption, err := chain.Client.GetMinimumBalanceForRentExemption(b.GetContext(),
		accessControllerAccountSize, rpc.CommitmentConfirmed)
	if err != nil {
		return fmt.Errorf("failed to get minimum balance for rent exemption: %w", err)
	}

	createAccountInstruction, err := system.NewCreateAccountInstruction(rentExemption, accessControllerAccountSize,
		programID, chain.DeployerKey.PublicKey(), roleAccount.PublicKey()).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to create CreateAccount instruction: %w", err)
	}

	initializeInstruction, err := access_controller.NewInitializeInstruction(
		roleAccount.PublicKey(),
		chain.DeployerKey.PublicKey(),
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to build instruction: %w", err)
	}

	instructions := []solana.Instruction{createAccountInstruction, initializeInstruction}
	err = chain.Confirm(instructions, common.AddSigners(roleAccount))
	if err != nil {
		return fmt.Errorf("failed to confirm CreateAccount and InitializeAccessController instructions: %w", err)
	}

	var data access_controller.AccessController
	err = common.GetAccountDataBorshInto(b.GetContext(), chain.Client, roleAccount.PublicKey(), rpc.CommitmentConfirmed, &data)
	if err != nil {
		return fmt.Errorf("failed to read access controller roleAccount: %w", err)
	}

	return nil
}

func initMCM(b operations.Bundle, deps Deps, in InitMCMInput) (cldf_datastore.AddressRef, error) {
	mcm.SetProgramID(in.MCM)
	// Should be one of:
	// BypasserManyChainMultisig
	// CancellerManyChainMultisig
	// ProposerManyChainMultisig
	ref := datastore.GetAddressRef(
		deps.ExistingAddresses,
		in.ChainSel,
		in.ContractType,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)

	var mcmSeed state.PDASeed
	if ref.Address != "" {
		mcmSeed = state.PDASeed([]byte(ref.Address))
		mcmConfigPDA := state.GetMCMConfigPDA(in.MCM, mcmSeed)
		var data mcm.MultisigConfig
		err := common.GetAccountDataBorshInto(b.GetContext(), deps.Chain.Client, mcmConfigPDA, rpc.CommitmentConfirmed, &data)
		if err == nil {
			b.Logger.Infow("mcm config already initialized, skipping initialization", "chain", deps.Chain.String())
			return cldf_datastore.AddressRef{}, nil
		}
		return cldf_datastore.AddressRef{}, fmt.Errorf("unable to read mcm ConfigPDA account config %q", mcmConfigPDA.String())
	}

	b.Logger.Infow("mcm config not initialized, initializing", "chain", deps.Chain.String())

	seed := randomSeed()
	b.Logger.Infow("generated MCM seed", "seed", string(seed[:]))
	err := initializeMCM(b, deps, in.MCM, seed)
	if err != nil {
		return cldf_datastore.AddressRef{}, fmt.Errorf("failed to initialize mcm: %w", err)
	}

	encodedAddress := mcms_solana.ContractAddress(in.MCM, mcms_solana.PDASeed(seed))

	configurer := mcms_solana.NewConfigurer(deps.Chain.Client, *deps.Chain.DeployerKey, types.ChainSelector(deps.Chain.ChainSelector()))
	tx, err := configurer.SetConfig(b.GetContext(), encodedAddress, &in.MCMConfig, false)
	if err != nil {
		return cldf_datastore.AddressRef{}, fmt.Errorf("failed to set config on mcm: %w", err)
	}
	b.Logger.Infow("called SetConfig on MCM", "transaction", tx.Hash)

	return cldf_datastore.AddressRef{
		Address:       string(seed[:]),
		ChainSelector: deps.Chain.Selector,
		Type:          cldf_datastore.ContractType(in.ContractType),
		Qualifier:     deps.Qualifier,
		Version:       common_utils.Version_1_6_0,
	}, nil
}

func initializeMCM(b operations.Bundle, deps Deps, mcmProgram solana.PublicKey, multisigID state.PDASeed) error {
	var mcmConfig mcm.MultisigConfig
	err := deps.Chain.GetAccountDataBorshInto(b.GetContext(), state.GetMCMConfigPDA(mcmProgram, multisigID), &mcmConfig)
	if err == nil {
		b.Logger.Infow("MCM already initialized, skipping initialization", "chain", deps.Chain.String())
		return nil
	}

	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	opts := &rpc.GetAccountInfoOpts{Commitment: rpc.CommitmentConfirmed}

	data, err := deps.Chain.Client.GetAccountInfoWithOpts(b.GetContext(), mcmProgram, opts)
	if err != nil {
		return fmt.Errorf("failed to get mcm program account info: %w", err)
	}
	err = bin.UnmarshalBorsh(&programData, data.Bytes())
	if err != nil {
		return fmt.Errorf("failed to unmarshal program data: %w", err)
	}

	instruction, err := mcm.NewInitializeInstruction(
		deps.Chain.Selector,
		multisigID,
		state.GetMCMConfigPDA(mcmProgram, multisigID),
		deps.Chain.DeployerKey.PublicKey(),
		solana.SystemProgramID,
		mcmProgram,
		programData.Address,
		state.GetMCMRootMetadataPDA(mcmProgram, multisigID),
		state.GetMCMExpiringRootAndOpCountPDA(mcmProgram, multisigID),
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to build instruction: %w", err)
	}

	err = deps.Chain.Confirm([]solana.Instruction{instruction})
	if err != nil {
		return fmt.Errorf("failed to confirm instructions: %w", err)
	}

	return nil
}

func initTimelock(b operations.Bundle, deps Deps, in InitTimelockInput) ([]cldf_datastore.AddressRef, error) {
	timelock.SetProgramID(in.Timelock)
	// Should be one of:
	// RBACTimelock
	ref := datastore.GetAddressRef(
		deps.ExistingAddresses,
		in.ChainSel,
		in.ContractType,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)

	var timelockSeed state.PDASeed
	if ref.Address != "" {
		timelockSeed = state.PDASeed([]byte(ref.Address))
		timelockConfigPDA := state.GetTimelockConfigPDA(in.Timelock, timelockSeed)
		var timelockConfig timelock.Config
		err := deps.Chain.GetAccountDataBorshInto(b.GetContext(), timelockConfigPDA, &timelockConfig)
		if err == nil {
			b.Logger.Infow("timelock config already initialized, skipping initialization", "chain", deps.Chain.String())
			return []cldf_datastore.AddressRef{}, nil
		}
		return []cldf_datastore.AddressRef{}, fmt.Errorf("unable to read timelock ConfigPDA account config %s", timelockConfigPDA.String())
	}

	b.Logger.Infow("timelock config not initialized, initializing", "chain", deps.Chain.String())

	seed := randomSeed()
	b.Logger.Infow("generated Timelock seed", "seed", string(seed[:]))

	err := initializeTimelock(b, deps, in.Timelock, seed, in.MinDelay)
	if err != nil {
		return []cldf_datastore.AddressRef{}, fmt.Errorf("failed to initialize timelock: %w", err)
	}

	return []cldf_datastore.AddressRef{
		{
			Address: mcms_solana.ContractAddress(
				solana.MustPublicKeyFromBase58(in.Timelock.String()),
				mcms_solana.PDASeed([]byte(seed[:])),
			),
			ChainSelector: deps.Chain.Selector,
			Type:          cldf_datastore.ContractType(utils.TimelockCompositeAddress),
			Qualifier:     deps.Qualifier,
			Version:       common_utils.Version_1_6_0,
		},
		{
			Address:       string(seed[:]),
			ChainSelector: deps.Chain.Selector,
			Type:          cldf_datastore.ContractType(in.ContractType),
			Qualifier:     deps.Qualifier,
			Version:       common_utils.Version_1_6_0,
		},
	}, nil
}

func initializeTimelock(b operations.Bundle, deps Deps, timelockProgram solana.PublicKey,
	timelockID state.PDASeed, minDelay *big.Int) error {
	if minDelay == nil {
		minDelay = big.NewInt(0)
	}

	var timelockConfig timelock.Config
	err := deps.Chain.GetAccountDataBorshInto(b.GetContext(), state.GetTimelockConfigPDA(timelockProgram, timelockID),
		&timelockConfig)
	if err == nil {
		b.Logger.Infow("Timelock already initialized, skipping initialization", "chain", deps.Chain.String())
		return nil
	}

	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	opts := &rpc.GetAccountInfoOpts{Commitment: rpc.CommitmentConfirmed}

	data, err := deps.Chain.Client.GetAccountInfoWithOpts(b.GetContext(), timelockProgram, opts)
	if err != nil {
		return fmt.Errorf("failed to get timelock program account info: %w", err)
	}
	err = bin.UnmarshalBorsh(&programData, data.Bytes())
	if err != nil {
		return fmt.Errorf("failed to unmarshal program data: %w", err)
	}

	accessControllerProgram := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		AccessControllerProgramType,
		common_utils.Version_1_6_0,
		"",
	)
	proposerAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		ProposerAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	executorAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		ExecutorAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	cancellerAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		CancellerAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	bypasserAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		BypasserAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)

	instruction, err := timelock.NewInitializeInstruction(
		timelockID,
		minDelay.Uint64(),
		state.GetTimelockConfigPDA(timelockProgram, timelockID),
		deps.Chain.DeployerKey.PublicKey(),
		solana.SystemProgramID,
		timelockProgram,
		programData.Address,
		solana.MustPublicKeyFromBase58(accessControllerProgram.Address),
		solana.MustPublicKeyFromBase58(proposerAccount.Address),
		solana.MustPublicKeyFromBase58(executorAccount.Address),
		solana.MustPublicKeyFromBase58(cancellerAccount.Address),
		solana.MustPublicKeyFromBase58(bypasserAccount.Address),
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to build instruction: %w", err)
	}

	err = deps.Chain.Confirm([]solana.Instruction{instruction})
	if err != nil {
		return fmt.Errorf("failed to confirm instructions: %w", err)
	}

	return nil
}

func addAccess(b operations.Bundle, deps Deps, in AddAccessInput) (cldf_datastore.AddressRef, error) {
	accessControllerProgram := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		AccessControllerProgramType,
		common_utils.Version_1_6_0,
		in.Qualifier,
	)
	timelockProgram := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.TimelockProgramType,
		common_utils.Version_1_6_0,
		in.Qualifier,
	)
	// timelock seeds stored as a separate program type
	// qualifier will identify the correct timelock instance
	timelockSeed := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0,
		in.Qualifier,
	)
	timelockConfigPDA := state.GetTimelockConfigPDA(
		solana.MustPublicKeyFromBase58(timelockProgram.Address),
		state.PDASeed([]byte(timelockSeed.Address)),
	)
	var roleAccessController cldf_deployment.ContractType
	switch in.Role {
	case timelock.Proposer_Role:
		roleAccessController = ProposerAccessControllerAccount
	case timelock.Executor_Role:
		roleAccessController = ExecutorAccessControllerAccount
	case timelock.Canceller_Role:
		roleAccessController = CancellerAccessControllerAccount
	case timelock.Bypasser_Role:
		roleAccessController = BypasserAccessControllerAccount
	default:
		return cldf_datastore.AddressRef{}, fmt.Errorf("unknown role: %d", in.Role)
	}
	roleAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		roleAccessController,
		common_utils.Version_1_6_0,
		in.Qualifier,
	)
	instructionBuilder := timelock.NewBatchAddAccessInstruction([32]uint8(
		state.PDASeed([]byte(timelockSeed.Address))),
		in.Role,
		timelockConfigPDA,
		solana.MustPublicKeyFromBase58(accessControllerProgram.Address),
		solana.MustPublicKeyFromBase58(roleAccount.Address),
		deps.Chain.DeployerKey.PublicKey())

	for _, account := range in.Accounts {
		instructionBuilder.Append(solana.Meta(account))
	}

	instruction, err := instructionBuilder.ValidateAndBuild()
	if err != nil {
		return cldf_datastore.AddressRef{}, fmt.Errorf("failed to build BatchAddAccess instruction: %w", err)
	}

	err = deps.Chain.Confirm([]solana.Instruction{instruction})
	if err != nil {
		return cldf_datastore.AddressRef{}, fmt.Errorf("failed to confirm BatchAddAccess instruction: %w", err)
	}
	return cldf_datastore.AddressRef{}, nil
}

func randomSeed() state.PDASeed {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	seed := state.PDASeed{}
	for i := range seed {
		seed[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return state.PDASeed(bytes.Trim(seed[:], "\x00"))
}
