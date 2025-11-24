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
	AccessControllerProgramName = "access_controller"
	AccessControllerProgramSize = 1 * 1024 * 1024

	TimelockProgramName = "timelock"
	TimelockProgramSize = 1 * 1024 * 1024

	McmProgramName = "mcm"
	McmProgramSize = 1 * 1024 * 1024
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
			utils.AccessControllerProgramType,
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

var TransferOwnershipOp = operations.NewOperation(
	"transfer-ownership-op",
	common_utils.Version_1_6_0,
	"Transfers ownership of programs to timelock",
	transferToTimelockSolanaOp,
)

var AcceptOwnershipOp = operations.NewOperation(
	"accept-ownership-op",
	common_utils.Version_1_6_0,
	"Accepts ownership of programs from timelock",
	acceptOwnershipTimelockSolanaOp,
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
		Qualifier        string
	}

	InitMCMInput struct {
		ContractType cldf_deployment.ContractType
		MCMConfig    types.Config
		ChainSel     uint64
		MCM          solana.PublicKey
		Qualifier    string
	}

	InitTimelockInput struct {
		ContractType cldf_deployment.ContractType
		ChainSel     uint64
		MinDelay     *big.Int
		Timelock     solana.PublicKey
		Qualifier    string
	}

	MCMOutput struct {
		NewAddresses []cldf_datastore.AddressRef
		BatchOps     []types.BatchOperation
	}

	AddAccessInput struct {
		Role      timelock.Role
		Accounts  []solana.PublicKey
		ChainSel  uint64
		Qualifier string
	}

	OwnableContract struct {
		ProgramID solana.PublicKey
		Seed      [32]byte
		OwnerPDA  solana.PublicKey
		Type      cldf_deployment.ContractType
	}

	TransferToTimelockInput struct {
		Contract     OwnableContract
		Qualifier    string
		CurrentOwner solana.PublicKey
		NewOwner     solana.PublicKey
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

func initMCM(b operations.Bundle, deps Deps, in InitMCMInput) (MCMOutput, error) {
	mcm.SetProgramID(in.MCM)
	// Should be one of:
	// BypasserSeed
	// CancellerSeed
	// ProposerSeed
	ref := datastore.GetAddressRef(
		deps.ExistingAddresses,
		in.ChainSel,
		in.ContractType,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	var outType cldf_datastore.ContractType
	switch in.ContractType {
	case utils.BypasserSeed:
		outType = cldf_datastore.ContractType(common_utils.BypasserManyChainMultisig)
	case utils.CancellerSeed:
		outType = cldf_datastore.ContractType(common_utils.CancellerManyChainMultisig)
	case utils.ProposerSeed:
		outType = cldf_datastore.ContractType(common_utils.ProposerManyChainMultisig)
	default:
		return MCMOutput{}, fmt.Errorf("unsupported mcm contract type: %s", in.ContractType)
	}

	var mcmSeed state.PDASeed
	if ref.Address != "" {
		mcmSeed = state.PDASeed([]byte(ref.Address))
		mcmConfigPDA := state.GetMCMConfigPDA(in.MCM, mcmSeed)
		var data mcm.MultisigConfig
		err := common.GetAccountDataBorshInto(b.GetContext(), deps.Chain.Client, mcmConfigPDA, rpc.CommitmentConfirmed, &data)
		if err == nil {
			b.Logger.Infow("mcm config already initialized, skipping initialization", "chain", deps.Chain.String())
			return MCMOutput{}, nil
		}
		return MCMOutput{}, fmt.Errorf("unable to read mcm ConfigPDA account config %q", mcmConfigPDA.String())
	}

	b.Logger.Infow("mcm config not initialized, initializing", "chain", deps.Chain.String())

	seed := randomSeed()
	b.Logger.Infow("generated MCM seed", "seed", string(seed[:]))
	ixns, err := initializeMCM(b, deps, in.MCM, seed)
	if err != nil {
		return MCMOutput{}, fmt.Errorf("failed to initialize mcm: %w", err)
	}

	var batches []types.BatchOperation
	if len(ixns) > 0 {
		batch, err := utils.BuildMCMSBatchOperation(
			deps.Chain.Selector,
			ixns,
			in.MCM.String(),
			in.ContractType.String(),
		)
		if err != nil {
			return MCMOutput{}, fmt.Errorf("failed to build timelock initialization batch operation: %w", err)
		}
		batches = append(batches, batch)
	}

	encodedAddress := mcms_solana.ContractAddress(in.MCM, mcms_solana.PDASeed(seed))

	configurer := mcms_solana.NewConfigurer(deps.Chain.Client, *deps.Chain.DeployerKey, types.ChainSelector(deps.Chain.ChainSelector()))
	tx, err := configurer.SetConfig(b.GetContext(), encodedAddress, &in.MCMConfig, false)
	if err != nil {
		return MCMOutput{}, fmt.Errorf("failed to set config on mcm: %w", err)
	}
	b.Logger.Infow("called SetConfig on MCM", "transaction", tx.Hash)

	return MCMOutput{
		BatchOps: batches,
		NewAddresses: []cldf_datastore.AddressRef{
			{
				Address: mcms_solana.ContractAddress(
					solana.MustPublicKeyFromBase58(in.MCM.String()),
					mcms_solana.PDASeed([]byte(seed[:])),
				),
				ChainSelector: deps.Chain.Selector,
				Type:          outType,
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
		}}, nil
}

func initializeMCM(b operations.Bundle, deps Deps, mcmProgram solana.PublicKey, multisigID state.PDASeed) ([]solana.Instruction, error) {
	var mcmConfig mcm.MultisigConfig
	err := deps.Chain.GetAccountDataBorshInto(b.GetContext(), state.GetMCMConfigPDA(mcmProgram, multisigID), &mcmConfig)
	if err == nil {
		b.Logger.Infow("MCM already initialized, skipping initialization", "chain", deps.Chain.String())
		return nil, nil
	}

	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	opts := &rpc.GetAccountInfoOpts{Commitment: rpc.CommitmentConfirmed}

	data, err := deps.Chain.Client.GetAccountInfoWithOpts(b.GetContext(), mcmProgram, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get mcm program account info: %w", err)
	}
	err = bin.UnmarshalBorsh(&programData, data.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal program data: %w", err)
	}

	upgradeAuthority, err := utils.GetUpgradeAuthority(deps.Chain.Client, mcmProgram)
	if err != nil {
		return nil, fmt.Errorf("failed to get upgrade authority: %w", err)
	}

	instruction, err := mcm.NewInitializeInstruction(
		deps.Chain.Selector,
		multisigID,
		state.GetMCMConfigPDA(mcmProgram, multisigID),
		upgradeAuthority,
		solana.SystemProgramID,
		mcmProgram,
		programData.Address,
		state.GetMCMRootMetadataPDA(mcmProgram, multisigID),
		state.GetMCMExpiringRootAndOpCountPDA(mcmProgram, multisigID),
	).ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to build instruction: %w", err)
	}

	if upgradeAuthority == deps.Chain.DeployerKey.PublicKey() {
		err = deps.Chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return nil, fmt.Errorf("failed to confirm instructions: %w", err)
		}
	} else {
		b.Logger.Infow("skipping confirm of initialize instruction as upgrade authority is not deployer key	", "chain", deps.Chain.String())
		return []solana.Instruction{instruction}, nil
	}
	return nil, nil
}

func initTimelock(b operations.Bundle, deps Deps, in InitTimelockInput) (MCMOutput, error) {
	timelock.SetProgramID(in.Timelock)
	// Should be one of:
	// RBACTimelockSeed
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
			return MCMOutput{}, nil
		}
		return MCMOutput{}, fmt.Errorf("unable to read timelock ConfigPDA account config %s", timelockConfigPDA.String())
	}

	b.Logger.Infow("timelock config not initialized, initializing", "chain", deps.Chain.String())

	seed := randomSeed()
	b.Logger.Infow("generated Timelock seed", "seed", string(seed[:]))

	ixns, err := initializeTimelock(b, deps, in.Timelock, seed, in.MinDelay)
	if err != nil {
		return MCMOutput{}, fmt.Errorf("failed to initialize timelock: %w", err)
	}
	var batches []types.BatchOperation
	if len(ixns) > 0 {
		batch, err := utils.BuildMCMSBatchOperation(
			deps.Chain.Selector,
			ixns,
			in.Timelock.String(),
			in.ContractType.String(),
		)
		if err != nil {
			return MCMOutput{}, fmt.Errorf("failed to build timelock initialization batch operation: %w", err)
		}
		batches = append(batches, batch)
	}

	return MCMOutput{
		BatchOps: batches,
		NewAddresses: []cldf_datastore.AddressRef{
			{
				Address: mcms_solana.ContractAddress(
					solana.MustPublicKeyFromBase58(in.Timelock.String()),
					mcms_solana.PDASeed([]byte(seed[:])),
				),
				ChainSelector: deps.Chain.Selector,
				Type:          cldf_datastore.ContractType(common_utils.RBACTimelock),
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
		}}, nil
}

