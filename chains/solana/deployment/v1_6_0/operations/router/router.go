package router

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/mcms/types"

	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	ContractType             cldf_deployment.ContractType = "Router"
	DestChainType            cldf_deployment.ContractType = "RemoteDest"
	TokenPoolLookupTableType cldf_deployment.ContractType = "TokenPoolLookupTable"
	ProgramName                                           = "ccip_router"
	Version                  *semver.Version              = semver.MustParse("1.6.0")
)

type ConnectChainsParams struct {
	Router              solana.PublicKey
	OffRamp             solana.PublicKey
	RemoteChainSelector uint64
	AllowlistEnabled    bool
	AllowedSenders      []solana.PublicKey
}

var Deploy = operations.NewOperation(
	"router:deploy",
	Version,
	"Deploys the Router program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			ContractType,
			Version,
			"",
			ProgramName)
	},
)

var Initialize = operations.NewOperation(
	"router:initialize",
	Version,
	"Initializes the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		programData, err := utils.GetSolProgramData(chain.Client, input.Router)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		authority := GetAuthority(chain, input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		instruction, err := ccip_router.NewInitializeInstruction(
			chain.Selector,
			solana.PublicKey{},
			input.FeeQuoter,
			input.LinkToken,
			input.RMNRemote,
			routerConfigPDA,
			authority,
			solana.SystemProgramID,
			input.Router,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm router initialization: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var ConnectChains = operations.NewOperation(
	"router:connect-chains",
	Version,
	"Connects the Router 1.6.0 contract to other chains",
	func(b operations.Bundle, chain cldf_solana.Chain, input ConnectChainsParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		isUpdate := false
		authority := GetAuthority(chain, input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		routerDestChainPDA, _ := state.FindDestChainStatePDA(input.RemoteChainSelector, input.Router)
		var destChainAccount ccip_router.DestChain
		err := chain.GetAccountDataBorshInto(context.Background(), routerDestChainPDA, &destChainAccount)
		if err == nil {
			b.Logger.Infof("Remote chain state account found: %+v", destChainAccount)
			isUpdate = true
		}
		destChainConfig := ccip_router.DestChainConfig{
			AllowedSenders:   input.AllowedSenders,
			AllowListEnabled: input.AllowlistEnabled,
		}
		var ixn solana.Instruction
		var addressRefs []datastore.AddressRef
		batches := make([]types.BatchOperation, 0)
		if isUpdate {
			ixn, err = ccip_router.NewUpdateDestChainConfigInstruction(
				input.RemoteChainSelector,
				destChainConfig,
				routerDestChainPDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build update dest chain instruction: %w", err)
			}
		} else {
			ixn, err = ccip_router.NewAddChainSelectorInstruction(
				input.RemoteChainSelector,
				destChainConfig,
				routerDestChainPDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build add source chain instruction: %w", err)
			}
			err = utils.ExtendLookupTable(chain, input.OffRamp, []solana.PublicKey{routerDestChainPDA})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to extend OffRamp lookup table: %w", err)
			}
			addressRefs = append(addressRefs, datastore.AddressRef{
				Address:       routerDestChainPDA.String(),
				ChainSelector: chain.Selector,
				Type:          datastore.ContractType(DestChainType),
				Version:       Version,
				Qualifier:     strconv.FormatUint(input.RemoteChainSelector, 10),
			})
		}
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
		} else {
			err = chain.Confirm([]solana.Instruction{ixn})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add dest chain instruction: %w", err)
			}
		}
		return sequences.OnChainOutput{
			BatchOps:  batches,
			Addresses: addressRefs,
		}, nil
	},
)

