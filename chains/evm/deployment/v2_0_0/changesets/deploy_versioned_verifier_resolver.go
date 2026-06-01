package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// DeployVersionedVerifierResolverChainCfg configures deployment on a single EVM chain.
type DeployVersionedVerifierResolverChainCfg struct {
	ResolverType datastore.ContractType
	// Qualifier is used to differentiate multiple resolver deployments of the same type on the same chain.
	// As a note, since all resolver types share the same bytecode, the CREATE2 salt is solely determined by the qualifier. Therefore, different resolver types deployed on the same chain must use different qualifiers to avoid CREATE2 address collisions.
	// These are enforced at the sequence level.
	Qualifier      string
	CREATE2Factory common.Address
}

// DeployVersionedVerifierResolverCfg configures standalone VersionedVerifierResolver deployments.
type DeployVersionedVerifierResolverCfg struct {
	Chains map[uint64]DeployVersionedVerifierResolverChainCfg
}

// DeployVersionedVerifierResolver deploys VersionedVerifierResolver contracts via CREATE2 on one or more EVM chains.
func DeployVersionedVerifierResolver(mcmsRegistry *cs_changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[cs_changesets.WithMCMS[DeployVersionedVerifierResolverCfg]] {
	return cldf_deployment.CreateChangeSet(
		makeApplyDeployVersionedVerifierResolver(mcmsRegistry),
		makeVerifyDeployVersionedVerifierResolver(),
	)
}

func makeVerifyDeployVersionedVerifierResolver() func(cldf_deployment.Environment, cs_changesets.WithMCMS[DeployVersionedVerifierResolverCfg]) error {
	return func(e cldf_deployment.Environment, cfg cs_changesets.WithMCMS[DeployVersionedVerifierResolverCfg]) error {
		if len(cfg.Cfg.Chains) == 0 {
			return fmt.Errorf("at least one chain must be configured")
		}
		for chainSel, chainCfg := range cfg.Cfg.Chains {
			if chainCfg.Qualifier == "" {
				return fmt.Errorf("chain %d: qualifier must not be empty", chainSel)
			}
			if chainCfg.CREATE2Factory == (common.Address{}) {
				return fmt.Errorf("chain %d: CREATE2Factory is required", chainSel)
			}
			if _, ok := e.BlockChains.EVMChains()[chainSel]; !ok {
				return fmt.Errorf("chain %d not found in environment", chainSel)
			}
		}
		return nil
	}
}

func makeApplyDeployVersionedVerifierResolver(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
) func(cldf_deployment.Environment, cs_changesets.WithMCMS[DeployVersionedVerifierResolverCfg]) (cldf_deployment.ChangesetOutput, error) {
	return func(e cldf_deployment.Environment, cfg cs_changesets.WithMCMS[DeployVersionedVerifierResolverCfg]) (cldf_deployment.ChangesetOutput, error) {
		if err := makeVerifyDeployVersionedVerifierResolver()(e, cfg); err != nil {
			return cldf_deployment.ChangesetOutput{}, err
		}

		newDS := datastore.NewMemoryDataStore()
		var reports []cldf_ops.Report[any, any]
		var batchOps []mcms_types.BatchOperation

		for chainSel, chainCfg := range cfg.Cfg.Chains {
			chain, ok := e.BlockChains.EVMChains()[chainSel]
			if !ok {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d not found in environment", chainSel)
			}

			contractType := chainCfg.ResolverType
			qualifier := chainCfg.Qualifier
			existing := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))

			if ref := findExistingVerifierResolver(existing, contractType, qualifier); ref != nil {
				if err := newDS.Addresses().Add(*ref); err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add existing resolver ref on chain %d: %w", chainSel, err)
				}
				continue
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployVerifierResolverViaCREATE2, chain, sequences.DeployVerifierResolverViaCREATE2Input{
				CREATE2Factory: chainCfg.CREATE2Factory,
				ChainSelector:  chainSel,
				Qualifier:      qualifier,
				Type:           contractType,
				Version:        versioned_verifier_resolver.Version,
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy %s on chain %d: %w", contractType, chainSel, err)
			}
			reports = append(reports, report.ExecutionReports...)

			for _, addr := range report.Output.Addresses {
				if err := newDS.Addresses().Add(addr); err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add %s on chain %d to datastore: %w", contractType, chainSel, err)
				}
			}

			if len(report.Output.Writes) > 0 {
				batchOp, err := contract_utils.NewBatchOperationFromWrites(report.Output.Writes)
				if err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to build batch operation on chain %d: %w", chainSel, err)
				}
				batchOps = append(batchOps, batchOp)
			}
		}

		return cs_changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(newDS).
			WithSingleBatchOpPerChain(batchOps).
			Build(cfg.MCMS)
	}
}

func findExistingVerifierResolver(
	refs []datastore.AddressRef,
	contractType datastore.ContractType,
	qualifier string,
) *datastore.AddressRef {
	for i := range refs {
		ref := refs[i]
		if ref.Type == contractType &&
			ref.Qualifier == qualifier &&
			ref.Version.String() == versioned_verifier_resolver.Version.String() {
			return &ref
		}
	}
	return nil
}