func initializeTimelock(b operations.Bundle, deps Deps, timelockProgram solana.PublicKey,
	timelockID state.PDASeed, minDelay *big.Int) ([]solana.Instruction, error) {
	if minDelay == nil {
		minDelay = big.NewInt(0)
	}

	var timelockConfig timelock.Config
	err := deps.Chain.GetAccountDataBorshInto(b.GetContext(), state.GetTimelockConfigPDA(timelockProgram, timelockID),
		&timelockConfig)
	if err == nil {
		b.Logger.Infow("Timelock already initialized, skipping initialization", "chain", deps.Chain.String())
		return nil, nil
	}

	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	opts := &rpc.GetAccountInfoOpts{Commitment: rpc.CommitmentConfirmed}

	data, err := deps.Chain.Client.GetAccountInfoWithOpts(b.GetContext(), timelockProgram, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get timelock program account info: %w", err)
	}
	err = bin.UnmarshalBorsh(&programData, data.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal program data: %w", err)
	}

	accessControllerProgram := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.AccessControllerProgramType,
		common_utils.Version_1_6_0,
		"",
	)
	proposerAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.ProposerAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	executorAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.ExecutorAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	cancellerAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.CancellerAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)
	bypasserAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.BypasserAccessControllerAccount,
		common_utils.Version_1_6_0,
		deps.Qualifier,
	)

	upgradeAuthority, err := utils.GetUpgradeAuthority(deps.Chain.Client, timelockProgram)
	if err != nil {
		return nil, fmt.Errorf("failed to get upgrade authority: %w", err)
	}

	instruction, err := timelock.NewInitializeInstruction(
		timelockID,
		minDelay.Uint64(),
		state.GetTimelockConfigPDA(timelockProgram, timelockID),
		upgradeAuthority,
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
		return nil, fmt.Errorf("failed to build instruction: %w", err)
	}

	if upgradeAuthority == deps.Chain.DeployerKey.PublicKey() {
		err = deps.Chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return nil, fmt.Errorf("failed to confirm instructions: %w", err)
		}
	} else {
		b.Logger.Infow("skipping confirm of initialize instruction as upgrade authority is not deployer key")
		return []solana.Instruction{instruction}, nil
	}

	return nil, nil
}