var AddOffRamp = operations.NewOperation(
	"router:add-off-ramp",
	Version,
	"Adds an OffRamp to the Router 1.6.0 contract for a given chain",
	func(b operations.Bundle, chain cldf_solana.Chain, input ConnectChainsParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		authority := GetAuthority(chain, input.Router)
		var sourceChainAccount ccip_offramp.SourceChain
		offRampSourceChainPDA, _, _ := state.FindOfframpSourceChainPDA(input.RemoteChainSelector, input.OffRamp)
		err := chain.GetAccountDataBorshInto(context.Background(), offRampSourceChainPDA, &sourceChainAccount)
		if err == nil {
			b.Logger.Infof("Remote chain state account found: %+v", sourceChainAccount)
			return sequences.OnChainOutput{}, nil
		}
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		allowedOffRampRemotePDA, _ := state.FindAllowedOfframpPDA(input.RemoteChainSelector, input.OffRamp, input.Router)
		ixn, err := ccip_router.NewAddOfframpInstruction(
			input.RemoteChainSelector,
			input.OffRamp,
			allowedOffRampRemotePDA,
			routerConfigPDA,
			authority,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if authority != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add off ramp instruction: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var TransferOwnership = operations.NewOperation(
	"router:transfer-ownership",
	Version,
	"Transfers ownership of the Router 1.6.0 contract to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Program)
		authority := GetAuthority(chain, input.Program)
		if authority != input.CurrentOwner {
			return sequences.OnChainOutput{}, fmt.Errorf("current owner %s does not match on-chain authority %s", input.CurrentOwner.String(), authority.String())
		}
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := ccip_router.NewTransferOwnershipInstruction(
			input.NewOwner,
			configPDA,
			authority,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if authority != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Program.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm transfer ownership instruction: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var AcceptOwnership = operations.NewOperation(
	"router:accept-ownership",
	Version,
	"Accepts ownership of the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Program)
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := ccip_router.NewAcceptOwnershipInstruction(
			configPDA,
			input.NewOwner,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if input.NewOwner != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Program.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm accept ownership instruction: %w", err)

			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm accept ownership: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var SetPool = operations.NewOperation(
	"router:set-pool",
	Version,
	"Sets the pool of the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input PoolParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		addresses := make([]datastore.AddressRef, 0)
		// Only works for BnM and LnR pools for now
		tokenAdminRegistryPDA, _, _ := state.FindTokenAdminRegistryPDA(input.TokenMint, input.Router)
		if input.TokenPoolLookupTable.IsZero() {
			tokenPoolConfigPDA, _ := tokens.TokenPoolConfigAddress(input.TokenMint, input.TokenPool)
			tokenPoolSigner, _ := tokens.TokenPoolSignerAddress(input.TokenMint, input.TokenPool)
			poolTokenAccount, _, _ := tokens.FindAssociatedTokenAddress(input.TokenProgramID, input.TokenMint, tokenPoolSigner)
			feeTokenConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(input.TokenMint, input.FeeQuoter)
			routerPoolSignerPDA, _, _ := state.FindExternalTokenPoolsSignerPDA(input.TokenPool, input.Router)
			table, err := common.CreateLookupTable(b.GetContext(), chain.Client, *chain.DeployerKey)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create lookup table: %w", err)
			}
			// link the token + token pool lookup table + token mint
			labels := datastore.NewLabelSet(input.TokenPool.String())
			labels.Add(input.TokenPoolType)
			addresses = append(addresses, datastore.AddressRef{
				Address:       table.String(),
				ChainSelector: chain.Selector,
				Labels:        labels,
				Type:          datastore.ContractType(TokenPoolLookupTableType),
				Version:       Version,
				Qualifier:     input.TokenMint.String(),
			})
			list := solana.PublicKeySlice{
				table,                 // 0
				tokenAdminRegistryPDA, // 1
				input.TokenPool,       // 2
				tokenPoolConfigPDA,    // 3 - writable
				poolTokenAccount,      // 4 - writable
				tokenPoolSigner,       // 5
				input.TokenProgramID,  // 6
				input.TokenMint,       // 7 - writable
				feeTokenConfigPDA,     // 8
				routerPoolSignerPDA,   // 9
			}
			if err = common.ExtendLookupTable(b.GetContext(), chain.Client, table, *chain.DeployerKey, list); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to extend lookup table for token pool (mint: %s): %w", input.TokenMint.String(), err)
			}
			if err := common.AwaitSlotChange(b.GetContext(), chain.Client); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to await slot change while extending lookup table: %w", err)
			}
			input.TokenPoolLookupTable = table
		}
		writableIndexes := []uint8{3, 4, 7}
		// we can only sign as either the deployer or the token admin
		// if there is no admin set, we assume the router authority is timelock
		currentAdmin := GetAuthority(chain, input.Router)
		var tokenAdminRegistryAccount ccip_common.TokenAdminRegistry
		if err := chain.GetAccountDataBorshInto(context.Background(), tokenAdminRegistryPDA, &tokenAdminRegistryAccount); err != nil {
			currentAdmin = tokenAdminRegistryAccount.Administrator
		}

		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		base := ccip_router.NewSetPoolInstruction(
			writableIndexes,
			routerConfigPDA,
			tokenAdminRegistryPDA,
			input.TokenMint,
			input.TokenPoolLookupTable,
			currentAdmin,
		)
		base.AccountMetaSlice = append(base.AccountMetaSlice, solana.Meta(input.TokenPoolLookupTable))
		tempIx, err := base.ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build router set pool instruction: %w", err)
		}
		ixData, err := tempIx.Data()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to extract data payload from router set pool instruction: %w", err)
		}
		instruction := solana.NewInstruction(input.Router, tempIx.Accounts(), ixData)
		if currentAdmin != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{instruction},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps:  []types.BatchOperation{batches},
				Addresses: addresses,
			}, nil
		}

		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm set pool: %w", err)
		}
		return sequences.OnChainOutput{
			Addresses: addresses,
		}, nil
	},
)

