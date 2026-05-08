package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ManualRegistrationInput struct {
	ChainAdapterVersion *semver.Version       `yaml:"chainAdapterVersion" json:"chainAdapterVersion"`
	Registrations       []RegisterTokenConfig `yaml:"registrations" json:"registrations"`
	MCMS                mcms.Input            `yaml:"mcms,omitempty" json:"mcms"`
}

type ManualRegistrationSequenceInput struct {
	RegisterTokenConfig
	ExistingDataStore datastore.DataStore
}

type RegisterTokenConfig struct {
	// A reference to the token pool. The ChainSelector property should be omitted
	// from the AddressRef, as it is already a property of the RegisterTokenConfig
	// struct. It is NOT necessary to define all fields in the AddressRef. Instead
	// only pass in the minimal set of fields needed to uniquely identify the pool
	// in the datastore. TokenPoolRef is conditionally required based on chain:
	// --
	//  EVM: if no token was found using TokenRef (or it wasn't provided), then we
	//  use the TokenPoolRef to find the token pool, and derive the token address.
	// --
	//  SVM: this field is always required.
	// --
	TokenPoolRef datastore.AddressRef `yaml:"tokenPoolRef" json:"tokenPoolRef"`

	// A reference to the token. The ChainSelector property should be omitted from
	// the AddressRef, as it is already present in the RegisterTokenConfig struct.
	// It is NOT necessary to define every field of AddressRef. Instead, only pass
	// in the *minimal set of fields* needed to uniquely identify the token in the
	// datastore. TokenRef is conditionally required based on chain:
	// --
	//  EVM: if this is not provided or if no token is found using this reference,
	//  then TokenPoolRef will be used as a fallback to derive the token address.
	// --
	//  SVM: this field is always required.
	// --
	TokenRef datastore.AddressRef `yaml:"tokenRef" json:"tokenRef"`

	// The chain selector for the token being registered (required).
	ChainSelector uint64 `yaml:"chainSelector" json:"chainSelector"`

	// The proposed owner of the token (required).
	ProposedOwner string `yaml:"proposedOwner" json:"proposedOwner"`

	// Extra args specific to SVM manual registration. Only required for SVM chains.
	SVMExtraArgs *SVMExtraArgs `yaml:"svmExtraArgs,omitempty" json:"svmExtraArgs,omitempty"`
}

type SVMExtraArgs struct {
	CustomerMintAuthorities []solana.PublicKey `yaml:"customerMintAuthorities,omitempty" json:"customerMintAuthorities,omitempty"`
	SkipTokenPoolInit       bool               `yaml:"skipTokenPoolInit" json:"skipTokenPoolInit"`
}

func ManualRegistration() cldf.ChangeSetV2[ManualRegistrationInput] {
	return cldf.CreateChangeSet(manualRegistrationApply(), manualRegistrationVerify())
}

func manualRegistrationVerify() func(cldf.Environment, ManualRegistrationInput) error {
	return func(e cldf.Environment, cfg ManualRegistrationInput) error {
		// TODO: implement
		return nil
	}
}

func manualRegistrationApply() func(cldf.Environment, ManualRegistrationInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ManualRegistrationInput) (cldf.ChangesetOutput, error) {
		// ds to collect all addresses created during this changeset
		// this gets passed as output
		ds := datastore.NewMemoryDataStore()
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		err := ds.Merge(e.DataStore) // start with existing datastore state from environment
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge existing datastore from environment: %w", err)
		}

		for i, registration := range cfg.Registrations {
			chainfam, err := chain_selectors.GetSelectorFamily(registration.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			// Safeguard: always prevent chain selector mismatches
			if registration.TokenPoolRef.ChainSelector != 0 && registration.TokenPoolRef.ChainSelector != registration.ChainSelector {
				return cldf.ChangesetOutput{}, fmt.Errorf("chain selector mismatch in TokenPoolRef for registration index %d: expected %d, got %d", i, registration.ChainSelector, registration.TokenPoolRef.ChainSelector)
			}
			if registration.TokenRef.ChainSelector != 0 && registration.TokenRef.ChainSelector != registration.ChainSelector {
				return cldf.ChangesetOutput{}, fmt.Errorf("chain selector mismatch in TokenRef for registration index %d: expected %d, got %d", i, registration.ChainSelector, registration.TokenRef.ChainSelector)
			}

			registration.TokenPoolRef, err = TryNormalizeAddressRef(registration.ChainSelector, registration.TokenPoolRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to normalize token pool ref for registration index %d: %w", i, err)
			}
			registration.TokenRef, err = TryNormalizeAddressRef(registration.ChainSelector, registration.TokenRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to normalize token ref for registration index %d: %w", i, err)
			}

			var adapterVersion *semver.Version
			fullPool, findErr := datastore_utils.FindAndFormatRef(ds.Seal(), registration.TokenPoolRef, registration.ChainSelector, datastore_utils.FullRef)
			if findErr == nil {
				adapterVersion = fullPool.Version
			}
			if adapterVersion == nil {
				switch {
				case registration.TokenPoolRef.Version != nil:
					adapterVersion = registration.TokenPoolRef.Version
				case cfg.ChainAdapterVersion != nil:
					adapterVersion = cfg.ChainAdapterVersion
				default:
					return cldf.ChangesetOutput{}, fmt.Errorf("registration[%d]: cannot determine token pool adapter version: pool not found in datastore under tokenPoolRef and neither tokenPoolRef.version nor chainAdapterVersion is set", i)
				}
			}

			adapter, exists := tokenPoolRegistry.GetTokenAdapter(chainfam, adapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s' and version '%v'", chainfam, adapterVersion)
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.ManualRegistration(),
				e.BlockChains,
				ManualRegistrationSequenceInput{
					RegisterTokenConfig: registration,
					ExistingDataStore:   ds.Seal(),
				},
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to execute ManualRegistration sequence for registration index %d: %w", i, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
			for j, addrRef := range report.Output.Addresses {
				if err = ds.Addresses().Upsert(addrRef); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add address ref (%+v) from report output to datastore for registration index %d, address index %d: %w", addrRef, i, j, err)
				}
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