func addAccess(b operations.Bundle, deps Deps, in AddAccessInput) (MCMOutput, error) {
	accessControllerProgram := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		utils.AccessControllerProgramType,
		common_utils.Version_1_6_0,
		"",
	)
	timelockProgram := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0,
		in.Qualifier,
	)
	id, seed, _ := mcms_solana.ParseContractAddress(timelockProgram.Address)
	timelockConfigPDA := state.GetTimelockConfigPDA(
		id,
		state.PDASeed([]byte(seed[:])),
	)
	var roleAccessController cldf_deployment.ContractType
	switch in.Role {
	case timelock.Proposer_Role:
		roleAccessController = utils.ProposerAccessControllerAccount
	case timelock.Executor_Role:
		roleAccessController = utils.ExecutorAccessControllerAccount
	case timelock.Canceller_Role:
		roleAccessController = utils.CancellerAccessControllerAccount
	case timelock.Bypasser_Role:
		roleAccessController = utils.BypasserAccessControllerAccount
	default:
		return MCMOutput{}, fmt.Errorf("unknown role: %d", in.Role)
	}
	roleAccount := datastore.GetAddressRef(
		deps.ExistingAddresses,
		deps.Chain.Selector,
		roleAccessController,
		common_utils.Version_1_6_0,
		in.Qualifier,
	)
	upgradeAuthority, err := utils.GetUpgradeAuthority(deps.Chain.Client, id)
	if err != nil {
		return MCMOutput{}, fmt.Errorf("failed to get upgrade authority: %w", err)
	}

	instructionBuilder := timelock.NewBatchAddAccessInstruction([32]uint8(
		state.PDASeed([]byte(seed[:]))),
		in.Role,
		timelockConfigPDA,
		solana.MustPublicKeyFromBase58(accessControllerProgram.Address),
		solana.MustPublicKeyFromBase58(roleAccount.Address),
		upgradeAuthority)

	for _, account := range in.Accounts {
		instructionBuilder.Append(solana.Meta(account))
	}

	instruction, err := instructionBuilder.ValidateAndBuild()
	if err != nil {
		return MCMOutput{}, fmt.Errorf("failed to build BatchAddAccess instruction: %w", err)
	}

	if upgradeAuthority != deps.Chain.DeployerKey.PublicKey() {
		batch, err := utils.BuildMCMSBatchOperation(
			deps.Chain.Selector,
			[]solana.Instruction{instruction},
			id.String(),
			utils.TimelockProgramType.String(),
		)
		if err != nil {
			return MCMOutput{}, fmt.Errorf("failed to build timelock initialization batch operation: %w", err)
		}
		return MCMOutput{
			BatchOps: []types.BatchOperation{batch},
		}, nil
	}

	err = deps.Chain.Confirm([]solana.Instruction{instruction})
	if err != nil {
		return MCMOutput{}, fmt.Errorf("failed to confirm BatchAddAccess instruction: %w", err)
	}
	return MCMOutput{}, nil
}