var RegisterTokenAdminRegistry = operations.NewOperation(
	"router:register-token-admin-registry",
	Version,
	"Registers a Token Admin Registry with the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenAdminRegistryParams) (TokenAdminRegistryOut, error) {
		ccip_router.SetProgramID(input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		tokenAdminRegistryPDA, _, _ := state.FindTokenAdminRegistryPDA(input.TokenMint, input.Router)
		tokenMintAuthority := utils.GetTokenMintAuthority(chain, input.TokenMint)
		// we can only sign as either the deployer or timelock
		// ccip admin should be timelock
		ccipAdmin := GetAuthority(chain, input.Router)
		// if no admin provided, use ccip admin
		if input.Admin.IsZero() {
			input.Admin = ccipAdmin
		}
		var pendingAdmin solana.PublicKey
		var tokenAdminRegistryAccount ccip_common.TokenAdminRegistry
		if err := chain.GetAccountDataBorshInto(context.Background(), tokenAdminRegistryPDA, &tokenAdminRegistryAccount); err != nil {
			if input.Admin == tokenAdminRegistryAccount.Administrator {
				b.Logger.Info("Token admin registry already registered with the given admin:", tokenAdminRegistryAccount)
				return TokenAdminRegistryOut{}, nil
			}
			pendingAdmin = tokenAdminRegistryAccount.PendingAdministrator
		}
		// this is the key that will need to accept the admin registration
		pendingSigner := input.Admin
		var ixn solana.Instruction
		var needProposal bool
		// we need to override the admin if there is a pending admin
		if !pendingAdmin.IsZero() {
			// we need to register via ccip admin if the mint authority is not the deployer
			if tokenMintAuthority != chain.DeployerKey.PublicKey() {
				tmp, err := ccip_router.NewCcipAdminOverridePendingAdministratorInstruction(
					input.Admin, // admin of the tokenAdminRegistry PDA
					routerConfigPDA,
					tokenAdminRegistryPDA, // this gets created
					input.TokenMint,
					ccipAdmin,
					solana.SystemProgramID,
				).ValidateAndBuild()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to build override pending admin instruction: %w", err)
				}
				ixData, err := tmp.Data()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to extract data payload from override pending admin instruction: %w", err)
				}
				ixn = solana.NewInstruction(input.Router, tmp.Accounts(), ixData)
				if ccipAdmin != chain.DeployerKey.PublicKey() {
					needProposal = true
				}
			} else {
				// the token mint authority can register directly
				tmp, err := ccip_router.NewOwnerOverridePendingAdministratorInstruction(
					input.Admin, // admin of the tokenAdminRegistry PDA
					routerConfigPDA,
					tokenAdminRegistryPDA, // this gets created
					input.TokenMint,
					tokenMintAuthority,
					solana.SystemProgramID,
				).ValidateAndBuild()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to build override pending admin instruction: %w", err)
				}
				ixData, err := tmp.Data()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to extract data payload from override pending admin instruction: %w", err)
				}
				ixn = solana.NewInstruction(input.Router, tmp.Accounts(), ixData)
				if tokenMintAuthority != chain.DeployerKey.PublicKey() {
					needProposal = true
				}
			}
		} else {
			// no pending admin, normal registration
			// we need to register via ccip admin if the mint authority is not the deployer
			if tokenMintAuthority != chain.DeployerKey.PublicKey() {
				tmp, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
					input.Admin, // admin of the tokenAdminRegistry PDA
					routerConfigPDA,
					tokenAdminRegistryPDA, // this gets created
					input.TokenMint,
					ccipAdmin,
					solana.SystemProgramID,
				).ValidateAndBuild()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to build propose administrator instruction: %w", err)
				}
				ixData, err := tmp.Data()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to extract data payload from propose administrator instruction: %w", err)
				}
				ixn = solana.NewInstruction(input.Router, tmp.Accounts(), ixData)
				if ccipAdmin != chain.DeployerKey.PublicKey() {
					needProposal = true
				}
			} else {
				// the token mint authority can register directly
				tmp, err := ccip_router.NewOwnerProposeAdministratorInstruction(
					input.Admin, // admin of the tokenAdminRegistry PDA
					routerConfigPDA,
					tokenAdminRegistryPDA, // this gets created
					input.TokenMint,
					tokenMintAuthority,
					solana.SystemProgramID,
				).ValidateAndBuild()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to build propose administrator instruction: %w", err)
				}
				ixData, err := tmp.Data()
				if err != nil {
					return TokenAdminRegistryOut{}, fmt.Errorf("failed to extract data payload from propose administrator instruction: %w", err)
				}
				ixn = solana.NewInstruction(input.Router, tmp.Accounts(), ixData)
				if tokenMintAuthority != chain.DeployerKey.PublicKey() {
					needProposal = true
				}
			}
		}
		if needProposal {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return TokenAdminRegistryOut{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return TokenAdminRegistryOut{
				PendingSigner: pendingSigner,
				OnChainOutput: sequences.OnChainOutput{
					BatchOps: []types.BatchOperation{batches},
				},
			}, nil
		}

		err := chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return TokenAdminRegistryOut{}, fmt.Errorf("failed to confirm register token admin registry: %w", err)
		}
		return TokenAdminRegistryOut{
			PendingSigner: pendingSigner,
		}, nil
	},
)

