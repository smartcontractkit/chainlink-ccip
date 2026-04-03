package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	mcmsops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/mcms"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	tokenpoolops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/token_pools"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// contractTypeToProgram maps contract types to their compiled program binary names.
// NOTE: The legacy system had a separate "redeploy" path for the offramp (fresh deploy +
// re-wire fee quoter). In practice, the offramp is always upgraded in place like other
// programs, so the redeploy path was intentionally not carried forward.
var contractTypeToProgram = map[cldf.ContractType]string{
	routerops.ContractType:            routerops.ProgramName,
	fqops.ContractType:                fqops.ProgramName,
	offrampops.ContractType:           offrampops.ProgramName,
	rmnremoteops.ContractType:         rmnremoteops.ProgramName,
	common_utils.BurnMintTokenPool:    tokenpoolops.BurnMintProgramName,
	common_utils.LockReleaseTokenPool: tokenpoolops.LockReleaseProgramName,
	utils.McmProgramType:              mcmsops.McmProgramName,
	utils.TimelockProgramType:         mcmsops.TimelockProgramName,
	utils.AccessControllerProgramType: mcmsops.AccessControllerProgramName,
}

func (a *SolanaAdapter) UpgradeChainContracts() *cldf_ops.Sequence[deployapi.ContractUpgradeConfigWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return UpgradeChainContracts
}

var UpgradeChainContracts = cldf_ops.NewSequence(
	"upgrade-chain-contracts",
	semver.MustParse("1.6.0"),
	"Upgrades deployed Solana CCIP programs in place via BPF Loader Upgradeable",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deployapi.ContractUpgradeConfigWithAddress) (sequences.OnChainOutput, error) {
		chain, ok := chains.SolanaChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("solana chain selector %d not found in environment", input.ChainSelector)
		}

		addresses := make([]datastore.AddressRef, 0)
		allBatchOps := make([]mcmsTypes.BatchOperation, 0)

		timelockSignerPDA := utils.GetTimelockSignerPDA(
			input.ExistingAddresses,
			input.ChainSelector,
			common_utils.CLLQualifier,
		)

		for _, contractType := range input.Contracts {
			programName, ok := contractTypeToProgram[contractType]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("no program name mapping for contract type %s", contractType)
			}

			existingAddr := findExistingAddress(input.ExistingAddresses, contractType)
			if existingAddr == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("no existing address found for %s to upgrade", contractType)
			}
			programID := solana.MustPublicKeyFromBase58(existingAddr)

			upgradeAuthority, err := utils.GetUpgradeAuthority(chain.Client, programID)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to read upgrade authority for %s (%s): %w", contractType, programID.String(), err)
			}

			batchOps, err := upgradeProgram(
				b,
				chain,
				programID,
				programName,
				contractType,
				upgradeAuthority,
				timelockSignerPDA,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to upgrade %s: %w", contractType, err)
			}
			allBatchOps = append(allBatchOps, batchOps...)

			addresses = append(addresses, datastore.AddressRef{
				Address:       programID.String(),
				ChainSelector: chain.Selector,
				Type:          datastore.ContractType(contractType),
				Version:       input.Version,
			})
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  allBatchOps,
		}, nil
	},
)

func upgradeProgram(
	b cldf_ops.Bundle,
	chain cldf_solana.Chain,
	programID solana.PublicKey,
	programName string,
	contractType cldf.ContractType,
	upgradeAuthority solana.PublicKey,
	timelockSignerPDA solana.PublicKey,
) ([]mcmsTypes.BatchOperation, error) {
	b.Logger.Infow("Deploying upgrade buffer", "program", contractType, "name", programName)
	bufferAddress, err := utils.DeployToBuffer(chain, b.Logger, programName)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy buffer: %w", err)
	}
	b.Logger.Infow("Buffer deployed", "program", contractType, "buffer", bufferAddress.String())

	// Transfer buffer authority to the on-chain upgrade authority
	setAuthIxn := utils.SetUpgradeAuthority(bufferAddress, chain.DeployerKey.PublicKey(), upgradeAuthority, true)
	if err := chain.Confirm([]solana.Instruction{setAuthIxn}); err != nil {
		return nil, fmt.Errorf("failed to set buffer authority: %w", err)
	}

	// Extend program if the new binary is larger (permissionless — any payer)
	bufferSize, err := utils.GetBufferSize(chain, bufferAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get buffer size: %w", err)
	}
	extendIxn, err := utils.GenerateExtendInstruction(chain, programID, chain.DeployerKey.PublicKey(), bufferSize)
	if err != nil {
		return nil, fmt.Errorf("failed to generate extend instruction: %w", err)
	}
	if extendIxn != nil {
		if err := chain.Confirm([]solana.Instruction{extendIxn}); err != nil {
			return nil, fmt.Errorf("failed to extend program: %w", err)
		}
	}

	upgradeIxn := utils.GenerateUpgradeInstruction(programID, bufferAddress, chain.DeployerKey.PublicKey(), upgradeAuthority)
	closeIxn := utils.GenerateCloseBufferInstruction(bufferAddress, chain.DeployerKey.PublicKey(), upgradeAuthority)

	if upgradeAuthority == chain.DeployerKey.PublicKey() {
		if err := chain.Confirm([]solana.Instruction{upgradeIxn, closeIxn}); err != nil {
			return nil, fmt.Errorf("failed to confirm upgrade: %w", err)
		}
		return nil, nil
	}
	if upgradeAuthority != timelockSignerPDA {
		return nil, fmt.Errorf(
			"unsupported upgrade authority %s: expected deployer %s for direct execution or timelock signer PDA %s for MCMS",
			upgradeAuthority.String(),
			chain.DeployerKey.PublicKey().String(),
			timelockSignerPDA.String(),
		)
	}

	batchOp, err := utils.BuildMCMSBatchOperation(
		chain.Selector,
		[]solana.Instruction{upgradeIxn, closeIxn},
		solana.BPFLoaderUpgradeableProgramID.String(),
		string(contractType),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build MCMS batch: %w", err)
	}
	return []mcmsTypes.BatchOperation{batchOp}, nil
}

func findExistingAddress(refs []datastore.AddressRef, contractType cldf.ContractType) string {
	for _, ref := range refs {
		if ref.Type == datastore.ContractType(contractType) {
			return ref.Address
		}
	}
	return ""
}