func transferToTimelockSolanaOp(b operations.Bundle, deps Deps, in TransferToTimelockInput) ([]types.BatchOperation, error) {
	out := make([]types.BatchOperation, 0)

	solChain := deps.Chain

	chainSelector := solChain.ChainSelector()

	contract := in.Contract
	transferInstruction, err := transferOwnershipInstruction(
		contract.ProgramID,
		contract.Seed,
		in.NewOwner,
		contract.OwnerPDA,
		in.CurrentOwner)
	if err != nil {
		return out, fmt.Errorf("failed to create transfer ownership instruction: %w", err)
	}
	if in.CurrentOwner != solChain.DeployerKey.PublicKey() {
		transferBatch, err := utils.BuildMCMSBatchOperation(
			chainSelector,
			[]solana.Instruction{transferInstruction},
			contract.ProgramID.String(),
			string(contract.Type),
		)
		if err != nil {
			return out, fmt.Errorf("failed to build accept ownership mcms transaction: %w", err)
		}
		out = append(out, transferBatch)
	} else {
		err = solChain.Confirm([]solana.Instruction{transferInstruction})
		if err != nil {
			return out, fmt.Errorf("failed to confirm instruction: %w", err)
		}
	}

	return out, nil
}

func acceptOwnershipTimelockSolanaOp(b operations.Bundle, deps Deps, in TransferToTimelockInput) ([]types.BatchOperation, error) {
	out := make([]types.BatchOperation, 0)

	solChain := deps.Chain

	chainSelector := solChain.ChainSelector()

	contract := in.Contract
	acceptMCMSTransaction, err := acceptMCMSTransaction(chainSelector, contract, in.NewOwner)
	if err != nil {
		return out, fmt.Errorf("failed to create accept ownership mcms transaction: %w", err)
	}
	out = append(out, acceptMCMSTransaction)

	return out, nil
}