var AcceptTokenAdminRegistry = operations.NewOperation(
	"router:accept-token-admin-registry",
	Version,
	"Accepts a Token Admin Registry with the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenAdminRegistryParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		tokenAdminRegistryPDA, _, _ := state.FindTokenAdminRegistryPDA(input.TokenMint, input.Router)
		var pendingAdmin solana.PublicKey
		var tokenAdminRegistryAccount ccip_common.TokenAdminRegistry
		if err := chain.GetAccountDataBorshInto(context.Background(), tokenAdminRegistryPDA, &tokenAdminRegistryAccount); err == nil {
			// NOTE: only set pendingAdmin if the account exists / fetch succeeds
			pendingAdmin = tokenAdminRegistryAccount.PendingAdministrator
		}
		timelockSigner := utils.GetTimelockSignerPDA(
			input.ExistingAddresses,
			chain.Selector,
			common_utils.CLLQualifier,
		)
		if pendingAdmin.IsZero() {
			// if there is no pending admin, we assume the authority is timelock
			// but we need to confirm that timelock is indeed the authority
			if input.Admin != timelockSigner {
				return sequences.OnChainOutput{}, fmt.Errorf("no pending admin found for token admin registry, expected timelock signer %s but got %s", timelockSigner.String(), input.Admin.String())
			}
		} else if pendingAdmin != timelockSigner && pendingAdmin != chain.DeployerKey.PublicKey() {
			// we can only sign as either the deployer or timelock
			return sequences.OnChainOutput{}, fmt.Errorf("pending admin %s does not match timelock signer %s or deployer %s", pendingAdmin.String(), timelockSigner.String(), chain.DeployerKey.PublicKey().String())
		}
		// sign as the pending admin to accept
		// when there is no pending admin, we assume the authority is timelock
		tempIx, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
			routerConfigPDA,
			tokenAdminRegistryPDA,
			input.TokenMint,
			pendingAdmin,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to generate instructions: %w", err)
		}
		ixData, err := tempIx.Data()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to extract data payload router accept admin role token admin registry instruction: %w", err)
		}
		ixn := solana.NewInstruction(input.Router, tempIx.Accounts(), ixData)
		// pending admin unset = we're proposing and accepting in the same batch, so timelock must be the signer
		// pending admin set = we need a proposal if the admin is not the deployer
		if pendingAdmin.IsZero() || pendingAdmin != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []types.BatchOperation{batches},
			}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm register token admin registry: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var TransferTokenAdminRegistry = operations.NewOperation(
	"router:transfer-token-admin-registry",
	Version,
	"Transfers a Token Admin Registry with the Router 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input TokenAdminRegistryParams) (sequences.OnChainOutput, error) {
		ccip_router.SetProgramID(input.Router)
		routerConfigPDA, _, _ := state.FindConfigPDA(input.Router)
		tokenAdminRegistryPDA, _, _ := state.FindTokenAdminRegistryPDA(input.TokenMint, input.Router)
		var currentAdmin solana.PublicKey
		var tokenAdminRegistryAccount ccip_common.TokenAdminRegistry
		if err := chain.GetAccountDataBorshInto(context.Background(), tokenAdminRegistryPDA, &tokenAdminRegistryAccount); err != nil {
			currentAdmin = tokenAdminRegistryAccount.Administrator
		}
		// we can only sign as either the deployer or timelock
		// ccip admin should be timelock
		ccipAdmin := GetAuthority(chain, input.Router)
		// if no admin provided, use ccip admin
		if input.Admin.IsZero() {
			input.Admin = ccipAdmin
		}
		// sign as the current admin to transfer
		tempIx, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
			input.Admin,
			routerConfigPDA,
			tokenAdminRegistryPDA,
			input.TokenMint,
			currentAdmin,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to generate instructions: %w", err)
		}
		ixData, err := tempIx.Data()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to extract data payload router accept admin role token admin registry instruction: %w", err)
		}
		ixn := solana.NewInstruction(input.Router, tempIx.Accounts(), ixData)
		// now we need a proposal if the admin is not the deployer
		if currentAdmin != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Router.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []types.BatchOperation{batches},
			}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm register token admin registry: %w", err)
		}

		return sequences.OnChainOutput{}, nil
	},
)

func GetAuthority(chain cldf_solana.Chain, program solana.PublicKey) solana.PublicKey {
	programData := ccip_router.Config{}
	routerConfigPDA, _, _ := state.FindConfigPDA(program)
	err := chain.GetAccountDataBorshInto(context.Background(), routerConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Owner
}

type Params struct {
	FeeQuoter solana.PublicKey
	Router    solana.PublicKey
	LinkToken solana.PublicKey
	RMNRemote solana.PublicKey
}

type PoolParams struct {
	Router               solana.PublicKey
	FeeQuoter            solana.PublicKey
	TokenMint            solana.PublicKey
	TokenProgramID       solana.PublicKey
	TokenPool            solana.PublicKey
	TokenPoolLookupTable solana.PublicKey
	TokenPoolType        string
}

type TokenAdminRegistryParams struct {
	Router            solana.PublicKey
	TokenMint         solana.PublicKey
	Admin             solana.PublicKey
	ExistingAddresses []datastore.AddressRef
}

type TokenAdminRegistryOut struct {
	sequences.OnChainOutput
	PendingSigner solana.PublicKey
}