func randomSeed() state.PDASeed {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	seed := state.PDASeed{}
	for i := range seed {
		seed[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return state.PDASeed(bytes.Trim(seed[:], "\x00"))
}

func transferOwnershipInstruction(
	programID solana.PublicKey, seed state.PDASeed, proposedOwner, ownerPDA, auth solana.PublicKey,
) (solana.Instruction, error) {
	if (seed == state.PDASeed{}) {
		return newSeedlessTransferOwnershipInstruction(programID, proposedOwner, ownerPDA, auth)
	}
	return newSeededTransferOwnershipInstruction(programID, seed, proposedOwner, ownerPDA, auth)
}

func acceptMCMSTransaction(
	selector uint64,
	contract OwnableContract,
	authority solana.PublicKey,
) (types.BatchOperation, error) {
	acceptInstruction, err := acceptOwnershipInstruction(contract.ProgramID, contract.Seed, contract.OwnerPDA, authority)
	if err != nil {
		return types.BatchOperation{}, fmt.Errorf("failed to build accept ownership instruction: %w", err)
	}
	acceptMCMSTx, err := utils.BuildMCMSBatchOperation(
		selector,
		[]solana.Instruction{acceptInstruction},
		contract.ProgramID.String(),
		string(contract.Type),
	)
	if err != nil {
		return types.BatchOperation{}, fmt.Errorf("failed to build accept ownership mcms transaction: %w", err)
	}
	return acceptMCMSTx, nil
}

func acceptOwnershipInstruction(programID solana.PublicKey, seed state.PDASeed, ownerPDA, auth solana.PublicKey,
) (solana.Instruction, error) {
	if (seed == state.PDASeed{}) {
		return newSeedlessAcceptOwnershipInstruction(programID, ownerPDA, auth)
	}
	return newSeededAcceptOwnershipInstruction(programID, seed, ownerPDA, auth)
}

func newSeededTransferOwnershipInstruction(
	programID solana.PublicKey, seed state.PDASeed, proposedOwner, config, authority solana.PublicKey,
) (solana.Instruction, error) {
	ix, err := mcm.NewTransferOwnershipInstruction(seed, proposedOwner, config, authority).ValidateAndBuild()
	return &seededInstruction{ix, programID}, err
}

func newSeededAcceptOwnershipInstruction(
	programID solana.PublicKey, seed state.PDASeed, config, authority solana.PublicKey,
) (solana.Instruction, error) {
	ix, err := mcm.NewAcceptOwnershipInstruction(seed, config, authority).ValidateAndBuild()
	return &seededInstruction{ix, programID}, err
}

func newSeedlessTransferOwnershipInstruction(
	programID, proposedOwner, config, authority solana.PublicKey,
) (solana.Instruction, error) {
	ix, err := access_controller.NewTransferOwnershipInstruction(proposedOwner, config, authority).ValidateAndBuild()
	return &seedlessInstruction{ix, programID}, err
}

func newSeedlessAcceptOwnershipInstruction(
	programID, config, authority solana.PublicKey,
) (solana.Instruction, error) {
	ix, err := access_controller.NewAcceptOwnershipInstruction(config, authority).ValidateAndBuild()
	return &seedlessInstruction{ix, programID}, err
}

type seedlessInstruction struct {
	*access_controller.Instruction
	programID solana.PublicKey
}

func (s *seedlessInstruction) ProgramID() solana.PublicKey {
	return s.programID
}

type seededInstruction struct {
	*mcm.Instruction
	programID solana.PublicKey
}

func (s *seededInstruction) ProgramID() solana.PublicKey {
	return s.programID
}
